// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var (
	DataSafeAlertPolicyRuleRequiredOnlyResource = DataSafeAlertPolicyRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy_rule", "test_alert_policy_rule", acctest.Required, acctest.Create, DataSafeAlertPolicyRuleRepresentation)

	DataSafeAlertPolicyRuleResourceConfig = DataSafeAlertPolicyRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy_rule", "test_alert_policy_rule", acctest.Required, acctest.Create, DataSafeAlertPolicyRuleRepresentation)

	DataSafeAlertPolicyRuleDataSourceRepresentation = map[string]interface{}{
		"alert_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_alert_policy.test_alert_policy.id}`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeAlertPolicyRuleDataSourceFilterRepresentation},
	}

	DataSafeAlertPolicyRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_alert_policy_rule.test_alert_policy_rule.key}`}},
	}

	DataSafeAlertPolicyRuleRepresentation = map[string]interface{}{
		"alert_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_alert_policy.test_alert_policy.id}`},
		"expression":      acctest.Representation{RepType: acctest.Required, Create: `operation eq \"abc\"`, Update: `operation eq \"abcd\"`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `Check if remote login password file is exclusive`, Update: `description2`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	DataSafeAlertPolicyRuleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Required, acctest.Create, DataSafeAlertPolicyRepresentation)

	DataSafeAlertPolicyRuleSingularDataSourceRepresentation = map[string]interface{}{
		"alert_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_alert_policy.test_alert_policy.id}`},
		"rule_key":        acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_alert_policy_rule.test_alert_policy_rule.key}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeAlertPolicyRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAlertPolicyRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_safe_alert_policy_rule.test_alert_policy_rule"
	datasourceName := "data.oci_data_safe_alert_policy_rules.test_alert_policy_rules"
	singularDatasourceName := "data.oci_data_safe_alert_policy_rule.test_alert_policy_rule"
	policyId := utils.GetEnvSettingWithBlankDefault("data_safe_alert_policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"policy_id\" { default = \"%s\" }\n", policyId)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeAlertPolicyRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy_rule", "test_alert_policy_rule", acctest.Optional, acctest.Create, DataSafeAlertPolicyRuleRepresentation), "datasafe", "alertPolicyRule", t)

	acctest.ResourceTest(t, testAccCheckDataSafeAlertPolicyRuleDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeAlertPolicyRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy_rule", "test_alert_policy_rule", acctest.Required, acctest.Create, DataSafeAlertPolicyRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "expression", "operation eq \"abc\""),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeAlertPolicyRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeAlertPolicyRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy_rule", "test_alert_policy_rule", acctest.Optional, acctest.Create, DataSafeAlertPolicyRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "Check if remote login password file is exclusive"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "expression", "operation eq \"abc\""),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "key")
					var compositeId string
					compositeId, err = acctest.FromInstanceState(s, resourceName, "id")
					prefix := "oci_data_safe_alert_policy_rule:"
					fullPath := prefix + compositeId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&fullPath, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataSafeAlertPolicyRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy_rule", "test_alert_policy_rule", acctest.Optional, acctest.Update, DataSafeAlertPolicyRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "expression", "operation eq \"abcd\""),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "key")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policy_rules", "test_alert_policy_rules", acctest.Required, acctest.Create, DataSafeAlertPolicyRuleDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeAlertPolicyRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_rule_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_rule_collection.0.items.0.expression"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policy_rule", "test_alert_policy_rule", acctest.Required, acctest.Create, DataSafeAlertPolicyRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeAlertPolicyRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expression", "operation eq \"abc\""),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + policyIdVariableStr + DataSafeAlertPolicyRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeAlertPolicyRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_alert_policy_rule" {
			noResourceFound = false
			request := oci_data_safe.GetAlertPolicyRuleRequest{}

			if value, ok := rs.Primary.Attributes["alert_policy_id"]; ok {
				request.AlertPolicyId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.RuleKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetAlertPolicyRule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeAlertPolicyRule") {
		resource.AddTestSweepers("DataSafeAlertPolicyRule", &resource.Sweeper{
			Name:         "DataSafeAlertPolicyRule",
			Dependencies: acctest.DependencyGraph["alertPolicyRule"],
			F:            sweepDataSafeAlertPolicyRuleResource,
		})
	}
}

func sweepDataSafeAlertPolicyRuleResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	alertPolicyRuleIds, err := getDataSafeAlertPolicyRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, alertPolicyRuleId := range alertPolicyRuleIds {
		if ok := acctest.SweeperDefaultResourceId[alertPolicyRuleId]; !ok {
			deleteAlertPolicyRuleRequest := oci_data_safe.DeleteAlertPolicyRuleRequest{}

			deleteAlertPolicyRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteAlertPolicyRule(context.Background(), deleteAlertPolicyRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting AlertPolicyRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", alertPolicyRuleId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeAlertPolicyRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AlertPolicyRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listAlertPolicyRulesRequest := oci_data_safe.ListAlertPolicyRulesRequest{}

	alertPolicyIds, error := getDataSafeAlertPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting alertPolicyId required for AlertPolicyRule resource requests \n")
	}
	for _, alertPolicyId := range alertPolicyIds {
		listAlertPolicyRulesRequest.AlertPolicyId = &alertPolicyId

		listAlertPolicyRulesResponse, err := dataSafeClient.ListAlertPolicyRules(context.Background(), listAlertPolicyRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting AlertPolicyRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, alertPolicyRule := range listAlertPolicyRulesResponse.Items {
			id := *alertPolicyRule.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AlertPolicyRuleId", id)
		}

	}
	return resourceIds, nil
}
