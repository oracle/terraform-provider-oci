// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
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
	DatabaseAutonomousDatabaseRequiredOnlyResource = DatabaseAutonomousDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseRepresentation)

	DatabaseAutonomousDatabaseResourceConfig = DatabaseAutonomousDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseRepresentation)

	DatabaseDatabaseAutonomousDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	DatabaseDatabaseAutonomousDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_version":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}`},
		"db_workload":    acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseDataSourceFilterRepresentation}}
	DatabaseAutonomousDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_database.test_autonomous_database.id}`}},
	}

	adbName                          = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	longAdbName1                     = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(14, utils.Charset)
	longAdbName2                     = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(28, utils.Charset)
	adbCloneName                     = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	tfStaticCompartmentId            = utils.GetEnvSettingWithBlankDefault("compartment_id_for_static_resource")
	tfStaticCompartmentIdVariableStr = fmt.Sprintf("variable \"compartment_id_for_static_resource\" { default = \"%s\" }\n", tfStaticCompartmentId)
	adbMemberName                    = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	DatabaseAutonomousDatabaseRepresentation = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":                       acctest.Representation{RepType: acctest.Required, Create: `1`},
		"data_storage_size_in_tbs":             acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_name":                              acctest.Representation{RepType: acctest.Required, Create: adbName},
		"admin_password":                       acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"db_version":                           acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}`},
		"db_workload":                          acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"character_set":                        acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"defined_tags":                         acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"display_name":                         acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`, Update: `displayName2`},
		"freeform_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_scaling_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_auto_scaling_for_storage_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_dedicated":                         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_mtls_connection_required":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"autonomous_maintenance_schedule_type": acctest.Representation{RepType: acctest.Optional, Create: `REGULAR`},
		"is_preview_version_with_service_terms_accepted": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"customer_contacts":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousDatabaseCustomerContactsRepresentation},
		"kms_key_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"license_model":              acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"scheduled_operations":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousDatabaseScheduledOperationsRepresentation},
		"vault_id":                   acctest.Representation{RepType: acctest.Optional, Create: kmsVaultId, Update: kmsVaultId},
		"whitelisted_ips":            acctest.Representation{RepType: acctest.Optional, Create: []string{`1.1.1.1/28`}},
		"operations_insights_status": acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`, Update: `ENABLED`},
		"timeouts":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseTimeoutsRepresentation},
		"ncharacter_set":             acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}
	KmsKeyResourceDependenciesDbaas = KmsVaultIdVariableStr + `
    	data "oci_kms_vault" "test_vault" {
    		#Required
    		vault_id = "${var.kms_vault_id}"
    	}
    	`
	KeyResourceDependencyConfigDbaas = tfStaticCompartmentIdVariableStr + KmsKeyResourceDependenciesDbaas + `
	data "oci_kms_keys" "test_keys_dependency" {
		#Required
		compartment_id = "${var.compartment_id_for_static_resource}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
		algorithm = "AES"

		filter {
    		name = "state"
    		values = ["ENABLED", "UPDATING"]
        }
	}`

	autonomousDatabaseRepresentationBYOL = acctest.GetUpdatedRepresentationCopy("license_model", acctest.Representation{RepType: acctest.Optional, Create: `BRING_YOUR_OWN_LICENSE`}, DatabaseAutonomousDatabaseRepresentation)

	autonomousDatabaseRepresentationMaxCpu = acctest.RepresentationCopyWithNewProperties(
		acctest.GetUpdatedRepresentationCopy("is_auto_scaling_enabled", acctest.Representation{RepType: acctest.Required, Create: `true`}, DatabaseAutonomousDatabaseRepresentation),
		map[string]interface{}{"max_cpu_core_count": acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `3`}})

	DatabaseAutonomousDatabaseCustomerContactsRepresentation = map[string]interface{}{
		"email": acctest.Representation{RepType: acctest.Optional, Create: `test@oracle.com`, Update: `test2@oracle.com`},
	}
	DatabaseAutonomousDatabaseResourcePoolSummaryRepresentation = map[string]interface{}{
		"is_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"pool_size":   acctest.Representation{RepType: acctest.Optional, Create: `128`, Update: `256`},
	}

	DatabaseAutonomousDatabaseScheduledOperationsRepresentation = map[string]interface{}{
		"day_of_week":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseScheduledOperationsDayOfWeekRepresentation},
		"scheduled_start_time": acctest.Representation{RepType: acctest.Optional, Create: `09:00`, Update: `10:00`},
		"scheduled_stop_time":  acctest.Representation{RepType: acctest.Optional, Create: `19:00`, Update: `20:00`},
	}
	DatabaseAutonomousDatabaseScheduledOperationsDayOfWeekRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `MONDAY`, Update: `TUESDAY`},
	}

	autonomousDatabaseTimeoutsRepresentation = map[string]interface{}{
		"create": acctest.Representation{RepType: acctest.Required, Create: `45m`},
		"update": acctest.Representation{RepType: acctest.Required, Create: `45m`},
		"delete": acctest.Representation{RepType: acctest.Required, Create: `45m`},
	}
	autonomousDatabaseCopyWithUpdatedIPsRepresentation = acctest.GetUpdatedRepresentationCopy("whitelisted_ips", acctest.Representation{RepType: acctest.Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}, Update: []string{}}, DatabaseAutonomousDatabaseRepresentation)

	autonomousDatabaseRepresentationForClone = acctest.RepresentationCopyWithNewProperties(
		acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbCloneName}, DatabaseAutonomousDatabaseRepresentation),
		map[string]interface{}{
			"clone_type": acctest.Representation{RepType: acctest.Optional, Create: `FULL`},
			"source":     acctest.Representation{RepType: acctest.Optional, Create: `DATABASE`},
			"source_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		})

	DatabaseAutonomousDatabaseResourceDependencies = DefinedTagsDependencies + KeyResourceDependencyConfigDbaas +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	okvSecret = utils.GetEnvSettingWithBlankDefault("okv_secret")
	OkvSecretVariableStr = fmt.Sprintf("variable \"okv_secret\" { default = \"%s\" }\n", okvSecret)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseRepresentation), "database", "autonomousDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. Verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				// verify computed field db_workload to be defaulted to OLTP
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. Delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		//2. Verify Create with secretId and secretVersionNumber
		{
			Config: config + compartmentIdVariableStr + OkvSecretVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"secret_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.okv_secret}`},
						"secret_version_number": acctest.Representation{RepType: acctest.Required, Create: `1`},
					}), []string{"admin_password"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "secret_id", okvSecret),
				resource.TestCheckResourceAttr(resourceName, "secret_version_number", "1"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				// verify computed field db_workload to be defaulted to OLTP
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//3. Verify Update with secretId and secretVersionNumber
		{
			Config: config + compartmentIdVariableStr + OkvSecretVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"secret_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.okv_secret}`, Update: `${var.okv_secret}`},
						"secret_version_number": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
					}), []string{"admin_password"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "secret_id", okvSecret),
				resource.TestCheckResourceAttr(resourceName, "secret_version_number", "2"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				// verify computed field db_workload to be defaulted to OLTP
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//4. Verify update to ECPU from OCPU
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"backup_retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `30`},
						"compute_model":                   acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
						"compute_count":                   acctest.Representation{RepType: acctest.Required, Create: `4.0`},
						"cpu_core_count":                  acctest.Representation{RepType: acctest.Required, Create: nil},
					}), []string{"admin_password"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceName, "backup_retention_period_in_days", "30"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//5. Delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		//6. Verify Create with optionals and long dbName
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.GetUpdatedRepresentationCopy("license_model", acctest.Representation{RepType: acctest.Optional, Create: `BRING_YOUR_OWN_LICENSE`}, autonomousDatabaseRepresentationMaxCpu), map[string]interface{}{
						"open_mode":                   acctest.Representation{RepType: acctest.Optional, Create: `READ_ONLY`, Update: `READ_ONLY`},
						"permission_level":            acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`, Update: `RESTRICTED`},
						"database_edition":            acctest.Representation{RepType: acctest.Optional, Create: `STANDARD_EDITION`, Update: `STANDARD_EDITION`},
						"db_name":                     acctest.Representation{RepType: acctest.Required, Create: adbName},
						"is_local_data_guard_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
						"local_adg_auto_failover_max_data_loss_limit": acctest.Representation{RepType: acctest.Required, Create: `20`},
					}),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "autonomous_maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_edition", "STANDARD_EDITION"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_for_storage_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "local_adg_auto_failover_max_data_loss_limit", "20"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "max_cpu_core_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.day_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.day_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.scheduled_start_time", "09:00"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.scheduled_stop_time", "19:00"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "open_mode", "READ_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "operations_insights_status", "NOT_ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "permission_level", "RESTRICTED"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//7. Verify DBMS status
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
						"database_management_status": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `ENABLED`},
						"operations_insights_status": acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`, Update: `NOT_ENABLED`},
						"database_edition":           acctest.Representation{RepType: acctest.Optional, Create: `STANDARD_EDITION`, Update: `STANDARD_EDITION`},
						"db_name":                    acctest.Representation{RepType: acctest.Required, Create: adbName},
					}),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "autonomous_maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_edition", "STANDARD_EDITION"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_management_status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//8. Verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
						"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"open_mode":        acctest.Representation{RepType: acctest.Optional, Create: `READ_WRITE`, Update: `READ_WRITE`},
						"permission_level": acctest.Representation{RepType: acctest.Optional, Create: `UNRESTRICTED`, Update: `UNRESTRICTED`},
						"database_edition": acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`, Update: `ENTERPRISE_EDITION`},
						"db_name":          acctest.Representation{RepType: acctest.Required, Create: adbName},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "autonomous_maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_for_storage_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.day_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.day_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.scheduled_start_time", "09:00"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.scheduled_stop_time", "19:00"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "open_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "permission_level", "UNRESTRICTED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//9. Verify rename-database to longer dbName
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
						"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"open_mode":        acctest.Representation{RepType: acctest.Optional, Create: `READ_WRITE`, Update: `READ_WRITE`},
						"permission_level": acctest.Representation{RepType: acctest.Optional, Create: `UNRESTRICTED`, Update: `UNRESTRICTED`},
						"database_edition": acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`, Update: `ENTERPRISE_EDITION`},
						"db_name":          acctest.Representation{RepType: acctest.Required, Create: adbName},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//10. Verify rename-database to smaller dbName
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
						"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"open_mode":        acctest.Representation{RepType: acctest.Optional, Create: `READ_WRITE`, Update: `READ_WRITE`},
						"permission_level": acctest.Representation{RepType: acctest.Optional, Create: `UNRESTRICTED`, Update: `UNRESTRICTED`},
						"database_edition": acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`, Update: `ENTERPRISE_EDITION`},
						"db_name":          acctest.Representation{RepType: acctest.Required, Create: adbName},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// update auto failover data loss limit
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
						"local_adg_auto_failover_max_data_loss_limit": acctest.Representation{RepType: acctest.Required, Create: `125`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "local_adg_auto_failover_max_data_loss_limit", "125"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//disable local adg
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
						"is_local_data_guard_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//11. Verify update to backup retention in days //  Backup Retention Period can only be updated for Autonomous Databases using ECPU compute model
		// 		{
		// 			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
		// 				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
		// 					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
		// 						"backup_retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `20`},
		// 					})),
		// 			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 				resource.TestCheckResourceAttr(resourceName, "backup_retention_period_in_days", "20"),
		//
		// 				func(s *terraform.State) (err error) {
		// 					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
		// 					if resId != resId2 {
		// 						return fmt.Errorf("resource recreated when it was supposed to be updated")
		// 					}
		// 					return err
		// 				},
		// 			),
		// 		},

		//12. Verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
						"database_management_status":  acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`, Update: `NOT_ENABLED`},
						"is_mtls_connection_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
					}), []string{"scheduled_operations"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "autonomous_maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "open_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "operations_insights_status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "database_management_status", "NOT_ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "permission_level", "UNRESTRICTED"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//13. Verify stop the autonomous database
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationBYOL, map[string]interface{}{
						"state":                       acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`},
						"is_mtls_connection_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//14. Verify start the autonomous database
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("state", acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`}, autonomousDatabaseRepresentationBYOL)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//15. Verify updates to whitelisted_ips
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("whitelisted_ips", acctest.Representation{RepType: acctest.Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}}, autonomousDatabaseRepresentationBYOL)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//16. Verify remove whitelisted_ips
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseCopyWithUpdatedIPsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//17. Verify autoscaling
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseCopyWithUpdatedIPsRepresentation, map[string]interface{}{
						"is_auto_scaling_enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
						"max_cpu_core_count":      acctest.Representation{RepType: acctest.Optional, Update: `3`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_for_storage_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_cpu_core_count", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.day_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.day_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.scheduled_start_time", "10:00"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.scheduled_stop_time", "20:00"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//18. Verify autoscaling for storage
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseCopyWithUpdatedIPsRepresentation, map[string]interface{}{
						"is_auto_scaling_for_storage_enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_for_storage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.day_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.day_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.scheduled_start_time", "10:00"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_operations.0.scheduled_stop_time", "20:00"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//19. Verify shrink
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseCopyWithUpdatedIPsRepresentation, map[string]interface{}{
						"is_shrink_only": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//20. Verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", acctest.Optional, acctest.Update, DatabaseDatabaseAutonomousDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseCopyWithUpdatedIPsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.actual_used_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.allocated_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.apex_details.#"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.autonomous_maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.backup_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_urls.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.cpu_core_count", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.data_storage_size_in_gb"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbName),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_for_storage_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_dedicated", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.is_preview"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.is_reconnect_clone_enabled"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.is_refreshable_clone"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.kms_key_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.local_adg_auto_failover_max_data_loss_limit"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.local_disaster_recovery_type"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.local_standby_db.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.max_cpu_core_count", "3"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.local_standby_db.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.open_mode"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.operations_insights_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.database_management_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.permission_level"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.remote_disaster_recovery_configuration.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.scheduled_operations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.scheduled_operations.0.day_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.scheduled_operations.0.day_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.scheduled_operations.0.scheduled_start_time", "10:00"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.scheduled_operations.0.scheduled_stop_time", "20:00"),
				// @Codegen: Can't test private_endpoint with fake resource
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.private_endpoint"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.private_endpoint_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.supported_regions_to_clone_to.#", "5"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_maintenance_begin"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_maintenance_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.used_data_storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "0"),
			),
		},
		//21. Verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_strings.0.all_connection_strings.%"),

				resource.TestCheckResourceAttr(singularDatasourceName, "connection_urls.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.apex_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.machine_learning_user_management_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.sql_dev_web_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.graph_studio_url"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_gb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_for_storage_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_local_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "local_standby_db.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_preview"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "local_disaster_recovery_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "local_standby_db.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_cpu_core_count", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "local_standby_db.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "open_mode"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_management_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "permission_level"),
				resource.TestCheckResourceAttr(singularDatasourceName, "remote_disaster_recovery_configuration.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_operations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_operations.0.day_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_operations.0.day_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_operations.0.scheduled_start_time", "10:00"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_operations.0.scheduled_stop_time", "20:00"),
				// @Codegen: Can't test private_endpointTestResourceDatabaseAutonomousDatabaseResource_preview with fake resource
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "supported_regions_to_clone_to.#", "5"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_maintenance_begin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_maintenance_end"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "0"),
			),
		},
		//22. Verify resource import
		{
			Config:            config + DatabaseAutonomousDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"admin_password",
				"autonomous_database_backup_id",
				"clone_type",
				"is_preview_version_with_service_terms_accepted",
				"secret_id",
				"secret_version_number",
				"source",
				"source_id",
				"lifecycle_details",
				"timestamp",
				// Need this workaround due to import behavior change introduced by https://github.com/hashicorp/terraform/issues/20985
				"used_data_storage_size_in_tbs",
				"is_shrink_only",
				"character_set",
				"ncharacter_set",
				"time_local_data_guard_enabled",
				"local_adg_auto_failover_max_data_loss_limit",
			},
			ResourceName: resourceName,
		},
		//23. Remove required only resource
		{
			Config: config + compartmentIdVariableStr,
		},
		//Test ADW db_workload
		//24. Verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `DW`},
							acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_autonomous_db_versions.test_autonomous_dw_versions.autonomous_db_versions.0.version}`}}, DatabaseAutonomousDatabaseRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be re-created.")
					}
					return err
				},
			),
		},
		//25. Verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `DW`},
							acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_autonomous_db_versions.test_autonomous_dw_versions.autonomous_db_versions.0.version}`},
							acctest.Representation{RepType: acctest.Optional, Create: `false`}}, DatabaseAutonomousDatabaseRepresentation), []string{"scheduled_operations"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//26. Verify autoscaling with DW workload
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `DW`},
							acctest.Representation{RepType: acctest.Optional, Update: `true`},
							acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_autonomous_db_versions.test_autonomous_dw_versions.autonomous_db_versions.0.version}`},
							acctest.Representation{RepType: acctest.Optional, Create: `false`}}, autonomousDatabaseRepresentationMaxCpu), []string{"scheduled_operations"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "max_cpu_core_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
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
		//27. Remove any previously created resources
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		//28. Verify ADB clone from a source ADB
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseRepresentationForClone),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "clone_type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbCloneName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be re-created.")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckDatabaseAutonomousDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_database" {
			noResourceFound = false
			request := oci_database.GetAutonomousDatabaseRequest{}

			tmp := rs.Primary.ID
			request.AutonomousDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetAutonomousDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousDatabaseLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseAutonomousDatabase") {
		resource.AddTestSweepers("DatabaseAutonomousDatabase", &resource.Sweeper{
			Name:         "DatabaseAutonomousDatabase",
			Dependencies: acctest.DependencyGraph["autonomousDatabase"],
			F:            sweepDatabaseAutonomousDatabaseResource,
		})
	}
}

func sweepDatabaseAutonomousDatabaseResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	autonomousDatabaseIds, err := getDatabaseAutonomousDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousDatabaseId := range autonomousDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[autonomousDatabaseId]; !ok {
			deleteAutonomousDatabaseRequest := oci_database.DeleteAutonomousDatabaseRequest{}

			deleteAutonomousDatabaseRequest.AutonomousDatabaseId = &autonomousDatabaseId

			deleteAutonomousDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteAutonomousDatabase(context.Background(), deleteAutonomousDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &autonomousDatabaseId, DatabaseAutonomousDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseAutonomousDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseAutonomousDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AutonomousDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listAutonomousDatabasesRequest := oci_database.ListAutonomousDatabasesRequest{}
	listAutonomousDatabasesRequest.CompartmentId = &compartmentId
	listAutonomousDatabasesResponse, err := databaseClient.ListAutonomousDatabases(context.Background(), listAutonomousDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousDatabase := range listAutonomousDatabasesResponse.Items {
		// if autonomousDatabase is in unavailable state, it also needs to be deleted, otherwise other resources which has dependency on it can not be deleted.
		if autonomousDatabase.LifecycleState == oci_database.AutonomousDatabaseSummaryLifecycleStateAvailable ||
			autonomousDatabase.LifecycleState == oci_database.AutonomousDatabaseSummaryLifecycleStateUnavailable {
			id := *autonomousDatabase.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousDatabaseId", id)
		}
	}
	return resourceIds, nil
}

func DatabaseAutonomousDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDatabaseResponse, ok := response.Response.(oci_database.GetAutonomousDatabaseResponse); ok {
		return autonomousDatabaseResponse.LifecycleState != oci_database.AutonomousDatabaseLifecycleStateTerminated
	}
	return false
}

func DatabaseAutonomousDatabaseSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetAutonomousDatabase(context.Background(), oci_database.GetAutonomousDatabaseRequest{
		AutonomousDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
