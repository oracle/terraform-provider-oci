// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	AlertRuleRequiredOnlyResource = AlertRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", acctest.Required, acctest.Create, alertRuleRepresentation)

	AlertRuleResourceConfig = AlertRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", acctest.Optional, acctest.Update, alertRuleRepresentation)

	alertRuleSingularDataSourceRepresentation = map[string]interface{}{
		"alert_rule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_alert_rule.test_alert_rule.id}`},
		"budget_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_budget.test_budget.id}`},
	}

	alertRuleDataSourceRepresentation = map[string]interface{}{
		"budget_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_budget.test_budget.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: alertRuleDataSourceFilterRepresentation}}
	alertRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_budget_alert_rule.test_alert_rule.id}`}},
	}

	alertRuleRepresentation = map[string]interface{}{
		"budget_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_budget.test_budget.id}`},
		"threshold":      acctest.Representation{RepType: acctest.Required, Create: `100`, Update: `200`},
		"threshold_type": acctest.Representation{RepType: acctest.Required, Create: `PERCENTAGE`, Update: `ABSOLUTE`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `ACTUAL`, Update: `FORECAST`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"message":        acctest.Representation{RepType: acctest.Optional, Create: `message`, Update: `message2`},
		"recipients":     acctest.Representation{RepType: acctest.Optional, Create: `JohnSmith@example.com`, Update: `SmithJohn@example.com`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagChange},
	}

	AlertRuleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Required, acctest.Create, budgetRepresentationWithTargetCompartmentId) +
		DefinedTagsDependencies
)

// issue-routing-tag: budget/default
func TestBudgetAlertRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBudgetAlertRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_budget_alert_rule.test_alert_rule"
	datasourceName := "data.oci_budget_alert_rules.test_alert_rules"
	singularDatasourceName := "data.oci_budget_alert_rule.test_alert_rule"

	var resId, resId2 string
	var compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AlertRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", acctest.Optional, acctest.Create, alertRuleRepresentation), "budget", "alertRule", t)

	acctest.ResourceTest(t, testAccCheckBudgetAlertRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AlertRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", acctest.Required, acctest.Create, alertRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "budget_id"),
				resource.TestCheckResourceAttr(resourceName, "threshold", "100"),
				resource.TestCheckResourceAttr(resourceName, "threshold_type", "PERCENTAGE"),
				resource.TestCheckResourceAttr(resourceName, "type", "ACTUAL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AlertRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AlertRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", acctest.Optional, acctest.Create, alertRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "budget_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "message", "message"),
				resource.TestCheckResourceAttr(resourceName, "recipients", "JohnSmith@example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "threshold", "100"),
				resource.TestCheckResourceAttr(resourceName, "threshold_type", "PERCENTAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ACTUAL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					budgetId, _ := acctest.FromInstanceState(s, resourceName, "budget_id")
					compositeId = "budgets/" + budgetId + "/alertRules/" + resId
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AlertRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", acctest.Optional, acctest.Update, alertRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "budget_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "message", "message2"),
				resource.TestCheckResourceAttr(resourceName, "recipients", "SmithJohn@example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "threshold", "200"),
				resource.TestCheckResourceAttr(resourceName, "threshold_type", "ABSOLUTE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "FORECAST"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_alert_rules", "test_alert_rules", acctest.Optional, acctest.Update, alertRuleDataSourceRepresentation) +
				compartmentIdVariableStr + AlertRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", acctest.Optional, acctest.Update, alertRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "budget_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "alert_rules.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_rules.0.budget_id"),
				resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_rules.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.message", "message2"),
				resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.recipients", "SmithJohn@example.com"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_rules.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.threshold", "200"),
				resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.threshold_type", "ABSOLUTE"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_rules.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_rules.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.type", "FORECAST"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_rules.0.version"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", acctest.Required, acctest.Create, alertRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AlertRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_rule_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "budget_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "message", "message2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recipients", "SmithJohn@example.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "threshold", "200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "threshold_type", "ABSOLUTE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "FORECAST"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config:                  config + AlertRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getAlertRuleImportId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getAlertRuleImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("budgets/" + rs.Primary.Attributes["budget_id"] + "/alertRules/" + rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckBudgetAlertRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BudgetClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_budget_alert_rule" {
			noResourceFound = false
			request := oci_budget.GetAlertRuleRequest{}

			tmp := rs.Primary.ID
			request.AlertRuleId = &tmp

			if value, ok := rs.Primary.Attributes["budget_id"]; ok {
				request.BudgetId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "budget")

			_, err := client.GetAlertRule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("BudgetAlertRule") {
		resource.AddTestSweepers("BudgetAlertRule", &resource.Sweeper{
			Name:         "BudgetAlertRule",
			Dependencies: acctest.DependencyGraph["alertRule"],
			F:            sweepBudgetAlertRuleResource,
		})
	}
}

func sweepBudgetAlertRuleResource(compartment string) error {
	budgetClient := acctest.GetTestClients(&schema.ResourceData{}).BudgetClient()
	alertRuleIds, err := getAlertRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, alertRuleId := range alertRuleIds {
		if ok := acctest.SweeperDefaultResourceId[alertRuleId]; !ok {
			deleteAlertRuleRequest := oci_budget.DeleteAlertRuleRequest{}

			deleteAlertRuleRequest.AlertRuleId = &alertRuleId

			deleteAlertRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "budget")
			_, error := budgetClient.DeleteAlertRule(context.Background(), deleteAlertRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting AlertRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", alertRuleId, error)
				continue
			}
		}
	}
	return nil
}

func getAlertRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AlertRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	budgetClient := acctest.GetTestClients(&schema.ResourceData{}).BudgetClient()

	listAlertRulesRequest := oci_budget.ListAlertRulesRequest{}

	budgetIds, error := getBudgetIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting budgetId required for AlertRule resource requests \n")
	}
	for _, budgetId := range budgetIds {
		listAlertRulesRequest.BudgetId = &budgetId

		listAlertRulesResponse, err := budgetClient.ListAlertRules(context.Background(), listAlertRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting AlertRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, alertRule := range listAlertRulesResponse.Items {
			id := *alertRule.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AlertRuleId", id)
		}

	}
	return resourceIds, nil
}
