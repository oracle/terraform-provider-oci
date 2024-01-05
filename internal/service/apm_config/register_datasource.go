// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_config

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_apm_config_config", ApmConfigConfigDataSource())
	tfresource.RegisterDatasource("oci_apm_config_configs", ApmConfigConfigsDataSource())
}
