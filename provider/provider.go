// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"crypto/rsa"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_database "github.com/oracle/oci-go-sdk/database"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

var descriptions map[string]string
var disableAutoRetries bool

func init() {
	descriptions = map[string]string{
		"tenancy_ocid": "(Required) The tenancy OCID for a user. The tenancy OCID can be found at the bottom of user settings in the Oracle Cloud Infrastructure console.",
		"user_ocid":    "(Required) The user OCID. This can be found in user settings in the Oracle Cloud Infrastructure console.",
		"fingerprint":  "(Required) The fingerprint for the user's RSA key. This can be found in user settings in the Oracle Cloud Infrastructure console.",
		"region":       "(Required) The region for API connections (e.g. us-ashburn-1).",
		"private_key": "(Optional) A PEM formatted RSA private key for the user.\n" +
			"A private_key or a private_key_path must be provided.",
		"private_key_path": "(Optional) The path to the user's PEM formatted private key.\n" +
			"A private_key or a private_key_path must be provided.",
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
		"tenancy_ocid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["tenancy_ocid"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_TENANCY_OCID", nil),
		},
		"user_ocid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["user_ocid"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_USER_OCID", nil),
		},
		"fingerprint": {
			Type:        schema.TypeString,
			Required:    true,
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
		"oci_core_console_history_data":       ConsoleHistoryContentDataSource(),
		"oci_core_cpes":                       CpesDataSource(),
		"oci_core_dhcp_options":               DhcpOptionsDataSource(),
		"oci_core_drg_attachments":            DrgAttachmentsDataSource(),
		"oci_core_drgs":                       DrgsDataSource(),
		"oci_core_images":                     ImagesDataSource(),
		"oci_core_instance_credentials":       InstanceCredentialsDataSource(),
		"oci_core_instances":                  InstancesDataSource(),
		"oci_core_internet_gateways":          InternetGatewaysDataSource(),
		"oci_core_ipsec_config":               IpSecConnectionDeviceConfigsDataSource(),
		"oci_core_ipsec_connections":          IpSecConnectionsDataSource(),
		"oci_core_ipsec_status":               IpSecConnectionDeviceStatusDataSource(),
		"oci_core_private_ips":                PrivateIpsDataSource(),
		"oci_core_route_tables":               RouteTablesDataSource(),
		"oci_core_security_lists":             SecurityListsDataSource(),
		"oci_core_shape":                      InstanceShapesDataSource(),
		"oci_core_subnets":                    SubnetsDataSource(),
		"oci_core_virtual_networks":           VcnsDataSource(), //This is a legacy name for VCN, removing it can cause breaking changes
		"oci_core_vcns":                       VcnsDataSource(),
		"oci_core_vnic":                       VnicsDataSource(),
		"oci_core_vnic_attachments":           VnicAttachmentsDataSource(),
		"oci_core_volume_attachments":         VolumeAttachmentsDataSource(),
		"oci_core_volume_backups":             VolumeBackupsDataSource(),
		"oci_core_volumes":                    VolumesDataSource(),
		"oci_database_database":               DatabaseDataSource(),
		"oci_database_databases":              DatabasesDataSource(),
		"oci_database_db_home":                DbHomeDataSource(),
		"oci_database_db_homes":               DbHomesDataSource(),
		"oci_database_db_node":                DbNodeDataSource(),
		"oci_database_db_nodes":               DbNodesDataSource(),
		"oci_database_db_system_shapes":       DbSystemShapesDataSource(),
		"oci_database_db_systems":             DbSystemsDataSource(),
		"oci_database_db_versions":            DbVersionsDataSource(),
		"oci_identity_api_keys":               ApiKeysDataSource(),
		"oci_identity_availability_domains":   AvailabilityDomainsDataSource(),
		"oci_identity_compartments":           CompartmentsDataSource(),
		"oci_identity_groups":                 GroupsDataSource(),
		"oci_identity_policies":               IdentityPoliciesDataSource(),
		"oci_identity_swift_passwords":        SwiftPasswordsDataSource(),
		"oci_identity_user_group_memberships": UserGroupMembershipsDataSource(),
		"oci_identity_users":                  UsersDataSource(),
		"oci_load_balancer_backends":          BackendsDataSource(),
		"oci_load_balancer_backendsets":       BackendSetsDataSource(),
		"oci_load_balancer_certificates":      CertificatesDataSource(),
		"oci_load_balancer_policies":          LoadBalancerPoliciesDataSource(),
		"oci_load_balancer_protocols":         LoadBalancerProtocolsDataSource(),
		"oci_load_balancer_shapes":            LoadBalancerShapesDataSource(),
		"oci_load_balancer_load_balancers":    LoadBalancersDataSource(),
		"oci_load_balancers":                  LoadBalancersDataSource(),
		"oci_objectstorage_bucket_summaries":  BucketsDataSource(),
		"oci_objectstorage_namespace":         NamespacesDataSource(),
		"oci_objectstorage_object_head":       ObjectHeadDataSource(),
		"oci_objectstorage_objects":           ObjectsDataSource(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"oci_core_console_history":           ConsoleHistoryResource(),
		"oci_core_cpe":                       CpeResource(),
		"oci_core_default_dhcp_options":      DefaultDhcpOptionsResource(),
		"oci_core_dhcp_options":              DhcpOptionsResource(),
		"oci_core_drg":                       DrgResource(),
		"oci_core_drg_attachment":            DrgAttachmentResource(),
		"oci_core_image":                     ImageResource(),
		"oci_core_instance":                  InstanceResource(),
		"oci_core_internet_gateway":          InternetGatewayResource(),
		"oci_core_ipsec":                     IpSecConnectionResource(),
		"oci_core_private_ip":                PrivateIpResource(),
		"oci_core_default_route_table":       DefaultRouteTableResource(),
		"oci_core_route_table":               RouteTableResource(),
		"oci_core_default_security_list":     DefaultSecurityListResource(),
		"oci_core_security_list":             SecurityListResource(),
		"oci_core_subnet":                    SubnetResource(),
		"oci_core_virtual_network":           VcnResource(), //This is a legacy name for VCN, removing it can cause breaking changes
		"oci_core_vcn":                       VcnResource(),
		"oci_core_vnic_attachment":           VnicAttachmentResource(),
		"oci_core_volume":                    VolumeResource(),
		"oci_core_volume_attachment":         VolumeAttachmentResource(),
		"oci_core_volume_backup":             VolumeBackupResource(),
		"oci_database_db_system":             DbSystemResource(),
		"oci_identity_api_key":               ApiKeyResource(),
		"oci_identity_compartment":           CompartmentResource(),
		"oci_identity_group":                 GroupResource(),
		"oci_identity_policy":                PolicyResource(),
		"oci_identity_swift_password":        SwiftPasswordResource(),
		"oci_identity_ui_password":           UiPasswordResource(),
		"oci_identity_user":                  UserResource(),
		"oci_identity_user_group_membership": UserGroupMembershipResource(),
		"oci_load_balancer":                  LoadBalancerResource(),
		"oci_load_balancer_load_balancer":    LoadBalancerResource(),
		"oci_load_balancer_backend":          BackendResource(),
		"oci_load_balancer_backendset":       BackendSetResource(),
		"oci_load_balancer_certificate":      CertificateResource(),
		"oci_load_balancer_listener":         ListenerResource(),
		"oci_objectstorage_bucket":           BucketResource(),
		"oci_objectstorage_object":           ObjectResource(),
		"oci_objectstorage_preauthrequest":   PreauthenticatedRequestResource(),
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

func ProviderConfig(d *schema.ResourceData) (clients interface{}, err error) {
	tenancyOCID := d.Get("tenancy_ocid").(string)
	userOCID := d.Get("user_ocid").(string)
	fingerprint := d.Get("fingerprint").(string)
	privateKeyBuffer, hasKey := d.Get("private_key").(string)
	privateKeyPath, hasKeyPath := d.Get("private_key_path").(string)
	privateKeyPassword, hasKeyPass := d.Get("private_key_password").(string)
	region, hasRegion := d.Get("region").(string)
	disableAutoRetriesLocal, hasDisableRetries := d.Get("disable_auto_retries").(bool)
	disableAutoRetries = disableAutoRetriesLocal

	// for internal use
	urlTemplate := getEnvSetting("url_template", "")
	allowInsecureTls := getEnvSetting("allow_insecure_tls", "")

	clientOpts := []baremetal.NewClientOptionsFunc{
		func(o *baremetal.NewClientOptions) {
			o.UserAgent = fmt.Sprintf("Oracle-GoSDK/%s (go/%s; %s/%s; terraform/%s) Oracle-TerraformProvider/%s",
				baremetal.SDKVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH, terraform.VersionString(), Version)
		},
	}

	if allowInsecureTls == "true" {
		log.Println("[WARN] USING INSECURE TLS")
		clientOpts = append(clientOpts, baremetal.CustomTransport(
			&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		))
	} else {
		clientOpts = append(clientOpts, baremetal.CustomTransport(
			&http.Transport{Proxy: http.ProxyFromEnvironment}),
		)
	}

	clientOpts = append(clientOpts, baremetal.CustomTransport(
		&http.Transport{TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS12}},
	))

	if hasKey && privateKeyBuffer != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyBytes([]byte(privateKeyBuffer)))
	} else if hasKeyPath && privateKeyPath != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyFilePath(privateKeyPath))
	} else {
		return nil, errors.New("One of private_key or private_key_path is required")
	}

	if hasKeyPass && privateKeyPassword != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyPassword(privateKeyPassword))
	}

	if hasRegion && region != "" {
		clientOpts = append(clientOpts, baremetal.Region(region))
	}

	if hasDisableRetries {
		clientOpts = append(clientOpts, baremetal.DisableAutoRetries(disableAutoRetries))
	}

	if urlTemplate != "" {
		clientOpts = append(clientOpts, baremetal.UrlTemplate(urlTemplate))
	}

	client, err := baremetal.NewClient(userOCID, tenancyOCID, fingerprint, clientOpts...)

	clientOpts = append(clientOpts, baremetal.DisableNotFoundRetries(true))
	clientWithoutNotFoundRetries, err := baremetal.NewClient(userOCID, tenancyOCID, fingerprint, clientOpts...)

	clients = &OracleClients{
		client: client,
		clientWithoutNotFoundRetries: clientWithoutNotFoundRetries,
	}

	tfConfigProvider := ResourceDataConfigProvider{d}

	// TODO: DefaultConfigProvider will return us a composingConfigurationProvider that reads from SDK config files,
	// and then from the environment variables ("TF_VAR" prefix). References to "TF_VAR" prefix should be removed from
	// the SDK, since it's Terraform specific. When that happens, we need to update this to pass in the right prefix.
	defaultConfigProvider := oci_common.DefaultConfigProvider()

	officialSdkConfigProvider, err := oci_common.ComposingConfigurationProvider([]oci_common.ConfigurationProvider{tfConfigProvider, defaultConfigProvider})
	if err != nil {
		return nil, err
	}

	err = setGoSDKClients(clients.(*OracleClients), officialSdkConfigProvider)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func setGoSDKClients(clients *OracleClients, officialSdkConfigProvider oci_common.ConfigurationProvider) (err error) {
	// Official Go SDK clients:
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(officialSdkConfigProvider)
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

	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	virtualNetworkClient, err := oci_core.NewVirtualNetworkClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	objectStorageClient, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	loadBalancerClient, err := oci_load_balancer.NewLoadBalancerClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	clients.blockStorageClient = &blockStorageClient
	clients.computeClient = &computeClient
	clients.databaseClient = &databaseClient
	clients.identityClient = &identityClient
	clients.virtualNetworkClient = &virtualNetworkClient
	clients.objectStorageClient = &objectStorageClient
	clients.loadBalancerClient = &loadBalancerClient

	return
}

