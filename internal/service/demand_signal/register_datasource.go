// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package demand_signal

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_demand_signal_occ_demand_signal", DemandSignalOccDemandSignalDataSource())
	tfresource.RegisterDatasource("oci_demand_signal_occ_demand_signals", DemandSignalOccDemandSignalsDataSource())
}
