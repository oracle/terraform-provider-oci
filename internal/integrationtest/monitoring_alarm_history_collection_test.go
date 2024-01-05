// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	MonitoringMonitoringAlarmHistoryCollectionSingularDataSourceRepresentation = map[string]interface{}{
		"alarm_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_monitoring_alarm.test_alarm.id}`},
		"alarm_historytype":                  acctest.Representation{RepType: acctest.Optional, Create: `STATE_TRANSITION_HISTORY`},
		"timestamp_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-12-01T01:00:00.001Z`},
		"timestamp_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `${timestamp()}`},
	}

	MonitoringAlarmHistoryCollectionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Required, acctest.Create, MonitoringAlarmRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: monitoring/default
func TestMonitoringAlarmHistoryCollectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringAlarmHistoryCollectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_monitoring_alarm_history_collection.test_alarm_history_collection"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_alarm_history_collection", "test_alarm_history_collection", acctest.Optional, acctest.Create, MonitoringMonitoringAlarmHistoryCollectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoringAlarmHistoryCollectionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
