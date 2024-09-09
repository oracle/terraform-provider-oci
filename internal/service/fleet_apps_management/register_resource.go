// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_fleet_apps_management_fleet", FleetAppsManagementFleetResource())
	tfresource.RegisterResource("oci_fleet_apps_management_fleet_credential", FleetAppsManagementFleetCredentialResource())
	tfresource.RegisterResource("oci_fleet_apps_management_fleet_property", FleetAppsManagementFleetPropertyResource())
	tfresource.RegisterResource("oci_fleet_apps_management_fleet_resource", FleetAppsManagementFleetResourceResource())
	tfresource.RegisterResource("oci_fleet_apps_management_maintenance_window", FleetAppsManagementMaintenanceWindowResource())
	tfresource.RegisterResource("oci_fleet_apps_management_onboarding", FleetAppsManagementOnboardingResource())
	tfresource.RegisterResource("oci_fleet_apps_management_property", FleetAppsManagementPropertyResource())
	tfresource.RegisterResource("oci_fleet_apps_management_scheduler_definition", FleetAppsManagementSchedulerDefinitionResource())
}
