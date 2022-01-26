// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BuildRunRequiredOnlyResource = BuildRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_run", "test_build_run", acctest.Required, acctest.Create, buildRunRepresentation)

	BuildRunResourceConfig = BuildRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_run", "test_build_run", acctest.Optional, acctest.Update, buildRunRepresentation)

	buildRunSingularDataSourceRepresentation = map[string]interface{}{
		"build_run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_run.test_build_run.id}`},
	}

	buildRunDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_build_run.test_build_run.id}`},
		"project_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: buildRunDataSourceFilterRepresentation}}
	buildRunDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_build_run.test_build_run.id}`}},
	}

	buildRunRepresentation = map[string]interface{}{
		"build_pipeline_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"build_run_arguments": acctest.RepresentationGroup{RepType: acctest.Optional, Group: buildRunBuildRunArgumentsRepresentation},
		"commit_info":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: buildRunCommitInfoRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	buildRunBuildRunArgumentsRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildRunBuildRunArgumentsItemsRepresentation},
	}
	buildRunCommitInfoRepresentation = map[string]interface{}{
		"commit_hash":       acctest.Representation{RepType: acctest.Required, Create: `commitHash`},
		"repository_branch": acctest.Representation{RepType: acctest.Required, Create: `repositoryBranch`},
		"repository_url":    acctest.Representation{RepType: acctest.Required, Create: `repositoryUrl`},
	}
	buildRunBuildRunArgumentsItemsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	BuildRunResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Required, acctest.Create, buildPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsBuildRunResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsBuildRunResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_build_run.test_build_run"
	datasourceName := "data.oci_devops_build_runs.test_build_runs"
	singularDatasourceName := "data.oci_devops_build_run.test_build_run"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BuildRunResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_run", "test_build_run", acctest.Optional, acctest.Create, buildRunRepresentation), "devops", "buildRun", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BuildRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_run", "test_build_run", acctest.Required, acctest.Create, buildRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BuildRunResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BuildRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_run", "test_build_run", acctest.Optional, acctest.Create, buildRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_run_arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_run_arguments.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_run_arguments.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "build_run_arguments.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "build_run_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "commit_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "commit_info.0.commit_hash", "commitHash"),
				resource.TestCheckResourceAttr(resourceName, "commit_info.0.repository_branch", "repositoryBranch"),
				resource.TestCheckResourceAttr(resourceName, "commit_info.0.repository_url", "repositoryUrl"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + BuildRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_run", "test_build_run", acctest.Optional, acctest.Update, buildRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_run_arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_run_arguments.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_run_arguments.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "build_run_arguments.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "build_run_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "commit_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "commit_info.0.commit_hash", "commitHash"),
				resource.TestCheckResourceAttr(resourceName, "commit_info.0.repository_branch", "repositoryBranch"),
				resource.TestCheckResourceAttr(resourceName, "commit_info.0.repository_url", "repositoryUrl"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_build_runs", "test_build_runs", acctest.Optional, acctest.Update, buildRunDataSourceRepresentation) +
				compartmentIdVariableStr + BuildRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_run", "test_build_run", acctest.Optional, acctest.Update, buildRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

				resource.TestCheckResourceAttr(datasourceName, "build_run_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "build_run_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_build_run", "test_build_run", acctest.Required, acctest.Create, buildRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BuildRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "build_run_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "build_outputs.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_run_arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_run_arguments.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_run_arguments.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_run_arguments.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_run_progress.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_run_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "commit_info.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "commit_info.0.commit_hash", "commitHash"),
				resource.TestCheckResourceAttr(singularDatasourceName, "commit_info.0.repository_branch", "repositoryBranch"),
				resource.TestCheckResourceAttr(singularDatasourceName, "commit_info.0.repository_url", "repositoryUrl"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BuildRunResourceConfig,
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
