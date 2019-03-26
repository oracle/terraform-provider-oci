// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"crypto/rsa"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/hashicorp/terraform/terraform"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/common/auth"
)

var descriptions map[string]string
var apiKeyConfigAttributes = [...]string{userOcidAttrName, fingerprintAttrName, privateKeyAttrName, privateKeyPathAttrName, privateKeyPasswordAttrName}
var avoidWaitingForDeleteTarget bool

const (
	authAPIKeySetting                     = "ApiKey"
	authInstancePrincipalSetting          = "InstancePrincipal"
	authInstancePrincipalWithCertsSetting = "InstancePrincipalWithCerts"
	requestHeaderOpcOboToken              = "opc-obo-token"
	requestHeaderOpcHostSerial            = "opc-host-serial"
	defaultRequestTimeout                 = 0
	defaultConnectionTimeout              = 10 * time.Second
	defaultTLSHandshakeTimeout            = 5 * time.Second
	defaultUserAgentProviderName          = "Oracle-TerraformProvider"
	userAgentFormatter                    = "Oracle-GoSDK/%s (go/%s; %s/%s; terraform/%s) %s/%s"
	userAgentProviderNameEnv              = "USER_AGENT_PROVIDER_NAME"
	domainNameOverrideEnv                 = "domain_name_override"
	customCertLocationEnv                 = "custom_cert_location"
	oracleR1DomainNameEnv                 = "oracle_r1_domain_name" // deprecate
	r1CertLocationEnv                     = "R1_CERT_LOCATION"      // deprecate

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

	tfEnvPrefix  = "TF_VAR_"
	ociEnvPrefix = "OCI_"
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
		authAttrName:        fmt.Sprintf("(Optional) The type of auth to use. Options are '%s' and '%s'. By default, '%s' will be used.", authAPIKeySetting, authInstancePrincipalSetting, authAPIKeySetting),
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
		authAttrName: {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  descriptions[authAttrName],
			DefaultFunc:  schema.MultiEnvDefaultFunc([]string{tfVarName(authAttrName), ociVarName(authAttrName)}, authAPIKeySetting),
			ValidateFunc: validation.StringInSlice([]string{authAPIKeySetting, authInstancePrincipalSetting, authInstancePrincipalWithCertsSetting}, true),
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
			Required:    true,
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
	}
}

func dataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"oci_audit_configuration":                        AuditConfigurationDataSource(),
		"oci_audit_events":                               AuditAuditEventsDataSource(),
		"oci_budget_budget":                              BudgetBudgetDataSource(),
		"oci_budget_budgets":                             BudgetBudgetsDataSource(),
		"oci_budget_alert_rule":                          BudgetAlertRuleDataSource(),
		"oci_budget_alert_rules":                         BudgetAlertRulesDataSource(),
		"oci_autoscaling_auto_scaling_configuration":     AutoscalingAutoScalingConfigurationDataSource(),
		"oci_autoscaling_auto_scaling_configurations":    AutoscalingAutoScalingConfigurationsDataSource(),
		"oci_containerengine_clusters":                   ContainerengineClustersDataSource(),
		"oci_containerengine_cluster_option":             ContainerengineClusterOptionDataSource(),
		"oci_containerengine_node_pool":                  ContainerengineNodePoolDataSource(),
		"oci_containerengine_node_pools":                 ContainerengineNodePoolsDataSource(),
		"oci_containerengine_node_pool_option":           ContainerengineNodePoolOptionDataSource(),
		"oci_containerengine_cluster_kube_config":        ContainerengineClusterKubeConfigDataSource(),
		"oci_containerengine_work_requests":              ContainerengineWorkRequestsDataSource(),
		"oci_containerengine_work_request_errors":        ContainerengineWorkRequestErrorsDataSource(),
		"oci_containerengine_work_request_log_entries":   ContainerengineWorkRequestLogEntriesDataSource(),
		"oci_core_app_catalog_listing":                   CoreAppCatalogListingDataSource(),
		"oci_core_app_catalog_listings":                  CoreAppCatalogListingsDataSource(),
		"oci_core_app_catalog_listing_resource_versions": CoreAppCatalogListingResourceVersionsDataSource(),
		"oci_core_app_catalog_listing_resource_version":  CoreAppCatalogListingResourceVersionDataSource(),
		"oci_core_listing_resource_versions":             CoreAppCatalogListingResourceVersionsDataSource(),
		"oci_core_listing_resource_version":              CoreAppCatalogListingResourceVersionDataSource(),
		"oci_core_app_catalog_subscriptions":             CoreAppCatalogSubscriptionsDataSource(),
		"oci_core_boot_volume_attachments":               CoreBootVolumeAttachmentsDataSource(),
		"oci_core_boot_volume":                           CoreBootVolumeDataSource(),
		"oci_core_boot_volumes":                          CoreBootVolumesDataSource(),
		"oci_core_boot_volume_backup":                    CoreBootVolumeBackupDataSource(),
		"oci_core_boot_volume_backups":                   CoreBootVolumeBackupsDataSource(),
		"oci_core_console_histories":                     CoreConsoleHistoriesDataSource(),
		"oci_core_console_history_data":                  CoreConsoleHistoryContentDataSource(),
		"oci_core_cpes":                                  CoreCpesDataSource(),
		"oci_core_cross_connect_group":                   CoreCrossConnectGroupDataSource(),
		"oci_core_cross_connect_groups":                  CoreCrossConnectGroupsDataSource(),
		"oci_core_cross_connect_locations":               CoreCrossConnectLocationsDataSource(),
		"oci_core_cross_connect_port_speed_shapes":       CoreCrossConnectPortSpeedShapesDataSource(),
		"oci_core_cross_connect_status":                  CoreCrossConnectStatusDataSource(),
		"oci_core_cross_connect":                         CoreCrossConnectDataSource(),
		"oci_core_cross_connects":                        CoreCrossConnectsDataSource(),
		"oci_core_dhcp_options":                          CoreDhcpOptionsDataSource(),
		"oci_core_drg_attachments":                       CoreDrgAttachmentsDataSource(),
		"oci_core_drgs":                                  CoreDrgsDataSource(),
		"oci_core_fast_connect_provider_service":         CoreFastConnectProviderServiceDataSource(),
		"oci_core_fast_connect_provider_services":        CoreFastConnectProviderServicesDataSource(),
		"oci_core_images":                                CoreImagesDataSource(),
		"oci_core_instance":                              CoreInstanceDataSource(),
		"oci_core_instance_credentials":                  CoreInstanceCredentialDataSource(),
		"oci_core_instance_configuration":                CoreInstanceConfigurationDataSource(),
		"oci_core_instance_configurations":               CoreInstanceConfigurationsDataSource(),
		"oci_core_instance_pool":                         CoreInstancePoolDataSource(),
		"oci_core_instance_pools":                        CoreInstancePoolsDataSource(),
		"oci_core_instance_pool_instances":               CoreInstancePoolInstancesDataSource(),
		"oci_core_instance_devices":                      CoreInstanceDevicesDataSource(),
		"oci_core_instances":                             CoreInstancesDataSource(),
		"oci_core_instance_console_connections":          CoreInstanceConsoleConnectionsDataSource(),
		"oci_core_internet_gateways":                     CoreInternetGatewaysDataSource(),
		"oci_core_ipsec_config":                          CoreIpSecConnectionDeviceConfigDataSource(),
		"oci_core_ipsec_connections":                     CoreIpSecConnectionsDataSource(),
		"oci_core_ipsec_status":                          CoreIpSecConnectionDeviceStatusDataSource(),
		"oci_core_letter_of_authority":                   CoreLetterOfAuthorityDataSource(),
		"oci_core_local_peering_gateways":                CoreLocalPeeringGatewaysDataSource(),
		"oci_core_nat_gateway":                           CoreNatGatewayDataSource(),
		"oci_core_nat_gateways":                          CoreNatGatewaysDataSource(),
		"oci_core_peer_region_for_remote_peerings":       CorePeerRegionForRemotePeeringsDataSource(),
		"oci_core_private_ips":                           CorePrivateIpsDataSource(),
		"oci_core_public_ip":                             CorePublicIpDataSource(),
		"oci_core_public_ips":                            CorePublicIpsDataSource(),
		"oci_core_remote_peering_connections":            CoreRemotePeeringConnectionsDataSource(),
		"oci_core_route_tables":                          CoreRouteTablesDataSource(),
		"oci_core_security_lists":                        CoreSecurityListsDataSource(),
		"oci_core_service_gateways":                      CoreServiceGatewaysDataSource(),
		"oci_core_services":                              CoreServicesDataSource(),
		"oci_core_shape":                                 CoreShapesDataSource(),
		"oci_core_shapes":                                CoreShapesDataSource(),
		"oci_core_subnet":                                CoreSubnetDataSource(),
		"oci_core_subnets":                               CoreSubnetsDataSource(),
		"oci_core_virtual_circuit_bandwidth_shapes":      CoreVirtualCircuitBandwidthShapesDataSource(),
		"oci_core_virtual_circuit_public_prefixes":       CoreVirtualCircuitPublicPrefixesDataSource(),
		"oci_core_virtual_circuit":                       CoreVirtualCircuitDataSource(),
		"oci_core_virtual_circuits":                      CoreVirtualCircuitsDataSource(),
		"oci_core_virtual_networks":                      CoreVcnsDataSource(), //This is a legacy name for VCN, removing it can cause breaking changes
		"oci_core_vcns":                                  CoreVcnsDataSource(),
		"oci_core_vnic":                                  CoreVnicDataSource(),
		"oci_core_vnic_attachments":                      CoreVnicAttachmentsDataSource(),
		"oci_core_volume":                                CoreVolumeDataSource(),
		"oci_core_volume_attachments":                    CoreVolumeAttachmentsDataSource(),
		"oci_core_volume_backup_policies":                CoreVolumeBackupPoliciesDataSource(),
		"oci_core_volume_backup_policy_assignments":      CoreVolumeBackupPolicyAssignmentsDataSource(),
		"oci_core_volume_backups":                        CoreVolumeBackupsDataSource(),
		"oci_core_volumes":                               CoreVolumesDataSource(),
		"oci_core_volume_groups":                         CoreVolumeGroupsDataSource(),
		"oci_core_volume_group_backups":                  CoreVolumeGroupBackupsDataSource(),
		"oci_database_autonomous_data_warehouse":         DatabaseAutonomousDataWarehouseDataSource(),
		"oci_database_autonomous_data_warehouses":        DatabaseAutonomousDataWarehousesDataSource(),
		"oci_database_autonomous_data_warehouse_wallet":  DatabaseAutonomousDataWarehouseWalletDataSource(),
		"oci_database_autonomous_data_warehouse_backup":  DatabaseAutonomousDataWarehouseBackupDataSource(),
		"oci_database_autonomous_data_warehouse_backups": DatabaseAutonomousDataWarehouseBackupsDataSource(),
		"oci_database_autonomous_database":               DatabaseAutonomousDatabaseDataSource(),
		"oci_database_autonomous_databases":              DatabaseAutonomousDatabasesDataSource(),
		"oci_database_autonomous_database_wallet":        DatabaseAutonomousDatabaseWalletDataSource(),
		"oci_database_autonomous_database_backup":        DatabaseAutonomousDatabaseBackupDataSource(),
		"oci_database_autonomous_database_backups":       DatabaseAutonomousDatabaseBackupsDataSource(),
		"oci_database_backups":                           DatabaseBackupsDataSource(),
		"oci_database_database":                          DatabaseDatabaseDataSource(),
		"oci_database_databases":                         DatabaseDatabasesDataSource(),
		"oci_database_data_guard_association":            DatabaseDataGuardAssociationDataSource(),
		"oci_database_data_guard_associations":           DatabaseDataGuardAssociationsDataSource(),
		"oci_database_db_home":                           DatabaseDbHomeDataSource(),
		"oci_database_db_homes":                          DatabaseDbHomesDataSource(),
		"oci_database_db_node":                           DatabaseDbNodeDataSource(),
		"oci_database_db_nodes":                          DatabaseDbNodesDataSource(),
		"oci_database_db_system_shapes":                  DatabaseDbSystemShapesDataSource(),
		"oci_database_db_systems":                        DatabaseDbSystemsDataSource(),
		"oci_database_db_system_patches":                 DatabaseDbSystemPatchesDataSource(),
		"oci_database_db_system_patch_history_entries":   DatabaseDbSystemPatchHistoryEntriesDataSource(),
		"oci_database_db_versions":                       DatabaseDbVersionsDataSource(),
		"oci_database_db_home_patches":                   DatabaseDbHomePatchesDataSource(),
		"oci_database_db_home_patch_history_entries":     DatabaseDbHomePatchHistoryEntriesDataSource(),
		"oci_database_exadata_iorm_config":               DatabaseExadataIormConfigDataSource(),
		"oci_dns_records":                                DnsRecordsDataSource(),
		"oci_dns_zones":                                  DnsZonesDataSource(),
		"oci_dns_steering_policies":                      DnsSteeringPoliciesDataSource(),
		"oci_dns_steering_policy":                        DnsSteeringPolicyDataSource(),
		"oci_dns_steering_policy_attachment":             DnsSteeringPolicyAttachmentDataSource(),
		"oci_dns_steering_policy_attachments":            DnsSteeringPolicyAttachmentsDataSource(),
		"oci_email_senders":                              EmailSendersDataSource(),
		"oci_email_sender":                               EmailSenderDataSource(),
		"oci_email_suppressions":                         EmailSuppressionsDataSource(),
		"oci_email_suppression":                          EmailSuppressionDataSource(),
		"oci_file_storage_exports":                       FileStorageExportsDataSource(),
		"oci_file_storage_export_sets":                   FileStorageExportSetsDataSource(),
		"oci_file_storage_file_systems":                  FileStorageFileSystemsDataSource(),
		"oci_file_storage_mount_targets":                 FileStorageMountTargetsDataSource(),
		"oci_file_storage_snapshots":                     FileStorageSnapshotsDataSource(),
		"oci_health_checks_http_monitor":                 HealthChecksHttpMonitorDataSource(),
		"oci_health_checks_http_monitors":                HealthChecksHttpMonitorsDataSource(),
		"oci_health_checks_ping_monitor":                 HealthChecksPingMonitorDataSource(),
		"oci_health_checks_ping_monitors":                HealthChecksPingMonitorsDataSource(),
		"oci_health_checks_http_probe_results":           HealthChecksHttpProbeResultsDataSource(),
		"oci_health_checks_ping_probe_results":           HealthChecksPingProbeResultsDataSource(),
		"oci_health_checks_vantage_points":               HealthChecksVantagePointsDataSource(),
		"oci_identity_api_keys":                          IdentityApiKeysDataSource(),
		"oci_identity_authentication_policy":             IdentityAuthenticationPolicyDataSource(),
		"oci_identity_auth_tokens":                       IdentityAuthTokensDataSource(),
		"oci_identity_availability_domain":               IdentityAvailabilityDomainDataSource(),
		"oci_identity_availability_domains":              IdentityAvailabilityDomainsDataSource(),
		"oci_identity_compartment":                       IdentityCompartmentDataSource(),
		"oci_identity_compartments":                      IdentityCompartmentsDataSource(),
		"oci_identity_customer_secret_keys":              IdentityCustomerSecretKeysDataSource(),
		"oci_identity_dynamic_groups":                    IdentityDynamicGroupsDataSource(),
		"oci_identity_fault_domains":                     IdentityFaultDomainsDataSource(),
		"oci_identity_groups":                            IdentityGroupsDataSource(),
		"oci_identity_identity_providers":                IdentityIdentityProvidersDataSource(),
		"oci_identity_identity_provider_groups":          IdentityIdentityProviderGroupsDataSource(),
		"oci_identity_idp_group_mappings":                IdentityIdpGroupMappingsDataSource(),
		"oci_identity_policies":                          IdentityPoliciesDataSource(),
		"oci_identity_regions":                           IdentityRegionsDataSource(),
		"oci_identity_smtp_credentials":                  IdentitySmtpCredentialsDataSource(),
		"oci_identity_swift_passwords":                   IdentitySwiftPasswordsDataSource(),
		"oci_identity_tag_default":                       IdentityTagDefaultDataSource(),
		"oci_identity_tag_defaults":                      IdentityTagDefaultsDataSource(),
		"oci_identity_tag_namespaces":                    IdentityTagNamespacesDataSource(),
		"oci_identity_tags":                              IdentityTagsDataSource(),
		"oci_identity_tenancy":                           IdentityTenancyDataSource(),
		"oci_identity_user_group_memberships":            IdentityUserGroupMembershipsDataSource(),
		"oci_identity_users":                             IdentityUsersDataSource(),
		"oci_identity_region_subscriptions":              IdentityRegionSubscriptionsDataSource(),
		"oci_kms_decrypted_data":                         KmsDecryptedDataDataSource(),
		"oci_kms_encrypted_data":                         KmsEncryptedDataDataSource(),
		"oci_kms_key":                                    KmsKeyDataSource(),
		"oci_kms_keys":                                   KmsKeysDataSource(),
		"oci_kms_key_version":                            KmsKeyVersionDataSource(),
		"oci_kms_key_versions":                           KmsKeyVersionsDataSource(),
		"oci_kms_vault":                                  KmsVaultDataSource(),
		"oci_kms_vaults":                                 KmsVaultsDataSource(),
		"oci_load_balancer_backend_health":               LoadBalancerBackendHealthDataSource(),
		"oci_load_balancer_backends":                     LoadBalancerBackendsDataSource(),
		"oci_load_balancer_backend_set_health":           LoadBalancerBackendSetHealthDataSource(),
		"oci_load_balancer_backend_sets":                 LoadBalancerBackendSetsDataSource(),
		"oci_load_balancer_backendsets":                  LoadBalancerBackendSetsDataSource(),
		"oci_load_balancer_certificates":                 LoadBalancerCertificatesDataSource(),
		"oci_load_balancer_health":                       LoadBalancerLoadBalancerHealthDataSource(),
		"oci_load_balancer_hostnames":                    LoadBalancerHostnamesDataSource(),
		"oci_load_balancer_policies":                     LoadBalancerLoadBalancerPoliciesDataSource(),
		"oci_load_balancer_protocols":                    LoadBalancerLoadBalancerProtocolsDataSource(),
		"oci_load_balancer_shapes":                       LoadBalancerLoadBalancerShapesDataSource(),
		"oci_load_balancer_load_balancers":               LoadBalancerLoadBalancersDataSource(),
		"oci_load_balancers":                             LoadBalancerLoadBalancersDataSource(),
		"oci_load_balancer_path_route_sets":              LoadBalancerPathRouteSetsDataSource(),
		"oci_load_balancer_rule_sets":                    LoadBalancerRuleSetsDataSource(),
		"oci_load_balancer_rule_set":                     LoadBalancerRuleSetDataSource(),
		"oci_monitoring_alarm":                           MonitoringAlarmDataSource(),
		"oci_monitoring_alarms":                          MonitoringAlarmsDataSource(),
		"oci_monitoring_alarm_statuses":                  MonitoringAlarmStatusesDataSource(),
		"oci_monitoring_alarm_history_collection":        MonitoringAlarmHistoryCollectionDataSource(),
		"oci_monitoring_metrics":                         MonitoringMetricsDataSource(),
		"oci_monitoring_metric_data":                     MonitoringMetricDataDataSource(),
		"oci_objectstorage_bucket":                       ObjectStorageBucketDataSource(),
		"oci_objectstorage_bucket_summaries":             ObjectStorageBucketsDataSource(),
		"oci_objectstorage_object_lifecycle_policy":      ObjectStorageObjectLifecyclePolicyDataSource(),
		"oci_objectstorage_namespace":                    ObjectStorageNamespaceDataSource(),
		"oci_objectstorage_namespace_metadata":           ObjectStorageNamespaceMetadataDataSource(),
		"oci_objectstorage_object_head":                  ObjectStorageObjectHeadDataSource(),
		"oci_objectstorage_object":                       ObjectStorageObjectDataSource(),
		"oci_objectstorage_objects":                      ObjectStorageObjectsDataSource(),
		"oci_objectstorage_preauthrequest":               ObjectStoragePreauthenticatedRequestDataSource(),
		"oci_objectstorage_preauthrequests":              ObjectStoragePreauthenticatedRequestsDataSource(),
		"oci_ons_notification_topic":                     OnsNotificationTopicDataSource(),
		"oci_ons_notification_topics":                    OnsNotificationTopicsDataSource(),
		"oci_ons_subscription":                           OnsSubscriptionDataSource(),
		"oci_ons_subscriptions":                          OnsSubscriptionsDataSource(),
		"oci_streaming_stream":                           StreamingStreamDataSource(),
		"oci_streaming_streams":                          StreamingStreamsDataSource(),
		"oci_waas_waas_policy":                           WaasWaasPolicyDataSource(),
		"oci_waas_waas_policies":                         WaasWaasPoliciesDataSource(),
		"oci_waas_certificate":                           WaasCertificateDataSource(),
		"oci_waas_certificates":                          WaasCertificatesDataSource(),
		"oci_waas_edge_subnets":                          WaasEdgeSubnetsDataSource(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"oci_autoscaling_auto_scaling_configuration":              AutoscalingAutoScalingConfigurationResource(),
		"oci_budget_budget":                                       BudgetBudgetResource(),
		"oci_budget_alert_rule":                                   BudgetAlertRuleResource(),
		"oci_core_app_catalog_listing_resource_version_agreement": AppCatalogListingResourceVersionAgreementResource(),
		"oci_core_listing_resource_version_agreement":             AppCatalogListingResourceVersionAgreementResource(),
		"oci_core_app_catalog_subscription":                       CoreAppCatalogSubscriptionResource(),
		"oci_core_boot_volume":                                    CoreBootVolumeResource(),
		"oci_core_boot_volume_backup":                             CoreBootVolumeBackupResource(),
		"oci_audit_configuration":                                 AuditConfigurationResource(),
		"oci_containerengine_cluster":                             ContainerengineClusterResource(),
		"oci_containerengine_node_pool":                           ContainerengineNodePoolResource(),
		"oci_core_console_history":                                CoreConsoleHistoryResource(),
		"oci_core_cpe":                                            CoreCpeResource(),
		"oci_core_cross_connect":                                  CoreCrossConnectResource(),
		"oci_core_cross_connect_group":                            CoreCrossConnectGroupResource(),
		"oci_core_default_dhcp_options":                           DefaultCoreDhcpOptionsResource(),
		"oci_core_dhcp_options":                                   CoreDhcpOptionsResource(),
		"oci_core_drg":                                            CoreDrgResource(),
		"oci_core_drg_attachment":                                 CoreDrgAttachmentResource(),
		"oci_core_image":                                          CoreImageResource(),
		"oci_core_instance":                                       CoreInstanceResource(),
		"oci_core_instance_console_connection":                    CoreInstanceConsoleConnectionResource(),
		"oci_core_instance_configuration":                         CoreInstanceConfigurationResource(),
		"oci_core_instance_pool":                                  CoreInstancePoolResource(),
		"oci_core_internet_gateway":                               CoreInternetGatewayResource(),
		"oci_core_ipsec":                                          CoreIpSecConnectionResource(),
		"oci_core_local_peering_gateway":                          CoreLocalPeeringGatewayResource(),
		"oci_core_nat_gateway":                                    CoreNatGatewayResource(),
		"oci_core_private_ip":                                     CorePrivateIpResource(),
		"oci_core_public_ip":                                      CorePublicIpResource(),
		"oci_core_default_route_table":                            DefaultCoreRouteTableResource(),
		"oci_core_route_table":                                    CoreRouteTableResource(),
		"oci_core_route_table_attachment":                         CoreRouteTableAttachmentResource(),
		"oci_core_remote_peering_connection":                      CoreRemotePeeringConnectionResource(),
		"oci_core_default_security_list":                          CoreDefaultSecurityListResource(),
		"oci_core_security_list":                                  CoreSecurityListResource(),
		"oci_core_service_gateway":                                CoreServiceGatewayResource(),
		"oci_core_subnet":                                         CoreSubnetResource(),
		"oci_core_virtual_circuit":                                CoreVirtualCircuitResource(),
		"oci_core_virtual_network":                                CoreVcnResource(), //This is a legacy name for VCN, removing it can cause breaking changes
		"oci_core_vcn":                                            CoreVcnResource(),
		"oci_core_vnic_attachment":                                CoreVnicAttachmentResource(),
		"oci_core_volume":                                         CoreVolumeResource(),
		"oci_core_volume_group":                                   CoreVolumeGroupResource(),
		"oci_core_volume_group_backup":                            CoreVolumeGroupBackupResource(),
		"oci_core_volume_attachment":                              CoreVolumeAttachmentResource(),
		"oci_core_volume_backup":                                  CoreVolumeBackupResource(),
		"oci_core_volume_backup_policy_assignment":                CoreVolumeBackupPolicyAssignmentResource(),
		"oci_database_autonomous_data_warehouse":                  DatabaseAutonomousDataWarehouseResource(),
		"oci_database_autonomous_data_warehouse_backup":           DatabaseAutonomousDataWarehouseBackupResource(),
		"oci_database_autonomous_database":                        DatabaseAutonomousDatabaseResource(),
		"oci_database_autonomous_database_backup":                 DatabaseAutonomousDatabaseBackupResource(),
		"oci_database_data_guard_association":                     DatabaseDataGuardAssociationResource(),
		//Do remember to enable database_db_home_test if you are enabling DB Home resource
		//"oci_database_db_home":                     DatabaseDbHomeResource(),
		"oci_database_db_system":                    DatabaseDbSystemResource(),
		"oci_database_exadata_iorm_config":          DatabaseExadataIormConfigResource(),
		"oci_database_backup":                       DatabaseBackupResource(),
		"oci_dns_record":                            DnsRecordResource(),
		"oci_dns_steering_policy":                   DnsSteeringPolicyResource(),
		"oci_dns_steering_policy_attachment":        DnsSteeringPolicyAttachmentResource(),
		"oci_dns_zone":                              DnsZoneResource(),
		"oci_email_sender":                          EmailSenderResource(),
		"oci_email_suppression":                     EmailSuppressionResource(),
		"oci_file_storage_export":                   FileStorageExportResource(),
		"oci_file_storage_export_set":               FileStorageExportSetResource(),
		"oci_file_storage_file_system":              FileStorageFileSystemResource(),
		"oci_file_storage_mount_target":             FileStorageMountTargetResource(),
		"oci_file_storage_snapshot":                 FileStorageSnapshotResource(),
		"oci_health_checks_http_monitor":            HealthChecksHttpMonitorResource(),
		"oci_health_checks_ping_monitor":            HealthChecksPingMonitorResource(),
		"oci_health_checks_http_probe":              HealthChecksHttpProbeResource(),
		"oci_health_checks_ping_probe":              HealthChecksPingProbeResource(),
		"oci_identity_api_key":                      IdentityApiKeyResource(),
		"oci_identity_authentication_policy":        IdentityAuthenticationPolicyResource(),
		"oci_identity_auth_token":                   IdentityAuthTokenResource(),
		"oci_identity_compartment":                  IdentityCompartmentResource(),
		"oci_identity_customer_secret_key":          IdentityCustomerSecretKeyResource(),
		"oci_identity_dynamic_group":                IdentityDynamicGroupResource(),
		"oci_identity_group":                        IdentityGroupResource(),
		"oci_identity_identity_provider":            IdentityIdentityProviderResource(),
		"oci_identity_idp_group_mapping":            IdentityIdpGroupMappingResource(),
		"oci_identity_policy":                       IdentityPolicyResource(),
		"oci_identity_smtp_credential":              IdentitySmtpCredentialResource(),
		"oci_identity_swift_password":               IdentitySwiftPasswordResource(),
		"oci_identity_tag":                          IdentityTagResource(),
		"oci_identity_tag_default":                  IdentityTagDefaultResource(),
		"oci_identity_tag_namespace":                IdentityTagNamespaceResource(),
		"oci_identity_ui_password":                  IdentityUiPasswordResource(),
		"oci_identity_user":                         IdentityUserResource(),
		"oci_identity_user_capabilities_management": IdentityUserCapabilitiesManagementResource(),
		"oci_identity_user_group_membership":        IdentityUserGroupMembershipResource(),
		"oci_kms_encrypted_data":                    KmsEncryptedDataResource(),
		"oci_kms_generated_key":                     KmsGeneratedKeyResource(),
		"oci_kms_key":                               KmsKeyResource(),
		"oci_kms_key_version":                       KmsKeyVersionResource(),
		"oci_kms_vault":                             KmsVaultResource(),
		"oci_load_balancer":                         LoadBalancerLoadBalancerResource(),
		"oci_load_balancer_load_balancer":           LoadBalancerLoadBalancerResource(),
		"oci_load_balancer_backend":                 LoadBalancerBackendResource(),
		"oci_load_balancer_backend_set":             LoadBalancerBackendSetResource(),
		"oci_load_balancer_backendset":              LoadBalancerBackendSetResource(),
		"oci_load_balancer_certificate":             LoadBalancerCertificateResource(),
		"oci_load_balancer_listener":                LoadBalancerListenerResource(),
		"oci_load_balancer_hostname":                LoadBalancerHostnameResource(),
		"oci_load_balancer_path_route_set":          LoadBalancerPathRouteSetResource(),
		"oci_load_balancer_rule_set":                LoadBalancerRuleSetResource(),
		"oci_monitoring_alarm":                      MonitoringAlarmResource(),
		"oci_objectstorage_bucket":                  ObjectStorageBucketResource(),
		"oci_objectstorage_object_lifecycle_policy": ObjectStorageObjectLifecyclePolicyResource(),
		"oci_objectstorage_object":                  ObjectStorageObjectResource(),
		"oci_objectstorage_namespace_metadata":      ObjectStorageNamespaceMetadataResource(),
		"oci_objectstorage_preauthrequest":          ObjectStoragePreauthenticatedRequestResource(),
		"oci_ons_notification_topic":                OnsNotificationTopicResource(),
		"oci_ons_subscription":                      OnsSubscriptionResource(),
		"oci_streaming_stream":                      StreamingStreamResource(),
		"oci_waas_waas_policy":                      WaasWaasPolicyResource(),
		"oci_waas_certificate":                      WaasCertificateResource(),
	}
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

