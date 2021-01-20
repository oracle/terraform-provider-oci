// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v33/common"
	oci_database "github.com/oracle/oci-go-sdk/v33/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CloudExadataInfrastructureRequiredOnlyResource = CloudExadataInfrastructureResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Required, Create, cloudExadataInfrastructureRepresentation)

	CloudExadataInfrastructureResourceConfig = CloudExadataInfrastructureResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Optional, Update, cloudExadataInfrastructureRepresentation)

	cloudExadataInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
	}

	cloudExadataInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `tstExaInfra`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, cloudExadataInfrastructureDataSourceFilterRepresentation}}
	cloudExadataInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`}},
	}

	cloudExadataInfrastructureRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Required, create: `tstExaInfra`, update: `displayName2`},
		"shape":               Representation{repType: Required, create: `Exadata.X8M`},
		"compute_count":       Representation{repType: Required, create: `2`}, // required for shape Exadata.X8M
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":  RepresentationGroup{Optional, cloudExadataInfrastructureMaintenanceWindowRepresentation},
		"storage_count":       Representation{repType: Required, create: `3`}, // required for shape Exadata.X8M
	}
	cloudExadataInfrastructureMaintenanceWindowRepresentation = map[string]interface{}{
		"preference":         Representation{repType: Required, create: `CUSTOM_PREFERENCE`},
		"days_of_week":       RepresentationGroup{Optional, cloudExadataInfrastructureMaintenanceWindowDaysOfWeekRepresentation},
		"hours_of_day":       Representation{repType: Optional, create: []string{`4`}, update: []string{`8`}},
		"lead_time_in_weeks": Representation{repType: Optional, create: `10`, update: `11`},
		"months":             RepresentationGroup{Optional, cloudExadataInfrastructureMaintenanceWindowMonthsRepresentation},
		"weeks_of_month":     Representation{repType: Optional, create: []string{`1`}, update: []string{`2`}},
	}
	cloudExadataInfrastructureMaintenanceWindowDaysOfWeekRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `MONDAY`, update: `TUESDAY`},
	}
	cloudExadataInfrastructureMaintenanceWindowMonthsRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `APRIL`, update: `MAY`},
	}

	CloudExadataInfrastructureResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

func TestDatabaseCloudExadataInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure"
	datasourceName := "data.oci_database_cloud_exadata_infrastructures.test_cloud_exadata_infrastructures"
	singularDatasourceName := "data.oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseCloudExadataInfrastructureDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CloudExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Required, Create, cloudExadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CloudExadataInfrastructureResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CloudExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Optional, Create, cloudExadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "MONDAY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "10"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "APRIL"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Optional, Create,
						representationCopyWithNewProperties(cloudExadataInfrastructureRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "MONDAY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "10"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "APRIL"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CloudExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Optional, Update, cloudExadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "11"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "MAY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_cloud_exadata_infrastructures", "test_cloud_exadata_infrastructures", Optional, Update, cloudExadataInfrastructureDataSourceRepresentation) +
					compartmentIdVariableStr + CloudExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Optional, Update, cloudExadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.available_storage_size_in_gbs"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.compute_count", "2"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.id"),
					//resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.last_maintenance_run_id"), // null for fake resource
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.lead_time_in_weeks", "11"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.months.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.months.0.name", "MAY"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.weeks_of_month.#", "1"),
					//resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.next_maintenance_run_id"), // null for fake resource
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.shape", "Exadata.X8M"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.storage_count", "3"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.total_storage_size_in_gbs"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Required, Create, cloudExadataInfrastructureSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CloudExadataInfrastructureResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_exadata_infrastructure_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "available_storage_size_in_gbs"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "compute_count", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "last_maintenance_run_id"), // null for fake resource
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.lead_time_in_weeks", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "MAY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "1"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "next_maintenance_run_id"), // null for fake resource
					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "Exadata.X8M"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "storage_count", "3"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "total_storage_size_in_gbs"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CloudExadataInfrastructureResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDatabaseCloudExadataInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_cloud_exadata_infrastructure" {
			noResourceFound = false
			request := oci_database.GetCloudExadataInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.CloudExadataInfrastructureId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetCloudExadataInfrastructure(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.CloudExadataInfrastructureLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseCloudExadataInfrastructure") {
		resource.AddTestSweepers("DatabaseCloudExadataInfrastructure", &resource.Sweeper{
			Name:         "DatabaseCloudExadataInfrastructure",
			Dependencies: DependencyGraph["cloudExadataInfrastructure"],
			F:            sweepDatabaseCloudExadataInfrastructureResource,
		})
	}
}

func sweepDatabaseCloudExadataInfrastructureResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	cloudExadataInfrastructureIds, err := getCloudExadataInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudExadataInfrastructureId := range cloudExadataInfrastructureIds {
		if ok := SweeperDefaultResourceId[cloudExadataInfrastructureId]; !ok {
			deleteCloudExadataInfrastructureRequest := oci_database.DeleteCloudExadataInfrastructureRequest{}

			deleteCloudExadataInfrastructureRequest.CloudExadataInfrastructureId = &cloudExadataInfrastructureId

			deleteCloudExadataInfrastructureRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteCloudExadataInfrastructure(context.Background(), deleteCloudExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudExadataInfrastructureId, error)
				continue
			}
			waitTillCondition(testAccProvider, &cloudExadataInfrastructureId, cloudExadataInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				cloudExadataInfrastructureSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getCloudExadataInfrastructureIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CloudExadataInfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listCloudExadataInfrastructuresRequest := oci_database.ListCloudExadataInfrastructuresRequest{}
	listCloudExadataInfrastructuresRequest.CompartmentId = &compartmentId
	listCloudExadataInfrastructuresRequest.LifecycleState = oci_database.CloudExadataInfrastructureSummaryLifecycleStateAvailable
	listCloudExadataInfrastructuresResponse, err := databaseClient.ListCloudExadataInfrastructures(context.Background(), listCloudExadataInfrastructuresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudExadataInfrastructure list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudExadataInfrastructure := range listCloudExadataInfrastructuresResponse.Items {
		id := *cloudExadataInfrastructure.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "CloudExadataInfrastructureId", id)
	}
	return resourceIds, nil
}

func cloudExadataInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudExadataInfrastructureResponse, ok := response.Response.(oci_database.GetCloudExadataInfrastructureResponse); ok {
		return cloudExadataInfrastructureResponse.LifecycleState != oci_database.CloudExadataInfrastructureLifecycleStateTerminated
	}
	return false
}

func cloudExadataInfrastructureSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetCloudExadataInfrastructure(context.Background(), oci_database.GetCloudExadataInfrastructureRequest{
		CloudExadataInfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
