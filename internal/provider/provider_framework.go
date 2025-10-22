package provider

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	gopath "path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/v65/common/auth"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	tf_resource "github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

type ociPluginProvider struct {
	configured                                  bool
	auth                                        string
	tenancyOcid                                 string
	userOcid                                    string
	fingerprint                                 string
	privateKey                                  string
	privateKeyPath                              string
	privateKeyPassword                          string
	region                                      string
	disableAutoRetries                          bool
	retryDurationSeconds                        int64
	configFileProfile                           string
	ignoreDefinedTags                           []string
	realmSpecificServiceEndpointTemplateEnabled bool
	testTimeMaintenanceRebootDue                string
	dualStackEndpointEnabled                    bool
	retriesConfigFile                           string
}

type ociProviderModel struct {
	Auth                                        types.String `tfsdk:"auth"`
	TenancyOcid                                 types.String `tfsdk:"tenancy_ocid"`
	UserOcid                                    types.String `tfsdk:"user_ocid"`
	Fingerprint                                 types.String `tfsdk:"fingerprint"`
	PrivateKey                                  types.String `tfsdk:"private_key"`
	PrivateKeyPath                              types.String `tfsdk:"private_key_path"`
	PrivateKeyPassword                          types.String `tfsdk:"private_key_password"`
	Region                                      types.String `tfsdk:"region"`
	DisableAutoRetries                          types.Bool   `tfsdk:"disable_auto_retries"`
	RetryDurationSeconds                        types.Int64  `tfsdk:"retry_duration_seconds"`
	ConfigFileProfile                           types.String `tfsdk:"config_file_profile"`
	IgnoreDefinedTags                           types.List   `tfsdk:"ignore_defined_tags"`
	RealmSpecificServiceEndpointTemplateEnabled types.Bool   `tfsdk:"realm_specific_service_endpoint_template_enabled"`
	DualStackEndpointEnabled                    types.Bool   `tfsdk:"dual_stack_endpoint_enabled"`
	TestTimeMaintenanceRebootDue                types.String `tfsdk:"test_time_maintenance_reboot_due"`
	RetriesConfigFile                           types.String `tfsdk:"retries_config_file"`
}

func New() provider.Provider {
	return &ociPluginProvider{}
}

func (p *ociPluginProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "oci"
}

