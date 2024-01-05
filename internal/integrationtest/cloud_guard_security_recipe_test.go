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
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudGuardSecurityRecipeRequiredOnlyResource = CloudGuardSecurityRecipeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Required, acctest.Create, CloudGuardSecurityRecipeRepresentation)

	CloudGuardSecurityRecipeResourceConfig = CloudGuardSecurityRecipeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Optional, acctest.Update, CloudGuardSecurityRecipeRepresentation)

	CloudGuardCloudGuardSecurityRecipeSingularDataSourceRepresentation = map[string]interface{}{
		"security_recipe_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_security_recipe.test_security_recipe.id}`},
	}

	CloudGuardCloudGuardSecurityRecipeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_security_recipe.test_security_recipe.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardSecurityRecipeDataSourceFilterRepresentation}}
	CloudGuardSecurityRecipeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_security_recipe.test_security_recipe.id}`}},
	}

	CloudGuardSecurityRecipeRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"security_policies": acctest.Representation{RepType: acctest.Required, Create: []string{securityPolicyId}, Update: []string{securityPolicyId}},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CloudGuardSecurityRecipeResourceDependencies = DefinedTagsDependencies + SecurityPolicyResourceDependencies

	CloudGuardSecurityPolicyDataSourceRepresentationPluralDataSource = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("compartment_ocid")},
	}
	securityPolicyId                   = `${data.oci_cloud_guard_security_policies.oracle_security_policy.security_policy_collection.0.items.0.id}`
	SecurityPolicyResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_policies", "oracle_security_policy", acctest.Required, acctest.Create, CloudGuardSecurityPolicyDataSourceRepresentationPluralDataSource)
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardSecurityRecipeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardSecurityRecipeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_security_recipe.test_security_recipe"
	datasourceName := "data.oci_cloud_guard_security_recipes.test_security_recipes"
	singularDatasourceName := "data.oci_cloud_guard_security_recipe.test_security_recipe"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardSecurityRecipeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Optional, acctest.Create, CloudGuardSecurityRecipeRepresentation), "cloudguard", "securityRecipe", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardSecurityRecipeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardSecurityRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Required, acctest.Create, CloudGuardSecurityRecipeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "security_policies.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardSecurityRecipeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardSecurityRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Optional, acctest.Create, CloudGuardSecurityRecipeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "owner"),
				resource.TestCheckResourceAttr(resourceName, "security_policies.#", "1"),

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
			Config: config + compartmentIdVariableStr + CloudGuardSecurityRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudGuardSecurityRecipeRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "owner"),
				resource.TestCheckResourceAttr(resourceName, "security_policies.#", "1"),

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
			Config: config + compartmentIdVariableStr + CloudGuardSecurityRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Optional, acctest.Update, CloudGuardSecurityRecipeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "owner"),
				resource.TestCheckResourceAttr(resourceName, "security_policies.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_recipes", "test_security_recipes", acctest.Optional, acctest.Update, CloudGuardCloudGuardSecurityRecipeDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardSecurityRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Optional, acctest.Update, CloudGuardSecurityRecipeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "security_recipe_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_recipe_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Required, acctest.Create, CloudGuardCloudGuardSecurityRecipeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardSecurityRecipeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_recipe_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_policies.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudGuardSecurityRecipeRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudGuardSecurityRecipeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_security_recipe" {
			noResourceFound = false
			request := oci_cloud_guard.GetSecurityRecipeRequest{}

			tmp := rs.Primary.ID
			request.SecurityRecipeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetSecurityRecipe(context.Background(), request)

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CloudGuardSecurityRecipe") {
		resource.AddTestSweepers("CloudGuardSecurityRecipe", &resource.Sweeper{
			Name:         "CloudGuardSecurityRecipe",
			Dependencies: acctest.DependencyGraph["securityRecipe"],
			F:            sweepCloudGuardSecurityRecipeResource,
		})
	}
}

func sweepCloudGuardSecurityRecipeResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	securityRecipeIds, err := getCloudGuardSecurityRecipeIds(compartment)
	if err != nil {
		return err
	}
	for _, securityRecipeId := range securityRecipeIds {
		if ok := acctest.SweeperDefaultResourceId[securityRecipeId]; !ok {
			deleteSecurityRecipeRequest := oci_cloud_guard.DeleteSecurityRecipeRequest{}

			deleteSecurityRecipeRequest.SecurityRecipeId = &securityRecipeId

			deleteSecurityRecipeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteSecurityRecipe(context.Background(), deleteSecurityRecipeRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityRecipe %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityRecipeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &securityRecipeId, CloudGuardSecurityRecipeSweepWaitCondition, time.Duration(3*time.Minute),
				CloudGuardSecurityRecipeSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getCloudGuardSecurityRecipeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityRecipeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listSecurityRecipesRequest := oci_cloud_guard.ListSecurityRecipesRequest{}
	listSecurityRecipesRequest.CompartmentId = &compartmentId
	listSecurityRecipesRequest.LifecycleState = oci_cloud_guard.ListSecurityRecipesLifecycleStateActive
	listSecurityRecipesResponse, err := cloudGuardClient.ListSecurityRecipes(context.Background(), listSecurityRecipesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityRecipe list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityRecipe := range listSecurityRecipesResponse.Items {
		id := *securityRecipe.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityRecipeId", id)
	}
	return resourceIds, nil
}

func CloudGuardSecurityRecipeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityRecipeResponse, ok := response.Response.(oci_cloud_guard.GetSecurityRecipeResponse); ok {
		return securityRecipeResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func CloudGuardSecurityRecipeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CloudGuardClient().GetSecurityRecipe(context.Background(), oci_cloud_guard.GetSecurityRecipeRequest{
		SecurityRecipeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
