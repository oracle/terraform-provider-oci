// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const (
	ociConfigFileEnv        = "OCI_CONFIG_FILE"
	ociSecondaryConfigDir   = ".oraclebmc"
	terraformEnvPrefix      = "TF_VAR"
	defaultOCIConfigProfile = "DEFAULT"
)

type fileCredentialSnapshot struct {
	tenancy        string
	user           string
	region         string
	fingerprint    string
	privateKey     []byte
	privateKeyErr  error
	privateKeyPass string
	securityToken  string
	sessionToken   bool
}

func (p *fileCredentialSnapshot) TenancyOCID() (string, error) {
	if p.tenancy == "" {
		return "", fmt.Errorf("tenancy OCID cannot be empty")
	}
	return p.tenancy, nil
}

func (p *fileCredentialSnapshot) UserOCID() (string, error) {
	// The SDK's ordinary file provider prefers a configured user even when
	// security_token_file is also present. The dedicated session-token provider
	// always omits the user. Preserve that distinction for CLI parity.
	if p.sessionToken || (p.user == "" && p.securityToken != "") {
		return "", nil
	}
	if p.user == "" {
		return "", fmt.Errorf("user OCID cannot be empty")
	}
	return p.user, nil
}

func (p *fileCredentialSnapshot) KeyFingerprint() (string, error) {
	if p.fingerprint == "" {
		return "", fmt.Errorf("fingerprint cannot be empty")
	}
	return p.fingerprint, nil
}

func (p *fileCredentialSnapshot) Region() (string, error) {
	if p.region == "" {
		return "", fmt.Errorf("region cannot be empty")
	}
	return p.region, nil
}

func (p *fileCredentialSnapshot) PrivateRSAKey() (*rsa.PrivateKey, error) {
	if p.privateKeyErr != nil {
		return nil, p.privateKeyErr
	}
	return oci_common.PrivateKeyFromBytesWithPassword(p.privateKey, []byte(p.privateKeyPass))
}

func (p *fileCredentialSnapshot) KeyID() (string, error) {
	// Match the OCI SDK providers: an explicit session-token provider always
	// uses the token, while an ordinary file provider uses it only when the
	// profile does not define a user.
	if p.sessionToken || (p.user == "" && p.securityToken != "") {
		token, err := os.ReadFile(p.securityToken)
		if err != nil {
			return "", fmt.Errorf("cannot read security token file %q: %w", p.securityToken, err)
		}
		return "ST$" + string(token), nil
	}
	tenancy, err := p.TenancyOCID()
	if err != nil {
		return "", err
	}
	user, err := p.UserOCID()
	if err != nil {
		return "", err
	}
	fingerprint, err := p.KeyFingerprint()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", tenancy, user, fingerprint), nil
}

func (p *fileCredentialSnapshot) AuthType() (oci_common.AuthConfig, error) {
	return oci_common.AuthConfig{AuthType: oci_common.UserPrincipal, IsFromConfigFile: true}, nil
}

func (p *fileCredentialSnapshot) Refreshable() bool {
	return p.sessionToken
}

func inProcessFileConfigurationProviders(profile string) ([]oci_common.ConfigurationProvider, error) {
	providers := make([]oci_common.ConfigurationProvider, 0, 3)
	if profile != "" {
		configPath := defaultOCIConfigPath()
		selected, err := loadFileCredentialSnapshot(configPath, profile, "")
		if err != nil {
			return nil, err
		}
		providers = append(providers, selected)
		if profile != defaultOCIConfigProfile {
			// OCI SDK CustomProfileConfigProvider composes the selected profile,
			// DEFAULT, then TF_VAR values. Keep the same fallback order here.
			if fallback, err := loadFileCredentialSnapshot(configPath, defaultOCIConfigProfile, ""); err == nil {
				providers = append(providers, fallback)
			}
		}
	} else {
		for _, path := range defaultOCIConfigPaths() {
			provider, err := loadFileCredentialSnapshot(path, defaultOCIConfigProfile, "")
			if err == nil {
				providers = append(providers, provider)
			}
		}
	}
	providers = append(providers, oci_common.ConfigurationProviderEnvironmentVariables(terraformEnvPrefix, ""))
	return providers, nil
}

