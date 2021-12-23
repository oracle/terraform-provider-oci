package provider

import (
	tf_ai_anomaly_detection "github.com/terraform-providers/terraform-provider-oci/internal/service/ai_anomaly_detection"
	tf_analytics "github.com/terraform-providers/terraform-provider-oci/internal/service/analytics"
	tf_apigateway "github.com/terraform-providers/terraform-provider-oci/internal/service/apigateway"
	tf_apm "github.com/terraform-providers/terraform-provider-oci/internal/service/apm"
	tf_artifacts "github.com/terraform-providers/terraform-provider-oci/internal/service/artifacts"
	tf_audit "github.com/terraform-providers/terraform-provider-oci/internal/service/audit"
	tf_autoscaling "github.com/terraform-providers/terraform-provider-oci/internal/service/autoscaling"
	tf_bds "github.com/terraform-providers/terraform-provider-oci/internal/service/bds"
	tf_blockchain "github.com/terraform-providers/terraform-provider-oci/internal/service/blockchain"
	tf_budget "github.com/terraform-providers/terraform-provider-oci/internal/service/budget"
	tf_cloud_guard "github.com/terraform-providers/terraform-provider-oci/internal/service/cloud_guard"
	tf_containerengine "github.com/terraform-providers/terraform-provider-oci/internal/service/containerengine"
	tf_core "github.com/terraform-providers/terraform-provider-oci/internal/service/core"
	tf_data_labeling_service "github.com/terraform-providers/terraform-provider-oci/internal/service/data_labeling_service"
	tf_database "github.com/terraform-providers/terraform-provider-oci/internal/service/database"
	tf_database_management "github.com/terraform-providers/terraform-provider-oci/internal/service/database_management"
	tf_database_tools "github.com/terraform-providers/terraform-provider-oci/internal/service/database_tools"
	tf_datacatalog "github.com/terraform-providers/terraform-provider-oci/internal/service/datacatalog"
	tf_dataflow "github.com/terraform-providers/terraform-provider-oci/internal/service/dataflow"
	tf_dataintegration "github.com/terraform-providers/terraform-provider-oci/internal/service/dataintegration"
	tf_email "github.com/terraform-providers/terraform-provider-oci/internal/service/email"
	tf_functions "github.com/terraform-providers/terraform-provider-oci/internal/service/functions"
	tf_generic_artifacts_content "github.com/terraform-providers/terraform-provider-oci/internal/service/generic_artifacts_content"
	tf_health_checks "github.com/terraform-providers/terraform-provider-oci/internal/service/health_checks"
	tf_identity "github.com/terraform-providers/terraform-provider-oci/internal/service/identity"
	tf_integration "github.com/terraform-providers/terraform-provider-oci/internal/service/integration"
	tf_jms "github.com/terraform-providers/terraform-provider-oci/internal/service/jms"
	tf_kms "github.com/terraform-providers/terraform-provider-oci/internal/service/kms"
	tf_limits "github.com/terraform-providers/terraform-provider-oci/internal/service/limits"
	tf_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/load_balancer"
	tf_management_agent "github.com/terraform-providers/terraform-provider-oci/internal/service/management_agent"
	tf_management_dashboard "github.com/terraform-providers/terraform-provider-oci/internal/service/management_dashboard"
	tf_marketplace "github.com/terraform-providers/terraform-provider-oci/internal/service/marketplace"
	tf_metering_computation "github.com/terraform-providers/terraform-provider-oci/internal/service/metering_computation"
	tf_monitoring "github.com/terraform-providers/terraform-provider-oci/internal/service/monitoring"
	tf_mysql "github.com/terraform-providers/terraform-provider-oci/internal/service/mysql"
	tf_nosql "github.com/terraform-providers/terraform-provider-oci/internal/service/nosql"
	tf_objectstorage "github.com/terraform-providers/terraform-provider-oci/internal/service/objectstorage"
	tf_ocvp "github.com/terraform-providers/terraform-provider-oci/internal/service/ocvp"
	tf_oda "github.com/terraform-providers/terraform-provider-oci/internal/service/oda"
	tf_ons "github.com/terraform-providers/terraform-provider-oci/internal/service/ons"
	tf_operator_access_control "github.com/terraform-providers/terraform-provider-oci/internal/service/operator_access_control"
	tf_optimizer "github.com/terraform-providers/terraform-provider-oci/internal/service/optimizer"
	tf_osmanagement "github.com/terraform-providers/terraform-provider-oci/internal/service/osmanagement"
	tf_service_catalog "github.com/terraform-providers/terraform-provider-oci/internal/service/service_catalog"
	tf_usage_proxy "github.com/terraform-providers/terraform-provider-oci/internal/service/usage_proxy"
)

