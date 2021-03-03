// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	sdkMeta "github.com/hashicorp/terraform-plugin-sdk/meta"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"runtime"

	oci_common "github.com/oracle/oci-go-sdk/v36/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/v36/common/auth"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var descriptions map[string]string
var apiKeyConfigAttributes = [...]string{userOcidAttrName, fingerprintAttrName, privateKeyAttrName, privateKeyPathAttrName, privateKeyPasswordAttrName}
var ociProvider *schema.Provider

var terraformCLIVersion = unknownTerraformCLIVersion
var avoidWaitingForDeleteTarget bool

type ConfigureClient func(client *oci_common.BaseClient) error

var configureClient ConfigureClient // global fn ref used to configure all clients initially and others later on

var OciResources map[string]*schema.Resource
var OciDatasources map[string]*schema.Resource

const (
	authAPIKeySetting                     = "ApiKey"
	authInstancePrincipalSetting          = "InstancePrincipal"
	authInstancePrincipalWithCertsSetting = "InstancePrincipalWithCerts"
	authSecurityToken                     = "SecurityToken"
	requestHeaderOpcOboToken              = "opc-obo-token"
	requestHeaderOpcHostSerial            = "opc-host-serial"
	defaultRequestTimeout                 = 0
	defaultConnectionTimeout              = 10 * time.Second
	defaultTLSHandshakeTimeout            = 10 * time.Second
	defaultUserAgentProviderName          = "Oracle-TerraformProvider"
	unknownTerraformCLIVersion            = "unknown"
	testTerraformCLIVersion               = "test"
	userAgentFormatter                    = "Oracle-GoSDK/%s (go/%s; %s/%s; terraform/%s; terraform-cli/%s) %s/%s"
	userAgentProviderNameEnv              = "USER_AGENT_PROVIDER_NAME"
	domainNameOverrideEnv                 = "domain_name_override"
	clientHostOverridesEnv                = "CLIENT_HOST_OVERRIDES"
	customCertLocationEnv                 = "custom_cert_location"
	acceptLocalCerts                      = "accept_local_certs"

	authAttrName                 = "auth"
	tenancyOcidAttrName          = "tenancy_ocid"
	userOcidAttrName             = "user_ocid"
	fingerprintAttrName          = "fingerprint"
	privateKeyAttrName           = "private_key"
	privateKeyPathAttrName       = "private_key_path"
	privateKeyPasswordAttrName   = "private_key_password"
	regionAttrName               = "region"
	disableAutoRetriesAttrName   = "disable_auto_retries"
	retryDurationSecondsAttrName = "retry_duration_seconds"
	oboTokenAttrName             = "obo_token"
	oboTokenPath                 = "obo_token_path"
	configFileProfileAttrName    = "config_file_profile"

	tfEnvPrefix              = "TF_VAR_"
	ociEnvPrefix             = "OCI_"
	defaultConfigFileName    = "config"
	defaultConfigDirName     = ".oci"
	colonDelimiter           = ";"
	equalToOperatorDelimiter = "="
)

// OboTokenProvider interface that wraps information about auth tokens so the sdk client can make calls
// on behalf of a different authorized user
type OboTokenProvider interface {
	OboToken() (string, error)
}

//EmptyOboTokenProvider always provides an empty obo token
type emptyOboTokenProvider struct{}

//OboToken provides the obo token
func (provider emptyOboTokenProvider) OboToken() (string, error) {
	return "", nil
}

type oboTokenProviderFromEnv struct{}

