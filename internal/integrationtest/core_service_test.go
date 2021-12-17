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
	serviceDataSourceRepresentation = map[string]interface{}{}

	ServiceResourceConfig = ""
)

// issue-routing-tag: core/serviceGateway
func TestCoreServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_services.test_services"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, serviceDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "services.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.cidr_block"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.name"),
			),
		},
	})
}
