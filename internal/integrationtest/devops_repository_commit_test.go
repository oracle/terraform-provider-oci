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
	DevopsDevopsRepositoryCommitSingularDataSourceRepresentation = map[string]interface{}{
		"commit_id":     acctest.Representation{RepType: acctest.Required, Create: `main`},
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	DevopsDevopsRepositoryCommitDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	DevopsRepositoryCommitResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryCommitResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryCommitResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_devops_repository_commits.test_repository_commits"
	singularDatasourceName := "data.oci_devops_repository_commit.test_repository_commit"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_commits", "test_repository_commits", acctest.Required, acctest.Create, DevopsDevopsRepositoryCommitDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsRepositoryCommitResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "repository_commit_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_commit", "test_repository_commit", acctest.Required, acctest.Create, DevopsDevopsRepositoryCommitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsRepositoryCommitResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
