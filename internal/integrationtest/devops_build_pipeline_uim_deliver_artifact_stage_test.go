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
	BuildPipelineStageUIMDeliverArtifactRequiredOnlyResource = BuildPipelineStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Required, acctest.Create, buildPipelineStageUIMDeliverArtifactRepresentation)

	BuildPipelineStageUIMDeliverArtifactResourceConfig = BuildPipelineStageUIMDeliverArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Update, buildPipelineStageUIMDeliverArtifactRepresentation)

	buildPipelineStageUIMDeliverArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
	}

	buildPipelineStageUIMDeliverArtifactDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageUIMDeliverArtifactDataSourceFilterRepresentation}}
	buildPipelineStageUIMDeliverArtifactDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`}},
	}

	// TODO: stage timeout is missing in alll the steps
	buildPipelineStageUIMDeliverArtifactRepresentation = map[string]interface{}{
		"build_pipeline_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"build_pipeline_stage_predecessor_collection": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionRepresentation},
		"build_pipeline_stage_type":                   acctest.Representation{RepType: acctest.Required, Create: `DELIVER_ARTIFACT`},
		"defined_tags":                                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"deliver_artifact_collection":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageDeliverArtifactCollectionRepresentation},
		"description":                                 acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	buildPipelineWaitStageRepresentationForDeliverArtifact = map[string]interface{}{
		"build_pipeline_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"build_pipeline_stage_predecessor_collection": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionRepresentationWait},
		"build_pipeline_stage_type":                   acctest.Representation{RepType: acctest.Required, Create: `WAIT`},
		"defined_tags":                                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                                 acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"wait_criteria":                               acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageWaitCriteriaRepresentation},
	}
	buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionItemsRepresentation},
	}

	buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionRepresentationWait = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionItemsRepresentationWait},
	}

	buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionItemsRepresentationWait = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
	}
	buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionItemsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage_wait.id}`},
	}

	buildPipelineStageDeliverArtifactCollectionRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageDeliverArtifactCollectionItemsRepresentation},
	}

	buildPipelineStageDeliverArtifactCollectionItemsRepresentation = map[string]interface{}{
		"artifact_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
		"artifact_name": acctest.Representation{RepType: acctest.Required, Create: `artifactName`, Update: `artifactName2`},
	}

	BuildPipelineStageUIMDeliverArtifactResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Required, acctest.Create, buildPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployGenericArtifactRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage_wait", acctest.Required, acctest.Create, buildPipelineWaitStageRepresentationForDeliverArtifact)
)

// issue-routing-tag: devops/default
func TestDevopsBuildPipelineStageUIMDeliverArtifactResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsBuildPipelineStageUIMDeliverArtifactResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_build_pipeline_stage.test_build_pipeline_stage"
	datasourceName := "data.oci_devops_build_pipeline_stages.test_build_pipeline_stages"
	singularDatasourceName := "data.oci_devops_build_pipeline_stage.test_build_pipeline_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BuildPipelineStageUIMDeliverArtifactResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Create, buildPipelineStageUIMDeliverArtifactRepresentation), "devops", "buildPipelineStage", t)

	acctest.ResourceTest(t, testAccCheckDevopsBuildPipelineStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Required, acctest.Create, buildPipelineStageUIMDeliverArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "DELIVER_ARTIFACT"),

				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deliver_artifact_collection.0.items.0.artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.0.artifact_name", "artifactName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Create, buildPipelineStageUIMDeliverArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "DELIVER_ARTIFACT"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deliver_artifact_collection.0.items.0.artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.0.artifact_name", "artifactName"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Update, buildPipelineStageUIMDeliverArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "DELIVER_ARTIFACT"),

				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deliver_artifact_collection.0.items.0.artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.0.artifact_name", "artifactName2"),

				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stages", "test_build_pipeline_stages", acctest.Optional, acctest.Update, buildPipelineStageUIMDeliverArtifactDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Update, buildPipelineStageUIMDeliverArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "build_pipeline_stage_collection.#", "1"),
			),
		},
		//verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Required, acctest.Create, buildPipelineStageUIMDeliverArtifactSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "build_pipeline_stage_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_type", "DELIVER_ARTIFACT"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deliver_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deliver_artifact_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deliver_artifact_collection.0.items.0.artifact_name", "artifactName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceConfig,
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
