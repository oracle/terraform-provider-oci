// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	computeInstanceUpdateSyncOcid                  = utils.GetEnvSettingWithBlankDefault("computeInstance_ocid")
	ResourceScheduleUpdateSyncRequiredOnlyResource = ResourceSchedulerScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, ResourceScheduleUpdateSyncRepresentation)

	ResourceScheduleUpdateSyncSingularDataSourceRepresentation = map[string]interface{}{
		"schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_scheduler_schedule.test_schedule.id}`},
	}
	ResourceScheduleUpdateSyncDataSourceRepresentation = map[string]interface{}{
		// must include at least one of `compartmentId` and `schedule_id`
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"schedule_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_scheduler_schedule.test_schedule.id}`},
	}

	ResourceScheduleUpdateSyncRepresentation = map[string]interface{}{
		// Required
		"action":             acctest.Representation{RepType: acctest.Required, Create: `START_RESOURCE`, Update: `START_RESOURCE`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"recurrence_details": acctest.Representation{RepType: acctest.Required, Create: `FREQ=DAILY;INTERVAL=1`, Update: `FREQ=DAILY;INTERVAL=1`},
		"recurrence_type":    acctest.Representation{RepType: acctest.Required, Create: `ICAL`, Update: `ICAL`},
		// Must include either `resources` or `resource_filters` when creating schedules
		"resources": acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourceSchedulerScheduleResourcesUpdateSyncRepresentation},
		// Optionals
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `provider description1`, Update: `provider description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `provider displayName1`, Update: `provider displayName2`},
		"time_ends":     acctest.Representation{RepType: acctest.Optional, Create: `2024-11-20T00:00:00Z`, Update: `2024-11-20T00:00:00Z`},
		"time_starts":   acctest.Representation{RepType: acctest.Optional, Create: `2024-11-15T00:00:00Z`, Update: `2024-11-15T00:00:00Z`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Finance"}},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangesDefinedTagsResourceSchedulerRepresentation},
	}

	ResourceSchedulerScheduleResourcesUpdateSyncRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: computeInstanceUpdateSyncOcid, Update: computeInstanceUpdateSyncOcid},
		// mimic customer's behavior
		"metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"metadata": "metadata"}, Update: map[string]string{"metadata": "metadata"}},
	}
)

func TestResourceScheduleUpdateSync(t *testing.T) {
	httpreplay.SetScenario("TestResourceSchedulerScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_resource_scheduler_schedule.test_schedule"
	singularDatasourceName := "data.oci_resource_scheduler_schedule.test_schedule"
	datasourceName := "data.oci_resource_scheduler_schedules.test_schedules"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ResourceSchedulerScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Optional, acctest.Create, ResourceScheduleUpdateSyncRepresentation), "resourcescheduler", "schedule", t)

	acctest.ResourceTest(t, testAccCheckResourceSchedulerScheduleDestroy, []resource.TestStep{
		//verify Create with Required - resource ocids
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, ResourceScheduleUpdateSyncRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),

				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//delete before next Create
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies,
		},

		// verify create with optionals - resourceOCID
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Optional, acctest.Create, ResourceScheduleUpdateSyncRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),
				resource.TestCheckResourceAttr(resourceName, "description", "provider description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provider displayName1"),
				resource.TestCheckResourceAttr(resourceName, "time_ends", "2024-11-20T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "time_starts", "2024-11-15T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.id", computeInstanceUpdateSyncOcid),
				resource.TestCheckResourceAttr(resourceName, "resources.0.metadata.%", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Optional, acctest.Update, ResourceScheduleUpdateSyncRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),
				resource.TestCheckResourceAttr(resourceName, "description", "provider description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provider displayName2"),
				resource.TestCheckResourceAttr(resourceName, "time_ends", "2024-11-20T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "time_starts", "2024-11-15T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.id", computeInstanceOcid),
				resource.TestCheckResourceAttr(resourceName, "resources.0.metadata.%", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Printf("xiaotong printing resId and resId2, %s, %s", resId, resId2)
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, ResourceScheduleUpdateSyncSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResourceScheduleUpdateSyncRequiredOnlyResource,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recurrence_type", "ICAL"),

				resource.TestCheckResourceAttr(singularDatasourceName, "resources.#", "1"),
			),
		},

		// verify datasources
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, ResourceScheduleUpdateSyncRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_scheduler_schedules", "test_schedules", acctest.Required, acctest.Create, ResourceScheduleUpdateSyncDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "schedule_collection.#", "1"),
			),
		},

		// verify resource import
		{
			Config:                  config + ResourceScheduleUpdateSyncRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"time_next_run"},
			ResourceName:            resourceName,
		},
	})
}
