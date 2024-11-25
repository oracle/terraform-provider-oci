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
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementCompliancePolicyRuleRequiredOnlyResource = FleetAppsManagementCompliancePolicyRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rule", "test_compliance_policy_rule", acctest.Required, acctest.Create, FleetAppsManagementCompliancePolicyRuleRepresentation)

	FleetAppsManagementCompliancePolicyRuleResourceConfig = FleetAppsManagementCompliancePolicyRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rule", "test_compliance_policy_rule", acctest.Optional, acctest.Update, FleetAppsManagementCompliancePolicyRuleRepresentation)

	FleetAppsManagementCompliancePolicyRuleSingularDataSourceRepresentation = map[string]interface{}{
		"compliance_policy_rule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_compliance_policy_rule.test_compliance_policy_rule.id}`},
	}

	FleetAppsManagementCompliancePolicyRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compliance_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compliance_policy_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_compliance_policy_rule.test_compliance_policy_rule.id}`},
		"patch_name":           acctest.Representation{RepType: acctest.Optional, Create: `BUG_FIX`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementCompliancePolicyRuleDataSourceFilterRepresentation}}
	FleetAppsManagementCompliancePolicyRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_compliance_policy_rule.test_compliance_policy_rule.id}`}},
	}

	FleetAppsManagementCompliancePolicyRuleRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"patch_selection":      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementCompliancePolicyRulePatchSelectionRepresentation},
		"patch_type":           acctest.Representation{RepType: acctest.Required, Create: []string{`BUG`}, Update: []string{`Security`}},
		"product_version":      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementCompliancePolicyRuleProductVersionRepresentation},
		"compliance_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compliance_policy_id}`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"grace_period":         acctest.Representation{RepType: acctest.Optional, Create: `gracePeriod`, Update: `gracePeriod2`},
		"severity":             acctest.Representation{RepType: acctest.Optional, Create: []string{`MEDIUM`}, Update: []string{`LOW`}},
	}
	FleetAppsManagementCompliancePolicyRulePatchSelectionRepresentation = map[string]interface{}{
		"selection_type":     acctest.Representation{RepType: acctest.Required, Create: `PATCH_LEVEL`},
		"days_since_release": acctest.Representation{RepType: acctest.Optional, Create: `0`},
		"patch_level":        acctest.Representation{RepType: acctest.Required, Create: `LATEST`},
	}
	FleetAppsManagementCompliancePolicyRuleProductVersionRepresentation = map[string]interface{}{
		"version":                               acctest.Representation{RepType: acctest.Required, Create: `8`, Update: `9`},
		"is_applicable_for_all_higher_versions": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	FleetAppsManagementCompliancePolicyRuleResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_compliance_policies", "test_compliance_policies", acctest.Required, acctest.Create, FleetAppsManagementCompliancePolicyDataSourceRepresentation)
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementCompliancePolicyRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementCompliancePolicyRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compliancePolicyId := utils.GetEnvSettingWithBlankDefault("compliance_policy_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compliancePolicyIdVariableStr := fmt.Sprintf("variable \"compliance_policy_id\" { default = \"%s\" }\n", compliancePolicyId)

	resourceName := "oci_fleet_apps_management_compliance_policy_rule.test_compliance_policy_rule"
	datasourceName := "data.oci_fleet_apps_management_compliance_policy_rules.test_compliance_policy_rules"
	singularDatasourceName := "data.oci_fleet_apps_management_compliance_policy_rule.test_compliance_policy_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compliancePolicyIdVariableStr+FleetAppsManagementCompliancePolicyRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rule", "test_compliance_policy_rule", acctest.Optional, acctest.Create, FleetAppsManagementCompliancePolicyRuleRepresentation), "fleetappsmanagement", "compliancePolicyRule", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementCompliancePolicyRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + compliancePolicyIdVariableStr + FleetAppsManagementCompliancePolicyRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rule", "test_compliance_policy_rule", acctest.Required, acctest.Create, FleetAppsManagementCompliancePolicyRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "patch_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_selection.0.days_since_release"),
				resource.TestCheckResourceAttr(resourceName, "patch_selection.0.patch_level", "LATEST"),
				resource.TestCheckResourceAttr(resourceName, "patch_selection.0.selection_type", "PATCH_LEVEL"),
				resource.TestCheckResourceAttr(resourceName, "patch_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "product_version.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "product_version.0.version", "8"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + compliancePolicyIdVariableStr + FleetAppsManagementCompliancePolicyRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + compliancePolicyIdVariableStr + FleetAppsManagementCompliancePolicyRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rule", "test_compliance_policy_rule", acctest.Optional, acctest.Create, FleetAppsManagementCompliancePolicyRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compliance_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "grace_period", "gracePeriod"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "patch_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_selection.0.days_since_release"),
				resource.TestCheckResourceAttr(resourceName, "patch_selection.0.patch_level", "LATEST"),
				resource.TestCheckResourceAttr(resourceName, "patch_selection.0.selection_type", "PATCH_LEVEL"),
				resource.TestCheckResourceAttr(resourceName, "patch_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "product_version.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "product_version.0.is_applicable_for_all_higher_versions", "false"),
				resource.TestCheckResourceAttr(resourceName, "product_version.0.version", "8"),
				resource.TestCheckResourceAttr(resourceName, "severity.#", "1"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + compliancePolicyIdVariableStr + FleetAppsManagementCompliancePolicyRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rule", "test_compliance_policy_rule", acctest.Optional, acctest.Update, FleetAppsManagementCompliancePolicyRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compliance_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "grace_period", "gracePeriod2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "patch_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_selection.0.days_since_release"),
				resource.TestCheckResourceAttr(resourceName, "patch_selection.0.selection_type", "PATCH_LEVEL"),
				resource.TestCheckResourceAttr(resourceName, "patch_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "product_version.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "product_version.0.is_applicable_for_all_higher_versions"),
				resource.TestCheckResourceAttr(resourceName, "product_version.0.version", "9"),
				resource.TestCheckResourceAttr(resourceName, "severity.#", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rules", "test_compliance_policy_rules", acctest.Required, acctest.Create, FleetAppsManagementCompliancePolicyRuleDataSourceRepresentation) +
				compartmentIdVariableStr + compliancePolicyIdVariableStr + FleetAppsManagementCompliancePolicyRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rule", "test_compliance_policy_rule", acctest.Optional, acctest.Update, FleetAppsManagementCompliancePolicyRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "compliance_policy_rule_collection.0.items.0.compliance_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "compliance_policy_rule_collection.0.items.0.display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "compliance_policy_rule_collection.0.items.0.state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "compliance_policy_rule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compliance_policy_rule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_compliance_policy_rule", "test_compliance_policy_rule", acctest.Required, acctest.Create, FleetAppsManagementCompliancePolicyRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + compliancePolicyIdVariableStr + FleetAppsManagementCompliancePolicyRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compliance_policy_rule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "grace_period", "gracePeriod2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_selection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_selection.0.days_since_release"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_selection.0.patch_level", "LATEST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_selection.0.selection_type", "PATCH_LEVEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_type.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product_version.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product_version.0.is_applicable_for_all_higher_versions", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product_version.0.version", "9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "severity.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementCompliancePolicyRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementCompliancePolicyRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_compliance_policy_rule" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetCompliancePolicyRuleRequest{}

			tmp := rs.Primary.ID
			request.CompliancePolicyRuleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetCompliancePolicyRule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementCompliancePolicyRule") {
		resource.AddTestSweepers("FleetAppsManagementCompliancePolicyRule", &resource.Sweeper{
			Name:         "FleetAppsManagementCompliancePolicyRule",
			Dependencies: acctest.DependencyGraph["compliancePolicyRule"],
			F:            sweepFleetAppsManagementCompliancePolicyRuleResource,
		})
	}
}

