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
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_core "github.com/oracle/oci-go-sdk/v54/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ComputeCapacityReservationRequiredOnlyResource = ComputeCapacityReservationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservTestCoreComputeCapacityReservationResource_basication", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation)

	ComputeCapacityReservationResourceConfig = ComputeCapacityReservationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Update, computeCapacityReservationRepresentation)

	computeCapacityReservationSingularDataSourceRepresentation = map[string]interface{}{
		"capacity_reservation_id": Representation{RepType: Required, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
	}

	computeCapacityReservationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{RepType: Optional, Create: `displayNameDataSourceCreate`, Update: `displayNameDataSourceUpdate`},
		"state":               Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":              RepresentationGroup{Required, computeCapacityReservationDataSourceFilterRepresentation}}
	computeCapacityReservationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`}},
	}

	computeCapacityReservationRepresentation = map[string]interface{}{
		"availability_domain":          Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":               Representation{RepType: Required, Create: `${var.compartment_id}`},
		"defined_tags":                 Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                 Representation{RepType: Optional, Create: `displayNameResourceCreate`, Update: `displayNameResourceUpdate`},
		"freeform_tags":                Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_reservation_configs": RepresentationGroup{Required, computeCapacityReservationInstanceReservationConfigsRepresentation},
		"is_default_reservation":       Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	computeCapacityReservationInstanceReservationConfigsRepresentation = map[string]interface{}{
		"instance_shape":        Representation{RepType: Required, Create: `VM.Standard2.1`},
		"fault_domain":          Representation{RepType: Optional, Create: `FAULT-DOMAIN-1`},
		"reserved_count":        Representation{RepType: Required, Create: `1`, Update: `2`},
		"instance_shape_config": RepresentationGroup{Optional, computeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation},
	}
	computeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": Representation{RepType: Optional, Create: `15`},
		"ocpus":         Representation{RepType: Optional, Create: `1`},
	}

	ComputeCapacityReservationResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeCapacityReservationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_compute_capacity_reservation.test_compute_capacity_reservation"
	datasourceName := "data.oci_core_compute_capacity_reservations.test_compute_capacity_reservations"
	singularDatasourceName := "data.oci_core_compute_capacity_reservation.test_compute_capacity_reservation"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ComputeCapacityReservationResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Create, computeCapacityReservationRepresentation), "core", "computeCapacityReservation", t)

	ResourceTest(t, testAccCheckCoreComputeCapacityReservationDestroy, []resource.TestStep{
		// Step 0: verify Create
		{
			Config: config + compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 1: delete before next Create
		{
			Config: config + compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies,
		},
		// Step 2: verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Create, computeCapacityReservationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// Step 3: verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ComputeCapacityReservationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Create,
					RepresentationCopyWithNewProperties(computeCapacityReservationRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Update, computeCapacityReservationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservations", "test_compute_capacity_reservations", Optional, Update, computeCapacityReservationDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeCapacityReservationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Optional, Update, computeCapacityReservationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayNameDataSourceUpdate"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "compute_capacity_reservations.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservations.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compute_capacity_reservations.0.compartment_id", compartmentId),
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
				GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeCapacityReservationResourceConfig,
			// Check: ComposeAggregateTestCheckFuncWrapper(),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "capacity_reservation_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
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
	})
}

func testAccCheckCoreComputeCapacityReservationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).computeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_compute_capacity_reservation" {
			noResourceFound = false
			request := oci_core.GetComputeCapacityReservationRequest{}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.CapacityReservationId = &value
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

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
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("CoreComputeCapacityReservation") {
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

			deleteComputeCapacityReservationRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteComputeCapacityReservation(context.Background(), deleteComputeCapacityReservationRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeCapacityReservation %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeCapacityReservationId, error)
				continue
			}
			WaitTillCondition(TestAccProvider, &computeCapacityReservationId, computeCapacityReservationSweepWaitCondition, time.Duration(3*time.Minute),
				computeCapacityReservationSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getComputeCapacityReservationIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ComputeCapacityReservationId")
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
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ComputeCapacityReservationId", id)
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
