// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LustreFileStorageAvailableMaintenanceScheduleStartTimeDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	LustreFileStorageAvailableMaintenanceScheduleStartTimeByFsIdDataSourceRepresentation = map[string]interface{}{
		"id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id}`},
		"day_of_week": acctest.Representation{RepType: acctest.Required, Create: `MONDAY`},
	}

	LustreFileStorageAvailableMaintenanceScheduleStartTimeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: lustre_file_storage/default
func TestLustreFileStorageAvailableMaintenanceScheduleStartTimeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLustreFileStorageAvailableMaintenanceScheduleStartTimeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_lustre_file_storage_available_maintenance_schedule_start_times.test_available_maintenance_schedule_start_times"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Step 1: compartment_id + availability_domain
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_lustre_file_storage_available_maintenance_schedule_start_times", "test_available_maintenance_schedule_start_times", acctest.Required, acctest.Create, LustreFileStorageAvailableMaintenanceScheduleStartTimeDataSourceRepresentation) +
				compartmentIdVariableStr + LustreFileStorageAvailableMaintenanceScheduleStartTimeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "available_maintenance_schedule_start_time_collection.#"),
			),
		},

		// Step 2: lustre_file_system_id (+ day_of_week)
		{
			Config: config +
				//LustreFileStorageLustreFileSystemRequiredOnlyResource +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_lustre_file_storage_available_maintenance_schedule_start_times",
					"test_available_maintenance_schedule_start_times",
					acctest.Required,
					acctest.Create,
					LustreFileStorageAvailableMaintenanceScheduleStartTimeByFsIdDataSourceRepresentation,
				) +
				compartmentIdVariableStr + LustreFileStorageAvailableMaintenanceScheduleStartTimeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "available_maintenance_schedule_start_time_collection.#"),
				resource.TestCheckResourceAttr(
					datasourceName,
					"available_maintenance_schedule_start_time_collection.0.items.0.day_of_week",
					"MONDAY",
				),
			),
		},
	})
}
