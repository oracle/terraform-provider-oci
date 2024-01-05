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
	UsageProxyResourceQuotaDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_name":        acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
		"service_entitlement": acctest.Representation{RepType: acctest.Optional, Create: `serviceEntitlement`},
	}

	UsageProxyResourceQuotaResourceConfig = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxyResourceQuotaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxyResourceQuotaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	datasourceName := "data.oci_usage_proxy_resource_quotas.test_resource_quotas"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_resource_quotas", "test_resource_quotas", acctest.Required, acctest.Create, UsageProxyResourceQuotaDataSourceRepresentation) +
				compartmentIdVariableStr + serviceNameVariableStr + UsageProxyResourceQuotaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "service_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "resource_quotum_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_quotum_collection.0.is_allowed"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_quotum_collection.0.items.#"),
			),
		},
	})
}
