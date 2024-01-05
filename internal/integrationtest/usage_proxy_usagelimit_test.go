// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	UsageProxyUsagelimitDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"limit_type":      acctest.Representation{RepType: acctest.Optional, Create: `limitType`},
		"resource_type":   acctest.Representation{RepType: acctest.Optional, Create: `resourceType`},
		"service_type":    acctest.Representation{RepType: acctest.Optional, Create: `serviceType`},
	}

	UsageProxyUsagelimitResourceConfig = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxyUsagelimitResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxyUsagelimitResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_usage_proxy_usagelimits.test_usagelimits"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_usagelimits", "test_usagelimits", acctest.Required, acctest.Create, UsageProxyUsagelimitDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + UsageProxyUsagelimitResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "usage_limit_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "usage_limit_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "usage_limit_collection.0.items.0.entitlement_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "usage_limit_collection.0.items.0.limit_type"),
			),
		},
	})
}
