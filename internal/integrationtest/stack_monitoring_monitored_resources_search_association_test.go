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
	StackMonitoringMonitoredResourcesSearchAssociationRequiredOnlyResource = StackMonitoringMonitoredResourcesSearchAssociationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_search_association", "test_monitored_resources_search_association", acctest.Required, acctest.Create, StackMonitoringMonitoredResourcesSearchAssociationRepresentation)

	StackMonitoringMonitoredResourcesSearchAssociationRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"association_type":          acctest.Representation{RepType: acctest.Optional, Create: `uses`},
		"destination_resource_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_stack_monitoring_monitored_resource.test_destination_resource.id}`},
		"destination_resource_name": acctest.Representation{RepType: acctest.Optional, Create: `terraformResource`},
		"destination_resource_type": acctest.Representation{RepType: acctest.Optional, Create: `host`},
		"source_resource_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_stack_monitoring_monitored_resource.test_source_resource.id}`},
		"source_resource_name":      acctest.Representation{RepType: acctest.Optional, Create: `terraformSecondaryResource`},
		"source_resource_type":      acctest.Representation{RepType: acctest.Optional, Create: `host`},
	}

	StackMonitoringMonitoredResourcesSearchAssociationResourceDependencies = StackMonitoredResourcesAssociateMonitoredResourceConfig
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourcesSearchAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourcesSearchAssociationResource_basic")
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

	resourceName := "oci_stack_monitoring_monitored_resources_search_association.test_monitored_resources_search_association"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringMonitoredResourcesSearchAssociationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_search_association", "test_monitored_resources_search_association", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourcesSearchAssociationRepresentation), "stackmonitoring", "monitoredResourcesSearchAssociation", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + managementAgentId2VariableStr + hostname2VariableStr + StackMonitoringMonitoredResourcesSearchAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_search_association", "test_monitored_resources_search_association", acctest.Required, acctest.Create, StackMonitoringMonitoredResourcesSearchAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + managementAgentId2VariableStr + hostname2VariableStr + StackMonitoringMonitoredResourcesSearchAssociationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + managementAgentId2VariableStr + hostname2VariableStr + StackMonitoringMonitoredResourcesSearchAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_search_association", "test_monitored_resources_search_association", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourcesSearchAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "association_type", "uses"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "destination_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "destination_resource_name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "destination_resource_type", "host"),
				resource.TestCheckResourceAttrSet(resourceName, "source_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "source_resource_name", "terraformSecondaryResource"),
				resource.TestCheckResourceAttr(resourceName, "source_resource_type", "host"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "items.0.source_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "items.0.source_resource_details.0.name", "terraformSecondaryResource"),
				resource.TestCheckResourceAttr(resourceName, "items.0.source_resource_details.0.type", "host"),
				resource.TestCheckResourceAttr(resourceName, "items.0.source_resource_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "items.0.destination_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "items.0.destination_resource_details.0.name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "items.0.destination_resource_details.0.type", "host"),
				resource.TestCheckResourceAttr(resourceName, "items.0.destination_resource_details.0.compartment_id", compartmentId),

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
