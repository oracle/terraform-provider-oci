// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	usageapiConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"tenant_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
	}

	usageapiConfigurationResourceConfig = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	singularDatasourceName := "data.oci_metering_computation_configuration.test_configuration"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + tenancyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_configuration", "test_configuration", acctest.Required, acctest.Create, usageapiConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + usageapiConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
