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
	OsManagementHubSoftwareSourceVendorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorDataSourceFilterRepresentation}}
	OsManagementHubSoftwareSourceVendorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`ORACLE`}},
	}

	OsManagementHubSoftwareSourceVendorResourceConfig = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceVendorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceVendorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_software_source_vendors.test_software_source_vendors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubSoftwareSourceVendorResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_vendors", "test_software_source_vendors", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceVendorDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "software_source_vendor_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_vendor_collection.0.items.#", "1"),
			),
		},
	})
}
