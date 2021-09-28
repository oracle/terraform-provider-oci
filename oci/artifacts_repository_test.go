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
	oci_artifacts "github.com/oracle/oci-go-sdk/v48/artifacts"
	"github.com/oracle/oci-go-sdk/v48/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RepositoryRequiredOnlyResource = RepositoryResourceDependencies +
		generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Required, Create, repositoryRepresentation)

	RepositoryResourceConfig = RepositoryResourceDependencies +
		generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Optional, Update, repositoryRepresentation)

	repositorySingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": Representation{repType: Required, create: `${oci_artifacts_repository.test_repository.id}`},
	}

	repositoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"id":             Representation{repType: Optional, create: `${oci_artifacts_repository.test_repository.id}`},
		"is_immutable":   Representation{repType: Optional, create: `false`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, repositoryDataSourceFilterRepresentation}}
	repositoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_artifacts_repository.test_repository.id}`}},
	}

	repositoryRepresentation = map[string]interface{}{
		"compartment_id":  Representation{repType: Required, create: `${var.compartment_id}`},
		"is_immutable":    Representation{repType: Required, create: `false`},
		"repository_type": Representation{repType: Required, create: `GENERIC`},
		"defined_tags":    Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":    Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":   Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	RepositoryResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: artifacts/default
func TestArtifactsRepositoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestArtifactsRepositoryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_artifacts_repository.test_repository"
	datasourceName := "data.oci_artifacts_repositories.test_repositories"
	singularDatasourceName := "data.oci_artifacts_repository.test_repository"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+RepositoryResourceDependencies+
		generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Optional, Create, repositoryRepresentation), "artifacts", "repository", t)

	ResourceTest(t, testAccCheckArtifactsRepositoryDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + RepositoryResourceDependencies +
				generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Required, Create, repositoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(resourceName, "repository_type", "GENERIC"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + RepositoryResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + RepositoryResourceDependencies +
				generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Optional, Create, repositoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(resourceName, "repository_type", "GENERIC"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RepositoryResourceDependencies +
				generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Optional, Create,
					representationCopyWithNewProperties(repositoryRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(resourceName, "repository_type", "GENERIC"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + RepositoryResourceDependencies +
				generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Optional, Update, repositoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(resourceName, "repository_type", "GENERIC"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				generateDataSourceFromRepresentationMap("oci_artifacts_repositories", "test_repositories", Optional, Update, repositoryDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryResourceDependencies +
				generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Optional, Update, repositoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "repository_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "repository_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Required, Create, repositorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_immutable", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + RepositoryResourceConfig,
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

func testAccCheckArtifactsRepositoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).artifactsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_artifacts_repository" {
			noResourceFound = false
			request := oci_artifacts.GetRepositoryRequest{}

			tmp := rs.Primary.ID
			request.RepositoryId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "artifacts")

			response, err := client.GetRepository(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_artifacts.RepositoryLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !inSweeperExcludeList("ArtifactsRepository") {
		resource.AddTestSweepers("ArtifactsRepository", &resource.Sweeper{
			Name:         "ArtifactsRepository",
			Dependencies: DependencyGraph["repository"],
			F:            sweepArtifactsRepositoryResource,
		})
	}
}

func sweepArtifactsRepositoryResource(compartment string) error {
	artifactsClient := GetTestClients(&schema.ResourceData{}).artifactsClient()
	repositoryIds, err := getRepositoryIds(compartment)
	if err != nil {
		return err
	}
	for _, repositoryId := range repositoryIds {
		if ok := SweeperDefaultResourceId[repositoryId]; !ok {
			deleteRepositoryRequest := oci_artifacts.DeleteRepositoryRequest{}

			deleteRepositoryRequest.RepositoryId = &repositoryId

			deleteRepositoryRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "artifacts")
			_, error := artifactsClient.DeleteRepository(context.Background(), deleteRepositoryRequest)
			if error != nil {
				fmt.Printf("Error deleting Repository %s %s, It is possible that the resource is already deleted. Please verify manually \n", repositoryId, error)
				continue
			}
			waitTillCondition(testAccProvider, &repositoryId, repositorySweepWaitCondition, time.Duration(3*time.Minute),
				repositorySweepResponseFetchOperation, "artifacts", true)
		}
	}
	return nil
}

func getRepositoryIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "RepositoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	artifactsClient := GetTestClients(&schema.ResourceData{}).artifactsClient()

	listRepositoriesRequest := oci_artifacts.ListRepositoriesRequest{}
	listRepositoriesRequest.CompartmentId = &compartmentId
	state := oci_artifacts.RepositoryLifecycleStateAvailable
	listRepositoriesRequest.LifecycleState = (*string)(&state)
	listRepositoriesResponse, err := artifactsClient.ListRepositories(context.Background(), listRepositoriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Repository list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, repository := range listRepositoriesResponse.Items {
		id := *repository.GetId()
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "RepositoryId", id)
	}
	return resourceIds, nil
}

func repositorySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if repositoryResponse, ok := response.Response.(oci_artifacts.GetRepositoryResponse); ok {
		return repositoryResponse.GetLifecycleState() != oci_artifacts.RepositoryLifecycleStateDeleted
	}
	return false
}

func repositorySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.artifactsClient().GetRepository(context.Background(), oci_artifacts.GetRepositoryRequest{
		RepositoryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
