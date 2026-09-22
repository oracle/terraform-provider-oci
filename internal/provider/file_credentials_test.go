// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

func TestUnitInProcessPrivateKeyPathFingerprintChangesWithContents(t *testing.T) {
	keyPath := filepath.Join(t.TempDir(), "oci_api_key.pem")
	writeTestPrivateKey(t, keyPath, generateTestPrivateKey(t))
	config := completeAPIKeyConfiguration()
	config[globalvar.PrivateKeyPathAttrName] = keyPath

	first, err := InProcessFileCredentialFingerprint(config)
	if err != nil {
		t.Fatalf("InProcessFileCredentialFingerprint() unexpected error: %v", err)
	}
	writeTestPrivateKey(t, keyPath, generateTestPrivateKey(t))
	second, err := InProcessFileCredentialFingerprint(config)
	if err != nil {
		t.Fatalf("InProcessFileCredentialFingerprint() after rotation unexpected error: %v", err)
	}
	if first == second {
		t.Fatal("credential fingerprint did not change after replacing private key contents at the same path")
	}
}

func TestUnitLoadFileCredentialSnapshotReloadsConfigAndKey(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "config")
	keyPath := filepath.Join(dir, "oci_api_key.pem")
	firstKey := generateTestPrivateKey(t)
	writeTestPrivateKey(t, keyPath, firstKey)
	writeTestOCIConfig(t, configPath, keyPath, "fingerprint-a", "")

	first, err := loadFileCredentialSnapshot(configPath, defaultOCIConfigProfile, "")
	if err != nil {
		t.Fatalf("loadFileCredentialSnapshot() unexpected error: %v", err)
	}
	firstParsed, err := first.PrivateRSAKey()
	if err != nil {
		t.Fatalf("first snapshot PrivateRSAKey() unexpected error: %v", err)
	}
	if firstParsed.PublicKey.N.Cmp(firstKey.PublicKey.N) != 0 {
		t.Fatal("first snapshot did not load the expected private key")
	}

	secondKey := generateTestPrivateKey(t)
	writeTestPrivateKey(t, keyPath, secondKey)
	writeTestOCIConfig(t, configPath, keyPath, "fingerprint-b", "")
	second, err := loadFileCredentialSnapshot(configPath, defaultOCIConfigProfile, "")
	if err != nil {
		t.Fatalf("loadFileCredentialSnapshot() after rotation unexpected error: %v", err)
	}
	secondParsed, err := second.PrivateRSAKey()
	if err != nil {
		t.Fatalf("second snapshot PrivateRSAKey() unexpected error: %v", err)
	}
	if secondParsed.PublicKey.N.Cmp(secondKey.PublicKey.N) != 0 {
		t.Fatal("second snapshot retained stale private key contents")
	}
	if second.fingerprint != "fingerprint-b" {
		t.Fatalf("second snapshot fingerprint = %q, want fingerprint-b", second.fingerprint)
	}
	if first.fingerprint == second.fingerprint {
		t.Fatal("second snapshot retained stale config profile contents")
	}
}

