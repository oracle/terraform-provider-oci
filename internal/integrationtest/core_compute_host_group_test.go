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
	CoreComputeHostGroupRequiredOnlyResource = CoreComputeHostGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Required, acctest.Create, CoreComputeHostGroupRepresentation)

	CoreComputeHostGroupResourceConfig = CoreComputeHostGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Optional, acctest.Update, CoreComputeHostGroupRepresentation)

	CoreComputeHostGroupSingularDataSourceRepresentation = map[string]interface{}{
		"compute_host_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_host_group.test_compute_host_group.id}`},
	}

	CoreComputeHostGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeHostGroupDataSourceFilterRepresentation}}
	CoreComputeHostGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_compute_host_group.test_compute_host_group.id}`}},
	}

	CoreComputeHostGroupRepresentation = map[string]interface{}{
		"availability_domain":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_targeted_placement_required": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"configurations":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreComputeHostGroupConfigurationsRepresentation},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	CoreComputeHostGroupConfigurationsRepresentation = map[string]interface{}{
		"recycle_level": acctest.Representation{RepType: acctest.Optional, Create: `SKIP_RECYCLE`, Update: `FULL_RECYCLE`},
		"target":        acctest.Representation{RepType: acctest.Optional, Create: `BM.GPU.H100.8`, Update: `BM.Standard3.64`},
	}

	CoreComputeHostGroupResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeHostGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeHostGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_compute_host_group.test_compute_host_group"
	datasourceName := "data.oci_core_compute_host_groups.test_compute_host_groups"
	singularDatasourceName := "data.oci_core_compute_host_group.test_compute_host_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreComputeHostGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Optional, acctest.Create, CoreComputeHostGroupRepresentation), "core", "computeHostGroup", t)

	acctest.ResourceTest(t, testAccCheckCoreComputeHostGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeHostGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Required, acctest.Create, CoreComputeHostGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "is_targeted_placement_required", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeHostGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreComputeHostGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Optional, acctest.Create, CoreComputeHostGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.recycle_level", "SKIP_RECYCLE"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.target", "BM.GPU.H100.8"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_targeted_placement_required", "false"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreComputeHostGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreComputeHostGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.recycle_level", "SKIP_RECYCLE"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.target", "BM.GPU.H100.8"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_targeted_placement_required", "false"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreComputeHostGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Optional, acctest.Update, CoreComputeHostGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.recycle_level", "FULL_RECYCLE"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.target", "BM.Standard3.64"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_targeted_placement_required", "true"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_host_groups", "test_compute_host_groups", acctest.Optional, acctest.Update, CoreComputeHostGroupDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeHostGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Optional, acctest.Update, CoreComputeHostGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "compute_host_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_host_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Required, acctest.Create, CoreComputeHostGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeHostGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_host_group_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.recycle_level", "FULL_RECYCLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.target", "BM.Standard3.64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_targeted_placement_required", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreComputeHostGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreComputeHostGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_compute_host_group" {
			noResourceFound = false
			request := oci_core.GetComputeHostGroupRequest{}

			tmp := rs.Primary.ID
			request.ComputeHostGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetComputeHostGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ComputeHostGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CoreComputeHostGroup") {
		resource.AddTestSweepers("CoreComputeHostGroup", &resource.Sweeper{
			Name:         "CoreComputeHostGroup",
			Dependencies: acctest.DependencyGraph["computeHostGroup"],
			F:            sweepCoreComputeHostGroupResource,
		})
	}
}

func sweepCoreComputeHostGroupResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	computeHostGroupIds, err := getCoreComputeHostGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, computeHostGroupId := range computeHostGroupIds {
		if ok := acctest.SweeperDefaultResourceId[computeHostGroupId]; !ok {
			deleteComputeHostGroupRequest := oci_core.DeleteComputeHostGroupRequest{}

			deleteComputeHostGroupRequest.ComputeHostGroupId = &computeHostGroupId

			deleteComputeHostGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteComputeHostGroup(context.Background(), deleteComputeHostGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeHostGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeHostGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &computeHostGroupId, CoreComputeHostGroupSweepWaitCondition, time.Duration(3*time.Minute),
				CoreComputeHostGroupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreComputeHostGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ComputeHostGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listComputeHostGroupsRequest := oci_core.ListComputeHostGroupsRequest{}
	listComputeHostGroupsRequest.CompartmentId = &compartmentId
	listComputeHostGroupsResponse, err := computeClient.ListComputeHostGroups(context.Background(), listComputeHostGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ComputeHostGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, computeHostGroup := range listComputeHostGroupsResponse.Items {
		id := *computeHostGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ComputeHostGroupId", id)
	}
	return resourceIds, nil
}

func CoreComputeHostGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if computeHostGroupResponse, ok := response.Response.(oci_core.GetComputeHostGroupResponse); ok {
		return computeHostGroupResponse.LifecycleState != oci_core.ComputeHostGroupLifecycleStateDeleted
	}
	return false
}

func CoreComputeHostGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetComputeHostGroup(context.Background(), oci_core.GetComputeHostGroupRequest{
		ComputeHostGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
