// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	acbDBName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	ACDatabaseResourceConfig = ACDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, ACDatabaseRepresentation)

	ACDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `containerdatabases2`},
		"infrastructure_type":      acctest.Representation{RepType: acctest.Optional, Create: `CLOUD_AT_CUSTOMER`},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousContainerDatabaseDataSourceFilterRepresentation},
	}

	ACDatabaseRepresentation = map[string]interface{}{
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `containerdatabases2`},
		"patch_model":                  acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousContainerDatabaseBackupConfigRepresentation},
		"key_store_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_unique_name":               acctest.Representation{RepType: acctest.Optional, Create: acbDBName},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
	}

	ACDatabaseBackupConfigRepresentation = map[string]interface{}{
		"recovery_window_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `RECOVERY_APPLIANCE`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
		"internet_proxy": acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"vpc_password":   acctest.Representation{RepType: acctest.Optional, Create: `vpcPassword`, Update: `vpcPassword2`},
		"vpc_user":       acctest.Representation{RepType: acctest.Optional, Create: `bkupUser1`},
	}

	acdBackupConfigLocalRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
			"type": acctest.Representation{RepType: acctest.Required, Create: `LOCAL`}}},
		"recovery_window_in_days": acctest.Representation{RepType: acctest.Optional, Create: `7`},
	}

	ACDatabaseResourceDependencies = AutonomousExadataInfrastructureResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, map[string]interface{}{"activation_file": acctest.Representation{RepType: acctest.Required, Create: activationFilePath}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, autonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(vmClusterNetworkRepresentation, map[string]interface{}{"validate_vm_cluster_network": acctest.Representation{RepType: acctest.Required, Create: "true"}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, keyStoreRepresentation) +
		KmsVaultIdVariableStr + OkvSecretVariableStr

	dgDbUniqueName = utils.RandomString(10, utils.CharsetWithoutDigits)
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabase_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabase_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	datasourceName := "data.oci_database_autonomous_container_databases.test_autonomous_container_databases"
	singularDatasourceName := "data.oci_database_autonomous_container_database.test_autonomous_container_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, ACDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_databases", "test_autonomous_container_databases", acctest.Optional, acctest.Create, ACDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies + ExaccAcdResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.#", "1"),
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
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, autonomousContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies + ExaccAcdResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
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
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies + ExaccAcdResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"maintenance_window_details",
				"backup_config.0.backup_destination_details.0.vpc_password",
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabase_rotateDatabase(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabase_rotateDatabase")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	datasourceName := "data.oci_database_autonomous_container_databases.test_autonomous_container_databases"
	singularDatasourceName := "data.oci_database_autonomous_container_database.test_autonomous_container_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals and rotate key
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify no rotation of key
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify rotate key
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_databases", "test_autonomous_container_databases", acctest.Optional, acctest.Update, autonomousContainerDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, autonomousContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.kms_key_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.#", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, autonomousContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
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
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, autonomousContainerDatabaseRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
			},
			ResourceName: resourceName,
		},
	})
}
