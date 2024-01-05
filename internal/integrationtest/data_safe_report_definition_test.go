// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeReportDefinitionRequiredOnlyResource = DataSafeReportDefinitionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Required, acctest.Create, reportDefinitionRepresentation)

	DataSafeReportDefinitionResourceConfig = DataSafeReportDefinitionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Optional, acctest.Update, reportDefinitionRepresentation)

	DataSafereportDefinitionSingularDataSourceRepresentation = map[string]interface{}{
		"report_definition_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_report_definition.test_report_definition.id}`},
	}

	DataSafereportDefinitionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	reportDefinitionRepresentation = map[string]interface{}{
		"column_filters":  acctest.RepresentationGroup{RepType: acctest.Required, Group: reportDefinitionColumnFiltersRepresentation},
		"column_info":     acctest.RepresentationGroup{RepType: acctest.Required, Group: reportDefinitionColumnInfoRepresentation},
		"column_sortings": acctest.RepresentationGroup{RepType: acctest.Required, Group: reportDefinitionColumnSortingsRepresentation},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName18`, Update: `displayName19`},
		"parent_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.report_ocid}`},
		"summary":         acctest.RepresentationGroup{RepType: acctest.Required, Group: reportDefinitionSummaryRepresentation},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreReportDefinitionSystemTagsChangesRep},
	}
	reportDefinitionColumnFiltersRepresentation = map[string]interface{}{
		"expressions": acctest.Representation{RepType: acctest.Required, Create: []string{`expressions`}, Update: []string{`expressions2`}},
		"field_name":  acctest.Representation{RepType: acctest.Required, Create: `operation`, Update: `operation`},
		"is_enabled":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_hidden":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"operator":    acctest.Representation{RepType: acctest.Required, Create: `IN`, Update: `EQ`},
	}
	reportDefinitionColumnInfoRepresentation = map[string]interface{}{
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `Target Id`, Update: `Target Id`},
		"display_order": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"field_name":    acctest.Representation{RepType: acctest.Required, Create: `targetId`, Update: `targetId`},
		"is_hidden":     acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"data_type":     acctest.Representation{RepType: acctest.Optional, Create: `String`, Update: `String`},
	}
	reportDefinitionColumnSortingsRepresentation = map[string]interface{}{
		"field_name":    acctest.Representation{RepType: acctest.Required, Create: `operation`, Update: `operation`},
		"is_ascending":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"sorting_order": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	reportDefinitionSummaryRepresentation = map[string]interface{}{
		"display_order":       acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name`},
		"count_of":            acctest.Representation{RepType: acctest.Required, Create: `creates`, Update: `creates`},
		"group_by_field_name": acctest.Representation{RepType: acctest.Optional, Create: `operation`, Update: `operation`},
		"is_hidden":           acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"scim_filter":         acctest.Representation{RepType: acctest.Optional, Create: `scimFilter`, Update: `scimFilter2`},
	}
	ignoreReportDefinitionSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `compliance_standards`}},
	}

	DataSafeReportDefinitionResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeReportDefinitionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeReportDefinitionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	reportDefId := utils.GetEnvSettingWithBlankDefault("report_ocid")
	reportDefIdVariableStr := fmt.Sprintf("variable \"report_ocid\" { default = \"%s\" }\n", reportDefId)

	resourceName := "oci_data_safe_report_definition.test_report_definition"
	datasourceName := "data.oci_data_safe_report_definitions.test_report_definitions"
	singularDatasourceName := "data.oci_data_safe_report_definition.test_report_definition"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+reportDefIdVariableStr+DataSafeReportDefinitionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Optional, acctest.Create, reportDefinitionRepresentation), "datasafe", "reportDefinition", t)

	acctest.ResourceTest(t, testAccCheckDataSafeReportDefinitionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + reportDefIdVariableStr + DataSafeReportDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Required, acctest.Create, reportDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.expressions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.operator", "IN"),
				resource.TestCheckResourceAttr(resourceName, "column_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.display_name", "Target Id"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.display_order", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.field_name", "targetId"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.is_ascending", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.sorting_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName18"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_id"),
				resource.TestCheckResourceAttr(resourceName, "summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.display_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + reportDefIdVariableStr + DataSafeReportDefinitionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + reportDefIdVariableStr + DataSafeReportDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Optional, acctest.Create, reportDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.expressions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.operator", "IN"),
				resource.TestCheckResourceAttr(resourceName, "column_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.data_type", "String"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.display_name", "Target Id"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.display_order", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.field_name", "targetId"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.is_ascending", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.sorting_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName18"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.count_of", "creates"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.display_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.group_by_field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.scim_filter", "scimFilter"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + reportDefIdVariableStr + compartmentIdUVariableStr + DataSafeReportDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(reportDefinitionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.expressions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.operator", "IN"),
				resource.TestCheckResourceAttr(resourceName, "column_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.data_type", "String"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.display_name", "Target Id"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.display_order", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.field_name", "targetId"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.is_ascending", "false"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.sorting_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName18"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.count_of", "creates"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.display_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.group_by_field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.scim_filter", "scimFilter"),

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
			Config: config + compartmentIdVariableStr + reportDefIdVariableStr + DataSafeReportDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Optional, acctest.Update, reportDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.expressions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceName, "column_filters.0.operator", "EQ"),
				resource.TestCheckResourceAttr(resourceName, "column_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.data_type", "String"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.display_name", "Target Id"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.display_order", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.field_name", "targetId"),
				resource.TestCheckResourceAttr(resourceName, "column_info.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.is_ascending", "true"),
				resource.TestCheckResourceAttr(resourceName, "column_sortings.0.sorting_order", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName19"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.count_of", "creates"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.display_order", "11"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.group_by_field_name", "operation"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "summary.0.scim_filter", "scimFilter2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_report_definitions", "test_report_definitions", acctest.Optional, acctest.Update, DataSafereportDefinitionDataSourceRepresentation) +
				compartmentIdVariableStr + reportDefIdVariableStr + DataSafeReportDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Optional, acctest.Update, reportDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "report_definition_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Required, acctest.Create, DataSafereportDefinitionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + reportDefIdVariableStr + DataSafeReportDefinitionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "report_definition_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "column_filters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_filters.0.expressions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_filters.0.field_name", "operation"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_filters.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_filters.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_filters.0.operator", "EQ"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_info.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_info.0.data_type", "String"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_info.0.display_name", "Target Id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_info.0.display_order", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_info.0.field_name", "targetId"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_info.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_sortings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_sortings.0.field_name", "operation"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_sortings.0.is_ascending", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "column_sortings.0.sorting_order", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_order"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_seeded"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.0.count_of", "creates"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.0.display_order", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.0.group_by_field_name", "operation"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.0.scim_filter", "scimFilter2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + reportDefIdVariableStr + DataSafeReportDefinitionResourceConfig,
		},
		// verify resource import
		{
			Config:                  config + reportDefIdVariableStr + DataSafeReportDefinitionResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeReportDefinitionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_report_definition" {
			noResourceFound = false
			request := oci_data_safe.GetReportDefinitionRequest{}

			tmp := rs.Primary.ID
			request.ReportDefinitionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetReportDefinition(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.ReportDefinitionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeReportDefinition") {
		resource.AddTestSweepers("DataSafeReportDefinition", &resource.Sweeper{
			Name:         "DataSafeReportDefinition",
			Dependencies: acctest.DependencyGraph["reportDefinition"],
			F:            sweepDataSafeReportDefinitionResource,
		})
	}
}

func sweepDataSafeReportDefinitionResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	reportDefinitionIds, err := getDataSafeReportDefinitionIds(compartment)
	if err != nil {
		return err
	}
	for _, reportDefinitionId := range reportDefinitionIds {
		if ok := acctest.SweeperDefaultResourceId[reportDefinitionId]; !ok {
			deleteReportDefinitionRequest := oci_data_safe.DeleteReportDefinitionRequest{}

			deleteReportDefinitionRequest.ReportDefinitionId = &reportDefinitionId

			deleteReportDefinitionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteReportDefinition(context.Background(), deleteReportDefinitionRequest)
			if error != nil {
				fmt.Printf("Error deleting ReportDefinition %s %s, It is possible that the resource is already deleted. Please verify manually \n", reportDefinitionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &reportDefinitionId, DataSafereportDefinitionsSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafereportDefinitionsSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeReportDefinitionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ReportDefinitionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listReportDefinitionsRequest := oci_data_safe.ListReportDefinitionsRequest{}
	listReportDefinitionsRequest.CompartmentId = &compartmentId
	listReportDefinitionsRequest.LifecycleState = oci_data_safe.ListReportDefinitionsLifecycleStateActive
	listReportDefinitionsResponse, err := dataSafeClient.ListReportDefinitions(context.Background(), listReportDefinitionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ReportDefinition list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, reportDefinition := range listReportDefinitionsResponse.Items {
		id := *reportDefinition.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ReportDefinitionId", id)
	}
	return resourceIds, nil
}

func DataSafereportDefinitionsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if reportDefinitionResponse, ok := response.Response.(oci_data_safe.GetReportDefinitionResponse); ok {
		return reportDefinitionResponse.LifecycleState != oci_data_safe.ReportDefinitionLifecycleStateDeleted
	}
	return false
}

func DataSafereportDefinitionsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetReportDefinition(context.Background(), oci_data_safe.GetReportDefinitionRequest{
		ReportDefinitionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
