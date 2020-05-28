// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_bds "github.com/oracle/oci-go-sdk/bds"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_database "github.com/oracle/oci-go-sdk/database"
	oci_events "github.com/oracle/oci-go-sdk/events"
	oci_functions "github.com/oracle/oci-go-sdk/functions"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
	oci_limits "github.com/oracle/oci-go-sdk/limits"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
	oci_nosql "github.com/oracle/oci-go-sdk/nosql"
	oci_osmanagement "github.com/oracle/oci-go-sdk/osmanagement"
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
)

// Hints for discovering and exporting this resource to configuration and state files
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

var exportCoreBootVolumeHints = &TerraformResourceHints{
	resourceClass:        "oci_core_boot_volume",
	datasourceClass:      "oci_core_boot_volumes",
	datasourceItemsAttr:  "boot_volumes",
	resourceAbbreviation: "boot_volume",
	discoverableLifecycleStates: []string{
		string(oci_core.BootVolumeLifecycleStateAvailable),
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
