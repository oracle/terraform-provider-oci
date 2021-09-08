// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeployManualApprovalStageRequiredOnlyResource = DeployStageResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployManualApprovalStageRepresentation)

	DeployManualApprovalStageResourceConfig = DeployStageResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployManualApprovalStageRepresentation)

	deployManualApprovalStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": Representation{repType: Required, create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployManualApprovalStageRepresentation = getUpdatedRepresentationCopy("deploy_stage_type", Representation{repType: Required, create: `MANUAL_APPROVAL`},
		representationCopyWithNewProperties(representationCopyWithRemovedProperties(deployStageRepresentation, []string{"wait_criteria"}), map[string]interface{}{
			"approval_policy": RepresentationGroup{Required, deployStageApprovalPolicyRepresentation},
		}))

	deployStageApprovalPolicyRepresentation = map[string]interface{}{
		"approval_policy_type":         Representation{repType: Required, create: `COUNT_BASED_APPROVAL`},
		"number_of_approvals_required": Representation{repType: Required, create: `1`, update: `2`},
	}
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_ManualApproval(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_ManualApproval")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DeployStageResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployManualApprovalStageRepresentation), "devops", "deployStage", t)

	ResourceTest(t, testAccCheckDevopsDeployStageDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + DeployStageResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployManualApprovalStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "MANUAL_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.approval_policy_type", "COUNT_BASED_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.number_of_approvals_required", "1"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + DeployStageResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployStageResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployManualApprovalStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "MANUAL_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.approval_policy_type", "COUNT_BASED_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.number_of_approvals_required", "1"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DeployStageResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployManualApprovalStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.approval_policy_type", "COUNT_BASED_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.number_of_approvals_required", "2"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", Optional, Update, deployStageDataSourceRepresentation) +
				compartmentIdVariableStr + DeployStageResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployManualApprovalStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_stage_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployManualApprovalStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployManualApprovalStageResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "MANUAL_APPROVAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.approval_policy_type", "COUNT_BASED_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.number_of_approvals_required", "2"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DeployManualApprovalStageResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
