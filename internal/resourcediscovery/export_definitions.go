// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v56/aianomalydetection"
	oci_analytics "github.com/oracle/oci-go-sdk/v56/analytics"
	oci_apigateway "github.com/oracle/oci-go-sdk/v56/apigateway"
	oci_apm "github.com/oracle/oci-go-sdk/v56/apmcontrolplane"
	oci_artifacts "github.com/oracle/oci-go-sdk/v56/artifacts"
	oci_bds "github.com/oracle/oci-go-sdk/v56/bds"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v56/certificatesmanagement"
	oci_containerengine "github.com/oracle/oci-go-sdk/v56/containerengine"
	oci_database_migration "github.com/oracle/oci-go-sdk/v56/databasemigration"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v56/datacatalog"
	oci_dataflow "github.com/oracle/oci-go-sdk/v56/dataflow"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v56/dataintegration"
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v56/datalabelingservice"
	oci_data_safe "github.com/oracle/oci-go-sdk/v56/datasafe"
	oci_datascience "github.com/oracle/oci-go-sdk/v56/datascience"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"
	oci_dns "github.com/oracle/oci-go-sdk/v56/dns"
	oci_file_storage "github.com/oracle/oci-go-sdk/v56/filestorage"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v56/goldengate"
	oci_integration "github.com/oracle/oci-go-sdk/v56/integration"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v56/loganalytics"
	oci_logging "github.com/oracle/oci-go-sdk/v56/logging"
	oci_management_agent "github.com/oracle/oci-go-sdk/v56/managementagent"
	oci_marketplace "github.com/oracle/oci-go-sdk/v56/marketplace"
	oci_monitoring "github.com/oracle/oci-go-sdk/v56/monitoring"
	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v56/networkloadbalancer"
	oci_nosql "github.com/oracle/oci-go-sdk/v56/nosql"
	oci_oce "github.com/oracle/oci-go-sdk/v56/oce"
	oci_ons "github.com/oracle/oci-go-sdk/v56/ons"
	oci_opsi "github.com/oracle/oci-go-sdk/v56/opsi"
	oci_sch "github.com/oracle/oci-go-sdk/v56/sch"
	oci_streaming "github.com/oracle/oci-go-sdk/v56/streaming"
	oci_visual_builder "github.com/oracle/oci-go-sdk/v56/visualbuilder"
	oci_vulnerability_scanning "github.com/oracle/oci-go-sdk/v56/vulnerabilityscanning"
	oci_waas "github.com/oracle/oci-go-sdk/v56/waas"
	oci_waf "github.com/oracle/oci-go-sdk/v56/waf"

	//oci_artifacts "github.com/oracle/oci-go-sdk/v56/artifacts"
	//oci_bastion "github.com/oracle/oci-go-sdk/v56/bastion"
	//oci_bds "github.com/oracle/oci-go-sdk/v56/bds"
	oci_blockchain "github.com/oracle/oci-go-sdk/v56/blockchain"

	//oci_dataflow "github.com/oracle/oci-go-sdk/v56/dataflow"
	//oci_dataintegration "github.com/oracle/oci-go-sdk/v56/dataintegration"
	//oci_data_labeling_service "github.com/oracle/oci-go-sdk/v56/datalabelingservice"
	//oci_data_safe "github.com/oracle/oci-go-sdk/v56/datasafe"
	//oci_datascience "github.com/oracle/oci-go-sdk/v56/datascience"
	//oci_devops "github.com/oracle/oci-go-sdk/v56/devops"
	//oci_dns "github.com/oracle/oci-go-sdk/v56/dns"
	//oci_email "github.com/oracle/oci-go-sdk/v56/email"
	//oci_events "github.com/oracle/oci-go-sdk/v56/events"
	//oci_file_storage "github.com/oracle/oci-go-sdk/v56/filestorage"
	//oci_functions "github.com/oracle/oci-go-sdk/v56/functions"
	//oci_golden_gate "github.com/oracle/oci-go-sdk/v56/goldengate"
	//oci_core "github.com/oracle/oci-go-sdk/v56/core"
	oci_bastion "github.com/oracle/oci-go-sdk/v56/bastion"
	oci_budget "github.com/oracle/oci-go-sdk/v56/budget"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v56/cloudguard"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
	oci_database_tools "github.com/oracle/oci-go-sdk/v56/databasetools"
	oci_email "github.com/oracle/oci-go-sdk/v56/email"
	oci_events "github.com/oracle/oci-go-sdk/v56/events"
	oci_functions "github.com/oracle/oci-go-sdk/v56/functions"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"
	oci_kms "github.com/oracle/oci-go-sdk/v56/keymanagement"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v56/loadbalancer"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v56/osmanagement"

	//oci_load_balancer "github.com/oracle/oci-go-sdk/v56/loadbalancer"
	//oci_integration "github.com/oracle/oci-go-sdk/v56/integration"
	//oci_jms "github.com/oracle/oci-go-sdk/v56/jms"
	oci_jms "github.com/oracle/oci-go-sdk/v56/jms"
	//oci_kms "github.com/oracle/oci-go-sdk/v56/keymanagement"
	//oci_limits "github.com/oracle/oci-go-sdk/v56/limits"
	oci_limits "github.com/oracle/oci-go-sdk/v56/limits"

	//oci_monitoring "github.com/oracle/oci-go-sdk/v56/monitoring"
	//oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"
	//oci_network_load_balancer "github.com/oracle/oci-go-sdk/v56/networkloadbalancer"
	//oci_nosql "github.com/oracle/oci-go-sdk/v56/nosql"
	//oci_oce "github.com/oracle/oci-go-sdk/v56/oce"
	oci_ocvp "github.com/oracle/oci-go-sdk/v56/ocvp"
	oci_oda "github.com/oracle/oci-go-sdk/v56/oda"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v56/operatoraccesscontrol"

	//oci_opsi "github.com/oracle/oci-go-sdk/v56/opsi"
	oci_optimizer "github.com/oracle/oci-go-sdk/v56/optimizer"
	//oci_osmanagement "github.com/oracle/oci-go-sdk/v56/osmanagement"
	//oci_sch "github.com/oracle/oci-go-sdk/v56/sch"
	//oci_streaming "github.com/oracle/oci-go-sdk/v56/streaming"
	//oci_vulnerability_scanning "github.com/oracle/oci-go-sdk/v56/vulnerabilityscanning"
	//oci_waas "github.com/oracle/oci-go-sdk/v56/waas"
	//oci_waf "github.com/oracle/oci-go-sdk/v56/waf"
)

// Hints for discovering and exporting this resource to configuration and state files

var exportAiAnomalyDetectionDataAssetHints = &TerraformResourceHints{
	resourceClass:          "oci_ai_anomaly_detection_data_asset",
	datasourceClass:        "oci_ai_anomaly_detection_data_assets",
	datasourceItemsAttr:    "data_asset_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "data_asset",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_ai_anomaly_detection.DataAssetLifecycleStateActive),
	},
}

var exportAiAnomalyDetectionModelHints = &TerraformResourceHints{
	resourceClass:          "oci_ai_anomaly_detection_model",
	datasourceClass:        "oci_ai_anomaly_detection_models",
	datasourceItemsAttr:    "model_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "model",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_ai_anomaly_detection.ModelLifecycleStateActive),
	},
}

var exportAiAnomalyDetectionProjectHints = &TerraformResourceHints{
	resourceClass:          "oci_ai_anomaly_detection_project",
	datasourceClass:        "oci_ai_anomaly_detection_projects",
	datasourceItemsAttr:    "project_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "project",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_ai_anomaly_detection.ProjectLifecycleStateActive),
	},
}

var exportAiAnomalyDetectionAiPrivateEndpointHints = &TerraformResourceHints{
	resourceClass:          "oci_ai_anomaly_detection_ai_private_endpoint",
	datasourceClass:        "oci_ai_anomaly_detection_ai_private_endpoints",
	datasourceItemsAttr:    "ai_private_endpoint_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "ai_private_endpoint",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateActive),
	},
}

var exportAnalyticsAnalyticsInstanceHints = &TerraformResourceHints{
	resourceClass:          "oci_analytics_analytics_instance",
	datasourceClass:        "oci_analytics_analytics_instances",
	datasourceItemsAttr:    "analytics_instances",
	resourceAbbreviation:   "analytics_instance",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_analytics.AnalyticsInstanceLifecycleStateActive),
	},
}

var exportApigatewayApiHints = &TerraformResourceHints{
	resourceClass:          "oci_apigateway_api",
	datasourceClass:        "oci_apigateway_apis",
	datasourceItemsAttr:    "api_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "api",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_apigateway.ApiLifecycleStateActive),
	},
}

var exportApigatewayGatewayHints = &TerraformResourceHints{
	resourceClass:          "oci_apigateway_gateway",
	datasourceClass:        "oci_apigateway_gateways",
	datasourceItemsAttr:    "gateway_collection",
	resourceAbbreviation:   "gateway",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_apigateway.GatewayLifecycleStateActive),
	},
}

var exportApigatewayDeploymentHints = &TerraformResourceHints{
	resourceClass:          "oci_apigateway_deployment",
	datasourceClass:        "oci_apigateway_deployments",
	datasourceItemsAttr:    "deployment_collection",
	resourceAbbreviation:   "deployment",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_apigateway.DeploymentLifecycleStateActive),
	},
}

var exportApigatewayCertificateHints = &TerraformResourceHints{
	resourceClass:          "oci_apigateway_certificate",
	datasourceClass:        "oci_apigateway_certificates",
	datasourceItemsAttr:    "certificate_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "certificate",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_apigateway.CertificateLifecycleStateActive),
	},
}

var exportApmApmDomainHints = &TerraformResourceHints{
	resourceClass:        "oci_apm_apm_domain",
	datasourceClass:      "oci_apm_apm_domains",
	datasourceItemsAttr:  "apm_domains",
	resourceAbbreviation: "apm_domain",
	discoverableLifecycleStates: []string{
		string(oci_apm.LifecycleStatesActive),
	},
}

var exportApmConfigConfigHints = &TerraformResourceHints{
	resourceClass:          "oci_apm_config_config",
	datasourceClass:        "oci_apm_config_configs",
	datasourceItemsAttr:    "config_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "config",
	requireResourceRefresh: true,
}

var exportApmSyntheticsScriptHints = &TerraformResourceHints{
	resourceClass:          "oci_apm_synthetics_script",
	datasourceClass:        "oci_apm_synthetics_scripts",
	datasourceItemsAttr:    "script_collection",
	resourceAbbreviation:   "script",
	requireResourceRefresh: true,
}

var exportApmSyntheticsMonitorHints = &TerraformResourceHints{
	resourceClass:          "oci_apm_synthetics_monitor",
	datasourceClass:        "oci_apm_synthetics_monitors",
	datasourceItemsAttr:    "monitor_collection",
	resourceAbbreviation:   "monitor",
	requireResourceRefresh: true,
}

var exportArtifactsContainerConfigurationHints = &TerraformResourceHints{
	resourceClass:        "oci_artifacts_container_configuration",
	datasourceClass:      "oci_artifacts_container_configuration",
	resourceAbbreviation: "container_configuration",
}

var exportArtifactsContainerRepositoryHints = &TerraformResourceHints{
	resourceClass:          "oci_artifacts_container_repository",
	datasourceClass:        "oci_artifacts_container_repositories",
	datasourceItemsAttr:    "container_repository_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "container_repository",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_artifacts.ContainerRepositoryLifecycleStateAvailable),
	},
}

var exportArtifactsContainerImageSignatureHints = &TerraformResourceHints{
	resourceClass:          "oci_artifacts_container_image_signature",
	datasourceClass:        "oci_artifacts_container_image_signatures",
	datasourceItemsAttr:    "container_image_signature_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "container_image_signature",
	requireResourceRefresh: true,
}

var exportArtifactsRepositoryHints = &TerraformResourceHints{
	resourceClass:          "oci_artifacts_repository",
	datasourceClass:        "oci_artifacts_repositories",
	datasourceItemsAttr:    "repository_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "repository",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_artifacts.RepositoryLifecycleStateAvailable),
	},
}

var exportAutoScalingAutoScalingConfigurationHints = &TerraformResourceHints{
	resourceClass:          "oci_autoscaling_auto_scaling_configuration",
	datasourceClass:        "oci_autoscaling_auto_scaling_configurations",
	datasourceItemsAttr:    "auto_scaling_configurations",
	resourceAbbreviation:   "auto_scaling_configuration",
	requireResourceRefresh: true,
}

