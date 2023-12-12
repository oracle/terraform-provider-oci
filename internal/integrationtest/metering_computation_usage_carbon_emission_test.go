// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MeteringComputationUsageCarbonEmissionRequiredOnlyResource = MeteringComputationUsageCarbonEmissionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emission", "test_usage_carbon_emission", acctest.Required, acctest.Create, MeteringComputationUsageCarbonEmissionRepresentation)

	MeteringComputationUsageCarbonEmissionRepresentation = map[string]interface{}{
		"tenant_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"time_usage_ended":             acctest.Representation{RepType: acctest.Required, Create: `2023-07-01T00:00:00Z`},
		"time_usage_started":           acctest.Representation{RepType: acctest.Required, Create: `2023-01-01T00:00:00Z`},
		"compartment_depth":            acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"group_by":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`service`}},
		"is_aggregate_by_time":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"usage_carbon_emission_filter": acctest.Representation{RepType: acctest.Optional, Create: `{\"operator\":\"OR\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"dxterraformtest\"}],\"tags\":[],\"filters\":[]}`},
	}

	MeteringComputationUsageCarbonEmissionResourceDependencies = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationUsageCarbonEmissionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationUsageCarbonEmissionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)
	usageEndTimeStr, usageStartTimeStr := generateCarbonEmissionsUsageRepresentationWithCurrentTimeInputs()
	usageEndTimeVariableStr := fmt.Sprintf("variable \"time_usage_ended\" { default = \"%s\" }\n", usageEndTimeStr)
	usageStartTimeVariableStr := fmt.Sprintf("variable \"time_usage_started\" { default = \"%s\" }\n", usageStartTimeStr)

	resourceName := "oci_metering_computation_usage_carbon_emission.test_usage_carbon_emission"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MeteringComputationUsageCarbonEmissionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emission", "test_usage_carbon_emission", acctest.Optional, acctest.Create, MeteringComputationUsageCarbonEmissionRepresentation), "usageapi", "usageCarbonEmission", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usageEndTimeVariableStr + usageStartTimeVariableStr + MeteringComputationUsageCarbonEmissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emission", "test_usage_carbon_emission", acctest.Required, acctest.Create, MeteringComputationUsageCarbonEmissionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_usage_ended"),
				resource.TestCheckResourceAttrSet(resourceName, "time_usage_started"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usageEndTimeVariableStr + usageStartTimeVariableStr + MeteringComputationUsageCarbonEmissionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usageEndTimeVariableStr + usageStartTimeVariableStr + MeteringComputationUsageCarbonEmissionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emission", "test_usage_carbon_emission", acctest.Optional, acctest.Create, MeteringComputationUsageCarbonEmissionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_depth", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_usage_ended"),
				resource.TestCheckResourceAttrSet(resourceName, "time_usage_started"),
				resource.TestCheckResourceAttrSet(resourceName, "usage_carbon_emission_filter"),

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
	})
}

func generateCarbonEmissionsUsageRepresentationWithCurrentTimeInputs() (string, string) {
	t := time.Now()
	year, month, day := t.Date()
	endTime := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	startTime := endTime.Add(-24 * time.Hour)
	usgaeEndTimeStr := endTime.Format("2006-01-02T15:04:05Z")
	usageStartTimeStr := startTime.Format("2006-01-02T15:04:05Z")
	return usgaeEndTimeStr, usageStartTimeStr
}
