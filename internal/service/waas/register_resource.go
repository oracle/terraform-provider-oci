// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_waas_address_list", WaasAddressListResource())
	tfresource.RegisterResource("oci_waas_certificate", WaasCertificateResource())
	tfresource.RegisterResource("oci_waas_custom_protection_rule", WaasCustomProtectionRuleResource())
	tfresource.RegisterResource("oci_waas_http_redirect", WaasHttpRedirectResource())
	tfresource.RegisterResource("oci_waas_protection_rule", WaasProtectionRuleResource())
	tfresource.RegisterResource("oci_waas_purge_cache", WaasPurgeCacheResource())
	tfresource.RegisterResource("oci_waas_waas_policy", WaasWaasPolicyResource())
}
