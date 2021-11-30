// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_devops "github.com/oracle/oci-go-sdk/v53/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DevopsRepositoryResourceConfig = DevopsRepositoryResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Optional, Update, devopsRepositoryRepresentation)

	devopsRepositorySingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"fields":        Representation{RepType: Required, Create: []string{`branchCount`, `commitCount`, `sizeInBytes`}},
	}

	devopsRepositoryDataSourceRepresentation = map[string]interface{}{
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	devopsRepositoryRepresentation = map[string]interface{}{
		"name":       Representation{RepType: Required, Create: `name`, Update: `name2`},
		"project_id": Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		//"default_branch":  Representation{RepType: Optional, Create: `defaultBranch`},
		"defined_tags":    Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"freeform_tags":   Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"repository_type": Representation{RepType: Required, Create: `HOSTED`},
		"lifecycle":       RepresentationGroup{Required, ignoreChangesRepositoryRepresentation},
	}

	ignoreChangesRepositoryRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}

	DevopsRepositoryResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsConnectionRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	githubAccessTokenVaultId := getEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr := fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	resourceName := "oci_devops_repository.test_repository"
	datasourceName := "data.oci_devops_repositories.test_repositories"
	singularDatasourceName := "data.oci_devops_repository.test_repository"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+githubAccessTokenVaultIdStr+DevopsRepositoryResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Optional, Create, devopsRepositoryRepresentation), "devops", "repository", t)

	ResourceTest(t, testAccCheckDevopsRepositoryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsRepositoryResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsRepositoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsRepositoryResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsRepositoryResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Optional, Create, devopsRepositoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				//resource.TestCheckResourceAttr(resourceName, "default_branch", "defaultBranch"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "repository_type", "HOSTED"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsRepositoryResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Optional, Update, devopsRepositoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				//resource.TestCheckResourceAttr(resourceName, "default_branch", "defaultBranch"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "repository_type", "HOSTED"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repositories", "test_repositories", Optional, Update, devopsRepositoryDataSourceRepresentation) +
				compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsRepositoryResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Optional, Update, devopsRepositoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "repository_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "repository_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsRepositorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsRepositoryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "branch_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "commit_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "default_branch", "defaultBranch"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "http_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repository_type", "HOSTED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ssh_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trigger_build_events.#", "1"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + DevopsRepositoryResourceConfig,
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

func testAccCheckDevopsRepositoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_repository" {
			noResourceFound = false
			request := oci_devops.GetRepositoryRequest{}

			if value, ok := rs.Primary.Attributes["fields"]; ok {
				interfaces := strings.Split(value, " ")
				tmp := make([]oci_devops.GetRepositoryFieldsEnum, len(interfaces))
				for i := range interfaces {
					tmp[i] = interface{}(interfaces[i]).(oci_devops.GetRepositoryFieldsEnum)
				}
				if len(tmp) != 0 {
					request.Fields = tmp
				}
				request.Fields = tmp
			}

			tmp := rs.Primary.ID
			request.RepositoryId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")

			response, err := client.GetRepository(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.RepositoryLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("DevopsRepository") {
		resource.AddTestSweepers("DevopsRepository", &resource.Sweeper{
			Name:         "DevopsRepository",
			Dependencies: DependencyGraph["repository"],
			F:            sweepDevopsRepositoryResource,
		})
	}
}

func sweepDevopsRepositoryResource(compartment string) error {
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	repositoryIds, err := devopsGetRepositoryIds(compartment)
	if err != nil {
		return err
	}
	for _, repositoryId := range repositoryIds {
		if ok := SweeperDefaultResourceId[repositoryId]; !ok {
			deleteRepositoryRequest := oci_devops.DeleteRepositoryRequest{}

			deleteRepositoryRequest.RepositoryId = &repositoryId

			deleteRepositoryRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteRepository(context.Background(), deleteRepositoryRequest)
			if error != nil {
				fmt.Printf("Error deleting Repository %s %s, It is possible that the resource is already deleted. Please verify manually \n", repositoryId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &repositoryId, devopsRepositorySweepWaitCondition, time.Duration(3*time.Minute),
				devopsRepositorySweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func devopsGetRepositoryIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "RepositoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()

	listRepositoriesRequest := oci_devops.ListRepositoriesRequest{}
	listRepositoriesRequest.CompartmentId = &compartmentId
	listRepositoriesRequest.LifecycleState = oci_devops.RepositoryLifecycleStateActive
	listRepositoriesResponse, err := devopsClient.ListRepositories(context.Background(), listRepositoriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Repository list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, repository := range listRepositoriesResponse.Items {
		id := *repository.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "RepositoryId", id)
		SweeperDefaultResourceId[*repository.DefaultBranch] = true

	}
	return resourceIds, nil
}

func devopsRepositorySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if repositoryResponse, ok := response.Response.(oci_devops.GetRepositoryResponse); ok {
		return repositoryResponse.LifecycleState != oci_devops.RepositoryLifecycleStateDeleted
	}
	return false
}

func devopsRepositorySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.devopsClient().GetRepository(context.Background(), oci_devops.GetRepositoryRequest{
		RepositoryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
