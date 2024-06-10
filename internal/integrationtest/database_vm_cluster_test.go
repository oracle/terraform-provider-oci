// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseVmClusterRequiredOnlyResource = DatabaseVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation)

	DatabaseVmClusterResourceConfig = DatabaseVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Update, DatabaseVmClusterRepresentation)

	DatabaseDatabaseVmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	DatabaseDatabaseVmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `vmCluster`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseVmClusterDataSourceFilterRepresentation}}
	DatabaseVmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_vm_cluster.test_vm_cluster.id}`}},
	}

	DatabaseVmClusterRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":                  acctest.Representation{RepType: acctest.Required, Create: `4`, Update: `6`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `vmCluster`},
		"exadata_infrastructure_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"gi_version":                      acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0.0`},
		"ssh_public_keys":                 acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"vm_cluster_network_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"cloud_automation_update_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseVmClusterCloudAutomationUpdateDetailsRepresentation},
		"data_collection_options":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseVmClusterDataCollectionOptionsRepresentation},
		"data_storage_size_in_tbs":        acctest.Representation{RepType: acctest.Optional, Create: `84`, Update: `86`},
		"db_node_storage_size_in_gbs":     acctest.Representation{RepType: acctest.Optional, Create: `120`, Update: `160`},
		"db_servers":                      acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`}},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"file_system_configuration_details": []acctest.RepresentationGroup{
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation},
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation1},
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation2},
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation3},
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation4},
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation5},
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation6},
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation7},
			{RepType: acctest.Optional, Group: DatabaseVmClusterFileSystemConfigurationDetailsRepresentation8}},

		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_sparse_diskgroup_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":               acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterIgnoreDefinedTagsSystemVersionRepresentation},
		"memory_size_in_gbs":          acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `90`},
		"time_zone":                   acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
	}
	DatabaseVmClusterCloudAutomationUpdateDetailsRepresentation = map[string]interface{}{
		"apply_update_time_preference": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseVmClusterCloudAutomationUpdateDetailsApplyUpdateTimePreferenceRepresentation},
		"freeze_period":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseVmClusterCloudAutomationUpdateDetailsFreezePeriodRepresentation},
		"is_early_adoption_enabled":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_freeze_period_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
	}

	vmClusterIgnoreDefinedTagsSystemVersionRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatabaseVmClusterDataCollectionOptionsRepresentation = map[string]interface{}{
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	DatabaseVmClusterCloudAutomationUpdateDetailsApplyUpdateTimePreferenceRepresentation = map[string]interface{}{
		"apply_update_preferred_end_time":   acctest.Representation{RepType: acctest.Optional, Create: `06:00`, Update: `08:00`},
		"apply_update_preferred_start_time": acctest.Representation{RepType: acctest.Optional, Create: `00:00`, Update: `02:00`},
	}
	DatabaseVmClusterCloudAutomationUpdateDetailsFreezePeriodRepresentation = map[string]interface{}{
		"freeze_period_end_time":   acctest.Representation{RepType: acctest.Optional, Create: `2026-02-15`, Update: `2026-03-15`},
		"freeze_period_start_time": acctest.Representation{RepType: acctest.Optional, Create: `2026-02-13`, Update: `2026-02-13`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `250`, Update: `260`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/u01`, Update: `/u01`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation1 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `15`, Update: `25`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/`, Update: `/`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation2 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `20`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/tmp`, Update: `/tmp`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation3 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `20`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/var`, Update: `/var`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation4 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `40`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/var/log`, Update: `/var/log`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation5 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `4`, Update: `14`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/home`, Update: `/home`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation6 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/var/log/audit`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation7 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `9`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `reserved`},
	}

	DatabaseVmClusterFileSystemConfigurationDetailsRepresentation8 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `16`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `swap`},
	}

	//DatabaseVmClusterResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation) +
	//	DefinedTagsDependencies

	DatabaseVmClusterResourceDependencies = VmClusterNetworkValidatedResourceConfig
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_vm_cluster.test_vm_cluster"
	datasourceName := "data.oci_database_vm_clusters.test_vm_clusters"
	singularDatasourceName := "data.oci_database_vm_cluster.test_vm_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseVmClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Create, DatabaseVmClusterRepresentation), "database", "vmCluster", t)

	acctest.ResourceTest(t, testAccCheckDatabaseVmClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Create, DatabaseVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "06:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "00:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-02-15"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-02-13"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_early_adoption_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "84"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "120"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.file_system_size_gb", "250"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.mount_point", "/u01"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "system_version", "19.2.12.0.0.200317"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseVmClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "06:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "00:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-02-15"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-02-13"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_early_adoption_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "84"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "120"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.file_system_size_gb", "250"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.mount_point", "/u01"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "system_version", "19.2.12.0.0.200317"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Update, DatabaseVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "08:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "02:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-03-15"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-02-13"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_early_adoption_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "6"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.file_system_size_gb", "260"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.mount_point", "/u01"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttr(resourceName, "system_version", "19.2.12.0.0.200317"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_clusters", "test_vm_clusters", acctest.Optional, acctest.Update, DatabaseDatabaseVmClusterDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Update, DatabaseVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "08:00"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "02:00"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-03-15"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-02-13"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.0.is_early_adoption_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.cpus_enabled"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.db_servers.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.file_system_configuration_details.0.file_system_size_gb", "260"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.file_system_configuration_details.0.mount_point", "/u01"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.shape"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.system_version", "19.2.12.0.0.200317"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.vm_cluster_network_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseDatabaseVmClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Update, DatabaseVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "08:00"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "02:00"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-03-15"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-02-13"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.is_early_adoption_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpus_enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_system_configuration_details.0.file_system_size_gb", "260"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_system_configuration_details.0.mount_point", "/u01"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_version", "19.2.12.0.0.200317"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseVmClusterRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cpu_core_count",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetVmClusterRequest{}

			tmp := rs.Primary.ID
			request.VmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetVmCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.VmClusterLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseVmCluster") {
		resource.AddTestSweepers("DatabaseVmCluster", &resource.Sweeper{
			Name:         "DatabaseVmCluster",
			Dependencies: acctest.DependencyGraph["vmCluster"],
			F:            sweepDatabaseVmClusterResource,
		})
	}
}

func sweepDatabaseVmClusterResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	vmClusterIds, err := getDatabaseVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterId := range vmClusterIds {
		if ok := acctest.SweeperDefaultResourceId[vmClusterId]; !ok {
			deleteVmClusterRequest := oci_database.DeleteVmClusterRequest{}

			deleteVmClusterRequest.VmClusterId = &vmClusterId

			deleteVmClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteVmCluster(context.Background(), deleteVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting VmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vmClusterId, DatabaseVmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseVmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseVmClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listVmClustersRequest := oci_database.ListVmClustersRequest{}
	listVmClustersRequest.CompartmentId = &compartmentId
	listVmClustersRequest.LifecycleState = oci_database.VmClusterSummaryLifecycleStateAvailable
	listVmClustersResponse, err := databaseClient.ListVmClusters(context.Background(), listVmClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VmCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vmCluster := range listVmClustersResponse.Items {
		id := *vmCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterId", id)
	}
	return resourceIds, nil
}

func DatabaseVmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vmClusterResponse, ok := response.Response.(oci_database.GetVmClusterResponse); ok {
		return vmClusterResponse.LifecycleState != oci_database.VmClusterLifecycleStateTerminated
	}
	return false
}

func DatabaseVmClusterSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetVmCluster(context.Background(), oci_database.GetVmClusterRequest{
		VmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
