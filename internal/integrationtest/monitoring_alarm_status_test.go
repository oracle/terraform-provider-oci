// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	alarmStatusDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_monitoring_alarm.test_alarm.display_name}`},
	}

	AlarmStatusResourceConfig = DefinedTagsDependencies + AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, getTopicRepresentationCopyWithRandomNameOrHttpReplayValue(10, utils.CharsetWithoutDigits, "talarmstatus")) +
		acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Required, acctest.Create, alarmRepresentation)
)

// issue-routing-tag: monitoring/default
func TestMonitoringAlarmStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringAlarmStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_monitoring_alarm_statuses.test_alarm_statuses"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_alarm_statuses", "test_alarm_statuses", acctest.Optional, acctest.Create, alarmStatusDataSourceRepresentation) +
				compartmentIdVariableStr + AlarmStatusResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}
