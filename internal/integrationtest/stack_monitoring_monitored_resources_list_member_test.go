// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	StackMonitoringStackMonitoringMonitoredResourcesListMemberDataSourceRepresentation = map[string]interface{}{
		"monitored_resource_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.source_resource_id}`},
		"destination_resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.destination_resource_id}`},
	}
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourcesListMemberResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourcesListMemberResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sourceResourceId := utils.GetEnvSettingWithBlankDefault("stack_mon_source_resource_id")
	if sourceResourceId == "" {
		t.Skip("Setting environmental variable source_resource_id which is a pre-requisite for this test")
	}
	sourceResourceIdVariableStr := fmt.Sprintf("variable \"source_resource_id\" { default = \"%s\" }\n", sourceResourceId)

	destinationResourceId := utils.GetEnvSettingWithBlankDefault("stack_mon_destination_resource_id")
	if destinationResourceId == "" {
		t.Skip("Setting environmental variable destination_resource_id which is a pre-requisite for this test")
	}
	destinationResourceIdVariableStr := fmt.Sprintf("variable \"destination_resource_id\" { default = \"%s\" }\n", destinationResourceId)

	destinationResourceName := utils.GetEnvSettingWithBlankDefault("stack_mon_destination_resource_name")
	if destinationResourceName == "" {
		t.Skip("Setting environmental variable destination_resource_name which is a pre-requisite for this test")
	}
	destinationResourceNameVariableStr := fmt.Sprintf("variable \"destination_resource_name\" { default = \"%s\" }\n", destinationResourceName)

	destinationResourceType := utils.GetEnvSettingWithBlankDefault("stack_mon_destination_resource_type")
	if destinationResourceType == "" {
		t.Skip("Setting environmental variable destination_resource_type which is a pre-requisite for this test")
	}
	destinationResourceTypeVariableStr := fmt.Sprintf("variable \"destination_resource_type\" { default = \"%s\" }\n", destinationResourceType)

	resourceName := "oci_stack_monitoring_monitored_resources_list_member.test_monitored_resources_list_member"

	var resId string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + sourceResourceIdVariableStr + destinationResourceIdVariableStr + destinationResourceNameVariableStr + destinationResourceTypeVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_list_member", "test_monitored_resources_list_member", acctest.Optional, acctest.Create, StackMonitoringStackMonitoringMonitoredResourcesListMemberDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "destination_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "items.0.resource_name", destinationResourceName),
				resource.TestCheckResourceAttr(resourceName, "items.0.resource_type", destinationResourceType),
				resource.TestCheckResourceAttr(resourceName, "items.0.compartment_id", compartmentId),
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
