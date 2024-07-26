// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_fleet_software_update_fsu_collection", FleetSoftwareUpdateFsuCollectionResource())
	tfresource.RegisterResource("oci_fleet_software_update_fsu_cycle", FleetSoftwareUpdateFsuCycleResource())
}
