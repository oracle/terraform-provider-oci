// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
)

var (
	regionSubscriptionDataSourceRepresentation = map[string]interface{}{
		"tenancy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"filter":     acctest.RepresentationGroup{RepType: acctest.Required, Group: regionSubscriptionDataSourceFilterRepresentation},
	}

	regionSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `is_home_region`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`true`}},
	}

	RegionSubscriptionResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityRegionSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityRegionSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_identity_region_subscriptions.test_region_subscriptions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_region_subscriptions", "test_region_subscriptions", acctest.Required, acctest.Create, regionSubscriptionDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),
				resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.0.is_home_region", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "region_subscriptions.0.region_key"),
				resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.0.region_name", utils.GetRequiredEnvSetting("region")),
				resource.TestCheckResourceAttrSet(datasourceName, "region_subscriptions.0.state"),
			),
		},
	})
}
