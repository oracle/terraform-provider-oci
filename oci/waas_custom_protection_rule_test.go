// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_waas "github.com/oracle/oci-go-sdk/waas"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CustomProtectionRuleRequiredOnlyResource = CustomProtectionRuleResourceDependencies +
		generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Required, Create, customProtectionRuleRepresentation)

	CustomProtectionRuleRequiredResourceWithoutDependencies = generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Required, Create, customProtectionRuleRepresentation) +
		generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule2", Optional, Update, customProtectionRuleRepresentation)

	CustomProtectionRuleResourceConfig = CustomProtectionRuleResourceDependencies +
		generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Update, customProtectionRuleRepresentation)

	customProtectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"custom_protection_rule_id": Representation{repType: Required, create: `${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`},
	}

	customProtectionRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        Representation{repType: Required, create: `${var.compartment_id}`},
		"display_names":                         Representation{repType: Optional, create: []string{`displayName`, `displayName2`}, update: []string{`displayName2`, `displayName3`}},
		"ids":                                   Representation{repType: Optional, create: []string{`${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`}},
		"states":                                Representation{repType: Optional, create: []string{`ACTIVE`}},
		"time_created_greater_than_or_equal_to": Representation{repType: Optional, create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                Representation{repType: Optional, create: `2038-01-01T00:00:00.000Z`},
		"filter":                                RepresentationGroup{Required, customProtectionRuleDataSourceFilterRepresentation}}
	customProtectionRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`}},
	}

	template1 = `SecRule REQUEST_URI / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%{TX.0}   found within %%{MATCHED_VAR_NAME}: %%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\"`
	template2 = `SecRule REQUEST_COOKIES / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%{TX.0}   found within %%{MATCHED_VAR_NAME}: %%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\"`

	customProtectionRuleRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"template":       Representation{repType: Required, create: template1, update: template2},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	CustomProtectionRuleResourceDependencies = DefinedTagsDependencies
)

func TestWaasCustomProtectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasCustomProtectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waas_custom_protection_rule.test_custom_protection_rule"
	datasourceName := "data.oci_waas_custom_protection_rules.test_custom_protection_rules"
	singularDatasourceName := "data.oci_waas_custom_protection_rule.test_custom_protection_rule"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckWaasCustomProtectionRuleDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CustomProtectionRuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Required, Create, customProtectionRuleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "template"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CustomProtectionRuleResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CustomProtectionRuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Create, customProtectionRuleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "template"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CustomProtectionRuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Create,
						representationCopyWithNewProperties(customProtectionRuleRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "template"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Update, customProtectionRuleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "template"),

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
					generateDataSourceFromRepresentationMap("oci_waas_custom_protection_rules", "test_custom_protection_rules", Optional, Update, customProtectionRuleDataSourceRepresentation) +
					compartmentIdVariableStr + CustomProtectionRuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Optional, Update, customProtectionRuleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_waas_custom_protection_rule", "test_custom_protection_rule", Required, Create, customProtectionRuleSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CustomProtectionRuleResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}

func testAccCheckWaasCustomProtectionRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).waasClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waas_custom_protection_rule" {
			noResourceFound = false
			request := oci_waas.GetCustomProtectionRuleRequest{}

			tmp := rs.Primary.ID
			request.CustomProtectionRuleId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "waas")

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
	if !inSweeperExcludeList("WaasCustomProtectionRule") {
		resource.AddTestSweepers("WaasCustomProtectionRule", &resource.Sweeper{
			Name:         "WaasCustomProtectionRule",
			Dependencies: DependencyGraph["customProtectionRule"],
			F:            sweepWaasCustomProtectionRuleResource,
		})
	}
}

func sweepWaasCustomProtectionRuleResource(compartment string) error {
	waasClient := GetTestClients(&schema.ResourceData{}).waasClient
	customProtectionRuleIds, err := getCustomProtectionRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, customProtectionRuleId := range customProtectionRuleIds {
		if ok := SweeperDefaultResourceId[customProtectionRuleId]; !ok {
			deleteCustomProtectionRuleRequest := oci_waas.DeleteCustomProtectionRuleRequest{}

			deleteCustomProtectionRuleRequest.CustomProtectionRuleId = &customProtectionRuleId

			deleteCustomProtectionRuleRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "waas")
			_, error := waasClient.DeleteCustomProtectionRule(context.Background(), deleteCustomProtectionRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting CustomProtectionRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", customProtectionRuleId, error)
				continue
			}
			waitTillCondition(testAccProvider, &customProtectionRuleId, customProtectionRuleSweepWaitCondition, time.Duration(3*time.Minute),
				customProtectionRuleSweepResponseFetchOperation, "waas", true)
		}
	}
	return nil
}

func getCustomProtectionRuleIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CustomProtectionRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	waasClient := GetTestClients(&schema.ResourceData{}).waasClient

	listCustomProtectionRulesRequest := oci_waas.ListCustomProtectionRulesRequest{}
	listCustomProtectionRulesRequest.CompartmentId = &compartmentId
	listCustomProtectionRulesRequest.LifecycleState = []oci_waas.ListCustomProtectionRulesLifecycleStateEnum{oci_waas.ListCustomProtectionRulesLifecycleStateActive}
	listCustomProtectionRulesResponse, err := waasClient.ListCustomProtectionRules(context.Background(), listCustomProtectionRulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CustomProtectionRule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, customProtectionRule := range listCustomProtectionRulesResponse.Items {
		id := *customProtectionRule.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "CustomProtectionRuleId", id)
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
	_, err := client.waasClient.GetCustomProtectionRule(context.Background(), oci_waas.GetCustomProtectionRuleRequest{
		CustomProtectionRuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
