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
	serviceDataSourceRepresentation = map[string]interface{}{}

	ServiceResourceConfig = ""
)

// issue-routing-tag: core/serviceGateway
func TestCoreServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_services.test_services"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "services.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.cidr_block"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.name"),
			),
		},
	})
}
