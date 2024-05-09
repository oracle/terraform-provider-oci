// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationDisableDryRun = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("managed_instance_group_ocid")},
		"disable":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementDisableRepresentation},
		"is_dry_run":                acctest.Representation{RepType: acctest.Required, Create: `true`},
		"work_request_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationDisable = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("managed_instance_group_ocid")},
		"disable":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementDisableRepresentation},
		"is_dry_run":                acctest.Representation{RepType: acctest.Required, Create: `false`},
		"work_request_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationEnable = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("managed_instance_group_ocid")},
		"enable":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementEnableRepresentation},
		"is_dry_run":                acctest.Representation{RepType: acctest.Required, Create: `false`},
		"work_request_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationInstall = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("managed_instance_group_ocid")},
		"install":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementInstallRepresentation},
		"is_dry_run":                acctest.Representation{RepType: acctest.Required, Create: `false`},
		"work_request_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationRemove = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("managed_instance_group_ocid")},
		"remove":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRemoveRepresentation},
		"is_dry_run":                acctest.Representation{RepType: acctest.Required, Create: `false`},
		"work_request_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupManageModuleStreamsManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementDisableRepresentation = map[string]interface{}{
		"module_name":        acctest.Representation{RepType: acctest.Required, Create: `php`},
		"stream_name":        acctest.Representation{RepType: acctest.Required, Create: `7.2`},
		"software_source_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementEnableRepresentation = map[string]interface{}{
		"module_name":        acctest.Representation{RepType: acctest.Required, Create: `php`},
		"stream_name":        acctest.Representation{RepType: acctest.Required, Create: `7.2`},
		"software_source_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementInstallRepresentation = map[string]interface{}{
		"module_name":        acctest.Representation{RepType: acctest.Required, Create: `php`},
		"profile_name":       acctest.Representation{RepType: acctest.Required, Create: `minimal`},
		"stream_name":        acctest.Representation{RepType: acctest.Required, Create: `7.2`},
		"software_source_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRemoveRepresentation = map[string]interface{}{
		"module_name":        acctest.Representation{RepType: acctest.Required, Create: `php`},
		"profile_name":       acctest.Representation{RepType: acctest.Required, Create: `minimal`},
		"stream_name":        acctest.Representation{RepType: acctest.Required, Create: `7.2`},
		"software_source_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
	}
	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies = OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupManageModuleStreamsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupManageModuleStreamsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instance_group_manage_module_streams_management.test_managed_instance_group_manage_module_streams_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_manage_module_streams_management", "test_managed_instance_group_manage_module_streams_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationDisable), "osmanagementhub", "managedInstanceGroupManageModuleStreamsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Disable dry run
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_manage_module_streams_management", "test_managed_instance_group_manage_module_streams_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationDisableDryRun),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
			),
		},
		// Disable
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_manage_module_streams_management", "test_managed_instance_group_manage_module_streams_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationDisable),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies,
		},
		// verify Disable Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_manage_module_streams_management", "test_managed_instance_group_manage_module_streams_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationDisable),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "disable.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "disable.0.module_name", "php"),
				resource.TestCheckResourceAttrSet(resourceName, "disable.0.software_source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "disable.0.stream_name"),
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
		// Enable
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_manage_module_streams_management", "test_managed_instance_group_manage_module_streams_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationEnable),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
			),
		},
		// Install
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_manage_module_streams_management", "test_managed_instance_group_manage_module_streams_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationInstall),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
			),
		},
		// Remove
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_manage_module_streams_management", "test_managed_instance_group_manage_module_streams_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupManageModuleStreamsManagementRepresentationRemove),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
			),
		},
	})
}
