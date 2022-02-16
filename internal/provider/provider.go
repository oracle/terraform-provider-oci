// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	tf_core "github.com/terraform-providers/terraform-provider-oci/internal/service/core"
	tf_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/load_balancer"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	sdkMeta "github.com/hashicorp/terraform-plugin-sdk/meta"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"runtime"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/v58/common/auth"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	tf_resource "github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var descriptions map[string]string
var ApiKeyConfigAttributes = [5]string{globalvar.UserOcidAttrName, globalvar.FingerprintAttrName, globalvar.PrivateKeyAttrName, globalvar.PrivateKeyPathAttrName, globalvar.PrivateKeyPasswordAttrName}
var ociProvider *schema.Provider

var TerraformCLIVersion = globalvar.UnknownTerraformCLIVersion
var AvoidWaitingForDeleteTarget bool

var OciResources map[string]*schema.Resource
var OciDatasources map[string]*schema.Resource

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
	if path := utils.GetEnvSettingWithBlankDefault(globalvar.OboTokenPath); path != "" {
		token, err := utils.GetTokenFromFile(path)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return utils.GetEnvSettingWithBlankDefault(globalvar.OboTokenAttrName), nil
}

func tfVarName(attrName string) string {
	return utils.TfEnvPrefix + attrName
}

func ociVarName(attrName string) string {
	return utils.OciEnvPrefix + strings.ToUpper(attrName)
}

func init() {
	descriptions = map[string]string{
		globalvar.AuthAttrName:        fmt.Sprintf("(Optional) The type of auth to use. Options are '%s', '%s' and '%s'. By default, '%s' will be used.", globalvar.AuthAPIKeySetting, globalvar.AuthSecurityToken, globalvar.AuthInstancePrincipalSetting, globalvar.AuthAPIKeySetting),
		globalvar.TenancyOcidAttrName: fmt.Sprintf("(Optional) The tenancy OCID for a user. The tenancy OCID can be found at the bottom of user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", globalvar.AuthAPIKeySetting),
		globalvar.UserOcidAttrName:    fmt.Sprintf("(Optional) The user OCID. This can be found in user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", globalvar.AuthAPIKeySetting),
		globalvar.FingerprintAttrName: fmt.Sprintf("(Optional) The fingerprint for the user's RSA key. This can be found in user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", globalvar.AuthAPIKeySetting),
		globalvar.RegionAttrName:      "(Required) The region for API connections (e.g. us-ashburn-1).",
		globalvar.PrivateKeyAttrName: "(Optional) A PEM formatted RSA private key for the user.\n" +
			fmt.Sprintf("A private_key or a private_key_path must be provided if auth is set to '%s', ignored otherwise.", globalvar.AuthAPIKeySetting),
		globalvar.PrivateKeyPathAttrName: "(Optional) The path to the user's PEM formatted private key.\n" +
			fmt.Sprintf("A private_key or a private_key_path must be provided if auth is set to '%s', ignored otherwise.", globalvar.AuthAPIKeySetting),
		globalvar.PrivateKeyPasswordAttrName: "(Optional) The password used to secure the private key.",
		globalvar.DisableAutoRetriesAttrName: "(Optional) Disable automatic retries for retriable errors.\n" +
			"Automatic retries were introduced to solve some eventual consistency problems but it also introduced performance issues on destroy operations.",
		globalvar.RetryDurationSecondsAttrName: "(Optional) The minimum duration (in seconds) to retry a resource operation in response to an error.\n" +
			"The actual retry duration may be longer due to jittering of retry operations. This value is ignored if the `disable_auto_retries` field is set to true.",
		globalvar.ConfigFileProfileAttrName: "(Optional) The profile name to be used from config file, if not set it will be DEFAULT.",
	}
}

func Provider() terraform.ResourceProvider {
	ociProvider = &schema.Provider{
		DataSourcesMap: DataSourcesMap(),
		Schema:         SchemaMap(),
		ResourcesMap:   ResourcesMap(),
		ConfigureFunc:  ProviderConfig,
	}
	return ociProvider
}

func SchemaMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		globalvar.AuthAttrName: {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  descriptions[globalvar.AuthAttrName],
			DefaultFunc:  schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.AuthAttrName), ociVarName(globalvar.AuthAttrName)}, globalvar.AuthAPIKeySetting),
			ValidateFunc: validation.StringInSlice([]string{globalvar.AuthAPIKeySetting, globalvar.AuthInstancePrincipalSetting, globalvar.AuthInstancePrincipalWithCertsSetting, globalvar.AuthSecurityToken}, true),
		},
		globalvar.TenancyOcidAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[globalvar.TenancyOcidAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.TenancyOcidAttrName), ociVarName(globalvar.TenancyOcidAttrName)}, nil),
		},
		globalvar.UserOcidAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[globalvar.UserOcidAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.UserOcidAttrName), ociVarName(globalvar.UserOcidAttrName)}, nil),
		},
		globalvar.FingerprintAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[globalvar.FingerprintAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.FingerprintAttrName), ociVarName(globalvar.FingerprintAttrName)}, nil),
		},
		// Mostly used for testing. Don't put keys in your .tf files
		globalvar.PrivateKeyAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: descriptions[globalvar.PrivateKeyAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.PrivateKeyAttrName), ociVarName(globalvar.PrivateKeyAttrName)}, nil),
		},
		globalvar.PrivateKeyPathAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[globalvar.PrivateKeyPathAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.PrivateKeyPathAttrName), ociVarName(globalvar.PrivateKeyPathAttrName)}, nil),
		},
		globalvar.PrivateKeyPasswordAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Default:     "",
			Description: descriptions[globalvar.PrivateKeyPasswordAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.PrivateKeyPasswordAttrName), ociVarName(globalvar.PrivateKeyPasswordAttrName)}, nil),
		},
		globalvar.RegionAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[globalvar.RegionAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.RegionAttrName), ociVarName(globalvar.RegionAttrName)}, nil),
		},
		globalvar.DisableAutoRetriesAttrName: {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: descriptions[globalvar.DisableAutoRetriesAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.DisableAutoRetriesAttrName), ociVarName(globalvar.DisableAutoRetriesAttrName)}, nil),
		},
		globalvar.RetryDurationSecondsAttrName: {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: descriptions[globalvar.RetryDurationSecondsAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.RetryDurationSecondsAttrName), ociVarName(globalvar.RetryDurationSecondsAttrName)}, nil),
		},
		globalvar.ConfigFileProfileAttrName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions[globalvar.ConfigFileProfileAttrName],
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.ConfigFileProfileAttrName), ociVarName(globalvar.ConfigFileProfileAttrName)}, nil),
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
	RegisterDatasource("oci_core_listing_resource_version", tf_core.CoreAppCatalogListingResourceVersionDataSource())
	RegisterDatasource("oci_core_listing_resource_versions", tf_core.CoreAppCatalogListingResourceVersionsDataSource())
	RegisterDatasource("oci_core_shape", tf_core.CoreShapesDataSource())
	RegisterDatasource("oci_core_virtual_networks", tf_core.CoreVcnsDataSource())
	RegisterDatasource("oci_load_balancers", tf_load_balancer.LoadBalancerLoadBalancersDataSource())
	RegisterDatasource("oci_load_balancer_backendsets", tf_load_balancer.LoadBalancerBackendSetsDataSource())
	return OciDatasources
}

// This returns a map of all resources to register with Terraform
// The OciResource map is populated by each resource's init function being invoked before it gets here
func ResourcesMap() map[string]*schema.Resource {
	// Register some aliases of registered resources. These are registered for convenience and legacy reasons.
	RegisterResource("oci_core_virtual_network", tf_core.CoreVcnResource())
	RegisterResource("oci_load_balancer", tf_load_balancer.LoadBalancerLoadBalancerResource())
	RegisterResource("oci_load_balancer_backendset", tf_load_balancer.LoadBalancerBackendSetResource())
	return OciResources
}

func ProviderConfig(d *schema.ResourceData) (interface{}, error) {
	clients := &tf_client.OracleClients{
		SdkClientMap:  make(map[string]interface{}, len(tf_client.OracleClientRegistrationsVar.RegisteredClients)),
		Configuration: make(map[string]string),
	}

	if d.Get(globalvar.DisableAutoRetriesAttrName).(bool) {
		tf_resource.ShortRetryTime = 0
		tf_resource.LongRetryTime = 0
	} else if retryDurationSeconds, exists := d.GetOkExists(globalvar.RetryDurationSecondsAttrName); exists {
		val := time.Duration(retryDurationSeconds.(int)) * time.Second
		if retryDurationSeconds.(int) < 0 {
			// Retry for maximum amount of time, if a negative value was specified
			val = time.Duration(globalvar.MaxInt64)
		}
		tf_resource.ConfiguredRetryDuration = &val
	}

	sdkConfigProvider, err := GetSdkConfigProvider(d, clients)
	if err != nil {
		return nil, err
	}

	httpClient := BuildHttpClient()

	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	tf_client.ConfigureClientVar, err = BuildConfigureClientFn(sdkConfigProvider, httpClient)
	if err != nil {
		return nil, err
	}

	err = tf_client.CreateSDKClients(clients, sdkConfigProvider, tf_client.ConfigureClientVar)
	if err != nil {
		return nil, err
	}

	AvoidWaitingForDeleteTarget, _ = strconv.ParseBool(utils.GetEnvSettingWithDefault("avoid_waiting_for_delete_target", "false"))

	return clients, nil
}

