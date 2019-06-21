package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	adbDedicatedName       = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbDedicatedUpdateName = randomString(1, charsetWithoutDigits) + randomString(13, charset)
	adbDedicatedCloneName  = randomString(1, charsetWithoutDigits) + randomString(13, charset)

	AutonomousDatabaseDedicatedRequiredOnlyResource = AutonomousDatabaseDedicatedResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseDedicatedRepresentation)

	AutonomousDatabaseDedicatedResourceConfig = AutonomousDatabaseDedicatedResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDedicatedRepresentation)

	autonomousDatabaseDedicatedDataSourceRepresentation = representationCopyWithNewProperties(
		autonomousDatabaseDataSourceRepresentation,
		map[string]interface{}{
			"autonomous_container_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		})

	autonomousDatabaseDedicatedRepresentation = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbDedicatedName}, autonomousDatabaseRepresentation), []string{"license_model"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     Representation{repType: Optional, create: `true`},
		})

	autonomousDatabaseDedicatedRepresentationForClone = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbDedicatedCloneName}, autonomousDatabaseDedicatedRepresentation), []string{"license_model"}),
		map[string]interface{}{
			"clone_type": Representation{repType: Optional, create: `FULL`},
			"source":     Representation{repType: Optional, create: `DATABASE`},
			"source_id":  Representation{repType: Optional, create: `${oci_database_autonomous_database.test_autonomous_database_source.id}`},
		})

	AutonomousDatabaseDedicatedResourceDependencies = AutonomousContainerDatabaseResourceConfig
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDedicatedRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabaseDedicatedRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbDedicatedName),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_dedicated", "true"),
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
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
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
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
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
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
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
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"admin_password",
					"clone_type",
					"source",
					"source_id",
					"lifecycle_details",
					"is_auto_scaling_enabled",
				},
				ResourceName: resourceName,
			},

			// remove any previously created resources
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseDedicatedResourceDependencies,
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
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
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
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_preview")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	autonomousDatabasePreviewRepresentation := getUpdatedRepresentationCopy("is_preview_version_with_service_terms_accepted", Representation{repType: Optional, create: `true`}, autonomousDatabaseRepresentation)
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabasePreviewRepresentation),
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
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
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
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, representationCopyWithNewProperties(autonomousDatabasePreviewRepresentation, map[string]interface{}{"is_auto_scaling_enabled": Representation{repType: Optional, update: `true`}})),
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Update, autonomousDatabasePreviewRepresentation),
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
						getUpdatedRepresentationCopy("db_workload", Representation{repType: Optional, create: "DW"}, autonomousDatabasePreviewRepresentation)),
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
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
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
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
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