type OracleClients struct {
	client                       *baremetal.Client
	clientWithoutNotFoundRetries *baremetal.Client

	blockStorageClient   *oci_core.BlockstorageClient
	computeClient        *oci_core.ComputeClient
	databaseClient       *oci_database.DatabaseClient
	identityClient       *oci_identity.IdentityClient
	virtualNetworkClient *oci_core.VirtualNetworkClient
	objectStorageClient  *oci_object_storage.ObjectStorageClient
	loadBalancerClient   *oci_load_balancer.LoadBalancerClient
}

type ResourceDataConfigProvider struct {
	D *schema.ResourceData
}

// TODO: The error messages returned by following methods get swallowed up by the ComposingConfigurationProvider,
// since it only checks whether an error exists or not.
// The ComposingConfigurationProvider in SDK should log the errors as debug statements instead.

func (p ResourceDataConfigProvider) TenancyOCID() (string, error) {
	if tenancyOCID, ok := p.D.GetOk("tenancy_ocid"); ok {
		return tenancyOCID.(string), nil
	}
	return "", fmt.Errorf("Can not get tenancy_ocid from Terraform configuration")
}

func (p ResourceDataConfigProvider) UserOCID() (string, error) {
	if userOCID, ok := p.D.GetOk("user_ocid"); ok {
		return userOCID.(string), nil
	}
	return "", fmt.Errorf("Can not get user_ocid from Terraform configuration")
}

