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
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_database "github.com/oracle/oci-go-sdk/v53/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	exaVcnRepresentation = map[string]interface{}{
		"cidr_block":     Representation{RepType: Required, Create: `10.1.0.0/16`},
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `-tf-vcn`},
		"dns_label":      Representation{RepType: Optional, Create: `tfvcn`},
	}

	exaSecurityListRepresentation = map[string]interface{}{
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"vcn_id":                 Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"display_name":           Representation{RepType: Optional, Create: `ExadataSecurityList`},
		"egress_security_rules":  []RepresentationGroup{{Required, exaSecurityListEgressSecurityRulesICMPRepresentation}, {Optional, exaSecurityListEgressSecurityRulesTCPRepresentation}},
		"ingress_security_rules": []RepresentationGroup{{Required, exaSecurityListIngressSecurityRulesICMPRepresentation}, {Optional, exaSecurityListIngressSecurityRulesTCPRepresentation}},
	}

	exaSecurityListIngressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol": Representation{RepType: Required, Create: `1`},
		"source":   Representation{RepType: Required, Create: `10.1.22.0/24`},
	}
	exaSecurityListIngressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol": Representation{RepType: Required, Create: `6`},
		"source":   Representation{RepType: Required, Create: `10.1.22.0/24`},
	}
	exaSecurityListEgressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol":    Representation{RepType: Required, Create: `1`},
		"destination": Representation{RepType: Required, Create: `10.1.22.0/24`},
	}
	exaSecurityListEgressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol":    Representation{RepType: Required, Create: `6`},
		"destination": Representation{RepType: Required, Create: `10.1.22.0/24`},
	}

	exaSubnetRepresentation = map[string]interface{}{
		"cidr_block":          Representation{RepType: Required, Create: `10.1.22.0/24`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"vcn_id":              Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name}")}`},
		"dhcp_options_id":     Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"display_name":        Representation{RepType: Optional, Create: `ExadataSubnet`},
		"dns_label":           Representation{RepType: Optional, Create: `subnetexadata1`},
		"route_table_id":      Representation{RepType: Optional, Create: `${oci_core_route_table.exadata_route_table.id}`},
		"security_list_ids":   Representation{RepType: Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`, `${oci_core_security_list.exadata_shapes_security_list.id}`}},
	}
	exaBackupSubnetRepresentation = map[string]interface{}{
		"cidr_block":          Representation{RepType: Required, Create: `10.1.23.0/24`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"vcn_id":              Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name}")}`},
		"dhcp_options_id":     Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"display_name":        Representation{RepType: Optional, Create: `ExadataBackupSubnet`},
		"dns_label":           Representation{RepType: Optional, Create: `subnetexadata2`},
		"route_table_id":      Representation{RepType: Optional, Create: `${oci_core_route_table.exadata_route_table.id}`},
		"security_list_ids":   Representation{RepType: Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}},
	}

	exadbSystemRepresentation = map[string]interface{}{
		"availability_domain":     Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name}")}`},
		"backup_subnet_id":        Representation{RepType: Required, Create: `${oci_core_subnet.exadata_backup_subnet.id}`},
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"database_edition":        Representation{RepType: Required, Create: `ENTERPRISE_EDITION_EXTREME_PERFORMANCE`},
		"db_home":                 RepresentationGroup{Required, exadbSystemDbHomeRepresentation},
		"hostname":                Representation{RepType: Required, Create: `myOracleDB`},
		"shape":                   Representation{RepType: Required, Create: `Exadata.Quarter1.84`},
		"ssh_public_keys":         Representation{RepType: Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"subnet_id":               Representation{RepType: Required, Create: `${oci_core_subnet.exadata_subnet.id}`},
		"cpu_core_count":          Representation{RepType: Optional, Create: `22`},
		"data_storage_size_in_gb": Representation{RepType: Optional, Create: `256`},
		"disk_redundancy":         Representation{RepType: Optional, Create: `HIGH`},
		"display_name":            Representation{RepType: Optional, Create: `tfDbSystemTestExadata`},
		"domain":                  Representation{RepType: Optional, Create: `${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_vcn.test_vcn.dns_label}.oraclevcn.com`},
		"license_model":           Representation{RepType: Optional, Create: `LICENSE_INCLUDED`},
		"node_count":              Representation{RepType: Optional, Create: `1`},
	}
	exadbSystemDbHomeRepresentation = map[string]interface{}{
		"database":     RepresentationGroup{Required, exadbSystemDbHomeDatabaseRepresentation},
		"db_version":   Representation{RepType: Optional, Create: `12.1.0.2`},
		"display_name": Representation{RepType: Optional, Create: `dbHome1`},
	}
	exadbSystemDbHomeDatabaseRepresentation = map[string]interface{}{
		"admin_password": Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"db_name":        Representation{RepType: Optional, Create: `tfDbName`},
	}

	ExaBaseDependencies = GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, exaVcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_security_list", "exadata_shapes_security_list", Optional, Create, exaSecurityListRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "exadata_subnet", Optional, Create, exaSubnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "exadata_backup_subnet", Optional, Create, exaBackupSubnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", Optional, Create, exadbSystemRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "exadata_route_table", Optional, Create, routeTableRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Optional, Create, internetGatewayRepresentation)

	DatabaseRequiredOnlyResource = DatabaseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseRepresentation)

	DatabaseResourceConfig = DatabaseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseRepresentation)

	databaseSingularDataSourceRepresentation = map[string]interface{}{
		"database_id": Representation{RepType: Required, Create: `${oci_database_database.test_database.id}`},
	}

	databaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"db_home_id":     Representation{RepType: Optional, Create: `${oci_database_db_home.test_db_home.id}`},
		"db_name":        Representation{RepType: Optional, Create: `myTestDb`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, databaseDataSourceFilterRepresentation}}
	databaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_database.test_database.id}`}},
	}

	databaseRepresentation = map[string]interface{}{
		"database":         RepresentationGroup{Required, databaseDatabaseRepresentation},
		"db_home_id":       Representation{RepType: Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":           Representation{RepType: Required, Create: `NONE`},
		"db_version":       Representation{RepType: Optional, Create: `12.1.0.2`},
		"kms_key_id":       Representation{RepType: Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_rotation": Representation{RepType: Optional, Update: `1`},
	}

	databaseRepresentationMigration = map[string]interface{}{
		"database":          RepresentationGroup{Required, databaseDatabaseRepresentation},
		"db_home_id":        Representation{RepType: Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":            Representation{RepType: Required, Create: `NONE`},
		"kms_key_migration": Representation{RepType: Required, Create: `true`},
		"kms_key_id":        Representation{RepType: Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
	}

	databaseDatabaseRepresentation = map[string]interface{}{
		"admin_password":   Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"db_name":          Representation{RepType: Required, Create: `myTestDb`},
		"character_set":    Representation{RepType: Optional, Create: `AL32UTF8`},
		"db_backup_config": RepresentationGroup{Optional, databaseDatabaseDbBackupConfigRepresentation},
		"db_unique_name":   Representation{RepType: Optional, Create: `myTestDb_12`},
		"db_workload":      Representation{RepType: Optional, Create: `OLTP`},
		"defined_tags":     Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   Representation{RepType: Optional, Create: `AL16UTF16`},
		"pdb_name":         Representation{RepType: Optional, Create: `pdbName`},
		// "tde_wallet_password": Representation{RepType: Optional, Create: `tdeWalletPassword`},	exadata doesn't support it.
	}
	databaseDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":     Representation{RepType: Optional, Create: `true`},
		"auto_backup_window":      Representation{RepType: Optional, Create: `SLOT_TWO`, Update: `SLOT_THREE`},
		"recovery_window_in_days": Representation{RepType: Optional, Create: `10`, Update: `30`},
	}
	databaseDatabaseDbBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type": Representation{RepType: Required, Create: `NFS`},
		"id":   Representation{RepType: Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
	}

	DatabaseResourceDependencies = ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig + KeyResourceDependencyConfig +
		GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Required, Create, dbHomeRepresentationSourceNone)
)

// issue-routing-tag: database/default
func TestDatabaseDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_database.test_database"
	datasourceName := "data.oci_database_databases.test_databases"
	singularDatasourceName := "data.oci_database_database.test_database"

	var resId, resId2 string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DatabaseResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Create, databaseRepresentation), "database", "database", t)

	ResourceTest(t, testAccCheckDatabaseDatabaseDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + DatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
			),
		},
		// verify migrate kms_key
		{
			Config: config + compartmentIdVariableStr + DatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseRepresentationMigration),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
			),
		},
		// delete before next create
		{
			Config: config + compartmentIdVariableStr + DatabaseResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Create, databaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_TWO"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "myTestDb"),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", "myTestDb_12"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttrSet(resourceName, "db_unique_name"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_enabled", "true"),
				//resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_THREE"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "30"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "myTestDb"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttrSet(resourceName, "db_unique_name"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_databases", Optional, Update, databaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),
				resource.TestCheckResourceAttr(datasourceName, "db_name", "myTestDb"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.character_set"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "databases.0.db_backup_config.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_home_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_unique_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_workload"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.is_cdb"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.ncharacter_set"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.pdb_name"),
				//resource.TestCheckResourceAttrSet(datasourceName, "databases.0.source_database_point_in_time_recovery_timestamp"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "character_set"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_unique_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_workload"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cdb"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_backup_timestamp"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ncharacter_set"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pdb_name"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "source_database_point_in_time_recovery_timestamp"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DatabaseResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"database",
				"db_version",
				"kms_key_migration",
				"kms_key_rotation",
				"kms_key_version_id",
				"source",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_database" {
			noResourceFound = false
			request := oci_database.GetDatabaseRequest{}

			tmp := rs.Primary.ID
			request.DatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")

			response, err := client.GetDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DatabaseLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("DatabaseDatabase") {
		resource.AddTestSweepers("DatabaseDatabase", &resource.Sweeper{
			Name:         "DatabaseDatabase",
			Dependencies: DependencyGraph["database"],
			F:            sweepDatabaseDatabaseResource,
		})
	}
}

func sweepDatabaseDatabaseResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	databaseIds, err := getDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseId := range databaseIds {
		if ok := SweeperDefaultResourceId[databaseId]; !ok {
			deleteDatabaseRequest := oci_database.DeleteDatabaseRequest{}

			deleteDatabaseRequest.DatabaseId = &databaseId

			deleteDatabaseRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDatabase(context.Background(), deleteDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting Database %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &databaseId, databaseSweepWaitCondition, time.Duration(3*time.Minute),
				databaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listDatabasesRequest := oci_database.ListDatabasesRequest{}
	listDatabasesRequest.CompartmentId = &compartmentId

	dbHomeIds, err := getDbHomeIds(compartment)
	if err != nil {
		return resourceIds, err
	}
	for _, dbHomeId := range dbHomeIds {
		listDatabasesRequest.DbHomeId = &dbHomeId
		listDatabasesRequest.LifecycleState = oci_database.DatabaseSummaryLifecycleStateAvailable
		listDatabasesResponse, err := databaseClient.ListDatabases(context.Background(), listDatabasesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Database list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, database := range listDatabasesResponse.Items {
			id := *database.Id
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseId", id)
		}
	}
	return resourceIds, nil
}

func databaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseResponse, ok := response.Response.(oci_database.GetDatabaseResponse); ok {
		return databaseResponse.LifecycleState != oci_database.DatabaseLifecycleStateTerminated
	}
	return false
}

func databaseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetDatabase(context.Background(), oci_database.GetDatabaseRequest{
		DatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
