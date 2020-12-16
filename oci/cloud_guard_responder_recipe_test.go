// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v31/cloudguard"
	"github.com/oracle/oci-go-sdk/v31/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ResponderRecipeRequiredOnlyResource = ResponderRecipeResourceDependencies +
		generateResourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", Required, Create, responderRecipeRepresentation)

	ResponderRecipeResourceConfig = ResponderRecipeResourceDependencies +
		generateResourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", Optional, Update, responderRecipeRepresentation)

	responderRecipeSingularDataSourceRepresentation = map[string]interface{}{
		"responder_recipe_id": Representation{repType: Required, create: `${oci_cloud_guard_responder_recipe.test_responder_recipe.id}`},
	}

	responderRecipeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"access_level":              Representation{repType: Optional, create: `ACCESSIBLE`},
		"compartment_id_in_subtree": Representation{repType: Optional, create: `true`},
		"display_name":              Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"resource_metadata_only":    Representation{repType: Optional, create: `false`},
		"state":                     Representation{repType: Optional, create: `ACTIVE`},
		"filter":                    RepresentationGroup{Required, responderRecipeDataSourceFilterRepresentation}}
	responderRecipeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_cloud_guard_responder_recipe.test_responder_recipe.id}`}},
	}

	//Making a list call and getting a source responderRecipeId
	responderRecipeId             = `${data.oci_cloud_guard_responder_recipes.oracle_responder_recipe.responder_recipe_collection.0.items.0.id}`
	responderRecipeRepresentation = map[string]interface{}{
		"compartment_id":             Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":               Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"source_responder_recipe_id": Representation{repType: Required, create: responderRecipeId},
		"defined_tags":               Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":              Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"responder_rules":            RepresentationGroup{Optional, responderRecipeResponderRulesRepresentation},
	}
	//hardcoding a responder-rule-id for testing purposes
	responderRecipeResponderRulesRepresentation = map[string]interface{}{
		"details":           RepresentationGroup{Required, responderRecipeResponderRulesDetailsRepresentation},
		"responder_rule_id": Representation{repType: Required, create: `MAKE_BUCKET_PRIVATE`},
	}
	responderRecipeResponderRulesDetailsRepresentation = map[string]interface{}{
		"is_enabled": Representation{repType: Required, create: `false`, update: `true`},
	}
	//Make a representation for plural datasource
	responderRecipeDataSourceRepresentationPluralDataSource = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: getEnvSettingWithBlankDefault("tenancy_ocid")},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
	}

	//Corrected the dependencies.
	//Removed tag dependencies and put in individual calls as the same is used in target and target will have from detectorRecipeDependencies as well so tags will be duplicated.
	ResponderRecipeResourceDependencies = generateDataSourceFromRepresentationMap("oci_cloud_guard_responder_recipes", "oracle_responder_recipe", Required, Create, responderRecipeDataSourceRepresentationPluralDataSource)
)

func TestCloudGuardResponderRecipeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardResponderRecipeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_guard_responder_recipe.test_responder_recipe"
	datasourceName := "data.oci_cloud_guard_responder_recipes.test_responder_recipes"
	singularDatasourceName := "data.oci_cloud_guard_responder_recipe.test_responder_recipe"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCloudGuardResponderRecipeDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ResponderRecipeResourceDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", Required, Create, responderRecipeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "source_responder_recipe_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ResponderRecipeResourceDependencies + DefinedTagsDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ResponderRecipeResourceDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", Optional, Create, responderRecipeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.0.details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.0.details.0.is_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "responder_rules.0.responder_rule_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source_responder_recipe_id"),

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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ResponderRecipeResourceDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", Optional, Create,
						representationCopyWithNewProperties(responderRecipeRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.0.details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.0.details.0.is_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "responder_rules.0.responder_rule_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source_responder_recipe_id"),

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
				Config: config + compartmentIdVariableStr + ResponderRecipeResourceDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", Optional, Update, responderRecipeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.0.details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "responder_rules.0.details.0.is_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "responder_rules.0.responder_rule_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source_responder_recipe_id"),

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
					generateDataSourceFromRepresentationMap("oci_cloud_guard_responder_recipes", "test_responder_recipes", Optional, Update, responderRecipeDataSourceRepresentation) +
					compartmentIdVariableStr + ResponderRecipeResourceDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", Optional, Update, responderRecipeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "resource_metadata_only", "false"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "responder_recipe_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "responder_recipe_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", Required, Create, responderRecipeSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ResponderRecipeResourceConfig + DefinedTagsDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "responder_recipe_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "responder_rules.#", "1"),
					//These are effective rules, after applying defaults over user input so here the count is more and can increase on addition of more rules.
					resource.TestCheckResourceAttrSet(singularDatasourceName, "effective_responder_rules.#"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "owner"),
					resource.TestCheckResourceAttr(singularDatasourceName, "responder_rules.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "responder_rules.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "responder_rules.0.description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "responder_rules.0.details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "responder_rules.0.details.0.is_enabled", "true"),
					//Since these are not passed in input, they can't be in input.
					//But these will be in effective_rules
					resource.TestCheckResourceAttr(singularDatasourceName, "effective_responder_rules.0.details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "effective_responder_rules.0.details.0.is_enabled", "true"),
					//Conditions and Configurations can be added from target level, hence if no I/P is there, no O/P will be there.
					resource.TestCheckResourceAttrSet(singularDatasourceName, "effective_responder_rules.0.details.0.mode"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "responder_rules.0.display_name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "responder_rules.0.policies.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "responder_rules.0.state"),
					//There are 2 supported modes.
					resource.TestCheckResourceAttr(singularDatasourceName, "responder_rules.0.supported_modes.#", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "responder_rules.0.time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "responder_rules.0.time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "responder_rules.0.type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ResponderRecipeResourceConfig + DefinedTagsDependencies,
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

func testAccCheckCloudGuardResponderRecipeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).cloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_responder_recipe" {
			noResourceFound = false
			request := oci_cloud_guard.GetResponderRecipeRequest{}

			tmp := rs.Primary.ID
			request.ResponderRecipeId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "cloud_guard")

			response, err := client.GetResponderRecipe(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_guard.LifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("CloudGuardResponderRecipe") {
		resource.AddTestSweepers("CloudGuardResponderRecipe", &resource.Sweeper{
			Name:         "CloudGuardResponderRecipe",
			Dependencies: DependencyGraph["responderRecipe"],
			F:            sweepCloudGuardResponderRecipeResource,
		})
	}
}

func sweepCloudGuardResponderRecipeResource(compartment string) error {
	cloudGuardClient := GetTestClients(&schema.ResourceData{}).cloudGuardClient()
	responderRecipeIds, err := getResponderRecipeIds(compartment)
	if err != nil {
		return err
	}
	for _, responderRecipeId := range responderRecipeIds {
		if ok := SweeperDefaultResourceId[responderRecipeId]; !ok {
			deleteResponderRecipeRequest := oci_cloud_guard.DeleteResponderRecipeRequest{}

			deleteResponderRecipeRequest.ResponderRecipeId = &responderRecipeId

			deleteResponderRecipeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteResponderRecipe(context.Background(), deleteResponderRecipeRequest)
			if error != nil {
				fmt.Printf("Error deleting ResponderRecipe %s %s, It is possible that the resource is already deleted. Please verify manually \n", responderRecipeId, error)
				continue
			}
			waitTillCondition(testAccProvider, &responderRecipeId, responderRecipeSweepWaitCondition, time.Duration(3*time.Minute),
				responderRecipeSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getResponderRecipeIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ResponderRecipeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := GetTestClients(&schema.ResourceData{}).cloudGuardClient()

	listResponderRecipesRequest := oci_cloud_guard.ListResponderRecipesRequest{}
	listResponderRecipesRequest.CompartmentId = &compartmentId
	listResponderRecipesRequest.LifecycleState = oci_cloud_guard.ListResponderRecipesLifecycleStateActive
	listResponderRecipesResponse, err := cloudGuardClient.ListResponderRecipes(context.Background(), listResponderRecipesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ResponderRecipe list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, responderRecipe := range listResponderRecipesResponse.Items {
		id := *responderRecipe.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ResponderRecipeId", id)
	}
	return resourceIds, nil
}

func responderRecipeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if responderRecipeResponse, ok := response.Response.(oci_cloud_guard.GetResponderRecipeResponse); ok {
		return responderRecipeResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func responderRecipeSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.cloudGuardClient().GetResponderRecipe(context.Background(), oci_cloud_guard.GetResponderRecipeRequest{
		ResponderRecipeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
