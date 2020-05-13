package oci

var tenancyResourceGraphs = map[string]TerraformResourceGraph{
	"identity": identityResourceGraph,
	"limits":   limitsResourceGraph,
}

var availabilityDomainsGraph = TerraformResourceGraph{
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
	},
}

var compartmentResourceGraphs = map[string]TerraformResourceGraph{
	"availability_domain": availabilityDomainsGraph,
	"auto_scaling":        autoScalingResourceGraph,
	"bds":                 bdsResourceGraph,
	"core":                coreResourceGraph,
	"database":            databaseResourceGraph,
	"functions":           functionsResourceGraph,
	"load_balancer":       loadBalancerResourceGraph,
	"object_storage":      objectStorageResourceGraph,
	"tagging":             taggingResourceGraph,
}

var autoScalingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAutoScalingAutoScalingConfigurationHints},
	},
}

var bdsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBdsBdsInstanceHints},
	},
}

var coreResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCoreCpeHints},
		{TerraformResourceHints: exportCoreCrossConnectGroupHints},
		{TerraformResourceHints: exportCoreCrossConnectHints},
		{TerraformResourceHints: exportCoreDrgAttachmentHints},
		{TerraformResourceHints: exportCoreDrgHints},
		{TerraformResourceHints: exportCoreImageHints},
		{TerraformResourceHints: exportCoreInstanceConfigurationHints},
		{TerraformResourceHints: exportCoreInstancePoolHints},
		{TerraformResourceHints: exportCoreInstanceHints},
		{TerraformResourceHints: exportCoreIpSecConnectionHints},
		{TerraformResourceHints: exportCoreNetworkSecurityGroupHints},
		{TerraformResourceHints: exportCoreRemotePeeringConnectionHints},
		{TerraformResourceHints: exportCoreServiceGatewayHints},
		{TerraformResourceHints: exportCoreVcnHints},
		{TerraformResourceHints: exportCoreVirtualCircuitHints},
		{TerraformResourceHints: exportCoreVnicAttachmentHints},
		{TerraformResourceHints: exportCoreVolumeAttachmentHints},
		{TerraformResourceHints: exportCoreVolumeGroupHints},
		{TerraformResourceHints: exportCoreVolumeHints},
	},
	"oci_core_boot_volume": {
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			datasourceQueryParams: map[string]string{
				"asset_id": "id",
			},
		},
	},
	"oci_core_instance": {
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
	"oci_core_vcn": {
		{
			TerraformResourceHints: exportCoreDhcpOptionsHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreInternetGatewayHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreLocalPeeringGatewayHints,
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
			TerraformResourceHints: exportCoreSubnetHints,
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
}

var databaseResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseAutonomousContainerDatabaseHints},
		{TerraformResourceHints: exportDatabaseAutonomousDatabaseHints},
		{TerraformResourceHints: exportDatabaseAutonomousExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseDbSystemHints},
	},
	"oci_database_db_system": {
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			datasourceQueryParams: map[string]string{
				"db_system_id": "id",
			},
		},
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

var identityResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportIdentityAuthenticationPolicyHints},
		{
			TerraformResourceHints: exportIdentityCompartmentHints,
			datasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
		{TerraformResourceHints: exportIdentityDynamicGroupHints},
		{TerraformResourceHints: exportIdentityGroupHints},
		{
			TerraformResourceHints: exportIdentityIdentityProviderHints,
			datasourceQueryParams:  map[string]string{"protocol": "'SAML2'"},
		},
		{
			TerraformResourceHints: exportIdentityPolicyHints,
			datasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
		{TerraformResourceHints: exportIdentityUserHints},
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
			TerraformResourceHints: exportIdentitySmtpCredentialHints,
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
		{
			TerraformResourceHints: exportLoadBalancerBackendSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerCertificateHints,
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

var objectStorageResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportObjectStorageNamespaceHints},
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
