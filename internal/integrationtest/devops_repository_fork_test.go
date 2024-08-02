// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsForkRepositoryRepresentation = map[string]interface{}{
		"name":                 acctest.Representation{RepType: acctest.Required, Create: `fork-name`, Update: `fork-name-update`},
		"project_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"repository_type":      acctest.Representation{RepType: acctest.Required, Create: `FORKED`},
		"parent_repository_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_repository.test_repository.id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesForkRepositoryRepresentation},
	}
	ignoreChangesForkRepositoryRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	DevopsForkRepositoryResource = acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_fork_repository", acctest.Optional, acctest.Create, DevopsForkRepositoryRepresentation)

	DevopsForkRepositoryResourceConfig = DevopsForkRepositoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_fork_repository", acctest.Optional, acctest.Update, DevopsForkRepositoryRepresentation)

	DevopsForkRepositorySingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_fork_repository.id}`},
		"fields":        acctest.Representation{RepType: acctest.Required, Create: []string{`branchCount`, `commitCount`, `sizeInBytes`}},
	}

	DevopsForkRepositoryDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_fork_repository.id}`},
	}

	DevopsForkRepositoryResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsRepositoryRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsForkRepositoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsForkRepositoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_repository.test_fork_repository"
	datasourceName := "data.oci_devops_repositories.test_fork_repositories"
	singularDatasourceName := "data.oci_devops_repository.test_fork_repository"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsForkRepositoryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_fork_repository", acctest.Optional, acctest.Create, DevopsForkRepositoryRepresentation), "devops", "repository", t)

	acctest.ResourceTest(t, testAccCheckDevopsForkRepositoryDestroy, []resource.TestStep{
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsForkRepositoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_fork_repository", acctest.Optional, acctest.Create, DevopsForkRepositoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "fork-name"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "repository_type", "FORKED"),

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
			Config: config + compartmentIdVariableStr + DevopsForkRepositoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_fork_repository", acctest.Optional, acctest.Update, DevopsForkRepositoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "fork-name-update"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "repository_type", "FORKED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repositories", "test_fork_repositories", acctest.Optional, acctest.Update, DevopsForkRepositoryDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsForkRepositoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_fork_repository", acctest.Optional, acctest.Update, DevopsForkRepositoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "repository_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "repository_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository", "test_fork_repository", acctest.Optional, acctest.Create, DevopsForkRepositorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsForkRepositoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "fork-name-update"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ssh_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "http_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repository_type", "FORKED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "branch_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "commit_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trigger_build_events.#", "1"),
			),
		},

		// verify resource import
		{
			Config:                  config + DevopsForkRepositoryResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDevopsForkRepositoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
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

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DevopsForkRepository") {
		resource.AddTestSweepers("DevopsForkRepository", &resource.Sweeper{
			Name:         "DevopsForkRepository",
			Dependencies: acctest.DependencyGraph["repository"],
			F:            sweepDevopsForkRepositoryResource,
		})
	}
}

func sweepDevopsForkRepositoryResource(compartment string) error {
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	repositoryIds, err := getDevopsForkRepositoryIds(compartment)
	if err != nil {
		return err
	}
	for _, repositoryId := range repositoryIds {
		if ok := acctest.SweeperDefaultResourceId[repositoryId]; !ok {
			deleteRepositoryRequest := oci_devops.DeleteRepositoryRequest{}

			deleteRepositoryRequest.RepositoryId = &repositoryId

			deleteRepositoryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteRepository(context.Background(), deleteRepositoryRequest)
			if error != nil {
				fmt.Printf("Error deleting Repository %s %s, It is possible that the resource is already deleted. Please verify manually \n", repositoryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &repositoryId, DevopsForkRepositorySweepResponseFetchOperation, time.Duration(3*time.Minute),
				devopsForkRepositorySweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getDevopsForkRepositoryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RepositoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RepositoryId", id)
		acctest.SweeperDefaultResourceId[*repository.DefaultBranch] = true

	}
	return resourceIds, nil
}

func DevopsForkRepositorySweepResponseFetchOperation(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if repositoryResponse, ok := response.Response.(oci_devops.GetRepositoryResponse); ok {
		return repositoryResponse.LifecycleState != oci_devops.RepositoryLifecycleStateDeleted
	}
	return false
}

func devopsForkRepositorySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DevopsClient().GetRepository(context.Background(), oci_devops.GetRepositoryRequest{
		RepositoryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
