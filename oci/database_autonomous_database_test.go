// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v30/common"
	oci_database "github.com/oracle/oci-go-sdk/v30/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousDatabaseRequiredOnlyResource = AutonomousDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation)

	AutonomousDatabaseResourceConfig = AutonomousDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseRepresentation)

	autonomousDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	autonomousDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"db_version":     Representation{repType: Optional, create: `${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}`},
		"db_workload":    Representation{repType: Optional, create: `OLTP`},
		"display_name":   Representation{repType: Optional, create: `example_autonomous_database`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, autonomousDatabaseDataSourceFilterRepresentation}}
	autonomousDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_autonomous_database.test_autonomous_database.id}`}},
	}

	adbName      = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbCloneName = randomString(1, charsetWithoutDigits) + randomString(13, charset)

	autonomousDatabaseRepresentation = map[string]interface{}{
		"compartment_id":           Representation{repType: Required, create: `${var.compartment_id}`},
		"cpu_core_count":           Representation{repType: Required, create: `1`},
		"data_storage_size_in_tbs": Representation{repType: Required, create: `1`},
		"db_name":                  Representation{repType: Required, create: adbName},
		"admin_password":           Representation{repType: Required, create: `BEstrO0ng_#11`, update: `BEstrO0ng_#12`},
		"db_version":               Representation{repType: Optional, create: `${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}`},
		"db_workload":              Representation{repType: Optional, create: `OLTP`},
		"defined_tags":             Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":             Representation{repType: Optional, create: `example_autonomous_database`, update: `displayName2`},
		"freeform_tags":            Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_auto_scaling_enabled":  Representation{repType: Optional, create: `false`},
		"is_dedicated":             Representation{repType: Optional, create: `false`},
		"is_preview_version_with_service_terms_accepted": Representation{repType: Optional, create: `false`},
		"license_model":              Representation{repType: Optional, create: `LICENSE_INCLUDED`},
		"whitelisted_ips":            Representation{repType: Optional, create: []string{`1.1.1.1/28`}},
		"operations_insights_status": Representation{repType: Optional, create: `NOT_ENABLED`, update: `ENABLED`},
		"timeouts":                   RepresentationGroup{Required, autonomousDatabaseTimeoutsRepresentation},
		"state":                      Representation{repType: Optional, create: `AVAILABLE`},
	}

	autonomousDatabaseTimeoutsRepresentation = map[string]interface{}{
		"create": Representation{repType: Required, create: `20m`},
		"update": Representation{repType: Required, create: `20m`},
		"delete": Representation{repType: Required, create: `20m`},
	}
	autonomousDatabaseCopyWithUpdatedIPsRepresentation = getUpdatedRepresentationCopy("whitelisted_ips", Representation{repType: Optional, create: []string{"1.1.1.1/28", "1.1.1.29"}, update: []string{}}, autonomousDatabaseRepresentation)

	autonomousDatabaseRepresentationForClone = representationCopyWithNewProperties(
		getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbCloneName}, autonomousDatabaseRepresentation),
		map[string]interface{}{
			"clone_type": Representation{repType: Optional, create: `FULL`},
			"source":     Representation{repType: Optional, create: `DATABASE`},
			"source_id":  Representation{repType: Optional, create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		})

	AutonomousDatabaseResourceDependencies = DefinedTagsDependencies +
		generateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", Required, Create, autonomousDbVersionDataSourceRepresentation) +
		generateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", Required, Create,
			representationCopyWithNewProperties(autonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": Representation{repType: Required, create: `DW`}}))
)

func TestDatabaseAutonomousDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseResource_basic")
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
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"open_mode":        Representation{repType: Optional, create: `READ_ONLY`, update: `READ_ONLY`},
							"permission_level": Representation{repType: Optional, create: `RESTRICTED`, update: `RESTRICTED`},
						}),
					),
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
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "open_mode", "READ_ONLY"),
					resource.TestCheckResourceAttr(resourceName, "operations_insights_status", "NOT_ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "permission_level", "RESTRICTED"),

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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						representationCopyWithNewProperties(autonomousDatabaseRepresentation, map[string]interface{}{
							"compartment_id":   Representation{repType: Required, create: `${var.compartment_id_for_update}`},
							"open_mode":        Representation{repType: Optional, create: `READ_WRITE`, update: `READ_WRITE`},
							"permission_level": Representation{repType: Optional, create: `UNRESTRICTED`, update: `UNRESTRICTED`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
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
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "open_mode", "READ_WRITE"),
					resource.TestCheckResourceAttr(resourceName, "permission_level", "UNRESTRICTED"),

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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseRepresentation),
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
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "open_mode", "READ_WRITE"),
					resource.TestCheckResourceAttr(resourceName, "operations_insights_status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "permission_level", "UNRESTRICTED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify stop the autonomous database
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getUpdatedRepresentationCopy("state", Representation{repType: Optional, create: "STOPPED"}, autonomousDatabaseRepresentation)),
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
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify start the autonomous database
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getUpdatedRepresentationCopy("state", Representation{repType: Optional, create: "AVAILABLE"}, autonomousDatabaseRepresentation)),
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
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),

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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getUpdatedRepresentationCopy("whitelisted_ips", Representation{repType: Optional, create: []string{"1.1.1.1/28", "1.1.1.29"}}, autonomousDatabaseRepresentation)),
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
					resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
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
			// verify remove whitelisted_ips
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseCopyWithUpdatedIPsRepresentation),
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
					resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "0"),

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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, representationCopyWithNewProperties(autonomousDatabaseCopyWithUpdatedIPsRepresentation, map[string]interface{}{"is_auto_scaling_enabled": Representation{repType: Optional, update: `true`}})),
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
					resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "0"),

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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseCopyWithUpdatedIPsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.apex_details.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.available_upgrade_versions.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.backup_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_strings.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_urls.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.cpu_core_count", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.data_storage_size_in_gb"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_safe_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.db_version"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", adbName),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_workload", "OLTP"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.is_dedicated", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.is_preview"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.open_mode"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.operations_insights_status"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.permission_level"),
					// @Codegen: Can't test private_endpoint with fake resource
					//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.private_endpoint"),
					//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.private_endpoint_ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_maintenance_begin"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_maintenance_end"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "apex_details.#"),
					resource.TestCheckResourceAttr(singularDatasourceName, "available_upgrade_versions.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_strings.0.all_connection_strings.%"),

					resource.TestCheckResourceAttr(singularDatasourceName, "connection_urls.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.apex_url"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.machine_learning_user_management_url"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.sql_dev_web_url"),

					resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "data_safe_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_gb"),
					resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adbName),
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
					resource.TestCheckResourceAttrSet(singularDatasourceName, "open_mode"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "permission_level"),
					// @Codegen: Can't test private_endpoint with fake resource
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_maintenance_begin"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_maintenance_end"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"admin_password",
					"autonomous_database_backup_id",
					"clone_type",
					"is_preview_version_with_service_terms_accepted",
					"source",
					"source_id",
					"lifecycle_details",
					"timestamp",
					// Need this workaround due to import behavior change introduced by https://github.com/hashicorp/terraform/issues/20985
					"used_data_storage_size_in_tbs",
				},
				ResourceName: resourceName,
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr,
			},
			// test ADW db_workload
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
							[]interface{}{Representation{repType: Optional, create: "DW"},
								Representation{repType: Optional, create: `${data.oci_database_autonomous_db_versions.test_autonomous_dw_versions.autonomous_db_versions.0.version}`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
					resource.TestCheckResourceAttrSet(resourceName, "db_version"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
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
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "db_version"},
							[]interface{}{Representation{repType: Optional, create: "DW"},
								Representation{repType: Optional, create: `${data.oci_database_autonomous_db_versions.test_autonomous_dw_versions.autonomous_db_versions.0.version}`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
					resource.TestCheckResourceAttrSet(resourceName, "db_version"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
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

			// verify autoscaling with DW workload
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update,
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled", "db_version"},
							[]interface{}{Representation{repType: Optional, create: "DW"},
								Representation{repType: Optional, update: `true`},
								Representation{repType: Optional, create: `${data.oci_database_autonomous_db_versions.test_autonomous_dw_versions.autonomous_db_versions.0.version}`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
					resource.TestCheckResourceAttrSet(resourceName, "db_version"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "DW"),
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

			// remove any previously created resources
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
			},
			// verify ADB clone from a source ADB
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_source", Optional, Create, autonomousDatabaseRepresentation) +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseRepresentationForClone),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "clone_type", "FULL"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbCloneName),
					resource.TestCheckResourceAttrSet(resourceName, "db_version"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
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

func testAccCheckDatabaseAutonomousDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_database" {
			noResourceFound = false
			request := oci_database.GetAutonomousDatabaseRequest{}

			tmp := rs.Primary.ID
			request.AutonomousDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseAutonomousDatabase") {
		resource.AddTestSweepers("DatabaseAutonomousDatabase", &resource.Sweeper{
			Name:         "DatabaseAutonomousDatabase",
			Dependencies: DependencyGraph["autonomousDatabase"],
			F:            sweepDatabaseAutonomousDatabaseResource,
		})
	}
}

func sweepDatabaseAutonomousDatabaseResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	autonomousDatabaseIds, err := getAutonomousDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousDatabaseId := range autonomousDatabaseIds {
		if ok := SweeperDefaultResourceId[autonomousDatabaseId]; !ok {
			deleteAutonomousDatabaseRequest := oci_database.DeleteAutonomousDatabaseRequest{}

			deleteAutonomousDatabaseRequest.AutonomousDatabaseId = &autonomousDatabaseId

			deleteAutonomousDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteAutonomousDatabase(context.Background(), deleteAutonomousDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousDatabaseId, error)
				continue
			}
			waitTillCondition(testAccProvider, &autonomousDatabaseId, autonomousDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				autonomousDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getAutonomousDatabaseIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "AutonomousDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

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
			addResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousDatabaseId", id)
		}
	}
	return resourceIds, nil
}

func autonomousDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDatabaseResponse, ok := response.Response.(oci_database.GetAutonomousDatabaseResponse); ok {
		return autonomousDatabaseResponse.LifecycleState != oci_database.AutonomousDatabaseLifecycleStateTerminated
	}
	return false
}

func autonomousDatabaseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetAutonomousDatabase(context.Background(), oci_database.GetAutonomousDatabaseRequest{
		AutonomousDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
