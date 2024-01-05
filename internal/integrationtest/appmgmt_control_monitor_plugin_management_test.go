// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	AppmgmtControlPluginResourceConfig    = MonitorPluginManagementResourceDependencies
	monitorPluginManagementRepresentation = map[string]interface{}{
		"monitored_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	appmgmtControlPluginInstanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_management_disabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_monitoring_disabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
	}

	privateVnicDetailsAppmgmtPluginRepresentation = map[string]interface{}{
		"assign_public_ip": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"subnet_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.appmgmt_control_subnet_id}`},
	}

	appmgmtControlPluginSourceDetailsRepresentation = map[string]interface{}{
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `image`},
		"source_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.appmgmt_control_image_id}`},
	}

	MonitorPluginManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: privateVnicDetailsAppmgmtPluginRepresentation},
		"source_details":      acctest.RepresentationGroup{RepType: acctest.Required, Group: appmgmtControlPluginSourceDetailsRepresentation},
		"agent_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: appmgmtControlPluginInstanceAgentConfigRepresentation},
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.appmgmt_control_image_id}`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.appmgmt_control_subnet_id}`}, //variable dependency taken from terraform.tfvars.json
	})) + AvailabilityDomainConfig
)

// issue-routing-tag: appmgmt_control/default
func TestAppmgmtControlMonitorPluginManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAppmgmtControlMonitorPluginManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("appmgmt_control_subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"appmgmt_control_subnet_id\" { default = \"%s\" }\n", subnetId)

	imageId := utils.GetEnvSettingWithBlankDefault("appmgmt_control_image_id")
	imageIdVariableStr := fmt.Sprintf("variable \"appmgmt_control_image_id\" { default = \"%s\" }\n", imageId)

	resourceName := "oci_appmgmt_control_monitor_plugin_management.test_monitor_plugin_management"

	//var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+imageIdVariableStr+MonitorPluginManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_appmgmt_control_monitor_plugin_management", "test_monitor_plugin_management", acctest.Required, acctest.Create, monitorPluginManagementRepresentation), "appmgmtcontrol", "monitorPluginManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + imageIdVariableStr + MonitorPluginManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_appmgmt_control_monitor_plugin_management", "test_monitor_plugin_management", acctest.Required, acctest.Create, monitorPluginManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "monitored_instance_id"),
			),
		},
	})
}