func (p *ociPluginProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			globalvar.AuthAttrName: schema.StringAttribute{
				Optional:    true,
				Description: descriptions[globalvar.AuthAttrName],
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(globalvar.AuthAPIKeySetting, globalvar.AuthInstancePrincipalSetting, globalvar.AuthInstancePrincipalWithCertsSetting, globalvar.AuthSecurityToken, globalvar.ResourcePrincipal, globalvar.AuthOKEWorkloadIdentity),
				},
			},
			globalvar.TenancyOcidAttrName: schema.StringAttribute{
				Optional:    true,
				Description: descriptions[globalvar.TenancyOcidAttrName],
			},
			globalvar.UserOcidAttrName: schema.StringAttribute{
				Optional:    true,
				Description: descriptions[globalvar.UserOcidAttrName],
			},
			globalvar.FingerprintAttrName: schema.StringAttribute{
				Optional:    true,
				Description: descriptions[globalvar.FingerprintAttrName],
			},
			// Mostly used for testing. Don't put keys in your .tf files
			globalvar.PrivateKeyAttrName: schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: descriptions[globalvar.PrivateKeyAttrName],
			},
			globalvar.PrivateKeyPathAttrName: schema.StringAttribute{
				Optional:    true,
				Description: descriptions[globalvar.PrivateKeyPathAttrName],
			},
			globalvar.PrivateKeyPasswordAttrName: schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: descriptions[globalvar.PrivateKeyPasswordAttrName],
			},
			globalvar.RegionAttrName: schema.StringAttribute{
				Optional:    true,
				Description: descriptions[globalvar.RegionAttrName],
			},
			globalvar.DisableAutoRetriesAttrName: schema.BoolAttribute{
				Optional: true,
				//Default:     false,
				Description: descriptions[globalvar.DisableAutoRetriesAttrName],
				//DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.DisableAutoRetriesAttrName), ociVarName(globalvar.DisableAutoRetriesAttrName)}, nil),
			},
			globalvar.RetryDurationSecondsAttrName: schema.Int64Attribute{
				Optional:    true,
				Description: descriptions[globalvar.RetryDurationSecondsAttrName],
				//DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.RetryDurationSecondsAttrName), ociVarName(globalvar.RetryDurationSecondsAttrName)}, nil),
			},
			globalvar.ConfigFileProfileAttrName: schema.StringAttribute{
				Optional:    true,
				Description: descriptions[globalvar.ConfigFileProfileAttrName],
				//DefaultFunc: schema.MultiEnvDefaultFunc([]string{tfVarName(globalvar.ConfigFileProfileAttrName), ociVarName(globalvar.ConfigFileProfileAttrName)}, nil),
			},
			globalvar.DefinedTagsToIgnore: schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: descriptions[globalvar.DefinedTagsToIgnore],
				//MaxItems:    100,
			},
			globalvar.RealmSpecificServiceEndpointTemplateEnabled: schema.BoolAttribute{
				Optional:    true,
				Description: descriptions[globalvar.RealmSpecificServiceEndpointTemplateEnabled],
			},
			globalvar.DualStackEndpointEnabled: schema.BoolAttribute{
				Optional:    true,
				Description: descriptions[globalvar.DualStackEndpointEnabled],
			},
			// test_time_maintenance_reboot_due is used only in some acceptance tests to simulate some scenario
			globalvar.TestTimeMaintenanceRebootDue: schema.StringAttribute{
				Optional: true,
			},
			globalvar.RetriesConfigFile: schema.StringAttribute{
				Optional:    true,
				Description: descriptions[globalvar.RetriesConfigFile],
			},
		},
	}
}

func (p *ociPluginProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config ociProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set Defaults from Env variables, if not provided in tf config
	p.SetDefaults(&config)
	clients, err := p.SetProviderConfig()
	if err != nil {
		log.Println("PluginProvider Configure err....", err)
		return
	}

	// Set the provider data with custom field value
	resp.DataSourceData = clients
	resp.ResourceData = clients
}

func (p *ociPluginProvider) Resources(_ context.Context) []func() resource.Resource {
	return globalvar.OciFrameworkResources
}

func (p *ociPluginProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return globalvar.OciFrameworkDataSources
}

