// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	OsManagementHubEntitlementDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"csi":            acctest.Representation{RepType: acctest.Optional, Create: `csi`},
		"vendor_name":    acctest.Representation{RepType: acctest.Optional, Create: `Oracle`},
	}

	OsManagementHubEntitlementResourceConfig = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubEntitlementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubEntitlementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_entitlements.test_entitlements"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_entitlements", "test_entitlements", acctest.Required, acctest.Create, OsManagementHubEntitlementDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubEntitlementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "entitlement_collection.#"),
			),
		},
	})
}
