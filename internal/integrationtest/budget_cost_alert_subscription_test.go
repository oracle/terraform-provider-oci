// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BudgetCostAlertSubscriptionRequiredOnlyResource = BudgetCostAlertSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Required, acctest.Create, BudgetCostAlertSubscriptionRepresentation)

	BudgetCostAlertSubscriptionResourceConfig = BudgetCostAlertSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Optional, acctest.Update, BudgetCostAlertSubscriptionRepresentation)

	BudgetCostAlertSubscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"cost_alert_subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_cost_alert_subscription.test_cost_alert_subscription.id}`},
	}

	BudgetCostAlertSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `Department A email list`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BudgetCostAlertSubscriptionDataSourceFilterRepresentation}}
	BudgetCostAlertSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_budget_cost_alert_subscription.test_cost_alert_subscription.id}`}},
	}

	BudgetCostAlertSubscriptionRepresentation = map[string]interface{}{
		"channels":       acctest.Representation{RepType: acctest.Required, Create: `{\"email\":\"test@example.com\"}`, Update: `{\"email\":\"updated@example.com\"}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `Department A email list`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	//BudgetCostAlertSubscriptionResourceDependencies = DefinedTagsDependencies
	BudgetCostAlertSubscriptionResourceDependencies = ""
)

// issue-routing-tag: budget/default
func TestBudgetCostAlertSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBudgetCostAlertSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_budget_cost_alert_subscription.test_cost_alert_subscription"
	datasourceName := "data.oci_budget_cost_alert_subscriptions.test_cost_alert_subscriptions"
	singularDatasourceName := "data.oci_budget_cost_alert_subscription.test_cost_alert_subscription"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BudgetCostAlertSubscriptionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Optional, acctest.Create, BudgetCostAlertSubscriptionRepresentation), "budget", "costAlertSubscription", t)

	acctest.ResourceTest(t, testAccCheckBudgetCostAlertSubscriptionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BudgetCostAlertSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Required, acctest.Create, BudgetCostAlertSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channels", `{"email":"test@example.com"}`),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "Department A email list"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BudgetCostAlertSubscriptionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BudgetCostAlertSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Optional, acctest.Create, BudgetCostAlertSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channels", `{"email":"test@example.com"}`),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "Department A email list"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BudgetCostAlertSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Optional, acctest.Update, BudgetCostAlertSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channels", `{"email":"updated@example.com"}`),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "Department A email list"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_alert_subscriptions", "test_cost_alert_subscriptions", acctest.Optional, acctest.Update, BudgetCostAlertSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + BudgetCostAlertSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Optional, acctest.Update, BudgetCostAlertSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "Department A email list"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "cost_alert_subscription_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cost_alert_subscription_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Required, acctest.Create, BudgetCostAlertSubscriptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BudgetCostAlertSubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_alert_subscription_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "channels", `{"email":"updated@example.com"}`),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "Department A email list"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + BudgetCostAlertSubscriptionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBudgetCostAlertSubscriptionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CostAdClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_budget_cost_alert_subscription" {
			noResourceFound = false
			request := oci_budget.GetCostAlertSubscriptionRequest{}

			tmp := rs.Primary.ID
			request.CostAlertSubscriptionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "budget")

			_, err := client.GetCostAlertSubscription(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("BudgetCostAlertSubscription") {
		resource.AddTestSweepers("BudgetCostAlertSubscription", &resource.Sweeper{
			Name:         "BudgetCostAlertSubscription",
			Dependencies: acctest.DependencyGraph["costAlertSubscription"],
			F:            sweepBudgetCostAlertSubscriptionResource,
		})
	}
}

func sweepBudgetCostAlertSubscriptionResource(compartment string) error {
	costAdClient := acctest.GetTestClients(&schema.ResourceData{}).CostAdClient()
	costAlertSubscriptionIds, err := getBudgetCostAlertSubscriptionIds(compartment)
	if err != nil {
		return err
	}
	for _, costAlertSubscriptionId := range costAlertSubscriptionIds {
		if ok := acctest.SweeperDefaultResourceId[costAlertSubscriptionId]; !ok {
			deleteCostAlertSubscriptionRequest := oci_budget.DeleteCostAlertSubscriptionRequest{}

			deleteCostAlertSubscriptionRequest.CostAlertSubscriptionId = &costAlertSubscriptionId

			deleteCostAlertSubscriptionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "budget")
			_, error := costAdClient.DeleteCostAlertSubscription(context.Background(), deleteCostAlertSubscriptionRequest)
			if error != nil {
				fmt.Printf("Error deleting CostAlertSubscription %s %s, It is possible that the resource is already deleted. Please verify manually \n", costAlertSubscriptionId, error)
				continue
			}
		}
	}
	return nil
}

func getBudgetCostAlertSubscriptionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CostAlertSubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	costAdClient := acctest.GetTestClients(&schema.ResourceData{}).CostAdClient()

	listCostAlertSubscriptionsRequest := oci_budget.ListCostAlertSubscriptionsRequest{}
	listCostAlertSubscriptionsRequest.CompartmentId = &compartmentId
	listCostAlertSubscriptionsResponse, err := costAdClient.ListCostAlertSubscriptions(context.Background(), listCostAlertSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CostAlertSubscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, costAlertSubscription := range listCostAlertSubscriptionsResponse.Items {
		id := *costAlertSubscription.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CostAlertSubscriptionId", id)
	}
	return resourceIds, nil
}
