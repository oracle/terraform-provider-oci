package network_firewall

import (
	"fmt"

	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportNetworkFirewallNetworkFirewallPolicyAddressListHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyAddressListId
	exportNetworkFirewallNetworkFirewallPolicyUrlListHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyUrlListId
	exportNetworkFirewallNetworkFirewallPolicyMappedSecretHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyMappedSecretId
	exportNetworkFirewallNetworkFirewallPolicyApplicationGroupHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyApplicationGroupId
	exportNetworkFirewallNetworkFirewallPolicyDecryptionRuleHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyDecryptionRuleId
	exportNetworkFirewallNetworkFirewallPolicySecurityRuleHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicySecurityRuleId
	exportNetworkFirewallNetworkFirewallPolicyApplicationHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyApplicationId
	exportNetworkFirewallNetworkFirewallPolicyServiceListHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyServiceListId
	exportNetworkFirewallNetworkFirewallPolicyServiceHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyServiceId
	exportNetworkFirewallNetworkFirewallPolicyDecryptionProfileHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyDecryptionProfileId
	exportNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleHints.GetIdFn = getNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleId
	tf_export.RegisterCompartmentGraphs("network_firewall", networkFirewallResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getNetworkFirewallNetworkFirewallPolicyAddressListId(resource *tf_export.OCIResource) (string, error) {

	addressListName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find addressListName for NetworkFirewall NetworkFirewallPolicyAddressList")
	}
	networkFirewallPolicyId := resource.Parent.Id
	return GetNetworkFirewallPolicySubResourceCompositeId(addressListName, networkFirewallPolicyId, "addressLists"), nil
}

func getNetworkFirewallNetworkFirewallPolicyUrlListId(resource *tf_export.OCIResource) (string, error) {

	networkFirewallPolicyId := resource.Parent.Id
	urlListName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find urlListName for NetworkFirewall NetworkFirewallPolicyUrlList")
	}
	return GetNetworkFirewallPolicySubResourceCompositeId(urlListName, networkFirewallPolicyId, "urlLists"), nil
}

func getNetworkFirewallNetworkFirewallPolicyMappedSecretId(resource *tf_export.OCIResource) (string, error) {

	mappedSecretName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find mappedSecretName for NetworkFirewall NetworkFirewallPolicyMappedSecret")
	}
	networkFirewallPolicyId := resource.Parent.Id
	return GetNetworkFirewallPolicySubResourceCompositeId(mappedSecretName, networkFirewallPolicyId, "mappedSecrets"), nil
}

func getNetworkFirewallNetworkFirewallPolicyApplicationGroupId(resource *tf_export.OCIResource) (string, error) {

	applicationGroupName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find applicationGroupName for NetworkFirewall NetworkFirewallPolicyApplicationGroup")
	}
	networkFirewallPolicyId := resource.Parent.Id
	return GetNetworkFirewallPolicySubResourceCompositeId(applicationGroupName, networkFirewallPolicyId, "applicationGroups"), nil
}

func getNetworkFirewallNetworkFirewallPolicyDecryptionRuleId(resource *tf_export.OCIResource) (string, error) {

	decryptionRuleName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find decryptionRuleName for NetworkFirewall NetworkFirewallPolicyDecryptionRule")
	}
	networkFirewallPolicyId := resource.Parent.Id
	return GetNetworkFirewallPolicySubResourceCompositeId(decryptionRuleName, networkFirewallPolicyId, "decryptionRules"), nil
}

func getNetworkFirewallNetworkFirewallPolicySecurityRuleId(resource *tf_export.OCIResource) (string, error) {

	networkFirewallPolicyId := resource.Parent.Id
	securityRuleName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find securityRuleName for NetworkFirewall NetworkFirewallPolicySecurityRule")
	}
	return GetNetworkFirewallPolicySubResourceCompositeId(securityRuleName, networkFirewallPolicyId, "securityRules"), nil
}

func getNetworkFirewallNetworkFirewallPolicyApplicationId(resource *tf_export.OCIResource) (string, error) {

	applicationName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find applicationName for NetworkFirewall NetworkFirewallPolicyApplication")
	}
	networkFirewallPolicyId := resource.Parent.Id
	return GetNetworkFirewallPolicySubResourceCompositeId(applicationName, networkFirewallPolicyId, "applications"), nil
}

func getNetworkFirewallNetworkFirewallPolicyServiceListId(resource *tf_export.OCIResource) (string, error) {

	networkFirewallPolicyId := resource.Parent.Id
	serviceListName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find serviceListName for NetworkFirewall NetworkFirewallPolicyServiceList")
	}
	return GetNetworkFirewallPolicySubResourceCompositeId(serviceListName, networkFirewallPolicyId, "serviceLists"), nil
}

