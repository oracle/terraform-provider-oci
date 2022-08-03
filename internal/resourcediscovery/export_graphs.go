// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

var tenancyResourceGraphs = map[string]TerraformResourceGraph{
	"budget":               budgetResourceGraph,
	"email_tenancy":        emailTenancyResourceGraph,
	"cloud_guard_tenancy":  cloudGuardTenancyResourceGraph,
	"identity":             identityResourceGraph,
	"limits":               limitsResourceGraph,
	"metering_computation": meteringComputationResourceGraph,
	"optimizer":            optimizerResourceGraph,
	"usage_proxy":          usageProxyResourceGraph,
}

var compartmentResourceGraphs = map[string]TerraformResourceGraph{
	"adm":                     admResourceGraph,
	"ai_anomaly_detection":    aiAnomalyDetectionResourceGraph,
	"ai_vision":               aiVisionResourceGraph,
	"analytics":               analyticsResourceGraph,
	"announcements_service":   announcementsServiceResourceGraph,
	"apigateway":              apigatewayResourceGraph,
	"apm":                     apmResourceGraph,
	"apm_config":              apmConfigResourceGraph,
	"apm_synthetics":          apmSyntheticsResourceGraph,
	"artifacts":               artifactsResourceGraph,
	"auto_scaling":            autoScalingResourceGraph,
	"availability_domain":     availabilityDomainResourceGraph,
	"bastion":                 bastionResourceGraph,
	"bds":                     bdsResourceGraph,
	"blockchain":              blockchainResourceGraph,
	"certificates_management": certificatesManagementResourceGraph,
	"cloud_guard":             cloudGuardResourceGraph,
	"containerengine":         containerengineResourceGraph,
	"core":                    coreResourceGraph,
	"data_connectivity":       dataConnectivityResourceGraph,
	"data_labeling_service":   dataLabelingServiceResourceGraph,
	"data_safe":               dataSafeResourceGraph,
	"database":                databaseResourceGraph,
	"database_migration":      databaseMigrationResourceGraph,
	"database_tools":          databaseToolsResourceGraph,
	"datacatalog":             datacatalogResourceGraph,
	"dataflow":                dataflowResourceGraph,
	"dataintegration":         dataintegrationResourceGraph,
	"datascience":             datascienceResourceGraph,
	"devops":                  devopsResourceGraph,
	"dns":                     dnsResourceGraph,
	"em_warehouse":            emWarehouseResourceGraph,
	"email":                   emailResourceGraph,
	"events":                  eventsResourceGraph,
	"file_storage":            fileStorageResourceGraph,
	"functions":               functionsResourceGraph,
	"golden_gate":             goldenGateResourceGraph,
	"health_checks":           healthChecksResourceGraph,
	"identity_data_plane":     identityDataPlaneResourceGraph,
	"integration":             integrationResourceGraph,
	"jms":                     jmsResourceGraph,
	"kms":                     kmsResourceGraph,
	"license_manager":         licenseManagerResourceGraph,
	"load_balancer":           loadBalancerResourceGraph,
	"log_analytics":           logAnalyticsResourceGraph,
	"logging":                 loggingResourceGraph,
	"management_agent":        managementAgentResourceGraph,
	"marketplace":             marketplaceResourceGraph,
	"monitoring":              monitoringResourceGraph,
	"mysql":                   mysqlResourceGraph,
	"network_firewall":        networkFirewallResourceGraph,
	"network_load_balancer":   networkLoadBalancerResourceGraph,
	"nosql":                   nosqlResourceGraph,
	"object_storage":          objectStorageResourceGraph,
	"oce":                     oceResourceGraph,
	"ocvp":                    ocvpResourceGraph,
	"oda":                     odaResourceGraph,
	"ons":                     onsResourceGraph,
	"operator_access_control": operatorAccessControlResourceGraph,
	"opsi":                    opsiResourceGraph,
	"osmanagement":            osmanagementResourceGraph,
	"osp_gateway":             ospGatewayResourceGraph,
	"resourcemanager":         resourcemanagerResourceGraph,
	"sch":                     schResourceGraph,
	"service_mesh":            serviceMeshResourceGraph,
	"stack_monitoring":        stackMonitoringResourceGraph,
	"streaming":               streamingResourceGraph,
	"tagging":                 taggingResourceGraph,
	"vault":                   vaultResourceGraph,
	"visual_builder":          visualBuilderResourceGraph,
	"vn_monitoring":           vnMonitoringResourceGraph,
	"vulnerability_scanning":  vulnerabilityScanningResourceGraph,
	"waa":                     waaResourceGraph,
	"waas":                    waasResourceGraph,
	"waf":                     wafResourceGraph,
}

var availabilityDomainResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportIdentityAvailabilityDomainHints,
		},
	},
	"oci_identity_availability_domain": {
		{
			TerraformResourceHints: exportCoreBootVolumeHints,
			datasourceQueryParams: map[string]string{
				"availability_domain": "name",
			},
		},
		{
			TerraformResourceHints: exportFileStorageFileSystemHints,
			datasourceQueryParams: map[string]string{
				"availability_domain": "name",
			},
		},
		{
			TerraformResourceHints: exportFileStorageMountTargetHints,
			datasourceQueryParams: map[string]string{
				"availability_domain": "name",
			},
		},
	},
	"oci_file_storage_file_system": {
		{
			TerraformResourceHints: exportFileStorageSnapshotHints,
			datasourceQueryParams: map[string]string{
				"file_system_id": "id",
			},
		},
	},
}

var admResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAdmVulnerabilityAuditHints},
		{TerraformResourceHints: exportAdmKnowledgeBaseHints},
	},
}

var aiAnomalyDetectionResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAiAnomalyDetectionDataAssetHints},
		{TerraformResourceHints: exportAiAnomalyDetectionModelHints},
		{TerraformResourceHints: exportAiAnomalyDetectionProjectHints},
		{TerraformResourceHints: exportAiAnomalyDetectionAiPrivateEndpointHints},
	},
}

var aiVisionResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAiVisionProjectHints},
		{TerraformResourceHints: exportAiVisionModelHints},
	},
	"oci_ai_vision_document_job": {
		{
			TerraformResourceHints: exportAiVisionDocumentJobHints,
			datasourceQueryParams: map[string]string{
				"document_job_id": "id",
			},
		},
	},
	"oci_ai_vision_image_job": {
		{
			TerraformResourceHints: exportAiVisionImageJobHints,
			datasourceQueryParams: map[string]string{
				"image_job_id": "id",
			},
		},
	},
}

var analyticsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAnalyticsAnalyticsInstanceHints},
	},
}

var announcementsServiceResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAnnouncementsServiceAnnouncementSubscriptionHints},
	},
}

var apigatewayResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportApigatewayApiHints},
		{TerraformResourceHints: exportApigatewayGatewayHints},
		{TerraformResourceHints: exportApigatewayDeploymentHints},
		{TerraformResourceHints: exportApigatewayCertificateHints},
		{TerraformResourceHints: exportApigatewaySubscriberHints},
		{TerraformResourceHints: exportApigatewayUsagePlanHints},
	},
}

var apmResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportApmApmDomainHints},
	},
}

var apmConfigResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {},
}

var apmSyntheticsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {},
}
var artifactsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportArtifactsContainerRepositoryHints},
		{TerraformResourceHints: exportArtifactsContainerImageSignatureHints},
		{TerraformResourceHints: exportArtifactsRepositoryHints},
	},
}

var autoScalingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAutoScalingAutoScalingConfigurationHints},
	},
}

var bastionResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBastionBastionHints},
	},
	"oci_bastion_bastion": {
		{
			TerraformResourceHints: exportBastionSessionHints,
			datasourceQueryParams: map[string]string{
				"bastion_id": "id",
			},
		},
	},
}

var bdsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBdsBdsInstanceHints},
	},
	"oci_bds_bds_instance": {
		{
			TerraformResourceHints: exportBdsBdsInstanceApiKeyHints,
			datasourceQueryParams: map[string]string{
				"bds_instance_id": "id",
			},
		},
		{
			TerraformResourceHints: exportBdsBdsInstanceMetastoreConfigHints,
			datasourceQueryParams: map[string]string{
				"bds_instance_id": "id",
			},
		},
	},
}

var blockchainResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBlockchainBlockchainPlatformHints},
	},
	"oci_blockchain_blockchain_platform": {
		{
			TerraformResourceHints: exportBlockchainOsnHints,
			datasourceQueryParams: map[string]string{
				"blockchain_platform_id": "id",
			},
		},
		{
			TerraformResourceHints: exportBlockchainPeerHints,
			datasourceQueryParams: map[string]string{
				"blockchain_platform_id": "id",
			},
		},
	},
}

var budgetResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{
			TerraformResourceHints: exportBudgetBudgetHints,
			datasourceQueryParams: map[string]string{
				"target_type": "'ALL'",
			},
		},
	},
	"oci_budget_budget": {
		{
			TerraformResourceHints: exportBudgetAlertRuleHints,
			datasourceQueryParams: map[string]string{
				"budget_id": "id",
			},
		},
	},
}

var certificatesManagementResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCertificatesManagementCaBundleHints},
		{TerraformResourceHints: exportCertificatesManagementCertificateAuthorityHints},
		{TerraformResourceHints: exportCertificatesManagementCertificateHints},
	},
}

var cloudGuardResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCloudGuardTargetHints},
		{TerraformResourceHints: exportCloudGuardManagedListHints},
		{TerraformResourceHints: exportCloudGuardResponderRecipeHints},
		{TerraformResourceHints: exportCloudGuardDetectorRecipeHints},
		{TerraformResourceHints: exportCloudGuardSecurityRecipeHints},
		{TerraformResourceHints: exportCloudGuardSecurityZoneHints},
	},
}

var cloudGuardTenancyResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportCloudGuardDataMaskRuleHints},
	},
}

var containerengineResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportContainerengineClusterHints},
		{TerraformResourceHints: exportContainerengineNodePoolHints},
	},
}

var coreResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCoreBootVolumeBackupHints},
		{TerraformResourceHints: exportCoreBootVolumeHints},
		{TerraformResourceHints: exportCoreConsoleHistoryHints},
		{TerraformResourceHints: exportCoreClusterNetworkHints},
		{TerraformResourceHints: exportCoreComputeImageCapabilitySchemaHints},
		{TerraformResourceHints: exportCoreCpeHints},
		{TerraformResourceHints: exportCoreCrossConnectGroupHints},
		{TerraformResourceHints: exportCoreCrossConnectHints},
		{TerraformResourceHints: exportCoreDrgAttachmentHints},
		{TerraformResourceHints: exportCoreDrgHints},
		{TerraformResourceHints: exportCoreDedicatedVmHostHints},
		{TerraformResourceHints: exportCoreImageHints},
		{TerraformResourceHints: exportCoreInstanceConfigurationHints},
		{TerraformResourceHints: exportCoreInstanceConsoleConnectionHints},
		{TerraformResourceHints: exportCoreInstancePoolHints},
		{TerraformResourceHints: exportCoreInstanceHints},
		{TerraformResourceHints: exportCoreInternetGatewayHints},
		{TerraformResourceHints: exportCoreIpSecConnectionHints},
		{TerraformResourceHints: exportCoreLocalPeeringGatewayHints},
		{TerraformResourceHints: exportCoreNetworkSecurityGroupHints},
		{
			TerraformResourceHints: exportCorePublicIpHints,
			datasourceQueryParams: map[string]string{
				"scope": "'REGION'",
			},
		},
		{TerraformResourceHints: exportCoreRemotePeeringConnectionHints},
		{TerraformResourceHints: exportCoreServiceGatewayHints},
		{TerraformResourceHints: exportCoreSubnetHints},
		{TerraformResourceHints: exportCoreVcnHints},
		{TerraformResourceHints: exportCoreVirtualCircuitHints},
		{TerraformResourceHints: exportCoreVolumeAttachmentHints},
		{TerraformResourceHints: exportCoreVolumeBackupHints},
		{TerraformResourceHints: exportCoreVolumeBackupPolicyHints},
		{TerraformResourceHints: exportCoreVolumeGroupHints},
		{TerraformResourceHints: exportCoreVolumeGroupBackupHints},
		{TerraformResourceHints: exportCoreVolumeHints},
		{TerraformResourceHints: exportCorePublicIpPoolHints},
		{TerraformResourceHints: exportCoreCaptureFilterHints},
		{TerraformResourceHints: exportCoreVtapHints},
	},
	"oci_core_boot_volume": {
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			datasourceQueryParams: map[string]string{
				"asset_id": "id",
			},
		},
	},
	"oci_core_drg": {
		{
			TerraformResourceHints: exportCoreDrgRouteDistributionHints,
			datasourceQueryParams: map[string]string{
				"drg_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreDrgRouteTableHints,
			datasourceQueryParams: map[string]string{
				"drg_id": "id",
			},
		},
	},
	"oci_core_instance_pool": {
		{
			TerraformResourceHints: exportCoreInstancePoolInstanceHints,
			datasourceQueryParams: map[string]string{
				"instance_pool_id": "id",
			},
		},
	},
	"oci_core_instance": {
		{
			TerraformResourceHints: exportCoreVnicAttachmentHints,
			datasourceQueryParams: map[string]string{
				"instance_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			datasourceQueryParams: map[string]string{
				"asset_id": "boot_volume_id",
			},
		},
	},
	"oci_core_network_security_group": {
		{
			TerraformResourceHints: exportCoreNetworkSecurityGroupSecurityRuleHints,
			datasourceQueryParams: map[string]string{
				"network_security_group_id": "id",
			},
		},
	},
	"oci_core_subnet": {
		{
			TerraformResourceHints: exportCorePrivateIpHints,
			datasourceQueryParams: map[string]string{
				"subnet_id": "id",
			},
		},
	},
	"oci_core_vcn": {
		{
			TerraformResourceHints: exportCoreDhcpOptionsHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreNatGatewayHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreRouteTableHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreSecurityListHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreVlanHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
	},
	"oci_core_volume": {
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			datasourceQueryParams: map[string]string{
				"asset_id": "id",
			},
		},
	},
	"oci_core_drg_route_table": {
		{
			TerraformResourceHints: exportCoreDrgRouteTableRouteRuleHints,
			datasourceQueryParams: map[string]string{
				"drg_route_table_id": "id",
			},
		},
	},
}

var dataConnectivityResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataConnectivityRegistryHints},
	},
	"oci_data_connectivity_registry": {
		{
			TerraformResourceHints: exportDataConnectivityRegistryConnectionHints,
			datasourceQueryParams: map[string]string{
				"data_asset_key": "data_asset_key",
				"registry_id":    "id",
			},
		},
		{
			TerraformResourceHints: exportDataConnectivityRegistryDataAssetHints,
			datasourceQueryParams: map[string]string{
				"registry_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDataConnectivityRegistryFolderHints,
			datasourceQueryParams: map[string]string{
				"registry_id": "id",
			},
		},
	},
	"oci_data_connectivity_registry_data_asset": {
		{
			TerraformResourceHints: exportDataConnectivityRegistryConnectionHints,
			datasourceQueryParams: map[string]string{
				"data_asset_key": "key",
				"registry_id":    "registry_id",
			},
		},
	},
}

var dataLabelingServiceResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataLabelingServiceDatasetHints},
	},
}

var dataSafeResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataSafeDataSafePrivateEndpointHints},
		{TerraformResourceHints: exportDataSafeOnPremConnectorHints},
		{TerraformResourceHints: exportDataSafeTargetDatabaseHints},
		{TerraformResourceHints: exportDataSafeSecurityAssessmentHints},
		{TerraformResourceHints: exportDataSafeUserAssessmentHints},
		{TerraformResourceHints: exportDataSafeReportDefinitionHints},
		{TerraformResourceHints: exportDataSafeAuditTrailHints},
		{TerraformResourceHints: exportDataSafeAlertHints},
		{TerraformResourceHints: exportDataSafeAuditArchiveRetrievalHints},
		{TerraformResourceHints: exportDataSafeAuditProfileHints},
		{TerraformResourceHints: exportDataSafeAuditPolicyHints},
		{TerraformResourceHints: exportDataSafeTargetAlertPolicyAssociationHints},
		{TerraformResourceHints: exportDataSafeSensitiveTypeHints},
		{TerraformResourceHints: exportDataSafeMaskingPolicyHints},
		{TerraformResourceHints: exportDataSafeLibraryMaskingFormatHints},
		{TerraformResourceHints: exportDataSafeSensitiveDataModelHints},
		{TerraformResourceHints: exportDataSafeDiscoveryJobHints},
	},
	"oci_data_safe_masking_policy": {
		{
			TerraformResourceHints: exportDataSafeMaskingPoliciesMaskingColumnHints,
			datasourceQueryParams: map[string]string{
				"masking_policy_id": "id",
			},
		},
	},
	"oci_data_safe_sensitive_data_model": {
		{
			TerraformResourceHints: exportDataSafeSensitiveDataModelsSensitiveColumnHints,
			datasourceQueryParams: map[string]string{
				"sensitive_data_model_id": "id",
			},
		},
	},
	"oci_data_safe_discovery_job": {
		{
			TerraformResourceHints: exportDataSafeDiscoveryJobsResultHints,
			datasourceQueryParams: map[string]string{
				"discovery_job_id": "id",
			},
		},
	},
}

var databaseResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseAutonomousContainerDatabaseHints},
		{TerraformResourceHints: exportDatabaseAutonomousDatabaseHints},
		{TerraformResourceHints: exportDatabaseAutonomousExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseAutonomousVmClusterHints},
		{TerraformResourceHints: exportDatabaseBackupDestinationHints},
		{TerraformResourceHints: exportDatabaseBackupHints},
		{TerraformResourceHints: exportDatabaseDbSystemHints},
		{TerraformResourceHints: exportDatabaseExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseVmClusterHints},
		{TerraformResourceHints: exportDatabaseDatabaseSoftwareImageHints},
		{TerraformResourceHints: exportDatabaseCloudExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseCloudVmClusterHints},
		{TerraformResourceHints: exportDatabaseKeyStoreHints},
		{TerraformResourceHints: exportDatabaseExternalContainerDatabaseHints},
		{TerraformResourceHints: exportDatabaseExternalPluggableDatabaseHints},
		{TerraformResourceHints: exportDatabaseExternalNonContainerDatabaseHints},
		{TerraformResourceHints: exportDatabasePluggableDatabaseHints},
		{TerraformResourceHints: exportDatabaseCloudAutonomousVmClusterHints},
	},
	"oci_database_autonomous_container_database": {
		{
			TerraformResourceHints: exportDatabaseAutonomousContainerDatabaseDataguardAssociationHints,
			datasourceQueryParams: map[string]string{
				"autonomous_container_database_id": "id",
			},
		},
	},
	"oci_database_db_home": {
		{
			TerraformResourceHints: exportDatabaseDatabaseHints,
			datasourceQueryParams: map[string]string{
				"db_home_id": "id",
			},
		},
	},
	"oci_database_db_system": {
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			datasourceQueryParams: map[string]string{
				"db_system_id": "id",
			},
		},
	},
	"oci_database_exadata_infrastructure": {
		{
			TerraformResourceHints: exportDatabaseVmClusterNetworkHints,
			datasourceQueryParams: map[string]string{
				"exadata_infrastructure_id": "id",
			},
		},
	},
	"oci_database_vm_cluster": {
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			datasourceQueryParams: map[string]string{
				"vm_cluster_id": "id",
			},
		},
	},
	"oci_database_cloud_vm_cluster": {
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			datasourceQueryParams: map[string]string{
				"vm_cluster_id": "id",
			},
		},
	},
}

var databaseMigrationResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseMigrationMigrationHints},
		{TerraformResourceHints: exportDatabaseMigrationConnectionHints},
	},
}

var databaseToolsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsPrivateEndpointHints},
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsConnectionHints},
	},
}

var datacatalogResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatacatalogCatalogHints},
		{TerraformResourceHints: exportDatacatalogCatalogPrivateEndpointHints},
		{TerraformResourceHints: exportDatacatalogMetastoreHints},
	},
	"oci_datacatalog_catalog": {
		{
			TerraformResourceHints: exportDatacatalogConnectionHints,
			datasourceQueryParams: map[string]string{
				"catalog_id":     "id",
				"data_asset_key": "data_asset_key",
			},
		},
		{
			TerraformResourceHints: exportDatacatalogDataAssetHints,
			datasourceQueryParams: map[string]string{
				"catalog_id": "id",
			},
		},
	},
	"oci_datacatalog_data_asset": {
		{
			TerraformResourceHints: exportDatacatalogConnectionHints,
			datasourceQueryParams: map[string]string{
				"data_asset_key": "key",
				"catalog_id":     "catalog_id",
			},
		},
	},
}

var dataflowResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataflowApplicationHints},
		{TerraformResourceHints: exportDataflowPrivateEndpointHints},
	},
}

var dataintegrationResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataintegrationWorkspaceHints},
	},
}

var datascienceResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatascienceProjectHints},
		{TerraformResourceHints: exportDatascienceNotebookSessionHints},
		{TerraformResourceHints: exportDatascienceModelHints},
		{TerraformResourceHints: exportDatascienceModelDeploymentHints},
		{TerraformResourceHints: exportDatascienceJobHints},
		{TerraformResourceHints: exportDatascienceJobRunHints},
	},
	"oci_datascience_model": {
		{
			TerraformResourceHints: exportDatascienceModelProvenanceHints,
			datasourceQueryParams: map[string]string{
				"model_id": "id",
			},
		},
	},
}

var devopsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDevopsProjectHints},
		{TerraformResourceHints: exportDevopsDeployEnvironmentHints},
		{TerraformResourceHints: exportDevopsDeployArtifactHints},
		{TerraformResourceHints: exportDevopsDeployPipelineHints},
		{TerraformResourceHints: exportDevopsDeployStageHints},
		{TerraformResourceHints: exportDevopsDeploymentHints},
		{TerraformResourceHints: exportDevopsRepositoryHints},
		{TerraformResourceHints: exportDevopsBuildPipelineHints},
		{TerraformResourceHints: exportDevopsBuildRunHints},
		{TerraformResourceHints: exportDevopsConnectionHints},
		{TerraformResourceHints: exportDevopsBuildPipelineStageHints},
		{TerraformResourceHints: exportDevopsTriggerHints},
	},
	"oci_devops_repository": {
		{
			TerraformResourceHints: exportDevopsRepositoryRefHints,
			datasourceQueryParams: map[string]string{
				"repository_id": "id",
			},
		},
	},
}

var dnsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDnsZoneHints},
		{TerraformResourceHints: exportDnsSteeringPolicyHints},
		{TerraformResourceHints: exportDnsSteeringPolicyAttachmentHints},
		{TerraformResourceHints: exportDnsTsigKeyHints},
	},
	"oci_dns_zone": {
		{
			TerraformResourceHints: exportDnsRrsetHints,
			datasourceQueryParams: map[string]string{
				"zone_name_or_id": "id",
			},
		},
	},
}

var emWarehouseResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportEmWarehouseEmWarehouseHints},
	},
}

var emailResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportEmailSenderHints},
		{TerraformResourceHints: exportEmailEmailDomainHints},
	},
	"oci_email_email_domain": {
		{
			TerraformResourceHints: exportEmailDkimHints,
			datasourceQueryParams: map[string]string{
				"email_domain_id": "id",
			},
		},
	},
}

var emailTenancyResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportEmailSuppressionHints},
	},
}

var eventsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportEventsRuleHints},
	},
}

var fileStorageResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFileStorageExportHints},
	},
}

var functionsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFunctionsApplicationHints},
	},
	"oci_functions_application": {
		{
			TerraformResourceHints: exportFunctionsFunctionHints,
			datasourceQueryParams: map[string]string{
				"application_id": "id",
			},
		},
	},
}

var goldenGateResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportGoldenGateDatabaseRegistrationHints},
		{TerraformResourceHints: exportGoldenGateDeploymentHints},
		{TerraformResourceHints: exportGoldenGateDeploymentBackupHints},
	},
}

var healthChecksResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportHealthChecksHttpMonitorHints},
		{TerraformResourceHints: exportHealthChecksPingMonitorHints},
	},
}

var identityResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportIdentityAuthenticationPolicyHints},
		{TerraformResourceHints: exportIdentityCompartmentHints},
		{TerraformResourceHints: exportIdentityDynamicGroupHints},
		{TerraformResourceHints: exportIdentityGroupHints},
		{
			TerraformResourceHints: exportIdentityIdentityProviderHints,
			datasourceQueryParams:  map[string]string{"protocol": "'SAML2'"},
		},
		{TerraformResourceHints: exportIdentityPolicyHints},
		{TerraformResourceHints: exportIdentityUserHints},
		{TerraformResourceHints: exportIdentityNetworkSourceHints},
		{TerraformResourceHints: exportIdentityDomainHints},
	},
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportIdentityCompartmentHints,
			datasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
		{
			TerraformResourceHints: exportIdentityPolicyHints,
			datasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
	},
	"oci_identity_identity_provider": {
		{
			TerraformResourceHints: exportIdentityIdpGroupMappingHints,
			datasourceQueryParams: map[string]string{
				"identity_provider_id": "id",
			},
		},
	},
	"oci_identity_user": {
		{
			TerraformResourceHints: exportIdentityApiKeyHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityAuthTokenHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityCustomerSecretKeyHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityDbCredentialHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentitySmtpCredentialHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentitySwiftPasswordHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityUiPasswordHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityUserGroupMembershipHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
	},
}
var identityDataPlaneResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {},
}

var integrationResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportIntegrationIntegrationInstanceHints},
	},
}

var jmsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportJmsFleetHints},
	},
}

var kmsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportKmsVaultHints},
	},
	"oci_kms_key": {
		{
			TerraformResourceHints: exportKmsKeyVersionHints,
			datasourceQueryParams: map[string]string{
				"key_id":              "id",
				"management_endpoint": "management_endpoint",
			},
		},
	},
	"oci_kms_vault": {
		{
			TerraformResourceHints: exportKmsKeyHints,
			datasourceQueryParams: map[string]string{
				"management_endpoint": "management_endpoint",
			},
		},
	},
}

var licenseManagerResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLicenseManagerConfigurationHints},
		{TerraformResourceHints: exportLicenseManagerProductLicenseHints},
	},
	"oci_license_manager_product_license": {
		{
			TerraformResourceHints: exportLicenseManagerLicenseRecordHints,
			datasourceQueryParams: map[string]string{
				"product_license_id": "id",
			},
		},
	},
}

var limitsResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportLimitsQuotaHints},
	},
}

var loadBalancerResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLoadBalancerLoadBalancerHints},
	},
	"oci_load_balancer_backend_set": {
		{
			TerraformResourceHints: exportLoadBalancerBackendHints,
			datasourceQueryParams: map[string]string{
				"backendset_name":  "name",
				"load_balancer_id": "load_balancer_id",
			},
		},
		{TerraformResourceHints: exportLoadBalancerListenerHints},
	},
	"oci_load_balancer_load_balancer": {
		// certificates have to be discovered before listeners in order to populate
		// the references for certificate_name in listeners (dependency)
		// If moving to parallel execution in future, this dependency needs to be maintained
		{
			TerraformResourceHints: exportLoadBalancerCertificateHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerBackendSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerHostnameHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerLoadBalancerRoutingPolicyHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerPathRouteSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerRuleSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
	},
}

var logAnalyticsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints},
	},
}

var loggingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLoggingLogGroupHints},
		{TerraformResourceHints: exportLoggingUnifiedAgentConfigurationHints},
	},
	"oci_logging_log_group": {
		{
			TerraformResourceHints: exportLoggingLogHints,
			datasourceQueryParams: map[string]string{
				"log_group_id": "id",
			},
		},
	},
}

var managementAgentResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportManagementAgentManagementAgentHints},
		{TerraformResourceHints: exportManagementAgentManagementAgentInstallKeyHints},
	},
}

var marketplaceResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMarketplaceAcceptedAgreementHints},
	},
}

var meteringComputationResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportMeteringComputationQueryHints},
		{TerraformResourceHints: exportMeteringComputationScheduleHints},
	},
}

var monitoringResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMonitoringAlarmHints},
	},
}

var mysqlResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMysqlMysqlBackupHints},
		{TerraformResourceHints: exportMysqlMysqlDbSystemHints},
		{TerraformResourceHints: exportMysqlChannelHints},
	},
}

var networkFirewallResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyHints},
		{TerraformResourceHints: exportNetworkFirewallNetworkFirewallHints},
	},
}

var networkLoadBalancerResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportNetworkLoadBalancerNetworkLoadBalancerHints},
	},
	"oci_network_load_balancer_network_load_balancer": {
		{
			TerraformResourceHints: exportNetworkLoadBalancerBackendSetHints,
			datasourceQueryParams: map[string]string{
				"network_load_balancer_id": "id",
			},
		},
	},
	"oci_network_load_balancer_backend_set": {
		{
			TerraformResourceHints: exportNetworkLoadBalancerBackendHints,
			datasourceQueryParams: map[string]string{
				"backend_set_name":         "name",
				"network_load_balancer_id": "network_load_balancer_id",
			},
		},
		{
			TerraformResourceHints: exportNetworkLoadBalancerListenerHints,
			datasourceQueryParams: map[string]string{
				"network_load_balancer_id": "network_load_balancer_id",
			},
		},
	},
}
var nosqlResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportNosqlTableHints},
	},
	"oci_nosql_table": {
		{
			TerraformResourceHints: exportNosqlIndexHints,
			datasourceQueryParams: map[string]string{
				"table_name_or_id": "id",
			},
		},
	},
}

var objectStorageResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportObjectStorageNamespaceHints},
	},
	"oci_objectstorage_bucket": {
		{
			TerraformResourceHints: exportObjectStorageObjectHints,
			datasourceQueryParams: map[string]string{
				"bucket":    "name",
				"namespace": "namespace",
			},
		},
		{
			TerraformResourceHints: exportObjectStorageObjectLifecyclePolicyHints,
			datasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
		{
			TerraformResourceHints: exportObjectStoragePreauthenticatedRequestHints,
			datasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
		{
			TerraformResourceHints: exportObjectStorageReplicationPolicyHints,
			datasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
	},
	"oci_objectstorage_namespace": {
		{
			TerraformResourceHints: exportObjectStorageBucketHints,
			datasourceQueryParams: map[string]string{
				"namespace": "namespace",
			},
		},
	},
}

var oceResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOceOceInstanceHints},
	},
}

var ocvpResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOcvpSddcHints},
	},
	"oci_ocvp_sddc": {
		{
			TerraformResourceHints: exportOcvpEsxiHostHints,
			datasourceQueryParams: map[string]string{
				"sddc_id": "id",
			},
		},
	},
}

var odaResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOdaOdaInstanceHints},
	},
}

var onsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOnsNotificationTopicHints},
		{TerraformResourceHints: exportOnsSubscriptionHints},
	},
}

var operatorAccessControlResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOperatorAccessControlOperatorControlHints},
		{TerraformResourceHints: exportOperatorAccessControlOperatorControlAssignmentHints},
	},
}

var opsiResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOpsiEnterpriseManagerBridgeHints},
		{TerraformResourceHints: exportOpsiDatabaseInsightHints},
		{TerraformResourceHints: exportOpsiHostInsightHints},
		{TerraformResourceHints: exportOpsiExadataInsightHints},
		{TerraformResourceHints: exportOpsiOperationsInsightsWarehouseHints},
		{TerraformResourceHints: exportOpsiOperationsInsightsPrivateEndpointHints},
	},
	"oci_opsi_operations_insights_warehouse": {
		{
			TerraformResourceHints: exportOpsiAwrHubHints,
			datasourceQueryParams: map[string]string{
				"operations_insights_warehouse_id": "id",
			},
		},
		{
			TerraformResourceHints: exportOpsiOperationsInsightsWarehouseUserHints,
			datasourceQueryParams: map[string]string{
				"operations_insights_warehouse_id": "id",
			},
		},
	},
}
var optimizerResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportOptimizerProfileHints},
	},
}

var osmanagementResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOsmanagementManagedInstanceHints},
		{TerraformResourceHints: exportOsmanagementManagedInstanceGroupHints},
		{TerraformResourceHints: exportOsmanagementSoftwareSourceHints},
	},
}

var ospGatewayResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {},
}

var resourcemanagerResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportResourcemanagerPrivateEndpointHints},
	},
}

var schResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportSchServiceConnectorHints},
	},
}

var serviceMeshResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportServiceMeshVirtualServiceHints},
		{TerraformResourceHints: exportServiceMeshAccessPolicyHints},
		{TerraformResourceHints: exportServiceMeshMeshHints},
		{TerraformResourceHints: exportServiceMeshIngressGatewayRouteTableHints},
		{TerraformResourceHints: exportServiceMeshVirtualServiceRouteTableHints},
		{TerraformResourceHints: exportServiceMeshVirtualDeploymentHints},
		{TerraformResourceHints: exportServiceMeshIngressGatewayHints},
	},
}

var stackMonitoringResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportStackMonitoringDiscoveryJobHints},
	},
	"oci_stack_monitoring_monitored_resource": {
		{
			TerraformResourceHints: exportStackMonitoringMonitoredResourceHints,
			datasourceQueryParams: map[string]string{
				"monitored_resource_id": "id",
			},
		},
	},
}

var streamingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportStreamingConnectHarnessHints},
		{TerraformResourceHints: exportStreamingStreamPoolHints},
		{TerraformResourceHints: exportStreamingStreamHints},
	},
}

var usageProxyResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {},
}

var vaultResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportVaultSecretHints},
	},
}

var visualBuilderResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportVisualBuilderVbInstanceHints},
	},
}

var vnMonitoringResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportVnMonitoringPathAnalyzerTestHints},
	},
}

var vulnerabilityScanningResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportVulnerabilityScanningHostScanRecipeHints},
		{TerraformResourceHints: exportVulnerabilityScanningHostScanTargetHints},
		{TerraformResourceHints: exportVulnerabilityScanningContainerScanRecipeHints},
		{TerraformResourceHints: exportVulnerabilityScanningContainerScanTargetHints},
	},
}

var waaResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportWaaWebAppAccelerationPolicyHints},
		{TerraformResourceHints: exportWaaWebAppAccelerationHints},
	},
}

var waasResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportWaasAddressListHints},
		{TerraformResourceHints: exportWaasCustomProtectionRuleHints},
		{TerraformResourceHints: exportWaasHttpRedirectHints},
		{TerraformResourceHints: exportWaasWaasPolicyHints},
	},
}

var wafResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportWafWebAppFirewallPolicyHints},
		{TerraformResourceHints: exportWafWebAppFirewallHints},
		{TerraformResourceHints: exportWafNetworkAddressListHints},
	},
}

var taggingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportIdentityTagDefaultHints},
		{TerraformResourceHints: exportIdentityTagNamespaceHints},
	},
	"oci_identity_tag_namespace": {
		{
			TerraformResourceHints: exportIdentityTagHints,
			datasourceQueryParams: map[string]string{
				"tag_namespace_id": "id",
			},
		},
	},
}
