package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	cloudDatabaseManagementPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `TFPEforTCPS`, Update: `name2`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	cloudDatabaseManagementRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.id}`},
		//Update: `ADVANCED` to be uncommented to explicitly call ModifyDatabaseManagement
		"management_type":      acctest.Representation{RepType: acctest.Required, Create: `BASIC`, Update: `ADVANCED`},
		"private_end_point_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
		"service_name":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.db_unique_name}.${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"credentialdetails":    acctest.RepresentationGroup{RepType: acctest.Required, Group: cloudDatabaseManagementCredentialDetailsRepresentation},
		//Update: `false` to be uncommented to explicitly call DisableDatabaseManagement
		"enable_management": acctest.Representation{RepType: acctest.Required, Create: `true` /*, Update: `false`*/},
		"protocol":          acctest.Representation{RepType: acctest.Optional, Create: `TCP`, Update: `TCPS`},
		"port":              acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"role":              acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`, Update: `SYSDBA`},
		"ssl_secret_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.sslSecretId}`, Update: `${var.sslSecretId}`},
	}

	cloudDatabaseManagementCredentialDetailsRepresentation = map[string]interface{}{
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `sys`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sslSecretId}`},
	}

<<<<<<< ours
	DatabaseCloudDatabaseManagementResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, serviceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, backupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, databaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentation) +
		BackupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Create, backupDestinationNFSRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{RepType: Optional, Update: activationFilePath}})) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, exadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, vmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, vmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, userRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, secretRepresentation)
=======
	cloudDatabaseManagementPrivateEndpointConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, cloudDatabaseManagementPrivateEndpointRepresentation)

	CloudDatabaseManagementResourceDependenciesBase = AvailabilityDomainConfig + DefinedTagsDependencies + CoreVcnResourceConfig +
		`
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
		`

	CloudDatabaseManagementResourceDependencies = CloudDatabaseManagementResourceDependenciesBase + cloudDatabaseManagementPrivateEndpointConfig + `
		resource "oci_database_db_system" "t" {
			availability_domain = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
			compartment_id = "${var.compartment_id}"
			subnet_id = "${oci_core_subnet.test_subnet.id}"
			database_edition = "ENTERPRISE_EDITION"
			disk_redundancy = "NORMAL"
			shape = "VM.Standard2.1"
			ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
			display_name = "TFVMDBSystemForTCPS"
			domain = "${oci_core_subnet.test_subnet.subnet_domain_name}"
			hostname = "myOracleDB" // this will be lowercased server side
			data_storage_size_in_gb = "256"
			license_model = "LICENSE_INCLUDED"
			node_count = "1"
			fault_domains = ["FAULT-DOMAIN-1"]
			db_home {
				db_version = "19.0.0.0"
				display_name = "aTFdbhome"
				database {
					admin_password = "FIpassword12##"
					db_name = "aTFdb"
					character_set = "AL32UTF8"
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					ncharacter_set = "AL16UTF16"
					db_workload = "OLTP"
					pdb_name = "pdbName"
				}
			}
			db_system_options {
				storage_management = "LVM"
			}
			defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
			freeform_tags = {"Department" = "Finance"}
			lifecycle {
				ignore_changes = [
					db_home.0.db_version,
					defined_tags,
					db_home.0.database.0.defined_tags,
				]
			}
		}
		data "oci_database_db_systems" "t" {
			compartment_id = "${var.compartment_id}"
			filter {
				name   = "id"
				values = ["${oci_database_db_system.t.id}"]
			}
		}
		data "oci_database_db_homes" "t" {
			compartment_id = "${var.compartment_id}"
			db_system_id = "${oci_database_db_system.t.id}"
			filter {
				name   = "db_system_id"
				values = ["${oci_database_db_system.t.id}"]
			}
		}
		data "oci_database_db_home" "t" {
			db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
		}
		data "oci_database_databases" "t" {
			compartment_id = "${var.compartment_id}"
			db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
			filter {
				name   = "db_name"
				values = ["${oci_database_db_system.t.db_home.0.database.0.db_name}"]
			}
		}
		data "oci_database_database" "t" {
			  database_id = "${data.oci_database_databases.t.databases.0.id}"
		}`
>>>>>>> theirs
)

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sslSecretId := utils.GetEnvSettingWithDefault("ssl_secret_id", "test_secret_id")
	sslSecretIdVariableStr := fmt.Sprintf("variable \"sslSecretId\" { default = \"%s\" }\n", sslSecretId)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// Create dbSystem
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies,
			},
			// enable database management with required fields
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, cloudDatabaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
				),
			},
			// verify enable database management with required fields
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, cloudDatabaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
					resource.TestCheckResourceAttrSet("data.oci_database_database.t", "db_name"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_status", "ENABLED"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_type", "BASIC"),
				),
			},
			// Update / disable database management with required fields
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Update, cloudDatabaseManagementRepresentation),
			},
			// verify Update / disable database management with required fields
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Update, cloudDatabaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
					resource.TestCheckResourceAttrSet("data.oci_database_database.t", "db_name"),
					//resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_status", "DISABLED"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_type", "ADVANCED"),
				),
			},
			// enable database management with optional fields (for TCPS)
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Create, cloudDatabaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
				),
			},
			// verify enable database management with optional fields (for TCPS)
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Create, cloudDatabaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
					resource.TestCheckResourceAttrSet("data.oci_database_database.t", "db_name"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_status", "ENABLED"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_type", "BASIC"),
				),
			},
			// Update / disable database management with optional fields (for TCPS)
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Update, cloudDatabaseManagementRepresentation),
			},
			// verify Update / disable database management with optional fields (for TCPS)
			{
				Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + CloudDatabaseManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Update, cloudDatabaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
					resource.TestCheckResourceAttrSet("data.oci_database_database.t", "db_name"),
					//resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_status", "DISABLED"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_type", "ADVANCED"),
				),
			},
		},
	})
}
