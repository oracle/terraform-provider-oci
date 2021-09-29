// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v48/databasemigration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConnectionResourceConfigTarget = ConnectionResourceDependenciesTarget +
		generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Optional, Update, connectionRepresentationTarget)

	connectionSingularDataSourceRepresentationCon = map[string]interface{}{
		"connection_id": Representation{repType: Required, create: `${oci_database_migration_connection.test_connection.id}`},
	}

	connectionDataSourceRepresentationCon = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, connectionDataSourceFilterRepresentationCon}}
	connectionDataSourceFilterRepresentationCon = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_migration_connection.test_connection.id}`}},
	}

	connectionRepresentationCon = map[string]interface{}{
		"admin_credentials":  RepresentationGroup{Required, connectionAdminCredentialsRepresentation},
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"database_type":      Representation{repType: Required, create: `MANUAL`},
		"vault_details":      RepresentationGroup{Required, connectionVaultDetailsRepresentation},
		"certificate_tdn":    Representation{repType: Optional, create: `certificateTdn`, update: `certificateTdn2`},
		"connect_descriptor": RepresentationGroup{Optional, connectionConnectDescriptorRepresentation},
		"database_id":        Representation{repType: Optional, create: `${data.oci_database_databases.db.databases.0.id}`},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":       Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"private_endpoint":   RepresentationGroup{Optional, connectionPrivateEndpointRepresentation},
		"ssh_details":        RepresentationGroup{Optional, connectionSshDetailsRepresentation},
		"tls_keystore":       Representation{repType: Optional, create: `tlsKeystore`, update: `tlsKeystore2`},
		"tls_wallet":         Representation{repType: Optional, create: `tlsWallet`, update: `tlsWallet2`},
	}

	connectionRepresentationTarget = map[string]interface{}{
		"admin_credentials": RepresentationGroup{Required, connectionAdminCredentialsRepresentation},
		"compartment_id":    Representation{repType: Required, create: `${var.compartment_id}`},
		"database_type":     Representation{repType: Required, create: `AUTONOMOUS`},
		"vault_details":     RepresentationGroup{Required, connectionVaultDetailsRepresentation},
		"database_id":       Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":      Representation{repType: Required, create: `TF_display_test_create`, update: `TF_display_test_update`},
	}

	connectionConnectDescriptorRepresentationMIG = map[string]interface{}{
		"connect_string": Representation{repType: Required, create: `(description=(address=(port=1521)(host=10.0.0.125))(connect_data=(service_name=pdb1120.exadbpriv.exadbvcn.oraclevcn.com)))`, update: `(description=(address=(port=1521)(host=10.0.0.125))(connect_data=(service_name=pdb1120.exadbpriv.exadbvcn.oraclevcn.com)))`},
	}
	connectionRepresentationSource = map[string]interface{}{
		"admin_credentials":  RepresentationGroup{Required, connectionAdminCredentialsRepresentation},
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"database_type":      Representation{repType: Required, create: `MANUAL`},
		"vault_details":      RepresentationGroup{Required, connectionVaultDetailsRepresentation},
		"connect_descriptor": RepresentationGroup{Required, connectionConnectDescriptorRepresentationMIG},
		"database_id":        Representation{repType: Optional, create: `${data.oci_database_databases.t.databases.0.id}`},
		"display_name":       Representation{repType: Required, create: `TF_display_test_create_source`, update: `TF_display_test_update_source`},
		"ssh_details":        RepresentationGroup{Required, connectionSshDetailsRepresentation},
		"private_endpoint":   RepresentationGroup{Required, connectionPrivateEndpointRepresentation},
	}

	connectionAdminCredentialsRepresentation = map[string]interface{}{
		"password": Representation{repType: Required, create: `ORcl##4567890`, update: `ORcl##4567890`},
		"username": Representation{repType: Required, create: `admin`, update: `admin`},
	}

	connectionAdminCredentialsRepresentationUPDATE = map[string]interface{}{
		"password": Representation{repType: Required, create: `ORcl##4567890`},
		"username": Representation{repType: Required, create: `admin`},
	}

	connectionVaultDetailsRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"key_id":         Representation{repType: Required, create: `${var.kms_key_id}`},
		"vault_id":       Representation{repType: Required, create: `${var.kms_vault_id}`},
	}

	connectionVaultDetailsRepresentationUPDATE = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"key_id":         Representation{repType: Required, create: `${var.kms_key_id}`},
		"vault_id":       Representation{repType: Required, create: `${var.kms_vault_id}`},
	}

	connectionConnectDescriptorRepresentation = map[string]interface{}{
		"connect_string":        Representation{repType: Optional, create: `connectString`, update: `connectString2`},
		"database_service_name": Representation{repType: Optional, create: `${oci_core_services.test_services.name}`},
		"host":                  Representation{repType: Optional, create: `host`, update: `host2`},
		"port":                  Representation{repType: Optional, create: `10`, update: `11`},
	}
	connectionPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"subnet_id":      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":         Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
	}
	connectionSshDetailsRepresentation = map[string]interface{}{
		"host":          Representation{repType: Required, create: `10.0.0.125`, update: `10.0.0.125`},
		"sshkey":        Representation{repType: Required, create: `ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`, update: `ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`},
		"user":          Representation{repType: Required, create: `opc`, update: `opc`},
		"sudo_location": Representation{repType: Required, create: `/usr/bin/sudo`, update: `/usr/bin/sudo`},
	}

	databaseRepresentationConnectionResource = map[string]interface{}{
		"database":   RepresentationGroup{Required, databaseDatabaseRepresentationConnectionResource},
		"db_version": Representation{repType: Required, create: `21.1.0.0`},
	}

	databaseDatabaseRepresentationConnectionResource = map[string]interface{}{
		"admin_password": Representation{repType: Required, create: `BEstrO0ng_#11`},
		"db_name":        Representation{repType: Required, create: `myDB`},
		"pdb_name":       Representation{repType: Required, create: `pdbName`},
	}

	SubnetData = `
	data "oci_core_subnet" "test_subnet" {
    subnet_id = "${oci_core_subnet.test_subnet.id}"
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
	AutonomousDatabaseResourceDependenciesCON = //DefinedTagsDependencies +
	generateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", Required, Create, autonomousDbVersionDataSourceRepresentation) +
		generateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", Required, Create,
			representationCopyWithNewProperties(autonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": Representation{repType: Required, create: `DW`}}))

	AutonomousDatabaseResourceDependenciesCONSOURCE = //DefinedTagsDependencies +
	generateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions_source", Required, Create, autonomousDbVersionDataSourceRepresentation) +
		generateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions_source", Required, Create,
			representationCopyWithNewProperties(autonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": Representation{repType: Required, create: `DW`}}))

	goldenGateDbSystemRepresentationSOURCE = map[string]interface{}{
		"availability_domain":     Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name}")}`},
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"database_edition":        Representation{repType: Required, create: `ENTERPRISE_EDITION`},
		"db_home":                 RepresentationGroup{Required, goldenGateDbSystemDbHomeRepresentation},
		"hostname":                Representation{repType: Required, create: `myDB`},
		"shape":                   Representation{repType: Required, create: `VM.Standard2.2`},
		"ssh_public_keys":         Representation{repType: Required, create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"subnet_id":               Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": Representation{repType: Optional, create: `256`},
		"display_name":            Representation{repType: Optional, create: `tfGGmyDB`},
		"domain":                  Representation{repType: Optional, create: `myDB`},
		"node_count":              Representation{repType: Optional, create: `1`},
		"db_system_options":       RepresentationGroup{Optional, goldenGateDbSystemOption},
		"private_ip":              Representation{repType: Required, create: `10.0.0.125`},
	}

	ConnectionResourceDependenciesTarget = generateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation) +
		AutonomousDatabaseResourceDependenciesCON +
		KmsKeyIdVariableStr +
		KmsVaultIdVariableStr

	ConnectionResourceDependenciesTargetCommon = generateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		SubnetData +
		generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation) +
		AutonomousDatabaseResourceDependenciesCON //+

	ConnectionResourceDependenciesSource = generateResourceFromRepresentationMap("oci_database_db_system", "t", Optional, Create, goldenGateDbSystemRepresentationSOURCE) +
		DatabaseData +
		DatabaseHomeConfig +
		KmsKeyIdVariableStr +
		AvailabilityDomainConfig +
		KmsVaultIdVariableStr
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithBlankDefault("compartment_id_for_update")

	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_migration_connection.test_connection"
	datasourceName := "data.oci_database_migration_connections.test_connections"
	singularDatasourceName := "data.oci_database_migration_connection.test_connection"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ConnectionResourceDependenciesTarget+
		generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Optional, Create, connectionRepresentationTarget), "databasemigration", "connection", t)

	ResourceTest(t, testAccCheckDatabaseMigrationConnectionDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + ConnectionResourceDependenciesTarget +
				generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Required, Create, connectionRepresentationTarget),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_type", "AUTONOMOUS"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + ConnectionResourceDependenciesTarget,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + ConnectionResourceDependenciesTarget +
				generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Optional, Create, connectionRepresentationTarget),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.password", "ORcl##4567890"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_type", "AUTONOMOUS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_display_test_create"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ConnectionResourceDependenciesTarget +
				generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Optional, Create,
					representationCopyWithNewProperties(connectionRepresentationTarget, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.password", "ORcl##4567890"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_type", "AUTONOMOUS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_display_test_create"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),
				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ConnectionResourceDependenciesTarget +
				generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Optional, Update, connectionRepresentationTarget),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.password", "ORcl##4567890"),
				resource.TestCheckResourceAttr(resourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_type", "AUTONOMOUS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_display_test_update"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),
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
				generateDataSourceFromRepresentationMap("oci_database_migration_connections", "test_connections", Optional, Update, connectionDataSourceRepresentationCon) +
				compartmentIdVariableStr + ConnectionResourceDependenciesTarget +
				generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Optional, Update, connectionRepresentationTarget),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "connection_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Required, Create, connectionSingularDataSourceRepresentationCon) +
				compartmentIdVariableStr + ConnectionResourceConfigTarget,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "admin_credentials.0.username", "admin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_type", "AUTONOMOUS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_display_test_update"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_details.0.compartment_id", compartmentId),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ConnectionResourceConfigTarget,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"tls_keystore",
				"tls_wallet",
				"admin_credentials.0.password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseMigrationConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_connection" {
			noResourceFound = false
			request := oci_database_migration.GetConnectionRequest{}

			tmp := rs.Primary.ID
			request.ConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_migration")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseMigrationConnection") {
		resource.AddTestSweepers("DatabaseMigrationConnection", &resource.Sweeper{
			Name:         "DatabaseMigrationConnection",
			Dependencies: DependencyGraph["connection"],
			F:            sweepDatabaseMigrationConnectionResource,
		})
	}
}

func sweepDatabaseMigrationConnectionResource(compartment string) error {
	databaseMigrationClient := GetTestClients(&schema.ResourceData{}).databaseMigrationClient()
	connectionIds, err := getConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, connectionId := range connectionIds {
		if ok := SweeperDefaultResourceId[connectionId]; !ok {
			deleteConnectionRequest := oci_database_migration.DeleteConnectionRequest{}

			deleteConnectionRequest.ConnectionId = &connectionId

			deleteConnectionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteConnection(context.Background(), deleteConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting Connection %s %s, It is possible that the resource is already deleted. Please verify manually \n", connectionId, error)
				continue
			}
			waitTillCondition(testAccProvider, &connectionId, connectionSweepWaitCondition, time.Duration(3*time.Minute),
				connectionSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getConnectionIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := GetTestClients(&schema.ResourceData{}).databaseMigrationClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ConnectionId", id)
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

func connectionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseMigrationClient().GetConnection(context.Background(), oci_database_migration.GetConnectionRequest{
		ConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
