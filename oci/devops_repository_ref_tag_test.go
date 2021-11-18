// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_devops "github.com/oracle/oci-go-sdk/v52/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RepositoryTagRefResourceConfig = RepositoryTagRefResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", Optional, Update, repositoryTagRefRepresentation)

	repositoryTagRefSingularDataSourceRepresentation = map[string]interface{}{
		"ref_name":      Representation{RepType: Required, Create: `refName`},
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	repositoryTagRefDataSourceRepresentation = map[string]interface{}{
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"object_id":     Representation{RepType: Required, Create: `${var.object_id}`},
		"ref_name":      Representation{RepType: Optional, Create: `refName`},
		"ref_type":      Representation{RepType: Optional, Create: `TAG`},
		"filter":        RepresentationGroup{Required, repositoryTagRefDataSourceFilterRepresentation}}
	repositoryTagRefDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_repository_ref.test_repository_ref.id}`}},
	}

	repositoryTagRefRepresentation = map[string]interface{}{
		"ref_name":      Representation{RepType: Required, Create: `refName`},
		"ref_type":      Representation{RepType: Required, Create: `TAG`},
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"object_id":     Representation{RepType: Required, Create: `${var.commit_id}`},
	}

	RepositoryTagRefResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsRepositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryTagRefResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryTagRefResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	objectId := getEnvSettingWithBlankDefault("object_id")
	objectIdStr := fmt.Sprintf("variable \"object_id\" { default = \"%s\" }\n", objectId)

	resourceName := "oci_devops_repository_ref.test_repository_ref"
	datasourceName := "data.oci_devops_repository_refs.test_repository_refs"
	singularDatasourceName := "data.oci_devops_repository_ref.test_repository_ref"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+objectIdStr+RepositoryTagRefResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", Required, Create, repositoryTagRefRepresentation), "devops", "repositoryRef", t)

	ResourceTest(t, testAccCheckDevopsRepositoryTagRefDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + objectIdStr + RepositoryTagRefResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", Required, Create, repositoryTagRefRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "object_id"),
				resource.TestCheckResourceAttr(resourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttr(resourceName, "ref_type", "TAG"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

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
			Config: config + compartmentIdVariableStr + objectIdStr + RepositoryTagRefResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", Optional, Update, repositoryTagRefRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "object_id"),
				resource.TestCheckResourceAttrSet(resourceName, "full_ref_name"),
				resource.TestCheckResourceAttr(resourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttr(resourceName, "ref_type", "TAG"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_refs", "test_repository_refs", Optional, Update, repositoryTagRefDataSourceRepresentation) +
				compartmentIdVariableStr + objectIdStr + RepositoryTagRefResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", Optional, Update, repositoryTagRefRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "object_id"),
				resource.TestCheckResourceAttr(datasourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttr(datasourceName, "ref_type", "TAG"),
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),

				resource.TestCheckResourceAttr(datasourceName, "repository_ref_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "repository_ref_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_ref", "test_repository_ref", Required, Create, repositoryTagRefSingularDataSourceRepresentation) +
				compartmentIdVariableStr + objectIdStr + RepositoryTagRefResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "full_ref_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_type", "TAG"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + objectIdStr + RepositoryTagRefResourceConfig,
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

func testAccCheckDevopsRepositoryTagRefDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
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

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DevopsRepositoryTagRef") {
		resource.AddTestSweepers("DevopsRepositoryTagRef", &resource.Sweeper{
			Name:         "DevopsRepositoryTagRef",
			Dependencies: DependencyGraph["repositoryRef"],
			F:            sweepDevopsRepositoryTagRefResource,
		})
	}
}

func sweepDevopsRepositoryTagRefResource(compartment string) error {
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	repositoryRefIds, err := getRepositoryTagRefIds(compartment)
	if err != nil {
		return err
	}
	for _, repositoryRefId := range repositoryRefIds {
		if ok := SweeperDefaultResourceId[repositoryRefId]; !ok {
			deleteRefRequest := oci_devops.DeleteRefRequest{}
			deleteRefRequest.RefName = &repositoryRefId

			deleteRefRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteRef(context.Background(), deleteRefRequest)
			if error != nil {
				fmt.Printf("Error deleting RepositoryRef %s %s, It is possible that the resource is already deleted. Please verify manually \n", repositoryRefId, error)
				continue
			}
		}
	}
	return nil
}

func getRepositoryTagRefIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "RepositoryTagRefId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()

	listRefsRequest := oci_devops.ListRefsRequest{}

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
			id := *repositoryRef.GetFullRefName()
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "RepositoryTagRefId", id)
		}

	}
	return resourceIds, nil
}
