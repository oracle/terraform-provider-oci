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
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApigatewayUsagePlanRequiredOnlyResource = ApigatewayUsagePlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Required, acctest.Create, ApigatewayUsagePlanRepresentation)

	ApigatewayUsagePlanResourceConfig = ApigatewayUsagePlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Optional, acctest.Update, ApigatewayUsagePlanRepresentation)

	ApigatewayUsagePlanSingularDataSourceRepresentation = map[string]interface{}{
		"usage_plan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_usage_plan.test_usage_plan.id}`},
	}

	ApigatewayUsagePlanDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayUsagePlanDataSourceFilterRepresentation}}
	ApigatewayUsagePlanDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apigateway_usage_plan.test_usage_plan.id}`}},
	}

	ApigatewayUsagePlanRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entitlements":   acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayusagePlanEntitlementsRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesUsagePlanRepresentation},
	}
	ignoreChangesUsagePlanRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	ApigatewayusagePlanEntitlementsRepresentation = map[string]interface{}{
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"quota":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayusagePlanEntitlementsQuotaRepresentation},
		"rate_limit":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayUsagePlanEntitlementsRateLimitRepresentation},
		"targets":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApigatewayUsagePlanEntitlementsTargetsRepresentation},
	}
	ApigatewayusagePlanEntitlementsQuotaRepresentation = map[string]interface{}{
		"operation_on_breach": acctest.Representation{RepType: acctest.Required, Create: `REJECT`, Update: `ALLOW`},
		"reset_policy":        acctest.Representation{RepType: acctest.Required, Create: `CALENDAR`},
		"unit":                acctest.Representation{RepType: acctest.Required, Create: `MINUTE`, Update: `HOUR`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	ApigatewayUsagePlanEntitlementsRateLimitRepresentation = map[string]interface{}{
		"unit":  acctest.Representation{RepType: acctest.Required, Create: `SECOND`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	ApigatewayUsagePlanEntitlementsTargetsRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_deployment.test_deployment.id}`},
	}

	ApigatewayDeploymentRepresentationWithUsagePlan = acctest.GetUpdatedRepresentationCopy(
		"specification.request_policies",
		acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"usage_plans": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
				"token_locations": acctest.Representation{RepType: acctest.Required, Create: []string{`request.headers[apiKeyLocation]`}, Update: []string{`request.path[apiKeyLocation]`}},
			}},
		}},
		ApigatewayDeploymentRepresentation)

	ApigatewayUsagePlanResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", acctest.Required, acctest.Create, ApigatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, ApigatewayDeploymentRepresentationWithUsagePlan) +
		DefinedTagsDependencies
)

