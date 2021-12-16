package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	databaseManagementRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.id}`},
		//Update: `ADVANCED` to be uncommented to explicitly call ModifyDatabaseManagement
		"management_type":      acctest.Representation{RepType: acctest.Required, Create: `BASIC` /*, Update: `ADVANCED`*/},
		"private_end_point_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.dbmgmtprivateendpoint.oc1.ap-hyderabad-1.amaaaaaacsc5xjaamlmllhfxmxict6jf3irizwsydralyklninmwsrovggkq`},
		"service_name":         acctest.Representation{RepType: acctest.Required, Create: `DB0809_hyd17q.sub02231620340.dbmgmtcustomer.oraclevcn.com`},
		"credentialdetails":    acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseCredentialDetailsRepresentation},
		//Update: `false` to be uncommented to explicitly call DisableDatabaseManagement
		"enable_management": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	databaseCredentialDetailsRepresentation = map[string]interface{}{
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `dbsnmp`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.vaultsecret.oc1.ap-hyderabad-1.amaaaaaacsc5xjaa2q7r6kfzdm44ylxqwomht6uinb5zyhezka7sl2t62ecq`},
	}

	dbSystemForDatabaseManagementRepresentation = `
		resource "oci_database_db_system" "t" {
			availability_domain = "${var.availability_domain}"
			compartment_id = "${var.compartment_id}"
			subnet_id = "${var.subnet_id}"
			database_edition = "ENTERPRISE_EDITION"
			disk_redundancy = "NORMAL"
			shape = "VM.Standard2.1"
			ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
			display_name = "-tf-dbSystem-001"
			domain = "${var.domain_name}"
			hostname = "myOracleDB" // this will be lowercased server side
			data_storage_size_in_gb = "256"
			license_model = "LICENSE_INCLUDED"
			node_count = "1"
			fault_domains = ["FAULT-DOMAIN-1"]
			db_home {
				db_version = "19.0.0.0"
				display_name = "-tf-db-home"
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
)

func TestDatabaseCloudDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	domainName := utils.GetEnvSettingWithBlankDefault("domain_name")
	domainNameVariableStr := fmt.Sprintf("variable \"domain_name\" { default = \"%s\" }\n", domainName)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	availabilityDomain := utils.GetEnvSettingWithBlankDefault("availability_domain")
	availabilityDomainVariableStr := fmt.Sprintf("variable \"availability_domain\" { default = \"%s\" }\n", availabilityDomain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// Create dbSystem
			{
				Config: config + compartmentIdVariableStr + domainNameVariableStr + subnetIdVariableStr +
					availabilityDomainVariableStr + dbSystemForDatabaseManagementRepresentation,
			},
			// enable database management
			{
				Config: config + compartmentIdVariableStr + domainNameVariableStr + subnetIdVariableStr +
					availabilityDomainVariableStr + dbSystemForDatabaseManagementRepresentation +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, databaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
				),
			},
			// verify enable database management
			{
				Config: config + compartmentIdVariableStr + domainNameVariableStr + subnetIdVariableStr +
					availabilityDomainVariableStr + dbSystemForDatabaseManagementRepresentation +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, databaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
					resource.TestCheckResourceAttrSet("data.oci_database_database.t", "db_name"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_status", "ENABLED"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_type", "BASIC"),
				),
			},
			// Update / disable database management
			{
				Config: config + compartmentIdVariableStr + domainNameVariableStr + subnetIdVariableStr +
					availabilityDomainVariableStr + dbSystemForDatabaseManagementRepresentation +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Update, databaseManagementRepresentation),
			},
			// verify Update / disable database management
			{
				Config: config + compartmentIdVariableStr + domainNameVariableStr + subnetIdVariableStr +
					availabilityDomainVariableStr + dbSystemForDatabaseManagementRepresentation +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Update, databaseManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
					resource.TestCheckResourceAttrSet("data.oci_database_database.t", "db_name"),
					resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_status", "DISABLED"),
					//resource.TestCheckResourceAttr("data.oci_database_database.t", "database_management_config.0.management_type", "ADVANCED"),
				),
			},
		},
	})
}
