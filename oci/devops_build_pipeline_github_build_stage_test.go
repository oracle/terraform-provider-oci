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
	BuildPipelineBuildStageRequiredOnlyResource = BuildPipelineBuildStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineBuildStageRepresentation)

	BuildPipelineBuildStageResourceConfig = BuildPipelineBuildStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineBuildStageRepresentation)

	buildPipelineBuildStageSingularDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_stage_id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
	}

	buildPipelineBuildStageDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_id": Representation{RepType: Optional, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"compartment_id":    Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"display_name":      Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":                Representation{RepType: Optional, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
		"state":             Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":            RepresentationGroup{Required, buildPipelineBuildStageDataSourceFilterRepresentation}}
	buildPipelineBuildStageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`}},
	}

	buildPipelineBuildStageRepresentation = map[string]interface{}{
		"build_pipeline_id":                           Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"build_pipeline_stage_predecessor_collection": RepresentationGroup{Required, buildPipelineBuildStageBuildPipelineStagePredecessorCollectionRepresentation},
		"build_pipeline_stage_type":                   Representation{RepType: Required, Create: `BUILD`},
		"defined_tags":                                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                                 Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":                                Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                               Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"build_spec_file":                             Representation{RepType: Required, Create: `buildSpecFile`, Update: `buildSpecFile2`},
		"image":                                       Representation{RepType: Required, Create: `OL7_X86_64_STANDARD_10`},
		"primary_build_source":                        Representation{RepType: Required, Create: `primaryBuildSource`, Update: `primaryBuildSource2`},
		"build_source_collection":                     RepresentationGroup{Required, buildPipelineStageBuildSourceCollectionRepresentation},
		"stage_execution_timeout_in_seconds":          Representation{RepType: Optional, Create: `10`, Update: `11`},
	}
	buildPipelineBuildStageBuildPipelineStagePredecessorCollectionRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, buildPipelineBuildStageBuildPipelineStagePredecessorCollectionItemsRepresentation},
	}

	buildPipelineBuildStageBuildPipelineStagePredecessorCollectionItemsRepresentation = map[string]interface{}{
		"id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
	}

	buildPipelineStageBuildSourceCollectionRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, buildPipelineStageBuildSourceCollectionItemsRepresentation},
	}

	buildPipelineStageBuildSourceCollectionItemsRepresentation = map[string]interface{}{
		"connection_type": Representation{RepType: Required, Create: `GITHUB`},
		"branch":          Representation{RepType: Required, Create: `branch`, Update: `branch2`},
		"connection_id":   Representation{RepType: Required, Create: `${oci_devops_connection.test_connection.id}`},
		"name":            Representation{RepType: Required, Create: `primaryBuildSource`, Update: `primaryBuildSource2`},
		"repository_url":  Representation{RepType: Required, Create: `repositoryUrl`, Update: `repositoryUrl2`},
	}

	githubAccessTokenVaultId    = getEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr = fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	BuildPipelineBuildStageResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Required, Create, buildPipelineRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation) +
		githubAccessTokenVaultIdStr +
		GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsConnectionRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsBuildPipelineBuildStageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsBuildPipelineBuildStageResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_build_pipeline_stage.test_build_pipeline_stage"
	datasourceName := "data.oci_devops_build_pipeline_stages.test_build_pipeline_stages"
	singularDatasourceName := "data.oci_devops_build_pipeline_stage.test_build_pipeline_stage"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+BuildPipelineBuildStageResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Create, buildPipelineBuildStageRepresentation), "devops", "buildPipelineStage", t)

	conf := config + compartmentIdVariableStr + BuildPipelineBuildStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Create, buildPipelineBuildStageRepresentation)
	print(conf)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDevopsBuildPipelineStageDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineBuildStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "BUILD"),

					resource.TestCheckResourceAttr(resourceName, "build_spec_file", "buildSpecFile"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.branch", "branch"),
					resource.TestCheckResourceAttrSet(resourceName, "build_source_collection.0.items.0.connection_id"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.connection_type", "GITHUB"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.name", "primaryBuildSource"),
					resource.TestCheckResourceAttr(resourceName, "primary_build_source", "primaryBuildSource"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.repository_url", "repositoryUrl"),
					resource.TestCheckResourceAttr(resourceName, "image", "OL7_X86_64_STANDARD_10"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						print("// verify create" + resId)
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Create, buildPipelineBuildStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "BUILD"),

					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "BUILD"),

					resource.TestCheckResourceAttr(resourceName, "build_source_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.branch", "branch"),
					resource.TestCheckResourceAttrSet(resourceName, "build_source_collection.0.items.0.connection_id"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.connection_type", "GITHUB"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.name", "primaryBuildSource"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.repository_url", "repositoryUrl"),
					resource.TestCheckResourceAttr(resourceName, "image", "OL7_X86_64_STANDARD_10"),
					resource.TestCheckResourceAttr(resourceName, "stage_execution_timeout_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "build_spec_file", "buildSpecFile"),
					func(s *terraform.State) (err error) {
						resId, _ = FromInstanceState(s, resourceName, "id")
						print("// verify create optionals" + resId)
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineBuildStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "BUILD"),

					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttr(resourceName, "stage_execution_timeout_in_seconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.branch", "branch2"),
					resource.TestCheckResourceAttrSet(resourceName, "build_source_collection.0.items.0.connection_id"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.connection_type", "GITHUB"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.name", "primaryBuildSource2"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.repository_url", "repositoryUrl2"),
					resource.TestCheckResourceAttr(resourceName, "image", "OL7_X86_64_STANDARD_10"),
					resource.TestCheckResourceAttr(resourceName, "build_spec_file", "buildSpecFile2"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						print("// verify update" + resId2)
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
					GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stages", "test_build_pipeline_stages", Optional, Update, buildPipelineBuildStageDataSourceRepresentation) +
					compartmentIdVariableStr + BuildPipelineBuildStageResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineBuildStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "build_pipeline_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "build_pipeline_stage_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineBuildStageSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BuildPipelineBuildStageResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "build_pipeline_stage_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_type", "BUILD"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.0.branch", "branch2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.0.connection_type", "GITHUB"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.0.name", "primaryBuildSource2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.0.repository_url", "repositoryUrl2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_spec_file", "buildSpecFile2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "image", "OL7_X86_64_STANDARD_10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "primary_build_source", "primaryBuildSource2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "stage_execution_timeout_in_seconds", "11"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					func(s *terraform.State) (err error) {
						return nil
					},
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
