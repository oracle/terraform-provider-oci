// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagedInstanceManagementRequiredOnlyResource = ManagedInstanceGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, managedInstanceGroupRepresentation)

	ManagedInstanceManagementResourceConfig = ManagedInstanceGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Optional, Update, managedInstanceGroupRepresentation)

	ManagedInstanceManagementRepresentation = map[string]interface{}{
		"managed_instance_id":    Representation{repType: Required, create: `${var.managed_instance_id}`},
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

	ManagedInstanceManagementResourceDependencies = generateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_parent_software_source", Required, Create, getUpdatedRepresentationCopy("arch_type", Representation{repType: Required, create: `X86_64`}, softwareSourceRepresentation)) +
		generateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_child_software_source", Required, Create, representationCopyWithNewProperties(getMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
			[]interface{}{
				Representation{repType: Required, create: `X86_64`},
				Representation{repType: Required, create: softwareSourceUpdateDisplayName},
			}, softwareSourceRepresentation),
			map[string]interface{}{
				"parent_id": Representation{repType: Required, create: `${oci_osmanagement_software_source.test_parent_software_source.id}`},
			})) + generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, managedInstanceGroupRepresentation)
)

func TestOsmanagementManagedInstanceManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedInstanceId := getEnvSettingWithBlankDefault("managed_instance_id")
	managedInstanceIdVariableStr := fmt.Sprintf("variable \"managed_instance_id\" { default = \"%s\" }\n", managedInstanceId)

	resourceName := "oci_osmanagement_managed_instance_management.test_managed_instance_management"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + ManagedInstanceManagementResourceDependencies +
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
				Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + ManagedInstanceManagementResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + ManagedInstanceManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Optional, Create, ManagedInstanceManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				),
			},
			// verify update with optionals
			{
				Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + ManagedInstanceManagementResourceDependencies +
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
