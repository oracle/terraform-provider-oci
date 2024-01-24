// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringMonitoredResourcesSearchRequiredOnlyResource = StackMonitoringMonitoredResourcesSearchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_search", "test_monitored_resources_search", acctest.Required, acctest.Create, StackMonitoringMonitoredResourcesSearchRepresentation)

	StackMonitoringMonitoredResourcesSearchRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"host_name":           acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_hostname_resource1}`},
		"license":             acctest.Representation{RepType: acctest.Optional, Create: `STANDARD_EDITION`},
		"management_agent_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_management_agent_id_resource1}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `terraformResource`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                acctest.Representation{RepType: acctest.Optional, Create: `host`},
	}

	StackMonitoringMonitoredResourcesSearchResourceDependencies = StackMonitoringMonitoredResourceResourceConfig
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourcesSearchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourcesSearchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId1 := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_resource1")
	if managementAgentId1 == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_resource1 that represents management agent with resource monitoring plugin is pre-requisite for this test")
	}
	managementAgentId1VariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_resource1\" { default = \"%s\" }\n", managementAgentId1)

	hostname1 := utils.GetEnvSettingWithBlankDefault("stack_mon_hostname_resource1")
	if hostname1 == "" {
		t.Skip("Setting environmental variable stack_mon_hostname_resource1 that host accessible by agent defined by stack_mon_management_agent_id_resource1 variable is pre-requisite for this test")
	}
	hostname1VariableStr := fmt.Sprintf("variable \"stack_mon_hostname_resource1\" { default = \"%s\" }\n", hostname1)

	resourceName := "oci_stack_monitoring_monitored_resources_search.test_monitored_resources_search"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringMonitoredResourcesSearchResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_search", "test_monitored_resources_search", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourcesSearchRepresentation), "stackmonitoring", "monitoredResourcesSearch", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + StackMonitoringMonitoredResourcesSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_search", "test_monitored_resources_search", acctest.Required, acctest.Create, StackMonitoringMonitoredResourcesSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + StackMonitoringMonitoredResourcesSearchResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + StackMonitoringMonitoredResourcesSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_search", "test_monitored_resources_search", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourcesSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "host_name", hostname1),
				resource.TestCheckResourceAttr(resourceName, "license", "STANDARD_EDITION"),
				resource.TestCheckResourceAttrSet(resourceName, "items.#"),
				resource.TestCheckResourceAttr(resourceName, "management_agent_id", managementAgentId1),
				resource.TestCheckResourceAttr(resourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
