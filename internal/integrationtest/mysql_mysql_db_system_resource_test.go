// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	mysqlDbSystemSourceRepresentation = map[string]interface{}{
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `BACKUP`},
		"backup_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_backup.test_mysql_backup.id}`},
	}

	mysqlDbSystemSourcePitrRepresentation = map[string]interface{}{
		"source_type":  acctest.Representation{RepType: acctest.Required, Create: `PITR`},
		"db_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_pitr_db_system.id}`},
	}

	mysqlDbSystemSourceImportRepresentation = map[string]interface{}{
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `IMPORTURL`},
		"source_url":  acctest.Representation{RepType: acctest.Optional, Create: `https://objectstorage.us-ashburn-1.oraclecloud.com/p/KQQJ9Sip5EmCfLw7htHwY5ZIQbHmI5-GmHeLojK_vT0OtezjV6s0tRFOBtX2ybXh/n/mysql2/b/bucket_for_canary_import_tests_1-DO_NOT_DELETE/o/@.manifest.json`},
	}

	mysqlHADbSystemRepresentation = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlHAConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemBackupPolicyRepresentation},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `MySQL Database Service`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `DBSystem001`, Update: `displayName2`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":          acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"is_highly_available":     acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"maintenance":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemMaintenanceRepresentation},
		"port":                    acctest.Representation{RepType: acctest.Optional, Create: `3306`},
		"port_x":                  acctest.Representation{RepType: acctest.Optional, Create: `33306`},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForMysqlRep},
	}

	ignoreDefinedTagsChangesForMysqlRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags"}},
	}

	mysqlDbSystemRepresentationStartStandaloneUpdateToHa = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlHAConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlDbSystemBackupPolicyNotUpdateableRepresentation},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `MySQL Database Service`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `DBSystem001`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":          acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"is_highly_available":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"maintenance":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemMaintenanceRepresentation},
		"port":                    acctest.Representation{RepType: acctest.Optional, Create: `3306`},
		"port_x":                  acctest.Representation{RepType: acctest.Optional, Create: `33306`},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForMysqlRepHA},
	}

	ignoreDefinedTagsChangesForMysqlRepHA = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags"}},
	}

	mysqlDbSystemRepresentationFromPARUrl = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"source":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemSourceImportRepresentation},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSourceChangesForMysqlFromPARUrl},
	}

	ignoreSourceChangesForMysqlFromPARUrl = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"source"}},
	}

	mysqlDbSystemRepresentationShapeAndConfigUpdate = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`, Update: `${var.MysqlConfigurationE3_2_32_OCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`, Update: `MySQL.VM.Standard.E3.2.32GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
	}

	MysqlDbSystemSourceBackupResourceDependencies = MysqlMysqlDbSystemResourceDependencies + utils.MysqlHAConfigurationIdVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Required, acctest.Create, MysqlMysqlBackupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_backup_db_system", acctest.Required, acctest.Create, MysqlMysqlDbSystemRepresentation)

	MysqlDbSystemSourcePitrResourceDependencies = MysqlMysqlDbSystemResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_pitr_db_system", acctest.Optional, acctest.Update, MysqlMysqlDbSystemRepresentation)

	MysqlMysqlDbSystemSecureConnectionsWithBYOCRepresentation = map[string]interface{}{
		"certificate_generation_type": acctest.Representation{RepType: acctest.Required, Create: `SYSTEM`, Update: `BYOC`},
		"certificate_id":              acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `${var.certificate_id}`},
	}
)

