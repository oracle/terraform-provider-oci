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
	oci_common "github.com/oracle/oci-go-sdk/v41/common"
	oci_database "github.com/oracle/oci-go-sdk/v41/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	adbDedicatedName                   = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbDedicatedUpdateName             = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbDedicatedCloneName              = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adDedicatedName                    = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adDedicatedUpdateName              = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbExaccName                       = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbBackupSourceName                = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbBackupIdName                    = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbBackupTimestampName             = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbPreviewDbName                   = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbDataSafeName                    = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbDbVersionName                   = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbDbRefreshableCloneName          = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbDbRefreshableCloneSourceADBName = randomString(1, charsetWithoutDigits) + randomString(13, charset)

	AutonomousDatabaseDedicatedRequiredOnlyResource = AutonomousDatabaseDedicatedResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseDedicatedRepresentation)

	AutonomousDatabaseDedicatedResourceConfig = AutonomousDatabaseDedicatedResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDedicatedRepresentation)

	autonomousDatabaseDedicatedDataSourceRepresentation = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(autonomousDatabaseDataSourceRepresentation, []string{"db_version"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"display_name":                     Representation{repType: Optional, create: adDedicatedName, update: adDedicatedUpdateName},
		})

	autonomousDatabaseDedicatedRepresentation = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbDedicatedName}, autonomousDatabaseRepresentation), []string{"license_model", "whitelisted_ips", "db_version", "is_auto_scaling_enabled", "customer_contacts"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     Representation{repType: Optional, create: `true`},
			"display_name":                     Representation{repType: Optional, create: adDedicatedName, update: adDedicatedUpdateName},
			"data_safe_status":                 Representation{repType: Optional, create: `REGISTERED`, update: `NOT_REGISTERED`},
		})

	autonomousDatabaseDedicatedRepresentationForClone = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbDedicatedCloneName}, autonomousDatabaseDedicatedRepresentation), []string{"license_model"}),
		map[string]interface{}{
			"clone_type":   Representation{repType: Optional, create: `FULL`},
			"display_name": Representation{repType: Optional, create: "example_autonomous_database_dedicated"},
			"source":       Representation{repType: Optional, create: `DATABASE`},
			"source_id":    Representation{repType: Optional, create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		})

	autonomousDatabaseDtaSafeStatusRepresentation = map[string]interface{}{
		"admin_password":           Representation{repType: Required, create: `BEstrO0ng_#11`, update: `BEstrO0ng_#12`},
		"compartment_id":           Representation{repType: Required, create: `${var.compartment_id}`},
		"cpu_core_count":           Representation{repType: Required, create: `1`},
		"data_storage_size_in_tbs": Representation{repType: Required, create: `1`},
		"db_name":                  Representation{repType: Required, create: adbDataSafeName},
		"db_workload":              Representation{repType: Optional, create: `OLTP`},
		"defined_tags":             Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":             Representation{repType: Optional, create: `example_autonomous_database`, update: `displayName2`},
		"freeform_tags":            Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_auto_scaling_enabled":  Representation{repType: Optional, create: `false`},
		"is_dedicated":             Representation{repType: Optional, create: `false`},
		"is_preview_version_with_service_terms_accepted": Representation{repType: Optional, create: `false`},
		"license_model":    Representation{repType: Optional, create: `LICENSE_INCLUDED`},
		"data_safe_status": Representation{repType: Optional, create: `REGISTERED`, update: `not_REGISTERED`},
		"timeouts":         RepresentationGroup{Required, autonomousDatabaseTimeoutsRepresentation},
	}

	autonomousDatabaseRepresentationForSourceFromBackupId = representationCopyWithNewProperties(
		getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbBackupIdName}, autonomousDatabaseRepresentation),
		map[string]interface{}{
			"clone_type":                    Representation{repType: Required, create: `FULL`},
			"source":                        Representation{repType: Required, create: `BACKUP_FROM_ID`},
			"autonomous_database_backup_id": Representation{repType: Required, create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`},
		})

	autonomousDatabaseRepresentationForSourceFromBackupTimestamp = representationCopyWithNewProperties(
		getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbBackupTimestampName}, autonomousDatabaseRepresentation),
		map[string]interface{}{
			"clone_type":             Representation{repType: Required, create: `FULL`},
			"source":                 Representation{repType: Required, create: `BACKUP_FROM_TIMESTAMP`},
			"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.autonomous_database_id}`},
			"timestamp":              Representation{repType: Required, create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.time_ended}`},
		})

	autonomousDatabaseDataGuardRepresentation = representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
		"db_version": Representation{repType: Optional, create: `19c`},
	})

	AutonomousDatabaseDedicatedResourceDependencies = AutonomousContainerDatabaseResourceConfig

	autonomousDatabaseRefreshableCloneSourceADBRepresentation = representationCopyWithNewProperties(
		autonomousDatabaseRepresentation, map[string]interface{}{
			"db_name":    Representation{repType: Required, create: adbDbRefreshableCloneSourceADBName},
			"db_version": Representation{repType: Optional, create: `19c`},
		})

	autonomousDatabaseRefreshableCloneRepresentation = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(autonomousDatabaseRepresentation, []string{"timeouts"}), map[string]interface{}{
			"admin_password":       Representation{repType: Optional, create: ``},
			"source":               Representation{repType: Required, create: `CLONE_TO_REFRESHABLE`},
			"db_name":              Representation{repType: Required, create: adbDbRefreshableCloneName},
			"source_id":            Representation{repType: Optional, create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
			"is_refreshable_clone": Representation{repType: Optional, create: `true`},
			"refreshable_mode":     Representation{repType: Optional, create: `MANUAL`},
			"db_version":           Representation{repType: Optional, create: `19c`},
		})

	autonomousDatabasesCloneDataSourceRepresentation2 = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"clone_type":             Representation{repType: Optional, create: `REFRESHABLE_CLONE`},
		"display_name":           Representation{repType: Optional, create: `example_autonomous_database`},
		"state":                  Representation{repType: Optional, create: `AVAILABLE`},
	}

	autonomousDatabasePrivateEndpointRepresentation = representationCopyWithRemovedProperties(
		representationCopyWithNewProperties(
			autonomousDatabaseRepresentation,
			map[string]interface{}{
				"nsg_ids":                Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, update: []string{`${oci_core_network_security_group.test_network_security_group.id}`, `${oci_core_network_security_group.test_network_security_group2.id}`}},
				"private_endpoint_label": Representation{repType: Optional, create: `xlx4fcli`},
				"subnet_id":              Representation{repType: Optional, create: `${oci_core_subnet.test_subnet.id}`},
			}), []string{"whitelisted_ips"})

	AutonomousDatabasePrivateEndpointResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", Required, Create, networkSecurityGroupRepresentation) +
		AutonomousDatabaseResourceDependencies

	AutonomousDatabaseFromBackupDependencies = AutonomousDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Required, Create, autonomousDatabaseBackupRepresentation) +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
			representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
				"db_name": Representation{repType: Required, create: adbBackupSourceName},
			}))

	autonomousDatabaseExaccRepresentation = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbExaccName}, autonomousDatabaseRepresentation), []string{"license_model", "whitelisted_ips", "db_version", "is_auto_scaling_enabled", "operations_insights_status"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     Representation{repType: Optional, create: `true`},
			"display_name":                     Representation{repType: Optional, create: adbExaccName},
			"is_access_control_enabled":        Representation{repType: Optional, create: `false`, update: `true`},
		})
	autonomousDatabaseDGExaccRepresentation = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbExaccName}, autonomousDatabaseRepresentation), []string{"license_model", "db_version", "is_auto_scaling_enabled", "operations_insights_status", "admin_password"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
			"is_dedicated":                     Representation{repType: Optional, create: `true`},
			"display_name":                     Representation{repType: Optional, create: adbExaccName},
			"is_access_control_enabled":        Representation{repType: Optional, create: `true`, update: `true`},
			"whitelisted_ips":                  Representation{repType: Optional, create: []string{`1.1.1.1/28`}, update: []string{`1.1.1.1/28`, `2.2.2.2/28`}},
			"standby_whitelisted_ips":          Representation{repType: Optional, update: []string{`3.4.5.6/28`, `3.6.7.8/28`}},
			"are_primary_whitelisted_ips_used": Representation{repType: Optional, create: `true`, update: `false`},
			"admin_password":                   Representation{repType: Required, create: `BEstrO0ng_#11`},
		})
	autonomousDatabaseUpdateExaccRepresentation = map[string]interface{}{
		"admin_password":                   Representation{repType: Required, create: `BEstrO0ng_#11`},
		"autonomous_container_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   Representation{repType: Required, create: `${var.compartment_id}`},
		"cpu_core_count":                   Representation{repType: Required, create: `1`},
		"data_storage_size_in_tbs":         Representation{repType: Required, create: `1`},
		"db_name":                          Representation{repType: Required, create: adbExaccName},
		"db_workload":                      Representation{repType: Optional, create: `OLTP`},
		"display_name":                     Representation{repType: Optional, create: adbExaccName},
		"is_auto_scaling_enabled":          Representation{repType: Optional, create: `false`},
		"is_dedicated":                     Representation{repType: Optional, create: `true`},
		"is_access_control_enabled":        Representation{repType: Optional, create: `false`, update: `true`},
	}

	autonomousDatabaseExaccRequiredOnlyResource = ExaccADBDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseExaccRepresentation)

	autonomousDatabaseExaccResourceConfig = ExaccADBDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseUpdateExaccRepresentation)

	ExaccADBDatabaseResourceDependencies = ACDatabaseResourceConfig

	ExaccADBWithDataguardResourceDependencies = ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig
)