var exportBastionBastionHints = &TerraformResourceHints{
	resourceClass:          "oci_bastion_bastion",
	datasourceClass:        "oci_bastion_bastions",
	datasourceItemsAttr:    "bastions",
	resourceAbbreviation:   "bastion",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_bastion.BastionLifecycleStateActive),
	},
}

var exportBastionSessionHints = &TerraformResourceHints{
	resourceClass:          "oci_bastion_session",
	datasourceClass:        "oci_bastion_sessions",
	datasourceItemsAttr:    "sessions",
	resourceAbbreviation:   "session",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_bastion.SessionLifecycleStateActive),
	},
}

var exportBdsBdsInstanceHints = &TerraformResourceHints{
	resourceClass:          "oci_bds_bds_instance",
	datasourceClass:        "oci_bds_bds_instances",
	datasourceItemsAttr:    "bds_instances",
	resourceAbbreviation:   "bds_instance",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_bds.BdsInstanceLifecycleStateActive),
	},
}
var exportBdsBdsInstanceApiKeyHints = &TerraformResourceHints{
	resourceClass:          "oci_bds_bds_instance_api_key",
	datasourceClass:        "oci_bds_bds_instance_api_keys",
	datasourceItemsAttr:    "bds_api_keys",
	resourceAbbreviation:   "bds_instance_api_key",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_bds.BdsApiKeyLifecycleStateActive),
	},
}
var exportBlockchainBlockchainPlatformHints = &TerraformResourceHints{
	resourceClass:          "oci_blockchain_blockchain_platform",
	datasourceClass:        "oci_blockchain_blockchain_platforms",
	datasourceItemsAttr:    "blockchain_platform_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "blockchain_platform",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_blockchain.BlockchainPlatformLifecycleStateActive),
	},
}

var exportBlockchainPeerHints = &TerraformResourceHints{
	resourceClass:          "oci_blockchain_peer",
	datasourceClass:        "oci_blockchain_peers",
	datasourceItemsAttr:    "peer_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "peer",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_blockchain.PeerLifecycleStateActive),
	},
}

var exportBlockchainOsnHints = &TerraformResourceHints{
	resourceClass:          "oci_blockchain_osn",
	datasourceClass:        "oci_blockchain_osns",
	datasourceItemsAttr:    "osn_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "osn",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_blockchain.OsnLifecycleStateActive),
	},
}

var exportBudgetBudgetHints = &TerraformResourceHints{
	resourceClass:        "oci_budget_budget",
	datasourceClass:      "oci_budget_budgets",
	datasourceItemsAttr:  "budgets",
	resourceAbbreviation: "budget",
	discoverableLifecycleStates: []string{
		string(oci_budget.BudgetLifecycleStateActive),
	},
}

var exportBudgetAlertRuleHints = &TerraformResourceHints{
	resourceClass:        "oci_budget_alert_rule",
	datasourceClass:      "oci_budget_alert_rules",
	datasourceItemsAttr:  "alert_rules",
	resourceAbbreviation: "alert_rule",
	discoverableLifecycleStates: []string{
		string(oci_budget.AlertRuleLifecycleStateActive),
	},
}

var exportCertificatesManagementCaBundleHints = &TerraformResourceHints{
	resourceClass:          "oci_certificates_management_ca_bundle",
	datasourceClass:        "oci_certificates_management_ca_bundles",
	datasourceItemsAttr:    "ca_bundle_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "ca_bundle",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_certificates_management.CaBundleLifecycleStateActive),
	},
}

var exportCertificatesManagementCertificateAuthorityHints = &TerraformResourceHints{
	resourceClass:          "oci_certificates_management_certificate_authority",
	datasourceClass:        "oci_certificates_management_certificate_authorities",
	datasourceItemsAttr:    "certificate_authority_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "certificate_authority",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_certificates_management.CertificateAuthorityLifecycleStateActive),
	},
}

var exportCertificatesManagementCertificateHints = &TerraformResourceHints{
	resourceClass:          "oci_certificates_management_certificate",
	datasourceClass:        "oci_certificates_management_certificates",
	datasourceItemsAttr:    "certificate_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "certificate",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_certificates_management.CertificateLifecycleStateActive),
	},
}

var exportCloudGuardTargetHints = &TerraformResourceHints{
	resourceClass:          "oci_cloud_guard_target",
	datasourceClass:        "oci_cloud_guard_targets",
	datasourceItemsAttr:    "target_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "target",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardManagedListHints = &TerraformResourceHints{
	resourceClass:          "oci_cloud_guard_managed_list",
	datasourceClass:        "oci_cloud_guard_managed_lists",
	datasourceItemsAttr:    "managed_list_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "managed_list",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardResponderRecipeHints = &TerraformResourceHints{
	resourceClass:          "oci_cloud_guard_responder_recipe",
	datasourceClass:        "oci_cloud_guard_responder_recipes",
	datasourceItemsAttr:    "responder_recipe_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "responder_recipe",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardDataMaskRuleHints = &TerraformResourceHints{
	resourceClass:          "oci_cloud_guard_data_mask_rule",
	datasourceClass:        "oci_cloud_guard_data_mask_rules",
	datasourceItemsAttr:    "data_mask_rule_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "data_mask_rule",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardDetectorRecipeHints = &TerraformResourceHints{
	resourceClass:          "oci_cloud_guard_detector_recipe",
	datasourceClass:        "oci_cloud_guard_detector_recipes",
	datasourceItemsAttr:    "detector_recipe_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "detector_recipe",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportContainerengineClusterHints = &TerraformResourceHints{
	resourceClass:          "oci_containerengine_cluster",
	datasourceClass:        "oci_containerengine_clusters",
	datasourceItemsAttr:    "clusters",
	resourceAbbreviation:   "cluster",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_containerengine.ClusterLifecycleStateActive),
	},
}

var exportContainerengineNodePoolHints = &TerraformResourceHints{
	resourceClass:          "oci_containerengine_node_pool",
	datasourceClass:        "oci_containerengine_node_pools",
	datasourceItemsAttr:    "node_pools",
	resourceAbbreviation:   "node_pool",
	requireResourceRefresh: true,
}

var exportCoreBootVolumeBackupHints = &TerraformResourceHints{
	resourceClass:        "oci_core_boot_volume_backup",
	datasourceClass:      "oci_core_boot_volume_backups",
	datasourceItemsAttr:  "boot_volume_backups",
	resourceAbbreviation: "boot_volume_backup",
	discoverableLifecycleStates: []string{
		string(oci_core.BootVolumeBackupLifecycleStateAvailable),
	},
}

var exportCoreBootVolumeHints = &TerraformResourceHints{
	resourceClass:        "oci_core_boot_volume",
	datasourceClass:      "oci_core_boot_volumes",
	datasourceItemsAttr:  "boot_volumes",
	resourceAbbreviation: "boot_volume",
	discoverableLifecycleStates: []string{
		string(oci_core.BootVolumeLifecycleStateAvailable),
	},
}

var exportCoreConsoleHistoryHints = &TerraformResourceHints{
	resourceClass:        "oci_core_console_history",
	datasourceClass:      "oci_core_console_histories",
	datasourceItemsAttr:  "console_histories",
	resourceAbbreviation: "console_history",
	discoverableLifecycleStates: []string{
		string(oci_core.ConsoleHistoryLifecycleStateSucceeded),
	},
}

var exportCoreClusterNetworkHints = &TerraformResourceHints{
	resourceClass:          "oci_core_cluster_network",
	datasourceClass:        "oci_core_cluster_networks",
	datasourceItemsAttr:    "cluster_networks",
	resourceAbbreviation:   "cluster_network",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_core.ClusterNetworkLifecycleStateRunning),
	},
}

var exportCoreComputeImageCapabilitySchemaHints = &TerraformResourceHints{
	resourceClass:          "oci_core_compute_image_capability_schema",
	datasourceClass:        "oci_core_compute_image_capability_schemas",
	datasourceItemsAttr:    "compute_image_capability_schemas",
	resourceAbbreviation:   "compute_image_capability_schema",
	requireResourceRefresh: true,
}

var exportCoreCpeHints = &TerraformResourceHints{
	resourceClass:        "oci_core_cpe",
	datasourceClass:      "oci_core_cpes",
	datasourceItemsAttr:  "cpes",
	resourceAbbreviation: "cpe",
}

var exportCoreCrossConnectGroupHints = &TerraformResourceHints{
	resourceClass:        "oci_core_cross_connect_group",
	datasourceClass:      "oci_core_cross_connect_groups",
	datasourceItemsAttr:  "cross_connect_groups",
	resourceAbbreviation: "cross_connect_group",
	discoverableLifecycleStates: []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioned),
	},
}

var exportCoreCrossConnectHints = &TerraformResourceHints{
	resourceClass:        "oci_core_cross_connect",
	datasourceClass:      "oci_core_cross_connects",
	datasourceItemsAttr:  "cross_connects",
	resourceAbbreviation: "cross_connect",
	discoverableLifecycleStates: []string{
		string(oci_core.CrossConnectLifecycleStatePendingCustomer),
		string(oci_core.CrossConnectLifecycleStateProvisioned),
	},
}

var exportCoreDhcpOptionsHints = &TerraformResourceHints{
	resourceClass:        "oci_core_dhcp_options",
	datasourceClass:      "oci_core_dhcp_options",
	datasourceItemsAttr:  "options",
	resourceAbbreviation: "dhcp_options",
	discoverableLifecycleStates: []string{
		string(oci_core.DhcpOptionsLifecycleStateAvailable),
	},
}

var exportCoreDrgAttachmentHints = &TerraformResourceHints{
	resourceClass:        "oci_core_drg_attachment",
	datasourceClass:      "oci_core_drg_attachments",
	datasourceItemsAttr:  "drg_attachments",
	resourceAbbreviation: "drg_attachment",
	discoverableLifecycleStates: []string{
		string(oci_core.DrgAttachmentLifecycleStateAttached),
	},
}

var exportCoreDrgHints = &TerraformResourceHints{
	resourceClass:        "oci_core_drg",
	datasourceClass:      "oci_core_drgs",
	datasourceItemsAttr:  "drgs",
	resourceAbbreviation: "drg",
	discoverableLifecycleStates: []string{
		string(oci_core.DrgLifecycleStateAvailable),
	},
}

var exportCoreDedicatedVmHostHints = &TerraformResourceHints{
	resourceClass:          "oci_core_dedicated_vm_host",
	datasourceClass:        "oci_core_dedicated_vm_hosts",
	datasourceItemsAttr:    "dedicated_vm_hosts",
	resourceAbbreviation:   "dedicated_vm_host",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_core.DedicatedVmHostLifecycleStateActive),
	},
}

var exportCoreImageHints = &TerraformResourceHints{
	resourceClass:        "oci_core_image",
	datasourceClass:      "oci_core_images",
	datasourceItemsAttr:  "images",
	resourceAbbreviation: "image",
	discoverableLifecycleStates: []string{
		string(oci_core.ImageLifecycleStateAvailable),
	},
}

var exportCoreInstanceConfigurationHints = &TerraformResourceHints{
	resourceClass:          "oci_core_instance_configuration",
	datasourceClass:        "oci_core_instance_configurations",
	datasourceItemsAttr:    "instance_configurations",
	resourceAbbreviation:   "instance_configuration",
	requireResourceRefresh: true,
}

var exportCoreInstanceConsoleConnectionHints = &TerraformResourceHints{
	resourceClass:        "oci_core_instance_console_connection",
	datasourceClass:      "oci_core_instance_console_connections",
	datasourceItemsAttr:  "instance_console_connections",
	resourceAbbreviation: "instance_console_connection",
	discoverableLifecycleStates: []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateActive),
	},
}

var exportCoreInstancePoolInstanceHints = &TerraformResourceHints{
	resourceClass:        "oci_core_instance_pool_instance",
	datasourceClass:      "oci_core_instance_pool_instances",
	datasourceItemsAttr:  "instances",
	resourceAbbreviation: "instance_pool_instance",
	discoverableLifecycleStates: []string{
		string(oci_core.InstancePoolInstanceLifecycleStateActive),
	},
}

var exportCoreInstancePoolHints = &TerraformResourceHints{
	resourceClass:          "oci_core_instance_pool",
	datasourceClass:        "oci_core_instance_pools",
	datasourceItemsAttr:    "instance_pools",
	resourceAbbreviation:   "instance_pool",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_core.InstancePoolLifecycleStateRunning),
	},
}

var exportCoreInstanceHints = &TerraformResourceHints{
	resourceClass:        "oci_core_instance",
	datasourceClass:      "oci_core_instances",
	datasourceItemsAttr:  "instances",
	resourceAbbreviation: "instance",
	discoverableLifecycleStates: []string{
		string(oci_core.InstanceLifecycleStateRunning),
	},
}

