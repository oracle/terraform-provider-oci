// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	SelfPartnerDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	SelfPartnerResourceConfig = ""
)

// issue-routing-tag: self/default
func TestSelfPartnerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSelfPartnerResource_basic")
	defer httpreplay.SaveScenario()

	// config := acctest.ProviderTestConfig()

	// compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	// compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// datasourceName := "data.oci_self_partners.test_partners"

	// acctest.SaveConfigContent("", "", "", t)

	/*acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_self_partners", "test_partners", acctest.Optional, acctest.Create, SelfPartnerDataSourceRepresentation) +
				compartmentIdVariableStr + SelfPartnerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "partner_collection.#", "1"),
			),
		},
	})*/
}
