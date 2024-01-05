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
	DeployOkeEnvironmentRequiredOnlyResource = DeployOkeEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployPrivateOkeEnvironmentRepresentation)

	DeployOkeEnvironmentResourceConfig = DeployOkeEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployPrivateOkeEnvironmentRepresentation)

	deployOkeEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
	}

	deployOkeEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: deployOkeEnvironmentDataSourceFilterRepresentation}}
	deployOkeEnvironmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deploy_environment.test_deploy_environment.id}`}},
	}

	private_cluster_fake_id                   = utils.GetEnvSettingWithBlankDefault("cluster_id")
	deployPrivateOkeEnvironmentRepresentation = map[string]interface{}{
		"deploy_environment_type": acctest.Representation{RepType: acctest.Required, Create: `OKE_CLUSTER`},
		"project_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"cluster_id":              acctest.Representation{RepType: acctest.Required, Create: private_cluster_fake_id},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
		"network_channel":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: deployEnvironmentNetworkChannelRepresentation},
	}
	private_subnet_id                             = utils.GetEnvSettingWithBlankDefault("subnet_id")
	deployEnvironmentNetworkChannelRepresentation = map[string]interface{}{
		"network_channel_type": acctest.Representation{RepType: acctest.Required, Create: `PRIVATE_ENDPOINT_CHANNEL`},
		"subnet_id":            acctest.Representation{RepType: acctest.Required, Create: private_subnet_id},
		"nsg_ids":              acctest.Representation{RepType: acctest.Optional, Create: []string{}},
	}
	DeployOkeEnvironmentResourceDependencies = AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", acctest.Required, acctest.Create, DevopsLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: devops/default
func TestDevopsDeployEnvironmentResource_privateOke(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployEnvironmentResource_privateOke")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_environment.test_deploy_environment"
	datasourceName := "data.oci_devops_deploy_environments.test_deploy_environments"
	singularDatasourceName := "data.oci_devops_deploy_environment.test_deploy_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeployOkeEnvironmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Create, deployPrivateOkeEnvironmentRepresentation), "devops", "deployEnvironment", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployEnvironmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployOkeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployPrivateOkeEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployOkeEnvironmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployOkeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Create, deployPrivateOkeEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_channel.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_channel.0.network_channel_type", "PRIVATE_ENDPOINT_CHANNEL"),
				resource.TestCheckResourceAttrSet(resourceName, "network_channel.0.subnet_id"),
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
			Config: config + compartmentIdVariableStr + DeployOkeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployPrivateOkeEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_channel.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_channel.0.network_channel_type", "PRIVATE_ENDPOINT_CHANNEL"),
				resource.TestCheckResourceAttrSet(resourceName, "network_channel.0.subnet_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environments", "test_deploy_environments", acctest.Optional, acctest.Update, deployOkeEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + DeployOkeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, deployPrivateOkeEnvironmentRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployOkeEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployOkeEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_environment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_environment_type", "OKE_CLUSTER"),
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
			Config:                  config + DeployOkeEnvironmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
