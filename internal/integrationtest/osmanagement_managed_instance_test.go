// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OsmanagementManagedInstanceRequiredOnlyResource = OsmanagementManagedInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsmanagementManagedInstanceRepresentation)

	OsmanagementManagedInstanceResourceConfig = OsmanagementManagedInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, OsmanagementManagedInstanceRepresentation)

	OsmanagementOsmanagementManagedInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	OsmanagementOsmanagementManagedInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"os_family":      acctest.Representation{RepType: acctest.Optional, Create: `LINUX`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OsmanagementManagedInstanceDataSourceFilterRepresentation}}

	OsmanagementManagedInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance.test_instance.id}`}},
	}

	OsmanagementManagedInstanceRepresentation = map[string]interface{}{
		"managed_instance_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"is_data_collection_authorized": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"notification_topic_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}

	OsmanagementManagedInstanceResourceDependencies = ManagedInstanceManagementResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsmanagementManagedInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Create, OsmanagementManagedInstanceRepresentation), "osmanagement", "managedInstance", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] OS Management Resource should be created after 5 minutes as OS Agent takes time to activate")
				time.Sleep(5 * time.Minute)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsmanagementManagedInstanceRepresentation),
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
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Create, OsmanagementManagedInstanceRepresentation),
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
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, OsmanagementManagedInstanceRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instances", "test_managed_instances", acctest.Optional, acctest.Update, OsmanagementOsmanagementManagedInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + OsmanagementManagedInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, OsmanagementManagedInstanceRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsmanagementOsmanagementManagedInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsmanagementManagedInstanceResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", acctest.Required, acctest.Create, OsmanagementManagedInstanceManagementRepresentation),
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
		// verify resource import
		{
			Config:            config + OsmanagementManagedInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"managed_instance_id",
			},
			ResourceName: resourceName,
		},
	})
}
