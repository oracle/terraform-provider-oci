// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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

/*
	Dependency variables:
	    hostname = var.stack_mon_hostname_resource1
	    management_agent_id = var.stack_mon_management_agent_id_resource1
*/

var (
	StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementRepresentation = map[string]interface{}{
		"metric_extension_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_metric_extension.test_metric_extension_for_enable.id}`},
		"resource_ids":        acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitored_resource.test_monitored_resource_for_enable.id}`}},
		"enable_metric_extension_on_given_resources": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	StackMonitoringMetricExtensionForEnableRepresentation = map[string]interface{}{
		"collection_recurrences": acctest.Representation{RepType: acctest.Required, Create: `FREQ=HOURLY;INTERVAL=6`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `OS File System Utilization`},
		"metric_list":            []acctest.RepresentationGroup{{RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricList0ForEnableRepresentation}, {RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricList1ForEnableRepresentation}, {RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricList2ForEnableRepresentation}, {RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricList3ForEnableRepresentation}},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `ME_OsFileSystemUtilization`},
		"query_properties":       acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMetricExtensionQueryPropertiesForEnableRepresentation},
		"resource_type":          acctest.Representation{RepType: acctest.Required, Create: `host_linux`},
		"description":            acctest.Representation{RepType: acctest.Required, Create: `Computes File System Utilization Percentage of various mount points`},
		"publish_trigger":        acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	StackMonitoringMetricExtensionMetricList0ForEnableRepresentation = map[string]interface{}{
		"data_type":    acctest.Representation{RepType: acctest.Required, Create: `STRING`},
		"name":         acctest.Representation{RepType: acctest.Required, Create: `MountPoint`},
		"is_dimension": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_hidden":    acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	StackMonitoringMetricExtensionMetricList1ForEnableRepresentation = map[string]interface{}{
		"data_type":    acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":         acctest.Representation{RepType: acctest.Required, Create: `FileSystemSize`},
		"is_dimension": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_hidden":    acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	StackMonitoringMetricExtensionMetricList2ForEnableRepresentation = map[string]interface{}{
		"data_type":    acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":         acctest.Representation{RepType: acctest.Required, Create: `FileSystemUsed`},
		"is_dimension": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_hidden":    acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	StackMonitoringMetricExtensionMetricList3ForEnableRepresentation = map[string]interface{}{
		"data_type":          acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `FileSystemUsage`},
		"compute_expression": acctest.Representation{RepType: acctest.Required, Create: `(FileSystemUsed / FileSystemSize) * 100`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `File System Usage`},
		"is_dimension":       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_hidden":          acctest.Representation{RepType: acctest.Required, Create: `false`},
		"metric_category":    acctest.Representation{RepType: acctest.Required, Create: `UTILIZATION`},
		"unit":               acctest.Representation{RepType: acctest.Required, Create: `percent`},
	}

	StackMonitoringMetricExtensionQueryPropertiesForEnableRepresentation = map[string]interface{}{
		"collection_method": acctest.Representation{RepType: acctest.Required, Create: `OS_COMMAND`},
		"command":           acctest.Representation{RepType: acctest.Required, Create: `/bin/bash`},
		"delimiter":         acctest.Representation{RepType: acctest.Required, Create: `|`},
		"script_details":    acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMetricExtensionQueryPropertiesScriptDetailsForEnableRepresentation},
		"starts_with":       acctest.Representation{RepType: acctest.Required, Create: `oci_result=`},
	}

	StackMonitoringMetricExtensionQueryPropertiesScriptDetailsForEnableRepresentation = map[string]interface{}{
		"content": acctest.Representation{RepType: acctest.Required, Create: `IyEvYmluL2Jhc2gKIyBDb3B5cmlnaHQgKGMpIDIwMjIsIE9yYWNsZSBhbmQvb3IgaXRzIGFmZmlsaWF0ZXMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuCiMKIyBTdGFjayBNb25pdG9yaW5nIC8gSG9zdDogY29sbGVjdCBmaWxlc3lzdGVtIHN0YXRpc3RpY3MgZnJvbSBMaW51eCBob3N0cwojCiMgT3V0cHV0IGZvcm1hdDoKIwojIHJlc3VsdD1tb3VudHxzaXplfHVzZWQKCmV4ZWMgMTA+JjEKZXhlYyAxPiYyCgoKd2hpbGUgcmVhZCAtciBkZXYgc2l6ZSB1c2VkIGF2YWlsIHVzZWRwIG1vdW50IG90aGVyCmRvCiAgICBpZiBbWyAiJHtkZXZ9IiA9fiAvIF1dCiAgICB0aGVuCiAgICAgICAgaWYgWyAiJHt0b3R9IiA9PSAiMCIgXQogICAgICAgIHRoZW4KICAgICAgICAgICAgIyBQcmV2ZW50IGRldmlzaW9uIGJ5IHplcm8KICAgICAgICAgICAgdXNlZD0wCiAgICAgICAgICAgIHVzZWRwPTAKICAgICAgICBmaQoKICAgICAgICBwcmludGYgIm9jaV9yZXN1bHQ9JXN8JXN8JXNcbiIgIiR7bW91bnR9IiAiJHtzaXplfSIgIiR7dXNlZH0iID4mMTAKICAgIGZpCmRvbmUgPCA8KGRmIC1rIDI+L2Rldi9udWxsKQ==`},
		"name":    acctest.Representation{RepType: acctest.Required, Create: `fileSystem.sh`},
	}

	StackMonitoringMonitoredResourceRepresentationForMetricExtEnable = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `TerraHostForMetricExt`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `host`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `TerraHostForMetricExt`},
		"host_name":           acctest.Representation{RepType: acctest.Required, Create: `${var.stack_mon_hostname_resource1}`},
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${var.stack_mon_management_agent_id_resource1}`},
		"properties":          []acctest.RepresentationGroup{{RepType: acctest.Required, Group: StackMonitoringMonitoredResourceOSCreatePropertyForMetricExtEnable1}, {RepType: acctest.Required, Group: StackMonitoringMonitoredResourceOSCreatePropertyForMetricExtEnable2}},
		"resource_time_zone":  acctest.Representation{RepType: acctest.Required, Create: `en`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveDataRepresentationForMetricExtTest},
	}

	//Get API does not return sensitive data, it returns null
	ignoreSensitiveDataRepresentationForMetricExtEnable = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`credentials`, `properties`, `external_id`, `defined_tags`}},
	}

	StackMonitoringMonitoredResourceOSCreatePropertyForMetricExtEnable1 = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `osName`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `Linux`},
	}

	StackMonitoringMonitoredResourceOSCreatePropertyForMetricExtEnable2 = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `osVersion`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `7.0`},
	}

	MetricExtensionMetricExtensionOnGivenResourcesManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension_for_enable", acctest.Required, acctest.Create, StackMonitoringMetricExtensionForEnableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource_for_enable", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceRepresentationForMetricExtEnable)
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResource_basic")
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

	resourceName := "oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management.test_metric_extension_metric_extension_on_given_resources_management"
	parentResourceName := "oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management.test_metric_extension_metric_extension_on_given_resources_management"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementAgentId1VariableStr+hostname1VariableStr+MetricExtensionMetricExtensionOnGivenResourcesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management", "test_metric_extension_metric_extension_on_given_resources_management", acctest.Required, acctest.Create, StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementRepresentation), "stackmonitoring", "metricExtensionMetricExtensionOnGivenResourcesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create and verify enable
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MetricExtensionMetricExtensionOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management", "test_metric_extension_metric_extension_on_given_resources_management", acctest.Required, acctest.Create, StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "metric_extension_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_ids.#", "1"),
			),
		},

		// verify enable
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MetricExtensionMetricExtensionOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management", "test_metric_extension_metric_extension_on_given_resources_management", acctest.Required, acctest.Create, StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_metric_extension_on_given_resources", "true"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MetricExtensionMetricExtensionOnGivenResourcesManagementResourceDependencies,
		},

		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MetricExtensionMetricExtensionOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management", "test_metric_extension_metric_extension_on_given_resources_management", acctest.Required, acctest.Create, StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "metric_extension_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_ids.#", "1"),
			),
		},

		// update to disable
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MetricExtensionMetricExtensionOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management", "test_metric_extension_metric_extension_on_given_resources_management", acctest.Required, acctest.Update, StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "metric_extension_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_ids.#", "1"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MetricExtensionMetricExtensionOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management", "test_metric_extension_metric_extension_on_given_resources_management", acctest.Required, acctest.Update, StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_metric_extension_on_given_resources", "false"),
			),
		},
	})
}
