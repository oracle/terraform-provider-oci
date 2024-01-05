// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_apm_apm_domain", ApmApmDomainDataSource())
	tfresource.RegisterDatasource("oci_apm_apm_domains", ApmApmDomainsDataSource())
	tfresource.RegisterDatasource("oci_apm_data_keys", ApmDataKeysDataSource())
}
