// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	publisherDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"publisher_id":   Representation{repType: Optional, create: `${oci_marketplace_publisher.test_publisher.id}`},
	}

	PublisherResourceConfig = ``
)

// issue-routing-tag: marketplace/default
func TestMarketplacePublisherResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplacePublisherResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_publishers.test_publishers"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_marketplace_publishers", "test_publishers", Required, Create, publisherDataSourceRepresentation) +
				compartmentIdVariableStr + PublisherResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "publishers.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "publishers.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "publishers.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "publishers.0.name"),
			),
		},
	})
}
