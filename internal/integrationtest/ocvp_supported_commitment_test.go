// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	OcvpSupportedCommitmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"host_shape_name": acctest.Representation{RepType: acctest.Optional, Create: "BM.Standard2.52"},
	}

	OcvpSupportedCommitmentResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, CoreCoreShapeDataSourceRepresentation)
)

// issue-routing-tag: ocvp/default
func TestOcvpSupportedCommitmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpSupportedCommitmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_ocvp_supported_commitments.test_supported_commitments"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_supported_commitments", "test_supported_commitments", acctest.Required, acctest.Create, OcvpSupportedCommitmentDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpSupportedCommitmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.name"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_supported_commitments", "test_supported_commitments", acctest.Optional, acctest.Create, OcvpSupportedCommitmentDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpSupportedCommitmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "host_shape_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.name"),
			),
		},
	})
}