var exportCoreInternetGatewayHints = &TerraformResourceHints{
	resourceClass:        "oci_core_internet_gateway",
	datasourceClass:      "oci_core_internet_gateways",
	datasourceItemsAttr:  "gateways",
	resourceAbbreviation: "internet_gateway",
	discoverableLifecycleStates: []string{
		string(oci_core.InternetGatewayLifecycleStateAvailable),
	},
}

var exportCoreIpSecConnectionHints = &TerraformResourceHints{
	resourceClass:        "oci_core_ipsec",
	datasourceClass:      "oci_core_ipsec_connections",
	datasourceItemsAttr:  "connections",
	resourceAbbreviation: "ip_sec_connection",
	discoverableLifecycleStates: []string{
		string(oci_core.IpSecConnectionLifecycleStateAvailable),
	},
}

var exportCoreLocalPeeringGatewayHints = &TerraformResourceHints{
	resourceClass:        "oci_core_local_peering_gateway",
	datasourceClass:      "oci_core_local_peering_gateways",
	datasourceItemsAttr:  "local_peering_gateways",
	resourceAbbreviation: "local_peering_gateway",
	discoverableLifecycleStates: []string{
		string(oci_core.LocalPeeringGatewayLifecycleStateAvailable),
	},
}

var exportCoreNatGatewayHints = &TerraformResourceHints{
	resourceClass:        "oci_core_nat_gateway",
	datasourceClass:      "oci_core_nat_gateways",
	datasourceItemsAttr:  "nat_gateways",
	resourceAbbreviation: "nat_gateway",
	discoverableLifecycleStates: []string{
		string(oci_core.NatGatewayLifecycleStateAvailable),
	},
}

var exportCoreNetworkSecurityGroupHints = &TerraformResourceHints{
	resourceClass:        "oci_core_network_security_group",
	datasourceClass:      "oci_core_network_security_groups",
	datasourceItemsAttr:  "network_security_groups",
	resourceAbbreviation: "network_security_group",
	discoverableLifecycleStates: []string{
		string(oci_core.NetworkSecurityGroupLifecycleStateAvailable),
	},
}

var exportCoreNetworkSecurityGroupSecurityRuleHints = &TerraformResourceHints{
	resourceClass:        "oci_core_network_security_group_security_rule",
	resourceAbbreviation: "network_security_group_security_rule",
}

var exportCorePrivateIpHints = &TerraformResourceHints{
	resourceClass:        "oci_core_private_ip",
	datasourceClass:      "oci_core_private_ips",
	datasourceItemsAttr:  "private_ips",
	resourceAbbreviation: "private_ip",
}

var exportCorePublicIpHints = &TerraformResourceHints{
	resourceClass:        "oci_core_public_ip",
	datasourceClass:      "oci_core_public_ips",
	datasourceItemsAttr:  "public_ips",
	resourceAbbreviation: "public_ip",
	discoverableLifecycleStates: []string{
		string(oci_core.PublicIpLifecycleStateAvailable),
		string(oci_core.PublicIpLifecycleStateAssigned),
	},
}

var exportCoreRemotePeeringConnectionHints = &TerraformResourceHints{
	resourceClass:        "oci_core_remote_peering_connection",
	datasourceClass:      "oci_core_remote_peering_connections",
	datasourceItemsAttr:  "remote_peering_connections",
	resourceAbbreviation: "remote_peering_connection",
	discoverableLifecycleStates: []string{
		string(oci_core.RemotePeeringConnectionLifecycleStateAvailable),
	},
}

var exportCoreRouteTableHints = &TerraformResourceHints{
	resourceClass:        "oci_core_route_table",
	datasourceClass:      "oci_core_route_tables",
	datasourceItemsAttr:  "route_tables",
	resourceAbbreviation: "route_table",
	discoverableLifecycleStates: []string{
		string(oci_core.RouteTableLifecycleStateAvailable),
	},
}

var exportCoreSecurityListHints = &TerraformResourceHints{
	resourceClass:        "oci_core_security_list",
	datasourceClass:      "oci_core_security_lists",
	datasourceItemsAttr:  "security_lists",
	resourceAbbreviation: "security_list",
	discoverableLifecycleStates: []string{
		string(oci_core.SecurityListLifecycleStateAvailable),
	},
}

var exportCoreServiceGatewayHints = &TerraformResourceHints{
	resourceClass:        "oci_core_service_gateway",
	datasourceClass:      "oci_core_service_gateways",
	datasourceItemsAttr:  "service_gateways",
	resourceAbbreviation: "service_gateway",
	discoverableLifecycleStates: []string{
		string(oci_core.ServiceGatewayLifecycleStateAvailable),
	},
}

var exportCoreSubnetHints = &TerraformResourceHints{
	resourceClass:        "oci_core_subnet",
	datasourceClass:      "oci_core_subnets",
	datasourceItemsAttr:  "subnets",
	resourceAbbreviation: "subnet",
	discoverableLifecycleStates: []string{
		string(oci_core.SubnetLifecycleStateAvailable),
	},
}

var exportCoreVcnHints = &TerraformResourceHints{
	resourceClass:        "oci_core_vcn",
	datasourceClass:      "oci_core_vcns",
	datasourceItemsAttr:  "virtual_networks",
	resourceAbbreviation: "vcn",
	discoverableLifecycleStates: []string{
		string(oci_core.VcnLifecycleStateAvailable),
	},
}

var exportCoreVlanHints = &TerraformResourceHints{
	resourceClass:        "oci_core_vlan",
	datasourceClass:      "oci_core_vlans",
	datasourceItemsAttr:  "vlans",
	resourceAbbreviation: "vlan",
	discoverableLifecycleStates: []string{
		string(oci_core.VlanLifecycleStateAvailable),
	},
}

var exportCoreVirtualCircuitHints = &TerraformResourceHints{
	resourceClass:        "oci_core_virtual_circuit",
	datasourceClass:      "oci_core_virtual_circuits",
	datasourceItemsAttr:  "virtual_circuits",
	resourceAbbreviation: "virtual_circuit",
	discoverableLifecycleStates: []string{
		string(oci_core.VirtualCircuitLifecycleStatePendingProvider),
		string(oci_core.VirtualCircuitLifecycleStateProvisioned),
	},
}

var exportCoreVnicAttachmentHints = &TerraformResourceHints{
	resourceClass:        "oci_core_vnic_attachment",
	datasourceClass:      "oci_core_vnic_attachments",
	datasourceItemsAttr:  "vnic_attachments",
	resourceAbbreviation: "vnic_attachment",
	discoverableLifecycleStates: []string{
		string(oci_core.VnicAttachmentLifecycleStateAttached),
	},
}

var exportCoreVolumeAttachmentHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume_attachment",
	datasourceClass:      "oci_core_volume_attachments",
	datasourceItemsAttr:  "volume_attachments",
	resourceAbbreviation: "volume_attachment",
	discoverableLifecycleStates: []string{
		string(oci_core.VolumeAttachmentLifecycleStateAttached),
	},
}

var exportCoreVolumeBackupHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume_backup",
	datasourceClass:      "oci_core_volume_backups",
	datasourceItemsAttr:  "volume_backups",
	resourceAbbreviation: "volume_backup",
	discoverableLifecycleStates: []string{
		string(oci_core.VolumeBackupLifecycleStateAvailable),
	},
}

var exportCoreVolumeBackupPolicyHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume_backup_policy",
	datasourceClass:      "oci_core_volume_backup_policies",
	datasourceItemsAttr:  "volume_backup_policies",
	resourceAbbreviation: "volume_backup_policy",
}

var exportCoreVolumeBackupPolicyAssignmentHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume_backup_policy_assignment",
	datasourceClass:      "oci_core_volume_backup_policy_assignments",
	datasourceItemsAttr:  "volume_backup_policy_assignments",
	resourceAbbreviation: "volume_backup_policy_assignment",
}

var exportCoreVolumeGroupHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume_group",
	datasourceClass:      "oci_core_volume_groups",
	datasourceItemsAttr:  "volume_groups",
	resourceAbbreviation: "volume_group",
	discoverableLifecycleStates: []string{
		string(oci_core.VolumeGroupLifecycleStateAvailable),
	},
}

var exportCoreVolumeGroupBackupHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume_group_backup",
	datasourceClass:      "oci_core_volume_group_backups",
	datasourceItemsAttr:  "volume_group_backups",
	resourceAbbreviation: "volume_group_backup",
	discoverableLifecycleStates: []string{
		string(oci_core.VolumeGroupBackupLifecycleStateCommitted),
		string(oci_core.VolumeGroupBackupLifecycleStateAvailable),
	},
}

var exportCoreVolumeHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume",
	datasourceClass:      "oci_core_volumes",
	datasourceItemsAttr:  "volumes",
	resourceAbbreviation: "volume",
	discoverableLifecycleStates: []string{
		string(oci_core.VolumeLifecycleStateAvailable),
	},
}

var exportCorePublicIpPoolHints = &TerraformResourceHints{
	resourceClass:          "oci_core_public_ip_pool",
	datasourceClass:        "oci_core_public_ip_pools",
	datasourceItemsAttr:    "public_ip_pool_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "public_ip_pool",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_core.PublicIpPoolLifecycleStateActive),
	},
}

var exportCoreIpv6Hints = &TerraformResourceHints{
	resourceClass:        "oci_core_ipv6",
	datasourceClass:      "oci_core_ipv6s",
	datasourceItemsAttr:  "ipv6s",
	resourceAbbreviation: "ipv6",
	discoverableLifecycleStates: []string{
		string(oci_core.Ipv6LifecycleStateAvailable),
	},
}

var exportCoreDrgRouteTableHints = &TerraformResourceHints{
	resourceClass:        "oci_core_drg_route_table",
	datasourceClass:      "oci_core_drg_route_tables",
	datasourceItemsAttr:  "drg_route_tables",
	resourceAbbreviation: "drg_route_table",
	discoverableLifecycleStates: []string{
		string(oci_core.DrgRouteTableLifecycleStateAvailable),
	},
}

var exportCoreDrgRouteDistributionHints = &TerraformResourceHints{
	resourceClass:        "oci_core_drg_route_distribution",
	datasourceClass:      "oci_core_drg_route_distributions",
	datasourceItemsAttr:  "drg_route_distributions",
	resourceAbbreviation: "drg_route_distribution",
	discoverableLifecycleStates: []string{
		string(oci_core.DrgRouteDistributionLifecycleStateAvailable),
	},
}

var exportCoreDrgRouteTableRouteRuleHints = &TerraformResourceHints{
	resourceClass:        "oci_core_drg_route_table_route_rule",
	resourceAbbreviation: "drg_route_table_route_rule",
}

var exportDataSafeDataSafePrivateEndpointHints = &TerraformResourceHints{
	resourceClass:          "oci_data_safe_data_safe_private_endpoint",
	datasourceClass:        "oci_data_safe_data_safe_private_endpoints",
	datasourceItemsAttr:    "data_safe_private_endpoints",
	resourceAbbreviation:   "data_safe_private_endpoint",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_data_safe.ListDataSafePrivateEndpointsLifecycleStateActive),
	},
}
var exportDataLabelingServiceDatasetHints = &TerraformResourceHints{
	resourceClass:          "oci_data_labeling_service_dataset",
	datasourceClass:        "oci_data_labeling_service_datasets",
	datasourceItemsAttr:    "dataset_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "dataset",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_data_labeling_service.DatasetLifecycleStateActive),
		string(oci_data_labeling_service.DatasetLifecycleStateNeedsAttention),
	},
}

var exportDataSafeOnPremConnectorHints = &TerraformResourceHints{
	resourceClass:        "oci_data_safe_on_prem_connector",
	datasourceClass:      "oci_data_safe_on_prem_connectors",
	datasourceItemsAttr:  "on_prem_connectors",
	resourceAbbreviation: "on_prem_connector",
	discoverableLifecycleStates: []string{
		string(oci_data_safe.OnPremConnectorLifecycleStateInactive),
		string(oci_data_safe.OnPremConnectorLifecycleStateActive),
	},
}

var exportDataSafeTargetDatabaseHints = &TerraformResourceHints{
	resourceClass:          "oci_data_safe_target_database",
	datasourceClass:        "oci_data_safe_target_databases",
	datasourceItemsAttr:    "target_databases",
	resourceAbbreviation:   "target_database",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateActive),
	},
}

var exportDataSafeSecurityAssessmentHints = &TerraformResourceHints{
	resourceClass:        "oci_data_safe_security_assessment",
	datasourceClass:      "oci_data_safe_security_assessments",
	datasourceItemsAttr:  "security_assessments",
	resourceAbbreviation: "security_assessment",
	discoverableLifecycleStates: []string{
		string(oci_data_safe.SecurityAssessmentLifecycleStateSucceeded),
	},
}