func sweepFleetAppsManagementCompliancePolicyRuleResource(compartment string) error {
	fleetAppsManagementAdminClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementAdminClient()
	compliancePolicyRuleIds, err := getFleetAppsManagementCompliancePolicyRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, compliancePolicyRuleId := range compliancePolicyRuleIds {
		if ok := acctest.SweeperDefaultResourceId[compliancePolicyRuleId]; !ok {
			deleteCompliancePolicyRuleRequest := oci_fleet_apps_management.DeleteCompliancePolicyRuleRequest{}

			deleteCompliancePolicyRuleRequest.CompliancePolicyRuleId = &compliancePolicyRuleId

			deleteCompliancePolicyRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementAdminClient.DeleteCompliancePolicyRule(context.Background(), deleteCompliancePolicyRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting CompliancePolicyRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", compliancePolicyRuleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &compliancePolicyRuleId, FleetAppsManagementCompliancePolicyRuleSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementCompliancePolicyRuleSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementCompliancePolicyRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CompliancePolicyRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementAdminClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementAdminClient()

	listCompliancePolicyRulesRequest := oci_fleet_apps_management.ListCompliancePolicyRulesRequest{}
	listCompliancePolicyRulesRequest.CompartmentId = &compartmentId
	listCompliancePolicyRulesRequest.LifecycleState = oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateActive
	listCompliancePolicyRulesResponse, err := fleetAppsManagementAdminClient.ListCompliancePolicyRules(context.Background(), listCompliancePolicyRulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CompliancePolicyRule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, compliancePolicyRule := range listCompliancePolicyRulesResponse.Items {
		id := *compliancePolicyRule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CompliancePolicyRuleId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementCompliancePolicyRuleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if compliancePolicyRuleResponse, ok := response.Response.(oci_fleet_apps_management.GetCompliancePolicyRuleResponse); ok {
		return compliancePolicyRuleResponse.LifecycleState != oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementCompliancePolicyRuleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementAdminClient().GetCompliancePolicyRule(context.Background(), oci_fleet_apps_management.GetCompliancePolicyRuleRequest{
		CompliancePolicyRuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
