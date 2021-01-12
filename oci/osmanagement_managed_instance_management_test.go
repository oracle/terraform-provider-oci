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
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, managedInstanceGroupRepresentation)

	ManagedInstanceManagementResourceConfig = ManagedInstanceGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Optional, Update, managedInstanceGroupRepresentation)

	ManagedInstanceManagementRepresentation = map[string]interface{}{
		"managed_instance_id":    Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"parent_software_source": RepresentationGroup{Optional, parentSoftwareSourceRepresentation},
	}

	parentSoftwareSourceRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `${oci_osmanagement_software_source.test_parent_software_source.display_name}`},
		"id":   Representation{repType: Required, create: `${oci_osmanagement_software_source.test_parent_software_source.id}`},
	}

	childSoftwareSourceRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `${oci_osmanagement_software_source.test_child_software_source.display_name}`},
		"id":   Representation{repType: Required, create: `${oci_osmanagement_software_source.test_child_software_source.id}`},
	}

	managedInstanceGroupsRepresentation = map[string]interface{}{
		"display_name": Representation{repType: Required, create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.display_name}`},
		"id":           Representation{repType: Required, create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`},
	}

	vnicDetailsRepresentation = map[string]interface{}{
		"assign_public_ip": Representation{repType: Required, create: `true`},
		"subnet_id":        Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
	}

	routeTablesRepresentation = map[string]interface{}{
		"manage_default_resource_id": Representation{repType: Required, create: `${oci_core_vcn.test_vcn.default_route_table_id}`},
		"route_rules":                RepresentationGroup{Required, routeRulesRepresentation},
	}

	routeRulesRepresentation = map[string]interface{}{
		"cidr_block":        Representation{repType: Required, create: `0.0.0.0/0`},
		"network_entity_id": Representation{repType: Required, create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
	}

	sourceDetailsRepresentation = map[string]interface{}{
		"source_type": Representation{repType: Required, create: `image`},
		"source_id":   Representation{repType: Required, create: `${var.OsManagedImageOCID[var.region]}`},
	}

	parentSourceDisplayName = randomString(10, charsetWithoutDigits)
	childSourceDisplayName  = randomString(10, charsetWithoutDigits)
	groupDisplayName        = randomString(10, charsetWithoutDigits)

	osmanagementSoftwareSourceRepresentation = generateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_parent_software_source", Required, Create, getMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
		[]interface{}{
			Representation{repType: Required, create: `X86_64`},
			Representation{repType: Required, create: parentSourceDisplayName},
		}, softwareSourceRepresentation)) +
		generateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_child_software_source", Required, Create, representationCopyWithNewProperties(getMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
			[]interface{}{
				Representation{repType: Required, create: `X86_64`},
				Representation{repType: Required, create: childSourceDisplayName},
			}, softwareSourceRepresentation),
			map[string]interface{}{
				"parent_id": Representation{repType: Required, create: `${oci_osmanagement_software_source.test_parent_software_source.id}`},
			})) + generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, getUpdatedRepresentationCopy("display_name", Representation{repType: Required, create: groupDisplayName}, managedInstanceGroupRepresentation))

	ManagedInstanceManagementResourceDependencies = generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, representationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
		"create_vnic_details": RepresentationGroup{Required, vnicDetailsRepresentation},
		"source_details":      RepresentationGroup{Required, sourceDetailsRepresentation},
		"image":               Representation{repType: Required, create: `${var.OsManagedImageOCID[var.region]}`},
	})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label":  Representation{repType: Required, create: `testvcn`},
			"cidr_block": Representation{repType: Required, create: `10.1.0.0/16`},
		})) +
		generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		generateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", Required, Create, routeTablesRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.1.20.0/24`}, "dns_label": Representation{repType: Required, create: `testsubnet`}, "route_table_id": Representation{repType: Required, create: `${oci_core_vcn.test_vcn.default_route_table_id}`}})) +
		AvailabilityDomainConfig + OsManagedImageIdsVariable
)

func TestOsmanagementManagedInstanceManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_osmanagement_managed_instance_management.test_managed_instance_management"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create dependencies
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] OS Management Resource should be created after 5 minutes as OS Agent takes time to activate")
					time.Sleep(5 * time.Minute)
					return nil
				},
			},
			// verify create
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Optional, Create, ManagedInstanceManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				),
			},
			// verify update with optionals
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceManagementResourceDependencies + osmanagementSoftwareSourceRepresentation +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Optional, Create, representationCopyWithNewProperties(ManagedInstanceManagementRepresentation, map[string]interface{}{
						"child_software_sources":  RepresentationGroup{Optional, childSoftwareSourceRepresentation},
						"managed_instance_groups": RepresentationGroup{Optional, managedInstanceGroupsRepresentation},
					})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				),
			},
		},
	})
}
