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
	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ArtifactsContainerRepositoryRequiredOnlyResource = ArtifactsContainerRepositoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Required, acctest.Create, ArtifactscontainerRepositoryRepresentation)

	ArtifactsContainerRepositoryResourceConfig = ArtifactsContainerRepositoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Optional, acctest.Update, ArtifactscontainerRepositoryRepresentation)

	ArtifactsArtifactscontainerRepositorySingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_artifacts_container_repository.test_container_repository.id}`},
	}

	ArtifactsArtifactscontainerRepositoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_public":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"repository_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_artifacts_container_repository.test_container_repository.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ArtifactscontainerRepositoryDataSourceFilterRepresentation}}
	ArtifactscontainerRepositoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_artifacts_container_repository.test_container_repository.id}`}},
	}

	ArtifactscontainerRepositoryRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: utils.RandomString(15, utils.CharsetLowerCaseWithoutDigits)},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_immutable":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_public":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"readme":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ArtifactscontainerRepositoryReadmeRepresentation},
	}
	ArtifactscontainerRepositoryReadmeRepresentation = map[string]interface{}{
		"content": acctest.Representation{RepType: acctest.Required, Create: `content`, Update: `content2`},
		"format":  acctest.Representation{RepType: acctest.Required, Create: `TEXT_MARKDOWN`, Update: `TEXT_PLAIN`},
	}

	ArtifactsContainerRepositoryResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: artifacts/default
func TestArtifactsContainerRepositoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestArtifactsContainerRepositoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_artifacts_container_repository.test_container_repository"
	datasourceName := "data.oci_artifacts_container_repositories.test_container_repositories"
	singularDatasourceName := "data.oci_artifacts_container_repository.test_container_repository"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ArtifactsContainerRepositoryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Optional, acctest.Create, ArtifactscontainerRepositoryRepresentation), "artifacts", "containerRepository", t)

	acctest.ResourceTest(t, testAccCheckArtifactsContainerRepositoryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ArtifactsContainerRepositoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Required, acctest.Create, ArtifactscontainerRepositoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ArtifactsContainerRepositoryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ArtifactsContainerRepositoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Optional, acctest.Create, ArtifactscontainerRepositoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billable_size_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image_count"),
				resource.TestCheckResourceAttr(resourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "layer_count"),
				resource.TestCheckResourceAttrSet(resourceName, "layers_size_in_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "readme.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "readme.0.content", "content"),
				resource.TestCheckResourceAttr(resourceName, "readme.0.format", "TEXT_MARKDOWN"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ArtifactsContainerRepositoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ArtifactscontainerRepositoryRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billable_size_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image_count"),
				resource.TestCheckResourceAttr(resourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "layer_count"),
				resource.TestCheckResourceAttrSet(resourceName, "layers_size_in_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "readme.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "readme.0.content", "content"),
				resource.TestCheckResourceAttr(resourceName, "readme.0.format", "TEXT_MARKDOWN"),
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
			Config: config + compartmentIdVariableStr + ArtifactsContainerRepositoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Optional, acctest.Update, ArtifactscontainerRepositoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billable_size_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image_count"),
				resource.TestCheckResourceAttr(resourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "layer_count"),
				resource.TestCheckResourceAttrSet(resourceName, "layers_size_in_bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "readme.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "readme.0.content", "content2"),
				resource.TestCheckResourceAttr(resourceName, "readme.0.format", "TEXT_PLAIN"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_artifacts_container_repositories", "test_container_repositories", acctest.Optional, acctest.Update, ArtifactsArtifactscontainerRepositoryDataSourceRepresentation) +
				compartmentIdVariableStr + ArtifactsContainerRepositoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Optional, acctest.Update, ArtifactscontainerRepositoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(datasourceName, "is_public", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "container_repository_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_repository_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_repository_collection.0.items.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_repository_collection.0.items.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "container_repository_collection.0.image_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "container_repository_collection.0.layer_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "container_repository_collection.0.layers_size_in_bytes"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_artifacts_container_repository", "test_container_repository", acctest.Required, acctest.Create, ArtifactsArtifactscontainerRepositorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ArtifactsContainerRepositoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "billable_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_public", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "layer_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "layers_size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "readme.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "readme.0.content", "content2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "readme.0.format", "TEXT_PLAIN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + ArtifactsContainerRepositoryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckArtifactsContainerRepositoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ArtifactsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_artifacts_container_repository" {
			noResourceFound = false
			request := oci_artifacts.GetContainerRepositoryRequest{}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.RepositoryId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "artifacts")

			response, err := client.GetContainerRepository(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_artifacts.ContainerRepositoryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ArtifactsContainerRepository") {
		resource.AddTestSweepers("ArtifactsContainerRepository", &resource.Sweeper{
			Name:         "ArtifactsContainerRepository",
			Dependencies: acctest.DependencyGraph["containerRepository"],
			F:            sweepArtifactsContainerRepositoryResource,
		})
	}
}

func sweepArtifactsContainerRepositoryResource(compartment string) error {
	artifactsClient := acctest.GetTestClients(&schema.ResourceData{}).ArtifactsClient()
	containerRepositoryIds, err := getArtifactsContainerRepositoryIds(compartment)
	if err != nil {
		return err
	}
	for _, containerRepositoryId := range containerRepositoryIds {
		if ok := acctest.SweeperDefaultResourceId[containerRepositoryId]; !ok {
			deleteContainerRepositoryRequest := oci_artifacts.DeleteContainerRepositoryRequest{}

			deleteContainerRepositoryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "artifacts")
			_, error := artifactsClient.DeleteContainerRepository(context.Background(), deleteContainerRepositoryRequest)
			if error != nil {
				fmt.Printf("Error deleting ContainerRepository %s %s, It is possible that the resource is already deleted. Please verify manually \n", containerRepositoryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &containerRepositoryId, ArtifactscontainerRepositoriesSweepWaitCondition, time.Duration(3*time.Minute),
				ArtifactscontainerRepositoriesSweepResponseFetchOperation, "artifacts", true)
		}
	}
	return nil
}

func getArtifactsContainerRepositoryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ContainerRepositoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	artifactsClient := acctest.GetTestClients(&schema.ResourceData{}).ArtifactsClient()

	listContainerRepositoriesRequest := oci_artifacts.ListContainerRepositoriesRequest{}
	listContainerRepositoriesRequest.CompartmentId = &compartmentId
	state := string(oci_artifacts.ContainerRepositoryLifecycleStateAvailable)
	listContainerRepositoriesRequest.LifecycleState = &state
	listContainerRepositoriesResponse, err := artifactsClient.ListContainerRepositories(context.Background(), listContainerRepositoriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ContainerRepository list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, containerRepository := range listContainerRepositoriesResponse.Items {
		id := *containerRepository.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ContainerRepositoryId", id)
	}
	return resourceIds, nil
}

func ArtifactscontainerRepositoriesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if containerRepositoryResponse, ok := response.Response.(oci_artifacts.GetContainerRepositoryResponse); ok {
		return containerRepositoryResponse.LifecycleState != oci_artifacts.ContainerRepositoryLifecycleStateDeleted
	}
	return false
}

func ArtifactscontainerRepositoriesSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ArtifactsClient().GetContainerRepository(context.Background(), oci_artifacts.GetContainerRepositoryRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
