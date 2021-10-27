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
	repositoryFileLineSingularDataSourceRepresentation = map[string]interface{}{
		"file_path":     Representation{RepType: Required, Create: `filePath`},
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"revision":      Representation{RepType: Required, Create: `revision`},
	}

	RepositoryFileLineResourceConfig = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, projectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, repositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryFileLineResource_basic(t *testing.T) {
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "RepositoryFileLine") {
		t.Skip("TestDevopsRepositoryFileLineResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryFileLineResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_devops_repository_file_line.test_repository_file_line"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_file_line", "test_repository_file_line", Required, Create, repositoryFileLineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryFileLineResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "file_path", "filePath"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "revision", "revision"),
				resource.TestCheckResourceAttr(singularDatasourceName, "start_line_number", "10"),

				resource.TestCheckResourceAttr(singularDatasourceName, "lines.#", "1"),
			),
		},
	})
}
