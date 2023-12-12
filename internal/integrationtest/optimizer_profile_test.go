// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_optimizer "github.com/oracle/oci-go-sdk/v65/optimizer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OptimizerProfileRequiredOnlyResource = OptimizerProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", acctest.Required, acctest.Create, OptimizerProfileRepresentation)

	OptimizerProfileResourceConfig = OptimizerProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", acctest.Optional, acctest.Update, OptimizerProfileRepresentation)

	OptimizerOptimizerProfileSingularDataSourceRepresentation = map[string]interface{}{
		"profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_optimizer_profile.test_profile.id}`},
	}

	OptimizerOptimizerProfileDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	OptimizerProfileRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":                  acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"levels_configuration":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OptimizerProfileLevelsConfigurationRepresentation},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"aggregation_interval_in_days": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `7`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"target_compartments":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: OptimizerProfileTargetCompartmentsRepresentation},
		"target_tags":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: OptimizerProfileTargetTagsRepresentation},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreProfileChangesRep},
	}
	OptimizerProfileLevelsConfigurationRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: OptimizerProfileLevelsConfigurationItemsRepresentation},
	}
	OptimizerProfileTargetCompartmentsRepresentation = map[string]interface{}{
		"items": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.compartment_id}`}, Update: []string{`${var.compartment_id_for_update}`}},
	}
	OptimizerProfileTargetTagsRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: OptimizerProfileTargetTagsItemsRepresentation},
	}
	OptimizerProfileLevelsConfigurationItemsRepresentation = map[string]interface{}{
		"level":             acctest.Representation{RepType: acctest.Required, Create: `cost-compute_aggressive_average`, Update: `cost-compute_conservative_average`},
		"recommendation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_optimizer_recommendation.test_recommendation.recommendation_id}`},
	}
	OptimizerProfileTargetTagsItemsRepresentation = map[string]interface{}{
		"tag_definition_name": acctest.Representation{RepType: acctest.Required, Create: `tagDefinitionName`, Update: `tagDefinitionName2`},
		"tag_namespace_name":  acctest.Representation{RepType: acctest.Required, Create: `tagNamespaceName`, Update: `tagNamespaceName2`},
		"tag_value_type":      acctest.Representation{RepType: acctest.Required, Create: `VALUE`, Update: `ANY`},
		"tag_values":          acctest.Representation{RepType: acctest.Optional, Create: []string{`tagValue1`}, Update: []string{}},
	}

	OptimizerProfileResourceDependencies = DefinedTagsDependencies + OptimizerRecommendationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Required, acctest.Create, OptimizerRecommendationRepresentation)
	ignoreProfileChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}
)

