// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	TenantmanagercontrolplaneSubscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
	}

	TenantmanagercontrolplaneSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.subscription_compartment_id}`},
		"entity_version":  acctest.Representation{RepType: acctest.Optional, Create: `${var.entity_version}`},
		"subscription_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.subscription_id}`},
	}

	TenantmanagercontrolplaneSubscriptionResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscriptions", "test_subscriptions", acctest.Required, acctest.Create, TenantmanagercontrolplaneSubscriptionDataSourceRepresentation)
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("subscription_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"subscription_compartment_id\" { default = \"%s\" }\n", compartmentId)

	entityVersion := utils.GetEnvSettingWithBlankDefault("subscription_entity_version")
	entityVersionVariableStr := fmt.Sprintf("variable \"entity_version\" { default = \"%s\" }\n", entityVersion)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_tenantmanagercontrolplane_subscriptions.test_subscriptions"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_subscription.test_subscription"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + compartmentIdVariableStr + entityVersionVariableStr + subscriptionIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscriptions", "test_subscriptions", acctest.Optional, acctest.Create, TenantmanagercontrolplaneSubscriptionDataSourceRepresentation)
	singularDataSourceConfig := config + subscriptionIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_subscription", "test_subscription", acctest.Required, acctest.Create, TenantmanagercontrolplaneSubscriptionSingularDataSourceRepresentation)

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

				resource.TestCheckResourceAttrSet(datasourceName, "subscription_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "classic_subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_amount_currency"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "csi_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "customer_country_code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "end_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entity_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_classic_subscription"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_government_subscription"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "payment_model"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "promotion.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "purchase_entitlement_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "skus.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "start_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_tier"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}

func getTenantmanagercontrolplaneSubscriptionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	subscriptionClient := acctest.GetTestClients(&schema.ResourceData{}).OrganizationsSubscriptionClient()

	listSubscriptionsRequest := oci_tenantmanagercontrolplane.ListSubscriptionsRequest{}
	listSubscriptionsRequest.CompartmentId = &compartmentId
	listSubscriptionsResponse, err := subscriptionClient.ListSubscriptions(context.Background(), listSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Subscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subscription := range listSubscriptionsResponse.Items {
		id := *subscription.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubscriptionId", id)
	}
	return resourceIds, nil
}
