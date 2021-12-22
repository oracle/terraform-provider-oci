package provider

import (
	tf_ai_anomaly_detection "github.com/terraform-providers/terraform-provider-oci/internal/service/ai_anomaly_detection"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/analytics"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/audit"
	tf_bds "github.com/terraform-providers/terraform-provider-oci/internal/service/bds"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/budget"
	tf_core "github.com/terraform-providers/terraform-provider-oci/internal/service/core"
	tf_email "github.com/terraform-providers/terraform-provider-oci/internal/service/email"
	tf_health_checks "github.com/terraform-providers/terraform-provider-oci/internal/service/health_checks"
	tf_identity "github.com/terraform-providers/terraform-provider-oci/internal/service/identity"
	tf_jms "github.com/terraform-providers/terraform-provider-oci/internal/service/jms"
	tf_kms "github.com/terraform-providers/terraform-provider-oci/internal/service/kms"
	tf_limits "github.com/terraform-providers/terraform-provider-oci/internal/service/limits"
	tf_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/load_balancer"
	tf_management_agent "github.com/terraform-providers/terraform-provider-oci/internal/service/management_agent"
	tf_management_dashboard "github.com/terraform-providers/terraform-provider-oci/internal/service/management_dashboard"
	tf_marketplace "github.com/terraform-providers/terraform-provider-oci/internal/service/marketplace"
	tf_mysql "github.com/terraform-providers/terraform-provider-oci/internal/service/mysql"
	tf_nosql "github.com/terraform-providers/terraform-provider-oci/internal/service/nosql"
)

