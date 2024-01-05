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
	StackMonitoringStackMonitoringDiscoveryJobLogDataSourceRepresentation = map[string]interface{}{
		"discovery_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_discovery_job.test_discovery_job.id}`},
		"log_type":         acctest.Representation{RepType: acctest.Optional, Create: `INFO`},
	}

	StackMonitoringDiscoveryJobLogResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_discovery_job", "test_discovery_job", acctest.Optional, acctest.Create, StackMonitoringDiscoveryJobRepresentation)
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringDiscoveryJobLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringDiscoveryJobLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_discovery")
	if managementAgentId == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_discovery that represents management agent capable of running stack monitoring discovery is pre-requisite for this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_discovery\" { default = \"%s\" }\n", managementAgentId)

	datasourceName := "data.oci_stack_monitoring_discovery_job_logs.test_discovery_job_logs"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_discovery_job_logs", "test_discovery_job_logs", acctest.Optional, acctest.Update, StackMonitoringStackMonitoringDiscoveryJobLogDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + StackMonitoringDiscoveryJobLogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_job_id"),
				resource.TestCheckResourceAttr(datasourceName, "discovery_job_log_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_job_log_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_job_log_collection.0.items.0.log_message"),
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_job_log_collection.0.items.0.log_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_job_log_collection.0.items.0.time_created"),
			),
		},
	})
}
