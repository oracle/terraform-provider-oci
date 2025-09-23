// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsTaskScheduleFleetId           = utils.GetEnvSettingWithBlankDefault("fleet_ocid")
	JmsTaskScheduleManagedInstanceId = utils.GetEnvSettingWithBlankDefault("managed_instance_ocid")

	JmsTaskScheduleDataSourceRepresentation = map[string]interface{}{
		"fleet_id":                    acctest.Representation{RepType: acctest.Optional, Create: JmsTaskScheduleFleetId},
		"id":                          acctest.Representation{RepType: acctest.Optional, Create: `dummy_id`},
		"managed_instance_id":         acctest.Representation{RepType: acctest.Optional, Create: JmsTaskScheduleManagedInstanceId},
		"name":                        acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"task_schedule_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `taskScheduleNameContains`},
		"filter":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: JmsTaskScheduleDataSourceFilterRepresentation}}
	JmsTaskScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`dummy_id`}},
	}
)

// issue-routing-tag: jms/default
func TestJmsTaskScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsTaskScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_task_schedules.test_task_schedules"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_task_schedules",
					"test_task_schedules",
					acctest.Optional,
					acctest.Create,
					JmsTaskScheduleDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "name"),
				resource.TestCheckResourceAttrSet(datasourceName, "task_schedule_name_contains"),
			),

			// cannot verify create because it requires setup of managed instances with applications
		},
	})
}
