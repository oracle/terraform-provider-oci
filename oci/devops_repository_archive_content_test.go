// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	repositoryArchiveContentSingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
		"format":        Representation{RepType: Optional, Create: `zip`},
		"ref_name":      Representation{RepType: Optional, Create: `main`},
	}

	RepositoryArchiveContentResourceConfig = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsRepositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryArchiveContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryArchiveContentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_devops_repository_archive_content.test_repository_archive_content"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_archive_content", "test_repository_archive_content", Required, Create, repositoryArchiveContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryArchiveContentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "format", "zip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_name", "main"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
			),
		},
	})
}
