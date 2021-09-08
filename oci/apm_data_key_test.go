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
	dataKeyDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"data_key_type": Representation{repType: Optional, create: `PRIVATE`},
	}

	DataKeyResourceConfig = generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation)
)

// issue-routing-tag: apm/default
func TestApmDataKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmDataKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_apm_data_keys.test_data_keys"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_apm_data_keys", "test_data_keys", Required, Create, dataKeyDataSourceRepresentation) +
				compartmentIdVariableStr + DataKeyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),

				resource.TestCheckResourceAttr(datasourceName, "data_keys.#", "2"),

				resource.TestCheckResourceAttr(datasourceName, "data_keys.0.name", "auto_generated_private_datakey"),
				resource.TestCheckResourceAttr(datasourceName, "data_keys.0.type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_keys.0.value"),

				resource.TestCheckResourceAttr(datasourceName, "data_keys.1.name", "auto_generated_public_datakey"),
				resource.TestCheckResourceAttr(datasourceName, "data_keys.1.type", "PUBLIC"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_keys.1.value"),
			),
		},
	})
}
