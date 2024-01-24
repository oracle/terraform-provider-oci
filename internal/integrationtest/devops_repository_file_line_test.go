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
	repositoryFileLineSingularDataSourceRepresentation = map[string]interface{}{
		"file_path":     acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("test_file_for_static_resource")},
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("repository_id_for_static_resource")},
		"revision":      acctest.Representation{RepType: acctest.Required, Create: `main`},
	}
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryFileLineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryFileLineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id_for_static_resource")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_devops_repository_file_line.test_repository_file_line"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_file_line", "test_repository_file_line", acctest.Required, acctest.Create, repositoryFileLineSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "file_path", "testfile"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "revision", "main"),

				resource.TestCheckResourceAttr(singularDatasourceName, "lines.#", "2"),
			),
		},
	})
}
