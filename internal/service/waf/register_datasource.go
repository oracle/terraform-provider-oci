// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_waf_network_address_list", WafNetworkAddressListDataSource())
	tfresource.RegisterDatasource("oci_waf_network_address_lists", WafNetworkAddressListsDataSource())
	tfresource.RegisterDatasource("oci_waf_protection_capabilities", WafProtectionCapabilitiesDataSource())
	tfresource.RegisterDatasource("oci_waf_protection_capability_group_tags", WafProtectionCapabilityGroupTagsDataSource())
	tfresource.RegisterDatasource("oci_waf_web_app_firewall", WafWebAppFirewallDataSource())
	tfresource.RegisterDatasource("oci_waf_web_app_firewall_policies", WafWebAppFirewallPoliciesDataSource())
	tfresource.RegisterDatasource("oci_waf_web_app_firewall_policy", WafWebAppFirewallPolicyDataSource())
	tfresource.RegisterDatasource("oci_waf_web_app_firewalls", WafWebAppFirewallsDataSource())
}