func (p oboTokenProviderFromEnv) OboToken() (string, error) {
	// priority token from path than token from environment
	if path := getEnvSettingWithBlankDefault(oboTokenPath); path != "" {
		token, err := getTokenFromFile(path)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return getEnvSettingWithBlankDefault(oboTokenAttrName), nil
}

func tfVarName(attrName string) string {
	return tfEnvPrefix + attrName
}

func ociVarName(attrName string) string {
	return ociEnvPrefix + strings.ToUpper(attrName)
}

func init() {
	descriptions = map[string]string{
		authAttrName:        fmt.Sprintf("(Optional) The type of auth to use. Options are '%s', '%s' and '%s'. By default, '%s' will be used.", authAPIKeySetting, authSecurityToken, authInstancePrincipalSetting, authAPIKeySetting),
		tenancyOcidAttrName: fmt.Sprintf("(Optional) The tenancy OCID for a user. The tenancy OCID can be found at the bottom of user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		userOcidAttrName:    fmt.Sprintf("(Optional) The user OCID. This can be found in user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		fingerprintAttrName: fmt.Sprintf("(Optional) The fingerprint for the user's RSA key. This can be found in user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		regionAttrName:      "(Required) The region for API connections (e.g. us-ashburn-1).",
		privateKeyAttrName: "(Optional) A PEM formatted RSA private key for the user.\n" +
			fmt.Sprintf("A private_key or a private_key_path must be provided if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		privateKeyPathAttrName: "(Optional) The path to the user's PEM formatted private key.\n" +
			fmt.Sprintf("A private_key or a private_key_path must be provided if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		privateKeyPasswordAttrName: "(Optional) The password used to secure the private key.",
		disableAutoRetriesAttrName: "(Optional) Disable automatic retries for retriable errors.\n" +
			"Automatic retries were introduced to solve some eventual consistency problems but it also introduced performance issues on destroy operations.",
		retryDurationSecondsAttrName: "(Optional) The minimum duration (in seconds) to retry a resource operation in response to an error.\n" +
			"The actual retry duration may be longer due to jittering of retry operations. This value is ignored if the `disable_auto_retries` field is set to true.",
		configFileProfileAttrName: "(Optional) The profile name to be used from config file, if not set it will be DEFAULT.",
	}
}

func Provider() terraform.ResourceProvider {
	ociProvider = &schema.Provider{
		DataSourcesMap: DataSourcesMap(),
		Schema:         schemaMap(),
		ResourcesMap:   ResourcesMap(),
		ConfigureFunc:  ProviderConfig,
	}
	return ociProvider
}

func schemaMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		authAttrName: {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  descriptions[authAttrName],
			DefaultFunc:  schema.MultiEnvDefaultFunc([]string{tfVarName(authAttrName), ociVarName(authAttrName)}, authAPIKeySetting),
			ValidateFunc: validation.StringInSlice([]string{authAPIKeySetting, authInstancePrincipalSetting, authInstancePrincipalWithCertsSetting, authSecurityToken}, true),
		},
		tenancyOcidAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[tenancyOcidAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(tenancyOcidAttrName), ociVarName(tenancyOcidAttrName)}, nil),
		},
		userOcidAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[userOcidAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(userOcidAttrName), ociVarName(userOcidAttrName)}, nil),
		},
		fingerprintAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[fingerprintAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(fingerprintAttrName), ociVarName(fingerprintAttrName)}, nil),
		},
		// Mostly used for testing. Don't put keys in your .tf files
		privateKeyAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: descriptions[privateKeyAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(privateKeyAttrName), ociVarName(privateKeyAttrName)}, nil),
		},
		privateKeyPathAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[privateKeyPathAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(privateKeyPathAttrName), ociVarName(privateKeyPathAttrName)}, nil),
		},
		privateKeyPasswordAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Default:     "",
			Description: descriptions[privateKeyPasswordAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(privateKeyPasswordAttrName), ociVarName(privateKeyPasswordAttrName)}, nil),
		},
		regionAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[regionAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(regionAttrName), ociVarName(regionAttrName)}, nil),
		},
		disableAutoRetriesAttrName: {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: descriptions[disableAutoRetriesAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(disableAutoRetriesAttrName), ociVarName(disableAutoRetriesAttrName)}, nil),
		},
		retryDurationSecondsAttrName: {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: descriptions[retryDurationSecondsAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(retryDurationSecondsAttrName), ociVarName(retryDurationSecondsAttrName)}, nil),
		},
		configFileProfileAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[configFileProfileAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(configFileProfileAttrName), ociVarName(configFileProfileAttrName)}, nil),
		},
	}
}

func RegisterResource(name string, resourceSchema *schema.Resource) {
	if OciResources == nil {
		OciResources = make(map[string]*schema.Resource)
	}
	OciResources[name] = resourceSchema
}

func RegisterDatasource(name string, datasourceSchema *schema.Resource) {
	if OciDatasources == nil {
		OciDatasources = make(map[string]*schema.Resource)
	}
	OciDatasources[name] = datasourceSchema
}

