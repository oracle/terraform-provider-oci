// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	adbDedicatedName                   = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbDedicatedUpdateName             = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbDedicatedCloneName              = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adDedicatedName                    = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adDedicatedUpdateName              = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbExaccName                       = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbBackupSourceName                = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbBackupIdName                    = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbBackupTimestampName             = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbPreviewDbName                   = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbDataSafeName                    = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbDbVersionName                   = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbDbRefreshableCloneName          = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbDbRefreshableCloneSourceADBName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbCrossCloneName                  = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	adbCrossCloneNameWithOptionals     = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	AutonomousDatabaseDedicatedRequiredOnlyResource = AutonomousDatabaseDedicatedResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseDedicatedRepresentation)

	AutonomousDatabaseDedicatedResourceConfig = AutonomousDatabaseDedicatedResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseDedicatedRepresentation)

	autonomousDatabaseDedicatedDataSourceRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(DatabaseDatabaseAutonomousDatabaseDataSourceRepresentation, []string{"db_version"}),
		map[string]interface{}{
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: adDedicatedName, Update: adDedicatedUpdateName},
		})

	autonomousDatabaseDedicatedRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbDedicatedName}, DatabaseAutonomousDatabaseRepresentation), []string{"license_model", "whitelisted_ips", "db_version", "is_auto_scaling_enabled", "customer_contacts", "kms_key_id", "vault_id", "autonomous_maintenance_schedule_type", "scheduled_operations", "character_set", "ncharacter_set"}),
		map[string]interface{}{
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     acctest.Representation{RepType: acctest.Optional, Create: `true`},
			"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: adDedicatedName, Update: adDedicatedUpdateName},
			"data_safe_status":                 acctest.Representation{RepType: acctest.Optional, Create: `REGISTERED`, Update: `NOT_REGISTERED`},
			"is_mtls_connection_required":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
			"character_set":                    acctest.Representation{RepType: acctest.Optional, Create: `AR8ADOS710`},
			"ncharacter_set":                   acctest.Representation{RepType: acctest.Optional, Create: `UTF8`},
		})

	autonomousDatabaseDedicatedRepresentationForClone = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbDedicatedCloneName}, autonomousDatabaseDedicatedRepresentation), []string{"license_model", "character_set", "ncharacter_set"}),
		map[string]interface{}{
			"clone_type":     acctest.Representation{RepType: acctest.Optional, Create: `FULL`},
			"display_name":   acctest.Representation{RepType: acctest.Optional, Create: "example_autonomous_database_dedicated_1"},
			"source":         acctest.Representation{RepType: acctest.Optional, Create: `DATABASE`},
			"source_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
			"character_set":  acctest.Representation{RepType: acctest.Optional, Create: `AR8ADOS710`},
			"ncharacter_set": acctest.Representation{RepType: acctest.Optional, Create: `UTF8`},
		})

	autonomousDatabaseDtaSafeStatusRepresentation = map[string]interface{}{
		"admin_password":           acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":           acctest.Representation{RepType: acctest.Required, Create: `1`},
		"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_name":                  acctest.Representation{RepType: acctest.Required, Create: adbDataSafeName},
		"db_workload":              acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`, Update: `displayName2`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_scaling_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_dedicated":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_preview_version_with_service_terms_accepted": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":    acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"data_safe_status": acctest.Representation{RepType: acctest.Optional, Create: `REGISTERED`, Update: `not_REGISTERED`},
		"timeouts":         acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseTimeoutsRepresentation},
	}

	autonomousDatabaseRepresentationForSourceFromBackupId = acctest.RepresentationCopyWithNewProperties(
		acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbBackupIdName}, DatabaseAutonomousDatabaseRepresentation),
		map[string]interface{}{
			"clone_type":                    acctest.Representation{RepType: acctest.Required, Create: `FULL`},
			"source":                        acctest.Representation{RepType: acctest.Required, Create: `BACKUP_FROM_ID`},
			"autonomous_database_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`},
		})

	autonomousDatabaseRepresentationForSourceFromBackupTimestamp = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbBackupTimestampName}, DatabaseAutonomousDatabaseRepresentation), []string{"kms_key_id", "vault_id"}),
		map[string]interface{}{
			"clone_type":             acctest.Representation{RepType: acctest.Required, Create: `FULL`},
			"source":                 acctest.Representation{RepType: acctest.Required, Create: `BACKUP_FROM_TIMESTAMP`},
			"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.autonomous_database_id}`},
			"timestamp":              acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.time_ended}`},
		})

	autonomousDatabaseDataGuardRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"scheduled_operations"}),
		map[string]interface{}{
			"db_version":                           acctest.Representation{RepType: acctest.Optional, Create: `19c`},
			"autonomous_maintenance_schedule_type": acctest.Representation{RepType: acctest.Optional, Create: `REGULAR`},
			"is_mtls_connection_required":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		})

	AutonomousDatabaseDedicatedResourceDependencies = DatabaseAutonomousContainerDatabaseResourceConfig

	autonomousDatabaseRefreshableCloneSourceADBRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"kms_key_id", "vault_id", "scheduled_operations"}), map[string]interface{}{
			"db_name":    acctest.Representation{RepType: acctest.Required, Create: adbDbRefreshableCloneSourceADBName},
			"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19c`},
		})

	autonomousDatabaseRefreshableCloneRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"timeouts", "kms_key_id", "vault_id", "scheduled_operations"}), map[string]interface{}{
			"admin_password":              acctest.Representation{RepType: acctest.Optional, Create: ``},
			"source":                      acctest.Representation{RepType: acctest.Required, Create: `CLONE_TO_REFRESHABLE`},
			"db_name":                     acctest.Representation{RepType: acctest.Required, Create: adbDbRefreshableCloneName},
			"source_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
			"is_refreshable_clone":        acctest.Representation{RepType: acctest.Optional, Create: `true`},
			"refreshable_mode":            acctest.Representation{RepType: acctest.Optional, Create: `MANUAL`},
			"db_version":                  acctest.Representation{RepType: acctest.Optional, Create: `19c`},
			"is_mtls_connection_required": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		})

	autonomousDatabasesCloneDataSourceRepresentation2 = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"clone_type":             acctest.Representation{RepType: acctest.Optional, Create: `REFRESHABLE_CLONE`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	autonomousDatabaseRefreshableClonesDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
	}

	autonomousDatabasePrivateEndpointRepresentation = acctest.RepresentationCopyWithRemovedProperties(
		acctest.RepresentationCopyWithNewProperties(
			DatabaseAutonomousDatabaseRepresentation,
			map[string]interface{}{
				"nsg_ids":                acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group.id}`, `${oci_core_network_security_group.test_network_security_group2.id}`}},
				"private_endpoint_label": acctest.Representation{RepType: acctest.Optional, Create: `xlx4fc9y`},
				"subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
			}), []string{"whitelisted_ips", "scheduled_operations"})

	AutonomousDatabasePrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		DatabaseAutonomousDatabaseResourceDependencies

	AutonomousDatabaseFromBackupDependencies = DatabaseAutonomousDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseBackupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
				"db_name": acctest.Representation{RepType: acctest.Required, Create: adbBackupSourceName},
			}))

	autonomousDatabaseExaccRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbExaccName}, DatabaseAutonomousDatabaseRepresentation), []string{"license_model", "whitelisted_ips", "db_version", "is_auto_scaling_enabled", "operations_insights_status"}),
		map[string]interface{}{
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     acctest.Representation{RepType: acctest.Optional, Create: `true`},
			"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: adbExaccName},
			"is_access_control_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		})
	autonomousDatabaseDGExaccRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbExaccName}, DatabaseAutonomousDatabaseRepresentation), []string{"license_model", "db_version", "is_auto_scaling_enabled", "operations_insights_status", "admin_password", "kms_key_id", "vault_id", "autonomous_maintenance_schedule_type", "customer_contacts", "scheduled_operations"}),
		map[string]interface{}{
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
			"is_dedicated":                     acctest.Representation{RepType: acctest.Optional, Create: `true`},
			"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: adbExaccName},
			"is_access_control_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
			"whitelisted_ips":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`1.1.1.1/28`}, Update: []string{`1.1.1.1/28`, `2.2.2.2/28`}},
			"standby_whitelisted_ips":          acctest.Representation{RepType: acctest.Optional, Update: []string{`3.4.5.6/28`, `3.6.7.8/28`}},
			"are_primary_whitelisted_ips_used": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
			"admin_password":                   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		})
	autonomousDatabaseUpdateExaccRepresentation = map[string]interface{}{
		"admin_password":                   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":                   acctest.Representation{RepType: acctest.Required, Create: `1`},
		"data_storage_size_in_tbs":         acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_name":                          acctest.Representation{RepType: acctest.Required, Create: adbExaccName},
		"db_workload":                      acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: adbExaccName},
		"is_auto_scaling_enabled":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_dedicated":                     acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_access_control_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	autonomousDatabaseExaccRequiredOnlyResource = ExaccADBDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseExaccRepresentation)

	autonomousDatabaseExaccResourceConfig = ExaccADBDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseUpdateExaccRepresentation)

	ExaccADBDatabaseResourceDependencies = ACDatabaseResourceConfig

	ExaccADBWithDataguardResourceDependencies = DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig

	PrimarySourceId string
)

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseDedicated(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseDedicated")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	//ADBD specific ACD MaintenanceWindowRepresentation
	AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation := acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetUpdatedRepresentationCopy("months",
			[]acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4}},
			DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation), []string{"lead_time_in_weeks"})

	AutonomousContainerDatabaseDedicatedRepresentation := acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation}, DatabaseAutonomousContainerDatabaseRepresentation)
	AutonomousContainerDatabaseDedicatedResourceConfig := ATPDAutonomousContainerDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, AutonomousContainerDatabaseDedicatedRepresentation)
	AutonomousDatabaseDedicatedResourceDependencies := AutonomousContainerDatabaseDedicatedResourceConfig
	AutonomousDatabaseDedicatedResourceConfig := AutonomousDatabaseDedicatedResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseDedicatedRepresentation)

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AR8ADOS710"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "UTF8"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "false"),
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
		// verify rotate key
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
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
		// verify no rotation of key
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", acctest.Optional, acctest.Update, autonomousDatabaseDedicatedDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseDedicatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_urls.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.cpu_core_count", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbDedicatedName),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_dedicated", "true"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.provisionable_cpus.#", "108"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseBackupRepresentation, []string{"display_name"})) +
				compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.0.all_connection_strings.%", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_urls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseAutonomousDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"admin_password",
				"clone_type",
				"is_preview_version_with_service_terms_accepted",
				"source",
				"source_id",
				"lifecycle_details",
				"is_auto_scaling_enabled",
				"is_auto_scaling_for_storage_enabled",
				"rotate_key_trigger",
				"is_mtls_connection_required",
				"character_set",
				"is_auto_scaling_for_storage_enabled",
				"ncharacter_set",
			},
			ResourceName: resourceName,
		},

		// remove any previously created resources
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify ADB clone from a source ADB
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseDedicatedRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDedicatedRepresentationForClone),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "clone_type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database_dedicated_1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AR8ADOS710"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "UTF8"),

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

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_preview(t *testing.T) {
	t.Skip("Skip this test as this is a seasonal feature only when Dbaas has a preview to be released.")

	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_preview")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	autonomousDatabasePreviewRepresentation := acctest.GetUpdatedRepresentationCopy("is_preview_version_with_service_terms_accepted", acctest.Representation{RepType: acctest.Optional, Create: `true`},
		acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbPreviewDbName}, DatabaseAutonomousDatabaseRepresentation))
	autonomousDatabasePreviewRepresentationForClone := acctest.GetUpdatedRepresentationCopy("is_preview_version_with_service_terms_accepted", acctest.Representation{RepType: acctest.Optional, Create: `true`}, autonomousDatabaseRepresentationForClone)

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabasePreviewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				// verify computed field db_workload to be defaulted to OLTP
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabasePreviewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabasePreviewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
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

		// verify updates to whitelisted_ips
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{"whitelisted_ips": acctest.Representation{RepType: acctest.Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
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
		// verify autoscaling
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
						"whitelisted_ips":         acctest.Representation{RepType: acctest.Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}},
						"is_auto_scaling_enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", acctest.Optional, acctest.Update, DatabaseDatabaseAutonomousDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
						"whitelisted_ips": acctest.Representation{RepType: acctest.Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.cpu_core_count", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbPreviewDbName),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_dedicated", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.is_preview"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
						"whitelisted_ips": acctest.Representation{RepType: acctest.Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.0.all_connection_strings.%", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_preview"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseAutonomousDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"admin_password",
				"clone_type",
				"is_preview_version_with_service_terms_accepted",
				"source",
				"source_id",
				"lifecycle_details",
				// Need this workaround due to import behavior change introduced by https://github.com/hashicorp/terraform/issues/20985
				"used_data_storage_size_in_tbs",
			},
			ResourceName: resourceName,
		},

		// test ADW db_workload
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("db_workload", acctest.Representation{RepType: acctest.Optional, Create: `DW`}, autonomousDatabasePreviewRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be re-created.")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("db_workload", acctest.Representation{RepType: acctest.Optional, Create: `DW`}, autonomousDatabasePreviewRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
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

		// verify autoscaling with DW workload
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `DW`},
							acctest.Representation{RepType: acctest.Optional, Update: `true`}}, autonomousDatabasePreviewRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
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

		// remove any previously created resources
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		// verify ADB clone from a source ADB
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabasePreviewRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabasePreviewRepresentationForClone),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "clone_type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttr(resourceName, "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_dataSafeStatus(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_dataSafeStatus")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. verify create and register
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDtaSafeStatusRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "REGISTERED"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. Update: deregister data safe only
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("data_safe_status", acctest.Representation{RepType: acctest.Optional, Create: `not_registered`}, autonomousDatabaseDtaSafeStatusRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//2. Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.GetUpdatedRepresentationCopy("data_safe_status", acctest.Representation{RepType: acctest.Optional, Create: `not_registered`}, autonomousDatabaseDtaSafeStatusRepresentation),
						map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//3. Update: all except data safe
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseDtaSafeStatusRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//4. Update: all except compartment (register data safe)
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDtaSafeStatusRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "REGISTERED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//5. verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDtaSafeStatusRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_preview"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_safe_status", "REGISTERED"),
			),
		},
	})
}

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_FromBackupId(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_FromBackupFromId")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_from_backupid"

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. Create dependencies
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
		},
		//1. verify create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backupid", acctest.Required, acctest.Create, autonomousDatabaseRepresentationForSourceFromBackupId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),

				func(s *terraform.State) (err error) {
					resId, err := acctest.FromInstanceState(s, resourceName, "id")
					sourceresId, err := acctest.FromInstanceState(s, "oci_database_autonomous_database.test_autonomous_database", "id")
					if resId == sourceresId {
						return fmt.Errorf("resource not created when it was supposed to be created")
					}
					return err
				},
			),
		},
		//2. delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
		},
		//3. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backupid", acctest.Optional, acctest.Create, autonomousDatabaseRepresentationForSourceFromBackupId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
	})
}

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_FromBackupTimestamp(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_FromBackupTimestamp")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_from_backuptimestamp"

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//1. Create dependencies
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
		},
		//2. verify create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backuptimestamp", acctest.Required, acctest.Create, autonomousDatabaseRepresentationForSourceFromBackupTimestamp),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),

				func(s *terraform.State) (err error) {
					resId, err := acctest.FromInstanceState(s, resourceName, "id")
					sourceresId, err := acctest.FromInstanceState(s, "oci_database_autonomous_database.test_autonomous_database", "id")
					if resId == sourceresId {
						return fmt.Errorf("resource not created when it was supposed to be created")
					}
					return err
				},
			),
		},
		//3. delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
		},
		//4. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backuptimestamp", acctest.Optional, acctest.Create, autonomousDatabaseRepresentationForSourceFromBackupTimestamp),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
	})
}

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_privateEndpoint(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_privateEndPoint")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. verify create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabasePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fc9y"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "private_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "6"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabasePrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fc9y"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "6"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//2. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(autonomousDatabasePrivateEndpointRepresentation, []string{"is_mtls_connection_required", "scheduled_operations"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "6"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//3. verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", acctest.Optional, acctest.Update, DatabaseDatabaseAutonomousDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("nsg_ids", acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}}, autonomousDatabasePrivateEndpointRepresentation)), Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.cpu_core_count", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbName),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_dedicated", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.is_preview"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.nsg_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.private_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.private_endpoint_label", "xlx4fc9y"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),
			),
		},
		//4. verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("nsg_ids", acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}}, autonomousDatabasePrivateEndpointRepresentation)), Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_strings.0.all_connection_strings.%"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_preview"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_label", "xlx4fc9y"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),
			),
		},
		//5. delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies,
		},
		//6. verify Create with no private end point
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19c`},
					}), []string{"whitelisted_ips"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//7. verify turn on PE
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":             acctest.Representation{RepType: acctest.Optional, Create: `19c`},
						"nsg_ids":                acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group.id}`, `${oci_core_network_security_group.test_network_security_group2.id}`}},
						"private_endpoint_label": acctest.Representation{RepType: acctest.Optional, Create: `xlx4fc9y`},
						"subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
					}), []string{"whitelisted_ips"})), Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fc9y"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//8. delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies,
		},
		//9. Create ADB with private access and data safe registered
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabasePrivateEndpointRepresentation, map[string]interface{}{
						"db_version":       acctest.Representation{RepType: acctest.Optional, Create: `19c`},
						"data_safe_status": acctest.Representation{RepType: acctest.Optional, Create: `REGISTERED`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fc9y"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "private_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "6"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//10. Enable mtls connection
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabasePrivateEndpointRepresentation, map[string]interface{}{
						"db_version":                  acctest.Representation{RepType: acctest.Optional, Create: `19c`},
						"is_mtls_connection_required": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fc9y"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//11. change network access to public
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(autonomousDatabasePrivateEndpointRepresentation, []string{"nsg_ids", "private_endpoint_label", "subnet_id"}), map[string]interface{}{
						"nsg_ids":                acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
						"private_endpoint_label": acctest.Representation{RepType: acctest.Optional, Create: `null`},
						"subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `null`},
						"db_version":             acctest.Representation{RepType: acctest.Optional, Create: `19c`, Update: `19c`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "null"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),

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

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_dbVersion(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_dbVersion")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	autonomousDatabaseDbVersionUpdateRepresentation := acctest.GetUpdatedRepresentationCopy("admin_password", acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbDbVersionName},
			acctest.GetUpdatedRepresentationCopy("defined_tags", acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`},
				acctest.GetUpdatedRepresentationCopy("display_name", acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`},
					acctest.GetUpdatedRepresentationCopy("freeform_tags", acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
						acctest.GetUpdatedRepresentationCopy("db_version", acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}`, Update: `19c`},
							acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"is_mtls_connection_required", "scheduled_operations"})))))))

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDbVersionUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbVersionName),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),

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
		// verify Update to only db_version
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseDbVersionUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbVersionName),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify Update of parameters except db_version
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("db_version", acctest.Representation{RepType: acctest.Optional, Update: `19c`},
						acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbDbVersionName}, DatabaseAutonomousDatabaseRepresentation))),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbVersionName),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),

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

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_dataGuard(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_dataGuard")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDataGuardRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "autonomous_maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				//resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "1"),

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
		//1. verify updates to updatable parameters and enable dataGuard
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"is_local_data_guard_enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_local_data_guard_enabled"),
				resource.TestCheckResourceAttr(resourceName, "standby_db.0.state", "STANDBY"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//2. verify updates disable dataGuard
		{
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, ListAutonomousDatabasesWaitCondition, standbyDbWaitConditionDuration,
				listListAutonomousDatabasesFetchOperation, "database", true),
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"is_data_guard_enabled": acctest.Representation{RepType: acctest.Optional, Update: `false`},
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_local_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_remote_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "local_standby_db.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
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
	})
}

