package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	// MD config consists of creating dependencies and generate MD config
	DatascienceModelDeploymentModelExplorerResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create, DatascienceModelDeploymentModelExplorerRepresentation)
	DatascienceModelDeploymentModelExplorerRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_deployment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentModelExplorerModelDeploymentConfigurationDetailsRepresentation},
		"project_id":                             acctest.Representation{RepType: acctest.Required, Create: `${var.model_explorer_project_id}`},
		"description":                            acctest.Representation{RepType: acctest.Optional, Create: `description1`, Update: `description1`},
		"display_name":                           acctest.Representation{RepType: acctest.Optional, Create: `displayName1`, Update: `displayName2`},
	}

	// model deployment configuration details related.
	DatascienceModelDeploymentModelExplorerModelDeploymentConfigurationDetailsRepresentation = map[string]interface{}{
		"deployment_type":             acctest.Representation{RepType: acctest.Required, Create: `SINGLE_MODEL`},
		"model_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: ModelConfigurationDetailsRepresentation},
	}
	ModelConfigurationDetailsRepresentation = map[string]interface{}{
		"instance_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: InstanceConfigurationRepresentation},
		"model_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.model_explorer_model_id}`},
		"bandwidth_mbps":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"maximum_bandwidth_mbps": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"scaling_policy":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicyRepresentation},
	}
	InstanceConfigurationRepresentation = map[string]interface{}{
		"subnet_id":           acctest.Representation{RepType: acctest.Optional, Create: nil},
		"instance_shape_name": acctest.Representation{RepType: acctest.Required, Create: `VM.GPU.A10.1`},
	}
)

func TestDatascienceModelDeploymentModelExplorerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentModelExplorerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_deployment.test_model_deployment"
	singularDatasourceName := "data.oci_datascience_model_deployment.test_model_deployment"
	var resId string

	acctest.ResourceTest(t, testAccCheckDatascienceModelDeploymentDestroy, []resource.TestStep{
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create, DatascienceModelDeploymentModelExplorerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_system_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_system_data.0.model_type", "MANAGED_MODEL"),
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
			),
		},

		// verify singular datasource - get model deployment
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceDatascienceModelDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelDeploymentModelExplorerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_system_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_system_data.0.model_type", "MANAGED_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},
	})
}