// issue-routing-tag: optimizer/default
func TestOptimizerProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_optimizer_profile.test_profile"
	datasourceName := "data.oci_optimizer_profiles.test_profiles"
	singularDatasourceName := "data.oci_optimizer_profile.test_profile"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OptimizerProfileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", acctest.Optional, acctest.Create, OptimizerProfileRepresentation), "optimizer", "profile", t)

	acctest.ResourceTest(t, testAccCheckOptimizerProfileDestroy, []resource.TestStep{
		// Pre-requisite: There shouldn't be a profile with the same <recommendationId, targetCompartment, targetTags> combination or with same name existing for the compartmentId
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OptimizerProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", acctest.Required, acctest.Create, OptimizerProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "levels_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OptimizerProfileResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Service limits may take 2 minutes to be available post deletion")
				time.Sleep(2 * time.Minute)
				return nil
			},
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OptimizerProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", acctest.Optional, acctest.Create, OptimizerProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "aggregation_interval_in_days", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "levels_configuration.0.items.0.level", "cost-compute_aggressive_average"),
				resource.TestCheckResourceAttrSet(resourceName, "levels_configuration.0.items.0.recommendation_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.0.items.0", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "target_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.0.tag_definition_name", "tagDefinitionName"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.0.tag_namespace_name", "tagNamespaceName"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.0.tag_value_type", "VALUE"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.0.tag_values.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OptimizerProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", acctest.Optional, acctest.Update, OptimizerProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "aggregation_interval_in_days", "7"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "levels_configuration.0.items.0.level", "cost-compute_conservative_average"),
				resource.TestCheckResourceAttrSet(resourceName, "levels_configuration.0.items.0.recommendation_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.0.items.0", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "target_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.0.tag_definition_name", "tagDefinitionName2"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.0.tag_namespace_name", "tagNamespaceName2"),
				resource.TestCheckResourceAttr(resourceName, "target_tags.0.items.0.tag_value_type", "ANY"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_profiles", "test_profiles", acctest.Optional, acctest.Update, OptimizerOptimizerProfileDataSourceRepresentation) +
				compartmentIdVariableStr + compartmentIdUVariableStr + OptimizerProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", acctest.Optional, acctest.Update, OptimizerProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "profile_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_profile", "test_profile", acctest.Required, acctest.Create, OptimizerOptimizerProfileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + compartmentIdUVariableStr + OptimizerProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "aggregation_interval_in_days", "7"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "levels_configuration.0.items.0.level", "cost-compute_conservative_average"),
				resource.TestCheckResourceAttrSet(resourceName, "levels_configuration.0.items.0.recommendation_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_compartments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_compartments.0.items.0", compartmentIdU),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_tags.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_tags.0.items.0.tag_definition_name", "tagDefinitionName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_tags.0.items.0.tag_namespace_name", "tagNamespaceName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_tags.0.items.0.tag_value_type", "ANY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OptimizerProfileRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOptimizerProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OptimizerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_optimizer_profile" {
			noResourceFound = false
			request := oci_optimizer.GetProfileRequest{}

			tmp := rs.Primary.ID
			request.ProfileId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "optimizer")

			response, err := client.GetProfile(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_optimizer.ListProfilesLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OptimizerProfile") {
		resource.AddTestSweepers("OptimizerProfile", &resource.Sweeper{
			Name:         "OptimizerProfile",
			Dependencies: acctest.DependencyGraph["profile"],
			F:            sweepOptimizerProfileResource,
		})
	}
}

func sweepOptimizerProfileResource(compartment string) error {
	optimizerClient := acctest.GetTestClients(&schema.ResourceData{}).OptimizerClient()
	profileIds, err := getOptimizerProfileIds(compartment)
	if err != nil {
		return err
	}
	for _, profileId := range profileIds {
		if ok := acctest.SweeperDefaultResourceId[profileId]; !ok {
			deleteProfileRequest := oci_optimizer.DeleteProfileRequest{}

			deleteProfileRequest.ProfileId = &profileId

			deleteProfileRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "optimizer")
			_, error := optimizerClient.DeleteProfile(context.Background(), deleteProfileRequest)
			if error != nil {
				fmt.Printf("Error deleting Profile %s %s, It is possible that the resource is already deleted. Please verify manually \n", profileId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &profileId, OptimizerProfileSweepWaitCondition, time.Duration(3*time.Minute),
				OptimizerProfileSweepResponseFetchOperation, "optimizer", true)
		}
	}
	return nil
}

func getOptimizerProfileIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProfileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	optimizerClient := acctest.GetTestClients(&schema.ResourceData{}).OptimizerClient()

	listProfilesRequest := oci_optimizer.ListProfilesRequest{}
	listProfilesRequest.CompartmentId = &compartmentId
	listProfilesRequest.LifecycleState = oci_optimizer.ListProfilesLifecycleStateActive
	listProfilesResponse, err := optimizerClient.ListProfiles(context.Background(), listProfilesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Profile list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, profile := range listProfilesResponse.Items {
		id := *profile.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProfileId", id)
	}
	return resourceIds, nil
}

func OptimizerProfileSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if profileResponse, ok := response.Response.(oci_optimizer.GetProfileResponse); ok {
		return profileResponse.LifecycleState != oci_optimizer.LifecycleStateDeleted
	}
	return false
}

func OptimizerProfileSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OptimizerClient().GetProfile(context.Background(), oci_optimizer.GetProfileRequest{
		ProfileId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