func GetSdkConfigProvider(d *schema.ResourceData, clients *tf_client.OracleClients) (oci_common.ConfigurationProvider, error) {

	auth := strings.ToLower(d.Get(globalvar.AuthAttrName).(string))
	profile := d.Get(globalvar.ConfigFileProfileAttrName).(string)
	clients.Configuration[globalvar.AuthAttrName] = auth

	configProviders, err := getConfigProviders(d, auth)
	if err != nil {
		return nil, err
	}
	resourceDataConfigProvider := ResourceDataConfigProvider{d}
	if region, error := resourceDataConfigProvider.Region(); error == nil {
		clients.Configuration["region"] = region
	}

	//In GoSDK, the first step is to check if AuthType exists,
	//for composite provider, we only check the first provider in the list for the AuthType.
	//Then SDK will based on the AuthType to Create the actual provider if it's a valid value.
	//If not, then SDK will base on the order in the composite provider list to check for necessary info (tenancyid, userID, fingerprint, region, keyID).
	configProviders = append(configProviders, resourceDataConfigProvider)
	if profile == "" {
		configProviders = append(configProviders, oci_common.DefaultConfigProvider())
	} else {
		defaultPath := path.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
		err := utils.CheckProfile(profile, defaultPath)
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
	case strings.ToLower(globalvar.AuthAPIKeySetting):
		// No additional config providers needed
	case strings.ToLower(globalvar.AuthInstancePrincipalSetting):
		_, ok := utils.CheckIncompatibleAttrsForApiKeyAuth(d, ApiKeyConfigAttributes)
		if !ok {
			log.Printf("[DEBUG] Ignoring all user credentials for %v authentication", auth)
		}

		region, ok := d.GetOk(globalvar.RegionAttrName)
		if !ok {
			return nil, fmt.Errorf("can not get %s from Terraform configuration (InstancePrincipal)", globalvar.RegionAttrName)
		}

		// Used to modify InstancePrincipal auth clients so that `accept_local_certs` is honored for auth clients as well
		// These clients are created implicitly by SDK, and are not modified by the utils.BuildConfigureClientFn that usually does this for the other SDK clients
		instancePrincipalAuthClientModifier := func(client oci_common.HTTPRequestDispatcher) (oci_common.HTTPRequestDispatcher, error) {
			if acceptLocalCerts := utils.GetEnvSettingWithBlankDefault(globalvar.AcceptLocalCerts); acceptLocalCerts != "" {
				if bool, err := strconv.ParseBool(acceptLocalCerts); err == nil {
					modifiedClient := BuildHttpClient()
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
	case strings.ToLower(globalvar.AuthInstancePrincipalWithCertsSetting):
		_, ok := utils.CheckIncompatibleAttrsForApiKeyAuth(d, ApiKeyConfigAttributes)
		if !ok {
			log.Printf("[DEBUG] Ignoring all user credentials for %v authentication", auth)
		}

		region, ok := d.GetOkExists(globalvar.RegionAttrName)
		if !ok {
			return nil, fmt.Errorf("can not get %s from Terraform configuration (InstancePrincipalWithCerts)", globalvar.RegionAttrName)
		}

		defaultCertsDir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("can not get working directory for current os platform")
		}

		certsDir := filepath.Clean(utils.GetEnvSettingWithDefault("test_certificates_location", defaultCertsDir))
		leafCertificateBytes, err := utils.GetCertificateFileBytes(filepath.Join(certsDir, "ip_cert.pem"))
		if err != nil {
			return nil, fmt.Errorf("can not read leaf certificate from %s", filepath.Join(certsDir, "ip_cert.pem"))
		}

		leafPrivateKeyBytes, err := utils.GetCertificateFileBytes(filepath.Join(certsDir, "ip_key.pem"))
		if err != nil {
			return nil, fmt.Errorf("can not read leaf private key from %s", filepath.Join(certsDir, "ip_key.pem"))
		}

		leafPassphraseBytes := []byte{}
		if _, err := os.Stat(certsDir + "/leaf_passphrase"); !os.IsNotExist(err) {
			leafPassphraseBytes, err = utils.GetCertificateFileBytes(filepath.Join(certsDir + "leaf_passphrase"))
			if err != nil {
				return nil, fmt.Errorf("can not read leafPassphraseBytes from %s", filepath.Join(certsDir+"leaf_passphrase"))
			}
		}

		intermediateCertificateBytes, err := utils.GetCertificateFileBytes(filepath.Join(certsDir, "intermediate.pem"))
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

	case strings.ToLower(globalvar.AuthSecurityToken):
		_, ok := utils.CheckIncompatibleAttrsForApiKeyAuth(d, ApiKeyConfigAttributes)
		if !ok {
			log.Printf("[DEBUG] Ignoring all user credentials for %v authentication", auth)
		}

		region, ok := d.GetOk(globalvar.RegionAttrName)
		if !ok {
			return nil, fmt.Errorf("can not get %s from Terraform configuration (InstancePrincipal)", globalvar.RegionAttrName)
		}
		// if region is part of the provider block make sure it is part of the final configuration too, and overwrites the region in the profile. +
		regionProvider := oci_common.NewRawConfigurationProvider("", "", region.(string), "", "", nil)
		configProviders = append(configProviders, regionProvider)

		profile, ok := d.GetOk(globalvar.ConfigFileProfileAttrName)
		if !ok {
			return nil, fmt.Errorf("missing profile in provider block %v", globalvar.ConfigFileProfileAttrName)
		}
		profileString := profile.(string)
		defaultPath := path.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
		if err := utils.CheckProfile(profileString, defaultPath); err != nil {
			return nil, err
		}
		securityTokenBasedAuthConfigProvider := oci_common.CustomProfileConfigProvider(defaultPath, profileString)

		keyId, err := securityTokenBasedAuthConfigProvider.KeyID()
		if err != nil || !strings.HasPrefix(keyId, "ST$") {
			return nil, fmt.Errorf("Security token is invalid ")
		}
		configProviders = append(configProviders, securityTokenBasedAuthConfigProvider)
	default:
		return nil, fmt.Errorf("auth must be one of '%s' or '%s' or '%s' or '%s'", globalvar.AuthAPIKeySetting, globalvar.AuthInstancePrincipalSetting, globalvar.AuthInstancePrincipalWithCertsSetting, globalvar.AuthSecurityToken)
	}

	return configProviders, nil
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
	if boatTenancyOCID := utils.GetEnvSettingWithBlankDefault(globalvar.BoatTenancyOcidAttrName); boatTenancyOCID != "" {
		return boatTenancyOCID, nil
	}
	if tenancyOCID, ok := p.D.GetOkExists(globalvar.TenancyOcidAttrName); ok {
		return tenancyOCID.(string), nil
	}
	return "", fmt.Errorf("can not get %s from Terraform configuration", globalvar.TenancyOcidAttrName)
}

func (p ResourceDataConfigProvider) UserOCID() (string, error) {
	if userOCID, ok := p.D.GetOkExists(globalvar.UserOcidAttrName); ok {
		return userOCID.(string), nil
	}
	return "", fmt.Errorf("can not get %s from Terraform configuration", globalvar.UserOcidAttrName)
}

func (p ResourceDataConfigProvider) KeyFingerprint() (string, error) {
	if fingerprint, ok := p.D.GetOkExists(globalvar.FingerprintAttrName); ok {
		return fingerprint.(string), nil
	}
	return "", fmt.Errorf("can not get %s from Terraform configuration", globalvar.FingerprintAttrName)
}

func (p ResourceDataConfigProvider) Region() (string, error) {
	if region, ok := p.D.GetOkExists(globalvar.RegionAttrName); ok {
		return region.(string), nil
	}
	return "", fmt.Errorf("can not get %s from Terraform configuration", globalvar.RegionAttrName)
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
	if privateKeyPassword, hasPrivateKeyPassword := p.D.GetOkExists(globalvar.PrivateKeyPasswordAttrName); hasPrivateKeyPassword {
		password = privateKeyPassword.(string)
	}

	if privateKey, hasPrivateKey := p.D.GetOkExists(globalvar.PrivateKeyAttrName); hasPrivateKey {
		return oci_common.PrivateKeyFromBytes([]byte(privateKey.(string)), &password)
	}

	if privateKeyPath, hasPrivateKeyPath := p.D.GetOkExists(globalvar.PrivateKeyPathAttrName); hasPrivateKeyPath {
		resolvedPath := utils.ExpandPath(privateKeyPath.(string))
		pemFileContent, readFileErr := ioutil.ReadFile(resolvedPath)
		if readFileErr != nil {
			return nil, fmt.Errorf("can not read private key from: '%s', Error: %q", privateKeyPath, readFileErr)
		}
		return oci_common.PrivateKeyFromBytes(pemFileContent, &password)
	}

	return nil, fmt.Errorf("can not get private_key or private_key_path from Terraform configuration")
}
func BuildHttpClient() (httpClient *http.Client) {
	httpClient = &http.Client{
		Timeout: globalvar.DefaultRequestTimeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: globalvar.DefaultConnectionTimeout,
			}).DialContext,
			TLSHandshakeTimeout: globalvar.DefaultTLSHandshakeTimeout,
			TLSClientConfig:     &tls.Config{MinVersion: tls.VersionTLS12},
			Proxy:               http.ProxyFromEnvironment,
		},
	}
	return
}