func (p ResourceDataConfigProvider) KeyFingerprint() (string, error) {
	if fingerprint, ok := p.D.GetOk("fingerprint"); ok {
		return fingerprint.(string), nil
	}
	return "", fmt.Errorf("Can not get fingerprint from Terraform configuration")
}

func (p ResourceDataConfigProvider) Region() (string, error) {
	if region, ok := p.D.GetOk("region"); ok {
		return region.(string), nil
	}
	return "", fmt.Errorf("Can not get region from Terraform configuration")
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
	if privateKeyPassword, hasPrivateKeyPassword := p.D.GetOk("private_key_password"); hasPrivateKeyPassword {
		password = privateKeyPassword.(string)
	}

	if privateKey, hasPrivateKey := p.D.GetOk("private_key"); hasPrivateKey {
		return oci_common.PrivateKeyFromBytes([]byte(privateKey.(string)), &password)
	}

	if privateKeyPath, hasPrivateKeyPath := p.D.GetOk("private_key_path"); hasPrivateKeyPath {
		pemFileContent, readFileErr := ioutil.ReadFile(privateKeyPath.(string))
		if readFileErr != nil {
			return nil, fmt.Errorf("Can not read private key from: '%s', Error: %q", privateKeyPath, readFileErr)
		}
		return oci_common.PrivateKeyFromBytes(pemFileContent, &password)
	}

	return nil, fmt.Errorf("Can not get private_key or private_key_path from Terraform configuration")
}
