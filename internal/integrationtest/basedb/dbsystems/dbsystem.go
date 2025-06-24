package dbsystems

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	DbSystemResourceConfig         = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Create, DbSystemResourceRepresentation)
	DbSystemResourceRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDbSystem`},
		"database_edition":        acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Optional, Create: `4`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`, Update: `BRING_YOUR_OWN_LICENSE`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"fault_domains":           acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`${var.ssh_public_key}`}},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `tfDbHost`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeGroup},
	}

	dbHomeGroup = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `tfDbHome`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseGroup},
	}

	databaseGroup = map[string]interface{}{
		"db_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDb`},
		"pdb_name":           acctest.Representation{RepType: acctest.Optional, Create: `tfPdb`},
		"character_set":      acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"ncharacter_set":     acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"db_workload":        acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"kms_key_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"kms_key_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_version_id}`},
		"vault_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"admin_password":     acctest.Representation{RepType: acctest.Required, Create: `${var.admin_password}`},
	}
)
