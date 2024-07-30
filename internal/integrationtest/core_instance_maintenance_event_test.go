// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreInstanceMaintenanceEventSingularDataSourceRepresentation = map[string]interface{}{
		"instance_maintenance_event_id": acctest.Representation{RepType: acctest.Required, Create: `${var.instance_maintenance_event_id}`},
	}

	CoreInstanceMaintenanceEventDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_action": acctest.Representation{RepType: acctest.Optional, Create: `REBOOT_MIGRATION`},
		"instance_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.instance_id}`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `SCHEDULED`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceMaintenanceEventDataSourceFilterRepresentation}}
	CoreInstanceMaintenanceEventDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance_maintenance_event.test_instance_maintenance_event.id}`}},
	}

	CoreInstanceMaintenanceEventRepresentation = map[string]interface{}{
		"instance_maintenance_event_id": acctest.Representation{RepType: acctest.Required, Create: `${var.instance_maintenance_event_id}`},
		"can_delete_local_storage":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"time_window_start":             acctest.Representation{RepType: acctest.Optional, Create: `2006-01-02T15:04:05Z`, Update: `2025-01-12T15:04:05Z`},
	}

	CoreInstanceMaintenanceEventResourceDependencies = utils.OciImageIdsVariable
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_instance_maintenance_event", "test_instance_maintenance_event", acctest.Required, acctest.Create, CoreInstanceMaintenanceEventRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
	//		"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
	//	})) +
	//		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
	//			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
	//		})) +
	//		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
	//		AvailabilityDomainConfig +
	//		DefinedTagsDependencies
)

// issue-routing-tag: core/default
func TestCoreInstanceMaintenanceEventResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceMaintenanceEventResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	instanceId := utils.GetEnvSettingWithBlankDefault("instance_id")
	instanceIdVariableStr := fmt.Sprintf("variable \"instance_id\" { default = \"%s\" }\n", instanceId)

	instanceMaintenanceEventId := utils.GetEnvSettingWithBlankDefault("instance_maintenance_event_id")
	instanceMaintenanceEventIdVariableStr := fmt.Sprintf("variable \"instance_maintenance_event_id\" { default = \"%s\" }\n", instanceMaintenanceEventId)

	resourceName := "oci_core_instance_maintenance_event.test_instance_maintenance_event"
	acctest.SaveConfigContent(config+compartmentIdVariableStr+instanceIdVariableStr+instanceMaintenanceEventIdVariableStr+CoreInstanceMaintenanceEventResourceDependencies, "core", "instanceMaintenanceEvent", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + instanceIdVariableStr + instanceMaintenanceEventIdVariableStr +
				CoreInstanceMaintenanceEventResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_maintenance_event", "test_instance_maintenance_event", acctest.Optional, acctest.Update, CoreInstanceMaintenanceEventRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_maintenance_event_id"),
				resource.TestCheckResourceAttr(resourceName, "time_window_start", `2025-01-12T15:04:05Z`),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_maintenance_events", "test_instance_maintenance_event", acctest.Optional, acctest.Update, CoreInstanceMaintenanceEventDataSourceRepresentation) +
				compartmentIdVariableStr + instanceIdVariableStr + instanceMaintenanceEventIdVariableStr + CoreInstanceMaintenanceEventResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_maintenance_event", "test_instance_maintenance_event", acctest.Optional, acctest.Update, CoreInstanceMaintenanceEventRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_action", "REBOOT_MIGRATION"),
				resource.TestCheckResourceAttr(resourceName, "instance_id", instanceId),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_maintenance_event", "test_instance_maintenance_event", acctest.Required, acctest.Create, CoreInstanceMaintenanceEventRepresentation) +
				compartmentIdVariableStr + instanceIdVariableStr + instanceMaintenanceEventIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "instance_maintenance_event_id"),
				resource.TestCheckResourceAttr(resourceName, "can_delete_local_storage", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "can_reschedule"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "estimated_duration"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_action"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_category"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_reason"),
				resource.TestCheckResourceAttrSet(resourceName, "start_window_duration"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(resourceName, "time_hard_due_date"),
				resource.TestCheckResourceAttrSet(resourceName, "time_window_start"),
			),
		},
		// verify resource import
		//	{
		//	Config: config + acctest.GenerateResourceFromRepresentationMap("oci_core_instance_maintenance_event", "test_instance_maintenance_event", acctest.Required, acctest.Create, CoreInstanceMaintenanceEventRepresentation) +
		//		compartmentIdVariableStr + instanceIdVariableStr,
		//	ImportState:       true,
		//	ImportStateVerify: true,
		//	ImportStateVerifyIgnore: []string{
		//		"alternative_resolution_action",
		//	},
		//	ResourceName: resourceName,
		//},
	})
}
