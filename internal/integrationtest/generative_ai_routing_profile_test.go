// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiRoutingProfileRequiredOnlyResource = GenerativeAiRoutingProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Required, acctest.Create, GenerativeAiRoutingProfileRepresentation)

	GenerativeAiRoutingProfileResourceConfig = GenerativeAiRoutingProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Optional, acctest.Update, GenerativeAiRoutingProfileRepresentation)

	GenerativeAiRoutingProfileSingularDataSourceRepresentation = map[string]interface{}{
		"routing_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_routing_profile.test_routing_profile.id}`},
	}

	GenerativeAiRoutingProfileDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_routing_profile.test_routing_profile.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiRoutingProfileDataSourceFilterRepresentation}}
	GenerativeAiRoutingProfileDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_routing_profile.test_routing_profile.id}`}},
	}

	GenerativeAiRoutingProfileRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"model_routing_policy":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiRoutingProfileModelRoutingPolicyRepresentation},
		"region_routing_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiRoutingProfileRegionRoutingPolicyRepresentation},
	}
	GenerativeAiRoutingProfileModelRoutingPolicyRepresentation = map[string]interface{}{
		"allowed_models": acctest.Representation{RepType: acctest.Optional, Create: []string{`meta.llama-3-70b-instruct`}, Update: []string{`meta.llama-3-70b-instruct`}},
	}
	GenerativeAiRoutingProfileRegionRoutingPolicyRepresentation = map[string]interface{}{
		"allowed_regions": acctest.Representation{RepType: acctest.Optional, Create: []string{`us-chicago-1`}, Update: []string{`us-chicago-1`}},
	}

	GenerativeAiRoutingProfileResourceDependencies = ""
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiRoutingProfileResource_basic(t *testing.T) {
	GenerativeAiRoutingProfileDataSourceRepresentation["id"] = acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_routing_profile.test_routing_profile.id}`, Update: `${oci_generative_ai_routing_profile.test_routing_profile.id}`}
	GenerativeAiRoutingProfileDataSourceRepresentation["state"] = acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`}

	httpreplay.SetScenario("TestGenerativeAiRoutingProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_routing_profile.test_routing_profile"
	datasourceName := "data.oci_generative_ai_routing_profiles.test_routing_profiles"
	singularDatasourceName := "data.oci_generative_ai_routing_profile.test_routing_profile"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiRoutingProfileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Optional, acctest.Create, GenerativeAiRoutingProfileRepresentation), "generativeai", "routingProfile", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiRoutingProfileDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiRoutingProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Required, acctest.Create, GenerativeAiRoutingProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiRoutingProfileResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiRoutingProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Optional, acctest.Create, GenerativeAiRoutingProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_routing_policy.0.allowed_models.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "region_routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "region_routing_policy.0.allowed_regions.#", "1"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiRoutingProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiRoutingProfileRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_routing_policy.0.allowed_models.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "region_routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "region_routing_policy.0.allowed_regions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + GenerativeAiRoutingProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Optional, acctest.Update, GenerativeAiRoutingProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_routing_policy.0.allowed_models.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "region_routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "region_routing_policy.0.allowed_regions.#", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_routing_profiles", "test_routing_profiles", acctest.Optional, acctest.Update, GenerativeAiRoutingProfileDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiRoutingProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Optional, acctest.Update, GenerativeAiRoutingProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "routing_profile_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "routing_profile_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_routing_profile", "test_routing_profile", acctest.Required, acctest.Create, GenerativeAiRoutingProfileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiRoutingProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "routing_profile_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_routing_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_routing_policy.0.allowed_models.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "region_routing_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "region_routing_policy.0.allowed_regions.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiRoutingProfileRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiRoutingProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_routing_profile" {
			noResourceFound = false
			request := oci_generative_ai.GetRoutingProfileRequest{}

			tmp := rs.Primary.ID
			request.RoutingProfileId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetRoutingProfile(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.RoutingProfileLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiRoutingProfile") {
		resource.AddTestSweepers("GenerativeAiRoutingProfile", &resource.Sweeper{
			Name:         "GenerativeAiRoutingProfile",
			Dependencies: acctest.DependencyGraph["routingProfile"],
			F:            sweepGenerativeAiRoutingProfileResource,
		})
	}
}

func sweepGenerativeAiRoutingProfileResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	routingProfileIds, err := getGenerativeAiRoutingProfileIds(compartment)
	if err != nil {
		return err
	}
	for _, routingProfileId := range routingProfileIds {
		if ok := acctest.SweeperDefaultResourceId[routingProfileId]; !ok {
			deleteRoutingProfileRequest := oci_generative_ai.DeleteRoutingProfileRequest{}

			deleteRoutingProfileRequest.RoutingProfileId = &routingProfileId

			deleteRoutingProfileRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteRoutingProfile(context.Background(), deleteRoutingProfileRequest)
			if error != nil {
				fmt.Printf("Error deleting RoutingProfile %s %s, It is possible that the resource is already deleted. Please verify manually \n", routingProfileId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &routingProfileId, GenerativeAiRoutingProfileSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiRoutingProfileSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiRoutingProfileIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RoutingProfileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listRoutingProfilesRequest := oci_generative_ai.ListRoutingProfilesRequest{}
	listRoutingProfilesRequest.CompartmentId = &compartmentId
	listRoutingProfilesRequest.LifecycleState = oci_generative_ai.RoutingProfileLifecycleStateActive
	listRoutingProfilesResponse, err := generativeAiClient.ListRoutingProfiles(context.Background(), listRoutingProfilesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RoutingProfile list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, routingProfile := range listRoutingProfilesResponse.Items {
		id := *routingProfile.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RoutingProfileId", id)
	}
	return resourceIds, nil
}

func GenerativeAiRoutingProfileSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if routingProfileResponse, ok := response.Response.(oci_generative_ai.GetRoutingProfileResponse); ok {
		return routingProfileResponse.LifecycleState != oci_generative_ai.RoutingProfileLifecycleStateDeleted
	}
	return false
}

func GenerativeAiRoutingProfileSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetRoutingProfile(context.Background(), oci_generative_ai.GetRoutingProfileRequest{
		RoutingProfileId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
