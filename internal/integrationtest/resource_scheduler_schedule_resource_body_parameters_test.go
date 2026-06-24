// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	bodyParameterConfig = "resource \"oci_resource_scheduler_schedule\" \"test_schedule\" {\n\taction = \"START_RESOURCE\"\n\tcompartment_id = \"${var.compartment_id}\"\n\trecurrence_details = \"FREQ=DAILY;INTERVAL=1\"\n\trecurrence_type = \"ICAL\"\n\tdescription = \"provider description1\"\n\tdisplay_name = \"provider displayName1\"\n\ttime_starts = \"2085-05-01T00:00:00Z\"\n\ttime_ends = \"2085-12-31T00:00:00Z\"\n\tresources {\n\t\tid = \"${var.function_ocid}\"\n\t\tparameters {\n\t\t\tparameter_type = \"BODY\"\n\t\t\tvalue = [\"{\\\"dummyKey\\\":\\\"dummyValue\\\"}\"]\n\t\t}\n\t}\n}"
)

func TestResourceSchedulerScheduleBodyParameterResource(t *testing.T) {
	httpreplay.SetScenario("TestResourceSchedulerScheduleBodyParameterResource")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	functionOcid := utils.GetEnvSettingWithBlankDefault("function_ocid")
	if functionOcid == "" {
		t.Skip("Skipping body parameter test because TF_VAR_function_ocid is not set")
	}
	functionOcidVariableStr := fmt.Sprintf("variable \"function_ocid\" { default = \"%s\" }\n", functionOcid)

	resourceName := "oci_resource_scheduler_schedule.test_schedule"

	acctest.ResourceTest(t, testAccCheckResourceSchedulerScheduleDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + functionOcidVariableStr + bodyParameterConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),
				resource.TestCheckResourceAttr(resourceName, "description", "provider description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provider displayName1"),
				resource.TestCheckResourceAttr(resourceName, "time_starts", "2085-05-01T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "time_ends", "2085-12-31T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.id", functionOcid),
				resource.TestCheckResourceAttr(resourceName, "resources.0.parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.parameters.0.parameter_type", "BODY"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.parameters.0.value.#", "1"),
				resource.TestCheckTypeSetElemAttr(resourceName, "resources.0.parameters.0.value.*", `{"dummyKey":"dummyValue"}`),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}