// This returns a map of all data sources to register with Terraform
// The OciDatasources map is populated by each datasource's init function being invoked before it gets here
func DataSourcesMap() map[string]*schema.Resource {
	// Register some aliases of registered datasources. These are registered for convenience and legacy reasons.
	RegisterDatasource("oci_core_listing_resource_version", CoreAppCatalogListingResourceVersionDataSource())
	RegisterDatasource("oci_core_listing_resource_versions", CoreAppCatalogListingResourceVersionsDataSource())
	RegisterDatasource("oci_core_shape", CoreShapesDataSource())
	RegisterDatasource("oci_core_virtual_networks", CoreVcnsDataSource())
	RegisterDatasource("oci_load_balancers", LoadBalancerLoadBalancersDataSource())
	RegisterDatasource("oci_load_balancer_backendsets", LoadBalancerBackendSetsDataSource())
	return OciDatasources
}

// This returns a map of all resources to register with Terraform
// The OciResource map is populated by each resource's init function being invoked before it gets here
func ResourcesMap() map[string]*schema.Resource {
	// Register some aliases of registered resources. These are registered for convenience and legacy reasons.
	RegisterResource("oci_core_virtual_network", CoreVcnResource())
	RegisterResource("oci_load_balancer", LoadBalancerLoadBalancerResource())
	RegisterResource("oci_load_balancer_backendset", LoadBalancerBackendSetResource())
	return OciResources
}

