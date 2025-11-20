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
	MulticloudResourceAnchorsDataSourceRepresentation = map[string]interface{}{
		"subscription_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.subscription_id}`},
		"subscription_service_name": acctest.Representation{RepType: acctest.Optional, Create: `${var.subscription_service_name}`},
	}

	MulticloudResourceAnchorsResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudResourceAnchorsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudResourceAnchorsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subscriptionId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_id")
	subscriptionServiceName := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_service_name")

	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" {}\n")
	subscriptionServiceNameVariableStr := fmt.Sprintf("variable \"subscription_service_name\" {}\n")

	datasourceName := "data.oci_multicloud_resource_anchors.test_resource_anchors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + subscriptionIdVariableStr + subscriptionServiceNameVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_resource_anchors", "test_resource_anchors", acctest.Optional, acctest.Create, MulticloudResourceAnchorsDataSourceRepresentation) +
				MulticloudResourceAnchorsResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_service_name", subscriptionServiceName),

				resource.TestCheckResourceAttr(datasourceName, "resource_anchor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.partner_cloud_account_identifier"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.csp_resource_anchor_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.lifecycle_details"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.lifecycle_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.system_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_collection.0.items.0.linked_compartment_id"),
			),
		},
	})
}
