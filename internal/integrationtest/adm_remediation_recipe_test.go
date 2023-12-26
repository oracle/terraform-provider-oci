// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreRemediationRecipeDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	AdmRemediationRecipeRequiredOnlyResource = AdmRemediationRecipeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Required, acctest.Create, AdmRemediationRecipeRepresentation)
	AdmRemediationRecipeResourceConfig = AdmRemediationRecipeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Optional, acctest.Update, AdmRemediationRecipeRepresentation)

	AdmRemediationRecipeSingularDataSourceRepresentation = map[string]interface{}{
		"remediation_recipe_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_adm_remediation_recipe.test_remediation_recipe.id}`},
	}

	AdmRemediationRecipeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_adm_remediation_recipe.test_remediation_recipe.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AdmRemediationRecipeDataSourceFilterRepresentation}}
	AdmRemediationRecipeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_adm_remediation_recipe.test_remediation_recipe.id}`}},
	}

	AdmRemediationRecipeRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"detect_configuration":          acctest.RepresentationGroup{RepType: acctest.Required, Group: AdmRemediationRecipeDetectConfigurationRepresentation},
		"is_run_triggered_on_kb_change": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"knowledge_base_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_adm_knowledge_base.test_knowledge_base.id}`},
		"network_configuration":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AdmRemediationRecipeNetworkConfigurationRepresentation},
		"scm_configuration":             acctest.RepresentationGroup{RepType: acctest.Required, Group: AdmRemediationRecipeScmConfigurationRepresentation},
		"verify_configuration":          acctest.RepresentationGroup{RepType: acctest.Required, Group: AdmRemediationRecipeVerifyConfigurationRepresentation},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"state":                         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `INACTIVE`},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRemediationRecipeDefinedTagsChangesRepresentation},
	}
	AdmRemediationRecipeDetectConfigurationRepresentation = map[string]interface{}{
		"exclusions":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`exclusions`}, Update: []string{`exclusions2`}},
		"max_permissible_cvss_v2score": acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		"max_permissible_cvss_v3score": acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		"max_permissible_severity":     acctest.Representation{RepType: acctest.Optional, Create: `LOW`},
		"upgrade_policy":               acctest.Representation{RepType: acctest.Optional, Create: `NEAREST`},
	}
	AdmRemediationRecipeNetworkConfigurationRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"nsg_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}, Update: []string{`nsgIds2`}},
	}
	AdmRemediationRecipeScmConfigurationRepresentation = map[string]interface{}{
		"branch":                 acctest.Representation{RepType: acctest.Required, Create: "branch", Update: `branch2`},
		"is_automerge_enabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"scm_type":               acctest.Representation{RepType: acctest.Required, Create: `OCI_CODE_REPOSITORY`, Update: `EXTERNAL_SCM`},
		"build_file_location":    acctest.Representation{RepType: acctest.Optional, Create: `buildFileLocation`, Update: `buildFileLocation2`},
		"external_scm_type":      acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `GITLAB`},
		"oci_code_repository_id": acctest.Representation{RepType: acctest.Required, Create: `${var.devops_code_repository_ocid}`},
		"pat_secret_id":          acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `${var.kms_secret_ocid}`},
		"repository_url":         acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `repositoryUrl2`},
		"username":               acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `username2`},
	}
	AdmRemediationRecipeVerifyConfigurationRepresentation = map[string]interface{}{
		"build_service_type":    acctest.Representation{RepType: acctest.Required, Create: `OCI_DEVOPS_BUILD`, Update: `GITLAB_PIPELINE`},
		"additional_parameters": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"additionalParameters": "additionalParameters"}, Update: map[string]string{"additionalParameters2": "additionalParameters2"}},
		"pat_secret_id":         acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `${var.kms_secret_ocid}`},
		"pipeline_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.devops_build_pipeline_ocid}`},
		"repository_url":        acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `repositoryUrl2`},
		"trigger_secret_id":     acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `${var.kms_secret_ocid}`},
		"username":              acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `username2`},
	}

	AdmRemediationRecipeForUpdateRepresentation = acctest.GetRepresentationCopyWithMultipleRemovedProperties(
		[]string{"scm_configuration.oci_code_repository_id", "verify_configuration.pipeline_id"}, AdmRemediationRecipeRepresentation)

	AdmRemediationRecipeResourceForUpdateConfig = AdmRemediationRecipeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Optional, acctest.Update, AdmRemediationRecipeForUpdateRepresentation)

	AdmRemediationRecipeResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, knowledgeBaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) + DefinedTagsDependencies
)

