// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	regionSubscriptionDataSourceRepresentation = map[string]interface{}{
		"tenancy_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"filter":     RepresentationGroup{Required, regionSubscriptionDataSourceFilterRepresentation},
	}

	regionSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `is_home_region`},
		"values": Representation{RepType: Required, Create: []string{`true`}},
	}

	RegionSubscriptionResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityRegionSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityRegionSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	datasourceName := "data.oci_identity_region_subscriptions.test_region_subscriptions"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_region_subscriptions", "test_region_subscriptions", Required, Create, regionSubscriptionDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),
				resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.0.is_home_region", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "region_subscriptions.0.region_key"),
				resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.0.region_name", getRequiredEnvSetting("region")),
				resource.TestCheckResourceAttrSet(datasourceName, "region_subscriptions.0.state"),
			),
		},
	})
}
