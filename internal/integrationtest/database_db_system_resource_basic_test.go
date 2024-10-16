// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/database"
)

var (

	// Upstream Dependencies Start
	// Core VCN Resource Representation
	DbSystemCoreVcnRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfVcn`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.1.0.0/16`},
		"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `tfvcn`},
	}

	// Core Internet Gateway Resource Representation
	DbSystemCoreInternetGatewayResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfInternetGateway`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	// Route Table Resource Representation
	DbSystemRouteTableResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfRouteTable`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_rules":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemRouteRulesGroup},
	}

	DbSystemRouteRulesGroup = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"description":       acctest.Representation{RepType: acctest.Required, Create: `Internal traffic for OCI Services`},
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
	}

	// Subnet Resource Representation
	DbSystemSubnetResourceRepresentation = map[string]interface{}{
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `tfSubnet`},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `10.1.20.0/24`},
		"vcn_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_table_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}},
		"dhcp_options_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"dns_label":         acctest.Representation{RepType: acctest.Required, Create: `tfsubnet`},
	}

	// Network Security Group Resource Representation
	DbSystemNetworkSecurityGroupResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfNetworkSecurityGroup`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}
	// Upstream Dependencies End

	// 1. Main Db System Resource Representation: Start
	DbSystemResourceRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDbSystem`},
		"database_edition":        acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`, Update: `BRING_YOUR_OWN_LICENSE`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"fault_domains":           acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}},
		"security_attributes":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle-zpr.maxegresscount.value": "42", "oracle-zpr.maxegresscount.mode": "enforce"}, Update: map[string]string{"oracle-zpr.maxegresscount.value": "updatedValue", "oracle-zpr.maxegresscount.mode": "enforce"}},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		//"nsg_ids":              acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
		"ssh_public_keys":     acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`}},
		"hostname":            acctest.Representation{RepType: acctest.Required, Create: `tfOracleDb`},
		"db_home":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemDbHomeGroup},
	}

	DbSystemDbHomeGroup = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `tfDbHome`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.25.0.0`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemDatabaseGroup},
	}

	DbSystemDatabaseGroup = map[string]interface{}{
		"db_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDb`},
		"pdb_name":           acctest.Representation{RepType: acctest.Optional, Create: `tfPdb`},
		"character_set":      acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"ncharacter_set":     acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"db_workload":        acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"kms_key_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"kms_key_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_version_id}`},
		"vault_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"admin_password":     acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
	}

	// 1. Main Db System Resource Representation: End

	// 2. Source Db System Resource Representation: Start
	DbSystemSourceResourceRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDbSystemSource`},
		"database_edition":        acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"fault_domains":           acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}},
		"security_attributes":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle-zpr.maxegresscount.value": "42", "oracle-zpr.maxegresscount.mode": "enforce"}},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`}},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `tfOracleDb`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemSourceDbHomeGroup},
	}

	DbSystemSourceDbHomeGroup = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `tfDbHome`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.25.0.0`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemSourceDatabaseGroup},
	}

	DbSystemSourceDatabaseGroup = map[string]interface{}{
		"db_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDb`},
		"pdb_name":           acctest.Representation{RepType: acctest.Optional, Create: `tfPdb`},
		"kms_key_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"vault_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"kms_key_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_version_id}`},
		"character_set":      acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"ncharacter_set":     acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"db_workload":        acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"db_backup_config":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DbSystemSourceDbBackupConfigGroup},
		"admin_password":     acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
	}

	DbSystemSourceDbBackupConfigGroup = map[string]interface{}{
		"auto_backup_enabled":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"run_immediate_full_backup": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	// 2. Source Db System Resource Representation: End

	// 3. FromDatabase Db System Resource Representation: Start
	DbSystemFromDatabaseResourceRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDbSystemFromDatabase`},
		"database_edition":        acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`}},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `tfOracleDb`},
		"source":                  acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemFromDatabaseDbHomeGroup},
	}

	DbSystemFromDatabaseDbHomeGroup = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `tfDbHome`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.25.0.0`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemFromDatabaseDatabaseGroup},
	}

	DbSystemFromDatabaseDatabaseGroup = map[string]interface{}{
		"db_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDb`},
		"kms_key_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"vault_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"kms_key_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_version_id}`},
		"database_id":        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_databases.test_source_db_system.databases.0.id}`},
		"admin_password":     acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
	}
	// 3. FromDatabase End

	DbSystemSourceDatabaseDataSourceRepresentation = map[string]interface{}{
		"db_home_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_db_homes.test_source_db_system.db_homes.0.db_home_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DbSystemSourceDatabaseDbHomesDataSourceRepresentation = map[string]interface{}{
		"db_system_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_system.test_source_db_system.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	// Database Upstream Resource Dependencies Configs
	DbSystemCoreVcnConfig         = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, DbSystemCoreVcnRepresentation)
	DbSystemRouteTableConfig      = acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, DbSystemRouteTableResourceRepresentation)
	DbSystemInternetGatewayConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, DbSystemCoreInternetGatewayResourceRepresentation)
	DbSystemSubnetConfig          = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, DbSystemSubnetResourceRepresentation)

	// Base configuration for Db System - Tags, ADs, Vcn, Route Tables, Internet Gateway, Subnet
	DbSystemBaseConfig = DefinedTagsDependencies + AvailabilityDomainConfig + DbSystemCoreVcnConfig + DbSystemRouteTableConfig + DbSystemInternetGatewayConfig + DbSystemSubnetConfig

	DbSystemResourceName             = "oci_database_db_system.test_db_system"
	DbSystemSourceResourceName       = "oci_database_db_system.test_source_db_system"
	DbSystemFromDatabaseResourceName = "oci_database_db_system.test_db_system_from_database"

	// Downstream dependencies only (Not used in current file)
	ResourceDatabaseResourceName                   = "oci_database_db_system.t"
	ResourceDatabaseBaseConfig                     = acctest.LegacyTestProviderConfig() + DbSystemBaseConfig
	DbSystemResourceConfig                         = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Required, acctest.Create, DbSystemResourceRepresentation)
	DbSystemResourceDependencies                   = DbSystemBaseConfig
	ResourceDatabaseToken, ResourceDatabaseTokenFn = acctest.TokenizeWithHttpReplay("database_db")
)

// issue-routing-tag: database/default
func TestResourceDatabaseDBSystemFromDatabase(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_basic") {
		t.Skip("Skipping suppressed DBSystem_basic")
	}

	config := acctest.LegacyTestProviderConfig()

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	const dbWaitConditionDuration = time.Duration(6 * time.Minute)

	var resId string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create Source Db System
		{
			Config: config + DbSystemBaseConfig + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_source_db_system", acctest.Optional, acctest.Create, DbSystemSourceResourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "test_source_db_system", acctest.Optional, acctest.Create, DbSystemSourceDatabaseDbHomesDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_source_db_system", acctest.Optional, acctest.Create, DbSystemSourceDatabaseDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(DbSystemSourceResourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(DbSystemSourceResourceName, "security_attributes.oracle-zpr.maxegresscount.value", "42"),
				resource.TestCheckResourceAttr(DbSystemSourceResourceName, "security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, "data.oci_database_databases.test_source_db_system", "databases.0.id")
					return err
				},
			),
		},
		// Wait for backup to complete and create a new Db System from the Source Db System's database
		{
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, dbBackupAvailableWaitCondition, dbWaitConditionDuration,
				listBackupsFetchOperation, "core", false),
			Config: config + DbSystemBaseConfig + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_source_db_system", acctest.Optional, acctest.Create, DbSystemSourceResourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "test_source_db_system", acctest.Optional, acctest.Create, DbSystemSourceDatabaseDbHomesDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_source_db_system", acctest.Optional, acctest.Create, DbSystemSourceDatabaseDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system_from_database", acctest.Optional, acctest.Create, DbSystemFromDatabaseResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(DbSystemFromDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(DbSystemFromDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(DbSystemFromDatabaseResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(DbSystemFromDatabaseResourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(DbSystemFromDatabaseResourceName, "time_created"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "display_name", "tfDbSystemFromDatabase"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(DbSystemFromDatabaseResourceName, "hostname"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "db_home.0.db_version", "19.25.0.0"),
				resource.TestCheckResourceAttrSet(DbSystemFromDatabaseResourceName, "db_home.0.display_name"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "db_home.0.database.0.db_name", "tfDb"),
				resource.TestCheckResourceAttr(DbSystemFromDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.test_source_db_system", "databases.0.last_backup_timestamp"),
			),
		},
	})
}

func dbAutomaticBackupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if listBackupResponse, ok := response.Response.(database.ListBackupsResponse); ok {
		if len(listBackupResponse.Items) > 1 {
			return listBackupResponse.Items[1].LifecycleState != database.BackupSummaryLifecycleStateActive
		}
		return true
	}
	return false
}

// TestAccResourceDatabaseDBSystem_basic tests creation of a DBSystem with the minimum required properties
// to assert expected default values are set
// issue-routing-tag: database/default
func TestResourceDatabaseDBSystemBasic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemBasic")
	defer httpreplay.SaveScenario()

	// This test is a subset of TestAccResourceDatabaseDBSystem_allXX. It tests omitting optional params.
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_basic") {
		t.Skip("Skipping suppressed DBSystem_basic")
	}

	config := acctest.LegacyTestProviderConfig()
	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + DbSystemBaseConfig +
				//acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, DbSystemNetworkSecurityGroupResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Create, DbSystemResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "id"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "time_created"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "display_name", `tfDbSystem`),
				resource.TestCheckResourceAttr(DbSystemResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "hostname"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "db_home.0.db_version"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "db_home.0.display_name"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "db_home.0.database.0.db_name", "tfDb"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "security_attributes.oracle-zpr.maxegresscount.value", "42"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				//resource.TestCheckResourceAttr(DbSystemResourceName, "nsg_ids.#", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, "oci_database_db_system.test_db_system", "id")
					return err
				},
			),
		},
		// Verify Update without updating nsgIds
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + DbSystemBaseConfig +
				//acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, DbSystemNetworkSecurityGroupResourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Update, DbSystemResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "id"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "time_created"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "display_name", `tfDbSystem`),
				resource.TestCheckResourceAttr(DbSystemResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "hostname"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "db_home.0.db_version"),
				resource.TestCheckResourceAttrSet(DbSystemResourceName, "db_home.0.display_name"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "db_home.0.database.0.db_name", "tfDb"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "security_attributes.oracle-zpr.maxegresscount.value", "updatedValue"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				resource.TestCheckResourceAttr(DbSystemResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.test_db_system", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	resource.AddTestSweepers("DatabaseDbSystem", &resource.Sweeper{
		Name:         "DatabaseDbSystem",
		Dependencies: acctest.DependencyGraph["dbSystem"],
		F:            sweepDatabaseDbSystemResource,
	})
}
