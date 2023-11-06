// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsConditionRequiredOnlyResource = IdentityDomainsConditionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Required, acctest.Create, IdentityDomainsConditionRepresentation)

	IdentityDomainsConditionResourceConfig = IdentityDomainsConditionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Optional, acctest.Update, IdentityDomainsConditionRepresentation)

	IdentityDomainsConditionSingularDataSourceRepresentation = map[string]interface{}{
		"condition_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_condition.test_condition.id}`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsConditionDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"condition_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"condition_filter": acctest.Representation{RepType: acctest.Optional, Create: `name eq \"name2\"`},
		"attribute_sets":   acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":      acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"sort_by":          acctest.Representation{RepType: acctest.Optional, Create: `meta.created`},
		"sort_order":       acctest.Representation{RepType: acctest.Optional, Create: `descending`},
	}

	IdentityDomainsConditionRepresentation = map[string]interface{}{
		"attribute_name":        acctest.Representation{RepType: acctest.Required, Create: `attributeName`, Update: `attributeName2`},
		"attribute_value":       acctest.Representation{RepType: acctest.Required, Create: `attributeValue`, Update: `attributeValue2`},
		"idcs_endpoint":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":                  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"operator":              acctest.Representation{RepType: acctest.Required, Create: `eq`, Update: `ne`},
		"schemas":               acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:Condition`}},
		"attribute_sets":        acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"evaluate_condition_if": acctest.Representation{RepType: acctest.Optional, Create: `evaluateConditionIf`, Update: `evaluateConditionIf2`},
		"external_id":           acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"tags":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsConditionTagsRepresentation},
	}
	IdentityDomainsConditionTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsConditionResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsConditionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsConditionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_condition.test_condition"
	datasourceName := "data.oci_identity_domains_conditions.test_conditions"
	singularDatasourceName := "data.oci_identity_domains_condition.test_condition"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsConditionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Optional, acctest.Create, IdentityDomainsConditionRepresentation), "identitydomains", "condition", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsConditionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsConditionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Required, acctest.Create, IdentityDomainsConditionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_name", "attributeName"),
				resource.TestCheckResourceAttr(resourceName, "attribute_value", "attributeValue"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "operator", "eq"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsConditionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsConditionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Optional, acctest.Create, IdentityDomainsConditionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_name", "attributeName"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_value", "attributeValue"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "evaluate_condition_if", "evaluateConditionIf"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "operator", "eq"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "conditions", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsConditionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Optional, acctest.Update, IdentityDomainsConditionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_name", "attributeName2"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_value", "attributeValue2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "evaluate_condition_if", "evaluateConditionIf2"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "operator", "ne"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_conditions", "test_conditions", acctest.Optional, acctest.Update, IdentityDomainsConditionDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsConditionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Optional, acctest.Update, IdentityDomainsConditionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "condition_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "conditions.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.schemas.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.attribute_name", "attributeName2"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.attribute_value", "attributeValue2"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.evaluate_condition_if", "evaluateConditionIf2"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.external_id", "externalId"),
				resource.TestCheckResourceAttrSet(datasourceName, "conditions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.operator", "ne"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.tags.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.tags.0.key", "key2"),
				resource.TestCheckResourceAttr(datasourceName, "conditions.0.tags.0.value", "value2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_condition", "test_condition", acctest.Required, acctest.Create, IdentityDomainsConditionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsConditionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "condition_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "attribute_name", "attributeName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attribute_value", "attributeValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "evaluate_condition_if", "evaluateConditionIf2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "operator", "ne"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsConditionRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_condition", "conditions"),
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

func testAccCheckIdentityDomainsConditionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_condition" {
			noResourceFound = false
			request := oci_identity_domains.GetConditionRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			tmp := rs.Primary.ID
			request.ConditionId = &tmp

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetCondition(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsCondition") {
		resource.AddTestSweepers("IdentityDomainsCondition", &resource.Sweeper{
			Name:         "IdentityDomainsCondition",
			Dependencies: acctest.DependencyGraph["condition"],
			F:            sweepIdentityDomainsConditionResource,
		})
	}
}

func sweepIdentityDomainsConditionResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	conditionIds, err := getIdentityDomainsConditionIds(compartment)
	if err != nil {
		return err
	}
	for _, conditionId := range conditionIds {
		if ok := acctest.SweeperDefaultResourceId[conditionId]; !ok {
			deleteConditionRequest := oci_identity_domains.DeleteConditionRequest{}

			deleteConditionRequest.ConditionId = &conditionId

			deleteConditionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteCondition(context.Background(), deleteConditionRequest)
			if error != nil {
				fmt.Printf("Error deleting Condition %s %s, It is possible that the resource is already deleted. Please verify manually \n", conditionId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsConditionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConditionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listConditionsRequest := oci_identity_domains.ListConditionsRequest{}
	listConditionsResponse, err := identityDomainsClient.ListConditions(context.Background(), listConditionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Condition list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, condition := range listConditionsResponse.Resources {
		id := *condition.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConditionId", id)
	}
	return resourceIds, nil
}
