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
	MulticloudNetworkAnchorDataSourceRepresentation = map[string]interface{}{
		"network_anchor_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.network_anchor_id}`},
		"subscription_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"subscription_service_name": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_service_name}`},
	}

	MulticloudNetworkAnchorResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudNetworkAnchorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudNetworkAnchorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	networkAnchorId := utils.GetEnvSettingWithBlankDefault("TF_VAR_network_anchor_id")
	subscriptionId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_id")
	subscriptionServiceName := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_service_name")

	networkAnchorIdVariableStr := fmt.Sprintf("variable \"network_anchor_id\" {}\n")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" {}\n")
	subscriptionServiceNameVariableStr := fmt.Sprintf("variable \"subscription_service_name\" {}\n")

	datasourceName := "data.oci_multicloud_network_anchor.test_network_anchor"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + networkAnchorIdVariableStr + subscriptionIdVariableStr + subscriptionServiceNameVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_network_anchor", "test_network_anchor", acctest.Required, acctest.Create, MulticloudNetworkAnchorDataSourceRepresentation) +
				MulticloudNetworkAnchorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "network_anchor_id", networkAnchorId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_service_name", subscriptionServiceName),

				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_anchor_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_updated"),
				// Commented as lifecycleState was renamed to networkAnchorLifecycleState
				// resource.TestCheckResourceAttrSet(datasourceName, "network_anchor_lifecycle_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "defined_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "system_tags.%"),

				resource.TestCheckResourceAttr(datasourceName, "oci_metadata_item.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_metadata_item.0.network_anchor_connection_status"),

				resource.TestCheckResourceAttr(datasourceName, "oci_metadata_item.0.vcn.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_metadata_item.0.vcn.0.vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_metadata_item.0.vcn.0.cidr_blocks.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_metadata_item.0.vcn.0.backup_cidr_blocks.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_metadata_item.0.vcn.0.dns_label"),

				resource.TestCheckResourceAttr(datasourceName, "oci_metadata_item.0.dns.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "oci_metadata_item.0.subnets.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_metadata_item.0.subnets.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_metadata_item.0.subnets.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_metadata_item.0.subnets.0.label"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_service_provider_metadata_item.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_service_provider_metadata_item.0.region"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_service_provider_metadata_item.0.odb_network_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_service_provider_metadata_item.0.cidr_blocks.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_service_provider_metadata_item.0.network_anchor_uri"),
			),
		},
	})
}
