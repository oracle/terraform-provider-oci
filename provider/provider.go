// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/hashicorp/terraform/terraform"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/common/auth"
	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_database "github.com/oracle/oci-go-sdk/database"
	oci_dns "github.com/oracle/oci-go-sdk/dns"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

var descriptions map[string]string
var disableAutoRetries bool

const (
	authAPIKeySetting            = "ApiKey"
	authInstancePrincipalSetting = "InstancePrincipal"
	defaultRequestTimeout        = 0
	defaultConnectionTimeout     = 10 * time.Second
	defaultTLSHandshakeTimeout   = 5 * time.Second
	userAgentFormatter           = "Oracle-GoSDK/%s (go/%s; %s/%s; terraform/%s) Oracle-TerraformProvider/%s"
	r1CertLocationEnv            = "R1_CERT_LOCATION"
)

type oboTokenProviderFromEnv struct{}

func (p oboTokenProviderFromEnv) OboToken() (string, error) {
	return getEnvSetting("obo_token", ""), nil
}

func init() {
	descriptions = map[string]string{
		"auth":         fmt.Sprintf("(Optional) The type of auth to use. Options are '%s' and '%s'. By default, '%s' will be used.", authAPIKeySetting, authInstancePrincipalSetting, authAPIKeySetting),
		"tenancy_ocid": fmt.Sprintf("(Optional) The tenancy OCID for a user. The tenancy OCID can be found at the bottom of user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		"user_ocid":    fmt.Sprintf("(Optional) The user OCID. This can be found in user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		"fingerprint":  fmt.Sprintf("(Optional) The fingerprint for the user's RSA key. This can be found in user settings in the Oracle Cloud Infrastructure console. Required if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		"region":       "(Required) The region for API connections (e.g. us-ashburn-1).",
		"private_key": "(Optional) A PEM formatted RSA private key for the user.\n" +
			fmt.Sprintf("A private_key or a private_key_path must be provided if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		"private_key_path": "(Optional) The path to the user's PEM formatted private key.\n" +
			fmt.Sprintf("A private_key or a private_key_path must be provided if auth is set to '%s', ignored otherwise.", authAPIKeySetting),
		"private_key_password": "(Optional) The password used to secure the private key.",
		"disable_auto_retries": "(Optional) Disable Automatic retries for retriable errors.\n" +
			"Auto retries were introduced to solve some eventual consistency problems but it also introduced performance issues on destroy operations.",
	}
}

// Provider is the adapter for terraform, that gives access to all the resources
func Provider(configfn schema.ConfigureFunc) terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: dataSourcesMap(),
		Schema:         schemaMap(),
		ResourcesMap:   resourcesMap(),
		ConfigureFunc:  configfn,
	}
}

func schemaMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auth": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  descriptions["auth"],
			DefaultFunc:  schema.EnvDefaultFunc("OCI_AUTH", authAPIKeySetting),
			ValidateFunc: validation.StringInSlice([]string{authAPIKeySetting, authInstancePrincipalSetting}, true),
		},
		"tenancy_ocid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions["tenancy_ocid"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_TENANCY_OCID", nil),
		},
		"user_ocid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions["user_ocid"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_USER_OCID", nil),
		},
		"fingerprint": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions["fingerprint"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_FINGERPRINT", nil),
		},
		// Mostly used for testing. Don't put keys in your .tf files
		"private_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: descriptions["private_key"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_PRIVATE_KEY", nil),
		},
		"private_key_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions["private_key_path"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_PRIVATE_KEY_PATH", nil),
		},
		"private_key_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Default:     "",
			Description: descriptions["private_key_password"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_PRIVATE_KEY_PASSWORD", nil),
		},
		"region": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["region"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_REGION", nil),
		},
		"disable_auto_retries": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: descriptions["disable_auto_retries"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_DISABLE_AUTO_RETRIES", nil),
		},
	}
}

func dataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"oci_core_boot_volume_attachments":             BootVolumeAttachmentsDataSource(),
		"oci_core_boot_volumes":                        BootVolumesDataSource(),
		"oci_core_console_history_data":                ConsoleHistoryContentDataSource(),
		"oci_core_cpes":                                CpesDataSource(),
		"oci_core_dhcp_options":                        DhcpOptionsDataSource(),
		"oci_core_drg_attachments":                     DrgAttachmentsDataSource(),
		"oci_core_drgs":                                DrgsDataSource(),
		"oci_core_images":                              ImagesDataSource(),
		"oci_core_instance_credentials":                InstanceCredentialDataSource(),
		"oci_core_instances":                           InstancesDataSource(),
		"oci_core_instance_console_connections":        InstanceConsoleConnectionsDataSource(),
		"oci_core_internet_gateways":                   InternetGatewaysDataSource(),
		"oci_core_ipsec_config":                        IpSecConnectionDeviceConfigDataSource(),
		"oci_core_ipsec_connections":                   IpSecConnectionsDataSource(),
		"oci_core_ipsec_status":                        IpSecConnectionDeviceStatusDataSource(),
		"oci_core_local_peering_gateways":              LocalPeeringGatewaysDataSource(),
		"oci_core_peer_region_for_remote_peerings":     PeerRegionForRemotePeeringsDataSource(),
		"oci_core_private_ips":                         PrivateIpsDataSource(),
		"oci_core_public_ip":                           PublicIpDataSource(),
		"oci_core_public_ips":                          PublicIpsDataSource(),
		"oci_core_remote_peering_connections":          RemotePeeringConnectionsDataSource(),
		"oci_core_route_tables":                        RouteTablesDataSource(),
		"oci_core_security_lists":                      SecurityListsDataSource(),
		"oci_core_shape":                               InstanceShapesDataSource(),
		"oci_core_shapes":                              InstanceShapesDataSource(),
		"oci_core_subnets":                             SubnetsDataSource(),
		"oci_core_virtual_networks":                    VcnsDataSource(), //This is a legacy name for VCN, removing it can cause breaking changes
		"oci_core_vcns":                                VcnsDataSource(),
		"oci_core_vnic":                                VnicDataSource(),
		"oci_core_vnic_attachments":                    VnicAttachmentsDataSource(),
		"oci_core_volume_attachments":                  VolumeAttachmentsDataSource(),
		"oci_core_volume_backup_policies":              VolumeBackupPoliciesDataSource(),
		"oci_core_volume_backup_policy_assignments":    VolumeBackupPolicyAssignmentsDataSource(),
		"oci_core_volume_backups":                      VolumeBackupsDataSource(),
		"oci_core_volumes":                             VolumesDataSource(),
		"oci_core_volume_groups":                       VolumeGroupsDataSource(),
		"oci_core_volume_group_backups":                VolumeGroupBackupsDataSource(),
		"oci_database_backups":                         BackupsDataSource(),
		"oci_database_database":                        DatabaseDataSource(),
		"oci_database_databases":                       DatabasesDataSource(),
		"oci_database_db_home":                         DbHomeDataSource(),
		"oci_database_db_homes":                        DbHomesDataSource(),
		"oci_database_db_node":                         DbNodeDataSource(),
		"oci_database_db_nodes":                        DbNodesDataSource(),
		"oci_database_db_system_shapes":                DbSystemShapesDataSource(),
		"oci_database_db_systems":                      DbSystemsDataSource(),
		"oci_database_db_system_patches":               DbSystemPatchesDataSource(),
		"oci_database_db_system_patch_history_entries": DbSystemPatchHistoryEntriesDataSource(),
		"oci_database_db_versions":                     DbVersionsDataSource(),
		"oci_database_db_home_patches":                 DbHomePatchesDataSource(),
		"oci_database_db_home_patch_history_entries":   DbHomePatchHistoryEntriesDataSource(),
		"oci_dns_records":                              RecordsDataSource(),
		"oci_dns_zones":                                ZonesDataSource(),
		"oci_file_storage_exports":                     ExportsDataSource(),
		"oci_file_storage_export_sets":                 ExportSetsDataSource(),
		"oci_file_storage_file_systems":                FileSystemsDataSource(),
		"oci_file_storage_mount_targets":               MountTargetsDataSource(),
		"oci_file_storage_snapshots":                   SnapshotsDataSource(),
		"oci_identity_api_keys":                        ApiKeysDataSource(),
		"oci_identity_auth_tokens":                     AuthTokensDataSource(),
		"oci_identity_availability_domains":            AvailabilityDomainsDataSource(),
		"oci_identity_compartments":                    CompartmentsDataSource(),
		"oci_identity_customer_secret_keys":            CustomerSecretKeysDataSource(),
		"oci_identity_dynamic_groups":                  DynamicGroupsDataSource(),
		"oci_identity_groups":                          GroupsDataSource(),
		"oci_identity_policies":                        IdentityPoliciesDataSource(),
		"oci_identity_regions":                         RegionsDataSource(),
		"oci_identity_swift_passwords":                 SwiftPasswordsDataSource(),
		"oci_identity_tag_namespaces":                  TagNamespacesDataSource(),
		"oci_identity_tags":                            TagsDataSource(),
		"oci_identity_tenancy":                         TenancyDataSource(),
		"oci_identity_user_group_memberships":          UserGroupMembershipsDataSource(),
		"oci_identity_users":                           UsersDataSource(),
		"oci_identity_region_subscriptions":            RegionSubscriptionsDataSource(),
		"oci_load_balancer_backend_health":             BackendHealthDataSource(),
		"oci_load_balancer_backends":                   BackendsDataSource(),
		"oci_load_balancer_backend_set_health":         BackendSetHealthDataSource(),
		"oci_load_balancer_backend_sets":               BackendSetsDataSource(),
		"oci_load_balancer_backendsets":                BackendSetsDataSource(),
		"oci_load_balancer_certificates":               CertificatesDataSource(),
		"oci_load_balancer_health":                     LoadBalancerHealthDataSource(),
		"oci_load_balancer_hostnames":                  HostnamesDataSource(),
		"oci_load_balancer_policies":                   LoadBalancerPoliciesDataSource(),
		"oci_load_balancer_protocols":                  LoadBalancerProtocolsDataSource(),
		"oci_load_balancer_shapes":                     LoadBalancerShapesDataSource(),
		"oci_load_balancer_load_balancers":             LoadBalancersDataSource(),
		"oci_load_balancers":                           LoadBalancersDataSource(),
		"oci_load_balancer_path_route_sets":            PathRouteSetsDataSource(),
		"oci_objectstorage_bucket_summaries":           BucketsDataSource(),
		"oci_objectstorage_namespace":                  NamespaceDataSource(),
		"oci_objectstorage_namespace_metadata":         NamespaceMetadataDataSource(),
		"oci_objectstorage_object_head":                ObjectHeadDataSource(),
		"oci_objectstorage_objects":                    ObjectsDataSource(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"oci_core_console_history":                 ConsoleHistoryResource(),
		"oci_core_cpe":                             CpeResource(),
		"oci_core_default_dhcp_options":            DefaultDhcpOptionsResource(),
		"oci_core_dhcp_options":                    DhcpOptionsResource(),
		"oci_core_drg":                             DrgResource(),
		"oci_core_drg_attachment":                  DrgAttachmentResource(),
		"oci_core_image":                           ImageResource(),
		"oci_core_instance":                        InstanceResource(),
		"oci_core_instance_console_connection":     InstanceConsoleConnectionResource(),
		"oci_core_internet_gateway":                InternetGatewayResource(),
		"oci_core_ipsec":                           IpSecConnectionResource(),
		"oci_core_local_peering_gateway":           LocalPeeringGatewayResource(),
		"oci_core_private_ip":                      PrivateIpResource(),
		"oci_core_public_ip":                       PublicIpResource(),
		"oci_core_default_route_table":             DefaultRouteTableResource(),
		"oci_core_route_table":                     RouteTableResource(),
		"oci_core_remote_peering_connection":       RemotePeeringConnectionResource(),
		"oci_core_default_security_list":           DefaultSecurityListResource(),
		"oci_core_security_list":                   SecurityListResource(),
		"oci_core_subnet":                          SubnetResource(),
		"oci_core_virtual_network":                 VcnResource(), //This is a legacy name for VCN, removing it can cause breaking changes
		"oci_core_vcn":                             VcnResource(),
		"oci_core_vnic_attachment":                 VnicAttachmentResource(),
		"oci_core_volume":                          VolumeResource(),
		"oci_core_volume_group":                    VolumeGroupResource(),
		"oci_core_volume_group_backup":             VolumeGroupBackupResource(),
		"oci_core_volume_attachment":               VolumeAttachmentResource(),
		"oci_core_volume_backup":                   VolumeBackupResource(),
		"oci_core_volume_backup_policy_assignment": VolumeBackupPolicyAssignmentResource(),
		//"oci_database_db_home":                     DbHomeResource(),
		"oci_database_db_system":               DbSystemResource(),
		"oci_database_backup":                  BackupResource(),
		"oci_dns_record":                       RecordResource(),
		"oci_dns_zone":                         ZoneResource(),
		"oci_file_storage_export":              ExportResource(),
		"oci_file_storage_export_set":          ExportSetResource(),
		"oci_file_storage_file_system":         FileSystemResource(),
		"oci_file_storage_mount_target":        MountTargetResource(),
		"oci_file_storage_snapshot":            SnapshotResource(),
		"oci_identity_api_key":                 ApiKeyResource(),
		"oci_identity_auth_token":              AuthTokenResource(),
		"oci_identity_compartment":             CompartmentResource(),
		"oci_identity_customer_secret_key":     CustomerSecretKeyResource(),
		"oci_identity_dynamic_group":           DynamicGroupResource(),
		"oci_identity_group":                   GroupResource(),
		"oci_identity_policy":                  PolicyResource(),
		"oci_identity_swift_password":          SwiftPasswordResource(),
		"oci_identity_tag_namespace":           TagNamespaceResource(),
		"oci_identity_tag":                     TagResource(),
		"oci_identity_ui_password":             UiPasswordResource(),
		"oci_identity_user":                    UserResource(),
		"oci_identity_user_group_membership":   UserGroupMembershipResource(),
		"oci_load_balancer":                    LoadBalancerResource(),
		"oci_load_balancer_load_balancer":      LoadBalancerResource(),
		"oci_load_balancer_backend":            BackendResource(),
		"oci_load_balancer_backend_set":        BackendSetResource(),
		"oci_load_balancer_backendset":         BackendSetResource(),
		"oci_load_balancer_certificate":        CertificateResource(),
		"oci_load_balancer_listener":           ListenerResource(),
		"oci_load_balancer_hostname":           HostnameResource(),
		"oci_load_balancer_path_route_set":     PathRouteSetResource(),
		"oci_objectstorage_bucket":             BucketResource(),
		"oci_objectstorage_object":             ObjectResource(),
		"oci_objectstorage_namespace_metadata": NamespaceMetadataResource(),
		"oci_objectstorage_preauthrequest":     PreauthenticatedRequestResource(),
	}
}