// issue-routing-tag: adm/default
func TestAdmRemediationRecipeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAdmRemediationRecipeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	secretId := utils.GetEnvSettingWithBlankDefault("kms_secret_ocid")
	secretIdVariableStr := fmt.Sprintf("variable \"kms_secret_ocid\" { default = \"%s\" }\n", secretId)

	codeRepositoryId := utils.GetEnvSettingWithBlankDefault("devops_code_repository_ocid")
	codeRepositoryIdStr := fmt.Sprintf("variable \"devops_code_repository_ocid\" { default = \"%s\" }\n", codeRepositoryId)

	pipelineId := utils.GetEnvSettingWithBlankDefault("devops_build_pipeline_ocid")
	pipelineIdStr := fmt.Sprintf("variable \"devops_build_pipeline_ocid\" { default = \"%s\" }\n", pipelineId)

	resourceName := "oci_adm_remediation_recipe.test_remediation_recipe"
	datasourceName := "data.oci_adm_remediation_recipes.test_remediation_recipes"
	singularDatasourceName := "data.oci_adm_remediation_recipe.test_remediation_recipe"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AdmRemediationRecipeResourceDependencies+secretIdVariableStr+codeRepositoryIdStr+pipelineIdStr+
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Optional, acctest.Create, AdmRemediationRecipeRepresentation), "adm", "remediationRecipe", t)

	acctest.ResourceTest(t, testAccCheckAdmRemediationRecipeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + secretIdVariableStr + compartmentIdVariableStr + AdmRemediationRecipeResourceDependencies + codeRepositoryIdStr + pipelineIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Required, acctest.Create, AdmRemediationRecipeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_run_triggered_on_kb_change", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.is_automerge_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.branch", "branch"),
				resource.TestCheckResourceAttrSet(resourceName, "scm_configuration.0.oci_code_repository_id"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.scm_type", "OCI_CODE_REPOSITORY"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.build_service_type", "OCI_DEVOPS_BUILD"),
				resource.TestCheckResourceAttrSet(resourceName, "verify_configuration.0.pipeline_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr + AdmRemediationRecipeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr + compartmentIdVariableStr + AdmRemediationRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Optional, acctest.Create, AdmRemediationRecipeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.exclusions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.max_permissible_cvss_v2score", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.max_permissible_cvss_v3score", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.max_permissible_severity", "LOW"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_triggered_on_kb_change", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.build_file_location", "buildFileLocation"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.branch", "branch"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.is_automerge_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "scm_configuration.0.oci_code_repository_id"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.scm_type", "OCI_CODE_REPOSITORY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.additional_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.build_service_type", "OCI_DEVOPS_BUILD"),
				resource.TestCheckResourceAttrSet(resourceName, "verify_configuration.0.pipeline_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr + compartmentIdVariableStr + compartmentIdUVariableStr + AdmRemediationRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AdmRemediationRecipeRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.exclusions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.max_permissible_cvss_v2score", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.max_permissible_cvss_v3score", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.max_permissible_severity", "LOW"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_triggered_on_kb_change", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.branch", "branch"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.build_file_location", "buildFileLocation"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.is_automerge_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "scm_configuration.0.oci_code_repository_id"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.scm_type", "OCI_CODE_REPOSITORY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.additional_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.build_service_type", "OCI_DEVOPS_BUILD"),
				resource.TestCheckResourceAttrSet(resourceName, "verify_configuration.0.pipeline_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr + compartmentIdVariableStr + AdmRemediationRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Optional, acctest.Update, AdmRemediationRecipeForUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.exclusions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "detect_configuration.0.max_permissible_cvss_v2score"),
				resource.TestCheckResourceAttrSet(resourceName, "detect_configuration.0.max_permissible_cvss_v3score"),
				resource.TestCheckResourceAttrSet(resourceName, "detect_configuration.0.max_permissible_severity"),
				resource.TestCheckResourceAttr(resourceName, "detect_configuration.0.upgrade_policy", "NEAREST"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_triggered_on_kb_change", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.branch", "branch2"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.build_file_location", "buildFileLocation2"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.external_scm_type", "GITLAB"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.is_automerge_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "scm_configuration.0.pat_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.repository_url", "repositoryUrl2"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.scm_type", "EXTERNAL_SCM"),
				resource.TestCheckResourceAttr(resourceName, "scm_configuration.0.username", "username2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.additional_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.build_service_type", "GITLAB_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "verify_configuration.0.pat_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.repository_url", "repositoryUrl2"),
				resource.TestCheckResourceAttrSet(resourceName, "verify_configuration.0.trigger_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "verify_configuration.0.username", "username2"),

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
			Config: config + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_adm_remediation_recipes", "test_remediation_recipes", acctest.Optional, acctest.Update, AdmRemediationRecipeDataSourceRepresentation) +
				compartmentIdVariableStr + AdmRemediationRecipeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Optional, acctest.Update, AdmRemediationRecipeForUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "remediation_recipe_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "remediation_recipe_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Required, acctest.Create, AdmRemediationRecipeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AdmRemediationRecipeResourceForUpdateConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remediation_recipe_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "detect_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detect_configuration.0.exclusions.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detect_configuration.0.max_permissible_cvss_v2score"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detect_configuration.0.max_permissible_cvss_v3score"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detect_configuration.0.max_permissible_severity"),
				resource.TestCheckResourceAttr(singularDatasourceName, "detect_configuration.0.upgrade_policy", "NEAREST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_triggered_on_kb_change", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scm_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scm_configuration.0.branch", "branch2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scm_configuration.0.build_file_location", "buildFileLocation2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scm_configuration.0.external_scm_type", "GITLAB"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scm_configuration.0.is_automerge_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scm_configuration.0.repository_url", "repositoryUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scm_configuration.0.scm_type", "EXTERNAL_SCM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scm_configuration.0.username", "username2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "verify_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "verify_configuration.0.additional_parameters.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "verify_configuration.0.build_service_type", "GITLAB_PIPELINE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "verify_configuration.0.repository_url", "repositoryUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "verify_configuration.0.username", "username2"),
			),
		},
		// verify resource import
		{
			Config:                  config + AdmRemediationRecipeRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAdmRemediationRecipeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApplicationDependencyManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_adm_remediation_recipe" {
			noResourceFound = false
			request := oci_adm.GetRemediationRecipeRequest{}

			tmp := rs.Primary.ID
			request.RemediationRecipeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "adm")

			response, err := client.GetRemediationRecipe(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_adm.RemediationRecipeLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("AdmRemediationRecipe") {
		resource.AddTestSweepers("AdmRemediationRecipe", &resource.Sweeper{
			Name:         "AdmRemediationRecipe",
			Dependencies: acctest.DependencyGraph["remediationRecipe"],
			F:            sweepAdmRemediationRecipeResource,
		})
	}
}

func sweepAdmRemediationRecipeResource(compartment string) error {
	applicationDependencyManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ApplicationDependencyManagementClient()
	remediationRecipeIds, err := getAdmRemediationRecipeIds(compartment)
	if err != nil {
		return err
	}
	for _, remediationRecipeId := range remediationRecipeIds {
		if ok := acctest.SweeperDefaultResourceId[remediationRecipeId]; !ok {
			deleteRemediationRecipeRequest := oci_adm.DeleteRemediationRecipeRequest{}

			deleteRemediationRecipeRequest.RemediationRecipeId = &remediationRecipeId

			deleteRemediationRecipeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "adm")
			_, error := applicationDependencyManagementClient.DeleteRemediationRecipe(context.Background(), deleteRemediationRecipeRequest)
			if error != nil {
				fmt.Printf("Error deleting RemediationRecipe %s %s, It is possible that the resource is already deleted. Please verify manually \n", remediationRecipeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &remediationRecipeId, AdmRemediationRecipeSweepWaitCondition, time.Duration(3*time.Minute),
				AdmRemediationRecipeSweepResponseFetchOperation, "adm", true)
		}
	}
	return nil
}

func getAdmRemediationRecipeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RemediationRecipeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	applicationDependencyManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ApplicationDependencyManagementClient()

	listRemediationRecipesRequest := oci_adm.ListRemediationRecipesRequest{}
	listRemediationRecipesRequest.CompartmentId = &compartmentId
	listRemediationRecipesRequest.LifecycleState = oci_adm.RemediationRecipeLifecycleStateActive
	listRemediationRecipesResponse, err := applicationDependencyManagementClient.ListRemediationRecipes(context.Background(), listRemediationRecipesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RemediationRecipe list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, remediationRecipe := range listRemediationRecipesResponse.Items {
		id := *remediationRecipe.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RemediationRecipeId", id)
	}
	return resourceIds, nil
}

func AdmRemediationRecipeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if remediationRecipeResponse, ok := response.Response.(oci_adm.GetRemediationRecipeResponse); ok {
		return remediationRecipeResponse.LifecycleState != oci_adm.RemediationRecipeLifecycleStateDeleted
	}
	return false
}

func AdmRemediationRecipeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ApplicationDependencyManagementClient().GetRemediationRecipe(context.Background(), oci_adm.GetRemediationRecipeRequest{
		RemediationRecipeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