// Added for resource discovery AUTH
func getProviderEnvSettingWithDefault(s string, dv string) string {
	v := os.Getenv(tfEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(ociEnvPrefix + strings.ToUpper(s))
	if v != "" {
		return v
	}
	v = os.Getenv(s)
	if v != "" {
		return v
	}
	return dv
}

func getEnvSettingWithBlankDefault(s string) string {
	return getEnvSettingWithDefault(s, "")
}

func getEnvSettingWithDefault(s string, dv string) string {
	v := os.Getenv(tfEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(ociEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(s)
	if v != "" {
		return v
	}
	return dv
}

// Deprecated: There should be only no need to panic individually
func getRequiredEnvSetting(s string) string {
	v := getEnvSettingWithBlankDefault(s)
	if v == "" {
		panic(fmt.Sprintf("Required env setting %s is missing", s))
	}
	return v
}

func checkIncompatibleAttrsForApiKeyAuth(d *schema.ResourceData) ([]string, bool) {
	var apiKeyConfigAttributesToUnset []string
	for _, apiKeyConfigAttribute := range apiKeyConfigAttributes {
		apiKeyConfigAttributeValue, hasConfigVariable := d.GetOkExists(apiKeyConfigAttribute)
		if (hasConfigVariable && apiKeyConfigAttributeValue != "") || getEnvSettingWithBlankDefault(apiKeyConfigAttribute) != "" {
			apiKeyConfigAttributesToUnset = append(apiKeyConfigAttributesToUnset, apiKeyConfigAttribute)
		}
	}
	return apiKeyConfigAttributesToUnset, len(apiKeyConfigAttributesToUnset) == 0
}

func getCertificateFileBytes(certificateFileFullPath string) (pemRaw []byte, err error) {
	absFile, err := filepath.Abs(certificateFileFullPath)
	if err != nil {
		return nil, fmt.Errorf("can't form absolute path of %s: %v", certificateFileFullPath, err)
	}

	if pemRaw, err = ioutil.ReadFile(absFile); err != nil {
		return nil, fmt.Errorf("can't read %s: %v", certificateFileFullPath, err)
	}
	return
}

func ProviderConfig(d *schema.ResourceData) (interface{}, error) {
	clients := &OracleClients{
		sdkClientMap:  make(map[string]interface{}, len(oracleClientRegistrations.registeredClients)),
		configuration: make(map[string]string),
	}

	if d.Get(disableAutoRetriesAttrName).(bool) {
		shortRetryTime = 0
		longRetryTime = 0
	} else if retryDurationSeconds, exists := d.GetOkExists(retryDurationSecondsAttrName); exists {
		val := time.Duration(retryDurationSeconds.(int)) * time.Second
		if retryDurationSeconds.(int) < 0 {
			// Retry for maximum amount of time, if a negative value was specified
			val = time.Duration(math.MaxInt64)
		}
		configuredRetryDuration = &val
	}

	sdkConfigProvider, err := getSdkConfigProvider(d, clients)
	if err != nil {
		return nil, err
	}

	httpClient := buildHttpClient()

	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	configureClient, err = buildConfigureClientFn(sdkConfigProvider, httpClient)
	if err != nil {
		return nil, err
	}

	err = createSDKClients(clients, sdkConfigProvider, configureClient)
	if err != nil {
		return nil, err
	}

	avoidWaitingForDeleteTarget, _ = strconv.ParseBool(getEnvSettingWithDefault("avoid_waiting_for_delete_target", "false"))

	return clients, nil
}

func getSdkConfigProvider(d *schema.ResourceData, clients *OracleClients) (oci_common.ConfigurationProvider, error) {

	auth := strings.ToLower(d.Get(authAttrName).(string))
	profile := d.Get(configFileProfileAttrName).(string)
	clients.configuration[authAttrName] = auth

	configProviders, err := getConfigProviders(d, auth)
	if err != nil {
		return nil, err
	}
	resourceDataConfigProvider := ResourceDataConfigProvider{d}
	if region, error := resourceDataConfigProvider.Region(); error == nil {
		clients.configuration["region"] = region
	}

	// TODO: DefaultConfigProvider will return us a composingConfigurationProvider that reads from SDK config files,
	// and then from the environment variables ("TF_VAR" prefix). References to "TF_VAR" prefix should be removed from
	// the SDK, since it's Terraform specific. When that happens, we need to update this to pass in the right prefix.
	configProviders = append(configProviders, resourceDataConfigProvider)
	if profile == "" {
		configProviders = append(configProviders, oci_common.DefaultConfigProvider())
	} else {
		defaultPath := path.Join(getHomeFolder(), defaultConfigDirName, defaultConfigFileName)
		err := checkProfile(profile, defaultPath)
		if err != nil {
			return nil, err
		}
		configProviders = append(configProviders, oci_common.CustomProfileConfigProvider(defaultPath, profile))
	}
	sdkConfigProvider, err := oci_common.ComposingConfigurationProvider(configProviders)
	if err != nil {
		return nil, err
	}

	return sdkConfigProvider, nil
}

func getConfigProviders(d *schema.ResourceData, auth string) ([]oci_common.ConfigurationProvider, error) {
	var configProviders []oci_common.ConfigurationProvider

	switch auth {
	case strings.ToLower(authAPIKeySetting):
		// No additional config providers needed
	case strings.ToLower(authInstancePrincipalSetting):
		apiKeyConfigVariablesToUnset, ok := checkIncompatibleAttrsForApiKeyAuth(d)
		if !ok {
			return nil, fmt.Errorf(`user credentials %v should be removed from the configuration`, strings.Join(apiKeyConfigVariablesToUnset, ", "))
		}

		region, ok := d.GetOk(regionAttrName)
		if !ok {
			return nil, fmt.Errorf("can not get %s from Terraform configuration (InstancePrincipal)", regionAttrName)
		}

		// Used to modify InstancePrincipal auth clients so that `accept_local_certs` is honored for auth clients as well
		// These clients are created implicitly by SDK, and are not modified by the buildConfigureClientFn that usually does this for the other SDK clients
		instancePrincipalAuthClientModifier := func(client oci_common.HTTPRequestDispatcher) (oci_common.HTTPRequestDispatcher, error) {
			if acceptLocalCerts := getEnvSettingWithBlankDefault(acceptLocalCerts); acceptLocalCerts != "" {
				if bool, err := strconv.ParseBool(acceptLocalCerts); err == nil {
					modifiedClient := buildHttpClient()
					modifiedClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = bool
					return modifiedClient, nil
				}
			}
			return client, nil
		}

		cfg, err := oci_common_auth.InstancePrincipalConfigurationForRegionWithCustomClient(oci_common.StringToRegion(region.(string)), instancePrincipalAuthClientModifier)
		if err != nil {
			return nil, err
		}
		log.Printf("[DEBUG] Configuration provided by: %s", cfg)

		configProviders = append(configProviders, cfg)
	case strings.ToLower(authInstancePrincipalWithCertsSetting):
		apiKeyConfigVariablesToUnset, ok := checkIncompatibleAttrsForApiKeyAuth(d)
		if !ok {
			return nil, fmt.Errorf(`user credentials %v should be removed from the configuration`, strings.Join(apiKeyConfigVariablesToUnset, ", "))
		}

		region, ok := d.GetOkExists(regionAttrName)
		if !ok {
			return nil, fmt.Errorf("can not get %s from Terraform configuration (InstancePrincipalWithCerts)", regionAttrName)
		}

		defaultCertsDir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("can not get working directory for current os platform")
		}

		certsDir := filepath.Clean(getEnvSettingWithDefault("test_certificates_location", defaultCertsDir))
		leafCertificateBytes, err := getCertificateFileBytes(filepath.Join(certsDir, "ip_cert.pem"))
		if err != nil {
			return nil, fmt.Errorf("can not read leaf certificate from %s", filepath.Join(certsDir, "ip_cert.pem"))
		}

		leafPrivateKeyBytes, err := getCertificateFileBytes(filepath.Join(certsDir, "ip_key.pem"))
		if err != nil {
			return nil, fmt.Errorf("can not read leaf private key from %s", filepath.Join(certsDir, "ip_key.pem"))
		}

		leafPassphraseBytes := []byte{}
		if _, err := os.Stat(certsDir + "/leaf_passphrase"); !os.IsNotExist(err) {
			leafPassphraseBytes, err = getCertificateFileBytes(filepath.Join(certsDir + "leaf_passphrase"))
			if err != nil {
				return nil, fmt.Errorf("can not read leafPassphraseBytes from %s", filepath.Join(certsDir+"leaf_passphrase"))
			}
		}

		intermediateCertificateBytes, err := getCertificateFileBytes(filepath.Join(certsDir, "intermediate.pem"))
		if err != nil {
			return nil, fmt.Errorf("can not read intermediate certificate from %s", filepath.Join(certsDir, "intermediate.pem"))
		}

		intermediateCertificatesBytes := [][]byte{
			intermediateCertificateBytes,
		}

		cfg, err := oci_common_auth.InstancePrincipalConfigurationWithCerts(oci_common.StringToRegion(region.(string)), leafCertificateBytes, leafPassphraseBytes, leafPrivateKeyBytes, intermediateCertificatesBytes)
		if err != nil {
			return nil, err
		}
		log.Printf("[DEBUG] Configuration provided by: %s", cfg)

		configProviders = append(configProviders, cfg)

	case strings.ToLower(authSecurityToken):
		apiKeyConfigVariablesToUnset, ok := checkIncompatibleAttrsForApiKeyAuth(d)
		if !ok {
			return nil, fmt.Errorf(`user credentials %v should be removed from the configuration`, strings.Join(apiKeyConfigVariablesToUnset, ", "))
		}
		profile, ok := d.GetOk(configFileProfileAttrName)
		if !ok {
			return nil, fmt.Errorf("missing profile in provider block %v", configFileProfileAttrName)
		}
		profileString := profile.(string)
		defaultPath := path.Join(getHomeFolder(), defaultConfigDirName, defaultConfigFileName)
		if err := checkProfile(profileString, defaultPath); err != nil {
			return nil, err
		}
		securityTokenBasedAuthConfigProvider := oci_common.CustomProfileConfigProvider(defaultPath, profileString)

		keyId, err := securityTokenBasedAuthConfigProvider.KeyID()
		if err != nil || !strings.HasPrefix(keyId, "ST$") {
			return nil, fmt.Errorf("Security token is invalid ")
		}
		configProviders = append(configProviders, securityTokenBasedAuthConfigProvider)
	default:
		return nil, fmt.Errorf("auth must be one of '%s' or '%s' or '%s' or '%s'", authAPIKeySetting, authInstancePrincipalSetting, authInstancePrincipalWithCertsSetting, authSecurityToken)
	}

	return configProviders, nil
}

func buildHttpClient() (httpClient *http.Client) {
	httpClient = &http.Client{
		Timeout: defaultRequestTimeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: defaultConnectionTimeout,
			}).DialContext,
			TLSHandshakeTimeout: defaultTLSHandshakeTimeout,
			TLSClientConfig:     &tls.Config{MinVersion: tls.VersionTLS12},
			Proxy:               http.ProxyFromEnvironment,
		},
	}
	return
}

