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
	OspGatewayOspGatewayAddressSingularDataSourceRepresentation = map[string]interface{}{
		"address_id":      acctest.Representation{RepType: acctest.Required, Create: `MX|LP|A|608700|5065_SPA`},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"osp_home_region": acctest.Representation{RepType: acctest.Required, Create: `${var.home_region}`},
	}

	OspGatewayAddressResourceConfig = ""
)

// issue-routing-tag: osp_gateway/default
func TestOspGatewayAddressResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOspGatewayAddressResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	homeRegion := utils.GetEnvSettingWithBlankDefault("region")
	regionVariableStr := fmt.Sprintf("variable \"home_region\" { default = \"%s\" }\n", homeRegion)

	singularDatasourceName := "data.oci_osp_gateway_address.test_address"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_address", "test_address", acctest.Required, acctest.Create, OspGatewayOspGatewayAddressSingularDataSourceRepresentation) +
				compartmentIdVariableStr + regionVariableStr + OspGatewayAddressResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "address_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "osp_home_region", homeRegion),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "address_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "city"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "country"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "county"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "line1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "line2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "postal_code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
