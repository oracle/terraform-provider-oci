// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

const tenantmanagercontrolplaneAssignedSubscriptionAvailableRegionDataSourceConfig = `
data "oci_tenantmanagercontrolplane_assigned_subscription_available_regions" "test_assigned_subscription_available_regions" {
	assigned_subscription_id = %q
}
`

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneAssignedSubscriptionAvailableRegionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneAssignedSubscriptionAvailableRegionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	_, assignedSubscriptionId := requiredTenantmanagercontrolplaneAssignedSubscriptionFixture(t)

	datasourceName := "data.oci_tenantmanagercontrolplane_assigned_subscription_available_regions.test_assigned_subscription_available_regions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + fmt.Sprintf(tenantmanagercontrolplaneAssignedSubscriptionAvailableRegionDataSourceConfig, assignedSubscriptionId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "assigned_subscription_id", assignedSubscriptionId),
				resource.TestCheckResourceAttr(datasourceName, "available_region_collection.#", "1"),
			),
		},
	})
}