func usesInProcessFileConfiguration(auth string) bool {
	return strings.EqualFold(auth, globalvar.AuthAPIKeySetting)
}

func loadFileCredentialSnapshot(configPath, profile, password string) (*fileCredentialSnapshot, error) {
	return loadFileCredentialSnapshotForMode(configPath, profile, password, false)
}

func loadSessionTokenCredentialSnapshot(configPath, profile, password string) (*fileCredentialSnapshot, error) {
	return loadFileCredentialSnapshotForMode(configPath, profile, password, true)
}

func loadFileCredentialSnapshotForMode(configPath, profile, password string, sessionToken bool) (*fileCredentialSnapshot, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read OCI config file %q: %w", configPath, err)
	}
	values, err := parseOCIConfigProfile(data, profile)
	if err != nil {
		return nil, fmt.Errorf("cannot read OCI config profile %q from %q: %w", profile, configPath, err)
	}
	keyPath := expandOCIPath(values["key_file"])
	var privateKey []byte
	var privateKeyErr error
	if keyPath == "" {
		privateKeyErr = fmt.Errorf("OCI config profile %q in %q does not define key_file", profile, configPath)
	} else {
		privateKey, err = os.ReadFile(keyPath)
		if err != nil {
			privateKeyErr = fmt.Errorf("cannot read OCI private key file %q: %w", keyPath, err)
		}
	}
	if password == "" {
		password = firstNonEmpty(values["passphrase"], values["pass_phrase"])
	}
	return &fileCredentialSnapshot{
		tenancy:        values["tenancy"],
		user:           values["user"],
		region:         values["region"],
		fingerprint:    values["fingerprint"],
		privateKey:     privateKey,
		privateKeyErr:  privateKeyErr,
		privateKeyPass: password,
		securityToken:  expandOCIPath(values["security_token_file"]),
		sessionToken:   sessionToken,
	}, nil
}

func parseOCIConfigProfile(data []byte, profile string) (map[string]string, error) {
	values := map[string]string{}
	inProfile := false
	found := false
	for line := range strings.SplitSeq(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			current := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "["), "]"))
			inProfile = current == profile
			found = found || inProfile
			continue
		}
		if !inProfile {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		values[strings.ToLower(strings.TrimSpace(key))] = strings.TrimSpace(value)
	}
	if !found {
		return nil, fmt.Errorf("profile not found")
	}
	return values, nil
}

func defaultOCIConfigPath() string {
	primary := filepath.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
	if _, err := os.Stat(primary); err == nil {
		return primary
	}
	if fallback := os.Getenv(ociConfigFileEnv); fallback != "" {
		if _, err := os.Stat(fallback); err == nil {
			return fallback
		}
	}
	return primary
}

func defaultOCIConfigPaths() []string {
	paths := []string{
		defaultOCIConfigPath(),
		filepath.Join(utils.GetHomeFolder(), ociSecondaryConfigDir, globalvar.DefaultConfigFileName),
	}
	seen := map[string]struct{}{}
	result := make([]string, 0, len(paths))
	for _, path := range paths {
		if _, ok := seen[path]; ok {
			continue
		}
		seen[path] = struct{}{}
		result = append(result, path)
	}
	return result
}

