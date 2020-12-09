// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v33/common"
	oci_database "github.com/oracle/oci-go-sdk/v33/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousContainerDatabaseRequiredOnlyResource = AutonomousContainerDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Required, Create, autonomousContainerDatabaseRepresentation)

	AutonomousContainerDatabaseResourceConfig = AutonomousContainerDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update, autonomousContainerDatabaseRepresentation)

	autonomousContainerDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	autonomousContainerDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                       Representation{repType: Required, create: `${var.compartment_id}`},
		"autonomous_exadata_infrastructure_id": Representation{repType: Optional, create: `${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`},
		"availability_domain":                  Representation{repType: Optional, create: `${data.oci_identity_availability_domain.ad.name}`},
		"display_name":                         Representation{repType: Optional, create: `containerdatabases2`, update: `displayName2`},
		"state":                                Representation{repType: Optional, create: `AVAILABLE`},
		"filter":                               RepresentationGroup{Required, autonomousContainerDatabaseDataSourceFilterRepresentation}}
	autonomousContainerDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_autonomous_container_database.test_autonomous_container_database.id}`}},
	}

	autonomousContainerDatabaseRepresentation = map[string]interface{}{
		"display_name":                         Representation{repType: Required, create: `containerdatabases2`, update: `displayName2`},
		"patch_model":                          Representation{repType: Required, create: `RELEASE_UPDATES`, update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`},
		"backup_config":                        RepresentationGroup{Optional, ACDatabaseBackupConfigRepresentation},
		"compartment_id":                       Representation{repType: Optional, create: `${var.compartment_id}`},
		"defined_tags":                         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                        Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":                           Representation{repType: Optional, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"maintenance_window_details":           RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type":         Representation{repType: Optional, create: `STANDARD`},
		"vault_id":                             Representation{repType: Optional, create: `${data.oci_kms_vault.test_vault.id}`},
	}
	autonomousContainerDatabaseBackupConfigRepresentation = map[string]interface{}{
		"backup_destination_details": RepresentationGroup{Optional, autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation},
		"recovery_window_in_days":    Representation{repType: Optional, create: `10`, update: `11`},
	}
	autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation = map[string]interface{}{
		"preference": Representation{repType: Required, create: `NO_PREFERENCE`},
	}
	autonomousContainerDatabaseMaintenanceWindowDetailsRepresentation = map[string]interface{}{
		"preference":     Representation{repType: Required, create: `CUSTOM_PREFERENCE`},
		"days_of_week":   RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsDaysOfWeekRepresentation},
		"hours_of_day":   Representation{repType: Optional, create: []string{`4`}, update: []string{`8`}},
		"months":         []RepresentationGroup{{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation}, {Optional, autonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2}, {Optional, autonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3}, {Optional, autonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4}},
		"weeks_of_month": Representation{repType: Optional, create: []string{`1`}, update: []string{`2`}},
	}

	autonomousContainerDatabaseMaintenanceWindowDetailsDaysOfWeekRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `MONDAY`, update: `TUESDAY`},
	}
	autonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `JANUARY`, update: `FEBRUARY`},
	}
	autonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2 = map[string]interface{}{
		"name": Representation{repType: Required, create: `APRIL`, update: `MAY`},
	}

	autonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3 = map[string]interface{}{
		"name": Representation{repType: Required, create: `JULY`, update: `AUGUST`},
	}
	autonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4 = map[string]interface{}{
		"name": Representation{repType: Required, create: `OCTOBER`, update: `NOVEMBER`},
	}

	AutonomousContainerDatabaseResourceDependencies = AutonomousExadataInfrastructureResourceConfig +
		generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Required, Create, autonomousVmClusterRepresentation) +
		KeyResourceDependencyConfig +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create,
			representationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, map[string]interface{}{"activation_file": Representation{repType: Required, create: activationFilePath}})) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create,
			representationCopyWithNewProperties(vmClusterNetworkRepresentation, map[string]interface{}{"validate_vm_cluster_network": Representation{repType: Required, create: "true"}}))
)

func TestDatabaseAutonomousContainerDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	datasourceName := "data.oci_database_autonomous_container_databases.test_autonomous_container_databases"
	singularDatasourceName := "data.oci_database_autonomous_container_database.test_autonomous_container_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousContainerDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Required, Create, autonomousContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
					resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Create,
						getUpdatedRepresentationCopy("maintenance_window_details", RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, autonomousContainerDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
					// all peer related properties are not returned in GET, hence commented check on the below peer related properties
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.id", "id"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.type", "NFS"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.recovery_window_in_days", "10"),
					//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
					//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step) and maintenance_window_details
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Create,
						representationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "MONDAY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "JANUARY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
					// all peer related properties are not returned in GET, hence commented check on the below peer related properties
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.id", "id"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.type", "NFS"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.recovery_window_in_days", "10"),
					//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
					//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update, autonomousContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
					// all peer related properties are not returned in GET, hence commented check on the below peer related properties
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.id", "id"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.type", "NFS"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
					//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.recovery_window_in_days", "10"),
					//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
					//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_databases", "test_autonomous_container_databases", Optional, Update, autonomousContainerDatabaseDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update, autonomousContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.autonomous_exadata_infrastructure_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.recovery_window_in_days", "11"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.db_version"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.kms_key_id"),
					//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.last_maintenance_run_id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.#", "1"),
					//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.patch_id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.#", "4"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.0.name", "FEBRUARY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.weeks_of_month.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.patch_model", "RELEASE_UPDATE_REVISIONS"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.service_level_agreement_type", "STANDARD"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.vault_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Required, Create, autonomousContainerDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.recovery_window_in_days", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "4"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_level_agreement_type", "STANDARD"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},

			{
				Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update,
						getUpdatedRepresentationCopy("maintenance_window_details", RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, autonomousContainerDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
					resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
					resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"rotate_key_trigger",
					"maintenance_window_details",
					"peer_autonomous_container_database_backup_config",
					"peer_autonomous_container_database_compartment_id",
					"peer_autonomous_vm_cluster_id",
					"peer_autonomous_container_database_display_name",
					"peer_autonomous_exadata_infrastructure_id",
					"peer_db_unique_name",
					"protection_mode",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckDatabaseAutonomousContainerDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_container_database" {
			noResourceFound = false
			request := oci_database.GetAutonomousContainerDatabaseRequest{}

			tmp := rs.Primary.ID
			request.AutonomousContainerDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetAutonomousContainerDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminated): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseAutonomousContainerDatabase") {
		resource.AddTestSweepers("DatabaseAutonomousContainerDatabase", &resource.Sweeper{
			Name:         "DatabaseAutonomousContainerDatabase",
			Dependencies: DependencyGraph["autonomousContainerDatabase"],
			F:            sweepDatabaseAutonomousContainerDatabaseResource,
		})
	}
}

func sweepDatabaseAutonomousContainerDatabaseResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	autonomousContainerDatabaseIds, err := getAutonomousContainerDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousContainerDatabaseId := range autonomousContainerDatabaseIds {
		if ok := SweeperDefaultResourceId[autonomousContainerDatabaseId]; !ok {
			terminateAutonomousContainerDatabaseRequest := oci_database.TerminateAutonomousContainerDatabaseRequest{}

			terminateAutonomousContainerDatabaseRequest.AutonomousContainerDatabaseId = &autonomousContainerDatabaseId

			terminateAutonomousContainerDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.TerminateAutonomousContainerDatabase(context.Background(), terminateAutonomousContainerDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousContainerDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousContainerDatabaseId, error)
				continue
			}
			waitTillCondition(testAccProvider, &autonomousContainerDatabaseId, autonomousContainerDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				autonomousContainerDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getAutonomousContainerDatabaseIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "AutonomousContainerDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listAutonomousContainerDatabasesRequest := oci_database.ListAutonomousContainerDatabasesRequest{}
	listAutonomousContainerDatabasesRequest.CompartmentId = &compartmentId
	listAutonomousContainerDatabasesRequest.LifecycleState = oci_database.AutonomousContainerDatabaseSummaryLifecycleStateAvailable
	listAutonomousContainerDatabasesResponse, err := databaseClient.ListAutonomousContainerDatabases(context.Background(), listAutonomousContainerDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousContainerDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousContainerDatabase := range listAutonomousContainerDatabasesResponse.Items {
		id := *autonomousContainerDatabase.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousContainerDatabaseId", id)
	}
	return resourceIds, nil
}

func autonomousContainerDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousContainerDatabaseResponse, ok := response.Response.(oci_database.GetAutonomousContainerDatabaseResponse); ok {
		return autonomousContainerDatabaseResponse.LifecycleState != oci_database.AutonomousContainerDatabaseLifecycleStateTerminated
	}
	return false
}

func autonomousContainerDatabaseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetAutonomousContainerDatabase(context.Background(), oci_database.GetAutonomousContainerDatabaseRequest{
		AutonomousContainerDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
