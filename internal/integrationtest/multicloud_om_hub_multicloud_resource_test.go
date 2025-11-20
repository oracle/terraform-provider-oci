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
	MulticloudOmHubMulticloudResourceDataSourceRepresentation = map[string]interface{}{
		"subscription_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"subscription_service_name": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_service_name}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.subscription_compartment_id}`},
		"limit":                     acctest.Representation{RepType: acctest.Optional, Create: `${var.multicloud_resources_limit}`},
	}

	MulticloudOmHubMulticloudResourceResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudOmHubMulticloudResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudOmHubMulticloudResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subscriptionId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_id")
	subscriptionServiceName := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_service_name")
	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_compartment_id")
	limit := utils.GetEnvSettingWithBlankDefault("TF_VAR_multicloud_resources_limit")

	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" {}\n")
	subscriptionServiceNameVariableStr := fmt.Sprintf("variable \"subscription_service_name\" {}\n")
	compartmentIdVariableStr := fmt.Sprintf("variable \"subscription_compartment_id\" {}\n")
	limitVariableStr := fmt.Sprintf("variable \"multicloud_resources_limit\" {}\n")

	datasourceName := "data.oci_multicloud_om_hub_multicloud_resources.test_om_hub_multicloud_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_om_hub_multicloud_resources", "test_om_hub_multicloud_resources", acctest.Optional, acctest.Create, MulticloudOmHubMulticloudResourceDataSourceRepresentation) +
				subscriptionIdVariableStr + subscriptionServiceNameVariableStr + compartmentIdVariableStr + limitVariableStr + MulticloudOmHubMulticloudResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_service_name", subscriptionServiceName),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "limit", limit),

				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.#"),

				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.resource_display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.resource_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.compartment_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.lifecycle_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.defined_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_resource_collection.0.items.0.system_tags.%"),
			),
		},
	})
}
