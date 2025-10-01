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
	MulticloudOmHubMultiCloudsMetadataDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.om_hub_compartment_id}`},
	}

	MulticloudOmHubMultiCloudsMetadataResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudOmHubMultiCloudsMetadataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudOmHubMultiCloudsMetadataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_om_hub_compartment_id")

	compartmentIdVariableStr := fmt.Sprintf("variable \"om_hub_compartment_id\" {}\n")

	datasourceName := "data.oci_multicloud_om_hub_multi_clouds_metadata.test_om_hub_multi_clouds_metadata"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_om_hub_multi_clouds_metadata", "test_om_hub_multi_clouds_metadata", acctest.Required, acctest.Create, MulticloudOmHubMultiCloudsMetadataDataSourceRepresentation) +
				compartmentIdVariableStr + MulticloudOmHubMultiCloudsMetadataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "multi_cloud_metadata_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "multi_cloud_metadata_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "multi_cloud_metadata_collection.0.items.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "multi_cloud_metadata_collection.0.items.0.subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "multi_cloud_metadata_collection.0.items.0.time_created"),
			),
		},
	})
}
