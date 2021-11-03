// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagedInstanceManagementRequiredOnlyResource = ManagedInstanceGroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, managedInstanceGroupRepresentation)

	ManagedInstanceManagementResourceConfig = ManagedInstanceGroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Optional, Update, managedInstanceGroupRepresentation)

	ManagedInstanceManagementRepresentation = map[string]interface{}{
		"managed_instance_id":    Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
		"parent_software_source": RepresentationGroup{Optional, parentSoftwareSourceRepresentation},
	}

	parentSoftwareSourceRepresentation = map[string]interface{}{
		"name": Representation{RepType: Required, Create: `${oci_osmanagement_software_source.test_parent_software_source.display_name}`},
		"id":   Representation{RepType: Required, Create: `${oci_osmanagement_software_source.test_parent_software_source.id}`},
	}

	childSoftwareSourceRepresentation = map[string]interface{}{
		"name": Representation{RepType: Required, Create: `${oci_osmanagement_software_source.test_child_software_source.display_name}`},
		"id":   Representation{RepType: Required, Create: `${oci_osmanagement_software_source.test_child_software_source.id}`},
	}

	osmsInstanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": Representation{RepType: Required, Create: `false`, Update: `false`},
		"is_management_disabled":   Representation{RepType: Required, Create: `false`, Update: `false`},
		"is_monitoring_disabled":   Representation{RepType: Required, Create: `false`, Update: `false`},
		"plugins_config":           RepresentationGroup{Required, autonomousInstanceAgentConfigPluginsConfigRepresentation},
	}

	autonomousInstanceAgentConfigPluginsConfigRepresentation = map[string]interface{}{
		"desired_state": Representation{RepType: Required, Create: `ENABLED`},
		"name":          Representation{RepType: Required, Create: `Oracle Autonomous Linux`},
	}

	managedInstanceGroupsRepresentation = map[string]interface{}{
		"display_name": Representation{RepType: Required, Create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.display_name}`},
		"id":           Representation{RepType: Required, Create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`},
	}

	vnicDetailsRepresentation = map[string]interface{}{
		"assign_public_ip": Representation{RepType: Required, Create: `true`},
		"subnet_id":        Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	routeTablesRepresentation = map[string]interface{}{
		"manage_default_resource_id": Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`},
		"route_rules":                RepresentationGroup{Required, routeRulesRepresentation},
	}

	routeRulesRepresentation = map[string]interface{}{
		"cidr_block":        Representation{RepType: Required, Create: `0.0.0.0/0`},
		"network_entity_id": Representation{RepType: Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
	}

	sourceDetailsRepresentation = map[string]interface{}{
		"source_type": Representation{RepType: Required, Create: `image`},
		"source_id":   Representation{RepType: Required, Create: `${var.OsManagedImageOCID[var.region]}`},
	}

	parentSourceDisplayName = RandomString(10, CharsetWithoutDigits)
	childSourceDisplayName  = RandomString(10, CharsetWithoutDigits)
	groupDisplayName        = RandomString(10, CharsetWithoutDigits)

	osmanagementSoftwareSourceRepresentation = GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_parent_software_source", Required, Create, GetMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
		[]interface{}{
			Representation{RepType: Required, Create: `X86_64`},
			Representation{RepType: Required, Create: parentSourceDisplayName},
		}, softwareSourceRepresentation)) +
		GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_child_software_source", Required, Create, RepresentationCopyWithNewProperties(GetMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
			[]interface{}{
				Representation{RepType: Required, Create: `X86_64`},
				Representation{RepType: Required, Create: childSourceDisplayName},
			}, softwareSourceRepresentation),
			map[string]interface{}{
				"parent_id": Representation{RepType: Required, Create: `${oci_osmanagement_software_source.test_parent_software_source.id}`},
			})) + GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, GetUpdatedRepresentationCopy("display_name", Representation{RepType: Required, Create: groupDisplayName}, managedInstanceGroupRepresentation))

	ManagedInstanceManagementResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, RepresentationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
		"create_vnic_details": RepresentationGroup{Required, vnicDetailsRepresentation},
		"source_details":      RepresentationGroup{Required, sourceDetailsRepresentation},
		"agent_config":        RepresentationGroup{Required, osmsInstanceAgentConfigRepresentation},
		"image":               Representation{RepType: Required, Create: `${var.OsManagedImageOCID[var.region]}`},
	})) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(VcnRepresentation, map[string]interface{}{
			"dns_label":  Representation{RepType: Required, Create: `testvcn`},
			"cidr_block": Representation{RepType: Required, Create: `10.1.0.0/16`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", Required, Create, routeTablesRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, RepresentationCopyWithNewProperties(SubnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.1.20.0/24`}, "dns_label": Representation{RepType: Required, Create: `testsubnet`}, "route_table_id": Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`}})) +
		AvailabilityDomainConfig + OsManagedImageIdsVariable
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_osmanagement_managed_instance_management.test_managed_instance_management"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ManagedInstanceManagementResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation), "osmanagement", "managedInstanceManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] OS Management Resource should be created after 2 minutes as OS Agent takes time to activate")
				time.Sleep(5 * time.Minute)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation +
				GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),

				func(s *terraform.State) (err error) {
					_, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation +
				GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Optional, Create, ManagedInstanceManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
			),
		},
		// verify Update with optionals
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation +
				GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Optional, Create, RepresentationCopyWithNewProperties(ManagedInstanceManagementRepresentation, map[string]interface{}{
					"child_software_sources":  RepresentationGroup{Optional, childSoftwareSourceRepresentation},
					"managed_instance_groups": RepresentationGroup{Optional, managedInstanceGroupsRepresentation},
				})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
			),
		},
	})
}
