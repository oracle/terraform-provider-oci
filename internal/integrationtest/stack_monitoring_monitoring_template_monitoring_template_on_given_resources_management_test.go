// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementRepresentation = map[string]interface{}{
		"monitoring_template_id":                        acctest.Representation{RepType: acctest.Required, Create: monitoringTemplateId},
		"enable_monitoring_template_on_given_resources": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	MonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management.test_monitoring_template_monitoring_template_on_given_resources_management"
	parentResourceName := "oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management.test_monitoring_template_monitoring_template_on_given_resources_management"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management", "test_monitoring_template_monitoring_template_on_given_resources_management", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementRepresentation), "stackmonitoring", "monitoringTemplateMonitoringTemplateOnGivenResourcesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create with enable
		{
			Config: config + compartmentIdVariableStr + MonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management", "test_monitoring_template_monitoring_template_on_given_resources_management", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "monitoring_template_id"),
			),
		},
		// Verify enable
		{
			Config: config + compartmentIdVariableStr + MonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management", "test_monitoring_template_monitoring_template_on_given_resources_management", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_monitoring_template_on_given_resources", "true"),
			),
		},
		// Delete before next Create
		{
			Config: config + compartmentIdVariableStr + MonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceDependencies,
		},
		// Create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + MonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management", "test_monitoring_template_monitoring_template_on_given_resources_management", acctest.Optional, acctest.Create, StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "monitoring_template_id"),
			),
		},
		// Update to disable
		{
			Config: config + compartmentIdVariableStr + MonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management", "test_monitoring_template_monitoring_template_on_given_resources_management", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "monitoring_template_id"),
			),
		},
		// Verify disable
		{
			Config: config + compartmentIdVariableStr + MonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management", "test_monitoring_template_monitoring_template_on_given_resources_management", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_monitoring_template_on_given_resources", "false"),
			),
		},
	})
}
