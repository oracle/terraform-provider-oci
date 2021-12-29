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
	repositoryDiffSingularDataSourceRepresentation = map[string]interface{}{
		"base_version":   acctest.Representation{RepType: acctest.Required, Create: `main`},
		"file_path":      acctest.Representation{RepType: acctest.Required, Create: ``},
		"repository_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"target_version": acctest.Representation{RepType: acctest.Required, Create: `main`},
	}

	repositoryDiffDataSourceRepresentation = map[string]interface{}{
		"base_version":   acctest.Representation{RepType: acctest.Required, Create: `main`},
		"repository_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"target_version": acctest.Representation{RepType: acctest.Required, Create: `main`},
	}

	RepositoryDiffResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, devopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryDiffResource_basic(t *testing.T) {
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "RepositoryDiff") {
		t.Skip("TestDevopsRepositoryDiffResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryDiffResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_devops_repository_diffs.test_repository_diffs"
	singularDatasourceName := "data.oci_devops_repository_diff.test_repository_diff"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_diffs", "test_repository_diffs", acctest.Required, acctest.Create, repositoryDiffDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryDiffResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_diff", "test_repository_diff", acctest.Required, acctest.Create, repositoryDiffSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryDiffResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
