// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubSoftwareSourceRemovePackagesManagementRepresentation = map[string]interface{}{
		"packages":           acctest.Representation{RepType: acctest.Required, Create: []string{`R-core-4.0.5-1.0.1.el8.x86_64`}},
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
		"depends_on":         acctest.Representation{RepType: acctest.Required, Create: []string{`oci_os_management_hub_software_source_add_packages_management.test_software_source_add_packages_management`, `oci_os_management_hub_software_source.test_software_source`}},
	}

	ignoreVSSandIsAutoUpdatedChanges = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"vendor_software_sources", "is_automatically_updated"}},
	}

	OsManagementHubCustomSoftwareSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubIgnoreDefinedTagsRepresentation},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"software_source_type":         acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"software_source_sub_type":     acctest.Representation{RepType: acctest.Required, Create: `MANIFEST`},
		"vendor_software_sources":      []acctest.RepresentationGroup{{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation}, {RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation2}},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"is_auto_resolve_dependencies": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"is_automatically_updated":     acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_created_from_package_list": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_latest_content_only":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreVSSandIsAutoUpdatedChanges},
	}

	OsManagementHubCustomSoftwareSourceRepresentation2 = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubDefinedTagsRepresentation},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `tf-custom-vcss`, Update: `tf-custom-vcss2`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `tf-custom-vcss`, Update: `tf-custom-vcss2`},
		"is_latest_content_only":       acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"software_source_type":         acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"software_source_sub_type":     acctest.Representation{RepType: acctest.Required, Create: `MANIFEST`},
		"vendor_software_sources":      []acctest.RepresentationGroup{{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation}, {RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation2}, {RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation3}},
		"is_auto_resolve_dependencies": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"is_automatically_updated":     acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_created_from_package_list": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreVendorSSChangesRepresentation},
	}

	OsManagementHubSoftwareSourceRemovePackagesManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubCustomSoftwareSourceRepresentation2)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceRemovePackagesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceRemovePackagesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_software_source_remove_packages_management.test_software_source_remove_packages_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubSoftwareSourceRemovePackagesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_remove_packages_management", "test_software_source_remove_packages_management", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceRemovePackagesManagementRepresentation), "osmanagementhub", "softwareSourceRemovePackagesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubSoftwareSourceRemovePackagesManagementResourceDependencies +
				OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AddsOnX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_remove_packages_management", "test_software_source_remove_packages_management", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceRemovePackagesManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_add_packages_management", "test_software_source_add_packages_management", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OsManagementHubSoftwareSourceAddPackagesManagementRepresentation, map[string]interface{}{
						"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
						"packages":           acctest.Representation{RepType: acctest.Required, Create: []string{`R-core-4.0.5-1.0.1.el8.x86_64`, `NetworkManager-adsl-1:1.30.0-13.0.1.el8_4.x86_64`}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
