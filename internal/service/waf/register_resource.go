// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_waf_network_address_list", WafNetworkAddressListResource())
	tfresource.RegisterResource("oci_waf_web_app_firewall", WafWebAppFirewallResource())
	tfresource.RegisterResource("oci_waf_web_app_firewall_policy", WafWebAppFirewallPolicyResource())
}
