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
	UsageProxyResourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_name":   acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
		"entitlement_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.entitlement_id}`},
	}

	UsageProxyResourceResourceConfig = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxyResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxyResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	entitlementId := utils.GetEnvSettingWithBlankDefault("entitlement_id")
	entitlementIdVariableStr := fmt.Sprintf("variable \"entitlement_id\" { default = \"%s\" }\n", entitlementId)

	datasourceName := "data.oci_usage_proxy_resources.test_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_resources", "test_resources", acctest.Required, acctest.Create, UsageProxyResourceDataSourceRepresentation) +
				compartmentIdVariableStr + serviceNameVariableStr + entitlementIdVariableStr + UsageProxyResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "service_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "resources_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "resources_collection.0.items.#"),
			),
		},
	})
}