func (p *ociPluginProvider) SetDefaults(config *ociProviderModel) {
	if config.Auth.IsUnknown() || config.Auth.IsNull() {
		config.Auth = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.AuthAttrName), ociVarName(globalvar.AuthAttrName)}, globalvar.AuthAPIKeySetting))
	}

	if config.TenancyOcid.IsUnknown() || config.TenancyOcid.IsNull() {
		config.TenancyOcid = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.TenancyOcidAttrName), ociVarName(globalvar.TenancyOcidAttrName)}, ""))
	}

	if config.UserOcid.IsUnknown() || config.UserOcid.IsNull() {
		config.UserOcid = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.UserOcidAttrName), ociVarName(globalvar.UserOcidAttrName)}, ""))
	}

	if config.Fingerprint.IsUnknown() || config.Fingerprint.IsNull() {
		config.Fingerprint = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.FingerprintAttrName), ociVarName(globalvar.FingerprintAttrName)}, ""))
	}

	if config.PrivateKey.IsUnknown() || config.PrivateKey.IsNull() {
		config.PrivateKey = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.PrivateKeyAttrName), ociVarName(globalvar.PrivateKeyAttrName)}, ""))
	}

	if config.PrivateKeyPath.IsUnknown() || config.PrivateKeyPath.IsNull() {
		config.PrivateKeyPath = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.PrivateKeyPathAttrName), ociVarName(globalvar.PrivateKeyPathAttrName)}, ""))
	}

	if config.PrivateKeyPassword.IsUnknown() || config.PrivateKeyPath.IsNull() {
		config.PrivateKeyPassword = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.PrivateKeyPasswordAttrName), ociVarName(globalvar.PrivateKeyPasswordAttrName)}, ""))
	}

	if config.Region.IsUnknown() || config.Region.IsNull() {
		config.Region = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.RegionAttrName), ociVarName(globalvar.RegionAttrName)}, ""))
	}

	if config.DisableAutoRetries.IsUnknown() || config.DisableAutoRetries.IsNull() {
		config.DisableAutoRetries = types.BoolValue(MultiEnvDefaultBoolFunc([]string{tfVarName(globalvar.DisableAutoRetriesAttrName), ociVarName(globalvar.DisableAutoRetriesAttrName)}, false))
	}

	if config.RetryDurationSeconds.IsUnknown() || config.RetryDurationSeconds.IsNull() {
		config.RetryDurationSeconds = types.Int64Value(MultiEnvDefaultIntFunc([]string{tfVarName(globalvar.RetryDurationSecondsAttrName), ociVarName(globalvar.RetryDurationSecondsAttrName)}, 0))
	}

	if config.ConfigFileProfile.IsUnknown() || config.ConfigFileProfile.IsNull() {
		config.ConfigFileProfile = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.ConfigFileProfileAttrName), ociVarName(globalvar.ConfigFileProfileAttrName)}, ""))
	}

	if config.RetriesConfigFile.IsUnknown() || config.RetriesConfigFile.IsNull() {
		config.RetriesConfigFile = types.StringValue(MultiEnvDefaultFunc([]string{tfVarName(globalvar.RetriesConfigFile), ociVarName(globalvar.RetriesConfigFile)}, ""))
	}

	p.auth = config.Auth.ValueString()
	p.tenancyOcid = config.TenancyOcid.ValueString()
	p.userOcid = config.UserOcid.ValueString()
	p.fingerprint = config.Fingerprint.ValueString()
	p.privateKey = config.PrivateKey.ValueString()
	p.privateKeyPath = config.PrivateKeyPath.ValueString()
	p.privateKeyPassword = config.PrivateKeyPassword.ValueString()
	p.region = config.Region.ValueString()
	p.disableAutoRetries = config.DisableAutoRetries.ValueBool()
	p.retryDurationSeconds = config.RetryDurationSeconds.ValueInt64()
	p.configFileProfile = config.ConfigFileProfile.ValueString()
	p.configured = true
	tf_resource.RealmSpecificServiceEndpointTemplateEnabled = getStringFromFwBool(config.RealmSpecificServiceEndpointTemplateEnabled)
	tf_resource.DualStackEndpointTemplateEnabled = getStringFromFwBool(config.DualStackEndpointEnabled)
	p.retriesConfigFile = config.RetriesConfigFile.ValueString()

}

