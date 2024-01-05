// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MeteringComputationScheduleRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_schedule", "test_schedule", acctest.Required, acctest.Create, MeteringComputationScheduleRepresentation)

	MeteringComputationScheduleResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_schedule", "test_schedule", acctest.Optional, acctest.Update, MeteringComputationScheduleRepresentation)

	MeteringComputationMeteringComputationScheduleSingularDataSourceRepresentation = map[string]interface{}{
		"schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_metering_computation_schedule.test_schedule.id}`},
	}

	MeteringComputationMeteringComputationScheduleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationScheduleDataSourceFilterRepresentation}}
	MeteringComputationScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_metering_computation_schedule.test_schedule.id}`}},
	}

	MeteringComputationScheduleRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"name":                 acctest.Representation{RepType: acctest.Required, Create: `name`},
		"query_properties":     acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationScheduleQueryPropertiesRepresentation},
		"result_location":      acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationScheduleResultLocationRepresentation},
		"schedule_recurrences": acctest.Representation{RepType: acctest.Required, Create: `DAILY`},
		"time_scheduled":       acctest.Representation{RepType: acctest.Required, Create: `2022-10-19T00:00:00Z`},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"output_file_format":   acctest.Representation{RepType: acctest.Optional, Create: `CSV`, Update: `PDF`},
	}
	MeteringComputationScheduleQueryPropertiesRepresentation = map[string]interface{}{
		"date_range":  acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationScheduleQueryPropertiesDateRangeRepresentation},
		"granularity": acctest.Representation{RepType: acctest.Required, Create: `DAILY`},
	}
	MeteringComputationScheduleResultLocationRepresentation = map[string]interface{}{
		"bucket":        acctest.Representation{RepType: acctest.Required, Create: `usage-schedule-test-bucket`},
		"location_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		"namespace":     acctest.Representation{RepType: acctest.Required, Create: `idy3u7psgoxm`},
		"region":        acctest.Representation{RepType: acctest.Required, Create: `us-ashburn-1`},
	}
	MeteringComputationScheduleQueryPropertiesDateRangeRepresentation = map[string]interface{}{
		"date_range_type":         acctest.Representation{RepType: acctest.Required, Create: `DYNAMIC`},
		"dynamic_date_range_type": acctest.Representation{RepType: acctest.Required, Create: `LAST_7_DAYS`},
	}
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	resourceName := "oci_metering_computation_schedule.test_schedule"
	datasourceName := "data.oci_metering_computation_schedules.test_schedules"
	singularDatasourceName := "data.oci_metering_computation_schedule.test_schedule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+tenancyIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_schedule", "test_schedule", acctest.Optional, acctest.Create, MeteringComputationScheduleRepresentation), "usageapi", "schedule", t)

	acctest.ResourceTest(t, testAccCheckMeteringComputationScheduleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + tenancyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_schedule", "test_schedule", acctest.Required, acctest.Create, MeteringComputationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.0.date_range_type", "DYNAMIC"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.0.dynamic_date_range_type", "LAST_7_DAYS"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.granularity", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "result_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.bucket", "usage-schedule-test-bucket"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.namespace", "idy3u7psgoxm"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.region", "us-ashburn-1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_recurrences", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "time_scheduled", "2022-10-19T00:00:00Z"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + tenancyIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + tenancyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_schedule", "test_schedule", acctest.Optional, acctest.Create, MeteringComputationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.0.date_range_type", "DYNAMIC"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.0.dynamic_date_range_type", "LAST_7_DAYS"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.granularity", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "result_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.bucket", "usage-schedule-test-bucket"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.namespace", "idy3u7psgoxm"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.region", "us-ashburn-1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_recurrences", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "output_file_format", "CSV"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + tenancyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_schedule", "test_schedule", acctest.Optional, acctest.Update, MeteringComputationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.0.date_range_type", "DYNAMIC"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.date_range.0.dynamic_date_range_type", "LAST_7_DAYS"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.granularity", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "result_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.bucket", "usage-schedule-test-bucket"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.namespace", "idy3u7psgoxm"),
				resource.TestCheckResourceAttr(resourceName, "result_location.0.region", "us-ashburn-1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_recurrences", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "output_file_format", "PDF"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_schedules", "test_schedules", acctest.Optional, acctest.Update, MeteringComputationMeteringComputationScheduleDataSourceRepresentation) +
				tenancyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_schedule", "test_schedule", acctest.Optional, acctest.Update, MeteringComputationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),

				resource.TestCheckResourceAttr(datasourceName, "schedule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schedule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_schedule", "test_schedule", acctest.Required, acctest.Create, MeteringComputationMeteringComputationScheduleSingularDataSourceRepresentation) +
				tenancyIdVariableStr + MeteringComputationScheduleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.date_range.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.date_range.0.date_range_type", "DYNAMIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.date_range.0.dynamic_date_range_type", "LAST_7_DAYS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.granularity", "DAILY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "result_location.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "result_location.0.bucket", "usage-schedule-test-bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "result_location.0.location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "result_location.0.namespace", "idy3u7psgoxm"),
				resource.TestCheckResourceAttr(singularDatasourceName, "result_location.0.region", "us-ashburn-1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_recurrences", "DAILY"),
			),
		},
		// verify resource import
		{
			Config:                  config + MeteringComputationScheduleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMeteringComputationScheduleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).UsageapiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_metering_computation_schedule" {
			noResourceFound = false
			request := oci_metering_computation.GetScheduleRequest{}

			tmp := rs.Primary.ID
			request.ScheduleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "metering_computation")

			_, err := client.GetSchedule(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MeteringComputationSchedule") {
		resource.AddTestSweepers("MeteringComputationSchedule", &resource.Sweeper{
			Name:         "MeteringComputationSchedule",
			Dependencies: acctest.DependencyGraph["schedule"],
			F:            sweepMeteringComputationScheduleResource,
		})
	}
}

func sweepMeteringComputationScheduleResource(compartment string) error {
	usageapiClient := acctest.GetTestClients(&schema.ResourceData{}).UsageapiClient()
	scheduleIds, err := getMeteringComputationScheduleIds(compartment)
	if err != nil {
		return err
	}
	for _, scheduleId := range scheduleIds {
		if ok := acctest.SweeperDefaultResourceId[scheduleId]; !ok {
			deleteScheduleRequest := oci_metering_computation.DeleteScheduleRequest{}

			deleteScheduleRequest.ScheduleId = &scheduleId

			deleteScheduleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "metering_computation")
			_, error := usageapiClient.DeleteSchedule(context.Background(), deleteScheduleRequest)
			if error != nil {
				fmt.Printf("Error deleting Schedule %s %s, It is possible that the resource is already deleted. Please verify manually \n", scheduleId, error)
				continue
			}
		}
	}
	return nil
}

func getMeteringComputationScheduleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ScheduleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	usageapiClient := acctest.GetTestClients(&schema.ResourceData{}).UsageapiClient()

	listSchedulesRequest := oci_metering_computation.ListSchedulesRequest{}
	listSchedulesRequest.CompartmentId = &compartmentId
	listSchedulesResponse, err := usageapiClient.ListSchedules(context.Background(), listSchedulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Schedule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, schedule := range listSchedulesResponse.Items {
		id := *schedule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ScheduleId", id)
	}
	return resourceIds, nil
}
