// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	CloudBridgeSupportedCloudRegionDataSourceRepresentation = map[string]interface{}{
		"asset_source_type": acctest.Representation{RepType: acctest.Optional, Create: `AWS`},
		"name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `us-east-1`},
	}

	CloudBridgeSupportedCloudRegionResourceConfig = ""
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeSupportedCloudRegionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeSupportedCloudRegionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_cloud_bridge_supported_cloud_regions.test_supported_cloud_regions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_supported_cloud_regions", "test_supported_cloud_regions", acctest.Optional, acctest.Create, CloudBridgeSupportedCloudRegionDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeSupportedCloudRegionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "supported_cloud_region_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "supported_cloud_region_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "supported_cloud_region_collection.0.items.0.asset_source_type", "AWS"),
				resource.TestCheckResourceAttr(datasourceName, "supported_cloud_region_collection.0.items.0.name", "us-east-1"),
			),
		},
	})
}