func validateConfigForAPIKeyAuth(d *schema.ResourceData) error {
	_, hasTenancyOCID := d.GetOkExists(tenancyOcidAttrName)
	_, hasUserOCID := d.GetOkExists(userOcidAttrName)
	_, hasFingerprint := d.GetOkExists(fingerprintAttrName)
	if !hasTenancyOCID || !hasUserOCID || !hasFingerprint {
		return fmt.Errorf("when auth is set to '%s', tenancy_ocid, user_ocid, and fingerprint are required", authAPIKeySetting)
	}
	return nil
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

func ProviderConfig(d *schema.ResourceData) (clients interface{}, err error) {
	clients = &OracleClients{configuration: map[string]string{}}

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

	auth := strings.ToLower(d.Get(authAttrName).(string))
	clients.(*OracleClients).configuration[authAttrName] = auth

	userAgentProviderName := getEnvSettingWithDefault(userAgentProviderNameEnv, defaultUserAgentProviderName)
	userAgent := fmt.Sprintf(userAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, terraform.VersionString(), userAgentProviderName, Version)

	httpClient := &http.Client{
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

	var configProviders []oci_common.ConfigurationProvider

	switch auth {
	case strings.ToLower(authAPIKeySetting):
		if err := validateConfigForAPIKeyAuth(d); err != nil {
			return nil, err
		}
	case strings.ToLower(authInstancePrincipalSetting):
		apiKeyConfigVariablesToUnset, ok := checkIncompatibleAttrsForApiKeyAuth(d)
		if !ok {
			return nil, fmt.Errorf(`user credentials %v should be removed from the configuration`, strings.Join(apiKeyConfigVariablesToUnset, ", "))
		}

		region, ok := d.GetOkExists(regionAttrName)
		if !ok {
			return nil, fmt.Errorf("can not get %s from Terraform configuration (InstancePrincipal)", regionAttrName)
		}
		cfg, err := oci_common_auth.InstancePrincipalConfigurationProviderForRegion(oci_common.StringToRegion(region.(string)))
		if err != nil {
			return nil, err
		}
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
		configProviders = append(configProviders, cfg)
	default:
		return nil, fmt.Errorf("auth must be one of '%s' or '%s' or '%s'", authAPIKeySetting, authInstancePrincipalSetting, authInstancePrincipalWithCertsSetting)
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

	avoidWaitingForDeleteTarget, _ = strconv.ParseBool(getEnvSettingWithDefault("avoid_waiting_for_delete_target", "false"))

	return clients, nil
}

type ResourceDataConfigProvider struct {
	D *schema.ResourceData
}

// TODO: The error messages returned by following methods get swallowed up by the ComposingConfigurationProvider,
// since it only checks whether an error exists or not.
// The ComposingConfigurationProvider in SDK should log the errors as debug statements instead.

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
		pemFileContent, readFileErr := ioutil.ReadFile(privateKeyPath.(string))
		if readFileErr != nil {
			return nil, fmt.Errorf("Can not read private key from: '%s', Error: %q", privateKeyPath, readFileErr)
		}
		return oci_common.PrivateKeyFromBytes(pemFileContent, &password)
	}

	return nil, fmt.Errorf("can not get private_key or private_key_path from Terraform configuration")
}