// issue-routing-tag: database/ExaCC
func TestResourceDatabaseExaccAutonomousDatabaseResource_dataGuard(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseExaccAutonomousDatabaseResource_dataGuard")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExaccADBWithDataguardResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDGExaccRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbExaccName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adbExaccName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_access_control_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "are_primary_whitelisted_ips_used", "true"),
				resource.TestCheckResourceAttr(resourceName, "standby_whitelisted_ips.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "memory_per_oracle_compute_unit_in_gbs"),

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
		// verify updates to acl parameter for Exacc
		{
			Config: config + compartmentIdVariableStr + ExaccADBWithDataguardResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseDGExaccRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbExaccName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adbExaccName),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_access_control_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "are_primary_whitelisted_ips_used", "false"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "standby_whitelisted_ips.#", "2"),
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
	})
}

func ListAutonomousDatabasesWaitCondition(response oci_common.OCIOperationResponse) bool {
	if listListAutonomousDatabasesResponse, ok := response.Response.(oci_database.ListAutonomousDatabasesResponse); ok {
		if len(listListAutonomousDatabasesResponse.Items) > 0 {
			return listListAutonomousDatabasesResponse.Items[0].StandbyDb.LifecycleState != oci_database.AutonomousDatabaseStandbySummaryLifecycleStateAvailable
		}
		return true
	}
	return false
}

