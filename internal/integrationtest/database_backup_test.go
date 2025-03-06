// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

var (
	// *** RESOURCE REPRESENTATIONS ***

	// VCN
	dbBackupVcnResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfVcnForDatabaseBackup`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `tfvcn`},
	}

	// Internet Gateway
	dbBackupInternetGatewayResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfInternetGateway`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	// Route Table
	dbBackupRouteTableResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfRouteTable`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_rules":    acctest.RepresentationGroup{RepType: acctest.Required, Group: routeRulesGroup},
	}

	routeRulesGroup = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"description":       acctest.Representation{RepType: acctest.Required, Create: `Internal traffic for OCI Services`},
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"destination_type":  acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
	}

	// Public Subnet
	dbBackupSubnetResourceRepresentation = map[string]interface{}{
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `tfPublicSubnet`},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/24`},
		"vcn_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_table_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}},
		"dhcp_options_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"dns_label":         acctest.Representation{RepType: acctest.Required, Create: `tfpublicsubnet`},
	}

	// Db System used to create a database backup via Recovery Service
	dbBackupDbSystemResourceRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDbSystemWithDatabaseBackup`},
		"database_edition":        acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`}},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `tf-oracle-db`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbBackupDbHomeGroup},
		"depends_on":              acctest.Representation{RepType: acctest.Required, Create: []string{"oci_recovery_recovery_service_subnet.test_recovery_service_subnet_registration"}},
	}

	dbBackupDbHomeGroup = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `tfDbHome`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbBackupDatabaseGroup},
	}

	dbBackupDatabaseGroup = map[string]interface{}{
		"db_name":          acctest.Representation{RepType: acctest.Optional, Create: `tfDb`},
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbBackupConfigGroup},
	}

	dbBackupConfigGroup = map[string]interface{}{
		"auto_backup_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"auto_backup_window":         acctest.Representation{RepType: acctest.Optional, Create: `SLOT_TWO`},
		"backup_deletion_policy":     acctest.Representation{RepType: acctest.Optional, Create: `DELETE_IMMEDIATELY`},
		"run_immediate_full_backup":  acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: backupDestinationDetailsGroup},
	}

	backupDestinationDetailsGroup = map[string]interface{}{
		"dbrs_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_recovery_protection_policy.test_protection_policy.id}`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `DBRS`},
	}

	// Service Gateway used by Recovery Service's Private Subnet
	rsServiceGatewayResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfRecoveryServiceServiceGateway`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"services":       acctest.RepresentationGroup{RepType: acctest.Required, Group: rsServicesGroup},
	}

	rsServicesGroup = map[string]interface{}{
		"service_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_services.test_services.services.0.id}`},
	}

	// Route Table used by the Recovery Service's Private Subnet
	rsRouteTableRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfRecoveryServicePrivateSubnetRouteTable`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_rules":    acctest.RepresentationGroup{RepType: acctest.Required, Group: rsRouteRuleGroup},
	}

	rsRouteRuleGroup = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_service_gateway.test_service_gateway.id}`},
		"description":       acctest.Representation{RepType: acctest.Required, Create: `Recovery Service traffic for OCI Services`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_services.test_services.services[0].cidr_block}`},
		"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `SERVICE_CIDR_BLOCK`},
	}

	// Private Subnet used by the Recovery Service
	rsPrivateSubnetResourceRepresentation = map[string]interface{}{
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `tfPrivateSubnet`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/24`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_table_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_private_subnet_route_table.id}`},
		"security_list_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_private_subnet_security_list.id}`}},
		"dhcp_options_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"dns_label":                  acctest.Representation{RepType: acctest.Required, Create: `tfprivatesubnet`},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	// Security List used by the Recovery Service's Private Subnet
	rsSecurityListRepresentation = map[string]interface{}{
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `tfRecoveryServiceSecurityList`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: rsIngressSecurityRule1Group}, {RepType: acctest.Required, Group: rsIngressSecurityRule2Group}},
		"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: rsEgressSecurityRulesGroup}},
	}

	rsIngressSecurityRule1Group = map[string]interface{}{
		"source":      acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: rsTcpOptions1Group},
	}

	rsTcpOptions1Group = map[string]interface{}{
		"min": acctest.Representation{RepType: acctest.Required, Create: `8005`},
		"max": acctest.Representation{RepType: acctest.Required, Create: `8005`},
	}

	rsIngressSecurityRule2Group = map[string]interface{}{
		"source":      acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: rsTcpOptions2Group},
	}

	rsTcpOptions2Group = map[string]interface{}{
		"min": acctest.Representation{RepType: acctest.Required, Create: `2484`},
		"max": acctest.Representation{RepType: acctest.Required, Create: `2484`},
	}

	rsEgressSecurityRulesGroup = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `all`},
	}

	// Recovery Service Policy
	rsSubnetProtectionPolicyResourceRepresentation = map[string]interface{}{
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `tfRecoveryServiceSubnetProtectionPolicy`},
		"backup_retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `14`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	// Recovery Service Subnet Registration
	rsSubnetRegistrationResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfRecoveryServiceSubnetRegistration`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnets":        acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.test_private_subnet.id}`}},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	// Db System Database Backup
	dbBackupResourceRepresentation = map[string]interface{}{
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `RecoveryServiceMonthlyDatabaseBackup`},
		"database_id":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_db_system.databases.0.id}`},
		"retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `90`, Update: `180`},
		"depends_on":               acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.test_db_system"}},
	}

	// *** DATASOURCE REPRESENTATIONS ***
	dbBackupOciServicesDatasourceRepresentation = map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: regexFilterGroup},
	}

	regexFilterGroup = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`.*Oracle.*Services.*Network`}},
		"regex":  acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	dbBackupDatabaseDbHomesDatasourceRepresentation = map[string]interface{}{
		"db_system_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_system.test_db_system.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: filterByDbHomeDisplayNameGroup},
	}

	filterByDbHomeDisplayNameGroup = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `display_name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`tfDbHome`}},
	}

	dbBackupDatabaseDatasourceRepresentation = map[string]interface{}{
		"db_home_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_db_homes.test_db_system.db_homes.0.db_home_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	dbBackupDatasourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_databases.test_db_system.databases.0.id}`},
	}

	dbBackupDatasourceByShapeFamilyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_family":   acctest.Representation{RepType: acctest.Optional, Create: `VIRTUALMACHINE`},
	}

	// Network Datasource Config(s)
	dbBackupServiceDatasourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, dbBackupOciServicesDatasourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "test_db_system", acctest.Optional, acctest.Create, dbBackupDatabaseDbHomesDatasourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_db_system", acctest.Optional, acctest.Create, dbBackupDatabaseDatasourceRepresentation)

	// Network Resource Config(s)
	dbBackupVcnConfig             = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, dbBackupVcnResourceRepresentation)
	dbBackupRouteTableConfig      = acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, dbBackupRouteTableResourceRepresentation)
	dbBackupInternetGatewayConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Optional, acctest.Create, dbBackupInternetGatewayResourceRepresentation)
	dbBackupSubnetConfig          = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, dbBackupSubnetResourceRepresentation)

	// Network [Recovery Service] Resource Config(s)
	rsSecurityListConfig       = acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_private_subnet_security_list", acctest.Optional, acctest.Create, rsSecurityListRepresentation)
	rsRouteTableConfig         = acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_private_subnet_route_table", acctest.Optional, acctest.Create, rsRouteTableRepresentation)
	rsPrivateSubnetConfig      = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_private_subnet", acctest.Required, acctest.Create, rsPrivateSubnetResourceRepresentation)
	rsServiceGatewayConfig     = acctest.GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", acctest.Required, acctest.Create, rsServiceGatewayResourceRepresentation)
	rsProtectionPolicy         = acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Required, acctest.Create, rsSubnetProtectionPolicyResourceRepresentation)
	rsSubnetRegistrationConfig = acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet_registration", acctest.Required, acctest.Create, rsSubnetRegistrationResourceRepresentation)

	// Db System Config
	dbBackupDbSystemConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Create, dbBackupDbSystemResourceRepresentation)

	// Base Config(s)
	dbBackupBaseConfig = DefinedTagsDependencies + AvailabilityDomainConfig + dbBackupServiceDatasourceConfig +
		dbBackupVcnConfig + dbBackupRouteTableConfig + dbBackupInternetGatewayConfig + dbBackupSubnetConfig +
		rsProtectionPolicy + rsSecurityListConfig + rsServiceGatewayConfig + rsRouteTableConfig +
		rsPrivateSubnetConfig + rsSubnetRegistrationConfig + dbBackupDbSystemConfig
)

