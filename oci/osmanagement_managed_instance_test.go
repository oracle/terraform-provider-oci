// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagedInstanceRequiredOnlyResource = ManagedInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Required, Create, managedInstanceRepresentation)

	ManagedInstanceResourceConfig = ManagedInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Optional, Update, managedInstanceRepresentation)

	managedInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
	}

	managedInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"os_family":      Representation{repType: Optional, create: `LINUX`},
		"filter":         RepresentationGroup{Required, managedInstanceDataSourceFilterRepresentation}}

	managedInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_instance.test_instance.id}`}},
	}

	managedInstanceRepresentation = map[string]interface{}{
		"managed_instance_id":           Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"is_data_collection_authorized": Representation{repType: Optional, create: `false`, update: `true`},
		"notification_topic_id":         Representation{repType: Optional, create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}

	ManagedInstanceResourceDependencies = ManagedInstanceManagementResourceDependencies + generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_osmanagement_managed_instance.test_managed_instance"
	datasourceName := "data.oci_osmanagement_managed_instances.test_managed_instances"
	singularDatasourceName := "data.oci_osmanagement_managed_instance.test_managed_instance"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ManagedInstanceResourceDependencies+
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Optional, Create, managedInstanceRepresentation), "osmanagement", "managedInstance", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create dependencies
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] OS Management Resource should be created after 5 minutes as OS Agent takes time to activate")
					time.Sleep(5 * time.Minute)
					return nil
				},
			},
			// verify create
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Required, Create, managedInstanceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Optional, Create, managedInstanceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_data_collection_authorized", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Optional, Update, managedInstanceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_data_collection_authorized", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instances", "test_managed_instances", Optional, Update, managedInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Optional, Update, managedInstanceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "os_family", "LINUX"),

					resource.TestCheckResourceAttr(datasourceName, "managed_instances.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.is_reboot_required"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.last_boot"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.last_checkin"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.os_family"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.status"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_instances.0.updates_available"),
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

					resource.TestCheckResourceAttr(singularDatasourceName, "autonomous.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bug_updates_available"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "enhancement_updates_available"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_data_collection_authorized", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reboot_required"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ksplice_effective_kernel_version"),
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
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"managed_instance_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}
