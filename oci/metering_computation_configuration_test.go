// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	usageapiConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"tenant_id": Representation{repType: Required, create: `${var.tenancy_id}`},
	}

	usageapiConfigurationResourceConfig = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	singularDatasourceName := "data.oci_metering_computation_configuration.test_configuration"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + tenancyIdVariableStr +
				generateDataSourceFromRepresentationMap("oci_metering_computation_configuration", "test_configuration", Required, Create, usageapiConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + usageapiConfigurationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
