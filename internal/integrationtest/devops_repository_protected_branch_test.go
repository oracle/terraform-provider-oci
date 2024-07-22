// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DevopsRepositoryProtectedBranchDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	DevopsRepositoryProtectedBranchResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryProtectedBranchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryProtectedBranchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_devops_repository_protected_branch_management.test_repository_protected_branch_management"
	datasourceName := "data.oci_devops_repository_protected_branches.test_repository_protected_branches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsRepositoryProtectedBranchManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_protected_branch_management", "test_repository_protected_branch_management", acctest.Required, acctest.Create, DevopsRepositoryProtectedBranchManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "branch_name", "refs/heads/main"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_protected_branches", "test_repository_protected_branches", acctest.Required, acctest.Create, DevopsRepositoryProtectedBranchDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsRepositoryProtectedBranchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "protected_branch_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "protected_branch_collection.0.items.#", "1"),
			),
		},
	})
}
