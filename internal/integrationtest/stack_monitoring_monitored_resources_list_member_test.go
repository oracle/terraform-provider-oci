// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/resourcediscovery"

	"terraform-provider-oci/internal/utils"
)

var (
	MonitoredResourcesListMemberRequiredOnlyResource = MonitoredResourcesListMemberResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_list_member", "test_monitored_resources_list_member", acctest.Required, acctest.Create, monitoredResourcesListMemberRepresentation)

	monitoredResourcesListMemberRepresentation = map[string]interface{}{
		"monitored_resource_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource.test_source_resource.id}`},
		"destination_resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_stack_monitoring_monitored_resource.test_destination_resource.id}`},
	}

	MonitoredResourcesListMemberResourceDependencies = MonitoredResourcesAssociateMonitoredResourceConfig
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourcesListMemberResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourcesListMemberResource_basic")
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

	managementAgentId2 := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_resource2")
	if managementAgentId2 == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_resource2 that represents management agent with resource monitoring plugin is pre-requisite for this test")
	}
	managementAgentId2VariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_resource2\" { default = \"%s\" }\n", managementAgentId2)

	hostname2 := utils.GetEnvSettingWithBlankDefault("stack_mon_hostname_resource2")
	if hostname2 == "" {
		t.Skip("Setting environmental variable stack_mon_hostname_resource2 that host accessible by agent defined by stack_mon_management_agent_id_resource2 variable is pre-requisite for this test")
	}
	hostname2VariableStr := fmt.Sprintf("variable \"stack_mon_hostname_resource2\" { default = \"%s\" }\n", hostname2)

	resourceName := "oci_stack_monitoring_monitored_resources_list_member.test_monitored_resources_list_member"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitoredResourcesListMemberResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_list_member", "test_monitored_resources_list_member", acctest.Optional, acctest.Create, monitoredResourcesListMemberRepresentation), "stackmonitoring", "monitoredResourcesListMember", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + managementAgentId2VariableStr + hostname2VariableStr + MonitoredResourcesListMemberResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_list_member", "test_monitored_resources_list_member", acctest.Required, acctest.Create, monitoredResourcesListMemberRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "monitored_resource_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + managementAgentId2VariableStr + hostname2VariableStr + MonitoredResourcesListMemberResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + managementAgentId2VariableStr + hostname2VariableStr + MonitoredResourcesListMemberResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_list_member", "test_monitored_resources_list_member", acctest.Optional, acctest.Create, monitoredResourcesListMemberRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "destination_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "monitored_resource_id"),

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
