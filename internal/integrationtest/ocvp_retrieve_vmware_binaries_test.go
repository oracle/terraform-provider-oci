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
	OcvpRetrieveVmwareBinariesDataSourceRepresentation = map[string]interface{}{
		"sddc_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sddc_id}`},
	}
)

// issue-routing-tag: ocvp/default
func TestRetrieveVmwareBinariesDataSource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRetrieveVmwareBinariesDataSource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	sddcId := utils.GetEnvSettingWithBlankDefault("ocvp_sddc_id")
	if sddcId == "" {
		t.Skip("Skipping test because env var 'ocvp_sddc_id' is not set")
	}
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	sddcIdVariableStr := fmt.Sprintf("variable \"sddc_id\" { default = \"%s\" }\n", sddcId)

	datasourceName := "data.oci_ocvp_retrieve_vmware_binaries.test_retrieve_vmware_binaries"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + sddcIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_retrieve_vmware_binaries", "test_retrieve_vmware_binaries", acctest.Required, acctest.Create, OcvpRetrieveVmwareBinariesDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "sddc_id", sddcId),
				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.file_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.description"),
			),
		},
	})
}
