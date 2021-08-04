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
	managedInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
	}

	managedInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `osms-instance`},
		"os_family":      Representation{repType: Optional, create: `ALL`},
	}

	parentSourceCreateDisplayName = randomString(10, charsetWithoutDigits)
	childSourceCreateDisplayName  = randomString(10, charsetWithoutDigits)
	groupCreateDisplayName        = randomString(10, charsetWithoutDigits)

	ManagedInstanceResourceConfig = ManagedInstanceManagementResourceDependencies + generateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_parent_software_source", Required, Create, getMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
		[]interface{}{
			Representation{repType: Required, create: `X86_64`},
			Representation{repType: Required, create: parentSourceCreateDisplayName},
		}, softwareSourceRepresentation)) +
		generateResourceFromRepresentationMap("oci_osmanagement_software_source", "test_child_software_source", Required, Create, representationCopyWithNewProperties(getMultipleUpdatedRepresenationCopy([]string{"arch_type", "display_name"},
			[]interface{}{
				Representation{repType: Required, create: `X86_64`},
				Representation{repType: Required, create: childSourceCreateDisplayName},
			}, softwareSourceRepresentation),
			map[string]interface{}{
				"parent_id": Representation{repType: Required, create: `${oci_osmanagement_software_source.test_parent_software_source.id}`},
			})) + generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, getUpdatedRepresentationCopy("display_name", Representation{repType: Required, create: groupCreateDisplayName}, managedInstanceGroupRepresentation))
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_osmanagement_managed_instances.test_managed_instances"
	singularDatasourceName := "data.oci_osmanagement_managed_instance.test_managed_instance"

	resourceName := "oci_osmanagement_managed_instance_management.test_managed_instance_management"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create dependencies
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceConfig,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] OS Management Resource should be created after 5 minutes as OS Agent takes time to activate")
					time.Sleep(5 * time.Minute)
					return nil
				},
			},
			// verify create
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceConfig +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instances", "test_managed_instances", Required, Create, managedInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceResourceConfig +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Required, Create, managedInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceResourceConfig +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "bug_updates_available"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "enhancement_updates_available"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reboot_required"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "last_boot"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "last_checkin"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_family"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_kernel_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "other_updates_available"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_job_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "security_updates_available"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "updates_available"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "work_request_count"),
				),
			},
		},
	})
}
