// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	alarmStatusDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Optional, create: `false`},
		"display_name":              Representation{repType: Optional, create: `${oci_monitoring_alarm.test_alarm.display_name}`},
	}

	AlarmStatusResourceConfig = DefinedTagsDependencies + AvailabilityDomainConfig +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, getTopicRepresentationCopyWithRandomNameOrHttpReplayValue(10, charsetWithoutDigits, "talarmstatus")) +
		generateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Required, Create, alarmRepresentation)
)

func TestMonitoringAlarmStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringAlarmStatusResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_monitoring_alarm_statuses.test_alarm_statuses"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_monitoring_alarm_statuses", "test_alarm_statuses", Optional, Create, alarmStatusDataSourceRepresentation) +
					compartmentIdVariableStr + AlarmStatusResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "display_name"),

					resource.TestCheckResourceAttrSet(datasourceName, "alarm_statuses.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "alarm_statuses.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "alarm_statuses.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "alarm_statuses.0.severity"),
					resource.TestCheckResourceAttrSet(datasourceName, "alarm_statuses.0.status"),
					resource.TestCheckResourceAttrSet(datasourceName, "alarm_statuses.0.timestamp_triggered"),
				),
			},
		},
	})
}