var exportDataSafeUserAssessmentHints = &TerraformResourceHints{
	resourceClass:        "oci_data_safe_user_assessment",
	datasourceClass:      "oci_data_safe_user_assessments",
	datasourceItemsAttr:  "user_assessments",
	resourceAbbreviation: "user_assessment",
	discoverableLifecycleStates: []string{
		string(oci_data_safe.UserAssessmentLifecycleStateSucceeded),
	},
}

var exportDatabaseAutonomousContainerDatabaseHints = &TerraformResourceHints{
	resourceClass:        "oci_database_autonomous_container_database",
	datasourceClass:      "oci_database_autonomous_container_databases",
	datasourceItemsAttr:  "autonomous_container_databases",
	resourceAbbreviation: "autonomous_container_database",
	discoverableLifecycleStates: []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseAutonomousContainerDatabaseDataguardAssociationHints = &TerraformResourceHints{
	resourceClass:        "oci_database_autonomous_container_database_dataguard_association",
	datasourceClass:      "oci_database_autonomous_container_database_dataguard_associations",
	datasourceItemsAttr:  "autonomous_container_database_dataguard_associations",
	resourceAbbreviation: "autonomous_container_database_dataguard_association",
	discoverableLifecycleStates: []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateAvailable),
	},
}

var exportDatabaseAutonomousDatabaseHints = &TerraformResourceHints{
	resourceClass:        "oci_database_autonomous_database",
	datasourceClass:      "oci_database_autonomous_databases",
	datasourceItemsAttr:  "autonomous_databases",
	resourceAbbreviation: "autonomous_database",
	discoverableLifecycleStates: []string{
		string(oci_database.AutonomousDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseAutonomousExadataInfrastructureHints = &TerraformResourceHints{
	resourceClass:        "oci_database_autonomous_exadata_infrastructure",
	datasourceClass:      "oci_database_autonomous_exadata_infrastructures",
	datasourceItemsAttr:  "autonomous_exadata_infrastructures",
	resourceAbbreviation: "autonomous_exadata_infrastructure",
	discoverableLifecycleStates: []string{
		string(oci_database.AutonomousExadataInfrastructureLifecycleStateAvailable),
	},
}

var exportDatabaseAutonomousVmClusterHints = &TerraformResourceHints{
	resourceClass:        "oci_database_autonomous_vm_cluster",
	datasourceClass:      "oci_database_autonomous_vm_clusters",
	datasourceItemsAttr:  "autonomous_vm_clusters",
	resourceAbbreviation: "autonomous_vm_cluster",
	discoverableLifecycleStates: []string{
		string(oci_database.AutonomousVmClusterLifecycleStateAvailable),
	},
}

var exportDatabaseBackupDestinationHints = &TerraformResourceHints{
	resourceClass:        "oci_database_backup_destination",
	datasourceClass:      "oci_database_backup_destinations",
	datasourceItemsAttr:  "backup_destinations",
	resourceAbbreviation: "backup_destination",
	discoverableLifecycleStates: []string{
		string(oci_database.BackupDestinationLifecycleStateActive),
	},
}

var exportDatabaseBackupHints = &TerraformResourceHints{
	resourceClass:        "oci_database_backup",
	datasourceClass:      "oci_database_backups",
	datasourceItemsAttr:  "backups",
	resourceAbbreviation: "backup",
	discoverableLifecycleStates: []string{
		string(oci_database.BackupLifecycleStateActive),
	},
}

var exportDatabaseDatabaseHints = &TerraformResourceHints{
	resourceClass:        "oci_database_database",
	datasourceClass:      "oci_database_databases",
	datasourceItemsAttr:  "databases",
	resourceAbbreviation: "database",
	discoverableLifecycleStates: []string{
		string(oci_database.DatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseDbHomeHints = &TerraformResourceHints{
	resourceClass:        "oci_database_db_home",
	datasourceClass:      "oci_database_db_homes",
	datasourceItemsAttr:  "db_homes",
	resourceAbbreviation: "db_home",
	discoverableLifecycleStates: []string{
		string(oci_database.DbHomeLifecycleStateAvailable),
	},
}

var exportDatabaseDbSystemHints = &TerraformResourceHints{
	resourceClass:        "oci_database_db_system",
	datasourceClass:      "oci_database_db_systems",
	datasourceItemsAttr:  "db_systems",
	resourceAbbreviation: "db_system",
	discoverableLifecycleStates: []string{
		string(oci_database.DbSystemLifecycleStateAvailable),
	},
}

var exportDatabaseExadataInfrastructureHints = &TerraformResourceHints{
	resourceClass:        "oci_database_exadata_infrastructure",
	datasourceClass:      "oci_database_exadata_infrastructures",
	datasourceItemsAttr:  "exadata_infrastructures",
	resourceAbbreviation: "exadata_infrastructure",
	discoverableLifecycleStates: []string{
		string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
		string(oci_database.ExadataInfrastructureLifecycleStateActive),
	},
}

var exportDatabaseVmClusterNetworkHints = &TerraformResourceHints{
	resourceClass:        "oci_database_vm_cluster_network",
	datasourceClass:      "oci_database_vm_cluster_networks",
	datasourceItemsAttr:  "vm_cluster_networks",
	resourceAbbreviation: "vm_cluster_network",
	discoverableLifecycleStates: []string{
		string(oci_database.VmClusterNetworkLifecycleStateRequiresValidation),
		string(oci_database.VmClusterNetworkLifecycleStateValidated),
		string(oci_database.VmClusterNetworkLifecycleStateAllocated),
	},
}

var exportDatabaseVmClusterHints = &TerraformResourceHints{
	resourceClass:        "oci_database_vm_cluster",
	datasourceClass:      "oci_database_vm_clusters",
	datasourceItemsAttr:  "vm_clusters",
	resourceAbbreviation: "vm_cluster",
	discoverableLifecycleStates: []string{
		string(oci_database.VmClusterLifecycleStateAvailable),
	},
}

var exportDatabaseDatabaseSoftwareImageHints = &TerraformResourceHints{
	resourceClass:        "oci_database_database_software_image",
	datasourceClass:      "oci_database_database_software_images",
	datasourceItemsAttr:  "database_software_images",
	resourceAbbreviation: "database_software_image",
	discoverableLifecycleStates: []string{
		string(oci_database.DatabaseSoftwareImageLifecycleStateAvailable),
	},
}

var exportDatabaseCloudExadataInfrastructureHints = &TerraformResourceHints{
	resourceClass:        "oci_database_cloud_exadata_infrastructure",
	datasourceClass:      "oci_database_cloud_exadata_infrastructures",
	datasourceItemsAttr:  "cloud_exadata_infrastructures",
	resourceAbbreviation: "cloud_exadata_infrastructure",
	discoverableLifecycleStates: []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateAvailable),
	},
}

var exportDatabaseCloudVmClusterHints = &TerraformResourceHints{
	resourceClass:        "oci_database_cloud_vm_cluster",
	datasourceClass:      "oci_database_cloud_vm_clusters",
	datasourceItemsAttr:  "cloud_vm_clusters",
	resourceAbbreviation: "cloud_vm_cluster",
	discoverableLifecycleStates: []string{
		string(oci_database.CloudVmClusterLifecycleStateAvailable),
	},
}

var exportDatabaseKeyStoreHints = &TerraformResourceHints{
	resourceClass:        "oci_database_key_store",
	datasourceClass:      "oci_database_key_stores",
	datasourceItemsAttr:  "key_stores",
	resourceAbbreviation: "key_store",
	discoverableLifecycleStates: []string{
		string(oci_database.KeyStoreLifecycleStateActive),
	},
}

var exportDatabaseExternalContainerDatabaseHints = &TerraformResourceHints{
	resourceClass:        "oci_database_external_container_database",
	datasourceClass:      "oci_database_external_container_databases",
	datasourceItemsAttr:  "external_container_databases",
	resourceAbbreviation: "external_container_database",
	discoverableLifecycleStates: []string{
		string(oci_database.ExternalContainerDatabaseLifecycleStateNotConnected),
		string(oci_database.ExternalContainerDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseExternalPluggableDatabaseHints = &TerraformResourceHints{
	resourceClass:        "oci_database_external_pluggable_database",
	datasourceClass:      "oci_database_external_pluggable_databases",
	datasourceItemsAttr:  "external_pluggable_databases",
	resourceAbbreviation: "external_pluggable_database",
	discoverableLifecycleStates: []string{
		string(oci_database.ExternalPluggableDatabaseLifecycleStateNotConnected),
		string(oci_database.ExternalPluggableDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseExternalNonContainerDatabaseHints = &TerraformResourceHints{
	resourceClass:        "oci_database_external_non_container_database",
	datasourceClass:      "oci_database_external_non_container_databases",
	datasourceItemsAttr:  "external_non_container_databases",
	resourceAbbreviation: "external_non_container_database",
	discoverableLifecycleStates: []string{
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateNotConnected),
		string(oci_database.ExternalNonContainerDatabaseLifecycleStateAvailable),
	},
}

var exportDatabaseExternalDatabaseConnectorHints = &TerraformResourceHints{
	resourceClass:        "oci_database_external_database_connector",
	datasourceClass:      "oci_database_external_database_connectors",
	datasourceItemsAttr:  "external_database_connectors",
	resourceAbbreviation: "external_database_connector",
	discoverableLifecycleStates: []string{
		string(oci_database.ExternalDatabaseConnectorLifecycleStateAvailable),
	},
}

var exportDatabaseCloudAutonomousVmClusterHints = &TerraformResourceHints{
	resourceClass:        "oci_database_cloud_autonomous_vm_cluster",
	datasourceClass:      "oci_database_cloud_autonomous_vm_clusters",
	datasourceItemsAttr:  "cloud_autonomous_vm_clusters",
	resourceAbbreviation: "cloud_autonomous_vm_cluster",
	discoverableLifecycleStates: []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateAvailable),
	},
}

var exportDatabaseMigrationMigrationHints = &TerraformResourceHints{
	resourceClass:          "oci_database_migration_migration",
	datasourceClass:        "oci_database_migration_migrations",
	datasourceItemsAttr:    "migration_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "migration",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_database_migration.LifecycleStatesActive),
	},
}

var exportDatabaseMigrationConnectionHints = &TerraformResourceHints{
	resourceClass:          "oci_database_migration_connection",
	datasourceClass:        "oci_database_migration_connections",
	datasourceItemsAttr:    "connection_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "connection",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_database_migration.LifecycleStatesActive),
	},
}

var exportDatabaseToolsDatabaseToolsPrivateEndpointHints = &TerraformResourceHints{
	resourceClass:          "oci_database_tools_database_tools_private_endpoint",
	datasourceClass:        "oci_database_tools_database_tools_private_endpoints",
	datasourceItemsAttr:    "database_tools_private_endpoint_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "database_tools_private_endpoint",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_database_tools.LifecycleStateActive),
	},
}

var exportDatabaseToolsDatabaseToolsConnectionHints = &TerraformResourceHints{
	resourceClass:          "oci_database_tools_database_tools_connection",
	datasourceClass:        "oci_database_tools_database_tools_connections",
	datasourceItemsAttr:    "database_tools_connection_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "database_tools_connection",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_database_tools.LifecycleStateActive),
	},
}

var exportDatabasePluggableDatabaseHints = &TerraformResourceHints{
	resourceClass:        "oci_database_pluggable_database",
	datasourceClass:      "oci_database_pluggable_databases",
	datasourceItemsAttr:  "pluggable_databases",
	resourceAbbreviation: "pluggable_database",
	discoverableLifecycleStates: []string{
		string(oci_database.PluggableDatabaseLifecycleStateAvailable),
	},
}

var exportDatacatalogCatalogHints = &TerraformResourceHints{
	resourceClass:        "oci_datacatalog_catalog",
	datasourceClass:      "oci_datacatalog_catalogs",
	datasourceItemsAttr:  "catalogs",
	resourceAbbreviation: "catalog",
	discoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDatacatalogDataAssetHints = &TerraformResourceHints{
	resourceClass:          "oci_datacatalog_data_asset",
	datasourceClass:        "oci_datacatalog_data_assets",
	datasourceItemsAttr:    "data_asset_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "data_asset",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDatacatalogConnectionHints = &TerraformResourceHints{
	resourceClass:          "oci_datacatalog_connection",
	datasourceClass:        "oci_datacatalog_connections",
	datasourceItemsAttr:    "connection_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "connection",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDatacatalogCatalogPrivateEndpointHints = &TerraformResourceHints{
	resourceClass:        "oci_datacatalog_catalog_private_endpoint",
	datasourceClass:      "oci_datacatalog_catalog_private_endpoints",
	datasourceItemsAttr:  "catalog_private_endpoints",
	resourceAbbreviation: "catalog_private_endpoint",
	discoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDatacatalogMetastoreHints = &TerraformResourceHints{
	resourceClass:          "oci_datacatalog_metastore",
	datasourceClass:        "oci_datacatalog_metastores",
	datasourceItemsAttr:    "metastores",
	resourceAbbreviation:   "metastore",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDataflowApplicationHints = &TerraformResourceHints{
	resourceClass:          "oci_dataflow_application",
	datasourceClass:        "oci_dataflow_applications",
	datasourceItemsAttr:    "applications",
	resourceAbbreviation:   "application",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_dataflow.ApplicationLifecycleStateActive),
	},
}

var exportDataflowPrivateEndpointHints = &TerraformResourceHints{
	resourceClass:          "oci_dataflow_private_endpoint",
	datasourceClass:        "oci_dataflow_private_endpoints",
	datasourceItemsAttr:    "private_endpoint_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "private_endpoint",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_dataflow.PrivateEndpointLifecycleStateActive),
		string(oci_dataflow.PrivateEndpointLifecycleStateInactive),
	},
}

var exportDataintegrationWorkspaceHints = &TerraformResourceHints{
	resourceClass:          "oci_dataintegration_workspace",
	datasourceClass:        "oci_dataintegration_workspaces",
	datasourceItemsAttr:    "workspaces",
	resourceAbbreviation:   "workspace",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_dataintegration.WorkspaceLifecycleStateActive),
	},
}

var exportDatascienceProjectHints = &TerraformResourceHints{
	resourceClass:        "oci_datascience_project",
	datasourceClass:      "oci_datascience_projects",
	datasourceItemsAttr:  "projects",
	resourceAbbreviation: "project",
	discoverableLifecycleStates: []string{
		string(oci_datascience.ProjectLifecycleStateActive),
	},
}

var exportDatascienceNotebookSessionHints = &TerraformResourceHints{
	resourceClass:        "oci_datascience_notebook_session",
	datasourceClass:      "oci_datascience_notebook_sessions",
	datasourceItemsAttr:  "notebook_sessions",
	resourceAbbreviation: "notebook_session",
	discoverableLifecycleStates: []string{
		string(oci_datascience.NotebookSessionLifecycleStateActive),
	},
}

var exportDatascienceModelHints = &TerraformResourceHints{
	resourceClass:          "oci_datascience_model",
	datasourceClass:        "oci_datascience_models",
	datasourceItemsAttr:    "models",
	resourceAbbreviation:   "model",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_datascience.ModelLifecycleStateActive),
	},
}

