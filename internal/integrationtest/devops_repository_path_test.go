// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	repositoryPathSingularDataSourceRepresentation = map[string]interface{}{
		"repository_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"folder_path":      acctest.Representation{RepType: acctest.Optional, Create: `folderPath`},
		"paths_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"ref":              acctest.Representation{RepType: acctest.Optional, Create: `ref`},
	}

	repositoryPathDataSourceRepresentation = map[string]interface{}{
		"repository_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"folder_path":      acctest.Representation{RepType: acctest.Optional, Create: `folderPath`},
		"paths_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"ref":              acctest.Representation{RepType: acctest.Optional, Create: `ref`},
	}

	RepositoryPathResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, devopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryPathResource_basic(t *testing.T) {
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "RepositoryPath") {
		t.Skip("TestDevopsRepositoryPathResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryPathResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_devops_repository_paths.test_repository_paths"
	singularDatasourceName := "data.oci_devops_repository_path.test_repository_path"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_paths", "test_repository_paths", acctest.Required, acctest.Create, repositoryPathDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryPathResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "folder_path", "folderPath"),
				resource.TestCheckResourceAttr(datasourceName, "paths_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "ref", "ref"),
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "repository_path_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "repository_path_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_path", "test_repository_path", acctest.Required, acctest.Create, repositoryPathSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryPathResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "folder_path", "folderPath"),
				resource.TestCheckResourceAttr(singularDatasourceName, "paths_in_subtree", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref", "ref"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
	})
}
