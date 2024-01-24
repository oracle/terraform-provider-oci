// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_waas_address_list", WaasAddressListDataSource())
	tfresource.RegisterDatasource("oci_waas_address_lists", WaasAddressListsDataSource())
	tfresource.RegisterDatasource("oci_waas_certificate", WaasCertificateDataSource())
	tfresource.RegisterDatasource("oci_waas_certificates", WaasCertificatesDataSource())
	tfresource.RegisterDatasource("oci_waas_custom_protection_rule", WaasCustomProtectionRuleDataSource())
	tfresource.RegisterDatasource("oci_waas_custom_protection_rules", WaasCustomProtectionRulesDataSource())
	tfresource.RegisterDatasource("oci_waas_edge_subnets", WaasEdgeSubnetsDataSource())
	tfresource.RegisterDatasource("oci_waas_http_redirect", WaasHttpRedirectDataSource())
	tfresource.RegisterDatasource("oci_waas_http_redirects", WaasHttpRedirectsDataSource())
	tfresource.RegisterDatasource("oci_waas_protection_rules", WaasProtectionRulesDataSource())
	tfresource.RegisterDatasource("oci_waas_waas_policies", WaasWaasPoliciesDataSource())
	tfresource.RegisterDatasource("oci_waas_waas_policy", WaasWaasPolicyDataSource())
	tfresource.RegisterDatasource("oci_waas_protection_rule", WaasProtectionRuleDataSource())
}
