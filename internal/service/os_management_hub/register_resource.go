// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_os_management_hub_lifecycle_environment", OsManagementHubLifecycleEnvironmentResource())
	tfresource.RegisterResource("oci_os_management_hub_managed_instance_group", OsManagementHubManagedInstanceGroupResource())
	tfresource.RegisterResource("oci_os_management_hub_management_station", OsManagementHubManagementStationResource())
	tfresource.RegisterResource("oci_os_management_hub_profile", OsManagementHubProfileResource())
	tfresource.RegisterResource("oci_os_management_hub_software_source", OsManagementHubSoftwareSourceResource())
}
