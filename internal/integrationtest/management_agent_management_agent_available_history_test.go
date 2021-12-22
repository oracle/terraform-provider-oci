// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	managementAgentAvailableHistoryDataSourceRepresentation = map[string]interface{}{
		"management_agent_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id}`},
		"time_availability_status_ended_greater_than": acctest.Representation{RepType: acctest.Optional, Create: `2020-01-01T01:01:01.000Z`},
		"time_availability_status_started_less_than":  acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T01:01:01.000Z`},
	}

	ManagementAgentAvailableHistoryResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentAvailableHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentAvailableHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}

	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)
	datasourceName := "data.oci_management_agent_management_agent_available_histories.test_management_agent_available_histories"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_available_histories", "test_management_agent_available_histories", acctest.Required, acctest.Create, managementAgentAvailableHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.0.availability_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.0.management_agent_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.0.time_availability_status_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.0.time_availability_status_started"),
			),
		},

		// verify datasource with optionals
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_available_histories", "test_management_agent_available_histories", acctest.Optional, acctest.Create, managementAgentAvailableHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_availability_status_ended_greater_than"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_availability_status_started_less_than"),

				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.0.availability_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.0.management_agent_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.0.time_availability_status_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_histories.0.time_availability_status_started"),
			),
		},
	})
}
