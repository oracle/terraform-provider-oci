// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreComputeCapacityReservationRequiredOnlyResource = CoreComputeCapacityReservationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Required, acctest.Create, CoreComputeCapacityReservationRepresentation)

	CoreComputeCapacityReservationResourceConfig = CoreComputeCapacityReservationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Update, CoreComputeCapacityReservationRepresentation)

	CoreCoreComputeCapacityReservationSingularDataSourceRepresentation = map[string]interface{}{
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
	}

	CoreCoreComputeCapacityReservationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayNameDataSourceCreate`, Update: `displayNameDataSourceUpdate`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeCapacityReservationDataSourceFilterRepresentation}}
	CoreComputeCapacityReservationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`}},
	}

	CoreComputeCapacityReservationRepresentation = map[string]interface{}{
		"availability_domain":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayNameResourceCreate`, Update: `displayNameResourceUpdate`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_reservation_configs": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeCapacityReservationInstanceReservationConfigsRepresentation},
		"is_default_reservation":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	CoreComputeCapacityReservationInstanceReservationConfigsRepresentation = map[string]interface{}{
		"instance_shape":             acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"fault_domain":               acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"cluster_placement_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.cluster_placement_group_id}`},
		"reserved_count":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"instance_shape_config":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreComputeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation},
	}

	CoreComputeCapacityReservationInstanceReservationConfigsRepresentation2 = map[string]interface{}{
		"instance_shape":             acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`},
		"fault_domain":               acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"cluster_placement_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.cluster_placement_group_id2}`},
		"reserved_count":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"instance_shape_config":      acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation2},
	}

	CoreComputeCapacityReservationInstanceReservationConfigsRepresentation3 = map[string]interface{}{
		"instance_shape":             acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`},
		"fault_domain":               acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"cluster_placement_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.cluster_placement_group_id3}`},
		"reserved_count":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"instance_shape_config":      acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation3},
	}

	CoreComputeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `15`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	CoreComputeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation2 = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `24`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `2`},
	}

	CoreComputeCapacityReservationInstanceReservationConfigsInstanceShapeConfigRepresentation3 = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `36`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `3`},
	}

	CoreComputeCapacityReservationResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeCapacityReservationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	clusterPlacementGroupId := utils.GetEnvSettingWithBlankDefault("cluster_placement_group_ocid")
	clusterPlacementGroupIdStr := fmt.Sprintf("variable \"cluster_placement_group_id\" { default = \"%s\" }\n", clusterPlacementGroupId)

	resourceName := "oci_core_compute_capacity_reservation.test_compute_capacity_reservation"
	datasourceName := "data.oci_core_compute_capacity_reservations.test_compute_capacity_reservations"
	singularDatasourceName := "data.oci_core_compute_capacity_reservation.test_compute_capacity_reservation"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+clusterPlacementGroupIdStr+CoreComputeCapacityReservationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Create, CoreComputeCapacityReservationRepresentation), "core", "computeCapacityReservation", t)

	acctest.ResourceTest(t, testAccCheckCoreComputeCapacityReservationDestroy, []resource.TestStep{
		// Step 0: verify Create
		{
			Config: config + compartmentIdVariableStr + clusterPlacementGroupIdStr + CoreComputeCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Required, acctest.Create, CoreComputeCapacityReservationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 1: Add instance reservation config
		{
			Config: config + compartmentIdVariableStr + clusterPlacementGroupIdStr + CoreComputeCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreComputeCapacityReservationRepresentation, map[string]interface{}{
					"instance_reservation_configs": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: CoreComputeCapacityReservationInstanceReservationConfigsRepresentation}, {RepType: acctest.Optional, Group: CoreComputeCapacityReservationInstanceReservationConfigsRepresentation2}},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.#", "2"),
			),
		},

		// Step 2: verify unordered attach
		{
			Config: config + compartmentIdVariableStr + clusterPlacementGroupIdStr + CoreComputeCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreComputeCapacityReservationRepresentation, map[string]interface{}{
					"instance_reservation_configs": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: CoreComputeCapacityReservationInstanceReservationConfigsRepresentation}, {RepType: acctest.Optional, Group: CoreComputeCapacityReservationInstanceReservationConfigsRepresentation3}, {RepType: acctest.Optional, Group: CoreComputeCapacityReservationInstanceReservationConfigsRepresentation2}},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.#", "3"),
			),
		},

		// Step 3: delete before next Create
		{
			Config: config + compartmentIdVariableStr + clusterPlacementGroupIdStr + CoreComputeCapacityReservationResourceDependencies,
		},
		// Step 4: verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + clusterPlacementGroupIdStr + CoreComputeCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Create, CoreComputeCapacityReservationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameResourceCreate"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_reservation_configs.0.cluster_placement_group_id"),
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// Step 5: verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + clusterPlacementGroupIdStr + CoreComputeCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreComputeCapacityReservationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameResourceCreate"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_reservation_configs.0.cluster_placement_group_id"),
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 6: verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreComputeCapacityReservationResourceDependencies + clusterPlacementGroupIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Update, CoreComputeCapacityReservationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameResourceUpdate"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_reservation_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_reservation_configs.0.cluster_placement_group_id"),
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// Step 7: verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservations", "test_compute_capacity_reservations", acctest.Optional, acctest.Update, CoreCoreComputeCapacityReservationDataSourceRepresentation) +
				compartmentIdVariableStr + clusterPlacementGroupIdStr + CoreComputeCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Update, CoreComputeCapacityReservationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_placement_group_id", clusterPlacementGroupId),
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
		// Step 8: verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Required, acctest.Create, CoreCoreComputeCapacityReservationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + clusterPlacementGroupIdStr + CoreComputeCapacityReservationResourceConfig,
			// Check: acctest.ComposeAggregateTestCheckFuncWrapper(),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "capacity_reservation_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_placement_group_id", clusterPlacementGroupId),
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
		// verify resource import

		{
			Config:                  config + CoreComputeCapacityReservationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreComputeCapacityReservationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_compute_capacity_reservation" {
			noResourceFound = false
			request := oci_core.GetComputeCapacityReservationRequest{}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.CapacityReservationId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreComputeCapacityReservation") {
		resource.AddTestSweepers("CoreComputeCapacityReservation", &resource.Sweeper{
			Name:         "CoreComputeCapacityReservation",
			Dependencies: acctest.DependencyGraph["computeCapacityReservation"],
			F:            sweepCoreComputeCapacityReservationResource,
		})
	}
}

func sweepCoreComputeCapacityReservationResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	computeCapacityReservationIds, err := getCoreComputeCapacityReservationIds(compartment)
	if err != nil {
		return err
	}
	for _, computeCapacityReservationId := range computeCapacityReservationIds {
		if ok := acctest.SweeperDefaultResourceId[computeCapacityReservationId]; !ok {
			deleteComputeCapacityReservationRequest := oci_core.DeleteComputeCapacityReservationRequest{}

			deleteComputeCapacityReservationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteComputeCapacityReservation(context.Background(), deleteComputeCapacityReservationRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeCapacityReservation %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeCapacityReservationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &computeCapacityReservationId, CoreComputeCapacityReservationSweepWaitCondition, time.Duration(3*time.Minute),
				CoreComputeCapacityReservationSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreComputeCapacityReservationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ComputeCapacityReservationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ComputeCapacityReservationId", id)
	}
	return resourceIds, nil
}

func CoreComputeCapacityReservationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if computeCapacityReservationResponse, ok := response.Response.(oci_core.GetComputeCapacityReservationResponse); ok {
		return computeCapacityReservationResponse.LifecycleState != oci_core.ComputeCapacityReservationLifecycleStateDeleted
	}
	return false
}

func CoreComputeCapacityReservationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetComputeCapacityReservation(context.Background(), oci_core.GetComputeCapacityReservationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
