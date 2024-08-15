// // // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // // Licensed under the Mozilla Public License v2.0
package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseMigrationConnectionRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, DatabaseMigrationConnectionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Required, acctest.Create, DatabaseMigrationConnectionOracleRepresentation)

	DatabaseMigrationConnectionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, DatabaseMigrationConnectionRepresentation)

	DatabaseMigrationConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_connection.test_connection.id}`},
	}

	DatabaseMigrationConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`MYSQL`}},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"technology_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`AMAZON_RDS_MYSQL`}},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationConnectionDataSourceFilterRepresentation},
	}

	DatabaseMigrationConnectionDataSourceOracleRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE`}},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"technology_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`AMAZON_RDS_ORACLE`}},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationConnectionDataSourceFilterOracleRepresentation},
	}

	DatabaseMigrationConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_connection.test_connection.id}`}},
	}

	DatabaseMigrationConnectionDataSourceFilterOracleRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_connection.test_connection_rds.id}`}},
	}

	DatabaseMigrationConnectionRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type":       acctest.Representation{RepType: acctest.Required, Create: `MYSQL`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"key_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"password":              acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"technology_type":       acctest.Representation{RepType: acctest.Required, Create: `AMAZON_RDS_MYSQL`},
		"username":              acctest.Representation{RepType: acctest.Required, Create: `ggfe`, Update: `ggfe`},
		"vault_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.kms_vault_id}`},
		"additional_attributes": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationConnectionAdditionalAttributesRepresentation},
		"database_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.source_database_mysql_id}`},
		"database_name":         acctest.Representation{RepType: acctest.Required, Create: `ggfe`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"host":                  acctest.Representation{RepType: acctest.Required, Create: `254.249.0.0`, Update: `254.249.0.0`},
		"nsg_ids":               acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.nsg_mysql_id}`}},
		"port":                  acctest.Representation{RepType: acctest.Required, Create: `3306`, Update: `3306`},
		"replication_password":  acctest.Representation{RepType: acctest.Optional, Create: `replicationPassword`, Update: `replicationPassword2`},
		"replication_username":  acctest.Representation{RepType: acctest.Optional, Create: `replicationUsername`, Update: `replicationUsername`},
		"security_protocol":     acctest.Representation{RepType: acctest.Required, Create: `PLAIN`, Update: `PLAIN`},
		"ssh_host":              acctest.Representation{RepType: acctest.Required, Create: `sshHost`, Update: `sshHost2`},
		"ssh_key":               acctest.Representation{RepType: acctest.Required, Create: `sshKey`, Update: `sshKey2`},
		"ssh_sudo_location":     acctest.Representation{RepType: acctest.Optional, Create: `sshSudoLocation`, Update: `sshSudoLocation2`},
		"ssh_user":              acctest.Representation{RepType: acctest.Optional, Create: `sshUser`, Update: `sshUser2`},
		"subnet_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_mysql_id}`},
		"wallet":                acctest.Representation{RepType: acctest.Optional, Create: `wallet`, Update: `wallet2`},
	}

	DatabaseMigrationConnectionOracleRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type":       acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"key_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"password":              acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"technology_type":       acctest.Representation{RepType: acctest.Required, Create: `AMAZON_RDS_ORACLE`},
		"username":              acctest.Representation{RepType: acctest.Required, Create: `ggfe`},
		"vault_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.kms_vault_id}`},
		"additional_attributes": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationConnectionAdditionalAttributesRepresentation},
		"connection_string":     acctest.Representation{RepType: acctest.Required, Create: `10.10.10.10:10/test`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"nsg_ids":               acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.nsg_mysql_id}`, `${var.nsg_mysql_id2}`}, Update: []string{`${var.nsg_mysql_id}`, `${var.nsg_mysql_id2}`}},
		"replication_password":  acctest.Representation{RepType: acctest.Optional, Create: `replicationPassword`, Update: `replicationPassword2`},
		"replication_username":  acctest.Representation{RepType: acctest.Optional, Create: `replicationUsername`},
		"wallet":                acctest.Representation{RepType: acctest.Required, Create: `wallet`, Update: `wallet2`},
	}

	connectionRepresentationTarget = map[string]interface{}{
		"database_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.source_database_mysql_id}`},
		"password":             acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"replication_password": acctest.Representation{RepType: acctest.Required, Create: `replicationPassword`, Update: `replicationPassword2`},
		"replication_username": acctest.Representation{RepType: acctest.Required, Create: `replicationUsername`, Update: `replicationUsername2`},
		"ssh_host":             acctest.Representation{RepType: acctest.Required, Create: `sshHost`, Update: `sshHost2`},
		"ssh_key":              acctest.Representation{RepType: acctest.Required, Create: `sshKey`, Update: `sshKey2`},
		"ssh_sudo_location":    acctest.Representation{RepType: acctest.Required, Create: `sshSudoLocation`, Update: `sshSudoLocation2`},
		"ssh_user":             acctest.Representation{RepType: acctest.Required, Create: `sshUser`, Update: `sshUser2`},
		"username":             acctest.Representation{RepType: acctest.Required, Create: `ggfe`, Update: `ggfe`},
	}

	DatabaseMigrationConnectionAdditionalAttributesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
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

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsVaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_id")
	kmsVaultIdVariableStr := fmt.Sprintf("variable \"kms_vault_id\" { default = \"%s\" }\n", kmsVaultId)

	sourceDBId := utils.GetEnvSettingWithBlankDefault("source_database_mysql_id")
	sourceDBIdVariableStr := fmt.Sprintf("variable \"source_database_mysql_id\" { default = \"%s\" }\n", sourceDBId)

	nsgId := utils.GetEnvSettingWithBlankDefault("nsg_mysql_id")
	nsgIdVariableStr := fmt.Sprintf("variable \"nsg_mysql_id\" { default = \"%s\" }\n", nsgId)

	nsgId2 := utils.GetEnvSettingWithBlankDefault("nsg_mysql_id2")
	nsgIdVariableStr2 := fmt.Sprintf("variable \"nsg_mysql_id2\" { default = \"%s\" }\n", nsgId2)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_mysql_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_mysql_id\" { default = \"%s\" }\n", subnetId)

	resourceName := "oci_database_migration_connection.test_connection"
	datasourceName := "data.oci_database_migration_connections.test_connections"
	singularDatasourceName := "data.oci_database_migration_connection.test_connection"
	resourceNameRDS := "oci_database_migration_connection.test_connection_rds"
	datasourceRDSName := "data.oci_database_migration_connections.test_connections_rds"

	var resId, resId2, resId3 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+kmsKeyIdVariableStr+kmsVaultIdVariableStr+compartmentIdUVariableStr+sourceDBIdVariableStr+nsgIdVariableStr+subnetIdVariableStr+nsgIdVariableStr2+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create, DatabaseMigrationConnectionRepresentation)+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Create, DatabaseMigrationConnectionOracleRepresentation), "databasemigration", "connection", t)

	acctest.ResourceTest(t, testAccCheckDatabaseMigrationConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceDBIdVariableStr + nsgIdVariableStr + subnetIdVariableStr + nsgIdVariableStr2 +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, DatabaseMigrationConnectionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Required, acctest.Create, DatabaseMigrationConnectionOracleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "MYSQL"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "security_protocol", "PLAIN"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "AMAZON_RDS_MYSQL"),
				resource.TestCheckResourceAttr(resourceName, "username", "ggfe"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				resource.TestCheckResourceAttr(resourceNameRDS, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameRDS, "connection_type", "ORACLE"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "connection_string"),
				resource.TestCheckResourceAttr(resourceNameRDS, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "key_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceNameRDS, "technology_type", "AMAZON_RDS_ORACLE"),
				resource.TestCheckResourceAttr(resourceNameRDS, "username", "ggfe"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					resId2, err = acctest.FromInstanceState(s, resourceNameRDS, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceDBIdVariableStr + nsgIdVariableStr + subnetIdVariableStr + nsgIdVariableStr2,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + compartmentIdUVariableStr + sourceDBIdVariableStr + nsgIdVariableStr + subnetIdVariableStr + nsgIdVariableStr2 +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create, DatabaseMigrationConnectionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Create, DatabaseMigrationConnectionOracleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "MYSQL"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "host", "254.249.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "replication_password", "replicationPassword"),
				resource.TestCheckResourceAttr(resourceName, "replication_username", "replicationUsername"),
				resource.TestCheckResourceAttr(resourceName, "security_protocol", "PLAIN"),
				resource.TestCheckResourceAttr(resourceName, "ssh_host", "sshHost"),
				resource.TestCheckResourceAttr(resourceName, "ssh_key", "sshKey"),
				resource.TestCheckResourceAttr(resourceName, "ssh_sudo_location", "sshSudoLocation"),

				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "AMAZON_RDS_MYSQL"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "username", "ggfe"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "wallet", "wallet"),

				resource.TestCheckResourceAttr(resourceNameRDS, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameRDS, "connection_type", "ORACLE"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "connection_string"),
				resource.TestCheckResourceAttr(resourceNameRDS, "description", "description"),
				resource.TestCheckResourceAttr(resourceNameRDS, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "key_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceNameRDS, "replication_password", "replicationPassword"),
				resource.TestCheckResourceAttr(resourceNameRDS, "replication_username", "replicationUsername"),

				resource.TestCheckResourceAttrSet(resourceNameRDS, "state"),
				resource.TestCheckResourceAttr(resourceNameRDS, "technology_type", "AMAZON_RDS_ORACLE"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "time_updated"),
				resource.TestCheckResourceAttr(resourceNameRDS, "username", "ggfe"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "wallet", "wallet"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					resId2, err = acctest.FromInstanceState(s, resourceNameRDS, "id")
					time.Sleep(1 * time.Minute)
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceDBIdVariableStr + nsgIdVariableStr + subnetIdVariableStr + nsgIdVariableStr2 +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseMigrationConnectionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseMigrationConnectionOracleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "MYSQL"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "host", "254.249.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "replication_password", "replicationPassword"),
				resource.TestCheckResourceAttr(resourceName, "replication_username", "replicationUsername"),
				resource.TestCheckResourceAttr(resourceName, "security_protocol", "PLAIN"),
				resource.TestCheckResourceAttr(resourceName, "ssh_host", "sshHost"),
				resource.TestCheckResourceAttr(resourceName, "ssh_key", "sshKey"),
				resource.TestCheckResourceAttr(resourceName, "ssh_sudo_location", "sshSudoLocation"),
				resource.TestCheckResourceAttr(resourceName, "ssh_user", "sshUser"),

				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "AMAZON_RDS_MYSQL"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "username", "ggfe"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "wallet", "wallet"),

				resource.TestCheckResourceAttr(resourceNameRDS, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceNameRDS, "connection_type", "ORACLE"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "connection_string"),
				resource.TestCheckResourceAttr(resourceNameRDS, "description", "description"),
				resource.TestCheckResourceAttr(resourceNameRDS, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "key_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceNameRDS, "replication_password", "replicationPassword"),
				resource.TestCheckResourceAttr(resourceNameRDS, "replication_username", "replicationUsername"),

				resource.TestCheckResourceAttrSet(resourceNameRDS, "state"),
				resource.TestCheckResourceAttr(resourceNameRDS, "technology_type", "AMAZON_RDS_ORACLE"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "time_updated"),
				resource.TestCheckResourceAttr(resourceNameRDS, "username", "ggfe"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "wallet", "wallet"),

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
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceDBIdVariableStr + nsgIdVariableStr + subnetIdVariableStr + nsgIdVariableStr2 +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, DatabaseMigrationConnectionRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Update, DatabaseMigrationConnectionOracleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "additional_attributes.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "MYSQL"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "host", "254.249.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "replication_password", "replicationPassword2"),
				resource.TestCheckResourceAttr(resourceName, "replication_username", "replicationUsername"),
				resource.TestCheckResourceAttr(resourceName, "security_protocol", "PLAIN"),
				resource.TestCheckResourceAttr(resourceName, "ssh_host", "sshHost2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_key", "sshKey2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_sudo_location", "sshSudoLocation2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_user", "sshUser2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "technology_type", "AMAZON_RDS_MYSQL"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "username", "ggfe"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "wallet", "wallet2"),

				resource.TestCheckResourceAttr(resourceNameRDS, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameRDS, "connection_type", "ORACLE"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "connection_string"),
				resource.TestCheckResourceAttr(resourceNameRDS, "description", "description2"),
				resource.TestCheckResourceAttr(resourceNameRDS, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "id"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "key_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceNameRDS, "replication_password", "replicationPassword2"),
				resource.TestCheckResourceAttr(resourceNameRDS, "replication_username", "replicationUsername"),

				resource.TestCheckResourceAttrSet(resourceNameRDS, "state"),
				resource.TestCheckResourceAttr(resourceNameRDS, "technology_type", "AMAZON_RDS_ORACLE"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "time_updated"),
				resource.TestCheckResourceAttr(resourceNameRDS, "username", "ggfe"),
				resource.TestCheckResourceAttrSet(resourceNameRDS, "vault_id"),
				resource.TestCheckResourceAttr(resourceNameRDS, "wallet", "wallet2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_connections", "test_connections", acctest.Optional, acctest.Update, DatabaseMigrationConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceDBIdVariableStr + nsgIdVariableStr + subnetIdVariableStr + nsgIdVariableStr2 +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Optional, acctest.Update, DatabaseMigrationConnectionRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_connections", "test_connections_rds", acctest.Optional, acctest.Update, DatabaseMigrationConnectionDataSourceOracleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_rds", acctest.Optional, acctest.Update, DatabaseMigrationConnectionOracleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "connection_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "technology_type.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "connection_collection.#", "1"),

				resource.TestCheckResourceAttr(datasourceRDSName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceRDSName, "connection_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceRDSName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceRDSName, "technology_type.#", "1"),

				resource.TestCheckResourceAttr(datasourceRDSName, "connection_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, DatabaseMigrationConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + DatabaseMigrationConnectionResourceConfig + sourceDBIdVariableStr + nsgIdVariableStr + subnetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "additional_attributes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "additional_attributes.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "additional_attributes.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_type", "MYSQL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host", "254.249.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port", "3306"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_protocol", "PLAIN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "technology_type", "AMAZON_RDS_MYSQL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseMigrationConnectionRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"ssl_ca",
				"ssl_cert",
				"ssl_crl",
				"ssl_key",
				"wallet",
				"database_id",
				"password",
				"replication_password",
				"replication_username",
				"ssh_host",
				"ssh_key",
				"ssh_sudo_location",
				"ssh_user",
				"username",
				"private_endpoint_id",
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
					string(oci_database_migration.ConnectionLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	connectionIds, err := getDatabaseMigrationConnectionIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &connectionId, DatabaseMigrationConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseMigrationConnectionSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getDatabaseMigrationConnectionIds(compartment string) ([]string, error) {
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
		id := *connection.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConnectionId", id)
	}
	return resourceIds, nil
}

func DatabaseMigrationConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if connectionResponse, ok := response.Response.(oci_database_migration.GetConnectionResponse); ok {
		return connectionResponse.GetLifecycleState() != oci_database_migration.ConnectionLifecycleStateDeleted
	}
	return false
}

func DatabaseMigrationConnectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseMigrationClient().GetConnection(context.Background(), oci_database_migration.GetConnectionRequest{
		ConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