func buildConfigureClientFn(configProvider oci_common.ConfigurationProvider, httpClient *http.Client) (ConfigureClient, error) {

	if ociProvider != nil && len(ociProvider.TerraformVersion) > 0 {
		terraformCLIVersion = ociProvider.TerraformVersion
	}
	userAgentProviderName := getEnvSettingWithDefault(userAgentProviderNameEnv, defaultUserAgentProviderName)
	userAgent := fmt.Sprintf(userAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, sdkMeta.SDKVersionString(), terraformCLIVersion, userAgentProviderName, Version)

	useOboToken, err := strconv.ParseBool(getEnvSettingWithDefault("use_obo_token", "false"))
	if err != nil {
		return nil, err
	}

	simulateDb, _ := strconv.ParseBool(getEnvSettingWithDefault("simulate_db", "false"))

	requestSigner := oci_common.DefaultRequestSigner(configProvider)
	var oboTokenProvider OboTokenProvider
	oboTokenProvider = emptyOboTokenProvider{}
	if useOboToken {
		// Add Obo token to the default list and update the signer
		httpHeadersToSign := append(oci_common.DefaultGenericHeaders(), requestHeaderOpcOboToken)
		requestSigner = oci_common.RequestSigner(configProvider, httpHeadersToSign, oci_common.DefaultBodyHeaders())
		oboTokenProvider = oboTokenProviderFromEnv{}
	}

	configureClientFn := func(client *oci_common.BaseClient) error {
		client.HTTPClient = httpClient
		client.UserAgent = userAgent
		client.Signer = requestSigner
		client.Interceptor = func(r *http.Request) error {
			if oboToken, err := oboTokenProvider.OboToken(); err == nil && oboToken != "" {
				r.Header.Set(requestHeaderOpcOboToken, oboToken)
			}

			if simulateDb {
				if r.Method == http.MethodPost && (strings.Contains(r.URL.Path, "/dbSystems") ||
					strings.Contains(r.URL.Path, "/autonomousData") ||
					strings.Contains(r.URL.Path, "/dataGuardAssociations") ||
					strings.Contains(r.URL.Path, "/autonomousExadata") ||
					strings.Contains(r.URL.Path, "/autonomousContainer") ||
					strings.Contains(r.URL.Path, "/backupDestinations") ||
					strings.Contains(r.URL.Path, "/exadataInfrastructures") ||
					strings.Contains(r.URL.Path, "/vmClusters") ||
					strings.Contains(r.URL.Path, "/cloudExadataInfrastructures") ||
					strings.Contains(r.URL.Path, "/cloudVmClusters") ||
					strings.Contains(r.URL.Path, "/autonomousVmClusters") ||
					strings.Contains(r.URL.Path, "/externalnoncontainerdatabases") ||
					strings.Contains(r.URL.Path, "/externalcontainerdatabases") ||
					strings.Contains(r.URL.Path, "/externalpluggabledatabases") ||
					strings.Contains(r.URL.Path, "/externaldatabaseconnectors")) {
					r.Header.Set(requestHeaderOpcHostSerial, "FAKEHOSTSERIAL")
				}
			}
			return nil
		}

		domainNameOverride := getEnvSettingWithBlankDefault(domainNameOverrideEnv)

		if domainNameOverride != "" {
			re := regexp.MustCompile(`(.*?)[-\w]+\.\w+$`)                             // (capture: preamble) match: d0main-name . tld end-of-string
			client.Host = re.ReplaceAllString(client.Host, "${1}"+domainNameOverride) // non-match conveniently returns original string
		}

		customCertLoc := getEnvSettingWithBlankDefault(customCertLocationEnv)

		if customCertLoc != "" {
			cert, err := ioutil.ReadFile(customCertLoc)
			if err != nil {
				return err
			}
			pool := x509.NewCertPool()
			if ok := pool.AppendCertsFromPEM(cert); !ok {
				return fmt.Errorf("failed to append custom cert to the pool")
			}
			// install the certificates in the client
			httpClient.Transport.(*http.Transport).TLSClientConfig.RootCAs = pool
		}

		if acceptLocalCerts := getEnvSettingWithBlankDefault(acceptLocalCerts); acceptLocalCerts != "" {
			if bool, err := strconv.ParseBool(acceptLocalCerts); err == nil {
				httpClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = bool
			}
		}

		// install the hook for HTTP replaying
		if h, ok := client.HTTPClient.(*http.Client); ok {
			_, err := httpreplay.InstallRecorder(h)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return configureClientFn, nil
}

func getHomeFolder() string {
	if os.Getenv("TF_HOME_OVERRIDE") != "" {
		return os.Getenv("TF_HOME_OVERRIDE")
	}
	current, e := user.Current()
	if e != nil {
		//Give up and try to return something sensible
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return current.HomeDir
}

// cleans and expands the path if it contains a tilde , returns the expanded path or the input path as is if not expansion
// was performed
func expandPath(filepath string) string {
	if strings.HasPrefix(filepath, fmt.Sprintf("~%c", os.PathSeparator)) {
		filepath = path.Join(getHomeFolder(), filepath[2:])
	}
	return path.Clean(filepath)
}

func checkProfile(profile string, path string) (err error) {
	var profileRegex = regexp.MustCompile(`^\[(.*)\]`)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(data)
	splitContent := strings.Split(content, "\n")
	for _, line := range splitContent {
		if match := profileRegex.FindStringSubmatch(line); match != nil && len(match) > 1 && match[1] == profile {
			return nil
		}
	}

	return fmt.Errorf("configuration file did not contain profile: %s", profile)
}

type ResourceDataConfigProvider struct {
	D *schema.ResourceData
}

// TODO: The error messages returned by following methods get swallowed up by the ComposingConfigurationProvider,
// since it only checks whether an error exists or not.
// The ComposingConfigurationProvider in SDK should log the errors as debug statements instead.

func (p ResourceDataConfigProvider) AuthType() (oci_common.AuthConfig, error) {
	return oci_common.AuthConfig{
			AuthType:         oci_common.UnknownAuthenticationType,
			IsFromConfigFile: false,
			OboToken:         nil,
		},
		fmt.Errorf("unsupported, keep the interface")
}

func (p ResourceDataConfigProvider) TenancyOCID() (string, error) {
	if tenancyOCID, ok := p.D.GetOkExists(tenancyOcidAttrName); ok {
		return tenancyOCID.(string), nil
	}
	return "", fmt.Errorf("can not get %s from Terraform configuration", tenancyOcidAttrName)
}

func (p ResourceDataConfigProvider) UserOCID() (string, error) {
	if userOCID, ok := p.D.GetOkExists(userOcidAttrName); ok {
		return userOCID.(string), nil
	}
	return "", fmt.Errorf("can not get %s from Terraform configuration", userOcidAttrName)
}

func (p ResourceDataConfigProvider) KeyFingerprint() (string, error) {
	if fingerprint, ok := p.D.GetOkExists(fingerprintAttrName); ok {
		return fingerprint.(string), nil
	}
	return "", fmt.Errorf("can not get %s from Terraform configuration", fingerprintAttrName)
}

func (p ResourceDataConfigProvider) Region() (string, error) {
	if region, ok := p.D.GetOkExists(regionAttrName); ok {
		return region.(string), nil
	}
	return "", fmt.Errorf("can not get %s from Terraform configuration", regionAttrName)
}

func (p ResourceDataConfigProvider) KeyID() (string, error) {
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

func (p ResourceDataConfigProvider) PrivateRSAKey() (key *rsa.PrivateKey, err error) {
	password := ""
	if privateKeyPassword, hasPrivateKeyPassword := p.D.GetOkExists(privateKeyPasswordAttrName); hasPrivateKeyPassword {
		password = privateKeyPassword.(string)
	}

	if privateKey, hasPrivateKey := p.D.GetOkExists(privateKeyAttrName); hasPrivateKey {
		return oci_common.PrivateKeyFromBytes([]byte(privateKey.(string)), &password)
	}

	if privateKeyPath, hasPrivateKeyPath := p.D.GetOkExists(privateKeyPathAttrName); hasPrivateKeyPath {
		resolvedPath := expandPath(privateKeyPath.(string))
		pemFileContent, readFileErr := ioutil.ReadFile(resolvedPath)
		if readFileErr != nil {
			return nil, fmt.Errorf("can not read private key from: '%s', Error: %q", privateKeyPath, readFileErr)
		}
		return oci_common.PrivateKeyFromBytes(pemFileContent, &password)
	}

	return nil, fmt.Errorf("can not get private_key or private_key_path from Terraform configuration")
}