// issue-routing-tag: apigateway/default
func TestApigatewayUsagePlanResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayUsagePlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_apigateway_usage_plan.test_usage_plan"
	datasourceName := "data.oci_apigateway_usage_plans.test_usage_plans"
	singularDatasourceName := "data.oci_apigateway_usage_plan.test_usage_plan"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApigatewayUsagePlanResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Optional, acctest.Create, ApigatewayUsagePlanRepresentation), "apigateway", "usagePlan", t)

	acctest.ResourceTest(t, testAccCheckApigatewayUsagePlanDestroy, []resource.TestStep{
		// // verify Create
		// {
		// 	Config: config + compartmentIdVariableStr +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Required, acctest.Create, ApigatewayUsagePlanRepresentation),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		// 		resource.TestCheckResourceAttr(resourceName, "entitlements.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "entitlements.0.name", "name"),

		// 		func(s *terraform.State) (err error) {
		// 			resId, err = acctest.FromInstanceState(s, resourceName, "id")
		// 			return err
		// 		},
		// 	),
		// },

		// // delete before next Create
		// {
		// 	Config: config + compartmentIdVariableStr,
		// },

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApigatewayUsagePlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Optional, acctest.Create, ApigatewayUsagePlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.operation_on_breach", "REJECT"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.reset_policy", "CALENDAR"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.unit", "MINUTE"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.0.unit", "SECOND"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "entitlements.0.targets.0.deployment_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApigatewayUsagePlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ApigatewayUsagePlanRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.operation_on_breach", "REJECT"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.reset_policy", "CALENDAR"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.unit", "MINUTE"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.0.unit", "SECOND"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.0.value", "10"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "entitlements.0.targets.0.deployment_id"),
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
			Config: config + compartmentIdVariableStr + ApigatewayUsagePlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Optional, acctest.Update, ApigatewayUsagePlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.operation_on_breach", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.reset_policy", "CALENDAR"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.unit", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.quota.0.value", "11"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.0.unit", "SECOND"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.rate_limit.0.value", "11"),
				resource.TestCheckResourceAttr(resourceName, "entitlements.0.targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "entitlements.0.targets.0.deployment_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_usage_plans", "test_usage_plans", acctest.Optional, acctest.Update, ApigatewayUsagePlanDataSourceRepresentation) +
				compartmentIdVariableStr + ApigatewayUsagePlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Optional, acctest.Update, ApigatewayUsagePlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "usage_plan_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "usage_plan_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Required, acctest.Create, ApigatewayUsagePlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApigatewayUsagePlanResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "usage_plan_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.quota.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.quota.0.operation_on_breach", "ALLOW"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.quota.0.reset_policy", "CALENDAR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.quota.0.unit", "HOUR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.quota.0.value", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.rate_limit.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.rate_limit.0.unit", "SECOND"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.rate_limit.0.value", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.0.targets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + ApigatewayUsagePlanRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApigatewayUsagePlanDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).UsagePlansClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_usage_plan" {
			noResourceFound = false
			request := oci_apigateway.GetUsagePlanRequest{}

			tmp := rs.Primary.ID
			request.UsagePlanId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apigateway")

			response, err := client.GetUsagePlan(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apigateway.UsagePlanLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ApigatewayUsagePlan") {
		resource.AddTestSweepers("ApigatewayUsagePlan", &resource.Sweeper{
			Name:         "ApigatewayUsagePlan",
			Dependencies: acctest.DependencyGraph["usagePlan"],
			F:            sweepApigatewayUsagePlanResource,
		})
	}
}

func sweepApigatewayUsagePlanResource(compartment string) error {
	usagePlansClient := acctest.GetTestClients(&schema.ResourceData{}).UsagePlansClient()
	usagePlanIds, err := getUsagePlanIds(compartment)
	if err != nil {
		return err
	}
	for _, usagePlanId := range usagePlanIds {
		if ok := acctest.SweeperDefaultResourceId[usagePlanId]; !ok {
			deleteUsagePlanRequest := oci_apigateway.DeleteUsagePlanRequest{}

			deleteUsagePlanRequest.UsagePlanId = &usagePlanId

			deleteUsagePlanRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apigateway")
			_, error := usagePlansClient.DeleteUsagePlan(context.Background(), deleteUsagePlanRequest)
			if error != nil {
				fmt.Printf("Error deleting UsagePlan %s %s, It is possible that the resource is already deleted. Please verify manually \n", usagePlanId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &usagePlanId, usagePlanSweepWaitCondition, time.Duration(3*time.Minute),
				usagePlanSweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getUsagePlanIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UsagePlanId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	usagePlansClient := acctest.GetTestClients(&schema.ResourceData{}).UsagePlansClient()

	listUsagePlansRequest := oci_apigateway.ListUsagePlansRequest{}
	listUsagePlansRequest.CompartmentId = &compartmentId
	listUsagePlansRequest.LifecycleState = oci_apigateway.UsagePlanLifecycleStateActive
	listUsagePlansResponse, err := usagePlansClient.ListUsagePlans(context.Background(), listUsagePlansRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UsagePlan list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, usagePlan := range listUsagePlansResponse.Items {
		id := *usagePlan.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UsagePlanId", id)
	}
	return resourceIds, nil
}

func usagePlanSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if usagePlanResponse, ok := response.Response.(oci_apigateway.GetUsagePlanResponse); ok {
		return usagePlanResponse.LifecycleState != oci_apigateway.UsagePlanLifecycleStateDeleted
	}
	return false
}

func usagePlanSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.UsagePlansClient().GetUsagePlan(context.Background(), oci_apigateway.GetUsagePlanRequest{
		UsagePlanId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
