// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v49/common"
	oci_database "github.com/oracle/oci-go-sdk/v49/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	adbDedicatedName                   = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbDedicatedUpdateName             = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbDedicatedCloneName              = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adDedicatedName                    = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adDedicatedUpdateName              = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbExaccName                       = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbBackupSourceName                = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbBackupIdName                    = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbBackupTimestampName             = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbPreviewDbName                   = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbDataSafeName                    = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbDbVersionName                   = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbDbRefreshableCloneName          = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)
	adbDbRefreshableCloneSourceADBName = RandomString(1, CharsetWithoutDigits) + RandomString(13, Charset)

	AutonomousDatabaseDedicatedRequiredOnlyResource = AutonomousDatabaseDedicatedResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseDedicatedRepresentation)

	AutonomousDatabaseDedicatedResourceConfig = AutonomousDatabaseDedicatedResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDedicatedRepresentation)

	autonomousDatabaseDedicatedDataSourceRepresentation = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(autonomousDatabaseDataSourceRepresentation, []string{"db_version"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{RepType: Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"display_name":                     Representation{RepType: Optional, Create: adDedicatedName, Update: adDedicatedUpdateName},
		})

	autonomousDatabaseDedicatedRepresentation = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbDedicatedName}, autonomousDatabaseRepresentation), []string{"license_model", "whitelisted_ips", "db_version", "is_auto_scaling_enabled", "customer_contacts", "kms_key_id", "vault_id", "autonomous_maintenance_schedule_type"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{RepType: Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     Representation{RepType: Optional, Create: `true`},
			"display_name":                     Representation{RepType: Optional, Create: adDedicatedName, Update: adDedicatedUpdateName},
			"data_safe_status":                 Representation{RepType: Optional, Create: `REGISTERED`, Update: `NOT_REGISTERED`},
		})

	autonomousDatabaseDedicatedRepresentationForClone = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbDedicatedCloneName}, autonomousDatabaseDedicatedRepresentation), []string{"license_model"}),
		map[string]interface{}{
			"clone_type":   Representation{RepType: Optional, Create: `FULL`},
			"display_name": Representation{RepType: Optional, Create: "example_autonomous_database_dedicated"},
			"source":       Representation{RepType: Optional, Create: `DATABASE`},
			"source_id":    Representation{RepType: Optional, Create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		})

	autonomousDatabaseDtaSafeStatusRepresentation = map[string]interface{}{
		"admin_password":           Representation{RepType: Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"compartment_id":           Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cpu_core_count":           Representation{RepType: Required, Create: `1`},
		"data_storage_size_in_tbs": Representation{RepType: Required, Create: `1`},
		"db_name":                  Representation{RepType: Required, Create: adbDataSafeName},
		"db_workload":              Representation{RepType: Optional, Create: `OLTP`},
		"defined_tags":             Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":             Representation{RepType: Optional, Create: `example_autonomous_database`, Update: `displayName2`},
		"freeform_tags":            Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_scaling_enabled":  Representation{RepType: Optional, Create: `false`},
		"is_dedicated":             Representation{RepType: Optional, Create: `false`},
		"is_preview_version_with_service_terms_accepted": Representation{RepType: Optional, Create: `false`},
		"license_model":    Representation{RepType: Optional, Create: `LICENSE_INCLUDED`},
		"data_safe_status": Representation{RepType: Optional, Create: `REGISTERED`, Update: `not_REGISTERED`},
		"timeouts":         RepresentationGroup{Required, autonomousDatabaseTimeoutsRepresentation},
	}

	autonomousDatabaseRepresentationForSourceFromBackupId = RepresentationCopyWithNewProperties(
		GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbBackupIdName}, autonomousDatabaseRepresentation),
		map[string]interface{}{
			"clone_type":                    Representation{RepType: Required, Create: `FULL`},
			"source":                        Representation{RepType: Required, Create: `BACKUP_FROM_ID`},
			"autonomous_database_backup_id": Representation{RepType: Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`},
		})

	autonomousDatabaseRepresentationForSourceFromBackupTimestamp = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbBackupTimestampName}, autonomousDatabaseRepresentation), []string{"kms_key_id", "vault_id"}),
		map[string]interface{}{
			"clone_type":             Representation{RepType: Required, Create: `FULL`},
			"source":                 Representation{RepType: Required, Create: `BACKUP_FROM_TIMESTAMP`},
			"autonomous_database_id": Representation{RepType: Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.autonomous_database_id}`},
			"timestamp":              Representation{RepType: Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.time_ended}`},
		})

	autonomousDatabaseDataGuardRepresentation = RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
		"db_version":                           Representation{RepType: Optional, Create: `19c`},
		"autonomous_maintenance_schedule_type": Representation{RepType: Optional, Create: `REGULAR`},
		"is_mtls_connection_required":          Representation{RepType: Optional, Create: `false`},
	})

	AutonomousDatabaseDedicatedResourceDependencies = AutonomousContainerDatabaseResourceConfig

	autonomousDatabaseRefreshableCloneSourceADBRepresentation = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(autonomousDatabaseRepresentation, []string{"kms_key_id", "vault_id"}), map[string]interface{}{
			"db_name":    Representation{RepType: Required, Create: adbDbRefreshableCloneSourceADBName},
			"db_version": Representation{RepType: Optional, Create: `19c`},
		})

	autonomousDatabaseRefreshableCloneRepresentation = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(autonomousDatabaseRepresentation, []string{"timeouts", "kms_key_id", "vault_id"}), map[string]interface{}{
			"admin_password":              Representation{RepType: Optional, Create: ``},
			"source":                      Representation{RepType: Required, Create: `CLONE_TO_REFRESHABLE`},
			"db_name":                     Representation{RepType: Required, Create: adbDbRefreshableCloneName},
			"source_id":                   Representation{RepType: Optional, Create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
			"is_refreshable_clone":        Representation{RepType: Optional, Create: `true`},
			"refreshable_mode":            Representation{RepType: Optional, Create: `MANUAL`},
			"db_version":                  Representation{RepType: Optional, Create: `19c`},
			"is_mtls_connection_required": Representation{RepType: Optional, Create: `false`},
		})

	autonomousDatabasesCloneDataSourceRepresentation2 = map[string]interface{}{
		"autonomous_database_id": Representation{RepType: Required, Create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"clone_type":             Representation{RepType: Optional, Create: `REFRESHABLE_CLONE`},
		"display_name":           Representation{RepType: Optional, Create: `example_autonomous_database`},
		"state":                  Representation{RepType: Optional, Create: `AVAILABLE`},
	}

	autonomousDatabasePrivateEndpointRepresentation = RepresentationCopyWithRemovedProperties(
		RepresentationCopyWithNewProperties(
			autonomousDatabaseRepresentation,
			map[string]interface{}{
				"nsg_ids":                Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group.id}`, `${oci_core_network_security_group.test_network_security_group2.id}`}},
				"private_endpoint_label": Representation{RepType: Optional, Create: `xlx4fcli`},
				"subnet_id":              Representation{RepType: Optional, Create: `${oci_core_subnet.test_subnet.id}`},
			}), []string{"whitelisted_ips"})

	AutonomousDatabasePrivateEndpointResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, SubnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", Required, Create, networkSecurityGroupRepresentation) +
		AutonomousDatabaseResourceDependencies

	AutonomousDatabaseFromBackupDependencies = AutonomousDatabaseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Required, Create, autonomousDatabaseBackupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
			RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
				"db_name": Representation{RepType: Required, Create: adbBackupSourceName},
			}))

	autonomousDatabaseExaccRepresentation = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbExaccName}, autonomousDatabaseRepresentation), []string{"license_model", "whitelisted_ips", "db_version", "is_auto_scaling_enabled", "operations_insights_status"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{RepType: Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     Representation{RepType: Optional, Create: `true`},
			"display_name":                     Representation{RepType: Optional, Create: adbExaccName},
			"is_access_control_enabled":        Representation{RepType: Optional, Create: `false`, Update: `true`},
		})
	autonomousDatabaseDGExaccRepresentation = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbExaccName}, autonomousDatabaseRepresentation), []string{"license_model", "db_version", "is_auto_scaling_enabled", "operations_insights_status", "admin_password", "kms_key_id", "vault_id", "autonomous_maintenance_schedule_type", "customer_contacts"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{RepType: Optional, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
			"is_dedicated":                     Representation{RepType: Optional, Create: `true`},
			"display_name":                     Representation{RepType: Optional, Create: adbExaccName},
			"is_access_control_enabled":        Representation{RepType: Optional, Create: `true`, Update: `true`},
			"whitelisted_ips":                  Representation{RepType: Optional, Create: []string{`1.1.1.1/28`}, Update: []string{`1.1.1.1/28`, `2.2.2.2/28`}},
			"standby_whitelisted_ips":          Representation{RepType: Optional, Update: []string{`3.4.5.6/28`, `3.6.7.8/28`}},
			"are_primary_whitelisted_ips_used": Representation{RepType: Optional, Create: `true`, Update: `false`},
			"admin_password":                   Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		})
	autonomousDatabaseUpdateExaccRepresentation = map[string]interface{}{
		"admin_password":                   Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"autonomous_container_database_id": Representation{RepType: Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cpu_core_count":                   Representation{RepType: Required, Create: `1`},
		"data_storage_size_in_tbs":         Representation{RepType: Required, Create: `1`},
		"db_name":                          Representation{RepType: Required, Create: adbExaccName},
		"db_workload":                      Representation{RepType: Optional, Create: `OLTP`},
		"display_name":                     Representation{RepType: Optional, Create: adbExaccName},
		"is_auto_scaling_enabled":          Representation{RepType: Optional, Create: `false`},
		"is_dedicated":                     Representation{RepType: Optional, Create: `true`},
		"is_access_control_enabled":        Representation{RepType: Optional, Create: `false`, Update: `true`},
	}

	autonomousDatabaseExaccRequiredOnlyResource = ExaccADBDatabaseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseExaccRepresentation)

	autonomousDatabaseExaccResourceConfig = ExaccADBDatabaseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseUpdateExaccRepresentation)

	ExaccADBDatabaseResourceDependencies = ACDatabaseResourceConfig

	ExaccADBWithDataguardResourceDependencies = ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig
)

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseDedicated(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseDedicated")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": Representation{RepType: Optional, Create: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": Representation{RepType: Optional, Create: `false`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": Representation{RepType: Optional, Create: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": Representation{RepType: Optional, Create: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify updates to dbName parameter, should cause force new
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{"db_name": Representation{RepType: Optional, Update: adbDedicatedUpdateName}})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedUpdateName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Optional, Update, autonomousDatabaseDedicatedDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDedicatedRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_dedicated", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", adDedicatedUpdateName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceConfig,
		},

		// verify Create with optionals for Exacc
		{
			Config: config + compartmentIdVariableStr + ExaccADBDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseExaccRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbExaccName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", adbExaccName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_access_control_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to acl parameter for Exacc
		{
			Config: config + compartmentIdVariableStr + ExaccADBDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseUpdateExaccRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// remove any previously created resources
		{
			Config: config + compartmentIdVariableStr + autonomousDatabaseExaccResourceConfig,
		},

		// verify resource import
		{
			Config:            config,
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
				"rotate_key_trigger",
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
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseDedicatedRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDedicatedRepresentationForClone),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "clone_type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database_dedicated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	autonomousDatabasePreviewRepresentation := GetUpdatedRepresentationCopy("is_preview_version_with_service_terms_accepted", Representation{RepType: Optional, Create: `true`},
		GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbPreviewDbName}, autonomousDatabaseRepresentation))
	autonomousDatabasePreviewRepresentationForClone := GetUpdatedRepresentationCopy("is_preview_version_with_service_terms_accepted", Representation{RepType: Optional, Create: `true`}, autonomousDatabaseRepresentationForClone)

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabasePreviewRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				// verify computed field db_workload to be defaulted to OLTP
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabasePreviewRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabasePreviewRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify updates to whitelisted_ips
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{"whitelisted_ips": Representation{RepType: Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}}})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "2"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify autoscaling
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
						"whitelisted_ips":         Representation{RepType: Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}},
						"is_auto_scaling_enabled": Representation{RepType: Optional, Update: `true`}})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Optional, Update, autonomousDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
						"whitelisted_ips": Representation{RepType: Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.defined_tags.%", "1"),
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
						"whitelisted_ips": Representation{RepType: Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
						"whitelisted_ips": Representation{RepType: Optional, Create: []string{"1.1.1.1/28", "1.1.1.29"}},
					})),
		},
		// verify resource import
		{
			Config:            config,
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
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					GetUpdatedRepresentationCopy("db_workload", Representation{RepType: Optional, Create: "DW"}, autonomousDatabasePreviewRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be re-created.")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					GetUpdatedRepresentationCopy("db_workload", Representation{RepType: Optional, Create: "DW"}, autonomousDatabasePreviewRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify autoscaling with DW workload
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled"},
						[]interface{}{Representation{RepType: Optional, Create: "DW"},
							Representation{RepType: Optional, Update: `true`}}, autonomousDatabasePreviewRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// remove any previously created resources
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
		},
		// verify ADB clone from a source ADB
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabasePreviewRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabasePreviewRepresentationForClone),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "clone_type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "true"),
				resource.TestCheckResourceAttr(resourceName, "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create and register
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDtaSafeStatusRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "REGISTERED"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Update: deregister data safe only
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					GetUpdatedRepresentationCopy("data_safe_status", Representation{RepType: Optional, Create: `not_registered`}, autonomousDatabaseDtaSafeStatusRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					RepresentationCopyWithNewProperties(GetUpdatedRepresentationCopy("data_safe_status", Representation{RepType: Optional, Create: `not_registered`}, autonomousDatabaseDtaSafeStatusRepresentation),
						map[string]interface{}{
							"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
						})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Update: all except data safe
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDtaSafeStatusRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Update: all except compartment (register data safe)
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDtaSafeStatusRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "REGISTERED"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDtaSafeStatusRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_from_backupid"

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Source ADB should be at least 2hrs old. Time Sleep for 2hrs")
				time.Sleep(2 * time.Hour)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backupid", Required, Create, autonomousDatabaseRepresentationForSourceFromBackupId),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),

				func(s *terraform.State) (err error) {
					resId, err := FromInstanceState(s, resourceName, "id")
					sourceresId, err := FromInstanceState(s, "oci_database_autonomous_database.test_autonomous_database", "id")
					if resId == sourceresId {
						return fmt.Errorf("resource not created when it was supposed to be created")
					}
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backupid", Optional, Create, autonomousDatabaseRepresentationForSourceFromBackupId),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_from_backuptimestamp"

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// Create dependencies, To Create clone the source db must be atleast 2 hrs old
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Source ADB should be at least 2hrs old. Time Sleep for 2hrs")
				time.Sleep(2 * time.Hour)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backuptimestamp", Required, Create, autonomousDatabaseRepresentationForSourceFromBackupTimestamp),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),

				func(s *terraform.State) (err error) {
					resId, err := FromInstanceState(s, resourceName, "id")
					sourceresId, err := FromInstanceState(s, "oci_database_autonomous_database.test_autonomous_database", "id")
					if resId == sourceresId {
						return fmt.Errorf("resource not created when it was supposed to be created")
					}
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backuptimestamp", Optional, Create, autonomousDatabaseRepresentationForSourceFromBackupTimestamp),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabasePrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fcli"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "private_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "6"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousDatabasePrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fcli"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "6"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithRemovedProperties(autonomousDatabasePrivateEndpointRepresentation, []string{"is_mtls_connection_required"})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
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
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "6"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Optional, Update, autonomousDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					GetUpdatedRepresentationCopy("nsg_ids", Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}}, autonomousDatabasePrivateEndpointRepresentation)), Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_dedicated", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.is_preview"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.nsg_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.private_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.private_endpoint_label", "xlx4fcli"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					GetUpdatedRepresentationCopy("nsg_ids", Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}}, autonomousDatabasePrivateEndpointRepresentation)), Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_preview"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_label", "xlx4fcli"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies,
		},
		// verify Create with no private end point
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithRemovedProperties(RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version": Representation{RepType: Optional, Create: `19c`},
					}), []string{"whitelisted_ips"})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
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
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify turn on PE
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithRemovedProperties(RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":             Representation{RepType: Optional, Create: `19c`},
						"nsg_ids":                Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group.id}`, `${oci_core_network_security_group.test_network_security_group2.id}`}},
						"private_endpoint_label": Representation{RepType: Optional, Create: `xlx4fcli`},
						"subnet_id":              Representation{RepType: Optional, Create: `${oci_core_subnet.test_subnet.id}`},
					}), []string{"whitelisted_ips"})), Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
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
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fcli"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies,
		},
		//Create ADB with private access and data safe registered
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousDatabasePrivateEndpointRepresentation, map[string]interface{}{
						"db_version":       Representation{RepType: Optional, Create: `19c`},
						"data_safe_status": Representation{RepType: Optional, Create: `REGISTERED`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fcli"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "private_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "6"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//Enable mtls connection
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousDatabasePrivateEndpointRepresentation, map[string]interface{}{
						"db_version":                  Representation{RepType: Optional, Create: `19c`},
						"is_mtls_connection_required": Representation{RepType: Optional, Create: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "xlx4fcli"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//change network access to public
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(autonomousDatabasePrivateEndpointRepresentation, []string{"nsg_ids", "private_endpoint_label", "subnet_id"}), map[string]interface{}{
						"nsg_ids":                Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
						"private_endpoint_label": Representation{RepType: Optional, Create: `null`},
						"subnet_id":              Representation{RepType: Optional, Create: `null`},
						"db_version":             Representation{RepType: Optional, Create: `19c`, Update: `19c`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
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
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_label", "null"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_mtls_connection_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.0.profiles.#", "3"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	autonomousDatabaseDbVersionUpdateRepresentation := GetUpdatedRepresentationCopy("admin_password", Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbDbVersionName},
			GetUpdatedRepresentationCopy("defined_tags", Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
				GetUpdatedRepresentationCopy("display_name", Representation{RepType: Optional, Create: `example_autonomous_database`},
					GetUpdatedRepresentationCopy("freeform_tags", Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}},
						GetUpdatedRepresentationCopy("db_version", Representation{RepType: Optional, Create: "${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}", Update: `19c`},
							RepresentationCopyWithRemovedProperties(autonomousDatabaseRepresentation, []string{"is_mtls_connection_required"})))))))

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDbVersionUpdateRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbVersionName),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update to only db_version
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDbVersionUpdateRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbVersionName),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify Update of parameters except db_version
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					GetUpdatedRepresentationCopy("db_version", Representation{RepType: Optional, Update: `19c`},
						GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbDbVersionName}, autonomousDatabaseRepresentation))),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbVersionName),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDataGuardRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters and enable dataGuard
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"is_data_guard_enabled": Representation{RepType: Optional, Update: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify updates disable dataGuard
		{
			PreConfig: WaitTillCondition(TestAccProvider, &resId, ListAutonomousDatabasesWaitCondition, standbyDbWaitConditionDuration,
				listListAutonomousDatabasesFetchOperation, "database", true),
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"is_data_guard_enabled": Representation{RepType: Optional, Update: `false`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExaccADBWithDataguardResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDGExaccRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbExaccName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDGExaccRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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

func listListAutonomousDatabasesFetchOperation(client *OracleClients, databaseId *string, retryPolicy *oci_common.RetryPolicy) error {
	_, err := client.databaseClient().ListAutonomousDatabases(context.Background(), oci_database.ListAutonomousDatabasesRequest{
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDataGuardRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify enable dataGuard
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"is_data_guard_enabled": Representation{RepType: Optional, Create: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify no-op when switchover is PRIMARY for first time
		{
			PreConfig: WaitTillCondition(TestAccProvider, &resId, ListAutonomousDatabasesWaitCondition, standbyDbWaitConditionDuration,
				listListAutonomousDatabasesFetchOperation, "database", true),
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": Representation{RepType: Optional, Update: `PRIMARY`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "switchover_to", "PRIMARY"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify switchover to STANDBY
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": Representation{RepType: Optional, Update: `STANDBY`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_of_last_switchover"),
				resource.TestCheckResourceAttr(resourceName, "switchover_to", "STANDBY"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify switchover to PRIMARY
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": Representation{RepType: Optional, Update: `PRIMARY`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(resourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_of_last_switchover"),
				resource.TestCheckResourceAttr(resourceName, "switchover_to", "PRIMARY"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": Representation{RepType: Optional, Update: `PRIMARY`},
					})) +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Required, Create, autonomousDatabaseDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbName),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_of_last_switchover"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"switchover_to": Representation{RepType: Optional, Update: `PRIMARY`},
					})) +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_data_guard_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_last_switchover"),
			),
		},
	})
}

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_refreshableClone(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_refreshableClone")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	var resId, resId2 string

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"
	clonesDatasourceName := "data.oci_database_autonomous_databases_clones.test_autonomous_databases_clones"

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseRefreshableCloneRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify LIST clones given a source ADB
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseRefreshableCloneRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases_clones", "test_autonomous_databases_clones", Optional, Create, autonomousDatabasesCloneDataSourceRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(clonesDatasourceName, "autonomous_databases.0.standby_db.#", "0"),
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
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousDatabaseRefreshableCloneRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseRefreshableCloneRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseRefreshableCloneRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataSourceRepresentation, map[string]interface{}{
						"db_version": Representation{RepType: Required, Create: `19c`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.cpu_core_count", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbDbRefreshableCloneName),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.defined_tags.%", "1"),
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
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseRefreshableCloneRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		// verify detaching a refreshable clone
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseRefreshableCloneRepresentation, map[string]interface{}{
						"is_refreshable_clone": Representation{RepType: Optional, Update: `false`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with required
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
					RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":    Representation{RepType: Required, Create: `19c`},
						"db_workload":   Representation{RepType: Required, Create: `AJD`},
						"license_model": Representation{RepType: Required, Create: `LICENSE_INCLUDED`},
						"is_free_tier":  Representation{RepType: Required, Create: `false`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
						[]interface{}{Representation{RepType: Optional, Create: "AJD"},
							Representation{RepType: Optional, Create: `19c`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{Representation{RepType: Optional, Create: "AJD"},
							Representation{RepType: Optional, Create: `19c`},
							Representation{RepType: Optional, Create: `NOT_ENABLED`},
							Representation{RepType: Optional, Create: `false`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify autoscaling with AJD workload
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{Representation{RepType: Optional, Create: "AJD"},
							Representation{RepType: Optional, Update: `true`},
							Representation{RepType: Optional, Create: `19c`},
							Representation{RepType: Optional, Create: `NOT_ENABLED`},
							Representation{RepType: Optional, Create: `false`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify Update db_workload to OLTP
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{Representation{RepType: Optional, Create: "OLTP"},
							Representation{RepType: Optional, Update: `true`},
							Representation{RepType: Optional, Create: `19c`},
							Representation{RepType: Optional, Create: `NOT_ENABLED`},
							Representation{RepType: Optional, Create: `false`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
		},
		// verify Create OLTP with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":                           Representation{RepType: Required, Create: `19c`},
						"is_free_tier":                         Representation{RepType: Required, Create: `true`},
						"autonomous_maintenance_schedule_type": Representation{RepType: Optional, Create: `REGULAR`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "true"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithRemovedProperties(RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":                           Representation{RepType: Required, Create: `19c`},
						"is_free_tier":                         Representation{RepType: Required, Create: `true`},
						"autonomous_maintenance_schedule_type": Representation{RepType: Optional, Create: `REGULAR`},
						"is_mtls_connection_required":          Representation{RepType: Required, Create: `false`},
					}), []string{"operations_insights_status"})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify OLTP updated to AJD
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithRemovedProperties(RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":                           Representation{RepType: Required, Create: `19c`},
						"db_workload":                          Representation{RepType: Required, Create: `AJD`},
						"is_free_tier":                         Representation{RepType: Required, Create: `false`},
						"autonomous_maintenance_schedule_type": Representation{RepType: Optional, Create: `REGULAR`},
						"is_mtls_connection_required":          Representation{RepType: Required, Create: `false`},
					}), []string{"operations_insights_status"})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_apex"

	var resId, resId2 string
	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with required
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Required, Create,
					RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":    Representation{RepType: Required, Create: `19c`},
						"db_workload":   Representation{RepType: Required, Create: `APEX`},
						"license_model": Representation{RepType: Required, Create: `LICENSE_INCLUDED`},
						"is_free_tier":  Representation{RepType: Required, Create: `false`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Create,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "is_mtls_connection_required"},
						[]interface{}{Representation{RepType: Optional, Create: "APEX"},
							Representation{RepType: Optional, Create: `19c`},
							Representation{RepType: Optional, Create: `true`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{Representation{RepType: Optional, Create: "APEX"},
							Representation{RepType: Optional, Create: `19c`},
							Representation{RepType: Optional, Create: `NOT_ENABLED`},
							Representation{RepType: Optional, Create: `true`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify autoscaling with APEX workload
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{Representation{RepType: Optional, Create: "APEX"},
							Representation{RepType: Optional, Update: `true`},
							Representation{RepType: Optional, Create: `19c`},
							Representation{RepType: Optional, Create: `NOT_ENABLED`},
							Representation{RepType: Optional, Create: `true`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify Update db_workload to OLTP
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status", "is_mtls_connection_required"},
						[]interface{}{Representation{RepType: Optional, Create: "OLTP"},
							Representation{RepType: Optional, Update: `true`},
							Representation{RepType: Optional, Create: `19c`},
							Representation{RepType: Optional, Create: `NOT_ENABLED`},
							Representation{RepType: Optional, Create: `true`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
		},
		// verify Create OLTP with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":                           Representation{RepType: Required, Create: `19c`},
						"is_free_tier":                         Representation{RepType: Required, Create: `true`},
						"autonomous_maintenance_schedule_type": Representation{RepType: Optional, Create: `REGULAR`},
						"is_mtls_connection_required":          Representation{RepType: Required, Create: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "true"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
					RepresentationCopyWithRemovedProperties(RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":                           Representation{RepType: Required, Create: `19c`},
						"is_free_tier":                         Representation{RepType: Required, Create: `true`},
						"autonomous_maintenance_schedule_type": Representation{RepType: Optional, Create: `REGULAR`},
						"is_mtls_connection_required":          Representation{RepType: Optional, Create: `true`},
					}), []string{"operations_insights_status"})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify OLTP updated to APEX
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
					RepresentationCopyWithRemovedProperties(RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":                           Representation{RepType: Required, Create: `19c`},
						"db_workload":                          Representation{RepType: Required, Create: `APEX`},
						"is_free_tier":                         Representation{RepType: Required, Create: `false`},
						"autonomous_maintenance_schedule_type": Representation{RepType: Optional, Create: `REGULAR`},
						"is_mtls_connection_required":          Representation{RepType: Required, Create: `true`},
					}), []string{"operations_insights_status"})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create with required
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
					RepresentationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
						"db_version":    Representation{RepType: Required, Create: `19c`},
						"license_model": Representation{RepType: Required, Create: `LICENSE_INCLUDED`},
						"is_free_tier":  Representation{RepType: Required, Create: `false`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
					GetMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
						[]interface{}{Representation{RepType: Optional, Create: "OLTP"},
							Representation{RepType: Optional, Create: `19c`}}, autonomousDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
		},
		// Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDataGuardRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// test config key
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
						"kms_key_id": Representation{RepType: Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
						"vault_id":   Representation{RepType: Optional, Create: kmsVaultId},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}
