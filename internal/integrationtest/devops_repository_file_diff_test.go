// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsDevopsRepositoryFileDiffSingularDataSourceRepresentation = map[string]interface{}{
		"base_version":                  acctest.Representation{RepType: acctest.Required, Create: `main`},
		"repository_id":                 acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("repository_id_for_static_resource")},
		"target_version":                acctest.Representation{RepType: acctest.Required, Create: `main2`},
		"file_path":                     acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("test_file_for_static_resource")},
		"is_comparison_from_merge_base": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryFileDiffResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryFileDiffResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id_for_static_resource")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_devops_repository_file_diff.test_repository_file_diff"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_file_diff", "test_repository_file_diff", acctest.Required, acctest.Create, DevopsDevopsRepositoryFileDiffSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "base_version", "main"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_version", "main2"),

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
