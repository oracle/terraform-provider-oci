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
	oci_zpr "github.com/oracle/oci-go-sdk/v65/zpr"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreChangesZprPolicyRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags", "system_tags"}},
	}

	ZprZprPolicyRequiredOnlyResource = ZprZprPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_zpr_zpr_policy", "test_zpr_policy", acctest.Required, acctest.Create, ZprZprPolicyRepresentation)

	ZprZprPolicyResourceConfig = ZprZprPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_zpr_zpr_policy", "test_zpr_policy", acctest.Optional, acctest.Update, ZprZprPolicyRepresentation)

	ZprZprPolicySingularDataSourceRepresentation = map[string]interface{}{
		"zpr_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_zpr_zpr_policy.test_zpr_policy.id}`},
	}

	ZprZprPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ZprZprPolicyDataSourceFilterRepresentation}}
	ZprZprPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_zpr_zpr_policy.test_zpr_policy.id}`}},
	}

	ZprZprPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"statements":     acctest.Representation{RepType: acctest.Required, Create: []string{`in prod:red VCN allow app:front-end endpoints to connect to db:back-end endpoints`}, Update: []string{`in prod:red VCN allow app:front-end endpoints to connect to db:back-end endpoints`, `in prod:red VCN allow app:front-end endpoints to connect to db:front-end endpoints`}},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesZprPolicyRepresentation},
	}

	ZprZprPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: zpr/default
func TestZprZprPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestZprZprPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_zpr_zpr_policy.test_zpr_policy"
	datasourceName := "data.oci_zpr_zpr_policies.test_zpr_policies"
	singularDatasourceName := "data.oci_zpr_zpr_policy.test_zpr_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ZprZprPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_zpr_zpr_policy", "test_zpr_policy", acctest.Optional, acctest.Create, ZprZprPolicyRepresentation), "zpr", "zprPolicy", t)

	acctest.ResourceTest(t, testAccCheckZprZprPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ZprZprPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_zpr_zpr_policy", "test_zpr_policy", acctest.Required, acctest.Create, ZprZprPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ZprZprPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ZprZprPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_zpr_zpr_policy", "test_zpr_policy", acctest.Optional, acctest.Create, ZprZprPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ZprZprPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_zpr_zpr_policy", "test_zpr_policy", acctest.Optional, acctest.Update, ZprZprPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_zpr_zpr_policies", "test_zpr_policies", acctest.Optional, acctest.Update, ZprZprPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + ZprZprPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_zpr_zpr_policy", "test_zpr_policy", acctest.Optional, acctest.Update, ZprZprPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "zpr_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zpr_policies.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_zpr_zpr_policy", "test_zpr_policy", acctest.Required, acctest.Create, ZprZprPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ZprZprPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "zpr_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "statements.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ZprZprPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckZprZprPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ZprClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_zpr_zpr_policy" {
			noResourceFound = false
			request := oci_zpr.GetZprPolicyRequest{}

			tmp := rs.Primary.ID
			request.ZprPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "zpr")

			response, err := client.GetZprPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_zpr.ZprPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ZprZprPolicy") {
		resource.AddTestSweepers("ZprZprPolicy", &resource.Sweeper{
			Name:         "ZprZprPolicy",
			Dependencies: acctest.DependencyGraph["zprPolicy"],
			F:            sweepZprZprPolicyResource,
		})
	}
}

func sweepZprZprPolicyResource(compartment string) error {
	zprClient := acctest.GetTestClients(&schema.ResourceData{}).ZprClient()
	// ZprZprPolicyResource can only run on root compartment
	compartment = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	zprPolicyIds, err := getZprZprPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, zprPolicyId := range zprPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[zprPolicyId]; !ok {
			deleteZprPolicyRequest := oci_zpr.DeleteZprPolicyRequest{}

			deleteZprPolicyRequest.ZprPolicyId = &zprPolicyId

			deleteZprPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "zpr")
			_, error := zprClient.DeleteZprPolicy(context.Background(), deleteZprPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting ZprPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", zprPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &zprPolicyId, ZprZprPolicySweepWaitCondition, time.Duration(3*time.Minute),
				ZprZprPolicySweepResponseFetchOperation, "zpr", true)
		}
	}
	return nil
}

func getZprZprPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ZprPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	zprClient := acctest.GetTestClients(&schema.ResourceData{}).ZprClient()

	listZprPoliciesRequest := oci_zpr.ListZprPoliciesRequest{}
	listZprPoliciesRequest.CompartmentId = &compartmentId
	listZprPoliciesRequest.LifecycleState = oci_zpr.ZprPolicyLifecycleStateNeedsAttention
	listZprPoliciesResponse, err := zprClient.ListZprPolicies(context.Background(), listZprPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ZprPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, zprPolicy := range listZprPoliciesResponse.Items {
		id := *zprPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ZprPolicyId", id)
	}
	return resourceIds, nil
}

func ZprZprPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if zprPolicyResponse, ok := response.Response.(oci_zpr.GetZprPolicyResponse); ok {
		return zprPolicyResponse.LifecycleState != oci_zpr.ZprPolicyLifecycleStateDeleted
	}
	return false
}

func ZprZprPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ZprClient().GetZprPolicy(context.Background(), oci_zpr.GetZprPolicyRequest{
		ZprPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
