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
	OpsiOpsiConfigurationConfigurationItemSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_item_field":               acctest.Representation{RepType: acctest.Required, Create: []string{`metadata`, `name`, `value`, `defaultValue`, `valueSourceConfig`}},
		"config_items_applicable_context": acctest.Representation{RepType: acctest.Required, Create: []string{`DB_CAPACITY_PLANNING`}},
		"name":                            acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"opsi_config_type":                acctest.Representation{RepType: acctest.Required, Create: `UX_CONFIGURATION`},
	}

	OpsiOpsiConfigurationConfigurationItemResourceConfig = ""
)

// issue-routing-tag: opsi/default
func TestOpsiOpsiConfigurationConfigurationItemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOpsiConfigurationConfigurationItemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_opsi_opsi_configuration_configuration_item.test_opsi_configuration_configuration_item"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_opsi_configuration_configuration_item", "test_opsi_configuration_configuration_item", acctest.Required, acctest.Create, OpsiOpsiConfigurationConfigurationItemSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiOpsiConfigurationConfigurationItemResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_item_field.#", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_items_applicable_context.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_items.0.name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_items.0.value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_items.0.default_value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_items.0.value_source_config"),
				resource.TestCheckResourceAttr(singularDatasourceName, "opsi_config_type", "UX_CONFIGURATION"),
			),
		},
	})
}