func expandOCIPath(path string) string {
	if path == "" {
		return ""
	}
	return utils.ExpandPath(path)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

// effectiveInProcessConfigString resolves a provider value using the same
// precedence as the provider schemas: explicit configuration, TF_VAR_*,
// OCI_*, then the schema default. The fingerprint is computed before the
// provider schema applies its defaults, so it must resolve these values itself
// to describe the credentials the configured provider will actually use.
func effectiveInProcessConfigString(config map[string]any, attrName, defaultValue string) string {
	if value, _ := config[attrName].(string); value != "" {
		return value
	}
	return MultiEnvDefaultFunc([]string{tfVarName(attrName), ociVarName(attrName)}, defaultValue)
}

// InProcessFileCredentialFingerprint returns a stable digest of file contents
// that can affect an in-process provider configuration. It intentionally omits
// security-token contents because those are refreshed for every signed request.
func InProcessFileCredentialFingerprint(config map[string]any) (string, error) {
	files := map[string][]byte{}
	auth := effectiveInProcessConfigString(config, globalvar.AuthAttrName, globalvar.AuthAPIKeySetting)
	isAPIKey := auth == "" || strings.EqualFold(auth, globalvar.AuthAPIKeySetting)
	isSecurityToken := strings.EqualFold(auth, globalvar.AuthSecurityToken)
	if !isAPIKey && !isSecurityToken {
		return hashCredentialFiles(files), nil
	}

	if inlineKey := effectiveInProcessConfigString(config, globalvar.PrivateKeyAttrName, ""); isAPIKey && inlineKey == "" {
		if keyPath := effectiveInProcessConfigString(config, globalvar.PrivateKeyPathAttrName, ""); keyPath != "" {
			if err := addFingerprintFile(files, expandOCIPath(keyPath), true); err != nil {
				return "", err
			}
		}
	}

	profile := effectiveInProcessConfigString(config, globalvar.ConfigFileProfileAttrName, "")
	if profile == "" && !configurationUsesDefaultFile(config) {
		return hashCredentialFiles(files), nil
	}
	paths := defaultOCIConfigPaths()
	if profile != "" {
		paths = []string{defaultOCIConfigPath()}
	}
	for _, configPath := range paths {
		data, err := os.ReadFile(configPath)
		if err != nil {
			if profile != "" {
				return "", fmt.Errorf("cannot read OCI config file %q: %w", configPath, err)
			}
			continue
		}
		files[configPath] = data
		profiles := []string{defaultOCIConfigProfile}
		if profile != "" && profile != defaultOCIConfigProfile {
			profiles = append([]string{profile}, profiles...)
		}
		for _, name := range profiles {
			values, err := parseOCIConfigProfile(data, name)
			if err != nil {
				if name == profile && profile != "" {
					return "", fmt.Errorf("cannot read OCI config profile %q from %q: %w", name, configPath, err)
				}
				continue
			}
			if keyPath := expandOCIPath(values["key_file"]); keyPath != "" {
				if err := addFingerprintFile(files, keyPath, false); err != nil {
					return "", err
				}
			}
		}
	}

	return hashCredentialFiles(files), nil
}

func configurationUsesDefaultFile(config map[string]any) bool {
	auth := effectiveInProcessConfigString(config, globalvar.AuthAttrName, globalvar.AuthAPIKeySetting)
	if auth != "" && !strings.EqualFold(auth, globalvar.AuthAPIKeySetting) {
		return strings.EqualFold(auth, globalvar.AuthSecurityToken)
	}
	for _, key := range []string{
		globalvar.TenancyOcidAttrName,
		globalvar.UserOcidAttrName,
		globalvar.FingerprintAttrName,
		globalvar.RegionAttrName,
	} {
		if effectiveInProcessConfigString(config, key, "") == "" {
			return true
		}
	}
	privateKey := effectiveInProcessConfigString(config, globalvar.PrivateKeyAttrName, "")
	privateKeyPath := effectiveInProcessConfigString(config, globalvar.PrivateKeyPathAttrName, "")
	return privateKey == "" && privateKeyPath == ""
}

func hashCredentialFiles(files map[string][]byte) string {
	paths := make([]string, 0, len(files))
	for path := range files {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	hash := sha256.New()
	for _, path := range paths {
		hash.Write([]byte(path))
		hash.Write([]byte{0})
		hash.Write(files[path])
		hash.Write([]byte{0})
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func addFingerprintFile(files map[string][]byte, path string, required bool) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if !required {
			// Preserve the path in the digest so a later transition to a
			// readable file invalidates the host's configured metadata. A
			// lower-precedence profile key may legitimately be unavailable
			// when another configuration provider supplies the private key.
			files[path] = nil
			return nil
		}
		return fmt.Errorf("cannot read OCI credential file %q: %w", path, err)
	}
	files[path] = data
	return nil
}
