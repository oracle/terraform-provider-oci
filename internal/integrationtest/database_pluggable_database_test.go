// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabasePluggableDatabaseRequiredOnlyResource = DatabasePluggableDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation)

	DatabasePluggableDatabaseResourceConfig = DatabasePluggableDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Update, DatabasePluggableDatabaseRepresentation)

	DatabaseDatabasePluggableDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
	}
	ignoreChangesPdbepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	DatabaseDatabasePluggableDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"pdb_name":       acctest.Representation{RepType: acctest.Optional, Create: `SalesPdb`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabasePluggableDatabaseDataSourceFilterRepresentation}}
	DatabasePluggableDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_pluggable_database.test_pluggable_database.id}`}},
	}

	DatabasePluggableDatabaseRepresentation = map[string]interface{}{
		"container_database_id":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_database.t.id}`},
		"pdb_admin_password":                 acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pdb_name":                           acctest.Representation{RepType: acctest.Required, Create: `SalesPdb`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"should_pdb_admin_account_be_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"depends_on":                         acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.t"}},
		"kms_key_version_id":                 acctest.Representation{RepType: acctest.Required, Update: `${var.kms_key_version_id}`},
	}

	DatabasePluggableDatabaseLocalCloneRepresentation = map[string]interface{}{
		"container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_database.t.id}`},
		"pdb_admin_password":    acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pdb_name":              acctest.Representation{RepType: acctest.Required, Create: `LocalClonePdb`},
		//"tde_wallet_password":                acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"should_pdb_admin_account_be_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		//"container_database_admin_password":  acctest.Representation{RepType: acctest.Optional, Create: `containerDatabaseAdminPassword`},
		"pdb_creation_type_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabasePluggableDatabasePdbCreationTypeDetailsLocalCloneRepresentation},
		"should_create_pdb_backup":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"depends_on":                acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.t"}},
	}

	DatabasePluggableDatabasePdbCreationTypeDetailsLocalCloneRepresentation = map[string]interface{}{
		"creation_type":                            acctest.Representation{RepType: acctest.Required, Create: `LOCAL_CLONE_PDB`},
		"source_pluggable_database_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
		"source_container_database_admin_password": acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
	}

	DatabasePluggableDatabaseRemoteCloneRepresentation = map[string]interface{}{
		"container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_database.t.id}`},
		"pdb_admin_password":    acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pdb_name":              acctest.Representation{RepType: acctest.Required, Create: `RemoteClonePdb`},
		//"tde_wallet_password":                acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"should_pdb_admin_account_be_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		//"container_database_admin_password":  acctest.Representation{RepType: acctest.Optional, Create: `containerDatabaseAdminPassword`},
		"pdb_creation_type_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabasePluggableDatabasePdbCreationTypeDetailsRemoteCloneRepresentation},
		"should_create_pdb_backup":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"depends_on":                acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.t"}},
	}

	DatabasePluggableDatabasePdbCreationTypeDetailsRemoteCloneRepresentation = map[string]interface{}{
		"creation_type":                            acctest.Representation{RepType: acctest.Required, Create: `REMOTE_CLONE_PDB`},
		"source_pluggable_database_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
		"source_container_database_admin_password": acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
	}

	DatabasePluggableDatabaseRelocateRepresentation = map[string]interface{}{
		"container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_database.t.id}`},
		"pdb_admin_password":    acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pdb_name":              acctest.Representation{RepType: acctest.Required, Create: `RelocatePdb`},
		//"tde_wallet_password":                acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"should_pdb_admin_account_be_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		//"container_database_admin_password":  acctest.Representation{RepType: acctest.Optional, Create: `containerDatabaseAdminPassword`},
		"pdb_creation_type_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabasePluggableDatabasePdbCreationTypeDetailsRelocateRepresentation},
		"should_create_pdb_backup":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"depends_on":                acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.t"}},
	}

	DatabasePluggableDatabasePdbCreationTypeDetailsRelocateRepresentation = map[string]interface{}{
		"creation_type":                            acctest.Representation{RepType: acctest.Required, Create: `RELOCATE_PDB`},
		"source_pluggable_database_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
		"source_container_database_admin_password": acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
	}

	DatabasePluggableDatabaseRefreshableCloneRepresentation = map[string]interface{}{
		"container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_database.t.id}`},
		"pdb_admin_password":    acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pdb_name":              acctest.Representation{RepType: acctest.Required, Create: `RefreshablePdb`},
		//"tde_wallet_password":                acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"should_pdb_admin_account_be_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		//"container_database_admin_password":  acctest.Representation{RepType: acctest.Optional, Create: `containerDatabaseAdminPassword`},
		"pdb_creation_type_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabasePluggableDatabasePdbCreationTypeDetailsRelocateRepresentation},
		"should_create_pdb_backup":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"depends_on":                acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.t"}},
	}

	DatabasePluggableDatabasePdbCreationTypeDetailsRefreshableCloneRepresentation = map[string]interface{}{
		"creation_type":                            acctest.Representation{RepType: acctest.Required, Create: `REMOTE_CLONE`},
		"source_pluggable_database_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
		"dblink_user_password":                     acctest.Representation{RepType: acctest.Optional, Create: `DBLINKPWD`},
		"dblink_username":                          acctest.Representation{RepType: acctest.Optional, Create: `DBLINKUSER`},
		"refreshable_clone_details":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabasePluggableDatabasePdbCreationTypeDetailsRefreshableCloneDetailsRepresentation},
		"source_container_database_admin_password": acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
	}
	DatabasePluggableDatabasePdbCreationTypeDetailsRefreshableCloneDetailsRepresentation = map[string]interface{}{
		"is_refreshable_clone": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	ResourcePluggableDatabaseBaseConfig = `

	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.1.0.0/16"
		display_name = "-tf-vcn"
		dns_label = "tfvcn"
	}

	resource "oci_core_route_table" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		route_rules {
			cidr_block = "0.0.0.0/0"
			network_entity_id = "${oci_core_internet_gateway.t.id}"
		}
	}
	resource "oci_core_internet_gateway" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "-tf-internet-gateway"
	}

	resource "oci_core_subnet" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block          = "10.1.20.0/24"
		display_name        = "TFSubnet1"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_route_table.t.id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "tfsubnet"
	}
	resource "oci_core_subnet" "t2" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block          = "10.1.21.0/24"
		display_name        = "TFSubnet2"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_route_table.t.id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "tfsubnet2"
	}
	resource "oci_core_network_security_group" "test_network_security_group" {
         compartment_id  = "${var.compartment_id}"
		 vcn_id            = "${oci_core_virtual_network.t.id}"
         display_name      =  "displayName"
    }

	resource "oci_core_network_security_group" "test_network_security_group2" {
		compartment_id = "${var.compartment_id}"
		vcn_id            = "${oci_core_virtual_network.t.id}"
	}`

	dbSystemForPluggableDbRepresentation = `
		resource "oci_database_db_system" "t" {
			compartment_id = "${var.compartment_id}"
			subnet_id = "${oci_core_subnet.t.id}"
			database_edition = "ENTERPRISE_EDITION"
			availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
			disk_redundancy = "NORMAL"
			shape = "VM.Standard2.1"
			ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
			display_name = "-tf-dbSystem-001"
			domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
			hostname = "myOracleDB" // this will be lowercased server side
			data_storage_size_in_gb = "256"
			license_model = "LICENSE_INCLUDED"
			node_count = "1"
			cpu_core_count = "1"
			fault_domains = ["FAULT-DOMAIN-1"]
			kms_key_id = "${var.kms_key_id}"
			db_home {
				db_version = "21.14.0.0"
				display_name = "-tf-db-home"
				database {
					admin_password = "BEstrO0ng_#11"
					db_name = "aTFdb"
					character_set = "AL32UTF8"
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					ncharacter_set = "AL16UTF16"
					db_workload = "OLTP"
					pdb_name = "pdbName"
					db_backup_config {
						auto_backup_enabled = true
						auto_backup_window = "SLOT_TWO"
						recovery_window_in_days = 10
					}
					kms_key_id = "${var.kms_key_id}"
					vault_id = "${var.vault_id}"
				}
			}
			db_system_options {
				storage_management = "LVM"
			}
			defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
			freeform_tags = {"Department" = "Finance"}
			nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
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

	DatabasePluggableDatabaseResourceDependencies = AvailabilityDomainConfig + ResourcePluggableDatabaseBaseConfig +
		DefinedTagsDependencies + dbSystemForPluggableDbRepresentation

	DatabasePluggableDatabaseResourceDependeciesExacc = DatabaseExaccPluggableDatabaseResourceDependencies
)

// issue-routing-tag: database/default
func TestDatabasePluggableDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabasePluggableDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	resourceName := "oci_database_pluggable_database.test_pluggable_database"
	localCloneResourceName := "oci_database_pluggable_database.test_pluggable_databases_local_clone"
	//remoteCloneResourceName := "oci_database_pluggable_database.test_pluggable_databases_remote_clone"
	//relocateResourceName := "oci_database_pluggable_database.test_pluggable_databases_relocate"
	//refreshableCloneResourceName := "oci_database_pluggable_database.test_pluggable_databases_refreshable_clone"
	datasourceName := "data.oci_database_pluggable_databases.test_pluggable_databases"
	singularDatasourceName := "data.oci_database_pluggable_database.test_pluggable_database"

	var resId, resId2, compId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+kmsKeyIdVariableStr+kmsKeyVersionIdVariableStr+vaultIdVariableStr+DatabasePluggableDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRepresentation), "database", "pluggableDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabasePluggableDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				//resource.TestCheckResourceAttr(resourceName, "container_database_admin_password", "containerDatabaseAdminPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),

				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),

				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
				//resource.TestCheckResourceAttr(resourceName, "should_create_pdb_backup", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_pdb_admin_account_be_locked", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					compId = "oci_database_pluggable_database:" + resId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Update, DatabasePluggableDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				//resource.TestCheckResourceAttr(resourceName, "container_database_admin_password", "containerDatabaseAdminPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),

				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),

				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
				//resource.TestCheckResourceAttr(resourceName, "should_create_pdb_backup", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_pdb_admin_account_be_locked", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				//resource.TestCheckResourceAttr(resourceName, "kms_key_version_id", kmsKeyId),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// delete before local clone
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies,
		},

		// verify local clone
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_local_clone", acctest.Optional, acctest.Create, DatabasePluggableDatabaseLocalCloneRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(localCloneResourceName, "pdb_name", "LocalClonePdb"),
				resource.TestCheckResourceAttrSet(localCloneResourceName, "id"),
				resource.TestCheckResourceAttr(localCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
			),
		},

		//// delete before remote clone
		//{
		//	Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies,
		//},
		//
		//// verify remote clone
		//{
		//	Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_remote_clone", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRemoteCloneRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(remoteCloneResourceName, "pdb_name", "RemoteClonePdb"),
		//		resource.TestCheckResourceAttrSet(remoteCloneResourceName, "id"),
		//		resource.TestCheckResourceAttr(remoteCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
		//	),
		//},
		//
		//// delete before relocate
		//{
		//	Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies,
		//},
		//
		//// verify relocate
		//{
		//	Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_relocate", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRelocateRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(relocateResourceName, "pdb_name", "RelocatePdb"),
		//		resource.TestCheckResourceAttrSet(relocateResourceName, "id"),
		//		resource.TestCheckResourceAttr(relocateResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
		//	),
		//},
		//
		//// delete before refreshable clone
		//{
		//	Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies,
		//},
		//
		//// verify create refreshable clone
		//{
		//	Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_refreshable_clone", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRefreshableCloneRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_name", "RefreshablePdb"),
		//		resource.TestCheckResourceAttrSet(refreshableCloneResourceName, "id"),
		//		resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
		//	),
		//},
		//
		//// verify refresh
		//{
		//	Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_refreshable_clone", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabasePluggableDatabaseRefreshableCloneRepresentation, map[string]interface{}{
		//			"refresh_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`},
		//		})),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_name", "RefreshablePdb"),
		//		resource.TestCheckResourceAttrSet(refreshableCloneResourceName, "id"),
		//		resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
		//	),
		//},
		//
		//// verify convert to regular
		//{
		//	Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
		//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_refreshable_clone", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabasePluggableDatabaseRefreshableCloneRepresentation, map[string]interface{}{
		//			"convert_to_regular_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`},
		//		})),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_name", "RefreshablePdb"),
		//		resource.TestCheckResourceAttrSet(refreshableCloneResourceName, "id"),
		//		resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
		//	),
		//},

		// verify datasource
		{
			Config: config + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_pluggable_databases", "test_pluggable_databases", acctest.Optional, acctest.Update, DatabaseDatabasePluggableDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Update, DatabasePluggableDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.open_mode"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Update, DatabaseDatabasePluggableDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabasePluggableDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "open_mode"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pdb_name", "SalesPdb"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabasePluggableDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				//"container_database_admin_password",
				"pdb_admin_password",
				//"pdb_creation_type_details",
				//"should_create_pdb_backup",
				"should_pdb_admin_account_be_locked",
				"tde_wallet_password",
				"rotate_key_trigger",
				"kms_key_version_id",
			},
			ResourceName: resourceName,
		},
	})
}

