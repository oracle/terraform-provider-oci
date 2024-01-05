// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PluggableDatabasePluggabledatabasemanagementsManagementRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_pluggabledatabasemanagements_management", "test_pluggable_database_pluggabledatabasemanagements_management", acctest.Required, acctest.Create, DatabasePluggableDatabasePluggabledatabasemanagementsManagementRepresentation)

	DatabasePluggableDatabasePluggabledatabasemanagementsManagementRepresentation = map[string]interface{}{
		"credential_details":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabasePluggableDatabasePluggabledatabasemanagementsManagementCredentialDetailsRepresentation},
		"pluggable_database_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
		"private_end_point_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
		"service_name":                       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.db_unique_name}.${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"enable_pluggabledatabasemanagement": acctest.Representation{RepType: acctest.Required, Create: `true` /*Update: `false`*/},
		"port":                               acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"protocol":                           acctest.Representation{RepType: acctest.Optional, Create: `TCP`, Update: `TCPS`},
		"role":                               acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`, Update: `SYSDBA`},
		"ssl_secret_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${var.sslSecretId}`, Update: `${var.sslSecretId}`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabasePluggableIgnoreChangesLBRepresentation},
	}

	DatabaseManagementRepresentation = map[string]interface{}{
		"database_id":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.id}`},
		"management_type":      acctest.Representation{RepType: acctest.Required, Create: `ADVANCED`},
		"private_end_point_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
		"service_name":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.db_unique_name}.${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"credentialdetails":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabasePluggableDatabasePluggabledatabasemanagementsManagementCredentialDetailsRepresentation},
		"enable_management":    acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"protocol":             acctest.Representation{RepType: acctest.Optional, Create: `TCP`, Update: `TCPS`},
		"port":                 acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"role":                 acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`, Update: `SYSDBA`},
		"ssl_secret_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.sslSecretId}`, Update: `${var.sslSecretId}`},
	}

	DatabasePluggableDatabasePluggabledatabasemanagementsManagementCredentialDetailsRepresentation = map[string]interface{}{
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sslSecretId}`},
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `sys`},
	}

	DatabasePluggableManagementPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `TFPEforTCPS`, Update: `name2`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	DatabasePluggableRepresentation = map[string]interface{}{
		"container_database_id":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_database.t.id}`},
		"pdb_admin_password":                 acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pdb_name":                           acctest.Representation{RepType: acctest.Required, Create: `SalesPdb`},
		"tde_wallet_password":                acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"should_pdb_admin_account_be_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabasePluggableIgnoreChangesLBRepresentation},
		"depends_on":                         acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.t"}},
	}

	DatabasePluggableIgnoreChangesLBRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`, `pluggable_database_management_config`, `connection_strings`}},
	}

	DatabasePluggableDatabaseResourceBase = acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabasePluggableRepresentation)

	DatabasePluggableManagementPrivateEndpointConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, DatabasePluggableManagementPrivateEndpointRepresentation)

	DatabasePluggableManagementResourceDependenciesBase = AvailabilityDomainConfig + DefinedTagsDependencies + CoreVcnResourceConfig +
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

	DatabasePluggableDbSystemResourceDependencies = `
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
			cpu_core_count = "2"
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

	DatabasePluggableManagementResourceDependencies = DatabasePluggableManagementResourceDependenciesBase + DatabasePluggableManagementPrivateEndpointConfig + DatabasePluggableDbSystemResourceDependencies + DatabasePluggableDatabaseResourceBase
)

// issue-routing-tag: database/default
func TestDatabasePluggableDatabasePluggabledatabasemanagementsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabasePluggableDatabasePluggabledatabasemanagementsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sslSecretId := utils.GetEnvSettingWithDefault("ssl_secret_id", "test_secret_id")
	sslSecretIdVariableStr := fmt.Sprintf("variable \"sslSecretId\" { default = \"%s\" }\n", sslSecretId)

	resourceName := "oci_database_pluggable_database_pluggabledatabasemanagements_management.test_pluggable_database_pluggabledatabasemanagements_management"
	parentResourceName := "oci_database_pluggable_database_pluggabledatabasemanagements_management.test_pluggable_database_pluggabledatabasemanagements_management"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + DatabasePluggableManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, DatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_pluggabledatabasemanagements_management", "test_pluggable_database_pluggabledatabasemanagements_management", acctest.Required, acctest.Create, DatabasePluggableDatabasePluggabledatabasemanagementsManagementRepresentation),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + DatabasePluggableManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, DatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_pluggabledatabasemanagements_management", "test_pluggable_database_pluggabledatabasemanagements_management", acctest.Required, acctest.Create, DatabasePluggableDatabasePluggabledatabasemanagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.password_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.user_name"),
				resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_end_point_id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				resource.TestCheckResourceAttr("oci_database_pluggable_database.test_pluggable_database", "pluggable_database_management_config.0.management_status", "ENABLED"),
				resource.TestCheckResourceAttr(parentResourceName, "enable_pluggabledatabasemanagement", "true"),
			),
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + DatabasePluggableManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Create, DatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_pluggabledatabasemanagements_management", "test_pluggable_database_pluggabledatabasemanagements_management", acctest.Optional, acctest.Create, DatabasePluggableDatabasePluggabledatabasemanagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.password_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.user_name"),
				resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_end_point_id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				resource.TestCheckResourceAttr("oci_database_pluggable_database.test_pluggable_database", "pluggable_database_management_config.0.management_status", "ENABLED"),
			),
		},
		// verify update/disable fields
		{
			Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + DatabasePluggableManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Optional, acctest.Create, DatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_pluggabledatabasemanagements_management", "test_pluggable_database_pluggabledatabasemanagements_management", acctest.Optional, acctest.Update, DatabasePluggableDatabasePluggabledatabasemanagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.password_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.user_name"),
				resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_end_point_id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				resource.TestCheckResourceAttr("oci_database_pluggable_database.test_pluggable_database", "pluggable_database_management_config.0.management_status", "ENABLED"),
			),
		},
		// verify update/disable status
		{
			Config: config + compartmentIdVariableStr + sslSecretIdVariableStr + DatabasePluggableManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_database_management", "test_database_cloud_database_management", acctest.Required, acctest.Create, DatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_pluggabledatabasemanagements_management", "test_pluggable_database_pluggabledatabasemanagements_management", acctest.Optional, acctest.Update, DatabasePluggableDatabasePluggabledatabasemanagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_pluggabledatabasemanagement", "true"),
			),
		},
	})
}
