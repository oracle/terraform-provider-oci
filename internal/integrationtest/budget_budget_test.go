// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_budget "github.com/oracle/oci-go-sdk/v58/budget"
	"github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BudgetRequiredOnlyResource = BudgetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Required, acctest.Create, budgetRepresentationWithTargetCompartmentId)

	BudgetResourceConfig = BudgetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Update, budgetRepresentationWithTargetCompartmentId)

	budgetSingularDataSourceRepresentation = map[string]interface{}{
		"budget_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_budget.test_budget.id}`},
	}

	budgetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_type":    acctest.Representation{RepType: acctest.Optional, Create: `COMPARTMENT`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: budgetDataSourceFilterRepresentation}}
	budgetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_budget_budget.test_budget.id}`}},
	}

	//Service required target_compartment_id or targets to be set. Both cannot be empty
	budgetRepresentationWithTargetCompartmentId = map[string]interface{}{
		"amount":                                acctest.Representation{RepType: acctest.Required, Create: `100`, Update: `200`},
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"reset_period":                          acctest.Representation{RepType: acctest.Required, Create: `MONTHLY`},
		"budget_processing_period_start_offset": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"defined_tags":                          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"target_compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	//Budget with target_type = COMPARTMENT
	budgetRepresentationWithTargetTypeAsCompartmentAndTargets = map[string]interface{}{
		"amount":         acctest.Representation{RepType: acctest.Required, Create: `100`, Update: `200`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"reset_period":   acctest.Representation{RepType: acctest.Required, Create: `MONTHLY`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"target_type":    acctest.Representation{RepType: acctest.Required, Create: `COMPARTMENT`},
		"targets":        acctest.Representation{RepType: acctest.Required, Create: []string{`${var.compartment_id}`}},
	}

	//Budget with target_type = TAG
	budgetRepresentationWithTargetTypeAsTagAndTargets = map[string]interface{}{
		"amount":         acctest.Representation{RepType: acctest.Required, Create: `100`, Update: `200`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"reset_period":   acctest.Representation{RepType: acctest.Required, Create: `MONTHLY`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"target_type":    acctest.Representation{RepType: acctest.Required, Create: `TAG`},
		"targets":        acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_tag_namespace.tag-namespace1.name}.CostCenter.test`}},
	}

	BudgetResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: budget/default
func TestBudgetBudgetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBudgetBudgetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_budget_budget.test_budget"
	datasourceName := "data.oci_budget_budgets.test_budgets"
	singularDatasourceName := "data.oci_budget_budget.test_budget"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BudgetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Create, budgetRepresentationWithTargetTypeAsCompartmentAndTargets), "budget", "budget", t)

	acctest.ResourceTest(t, testAccCheckBudgetBudgetDestroy, []resource.TestStep{
		// verify Create for TargetType = Compartment
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Required, acctest.Create, budgetRepresentationWithTargetTypeAsCompartmentAndTargets),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "amount", "100"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies,
		},
		// verify Create with optionals for TargetType = Compartment
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Create, budgetRepresentationWithTargetTypeAsCompartmentAndTargets),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_rule_count"),
				resource.TestCheckResourceAttr(resourceName, "amount", "100"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters for TargetType = Compartment
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Update, budgetRepresentationWithTargetTypeAsCompartmentAndTargets),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_rule_count"),
				resource.TestCheckResourceAttr(resourceName, "amount", "200"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
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
		// verify Create for TargetType = Tag
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Required, acctest.Create, budgetRepresentationWithTargetTypeAsTagAndTargets),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "amount", "100"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies,
		},
		// verify Create with optionals for TargetType = Tag
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Create, budgetRepresentationWithTargetTypeAsTagAndTargets),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_rule_count"),
				resource.TestCheckResourceAttr(resourceName, "amount", "100"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TAG"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters for TargetType = Tag
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Update, budgetRepresentationWithTargetTypeAsTagAndTargets),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_rule_count"),
				resource.TestCheckResourceAttr(resourceName, "amount", "200"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TAG"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
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

		// verify Create
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Required, acctest.Create, budgetRepresentationWithTargetCompartmentId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "amount", "100"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Create, budgetRepresentationWithTargetCompartmentId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_rule_count"),
				resource.TestCheckResourceAttr(resourceName, "amount", "100"),
				resource.TestCheckResourceAttr(resourceName, "budget_processing_period_start_offset", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Update, budgetRepresentationWithTargetCompartmentId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_rule_count"),
				resource.TestCheckResourceAttr(resourceName, "amount", "200"),
				resource.TestCheckResourceAttr(resourceName, "budget_processing_period_start_offset", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "reset_period", "MONTHLY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_budgets", "test_budgets", acctest.Optional, acctest.Update, budgetDataSourceRepresentation) +
				compartmentIdVariableStr + BudgetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Optional, acctest.Update, budgetRepresentationWithTargetCompartmentId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "target_type", "COMPARTMENT"),

				resource.TestCheckResourceAttr(datasourceName, "budgets.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.actual_spend"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.alert_rule_count"),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.amount", "200"),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.budget_processing_period_start_offset", "11"),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.forecasted_spend"),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.reset_period", "MONTHLY"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.target_compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.target_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(datasourceName, "budgets.0.targets.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "budgets.0.version"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_budget", "test_budget", acctest.Required, acctest.Create, budgetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BudgetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "budget_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_rule_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "amount", "200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "budget_processing_period_start_offset", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "reset_period", "MONTHLY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_type", "COMPARTMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BudgetResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				// Need this workaround due to import behavior change introduced by https://github.com/hashicorp/terraform/issues/20985
				"actual_spend",
				"forecasted_spend",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckBudgetBudgetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BudgetClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_budget_budget" {
			noResourceFound = false
			request := oci_budget.GetBudgetRequest{}

			tmp := rs.Primary.ID
			request.BudgetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "budget")

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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("BudgetBudget") {
		resource.AddTestSweepers("BudgetBudget", &resource.Sweeper{
			Name:         "BudgetBudget",
			Dependencies: acctest.DependencyGraph["budget"],
			F:            sweepBudgetBudgetResource,
		})
	}
}

func sweepBudgetBudgetResource(compartment string) error {
	budgetClient := acctest.GetTestClients(&schema.ResourceData{}).BudgetClient()
	// BudgetBudgetResource can only run on root compartment
	compartment = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	budgetIds, err := getBudgetIds(compartment)
	if err != nil {
		return err
	}
	for _, budgetId := range budgetIds {
		if ok := acctest.SweeperDefaultResourceId[budgetId]; !ok {
			deleteBudgetRequest := oci_budget.DeleteBudgetRequest{}

			deleteBudgetRequest.BudgetId = &budgetId

			deleteBudgetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "budget")
			_, error := budgetClient.DeleteBudget(context.Background(), deleteBudgetRequest)
			if error != nil {
				fmt.Printf("Error deleting Budget %s %s, It is possible that the resource is already deleted. Please verify manually \n", budgetId, error)
				continue
			}
		}
	}
	return nil
}

func getBudgetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BudgetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	budgetClient := acctest.GetTestClients(&schema.ResourceData{}).BudgetClient()

	listBudgetsRequest := oci_budget.ListBudgetsRequest{}
	listBudgetsRequest.CompartmentId = &compartmentId
	listBudgetsResponse, err := budgetClient.ListBudgets(context.Background(), listBudgetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Budget list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, budget := range listBudgetsResponse.Items {
		id := *budget.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BudgetId", id)
	}
	return resourceIds, nil
}
