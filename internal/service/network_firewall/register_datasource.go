// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall", NetworkFirewallNetworkFirewallDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policies", NetworkFirewallNetworkFirewallPoliciesDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy", NetworkFirewallNetworkFirewallPolicyDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_address_list", NetworkFirewallNetworkFirewallPolicyAddressListDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_address_lists", NetworkFirewallNetworkFirewallPolicyAddressListsDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_application", NetworkFirewallNetworkFirewallPolicyApplicationDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_application_group", NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_application_groups", NetworkFirewallNetworkFirewallPolicyApplicationGroupsDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_applications", NetworkFirewallNetworkFirewallPolicyApplicationsDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_decryption_profile", NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_decryption_profiles", NetworkFirewallNetworkFirewallPolicyDecryptionProfilesDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_decryption_rule", NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_decryption_rules", NetworkFirewallNetworkFirewallPolicyDecryptionRulesDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_mapped_secret", NetworkFirewallNetworkFirewallPolicyMappedSecretDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_mapped_secrets", NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_nat_rule", NetworkFirewallNetworkFirewallPolicyNatRuleDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_nat_rules", NetworkFirewallNetworkFirewallPolicyNatRulesDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_security_rule", NetworkFirewallNetworkFirewallPolicySecurityRuleDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_security_rules", NetworkFirewallNetworkFirewallPolicySecurityRulesDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_service", NetworkFirewallNetworkFirewallPolicyServiceDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_service_list", NetworkFirewallNetworkFirewallPolicyServiceListDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_service_lists", NetworkFirewallNetworkFirewallPolicyServiceListsDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_services", NetworkFirewallNetworkFirewallPolicyServicesDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_tunnel_inspection_rule", NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_tunnel_inspection_rules", NetworkFirewallNetworkFirewallPolicyTunnelInspectionRulesDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_url_list", NetworkFirewallNetworkFirewallPolicyUrlListDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy_url_lists", NetworkFirewallNetworkFirewallPolicyUrlListsDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewalls", NetworkFirewallNetworkFirewallsDataSource())
}
