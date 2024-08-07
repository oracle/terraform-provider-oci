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
	DevopsRepositoryProtectedBranchManagementRequiredOnlyResource = DevopsRepositoryProtectedBranchManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_protected_branch_management", "test_repository_protected_branch_management", acctest.Required, acctest.Create, DevopsRepositoryProtectedBranchManagementRepresentation)

	DevopsRepositoryProtectedBranchManagementRepresentation = map[string]interface{}{
		"branch_name":       acctest.Representation{RepType: acctest.Required, Create: `refs/heads/main`},
		"repository_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"protection_levels": acctest.Representation{RepType: acctest.Optional, Create: []string{`READ_ONLY`}, Update: []string{`PULL_REQUEST_MERGE_ONLY`}},
	}

	DevopsRepositoryProtectedBranchManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryProtectedBranchManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryProtectedBranchManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_repository_protected_branch_management.test_repository_protected_branch_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsRepositoryProtectedBranchManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_protected_branch_management", "test_repository_protected_branch_management", acctest.Required, acctest.Create, DevopsRepositoryProtectedBranchManagementRepresentation), "devops", "repositoryProtectedBranchManagement", t)

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsRepositoryProtectedBranchManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsRepositoryProtectedBranchManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_protected_branch_management", "test_repository_protected_branch_management", acctest.Optional, acctest.Create, DevopsRepositoryProtectedBranchManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "branch_name", "refs/heads/main"),
				resource.TestCheckResourceAttr(resourceName, "protection_levels.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),
			),
		},
		// verify update with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsRepositoryProtectedBranchManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_protected_branch_management", "test_repository_protected_branch_management", acctest.Optional, acctest.Update, DevopsRepositoryProtectedBranchManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "branch_name", "refs/heads/main"),
				resource.TestCheckResourceAttr(resourceName, "protection_levels.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),
			),
		},
	})
}
