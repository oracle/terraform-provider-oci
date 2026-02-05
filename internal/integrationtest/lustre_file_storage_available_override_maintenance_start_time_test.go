// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LustreFileStorageAvailableOverrideMaintenanceStartTimeDataSourceRepresentation = map[string]interface{}{
		"id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id}`},
		"date": acctest.Representation{RepType: acctest.Optional, Create: getFutureDate(16)}, // 2 weeks + 2 days = 16 days
	}

	LustreFileStorageAvailableOverrideMaintenanceStartTimeResourceConfig = ""
)

// issue-routing-tag: lustre_file_storage/default
func TestLustreFileStorageAvailableOverrideMaintenanceStartTimeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLustreFileStorageAvailableOverrideMaintenanceStartTimeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_lustre_file_storage_available_override_maintenance_start_times.test_available_override_maintenance_start_times"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Step 1 – Create the Lustre File System only
		{
			Config: config +
				LustreFileStorageLustreFileSystemRequiredOnlyResource +
				compartmentIdVariableStr,
		},

		// Step 2 – wait 10 minutes to allow maintenance metadata to become available before triggering override
		{
			PreConfig: func() {
				t.Log("Waiting 10 minutes for maintenance metadata propagation...")
				time.Sleep(10 * time.Minute)
			},

			Config: config +
				LustreFileStorageLustreFileSystemRequiredOnlyResource +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_lustre_file_storage_available_override_maintenance_start_times",
					"test_available_override_maintenance_start_times",
					acctest.Required,
					acctest.Create,
					LustreFileStorageAvailableOverrideMaintenanceStartTimeDataSourceRepresentation,
				) +
				compartmentIdVariableStr +
				LustreFileStorageAvailableOverrideMaintenanceStartTimeResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "available_override_maintenance_start_time_collection.#"),
			),
		},
	})
}
