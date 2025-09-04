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
	MulticloudResourceAnchorDataSourceRepresentation = map[string]interface{}{
		"resource_anchor_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.resource_anchor_id}`},
		"subscription_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"subscription_service_name": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_service_name}`},
	}

	MulticloudResourceAnchorResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudResourceAnchorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudResourceAnchorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	resourceAnchorId := utils.GetEnvSettingWithBlankDefault("TF_VAR_resource_anchor_id")
	subscriptionId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_id")
	subscriptionServiceName := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_service_name")

	resourceAnchorIdVariableStr := fmt.Sprintf("variable \"resource_anchor_id\" {}\n")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" {}\n")
	subscriptionServiceNameVariableStr := fmt.Sprintf("variable \"subscription_service_name\" {}\n")

	datasourceName := "data.oci_multicloud_resource_anchor.test_resource_anchor"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + resourceAnchorIdVariableStr + subscriptionIdVariableStr + subscriptionServiceNameVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_resource_anchor", "test_resource_anchor", acctest.Required, acctest.Create, MulticloudResourceAnchorDataSourceRepresentation) +
				MulticloudResourceAnchorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "resource_anchor_id", resourceAnchorId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_service_name", subscriptionServiceName),

				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "lifecycle_details"),
				resource.TestCheckResourceAttrSet(datasourceName, "linked_compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "region"),
				resource.TestCheckResourceAttrSet(datasourceName, "lifecycle_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "system_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_updated"),

				// Checks for subscription_service_name: "ORACLEDBATAZURE"
				resource.TestCheckResourceAttr(datasourceName, "cloud_service_provider_metadata_item.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_service_provider_metadata_item.0.resource_anchor_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_service_provider_metadata_item.0.resource_anchor_uri"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_service_provider_metadata_item.0.resource_group"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_service_provider_metadata_item.0.subscription"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_service_provider_metadata_item.0.subscription_type", "ORACLEDBATAZURE"),
			),
		},
	})
}
