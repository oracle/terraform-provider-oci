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

	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DevopsCodeRepositoryTriggerRequiredOnlyResource = DevopsCodeRepositoryTriggerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Required, acctest.Create, devopsCodeRepositoryTriggerRepresentation)

	DevopsCodeRepositoryTriggerResourceConfig = DevopsCodeRepositoryTriggerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Update, devopsCodeRepositoryTriggerRepresentation)

	devopsCodeRepositoryTriggerSingularDataSourceRepresentation = map[string]interface{}{
		"trigger_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_trigger.test_trigger.id}`},
	}

	devopsCodeRepositoryTriggerDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_trigger.test_trigger.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: devopsCodeRepositoryTriggerDataSourceFilterRepresentation}}
	devopsCodeRepositoryTriggerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_trigger.test_trigger.id}`}},
	}

	devopsCodeRepositoryTriggerRepresentation = map[string]interface{}{
		"actions":        acctest.RepresentationGroup{RepType: acctest.Required, Group: devopsCodeRepositoryTriggerActionsRepresentation},
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"trigger_source": acctest.Representation{RepType: acctest.Required, Create: `DEVOPS_CODE_REPOSITORY`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"repository_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	devopsCodeRepositoryTriggerActionsRepresentation = map[string]interface{}{
		"build_pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"type":              acctest.Representation{RepType: acctest.Required, Create: `TRIGGER_BUILD_PIPELINE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: devopsCodeRepositoryTriggerActionsFilterRepresentation},
	}
	devopsCodeRepositoryTriggerActionsFilterRepresentation = map[string]interface{}{
		"trigger_source": acctest.Representation{RepType: acctest.Required, Create: `DEVOPS_CODE_REPOSITORY`},
		"events":         acctest.Representation{RepType: acctest.Optional, Create: []string{`PUSH`}},
		"include":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: devopsCodeRepositoryTriggerActionsFilterIncludeRepresentation},
	}
	devopsCodeRepositoryTriggerActionsFilterIncludeRepresentation = map[string]interface{}{
		"head_ref": acctest.Representation{RepType: acctest.Optional, Create: `master`, Update: `master2`},
	}

	devopsCodeRepositoryRepresentation = map[string]interface{}{
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"project_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"default_branch":           acctest.Representation{RepType: acctest.Optional, Create: `defaultBranch`, Update: `defaultBranch2`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"mirror_repository_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: repositoryMirrorRepositoryConfigRepresentation},
		"repository_type":          acctest.Representation{RepType: acctest.Required, Create: `HOSTED`, Update: `MIRRORED`},
	}

	repositoryMirrorRepositoryConfigRepresentation = map[string]interface{}{
		"trigger_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: repositoryMirrorRepositoryConfigTriggerScheduleRepresentation},
	}
	repositoryMirrorRepositoryConfigTriggerScheduleRepresentation = map[string]interface{}{
		"schedule_type":   acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `DEFAULT`},
		"custom_schedule": acctest.Representation{RepType: acctest.Optional, Create: `customSchedule`, Update: `customSchedule2`},
	}

	DevopsCodeRepositoryTriggerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, devopsCodeRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Required, acctest.Create, buildPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsCodeRepositoryTriggerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsCodeRepositoryTriggerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_trigger.test_trigger"
	datasourceName := "data.oci_devops_triggers.test_triggers"
	singularDatasourceName := "data.oci_devops_trigger.test_trigger"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsCodeRepositoryTriggerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Create, devopsCodeRepositoryTriggerRepresentation), "devops", "trigger", t)

	acctest.ResourceTest(t, testAccCheckDevopsTriggerDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + DevopsCodeRepositoryTriggerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Required, acctest.Create, devopsCodeRepositoryTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger_source", "DEVOPS_CODE_REPOSITORY"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + DevopsCodeRepositoryTriggerResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsCodeRepositoryTriggerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Create, devopsCodeRepositoryTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.head_ref", "master"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.trigger_source", "DEVOPS_CODE_REPOSITORY"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger_source", "DEVOPS_CODE_REPOSITORY"),

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

		//verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DevopsCodeRepositoryTriggerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Update, devopsCodeRepositoryTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "actions.0.build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.filter.0.include.0.head_ref", "master2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_triggers", "test_triggers", acctest.Optional, acctest.Update, devopsCodeRepositoryTriggerDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsCodeRepositoryTriggerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Optional, acctest.Update, devopsCodeRepositoryTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "trigger_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_trigger", "test_trigger", acctest.Required, acctest.Create, devopsCodeRepositoryTriggerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsCodeRepositoryTriggerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trigger_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.events.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.include.0.head_ref", "master2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.filter.0.trigger_source", "DEVOPS_CODE_REPOSITORY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.type", "TRIGGER_BUILD_PIPELINE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trigger_source", "DEVOPS_CODE_REPOSITORY"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DevopsCodeRepositoryTriggerResourceConfig,
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
