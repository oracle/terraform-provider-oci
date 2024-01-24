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
	MarketplaceMarketplacePublisherDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"publisher_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_marketplace_publisher.test_publisher.id}`},
	}

	MarketplacePublisherResourceConfig = ``
)

// issue-routing-tag: marketplace/default
func TestMarketplacePublisherResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplacePublisherResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_publishers.test_publishers"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_publishers", "test_publishers", acctest.Required, acctest.Create, MarketplaceMarketplacePublisherDataSourceRepresentation) +
				compartmentIdVariableStr + MarketplacePublisherResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "publishers.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "publishers.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "publishers.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "publishers.0.name"),
			),
		},
	})
}
