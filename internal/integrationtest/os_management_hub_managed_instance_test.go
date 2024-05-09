// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	//"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	//"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubManagedInstanceRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceRepresentation)

	//OsManagementHubManagedInstanceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, OsManagementHubManagedInstanceRepresentation)

	OsManagementHubManagedInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")},
	}

	OsManagementHubManagedInstanceDataSourceRepresentation = map[string]interface{}{
		"advisory_name":         acctest.Representation{RepType: acctest.Optional, Create: []string{`advisoryName`}},
		"arch_type":             acctest.Representation{RepType: acctest.Optional, Create: []string{`X86_64`}},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
		"group":                 acctest.Representation{RepType: acctest.Optional, Create: `group`},
		"group_not_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `groupNotEqualTo`},
		"is_attached_to_group_or_lifecycle_stage": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_managed_by_autonomous_linux":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_management_station":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_profile_attached":                     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle_environment":                   acctest.Representation{RepType: acctest.Optional, Create: `lifecycleEnvironment`},
		"lifecycle_environment_not_equal_to":      acctest.Representation{RepType: acctest.Optional, Create: `lifecycleEnvironmentNotEqualTo`},
		"location":                                acctest.Representation{RepType: acctest.Optional, Create: []string{`OCI_COMPUTE`}},
		"location_not_equal_to":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`OCI_COMPUTE`}},
		"managed_instance_id":                     acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")},
		"os_family":                               acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_LINUX_8`}},
		"software_source_id":                      acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("osmh_software_source_ocid")},
		"status":                                  acctest.Representation{RepType: acctest.Optional, Create: []string{`NORMAL`}},
		"filter":                                  acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagedInstanceDataSourceFilterRepresentation}}

	OsManagementHubManagedInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `managed_instance_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")}},
	}
	OsManagementHubManagedInstanceLifecycleEnvironmentDataSourceRepresentation = map[string]interface{}{}
	OsManagementHubManagedInstanceLifecycleStageDataSourceRepresentation       = map[string]interface{}{}

	OsManagementHubLifecycleEnvironmentStageRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"rank":         acctest.Representation{RepType: acctest.Required, Create: `1`},
		"defined_tags": acctest.Representation{RepType: acctest.Optional, Create: ignoreDefinedTagsChangesForOsmhLERep},
	}
	OsManagementHubManagedInstanceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}
	OsManagementHubManagedInstanceRegistrationFailureRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_failed_ocid")},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}
	OsManagementHubManagedInstanceWindowsRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_windows_ocid")},
	}
	OsManagementHubManagedInstanceAutonomousSettingsRepresentation = map[string]interface{}{
		"is_data_collection_authorized": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instance.test_managed_instance"
	datasourceName := "data.oci_os_management_hub_managed_instances.test_managed_instances"
	singularDatasourceName := "data.oci_os_management_hub_managed_instance.test_managed_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceRepresentation), "osmanagementhub", "managedInstance", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//verify resource import
		{
			Config:            config + OsManagementHubManagedInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"managed_instance_id", "time_last_boot", "time_last_checkin",
			},
			ResourceName: resourceName,
		},

		//delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		//// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

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
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, OsManagementHubManagedInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instances", "test_managed_instances", acctest.Optional, acctest.Update, OsManagementHubManagedInstanceDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Optional, acctest.Update, OsManagementHubManagedInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "advisory_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "arch_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNameContains"),
				resource.TestCheckResourceAttr(datasourceName, "group", "group"),
				resource.TestCheckResourceAttr(datasourceName, "group_not_equal_to", "groupNotEqualTo"),
				resource.TestCheckResourceAttr(datasourceName, "is_attached_to_group_or_lifecycle_stage", "false"),
				resource.TestCheckResourceAttr(datasourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(datasourceName, "is_management_station", "false"),
				resource.TestCheckResourceAttr(datasourceName, "is_profile_attached", "false"),
				resource.TestCheckResourceAttr(datasourceName, "lifecycle_environment_not_equal_to", "lifecycleEnvironmentNotEqualTo"),
				resource.TestCheckResourceAttr(datasourceName, "location.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "location_not_equal_to.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "os_family.#", "1"),
				//resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),
				resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "managed_instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "architecture"),
				resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_settings.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bug_updates_available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enhancement_updates_available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "installed_packages"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "installed_windows_updates"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_managed_by_autonomous_linux"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_management_station"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reboot_required"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ksplice_effective_kernel_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "location"),
				resource.TestCheckResourceAttr(singularDatasourceName, "managed_instance_group.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_family"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_kernel_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "other_updates_available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_job_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_updates_available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_sources.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_boot"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_checkin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "updates_available"),
			),
		},
	})
}
