// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RepositoryRefResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", acctest.Optional, acctest.Update, repositoryRefRepresentation)

	repositoryRefSingularDataSourceRepresentation = map[string]interface{}{
		"ref_name":      acctest.Representation{RepType: acctest.Required, Create: `refName`},
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	repositoryRefDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	repositoryRefRepresentation = map[string]interface{}{
		"ref_name":      acctest.Representation{RepType: acctest.Required, Create: `refName`},
		"ref_type":      acctest.Representation{RepType: acctest.Required, Create: `BRANCH`},
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"commit_id":     acctest.Representation{RepType: acctest.Required, Create: `commitId`, Update: `commitId1`},
	}

	RepositoryRefResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, devopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryRefResource_basic(t *testing.T) {
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "RepoRef") {
		t.Skip("TestDevopsRepositoryRefResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryRefResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	commitId := utils.GetEnvSettingWithBlankDefault("commit_id")
	commitIdStr := fmt.Sprintf("variable \"commit_id\" { default = \"%s\" }\n", commitId)

	resourceName := "oci_devops_repository_ref.test_repository_ref"
	datasourceName := "data.oci_devops_repository_refs.test_repository_refs"
	singularDatasourceName := "data.oci_devops_repository_ref.test_repository_ref"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+commitIdStr+RepositoryRefResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", acctest.Required, acctest.Create, repositoryRefRepresentation), "devops", "repositoryRef", t)

	acctest.ResourceTest(t, testAccCheckDevopsRepositoryRefDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + commitIdStr + RepositoryRefResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", acctest.Required, acctest.Create, repositoryRefRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "commit_id"),
				resource.TestCheckResourceAttr(resourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttr(resourceName, "ref_type", "BRANCH"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + commitIdStr + RepositoryRefResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", acctest.Optional, acctest.Update, repositoryRefRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "commit_id"),
				resource.TestCheckResourceAttrSet(resourceName, "full_ref_name"),
				resource.TestCheckResourceAttr(resourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttr(resourceName, "ref_type", "BRANCH"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

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
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", acctest.Optional, acctest.Update, repositoryRefRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_refs", "test_repository_refs", acctest.Optional, acctest.Update, repositoryRefDataSourceRepresentation) +
				compartmentIdVariableStr + commitIdStr + RepositoryRefResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),
				resource.TestCheckResourceAttr(datasourceName, "repository_ref_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "repository_ref_collection.0.items.#", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", acctest.Required, acctest.Create, repositoryRefSingularDataSourceRepresentation) +
				compartmentIdVariableStr + commitIdStr + RepositoryRefResourceDependencies + RepositoryRefResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "commit_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "full_ref_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_type", "BRANCH"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + commitIdStr + RepositoryRefResourceDependencies + RepositoryRefResourceConfig,
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

func testAccCheckDevopsRepositoryRefDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_repository_ref" {
			noResourceFound = false
			request := oci_devops.GetRefRequest{}

			if value, ok := rs.Primary.Attributes["ref_name"]; ok {
				request.RefName = &value
			}

			if value, ok := rs.Primary.Attributes["repository_id"]; ok {
				request.RepositoryId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

			_, err := client.GetRef(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DevopsRepositoryRef") {
		resource.AddTestSweepers("DevopsRepositoryRef", &resource.Sweeper{
			Name:         "DevopsRepositoryRef",
			Dependencies: acctest.DependencyGraph["repositoryRef"],
			F:            sweepDevopsRepositoryRefResource,
		})
	}
}

func sweepDevopsRepositoryRefResource(compartment string) error {
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	repositoryRefIds, err := getRepositoryRefIds(compartment)
	if err != nil {
		return err
	}
	for _, repositoryRefId := range repositoryRefIds {
		if ok := acctest.SweeperDefaultResourceId[repositoryRefId]; !ok {
			deleteRefRequest := oci_devops.DeleteRefRequest{}

			deleteRefRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteRef(context.Background(), deleteRefRequest)
			if error != nil {
				fmt.Printf("Error deleting RepositoryRef %s %s, It is possible that the resource is already deleted. Please verify manually \n", repositoryRefId, error)
				continue
			}
		}
	}
	return nil
}

func getRepositoryRefIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RepositoryRefId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

	listRefsRequest := oci_devops.ListRefsRequest{}
	//listRefsRequest.CompartmentId = &compartmentId

	repositoryIds, error := devopsGetRepositoryIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting repositoryId required for RepositoryRef resource requests \n")
	}
	for _, repositoryId := range repositoryIds {
		listRefsRequest.RepositoryId = &repositoryId

		listRefsResponse, err := devopsClient.ListRefs(context.Background(), listRefsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting RepositoryRef list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, repositoryRef := range listRefsResponse.Items {
			id := *repositoryRef.GetRepositoryId()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RepositoryRefId", id)
		}

	}
	return resourceIds, nil
}