func TestUnitInProcessProviderReloadsConfigAndKey(t *testing.T) {
	home := t.TempDir()
	t.Setenv("TF_HOME_OVERRIDE", home)
	configPath := filepath.Join(home, globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
	keyPath := filepath.Join(home, "oci_api_key.pem")
	if err := os.MkdirAll(filepath.Dir(configPath), 0o700); err != nil {
		t.Fatalf("cannot create OCI config directory: %v", err)
	}

	firstKey := generateTestPrivateKey(t)
	writeTestPrivateKey(t, keyPath, firstKey)
	writeTestOCIConfig(t, configPath, keyPath, "fingerprint-a", "")
	first := newInProcessSDKConfigProvider(t, defaultOCIConfigProfile)
	firstParsed, err := first.PrivateRSAKey()
	if err != nil {
		t.Fatalf("first provider PrivateRSAKey() unexpected error: %v", err)
	}

	secondKey := generateTestPrivateKey(t)
	writeTestPrivateKey(t, keyPath, secondKey)
	writeTestOCIConfig(t, configPath, keyPath, "fingerprint-b", "")
	second := newInProcessSDKConfigProvider(t, defaultOCIConfigProfile)
	secondParsed, err := second.PrivateRSAKey()
	if err != nil {
		t.Fatalf("second provider PrivateRSAKey() unexpected error: %v", err)
	}

	if firstParsed.PublicKey.N.Cmp(firstKey.PublicKey.N) != 0 {
		t.Fatal("first in-process provider did not load the original private key")
	}
	if secondParsed.PublicKey.N.Cmp(secondKey.PublicKey.N) != 0 {
		t.Fatal("second in-process provider retained stale private key contents")
	}
	if firstParsed.PublicKey.N.Cmp(secondParsed.PublicKey.N) == 0 {
		t.Fatal("reconfigured in-process provider reused the original private key")
	}
}

func TestUnitConfigProfileFingerprintChangesWithConfigOrKeyContents(t *testing.T) {
	home := t.TempDir()
	t.Setenv("TF_HOME_OVERRIDE", home)
	configPath := filepath.Join(home, globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
	keyPath := filepath.Join(home, "oci_api_key.pem")
	if err := os.MkdirAll(filepath.Dir(configPath), 0o700); err != nil {
		t.Fatalf("cannot create OCI config directory: %v", err)
	}
	writeTestPrivateKey(t, keyPath, generateTestPrivateKey(t))
	writeTestOCIConfig(t, configPath, keyPath, "fingerprint-a", "")
	config := map[string]any{globalvar.ConfigFileProfileAttrName: defaultOCIConfigProfile}

	first, err := InProcessFileCredentialFingerprint(config)
	if err != nil {
		t.Fatalf("first fingerprint unexpected error: %v", err)
	}
	writeTestOCIConfig(t, configPath, keyPath, "fingerprint-b", "")
	second, err := InProcessFileCredentialFingerprint(config)
	if err != nil {
		t.Fatalf("fingerprint after config rotation unexpected error: %v", err)
	}
	if first == second {
		t.Fatal("credential fingerprint did not change after replacing config contents")
	}

	writeTestPrivateKey(t, keyPath, generateTestPrivateKey(t))
	third, err := InProcessFileCredentialFingerprint(config)
	if err != nil {
		t.Fatalf("fingerprint after key rotation unexpected error: %v", err)
	}
	if second == third {
		t.Fatal("credential fingerprint did not change after replacing profile key contents")
	}
}

func TestUnitFileCredentialSnapshotReloadsSecurityToken(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "config")
	keyPath := filepath.Join(dir, "oci_api_key.pem")
	tokenPath := filepath.Join(dir, "security_token")
	writeTestPrivateKey(t, keyPath, generateTestPrivateKey(t))
	writeTestOCIConfig(t, configPath, keyPath, "fingerprint", tokenPath)
	if err := os.WriteFile(tokenPath, []byte("token-a"), 0o600); err != nil {
		t.Fatalf("cannot write security token: %v", err)
	}

	snapshot, err := loadFileCredentialSnapshot(configPath, defaultOCIConfigProfile, "")
	if err != nil {
		t.Fatalf("loadFileCredentialSnapshot() unexpected error: %v", err)
	}
	first, err := snapshot.KeyID()
	if err != nil {
		t.Fatalf("KeyID() unexpected error: %v", err)
	}
	if first != "ST$token-a" {
		t.Fatalf("KeyID() = %q, want ST$token-a", first)
	}
	if err := os.WriteFile(tokenPath, []byte("token-b"), 0o600); err != nil {
		t.Fatalf("cannot rotate security token: %v", err)
	}
	second, err := snapshot.KeyID()
	if err != nil {
		t.Fatalf("KeyID() after rotation unexpected error: %v", err)
	}
	if second != "ST$token-b" {
		t.Fatalf("KeyID() after rotation = %q, want ST$token-b", second)
	}
}

func TestUnitSecurityTokenContentsDoNotChangeCredentialFingerprint(t *testing.T) {
	home := t.TempDir()
	t.Setenv("TF_HOME_OVERRIDE", home)
	configPath := filepath.Join(home, globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
	keyPath := filepath.Join(home, "oci_api_key.pem")
	tokenPath := filepath.Join(home, "security_token")
	if err := os.MkdirAll(filepath.Dir(configPath), 0o700); err != nil {
		t.Fatalf("cannot create OCI config directory: %v", err)
	}
	writeTestPrivateKey(t, keyPath, generateTestPrivateKey(t))
	writeTestOCIConfig(t, configPath, keyPath, "fingerprint", tokenPath)
	if err := os.WriteFile(tokenPath, []byte("token-a"), 0o600); err != nil {
		t.Fatalf("cannot write security token: %v", err)
	}
	config := map[string]any{
		globalvar.AuthAttrName:              globalvar.AuthSecurityToken,
		globalvar.ConfigFileProfileAttrName: defaultOCIConfigProfile,
	}
	first, err := InProcessFileCredentialFingerprint(config)
	if err != nil {
		t.Fatalf("first fingerprint unexpected error: %v", err)
	}
	if err := os.WriteFile(tokenPath, []byte("token-b"), 0o600); err != nil {
		t.Fatalf("cannot rotate security token: %v", err)
	}
	second, err := InProcessFileCredentialFingerprint(config)
	if err != nil {
		t.Fatalf("second fingerprint unexpected error: %v", err)
	}
	if first != second {
		t.Fatal("security-token contents changed the metadata fingerprint; tokens should reload per request")
	}
}

func TestUnitParseOCIConfigProfile(t *testing.T) {
	data := []byte("[DEFAULT]\nregion=us-ashburn-1\n\n[OTHER]\nregion = us-phoenix-1\nkey_file=/tmp/key=with-equals.pem\n")
	values, err := parseOCIConfigProfile(data, "OTHER")
	if err != nil {
		t.Fatalf("parseOCIConfigProfile() unexpected error: %v", err)
	}
	if values["region"] != "us-phoenix-1" || values["key_file"] != "/tmp/key=with-equals.pem" {
		t.Fatalf("parseOCIConfigProfile() = %#v", values)
	}
	if _, err := parseOCIConfigProfile(data, "MISSING"); err == nil {
		t.Fatal("parseOCIConfigProfile() expected missing-profile error")
	}
}

func completeAPIKeyConfiguration() map[string]any {
	return map[string]any{
		globalvar.AuthAttrName:        globalvar.AuthAPIKeySetting,
		globalvar.TenancyOcidAttrName: "ocid1.tenancy.oc1..test",
		globalvar.UserOcidAttrName:    "ocid1.user.oc1..test",
		globalvar.FingerprintAttrName: "fingerprint",
		globalvar.RegionAttrName:      "us-ashburn-1",
	}
}

func newInProcessSDKConfigProvider(t *testing.T, profile string) interface {
	PrivateRSAKey() (*rsa.PrivateKey, error)
} {
	t.Helper()
	r := &schema.Resource{Schema: SchemaMap()}
	d := r.Data(nil)
	if err := d.Set(globalvar.AuthAttrName, globalvar.AuthAPIKeySetting); err != nil {
		t.Fatalf("cannot set auth: %v", err)
	}
	if err := d.Set(globalvar.ConfigFileProfileAttrName, profile); err != nil {
		t.Fatalf("cannot set config profile: %v", err)
	}
	clients := &tf_client.OracleClients{
		SdkClientMap:  make(map[string]interface{}),
		Configuration: make(map[string]string),
	}
	provider, err := getSdkConfigProvider(d, clients, true)
	if err != nil {
		t.Fatalf("getSdkConfigProvider() unexpected error: %v", err)
	}
	return provider
}

func generateTestPrivateKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("cannot generate test private key: %v", err)
	}
	return key
}

func writeTestPrivateKey(t *testing.T, path string, key *rsa.PrivateKey) {
	t.Helper()
	data := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("cannot write test private key: %v", err)
	}
}

func writeTestOCIConfig(t *testing.T, configPath, keyPath, fingerprint, tokenPath string) {
	t.Helper()
	content := fmt.Sprintf("[DEFAULT]\ntenancy=ocid1.tenancy.oc1..test\nuser=ocid1.user.oc1..test\nfingerprint=%s\nregion=us-ashburn-1\nkey_file=%s\n", fingerprint, keyPath)
	if tokenPath != "" {
		content += "security_token_file=" + tokenPath + "\n"
	}
	if err := os.WriteFile(configPath, []byte(content), 0o600); err != nil {
		t.Fatalf("cannot write OCI config: %v", err)
	}
}
