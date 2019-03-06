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
)

var (
	BudgetRequiredOnlyResource = BudgetResourceDependencies +
		generateResourceFromRepresentationMap("oci_budget_budget", "test_budget", Required, Create, budgetRepresentation)

	BudgetResourceConfig = BudgetResourceDependencies +
		generateResourceFromRepresentationMap("oci_budget_budget", "test_budget", Optional, Update, budgetRepresentation)

	budgetSingularDataSourceRepresentation = map[string]interface{}{
		"budget_id": Representation{repType: Required, create: `${oci_budget_budget.test_budget.id}`},
	}

	budgetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, budgetDataSourceFilterRepresentation}}
	budgetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_budget_budget.test_budget.id}`}},
	}

	budgetRepresentation = map[string]interface{}{
		"amount":                Representation{repType: Required, create: `100`, update: `200`},
		"compartment_id":        Representation{repType: Required, create: `${var.compartment_id}`},
		"reset_period":          Representation{repType: Required, create: `MONTHLY`, update: `MONTHLY`},
		"target_compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":          Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":          Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":         Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}

	BudgetResourceDependencies = DefinedTagsDependencies
)

func TestBudgetBudgetResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_budget_budget.test_budget"
	datasourceName := "data.oci_budget_budgets.test_budgets"
	singularDatasourceName := "data.oci_budget_budget.test_budget"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckBudgetBudgetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
					generateResourceFromRepresentationMap("oci_budget_budget", "test_budget", Required, Create, budgetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "amount", "100"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
					resource.TestCheckResourceAttrSet(resourceName, "target_compartment_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BudgetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
					generateResourceFromRepresentationMap("oci_budget_budget", "test_budget", Optional, Create, budgetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "alert_rule_count"),
					resource.TestCheckResourceAttr(resourceName, "amount", "100"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "target_compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
					generateResourceFromRepresentationMap("oci_budget_budget", "test_budget", Optional, Update, budgetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "alert_rule_count"),
					resource.TestCheckResourceAttr(resourceName, "amount", "200"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "target_compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
					generateDataSourceFromRepresentationMap("oci_budget_budgets", "test_budgets", Optional, Update, budgetDataSourceRepresentation) +
					compartmentIdVariableStr + BudgetResourceDependencies +
					generateResourceFromRepresentationMap("oci_budget_budget", "test_budget", Optional, Update, budgetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "budgets.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.alert_rule_count"),
					resource.TestCheckResourceAttr(datasourceName, "budgets.0.amount", "200"),
					resource.TestCheckResourceAttr(datasourceName, "budgets.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "budgets.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "budgets.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "budgets.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "budgets.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "budgets.0.reset_period", "MONTHLY"),
					resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.target_compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.time_updated"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_budget_budget", "test_budget", Required, Create, budgetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BudgetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "budget_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "target_compartment_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_rule_count"),
					resource.TestCheckResourceAttr(singularDatasourceName, "amount", "200"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "reset_period", "MONTHLY"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + BudgetResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckBudgetBudgetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).budgetClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_budget_budget" {
			noResourceFound = false
			request := oci_budget.GetBudgetRequest{}

			tmp := rs.Primary.ID
			request.BudgetId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "budget")

			_, err := client.GetBudget(context.Background(), request)

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
