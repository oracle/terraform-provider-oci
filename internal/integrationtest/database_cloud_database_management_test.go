package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/integrationtest/basedb"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	CloudDatabaseManagementPrivateEndpointResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `tfPrivateEndpoint`, Update: `tfPrivateEndpoint2`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	CloudDatabaseManagementResourceRepresentation = map[string]interface{}{
		"database_id":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_databases.databases.0.id}`},
		"management_type":      acctest.Representation{RepType: acctest.Required, Create: `BASIC`, Update: `ADVANCED`},
		"private_end_point_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
		"service_name":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_databases.databases.0.db_unique_name}.${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"enable_management":    acctest.Representation{RepType: acctest.Required, Create: `true` /*, Update: `false`*/},
		"protocol":             acctest.Representation{RepType: acctest.Optional, Create: `TCP`, Update: `TCPS`},
		"port":                 acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"role":                 acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`, Update: `SYSDBA`},
		"ssl_secret_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.ssl_secret_id}`, Update: `${var.ssl_secret_id}`},
		"credentialdetails":    acctest.RepresentationGroup{RepType: acctest.Required, Group: credentialDetailsGroup},
	}

	credentialDetailsGroup = map[string]interface{}{
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `sys`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.ssl_secret_id}`},
	}
)

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.BaseDBProviderTestConfig()

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dbSystem
		{
			Config: config + basedb.BaseConfig,
		},
		// Enable database management with required fields
		{
			Config: config + basedb.BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, CloudDatabaseManagementPrivateEndpointResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, CloudDatabaseManagementResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr("data.oci_database_databases.test_databases", "databases.#", "1"),
			),
		},
		// Verify enable database management with required fields
		{
			Config: config + basedb.BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, CloudDatabaseManagementPrivateEndpointResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, CloudDatabaseManagementResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr("data.oci_database_databases.test_databases", "databases.#", "1"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_database", "db_name"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_database", "database_management_config.0.management_status", "ENABLED"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_database", "database_management_config.0.management_type", "BASIC"),
			),
		},
		// Update / disable database management with required fields
		{
			Config: config + basedb.BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, CloudDatabaseManagementPrivateEndpointResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Update, CloudDatabaseManagementResourceRepresentation),
		},
		// Verify Update database management with required fields
		{
			Config: config + basedb.BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, CloudDatabaseManagementPrivateEndpointResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Update, CloudDatabaseManagementResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr("data.oci_database_databases.test_databases", "databases.#", "1"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_database", "db_name"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_database", "database_management_config.0.management_type", "ADVANCED"),
			),
		},
		{
			// Enable database management with optional fields (for TCPS)
			Config: config + basedb.BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, CloudDatabaseManagementPrivateEndpointResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Create, CloudDatabaseManagementResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr("data.oci_database_databases.test_databases", "databases.#", "1"),
			),
		},
		// verify enable database management with optional fields (for TCPS)
		{
			Config: config + basedb.BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, CloudDatabaseManagementPrivateEndpointResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Create, CloudDatabaseManagementResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr("data.oci_database_databases.test_databases", "databases.#", "1"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_database", "db_name"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_database", "database_management_config.0.management_status", "ENABLED"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_database", "database_management_config.0.management_type", "BASIC"),
			),
		},
		// Update / disable database management with optional fields (for TCPS)
		{
			Config: config + basedb.BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, CloudDatabaseManagementPrivateEndpointResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Update, CloudDatabaseManagementResourceRepresentation),
		},
		// Verify Update / disable database management with optional fields (for TCPS)
		{
			Config: config + basedb.BaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, CloudDatabaseManagementPrivateEndpointResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Update, CloudDatabaseManagementResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr("data.oci_database_databases.test_databases", "databases.#", "1"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_database", "db_name"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_database", "database_management_config.0.management_type", "ADVANCED"),
			),
		},
	})
}
