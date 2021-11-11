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
	ipsecAlgorithmSingularDataSourceRepresentation = map[string]interface{}{}

	IpsecAlgorithmResourceConfig = ""
)

// issue-routing-tag: core/default
func TestCoreIpsecAlgorithmResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpsecAlgorithmResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ipsec_algorithm.test_ipsec_algorithm"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_algorithm", "test_ipsec_algorithm", Required, Create, ipsecAlgorithmSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IpsecAlgorithmResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_phase_one_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_phase_two_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_phase_one_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_phase_two_parameters.#", "1"),
			),
		},
	})
}
