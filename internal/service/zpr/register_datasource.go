// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package zpr

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_zpr_configuration", ZprConfigurationDataSource())
	tfresource.RegisterDatasource("oci_zpr_zpr_policies", ZprZprPoliciesDataSource())
	tfresource.RegisterDatasource("oci_zpr_zpr_policy", ZprZprPolicyDataSource())
}
