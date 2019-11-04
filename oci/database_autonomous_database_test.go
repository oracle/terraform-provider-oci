// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"

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
		"admin_password":           Representation{repType: Required, create: `BEstrO0ng_#11`, update: `BEstrO0ng_#12`},
		"compartment_id":           Representation{repType: Required, create: `${var.compartment_id}`},
		"cpu_core_count":           Representation{repType: Required, create: `1`},
		"data_storage_size_in_tbs": Representation{repType: Required, create: `1`},
		"db_name":                  Representation{repType: Required, create: adbName},
		"db_workload":              Representation{repType: Optional, create: `OLTP`},
		"defined_tags":             Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":             Representation{repType: Optional, create: `example_autonomous_database`, update: `displayName2`},
		"freeform_tags":            Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_auto_scaling_enabled":  Representation{repType: Optional, create: `false`},
		"is_dedicated":             Representation{repType: Optional, create: `false`},
		"is_preview_version_with_service_terms_accepted": Representation{repType: Optional, create: `false`},
		"license_model":   Representation{repType: Optional, create: `LICENSE_INCLUDED`},
		"whitelisted_ips": Representation{repType: Optional, create: []string{`1.1.1.1/28`}},
	}

	autonomousDatabaseCopyWithUpdatedIPsRepresentation = getUpdatedRepresentationCopy("whitelisted_ips", Representation{repType: Optional, create: []string{"1.1.1.1/28", "1.1.1.29"}, update: []string{}}, autonomousDatabaseRepresentation)

	autonomousDatabaseRepresentationForClone = representationCopyWithNewProperties(
		getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbCloneName}, autonomousDatabaseRepresentation),
		map[string]interface{}{
			"clone_type": Representation{repType: Optional, create: `FULL`},
			"source":     Representation{repType: Optional, create: `DATABASE`},
			"source_id":  Representation{repType: Optional, create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		})

	AutonomousDatabaseResourceDependencies = DefinedTagsDependencies
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
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
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
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),

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
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_safe_status", "NOT_REGISTERED"),
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
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_ips.#", "1"),

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
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_strings.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.connection_urls.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.cpu_core_count", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.data_safe_status", "NOT_REGISTERED"),
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
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.0.all_connection_strings.%", "4"),

					resource.TestCheckResourceAttr(singularDatasourceName, "connection_urls.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.apex_url"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.machine_learning_user_management_url"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_urls.0.sql_dev_web_url"),

					resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "data_safe_status", "NOT_REGISTERED"),
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
					resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
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
						getUpdatedRepresentationCopy("db_workload", Representation{repType: Optional, create: "DW"}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
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
						getUpdatedRepresentationCopy("db_workload", Representation{repType: Optional, create: "DW"}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
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
						getMultipleUpdatedRepresenationCopy([]string{"db_workload", "is_auto_scaling_enabled"},
							[]interface{}{Representation{repType: Optional, create: "DW"},
								Representation{repType: Optional, update: `true`}}, autonomousDatabaseRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
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
	client := testAccProvider.Meta().(*OracleClients).databaseClient
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
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient
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
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient

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
	_, err := client.databaseClient.GetAutonomousDatabase(context.Background(), oci_database.GetAutonomousDatabaseRequest{
		AutonomousDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
