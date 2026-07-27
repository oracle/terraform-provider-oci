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

const tenantmanagercontrolplaneAssignedSubscriptionsDataSourceConfig = `
data "oci_tenantmanagercontrolplane_assigned_subscriptions" "test_assigned_subscriptions" {
	compartment_id  = %q
	subscription_id = %q
}
`

const tenantmanagercontrolplaneAssignedSubscriptionSingularDataSourceConfig = `
data "oci_tenantmanagercontrolplane_assigned_subscription" "test_assigned_subscription" {
	assigned_subscription_id = %q
	compartment_id            = %q
}
`

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneAssignedSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneAssignedSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId, assignedSubscriptionId := requiredTenantmanagercontrolplaneAssignedSubscriptionFixture(t)

	datasourceName := "data.oci_tenantmanagercontrolplane_assigned_subscriptions.test_assigned_subscriptions"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_assigned_subscription.test_assigned_subscription"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + fmt.Sprintf(tenantmanagercontrolplaneAssignedSubscriptionsDataSourceConfig, compartmentId, assignedSubscriptionId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", assignedSubscriptionId),
				resource.TestCheckResourceAttr(datasourceName, "assigned_subscription_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "assigned_subscription_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "assigned_subscription_collection.0.items.0.id", assignedSubscriptionId),
				resource.TestCheckResourceAttr(datasourceName, "assigned_subscription_collection.0.items.0.state", "ACTIVE"),
			),
		},
		{
			Config: config + fmt.Sprintf(tenantmanagercontrolplaneAssignedSubscriptionSingularDataSourceConfig, assignedSubscriptionId, compartmentId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "assigned_subscription_id", assignedSubscriptionId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entity_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", assignedSubscriptionId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}

func requiredTenantmanagercontrolplaneAssignedSubscriptionFixture(t *testing.T) (string, string) {
	t.Helper()

	compartmentId := utils.GetEnvSettingWithBlankDefault("assigned_subscription_compartment_id")
	if compartmentId == "" {
		compartmentId = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	}
	if compartmentId == "" {
		t.Fatal("TF_VAR_assigned_subscription_compartment_id or TF_VAR_tenancy_ocid must be set for this acceptance test")
	}

	assignedSubscriptionId := utils.GetEnvSettingWithBlankDefault("assigned_subscription_id")
	if assignedSubscriptionId == "" {
		t.Fatal("TF_VAR_assigned_subscription_id must be set for this acceptance test")
	}

	return compartmentId, assignedSubscriptionId
}
