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
	TenantmanagercontrolplaneSubscriptionLineItemDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.v2_subscription_id}`},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneSubscriptionLineItemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneSubscriptionLineItemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	v2SubscriptionId := utils.GetEnvSettingWithBlankDefault("v2_subscription_id")
	v2SubscriptionIdVariableStr := fmt.Sprintf("variable \"v2_subscription_id\" { default = \"%s\" }\n", v2SubscriptionId)

	datasourceName := "data.oci_tenantmanagercontrolplane_subscription_line_items.test_subscription_line_items"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + v2SubscriptionIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscription_line_items", "test_subscription_line_items", acctest.Required, acctest.Create, TenantmanagercontrolplaneSubscriptionLineItemDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "subscription_line_item_collection.#"),
			),
		},
	})
}
