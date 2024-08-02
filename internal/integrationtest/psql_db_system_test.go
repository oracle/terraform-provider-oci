// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	//"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	//"strconv"

	//"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	//"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PsqlDbSystemRequiredOnlyResource = PsqlDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create, PsqlDbSystemRepresentation)

	PsqlDbSystemResourceConfig = PsqlDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Optional, acctest.Update, PsqlDbSystemRepresentation)

	PsqlDbSystemSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_psql_db_system.test_db_system.id}`},
	}

	PsqlDbSystemDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `terraform-test-2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemDataSourceFilterRepresentation},
	}
	PsqlDbSystemIdOnlyDataSourceRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_psql_db_system.test_db_system.id}`},
	}

	PsqlDbSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_psql_db_system.test_db_system.id}`}},
	}

	PsqlDbSystemRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_version":                  acctest.Representation{RepType: acctest.Required, Create: `14`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `test-terraform`, Update: `terraform-test-2`},
		"network_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemNetworkDetailsRepresentation},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `PostgreSQL.VM.Standard.E4.Flex.2.32GB`},
		"storage_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemStorageDetailsRepresentation},
		"config_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.config_id}`, Update: `${var.update_config_id}`},
		"apply_config":                acctest.Representation{RepType: acctest.Optional, Update: `RESTART`},
		"credentials":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemCredentialsRepresentation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `terrafrom test dbSystem`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key-2": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"instance_memory_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `32`},
		"instance_ocpu_count":         acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"instances_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemInstancesDetailsRepresentation},
		"management_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemManagementPolicyRepresentation},
		"source":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemSourceRepresentation},
		"system_type":                 acctest.Representation{RepType: acctest.Required, Create: `OCI_OPTIMIZED_STORAGE`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignorePsqlDefinedTagsChangesRepresentation},
	}
	ignorePsqlDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	PsqlDbSystemCredentialsRepresentation = map[string]interface{}{
		"password_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemCredentialsPasswordDetailsRepresentation},
		"username":         acctest.Representation{RepType: acctest.Required, Create: `user`},
	}

	PsqlDbSystemNetworkDetailsRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"nsg_ids":   acctest.Representation{RepType: acctest.Required, Create: []string{`${var.nsg_id}`}, Update: []string{`${var.update_nsg_id}`}},
		//"primary_db_endpoint_private_ip": acctest.Representation{RepType: acctest.Optional, Create: `primaryDbEndpointPrivateIp`},
	}
	PsqlDbSystemStorageDetailsRepresentation = map[string]interface{}{
		"availability_domain":   acctest.Representation{RepType: acctest.Required, Create: `gXfg:PHX-AD-1`},
		"is_regionally_durable": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"system_type":           acctest.Representation{RepType: acctest.Required, Create: `OCI_OPTIMIZED_STORAGE`},
		"iops":                  acctest.Representation{RepType: acctest.Optional, Create: `300000`},
	}

	PsqlDbSystemInstancesDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `Terraform federated test dbSystem`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `dbsystem-instance`},
		"private_ip":   acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.42`},
	}
	PsqlDbSystemManagementPolicyRepresentation = map[string]interface{}{
		"backup_policy":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemManagementPolicyBackupPolicyRepresentation},
		"maintenance_window_start": acctest.Representation{RepType: acctest.Optional, Create: `SUN 12:00`},
	}
	PsqlDbSystemSourceRepresentation = map[string]interface{}{
		"source_type":                        acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"backup_id":                          acctest.Representation{RepType: acctest.Optional, Create: ``},
		"is_having_restore_config_overrides": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	PsqlDbSystemCredentialsPasswordDetailsRepresentation = map[string]interface{}{
		"password_type": acctest.Representation{RepType: acctest.Required, Create: `PLAIN_TEXT`, Update: `PLAIN_TEXT`},
		"password":      acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
	}
	PsqlDbSystemManagementPolicyBackupPolicyRepresentation = map[string]interface{}{
		"backup_start":     acctest.Representation{RepType: acctest.Optional, Create: `02:00`, Update: `03:00`},
		"kind":             acctest.Representation{RepType: acctest.Optional, Create: `WEEKLY`},
		"retention_days":   acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"days_of_the_week": acctest.Representation{RepType: acctest.Optional, Create: []string{`SUNDAY`}},
	}
	// Test with Vault Secret, Backup Source and Monthly backup, in a AD
	PsqlDbSystemRepresentationMonthlyBackupVault = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_version":                  acctest.Representation{RepType: acctest.Required, Create: `14`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `test-terraform`, Update: `terraform-test-2`},
		"network_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemIpNetworkDetailsRepresentation},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `PostgreSQL.VM.Standard.E4.Flex.2.32GB`},
		"storage_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemRegionalStorageDetailsRepresentation},
		"credentials":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemVaultCredentialsRepresentation},
		"instance_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"instance_memory_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `32`},
		"instance_ocpu_count":         acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"management_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemMonthlyManagementPolicyRepresentation},
		"source":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemBackupSourceRepresentation},
		"system_type":                 acctest.Representation{RepType: acctest.Required, Create: `OCI_OPTIMIZED_STORAGE`},
	}

	// FLEX DbSystem
	PsqlDbSystemRepresentationFlexShape = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_version":                  acctest.Representation{RepType: acctest.Required, Create: `14`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `test-terraform`, Update: `terraform-test-2`},
		"credentials":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemCredentialsRepresentation},
		"network_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlFlexDbSystemIpNetworkDetailsRepresentation},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `PostgreSQL.VM.Standard.E4.Flex`},
		"storage_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemRegionalStorageDetailsRepresentation},
		"instance_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"instance_memory_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `12`},
		"instance_ocpu_count":         acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `3`},
		"management_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemMonthlyManagementPolicyRepresentation},
		"config_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.flex_config_id}`, Update: `${var.flex_update_config_id}`},
	}

	PsqlDbSystemIpNetworkDetailsRepresentation = map[string]interface{}{
		"subnet_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"nsg_ids":                        acctest.Representation{RepType: acctest.Required, Create: []string{}},
		"primary_db_endpoint_private_ip": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.110`, Update: `10.0.0.111`},
	}

	PsqlFlexDbSystemIpNetworkDetailsRepresentation = map[string]interface{}{
		"subnet_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"nsg_ids":                        acctest.Representation{RepType: acctest.Required, Create: []string{`${var.nsg_id}`}, Update: []string{`${var.update_nsg_id}`}},
		"primary_db_endpoint_private_ip": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.160`},
	}

	PsqlDbSystemRegionalStorageDetailsRepresentation = map[string]interface{}{
		"is_regionally_durable": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"system_type":           acctest.Representation{RepType: acctest.Required, Create: `OCI_OPTIMIZED_STORAGE`},
		"iops":                  acctest.Representation{RepType: acctest.Optional, Create: `300000`},
	}

	PsqlDbSystemVaultCredentialsRepresentation = map[string]interface{}{
		"password_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemVaultCredentialsPasswordDetailsRepresentation},
		"username":         acctest.Representation{RepType: acctest.Required, Create: `user`},
	}
	PsqlDbSystemVaultCredentialsPasswordDetailsRepresentation = map[string]interface{}{
		"password_type":  acctest.Representation{RepType: acctest.Required, Create: `VAULT_SECRET`},
		"secret_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"secret_version": acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}
	PsqlDbSystemMonthlyManagementPolicyRepresentation = map[string]interface{}{
		"backup_policy":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemMonthlyManagementPolicyBackupPolicyRepresentation},
		"maintenance_window_start": acctest.Representation{RepType: acctest.Optional, Create: `SUN 12:00`},
	}
	PsqlDbSystemMonthlyManagementPolicyBackupPolicyRepresentation = map[string]interface{}{
		"backup_start":      acctest.Representation{RepType: acctest.Optional, Create: `02:00`, Update: `03:00`},
		"kind":              acctest.Representation{RepType: acctest.Optional, Create: `MONTHLY`},
		"retention_days":    acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"days_of_the_month": acctest.Representation{RepType: acctest.Optional, Create: []string{`1`}, Update: []string{`2`}},
	}
	PsqlDbSystemBackupSourceRepresentation = map[string]interface{}{
		"source_type":                        acctest.Representation{RepType: acctest.Required, Create: `BACKUP`},
		"backup_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${var.backup_id}`},
		"is_having_restore_config_overrides": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	PsqlDbSystemResourceDependencies = AvailabilityDomainConfig + DefinedTagsDependencies
)

