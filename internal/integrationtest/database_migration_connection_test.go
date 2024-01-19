// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	_ "strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseMigrationConnectionRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, connectionRepresentationTarget)

	goldenGateDbSystemOption = map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Required, Create: `LVM`},
	}

	goldenGateDbSystemDbHomeRepresentation = map[string]interface{}{
		"database":   acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDatabaseRepresentation},
		"db_version": acctest.Representation{RepType: acctest.Required, Create: `21.3.0.0`},
	}

	goldenGateDatabaseRepresentation = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `myDB`},
		"pdb_name":       acctest.Representation{RepType: acctest.Required, Create: `pdbName`},
	}

	kmsKeyId            = utils.GetEnvSettingWithBlankDefault("kms_key_id")
	KmsKeyIdVariableStr = fmt.Sprintf("\nvariable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	DatabaseHomeConfig = `
	data "oci_database_db_homes" "t" {
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.t.id}"
}`

	DatabaseData = `
	data "oci_database_databases" "t" {
	compartment_id = "${var.compartment_id}"
	db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
}`

	ConnectionResourceConfigTarget = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, connectionRepresentationTarget)

	ConnectionResourceConfigSource = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, DatabaseMigrationConnectionRepresentation)

	connectionSingularDataSourceRepresentationCon = map[string]interface{}{
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_connection.test_connection.id}`},
	}

	connectionDataSourceRepresentationCon = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionDataSourceFilterRepresentationCon},
	}
	connectionDataSourceFilterRepresentationCon = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_connection.test_connection.id}`}},
	}

	DatabaseMigrationConnectionRepresentation = map[string]interface{}{
		"admin_credentials":       acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":           acctest.Representation{RepType: acctest.Required, Create: `USER_MANAGED_OCI`},
		"vault_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"certificate_tdn":         acctest.Representation{RepType: acctest.Optional, Create: `certificateTdn`, Update: `certificateTdn2`},
		"connect_descriptor":      acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionConnectDescriptorRepresentationMIG},
		"database_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.database_source}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"private_endpoint":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: connectionPrivateEndpointRepresentation},
		"replication_credentials": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationConnectionReplicationCredentialsRepresentation},
		"ssh_details":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: connectionSshDetailsRepresentation},
		"tls_keystore":            acctest.Representation{RepType: acctest.Optional, Create: `tlsKeystore`, Update: `tlsKeystore2`},
		"tls_wallet":              acctest.Representation{RepType: acctest.Optional, Create: `tlsWallet`, Update: `tlsWallet2`},
	}

	connectionRepresentationCon = map[string]interface{}{
		"admin_credentials":  acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":      acctest.Representation{RepType: acctest.Required, Create: `USER_MANAGED_OCI`},
		"vault_details":      acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"certificate_tdn":    acctest.Representation{RepType: acctest.Optional, Create: `certificateTdn`, Update: `certificateTdn2`},
		"connect_descriptor": acctest.RepresentationGroup{RepType: acctest.Optional, Group: connectionConnectDescriptorRepresentationMIG},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `TF_display_test_create`, Update: `TF_display_test_create`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":            acctest.Representation{RepType: acctest.Required, Create: []string{`${var.nsg_id}`}},
		"private_endpoint":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: connectionPrivateEndpointRepresentation},
		"ssh_details":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: connectionSshDetailsRepresentation},
		"tls_keystore":       acctest.Representation{RepType: acctest.Optional, Create: `tlsKeystore`, Update: `tlsKeystore2`},
		"tls_wallet":         acctest.Representation{RepType: acctest.Optional, Create: `tlsWallet`, Update: `tlsWallet2`},
	}

	connectionRepresentationTarget = map[string]interface{}{
		"admin_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":     acctest.Representation{RepType: acctest.Required, Create: `AUTONOMOUS`},
		"vault_details":     acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"database_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.database_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `TF_display_test_create11`, Update: `TF_display_test_update`},
	}

	connectionRepresentationTargetOpc = map[string]interface{}{
		"admin_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":     acctest.Representation{RepType: acctest.Required, Create: `AUTONOMOUS`},
		"nsg_ids":           acctest.Representation{RepType: acctest.Required, Create: []string{`${var.nsg_id}`}},
		"private_endpoint":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: connectionPrivateEndpointRepresentation},
		"vault_details":     acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"database_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.database_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `TF_display_test_create`, Update: `TF_display_test_update`},
	}

	connectionRepresentationUserManagedOciTarget = map[string]interface{}{
		"admin_credentials":  acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":      acctest.Representation{RepType: acctest.Required, Create: `USER_MANAGED_OCI`},
		"vault_details":      acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"connect_descriptor": acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionConnectDescriptorRepresentationMIG},
		"database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.database_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `TF_display_test_create_target`, Update: `TF_tgt_display_test_update_target`},
	}

	connectionRepresentationSource = map[string]interface{}{
		"admin_credentials":        acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":            acctest.Representation{RepType: acctest.Required, Create: `MANUAL`},
		"manual_database_sub_type": acctest.Representation{RepType: acctest.Optional, Create: `RDS_ORACLE`},
		"vault_details":            acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"connect_descriptor":       acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionConnectDescriptorRepresentationMIG},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `TF_display_test_create_source`, Update: `TF_display_test_update_source`},
		"private_endpoint":         acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionPrivateEndpointRepresentation},
	}

	connectionRepresentationNoSshSource = map[string]interface{}{
		"admin_credentials":  acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":      acctest.Representation{RepType: acctest.Required, Create: `MANUAL`},
		"vault_details":      acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"connect_descriptor": acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionConnectDescriptorRepresentationMIG},
		"database_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.database_container_source_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `TF_display_test_create_source`, Update: `TF_display_test_update_source`},
	}

	connectionAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `Cr3dential_23#`, Update: `Cr3dential_23#`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `admin`, Update: `admin`},
	}

	connectionAdminCredentialsRepresentationUPDATE = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `DMS-pswd-2023#`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `admin`},
	}

	connectionVaultDetailsRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"key_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.kms_vault_id}`},
	}

	connectionVaultDetailsRepresentationUPDATE = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"key_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.kms_vault_id}`},
	}

	connectionConnectDescriptorRepresentation = map[string]interface{}{
		"connect_string":        acctest.Representation{RepType: acctest.Optional, Create: `(description=(address=(port=1521)(host=10.0.0.220))(connect_data=(service_name=DBSOURCE_pdb1.sub04181535190.acommonvcn.oraclevcn.com)))`, Update: `(description=(address=(port=1521)(host=10.0.0.220))(connect_data=(service_name=DBSOURCE_pdb1.sub04181535190.acommonvcn.oraclevcn.com)))`},
		"database_service_name": acctest.Representation{RepType: acctest.Optional, Create: `database_migration`},
		"host":                  acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.220`, Update: `10.0.0.220`},
		"port":                  acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1521`},
	}
	connectionPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.vcn_id}`},
	}
	connectionConnectDescriptorRepresentationUpdate = map[string]interface{}{
		"database_service_name": acctest.Representation{RepType: acctest.Required, Create: `DBSOURCE_phx1vk.sub04102006390.acommonvcn.oraclevcn.com`},
		"host":                  acctest.Representation{RepType: acctest.Required, Create: `10.0.0.119`, Update: `10.0.0.119`},
		"port":                  acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1521`},
	}

	connectionConnectDescriptorRepresentationMIG = map[string]interface{}{
		"connect_string": acctest.Representation{RepType: acctest.Required, Create: `(description=(address=(port=1521)(host=10.0.0.119))(connect_data=(service_name=DBSOURCE_phx1vk.sub04102006390.acommonvcn.oraclevcn.com)))`, Update: `(description=(address=(port=1521)(host=10.0.0.119))(connect_data=(service_name=DBSOURCE_phx1vk.sub04102006390.acommonvcn.oraclevcn.com)))`},
	}
	connectionConnectDescriptorRepresentationPDB = map[string]interface{}{
		"connect_string": acctest.Representation{RepType: acctest.Required, Create: `(description=(address=(port=1521)(host=10.0.0.119))(connect_data=(service_name=DBSOURCE_pdb1.sub04102006390.acommonvcn.oraclevcn.com)))`, Update: `(description=(address=(port=1521)(host=10.0.0.119))(connect_data=(service_name=DBSOURCE_pdb1.sub04102006390.acommonvcn.oraclevcn.com)))`},
	}

	DatabaseMigrationConnectionReplicationCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Optional, Create: `DMS-pswd-2023#`, Update: `DMS-pswd-2023#`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `admin`, Update: `admin2`},
	}
	DatabaseMigrationConnectionSshDetailsRepresentation = map[string]interface{}{
		"host":          acctest.Representation{RepType: acctest.Required, Create: `host`, Update: `host2`},
		"sshkey":        acctest.Representation{RepType: acctest.Required, Create: `sshkey`, Update: `sshkey2`},
		"user":          acctest.Representation{RepType: acctest.Required, Create: `user`, Update: `user2`},
		"sudo_location": acctest.Representation{RepType: acctest.Optional, Create: `sudoLocation`, Update: `sudoLocation2`},
	}

	connectionSshDetailsRepresentation = map[string]interface{}{
		"host":          acctest.Representation{RepType: acctest.Required, Create: `10.0.0.119`, Update: `10.0.0.119`},
		"sshkey":        acctest.Representation{RepType: acctest.Required, Create: `${var.ssh_key}`},
		"user":          acctest.Representation{RepType: acctest.Required, Create: `opc`, Update: `opc2`},
		"sudo_location": acctest.Representation{RepType: acctest.Required, Create: `/usr/bin/sudo`, Update: `/usr/bin/sudo`},
	}

	databaseRepresentationConnectionResource = map[string]interface{}{
		"database":   acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentationConnectionResource},
		"db_version": acctest.Representation{RepType: acctest.Required, Create: `21.1.0.0`},
	}

	databaseDatabaseRepresentationConnectionResource = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `myDB`},
		"pdb_name":       acctest.Representation{RepType: acctest.Required, Create: `pdbName`},
	}

	SubnetData = `
		data "oci_core_subnet" "test_subnet" {
		subnet_id = "${var.subnet_id}"
	}`
	SubnetDataDomainOutput = `
		output "oci_core_subnet_test_subnetSource_subnet_domain_name" {
	    value = "${oci_core_subnet.test_subnet.subnet_domain_name}"
	}`

	SubnetDataIDOutput = `
		output "oci_core_subnet_test_subnetSource_id" {
	    value = "${oci_core_subnet.test_subnet.id}"
	}`
	SubnetDataDNSOutput = `
		output "oci_core_subnet_test_subnet_DNS" {
	    value = "${oci_core_subnet.test_subnet.dns_label}"
	}`

	VCNDataDNSOutput = `
		output "oci_core_vcn_test_vcn_DNS" {
	    value = "${oci_core_vcn.test_vcn.dns_label}"
	}`
	VCNDataDomainNameOutput = `
		output "oci_core_vcn_test_vcn_domain_name" {
	    value = "${oci_core_vcn.test_vcn.vcn_domain_name}"
	}`

	DatabaseDataA = `
		data "oci_database_autonomous_database" "t" {
		compartment_id = "${var.compartment_id}"
		db_home_id = "${data.oci_database_autonomous_db_homes.t.db_homes.0.id}"
	}`
	DatabaseHomeConfigA = `
		data "oci_database_autonomous_db_homes" "t" {
		compartment_id = "${var.compartment_id}"
		db_system_id = "${oci_database_db_system.t.id}"
	}`
	AutonomousDatabaseResourceDependenciesCON = acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))

	AutonomousDatabaseResourceDependenciesCONSOURCE = acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions_source", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions_source", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))

	goldenGateDbSystemRepresentationSOURCE = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: "efde:phx-ad-1"},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDbSystemDbHomeRepresentation},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myDB`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfGGmyDB`},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `sub10031523100.vcnabmartin.oraclevcn.com`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"db_system_options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: goldenGateDbSystemOption},
		"private_ip":              acctest.Representation{RepType: acctest.Required, Create: `10.0.0.125`},
	}
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sshKey := utils.GetEnvSettingWithBlankDefault("ssh_key")
	sshKeyIdStr := fmt.Sprintf("variable \"ssh_key\" {\n type = \"string\"\n default = <<EOF\n%s\nEOF \n}\n", sshKey)

	databaseId := utils.GetEnvSettingWithBlankDefault("database_id")
	databaseIdVariableStr := fmt.Sprintf("variable \"database_id\" { default = \"%s\" }\n", databaseId)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	KmsKeyIdVariableStr := fmt.Sprintf("\nvariable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsVaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_id")
	KmsVaultIdVariableStr := fmt.Sprintf("\nvariable \"kms_vault_id\" { default = \"%s\" }\n", kmsVaultId)

	vcnId := utils.GetEnvSettingWithBlankDefault("vcn_id")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	databaseSourceId := utils.GetEnvSettingWithBlankDefault("database_source")
	databaseSourceStr := fmt.Sprintf("variable \"database_source\" { default = \"%s\" }\n", databaseSourceId)

	nsgId := utils.GetEnvSettingWithBlankDefault("nsg_id")
	nsgIdStr := fmt.Sprintf("variable \"nsg_id\" { default = \"%s\" }\n", nsgId)

	compartmentIdU := utils.GetEnvSettingWithBlankDefault("compartment_id_for_update")
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_migration_connection.test_connection"
	datasourceName := "data.oci_database_migration_connections.test_connections"
	singularDatasourceName := "data.oci_database_migration_connection.test_connection"
	resourceNameRDS := "oci_database_migration_connection.test_connection_rds"

	var resId, resId2, resId3 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+databaseSourceStr+subnetStr+vcnIdVariableStr+KmsVaultIdVariableStr+KmsKeyIdVariableStr+sshKeyIdStr+databaseIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create, DatabaseMigrationConnectionRepresentation)+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Create, connectionRepresentationSource), "databasemigration", "connection", t)

	acctest.ResourceTest(t, testAccCheckDatabaseMigrationConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + databaseIdVariableStr + KmsKeyIdVariableStr + KmsVaultIdVariableStr + vcnIdVariableStr + subnetStr + databaseSourceStr + nsgIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, connectionRepresentationTarget),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_type", "AUTONOMOUS"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + databaseSourceStr + subnetStr + vcnIdVariableStr + KmsVaultIdVariableStr + KmsKeyIdVariableStr + sshKeyIdStr + databaseIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + databaseSourceStr + subnetStr + vcnIdVariableStr + KmsVaultIdVariableStr + KmsKeyIdVariableStr + sshKeyIdStr + databaseIdVariableStr + nsgIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create, DatabaseMigrationConnectionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Create, connectionRepresentationSource),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.password", "Cr3dential_23#"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_type", "USER_MANAGED_OCI"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.password", "DMS-pswd-2023#"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.host", "10.0.0.119"),
				resource.TestCheckResourceAttrSet(resourceName, "ssh_details.0.sshkey"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.sudo_location", "/usr/bin/sudo"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.user", "opc"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.0.password", "Cr3dential_23#"),
				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceNameRDS, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameRDS, "database_type", "MANUAL"),
				resource.TestCheckResourceAttr(resourceNameRDS, "manual_database_sub_type", "RDS_ORACLE"),
				resource.TestCheckResourceAttr(resourceNameRDS, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "connect_descriptor.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRDS, "display_name", "TF_display_test_create_source"),
				resource.TestCheckResourceAttr(resourceNameRDS, "private_endpoint.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "private_endpoint.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "private_endpoint.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					resId2, err = acctest.FromInstanceState(s, resourceNameRDS, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId3, &compartmentId, resourceNameRDS); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + databaseIdVariableStr + KmsKeyIdVariableStr + KmsVaultIdVariableStr + vcnIdVariableStr + subnetStr + sshKeyIdStr + databaseSourceStr + nsgIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseMigrationConnectionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(connectionRepresentationSource, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.password", "Cr3dential_23#"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_type", "USER_MANAGED_OCI"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.password", "DMS-pswd-2023#"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.host", "10.0.0.119"),
				resource.TestCheckResourceAttrSet(resourceName, "ssh_details.0.sshkey"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.sudo_location", "/usr/bin/sudo"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.user", "opc"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.0.password", "Cr3dential_23#"),
				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceNameRDS, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceNameRDS, "database_type", "MANUAL"),
				resource.TestCheckResourceAttr(resourceNameRDS, "manual_database_sub_type", "RDS_ORACLE"),
				resource.TestCheckResourceAttr(resourceNameRDS, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "connect_descriptor.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRDS, "display_name", "TF_display_test_create_source"),
				resource.TestCheckResourceAttr(resourceNameRDS, "private_endpoint.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "private_endpoint.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "private_endpoint.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId3 {
						return fmt.Errorf("update to the compartment: resource %s recreated when it was supposed to be updated", resourceName)
					}
					resId3, err = acctest.FromInstanceState(s, resourceNameRDS, "id")
					if resId2 != resId3 {
						return fmt.Errorf("update to the compartment: resource %s recreated when it was supposed to be updated", resourceNameRDS)
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + databaseSourceStr + subnetStr + vcnIdVariableStr + KmsVaultIdVariableStr + KmsKeyIdVariableStr + sshKeyIdStr + databaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, DatabaseMigrationConnectionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Update, connectionRepresentationSource),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.password", "Cr3dential_23#"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_type", "USER_MANAGED_OCI"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.password", "DMS-pswd-2023#"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.username", "admin2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.host", "10.0.0.119"),
				resource.TestCheckResourceAttrSet(resourceName, "ssh_details.0.sshkey"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.sudo_location", "/usr/bin/sudo"),
				resource.TestCheckResourceAttr(resourceName, "ssh_details.0.user", "opc2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.0.password", "Cr3dential_23#"),
				resource.TestCheckResourceAttr(resourceNameRDS, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceNameRDS, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameRDS, "database_type", "MANUAL"),
				resource.TestCheckResourceAttr(resourceNameRDS, "manual_database_sub_type", "RDS_ORACLE"),
				resource.TestCheckResourceAttr(resourceNameRDS, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "connect_descriptor.#", "1"),
				resource.TestCheckResourceAttr(resourceNameRDS, "display_name", "TF_display_test_update_source"),
				resource.TestCheckResourceAttr(resourceNameRDS, "private_endpoint.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "private_endpoint.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "private_endpoint.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId3 {
						return fmt.Errorf("updates to updatable parameters: resource %s recreated when it was supposed to be updated", resourceName)
					}
					resId3, err = acctest.FromInstanceState(s, resourceNameRDS, "id")
					if resId2 != resId3 {
						return fmt.Errorf("updates to updatable parameters: resource %s recreated when it was supposed to be updated", resourceNameRDS)
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_connections", "test_connections", acctest.Optional, acctest.Update, connectionDataSourceRepresentationCon) +
				compartmentIdVariableStr + databaseSourceStr + subnetStr + vcnIdVariableStr + KmsVaultIdVariableStr + KmsKeyIdVariableStr + sshKeyIdStr + databaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, connectionRepresentationTarget),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "connection_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, connectionSingularDataSourceRepresentationCon) +
				compartmentIdVariableStr + databaseIdVariableStr + KmsKeyIdVariableStr + KmsVaultIdVariableStr + vcnIdVariableStr + subnetStr + sshKeyIdStr + databaseSourceStr + nsgIdStr + ConnectionResourceConfigSource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_type", "USER_MANAGED_OCI"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_credentials.0.username", "admin2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_details.0.host", "10.0.0.119"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_details.0.sudo_location", "/usr/bin/sudo"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_details.0.user", "opc2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_details.0.compartment_id", compartmentId),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseMigrationConnectionRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"tls_keystore",
				"tls_wallet",
				"admin_credentials.0.password",
				"replication_credentials.0.password",
				"ssh_details.0.sshkey",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseMigrationConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_connection" {
			noResourceFound = false
			request := oci_database_migration.GetConnectionRequest{}

			tmp := rs.Primary.ID
			request.ConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")

			response, err := client.GetConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_migration.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseMigrationConnection") {
		resource.AddTestSweepers("DatabaseMigrationConnection", &resource.Sweeper{
			Name:         "DatabaseMigrationConnection",
			Dependencies: acctest.DependencyGraph["connection"],
			F:            sweepDatabaseMigrationConnectionResource,
		})
	}
}

func sweepDatabaseMigrationConnectionResource(compartment string) error {
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()
	connectionIds, err := getConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, connectionId := range connectionIds {
		if ok := acctest.SweeperDefaultResourceId[connectionId]; !ok {
			deleteConnectionRequest := oci_database_migration.DeleteConnectionRequest{}

			deleteConnectionRequest.ConnectionId = &connectionId

			deleteConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteConnection(context.Background(), deleteConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting Connection %s %s, It is possible that the resource is already deleted. Please verify manually \n", connectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &connectionId, connectionSweepWaitCondition, time.Duration(3*time.Minute),
				connectionSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()
	listConnectionsRequest := oci_database_migration.ListConnectionsRequest{}
	listConnectionsRequest.CompartmentId = &compartmentId
	listConnectionsRequest.LifecycleState = oci_database_migration.ListConnectionsLifecycleStateActive
	listConnectionsResponse, err := databaseMigrationClient.ListConnections(context.Background(), listConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Connection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, connection := range listConnectionsResponse.Items {
		id := *connection.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConnectionId", id)
	}
	return resourceIds, nil
}

func connectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if connectionResponse, ok := response.Response.(oci_database_migration.GetConnectionResponse); ok {
		return connectionResponse.LifecycleState != oci_database_migration.LifecycleStatesDeleted
	}
	return false
}

func connectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseMigrationClient().GetConnection(context.Background(), oci_database_migration.GetConnectionRequest{
		ConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
