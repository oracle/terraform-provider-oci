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
	OsManagementHubProfileWithAttachedSSRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"profile_type":        acctest.Representation{RepType: acctest.Required, Create: `SOFTWARESOURCE`},
		"registration_type":   acctest.Representation{RepType: acctest.Required, Create: `OCI_LINUX`},
		"software_source_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`, `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`}},
		"arch_type":           acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
		"os_family":           acctest.Representation{RepType: acctest.Required, Create: `ORACLE_LINUX_8`},
		"vendor_name":         acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubProfileIgnoreDefinedTagsRepresentation},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Finance"}},
	}

	OsManagementHubProfileDetachSoftwareSourcesManagementRepresentation = map[string]interface{}{
		"profile_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_profile.test_profile.id}`},
		"software_sources": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`}},
	}

	OsManagementHubProfileDetachSoftwareSourcesManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubProfileWithAttachedSSRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubProfileDetachSoftwareSourcesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubProfileDetachSoftwareSourcesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_profile_detach_software_sources_management.test_profile_detach_software_sources_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubProfileDetachSoftwareSourcesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile_detach_software_sources_management", "test_profile_detach_software_sources_management", acctest.Required, acctest.Create, OsManagementHubProfileDetachSoftwareSourcesManagementRepresentation), "osmanagementhub", "profileDetachSoftwareSourcesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileDetachSoftwareSourcesManagementResourceDependencies +
				OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile_detach_software_sources_management", "test_profile_detach_software_sources_management", acctest.Required, acctest.Create, OsManagementHubProfileDetachSoftwareSourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "profile_id"),
				resource.TestCheckResourceAttr(resourceName, "software_sources.#", "1"),

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
