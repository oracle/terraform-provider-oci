// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v48/datascience"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ModelDeploymentRequiredOnlyResource = ModelDeploymentResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", Required, Create, modelDeploymentRepresentation)

	ModelDeploymentResourceConfig = ModelDeploymentResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", Optional, Update, modelDeploymentRepresentation)

	modelDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"model_deployment_id": Representation{repType: Required, create: `${oci_datascience_model_deployment.test_model_deployment.id}`},
	}

	modelForModelDeploymentRepresentation = map[string]interface{}{
		"artifact_content_length":      Representation{repType: Required, create: `6954`},
		"model_artifact":               Representation{repType: Required, create: `../examples/datascience/artifact.zip`},
		"compartment_id":               Representation{repType: Required, create: `${var.compartment_id}`},
		"project_id":                   Representation{repType: Required, create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": Representation{repType: Optional, create: `attachment; filename=tfTestArtifact.zip`},
	}

	modelForUpdateModelDeploymentRepresentation = map[string]interface{}{
		"artifact_content_length":      Representation{repType: Required, create: `6954`},
		"model_artifact":               Representation{repType: Required, create: `../examples/datascience/artifact.zip`},
		"compartment_id":               Representation{repType: Required, create: `${var.compartment_id}`},
		"project_id":                   Representation{repType: Required, create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": Representation{repType: Optional, create: `attachment; filename=tfTestArtifact.zip`},
	}

	modelDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"created_by":     Representation{repType: Optional, create: `${var.user_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"id":             Representation{repType: Optional, create: `${oci_datascience_model_deployment.test_model_deployment.id}`},
		"project_id":     Representation{repType: Required, create: `${oci_datascience_project.test_project.id}`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, modelDeploymentDataSourceFilterRepresentation}}
	modelDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_datascience_model_deployment.test_model_deployment.id}`}},
	}

	modelDeploymentRepresentation = map[string]interface{}{
		"compartment_id":                         Representation{repType: Required, create: `${var.compartment_id}`},
		"model_deployment_configuration_details": RepresentationGroup{Required, modelDeploymentModelDeploymentConfigurationDetailsRepresentation},
		"project_id":                             Representation{repType: Required, create: `${oci_datascience_project.test_project.id}`},
		"category_log_details":                   RepresentationGroup{Optional, modelDeploymentCategoryLogDetailsRepresentation},
		"defined_tags":                           Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                            Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":                           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":                          Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	modelDeploymentModelDeploymentConfigurationDetailsRepresentation = map[string]interface{}{
		"deployment_type":             Representation{repType: Required, create: `SINGLE_MODEL`},
		"model_configuration_details": RepresentationGroup{Required, modelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation},
	}
	modelDeploymentCategoryLogDetailsRepresentation = map[string]interface{}{
		"access":  RepresentationGroup{Optional, modelDeploymentCategoryLogDetailsAccessRepresentation},
		"predict": RepresentationGroup{Optional, modelDeploymentCategoryLogDetailsPredictRepresentation},
	}
	modelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation = map[string]interface{}{
		"instance_configuration": RepresentationGroup{Required, modelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationRepresentation},
		"model_id":               Representation{repType: Required, create: `${oci_datascience_model.test_model.id}`, update: `${oci_datascience_model.test_model_update.id}`},
		"bandwidth_mbps":         Representation{repType: Optional, create: `10`},
		"scaling_policy":         RepresentationGroup{Optional, modelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicyRepresentation},
	}
	modelDeploymentCategoryLogDetailsAccessRepresentation = map[string]interface{}{
		"log_group_id": Representation{repType: Required, create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       Representation{repType: Required, create: `${oci_logging_log.test_access_log.id}`},
	}
	modelDeploymentCategoryLogDetailsPredictRepresentation = map[string]interface{}{
		"log_group_id": Representation{repType: Required, create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       Representation{repType: Required, create: `${oci_logging_log.test_predict_log.id}`},
	}
	modelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationRepresentation = map[string]interface{}{
		"instance_shape_name": Representation{repType: Required, create: `VM.Standard2.1`, update: `VM.Standard2.2`},
	}
	modelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicyRepresentation = map[string]interface{}{
		"instance_count": Representation{repType: Required, create: `1`},
		"policy_type":    Representation{repType: Required, create: `FIXED_SIZE`},
	}

	logGroupMDRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `tf_testing_log_group`, update: `tf_testing_log_group_update`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	logGroupUpdateMDRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `tf_update_testing_log_group`, update: `tf_update_testing_log_group_update`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	customAccessLogRepresentation = map[string]interface{}{
		"display_name":       Representation{repType: Required, create: `tf-testing-access-log`, update: `tf-testing-update-access-log`},
		"log_group_id":       Representation{repType: Required, create: `${oci_logging_log_group.test_log_group.id}`, update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           Representation{repType: Required, create: `CUSTOM`},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         Representation{repType: Optional, create: `false`, update: `true`},
		"retention_duration": Representation{repType: Optional, create: `30`, update: `60`},
	}

	customPredictLogRepresentation = map[string]interface{}{
		"display_name":       Representation{repType: Required, create: `tf-testing-predict-log`, update: `tf-testing-update-predict-log`},
		"log_group_id":       Representation{repType: Required, create: `${oci_logging_log_group.test_log_group.id}`, update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           Representation{repType: Required, create: `CUSTOM`},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         Representation{repType: Optional, create: `false`, update: `true`},
		"retention_duration": Representation{repType: Optional, create: `30`, update: `60`},
	}

	ModelDeploymentResourceDependencies = generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Optional, Create, modelForModelDeploymentRepresentation) +
		generateResourceFromRepresentationMap("oci_datascience_model", "test_model_update", Optional, Create, modelForUpdateModelDeploymentRepresentation) +
		generateResourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectRepresentation) +

		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupMDRepresentation) +
		generateResourceFromRepresentationMap("oci_logging_log", "test_access_log", Required, Create, customAccessLogRepresentation) +
		generateResourceFromRepresentationMap("oci_logging_log", "test_predict_log", Required, Create, customPredictLogRepresentation) +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", Required, Create, logGroupUpdateMDRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatascienceModelDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	userId := getEnvSettingWithBlankDefault("user_ocid")
	userIdVariableStr := fmt.Sprintf("variable \"user_id\" { default = \"%s\" }\n", userId)

	resourceName := "oci_datascience_model_deployment.test_model_deployment"
	datasourceName := "data.oci_datascience_model_deployments.test_model_deployments"
	singularDatasourceName := "data.oci_datascience_model_deployment.test_model_deployment"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatascienceModelDeploymentDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + ModelDeploymentResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", Required, Create, modelDeploymentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + ModelDeploymentResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + ModelDeploymentResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", Optional, Create, modelDeploymentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ModelDeploymentResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", Optional, Create,
					representationCopyWithNewProperties(modelDeploymentRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ModelDeploymentResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", Optional, Update, modelDeploymentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource - list model deployments
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_datascience_model_deployments", "test_model_deployments", Optional, Update, modelDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + userIdVariableStr + ModelDeploymentResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", Optional, Update, modelDeploymentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "created_by", userId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "model_deployments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.category_log_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.defined_tags.%", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.time_created"),
			),
		},
		// verify singular datasource - get model deployment
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", Required, Create, modelDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ModelDeploymentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ModelDeploymentResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"description",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatascienceModelDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model_deployment" {
			noResourceFound = false
			request := oci_datascience.GetModelDeploymentRequest{}

			tmp := rs.Primary.ID
			request.ModelDeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datascience")

			response, err := client.GetModelDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ModelDeploymentLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatascienceModelDeployment") {
		resource.AddTestSweepers("DatascienceModelDeployment", &resource.Sweeper{
			Name:         "DatascienceModelDeployment",
			Dependencies: DependencyGraph["modelDeployment"],
			F:            sweepDatascienceModelDeploymentResource,
		})
	}
}

