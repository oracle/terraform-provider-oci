// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_apigateway "github.com/oracle/oci-go-sdk/apigateway"
	oci_bds "github.com/oracle/oci-go-sdk/bds"
	oci_budget "github.com/oracle/oci-go-sdk/budget"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_database "github.com/oracle/oci-go-sdk/database"
	oci_datacatalog "github.com/oracle/oci-go-sdk/datacatalog"
	oci_dataflow "github.com/oracle/oci-go-sdk/dataflow"
	oci_dns "github.com/oracle/oci-go-sdk/dns"
	oci_email "github.com/oracle/oci-go-sdk/email"
	oci_events "github.com/oracle/oci-go-sdk/events"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
	oci_functions "github.com/oracle/oci-go-sdk/functions"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
	oci_limits "github.com/oracle/oci-go-sdk/limits"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
	oci_monitoring "github.com/oracle/oci-go-sdk/monitoring"
	oci_nosql "github.com/oracle/oci-go-sdk/nosql"
	oci_osmanagement "github.com/oracle/oci-go-sdk/osmanagement"
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
	oci_waas "github.com/oracle/oci-go-sdk/waas"
)

// Hints for discovering and exporting this resource to configuration and state files
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

var exportAutoScalingAutoScalingConfigurationHints = &TerraformResourceHints{
	resourceClass:          "oci_autoscaling_auto_scaling_configuration",
	datasourceClass:        "oci_autoscaling_auto_scaling_configurations",
	datasourceItemsAttr:    "auto_scaling_configurations",
	resourceAbbreviation:   "auto_scaling_configuration",
	requireResourceRefresh: true,
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

var exportDatabaseAutonomousContainerDatabaseHints = &TerraformResourceHints{
	resourceClass:        "oci_database_autonomous_container_database",
	datasourceClass:      "oci_database_autonomous_container_databases",
	datasourceItemsAttr:  "autonomous_container_databases",
	resourceAbbreviation: "autonomous_container_database",
	discoverableLifecycleStates: []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
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

var exportDnsRecordHints = &TerraformResourceHints{
	resourceClass:        "oci_dns_record",
	datasourceClass:      "oci_dns_records",
	datasourceItemsAttr:  "records",
	resourceAbbreviation: "record",
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
	resourceAbbreviation:   "connection",
	requireResourceRefresh: true,
	discoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
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
	resourceAbbreviation: "group",
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

var exportLoadBalancerRuleSetHints = &TerraformResourceHints{
	resourceClass:        "oci_load_balancer_rule_set",
	datasourceClass:      "oci_load_balancer_rule_sets",
	datasourceItemsAttr:  "rule_sets",
	resourceAbbreviation: "rule_set",
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

var exportObjectStorageNamespaceHints = &TerraformResourceHints{
	resourceClass:        "oci_objectstorage_namespace",
	datasourceClass:      "oci_objectstorage_namespace",
	resourceAbbreviation: "namespace",
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
