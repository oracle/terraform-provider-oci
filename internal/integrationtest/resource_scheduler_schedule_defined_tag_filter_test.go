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
	ResourceSchedulerScheduleWithResourceFilterRequiredOnlyResource = ResourceSchedulerScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, DefinedTagsFilterRepresentation)

	DefinedTagsFilterSingularDataSourceRepresentation = map[string]interface{}{
		"schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_scheduler_schedule.test_schedule.id}`},
	}

	DefinedTagsFilterDataSourceRepresentation = map[string]interface{}{
		// must include at least one of `compartmentId` and `schedule_id`
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"schedule_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_scheduler_schedule.test_schedule.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `provider displayName1`, Update: `provider displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ResourceSchedulerScheduleDataSourceFilterRepresentation}}

	ResourceSchedulerScheduleResourceDefinedTagsFiltersRepresentation = map[string]interface{}{
		"attribute": acctest.Representation{RepType: acctest.Required, Create: `DEFINED_TAGS`, Update: `DEFINED_TAGS`},
		"value":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourceSchedulerScheduleResourceDefinedTagsFiltersValueRepresentation},
	}
	ResourceSchedulerScheduleResourceDefinedTagsFiltersValueRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `ResourceSchedulerCanary`, Update: `Test`},
		"tag_key":   acctest.Representation{RepType: acctest.Required, Create: `ScheduleTagFilterTestKey`, Update: `ORM-123837`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `foo`, Update: `inaccessible`},
	}

	DefinedTagsFilterRepresentation = map[string]interface{}{
		// Required
		"action":             acctest.Representation{RepType: acctest.Required, Create: `START_RESOURCE`, Update: `STOP_RESOURCE`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"recurrence_details": acctest.Representation{RepType: acctest.Required, Create: `FREQ=DAILY;INTERVAL=1`, Update: `FREQ=DAILY;INTERVAL=2`},
		"recurrence_type":    acctest.Representation{RepType: acctest.Required, Create: `ICAL`, Update: `ICAL`},
		"resource_filters":   acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourceSchedulerScheduleResourceDefinedTagsFiltersRepresentation},
		// Optionals
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `provider description1`, Update: `provider description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `provider displayName1`, Update: `provider displayName2`},
		"time_ends":     acctest.Representation{RepType: acctest.Optional, Create: `2024-06-22T00:00:00Z`, Update: `2024-06-24T00:00:00Z`},
		"time_starts":   acctest.Representation{RepType: acctest.Optional, Create: `2024-06-16T00:00:00Z`, Update: `2024-06-18T00:00:00Z`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"state":         acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangesDefinedTagsResourceSchedulerRepresentation},
	}
)

func TestResourceSchedulerScheduleDefinedTagResourceFilter(t *testing.T) {
	httpreplay.SetScenario("TestResourceSchedulerScheduleDefinedTagResourceFilter")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_resource_scheduler_schedule.test_schedule"
	singularDatasourceName := "data.oci_resource_scheduler_schedule.test_schedule"
	datasourceName := "data.oci_resource_scheduler_schedules.test_schedules"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ResourceSchedulerScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Optional, acctest.Create, DefinedTagsFilterRepresentation), "resourcescheduler", "schedule", t)

	acctest.ResourceTest(t, testAccCheckResourceSchedulerScheduleDestroy, []resource.TestStep{
		//verify Create with Required - resourceFilter
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, DefinedTagsFilterRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),

				resource.TestCheckResourceAttr(resourceName, "resources.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.attribute", "DEFINED_TAGS"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.namespace", "ResourceSchedulerCanary"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.tag_key", "ScheduleTagFilterTestKey"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.value", "foo"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//delete before next Create
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies,
		},

		// verify create with optionals - resourceFilter
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Optional, acctest.Create, DefinedTagsFilterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),

				resource.TestCheckResourceAttr(resourceName, "description", "provider description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provider displayName1"),
				resource.TestCheckResourceAttr(resourceName, "time_ends", "2024-06-22T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "time_starts", "2024-06-16T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

				resource.TestCheckResourceAttr(resourceName, "resources.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.attribute", "DEFINED_TAGS"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.namespace", "ResourceSchedulerCanary"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.tag_key", "ScheduleTagFilterTestKey"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.value", "foo"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, DefinedTagsFilterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResourceSchedulerScheduleWithResourceFilterRequiredOnlyResource,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recurrence_type", "ICAL"),

				resource.TestCheckResourceAttr(singularDatasourceName, "resources.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.attribute", "DEFINED_TAGS"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.namespace", "ResourceSchedulerCanary"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.tag_key", "ScheduleTagFilterTestKey"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.value", "foo"),
			),
		},

		// verify datasources
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, DefinedTagsFilterRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_scheduler_schedules", "test_schedules", acctest.Required, acctest.Create, DefinedTagsFilterDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "schedule_collection.#", "1"),
			),
		},

		// verify resource import
		{
			Config:                  config + ResourceSchedulerScheduleWithResourceFilterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"time_next_run"},
			ResourceName:            resourceName,
		},
	})
}
