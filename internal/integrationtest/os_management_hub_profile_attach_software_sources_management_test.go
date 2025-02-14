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
	OsManagementHubProfileAttachSoftwareSourcesManagementRepresentation = map[string]interface{}{
		"profile_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_profile.test_profile.id}`},
		"software_sources": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`}},
	}

	OsManagementHubProfileAttachSoftwareSourcesManagementResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stages", "test_lifecycle_stages", acctest.Required, acctest.Create, OsManagementHubLifecycleStageDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubProfileRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubProfileAttachSoftwareSourcesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubProfileAttachSoftwareSourcesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_profile_attach_software_sources_management.test_profile_attach_software_sources_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubProfileAttachSoftwareSourcesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile_attach_software_sources_management", "test_profile_attach_software_sources_management", acctest.Required, acctest.Create, OsManagementHubProfileAttachSoftwareSourcesManagementRepresentation), "osmanagementhub", "profileAttachSoftwareSourcesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileAttachSoftwareSourcesManagementResourceDependencies + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile_attach_software_sources_management", "test_profile_attach_software_sources_management", acctest.Required, acctest.Create, OsManagementHubProfileAttachSoftwareSourcesManagementRepresentation),
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
