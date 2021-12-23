// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagedInstanceRequiredOnlyResource = ManagedInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, managedInstanceRepresentation)

	ManagedInstanceResourceConfig = ManagedInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, managedInstanceRepresentation)

	managedInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	managedInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"os_family":      acctest.Representation{RepType: acctest.Optional, Create: `LINUX`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: managedInstanceDataSourceFilterRepresentation}}

	managedInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance.test_instance.id}`}},
	}

	managedInstanceRepresentation = map[string]interface{}{
		"managed_instance_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"is_data_collection_authorized": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"notification_topic_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}

	ManagedInstanceResourceDependencies = ManagedInstanceManagementResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_osmanagement_managed_instance.test_managed_instance"
	datasourceName := "data.oci_osmanagement_managed_instances.test_managed_instances"
	singularDatasourceName := "data.oci_osmanagement_managed_instance.test_managed_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagedInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Create, managedInstanceRepresentation), "osmanagement", "managedInstance", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] OS Management Resource should be created after 5 minutes as OS Agent takes time to activate")
				time.Sleep(5 * time.Minute)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, managedInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Create, managedInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_data_collection_authorized", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),

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
			Config: config + compartmentIdVariableStr + ManagedInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, managedInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_data_collection_authorized", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instances", "test_managed_instances", acctest.Optional, acctest.Update, managedInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, managedInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, managedInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedInstanceResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", acctest.Required, acctest.Create, ManagedInstanceManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}
