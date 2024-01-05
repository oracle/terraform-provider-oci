// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	CoreComputeCapacityTopologyRequiredOnlyResource = CoreComputeCapacityTopologyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Required, acctest.Create, CoreComputeCapacityTopologyRepresentation)

	CoreComputeCapacityTopologyResourceConfig = CoreComputeCapacityTopologyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Optional, acctest.Update, CoreComputeCapacityTopologyRepresentation)

	CoreComputeCapacityTopologySingularDataSourceRepresentation = map[string]interface{}{
		"compute_capacity_topology_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_capacity_topology.test_compute_capacity_topology.id}`},
	}

	CoreComputeCapacityTopologyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeCapacityTopologyDataSourceFilterRepresentation}}
	CoreComputeCapacityTopologyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_compute_capacity_topology.test_compute_capacity_topology.id}`}},
	}

	CoreComputeCapacityTopologyRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"capacity_source":     acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeCapacityTopologyCapacitySourceRepresentation},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	CoreComputeCapacityTopologyCapacitySourceRepresentation = map[string]interface{}{
		"capacity_type":  acctest.Representation{RepType: acctest.Required, Create: `DEDICATED`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	CoreComputeCapacityTopologyResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeBm
func TestCoreComputeCapacityTopologyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityTopologyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_compute_capacity_topology.test_compute_capacity_topology"
	datasourceName := "data.oci_core_compute_capacity_topologies.test_compute_capacity_topologies"
	singularDatasourceName := "data.oci_core_compute_capacity_topology.test_compute_capacity_topology"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreComputeCapacityTopologyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Optional, acctest.Create, CoreComputeCapacityTopologyRepresentation), "core", "computeCapacityTopology", t)

	acctest.ResourceTest(t, testAccCheckCoreComputeCapacityTopologyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeCapacityTopologyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Required, acctest.Create, CoreComputeCapacityTopologyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.0.capacity_type", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeCapacityTopologyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreComputeCapacityTopologyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Optional, acctest.Create, CoreComputeCapacityTopologyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.0.capacity_type", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + CoreComputeCapacityTopologyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Optional, acctest.Create, CoreComputeCapacityTopologyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.0.capacity_type", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreComputeCapacityTopologyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Optional, acctest.Update, CoreComputeCapacityTopologyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.0.capacity_type", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "capacity_source.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_topologies", "test_compute_capacity_topologies", acctest.Optional, acctest.Update, CoreComputeCapacityTopologyDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeCapacityTopologyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Optional, acctest.Update, CoreComputeCapacityTopologyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "compute_capacity_topology_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_capacity_topology_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Required, acctest.Create, CoreComputeCapacityTopologySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeCapacityTopologyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_capacity_topology_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity_source.0.capacity_type", "DEDICATED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity_source.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreComputeCapacityTopologyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreComputeCapacityTopologyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_compute_capacity_topology" {
			noResourceFound = false
			request := oci_core.GetComputeCapacityTopologyRequest{}

			tmp := rs.Primary.ID
			request.ComputeCapacityTopologyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetComputeCapacityTopology(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ComputeCapacityTopologyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CoreComputeCapacityTopology") {
		resource.AddTestSweepers("CoreComputeCapacityTopology", &resource.Sweeper{
			Name:         "CoreComputeCapacityTopology",
			Dependencies: acctest.DependencyGraph["computeCapacityTopology"],
			F:            sweepCoreComputeCapacityTopologyResource,
		})
	}
}

func sweepCoreComputeCapacityTopologyResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	computeCapacityTopologyIds, err := getCoreComputeCapacityTopologyIds(compartment)
	if err != nil {
		return err
	}
	for _, computeCapacityTopologyId := range computeCapacityTopologyIds {
		if ok := acctest.SweeperDefaultResourceId[computeCapacityTopologyId]; !ok {
			deleteComputeCapacityTopologyRequest := oci_core.DeleteComputeCapacityTopologyRequest{}

			deleteComputeCapacityTopologyRequest.ComputeCapacityTopologyId = &computeCapacityTopologyId

			deleteComputeCapacityTopologyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteComputeCapacityTopology(context.Background(), deleteComputeCapacityTopologyRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeCapacityTopology %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeCapacityTopologyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &computeCapacityTopologyId, CoreComputeCapacityTopologySweepWaitCondition, time.Duration(3*time.Minute),
				CoreComputeCapacityTopologySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreComputeCapacityTopologyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ComputeCapacityTopologyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listComputeCapacityTopologiesRequest := oci_core.ListComputeCapacityTopologiesRequest{}
	listComputeCapacityTopologiesRequest.CompartmentId = &compartmentId
	listComputeCapacityTopologiesResponse, err := computeClient.ListComputeCapacityTopologies(context.Background(), listComputeCapacityTopologiesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ComputeCapacityTopology list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, computeCapacityTopology := range listComputeCapacityTopologiesResponse.Items {
		id := *computeCapacityTopology.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ComputeCapacityTopologyId", id)
	}
	return resourceIds, nil
}

func CoreComputeCapacityTopologySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if computeCapacityTopologyResponse, ok := response.Response.(oci_core.GetComputeCapacityTopologyResponse); ok {
		return computeCapacityTopologyResponse.LifecycleState != oci_core.ComputeCapacityTopologyLifecycleStateDeleted
	}
	return false
}

func CoreComputeCapacityTopologySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetComputeCapacityTopology(context.Background(), oci_core.GetComputeCapacityTopologyRequest{
		ComputeCapacityTopologyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
