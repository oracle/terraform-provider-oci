// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v58/databasemigration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConnectionResourceConfigTarget = ConnectionResourceDependenciesTarget +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, connectionRepresentationTarget)

	connectionSingularDataSourceRepresentationCon = map[string]interface{}{
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_connection.test_connection.id}`},
	}

	connectionDataSourceRepresentationCon = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionDataSourceFilterRepresentationCon}}
	connectionDataSourceFilterRepresentationCon = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_connection.test_connection.id}`}},
	}

	connectionRepresentationCon = map[string]interface{}{
		"admin_credentials":  acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":      acctest.Representation{RepType: acctest.Required, Create: `MANUAL`},
		"vault_details":      acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"certificate_tdn":    acctest.Representation{RepType: acctest.Optional, Create: `certificateTdn`, Update: `certificateTdn2`},
		"connect_descriptor": acctest.RepresentationGroup{RepType: acctest.Optional, Group: connectionConnectDescriptorRepresentation},
		"database_id":        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
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
		"database_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `TF_display_test_create`, Update: `TF_display_test_update`},
	}

	connectionConnectDescriptorRepresentationMIG = map[string]interface{}{
		"connect_string": acctest.Representation{RepType: acctest.Required, Create: `(description=(address=(port=1521)(host=10.0.0.125))(connect_data=(service_name=pdb1120.exadbpriv.exadbvcn.oraclevcn.com)))`, Update: `(description=(address=(port=1521)(host=10.0.0.125))(connect_data=(service_name=pdb1120.exadbpriv.exadbvcn.oraclevcn.com)))`},
	}
	connectionRepresentationSource = map[string]interface{}{
		"admin_credentials":  acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionAdminCredentialsRepresentation},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_type":      acctest.Representation{RepType: acctest.Required, Create: `MANUAL`},
		"vault_details":      acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionVaultDetailsRepresentation},
		"connect_descriptor": acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionConnectDescriptorRepresentationMIG},
		"database_id":        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_databases.t.databases.0.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `TF_display_test_create_source`, Update: `TF_display_test_update_source`},
		"ssh_details":        acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionSshDetailsRepresentation},
		"private_endpoint":   acctest.RepresentationGroup{RepType: acctest.Required, Group: connectionPrivateEndpointRepresentation},
	}

	connectionAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `ORcl##4567890`, Update: `ORcl##4567890`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `admin`, Update: `admin`},
	}

	connectionAdminCredentialsRepresentationUPDATE = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `ORcl##4567890`},
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
		"connect_string":        acctest.Representation{RepType: acctest.Optional, Create: `connectString`, Update: `connectString2`},
		"database_service_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_services.test_services.name}`},
		"host":                  acctest.Representation{RepType: acctest.Optional, Create: `host`, Update: `host2`},
		"port":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	connectionPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}
	connectionSshDetailsRepresentation = map[string]interface{}{
		"host":          acctest.Representation{RepType: acctest.Required, Create: `10.0.0.125`, Update: `10.0.0.125`},
		"sshkey":        acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`, Update: `ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`},
		"user":          acctest.Representation{RepType: acctest.Required, Create: `opc`, Update: `opc`},
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
	acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Required, acctest.Create, autonomousDbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(autonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))

	AutonomousDatabaseResourceDependenciesCONSOURCE = //DefinedTagsDependencies +
	acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions_source", acctest.Required, acctest.Create, autonomousDbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions_source", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(autonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))

	goldenGateDbSystemRepresentationSOURCE = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name}")}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDbSystemDbHomeRepresentation},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myDB`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfGGmyDB`},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `myDB`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"db_system_options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: goldenGateDbSystemOption},
		"private_ip":              acctest.Representation{RepType: acctest.Required, Create: `10.0.0.125`},
	}

	ConnectionResourceDependenciesTarget = acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, serviceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation) +
		AutonomousDatabaseResourceDependenciesCON +
		KmsKeyIdVariableStr +
		KmsVaultIdVariableStr

	ConnectionResourceDependenciesTargetCommon = acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, serviceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		SubnetData +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation) +
		AutonomousDatabaseResourceDependenciesCON //+

	ConnectionResourceDependenciesSource = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "t", acctest.Optional, acctest.Create, goldenGateDbSystemRepresentationSOURCE) +
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

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithBlankDefault("compartment_id_for_update")

	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_migration_connection.test_connection"
	datasourceName := "data.oci_database_migration_connections.test_connections"
	singularDatasourceName := "data.oci_database_migration_connection.test_connection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ConnectionResourceDependenciesTarget+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create, connectionRepresentationTarget), "databasemigration", "connection", t)

	acctest.ResourceTest(t, testAccCheckDatabaseMigrationConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ConnectionResourceDependenciesTarget +
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
			Config: config + compartmentIdVariableStr + ConnectionResourceDependenciesTarget,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ConnectionResourceDependenciesTarget +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create, connectionRepresentationTarget),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ConnectionResourceDependenciesTarget +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(connectionRepresentationTarget, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, connectionRepresentationTarget),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_connections", "test_connections", acctest.Optional, acctest.Update, connectionDataSourceRepresentationCon) +
				compartmentIdVariableStr + ConnectionResourceDependenciesTarget +
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
				compartmentIdVariableStr + ConnectionResourceConfigTarget,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