func init() {
	// ai anomaly service
	RegisterResource("oci_ai_anomaly_detection_ai_private_endpoint", tf_ai_anomaly_detection.AiAnomalyDetectionAiPrivateEndpointResource())
	RegisterResource("oci_ai_anomaly_detection_data_asset", tf_ai_anomaly_detection.AiAnomalyDetectionDataAssetResource())
	RegisterResource("oci_ai_anomaly_detection_model", tf_ai_anomaly_detection.AiAnomalyDetectionModelResource())
	RegisterResource("oci_ai_anomaly_detection_project", tf_ai_anomaly_detection.AiAnomalyDetectionProjectResource())

	// analytics service
	RegisterResource("oci_analytics_analytics_instance_vanity_url", tf_analytics.AnalyticsAnalyticsInstanceVanityUrlResource())
	RegisterResource("oci_analytics_analytics_instance", tf_analytics.AnalyticsAnalyticsInstanceResource())
	RegisterResource("oci_analytics_analytics_instance_private_access_channel", tf_analytics.AnalyticsAnalyticsInstancePrivateAccessChannelResource())

	//apm service
	RegisterResource("oci_apm_apm_domain", tf_apm.ApmApmDomainResource())

	//autoscaling service
	RegisterResource("oci_autoscaling_auto_scaling_configuration", tf_autoscaling.AutoScalingAutoScalingConfigurationResource())

	// bds service
	RegisterResource("oci_bds_auto_scaling_configuration", tf_bds.BdsAutoScalingConfigurationResource())
	RegisterResource("oci_bds_bds_instance", tf_bds.BdsBdsInstanceResource())

	//Blockchain Service
	RegisterResource("oci_blockchain_blockchain_platform", tf_blockchain.BlockchainBlockchainPlatformResource())
	RegisterResource("oci_blockchain_peer", tf_blockchain.BlockchainPeerResource())
	RegisterResource("oci_blockchain_osn", tf_blockchain.BlockchainOsnResource())

	// budget service
	RegisterResource("oci_budget_alert_rule", tf_budget.BudgetAlertRuleResource())
	RegisterResource("oci_budget_budget", tf_budget.BudgetBudgetResource())

	// datacatalog service
	RegisterResource("oci_datacatalog_metastore", tf_datacatalog.DatacatalogMetastoreResource())
	RegisterResource("oci_datacatalog_data_asset", tf_datacatalog.DatacatalogDataAssetResource())
	RegisterResource("oci_datacatalog_connection", tf_datacatalog.DatacatalogConnectionResource())
	RegisterResource("oci_datacatalog_catalog", tf_datacatalog.DatacatalogCatalogResource())
	RegisterResource("oci_datacatalog_catalog_private_endpoint", tf_datacatalog.DatacatalogCatalogPrivateEndpointResource())

	// data labeling service
	RegisterResource("oci_data_labeling_service_dataset", tf_data_labeling_service.DataLabelingServiceDatasetResource())

	// email service
	RegisterResource("oci_email_dkim", tf_email.EmailDkimResource())
	RegisterResource("oci_email_suppression", tf_email.EmailSuppressionResource())
	RegisterResource("oci_email_email_domain", tf_email.EmailEmailDomainResource())
	RegisterResource("oci_email_sender", tf_email.EmailSenderResource())

	//function service
	RegisterResource("oci_functions_invoke_function", tf_functions.FunctionsInvokeFunctionResource())
	RegisterResource("oci_functions_function", tf_functions.FunctionsFunctionResource())
	RegisterResource("oci_functions_application", tf_functions.FunctionsApplicationResource())

	// identity service
	RegisterResource("oci_health_checks_ping_monitor", tf_health_checks.HealthChecksPingMonitorResource())
	RegisterResource("oci_health_checks_ping_probe", tf_health_checks.HealthChecksPingProbeResource())
	RegisterResource("oci_health_checks_http_probe", tf_health_checks.HealthChecksHttpProbeResource())
	RegisterResource("oci_health_checks_http_monitor", tf_health_checks.HealthChecksHttpMonitorResource())
	RegisterResource("oci_cloud_guard_responder_recipe", tf_cloud_guard.CloudGuardResponderRecipeResource())
	RegisterResource("oci_cloud_guard_data_mask_rule", tf_cloud_guard.CloudGuardDataMaskRuleResource())
	RegisterResource("oci_cloud_guard_cloud_guard_configuration", tf_cloud_guard.CloudGuardCloudGuardConfigurationResource())
	RegisterResource("oci_cloud_guard_target", tf_cloud_guard.CloudGuardTargetResource())
	RegisterResource("oci_cloud_guard_managed_list", tf_cloud_guard.CloudGuardManagedListResource())
	RegisterResource("oci_cloud_guard_detector_recipe", tf_cloud_guard.CloudGuardDetectorRecipeResource())
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
	RegisterResource("oci_audit_configuration", tf_audit.AuditConfigurationResource())

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

	// monitoring service
	RegisterResource("oci_monitoring_alarm", tf_monitoring.MonitoringAlarmResource())

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
	RegisterResource("oci_metering_computation_usage", tf_metering_computation.MeteringComputationUsageResource())
	RegisterResource("oci_metering_computation_query", tf_metering_computation.MeteringComputationQueryResource())
	RegisterResource("oci_metering_computation_custom_table", tf_metering_computation.MeteringComputationCustomTableResource())
	RegisterResource("oci_ocvp_esxi_host", tf_ocvp.OcvpEsxiHostResource())
	RegisterResource("oci_ocvp_sddc", tf_ocvp.OcvpSddcResource())
	RegisterResource("oci_oda_oda_instance", tf_oda.OdaOdaInstanceResource())
	RegisterResource("oci_operator_access_control_operator_control_assignment", tf_operator_access_control.OperatorAccessControlOperatorControlAssignmentResource())
	RegisterResource("oci_operator_access_control_operator_control", tf_operator_access_control.OperatorAccessControlOperatorControlResource())
	RegisterResource("oci_optimizer_enrollment_status", tf_optimizer.OptimizerEnrollmentStatusResource())
	RegisterResource("oci_optimizer_recommendation", tf_optimizer.OptimizerRecommendationResource())
	RegisterResource("oci_optimizer_profile", tf_optimizer.OptimizerProfileResource())
	RegisterResource("oci_optimizer_resource_action", tf_optimizer.OptimizerResourceActionResource())
	RegisterResource("oci_service_catalog_service_catalog_association", tf_service_catalog.ServiceCatalogServiceCatalogAssociationResource())
	RegisterResource("oci_service_catalog_service_catalog", tf_service_catalog.ServiceCatalogServiceCatalogResource())
	RegisterResource("oci_service_catalog_private_application", tf_service_catalog.ServiceCatalogPrivateApplicationResource())
	RegisterResource("oci_apigateway_certificate", tf_apigateway.ApigatewayCertificateResource())
	RegisterResource("oci_apigateway_gateway", tf_apigateway.ApigatewayGatewayResource())
	RegisterResource("oci_apigateway_deployment", tf_apigateway.ApigatewayDeploymentResource())
	RegisterResource("oci_apigateway_api", tf_apigateway.ApigatewayApiResource())
	RegisterResource("oci_artifacts_generic_artifact", tf_artifacts.ArtifactsGenericArtifactResource())
	RegisterResource("oci_artifacts_repository", tf_artifacts.ArtifactsRepositoryResource())
	RegisterResource("oci_artifacts_container_configuration", tf_artifacts.ArtifactsContainerConfigurationResource())
	RegisterResource("oci_artifacts_container_image_signature", tf_artifacts.ArtifactsContainerImageSignatureResource())
	RegisterResource("oci_artifacts_container_repository", tf_artifacts.ArtifactsContainerRepositoryResource())
	RegisterResource("oci_generic_artifacts_content_artifact_by_path", tf_generic_artifacts_content.GenericArtifactsContentArtifactByPathResource())

	//ONS Service
	RegisterResource("oci_ons_notification_topic", tf_ons.OnsNotificationTopicResource())
	RegisterResource("oci_ons_subscription", tf_ons.OnsSubscriptionResource())

	// database_tools service
	RegisterResource("oci_database_tools_database_tools_connection", tf_database_tools.DatabaseToolsDatabaseToolsConnectionResource())
	RegisterResource("oci_database_tools_database_tools_private_endpoint", tf_database_tools.DatabaseToolsDatabaseToolsPrivateEndpointResource())

	// usage proxy service
	RegisterResource("oci_usage_proxy_subscription_redeemable_user", tf_usage_proxy.UsageProxySubscriptionRedeemableUserResource())

	//Object Storage service
	RegisterResource("oci_objectstorage_object_lifecycle_policy", tf_objectstorage.ObjectStorageObjectLifecyclePolicyResource())
	RegisterResource("oci_objectstorage_bucket", tf_objectstorage.ObjectStorageBucketResource())
	RegisterResource("oci_objectstorage_preauthrequest", tf_objectstorage.ObjectStoragePreauthenticatedRequestResource())
	RegisterResource("oci_objectstorage_replication_policy", tf_objectstorage.ObjectStorageReplicationPolicyResource())
	RegisterResource("oci_objectstorage_object", tf_objectstorage.ObjectStorageObjectResource())
	RegisterResource("oci_objectstorage_namespace_metadata", tf_objectstorage.ObjectStorageNamespaceMetadataResource())

	//Databse Service
	RegisterResource("oci_database_autonomous_container_database", tf_database.DatabaseAutonomousContainerDatabaseResource())
	RegisterResource("oci_database_autonomous_container_database_dataguard_association_operation", tf_database.DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource())
	RegisterResource("oci_database_autonomous_database", tf_database.DatabaseAutonomousDatabaseResource())
	RegisterResource("oci_database_autonomous_database_backup", tf_database.DatabaseAutonomousDatabaseBackupResource())
	RegisterResource("oci_database_autonomous_database_instance_wallet_management", tf_database.DatabaseAutonomousDatabaseInstanceWalletManagementResource())
	RegisterResource("oci_database_autonomous_database_regional_wallet_management", tf_database.DatabaseAutonomousDatabaseRegionalWalletManagementResource())
	RegisterResource("oci_database_autonomous_database_wallet", tf_database.DatabaseAutonomousDatabaseWalletResource())
	RegisterResource("oci_database_autonomous_exadata_infrastructure", tf_database.DatabaseAutonomousExadataInfrastructureResource())
	RegisterResource("oci_database_autonomous_vm_cluster", tf_database.DatabaseAutonomousVmClusterResource())
	RegisterResource("oci_database_backup", tf_database.DatabaseBackupResource())
	RegisterResource("oci_database_backup_destination", tf_database.DatabaseBackupDestinationResource())
	RegisterResource("oci_database_cloud_autonomous_vm_cluster", tf_database.DatabaseCloudAutonomousVmClusterResource())
	RegisterResource("oci_database_cloud_database_management", tf_database.DatabaseCloudDatabaseManagementResource())
	RegisterResource("oci_database_cloud_exadata_infrastructure", tf_database.DatabaseCloudExadataInfrastructureResource())
	RegisterResource("oci_database_cloud_vm_cluster", tf_database.DatabaseCloudVmClusterResource())
	RegisterResource("oci_database_data_guard_association", tf_database.DatabaseDataGuardAssociationResource())
	RegisterResource("oci_database_database", tf_database.DatabaseDatabaseResource())
	RegisterResource("oci_database_database_software_image", tf_database.DatabaseDatabaseSoftwareImageResource())
	RegisterResource("oci_database_database_upgrade", tf_database.DatabaseDatabaseUpgradeResource())
	RegisterResource("oci_database_db_home", tf_database.DatabaseDbHomeResource())
	RegisterResource("oci_database_db_node_console_connection", tf_database.DatabaseDbNodeConsoleConnectionResource())
	RegisterResource("oci_database_db_system", tf_database.DatabaseDbSystemResource())
	RegisterResource("oci_database_exadata_infrastructure", tf_database.DatabaseExadataInfrastructureResource())
	RegisterResource("oci_database_exadata_infrastructure", tf_database.DatabaseExadataInfrastructureResource())
	RegisterResource("oci_database_exadata_infrastructure_storage", tf_database.DatabaseExadataInfrastructureStorageResource())
	RegisterResource("oci_database_exadata_iorm_config", tf_database.DatabaseExadataIormConfigResource())
	RegisterResource("oci_database_external_container_database", tf_database.DatabaseExternalContainerDatabaseResource())
	RegisterResource("oci_database_external_container_database_management", tf_database.DatabaseExternalContainerDatabaseManagementResource())
	RegisterResource("oci_database_external_database_connector", tf_database.DatabaseExternalDatabaseConnectorResource())
	RegisterResource("oci_database_external_non_container_database", tf_database.DatabaseExternalNonContainerDatabaseResource())
	RegisterResource("oci_database_external_non_container_database_management", tf_database.DatabaseExternalNonContainerDatabaseManagementResource())
	RegisterResource("oci_database_external_non_container_database_operations_insights_management", tf_database.DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource())
	RegisterResource("oci_database_external_pluggable_database", tf_database.DatabaseExternalPluggableDatabaseResource())
	RegisterResource("oci_database_external_pluggable_database_management", tf_database.DatabaseExternalPluggableDatabaseManagementResource())
	RegisterResource("oci_database_external_pluggable_database_operations_insights_management", tf_database.DatabaseExternalPluggableDatabaseOperationsInsightsManagementResource())
	RegisterResource("oci_database_key_store", tf_database.DatabaseKeyStoreResource())
	RegisterResource("oci_database_maintenance_run", tf_database.DatabaseMaintenanceRunResource())
	RegisterResource("oci_database_pluggable_database", tf_database.DatabasePluggableDatabaseResource())
	RegisterResource("oci_database_pluggable_databases_local_clone", tf_database.DatabasePluggableDatabasesLocalCloneResource())
	RegisterResource("oci_database_pluggable_databases_remote_clone", tf_database.DatabasePluggableDatabasesRemoteCloneResource())
	RegisterResource("oci_database_vm_cluster", tf_database.DatabaseVmClusterResource())
	RegisterResource("oci_database_vm_cluster_add_virtual_machine", tf_database.DatabaseVmClusterAddVirtualMachineResource())
	RegisterResource("oci_database_vm_cluster_network", tf_database.DatabaseVmClusterNetworkResource())
	RegisterResource("oci_database_vm_cluster_remove_virtual_machine", tf_database.DatabaseVmClusterRemoveVirtualMachineResource())

	// containerengine service
	RegisterResource("oci_containerengine_node_pool", tf_containerengine.ContainerengineNodePoolResource())
	RegisterResource("oci_containerengine_cluster", tf_containerengine.ContainerengineClusterResource())

	// database management service
	RegisterResource("oci_database_management_managed_database_group", tf_database_management.DatabaseManagementManagedDatabaseGroupResource())
	RegisterResource("oci_database_management_managed_databases_change_database_parameter", tf_database_management.DatabaseManagementManagedDatabasesChangeDatabaseParameterResource())
	RegisterResource("oci_database_management_managed_databases_reset_database_parameter", tf_database_management.DatabaseManagementManagedDatabasesResetDatabaseParameterResource())
	RegisterResource("oci_database_management_db_management_private_endpoint", tf_database_management.DatabaseManagementDbManagementPrivateEndpointResource())

	// dataflow service
	RegisterResource("oci_dataflow_application", tf_dataflow.DataflowApplicationResource())
	RegisterResource("oci_dataflow_private_endpoint", tf_dataflow.DataflowPrivateEndpointResource())
	RegisterResource("oci_dataflow_invoke_run", tf_dataflow.DataflowInvokeRunResource())

	// integration service
	RegisterResource("oci_integration_integration_instance", tf_integration.IntegrationIntegrationInstanceResource())

	// dataintegration service
	RegisterResource("oci_dataintegration_workspace", tf_dataintegration.DataintegrationWorkspaceResource())

	// osmanagement service
	RegisterResource("oci_osmanagement_managed_instance", tf_osmanagement.OsmanagementManagedInstanceResource())
	RegisterResource("oci_osmanagement_software_source", tf_osmanagement.OsmanagementSoftwareSourceResource())
	RegisterResource("oci_osmanagement_managed_instance_group", tf_osmanagement.OsmanagementManagedInstanceGroupResource())
	RegisterResource("oci_osmanagement_managed_instance_management", tf_osmanagement.OsmanagementManagedInstanceManagementResource())
}
