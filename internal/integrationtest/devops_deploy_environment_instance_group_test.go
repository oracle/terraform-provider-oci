// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
	DeployInstanceGroupEnvironmentRequiredOnlyResource = DevopsDeployEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployInstanceGroupEnvironmentRepresentation)

	DeployInstanceGroupEnvironmentResourceConfig = DevopsDeployEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployInstanceGroupEnvironmentRepresentation)

	deployInstanceGroupEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
	}

	deployInstanceGroupEnvironmentRepresentation = acctest.GetUpdatedRepresentationCopy("deploy_environment_type", acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_INSTANCE_GROUP`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DevopsdeployEnvironmentRepresentation, []string{"cluster_id"}), map[string]interface{}{
			"compute_instance_group_selectors": acctest.RepresentationGroup{RepType: acctest.Required, Group: deployComputeInstanceGroupEnvironmentSelectorCollectionRepresentation},
		}))

	deployComputeInstanceGroupEnvironmentSelectorCollectionRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deployComputeInstanceGroupEnvironmentSelectorCollectionItemsRepresentation},
	}
	deployComputeInstanceGroupEnvironmentSelectorCollectionItemsRepresentation = map[string]interface{}{
		"selector_type":        acctest.Representation{RepType: acctest.Required, Create: `INSTANCE_IDS`},
		"compute_instance_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`ocid1.instance.oc1.iad.anuwcljtnsx72macffe5fbkzbj4eerle5ot56g2cexj3jvfsr242pye44ghq`}},
	}
)

// issue-routing-tag: devops/default
func TestDevopsDeployEnvironmentResource_instanceGroup(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployEnvironmentResource_instanceGroup")
	defer httpreplay.SaveScenario()
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_environment.test_deploy_environment"
	datasourceName := "data.oci_devops_deploy_environments.test_deploy_environments"
	singularDatasourceName := "data.oci_devops_deploy_environment.test_deploy_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsDeployEnvironmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Create, deployInstanceGroupEnvironmentRepresentation), "devops", "deployEnvironment", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployEnvironmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployInstanceGroupEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "COMPUTE_INSTANCE_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Create, deployInstanceGroupEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "COMPUTE_INSTANCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
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
			Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployInstanceGroupEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "COMPUTE_INSTANCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environments", "test_deploy_environments", acctest.Optional, acctest.Update, DevopsDevopsDeployEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployInstanceGroupEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_environment_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployInstanceGroupEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployInstanceGroupEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_environment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_environment_type", "COMPUTE_INSTANCE_GROUP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DeployInstanceGroupEnvironmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
