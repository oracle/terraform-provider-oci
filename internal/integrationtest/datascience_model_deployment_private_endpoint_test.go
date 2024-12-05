package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatasciencePrivateModelDeploymentRequiredOnlyResource = DatasciencePrivateModelDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatasciencePrivateModelDeploymentRepresentation)

	DatasciencePrivateModelDeploymentResourceConfig = DatasciencePrivateModelDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatasciencePrivateModelDeploymentRepresentation)

	DatascienceDatasciencePrivateModelDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"model_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model_deployment.test_model_deployment.id}`},
	}

	DatascienceDatasciencePrivateModelDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_deployment.test_model_deployment.created_by}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_deployment.test_model_deployment.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentDataSourceFilterRepresentation}}

	DatasciencePrivateModelDeploymentRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_deployment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatasciencePrivateModelDeploymentModelDeploymentConfigurationDetailsRepresentation},
		"project_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"category_log_details":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentCategoryLogDetailsRepresentation},
		"defined_tags":                           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreMDRepresentation},
	}

	DatasciencePrivateModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationRepresentation = map[string]interface{}{
		"instance_shape_name":                            acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`, Update: `VM.Standard.E3.Flex`},
		"model_deployment_instance_shape_config_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationModelDeploymentInstanceShapeConfigDetailsRepresentation},
		"private_endpoint_id":                            acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_private_endpoint.test_md_private_endpoint.id}`},
		"subnet_id":                                      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	DatasciencePrivateModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation = map[string]interface{}{
		"instance_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatasciencePrivateModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationRepresentation},
		"model_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`, Update: `${oci_datascience_model.test_model_update.id}`},
		"bandwidth_mbps":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"maximum_bandwidth_mbps": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"scaling_policy":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicyRepresentation},
	}

	DatasciencePrivateModelDeploymentModelDeploymentConfigurationDetailsRepresentation = map[string]interface{}{
		"deployment_type":             acctest.Representation{RepType: acctest.Required, Create: `SINGLE_MODEL`},
		"model_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatasciencePrivateModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation},
	}

	DatasciencePrivateModelDeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, modelForModelDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model_update", acctest.Optional, acctest.Create, modelForUpdateModelDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, ModelDeploymentCoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_md_private_endpoint", acctest.Required, acctest.Create, DataScienceMDPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupMDRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_access_log", acctest.Required, acctest.Create, customAccessLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_predict_log", acctest.Required, acctest.Create, customPredictLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Create, logGroupUpdateMDRepresentation)
)

func TestDatasciencePrivateModelDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatasciencePrivateModelDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_model_deployment.test_model_deployment"
	datasourceName := "data.oci_datascience_model_deployments.test_model_deployments"
	singularDatasourceName := "data.oci_datascience_model_deployment.test_model_deployment"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatascienceModelDeploymentDestroy, []resource.TestStep{
		// 0. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatasciencePrivateModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create, DatasciencePrivateModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_8"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.private_endpoint_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),

				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.ocpus", "1"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			)},

		// 1. verify Deactivate model deployment
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatasciencePrivateModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceModelDeploymentRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_8"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// 2. verify Activate model deployment
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatasciencePrivateModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatasciencePrivateModelDeploymentRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_8"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// 3. verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatasciencePrivateModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatasciencePrivateModelDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_8"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.private_endpoint_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.subnet_id"),

				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),

				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.ocpus", "1"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// 4. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatasciencePrivateModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatasciencePrivateModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_2"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),

				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.memory_in_gbs", "12"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.ocpus", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.subnet_id"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// 5. verify datasource - list model deployments
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployments", "test_model_deployments", acctest.Optional, acctest.Update, DatascienceDatasciencePrivateModelDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DatasciencePrivateModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatasciencePrivateModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
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
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_2"),

				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),

				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.memory_in_gbs", "12"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.ocpus", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.subnet_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "2"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.project_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.time_created"),
			),
		},

		// 6. verify singular datasource - get model deployment
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceDatasciencePrivateModelDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatasciencePrivateModelDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_url"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "model_deployment_system_data.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		// 7. verify resource import
		{
			Config:            config + DatasciencePrivateModelDeploymentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"description",
			},
			ResourceName: resourceName,
		},
	})
}
