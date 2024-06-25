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
	multipleConfig = "resource \"oci_resource_scheduler_schedule\" \"test_schedule\" {\n\taction = \"START_RESOURCE\"\n\tcompartment_id = \"${var.compartment_id}\"\n\trecurrence_details = \"FREQ=DAILY;INTERVAL=1\"\n\trecurrence_type = \"ICAL\"\n\tdescription = \"provider description1\"\n\tdisplay_name = \"provider displayName1\"\n\ttime_starts = \"2024-06-16T00:00:00Z\"\n\ttime_ends = \"2024-06-22T00:00:00Z\"\n\tfreeform_tags = {\n\t\t\"Department\"=\"Finance\"\n\t}\n\tresource_filters {\n\t\tattribute = \"DEFINED_TAGS\"\n\t\tvalue {\n\t\t\tnamespace=\"ResourceSchedulerCanary\"\n\t\t\ttag_key=\"ScheduleTagFilterTestKey\"\n\t\t\tvalue=\"foo\"\n\t\t} \n\n\t\tvalue {\n\t\t\tnamespace=\"Test\"\n\t\t\ttag_key=\"ORM-123837\"\n\t\t\tvalue=\"inaccessible\"\n\t\t}\n\t}\n\tresource_filters {\n\t\tattribute = \"RESOURCE_TYPE\"\n\t\tvalue {\n\t\t\tvalue=\"instance\"\n\t\t} \n\t\tvalue {\n\t\t\tvalue=\"autonomousDatabase\"\n\t\t}\n\t}\n\tresource_filters {\n\t\tattribute = \"COMPARTMENT_ID\"\n\t\tvalue {\n\t\t\tvalue=\"ocid1.tenancy.oc1..aaaaaaaacijm644q2oxrhacoweompuuzccbvkqo4syimyapchsl6afxtuqlq\"\n\t\t} \n\t}\n}"
)

func TestResourceSchedulerScheduleMultipleFiltersResourceFilter(t *testing.T) {
	httpreplay.SetScenario("TestResourceSchedulerScheduleMultipleFiltersResourceFilter")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_resource_scheduler_schedule.test_schedule"

	acctest.ResourceTest(t, testAccCheckResourceSchedulerScheduleDestroy, []resource.TestStep{
		// verify create with multiple resourceFilter objects with optionals
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies + multipleConfig,

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
				resource.TestCheckResourceAttr(resourceName, "resource_filters.#", "3"),

				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.attribute", "DEFINED_TAGS"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.namespace", "ResourceSchedulerCanary"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.tag_key", "ScheduleTagFilterTestKey"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.value", "foo"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.1.namespace", "Test"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.1.tag_key", "ORM-123837"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.1.value", "inaccessible"),

				resource.TestCheckResourceAttr(resourceName, "resource_filters.1.attribute", "RESOURCE_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.0.value", "instance"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.1.value", "autonomousDatabase"),

				resource.TestCheckResourceAttr(resourceName, "resource_filters.1.attribute", "COMPARTMENT_ID"),
				resource.TestCheckResourceAttr(resourceName, "resource_filters.0.value.#", "1"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}