func BuildConfigureClientFn(configProvider oci_common.ConfigurationProvider, httpClient *http.Client) (tf_client.ConfigureClient, error) {

	if ociProvider != nil && len(ociProvider.TerraformVersion) > 0 {
		TerraformCLIVersion = ociProvider.TerraformVersion
	}
	userAgentProviderName := utils.GetEnvSettingWithDefault(globalvar.UserAgentProviderNameEnv, globalvar.DefaultUserAgentProviderName)
	userAgent := fmt.Sprintf(globalvar.UserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, sdkMeta.SDKVersionString(), TerraformCLIVersion, userAgentProviderName, globalvar.Version)

	useOboToken, err := strconv.ParseBool(utils.GetEnvSettingWithDefault("use_obo_token", "false"))
	if err != nil {
		return nil, err
	}

	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))

	requestSigner := oci_common.DefaultRequestSigner(configProvider)
	var oboTokenProvider OboTokenProvider
	oboTokenProvider = emptyOboTokenProvider{}
	if useOboToken {
		// Add Obo token to the default list and Update the signer
		httpHeadersToSign := append(oci_common.DefaultGenericHeaders(), globalvar.RequestHeaderOpcOboToken)
		requestSigner = oci_common.RequestSigner(configProvider, httpHeadersToSign, oci_common.DefaultBodyHeaders())
		oboTokenProvider = oboTokenProviderFromEnv{}
	}

	configureClientFn := func(client *oci_common.BaseClient) error {
		client.HTTPClient = httpClient
		client.UserAgent = userAgent
		client.Signer = requestSigner
		client.Interceptor = func(r *http.Request) error {
			if oboToken, err := oboTokenProvider.OboToken(); err == nil && oboToken != "" {
				r.Header.Set(globalvar.RequestHeaderOpcOboToken, oboToken)
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
					r.Header.Set(globalvar.RequestHeaderOpcHostSerial, "FAKEHOSTSERIAL")
				}
			}
			return nil
		}

		domainNameOverride := utils.GetEnvSettingWithBlankDefault(globalvar.DomainNameOverrideEnv)

		if domainNameOverride != "" {
			hasCorrectDomainName := utils.GetEnvSettingWithBlankDefault(globalvar.HasCorrectDomainNameEnv)
			re := regexp.MustCompile(`(.*?)[-\w]+\.\w+$`) // (capture: preamble) match: d0main-name . tld end-of-string
			if hasCorrectDomainName == "" || !strings.HasSuffix(client.Host, hasCorrectDomainName) {
				client.Host = re.ReplaceAllString(client.Host, "${1}"+domainNameOverride) // non-match conveniently returns original string
			}
		}

		customCertLoc := utils.GetEnvSettingWithBlankDefault(globalvar.CustomCertLocationEnv)

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

		if acceptLocalCerts := utils.GetEnvSettingWithBlankDefault(globalvar.AcceptLocalCerts); acceptLocalCerts != "" {
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
