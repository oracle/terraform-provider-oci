// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_osmanagement_managed_instance", OsmanagementManagedInstanceResource())
	tfresource.RegisterResource("oci_osmanagement_managed_instance_group", OsmanagementManagedInstanceGroupResource())
	tfresource.RegisterResource("oci_osmanagement_managed_instance_management", OsmanagementManagedInstanceManagementResource())
	tfresource.RegisterResource("oci_osmanagement_software_source", OsmanagementSoftwareSourceResource())
}
