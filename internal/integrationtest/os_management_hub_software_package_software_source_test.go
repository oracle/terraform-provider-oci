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
	OsManagementHubSoftwarePackageSoftwareSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"software_package_name": acctest.Representation{RepType: acctest.Required, Create: `ModemManager-glib-devel-1.10.4-1.el8.x86_64.rpm`},
		"arch_type":             acctest.Representation{RepType: acctest.Optional, Create: []string{`X86_64`}},
		"availability":          acctest.Representation{RepType: acctest.Optional, Create: []string{`SELECTED`}},
		"availability_anywhere": acctest.Representation{RepType: acctest.Optional, Create: []string{`SELECTED`}},
		"availability_at_oci":   acctest.Representation{RepType: acctest.Optional, Create: []string{`SELECTED`}},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `ol8_codeready_builder-x86_64`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `ol8_codeready_builder-x86_64`},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_LINUX_8`}},
		"software_source_type":  acctest.Representation{RepType: acctest.Optional, Create: []string{`VENDOR`}},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
	}

	OsManagementHubSoftwarePackageSoftwareSourceResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_packages", "test_software_packages", acctest.Required, acctest.Create, OsManagementHubSoftwarePackageDataSourceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwarePackageSoftwareSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwarePackageSoftwareSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_software_package_software_source.test_software_package_software_source"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_package_software_source", "test_software_package_software_source", acctest.Required, acctest.Create, OsManagementHubSoftwarePackageSoftwareSourceDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubSoftwarePackageSoftwareSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_collection.0.items.0.arch_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_collection.0.items.0.availability"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_collection.0.items.0.availability_at_oci"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.0.display_name", "ol8_codeready_builder-x86_64"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.0.os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_package_name"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.0.software_source_type", "VENDOR"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.0.state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_source_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.#", "1"),
			),
		},
		// verify data source with optional
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_package_software_source", "test_software_package_software_source", acctest.Optional, acctest.Create, OsManagementHubSoftwarePackageSoftwareSourceDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubSoftwarePackageSoftwareSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_collection.0.items.0.arch_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_collection.0.items.0.availability"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_collection.0.items.0.availability_at_oci"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.0.display_name", "ol8_codeready_builder-x86_64"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "ol8_codeready_builder-x86_64"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.0.os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_package_name"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.0.software_source_type", "VENDOR"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.0.state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_source_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.0.items.#", "1"),
			),
		},
	})
}
