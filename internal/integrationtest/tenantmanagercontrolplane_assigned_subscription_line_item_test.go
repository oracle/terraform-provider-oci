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
	TenantmanagercontrolplaneAssignedSubscriptionLineItemDataSourceRepresentation = map[string]interface{}{
		"assigned_subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.v2_assigned_subscription_id}`},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneAssignedSubscriptionLineItemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneAssignedSubscriptionLineItemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	v2AssignedSubscriptionId := utils.GetEnvSettingWithBlankDefault("v2_assigned_subscription_id")
	v2AssignedSubscriptionIdVariableStr := fmt.Sprintf("variable \"v2_assigned_subscription_id\" { default = \"%s\" }\n", v2AssignedSubscriptionId)

	datasourceName := "data.oci_tenantmanagercontrolplane_assigned_subscription_line_items.test_assigned_subscription_line_items"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + v2AssignedSubscriptionIdVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_assigned_subscription_line_items", "test_assigned_subscription_line_items", acctest.Required, acctest.Create, TenantmanagercontrolplaneAssignedSubscriptionLineItemDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "assigned_subscription_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "assigned_subscription_line_item_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "assigned_subscription_line_item_collection.0.items.#"),
			),
		},
	})
}
