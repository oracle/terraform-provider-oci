// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_fleet_software_update_fsu_collection", FleetSoftwareUpdateFsuCollectionDataSource())
	tfresource.RegisterDatasource("oci_fleet_software_update_fsu_collections", FleetSoftwareUpdateFsuCollectionsDataSource())
	tfresource.RegisterDatasource("oci_fleet_software_update_fsu_cycle", FleetSoftwareUpdateFsuCycleDataSource())
	tfresource.RegisterDatasource("oci_fleet_software_update_fsu_cycles", FleetSoftwareUpdateFsuCyclesDataSource())
}