func TestMysqlMysqlDbSystemResource_sourcePitr(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_sourcePitr")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_sytem_source_pitr"

	// changing admin_password RepType to Optional. As it is an optional parameter for restore operation.
	updatedAdminPasswordRepresentation := acctest.GetUpdatedRepresentationCopy("admin_password", acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		MysqlMysqlDbSystemRepresentation)

	// changing admin_username RepType to Optional. As it is an optional parameter for restore operation.
	updatedAdminUsernameRepresentation := acctest.GetUpdatedRepresentationCopy("admin_username", acctest.Representation{RepType: acctest.Optional, Create: `adminUser`},
		updatedAdminPasswordRepresentation)

	sourcePitrRepresentation := acctest.GetUpdatedRepresentationCopy("ip_address", acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.8`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(updatedAdminUsernameRepresentation, []string{"data_storage_size_in_gb"}), map[string]interface{}{
			"source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemSourcePitrRepresentation},
		}))

	acctest.ResourceTest(t, testAccCheckMysqlMysqlDbSystemDestroy, []resource.TestStep{
		// Verify PointInTimeRecovery

		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourcePitrResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_sytem_source_pitr", acctest.Optional, acctest.Create, sourcePitrRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "01:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "MySQL Database Service"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DBSystem001"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "PITR"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.resource_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.resource_id"),
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_sourceBackup(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_sourceBackup")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	// changing admin_password RepType to Optional. As it is an optional parameter for restore operation.
	updatedAdminPasswordRepresentation := acctest.GetUpdatedRepresentationCopy("admin_password", acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		MysqlMysqlDbSystemRepresentation)

	// changing admin_username RepType to Optional. As it is an optional parameter for restore operation.
	updatedAdminUsernameRepresentation := acctest.GetUpdatedRepresentationCopy("admin_username", acctest.Representation{RepType: acctest.Optional, Create: `adminUser`},
		updatedAdminPasswordRepresentation)

	updatedRepresentation := acctest.GetUpdatedRepresentationCopy("ip_address", acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.8`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(updatedAdminUsernameRepresentation, []string{"data_storage_size_in_gb"}), map[string]interface{}{
			"source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemSourceRepresentation},
		}))

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "01:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "MySQL Database Service"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DBSystem001"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.backup_id"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "BACKUP"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.resource_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.resource_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.backup_id"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "BACKUP"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.resource_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.resource_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify stop
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(updatedRepresentation, map[string]interface{}{
					"state":         acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
					"shutdown_type": acctest.Representation{RepType: acctest.Optional, Create: `FAST`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.backup_id"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "BACKUP"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.resource_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.resource_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify start
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(updatedRepresentation, map[string]interface{}{
					"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.8"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.backup_id"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "BACKUP"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.resource_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.resource_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_sourceImport(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_sourceImport")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system_import"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create DbSystem using import PAR URL
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system_import", acctest.Optional, acctest.Create, mysqlDbSystemRepresentationFromPARUrl),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "IMPORTURL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify stop
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system_import", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(mysqlDbSystemRepresentationFromPARUrl, map[string]interface{}{
					"state":         acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
					"shutdown_type": acctest.Representation{RepType: acctest.Optional, Create: `FAST`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "IMPORTURL"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify start
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system_import", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(mysqlDbSystemRepresentationFromPARUrl, map[string]interface{}{
					"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "IMPORTURL"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_HA(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_HA")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify HA Create
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlHADbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_highly_available", "true"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.resource_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.resource_id"),
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_crashRecovery(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_crashRecovery")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	updatedRepresentation := acctest.GetUpdatedRepresentationCopy("crash_recovery", acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		acctest.RepresentationCopyWithNewProperties(MysqlMysqlDbSystemRepresentation, map[string]interface{}{
			"backup_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlDbSystemBackupPolicyNotUpdateableRepresentation},
		}))

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "01:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "crash_recovery", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "description", "MySQL Database Service"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DBSystem001"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.resource_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.resource_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "01:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "crash_recovery", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.resource_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.resource_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_dataStorageSizeGB(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_dataStorageSizeGB")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	updatedRepresentation := acctest.GetUpdatedRepresentationCopy("data_storage_size_in_gb", acctest.Representation{RepType: acctest.Optional, Create: `50`, Update: `100`},
		acctest.RepresentationCopyWithNewProperties(MysqlMysqlDbSystemRepresentation, map[string]interface{}{
			"backup_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlDbSystemBackupPolicyNotUpdateableRepresentation},
		}))

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with data_storage_size_in_gb
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "50"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify update to data_storage_size_in_gb
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "100"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_HA_enable(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_HA")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify standalone Create
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlDbSystemRepresentationStartStandaloneUpdateToHa),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_highly_available", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify update to HA (enable HA)
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemRepresentationStartStandaloneUpdateToHa),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_highly_available", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_updateConfigurationIdAndShape(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_updateConfigurationIdAndShape")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	// would be nice to make this look up work, but since this is terraform syntax we can't evaluate it.
	// Luckily configs are realm wide and the ocid will be the same across all regions. We should remove the map
	// at some point.
	// mysqlConfigurationOCIDVal := `${var.MysqlConfigurationOCID[var.region]}`
	// change configuration and shape from
	// MySQL.VM.Standard.E3.1.8GB.Standalone
	mysqlConfigurationOCIDVal := "ocid1.mysqlconfiguration.oc1..aaaaaaaalwzc2a22xqm56fwjwfymixnulmbq3v77p5v4lcbb6qhkftxf2trq"
	// to
	// mysqlConfigurationOCIDVal2 := `${var.MysqlConfigurationE3_2_32_OCID[var.region]}`
	// MySQL.VM.Standard.E3.2.32GB.Standalone
	mysqlConfigurationE3_2_32_OCIDVal := "ocid1.mysqlconfiguration.oc1..aaaaaaaakremacvh2fizcznnja5rdxry2q4nyn27afjblyrimzjmrqblhfwa"

	// NOTE: for now this test has to be separate from TestMysqlMysqlDbSystemResource_HA because the configuration and is_highly_available fields
	// cannot be changed at the same time.  Once that is fixed, we should consolidate these two test cases
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify standalone Create
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlDbSystemRepresentationShapeAndConfigUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "configuration_id", mysqlConfigurationOCIDVal),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.VM.Standard.E3.1.8GB"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify update to new configuration and shape
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemRepresentationShapeAndConfigUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "configuration_id", mysqlConfigurationE3_2_32_OCIDVal),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.VM.Standard.E3.2.32GB"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_databaseManagement(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_databaseManagement")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	updatedRepresentation := acctest.GetUpdatedRepresentationCopy("database_management", acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`, Update: `ENABLED`},
		acctest.RepresentationCopyWithNewProperties(MysqlMysqlDbSystemRepresentation, map[string]interface{}{
			"backup_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlDbSystemBackupPolicyNotUpdateableRepresentation},
		}))

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optional fields
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_management", "DISABLED"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_management", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_secureConnections(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_secureConnections")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	certificateId := utils.GetEnvSettingWithBlankDefault("certificate_ocid")
	certificateIdVariableStr := fmt.Sprintf("variable \"certificate_id\" { default = \"%s\" }\n", certificateId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	updatedRepresentation := acctest.GetUpdatedRepresentationCopy("secure_connections", acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemSecureConnectionsWithBYOCRepresentation},
		acctest.RepresentationCopyWithNewProperties(MysqlMysqlDbSystemRepresentation, map[string]interface{}{
			"backup_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlDbSystemBackupPolicyNotUpdateableRepresentation},
		}))

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optional fields
		{
			Config: config + compartmentIdVariableStr + certificateIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "secure_connections.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "secure_connections.0.certificate_generation_type", "SYSTEM"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + certificateIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "secure_connections.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "secure_connections.0.certificate_generation_type", "BYOC"),
				resource.TestCheckResourceAttrSet(resourceName, "secure_connections.0.certificate_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func TestMysqlMysqlDbSystemResource_hostnameLabel(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_hostnameLabel")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	updatedRepresentation := acctest.GetUpdatedRepresentationCopy("hostname_label", acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`, Update: `hostnameLabel2`},
		acctest.RepresentationCopyWithNewProperties(MysqlMysqlDbSystemRepresentation, map[string]interface{}{
			"backup_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlDbSystemBackupPolicyNotUpdateableRepresentation},
		}))

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with hostname_label
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify update to hostname_label
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemSourceBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, updatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}
