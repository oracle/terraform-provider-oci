// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	OcvpRetrievePasswordDataSourceRepresentation = map[string]interface{}{
		"sddc_id": acctest.Representation{RepType: acctest.Required, Create: `${local.upgraded_sddc_id}`},
		"type":    acctest.Representation{RepType: acctest.Required, Create: "VCENTER"},
	}
)

// issue-routing-tag: ocvp/default
func TestRetrievePasswordDataSource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRetrievePasswordDataSource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_ocvp_retrieve_password.test_password"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + sddcDataSourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_retrieve_password", "test_password", acctest.Required, acctest.Create, OcvpRetrievePasswordDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_id"),
				resource.TestCheckResourceAttr(datasourceName, "type", "VCENTER"),

				resource.TestCheckResourceAttrSet(datasourceName, "sddc_password.value"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_password.passwordType", "VCENTER"),
			),
		},
	})
}
