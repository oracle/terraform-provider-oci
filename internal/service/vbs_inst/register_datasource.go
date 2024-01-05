// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vbs_inst

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_vbs_inst_vbs_instance", VbsInstVbsInstanceDataSource())
	tfresource.RegisterDatasource("oci_vbs_inst_vbs_instances", VbsInstVbsInstancesDataSource())
}
