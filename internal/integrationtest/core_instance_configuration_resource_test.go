// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	InstanceConfigurationWithPlatformConfigDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		utils.VolumeBackupPolicyDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig

	instanceConfigurationWithPlatformConfigInstanceDetailsLaunchRepresentation = map[string]interface{}{
		"instance_type":  acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationWithPlatformConfigInstanceDetailsLaunchDetailsRepresentation},
	}
	instanceConfigurationWithPlatformConfigInstanceDetailsLaunchDetailsRepresentation = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation, map[string]interface{}{
		"shape":           acctest.Representation{RepType: acctest.Optional, Create: `BM.DenseIO.E4.128`},
		"platform_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instancePlatformConfigRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceLaunchOptionsRepresentationForInstanceConfiguration = map[string]interface{}{
		"network_type": acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
	}
)

// issue-routing-tag: core/computeImaging
func TestAccCoreInstanceConfigurationResource_platformConfig(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAccCoreInstanceConfigurationResource_platformConfig") {
		t.Skip("Skipping suppressed TestAccCoreInstanceConfigurationResource_platformConfig")
	}

	config := `
        provider oci {
            test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
        }
    ` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_configuration.test_instance_configuration"

	acctest.ResourceTest(t, testAccCheckCoreInstanceConfigurationDestroy, []resource.TestStep{
		// Create with platform config
		{
			Config: config + compartmentIdVariableStr + InstanceConfigurationWithPlatformConfigDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationWithPlatformConfigInstanceDetailsLaunchRepresentation}, CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.0.numa_nodes_per_socket", "NPS1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.0.type", "AMD_MILAN_BM"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape", "BM.DenseIO.E4.128"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_configurations", "test_instance_configurations", acctest.Required, acctest.Create, CoreCoreInstanceConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceConfigurationWithPlatformConfigDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationWithPlatformConfigInstanceDetailsLaunchRepresentation}, CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.0.numa_nodes_per_socket", "NPS1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.0.type", "AMD_MILAN_BM"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape", "BM.DenseIO.E4.128"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_configurations", "test_instance_configurations", acctest.Required, acctest.Create, CoreCoreInstanceConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceConfigurationWithPlatformConfigDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationWithPlatformConfigInstanceDetailsLaunchRepresentation}, CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.0.numa_nodes_per_socket", "NPS1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.platform_config.0.type", "AMD_MILAN_BM"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape", "BM.DenseIO.E4.128"),
			),
		},
	})
}
