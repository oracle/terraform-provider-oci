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
	OspGatewayOspGatewayAddressRuleSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"country_code":    acctest.Representation{RepType: acctest.Required, Create: `cl`},
		"osp_home_region": acctest.Representation{RepType: acctest.Required, Create: `${var.home_region}`},
	}

	OspGatewayAddressRuleResourceConfig = ""
)

// issue-routing-tag: osp_gateway/default
func TestOspGatewayAddressRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOspGatewayAddressRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	homeRegion := utils.GetEnvSettingWithBlankDefault("region")
	regionVariableStr := fmt.Sprintf("variable \"home_region\" { default = \"%s\" }\n", homeRegion)

	singularDatasourceName := "data.oci_osp_gateway_address_rule.test_address_rule"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_address_rule", "test_address_rule", acctest.Required, acctest.Create, OspGatewayOspGatewayAddressRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + regionVariableStr + OspGatewayAddressRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "country_code"),
				resource.TestCheckResourceAttr(singularDatasourceName, "osp_home_region", homeRegion),

				resource.TestCheckResourceAttr(singularDatasourceName, "address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contact.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "country_code"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tax.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tax.0.value_set.#", "181"),
			),
		},
	})
}