func getEnvSetting(s string, dv string) string {
	v := os.Getenv("TF_VAR_" + s)
	if v != "" {
		return v
	}
	v = os.Getenv("OCI_" + s)
	if v != "" {
		return v
	}
	v = os.Getenv(s)
	if v != "" {
		return v
	}
	return dv
}

func getRequiredEnvSetting(s string) string {
	v := getEnvSetting(s, "")
	if v == "" {
		panic(fmt.Sprintf("Required env setting %s is missing", s))
	}
	return v
}

func validateConfigForAPIKeyAuth(d *schema.ResourceData) error {
	_, hasTenancyOCID := d.GetOkExists("tenancy_ocid")
	_, hasUserOCID := d.GetOkExists("user_ocid")
	_, hasFingerprint := d.GetOkExists("fingerprint")
	if !hasTenancyOCID || !hasUserOCID || !hasFingerprint {
		return fmt.Errorf("when auth is set to '%s', tenancy_ocid, user_ocid, and fingerprint are required", authAPIKeySetting)
	}
	return nil
}

func ProviderConfig(d *schema.ResourceData) (clients interface{}, err error) {
	clients = &OracleClients{}
	disableAutoRetries = d.Get("disable_auto_retries").(bool)
	auth := strings.ToLower(d.Get("auth").(string))

	userAgent := fmt.Sprintf(userAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, terraform.VersionString(), Version)

	httpClient := &http.Client{
		Timeout: defaultRequestTimeout,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: defaultConnectionTimeout,
			}).Dial,
			TLSHandshakeTimeout: defaultTLSHandshakeTimeout,
			TLSClientConfig:     &tls.Config{MinVersion: tls.VersionTLS12},
			Proxy:               http.ProxyFromEnvironment,
		},
	}

	var configProviders []oci_common.ConfigurationProvider

	switch auth {
	case strings.ToLower(authAPIKeySetting):
		if err := validateConfigForAPIKeyAuth(d); err != nil {
			return nil, err
		}
	case strings.ToLower(authInstancePrincipalSetting):
		region, ok := d.GetOkExists("region")
		if !ok {
			return nil, fmt.Errorf("can not get region from Terraform configuration (InstancePrincipal)")
		}
		cfg, err := oci_common_auth.InstancePrincipalConfigurationProviderForRegion(oci_common.StringToRegion(region.(string)))
		if err != nil {
			return nil, err
		}
		configProviders = append(configProviders, cfg)
	default:
		return nil, fmt.Errorf("auth must be one of '%s' or '%s'", authAPIKeySetting, authInstancePrincipalSetting)
	}

	configProviders = append(configProviders, ResourceDataConfigProvider{d})

	// TODO: DefaultConfigProvider will return us a composingConfigurationProvider that reads from SDK config files,
	// and then from the environment variables ("TF_VAR" prefix). References to "TF_VAR" prefix should be removed from
	// the SDK, since it's Terraform specific. When that happens, we need to update this to pass in the right prefix.
	configProviders = append(configProviders, oci_common.DefaultConfigProvider())

	officialSdkConfigProvider, err := oci_common.ComposingConfigurationProvider(configProviders)
	if err != nil {
		return nil, err
	}

	err = setGoSDKClients(clients.(*OracleClients), officialSdkConfigProvider, httpClient, userAgent)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func setGoSDKClients(clients *OracleClients, officialSdkConfigProvider oci_common.ConfigurationProvider, httpClient *http.Client, userAgent string) (err error) {
	// Official Go SDK clients:
	blockstorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	computeClient, err := oci_core.NewComputeClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	databaseClient, err := oci_database.NewDatabaseClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	dnsClient, err := oci_dns.NewDnsClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	fileStorageClient, err := oci_file_storage.NewFileStorageClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	loadBalancerClient, err := oci_load_balancer.NewLoadBalancerClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	objectStorageClient, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	virtualNetworkClient, err := oci_core.NewVirtualNetworkClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	useOboToken, err := strconv.ParseBool(getEnvSetting("use_obo_token", "false"))
	if err != nil {
		return
	}

	var oboTokenProvider oci_common.OboTokenProvider
	if useOboToken {
		oboTokenProvider = oboTokenProviderFromEnv{}
	} else {
		oboTokenProvider = oci_common.NewEmptyOboTokenProvider()
	}

	configureClient := func(client *oci_common.BaseClient) error {
		client.HTTPClient = httpClient
		client.UserAgent = userAgent
		client.Obo = oboTokenProvider

		// R1 Support
		if region, err := officialSdkConfigProvider.Region(); err == nil && strings.ToLower(region) == "r1" {
			service := strings.Split(client.Host, ".")[0]
			client.Host = fmt.Sprintf("%s.r1.oracleiaas.com", service)

			pool := x509.NewCertPool()
			//readCertPem reads the pem files to a []byte
			cert, err := readCertPem()
			if err != nil {
				return err
			}
			if ok := pool.AppendCertsFromPEM(cert); !ok {
				return fmt.Errorf("failed to append R1 cert to the cert pool")
			}
			//install the certificates to the client
			if h, ok := client.HTTPClient.(*http.Client); ok {
				tr := &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}
				h.Transport = tr
			} else {
				return fmt.Errorf("the client dispatcher is not of http.Client type. can not patch the tls config")
			}
		}
		return nil
	}

	err = configureClient(&blockstorageClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&computeClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&databaseClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&fileStorageClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&identityClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&loadBalancerClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&objectStorageClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&virtualNetworkClient.BaseClient)
	if err != nil {
		return
	}

	clients.blockstorageClient = &blockstorageClient
	clients.computeClient = &computeClient
	clients.databaseClient = &databaseClient
	clients.dnsClient = &dnsClient
	clients.fileStorageClient = &fileStorageClient
	clients.identityClient = &identityClient
	clients.loadBalancerClient = &loadBalancerClient
	clients.objectStorageClient = &objectStorageClient
	clients.virtualNetworkClient = &virtualNetworkClient

	return
}

