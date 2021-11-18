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
	virtualCircuitPublicPrefixDataSourceRepresentation = map[string]interface{}{
		"virtual_circuit_id": Representation{RepType: Required, Create: `${oci_core_virtual_circuit.test_virtual_circuit.id}`},
		"verification_state": Representation{RepType: Optional, Create: `COMPLETED`},
	}

	VirtualCircuitPublicPrefixResourceConfig = VirtualCircuitPublicPropertyVariables +
		VirtualCircuitResourceDependenciesCopyForVC +
		GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyRepresentation)
)

// issue-routing-tag: core/default
func TestCoreVirtualCircuitPublicPrefixResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVirtualCircuitPublicPrefixResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	secretIdCKN := getEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := getEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	secretVersionCAK := getEnvSettingWithBlankDefault("secret_version_cak")
	secretVersionStrCAK := fmt.Sprintf("variable \"secret_version_cak\" { default = \"%s\" }\n", secretVersionCAK)

	secretVersionCKN := getEnvSettingWithBlankDefault("secret_version_ckn")
	secretVersionStrCKN := fmt.Sprintf("variable \"secret_version_ckn\" { default = \"%s\" }\n", secretVersionCKN)

	datasourceName := "data.oci_core_virtual_circuit_public_prefixes.test_virtual_circuit_public_prefixes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuit_public_prefixes", "test_virtual_circuit_public_prefixes", Required, Create, virtualCircuitPublicPrefixDataSourceRepresentation) +
				compartmentIdVariableStr + VirtualCircuitPublicPrefixResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_public_prefixes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_public_prefixes.0.cidr_block"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_public_prefixes.0.verification_state"),
			),
		},
	})
}
