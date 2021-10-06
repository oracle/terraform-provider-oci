// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	alarmHistoryCollectionSingularDataSourceRepresentation = map[string]interface{}{
		"alarm_id":                           Representation{RepType: Required, Create: `${oci_monitoring_alarm.test_alarm.id}`},
		"alarm_historytype":                  Representation{RepType: Optional, Create: `STATE_TRANSITION_HISTORY`},
		"timestamp_greater_than_or_equal_to": Representation{RepType: Optional, Create: `2018-12-01T01:00:00.001Z`},
		"timestamp_less_than":                Representation{RepType: Optional, Create: `${timestamp()}`},
	}

	AlarmHistoryCollectionResourceConfig = GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Required, Create, alarmRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: monitoring/default
func TestMonitoringAlarmHistoryCollectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringAlarmHistoryCollectionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_monitoring_alarm_history_collection.test_alarm_history_collection"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_monitoring_alarm_history_collection", "test_alarm_history_collection", Optional, Create, alarmHistoryCollectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AlarmHistoryCollectionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "alarm_historytype", "STATE_TRANSITION_HISTORY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alarm_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timestamp_greater_than_or_equal_to", "2018-12-01T01:00:00.001Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "timestamp_less_than"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "entries.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_enabled"),
			),
			// Non empty plan expected because the data source input relies on interpolation syntax
			ExpectNonEmptyPlan: true,
		},
	})
}
