// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	filter = `<<EOF
{
	"operator": "AND",
	"dimentions": [],
	"tags": [],
	"filters": [
		"operator": "OR",
		"dimentions": [
			"key": "compartName"
			"value": "dxterraformtest"
		]
		"filters": []
		"tags": []
	]
}
EOF`

	usageRepresentationWithOptionals = `resource "oci_metering_computation_usage" "test_usage" {
compartment_depth = 1
filter = <<EOF
{
		"operator": "AND",
		"dimensions": [],
		"tags": [],
		"filters": [
			{
				"operator": "OR",
			 	"dimensions": [
					{
						"key": "compartmentName",
						"value": "dxterraformtest"
					}
				],
				"filters": [],
				"tags": []
			}
		]
}
EOF
granularity = "DAILY"
group_by = ["service"]
query_type = "COST"
tenant_id = "${var.tenancy_id}"
time_usage_ended = "${var.time_usage_ended}"
time_usage_started = "${var.time_usage_started}"
time_forecast_ended= "2021-03-21T00:00:00Z"
}
`

	UsageRequiredOnlyResource = MeteringComputationUsageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage", "test_usage", acctest.Required, acctest.Create, MeteringComputationUsageRepresentation)

	MeteringComputationUsageRepresentation = map[string]interface{}{
		"granularity":        acctest.Representation{RepType: acctest.Required, Create: `DAILY`},
		"tenant_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"time_usage_ended":   acctest.Representation{RepType: acctest.Required, Create: `2021-03-19T00:00:00Z`},
		"time_usage_started": acctest.Representation{RepType: acctest.Required, Create: `2021-03-18T00:00:00Z`},
		"compartment_depth":  acctest.Representation{RepType: acctest.Optional, Create: `1`},
		//"filter":               acctest.Representation{RepType: acctest.Optional, Create: },
		"filter":   acctest.Representation{RepType: acctest.Optional, Create: `{\"operator\":\"OR\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"dxterraformtest\"}],\"tags\":[],\"filters\":[]}`, Update: `{\"operator\":\"OR\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"dxterraformtest\"}],\"tags\":[],\"filters\":[]}`},
		"forecast": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MeteringComputationUsageForecastRepresentation},
		"group_by": acctest.Representation{RepType: acctest.Optional, Create: []string{`service`}},
		//"group_by_tag":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: MeteringComputationUsageGroupByTagRepresentation},
		"is_aggregate_by_time": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"query_type":           acctest.Representation{RepType: acctest.Optional, Create: `COST`},
	}
	MeteringComputationUsageForecastRepresentation = map[string]interface{}{
		"time_forecast_ended":   acctest.Representation{RepType: acctest.Required, Create: `2021-03-20T00:00:00Z`},
		"forecast_type":         acctest.Representation{RepType: acctest.Optional, Create: `BASIC`},
		"time_forecast_started": acctest.Representation{RepType: acctest.Optional, Create: `2021-03-19T00:00:00Z`},
	}
	MeteringComputationUsageGroupByTagRepresentation = map[string]interface{}{
		"key":       acctest.Representation{RepType: acctest.Optional, Create: `key`},
		"namespace": acctest.Representation{RepType: acctest.Optional, Create: `namespace`},
		"value":     acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}

	MeteringComputationUsageResourceDependencies = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)
	usgaeEndTimeStr, usageStartTimeStr := generateUsageRepresentationWithCurrentTimeInputs()
	usgaeEndTimeVariableStr := fmt.Sprintf("variable \"time_usage_ended\" { default = \"%s\" }\n", usgaeEndTimeStr)
	usageStartTimeVariableStr := fmt.Sprintf("variable \"time_usage_started\" { default = \"%s\" }\n", usageStartTimeStr)

	resourceName := "oci_metering_computation_usage.test_usage"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MeteringComputationUsageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage", "test_usage", acctest.Optional, acctest.Create, MeteringComputationUsageRepresentation), "usageapi", "usage", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			PreConfig: func() {
				fmt.Printf("config is : %s", acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage", "test_usage", acctest.Optional, acctest.Create, MeteringComputationUsageRepresentation))
			},
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usgaeEndTimeVariableStr + usageStartTimeVariableStr + MeteringComputationUsageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage", "test_usage", acctest.Required, acctest.Create, MeteringComputationUsageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "granularity", "DAILY"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_usage_ended"),
				resource.TestCheckResourceAttrSet(resourceName, "time_usage_started"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usgaeEndTimeVariableStr + usageStartTimeVariableStr + MeteringComputationUsageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usgaeEndTimeVariableStr + usageStartTimeVariableStr + MeteringComputationUsageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage", "test_usage", acctest.Optional, acctest.Create, MeteringComputationUsageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_depth", "1"),
				resource.TestCheckResourceAttr(resourceName, "forecast.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "forecast.0.forecast_type", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "forecast.0.time_forecast_ended", "2021-03-20T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "forecast.0.time_forecast_started", "2021-03-19T00:00:00Z"),
				resource.TestCheckResourceAttrSet(resourceName, "filter"),
				resource.TestCheckResourceAttr(resourceName, "granularity", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "group_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_type", "COST"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_usage_ended"),
				resource.TestCheckResourceAttrSet(resourceName, "time_usage_started"),

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

func generateUsageRepresentationWithCurrentTimeInputs() (string, string) {
	t := time.Now()
	year, month, day := t.Date()
	endTime := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	startTime := endTime.Add(-24 * time.Hour)
	usgaeEndTimeStr := endTime.Format("2006-01-02T15:04:05Z")
	usageStartTimeStr := startTime.Format("2006-01-02T15:04:05Z")
	return usgaeEndTimeStr, usageStartTimeStr
}