func getNetworkFirewallNetworkFirewallPolicyServiceId(resource *tf_export.OCIResource) (string, error) {

	networkFirewallPolicyId := resource.Parent.Id
	serviceName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find serviceName for NetworkFirewall NetworkFirewallPolicyService")
	}
	return GetNetworkFirewallPolicySubResourceCompositeId(serviceName, networkFirewallPolicyId, "services"), nil
}

func getNetworkFirewallNetworkFirewallPolicyDecryptionProfileId(resource *tf_export.OCIResource) (string, error) {

	decryptionProfileName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find decryptionProfileName for NetworkFirewall NetworkFirewallPolicyDecryptionProfile")
	}
	networkFirewallPolicyId := resource.Parent.Id
	return GetNetworkFirewallPolicySubResourceCompositeId(decryptionProfileName, networkFirewallPolicyId, "decryptionProfiles"), nil
}

func getNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleId(resource *tf_export.OCIResource) (string, error) {

	networkFirewallPolicyId := resource.Parent.Id
	tunnelInspectionRuleName, ok := resource.SourceAttributes["tunnel_inspection_rule_name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find tunnelInspectionRuleName for NetworkFirewall NetworkFirewallPolicyTunnelInspectionRule")
	}
	return GetNetworkFirewallPolicyTunnelInspectionRuleCompositeId(networkFirewallPolicyId, tunnelInspectionRuleName), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportNetworkFirewallNetworkFirewallPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy",
	DatasourceClass:        "oci_network_firewall_network_firewall_policies",
	DatasourceItemsAttr:    "network_firewall_policy_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_network_firewall.LifecycleStateActive),
		string(oci_network_firewall.LifecycleStateNeedsAttention),
	},
}

var exportNetworkFirewallNetworkFirewallHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall",
	DatasourceClass:        "oci_network_firewall_network_firewalls",
	DatasourceItemsAttr:    "network_firewall_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_network_firewall.LifecycleStateActive),
		string(oci_network_firewall.LifecycleStateNeedsAttention),
	},
}

var exportNetworkFirewallNetworkFirewallPolicyAddressListHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_address_list",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_address_lists",
	DatasourceItemsAttr:    "address_list_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_address_list",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyUrlListHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_url_list",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_url_lists",
	DatasourceItemsAttr:    "url_list_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_url_list",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyMappedSecretHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_mapped_secret",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_mapped_secrets",
	DatasourceItemsAttr:    "mapped_secret_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_mapped_secret",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyApplicationGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_application_group",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_application_groups",
	DatasourceItemsAttr:    "application_group_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_application_group",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyDecryptionRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_decryption_rule",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_decryption_rules",
	DatasourceItemsAttr:    "decryption_rule_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_decryption_rule",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicySecurityRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_security_rule",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_security_rules",
	DatasourceItemsAttr:    "security_rule_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_security_rule",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyApplicationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_application",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_applications",
	DatasourceItemsAttr:    "application_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_application",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyServiceListHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_service_list",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_service_lists",
	DatasourceItemsAttr:    "service_list_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_service_list",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyServiceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_service",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_services",
	DatasourceItemsAttr:    "service_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_service",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyDecryptionProfileHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_decryption_profile",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_decryption_profiles",
	DatasourceItemsAttr:    "decryption_profile_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_decryption_profile",
	RequireResourceRefresh: true,
}

var exportNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy_tunnel_inspection_rule",
	DatasourceClass:        "oci_network_firewall_network_firewall_policy_tunnel_inspection_rules",
	DatasourceItemsAttr:    "tunnel_inspection_rule_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy_tunnel_inspection_rule",
	RequireResourceRefresh: true,
}

var networkFirewallResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyHints},
		{TerraformResourceHints: exportNetworkFirewallNetworkFirewallHints},
	},
	"oci_network_firewall_network_firewall_policy": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyAddressListHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyApplicationHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyApplicationGroupHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyDecryptionProfileHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyDecryptionRuleHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyMappedSecretHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicySecurityRuleHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyServiceHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyServiceListHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyUrlListHints,
			DatasourceQueryParams: map[string]string{
				"network_firewall_policy_id": "id",
			},
		},
	},
	"oci_network_firewall_network_firewall_policy_application": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyApplicationGroupHints,
		},
	},
	"oci_network_firewall_network_firewall_policy_service": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyServiceListHints,
		},
	},
	"oci_network_firewall_network_firewall_policy_service_list": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicySecurityRuleHints,
		},
	},
	"oci_network_firewall_network_firewall_policy_application_group": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicySecurityRuleHints,
		},
	},
	"oci_network_firewall_network_firewall_policy_url_list": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicySecurityRuleHints,
		},
	},
	"oci_network_firewall_network_firewall_policy_address_list": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyDecryptionRuleHints,
		},
	},
	"oci_network_firewall_network_firewall_policy_decryption_profile": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyDecryptionRuleHints,
		},
	},
	"oci_network_firewall_network_firewall_policy_mapped_secret": {
		{
			TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyDecryptionRuleHints,
		},
	},
}
