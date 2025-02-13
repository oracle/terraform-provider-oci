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
	StackMonitoringDefinedMonitoringTemplateDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	StackMonitoringDefinedMonitoringTemplateResourceConfig = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringDefinedMonitoringTemplateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringDefinedMonitoringTemplateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_stack_monitoring_defined_monitoring_templates.test_defined_monitoring_templates"

	definedMonitoringTemplatesCount := utils.GetEnvSettingWithBlankDefault("defined_monitoring_templates_count")

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_defined_monitoring_templates", "test_defined_monitoring_templates", acctest.Required, acctest.Create, StackMonitoringDefinedMonitoringTemplateDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringDefinedMonitoringTemplateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "defined_monitoring_template_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "defined_monitoring_template_collection.0.items.#", definedMonitoringTemplatesCount),
			),
		},
	})
}
