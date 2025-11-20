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
	ServiceCatalogAllApplicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `Pacman Free Image Listing for Gov`},
	}

	ServiceCatalogAllApplicationResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_publishers", "test_publishers", acctest.Required, acctest.Create, MarketplaceMarketplacePublisherDataSourceRepresentation)
)

// issue-routing-tag: service_catalog/default
func TestServiceCatalogAllApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceCatalogAllApplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_service_catalog_all_applications.test_all_applications"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_catalog_all_applications", "test_all_applications", acctest.Optional, acctest.Create, ServiceCatalogAllApplicationDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogAllApplicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Pacman Free Image Listing for Gov"),

				resource.TestCheckResourceAttrSet(datasourceName, "application_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "application_collection.0.items.#", "1"),
			),
		},
	})
}
