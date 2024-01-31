package globalvar

import (
	"time"
)

const (
	AuthAPIKeySetting                     = "ApiKey"
	AuthInstancePrincipalSetting          = "InstancePrincipal"
	AuthInstancePrincipalWithCertsSetting = "InstancePrincipalWithCerts"
	AuthSecurityToken                     = "SecurityToken"
	ResourcePrincipal                     = "ResourcePrincipal"
	RequestHeaderOpcOboToken              = "opc-obo-token"
	RequestHeaderOpcHostSerial            = "opc-host-serial"
	DefaultRequestTimeout                 = 0
	DefaultConnectionTimeout              = 10 * time.Second
	DefaultTLSHandshakeTimeout            = 10 * time.Second
	DefaultUserAgentProviderName          = "Oracle-TerraformProvider"
	UnknownTerraformCLIVersion            = "unknown"
	TestTerraformCLIVersion               = "test"
	UserAgentFormatter                    = "Oracle-GoSDK/%s (go/%s; %s/%s; terraform/%s; terraform-cli/%s) %s/%s"
	UserAgentProviderNameEnv              = "USER_AGENT_PROVIDER_NAME"
	UserAgentTerraformNameEnv             = "TF_APPEND_USER_AGENT"
	UserAgentSDKNameEnv                   = "OCI_SDK_APPEND_USER_AGENT"
	DomainNameOverrideEnv                 = "domain_name_override"
	HasCorrectDomainNameEnv               = "has_correct_domain_name"
	ClientHostOverridesEnv                = "CLIENT_HOST_OVERRIDES"
	CustomCertLocationEnv                 = "custom_cert_location"
	AcceptLocalCerts                      = "accept_local_certs"
	JobOCID                               = "job-ocid"

	AuthAttrName                                = "auth"
	TenancyOcidAttrName                         = "tenancy_ocid"
	BoatTenancyOcidAttrName                     = "boat_tenancy_ocid"
	UserOcidAttrName                            = "user_ocid"
	FingerprintAttrName                         = "fingerprint"
	PrivateKeyAttrName                          = "private_key"
	PrivateKeyPathAttrName                      = "private_key_path"
	PrivateKeyPasswordAttrName                  = "private_key_password"
	RegionAttrName                              = "region"
	DisableAutoRetriesAttrName                  = "disable_auto_retries"
	RetryDurationSecondsAttrName                = "retry_duration_seconds"
	OboTokenAttrName                            = "obo_token"
	OboTokenPath                                = "obo_token_path"
	ConfigFileProfileAttrName                   = "config_file_profile"
	DefinedTagsToIgnore                         = "ignore_defined_tags"
	RealmSpecificServiceEndpointTemplateEnabled = "realm_specific_service_endpoint_template_enabled"

	DefaultConfigFileName    = "config"
	DefaultConfigDirName     = ".oci"
	ColonDelimiter           = ";"
	EqualToOperatorDelimiter = "="
	DotDelimiter             = "."

	// Resource Discovery
	ExportUserAgentFormatter        = "Oracle-GoSDK/%s (go/%s; %s/%s; terraform-oci-exporter/%s)"
	DefaultTmpStateFile             = "terraform.tfstate.tmp"
	DefaultStateFilename            = "terraform.tfstate"
	VarsFile                        = "vars.tf"
	ProviderFile                    = "provider.tf"
	MissingRequiredAttributeWarning = `

Warning: There are one or more 'Required' attributes for which a value could not be discovered.
This may be expected behavior from the service, which may prevent discovery of certain sensitive attributes or secrets.

Placeholder values have been added for such attributes with a comment "Required attribute not found in discovery, placeholder value set to avoid plan failure".
These missing attributes are also added to the lifecycle ignore_changes.
`
	PlaceholderValueForMissingAttribute = `<placeholder for missing required attribute>`
	EnvLogFile                          = "TF_LOG_PATH"
	EnvOCITFLogFile                     = "OCI_TF_LOG_PATH" // Log path for Custom TF logger - TFProviderLogger
	TerraformBinPathName                = "terraform_bin_path"
	MaxInt64                            = 1<<63 - 1 // TODO : Fix needed for GoLang SDK v1.17.2
	DiscoverAllStatesEnv                = "TF_DISCOVER_ALL_STATES"
)

const (
	SubnetService       = "subnet"
	CoreService         = "core"
	LoadBalancerService = "loadbalancer"
	WorkRequest         = "workrequests"
	DeleteResource      = "delete"
	TfEnvPrefix         = "TF_VAR_"
	OciEnvPrefix        = "OCI_"
)
const (
	DebugTestSteps               = "DEBUG_TEST_STEPS"
	DebugTestStepsShowConfigOnly = "DEBUG_TEST_STEPS_SHOW_CONFIG_ONLY"
	SecurityTokenProfileForTest  = "terraform-federation-test"
)
const TerraformDocumentLink = "https://registry.terraform.io/providers/oracle/oci/latest/docs/"
