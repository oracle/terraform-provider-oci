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
	repositoryCommitSingularDataSourceRepresentation = map[string]interface{}{
		"commit_id":     Representation{RepType: Required, Create: `main`},
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	repositoryCommitDataSourceRepresentation = map[string]interface{}{
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	RepositoryCommitResourceConfig = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsRepositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryCommitResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryCommitResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_devops_repository_commits.test_repository_commits"
	singularDatasourceName := "data.oci_devops_repository_commit.test_repository_commit"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_commits", "test_repository_commits", Required, Create, repositoryCommitDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryCommitResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "repository_commit_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_commit", "test_repository_commit", Required, Create, repositoryCommitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RepositoryCommitResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "commit_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "author_email"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "author_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "commit_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "commit_message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "committer_email"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "committer_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tree_id"),
			),
		},
	})
}