func init() {
	// ai anomaly service
	RegisterResource("oci_ai_anomaly_detection_ai_private_endpoint", tf_ai_anomaly_detection.AiAnomalyDetectionAiPrivateEndpointResource())
	RegisterResource("oci_ai_anomaly_detection_data_asset", tf_ai_anomaly_detection.AiAnomalyDetectionDataAssetResource())
	RegisterResource("oci_ai_anomaly_detection_model", tf_ai_anomaly_detection.AiAnomalyDetectionModelResource())
	RegisterResource("oci_ai_anomaly_detection_project", tf_ai_anomaly_detection.AiAnomalyDetectionProjectResource())

	// analytics service
	RegisterResource("oci_analytics_analytics_instance", analytics.AnalyticsAnalyticsInstanceResource())
	RegisterResource("oci_analytics_analytics_instance_private_access_channel", analytics.AnalyticsAnalyticsInstancePrivateAccessChannelResource())
	RegisterResource("oci_analytics_analytics_instance_vanity_url", analytics.AnalyticsAnalyticsInstanceVanityUrlResource())

	// bds service
	RegisterResource("oci_bds_auto_scaling_configuration", tf_bds.BdsAutoScalingConfigurationResource())
	RegisterResource("oci_bds_bds_instance", tf_bds.BdsBdsInstanceResource())

	// budget service
	RegisterResource("oci_budget_alert_rule", budget.BudgetAlertRuleResource())
	RegisterResource("oci_budget_budget", budget.BudgetBudgetResource())

	// email service
	RegisterResource("oci_email_dkim", tf_email.EmailDkimResource())
	RegisterResource("oci_email_suppression", tf_email.EmailSuppressionResource())
	RegisterResource("oci_email_email_domain", tf_email.EmailEmailDomainResource())
	RegisterResource("oci_email_sender", tf_email.EmailSenderResource())

	// identity service
	RegisterResource("oci_health_checks_ping_monitor", tf_health_checks.HealthChecksPingMonitorResource())
	RegisterResource("oci_health_checks_ping_probe", tf_health_checks.HealthChecksPingProbeResource())
	RegisterResource("oci_health_checks_http_probe", tf_health_checks.HealthChecksHttpProbeResource())
	RegisterResource("oci_health_checks_http_monitor", tf_health_checks.HealthChecksHttpMonitorResource())
	RegisterResource("oci_identity_api_key", tf_identity.IdentityApiKeyResource())
	RegisterResource("oci_identity_auth_token", tf_identity.IdentityAuthTokenResource())
	RegisterResource("oci_identity_authentication_policy", tf_identity.IdentityAuthenticationPolicyResource())
	RegisterResource("oci_identity_compartment", tf_identity.IdentityCompartmentResource())
	RegisterResource("oci_identity_customer_secret_key", tf_identity.IdentityCustomerSecretKeyResource())
	RegisterResource("oci_identity_dynamic_group", tf_identity.IdentityDynamicGroupResource())
	RegisterResource("oci_identity_smtp_credential", tf_identity.IdentitySmtpCredentialResource())
	RegisterResource("oci_identity_authentication_policy", tf_identity.IdentityAuthenticationPolicyResource())
	RegisterResource("oci_identity_idp_group_mapping", tf_identity.IdentityIdpGroupMappingResource())
	RegisterResource("oci_identity_ui_password", tf_identity.IdentityUiPasswordResource())
	RegisterResource("oci_identity_user_capabilities_management", tf_identity.IdentityUserCapabilitiesManagementResource())
	RegisterResource("oci_identity_tag_default", tf_identity.IdentityTagDefaultResource())
	RegisterResource("oci_identity_network_source", tf_identity.IdentityNetworkSourceResource())
	RegisterResource("oci_identity_identity_provider", tf_identity.IdentityIdentityProviderResource())
	RegisterResource("oci_identity_data_plane_generate_scoped_access_token", tf_identity.IdentityDataPlaneGenerateScopedAccessTokenResource())
	RegisterResource("oci_identity_swift_password", tf_identity.IdentitySwiftPasswordResource())
	RegisterResource("oci_identity_user_group_membership", tf_identity.IdentityUserGroupMembershipResource())
	RegisterResource("oci_identity_db_credential", tf_identity.IdentityDbCredentialResource())
	RegisterResource("oci_identity_domain_replication_to_region", tf_identity.IdentityDomainReplicationToRegionResource())
	RegisterResource("oci_identity_domain", tf_identity.IdentityDomainResource())
	RegisterResource("oci_identity_auth_token", tf_identity.IdentityAuthTokenResource())
	RegisterResource("oci_identity_tag", tf_identity.IdentityTagResource())
	RegisterResource("oci_identity_api_key", tf_identity.IdentityApiKeyResource())
	RegisterResource("oci_identity_group", tf_identity.IdentityGroupResource())
	RegisterResource("oci_identity_compartment", tf_identity.IdentityCompartmentResource())
	RegisterResource("oci_identity_tag_namespace", tf_identity.IdentityTagNamespaceResource())
	RegisterResource("oci_identity_policy", tf_identity.IdentityPolicyResource())
	RegisterResource("oci_identity_user", tf_identity.IdentityUserResource())
	RegisterResource("oci_identity_customer_secret_key", tf_identity.IdentityCustomerSecretKeyResource())
	RegisterResource("oci_audit_configuration", audit.AuditConfigurationResource())

	// kms service
	RegisterResource("oci_kms_verify", tf_kms.KmsVerifyResource())
	RegisterResource("oci_kms_vault_replication", tf_kms.KmsVaultReplicationResource())
	RegisterResource("oci_kms_generated_key", tf_kms.KmsGeneratedKeyResource())
	RegisterResource("oci_kms_key", tf_kms.KmsKeyResource())
	RegisterResource("oci_kms_key_version", tf_kms.KmsKeyVersionResource())
	RegisterResource("oci_kms_encrypted_data", tf_kms.KmsEncryptedDataResource())
	RegisterResource("oci_kms_sign", tf_kms.KmsSignResource())
	RegisterResource("oci_kms_vault", tf_kms.KmsVaultResource())

	//load_balancer service
	RegisterResource("oci_load_balancer_certificate", tf_load_balancer.LoadBalancerCertificateResource())
	RegisterResource("oci_load_balancer_rule_set", tf_load_balancer.LoadBalancerRuleSetResource())
	RegisterResource("oci_load_balancer_listener", tf_load_balancer.LoadBalancerListenerResource())
	RegisterResource("oci_load_balancer_path_route_set", tf_load_balancer.LoadBalancerPathRouteSetResource())
	RegisterResource("oci_load_balancer_hostname", tf_load_balancer.LoadBalancerHostnameResource())
	RegisterResource("oci_load_balancer_ssl_cipher_suite", tf_load_balancer.LoadBalancerSslCipherSuiteResource())
	RegisterResource("oci_load_balancer_backend_set", tf_load_balancer.LoadBalancerBackendSetResource())
	RegisterResource("oci_load_balancer_backend", tf_load_balancer.LoadBalancerBackendResource())
	RegisterResource("oci_load_balancer_load_balancer_routing_policy", tf_load_balancer.LoadBalancerLoadBalancerRoutingPolicyResource())
	RegisterResource("oci_load_balancer_load_balancer", tf_load_balancer.LoadBalancerLoadBalancerResource())

	// core service
	RegisterResource("oci_core_drg_route_table_route_rule", tf_core.CoreDrgRouteTableRouteRuleResource())
	RegisterResource("oci_core_public_ip", tf_core.CorePublicIpResource())
	RegisterResource("oci_core_drg_route_distribution_statement", tf_core.CoreDrgRouteDistributionStatementResource())
	RegisterResource("oci_core_service_gateway", tf_core.CoreServiceGatewayResource())
	RegisterResource("oci_core_route_table_attachment", tf_core.CoreRouteTableAttachmentResource())
	RegisterResource("oci_core_image", tf_core.CoreImageResource())
	RegisterResource("oci_core_ipsec_connection_tunnel_management", tf_core.CoreIpSecConnectionTunnelManagementResource())
	RegisterResource("oci_core_shape_management", tf_core.CoreShapeResource())
	RegisterResource("oci_core_drg_route_distribution", tf_core.CoreDrgRouteDistributionResource())
	RegisterResource("oci_core_volume", tf_core.CoreVolumeResource())
	RegisterResource("oci_core_boot_volume", tf_core.CoreBootVolumeResource())
	RegisterResource("oci_core_network_security_group_security_rule", tf_core.CoreNetworkSecurityGroupSecurityRuleResource())
	RegisterResource("oci_core_public_ip_pool_capacity", tf_core.PublicIpPoolCapacityResource())
	RegisterResource("oci_core_virtual_circuit", tf_core.CoreVirtualCircuitResource())
	RegisterResource("oci_core_local_peering_gateway", tf_core.CoreLocalPeeringGatewayResource())
	RegisterResource("oci_core_drg_attachment", tf_core.CoreDrgAttachmentResource())
	RegisterResource("oci_core_internet_gateway", tf_core.CoreInternetGatewayResource())
	RegisterResource("oci_core_vcn", tf_core.CoreVcnResource())
	RegisterResource("oci_core_app_catalog_subscription", tf_core.CoreAppCatalogSubscriptionResource())
	RegisterResource("oci_core_remote_peering_connection", tf_core.CoreRemotePeeringConnectionResource())
	RegisterResource("oci_core_default_dhcp_options", tf_core.DefaultCoreDhcpOptionsResource())
	RegisterResource("oci_core_console_history", tf_core.CoreConsoleHistoryResource())
	RegisterResource("oci_core_ipsec", tf_core.CoreIpSecConnectionResource())
	RegisterResource("oci_core_volume_group", tf_core.CoreVolumeGroupResource())
	RegisterResource("oci_core_network_security_group", tf_core.CoreNetworkSecurityGroupResource())
	RegisterResource("oci_core_volume_backup_policy", tf_core.CoreVolumeBackupPolicyResource())
	RegisterResource("oci_core_app_catalog_listing_resource_version_agreement", tf_core.AppCatalogListingResourceVersionAgreementResource())
	RegisterResource("oci_core_listing_resource_version_agreement", tf_core.AppCatalogListingResourceVersionAgreementResource())
	RegisterResource("oci_core_instance_pool_instance", tf_core.CoreInstancePoolInstanceResource())
	RegisterResource("oci_core_nat_gateway", tf_core.CoreNatGatewayResource())
	RegisterResource("oci_core_subnet", tf_core.CoreSubnetResource())
	RegisterResource("oci_core_default_route_table", tf_core.DefaultCoreRouteTableResource())
	RegisterResource("oci_core_compute_capacity_reservation", tf_core.CoreComputeCapacityReservationResource())
	RegisterResource("oci_core_route_table", tf_core.CoreRouteTableResource())
	RegisterResource("oci_core_vnic_attachment", tf_core.CoreVnicAttachmentResource())
	RegisterResource("oci_core_volume_attachment", tf_core.CoreVolumeAttachmentResource())
	RegisterResource("oci_core_drg_route_table", tf_core.CoreDrgRouteTableResource())
	RegisterResource("oci_core_instance", tf_core.CoreInstanceResource())
	RegisterResource("oci_core_boot_volume_backup", tf_core.CoreBootVolumeBackupResource())
	RegisterResource("oci_core_default_security_list", tf_core.CoreDefaultSecurityListResource())
	RegisterResource("oci_core_private_ip", tf_core.CorePrivateIpResource())
	RegisterResource("oci_core_dedicated_vm_host", tf_core.CoreDedicatedVmHostResource())
	RegisterResource("oci_core_vlan", tf_core.CoreVlanResource())
	RegisterResource("oci_core_instance_console_connection", tf_core.CoreInstanceConsoleConnectionResource())
	RegisterResource("oci_core_instance_configuration", tf_core.CoreInstanceConfigurationResource())
	RegisterResource("oci_core_cross_connect_group", tf_core.CoreCrossConnectGroupResource())
	RegisterResource("oci_core_public_ip_pool", tf_core.CorePublicIpPoolResource())
	RegisterResource("oci_core_dhcp_options", tf_core.CoreDhcpOptionsResource())
	RegisterResource("oci_core_cross_connect", tf_core.CoreCrossConnectResource())
	RegisterResource("oci_core_compute_image_capability_schema", tf_core.CoreComputeImageCapabilitySchemaResource())
	RegisterResource("oci_core_drg", tf_core.CoreDrgResource())
	RegisterResource("oci_core_security_list", tf_core.CoreSecurityListResource())
	RegisterResource("oci_core_cpe", tf_core.CoreCpeResource())
	RegisterResource("oci_core_ipv6", tf_core.CoreIpv6Resource())
	RegisterResource("oci_core_cluster_network", tf_core.CoreClusterNetworkResource())
	RegisterResource("oci_core_drg_attachment_management", tf_core.CoreDrgAttachmentManagementResource())
	RegisterResource("oci_core_volume_backup_policy_assignment", tf_core.CoreVolumeBackupPolicyAssignmentResource())
	RegisterResource("oci_core_volume_backup", tf_core.CoreVolumeBackupResource())
	RegisterResource("oci_core_instance_pool", tf_core.CoreInstancePoolResource())
	RegisterResource("oci_core_drg_attachments_list", tf_core.CoreDrgAttachmentsListResource())
	RegisterResource("oci_core_volume_group_backup", tf_core.CoreVolumeGroupBackupResource())

	// JMS Service
	RegisterResource("oci_jms_fleet", tf_jms.JmsFleetResource())

	// mysql service
	RegisterResource("oci_mysql_mysql_backup", tf_mysql.MysqlMysqlBackupResource())
	RegisterResource("oci_mysql_analytics_cluster", tf_mysql.MysqlAnalyticsClusterResource())
	RegisterResource("oci_mysql_channel", tf_mysql.MysqlChannelResource())
	RegisterResource("oci_mysql_heat_wave_cluster", tf_mysql.MysqlHeatWaveClusterResource())
	RegisterResource("oci_mysql_mysql_db_system", tf_mysql.MysqlMysqlDbSystemResource())

	// nosql service
	RegisterResource("oci_nosql_table", tf_nosql.NosqlTableResource())
	RegisterResource("oci_nosql_index", tf_nosql.NosqlIndexResource())
	RegisterResource("oci_limits_quota", tf_limits.LimitsQuotaResource())
	RegisterResource("oci_management_agent_management_agent", tf_management_agent.ManagementAgentManagementAgentResource())
	RegisterResource("oci_management_agent_management_agent_install_key", tf_management_agent.ManagementAgentManagementAgentInstallKeyResource())
	RegisterResource("oci_management_dashboard_management_dashboards_import", tf_management_dashboard.ManagementDashboardManagementDashboardsImportResource())
	RegisterResource("oci_marketplace_publication", tf_marketplace.MarketplacePublicationResource())
	RegisterResource("oci_marketplace_accepted_agreement", tf_marketplace.MarketplaceAcceptedAgreementResource())
	RegisterResource("oci_marketplace_listing_package_agreement", tf_marketplace.MarketplaceListingPackageAgreementResource())
}