func (p *ociPluginProvider) SetProviderConfig() (interface{}, error) {
	//tf_resource.DefinedTagsToSuppress = IgnoreDefinedTags(req)
	clients := &tf_client.OracleClients{
		SdkClientMap:  make(map[string]interface{}, len(tf_client.OracleClientRegistrationsVar.RegisteredClients)),
		Configuration: make(map[string]string),
	}

	disableAutoRetries := p.disableAutoRetries

	retryDurationSeconds := p.retryDurationSeconds

	if disableAutoRetries {
		tf_resource.ShortRetryTime = 0
		tf_resource.LongRetryTime = 0
	} else if retryDurationSeconds > 0 {
		val := time.Duration(retryDurationSeconds) * time.Second
		if retryDurationSeconds < 0 {
			// Retry for maximum amount of time, if a negative value was specified
			val = time.Duration(globalvar.MaxInt64)
		}
		tf_resource.ConfiguredRetryDuration = &val
	}

	retriesConfigFile := p.retriesConfigFile
	if len(retriesConfigFile) > 0 {
		err := tf_resource.SetRetriesConfig(retriesConfigFile)
		if err != nil {
			return nil, err
		}
	}

	sdkConfigProvider, err := p._GetSdkConfigProvider(clients)
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

func (p *ociPluginProvider) _GetSdkConfigProvider(clients *tf_client.OracleClients) (oci_common.ConfigurationProvider, error) {
	auth := strings.ToLower(p.auth)
	profile := p.configFileProfile
	clients.Configuration[globalvar.AuthAttrName] = auth

	configProviders, err := p._getConfigProviders()
	if err != nil {
		return nil, err
	}
	configProvider := ConfigProvider{p}
	if region, error := configProvider.Region(); error == nil {
		clients.Configuration["region"] = region
	}

	//In GoSDK, the first step is to check if AuthType exists,
	//for composite provider, we only check the first provider in the list for the AuthType.
	//Then SDK will based on the AuthType to Create the actual provider if it's a valid value.
	//If not, then SDK will base on the order in the composite provider list to check for necessary info (tenancyid, userID, fingerprint, region, keyID).
	configProviders = append(configProviders, configProvider)
	if profile == "" {
		configProviders = append(configProviders, oci_common.DefaultConfigProvider())
	} else {
		defaultPath := gopath.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
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

func (p *ociPluginProvider) CheckIncompatibleAttrsForApiKeyAuthFw() ([]string, bool) {
	var apiKeyConfigAttributesToUnset []string

	if p.userOcid != "" || utils.GetEnvSettingWithBlankDefault(globalvar.UserOcidAttrName) != "" {
		apiKeyConfigAttributesToUnset = append(apiKeyConfigAttributesToUnset, globalvar.UserOcidAttrName)
	}
	if p.fingerprint != "" || utils.GetEnvSettingWithBlankDefault(globalvar.FingerprintAttrName) != "" {
		apiKeyConfigAttributesToUnset = append(apiKeyConfigAttributesToUnset, globalvar.FingerprintAttrName)
	}
	if p.privateKey != "" || utils.GetEnvSettingWithBlankDefault(globalvar.PrivateKeyAttrName) != "" {
		apiKeyConfigAttributesToUnset = append(apiKeyConfigAttributesToUnset, globalvar.PrivateKeyAttrName)
	}
	if p.privateKeyPath != "" || utils.GetEnvSettingWithBlankDefault(globalvar.PrivateKeyPathAttrName) != "" {
		apiKeyConfigAttributesToUnset = append(apiKeyConfigAttributesToUnset, globalvar.PrivateKeyPathAttrName)
	}
	if p.privateKeyPassword != "" || utils.GetEnvSettingWithBlankDefault(globalvar.PrivateKeyPasswordAttrName) != "" {
		apiKeyConfigAttributesToUnset = append(apiKeyConfigAttributesToUnset, globalvar.PrivateKeyPasswordAttrName)
	}

	return apiKeyConfigAttributesToUnset, len(apiKeyConfigAttributesToUnset) == 0
}

func (p *ociPluginProvider) _getConfigProviders() ([]oci_common.ConfigurationProvider, error) {
	var configProviders []oci_common.ConfigurationProvider

	auth := strings.ToLower(p.auth)

	switch auth {
	case strings.ToLower(globalvar.AuthAPIKeySetting):
		// No additional config providers needed
	case strings.ToLower(globalvar.AuthInstancePrincipalSetting):
		_, ok := p.CheckIncompatibleAttrsForApiKeyAuthFw()
		if !ok {
			log.Printf("[DEBUG] Ignoring all user credentials for %v authentication", auth)
		}

		region, ok := p.region, p.region != ""
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

		cfg, err := oci_common_auth.InstancePrincipalConfigurationForRegionWithCustomClient(oci_common.StringToRegion(region), instancePrincipalAuthClientModifier)
		if err != nil {
			return nil, err
		}
		log.Printf("[DEBUG] Configuration provided by: %s", cfg)

		configProviders = append(configProviders, cfg)
	case strings.ToLower(globalvar.AuthInstancePrincipalWithCertsSetting):
		_, ok := p.CheckIncompatibleAttrsForApiKeyAuthFw()
		if !ok {
			log.Printf("[DEBUG] Ignoring all user credentials for %v authentication", auth)
		}

		region, ok := p.region, p.region != ""
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

		cfg, err := oci_common_auth.InstancePrincipalConfigurationWithCerts(oci_common.StringToRegion(region), leafCertificateBytes, leafPassphraseBytes, leafPrivateKeyBytes, intermediateCertificatesBytes)
		if err != nil {
			return nil, err
		}
		log.Printf("[DEBUG] Configuration provided by: %s", cfg)

		configProviders = append(configProviders, cfg)
	case strings.ToLower(globalvar.AuthSecurityToken):
		region, ok := p.region, p.region != ""
		if !ok {
			return nil, fmt.Errorf("can not get %s from Terraform configuration (SecurityToken)", globalvar.RegionAttrName)
		}
		// if region is part of the provider block make sure it is part of the final configuration too, and overwrites the region in the profile. +
		regionProvider := oci_common.NewRawConfigurationProvider("", "", region, "", "", nil)
		configProviders = append(configProviders, regionProvider)

		profile, ok := p.configFileProfile, p.configFileProfile != ""
		if !ok {
			return nil, fmt.Errorf("missing profile in provider block %v", globalvar.ConfigFileProfileAttrName)
		}
		privateKeyPassword := p.privateKeyPassword
		privateKeyPasswordString := privateKeyPassword
		profileString := profile
		defaultPath := gopath.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName, globalvar.DefaultConfigFileName)
		if err := utils.CheckProfile(profileString, defaultPath); err != nil {
			return nil, err
		}
		securityTokenBasedAuthConfigProvider, err := oci_common.ConfigurationProviderForSessionTokenWithProfile(defaultPath, profileString, privateKeyPasswordString)
		if err != nil {
			return nil, fmt.Errorf("could not create security token based auth config provider %v", err)
		}
		configProviders = append(configProviders, securityTokenBasedAuthConfigProvider)
	case strings.ToLower(globalvar.ResourcePrincipal):
		var err error
		var resourcePrincipalAuthConfigProvider oci_common_auth.ConfigurationProviderWithClaimAccess
		region, ok := p.region, p.region != ""
		if !ok {
			log.Printf("did not get %s from Terraform configuration (ResourcePrincipal), falling back to environment variable", globalvar.RegionAttrName)
			resourcePrincipalAuthConfigProvider, err = oci_common_auth.ResourcePrincipalConfigurationProvider()
		} else {
			resourcePrincipalAuthConfigProvider, err = oci_common_auth.ResourcePrincipalConfigurationProviderForRegion(oci_common.StringToRegion(region))
		}
		if err != nil {
			return nil, err
		}
		configProviders = append(configProviders, resourcePrincipalAuthConfigProvider)
	case strings.ToLower(globalvar.AuthOKEWorkloadIdentity):
		okeWorkloadIdentityConfigProvider, err := oci_common_auth.OkeWorkloadIdentityConfigurationProvider()
		if err != nil {
			return nil, fmt.Errorf("can not get oke workload indentity based auth config provider %v", err)
		}
		configProviders = append(configProviders, okeWorkloadIdentityConfigProvider)
	default:
		return nil, fmt.Errorf("auth must be one of '%s' or '%s' or '%s' or '%s' or '%s' or '%s'", globalvar.AuthAPIKeySetting, globalvar.AuthInstancePrincipalSetting, globalvar.AuthInstancePrincipalWithCertsSetting, globalvar.AuthSecurityToken, globalvar.ResourcePrincipal, globalvar.AuthOKEWorkloadIdentity)
	}

	return configProviders, nil
}

type ConfigProvider struct {
	D *ociPluginProvider
}

// TODO: The error messages returned by following methods get swallowed up by the ComposingConfigurationProvider,
// since it only checks whether an error exists or not.
// The ComposingConfigurationProvider in SDK should log the errors as debug statements instead.

func (p ConfigProvider) AuthType() (oci_common.AuthConfig, error) {
	return oci_common.AuthConfig{
			AuthType:         oci_common.UnknownAuthenticationType,
			IsFromConfigFile: false,
			OboToken:         nil,
		},
		fmt.Errorf("unsupported, keep the interface")
}

func (p ConfigProvider) TenancyOCID() (string, error) {
	if boatTenancyOCID := utils.GetEnvSettingWithBlankDefault(globalvar.BoatTenancyOcidAttrName); boatTenancyOCID != "" {
		return boatTenancyOCID, nil
	}
	tenancyOcid := p.D.tenancyOcid
	if tenancyOcid == "" {
		return "", fmt.Errorf("can not get %s from Terraform configuration", globalvar.TenancyOcidAttrName)
	}
	return tenancyOcid, nil
}

func (p ConfigProvider) UserOCID() (string, error) {
	userOcid := p.D.userOcid
	if userOcid == "" {
		return "", fmt.Errorf("can not get %s from Terraform configuration", globalvar.UserOcidAttrName)
	}
	return userOcid, nil
}

func (p ConfigProvider) KeyFingerprint() (string, error) {
	fingerprint := p.D.fingerprint
	if fingerprint == "" {
		return "", fmt.Errorf("can not get %s from Terraform configuration", globalvar.FingerprintAttrName)
	}
	return fingerprint, nil
}

func (p ConfigProvider) Region() (string, error) {
	region := p.D.region
	if region == "" {
		return "", fmt.Errorf("can not get %s from Terraform configuration", globalvar.RegionAttrName)
	}
	return region, nil
}

func (p ConfigProvider) KeyID() (string, error) {
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

func (p ConfigProvider) PrivateRSAKey() (key *rsa.PrivateKey, err error) {
	password := ""
	privateKeyPassword := p.D.privateKeyPassword
	if privateKeyPassword != "" {
		password = privateKeyPassword
	}

	privateKey := p.D.privateKey
	if privateKey != "" {
		privateKey := strings.ReplaceAll(p.D.privateKey, "\\n", "\n") // Ensure \n is replaced by actual newlines
		return oci_common.PrivateKeyFromBytesWithPassword([]byte(privateKey), []byte(password))
	}

	privateKeyPath := p.D.privateKeyPath
	if privateKeyPath != "" {
		resolvedPath := utils.ExpandPath(privateKeyPath)
		pemFileContent, readFileErr := ioutil.ReadFile(resolvedPath)
		if readFileErr != nil {
			return nil, fmt.Errorf("can not read private key from: '%s', Error: %q", privateKeyPath, readFileErr)
		}
		return oci_common.PrivateKeyFromBytes(pemFileContent, &password)
	}

	return nil, fmt.Errorf("can not get private_key or private_key_path from Terraform configuration")
}

func MultiEnvDefaultFunc(ks []string, dv string) string {
	for _, k := range ks {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return dv
}

func MultiEnvDefaultBoolFunc(ks []string, dv bool) bool {
	for _, k := range ks {
		v := os.Getenv(k)
		if strings.EqualFold(v, "true") {
			return true
		}
	}
	return dv
}

func MultiEnvDefaultIntFunc(ks []string, dv int64) int64 {
	for _, k := range ks {
		if v := os.Getenv(k); v != "" {
			tmpInt64, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Printf("[INFO] Invalid value for the environment variable %s. Given value is %s but expected an integer value\n", k, v)
				return dv
			}
			return tmpInt64
		}
	}
	return dv
}

func getStringFromFwBool(val types.Bool) string {
	if !(val.IsUnknown() || val.IsNull()) {
		return strconv.FormatBool(val.ValueBool())
	}
	return ""
}
