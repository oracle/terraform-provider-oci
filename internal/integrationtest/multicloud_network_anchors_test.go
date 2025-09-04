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
	MulticloudNetworkAnchorsDataSourceRepresentation = map[string]interface{}{
		"subscription_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"subscription_service_name": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_service_name}`},
		"external_location":         acctest.Representation{RepType: acctest.Required, Create: `${var.network_anchor_external_location}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.network_anchor_compartment_id}`},
	}

	MulticloudNetworkAnchorsResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudNetworkAnchorsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudNetworkAnchorsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subscriptionId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_id")
	subscriptionServiceName := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_service_name")
	externalLocation := utils.GetEnvSettingWithBlankDefault("TF_VAR_network_anchor_external_location")
	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_network_anchor_compartment_id")

	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" {}\n")
	subscriptionServiceNameVariableStr := fmt.Sprintf("variable \"subscription_service_name\" {}\n")
	externalLocationVariableStr := fmt.Sprintf("variable \"network_anchor_external_location\" {}\n")
	compartmentIdVariableStr := fmt.Sprintf("variable \"network_anchor_compartment_id\" {}\n")

	datasourceName := "data.oci_multicloud_network_anchors.test_network_anchors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + subscriptionIdVariableStr + subscriptionServiceNameVariableStr + externalLocationVariableStr + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_network_anchors", "test_network_anchors", acctest.Optional, acctest.Create, MulticloudNetworkAnchorsDataSourceRepresentation) +
				MulticloudNetworkAnchorsResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_service_name", subscriptionServiceName),
				resource.TestCheckResourceAttr(datasourceName, "external_location", externalLocation),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "network_anchor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.#"),

				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.resource_anchor_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.time_updated"),
				// Commented as lifecycleState was renamed to networkAnchorLifecycleState
				// resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.network_anchor_lifecycle_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.freeform_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.defined_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_collection.0.items.0.system_tags.%"),
			),
		},
	})
}