type OracleClients struct {
	blockstorageClient   *oci_core.BlockstorageClient
	computeClient        *oci_core.ComputeClient
	databaseClient       *oci_database.DatabaseClient
	dnsClient            *oci_dns.DnsClient
	identityClient       *oci_identity.IdentityClient
	virtualNetworkClient *oci_core.VirtualNetworkClient
	objectStorageClient  *oci_object_storage.ObjectStorageClient
	loadBalancerClient   *oci_load_balancer.LoadBalancerClient
	fileStorageClient    *oci_file_storage.FileStorageClient
}

type ResourceDataConfigProvider struct {
	D *schema.ResourceData
}

// TODO: The error messages returned by following methods get swallowed up by the ComposingConfigurationProvider,
// since it only checks whether an error exists or not.
// The ComposingConfigurationProvider in SDK should log the errors as debug statements instead.

func (p ResourceDataConfigProvider) TenancyOCID() (string, error) {
	if tenancyOCID, ok := p.D.GetOkExists("tenancy_ocid"); ok {
		return tenancyOCID.(string), nil
	}
	return "", fmt.Errorf("can not get tenancy_ocid from Terraform configuration")
}

func (p ResourceDataConfigProvider) UserOCID() (string, error) {
	if userOCID, ok := p.D.GetOkExists("user_ocid"); ok {
		return userOCID.(string), nil
	}
	return "", fmt.Errorf("can not get user_ocid from Terraform configuration")
}

