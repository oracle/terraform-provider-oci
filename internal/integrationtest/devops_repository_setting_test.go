// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DevopsRepositorySettingRequiredOnlyResource = DevopsRepositorySettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_setting", "test_repository_setting", acctest.Required, acctest.Create, DevopsRepositorySettingRepresentation)

	DevopsRepositorySettingResourceConfig = DevopsRepositorySettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_setting", "test_repository_setting", acctest.Optional, acctest.Update, DevopsRepositorySettingRepresentation)

	DevopsRepositorySettingSingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	DevopsRepositorySettingRepresentation = map[string]interface{}{
		"repository_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"approval_rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsRepositorySettingApprovalRulesRepresentation},
		"merge_checks":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsRepositorySettingMergeChecksRepresentation},
		"merge_settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsRepositorySettingMergeSettingsRepresentation},
	}
	DevopsRepositorySettingApprovalRulesRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsRepositorySettingApprovalRulesItemsRepresentation},
	}
	DevopsRepositorySettingMergeChecksRepresentation = map[string]interface{}{
		"last_build_succeeded": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `DISABLED`},
	}
	DevopsRepositorySettingMergeSettingsRepresentation = map[string]interface{}{
		"allowed_merge_strategies": acctest.Representation{RepType: acctest.Required, Create: []string{`MERGE_COMMIT`}, Update: []string{`FAST_FORWARD`}},
		"default_merge_strategy":   acctest.Representation{RepType: acctest.Required, Create: `MERGE_COMMIT`, Update: `FAST_FORWARD`},
	}
	DevopsRepositorySettingApprovalRulesItemsRepresentation = map[string]interface{}{
		"min_approvals_count": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"destination_branch":  acctest.Representation{RepType: acctest.Optional, Create: `main`, Update: `main`},
		"reviewers":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsRepositorySettingApprovalRulesItemsReviewersRepresentation},
	}
	DevopsRepositorySettingApprovalRulesItemsReviewersRepresentation = map[string]interface{}{
		"principal_id": acctest.Representation{RepType: acctest.Required, Create: `${var.principal_ocid}`},
	}

	DevopsRepositorySettingResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositorySettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositorySettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	principalId := utils.GetEnvSettingWithBlankDefault("principal_ocid")
	principalIdVariableStr := fmt.Sprintf("variable \"principal_ocid\" { default = \"%s\" }\n", principalId)

	resourceName := "oci_devops_repository_setting.test_repository_setting"

	singularDatasourceName := "data.oci_devops_repository_setting.test_repository_setting"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsRepositorySettingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_setting", "test_repository_setting", acctest.Optional, acctest.Create, DevopsRepositorySettingRepresentation), "devops", "repositorySetting", t)

	acctest.ResourceTest(t, testAccCheckDevopsRepositorySettingDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsRepositorySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_setting", "test_repository_setting", acctest.Required, acctest.Create, DevopsRepositorySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsRepositorySettingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + principalIdVariableStr + DevopsRepositorySettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_setting", "test_repository_setting", acctest.Optional, acctest.Create, DevopsRepositorySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approval_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.0.destination_branch", "main"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.0.min_approvals_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.0.reviewers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_rules.0.items.0.reviewers.0.principal_id"),
				resource.TestCheckResourceAttr(resourceName, "merge_checks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "merge_checks.0.last_build_succeeded", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "merge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "merge_settings.0.allowed_merge_strategies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "merge_settings.0.default_merge_strategy", "MERGE_COMMIT"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DevopsRepositorySettingResourceDependencies + principalIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_setting", "test_repository_setting", acctest.Optional, acctest.Update, DevopsRepositorySettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approval_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.0.destination_branch", "main"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.0.min_approvals_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "approval_rules.0.items.0.reviewers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_rules.0.items.0.reviewers.0.principal_id"),
				resource.TestCheckResourceAttr(resourceName, "merge_checks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "merge_checks.0.last_build_succeeded", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "merge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "merge_settings.0.allowed_merge_strategies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "merge_settings.0.default_merge_strategy", "FAST_FORWARD"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_setting", "test_repository_setting", acctest.Required, acctest.Create, DevopsRepositorySettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + principalIdVariableStr + DevopsRepositorySettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "approval_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approval_rules.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approval_rules.0.items.0.destination_branch", "main"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approval_rules.0.items.0.min_approvals_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approval_rules.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approval_rules.0.items.0.reviewers.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approval_rules.0.items.0.reviewers.0.principal_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approval_rules.0.items.0.reviewers.0.principal_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approval_rules.0.items.0.reviewers.0.principal_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "merge_checks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "merge_checks.0.last_build_succeeded", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "merge_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "merge_settings.0.allowed_merge_strategies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "merge_settings.0.default_merge_strategy", "FAST_FORWARD"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsRepositorySettingRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDevopsRepositorySettingDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_repository_setting" {
			noResourceFound = false
			request := oci_devops.GetRepositorySettingsRequest{}

			if value, ok := rs.Primary.Attributes["repository_id"]; ok {
				request.RepositoryId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

			_, err := client.GetRepositorySettings(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