// issue-routing-tag: database/default
func TestDatabaseBackupResource_basic(t *testing.T) {
	config := acctest.ProviderTestConfig()

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_backup.test_backup"
	datasourceName := "data.oci_database_backups.test_backups"
	datasourceByShapeFamilyName := "data.oci_database_backups.test_backups_by_shape_family"

	const dbWaitConditionDuration = time.Duration(6 * time.Minute)
	var resId, resId2, dbId string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create Db System
		{
			Config: config + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + compartmentIdVariableStr + dbBackupBaseConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet("oci_database_db_system.test_db_system", "display_name"),
				func(s *terraform.State) (err error) {
					dbId, err = acctest.FromInstanceState(s, "data.oci_database_databases.test_db_system", "databases.0.id")
					return err
				},
			),
		},
		// Wait for auto backup to complete and create a manual backup
		{
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &dbId, dbBackupAvailableWaitCondition, dbWaitConditionDuration, listBackupsFetchOperation, "core", false),
			Config: config + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + compartmentIdVariableStr + dbBackupBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, dbBackupResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "RecoveryServiceMonthlyDatabaseBackup"),
				resource.TestCheckResourceAttr(resourceName, "retention_period_in_days", "90"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_database_backups", "test_backup", acctest.Optional, acctest.Create, dbBackupDatasourceRepresentation) +
				kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + compartmentIdVariableStr + dbBackupBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Update, dbBackupResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "RecoveryServiceMonthlyDatabaseBackup"),
				resource.TestCheckResourceAttr(resourceName, "retention_period_in_days", "180"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource list by shape family
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_database_backups", "test_backups_by_shape_family", acctest.Optional, acctest.Update, dbBackupDatasourceByShapeFamilyRepresentation) +
				kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + compartmentIdVariableStr + dbBackupBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Update, dbBackupResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceByShapeFamilyName, "shape_family", "VIRTUALMACHINE"),
				resource.TestCheckResourceAttr(datasourceByShapeFamilyName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.database_id"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.time_ended"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.time_expiry_scheduled"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.database_edition"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.backup_destination_type"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.type"),
				resource.TestCheckResourceAttrSet(datasourceByShapeFamilyName, "backups.0.version"),
			),
		},
		// verify datasource
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_database_backups", "test_backups", acctest.Optional, acctest.Update, dbBackupDatasourceRepresentation) +
				kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + compartmentIdVariableStr + dbBackupBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Update, dbBackupResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_id"),

				resource.TestCheckResourceAttr(datasourceName, "backups.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.time_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.time_expiry_scheduled"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_edition"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.backup_destination_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.version"),

				resource.TestCheckResourceAttr(datasourceName, "backups.0.retention_period_in_days", "180"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.backup_destination_type", "DBRS"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.display_name", "RecoveryServiceMonthlyDatabaseBackup"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.type", "FULL"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.version", "19.25.0.0"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.state", "ACTIVE"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseBackup") {
		resource.AddTestSweepers("DatabaseBackup", &resource.Sweeper{
			Name:         "DatabaseBackup",
			Dependencies: acctest.DependencyGraph["backup"],
			F:            sweepDatabaseBackupResource,
		})
	}
}

func sweepDatabaseBackupResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	backupIds, err := getDatabaseBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, backupId := range backupIds {
		if ok := acctest.SweeperDefaultResourceId[backupId]; !ok {
			deleteBackupRequest := oci_database.DeleteBackupRequest{}

			deleteBackupRequest.BackupId = &backupId

			deleteBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteBackup(context.Background(), deleteBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting Backup %s %s, It is possible that the resource is already deleted. Please verify manually \n", backupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &backupId, DatabaseBackupSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseBackupSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listBackupsRequest := oci_database.ListBackupsRequest{}
	listBackupsRequest.CompartmentId = &compartmentId
	listBackupsRequest.LifecycleState = oci_database.BackupSummaryLifecycleStateActive
	listBackupsResponse, err := databaseClient.ListBackups(context.Background(), listBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Backup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, backup := range listBackupsResponse.Items {
		id := *backup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BackupId", id)
	}
	return resourceIds, nil
}

func DatabaseBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if backupResponse, ok := response.Response.(oci_database.GetBackupResponse); ok {
		return backupResponse.LifecycleState != oci_database.BackupLifecycleStateDeleted
	}
	return false
}

func DatabaseBackupSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetBackup(context.Background(), oci_database.GetBackupRequest{
		BackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
