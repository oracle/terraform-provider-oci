// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v25/common"
	oci_database "github.com/oracle/oci-go-sdk/v25/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousExadataInfrastructureRequiredOnlyResource = AutonomousExadataInfrastructureResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Required, Create, autonomousExadataInfrastructureRepresentation)

	AutonomousExadataInfrastructureResourceConfig = AutonomousExadataInfrastructureResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Optional, Update, autonomousExadataInfrastructureRepresentation)

	autonomousExadataInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`},
	}

	autonomousExadataInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domain.ad.name}`},
		"display_name":        Representation{repType: Optional, create: `tst3dbsys`, update: `displayName2`},
		"state":               Representation{repType: Optional, create: `AVAILABLE`},
		"filter":              RepresentationGroup{Required, autonomousExadataInfrastructureDataSourceFilterRepresentation}}
	autonomousExadataInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`}},
	}

	autonomousExadataInfrastructureRepresentation = map[string]interface{}{
		"availability_domain":        Representation{repType: Required, create: `${data.oci_identity_availability_domain.ad.name}`},
		"compartment_id":             Representation{repType: Required, create: `${var.compartment_id}`},
		"shape":                      Representation{repType: Required, create: `Exadata.Quarter2.92`},
		"subnet_id":                  Representation{repType: Required, create: `${oci_core_subnet.exadata_subnet.id}`},
		"defined_tags":               Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":               Representation{repType: Optional, create: `tst3dbsys`, update: `displayName2`},
		"domain":                     Representation{repType: Optional, create: `subnetexadata.tfvcn.oraclevcn.com`},
		"freeform_tags":              Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"license_model":              Representation{repType: Optional, create: `LICENSE_INCLUDED`},
		"maintenance_window_details": RepresentationGroup{Optional, autonomousExadataInfrastructureMaintenanceWindowDetailsRepresentation},
		"nsg_ids":                    Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, update: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsRepresentation = map[string]interface{}{
		"preference":     Representation{repType: Required, create: `NO_PREFERENCE`, update: `CUSTOM_PREFERENCE`},
		"days_of_week":   RepresentationGroup{Optional, autonomousExadataInfrastructureMaintenanceWindowDetailsDaysOfWeekRepresentation},
		"hours_of_day":   Representation{repType: Optional, create: []string{`4`}, update: []string{`8`}},
		"months":         []RepresentationGroup{{Optional, autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation}, {Optional, autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation2}, {Optional, autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation3}, {Optional, autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation4}},
		"weeks_of_month": Representation{repType: Optional, create: []string{`1`}, update: []string{`2`}},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsDaysOfWeekRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `MONDAY`, update: `TUESDAY`},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `JANUARY`, update: `FEBRUARY`},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation2 = map[string]interface{}{
		"name": Representation{repType: Required, create: `APRIL`, update: `MAY`},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation3 = map[string]interface{}{
		"name": Representation{repType: Required, create: `JULY`, update: `AUGUST`},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation4 = map[string]interface{}{
		"name": Representation{repType: Required, create: `OCTOBER`, update: `NOVEMBER`},
	}

	AutonomousExadataInfrastructureResourceDependencies = ExadataBaseDependencies +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, getUpdatedRepresentationCopy("vcn_id", Representation{repType: Required, create: `${oci_core_virtual_network.t.id}`}, networkSecurityGroupRepresentation)) +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", Required, Create, getUpdatedRepresentationCopy("vcn_id", Representation{repType: Required, create: `${oci_core_virtual_network.t.id}`}, networkSecurityGroupRepresentation))
)

func TestDatabaseAutonomousExadataInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure"
	datasourceName := "data.oci_database_autonomous_exadata_infrastructures.test_autonomous_exadata_infrastructures"
	singularDatasourceName := "data.oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousExadataInfrastructureDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Required, Create, autonomousExadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.Quarter2.92"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Optional, Create, autonomousExadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tst3dbsys"),
					resource.TestCheckResourceAttr(resourceName, "domain", "subnetexadata.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "hostname"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.Quarter2.92"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Optional, Create,
						representationCopyWithNewProperties(autonomousExadataInfrastructureRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tst3dbsys"),
					resource.TestCheckResourceAttr(resourceName, "domain", "subnetexadata.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "hostname"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.preference", "NO_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.Quarter2.92"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
				Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Optional, Update, autonomousExadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "domain", "subnetexadata.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "hostname"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.Quarter2.92"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructures", "test_autonomous_exadata_infrastructures", Optional, Update, autonomousExadataInfrastructureDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Optional, Update, autonomousExadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.domain", "subnetexadata.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.hostname"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.shape", "Exadata.Quarter2.92"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Required, Create, autonomousExadataInfrastructureSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousExadataInfrastructureResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_exadata_infrastructure_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "domain", "subnetexadata.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "hostname"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "Exadata.Quarter2.92"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceConfig,
			},
			// verify resource import
			{
				Config:            config + generateResourceImportConfig("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure"),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"maintenance_window_details",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckDatabaseAutonomousExadataInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_exadata_infrastructure" {
			noResourceFound = false
			request := oci_database.GetAutonomousExadataInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.AutonomousExadataInfrastructureId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetAutonomousExadataInfrastructure(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousExadataInfrastructureLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("DatabaseAutonomousExadataInfrastructure") {
		resource.AddTestSweepers("DatabaseAutonomousExadataInfrastructure", &resource.Sweeper{
			Name:         "DatabaseAutonomousExadataInfrastructure",
			Dependencies: DependencyGraph["autonomousExadataInfrastructure"],
			F:            sweepDatabaseAutonomousExadataInfrastructureResource,
		})
	}
}

func sweepDatabaseAutonomousExadataInfrastructureResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	autonomousExadataInfrastructureIds, err := getAutonomousExadataInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousExadataInfrastructureId := range autonomousExadataInfrastructureIds {
		if ok := SweeperDefaultResourceId[autonomousExadataInfrastructureId]; !ok {
			terminateAutonomousExadataInfrastructureRequest := oci_database.TerminateAutonomousExadataInfrastructureRequest{}

			terminateAutonomousExadataInfrastructureRequest.AutonomousExadataInfrastructureId = &autonomousExadataInfrastructureId

			terminateAutonomousExadataInfrastructureRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.TerminateAutonomousExadataInfrastructure(context.Background(), terminateAutonomousExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousExadataInfrastructureId, error)
				continue
			}
			waitTillCondition(testAccProvider, &autonomousExadataInfrastructureId, autonomousExadataInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				autonomousExadataInfrastructureSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getAutonomousExadataInfrastructureIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "AutonomousExadataInfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listAutonomousExadataInfrastructuresRequest := oci_database.ListAutonomousExadataInfrastructuresRequest{}
	listAutonomousExadataInfrastructuresRequest.CompartmentId = &compartmentId
	listAutonomousExadataInfrastructuresRequest.LifecycleState = oci_database.AutonomousExadataInfrastructureSummaryLifecycleStateAvailable
	listAutonomousExadataInfrastructuresResponse, err := databaseClient.ListAutonomousExadataInfrastructures(context.Background(), listAutonomousExadataInfrastructuresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousExadataInfrastructure list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousExadataInfrastructure := range listAutonomousExadataInfrastructuresResponse.Items {
		id := *autonomousExadataInfrastructure.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousExadataInfrastructureId", id)
	}
	return resourceIds, nil
}

func autonomousExadataInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousExadataInfrastructureResponse, ok := response.Response.(oci_database.GetAutonomousExadataInfrastructureResponse); ok {
		return autonomousExadataInfrastructureResponse.LifecycleState != oci_database.AutonomousExadataInfrastructureLifecycleStateTerminated
	}
	return false
}

func autonomousExadataInfrastructureSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetAutonomousExadataInfrastructure(context.Background(), oci_database.GetAutonomousExadataInfrastructureRequest{
		AutonomousExadataInfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