func listListAutonomousDatabasesFetchOperation(client *client.OracleClients, databaseId *string, retryPolicy *oci_common.RetryPolicy) error {
	_, err := client.DatabaseClient().ListAutonomousDatabases(context.Background(), oci_database.ListAutonomousDatabasesRequest{
		AutonomousContainerDatabaseId: databaseId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_switchover(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_switchover")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDataGuardRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),

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
		//1. verify enable dataGuard
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"is_data_guard_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_local_data_guard_enabled"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//2. verify no-op when switchover is PRIMARY for first time
		{
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, ListAutonomousDatabasesWaitCondition, standbyDbWaitConditionDuration,
				listListAutonomousDatabasesFetchOperation, "database", true),
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": acctest.Representation{RepType: acctest.Optional, Update: `PRIMARY`},
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "switchover_to", "PRIMARY"),
				resource.TestCheckResourceAttrSet(resourceName, "time_local_data_guard_enabled"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//3. verify switchover to STANDBY
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": acctest.Representation{RepType: acctest.Optional, Update: `STANDBY`},
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_of_last_switchover"),
				resource.TestCheckResourceAttr(resourceName, "switchover_to", "STANDBY"),
				resource.TestCheckResourceAttrSet(resourceName, "time_local_data_guard_enabled"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//4. verify switchover to PRIMARY
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": acctest.Representation{RepType: acctest.Optional, Update: `PRIMARY`},
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_of_last_switchover"),
				resource.TestCheckResourceAttr(resourceName, "switchover_to", "PRIMARY"),
				resource.TestCheckResourceAttrSet(resourceName, "time_local_data_guard_enabled"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//5. verify datasource
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": acctest.Representation{RepType: acctest.Optional, Update: `PRIMARY`},
					})) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbName),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_of_last_switchover"),
			),
		},
		//6. verify singular datasource
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": acctest.Representation{RepType: acctest.Optional, Update: `PRIMARY`},
					})) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_last_switchover"),
				resource.TestCheckResourceAttrSet(resourceName, "time_local_data_guard_enabled"),
			),
		},
	})
}

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_refreshableClone(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_refreshableClone")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	var resId, resId2 string

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"
	clonesDatasourceName := "data.oci_database_autonomous_databases_clones.test_autonomous_databases_clones"
	refreshableClonesDatasourceName := "data.oci_database_autonomous_database_refreshable_clones.test_autonomous_database_refreshable_clones"

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refreshable_clone", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
				resource.TestCheckResourceAttr(resourceName, "refreshable_mode", "MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "refreshable_status", "REFRESHING"),
				resource.TestCheckResourceAttr(resourceName, "source", "CLONE_TO_REFRESHABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				// time_of_last_refresh_point apply when refreshable_mode both MANUAL and AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh_point"),
				// time_of_last_refresh and time_of_next_refresh returned when refreshable_mode=AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh"),
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_next_refresh"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. verify LIST clones given a source ADB
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases_clones", "test_autonomous_databases_clones", acctest.Optional, acctest.Create, autonomousDatabasesCloneDataSourceRepresentation2) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_refreshable_clones", "test_autonomous_database_refreshable_clones", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableClonesDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(clonesDatasourceName, "clone_type", "REFRESHABLE_CLONE"),
				resource.TestCheckResourceAttr(clonesDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(clonesDatasourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(clonesDatasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.#"),
				resource.TestCheckResourceAttr(clonesDatasourceName, "autonomous_databases.0.available_upgrade_versions.#", "0"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.compartment_id"),
				resource.TestCheckResourceAttr(clonesDatasourceName, "autonomous_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttr(clonesDatasourceName, "autonomous_databases.0.connection_urls.#", "1"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.cpu_core_count"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.data_safe_status"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.db_name"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.db_workload"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.display_name"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.failed_data_recovery_in_seconds"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.is_auto_scaling_enabled"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.is_data_guard_enabled"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.is_dedicated"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.is_free_tier"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.is_preview"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.is_refreshable_clone"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.license_model"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.open_mode"),
				// todo: commented due to a bug in service, to be reverted after they fix it
				//resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.refreshable_mode"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.refreshable_status"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.source_id"),
				//resource.TestCheckResourceAttr(clonesDatasourceName, "autonomous_databases.0.standby_db.#", "0"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_created"),
				//resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_deletion_of_free_autonomous_database"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_maintenance_begin"),
				resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_maintenance_end"),
				// values are not available immediately
				//resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_of_last_failover"),
				//resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_of_last_refresh"),
				//resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_of_last_refresh_point"),
				//resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_of_last_switchover"),
				//resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_of_next_refresh"),
				//resource.TestCheckResourceAttrSet(clonesDatasourceName, "autonomous_databases.0.time_reclamation_of_free_autonomous_database"),
				resource.TestCheckResourceAttr(clonesDatasourceName, "autonomous_databases.0.whitelisted_ips.#", "1"),

				resource.TestCheckResourceAttr(refreshableClonesDatasourceName, "refreshable_clone_collection.#", "1"),
				resource.TestCheckResourceAttr(refreshableClonesDatasourceName, "refreshable_clone_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(refreshableClonesDatasourceName, "refreshable_clone_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(refreshableClonesDatasourceName, "refreshable_clone_collection.0.items.0.region"),
			),
		},
		//2. verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases_clones", "test_autonomous_databases_clones", acctest.Optional, acctest.Create, autonomousDatabasesCloneDataSourceRepresentation2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRefreshableCloneRepresentation, map[string]interface{}{
						"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Create: `1`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refreshable_clone", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
				resource.TestCheckResourceAttr(resourceName, "refreshable_mode", "MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "refreshable_status", "REFRESHING"),
				resource.TestCheckResourceAttr(resourceName, "source", "CLONE_TO_REFRESHABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				// time_of_last_refresh_point apply when refreshable_mode both MANUAL and AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh_point"),
				// time_of_last_refresh and time_of_next_refresh returned when refreshable_mode=AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh"),
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_next_refresh"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//3. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRefreshableCloneRepresentation, map[string]interface{}{
					"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Update: `1`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refreshable_clone", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
				resource.TestCheckResourceAttr(resourceName, "refreshable_mode", "MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "refreshable_status", "REFRESHING"),
				resource.TestCheckResourceAttr(resourceName, "source", "CLONE_TO_REFRESHABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				// time_of_last_refresh_point apply when refreshable_mode both MANUAL and AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh_point"),
				// time_of_last_refresh and time_of_next_refresh returned when refreshable_mode=AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh"),
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_next_refresh"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//4. verify datasource
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseRefreshableCloneRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousDatabaseDataSourceRepresentation, map[string]interface{}{
						"db_version": acctest.Representation{RepType: acctest.Required, Create: `19c`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.cpu_core_count", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_dedicated", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_free_tier", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_preview", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_refreshable_clone", "true"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.open_mode"),
				// todo: commented due to a bug in service, to be reverted after they fix it
				//resource.TestCheckResourceAttr(singularDatasourceName, "refreshable_mode", "MANUAL"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.refreshable_status", "REFRESHING"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.source_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
				// time_of_last_refresh_point apply when refreshable_mode both MANUAL and AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_of_last_refresh_point"),
				// time_of_last_refresh and time_of_next_refresh returned when refreshable_mode=AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_of_last_refresh"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_of_next_refresh"),
			),
		},
		//5. verify singular datasource
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, autonomousDatabaseRefreshableCloneRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_strings.0.all_connection_strings.%"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_free_tier", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_preview", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_refreshable_clone", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "open_mode"),
				// todo: commented due to a bug in service, to be reverted after they fix it
				//resource.TestCheckResourceAttr(singularDatasourceName, "refreshable_mode", "MANUAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "refreshable_status", "REFRESHING"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				// time_of_last_refresh_point apply when refreshable_mode both MANUAL and AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_databases.0.time_of_last_refresh_point"),
				// time_of_last_refresh and time_of_next_refresh returned when refreshable_mode=AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_databases.0.time_of_last_refresh"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_databases.0.time_of_next_refresh"),
			),
		},
		//6. verify detaching a refreshable clone
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRefreshableCloneRepresentation, map[string]interface{}{
						"is_refreshable_clone":     acctest.Representation{RepType: acctest.Optional, Update: `false`},
						"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Update: `1`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refreshable_clone", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
				resource.TestCheckResourceAttr(resourceName, "refreshable_mode", "MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "source", "CLONE_TO_REFRESHABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				// time_of_last_refresh_point apply when refreshable_mode both MANUAL and AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh_point"),
				// time_of_last_refresh and time_of_next_refresh returned when refreshable_mode=AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh"),
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_next_refresh"),
				resource.TestCheckResourceAttrSet(resourceName, "time_until_reconnect_clone_enabled"),
				resource.TestCheckResourceAttr(resourceName, "is_reconnect_clone_enabled", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//7. verify reconnecting a refreshable clone
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", acctest.Optional, acctest.Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRefreshableCloneRepresentation, map[string]interface{}{
					"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Update: `1`},
				})),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_refreshable_clone", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
				resource.TestCheckResourceAttr(resourceName, "refreshable_mode", "MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "source", "CLONE_TO_REFRESHABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				// time_of_last_refresh_point apply when refreshable_mode both MANUAL and AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh_point"),
				// time_of_last_refresh and time_of_next_refresh returned when refreshable_mode=AUTOMATIC, not available immediately
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_last_refresh"),
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_databases.0.time_of_next_refresh"),
				resource.TestCheckResourceAttr(resourceName, "is_reconnect_clone_enabled", "false"),
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

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_AJD(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_AJD")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. verify create with required
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":    acctest.Representation{RepType: acctest.Required, Create: `19c`},
						"db_workload":   acctest.Representation{RepType: acctest.Required, Create: `AJD`},
						"license_model": acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
						"is_free_tier":  acctest.Representation{RepType: acctest.Required, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		//2. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `AJD`},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`}}, DatabaseAutonomousDatabaseRepresentation), []string{"scheduled_operations", "ncharacter_set", "character_set"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//3. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `AJD`},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`},
							acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`},
							acctest.Representation{RepType: acctest.Optional, Create: `false`}}, DatabaseAutonomousDatabaseRepresentation), []string{"scheduled_operations", "ncharacter_set", "character_set"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
		//4. verify autoscaling with AJD workload
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `AJD`},
							acctest.Representation{RepType: acctest.Optional, Update: `true`},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`},
							acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`},
							acctest.Representation{RepType: acctest.Optional, Create: `false`}}, DatabaseAutonomousDatabaseRepresentation), []string{"scheduled_operations", "ncharacter_set", "character_set"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
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
		//5. verify Update db_workload to OLTP
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `AJD`, Update: `OLTP`},
							acctest.Representation{RepType: acctest.Optional, Update: `true`},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`},
							acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`},
							acctest.Representation{RepType: acctest.Optional, Create: `false`}}, DatabaseAutonomousDatabaseRepresentation), []string{"scheduled_operations", "ncharacter_set", "character_set"})),
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
	})
}

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_APEX(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_APEX")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_apex"

	var resId, resId2 string
	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. verify create with required
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":    acctest.Representation{RepType: acctest.Required, Create: `19c`},
						"db_workload":   acctest.Representation{RepType: acctest.Required, Create: `APEX`},
						"license_model": acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
						"is_free_tier":  acctest.Representation{RepType: acctest.Required, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		//2. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `APEX`},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`},
							acctest.Representation{RepType: acctest.Optional, Create: `true`}}, DatabaseAutonomousDatabaseRepresentation), []string{"ncharacter_set"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//3. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `APEX`},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`},
							acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`},
							acctest.Representation{RepType: acctest.Optional, Create: `true`}}, DatabaseAutonomousDatabaseRepresentation), []string{"scheduled_operations", "ncharacter_set", "character_set"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
		//4. verify autoscaling with APEX workload
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `APEX`},
							acctest.Representation{RepType: acctest.Optional, Update: `true`},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`},
							acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`},
							acctest.Representation{RepType: acctest.Optional, Create: `true`}}, DatabaseAutonomousDatabaseRepresentation), []string{"scheduled_operations", "ncharacter_set"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
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
		//5. verify Update db_workload to OLTP
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
							acctest.Representation{RepType: acctest.Optional, Update: `true`},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`},
							acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`},
							acctest.Representation{RepType: acctest.Optional, Create: `true`}}, DatabaseAutonomousDatabaseRepresentation), []string{"scheduled_operations", "ncharacter_set"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
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
	})
}

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_ConfigureKey(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_ConfigureKey")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. verify create with required
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":    acctest.Representation{RepType: acctest.Required, Create: `19c`},
						"license_model": acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
						"is_free_tier":  acctest.Representation{RepType: acctest.Required, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
					acctest.GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
						[]interface{}{acctest.Representation{RepType: acctest.Optional, Create: "OLTP"},
							acctest.Representation{RepType: acctest.Optional, Create: `19c`}}, DatabaseAutonomousDatabaseRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//2. delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		//3. Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDataGuardRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
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
		//4. test config key
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"kms_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
						"vault_id":   acctest.Representation{RepType: acctest.Optional, Create: kmsVaultId},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),

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

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_CrossRegionClone(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_CrossRegionClone")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_cross_region_clone"

	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	if sourceRegion == "" {
		t.Skip("Skipping TestResourceDatabaseAutonomousDatabaseResource_CrossRegionClone test because there is no source region specified")
	}

	var err error
	PrimarySourceId, _, err = createAdbInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to createAdbInRegion with the error %v", err)
	}

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_cross_region_clone", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"clone_type":   acctest.Representation{RepType: acctest.Required, Create: `FULL`},
						"source":       acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
						"db_name":      acctest.Representation{RepType: acctest.Required, Create: adbCrossCloneName},
						"source_id":    acctest.Representation{RepType: acctest.Required, Create: PrimarySourceId},
						"display_name": acctest.Representation{RepType: acctest.Required, Create: `example_cross_region_clone_source`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbCrossCloneName),
				resource.TestCheckResourceAttr(resourceName, "source", "DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "source_id", PrimarySourceId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_cross_region_clone_source"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_cross_region_clone", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentation, map[string]interface{}{
						"clone_type":   acctest.Representation{RepType: acctest.Required, Create: `FULL`},
						"source":       acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
						"db_name":      acctest.Representation{RepType: acctest.Required, Create: adbCrossCloneNameWithOptionals},
						"source_id":    acctest.Representation{RepType: acctest.Required, Create: PrimarySourceId},
						"display_name": acctest.Representation{RepType: acctest.Required, Create: `example_cross_region_clone_source`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
	})
}
