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
	CoreDrgNatPolicyRequiredOnlyResource = CoreDrgNatPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Required, acctest.Create, CoreDrgNatPolicyRepresentation)

	CoreDrgNatPolicyResourceConfig = CoreDrgNatPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Optional, acctest.Update, CoreDrgNatPolicyRepresentation)

	CoreDrgNatPolicySingularDataSourceRepresentation = map[string]interface{}{
		"drg_nat_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_nat_policy.test_drg_nat_policy.id}`},
	}

	CoreDrgNatPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgNatPolicyDataSourceFilterRepresentation}}
	CoreDrgNatPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_drg_nat_policy.test_drg_nat_policy.id}`}},
	}

	CoreDrgNatPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CoreDrgNatPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/pnp
func TestCoreDrgNatPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgNatPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_drg_nat_policy.test_drg_nat_policy"
	datasourceName := "data.oci_core_drg_nat_policies.test_drg_nat_policies"
	singularDatasourceName := "data.oci_core_drg_nat_policy.test_drg_nat_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDrgNatPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Optional, acctest.Create, CoreDrgNatPolicyRepresentation), "core", "drgNatPolicy", t)

	acctest.ResourceTest(t, testAccCheckCoreDrgNatPolicyDestroy, []resource.TestStep{

		// verify Create with no triggers
		{
			Config: config + compartmentIdVariableStr + CoreDrgNatPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Required, acctest.Create, CoreDrgNatPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgNatPolicyResourceDependencies,
		},

		// verify Create with policy fields optionals
		{
			Config: config + compartmentIdVariableStr + CoreDrgNatPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Optional, acctest.Create, CoreDrgNatPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreDrgNatPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreDrgNatPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreDrgNatPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Optional, acctest.Update, CoreDrgNatPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_nat_policies", "test_drg_nat_policies", acctest.Optional, acctest.Update, CoreDrgNatPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgNatPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Optional, acctest.Update, CoreDrgNatPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "drg_nat_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_policies.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "drg_nat_policies.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_nat_policies.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_nat_policies.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_nat_policies.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_nat_policy", "test_drg_nat_policy", acctest.Required, acctest.Create, CoreDrgNatPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgNatPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "drg_nat_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreDrgNatPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreDrgNatPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg_nat_policy" {
			noResourceFound = false
			request := oci_core.GetDrgNatPolicyRequest{}

			tmp := rs.Primary.ID
			request.DrgNatPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetDrgNatPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DrgNatPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CoreDrgNatPolicy") {
		resource.AddTestSweepers("CoreDrgNatPolicy", &resource.Sweeper{
			Name:         "CoreDrgNatPolicy",
			Dependencies: acctest.DependencyGraph["drgNatPolicy"],
			F:            sweepCoreDrgNatPolicyResource,
		})
	}
}

func sweepCoreDrgNatPolicyResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	drgNatPolicyIds, err := getCoreDrgNatPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, drgNatPolicyId := range drgNatPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[drgNatPolicyId]; !ok {
			deleteDrgNatPolicyRequest := oci_core.DeleteDrgNatPolicyRequest{}

			deleteDrgNatPolicyRequest.DrgNatPolicyId = &drgNatPolicyId

			deleteDrgNatPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDrgNatPolicy(context.Background(), deleteDrgNatPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting DrgNatPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", drgNatPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &drgNatPolicyId, CoreDrgNatPolicySweepWaitCondition, time.Duration(3*time.Minute),
				CoreDrgNatPolicySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreDrgNatPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DrgNatPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listDrgNatPoliciesRequest := oci_core.ListDrgNatPoliciesRequest{}
	listDrgNatPoliciesRequest.CompartmentId = &compartmentId
	listDrgNatPoliciesResponse, err := virtualNetworkClient.ListDrgNatPolicies(context.Background(), listDrgNatPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DrgNatPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, drgNatPolicy := range listDrgNatPoliciesResponse.Items {
		id := *drgNatPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DrgNatPolicyId", id)
	}
	return resourceIds, nil
}

func CoreDrgNatPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drgNatPolicyResponse, ok := response.Response.(oci_core.GetDrgNatPolicyResponse); ok {
		return drgNatPolicyResponse.LifecycleState != oci_core.DrgNatPolicyLifecycleStateDeleted
	}
	return false
}

func CoreDrgNatPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetDrgNatPolicy(context.Background(), oci_core.GetDrgNatPolicyRequest{
		DrgNatPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
