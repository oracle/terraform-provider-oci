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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DeployFunctionEnvironmentRequiredOnlyResource = DevopsDeployEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployFunctionEnvironmentRepresentation)

	DeployFunctionEnvironmentResourceConfig = DevopsDeployEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployFunctionEnvironmentRepresentation)

	deployFunctionEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
	}

	deployFunctionEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: deployFunctionEnvironmentDataSourceFilterRepresentation}}
	deployFunctionEnvironmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deploy_environment.test_deploy_environment.id}`}},
	}

	function_fake_id                        = "ocid1.fnfunc.oc1.us-ashburn-1.aaaaaaaaafqtkm3fg4zwgnlggmywkzdemi2dcyzymfrdqojygcstofunction1"
	deployFunctionEnvironmentRepresentation = acctest.GetUpdatedRepresentationCopy("deploy_environment_type", acctest.Representation{RepType: acctest.Required, Create: `FUNCTION`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DevopsdeployEnvironmentRepresentation, []string{"cluster_id"}), map[string]interface{}{
			"function_id": acctest.Representation{RepType: acctest.Required, Create: function_fake_id},
		}))
)

// issue-routing-tag: devops/default
func TestDevopsDeployEnvironmentResource_function(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployEnvironmentResource_function")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_environment.test_deploy_environment"
	datasourceName := "data.oci_devops_deploy_environments.test_deploy_environments"
	singularDatasourceName := "data.oci_devops_deploy_environment.test_deploy_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsDeployEnvironmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Create, deployFunctionEnvironmentRepresentation), "devops", "deployEnvironment", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDevopsDeployEnvironmentDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployFunctionEnvironmentRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "function_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "FUNCTION"),
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
					acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Create, deployFunctionEnvironmentRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "function_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "FUNCTION"),
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
					acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployFunctionEnvironmentRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "function_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "FUNCTION"),
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environments", "test_deploy_environments", acctest.Optional, acctest.Update, deployFunctionEnvironmentDataSourceRepresentation) +
					compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployFunctionEnvironmentRepresentation),
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployFunctionEnvironmentSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DeployFunctionEnvironmentResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_environment_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deploy_environment_type", "FUNCTION"),
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
				Config:                  config + DeployFunctionEnvironmentRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
