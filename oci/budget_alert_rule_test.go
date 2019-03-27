// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	oci_budget "github.com/oracle/oci-go-sdk/budget"
	"github.com/oracle/oci-go-sdk/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AlertRuleRequiredOnlyResource = AlertRuleResourceDependencies +
		generateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", Required, Create, alertRuleRepresentation)

	AlertRuleResourceConfig = AlertRuleResourceDependencies +
		generateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", Optional, Update, alertRuleRepresentation)

	alertRuleSingularDataSourceRepresentation = map[string]interface{}{
		"alert_rule_id": Representation{repType: Required, create: `${oci_budget_alert_rule.test_alert_rule.id}`},
		"budget_id":     Representation{repType: Required, create: `${oci_budget_budget.test_budget.id}`},
	}

	alertRuleDataSourceRepresentation = map[string]interface{}{
		"budget_id":    Representation{repType: Required, create: `${oci_budget_budget.test_budget.id}`},
		"display_name": Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":        Representation{repType: Optional, create: `AVAILABLE`},
		"filter":       RepresentationGroup{Required, alertRuleDataSourceFilterRepresentation}}
	alertRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_budget_alert_rule.test_alert_rule.id}`}},
	}

	alertRuleRepresentation = map[string]interface{}{
		"budget_id":      Representation{repType: Required, create: `${oci_budget_budget.test_budget.id}`},
		"recipients":     Representation{repType: Required, create: `JohnSmith@example.com`, update: `SmithJohn@example.com`},
		"threshold":      Representation{repType: Required, create: `100`, update: `200`},
		"threshold_type": Representation{repType: Required, create: `PERCENTAGE`, update: `ABSOLUTE`},
		"type":           Representation{repType: Required, create: `ACTUAL`, update: `FORECAST`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"message":        Representation{repType: Optional, create: `message`, update: `message2`},
	}

	AlertRuleResourceDependencies = BudgetResourceConfig
)

func TestBudgetAlertRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBudgetAlertRuleResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_budget_alert_rule.test_alert_rule"
	datasourceName := "data.oci_budget_alert_rules.test_alert_rules"
	singularDatasourceName := "data.oci_budget_alert_rule.test_alert_rule"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckBudgetAlertRuleDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AlertRuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", Required, Create, alertRuleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "budget_id"),
					resource.TestCheckResourceAttr(resourceName, "recipients", "JohnSmith@example.com"),
					resource.TestCheckResourceAttr(resourceName, "threshold", "100"),
					resource.TestCheckResourceAttr(resourceName, "threshold_type", "PERCENTAGE"),
					resource.TestCheckResourceAttr(resourceName, "type", "ACTUAL"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AlertRuleResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AlertRuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", Optional, Create, alertRuleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "budget_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AlertRuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", Optional, Update, alertRuleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "budget_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_budget_alert_rules", "test_alert_rules", Optional, Update, alertRuleDataSourceRepresentation) +
					compartmentIdVariableStr + AlertRuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", Optional, Update, alertRuleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "budget_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "alert_rules.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "alert_rules.0.budget_id"),
					resource.TestCheckResourceAttr(datasourceName, "alert_rules.0.defined_tags.%", "1"),
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
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_budget_alert_rule", "test_alert_rule", Required, Create, alertRuleSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AlertRuleResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_rule_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "budget_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AlertRuleResourceConfig,
			},
		},
	})
}

func testAccCheckBudgetAlertRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).budgetClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_budget_alert_rule" {
			noResourceFound = false
			request := oci_budget.GetAlertRuleRequest{}

			tmp := rs.Primary.ID
			request.AlertRuleId = &tmp

			if value, ok := rs.Primary.Attributes["budget_id"]; ok {
				request.BudgetId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "budget")

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
