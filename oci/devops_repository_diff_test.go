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
	repositoryDiffSingularDataSourceRepresentation = map[string]interface{}{
		"base_version":   Representation{RepType: Required, Create: `main`},
		"file_path":      Representation{RepType: Required, Create: ``},
		"repository_id":  Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"target_version": Representation{RepType: Required, Create: `main`},
	}

	repositoryDiffDataSourceRepresentation = map[string]interface{}{
		"base_version":   Representation{RepType: Required, Create: `main`},
		"repository_id":  Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"target_version": Representation{RepType: Required, Create: `main`},
	}

	RepositoryDiffResourceConfig = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsRepositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryDiffResource_basic(t *testing.T) {
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "RepositoryDiff") {
		t.Skip("TestDevopsRepositoryDiffResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryDiffResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_devops_repository_diffs.test_repository_diffs"
	singularDatasourceName := "data.oci_devops_repository_diff.test_repository_diff"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_diffs", "test_repository_diffs", Required, Create, repositoryDiffDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryDiffResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "base_version", "baseVersion"),
				resource.TestCheckResourceAttr(datasourceName, "is_comparison_from_merge_base", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),
				resource.TestCheckResourceAttr(datasourceName, "target_version", "targetVersion"),

				resource.TestCheckResourceAttrSet(datasourceName, "diff_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_diff", "test_repository_diff", Required, Create, repositoryDiffSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryDiffResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "base_version", "baseVersion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_path", "filePath"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_comparison_from_merge_base", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_version", "targetVersion"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "are_conflicts_in_file"),
				resource.TestCheckResourceAttr(singularDatasourceName, "changes.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_binary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_large"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "new_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "new_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "old_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "old_path"),
			),
		},
	})
}
