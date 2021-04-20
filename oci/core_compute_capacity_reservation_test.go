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
	"github.com/oracle/oci-go-sdk/v40/common"
	oci_core "github.com/oracle/oci-go-sdk/v40/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ComputeCapacityReservationRequiredOnlyResource = ComputeCapacityReservationResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_compute_capacity_reservTestCoreComputeCapacityReservationResource_basication", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation)

	ComputeCapacityReservationResourceConfig = ComputeCapacityReservationResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Update, computeCapacityReservationRepresentation)

	computeCapacityReservationSingularDataSourceRepresentation = map[string]interface{}{
		"capacity_reservation_id": Representation{repType: Required, create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
	}

	computeCapacityReservationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{repType: Optional, create: `displayNameDataSourceCreate`, update: `displayNameDataSourceUpdate`},
		"state":               Representation{repType: Optional, create: `ACTIVE`},
		"filter":              RepresentationGroup{Required, computeCapacityReservationDataSourceFilterRepresentation}}
	computeCapacityReservationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`}},
	}

	computeCapacityReservationRepresentation = map[string]interface{}{
		"availability_domain":          Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":               Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":                 Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                 Representation{repType: Optional, create: `displayNameResourceCreate`, update: `displayNameResourceUpdate`},
		"freeform_tags":                Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"instance_reservation_configs": RepresentationGroup{Required, computeCapacityReservationInstanceReservationConfigsRepresentation},
		"is_default_reservation":       Representation{repType: Optional, create: `false`, update: `true`},
	}
	computeCapacityReservationInstanceReservationConfigsRepresentation = map[string]interface{}{
		"instance_shape":        Representation{repType: Required, create: `VM.Standard2.1`},
		"fault_domain":          Representation{repType: Optional, create: `FAULT-DOMAIN-1`},
		"reserved_count":        Representation{repType: Required, create: `1`, update: `2`},
		"instance_shape_config": RepresentationGroup{Optional, computeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation},
	}
	computeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": Representation{repType: Optional, create: `15`},
		"ocpus":         Representation{repType: Optional, create: `1`},
	}

	ComputeCapacityReservationResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

func TestCoreComputeCapacityReservationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_compute_capacity_reservation.test_compute_capacity_reservation"
	datasourceName := "data.oci_core_compute_capacity_reservations.test_compute_capacity_reservations"
	singularDatasourceName := "data.oci_core_compute_capacity_reservation.test_compute_capacity_reservation"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreComputeCapacityReservationDestroy,
		Steps: []resource.TestStep{
			// Step 0: verify create
			{
				Config: config + compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// Step 1: delete before next create
			{
				Config: config + compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies,
			},
			// Step 2: verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Create, computeCapacityReservationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameResourceCreate"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.fault_domain", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.0.memory_in_gbs", "15"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.reserved_count", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_reservation_configs.0.used_count"),
					resource.TestCheckResourceAttr(resourceName, "is_default_reservation", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

			// Step 3: verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ComputeCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Create,
						representationCopyWithNewProperties(computeCapacityReservationRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameResourceCreate"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.fault_domain", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.0.memory_in_gbs", "15"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.reserved_count", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_reservation_configs.0.used_count"),
					resource.TestCheckResourceAttr(resourceName, "is_default_reservation", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// Step 4: verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Update, computeCapacityReservationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameResourceUpdate"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.fault_domain", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.0.memory_in_gbs", "15"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.instance_shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.0.reserved_count", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_reservation_configs.0.used_count"),
					resource.TestCheckResourceAttr(resourceName, "is_default_reservation", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Step 5: verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservations", "test_compute_capacity_reservations", Optional, Update, computeCapacityReservationDataSourceRepresentation) +
					compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Update, computeCapacityReservationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayNameDataSourceUpdate"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "compute_capacity_reservations.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservations.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compute_capacity_reservations.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "compute_capacity_reservations.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "compute_capacity_reservations.0.display_name", "displayNameResourceUpdate"),
					resource.TestCheckResourceAttr(datasourceName, "compute_capacity_reservations.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservations.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "compute_capacity_reservations.0.is_default_reservation", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservations.0.reserved_instance_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservations.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservations.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservations.0.used_instance_count"),
				),
			},
			// Step 6: verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ComputeCapacityReservationResourceConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "capacity_reservation_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayNameResourceUpdate"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_reservation_configs.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_reservation_configs.0.fault_domain", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_reservation_configs.0.instance_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_reservation_configs.0.instance_shape_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_reservation_configs.0.instance_shape_config.0.memory_in_gbs", "15"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_reservation_configs.0.instance_shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_reservation_configs.0.reserved_count", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_reservation_configs.0.used_count"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_default_reservation", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "reserved_instance_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					// TODO: investigate why time_updated is not set
					// resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "used_instance_count"),
				),
			},
			// Step 7: remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ComputeCapacityReservationResourceConfig,
			},
			// Step 8: verify resource import
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

func testAccCheckCoreComputeCapacityReservationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_compute_capacity_reservation" {
			noResourceFound = false
			request := oci_core.GetComputeCapacityReservationRequest{}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.CapacityReservationId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetComputeCapacityReservation(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ComputeCapacityReservationLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("CoreComputeCapacityReservation") {
		resource.AddTestSweepers("CoreComputeCapacityReservation", &resource.Sweeper{
			Name:         "CoreComputeCapacityReservation",
			Dependencies: DependencyGraph["computeCapacityReservation"],
			F:            sweepCoreComputeCapacityReservationResource,
		})
	}
}

func sweepCoreComputeCapacityReservationResource(compartment string) error {
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()
	computeCapacityReservationIds, err := getComputeCapacityReservationIds(compartment)
	if err != nil {
		return err
	}
	for _, computeCapacityReservationId := range computeCapacityReservationIds {
		if ok := SweeperDefaultResourceId[computeCapacityReservationId]; !ok {
			deleteComputeCapacityReservationRequest := oci_core.DeleteComputeCapacityReservationRequest{}

			deleteComputeCapacityReservationRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := computeClient.DeleteComputeCapacityReservation(context.Background(), deleteComputeCapacityReservationRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeCapacityReservation %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeCapacityReservationId, error)
				continue
			}
			waitTillCondition(testAccProvider, &computeCapacityReservationId, computeCapacityReservationSweepWaitCondition, time.Duration(3*time.Minute),
				computeCapacityReservationSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getComputeCapacityReservationIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ComputeCapacityReservationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()

	listComputeCapacityReservationsRequest := oci_core.ListComputeCapacityReservationsRequest{}
	listComputeCapacityReservationsRequest.CompartmentId = &compartmentId
	listComputeCapacityReservationsRequest.LifecycleState = oci_core.ComputeCapacityReservationLifecycleStateActive
	listComputeCapacityReservationsResponse, err := computeClient.ListComputeCapacityReservations(context.Background(), listComputeCapacityReservationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ComputeCapacityReservation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, computeCapacityReservation := range listComputeCapacityReservationsResponse.Items {
		id := *computeCapacityReservation.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ComputeCapacityReservationId", id)
	}
	return resourceIds, nil
}

func computeCapacityReservationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if computeCapacityReservationResponse, ok := response.Response.(oci_core.GetComputeCapacityReservationResponse); ok {
		return computeCapacityReservationResponse.LifecycleState != oci_core.ComputeCapacityReservationLifecycleStateDeleted
	}
	return false
}

func computeCapacityReservationSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.computeClient().GetComputeCapacityReservation(context.Background(), oci_core.GetComputeCapacityReservationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
