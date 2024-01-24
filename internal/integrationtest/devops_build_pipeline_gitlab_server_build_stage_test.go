// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BuildPipelineBuildStageGitlabServerRequiredOnlyResource = BuildPipelineBuildStageGitlabServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Required, acctest.Create, buildPipelineBuildStageGitlabServerRepresentation)

	BuildPipelineBuildStageGitlabServerResourceConfig = BuildPipelineBuildStageGitlabServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Update, buildPipelineBuildStageGitlabServerRepresentation)

	buildPipelineBuildStageGitlabServerSingularDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
	}

	buildPipelineBuildStageGitlabServerDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineBuildStageGitlabServerDataSourceFilterRepresentation}}
	buildPipelineBuildStageGitlabServerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`}},
	}

	buildPipelineBuildStageGitlabServerRepresentation = map[string]interface{}{
		"build_pipeline_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"build_pipeline_stage_predecessor_collection": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineBuildStageBuildPipelineStageGitlabServerPredecessorCollectionRepresentation},
		"build_pipeline_stage_type":                   acctest.Representation{RepType: acctest.Required, Create: `BUILD`},
		"defined_tags":                                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                                 acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"build_spec_file":                             acctest.Representation{RepType: acctest.Required, Create: `buildSpecFile`, Update: `buildSpecFile2`},
		"image":                                       acctest.Representation{RepType: acctest.Required, Create: `OL7_X86_64_STANDARD_10`},
		"primary_build_source":                        acctest.Representation{RepType: acctest.Required, Create: `primaryBuildSource`, Update: `primaryBuildSource2`},
		"build_source_collection":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageGitlabServerBuildSourceCollectionRepresentation},
		"stage_execution_timeout_in_seconds":          acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"lifecycle":                                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
		"private_access_config":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsBuildPipelineStagePrivateAccessConfigRepresentation},
	}

	//static values for testing
	gitlabServerSubnetId    = utils.GetEnvSettingWithBlankDefault("test_subnet_id")
	gitlabServerSubnetIdStr = fmt.Sprintf("variable \"test_subnet_id\" { default = \"%s\" }\n", gitlabServerSubnetId)

	gitlabNSGId    = utils.GetEnvSettingWithBlankDefault("test_nsg_id")
	gitlabNSGIdStr = fmt.Sprintf("variable \"test_nsg_id\" { default = \"%s\" }\n", gitlabNSGId)

	gitlabNSGId2    = utils.GetEnvSettingWithBlankDefault("test_nsg_id_update")
	gitlabNSGId2Str = fmt.Sprintf("variable \"test_nsg_id_update\" { default = \"%s\" }\n", gitlabNSGId2)

	DevopsBuildPipelineStagePrivateAccessConfigRepresentation = map[string]interface{}{
		"network_channel_type": acctest.Representation{RepType: acctest.Required, Create: `SERVICE_VNIC_CHANNEL`},
		"subnet_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.test_subnet_id}`},
		"nsg_ids":              acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.test_nsg_id}`}, Update: []string{`${var.test_nsg_id_update}`}},
	}
	buildPipelineBuildStageBuildPipelineStageGitlabServerPredecessorCollectionRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineBuildStageGitlabServerBuildPipelineStagePredecessorCollectionItemsRepresentation},
	}

	buildPipelineStageGitlabServerBuildSourceCollectionRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineStageGitlabServerBuildSourceCollectionItemsRepresentation},
	}

	buildPipelineBuildStageGitlabServerBuildPipelineStagePredecessorCollectionItemsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
	}

	gitlabServerRepoUrl  = "https://gitlabServer.com/dlcbld/docktest"
	gitlabServerRepoUrl2 = "https://gitlabServer.com/dlcbld/deliver-artifact-test"

	buildPipelineStageGitlabServerBuildSourceCollectionItemsRepresentation = map[string]interface{}{
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `GITLAB_SERVER`},
		"branch":          acctest.Representation{RepType: acctest.Required, Create: `branch`, Update: `branch2`},
		"connection_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_connection.test_connection.id}`},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `primaryBuildSource`, Update: `primaryBuildSource2`},
		"repository_url":  acctest.Representation{RepType: acctest.Required, Create: gitlabServerRepoUrl, Update: gitlabServerRepoUrl2},
	}

	gitlabServerAccessTokenVaultId    = utils.GetEnvSettingWithBlankDefault("github_access_token_vault_id")
	gitlabServerAccessTokenVaultIdStr = fmt.Sprintf("variable \"gitlab_access_token_vault_id\" { default = \"%s\" }\n", gitlabServerAccessTokenVaultId)

	BuildPipelineBuildStageGitlabServerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Required, acctest.Create, DevopsBuildPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		gitlabServerAccessTokenVaultIdStr +
		gitlabServerSubnetIdStr +
		gitlabNSGIdStr +
		gitlabNSGId2Str +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, devopsGitlabServerConnectionRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsBuildPipelineBuildStageGitlabServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsBuildPipelineBuildStageGitlabServerResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_build_pipeline_stage.test_build_pipeline_stage"
	datasourceName := "data.oci_devops_build_pipeline_stages.test_build_pipeline_stages"
	singularDatasourceName := "data.oci_devops_build_pipeline_stage.test_build_pipeline_stage"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BuildPipelineBuildStageGitlabServerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Create, buildPipelineBuildStageGitlabServerRepresentation), "devops", "buildPipelineStage", t)

	conf := config + compartmentIdVariableStr + BuildPipelineBuildStageGitlabServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Create, buildPipelineBuildStageGitlabServerRepresentation)
	print(conf)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDevopsBuildPipelineStageDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageGitlabServerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Required, acctest.Create, buildPipelineBuildStageGitlabServerRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "BUILD"),

					resource.TestCheckResourceAttr(resourceName, "build_source_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.branch", "branch"),
					resource.TestCheckResourceAttrSet(resourceName, "build_source_collection.0.items.0.connection_id"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.connection_type", "GITLAB_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.name", "primaryBuildSource"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.repository_url", "https://gitlabServer.com/dlcbld/docktest"),
					resource.TestCheckResourceAttr(resourceName, "image", "OL7_X86_64_STANDARD_10"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						print("// verify create" + resId)
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageGitlabServerResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageGitlabServerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Create, buildPipelineBuildStageGitlabServerRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "BUILD"),

					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "BUILD"),

					resource.TestCheckResourceAttr(resourceName, "build_source_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.branch", "branch"),
					resource.TestCheckResourceAttrSet(resourceName, "build_source_collection.0.items.0.connection_id"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.connection_type", "GITLAB_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.name", "primaryBuildSource"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.repository_url", "https://gitlabServer.com/dlcbld/docktest"),
					resource.TestCheckResourceAttr(resourceName, "image", "OL7_X86_64_STANDARD_10"),
					resource.TestCheckResourceAttr(resourceName, "stage_execution_timeout_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "private_access_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "private_access_config.0.network_channel_type", "SERVICE_VNIC_CHANNEL"),
					resource.TestCheckResourceAttrSet(resourceName, "private_access_config.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "private_access_config.0.subnet_id", gitlabServerSubnetId),
					resource.TestCheckResourceAttrSet(resourceName, "private_access_config.0.nsg_ids.0"),
					resource.TestCheckResourceAttr(resourceName, "private_access_config.0.nsg_ids.0", gitlabNSGId),

					func(s *terraform.State) (err error) {
						resId, _ = acctest.FromInstanceState(s, resourceName, "id")
						print("// verify create optionals" + resId)
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
				Config: config + compartmentIdVariableStr + BuildPipelineBuildStageGitlabServerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Update, buildPipelineBuildStageGitlabServerRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "BUILD"),

					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
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
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.connection_type", "GITLAB_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.name", "primaryBuildSource2"),
					resource.TestCheckResourceAttr(resourceName, "build_source_collection.0.items.0.repository_url", "https://gitlabServer.com/dlcbld/deliver-artifact-test"),
					resource.TestCheckResourceAttr(resourceName, "image", "OL7_X86_64_STANDARD_10"),
					resource.TestCheckResourceAttr(resourceName, "private_access_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "private_access_config.0.network_channel_type", "SERVICE_VNIC_CHANNEL"),
					resource.TestCheckResourceAttrSet(resourceName, "private_access_config.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "private_access_config.0.subnet_id", gitlabServerSubnetId),
					resource.TestCheckResourceAttrSet(resourceName, "private_access_config.0.nsg_ids.0"),
					resource.TestCheckResourceAttr(resourceName, "private_access_config.0.nsg_ids.0", gitlabNSGId2),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stages", "test_build_pipeline_stages", acctest.Optional, acctest.Update, buildPipelineBuildStageGitlabServerDataSourceRepresentation) +
					compartmentIdVariableStr + BuildPipelineBuildStageGitlabServerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Optional, acctest.Update, buildPipelineBuildStageGitlabServerRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "build_pipeline_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "build_pipeline_stage_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", acctest.Required, acctest.Create, buildPipelineBuildStageGitlabServerSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BuildPipelineBuildStageGitlabServerResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "build_pipeline_stage_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_type", "BUILD"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.0.branch", "branch2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.0.connection_type", "GITLAB_SERVER"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.0.name", "primaryBuildSource2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_source_collection.0.items.0.repository_url", "https://gitlabServer.com/dlcbld/deliver-artifact-test"),
					resource.TestCheckResourceAttr(singularDatasourceName, "build_spec_file", "buildSpecFile2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
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
					resource.TestCheckResourceAttr(singularDatasourceName, "private_access_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "private_access_config.0.network_channel_type", "SERVICE_VNIC_CHANNEL"),

					func(s *terraform.State) (err error) {
						return nil
					},
				),
			},
			// verify resource import
			{
				Config:                  config + BuildPipelineBuildStageGitlabServerRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
