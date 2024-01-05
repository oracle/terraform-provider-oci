// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	dataKeyDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"data_key_type": acctest.Representation{RepType: acctest.Optional, Create: `PRIVATE`},
	}

	DataKeyResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm/default
func TestApmDataKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmDataKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_apm_data_keys.test_data_keys"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_data_keys", "test_data_keys", acctest.Required, acctest.Create, dataKeyDataSourceRepresentation) +
				compartmentIdVariableStr + DataKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
