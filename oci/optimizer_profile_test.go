// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v35/common"
	oci_optimizer "github.com/oracle/oci-go-sdk/v35/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ProfileRequiredOnlyResource = ProfileResourceDependencies +
		generateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", Required, Create, profileRepresentation)

	ProfileResourceConfig = ProfileResourceDependencies +
		generateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", Optional, Update, profileRepresentation)

	profileSingularDataSourceRepresentation = map[string]interface{}{
		"profile_id": Representation{repType: Required, create: `${oci_optimizer_profile.test_profile.id}`},
	}

	profileDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"name":           Representation{repType: Optional, create: `name`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
	}

	profileRepresentation = map[string]interface{}{
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"description":          Representation{repType: Required, create: `description`, update: `description2`},
		"levels_configuration": RepresentationGroup{Required, profileLevelsConfigurationRepresentation},
		"name":                 Representation{repType: Required, create: `name`},
		"defined_tags":         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":        Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}
	profileLevelsConfigurationRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, profileLevelsConfigurationItemsRepresentation},
	}
	profileLevelsConfigurationItemsRepresentation = map[string]interface{}{
		"level":             Representation{repType: Required, create: `cost-compute_aggressive_average`, update: `cost-compute_conservative_average`},
		"recommendation_id": Representation{repType: Required, create: `${oci_optimizer_recommendation.test_recommendation.recommendation_id}`},
	}

	ProfileResourceDependencies = DefinedTagsDependencies + RecommendationResourceDependencies +
		generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation)
)

func TestOptimizerProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerProfileResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_optimizer_profile.test_profile"
	datasourceName := "data.oci_optimizer_profiles.test_profiles"
	singularDatasourceName := "data.oci_optimizer_profile.test_profile"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOptimizerProfileDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ProfileResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", Required, Create, profileRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "levels_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ProfileResourceDependencies,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] Service limits may take 2 minutes to be available post deletion")
					time.Sleep(2 * time.Minute)
					return nil
				},
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ProfileResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", Optional, Create, profileRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "levels_configuration.0.items.0.level", "cost-compute_aggressive_average"),
					resource.TestCheckResourceAttrSet(resourceName, "levels_configuration.0.items.0.recommendation_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ProfileResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", Optional, Update, profileRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "levels_configuration.0.items.0.level", "cost-compute_conservative_average"),
					resource.TestCheckResourceAttrSet(resourceName, "levels_configuration.0.items.0.recommendation_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
					generateDataSourceFromRepresentationMap("oci_optimizer_profiles", "test_profiles", Optional, Update, profileDataSourceRepresentation) +
					compartmentIdVariableStr + ProfileResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_profile", "test_profile", Optional, Update, profileRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "name"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttrSet(datasourceName, "profile_collection.#"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_optimizer_profile", "test_profile", Required, Create, profileSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ProfileResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "levels_configuration.0.items.0.level", "cost-compute_conservative_average"),
					resource.TestCheckResourceAttrSet(resourceName, "levels_configuration.0.items.0.recommendation_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ProfileResourceConfig,
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

func testAccCheckOptimizerProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).optimizerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_optimizer_profile" {
			noResourceFound = false
			request := oci_optimizer.GetProfileRequest{}

			tmp := rs.Primary.ID
			request.ProfileId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "optimizer")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("OptimizerProfile") {
		resource.AddTestSweepers("OptimizerProfile", &resource.Sweeper{
			Name:         "OptimizerProfile",
			Dependencies: DependencyGraph["profile"],
			F:            sweepOptimizerProfileResource,
		})
	}
}

func sweepOptimizerProfileResource(compartment string) error {
	optimizerClient := GetTestClients(&schema.ResourceData{}).optimizerClient()
	profileIds, err := getProfileIds(compartment)
	if err != nil {
		return err
	}
	for _, profileId := range profileIds {
		if ok := SweeperDefaultResourceId[profileId]; !ok {
			deleteProfileRequest := oci_optimizer.DeleteProfileRequest{}

			deleteProfileRequest.ProfileId = &profileId

			deleteProfileRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "optimizer")
			_, error := optimizerClient.DeleteProfile(context.Background(), deleteProfileRequest)
			if error != nil {
				fmt.Printf("Error deleting Profile %s %s, It is possible that the resource is already deleted. Please verify manually \n", profileId, error)
				continue
			}
			waitTillCondition(testAccProvider, &profileId, profileSweepWaitCondition, time.Duration(3*time.Minute),
				profileSweepResponseFetchOperation, "optimizer", true)
		}
	}
	return nil
}

func getProfileIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ProfileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	optimizerClient := GetTestClients(&schema.ResourceData{}).optimizerClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ProfileId", id)
	}
	return resourceIds, nil
}

func profileSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if profileResponse, ok := response.Response.(oci_optimizer.GetProfileResponse); ok {
		return profileResponse.LifecycleState != oci_optimizer.LifecycleStateDeleted
	}
	return false
}

func profileSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.optimizerClient().GetProfile(context.Background(), oci_optimizer.GetProfileRequest{
		ProfileId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
