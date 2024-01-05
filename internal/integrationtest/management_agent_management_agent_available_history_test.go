// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ManagementAgentAvailableHistoryResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentAvailableHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentAvailableHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentIds, err := getManagementAgentIds(compartmentId)
	if err != nil {
		t.Errorf("Failed to get agents in compartment %s", err)
	}
	if len(managementAgentIds) == 0 {
		t.Errorf("Failed to find any active agents in compartment %s", compartmentId)
	}
	managementAgentId := managementAgentIds[0]

	managementAgentAvailableHistoryDataSourceRepresentation := map[string]interface{}{
		"management_agent_id":                         acctest.Representation{RepType: acctest.Required, Create: managementAgentId},
		"time_availability_status_ended_greater_than": acctest.Representation{RepType: acctest.Optional, Create: `2020-01-01T01:01:01.000Z`},
		"time_availability_status_started_less_than":  acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T01:01:01.000Z`},
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