//func TestDatabasePluggableDatabaseResourceExacc_basic(t *testing.T) {
//	httpreplay.SetScenario("TestDatabasePluggableDatabaseResourceExacc_basic")
//	defer httpreplay.SaveScenario()
//
//	config := acctest.ProviderTestConfig()
//
//	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
//	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
//
//	resourceName := "oci_database_pluggable_database.test_pluggable_database"
//	localCloneResourceName := "oci_database_pluggable_database.test_pluggable_databases_local_clone"
//	remoteCloneResourceName := "oci_database_pluggable_database.test_pluggable_databases_remote_clone"
//	relocateResourceName := "oci_database_pluggable_database.test_pluggable_databases_relocate"
//	refreshableCloneResourceName := "oci_database_pluggable_database.test_pluggable_databases_refreshable_clone"
//	//datasourceName := "data.oci_database_pluggable_databases.test_pluggable_databases"
//	//singularDatasourceName := "data.oci_database_pluggable_database.test_pluggable_database"
//
//	var resId, resId2, compId string
//	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
//	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabasePluggableDatabaseResourceDependeciesExacc+
//		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRepresentation), "database", "pluggableDatabase", t)
//
//	acctest.ResourceTest(t, testAccCheckDatabasePluggableDatabaseDestroy, []resource.TestStep{
//		// verify Create
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
//				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
//				resource.TestCheckResourceAttr(resourceName, "tde_wallet_password", "BEstrO0ng_#11"),
//
//				func(s *terraform.State) (err error) {
//					resId, err = acctest.FromInstanceState(s, resourceName, "id")
//					return err
//				},
//			),
//		},
//
//		// delete before next Create
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc,
//		},
//		// verify Create with optionals
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRepresentation),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
//				//resource.TestCheckResourceAttr(resourceName, "container_database_admin_password", "containerDatabaseAdminPassword"),
//				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
//				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
//				resource.TestCheckResourceAttrSet(resourceName, "id"),
//				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
//
//				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//
//				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
//				//resource.TestCheckResourceAttr(resourceName, "should_create_pdb_backup", "true"),
//				resource.TestCheckResourceAttr(resourceName, "should_pdb_admin_account_be_locked", "false"),
//				resource.TestCheckResourceAttrSet(resourceName, "state"),
//				resource.TestCheckResourceAttr(resourceName, "tde_wallet_password", "BEstrO0ng_#11"),
//				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
//
//				func(s *terraform.State) (err error) {
//					resId, err = acctest.FromInstanceState(s, resourceName, "id")
//					compId = "oci_database_pluggable_database:" + resId
//					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
//						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compId, &compartmentId, resourceName); errExport != nil {
//							return errExport
//						}
//					}
//					return err
//				},
//			),
//		},
//
//		// verify updates to updatable parameters
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Update, DatabasePluggableDatabaseRepresentation),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
//				//resource.TestCheckResourceAttr(resourceName, "container_database_admin_password", "containerDatabaseAdminPassword"),
//				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
//				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
//				resource.TestCheckResourceAttrSet(resourceName, "id"),
//				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
//
//				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//
//				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
//				//resource.TestCheckResourceAttr(resourceName, "should_create_pdb_backup", "true"),
//				resource.TestCheckResourceAttr(resourceName, "should_pdb_admin_account_be_locked", "false"),
//				resource.TestCheckResourceAttrSet(resourceName, "state"),
//				resource.TestCheckResourceAttr(resourceName, "tde_wallet_password", "BEstrO0ng_#11"),
//				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
//
//				func(s *terraform.State) (err error) {
//					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
//					if resId != resId2 {
//						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
//					}
//					return err
//				},
//			),
//		},
//
//		// delete before local clone
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc,
//		},
//
//		// verify local clone
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_local_clone", acctest.Optional, acctest.Create, DatabasePluggableDatabaseLocalCloneRepresentation),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttr(localCloneResourceName, "pdb_name", "LocalClonePdb"),
//				resource.TestCheckResourceAttrSet(localCloneResourceName, "id"),
//				resource.TestCheckResourceAttr(localCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//			),
//		},
//
//		// delete before remote clone
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc,
//		},
//
//		// verify remote clone
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_remote_clone", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRemoteCloneRepresentation),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttr(remoteCloneResourceName, "pdb_name", "RemoteClonePdb"),
//				resource.TestCheckResourceAttrSet(remoteCloneResourceName, "id"),
//				resource.TestCheckResourceAttr(remoteCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//			),
//		},
//
//		// delete before relocate
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc,
//		},
//
//		// verify relocate
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_relocate", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRelocateRepresentation),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttr(relocateResourceName, "pdb_name", "RelocatePdb"),
//				resource.TestCheckResourceAttrSet(relocateResourceName, "id"),
//				resource.TestCheckResourceAttr(relocateResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//			),
//		},
//
//		// delete before refreshable clone
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc,
//		},
//
//		// verify create refreshable clone
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_refreshable_clone", acctest.Optional, acctest.Create, DatabasePluggableDatabaseRefreshableCloneRepresentation),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_name", "RefreshablePdb"),
//				resource.TestCheckResourceAttrSet(refreshableCloneResourceName, "id"),
//				resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//			),
//		},
//
//		// verify refresh
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_refreshable_clone", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabasePluggableDatabaseRefreshableCloneRepresentation, map[string]interface{}{
//					"refresh_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`},
//				})),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_name", "RefreshablePdb"),
//				resource.TestCheckResourceAttrSet(refreshableCloneResourceName, "id"),
//				resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//			),
//		},
//
//		// verify convert to regular
//		{
//			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseResourceDependeciesExacc +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableDatabaseRepresentation) +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_databases_refreshable_clone", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabasePluggableDatabaseRefreshableCloneRepresentation, map[string]interface{}{
//					"convert_to_regular_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`},
//				})),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_name", "RefreshablePdb"),
//				resource.TestCheckResourceAttrSet(refreshableCloneResourceName, "id"),
//				resource.TestCheckResourceAttr(refreshableCloneResourceName, "pdb_admin_password", "BEstrO0ng_#11"),
//			),
//		},
//	})
//}

func testAccCheckDatabasePluggableDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_pluggable_database" {
			noResourceFound = false
			request := oci_database.GetPluggableDatabaseRequest{}

			tmp := rs.Primary.ID
			request.PluggableDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetPluggableDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.PluggableDatabaseLifecycleStateTerminated): true,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabasePluggableDatabase") {
		resource.AddTestSweepers("DatabasePluggableDatabase", &resource.Sweeper{
			Name:         "DatabasePluggableDatabase",
			Dependencies: acctest.DependencyGraph["pluggableDatabase"],
			F:            sweepDatabasePluggableDatabaseResource,
		})
	}
}

func sweepDatabasePluggableDatabaseResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	pluggableDatabaseIds, err := getDatabasePluggableDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, pluggableDatabaseId := range pluggableDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[pluggableDatabaseId]; !ok {
			deletePluggableDatabaseRequest := oci_database.DeletePluggableDatabaseRequest{}

			deletePluggableDatabaseRequest.PluggableDatabaseId = &pluggableDatabaseId

			deletePluggableDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeletePluggableDatabase(context.Background(), deletePluggableDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting PluggableDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", pluggableDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &pluggableDatabaseId, DatabasePluggableDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				DatabasePluggableDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabasePluggableDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PluggableDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listPluggableDatabasesRequest := oci_database.ListPluggableDatabasesRequest{}
	listPluggableDatabasesRequest.CompartmentId = &compartmentId
	listPluggableDatabasesRequest.LifecycleState = oci_database.PluggableDatabaseSummaryLifecycleStateAvailable
	listPluggableDatabasesResponse, err := databaseClient.ListPluggableDatabases(context.Background(), listPluggableDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PluggableDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pluggableDatabase := range listPluggableDatabasesResponse.Items {
		id := *pluggableDatabase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PluggableDatabaseId", id)
	}
	return resourceIds, nil
}

func DatabasePluggableDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if pluggableDatabaseResponse, ok := response.Response.(oci_database.GetPluggableDatabaseResponse); ok {
		return pluggableDatabaseResponse.LifecycleState != oci_database.PluggableDatabaseLifecycleStateTerminated
	}
	return false
}

func DatabasePluggableDatabaseSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetPluggableDatabase(context.Background(), oci_database.GetPluggableDatabaseRequest{
		PluggableDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
