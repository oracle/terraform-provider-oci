// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PsaPsaServiceDataSourceRepresentation = map[string]interface{}{}

	PsaPsaServiceResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_psa_psa_services", "test_psa_services", acctest.Optional, acctest.Create, PsaPsaServiceDataSourceRepresentation)
)

// issue-routing-tag: psa/default
func TestPsaPsaServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsaPsaServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_psa_psa_services.test_psa_services"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + PsaPsaServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "psa_service_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "psa_service_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "psa_service_collection.0.items.0.id"),
			),
		},
	})
}
