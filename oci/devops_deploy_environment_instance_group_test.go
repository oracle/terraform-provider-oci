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
	DeployInstanceGroupEnvironmentRequiredOnlyResource = DeployEnvironmentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployInstanceGroupEnvironmentRepresentation)

	DeployInstanceGroupEnvironmentResourceConfig = DeployEnvironmentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Update, deployInstanceGroupEnvironmentRepresentation)

	deployInstanceGroupEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_environment_id": Representation{RepType: Required, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
	}

	deployInstanceGroupEnvironmentRepresentation = GetUpdatedRepresentationCopy("deploy_environment_type", Representation{RepType: Required, Create: `COMPUTE_INSTANCE_GROUP`},
		RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(deployEnvironmentRepresentation, []string{"cluster_id"}), map[string]interface{}{
			"compute_instance_group_selectors": RepresentationGroup{Required, deployComputeInstanceGroupEnvironmentSelectorCollectionRepresentation},
		}))

	deployComputeInstanceGroupEnvironmentSelectorCollectionRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, deployComputeInstanceGroupEnvironmentSelectorCollectionItemsRepresentation},
	}
	deployComputeInstanceGroupEnvironmentSelectorCollectionItemsRepresentation = map[string]interface{}{
		"selector_type":        Representation{RepType: Required, Create: `INSTANCE_IDS`},
		"compute_instance_ids": Representation{RepType: Required, Create: []string{`ocid1.instance.oc1.phx.anuwcljtnsx72macffe5fbkzbj4eerle5ot56g2cexj3jvfsr242pye44ghq`}},
	}
)

// issue-routing-tag: devops/default
func TestDevopsDeployEnvironmentResource_instanceGroup(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployEnvironmentResource_instanceGroup")
	defer httpreplay.SaveScenario()
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_environment.test_deploy_environment"
	datasourceName := "data.oci_devops_deploy_environments.test_deploy_environments"
	singularDatasourceName := "data.oci_devops_deploy_environment.test_deploy_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DeployEnvironmentResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Create, deployInstanceGroupEnvironmentRepresentation), "devops", "deployEnvironment", t)

	ResourceTest(t, testAccCheckDevopsDeployEnvironmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployInstanceGroupEnvironmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "COMPUTE_INSTANCE_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Create, deployInstanceGroupEnvironmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "COMPUTE_INSTANCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Update, deployInstanceGroupEnvironmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_instance_group_selectors.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "COMPUTE_INSTANCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environments", "test_deploy_environments", Optional, Update, deployEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + DeployEnvironmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Update, deployInstanceGroupEnvironmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployInstanceGroupEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployInstanceGroupEnvironmentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_environment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_environment_type", "COMPUTE_INSTANCE_GROUP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
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
			Config: config + compartmentIdVariableStr + DeployInstanceGroupEnvironmentResourceConfig,
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
