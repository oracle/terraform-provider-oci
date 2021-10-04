// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	repositoryPathSingularDataSourceRepresentation = map[string]interface{}{
		"repository_id":    Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"display_name":     Representation{RepType: Optional, Create: `displayName`},
		"folder_path":      Representation{RepType: Optional, Create: `folderPath`},
		"paths_in_subtree": Representation{RepType: Optional, Create: `false`},
		"ref":              Representation{RepType: Optional, Create: `ref`},
	}

	repositoryPathDataSourceRepresentation = map[string]interface{}{
		"repository_id":    Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"display_name":     Representation{RepType: Optional, Create: `displayName`},
		"folder_path":      Representation{RepType: Optional, Create: `folderPath`},
		"paths_in_subtree": Representation{RepType: Optional, Create: `false`},
		"ref":              Representation{RepType: Optional, Create: `ref`},
	}

	RepositoryPathResourceConfig = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsRepositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryPathResource_basic(t *testing.T) {
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "RepositoryPath") {
		t.Skip("TestDevopsRepositoryPathResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryPathResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_devops_repository_paths.test_repository_paths"
	singularDatasourceName := "data.oci_devops_repository_path.test_repository_path"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_paths", "test_repository_paths", Required, Create, repositoryPathDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryPathResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_path", "test_repository_path", Required, Create, repositoryPathSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryPathResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
