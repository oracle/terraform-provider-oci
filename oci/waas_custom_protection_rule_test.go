// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v51/common"
	oci_waas "github.com/oracle/oci-go-sdk/v51/waas"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CustomProtectionRuleRequiredOnlyResource = CustomProtectionRuleResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Required, Create, customProtectionRuleRepresentation)

	CustomProtectionRuleRequiredResourceWithoutDependencies = GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Required, Create, customProtectionRuleRepresentation) +
		GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule2", Optional, Update, customProtectionRuleRepresentation)

	CustomProtectionRuleResourceConfig = CustomProtectionRuleResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Update, customProtectionRuleRepresentation)

	customProtectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"custom_protection_rule_id": Representation{RepType: Required, Create: `${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`},
	}

	customProtectionRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_names":                         Representation{RepType: Optional, Create: []string{`displayName`, `displayName2`}, Update: []string{`displayName2`, `displayName3`}},
		"ids":                                   Representation{RepType: Optional, Create: []string{`${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`}},
		"states":                                Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"time_created_greater_than_or_equal_to": Representation{RepType: Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                Representation{RepType: Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                RepresentationGroup{Required, customProtectionRuleDataSourceFilterRepresentation}}
	customProtectionRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`}},
	}

	template1 = `SecRule REQUEST_URI / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%{TX.0}   found within %%{MATCHED_VAR_NAME}: %%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\"`
	template2 = `SecRule REQUEST_COOKIES / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%{TX.0}   found within %%{MATCHED_VAR_NAME}: %%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\"`

	customProtectionRuleRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"template":       Representation{RepType: Required, Create: template1, Update: template2},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CustomProtectionRuleResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: waas/default
func TestWaasCustomProtectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasCustomProtectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waas_custom_protection_rule.test_custom_protection_rule"
	datasourceName := "data.oci_waas_custom_protection_rules.test_custom_protection_rules"
	singularDatasourceName := "data.oci_waas_custom_protection_rule.test_custom_protection_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+CustomProtectionRuleResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Create, customProtectionRuleRepresentation), "waas", "customProtectionRule", t)

	ResourceTest(t, testAccCheckWaasCustomProtectionRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CustomProtectionRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Required, Create, customProtectionRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "template"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CustomProtectionRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CustomProtectionRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Create, customProtectionRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "template"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CustomProtectionRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Create,
					RepresentationCopyWithNewProperties(customProtectionRuleRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "template"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CustomProtectionRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Update, customProtectionRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "template"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_waas_custom_protection_rules", "test_custom_protection_rules", Optional, Update, customProtectionRuleDataSourceRepresentation) +
				compartmentIdVariableStr + CustomProtectionRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Update, customProtectionRuleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_names.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "states.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "custom_protection_rules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "custom_protection_rules.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "custom_protection_rules.0.defined_tags.%", "1"),
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
				GenerateDataSourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Required, Create, customProtectionRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CustomProtectionRuleResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_protection_rule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + CustomProtectionRuleResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckWaasCustomProtectionRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).waasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waas_custom_protection_rule" {
			noResourceFound = false
			request := oci_waas.GetCustomProtectionRuleRequest{}

			tmp := rs.Primary.ID
			request.CustomProtectionRuleId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waas")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("WaasCustomProtectionRule") {
		resource.AddTestSweepers("WaasCustomProtectionRule", &resource.Sweeper{
			Name:         "WaasCustomProtectionRule",
			Dependencies: DependencyGraph["customProtectionRule"],
			F:            sweepWaasCustomProtectionRuleResource,
		})
	}
}

func sweepWaasCustomProtectionRuleResource(compartment string) error {
	waasClient := GetTestClients(&schema.ResourceData{}).waasClient()
	customProtectionRuleIds, err := getCustomProtectionRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, customProtectionRuleId := range customProtectionRuleIds {
		if ok := SweeperDefaultResourceId[customProtectionRuleId]; !ok {
			deleteCustomProtectionRuleRequest := oci_waas.DeleteCustomProtectionRuleRequest{}

			deleteCustomProtectionRuleRequest.CustomProtectionRuleId = &customProtectionRuleId

			deleteCustomProtectionRuleRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waas")
			_, error := waasClient.DeleteCustomProtectionRule(context.Background(), deleteCustomProtectionRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting CustomProtectionRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", customProtectionRuleId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &customProtectionRuleId, customProtectionRuleSweepWaitCondition, time.Duration(3*time.Minute),
				customProtectionRuleSweepResponseFetchOperation, "waas", true)
		}
	}
	return nil
}

func getCustomProtectionRuleIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "CustomProtectionRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	waasClient := GetTestClients(&schema.ResourceData{}).waasClient()

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
		AddResourceIdToSweeperResourceIdMap(compartmentId, "CustomProtectionRuleId", id)
	}
	return resourceIds, nil
}

func customProtectionRuleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if customProtectionRuleResponse, ok := response.Response.(oci_waas.GetCustomProtectionRuleResponse); ok {
		return customProtectionRuleResponse.LifecycleState != oci_waas.LifecycleStatesDeleted
	}
	return false
}

func customProtectionRuleSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.waasClient().GetCustomProtectionRule(context.Background(), oci_waas.GetCustomProtectionRuleRequest{
		CustomProtectionRuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
