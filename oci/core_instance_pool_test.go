// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	InstancePoolRequiredOnlyResource = InstancePoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolRepresentation)

	InstancePoolResourceConfig = InstancePoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation)

	instancePoolSingularDataSourceRepresentation = map[string]interface{}{
		"instance_pool_id": Representation{repType: Required, create: `${oci_core_instance_pool.test_instance_pool.id}`},
	}

	instancePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `backend-servers-pool`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `RUNNING`},
		"filter":         RepresentationGroup{Required, instancePoolDataSourceFilterRepresentation}}
	instancePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_instance_pool.test_instance_pool.id}`}},
	}

	instancePoolRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_configuration_id": Representation{repType: Required, create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  RepresentationGroup{Required, instancePoolPlacementConfigurationsRepresentation},
		"size":                      Representation{repType: Required, create: `2`, update: `3`},
		"state":                     Representation{repType: Optional, create: `Running`},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              Representation{repType: Optional, create: `backend-servers-pool`, update: `displayName2`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	instancePoolPlacementConfigurationsRepresentation = map[string]interface{}{
		"availability_domain":    Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"secondary_vnic_subnets": RepresentationGroup{Optional, instancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation},
	}
	instancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation = map[string]interface{}{
		"subnet_id": Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		//the display_name should be the same as in the instance configuration
		"display_name": Representation{repType: Required, create: `backend-servers-pool`},
	}

	instanceConfigurationPoolRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_details": RepresentationGroup{Required, instanceConfigurationInstanceDetailsPoolRepresentation},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     Representation{repType: Optional, create: `backend-servers`, update: `displayName2`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	instanceConfigurationInstanceDetailsPoolRepresentation = map[string]interface{}{
		"instance_type":   Representation{repType: Required, create: `compute`},
		"secondary_vnics": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation},
		"launch_details":  RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentationn},
	}
	instanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation = map[string]interface{}{
		"create_vnic_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		//the display_name should be the same as in the secondary_vnic_subnets
		"display_name": Representation{repType: Optional, create: `backend-servers-pool`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentationn = map[string]interface{}{
		"compartment_id":      Representation{repType: Optional, create: `${var.compartment_id}`},
		"create_vnic_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{repType: Optional, create: `backend-servers`},
		"extended_metadata":   Representation{repType: Optional, create: map[string]string{"extendedMetadata": "extendedMetadata"}, update: map[string]string{"extendedMetadata2": "extendedMetadata2"}},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"ipxe_script":         Representation{repType: Optional, create: `ipxeScript`},
		"metadata":            Representation{repType: Optional, create: map[string]string{"metadata": "metadata"}, update: map[string]string{"metadata2": "metadata2"}},
		"shape":               Representation{repType: Optional, create: InstanceConfigurationVmShape},
		"source_details":      RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation = map[string]interface{}{
		"assign_public_ip":       Representation{repType: Optional, create: `true`},
		"display_name":           Representation{repType: Optional, create: `backend-servers`},
		"skip_source_dest_check": Representation{repType: Optional, create: `false`},
	}

	InstancePoolResourceDependenciesWithoutSecondaryVnic = ImageRequiredOnlyResource +
		`
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}` +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
			getUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional,
				representationCopyWithRemovedProperties(instanceConfigurationInstanceDetailsPoolRepresentation, []string{"secondary_vnics"})}, instanceConfigurationPoolRepresentation))

	InstancePoolResourceDependencies = ImageRequiredOnlyResource +
		`
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}` +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create, instanceConfigurationPoolRepresentation)
)

func TestCoreInstancePoolResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_pool.test_instance_pool"
	datasourceName := "data.oci_core_instance_pools.test_instance_pools"
	singularDatasourceName := "data.oci_core_instance_pool.test_instance_pool"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstancePoolDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependenciesWithoutSecondaryVnic +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "size", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Create, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers-pool"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "size", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "size", "3"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
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
			// verify stop the Instance Pool
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update,
						getUpdatedRepresentationCopy("state", Representation{repType: Optional, create: "Stopped"}, instancePoolRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "size", "3"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
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
			// verify start the Instance Pool
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "size", "3"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
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
			// verify datasource the state will be updated to RUNNING
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance_pools", "test_instance_pools", Optional, Update, instancePoolDataSourceRepresentation) +
					compartmentIdVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "instance_pools.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.instance_configuration_id"),
					resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.size", "3"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstancePoolResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "size", "3"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreInstancePoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeManagementClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance_pool" {
			noResourceFound = false
			request := oci_core.GetInstancePoolRequest{}

			tmp := rs.Primary.ID
			request.InstancePoolId = &tmp

			response, err := client.GetInstancePool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.InstancePoolLifecycleStateTerminated): true,
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