func TestResourceDatabaseAutonomousDatabaseDedicated(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseDedicated")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
							"rotate_key_trigger": Representation{repType: Optional, create: `true`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
							"rotate_key_trigger": Representation{repType: Optional, create: `false`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
							"rotate_key_trigger": Representation{repType: Optional, create: `true`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
							"rotate_key_trigger": Representation{repType: Optional, create: `true`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, representationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{"db_name": Representation{repType: Optional, update: adbDedicatedUpdateName}})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Optional, Update, autonomousDatabaseDedicatedDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDedicatedRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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

			// verify create with optionals for Exacc
			{
				Config: config + compartmentIdVariableStr + ExaccADBDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseExaccRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to acl parameter for Exacc
			{
				Config: config + compartmentIdVariableStr + ExaccADBDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseUpdateExaccRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseDedicatedRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDedicatedRepresentationForClone),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource updated when it was supposed to be re-created.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_preview(t *testing.T) {
	t.Skip("Skip this test as this is a seasonal feature only when Dbaas has a preview to be released.")

	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_preview")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	autonomousDatabasePreviewRepresentation := getUpdatedRepresentationCopy("is_preview_version_with_service_terms_accepted", Representation{repType: Optional, create: `true`},
		getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbPreviewDbName}, autonomousDatabaseRepresentation))
	autonomousDatabasePreviewRepresentationForClone := getUpdatedRepresentationCopy("is_preview_version_with_service_terms_accepted", Representation{repType: Optional, create: `true`}, autonomousDatabaseRepresentationForClone)

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabasePreviewRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbPreviewDbName),
					// verify computed field db_workload to be defaulted to OLTP
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabasePreviewRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabasePreviewRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, representationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{"whitelisted_ips": Representation{repType: Optional, create: []string{"1.1.1.1/28", "1.1.1.29"}}})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
							"whitelisted_ips":         Representation{repType: Optional, create: []string{"1.1.1.1/28", "1.1.1.29"}},
							"is_auto_scaling_enabled": Representation{repType: Optional, update: `true`}})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Optional, Update, autonomousDatabaseDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
							"whitelisted_ips": Representation{repType: Optional, create: []string{"1.1.1.1/28", "1.1.1.29"}},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
							"whitelisted_ips": Representation{repType: Optional, create: []string{"1.1.1.1/28", "1.1.1.29"}},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{
							"whitelisted_ips": Representation{repType: Optional, create: []string{"1.1.1.1/28", "1.1.1.29"}},
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
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						getUpdatedRepresentationCopy("db_workload", Representation{repType: Optional, create: "DW"}, autonomousDatabasePreviewRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getUpdatedRepresentationCopy("db_workload", Representation{repType: Optional, create: "DW"}, autonomousDatabasePreviewRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled"},
							[]interface{}{Representation{repType: Optional, create: "DW"},
								Representation{repType: Optional, update: `true`}}, autonomousDatabasePreviewRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabasePreviewRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabasePreviewRepresentationForClone),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource updated when it was supposed to be re-created.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_dataSafeStatus(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_dataSafeStatus")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create and register
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDtaSafeStatusRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "data_safe_status", "REGISTERED"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// update: deregister data safe only
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						getUpdatedRepresentationCopy("data_safe_status", Representation{repType: Optional, create: `not_registered`}, autonomousDatabaseDtaSafeStatusRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(getUpdatedRepresentationCopy("data_safe_status", Representation{repType: Optional, create: `not_registered`}, autonomousDatabaseDtaSafeStatusRepresentation),
							map[string]interface{}{
								"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
							})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// update: all except data safe
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDtaSafeStatusRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// update: all except compartment (register data safe)
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDtaSafeStatusRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbDataSafeName),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "data_safe_status", "REGISTERED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDtaSafeStatusRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_FromBackupId(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_FromBackupFromId")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_from_backupid"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// create dependencies
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] Source ADB should be at least 2hrs old. Time Sleep for 2hrs")
					time.Sleep(2 * time.Hour)
					return nil
				},
			},
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backupid", Required, Create, autonomousDatabaseRepresentationForSourceFromBackupId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "db_name"),

					func(s *terraform.State) (err error) {
						resId, err := fromInstanceState(s, resourceName, "id")
						sourceresId, err := fromInstanceState(s, "oci_database_autonomous_database.test_autonomous_database", "id")
						if resId == sourceresId {
							return fmt.Errorf("resource not created when it was supposed to be created")
						}
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backupid", Optional, Create, autonomousDatabaseRepresentationForSourceFromBackupId),
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_FromBackupTimestamp(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_FromBackupTimestamp")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_from_backuptimestamp"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// create dependencies, To create clone the source db must be atleast 2 hrs old
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] Source ADB should be at least 2hrs old. Time Sleep for 2hrs")
					time.Sleep(2 * time.Hour)
					return nil
				},
			},
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backuptimestamp", Required, Create, autonomousDatabaseRepresentationForSourceFromBackupTimestamp),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "db_name"),

					func(s *terraform.State) (err error) {
						resId, err := fromInstanceState(s, resourceName, "id")
						sourceresId, err := fromInstanceState(s, "oci_database_autonomous_database.test_autonomous_database", "id")
						if resId == sourceresId {
							return fmt.Errorf("resource not created when it was supposed to be created")
						}
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backuptimestamp", Optional, Create, autonomousDatabaseRepresentationForSourceFromBackupTimestamp),
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_privateEndpoint(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_privateEndPoint")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabasePrivateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabasePrivateEndpointRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabasePrivateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Optional, Update, autonomousDatabaseDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getUpdatedRepresentationCopy("nsg_ids", Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}}, autonomousDatabasePrivateEndpointRepresentation)), Check: resource.ComposeAggregateTestCheckFunc(
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
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getUpdatedRepresentationCopy("nsg_ids", Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}}, autonomousDatabasePrivateEndpointRepresentation)), Check: resource.ComposeAggregateTestCheckFunc(
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
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies,
			},
			// verify create with no private end point
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithRemovedProperties(representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version": Representation{repType: Optional, create: `19c`},
						}), []string{"whitelisted_ips"})),
				Check: resource.ComposeAggregateTestCheckFunc(
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

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify turn on PE
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithRemovedProperties(representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":             Representation{repType: Optional, create: `19c`},
							"nsg_ids":                Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, update: []string{`${oci_core_network_security_group.test_network_security_group.id}`, `${oci_core_network_security_group.test_network_security_group2.id}`}},
							"private_endpoint_label": Representation{repType: Optional, create: `xlx4fcli`},
							"subnet_id":              Representation{repType: Optional, create: `${oci_core_subnet.test_subnet.id}`},
						}), []string{"whitelisted_ips"})), Check: resource.ComposeAggregateTestCheckFunc(
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

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies,
			},
			//Create ADB with private access and data safe registered
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabasePrivateEndpointRepresentation, map[string]interface{}{
							"db_version":       Representation{repType: Optional, create: `19c`},
							"data_safe_status": Representation{repType: Optional, create: `REGISTERED`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			//change network access to public
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabasePrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(autonomousDatabasePrivateEndpointRepresentation, []string{"nsg_ids", "private_endpoint_label", "subnet_id"}), map[string]interface{}{
							"nsg_ids":                Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, update: []string{}},
							"private_endpoint_label": Representation{repType: Optional, create: `null`},
							"subnet_id":              Representation{repType: Optional, create: `null`},
							"db_version":             Representation{repType: Optional, create: `19c`, update: `19c`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_dbVersion(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_dbVersion")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	autonomousDatabaseDbVersionUpdateRepresentation := getUpdatedRepresentationCopy("admin_password", Representation{repType: Required, create: `BEstrO0ng_#11`},
		getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbDbVersionName},
			getUpdatedRepresentationCopy("defined_tags", Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
				getUpdatedRepresentationCopy("display_name", Representation{repType: Optional, create: `example_autonomous_database`},
					getUpdatedRepresentationCopy("freeform_tags", Representation{repType: Optional, create: map[string]string{"Department": "Finance"}},
						getUpdatedRepresentationCopy("db_version", Representation{repType: Optional, create: "${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}", update: `19c`}, autonomousDatabaseRepresentation))))))

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDbVersionUpdateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbDbVersionName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "18c"),
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
			// verify update to only db_version
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDbVersionUpdateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify update of parameters except db_version
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getUpdatedRepresentationCopy("db_version", Representation{repType: Optional, update: `19c`},
							getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbDbVersionName}, autonomousDatabaseRepresentation))),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_dataGuard(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_dataGuard")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDataGuardRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
			// verify updates to updatable parameters and enable dataGuard
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
							"is_data_guard_enabled": Representation{repType: Optional, update: `true`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates disable dataGuard
			{
				PreConfig: waitTillCondition(testAccProvider, &resId, ListAutonomousDatabasesWaitCondition, standbyDbWaitConditionDuration,
					listListAutonomousDatabasesFetchOperation, "database", true),
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
							"is_data_guard_enabled": Representation{repType: Optional, update: `false`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceDatabaseExaccAutonomousDatabaseResource_dataGuard(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_dataGuard")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ExaccADBWithDataguardResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDGExaccRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDGExaccRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
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

func TestResourceDatabaseAutonomousDatabaseResource_switchover(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_switchover")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	const standbyDbWaitConditionDuration = time.Duration(60 * time.Minute)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDataGuardRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
			// verify enable dataGuard
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
							"is_data_guard_enabled": Representation{repType: Optional, create: `true`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify no-op when switchover is PRIMARY for first time
			{
				PreConfig: waitTillCondition(testAccProvider, &resId, ListAutonomousDatabasesWaitCondition, standbyDbWaitConditionDuration,
					listListAutonomousDatabasesFetchOperation, "database", true),
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
							"switchover_to": Representation{repType: Optional, update: `PRIMARY`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
							"switchover_to": Representation{repType: Optional, update: `STANDBY`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
							"switchover_to": Representation{repType: Optional, update: `PRIMARY`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
							"switchover_to": Representation{repType: Optional, update: `PRIMARY`},
						})) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Required, Create, autonomousDatabaseDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDataGuardRepresentation, map[string]interface{}{
							"switchover_to": Representation{repType: Optional, update: `PRIMARY`},
						})) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_refreshableClone(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_refreshableClone")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	var resId, resId2 string

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"
	clonesDatasourceName := "data.oci_database_autonomous_databases_clones.test_autonomous_databases_clones"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseRefreshableCloneRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify LIST clones given a source ADB
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseRefreshableCloneRepresentation) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_databases_clones", "test_autonomous_databases_clones", Optional, Create, autonomousDatabasesCloneDataSourceRepresentation2),
				Check: resource.ComposeAggregateTestCheckFunc(
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
			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabaseRefreshableCloneRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseRefreshableCloneRepresentation),
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
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseRefreshableCloneRepresentation) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseDataSourceRepresentation, map[string]interface{}{
							"db_version": Representation{repType: Required, create: `19c`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseRefreshableCloneRepresentation) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRefreshableCloneSourceADBRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithNewProperties(autonomousDatabaseRefreshableCloneRepresentation, map[string]interface{}{
							"is_refreshable_clone": Representation{repType: Optional, update: `false`},
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_AJD(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_AJD")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create with required
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
						representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":    Representation{repType: Required, create: `19c`},
							"db_workload":   Representation{repType: Required, create: `AJD`},
							"license_model": Representation{repType: Required, create: `LICENSE_INCLUDED`},
							"is_free_tier":  Representation{repType: Required, create: `false`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "AJD"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
							[]interface{}{Representation{repType: Optional, create: "AJD"},
								Representation{repType: Optional, create: `19c`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "operations_insights_status"},
							[]interface{}{Representation{repType: Optional, create: "AJD"},
								Representation{repType: Optional, create: `19c`},
								Representation{repType: Optional, create: `NOT_ENABLED`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status"},
							[]interface{}{Representation{repType: Optional, create: "AJD"},
								Representation{repType: Optional, update: `true`},
								Representation{repType: Optional, create: `19c`},
								Representation{repType: Optional, create: `NOT_ENABLED`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify update db_workload to OLTP
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status"},
							[]interface{}{Representation{repType: Optional, create: "OLTP"},
								Representation{repType: Optional, update: `true`},
								Representation{repType: Optional, create: `19c`},
								Representation{repType: Optional, create: `NOT_ENABLED`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
			},
			// verify create OLTP with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":   Representation{repType: Required, create: `19c`},
							"is_free_tier": Representation{repType: Required, create: `true`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithRemovedProperties(representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":   Representation{repType: Required, create: `19c`},
							"is_free_tier": Representation{repType: Required, create: `true`},
						}), []string{"operations_insights_status"})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						representationCopyWithRemovedProperties(representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":   Representation{repType: Required, create: `19c`},
							"db_workload":  Representation{repType: Required, create: `AJD`},
							"is_free_tier": Representation{repType: Required, create: `false`},
						}), []string{"operations_insights_status"})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceDatabaseAutonomousDatabaseResource_APEX(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_APEX")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_apex"

	var resId, resId2 string
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create with required
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Required, Create,
						representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":    Representation{repType: Required, create: `19c`},
							"db_workload":   Representation{repType: Required, create: `APEX`},
							"license_model": Representation{repType: Required, create: `LICENSE_INCLUDED`},
							"is_free_tier":  Representation{repType: Required, create: `false`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "APEX"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Create,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
							[]interface{}{Representation{repType: Optional, create: "APEX"},
								Representation{repType: Optional, create: `19c`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version", "operations_insights_status"},
							[]interface{}{Representation{repType: Optional, create: "APEX"},
								Representation{repType: Optional, create: `19c`},
								Representation{repType: Optional, create: `NOT_ENABLED`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status"},
							[]interface{}{Representation{repType: Optional, create: "APEX"},
								Representation{repType: Optional, update: `true`},
								Representation{repType: Optional, create: `19c`},
								Representation{repType: Optional, create: `NOT_ENABLED`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify update db_workload to OLTP
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version", "operations_insights_status"},
							[]interface{}{Representation{repType: Optional, create: "OLTP"},
								Representation{repType: Optional, update: `true`},
								Representation{repType: Optional, create: `19c`},
								Representation{repType: Optional, create: `NOT_ENABLED`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
			},
			// verify create OLTP with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":   Representation{repType: Required, create: `19c`},
							"is_free_tier": Representation{repType: Required, create: `true`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
						representationCopyWithRemovedProperties(representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":   Representation{repType: Required, create: `19c`},
							"is_free_tier": Representation{repType: Required, create: `true`},
						}), []string{"operations_insights_status"})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_apex", Optional, Update,
						representationCopyWithRemovedProperties(representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"db_version":   Representation{repType: Required, create: `19c`},
							"db_workload":  Representation{repType: Required, create: `APEX`},
							"is_free_tier": Representation{repType: Required, create: `false`},
						}), []string{"operations_insights_status"})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}
