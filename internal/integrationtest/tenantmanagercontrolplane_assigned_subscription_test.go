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
	TenantmanagercontrolplaneAssignedSubscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"assigned_subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.assigned_subscription_id}`},
	}

	TenantmanagercontrolplaneAssignedSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.assigned_subscription_compartment_id}`},
		"entity_version":  acctest.Representation{RepType: acctest.Optional, Create: `${var.entity_version}`},
		"subscription_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.assigned_subscription_id}`},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneAssignedSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneAssignedSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("assigned_subscription_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"assigned_subscription_compartment_id\" { default = \"%s\" }\n", compartmentId)

	entityVersion := utils.GetEnvSettingWithBlankDefault("entity_version")
	entityVersionVariableStr := fmt.Sprintf("variable \"entity_version\" { default = \"%s\" }\n", entityVersion)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("assigned_subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"assigned_subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_tenantmanagercontrolplane_assigned_subscriptions.test_assigned_subscriptions"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_assigned_subscription.test_assigned_subscription"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + compartmentIdVariableStr + entityVersionVariableStr + subscriptionIdVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_assigned_subscriptions", "test_assigned_subscriptions", acctest.Optional, acctest.Create, TenantmanagercontrolplaneAssignedSubscriptionDataSourceRepresentation)
	singularDataSourceConfig := config + subscriptionIdVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_assigned_subscription", "test_assigned_subscription", acctest.Required, acctest.Create, TenantmanagercontrolplaneAssignedSubscriptionSingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "entity_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "assigned_subscription_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "assigned_subscription_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "classic_subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entity_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_classic_subscription"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_government_subscription"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "order_ids.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "promotion.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "purchase_entitlement_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "skus.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "start_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
