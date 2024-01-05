// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_waas "github.com/oracle/oci-go-sdk/v65/waas"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	WaasCustomProtectionRuleRequiredOnlyResource = WaasCustomProtectionRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Required, acctest.Create, WaasCustomProtectionRuleRepresentation)

	CustomProtectionRuleRequiredResourceWithoutDependencies = acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Required, acctest.Create, WaasCustomProtectionRuleRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule2", acctest.Optional, acctest.Update, WaasCustomProtectionRuleRepresentation)

	WaasCustomProtectionRuleResourceConfig = WaasCustomProtectionRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Optional, acctest.Update, WaasCustomProtectionRuleRepresentation)

	WaasWaasCustomProtectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"custom_protection_rule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`},
	}

	WaasWaasCustomProtectionRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_names":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`, `displayName2`}, Update: []string{`displayName2`, `displayName3`}},
		"ids":                                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`}},
		"states":                                acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: WaasCustomProtectionRuleDataSourceFilterRepresentation}}
	WaasCustomProtectionRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`}},
	}

	template1 = `SecRule REQUEST_URI / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%{TX.0}   found within %%{MATCHED_VAR_NAME}: %%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\"`
	template2 = `SecRule REQUEST_COOKIES / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%{TX.0}   found within %%{MATCHED_VAR_NAME}: %%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\"`

	WaasCustomProtectionRuleRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"template":       acctest.Representation{RepType: acctest.Required, Create: template1, Update: template2},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	WaasCustomProtectionRuleResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: waas/default
func TestWaasCustomProtectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasCustomProtectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waas_custom_protection_rule.test_custom_protection_rule"
	datasourceName := "data.oci_waas_custom_protection_rules.test_custom_protection_rules"
	singularDatasourceName := "data.oci_waas_custom_protection_rule.test_custom_protection_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+WaasCustomProtectionRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Optional, acctest.Create, WaasCustomProtectionRuleRepresentation), "waas", "customProtectionRule", t)

	acctest.ResourceTest(t, testAccCheckWaasCustomProtectionRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WaasCustomProtectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Required, acctest.Create, WaasCustomProtectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "template"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WaasCustomProtectionRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + WaasCustomProtectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Optional, acctest.Create, WaasCustomProtectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "template"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WaasCustomProtectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(WaasCustomProtectionRuleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "template"),

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
			Config: config + compartmentIdVariableStr + WaasCustomProtectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Optional, acctest.Update, WaasCustomProtectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "template"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_custom_protection_rules", "test_custom_protection_rules", acctest.Optional, acctest.Update, WaasWaasCustomProtectionRuleDataSourceRepresentation) +
				compartmentIdVariableStr + WaasCustomProtectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Optional, acctest.Update, WaasCustomProtectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_names.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "states.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "custom_protection_rules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "custom_protection_rules.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "custom_protection_rules.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "custom_protection_rules.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "custom_protection_rules.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "custom_protection_rules.0.mod_security_rule_ids.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "custom_protection_rules.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "custom_protection_rules.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", acctest.Required, acctest.Create, WaasWaasCustomProtectionRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + WaasCustomProtectionRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_protection_rule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mod_security_rule_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "template"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + WaasCustomProtectionRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckWaasCustomProtectionRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).WaasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waas_custom_protection_rule" {
			noResourceFound = false
			request := oci_waas.GetCustomProtectionRuleRequest{}

			tmp := rs.Primary.ID
			request.CustomProtectionRuleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waas")

			response, err := client.GetCustomProtectionRule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waas.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("WaasCustomProtectionRule") {
		resource.AddTestSweepers("WaasCustomProtectionRule", &resource.Sweeper{
			Name:         "WaasCustomProtectionRule",
			Dependencies: acctest.DependencyGraph["customProtectionRule"],
			F:            sweepWaasCustomProtectionRuleResource,
		})
	}
}

func sweepWaasCustomProtectionRuleResource(compartment string) error {
	waasClient := acctest.GetTestClients(&schema.ResourceData{}).WaasClient()
	customProtectionRuleIds, err := getWaasCustomProtectionRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, customProtectionRuleId := range customProtectionRuleIds {
		if ok := acctest.SweeperDefaultResourceId[customProtectionRuleId]; !ok {
			deleteCustomProtectionRuleRequest := oci_waas.DeleteCustomProtectionRuleRequest{}

			deleteCustomProtectionRuleRequest.CustomProtectionRuleId = &customProtectionRuleId

			deleteCustomProtectionRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waas")
			_, error := waasClient.DeleteCustomProtectionRule(context.Background(), deleteCustomProtectionRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting CustomProtectionRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", customProtectionRuleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &customProtectionRuleId, WaasCustomProtectionRuleSweepWaitCondition, time.Duration(3*time.Minute),
				WaasCustomProtectionRuleSweepResponseFetchOperation, "waas", true)
		}
	}
	return nil
}

func getWaasCustomProtectionRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CustomProtectionRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	waasClient := acctest.GetTestClients(&schema.ResourceData{}).WaasClient()

	listCustomProtectionRulesRequest := oci_waas.ListCustomProtectionRulesRequest{}
	listCustomProtectionRulesRequest.CompartmentId = &compartmentId
	listCustomProtectionRulesRequest.LifecycleState = []oci_waas.LifecycleStatesEnum{oci_waas.LifecycleStatesActive}
	listCustomProtectionRulesResponse, err := waasClient.ListCustomProtectionRules(context.Background(), listCustomProtectionRulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CustomProtectionRule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, customProtectionRule := range listCustomProtectionRulesResponse.Items {
		id := *customProtectionRule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CustomProtectionRuleId", id)
	}
	return resourceIds, nil
}

func WaasCustomProtectionRuleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if customProtectionRuleResponse, ok := response.Response.(oci_waas.GetCustomProtectionRuleResponse); ok {
		return customProtectionRuleResponse.LifecycleState != oci_waas.LifecycleStatesDeleted
	}
	return false
}

func WaasCustomProtectionRuleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.WaasClient().GetCustomProtectionRule(context.Background(), oci_waas.GetCustomProtectionRuleRequest{
		CustomProtectionRuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
