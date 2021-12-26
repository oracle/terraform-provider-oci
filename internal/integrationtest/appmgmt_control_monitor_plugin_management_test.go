// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	monitorPluginManagementRepresentation = map[string]interface{}{
		"monitored_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${var.monitored_instance_id}`},
	}
	MonitorPluginManagementResourceDependencies = ""
)

// issue-routing-tag: appmgmt_control/default
func TestAppmgmtControlMonitorPluginManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAppmgmtControlMonitorPluginManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	monitoredInstanceId := utils.GetEnvSettingWithBlankDefault("monitored_instance_id")
	if monitoredInstanceId == "" {
		t.Skip("Manually create vm instance with Management Agent and set TF_VAR_monitored_instance_id variable with OCID of such VM instance to run this test")
	}
	monitoredInstanceIdVariableStr := fmt.Sprintf("variable \"monitored_instance_id\" { default = \"%s\" }\n", monitoredInstanceId)

	resourceName := "oci_appmgmt_control_monitor_plugin_management.test_monitor_plugin_management"

	//var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitorPluginManagementResourceDependencies+monitoredInstanceIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_appmgmt_control_monitor_plugin_management", "test_monitor_plugin_management", acctest.Required, acctest.Create, monitorPluginManagementRepresentation), "appmgmtcontrol", "monitorPluginManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MonitorPluginManagementResourceDependencies + monitoredInstanceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_appmgmt_control_monitor_plugin_management", "test_monitor_plugin_management", acctest.Required, acctest.Create, monitorPluginManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "monitored_instance_id"),
			),
		},
	})
}