var exportDatascienceModelProvenanceHints = &TerraformResourceHints{
	resourceClass:        "oci_datascience_model_provenance",
	datasourceClass:      "oci_datascience_model_provenance",
	resourceAbbreviation: "model_provenance",
}

var exportDatascienceModelDeploymentHints = &TerraformResourceHints{
	resourceClass:        "oci_datascience_model_deployment",
	datasourceClass:      "oci_datascience_model_deployments",
	datasourceItemsAttr:  "model_deployments",
	resourceAbbreviation: "model_deployment",
	discoverableLifecycleStates: []string{
		string(oci_datascience.ModelDeploymentLifecycleStateActive),
		string(oci_datascience.ModelDeploymentLifecycleStateNeedsAttention),
	},
}

var exportDatascienceJobHints = &TerraformResourceHints{
	resourceClass:          "oci_datascience_job",
	datasourceClass:        "oci_datascience_jobs",
	datasourceItemsAttr:    "jobs",
	resourceAbbreviation:   "job",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_datascience.JobLifecycleStateActive),
	},
}

var exportDatascienceJobRunHints = &TerraformResourceHints{
	resourceClass:          "oci_datascience_job_run",
	datasourceClass:        "oci_datascience_job_runs",
	datasourceItemsAttr:    "job_runs",
	resourceAbbreviation:   "job_run",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_datascience.JobRunLifecycleStateSucceeded),
		string(oci_datascience.JobRunLifecycleStateNeedsAttention),
	},
}
var exportDevopsProjectHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_project",
	datasourceClass:        "oci_devops_projects",
	datasourceItemsAttr:    "project_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "project",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.ProjectLifecycleStateActive),
	},
}

var exportDevopsDeployEnvironmentHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_deploy_environment",
	datasourceClass:        "oci_devops_deploy_environments",
	datasourceItemsAttr:    "deploy_environment_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "deploy_environment",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.DeployEnvironmentLifecycleStateActive),
	},
}

var exportDevopsDeployArtifactHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_deploy_artifact",
	datasourceClass:        "oci_devops_deploy_artifacts",
	datasourceItemsAttr:    "deploy_artifact_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "deploy_artifact",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.DeployArtifactLifecycleStateActive),
	},
}

var exportDevopsDeployPipelineHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_deploy_pipeline",
	datasourceClass:        "oci_devops_deploy_pipelines",
	datasourceItemsAttr:    "deploy_pipeline_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "deploy_pipeline",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.DeployPipelineLifecycleStateActive),
	},
}

var exportDevopsDeployStageHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_deploy_stage",
	datasourceClass:        "oci_devops_deploy_stages",
	datasourceItemsAttr:    "deploy_stage_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "deploy_stage",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.DeployStageLifecycleStateActive),
	},
}

var exportDevopsDeploymentHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_deployment",
	datasourceClass:        "oci_devops_deployments",
	datasourceItemsAttr:    "deployment_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "deployment",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.DeploymentLifecycleStateSucceeded),
	},
}

var exportDevopsRepositoryHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_repository",
	datasourceClass:        "oci_devops_repositories",
	datasourceItemsAttr:    "repository_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "repository",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.RepositoryLifecycleStateActive),
	},
}

var exportDevopsRepositoryRefHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_repository_ref",
	datasourceClass:        "oci_devops_repository_refs",
	datasourceItemsAttr:    "repository_ref_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "repository_ref",
	requireResourceRefresh: true,
}

var exportDevopsBuildPipelineHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_build_pipeline",
	datasourceClass:        "oci_devops_build_pipelines",
	datasourceItemsAttr:    "build_pipeline_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "build_pipeline",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.BuildPipelineLifecycleStateActive),
	},
}

var exportDevopsBuildRunHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_build_run",
	datasourceClass:        "oci_devops_build_runs",
	datasourceItemsAttr:    "build_run_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "build_run",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.BuildRunLifecycleStateSucceeded),
	},
}

var exportDevopsConnectionHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_connection",
	datasourceClass:        "oci_devops_connections",
	datasourceItemsAttr:    "connection_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "connection",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.ConnectionLifecycleStateActive),
	},
}

var exportDevopsBuildPipelineStageHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_build_pipeline_stage",
	datasourceClass:        "oci_devops_build_pipeline_stages",
	datasourceItemsAttr:    "build_pipeline_stage_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "build_pipeline_stage",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.BuildPipelineStageLifecycleStateActive),
	},
}

var exportDevopsTriggerHints = &TerraformResourceHints{
	resourceClass:          "oci_devops_trigger",
	datasourceClass:        "oci_devops_triggers",
	datasourceItemsAttr:    "trigger_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "trigger",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_devops.TriggerLifecycleStateActive),
	},
}

var exportDevopsRepositoryMirrorHints = &TerraformResourceHints{
	resourceClass:        "oci_devops_repository_mirror",
	resourceAbbreviation: "repository_mirror",
}

var exportDnsZoneHints = &TerraformResourceHints{
	resourceClass:          "oci_dns_zone",
	datasourceClass:        "oci_dns_zones",
	datasourceItemsAttr:    "zones",
	resourceAbbreviation:   "zone",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_dns.ZoneLifecycleStateActive),
	},
}

var exportDnsSteeringPolicyHints = &TerraformResourceHints{
	resourceClass:          "oci_dns_steering_policy",
	datasourceClass:        "oci_dns_steering_policies",
	datasourceItemsAttr:    "steering_policies",
	resourceAbbreviation:   "steering_policy",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_dns.SteeringPolicyLifecycleStateActive),
	},
}

var exportDnsSteeringPolicyAttachmentHints = &TerraformResourceHints{
	resourceClass:        "oci_dns_steering_policy_attachment",
	datasourceClass:      "oci_dns_steering_policy_attachments",
	datasourceItemsAttr:  "steering_policy_attachments",
	resourceAbbreviation: "steering_policy_attachment",
	discoverableLifecycleStates: []string{
		string(oci_dns.SteeringPolicyAttachmentLifecycleStateActive),
	},
}

var exportDnsTsigKeyHints = &TerraformResourceHints{
	resourceClass:          "oci_dns_tsig_key",
	datasourceClass:        "oci_dns_tsig_keys",
	datasourceItemsAttr:    "tsig_keys",
	resourceAbbreviation:   "tsig_key",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_dns.TsigKeyLifecycleStateActive),
	},
}

var exportDnsRrsetHints = &TerraformResourceHints{
	resourceClass:        "oci_dns_rrset",
	datasourceClass:      "oci_dns_rrset",
	resourceAbbreviation: "rrset",
}

var exportEmailSuppressionHints = &TerraformResourceHints{
	resourceClass:        "oci_email_suppression",
	datasourceClass:      "oci_email_suppressions",
	datasourceItemsAttr:  "suppressions",
	resourceAbbreviation: "suppression",
}

var exportEmailSenderHints = &TerraformResourceHints{
	resourceClass:        "oci_email_sender",
	datasourceClass:      "oci_email_senders",
	datasourceItemsAttr:  "senders",
	resourceAbbreviation: "sender",
	discoverableLifecycleStates: []string{
		string(oci_email.SenderLifecycleStateActive),
	},
}

var exportEmailEmailDomainHints = &TerraformResourceHints{
	resourceClass:          "oci_email_email_domain",
	datasourceClass:        "oci_email_email_domains",
	datasourceItemsAttr:    "email_domain_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "email_domain",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_email.EmailDomainLifecycleStateActive),
	},
}

var exportEmailDkimHints = &TerraformResourceHints{
	resourceClass:          "oci_email_dkim",
	datasourceClass:        "oci_email_dkims",
	datasourceItemsAttr:    "dkim_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "dkim",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_email.DkimLifecycleStateActive),
		string(oci_email.DkimLifecycleStateNeedsAttention),
	},
}

var exportEventsRuleHints = &TerraformResourceHints{
	resourceClass:          "oci_events_rule",
	datasourceClass:        "oci_events_rules",
	datasourceItemsAttr:    "rules",
	resourceAbbreviation:   "rule",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_events.RuleLifecycleStateActive),
	},
}

var exportFileStorageFileSystemHints = &TerraformResourceHints{
	resourceClass:        "oci_file_storage_file_system",
	datasourceClass:      "oci_file_storage_file_systems",
	datasourceItemsAttr:  "file_systems",
	resourceAbbreviation: "file_system",
	discoverableLifecycleStates: []string{
		string(oci_file_storage.FileSystemLifecycleStateActive),
	},
}

var exportFileStorageMountTargetHints = &TerraformResourceHints{
	resourceClass:        "oci_file_storage_mount_target",
	datasourceClass:      "oci_file_storage_mount_targets",
	datasourceItemsAttr:  "mount_targets",
	resourceAbbreviation: "mount_target",
	discoverableLifecycleStates: []string{
		string(oci_file_storage.MountTargetLifecycleStateActive),
	},
}

var exportFileStorageExportHints = &TerraformResourceHints{
	resourceClass:          "oci_file_storage_export",
	datasourceClass:        "oci_file_storage_exports",
	datasourceItemsAttr:    "exports",
	resourceAbbreviation:   "export",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_file_storage.ExportLifecycleStateActive),
	},
}

var exportFileStorageSnapshotHints = &TerraformResourceHints{
	resourceClass:        "oci_file_storage_snapshot",
	datasourceClass:      "oci_file_storage_snapshots",
	datasourceItemsAttr:  "snapshots",
	resourceAbbreviation: "snapshot",
	discoverableLifecycleStates: []string{
		string(oci_file_storage.SnapshotLifecycleStateActive),
	},
}

var exportFunctionsApplicationHints = &TerraformResourceHints{
	resourceClass:          "oci_functions_application",
	datasourceClass:        "oci_functions_applications",
	datasourceItemsAttr:    "applications",
	resourceAbbreviation:   "application",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_functions.ApplicationLifecycleStateActive),
	},
}

