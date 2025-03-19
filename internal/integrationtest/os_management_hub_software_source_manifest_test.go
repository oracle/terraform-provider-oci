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
	OsManagementHubSoftwareSourceManifestResourceConfig = OsManagementHubSoftwareSourceManifestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_manifest", "test_software_source_manifest", acctest.Optional, acctest.Update, OsManagementHubSoftwareSourceManifestRepresentation)

	OsManagementHubSoftwareSourceManifestSingularDataSourceRepresentation = map[string]interface{}{
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
	}

	OsManagementHubSoftwareSourceManifestRepresentation = map[string]interface{}{
		"content":            acctest.Representation{RepType: acctest.Required, Create: "ed-1.14.2-4.el8.x86_64"},
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
	}

	OsManagementHubSoftwareSourceManifestResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceCustomRepresentation) + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceManifestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceManifestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_software_source_manifest.test_software_source_manifest"

	singularDatasourceName := "data.oci_os_management_hub_software_source_manifest.test_software_source_manifest"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubSoftwareSourceManifestResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_manifest", "test_software_source_manifest", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceManifestRepresentation), "osmanagementhub", "softwareSourceManifest", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubSoftwareSourceManifestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_manifest", "test_software_source_manifest", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceManifestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "ed-1.14.2-4.el8.x86_64"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + OsManagementHubSoftwareSourceManifestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_manifest", "test_software_source_manifest", acctest.Optional, acctest.Update, OsManagementHubSoftwareSourceManifestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "ed-1.14.2-4.el8.x86_64"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_manifest", "test_software_source_manifest", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceManifestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubSoftwareSourceManifestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},
	})
}
