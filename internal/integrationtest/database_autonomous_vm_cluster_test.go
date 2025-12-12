// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseAutonomousVmClusterRequiredOnlyResource = DatabaseAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterFinalRepresentation)

	DatabaseAutonomousVmClusterResourceConfig = DatabaseAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Optional, acctest.Update, DatabaseAutonomousVmClusterFinalRepresentation)

	DatabaseDatabaseAutonomousVmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
	}

	DatabaseDatabaseAutonomousVmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `autonomousVmCluster`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousVmClusterDataSourceFilterRepresentation}}
	DatabaseAutonomousVmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`}},
	}

	DatabaseAutonomousVmClusterFinalRepresentation = getExaccAutonomousVmClusterRepresentation()

	DatabaseAutonomousVmClusterRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                          acctest.Representation{RepType: acctest.Required, Create: `autonomousVmCluster`},
		"exadata_infrastructure_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"vm_cluster_network_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"compute_model":                         acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"autonomous_data_storage_size_in_tbs":   acctest.Representation{RepType: acctest.Required, Create: `5.0`, Update: `7.0`},
		"cpu_core_count_per_node":               acctest.Representation{RepType: acctest.Required, Create: `44`, Update: `48`},
		"db_servers":                            acctest.Representation{RepType: acctest.Optional, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`}},
		"defined_tags":                          acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_mtls_enabled":                       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"license_model":                         acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"scan_listener_port_non_tls":            acctest.Representation{RepType: acctest.Optional, Create: `1600`},
		"scan_listener_port_tls":                acctest.Representation{RepType: acctest.Optional, Create: `3600`},
		"maintenance_window_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousVmClusterMaintenanceWindowDetailsRepresentation},
		"memory_per_oracle_compute_unit_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `6`},
		"time_zone":                             acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"total_container_databases":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"lifecycle":                             acctest.RepresentationGroup{RepType: acctest.Required, Group: DbaasIgnoreDefinedTagsRepresentation},
	}

	DatabaseDevAutonomousVmClusterRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `autonomousVmCluster`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"vm_cluster_network_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		//"compute_model":                         acctest.Representation{RepType: acctest.Optional, Create: `OCPU`},
		"autonomous_data_storage_size_in_tbs":   acctest.Representation{RepType: acctest.Required, Create: `5.0`, Update: `7.0`},
		"cpu_core_count_per_node":               acctest.Representation{RepType: acctest.Required, Create: `44`, Update: `48`},
		"db_servers":                            acctest.Representation{RepType: acctest.Optional, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`}},
		"defined_tags":                          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_mtls_enabled":                       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"license_model":                         acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"scan_listener_port_non_tls":            acctest.Representation{RepType: acctest.Optional, Create: `1600`},
		"scan_listener_port_tls":                acctest.Representation{RepType: acctest.Optional, Create: `3600`},
		"maintenance_window_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousVmClusterMaintenanceWindowDetailsRepresentation},
		"memory_per_oracle_compute_unit_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `20`},
		"time_zone":                             acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"total_container_databases":             acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `4`},
	}

	DatabaseECPUAutonomousVmClusterRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseDevAutonomousVmClusterRepresentation, map[string]interface{}{
		"compute_model":                         acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"display_name":                          acctest.Representation{RepType: acctest.Required, Create: `ecpuAutonomousVmCluster`},
		"memory_per_oracle_compute_unit_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `5`},
	})

	DatabaseAutonomousVmClusterMaintenanceWindowDetailsRepresentation = map[string]interface{}{
		"preference":                       acctest.Representation{RepType: acctest.Optional, Create: `CUSTOM_PREFERENCE`, Update: `CUSTOM_PREFERENCE`},
		"days_of_week":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousVmClusterMaintenanceWindowDetailsDaysOfWeekRepresentation},
		"hours_of_day":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`0`}, Update: []string{`4`}},
		"lead_time_in_weeks":               acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"months":                           []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseAutonomousVmClusterMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseAutonomousVmClusterMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseAutonomousVmClusterMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseAutonomousVmClusterMaintenanceWindowDetailsMonthsRepresentation4}},
		"weeks_of_month":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`1`}, Update: []string{`2`}},
		"custom_action_timeout_in_mins":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_custom_action_timeout_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_monthly_patching_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"patching_mode":                    acctest.Representation{RepType: acctest.Optional, Create: `ROLLING`, Update: `NONROLLING`},
	}
	DatabaseAutonomousVmClusterMaintenanceWindowDetailsDaysOfWeekRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `MONDAY`, Update: `TUESDAY`},
	}
	DatabaseAutonomousVmClusterMaintenanceWindowDetailsMonthsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `JANUARY`, Update: `FEBRUARY`},
	}

	DatabaseAutonomousVmClusterMaintenanceWindowDetailsMonthsRepresentation2 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `APRIL`, Update: `MAY`},
	}
	DatabaseAutonomousVmClusterMaintenanceWindowDetailsMonthsRepresentation3 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `JULY`, Update: `AUGUST`},
	}
	DatabaseAutonomousVmClusterMaintenanceWindowDetailsMonthsRepresentation4 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `OCTOBER`, Update: `NOVEMBER`},
	}

	DatabaseAVMClusterWithSingleNetworkResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, map[string]interface{}{"activation_file": acctest.Representation{RepType: acctest.Required, Create: activationFilePath}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatabaseVmClusterNetworkRepresentation, map[string]interface{}{"validate_vm_cluster_network": acctest.Representation{RepType: acctest.Required, Create: "true"}})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, DatabaseDatabaseDbServerDataSourceRepresentation) +
		DefinedTagsDependencies

	DatabaseAutonomousVmClusterResourceDependencies = getExaccAutonomousVmClusterDependency()

	DatabaseAutonomousVmClusterResourceDependency = DatabaseAVMClusterWithSingleNetworkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network2", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(vmClusterNetwork2Representation, map[string]interface{}{"validate_vm_cluster_network": acctest.Representation{RepType: acctest.Required, Create: "true"}}))
)

// issue-routing-tag: database/ExaCC
func TestDatabaseAutonomousVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"
	datasourceName := "data.oci_database_autonomous_vm_clusters.test_autonomous_vm_clusters"
	singularDatasourceName := "data.oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"
	avmResourceUsageDatasourceName := "data.oci_database_autonomous_vm_cluster_resource_usage.test_autonomous_vm_cluster_resource_usage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousVmClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Optional, acctest.Create, DatabaseAutonomousVmClusterFinalRepresentation), "database", "autonomousVmCluster", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousVmClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Optional, acctest.Create, DatabaseAutonomousVmClusterFinalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "autonomousVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.lead_time_in_weeks", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.0.name", "JANUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Optional, acctest.Create, DatabaseAutonomousVmClusterFinalRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_vm_cluster_resource_usage", "test_autonomous_vm_cluster_resource_usage", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterResourceUsageSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "autonomous_vm_cluster_id"),

				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "autonomous_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "available_autonomous_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "available_cpus"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "exadata_storage_in_tbs"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "is_local_backup_enabled"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "memory_per_compute_unit_in_gbs"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "memory_per_oracle_compute_unit_in_gbs"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "non_provisionable_autonomous_container_databases"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "provisionable_autonomous_container_databases"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "provisioned_autonomous_container_databases"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "provisioned_cpus"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "reclaimable_cpus"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "reserved_cpus"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "total_container_databases"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "total_cpus"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "used_autonomous_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(avmResourceUsageDatasourceName, "used_cpus"),
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousVmClusterFinalRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "autonomousVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.lead_time_in_weeks", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.0.name", "JANUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// compartment back to previous
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Optional, acctest.Create, DatabaseAutonomousVmClusterFinalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "autonomousVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify to updatable parameter
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Optional, acctest.Update, DatabaseAutonomousVmClusterFinalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "autonomousVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.lead_time_in_weeks", "2"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_vm_clusters", "test_autonomous_vm_clusters", acctest.Optional, acctest.Update, getExaccAutonomousVmClusterDataSourceRepresentation()) +
				compartmentIdVariableStr + DatabaseAutonomousVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Optional, acctest.Update, DatabaseAutonomousVmClusterFinalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "autonomousVmCluster"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.available_autonomous_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.available_container_databases"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.available_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.available_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.cpus_enabled"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.cpus_lowest_scaled_value"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.display_name", "autonomousVmCluster"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.exadata_storage_in_tbs_lowest_scaled_value"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.max_acds_lowest_scaled_value"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.memory_per_compute_unit_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.reclaimable_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.vm_cluster_network_id"),
				resource.TestCheckResourceAttr(resourceName, "autonomous_vm_clusters.0.system_tags.%", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousVmClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousVmClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_autonomous_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_container_databases"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpus_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpus_lowest_scaled_value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "autonomousVmCluster"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_storage_in_tbs_lowest_scaled_value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_mtls_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_acds_lowest_scaled_value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_per_compute_unit_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reclaimable_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
			),
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + DatabaseAutonomousVmClusterRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"maintenance_window_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseAutonomousVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetAutonomousVmClusterRequest{}

			tmp := rs.Primary.ID
			request.AutonomousVmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetAutonomousVmCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousVmClusterLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseAutonomousVmCluster") {
		resource.AddTestSweepers("DatabaseAutonomousVmCluster", &resource.Sweeper{
			Name:         "DatabaseAutonomousVmCluster",
			Dependencies: acctest.DependencyGraph["autonomousVmCluster"],
			F:            sweepDatabaseAutonomousVmClusterResource,
		})
	}
}

func sweepDatabaseAutonomousVmClusterResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	autonomousVmClusterIds, err := getDatabaseAutonomousVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousVmClusterId := range autonomousVmClusterIds {
		if ok := acctest.SweeperDefaultResourceId[autonomousVmClusterId]; !ok {
			deleteAutonomousVmClusterRequest := oci_database.DeleteAutonomousVmClusterRequest{}

			deleteAutonomousVmClusterRequest.AutonomousVmClusterId = &autonomousVmClusterId

			deleteAutonomousVmClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteAutonomousVmCluster(context.Background(), deleteAutonomousVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousVmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousVmClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &autonomousVmClusterId, DatabaseAutonomousVmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseAutonomousVmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseAutonomousVmClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AutonomousVmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listAutonomousVmClustersRequest := oci_database.ListAutonomousVmClustersRequest{}
	listAutonomousVmClustersRequest.CompartmentId = &compartmentId
	listAutonomousVmClustersRequest.LifecycleState = oci_database.AutonomousVmClusterSummaryLifecycleStateAvailable
	listAutonomousVmClustersResponse, err := databaseClient.ListAutonomousVmClusters(context.Background(), listAutonomousVmClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousVmCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousVmCluster := range listAutonomousVmClustersResponse.Items {
		id := *autonomousVmCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousVmClusterId", id)
	}
	return resourceIds, nil
}

func DatabaseAutonomousVmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousVmClusterResponse, ok := response.Response.(oci_database.GetAutonomousVmClusterResponse); ok {
		return autonomousVmClusterResponse.LifecycleState != oci_database.AutonomousVmClusterLifecycleStateTerminated
	}
	return false
}

func DatabaseAutonomousVmClusterSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetAutonomousVmCluster(context.Background(), oci_database.GetAutonomousVmClusterRequest{
		AutonomousVmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getExaccAutonomousVmClusterDependency() string {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	if simulateDb {
		return DatabaseAutonomousVmClusterResourceDependency
	} else {
		infraId := utils.GetEnvSettingWithBlankDefault("exadata_infrastructure_id")
		vmClusterNetworkId := utils.GetEnvSettingWithBlankDefault("vm_cluster_network_id")
		return fmt.Sprintf("variable \"exadata_infrastructure_id\" { default = \"%s\" }\n", infraId) +
			fmt.Sprintf("variable \"vm_cluster_network_id\" { default = \"%s\" }\n", vmClusterNetworkId)
	}
}

func getExaccAutonomousVmClusterRepresentation() map[string]interface{} {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	if simulateDb {
		return DatabaseAutonomousVmClusterRepresentation
	} else {
		// Add the dynamic properties
		return acctest.RepresentationCopyWithNewProperties(
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousVmClusterRepresentation, []string{"defined_tags", "scan_listener_port_tls", "scan_listener_port_non_tls", "db_servers"}),
			map[string]interface{}{
				"exadata_infrastructure_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.exadata_infrastructure_id}`},
				"vm_cluster_network_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_id}`},
				"memory_per_oracle_compute_unit_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `4`},
			})
	}
}

func getExaccAutonomousVmClusterDataSourceRepresentation() map[string]interface{} {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	if simulateDb {
		return DatabaseDatabaseAutonomousVmClusterDataSourceRepresentation
	} else {
		// Add the dynamic properties
		return acctest.RepresentationCopyWithNewProperties(
			DatabaseDatabaseAutonomousVmClusterDataSourceRepresentation,
			map[string]interface{}{
				"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.exadata_infrastructure_id}`},
			})
	}
}