func (p ResourceDataConfigProvider) KeyFingerprint() (string, error) {
	if fingerprint, ok := p.D.GetOkExists("fingerprint"); ok {
		return fingerprint.(string), nil
	}
	return "", fmt.Errorf("can not get fingerprint from Terraform configuration")
}

func (p ResourceDataConfigProvider) Region() (string, error) {
	if region, ok := p.D.GetOkExists("region"); ok {
		return region.(string), nil
	}
	return "", fmt.Errorf("can not get region from Terraform configuration")
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
	if privateKeyPassword, hasPrivateKeyPassword := p.D.GetOkExists("private_key_password"); hasPrivateKeyPassword {
		password = privateKeyPassword.(string)
	}

	if privateKey, hasPrivateKey := p.D.GetOkExists("private_key"); hasPrivateKey {
		return oci_common.PrivateKeyFromBytes([]byte(privateKey.(string)), &password)
	}

	if privateKeyPath, hasPrivateKeyPath := p.D.GetOkExists("private_key_path"); hasPrivateKeyPath {
		pemFileContent, readFileErr := ioutil.ReadFile(privateKeyPath.(string))
		if readFileErr != nil {
			return nil, fmt.Errorf("Can not read private key from: '%s', Error: %q", privateKeyPath, readFileErr)
		}
		return oci_common.PrivateKeyFromBytes(pemFileContent, &password)
	}

	return nil, fmt.Errorf("can not get private_key or private_key_path from Terraform configuration")
}

func readCertPem() (file []byte, err error) {
	r1CertLoc := getEnvSetting(r1CertLocationEnv, "")
	if r1CertLoc == "" {
		err = fmt.Errorf("the R1 Certificate Location must be specified in the environment variable %s", r1CertLocationEnv)
		return
	}
	file, err = ioutil.ReadFile(r1CertLoc)
	return
}
