package dbsystems

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	DbSystemDatasourceConfig         = acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_systems", "test_db_systems", acctest.Optional, acctest.Create, DbSystemDatasourceRepresentation)
	DbSystemDatasourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: filterGroup},
	}

	filterGroup = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_db_system.test_db_system.id}`}},
	}

	DbHomesDatasourceConfig         = acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "test_db_homes", acctest.Optional, acctest.Create, DbHomesDatasourceRepresentation)
	DbHomesDatasourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_system_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_system.test_db_system.id}`},
	}

	DatabasesDatasourceConfig         = acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_databases", acctest.Optional, acctest.Create, DatabasesDatasourceRepresentation)
	DatabasesDatasourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_home_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_db_homes.test_db_homes.db_homes.0.db_home_id}`},
	}

	DatabaseDatasourceConfig         = acctest.GenerateDataSourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatasourceRepresentation)
	DatabaseDatasourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_databases.databases.0.id}`},
	}
)