var exportFunctionsFunctionHints = &TerraformResourceHints{
	resourceClass:          "oci_functions_function",
	datasourceClass:        "oci_functions_functions",
	datasourceItemsAttr:    "functions",
	resourceAbbreviation:   "function",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_functions.FunctionLifecycleStateActive),
	},
}

var exportGoldenGateDatabaseRegistrationHints = &TerraformResourceHints{
	resourceClass:          "oci_golden_gate_database_registration",
	datasourceClass:        "oci_golden_gate_database_registrations",
	datasourceItemsAttr:    "database_registration_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "database_registration",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	},
}

var exportGoldenGateDeploymentHints = &TerraformResourceHints{
	resourceClass:          "oci_golden_gate_deployment",
	datasourceClass:        "oci_golden_gate_deployments",
	datasourceItemsAttr:    "deployment_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "deployment",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	},
}

var exportGoldenGateDeploymentBackupHints = &TerraformResourceHints{
	resourceClass:          "oci_golden_gate_deployment_backup",
	datasourceClass:        "oci_golden_gate_deployment_backups",
	datasourceItemsAttr:    "deployment_backup_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "deployment_backup",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	},
}

var exportHealthChecksHttpMonitorHints = &TerraformResourceHints{
	resourceClass:          "oci_health_checks_http_monitor",
	datasourceClass:        "oci_health_checks_http_monitors",
	datasourceItemsAttr:    "http_monitors",
	resourceAbbreviation:   "http_monitor",
	requireResourceRefresh: true,
}

var exportHealthChecksPingMonitorHints = &TerraformResourceHints{
	resourceClass:          "oci_health_checks_ping_monitor",
	datasourceClass:        "oci_health_checks_ping_monitors",
	datasourceItemsAttr:    "ping_monitors",
	resourceAbbreviation:   "ping_monitor",
	requireResourceRefresh: true,
}

var exportIdentityApiKeyHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_api_key",
	datasourceClass:      "oci_identity_api_keys",
	datasourceItemsAttr:  "api_keys",
	resourceAbbreviation: "api_key",
	discoverableLifecycleStates: []string{
		string(oci_identity.ApiKeyLifecycleStateActive),
	},
}

var exportIdentityAvailabilityDomainHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_availability_domain",
	datasourceClass:      "oci_identity_availability_domains",
	datasourceItemsAttr:  "availability_domains",
	resourceAbbreviation: "availability_domain",
}

var exportIdentityAuthenticationPolicyHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_authentication_policy",
	datasourceClass:      "oci_identity_authentication_policy",
	resourceAbbreviation: "authentication_policy",
}

var exportIdentityAuthTokenHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_auth_token",
	datasourceClass:      "oci_identity_auth_tokens",
	datasourceItemsAttr:  "tokens",
	resourceAbbreviation: "auth_token",
	discoverableLifecycleStates: []string{
		string(oci_identity.AuthTokenLifecycleStateActive),
	},
}

var exportIdentityCompartmentHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_compartment",
	datasourceClass:      "oci_identity_compartments",
	datasourceItemsAttr:  "compartments",
	resourceAbbreviation: "compartment",
	discoverableLifecycleStates: []string{
		string(oci_identity.CompartmentLifecycleStateActive),
	},
}

var exportIdentityCustomerSecretKeyHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_customer_secret_key",
	datasourceClass:      "oci_identity_customer_secret_keys",
	datasourceItemsAttr:  "customer_secret_keys",
	resourceAbbreviation: "customer_secret_key",
	discoverableLifecycleStates: []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateActive),
	},
}

var exportIdentityDynamicGroupHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_dynamic_group",
	datasourceClass:      "oci_identity_dynamic_groups",
	datasourceItemsAttr:  "dynamic_groups",
	resourceAbbreviation: "dynamic_group",
	discoverableLifecycleStates: []string{
		string(oci_identity.DynamicGroupLifecycleStateActive),
	},
}

var exportIdentityGroupHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_group",
	datasourceClass:      "oci_identity_groups",
	datasourceItemsAttr:  "groups",
	resourceAbbreviation: "Group",
	discoverableLifecycleStates: []string{
		string(oci_identity.GroupLifecycleStateActive),
	},
}

var exportIdentityIdentityProviderHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_identity_provider",
	datasourceClass:      "oci_identity_identity_providers",
	datasourceItemsAttr:  "identity_providers",
	resourceAbbreviation: "identity_provider",
	discoverableLifecycleStates: []string{
		string(oci_identity.IdentityProviderLifecycleStateActive),
	},
}

var exportIdentityIdpGroupMappingHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_idp_group_mapping",
	datasourceClass:      "oci_identity_idp_group_mappings",
	datasourceItemsAttr:  "idp_group_mappings",
	resourceAbbreviation: "idp_group_mapping",
	discoverableLifecycleStates: []string{
		string(oci_identity.IdpGroupMappingLifecycleStateActive),
	},
}

var exportIdentityPolicyHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_policy",
	datasourceClass:      "oci_identity_policies",
	datasourceItemsAttr:  "policies",
	resourceAbbreviation: "policy",
	discoverableLifecycleStates: []string{
		string(oci_identity.PolicyLifecycleStateActive),
	},
}

var exportIdentitySmtpCredentialHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_smtp_credential",
	datasourceClass:      "oci_identity_smtp_credentials",
	datasourceItemsAttr:  "smtp_credentials",
	resourceAbbreviation: "smtp_credential",
	discoverableLifecycleStates: []string{
		string(oci_identity.SmtpCredentialLifecycleStateActive),
	},
}

var exportIdentitySwiftPasswordHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_swift_password",
	datasourceClass:      "oci_identity_swift_passwords",
	datasourceItemsAttr:  "passwords",
	resourceAbbreviation: "swift_password",
	discoverableLifecycleStates: []string{
		string(oci_identity.SwiftPasswordLifecycleStateActive),
	},
}

var exportIdentityUiPasswordHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_ui_password",
	datasourceClass:      "oci_identity_ui_password",
	resourceAbbreviation: "ui_password",
	discoverableLifecycleStates: []string{
		string(oci_identity.UiPasswordLifecycleStateActive),
	},
}

var exportIdentityUserGroupMembershipHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_user_group_membership",
	datasourceClass:      "oci_identity_user_group_memberships",
	datasourceItemsAttr:  "memberships",
	resourceAbbreviation: "user_group_membership",
	discoverableLifecycleStates: []string{
		string(oci_identity.UserGroupMembershipLifecycleStateActive),
	},
}

var exportIdentityUserHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_user",
	datasourceClass:      "oci_identity_users",
	datasourceItemsAttr:  "users",
	resourceAbbreviation: "user",
	discoverableLifecycleStates: []string{
		string(oci_identity.UserLifecycleStateActive),
	},
}

var exportIdentityTagDefaultHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_tag_default",
	datasourceClass:      "oci_identity_tag_defaults",
	datasourceItemsAttr:  "tag_defaults",
	resourceAbbreviation: "tag_default",
	discoverableLifecycleStates: []string{
		string(oci_identity.TagDefaultLifecycleStateActive),
	},
}

var exportIdentityTagNamespaceHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_tag_namespace",
	datasourceClass:      "oci_identity_tag_namespaces",
	datasourceItemsAttr:  "tag_namespaces",
	resourceAbbreviation: "tag_namespace",
	discoverableLifecycleStates: []string{
		string(oci_identity.TagNamespaceLifecycleStateActive),
	},
}

var exportIdentityTagHints = &TerraformResourceHints{
	resourceClass:          "oci_identity_tag",
	datasourceClass:        "oci_identity_tags",
	datasourceItemsAttr:    "tags",
	resourceAbbreviation:   "tag",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_identity.TagLifecycleStateActive),
	},
}

var exportIdentityDomainHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_domain",
	datasourceClass:      "oci_identity_domains",
	datasourceItemsAttr:  "domains",
	resourceAbbreviation: "domain",
	discoverableLifecycleStates: []string{
		string(oci_identity.DomainLifecycleStateActive),
	},
}

var exportIdentityDbCredentialHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_db_credential",
	datasourceClass:      "oci_identity_db_credentials",
	datasourceItemsAttr:  "db_credentials",
	resourceAbbreviation: "db_credential",
	discoverableLifecycleStates: []string{
		string(oci_identity.DbCredentialLifecycleStateActive),
	},
}

var exportIdentityImportStandardTagsManagementHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_import_standard_tags_management",
	resourceAbbreviation: "import_standard_tags_management",
}

var exportIdentityDataPlaneGenerateScopedAccessTokenHints = &TerraformResourceHints{
	resourceClass:        "oci_identity_data_plane_generate_scoped_access_token",
	resourceAbbreviation: "generate_scoped_access_token",
}

var exportIntegrationIntegrationInstanceHints = &TerraformResourceHints{
	resourceClass:          "oci_integration_integration_instance",
	datasourceClass:        "oci_integration_integration_instances",
	datasourceItemsAttr:    "integration_instances",
	resourceAbbreviation:   "integration_instance",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_integration.IntegrationInstanceLifecycleStateActive),
	},
}

var exportJmsFleetHints = &TerraformResourceHints{
	resourceClass:          "oci_jms_fleet",
	datasourceClass:        "oci_jms_fleets",
	datasourceItemsAttr:    "fleet_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "fleet",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_jms.LifecycleStateActive),
	},
}

var exportKmsKeyHints = &TerraformResourceHints{
	resourceClass:          "oci_kms_key",
	datasourceClass:        "oci_kms_keys",
	datasourceItemsAttr:    "keys",
	resourceAbbreviation:   "key",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_kms.KeyLifecycleStateEnabled),
	},
}

var exportKmsKeyVersionHints = &TerraformResourceHints{
	resourceClass:        "oci_kms_key_version",
	datasourceClass:      "oci_kms_key_versions",
	datasourceItemsAttr:  "key_versions",
	resourceAbbreviation: "key_version",
	discoverableLifecycleStates: []string{
		string(oci_kms.KeyVersionLifecycleStateEnabled),
	},
}

var exportKmsVaultHints = &TerraformResourceHints{
	resourceClass:        "oci_kms_vault",
	datasourceClass:      "oci_kms_vaults",
	datasourceItemsAttr:  "vaults",
	resourceAbbreviation: "vault",
	discoverableLifecycleStates: []string{
		string(oci_kms.VaultLifecycleStateActive),
	},
}

var exportKmsSignHints = &TerraformResourceHints{
	resourceClass:        "oci_kms_sign",
	resourceAbbreviation: "sign",
}

var exportKmsVerifyHints = &TerraformResourceHints{
	resourceClass:        "oci_kms_verify",
	resourceAbbreviation: "verify",
}

var exportKmsCreateReplicaHints = &TerraformResourceHints{
	resourceClass:        "oci_kms_vault_replication",
	resourceAbbreviation: "vault_replication",
}

var exportKmsDeleteReplicaHints = &TerraformResourceHints{
	resourceClass:        "oci_kms_vault_replication",
	resourceAbbreviation: "vault_replication",
}

var exportIdentityNetworkSourceHints = &TerraformResourceHints{
	resourceClass:          "oci_identity_network_source",
	datasourceClass:        "oci_identity_network_sources",
	datasourceItemsAttr:    "network_sources",
	resourceAbbreviation:   "network_source",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_identity.NetworkSourcesLifecycleStateActive),
	},
}

var exportLimitsQuotaHints = &TerraformResourceHints{
	resourceClass:          "oci_limits_quota",
	datasourceClass:        "oci_limits_quotas",
	datasourceItemsAttr:    "quotas",
	resourceAbbreviation:   "quota",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_limits.QuotaLifecycleStateActive),
	},
}

var exportLoadBalancerBackendHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_backend",
	datasourceClass:      "oci_load_balancer_backends",
	datasourceItemsAttr:  "backends",
	resourceAbbreviation: "backend",
}

var exportLoadBalancerBackendSetHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_backend_set",
	datasourceClass:      "oci_load_balancer_backend_sets",
	datasourceItemsAttr:  "backendsets",
	resourceAbbreviation: "backend_set",
}

var exportLoadBalancerCertificateHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_certificate",
	datasourceClass:      "oci_load_balancer_certificates",
	datasourceItemsAttr:  "certificates",
	resourceAbbreviation: "certificate",
}

var exportLoadBalancerHostnameHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_hostname",
	datasourceClass:      "oci_load_balancer_hostnames",
	datasourceItemsAttr:  "hostnames",
	resourceAbbreviation: "hostname",
}

var exportLoadBalancerListenerHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_listener",
	resourceAbbreviation: "listener",
}

var exportLoadBalancerLoadBalancerHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_load_balancer",
	datasourceClass:      "oci_load_balancer_load_balancers",
	datasourceItemsAttr:  "load_balancers",
	resourceAbbreviation: "load_balancer",
	discoverableLifecycleStates: []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateActive),
	},
}

var exportLoadBalancerPathRouteSetHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_path_route_set",
	datasourceClass:      "oci_load_balancer_path_route_sets",
	datasourceItemsAttr:  "path_route_sets",
	resourceAbbreviation: "path_route_set",
}

var exportLoadBalancerLoadBalancerRoutingPolicyHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_load_balancer_routing_policy",
	datasourceClass:      "oci_load_balancer_load_balancer_routing_policies",
	datasourceItemsAttr:  "routing_policies",
	resourceAbbreviation: "load_balancer_routing_policy",
}

var exportLoadBalancerRuleSetHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_rule_set",
	datasourceClass:      "oci_load_balancer_rule_sets",
	datasourceItemsAttr:  "rule_sets",
	resourceAbbreviation: "rule_set",
}

var exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints = &TerraformResourceHints{
	resourceClass:          "oci_log_analytics_log_analytics_object_collection_rule",
	datasourceClass:        "oci_log_analytics_log_analytics_object_collection_rules",
	datasourceItemsAttr:    "log_analytics_object_collection_rule_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "log_analytics_object_collection_rule",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_log_analytics.ObjectCollectionRuleLifecycleStatesActive),
	},
}

var exportLogAnalyticsNamespaceScheduledTaskHints = &TerraformResourceHints{
	resourceClass:          "oci_log_analytics_namespace_scheduled_task",
	datasourceClass:        "oci_log_analytics_namespace_scheduled_tasks",
	datasourceItemsAttr:    "scheduled_task_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "namespace_scheduled_task",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_log_analytics.ScheduledTaskLifecycleStateActive),
	},
}

var exportLogAnalyticsLogAnalyticsImportCustomContentHints = &TerraformResourceHints{
	resourceClass:        "oci_log_analytics_log_analytics_import_custom_content",
	resourceAbbreviation: "log_analytics_import_custom_content",
}

var exportLogAnalyticsLogAnalyticsPreferencesManagementHints = &TerraformResourceHints{
	resourceClass:        "oci_log_analytics_log_analytics_preferences_management",
	resourceAbbreviation: "log_analytics_preferences_management",
}

var exportLogAnalyticsLogAnalyticsUnprocessedDataBucketManagementHints = &TerraformResourceHints{
	resourceClass:        "oci_log_analytics_log_analytics_unprocessed_data_bucket_management",
	resourceAbbreviation: "log_analytics_unprocessed_data_bucket_management",
}

var exportLogAnalyticsLogAnalyticsResourceCategoriesManagementHints = &TerraformResourceHints{
	resourceClass:        "oci_log_analytics_log_analytics_resource_categories_management",
	resourceAbbreviation: "log_analytics_resource_categories_management",
}

var exportLoggingLogGroupHints = &TerraformResourceHints{
	resourceClass:        "oci_logging_log_group",
	datasourceClass:      "oci_logging_log_groups",
	datasourceItemsAttr:  "log_groups",
	resourceAbbreviation: "log_group",
	discoverableLifecycleStates: []string{
		string(oci_logging.LogGroupLifecycleStateActive),
	},
}

var exportLoggingLogHints = &TerraformResourceHints{
	resourceClass:        "oci_logging_log",
	datasourceClass:      "oci_logging_logs",
	datasourceItemsAttr:  "logs",
	resourceAbbreviation: "log",
	discoverableLifecycleStates: []string{
		string(oci_logging.LogLifecycleStateActive),
	},
}

var exportLoggingUnifiedAgentConfigurationHints = &TerraformResourceHints{
	resourceClass:          "oci_logging_unified_agent_configuration",
	datasourceClass:        "oci_logging_unified_agent_configurations",
	datasourceItemsAttr:    "unified_agent_configuration_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "unified_agent_configuration",
	requireResourceRefresh: true,
}

var exportManagementAgentManagementAgentHints = &TerraformResourceHints{
	resourceClass:        "oci_management_agent_management_agent",
	datasourceClass:      "oci_management_agent_management_agents",
	datasourceItemsAttr:  "management_agents",
	resourceAbbreviation: "management_agent",
	discoverableLifecycleStates: []string{
		string(oci_management_agent.LifecycleStatesActive),
	},
}

var exportManagementAgentManagementAgentInstallKeyHints = &TerraformResourceHints{
	resourceClass:        "oci_management_agent_management_agent_install_key",
	datasourceClass:      "oci_management_agent_management_agent_install_keys",
	datasourceItemsAttr:  "management_agent_install_keys",
	resourceAbbreviation: "management_agent_install_key",
	discoverableLifecycleStates: []string{
		string(oci_management_agent.LifecycleStatesActive),
	},
}

var exportMarketplaceAcceptedAgreementHints = &TerraformResourceHints{
	resourceClass:          "oci_marketplace_accepted_agreement",
	datasourceClass:        "oci_marketplace_accepted_agreements",
	datasourceItemsAttr:    "accepted_agreements",
	resourceAbbreviation:   "accepted_agreement",
	requireResourceRefresh: true,
}

var exportMarketplacePublicationHints = &TerraformResourceHints{
	resourceClass:          "oci_marketplace_publication",
	datasourceClass:        "oci_marketplace_publications",
	datasourceItemsAttr:    "publications",
	resourceAbbreviation:   "publication",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_marketplace.PublicationLifecycleStateActive),
	},
}

var exportMeteringComputationQueryHints = &TerraformResourceHints{
	resourceClass:          "oci_metering_computation_query",
	datasourceClass:        "oci_metering_computation_queries",
	isDatasourceCollection: true,
	datasourceItemsAttr:    "query_collection",
	resourceAbbreviation:   "query",
	requireResourceRefresh: true,
}

var exportMeteringComputationCustomTableHints = &TerraformResourceHints{
	resourceClass:          "oci_metering_computation_custom_table",
	datasourceClass:        "oci_metering_computation_custom_tables",
	datasourceItemsAttr:    "custom_table_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "custom_table",
	requireResourceRefresh: true,
}

var exportMonitoringAlarmHints = &TerraformResourceHints{
	resourceClass:          "oci_monitoring_alarm",
	datasourceClass:        "oci_monitoring_alarms",
	datasourceItemsAttr:    "alarms",
	resourceAbbreviation:   "alarm",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_monitoring.AlarmLifecycleStateActive),
	},
}

var exportMysqlHeatWaveClusterHints = &TerraformResourceHints{
	resourceClass:        "oci_mysql_heat_wave_cluster",
	datasourceClass:      "oci_mysql_heat_wave_cluster",
	resourceAbbreviation: "heat_wave_cluster",
	discoverableLifecycleStates: []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateActive),
	},
}

var exportMysqlMysqlBackupHints = &TerraformResourceHints{
	resourceClass:        "oci_mysql_mysql_backup",
	datasourceClass:      "oci_mysql_mysql_backups",
	datasourceItemsAttr:  "backups",
	resourceAbbreviation: "mysql_backup",
	discoverableLifecycleStates: []string{
		string(oci_mysql.BackupLifecycleStateActive),
	},
}

var exportMysqlMysqlDbSystemHints = &TerraformResourceHints{
	resourceClass:          "oci_mysql_mysql_db_system",
	datasourceClass:        "oci_mysql_mysql_db_systems",
	datasourceItemsAttr:    "db_systems",
	resourceAbbreviation:   "mysql_db_system",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_mysql.DbSystemLifecycleStateActive),
	},
}

var exportMysqlChannelHints = &TerraformResourceHints{
	resourceClass:          "oci_mysql_channel",
	datasourceClass:        "oci_mysql_channels",
	datasourceItemsAttr:    "channels",
	resourceAbbreviation:   "channel",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_mysql.ChannelLifecycleStateActive),
		string(oci_mysql.ChannelLifecycleStateNeedsAttention),
	},
}

var exportNetworkLoadBalancerNetworkLoadBalancerHints = &TerraformResourceHints{
	resourceClass:          "oci_network_load_balancer_network_load_balancer",
	datasourceClass:        "oci_network_load_balancer_network_load_balancers",
	datasourceItemsAttr:    "network_load_balancer_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "network_load_balancer",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_network_load_balancer.LifecycleStateActive),
	},
}

var exportNetworkLoadBalancerBackendSetHints = &TerraformResourceHints{
	resourceClass:          "oci_network_load_balancer_backend_set",
	datasourceClass:        "oci_network_load_balancer_backend_sets",
	datasourceItemsAttr:    "backend_set_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "backend_set",
	requireResourceRefresh: true,
}

var exportNetworkLoadBalancerBackendHints = &TerraformResourceHints{
	resourceClass:          "oci_network_load_balancer_backend",
	datasourceClass:        "oci_network_load_balancer_backends",
	datasourceItemsAttr:    "backend_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "backend",
	requireResourceRefresh: true,
}

var exportNetworkLoadBalancerListenerHints = &TerraformResourceHints{
	resourceClass:          "oci_network_load_balancer_listener",
	datasourceClass:        "oci_network_load_balancer_listeners",
	datasourceItemsAttr:    "listener_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "listener",
	requireResourceRefresh: true,
}

var exportNosqlTableHints = &TerraformResourceHints{
	resourceClass:          "oci_nosql_table",
	datasourceClass:        "oci_nosql_tables",
	datasourceItemsAttr:    "table_collection",
	resourceAbbreviation:   "table",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_nosql.TableLifecycleStateActive),
	},
}

var exportNosqlIndexHints = &TerraformResourceHints{
	resourceClass:          "oci_nosql_index",
	datasourceClass:        "oci_nosql_indexes",
	datasourceItemsAttr:    "index_collection",
	resourceAbbreviation:   "index",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_nosql.IndexLifecycleStateActive),
	},
}

var exportObjectStorageBucketHints = &TerraformResourceHints{
	resourceClass:          "oci_objectstorage_bucket",
	datasourceClass:        "oci_objectstorage_bucket_summaries",
	datasourceItemsAttr:    "bucket_summaries",
	resourceAbbreviation:   "bucket",
	requireResourceRefresh: true,
}

var exportObjectStorageObjectLifecyclePolicyHints = &TerraformResourceHints{
	resourceClass:        "oci_objectstorage_object_lifecycle_policy",
	datasourceClass:      "oci_objectstorage_object_lifecycle_policy",
	resourceAbbreviation: "object_lifecycle_policy",
}

var exportObjectStorageNamespaceHints = &TerraformResourceHints{
	resourceClass:        "oci_objectstorage_namespace",
	datasourceClass:      "oci_objectstorage_namespace",
	resourceAbbreviation: "namespace",
}

var exportObjectStorageObjectHints = &TerraformResourceHints{
	resourceClass:        "oci_objectstorage_object",
	datasourceClass:      "oci_objectstorage_objects",
	datasourceItemsAttr:  "objects",
	resourceAbbreviation: "object",
}

var exportObjectStoragePreauthenticatedRequestHints = &TerraformResourceHints{
	resourceClass:        "oci_objectstorage_preauthrequest",
	datasourceClass:      "oci_objectstorage_preauthrequests",
	datasourceItemsAttr:  "preauthenticated_requests",
	resourceAbbreviation: "preauthenticated_request",
}

var exportObjectStorageReplicationPolicyHints = &TerraformResourceHints{
	resourceClass:        "oci_objectstorage_replication_policy",
	datasourceClass:      "oci_objectstorage_replication_policies",
	datasourceItemsAttr:  "replication_policies",
	resourceAbbreviation: "replication_policy",
}

var exportOceOceInstanceHints = &TerraformResourceHints{
	resourceClass:          "oci_oce_oce_instance",
	datasourceClass:        "oci_oce_oce_instances",
	datasourceItemsAttr:    "oce_instances",
	resourceAbbreviation:   "oce_instance",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_oce.OceInstanceLifecycleStateActive),
	},
}

var exportOcvpSddcHints = &TerraformResourceHints{
	resourceClass:          "oci_ocvp_sddc",
	datasourceClass:        "oci_ocvp_sddcs",
	datasourceItemsAttr:    "sddc_collection",
	resourceAbbreviation:   "sddc",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_ocvp.LifecycleStatesActive),
		string(oci_ocvp.LifecycleStatesFailed),
	},
}

var exportOcvpEsxiHostHints = &TerraformResourceHints{
	resourceClass:          "oci_ocvp_esxi_host",
	datasourceClass:        "oci_ocvp_esxi_hosts",
	datasourceItemsAttr:    "esxi_host_collection",
	resourceAbbreviation:   "esxi_host",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_ocvp.LifecycleStatesActive),
		string(oci_ocvp.LifecycleStatesFailed),
	},
}

