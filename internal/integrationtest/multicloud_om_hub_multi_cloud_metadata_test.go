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
	MulticloudOmHubMultiCloudMetadataDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.om_hub_compartment_id}`},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
	}

	MulticloudOmHubMultiCloudMetadataResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudOmHubMultiCloudMetadataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudOmHubMultiCloudMetadataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_om_hub_compartment_id")
	subscriptionId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_id")

	compartmentIdVariableStr := fmt.Sprintf("variable \"om_hub_compartment_id\" {}\n")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" {}\n")

	datasourceName := "data.oci_multicloud_om_hub_multi_cloud_metadata.test_om_hub_multi_cloud_metadata"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_om_hub_multi_cloud_metadata", "test_om_hub_multi_cloud_metadata", acctest.Required, acctest.Create, MulticloudOmHubMultiCloudMetadataDataSourceRepresentation) +
				subscriptionIdVariableStr + compartmentIdVariableStr + MulticloudOmHubMultiCloudMetadataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttrSet(datasourceName, "base_compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "base_subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created"),
			),
		},
	})
}
