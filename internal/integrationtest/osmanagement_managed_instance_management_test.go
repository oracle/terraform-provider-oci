// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagedInstanceManagementRequiredOnlyResource = ManagedInstanceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, managedInstanceGroupRepresentation)

	ManagedInstanceManagementResourceConfig = ManagedInstanceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, managedInstanceGroupRepresentation)

	ManagedInstanceManagementRepresentation = map[string]interface{}{
		"managed_instance_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"parent_software_source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: parentSoftwareSourceRepresentation},
	}

	parentSoftwareSourceRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_software_source.test_parent_software_source.display_name}`},
		"id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_software_source.test_parent_software_source.id}`},
	}

	childSoftwareSourceRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_software_source.test_child_software_source.display_name}`},
		"id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_software_source.test_child_software_source.id}`},
	}

	osmsInstanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_management_disabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_monitoring_disabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"plugins_config":           acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousInstanceAgentConfigPluginsConfigRepresentation},
	}

	autonomousInstanceAgentConfigPluginsConfigRepresentation = map[string]interface{}{
		"desired_state": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
		"name":          acctest.Representation{RepType: acctest.Required, Create: `Oracle Autonomous Linux`},
	}

	managedInstanceGroupsRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.display_name}`},
		"id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`},
	}

	vnicDetailsRepresentation = map[string]interface{}{
		"assign_public_ip": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"subnet_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	routeTablesRepresentation = map[string]interface{}{
		"manage_default_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`},
		"route_rules":                acctest.RepresentationGroup{RepType: acctest.Required, Group: routeRulesRepresentation},
	}

	routeRulesRepresentation = map[string]interface{}{
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
	}

	sourceDetailsRepresentation = map[string]interface{}{
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `image`},
		"source_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.OsManagedImageOCID[var.region]}`},
	}

	parentSourceDisplayName = utils.RandomString(10, utils.CharsetWithoutDigits)
	childSourceDisplayName  = utils.RandomString(10, utils.CharsetWithoutDigits)
	groupDisplayName        = utils.RandomString(10, utils.CharsetWithoutDigits)

	osmanagementSoftwareSourceRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_parent_software_source", acctest.Required, acctest.Create, acctest.GetMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
		[]interface{}{
			acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
			acctest.Representation{RepType: acctest.Required, Create: parentSourceDisplayName},
		}, softwareSourceRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_child_software_source", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
			[]interface{}{
				acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
				acctest.Representation{RepType: acctest.Required, Create: childSourceDisplayName},
			}, softwareSourceRepresentation),
			map[string]interface{}{
				"parent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_software_source.test_parent_software_source.id}`},
			})) + acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("display_name", acctest.Representation{RepType: acctest.Required, Create: groupDisplayName}, managedInstanceGroupRepresentation))

	ManagedInstanceManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: vnicDetailsRepresentation},
		"source_details":      acctest.RepresentationGroup{RepType: acctest.Required, Group: sourceDetailsRepresentation},
		"agent_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: osmsInstanceAgentConfigRepresentation},
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.OsManagedImageOCID[var.region]}`},
	})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label":  acctest.Representation{RepType: acctest.Required, Create: `testvcn`},
			"cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.1.0.0/16`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", acctest.Required, acctest.Create, routeTablesRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.1.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `testsubnet`}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`}})) +
		AvailabilityDomainConfig + utils.OsManagedImageIdsVariable
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_osmanagement_managed_instance_management.test_managed_instance_management"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagedInstanceManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", acctest.Required, acctest.Create, ManagedInstanceManagementRepresentation), "osmanagement", "managedInstanceManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
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
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", acctest.Required, acctest.Create, ManagedInstanceManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", acctest.Optional, acctest.Create, ManagedInstanceManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
			),
		},
		// verify Update with optionals
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ManagedInstanceManagementRepresentation, map[string]interface{}{
					"child_software_sources":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: childSoftwareSourceRepresentation},
					"managed_instance_groups": acctest.RepresentationGroup{RepType: acctest.Optional, Group: managedInstanceGroupsRepresentation},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
			),
		},
	})
}