func sweepDatascienceModelDeploymentResource(compartment string) error {
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()
	modelDeploymentIds, err := getModelDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, modelDeploymentId := range modelDeploymentIds {
		if ok := SweeperDefaultResourceId[modelDeploymentId]; !ok {
			deleteModelDeploymentRequest := oci_datascience.DeleteModelDeploymentRequest{}

			deleteModelDeploymentRequest.ModelDeploymentId = &modelDeploymentId

			deleteModelDeploymentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModelDeployment(context.Background(), deleteModelDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting ModelDeployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelDeploymentId, error)
				continue
			}
			waitTillCondition(testAccProvider, &modelDeploymentId, modelDeploymentSweepWaitCondition, time.Duration(3*time.Minute),
				modelDeploymentSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getModelDeploymentIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ModelDeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()

	listModelDeploymentsRequest := oci_datascience.ListModelDeploymentsRequest{}
	listModelDeploymentsRequest.CompartmentId = &compartmentId
	listModelDeploymentsRequest.LifecycleState = oci_datascience.ListModelDeploymentsLifecycleStateNeedsAttention
	// listModelDeploymentsRequest.LifecycleState = oci_datascience.ListModelDeploymentsLifecycleStateActiveNeedsAttention
	listModelDeploymentsResponse, err := dataScienceClient.ListModelDeployments(context.Background(), listModelDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ModelDeployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, modelDeployment := range listModelDeploymentsResponse.Items {
		id := *modelDeployment.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ModelDeploymentId", id)
	}
	return resourceIds, nil
}

func modelDeploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelDeploymentResponse, ok := response.Response.(oci_datascience.GetModelDeploymentResponse); ok {
		return modelDeploymentResponse.LifecycleState != oci_datascience.ModelDeploymentLifecycleStateDeleted
	}
	return false
}

func modelDeploymentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataScienceClient().GetModelDeployment(context.Background(), oci_datascience.GetModelDeploymentRequest{
		ModelDeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
