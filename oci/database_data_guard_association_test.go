// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DataGuardAssociationRequiredOnlyResource = DataGuardAssociationResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", Required, Create, dataGuardAssociationRepresentationExistingDbSystem)

	DataGuardAssociationResourceConfig = DataGuardAssociationResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingDbSystem)

	dataGuardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"data_guard_association_id": Representation{repType: Required, create: `${oci_database_data_guard_association.test_data_guard_association.id}`},
		"database_id":               Representation{repType: Required, create: `${data.oci_database_databases.db.databases.0.id}`},
	}

	dataGuardAssociationDataSourceRepresentation = map[string]interface{}{
		"database_id": Representation{repType: Required, create: `${data.oci_database_databases.db.databases.0.id}`},
		"filter":      RepresentationGroup{Required, dataGuardAssociationDataSourceFilterRepresentation}}
	dataGuardAssociationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_data_guard_association.test_data_guard_association.id}`}},
	}

	dataGuardAssociationRepresentationBase = map[string]interface{}{
		"depends_on":                       Representation{repType: Required, create: []string{"oci_database_db_system.test_db_system"}},
		"database_admin_password":          Representation{repType: Required, create: `BEstrO0ng_#11`},
		"database_id":                      Representation{repType: Required, create: `${data.oci_database_databases.db.databases.0.id}`},
		"delete_standby_db_home_on_delete": Representation{repType: Required, create: `true`},
		"protection_mode":                  Representation{repType: Required, create: `MAXIMUM_PERFORMANCE`},
		"transport_type":                   Representation{repType: Required, create: `ASYNC`},
	}
	dataGuardAssociationRepresentationExistingDbSystem = representationCopyWithNewProperties(dataGuardAssociationRepresentationBase, map[string]interface{}{
		"depends_on":        Representation{repType: Required, create: []string{`oci_database_db_system.test_db_system`, `oci_database_db_system.test_db_system2`}},
		"creation_type":     Representation{repType: Required, create: `ExistingDbSystem`},
		"peer_db_system_id": Representation{repType: Required, create: `${oci_database_db_system.test_db_system2.id}`},
	})
	dataGuardAssociationRepresentationNewDbSystem = representationCopyWithNewProperties(dataGuardAssociationRepresentationBase, map[string]interface{}{
		"creation_type":       Representation{repType: Required, create: `NewDbSystem`},
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{repType: Required, create: `displayName`},
		"hostname":            Representation{repType: Required, create: `hostname`},
		"subnet_id":           Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
	})

	DataGuardAssociationResourceDependenciesBase = DefinedTagsDependencies + AvailabilityDomainConfig + VcnResourceConfig + `
#dataguard requires the some port to be open on the subnet
resource "oci_core_subnet" "test_subnet" {
  availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
  cidr_block          = "10.0.2.0/24"
  display_name        = "TFADSubnet"
  dns_label           = "adsubnet"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_vcn.test_vcn.id}"
  security_list_ids   = ["${oci_core_security_list.test_security_list.id}"]
  route_table_id      = "${oci_core_vcn.test_vcn.default_route_table_id}"
  dhcp_options_id     = "${oci_core_vcn.test_vcn.default_dhcp_options_id}"
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = "${var.compartment_id}"
  vcn_id         = "${oci_core_vcn.test_vcn.id}"
  display_name   = "TFExampleSecurityList"

  // allow outbound tcp traffic on all ports
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "6"
  }

  ingress_security_rules {
    protocol  = "6"
    source    = "0.0.0.0/0"
  }
}

data "oci_database_databases" "db" {
       compartment_id = "${var.compartment_id}"
       db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}

data "oci_database_db_homes" "t" {
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.test_db_system.id}"
	filter {
		name = "display_name"
		values = ["TFTestDbHome1"]
	}
}

`
	DataGuardAssociationResourceDependencies = DataGuardAssociationResourceDependenciesBase + `
resource "oci_database_db_system" "test_db_system" {
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	database_edition = "ENTERPRISE_EDITION"
	disk_redundancy = "NORMAL"
	shape = "BM.DenseIO2.52"
	cpu_core_count = "2"
	ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
	domain = "${oci_core_subnet.test_subnet.subnet_domain_name}"
	hostname = "myOracleDB"
	data_storage_size_in_gb = "256"
	license_model = "LICENSE_INCLUDED"
	node_count = "1"
	display_name = "TFTestDbSystemBM1"
	db_home {
		db_version = "12.1.0.2"
		display_name = "TFTestDbHome1"
		database {
			"admin_password" = "BEstrO0ng_#11"
			"db_name" = "tfDbName"
		}
	}
}

resource "oci_database_db_system" "test_db_system2" {
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	database_edition = "ENTERPRISE_EDITION"
	disk_redundancy = "NORMAL"
	shape = "BM.DenseIO2.52"
	cpu_core_count = "2"
	ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
	domain = "${oci_core_subnet.test_subnet.subnet_domain_name}"
	hostname = "myOracleDB"
	data_storage_size_in_gb = "256"
	license_model = "LICENSE_INCLUDED"
	node_count = "1"
	display_name = "TFTestDbSystemBM2"
	db_home {
		db_version = "12.1.0.2"
		display_name = "TFTestDbHome1"
		database {
			"admin_password" = "BEstrO0ng_#11"
			"db_name" = "db2"
		}
	}
}
`
	DataGuardAssociationResourceDependenciesNewDbSystem = DataGuardAssociationResourceDependenciesBase + `
resource "oci_database_db_system" "test_db_system" {
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	database_edition = "ENTERPRISE_EDITION"
	disk_redundancy = "NORMAL"
	shape = "VM.Standard2.1"
	cpu_core_count = "1"
	ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
	domain = "${oci_core_subnet.test_subnet.subnet_domain_name}"
	hostname = "myOracleDB"
	data_storage_size_in_gb = "256"
	license_model = "LICENSE_INCLUDED"
	node_count = "1"
	display_name = "TFTestDbSystemVM"
	db_home {
		db_version = "12.1.0.2"
		display_name = "TFTestDbHome1"
		database {
			"admin_password" = "BEstrO0ng_#11"
			"db_name" = "tfDbName"
		}
	}
}
`
)

func TestDatabaseDataGuardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDataGuardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_data_guard_association.test_data_guard_association"
	datasourceName := "data.oci_database_data_guard_associations.test_data_guard_associations"
	singularDatasourceName := "data.oci_database_data_guard_association.test_data_guard_association"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create existingDbSystem
			{
				Config: config + compartmentIdVariableStr + DataGuardAssociationResourceDependenciesNewDbSystem +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", Required, Create, dataGuardAssociationRepresentationNewDbSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "creation_type", "NewDbSystem"),
					resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
					resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DataGuardAssociationResourceDependencies,
			},
			// verify create with optionals NewDbSystem
			{
				Config: config + compartmentIdVariableStr + DataGuardAssociationResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", Optional, Create, dataGuardAssociationRepresentationExistingDbSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "creation_type", "ExistingDbSystem"),
					resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
					resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "role"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_data_guard_associations", "test_data_guard_associations", Optional, Update, dataGuardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + DataGuardAssociationResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingDbSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_id"),

					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_db_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_role"),
					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.role"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.transport_type", "ASYNC"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", Required, Create, dataGuardAssociationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DataGuardAssociationResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingDbSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_guard_association_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_db_system_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_role"),
					resource.TestCheckResourceAttr(singularDatasourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "transport_type", "ASYNC"),
				),
			},
			// @CODEGEN We need to remove the import because the import step removed the dependency relationship between the DGA and the peer dbSystem and the delete will fail without it.
		},
	})
}
