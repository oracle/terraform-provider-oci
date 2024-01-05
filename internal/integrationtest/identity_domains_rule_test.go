// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsRuleRequiredOnlyResource = IdentityDomainsRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Required, acctest.Create, IdentityDomainsRuleRepresentation)

	IdentityDomainsRuleResourceConfig = IdentityDomainsRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Optional, acctest.Update, IdentityDomainsRuleRepresentation)

	IdentityDomainsRuleSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"rule_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_rule.test_rule.id}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsRuleDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"rule_count":     acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"rule_filter":    acctest.Representation{RepType: acctest.Optional, Create: `id eq \"${oci_identity_domains_rule.test_rule.id}\"`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsRuleRepresentation = map[string]interface{}{
		"condition":       acctest.Representation{RepType: acctest.Required, Create: `condition`, Update: `condition2`},
		"idcs_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"policy_type":     acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsRulePolicyTypeRepresentation},
		"return":          acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsRuleReturnRepresentation},
		"schemas":         acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:Rule`}},
		"active":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"attribute_sets":  acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"condition_group": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsRuleConditionGroupRepresentation},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"external_id":     acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"locked":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"rule_groovy":     acctest.Representation{RepType: acctest.Optional, Create: `ruleGroovy`, Update: `ruleGroovy2`},
		"tags":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsRuleTagsRepresentation},
	}
	IdentityDomainsRulePolicyTypeRepresentation = map[string]interface{}{
		// policy type value is immutable
		"value": acctest.Representation{RepType: acctest.Required, Create: `SignOn`},
	}
	IdentityDomainsRuleReturnRepresentation = map[string]interface{}{
		"name":          acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"value":         acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"return_groovy": acctest.Representation{RepType: acctest.Optional, Create: `returnGroovy`, Update: `returnGroovy2`},
	}
	IdentityDomainsRuleConditionGroupRepresentation = map[string]interface{}{
		// The other supported type ConditionGroup is not available in Terraform yet
		"type":  acctest.Representation{RepType: acctest.Required, Create: `Condition`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_domains_condition.test_condition.id}`, Update: `${oci_identity_domains_condition.test_condition2.id}`},
	}
	IdentityDomainsRuleTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsRuleResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Required, acctest.Create, IdentityDomainsConditionRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition2", acctest.Optional, acctest.Create, IdentityDomainsConditionRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_rule.test_rule"
	datasourceName := "data.oci_identity_domains_rules.test_rules"
	singularDatasourceName := "data.oci_identity_domains_rule.test_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Optional, acctest.Create, IdentityDomainsRuleRepresentation), "identitydomains", "rule", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Required, acctest.Create, IdentityDomainsRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "condition", "condition"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.0.value", "SignOn"),
				resource.TestCheckResourceAttr(resourceName, "return.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "return.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "return.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Optional, acctest.Create, IdentityDomainsRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition", "condition"),
				resource.TestCheckResourceAttr(resourceName, "condition_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition_group.0.type", "Condition"),
				resource.TestCheckResourceAttrSet(resourceName, "condition_group.0.value"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "locked", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.0.value", "SignOn"),
				resource.TestCheckResourceAttr(resourceName, "return.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "return.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "return.0.return_groovy", "returnGroovy"),
				resource.TestCheckResourceAttr(resourceName, "return.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "rule_groovy", "ruleGroovy"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "rules", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Optional, acctest.Update, IdentityDomainsRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "true"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition", "condition2"),
				resource.TestCheckResourceAttr(resourceName, "condition_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition_group.0.type", "Condition"),
				resource.TestCheckResourceAttrSet(resourceName, "condition_group.0.value"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "locked", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_type.0.value", "SignOn"),
				resource.TestCheckResourceAttr(resourceName, "return.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "return.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "return.0.return_groovy", "returnGroovy2"),
				resource.TestCheckResourceAttr(resourceName, "return.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "rule_groovy", "ruleGroovy2"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_rules", "test_rules", acctest.Optional, acctest.Update, IdentityDomainsRuleDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Optional, acctest.Update, IdentityDomainsRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "rule_count", "10"),
				resource.TestCheckResourceAttrSet(datasourceName, "rule_filter"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.schemas.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.active", "true"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.condition", "condition2"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.condition_group.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.condition_group.0.type", "Condition"),
				resource.TestCheckResourceAttrSet(datasourceName, "rules.0.condition_group.0.value"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.description", "description2"),
				resource.TestCheckResourceAttrSet(datasourceName, "rules.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.locked", "true"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.policy_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.policy_type.0.value", "SignOn"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.return.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.return.0.name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.return.0.return_groovy", "returnGroovy2"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.return.0.value", "value2"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.rule_groovy", "ruleGroovy2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_rule", "test_rule", acctest.Required, acctest.Create, IdentityDomainsRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "rule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "active", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition", "condition2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition_group.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition_group.0.type", "Condition"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "condition_group.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locked", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_type.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_type.0.value", "SignOn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "return.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "return.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "return.0.return_groovy", "returnGroovy2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "return.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_groovy", "ruleGroovy2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsRuleRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_rule", "rules"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_rule" {
			noResourceFound = false
			request := oci_identity_domains.GetRuleRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			tmp := rs.Primary.ID
			request.RuleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetRule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsRule") {
		resource.AddTestSweepers("IdentityDomainsRule", &resource.Sweeper{
			Name:         "IdentityDomainsRule",
			Dependencies: acctest.DependencyGraph["rule"],
			F:            sweepIdentityDomainsRuleResource,
		})
	}
}

func sweepIdentityDomainsRuleResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	ruleIds, err := getIdentityDomainsRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, ruleId := range ruleIds {
		if ok := acctest.SweeperDefaultResourceId[ruleId]; !ok {
			deleteRuleRequest := oci_identity_domains.DeleteRuleRequest{}

			deleteRuleRequest.RuleId = &ruleId

			deleteRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteRule(context.Background(), deleteRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting Rule %s %s, It is possible that the resource is already deleted. Please verify manually \n", ruleId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listRulesRequest := oci_identity_domains.ListRulesRequest{}
	listRulesResponse, err := identityDomainsClient.ListRules(context.Background(), listRulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Rule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, rule := range listRulesResponse.Resources {
		id := *rule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RuleId", id)
	}
	return resourceIds, nil
}