var exportOdaOdaInstanceHints = &TerraformResourceHints{
	resourceClass:        "oci_oda_oda_instance",
	datasourceClass:      "oci_oda_oda_instances",
	datasourceItemsAttr:  "oda_instances",
	resourceAbbreviation: "oda_instance",
	discoverableLifecycleStates: []string{
		string(oci_oda.OdaInstanceLifecycleStateActive),
	},
}

var exportOnsNotificationTopicHints = &TerraformResourceHints{
	resourceClass:        "oci_ons_notification_topic",
	datasourceClass:      "oci_ons_notification_topics",
	datasourceItemsAttr:  "notification_topics",
	resourceAbbreviation: "notification_topic",
	discoverableLifecycleStates: []string{
		string(oci_ons.NotificationTopicLifecycleStateActive),
	},
}

var exportOnsSubscriptionHints = &TerraformResourceHints{
	resourceClass:        "oci_ons_subscription",
	datasourceClass:      "oci_ons_subscriptions",
	datasourceItemsAttr:  "subscriptions",
	resourceAbbreviation: "subscription",
	discoverableLifecycleStates: []string{
		string(oci_ons.SubscriptionLifecycleStatePending),
		string(oci_ons.SubscriptionLifecycleStateActive),
	},
}

var exportOperatorAccessControlOperatorControlHints = &TerraformResourceHints{
	resourceClass:          "oci_operator_access_control_operator_control",
	datasourceClass:        "oci_operator_access_control_operator_controls",
	datasourceItemsAttr:    "operator_control_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "operator_control",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_operator_access_control.OperatorControlLifecycleStatesCreated),
		string(oci_operator_access_control.OperatorControlLifecycleStatesAssigned),
	},
}

var exportOperatorAccessControlOperatorControlAssignmentHints = &TerraformResourceHints{
	resourceClass:          "oci_operator_access_control_operator_control_assignment",
	datasourceClass:        "oci_operator_access_control_operator_control_assignments",
	datasourceItemsAttr:    "operator_control_assignment_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "operator_control_assignment",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesCreated),
	},
}

var exportOpsiEnterpriseManagerBridgeHints = &TerraformResourceHints{
	resourceClass:          "oci_opsi_enterprise_manager_bridge",
	datasourceClass:        "oci_opsi_enterprise_manager_bridges",
	datasourceItemsAttr:    "enterprise_manager_bridge_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "enterprise_manager_bridge",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_opsi.LifecycleStateActive),
		string(oci_opsi.LifecycleStateNeedsAttention),
	},
}

var exportOpsiDatabaseInsightHints = &TerraformResourceHints{
	resourceClass:          "oci_opsi_database_insight",
	datasourceClass:        "oci_opsi_database_insights",
	datasourceItemsAttr:    "database_insights_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "database_insight",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_opsi.LifecycleStateActive),
	},
}

var exportOpsiHostInsightHints = &TerraformResourceHints{
	resourceClass:          "oci_opsi_host_insight",
	datasourceClass:        "oci_opsi_host_insights",
	datasourceItemsAttr:    "host_insight_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "host_insight",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_opsi.LifecycleStateActive),
	},
}

var exportOpsiAwrHubHints = &TerraformResourceHints{
	resourceClass:          "oci_opsi_awr_hub",
	datasourceClass:        "oci_opsi_awr_hubs",
	datasourceItemsAttr:    "awr_hub_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "awr_hub",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_opsi.AwrHubLifecycleStateActive),
	},
}

var exportOpsiExadataInsightHints = &TerraformResourceHints{
	resourceClass:          "oci_opsi_exadata_insight",
	datasourceClass:        "oci_opsi_exadata_insights",
	datasourceItemsAttr:    "exadata_insight_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "exadata_insight",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_opsi.ExadataInsightLifecycleStateActive),
	},
}

var exportOpsiOperationsInsightsWarehouseUserHints = &TerraformResourceHints{
	resourceClass:          "oci_opsi_operations_insights_warehouse_user",
	datasourceClass:        "oci_opsi_operations_insights_warehouse_users",
	datasourceItemsAttr:    "operations_insights_warehouse_user_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "operations_insights_warehouse_user",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_opsi.OperationsInsightsWarehouseUserLifecycleStateActive),
	},
}

var exportOpsiOperationsInsightsWarehouseHints = &TerraformResourceHints{
	resourceClass:          "oci_opsi_operations_insights_warehouse",
	datasourceClass:        "oci_opsi_operations_insights_warehouses",
	datasourceItemsAttr:    "operations_insights_warehouse_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "operations_insights_warehouse",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_opsi.OperationsInsightsWarehouseLifecycleStateActive),
	},
}

var exportOpsiOperationsInsightsWarehouseDownloadWarehouseWalletHints = &TerraformResourceHints{
	resourceClass:        "oci_opsi_operations_insights_warehouse_download_warehouse_wallet",
	resourceAbbreviation: "operations_insights_warehouse_download_warehouse_wallet",
}

var exportOpsiOperationsInsightsWarehouseRotateWarehouseWalletHints = &TerraformResourceHints{
	resourceClass:        "oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet",
	resourceAbbreviation: "operations_insights_warehouse_rotate_warehouse_wallet",
}

var exportOptimizerProfileHints = &TerraformResourceHints{
	resourceClass:          "oci_optimizer_profile",
	datasourceClass:        "oci_optimizer_profiles",
	datasourceItemsAttr:    "profile_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "profile",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_optimizer.LifecycleStateActive),
	},
}

var exportOsmanagementManagedInstanceHints = &TerraformResourceHints{
	resourceClass:          "oci_osmanagement_managed_instance",
	datasourceClass:        "oci_osmanagement_managed_instances",
	datasourceItemsAttr:    "managed_instances",
	resourceAbbreviation:   "managed_instance",
	requireResourceRefresh: true,
}

var exportOsmanagementManagedInstanceGroupHints = &TerraformResourceHints{
	resourceClass:        "oci_osmanagement_managed_instance_group",
	datasourceClass:      "oci_osmanagement_managed_instance_groups",
	datasourceItemsAttr:  "managed_instance_groups",
	resourceAbbreviation: "managed_instance_group",
	discoverableLifecycleStates: []string{
		string(oci_osmanagement.ListManagedInstanceGroupsLifecycleStateActive),
	},
}

var exportOsmanagementSoftwareSourceHints = &TerraformResourceHints{
	resourceClass:          "oci_osmanagement_software_source",
	datasourceClass:        "oci_osmanagement_software_sources",
	datasourceItemsAttr:    "software_sources",
	resourceAbbreviation:   "software_source",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_osmanagement.ListSoftwareSourcesLifecycleStateActive),
	},
}

var exportSchServiceConnectorHints = &TerraformResourceHints{
	resourceClass:          "oci_sch_service_connector",
	datasourceClass:        "oci_sch_service_connectors",
	datasourceItemsAttr:    "service_connector_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "service_connector",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_sch.LifecycleStateActive),
	},
}

var exportStreamingConnectHarnessHints = &TerraformResourceHints{
	resourceClass:        "oci_streaming_connect_harness",
	datasourceClass:      "oci_streaming_connect_harnesses",
	datasourceItemsAttr:  "connect_harness",
	resourceAbbreviation: "connect_harness",
	discoverableLifecycleStates: []string{
		string(oci_streaming.ConnectHarnessLifecycleStateActive),
	},
}

var exportStreamingStreamPoolHints = &TerraformResourceHints{
	resourceClass:          "oci_streaming_stream_pool",
	datasourceClass:        "oci_streaming_stream_pools",
	datasourceItemsAttr:    "stream_pools",
	resourceAbbreviation:   "stream_pool",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_streaming.StreamPoolLifecycleStateActive),
	},
}

var exportStreamingStreamHints = &TerraformResourceHints{
	resourceClass:          "oci_streaming_stream",
	datasourceClass:        "oci_streaming_streams",
	datasourceItemsAttr:    "streams",
	resourceAbbreviation:   "stream",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_streaming.StreamLifecycleStateActive),
	},
}
var exportUsageProxySubscriptionRedeemableUserHints = &TerraformResourceHints{
	resourceClass:          "oci_usage_proxy_subscription_redeemable_user",
	datasourceClass:        "oci_usage_proxy_subscription_redeemable_users",
	datasourceItemsAttr:    "redeemable_user_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "subscription_redeemable_user",
}

var exportVisualBuilderVbInstanceHints = &TerraformResourceHints{
	resourceClass:          "oci_visual_builder_vb_instance",
	datasourceClass:        "oci_visual_builder_vb_instances",
	datasourceItemsAttr:    "vb_instance_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "vb_instance",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_visual_builder.VbInstanceLifecycleStateActive),
	},
}

var exportVulnerabilityScanningHostScanRecipeHints = &TerraformResourceHints{
	resourceClass:          "oci_vulnerability_scanning_host_scan_recipe",
	datasourceClass:        "oci_vulnerability_scanning_host_scan_recipes",
	datasourceItemsAttr:    "host_scan_recipe_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "host_scan_recipe",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_vulnerability_scanning.LifecycleStateActive),
	},
}

var exportVulnerabilityScanningHostScanTargetHints = &TerraformResourceHints{
	resourceClass:          "oci_vulnerability_scanning_host_scan_target",
	datasourceClass:        "oci_vulnerability_scanning_host_scan_targets",
	datasourceItemsAttr:    "host_scan_target_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "host_scan_target",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_vulnerability_scanning.LifecycleStateActive),
	},
}

var exportVulnerabilityScanningContainerScanRecipeHints = &TerraformResourceHints{
	resourceClass:          "oci_vulnerability_scanning_container_scan_recipe",
	datasourceClass:        "oci_vulnerability_scanning_container_scan_recipes",
	datasourceItemsAttr:    "container_scan_recipe_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "container_scan_recipe",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_vulnerability_scanning.LifecycleStateActive),
	},
}

var exportVulnerabilityScanningContainerScanTargetHints = &TerraformResourceHints{
	resourceClass:          "oci_vulnerability_scanning_container_scan_target",
	datasourceClass:        "oci_vulnerability_scanning_container_scan_targets",
	datasourceItemsAttr:    "container_scan_target_summary_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "container_scan_target",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_vulnerability_scanning.LifecycleStateActive),
	},
}

var exportWaasAddressListHints = &TerraformResourceHints{
	resourceClass:          "oci_waas_address_list",
	datasourceClass:        "oci_waas_address_lists",
	datasourceItemsAttr:    "address_lists",
	resourceAbbreviation:   "address_list",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_waas.LifecycleStatesActive),
	},
}

var exportWaasCustomProtectionRuleHints = &TerraformResourceHints{
	resourceClass:          "oci_waas_custom_protection_rule",
	datasourceClass:        "oci_waas_custom_protection_rules",
	datasourceItemsAttr:    "custom_protection_rules",
	resourceAbbreviation:   "custom_protection_rule",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_waas.LifecycleStatesActive),
	},
}

var exportWaasHttpRedirectHints = &TerraformResourceHints{
	resourceClass:        "oci_waas_http_redirect",
	datasourceClass:      "oci_waas_http_redirects",
	datasourceItemsAttr:  "http_redirects",
	resourceAbbreviation: "http_redirect",
	discoverableLifecycleStates: []string{
		string(oci_waas.LifecycleStatesActive),
	},
}

var exportWaasWaasPolicyHints = &TerraformResourceHints{
	resourceClass:          "oci_waas_waas_policy",
	datasourceClass:        "oci_waas_waas_policies",
	datasourceItemsAttr:    "waas_policies",
	resourceAbbreviation:   "waas_policy",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_waas.WaasPolicyLifecycleStateActive),
	},
}

var exportWafWebAppFirewallPolicyHints = &TerraformResourceHints{
	resourceClass:          "oci_waf_web_app_firewall_policy",
	datasourceClass:        "oci_waf_web_app_firewall_policies",
	datasourceItemsAttr:    "web_app_firewall_policy_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "web_app_firewall_policy",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_waf.WebAppFirewallPolicyLifecycleStateActive),
	},
}

var exportWafWebAppFirewallHints = &TerraformResourceHints{
	resourceClass:          "oci_waf_web_app_firewall",
	datasourceClass:        "oci_waf_web_app_firewalls",
	datasourceItemsAttr:    "web_app_firewall_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "web_app_firewall",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_waf.WebAppFirewallLifecycleStateActive),
	},
}

var exportWafNetworkAddressListHints = &TerraformResourceHints{
	resourceClass:          "oci_waf_network_address_list",
	datasourceClass:        "oci_waf_network_address_lists",
	datasourceItemsAttr:    "network_address_list_collection",
	isDatasourceCollection: true,
	resourceAbbreviation:   "network_address_list",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_waf.NetworkAddressListLifecycleStateActive),
	},
}
