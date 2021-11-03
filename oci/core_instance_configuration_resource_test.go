// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var (
	InstanceConfigurationWithPlatformConfigDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, SubnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		OciImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		VolumeBackupPolicyDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig

	instanceConfigurationWithPlatformConfigInstanceDetailsLaunchRepresentation = map[string]interface{}{
		"instance_type":  Representation{RepType: Required, Create: `compute`},
		"launch_details": RepresentationGroup{Optional, instanceConfigurationWithPlatformConfigInstanceDetailsLaunchDetailsRepresentation},
	}
	instanceConfigurationWithPlatformConfigInstanceDetailsLaunchDetailsRepresentation = RepresentationCopyWithRemovedProperties(RepresentationCopyWithNewProperties(instanceConfigurationInstanceDetailsLaunchDetailsRepresentation, map[string]interface{}{
		"shape":           Representation{RepType: Optional, Create: `BM.DenseIO.E4.128`},
		"platform_config": RepresentationGroup{Optional, instancePlatformConfigRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceLaunchOptionsRepresentationForInstanceConfiguration = map[string]interface{}{
		"network_type": Representation{RepType: Optional, Create: `PARAVIRTUALIZED`},
	}
)

// issue-routing-tag: core/computeImaging
func TestAccCoreInstanceConfigurationResource_platformConfig(t *testing.T) {
	if strings.Contains(GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAccCoreInstanceConfigurationResource_platformConfig") {
		t.Skip("Skipping suppressed TestAccCoreInstanceConfigurationResource_platformConfig")
	}

	config := `
        provider oci {
            test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
        }
    ` + CommonTestVariables()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_configuration.test_instance_configuration"

	ResourceTest(t, testAccCheckCoreInstanceConfigurationDestroy, []resource.TestStep{
		// Create with platform config
		{
			Config: config + compartmentIdVariableStr + InstanceConfigurationWithPlatformConfigDependencies +
				GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
					GetUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional, instanceConfigurationWithPlatformConfigInstanceDetailsLaunchRepresentation}, instanceConfigurationRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_core_instance_configurations", "test_instance_configurations", Required, Create, instanceConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceConfigurationWithPlatformConfigDependencies +
				GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
					GetUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional, instanceConfigurationWithPlatformConfigInstanceDetailsLaunchRepresentation}, instanceConfigurationRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_core_instance_configurations", "test_instance_configurations", Required, Create, instanceConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceConfigurationWithPlatformConfigDependencies +
				GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
					GetUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional, instanceConfigurationWithPlatformConfigInstanceDetailsLaunchRepresentation}, instanceConfigurationRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
