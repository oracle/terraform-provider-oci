// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package appmgmt_control

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_appmgmt_control_monitored_instance", AppmgmtControlMonitoredInstanceDataSource())
	tfresource.RegisterDatasource("oci_appmgmt_control_monitored_instances", AppmgmtControlMonitoredInstancesDataSource())
}