// issue-routing-tag: psql/default
func TestPsqlDbSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlDbSystemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	nsgId := utils.GetEnvSettingWithBlankDefault("nsg_id")
	nsgIdVariableStr := fmt.Sprintf("variable \"nsg_id\" { default = \"%s\" }\n", nsgId)

	nsgIdU := utils.GetEnvSettingWithBlankDefault("update_nsg_id")
	nsgIdUVariableStr := fmt.Sprintf("variable \"update_nsg_id\" { default = \"%s\" }\n", nsgIdU)

	configId := utils.GetEnvSettingWithBlankDefault("config_id")
	configIdVariableStr := fmt.Sprintf("variable \"config_id\" { default = \"%s\" }\n", configId)

	configIdU := utils.GetEnvSettingWithBlankDefault("update_config_id")
	configIdUVariableStr := fmt.Sprintf("variable \"update_config_id\" { default = \"%s\" }\n", configIdU)

	flexConfigId := utils.GetEnvSettingWithBlankDefault("flex_config_id")
	flexConfigIdVariableStr := fmt.Sprintf("variable \"flex_config_id\" { default = \"%s\" }\n", flexConfigId)

	//flexConfigIdU := utils.GetEnvSettingWithBlankDefault("flex_update_config_id")
	//flexConfigIdUVariableStr := fmt.Sprintf("variable \"flex_update_config_id\" { default = \"%s\" }\n", flexConfigIdU)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_ocid")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	backupId := utils.GetEnvSettingWithBlankDefault("backup_id")
	backupIdVariableStr := fmt.Sprintf("variable \"backup_id\" { default = \"%s\" }\n", backupId)

	resourceName := "oci_psql_db_system.test_db_system"
	test2_resourceName := "oci_psql_db_system.test_db_system_2"
	datasourceName := "data.oci_psql_db_systems.test_db_systems"
	singularDatasourceName := "data.oci_psql_db_system.test_db_system"
	test_flex_resourceName := "oci_psql_db_system.test_flex_db_system"

	var resId string
	var resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PsqlDbSystemResourceDependencies+configIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Optional, acctest.Create, PsqlDbSystemRepresentation), "psql", "dbSystem", t)

	acctest.ResourceTest(t, testAccCheckPsqlDbSystemDestroy, []resource.TestStep{
		// Flex Test
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies + flexConfigIdVariableStr + nsgIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_flex_db_system", acctest.Optional, acctest.Create, PsqlDbSystemRepresentationFlexShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(test_flex_resourceName, "compartment_id", compartmentId),
			),
		},

		// verify updates to updatable parameters
		/*
			{
				Config: config + compartmentIdVariableStr + flexConfigIdUVariableStr + subnetIdVariableStr + flexConfigIdVariableStr + PsqlDbSystemResourceDependencies + backupIdVariableStr + nsgIdUVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_flex_db_system", acctest.Optional, acctest.Update, PsqlDbSystemRepresentationFlexShape),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(test_flex_resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, test_flex_resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		*/

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies + flexConfigIdVariableStr + backupIdVariableStr + nsgIdVariableStr,
		},

		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies + nsgIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create, PsqlDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.password_type", "PLAIN_TEXT"),
				//resource.TestCheckResourceAttrSet(resourceName, "credentials.0.password_details.0.secret_id"),
				//resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.secret_version", "secretVersion"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.username", "user"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + nsgIdVariableStr + PsqlDbSystemResourceDependencies,
		},

		// verify Create checks for feilds not in Create with options
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies + vaultIdVariableStr + backupIdVariableStr + nsgIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system_2", acctest.Optional, acctest.Create, PsqlDbSystemRepresentationMonthlyBackupVault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(test2_resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(test2_resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(test2_resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(test2_resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(test2_resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(test2_resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(test2_resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttr(test2_resourceName, "storage_details.0.is_regionally_durable", "true"),
				resource.TestCheckResourceAttr(test2_resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttr(test2_resourceName, "network_details.0.primary_db_endpoint_private_ip", "10.0.0.110"),

				resource.TestCheckResourceAttrSet(test2_resourceName, "source.0.backup_id"),
				//resource.TestCheckResourceAttr(test2_resourceName, "source.0.is_having_restore_config_overrides", "false"),
				//resource.TestCheckResourceAttr(test2_resourceName, "source.0.source_type", "BACKUP"),
				//resource.TestCheckResourceAttr(test2_resourceName, "management_policy.0.backup_policy.0.days_of_the_month.#", "1"),

				resource.TestCheckResourceAttrSet(test2_resourceName, "credentials.0.password_details.0.secret_id"),
				resource.TestCheckResourceAttr(test2_resourceName, "credentials.0.password_details.0.secret_version", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies + backupIdVariableStr + nsgIdVariableStr,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies + configIdVariableStr + nsgIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Optional, acctest.Create, PsqlDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "config_id"),
				resource.TestCheckResourceAttr(resourceName, "credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.password_type", "PLAIN_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.username", "user"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "description", "terrafrom test dbSystem"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_memory_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "instance_ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.description", "Terraform federated test dbSystem"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.display_name", "dbsystem-instance"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.private_ip", "10.0.0.42"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.backup_start", "02:00"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.days_of_the_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.kind", "WEEKLY"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.retention_days", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.maintenance_window_start", "SUN 12:00"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.primary_db_endpoint_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.iops", "300000"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + subnetIdVariableStr + configIdVariableStr + PsqlDbSystemResourceDependencies + nsgIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PsqlDbSystemRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "config_id"),
				resource.TestCheckResourceAttr(resourceName, "credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.password_type", "PLAIN_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.username", "user"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "description", "terrafrom test dbSystem"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_memory_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "instance_ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.description", "Terraform federated test dbSystem"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.display_name", "dbsystem-instance"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.private_ip", "10.0.0.42"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.backup_start", "02:00"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.days_of_the_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.kind", "WEEKLY"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.retention_days", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.maintenance_window_start", "SUN 12:00"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.primary_db_endpoint_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.iops", "300000"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + configIdUVariableStr + subnetIdVariableStr + configIdVariableStr + PsqlDbSystemResourceDependencies + nsgIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Optional, acctest.Update, PsqlDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "config_id"),
				resource.TestCheckResourceAttr(resourceName, "credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password_details.0.password_type", "PLAIN_TEXT"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.username", "user"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "description", "terrafrom test dbSystem"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terraform-test-2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_memory_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "instance_ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.description", "Terraform federated test dbSystem"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.display_name", "dbsystem-instance"),
				resource.TestCheckResourceAttr(resourceName, "instances_details.0.private_ip", "10.0.0.42"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.backup_start", "03:00"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.days_of_the_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.kind", "WEEKLY"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.maintenance_window_start", "SUN 12:00"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.retention_days", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.primary_db_endpoint_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.iops", "300000"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_db_systems", "test_db_systems", acctest.Optional, acctest.Update, PsqlDbSystemDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies + configIdVariableStr + configIdUVariableStr + nsgIdUVariableStr + nsgIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Optional, acctest.Update, PsqlDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "terraform-test-2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "db_system_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_system_collection.0.items.#", "1"),
			),
		},
		// verify datasource with only id
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_db_systems", "test_db_systems", acctest.Optional, acctest.Update, PsqlDbSystemIdOnlyDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies + configIdVariableStr + configIdUVariableStr + nsgIdUVariableStr + nsgIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Optional, acctest.Update, PsqlDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "db_system_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_system_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create, PsqlDbSystemSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + configIdVariableStr + configIdUVariableStr + nsgIdUVariableStr + nsgIdVariableStr + PsqlDbSystemResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "excluded_fields.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "admin_username"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_version", "14.11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "terrafrom test dbSystem"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "terraform-test-2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_memory_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_ocpu_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instances.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.backup_start", "03:00"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.days_of_the_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.kind", "WEEKLY"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.backup_policy.0.retention_days", "1"),
				resource.TestCheckResourceAttr(resourceName, "management_policy.0.maintenance_window_start", "SUN 12:00"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.primary_db_endpoint_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_details.0.iops", "300000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_type", "OCI_OPTIMIZED_STORAGE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + subnetIdVariableStr + PsqlDbSystemRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"credentials",
				"instances_details",
				"patch_operations",
				"apply_config",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckPsqlDbSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).PostgresqlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_psql_db_system" {
			noResourceFound = false
			request := oci_psql.GetDbSystemRequest{}

			tmp := rs.Primary.ID
			request.DbSystemId = &tmp

			if value, ok := rs.Primary.Attributes["excluded_fields"]; ok {
				tmp := []oci_psql.GetDbSystemExcludedFieldsEnum{}
				in := []byte(value)
				json.Unmarshal(in, &tmp)
				request.ExcludedFields = tmp
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psql")

			response, err := client.GetDbSystem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_psql.DbSystemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("PsqlDbSystem") {
		resource.AddTestSweepers("PsqlDbSystem", &resource.Sweeper{
			Name:         "PsqlDbSystem",
			Dependencies: acctest.DependencyGraph["dbSystem"],
			F:            sweepPsqlDbSystemResource,
		})
	}
}

func sweepPsqlDbSystemResource(compartment string) error {
	postgresqlClient := acctest.GetTestClients(&schema.ResourceData{}).PostgresqlClient()
	dbSystemIds, err := getPsqlDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, dbSystemId := range dbSystemIds {
		if ok := acctest.SweeperDefaultResourceId[dbSystemId]; !ok {
			deleteDbSystemRequest := oci_psql.DeleteDbSystemRequest{}

			deleteDbSystemRequest.DbSystemId = &dbSystemId

			deleteDbSystemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psql")
			_, error := postgresqlClient.DeleteDbSystem(context.Background(), deleteDbSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting DbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dbSystemId, PsqlDbSystemSweepWaitCondition, time.Duration(3*time.Minute),
				PsqlDbSystemSweepResponseFetchOperation, "psql", true)
		}
	}
	return nil
}

func getPsqlDbSystemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DbSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	postgresqlClient := acctest.GetTestClients(&schema.ResourceData{}).PostgresqlClient()

	listDbSystemsRequest := oci_psql.ListDbSystemsRequest{}
	listDbSystemsRequest.CompartmentId = &compartmentId
	listDbSystemsResponse, err := postgresqlClient.ListDbSystems(context.Background(), listDbSystemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DbSystem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dbSystem := range listDbSystemsResponse.Items {
		id := *dbSystem.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DbSystemId", id)
	}
	return resourceIds, nil
}

func PsqlDbSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dbSystemResponse, ok := response.Response.(oci_psql.GetDbSystemResponse); ok {
		return dbSystemResponse.LifecycleState != oci_psql.DbSystemLifecycleStateDeleted
	}
	return false
}

func PsqlDbSystemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.PostgresqlClient().GetDbSystem(context.Background(), oci_psql.GetDbSystemRequest{
		DbSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
