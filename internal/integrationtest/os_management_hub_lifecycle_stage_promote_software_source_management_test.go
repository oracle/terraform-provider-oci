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
	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRequiredOnlyResource = OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Required, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation)

	OsManagementHubVersionedCustomSoftwareSourceRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubDefinedTagsRepresentation},
		"software_source_version":       acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `tf-custom-vcss`, Update: `tf-custom-vcss2`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `tf-custom-vcss`, Update: `tf-custom-vcss2`},
		"software_source_type":          acctest.Representation{RepType: acctest.Required, Create: `VERSIONED`},
		"is_latest_content_only":        acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"vendor_software_sources":       []acctest.RepresentationGroup{{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation}, {RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation2}},
		"is_auto_resolve_dependencies":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_automatically_updated":      acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_created_from_package_list":  acctest.Representation{RepType: acctest.Required, Create: `false`},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreVendorSSChangesRepresentation},
		"custom_software_source_filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVersionedCustomSoftwareSourceFilterRepresentation},
	}

	OsManagementHubSoftwareSourceVersionedCustomSoftwareSourceFilterRepresentation = map[string]interface{}{
		"module_stream_profile_filters": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubSoftwareSourceCustomSoftwareSourceFilterModuleStreamProfileFiltersRepresentation},
		"package_filters":               acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVersionedCustomSoftwareSourceFilterPackageFiltersRepresentation},
		"package_group_filters":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubSoftwareSourceCustomSoftwareSourceFilterPackageGroupFiltersRepresentation},
	}

	OsManagementHubSoftwareSourceVersionedCustomSoftwareSourceFilterPackageFiltersRepresentation = map[string]interface{}{
		"filter_type":     acctest.Representation{RepType: acctest.Required, Create: `INCLUDE`},
		"package_name":    acctest.Representation{RepType: acctest.Required, Create: `zsh`},
		"package_version": acctest.Representation{RepType: acctest.Required, Create: `5.5.1-10.el8`},
	}

	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation = map[string]interface{}{
		"lifecycle_stage_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id}`},
		"software_source_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubLifecycleStagePromoteSoftwareSourceManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubLifecycleStagePromoteSoftwareSourceManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_lifecycle_stage_promote_software_source_management.test_lifecycle_stage_promote_software_source_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation), "osmanagementhub", "lifecycleStagePromoteSoftwareSourceManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubVersionedCustomSoftwareSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Required, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Optional, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubVersionedCustomSoftwareSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_id"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.display_name", "displayName"),

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
