// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
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
}
`

	UsageRequiredOnlyResource = UsageResourceDependencies +
		generateResourceFromRepresentationMap("oci_metering_computation_usage", "test_usage", Required, Create, usageRepresentation)

	usageRepresentation = map[string]interface{}{
		"granularity":        Representation{repType: Required, create: `DAILY`},
		"tenant_id":          Representation{repType: Required, create: `${var.tenancy_id}`},
		"time_usage_ended":   Representation{repType: Required, create: `${var.time_usage_ended}`},
		"time_usage_started": Representation{repType: Required, create: `${var.time_usage_started}`},
		"compartment_depth":  Representation{repType: Optional, create: `1`},
		"filter":             Representation{repType: Optional, create: `filter`},
		"group_by":           Representation{repType: Optional, create: []string{`service`}},
		"query_type":         Representation{repType: Optional, create: `COST`},
	}

	UsageResourceDependencies = ""
)

func TestMeteringComputationUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationUsageResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)
	usgaeEndTimeStr, usageStartTimeStr := generateUsageRepresentationWithCurrentTimeInputs()
	usgaeEndTimeVariableStr := fmt.Sprintf("variable \"time_usage_ended\" { default = \"%s\" }\n", usgaeEndTimeStr)
	usageStartTimeVariableStr := fmt.Sprintf("variable \"time_usage_started\" { default = \"%s\" }\n", usageStartTimeStr)

	resourceName := "oci_metering_computation_usage.test_usage"

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				PreConfig: func() {
					fmt.Printf("config is : %s", generateResourceFromRepresentationMap("oci_metering_computation_usage", "test_usage", Optional, Create, usageRepresentation))
				},
				Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usgaeEndTimeVariableStr + usageStartTimeVariableStr + UsageResourceDependencies +
					generateResourceFromRepresentationMap("oci_metering_computation_usage", "test_usage", Required, Create, usageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "granularity", "DAILY"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_usage_ended"),
					resource.TestCheckResourceAttrSet(resourceName, "time_usage_started"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usgaeEndTimeVariableStr + usageStartTimeVariableStr + UsageResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + tenancyIdVariableStr + usgaeEndTimeVariableStr + usageStartTimeVariableStr + UsageResourceDependencies +
					usageRepresentationWithOptionals,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_depth", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "filter"),
					resource.TestCheckResourceAttr(resourceName, "granularity", "DAILY"),
					resource.TestCheckResourceAttr(resourceName, "group_by.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "items.#"),
					resource.TestCheckResourceAttr(resourceName, "query_type", "COST"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_usage_ended"),
					resource.TestCheckResourceAttrSet(resourceName, "time_usage_started"),

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
