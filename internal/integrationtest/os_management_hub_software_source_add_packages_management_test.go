// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubSoftwareSourceAddPackagesManagementRequiredOnlyResource = OsManagementHubSoftwareSourceAddPackagesManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_add_packages_management", "test_software_source_add_packages_management", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceAddPackagesManagementRepresentation)

	ignoreVendorSSChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`vendor_software_sources`}},
	}

	OsManagementHubSoftwareSourceAddPackagesManagementRepresentation = map[string]interface{}{
		"packages":                        acctest.Representation{RepType: acctest.Required, Create: []string{`NetworkManager-adsl-1:1.30.0-13.0.1.el8_4.x86_64`}},
		"software_source_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
		"is_continue_on_missing_packages": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	OsManagementHubSoftwareSourceAddPackagesManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceCustomRepresentation) + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceAddPackagesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceAddPackagesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_software_source_add_packages_management.test_software_source_add_packages_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubSoftwareSourceAddPackagesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_add_packages_management", "test_software_source_add_packages_management", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceAddPackagesManagementRepresentation), "osmanagementhub", "softwareSourceAddPackagesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubSoftwareSourceAddPackagesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_add_packages_management", "test_software_source_add_packages_management", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceAddPackagesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "packages.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubSoftwareSourceAddPackagesManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubSoftwareSourceAddPackagesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_add_packages_management", "test_software_source_add_packages_management", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceAddPackagesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_continue_on_missing_packages", "false"),
				resource.TestCheckResourceAttr(resourceName, "packages.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
