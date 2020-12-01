// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	managementAgentAvailableHistoryDataSourceRepresentation = map[string]interface{}{
		"management_agent_id":                         Representation{repType: Required, create: `${var.managed_agent_id}`},
		"time_availability_status_ended_greater_than": Representation{repType: Optional, create: `2020-01-01T01:01:01.000Z`},
		"time_availability_status_started_less_than":  Representation{repType: Optional, create: `2030-01-01T01:01:01.000Z`},
	}

	ManagementAgentAvailableHistoryResourceConfig = ""
)

func TestManagementAgentManagementAgentAvailableHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentAvailableHistoryResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := getEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}

	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)
	datasourceName := "data.oci_management_agent_management_agent_available_histories.test_management_agent_available_histories"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_available_histories", "test_management_agent_available_histories", Required, Create, managementAgentAvailableHistoryDataSourceRepresentation) +
					compartmentIdVariableStr + managementAgentIdVariableStr,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_available_histories", "test_management_agent_available_histories", Optional, Create, managementAgentAvailableHistoryDataSourceRepresentation) +
					compartmentIdVariableStr + managementAgentIdVariableStr,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
