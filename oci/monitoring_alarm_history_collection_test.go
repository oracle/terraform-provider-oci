// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	alarmHistoryCollectionSingularDataSourceRepresentation = map[string]interface{}{
		"alarm_id":                           Representation{repType: Required, create: `${oci_monitoring_alarm.test_alarm.id}`},
		"alarm_historytype":                  Representation{repType: Optional, create: `STATE_TRANSITION_HISTORY`},
		"timestamp_greater_than_or_equal_to": Representation{repType: Optional, create: `2018-12-01T01:00:00.001Z`},
		"timestamp_less_than":                Representation{repType: Optional, create: `${timestamp()}`},
	}

	AlarmHistoryCollectionResourceConfig = DefinedTagsDependencies + AvailabilityDomainConfig +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, getTopicRepresentationCopyWithSuffixName("alarmhistorycollection")) +
		generateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Required, Create, alarmRepresentation)
)

func TestMonitoringAlarmHistoryCollectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringAlarmHistoryCollectionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_monitoring_alarm_history_collection.test_alarm_history_collection"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_monitoring_alarm_history_collection", "test_alarm_history_collection", Optional, Create, alarmHistoryCollectionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AlarmHistoryCollectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "alarm_historytype", "STATE_TRANSITION_HISTORY"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "alarm_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "timestamp_greater_than_or_equal_to", "2018-12-01T01:00:00.001Z"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "timestamp_less_than"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "entries.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_enabled"),
				),
			},
		},
	})
}
