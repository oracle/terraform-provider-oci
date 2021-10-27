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
	repositoryObjectSingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	RepositoryObjectResourceConfig = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsRepositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryObjectResource_basic(t *testing.T) {
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "RepositoryObject") {
		t.Skip("TestDevopsRepositoryObjectResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_devops_repository_object.test_repository_object"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_object", "test_repository_object", Required, Create, repositoryObjectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryObjectResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "file_path", "filePath"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_binary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sha"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
	})
}
