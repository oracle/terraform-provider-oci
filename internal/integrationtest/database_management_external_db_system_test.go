// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalDbSystemRequiredOnlyResource = DatabaseManagementExternalDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemRepresentation)

	DatabaseManagementExternalDbSystemResourceConfig = DatabaseManagementExternalDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemRepresentation)

	DatabaseManagementDatabaseManagementExternalDbSystemSingularDataSourceRepresentation = map[string]interface{}{
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_db_system.test_external_db_system.id}`},
	}

	DatabaseManagementDatabaseManagementExternalDbSystemDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `EXAMPLE-displayName-Value`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDataSourceFilterRepresentation}}
	DatabaseManagementExternalDbSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_db_system.test_external_db_system.id}`}},
	}

	DatabaseManagementExternalDbSystemDiscoveryCdbPatchOperationsRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseManagementExternalDbSystemDiscoveryRepresentation,
		map[string]interface{}{
			"patch_operations": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDiscoveryCdbPatchRepresentation}},
		},
	)
	DatabaseManagementExternalDbSystemDiscoveryPdbPatchOperationsRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseManagementExternalDbSystemDiscoveryRepresentation,
		map[string]interface{}{
			"patch_operations": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDiscoveryPdbPatchRepresentation}},
		},
	)
	DatabaseManagementExternalDbSystemDiscoveryCdbPatchRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `MERGE`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `discoveredComponents[?componentType == 'DATABASE'] | [0]`},
		"value":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDiscoveryPatchMergeOperationValueRepresentation},
	}
	DatabaseManagementExternalDbSystemDiscoveryPatchMergeOperationValueRepresentation = map[string]interface{}{
		"is_selected_for_monitoring": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"connector":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDatabaseConnectorRepresentation},
	}
	DatabaseManagementExternalDbSystemDiscoveryPdbPatchRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `MERGE`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `discoveredComponents[?componentType == 'DATABASE'].pluggableDatabases`},
		"value": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"is_selected_for_monitoring": acctest.Representation{RepType: acctest.Required, Create: `false`},
		}},
	}
	DatabaseManagementExternalDbSystemDatabaseConnectorRepresentation = map[string]interface{}{
		"connector_type":  acctest.Representation{RepType: acctest.Required, Create: `MACS`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `EXAMPLE-displayName-Value`},
		"agent_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"connection_info": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDatabaseConnectionInfoRepresentation},
	}
	DatabaseManagementDatabaseConnectionInfoRepresentation = map[string]interface{}{
		"component_type":         acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
		"connection_string":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDatabaseConnectionStringRepresentation},
		"connection_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDatabaseConnectionCredentialsRepresentation},
	}
	DatabaseManagementDatabaseConnectionStringRepresentation = map[string]interface{}{
		"host_name": acctest.Representation{RepType: acctest.Required, Create: `${var.db_host_name}`},
		"port":      acctest.Representation{RepType: acctest.Required, Create: `${var.db_port}`},
		"protocol":  acctest.Representation{RepType: acctest.Required, Create: `TCP`},
		"service":   acctest.Representation{RepType: acctest.Required, Create: `${var.db_service_name}`},
	}
	DatabaseManagementDatabaseConnectionCredentialsRepresentation = map[string]interface{}{
		"credential_type":    acctest.Representation{RepType: acctest.Required, Create: `DETAILS`},
		"credential_name":    acctest.Representation{RepType: acctest.Required, Create: `${var.db_credential_name}`},
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `${var.db_user_name}`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_password_secret_id}`},
		"role":               acctest.Representation{RepType: acctest.Required, Create: `NORMAL`},
	}
	DatabaseManagementExternalDbSystemDisableDatabaseManagementRepresentation = map[string]interface{}{
		"external_db_system_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_db_system.test_external_db_system.id}`},
		"license_model":              acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"enable_database_management": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	ignoreLicenseModelChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`database_management_config`, `stack_monitoring_config`, `defined_tags`}},
	}
	DatabaseManagementExternalDbSystemRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_system_discovery_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_db_system_discovery.test_external_db_system_discovery.id}`},
		"database_management_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDatabaseManagementConfigRepresentation},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		// the SM service's preprod env is pointing to prod DBM env causing failures.
		// Commenting out the stack monitoring related testing pending env availability.
		//"stack_monitoring_config":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemStackMonitoringConfigRepresentation},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `EXAMPLE-displayName-Value`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreLicenseModelChangesRepresentation},
	}
	DatabaseManagementExternalDbSystemDatabaseManagementConfigRepresentation = map[string]interface{}{
		"license_model": acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
	}
	DatabaseManagementExternalDbSystemStackMonitoringConfigRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	DatabaseManagementExternalDbSystemDisableStackMonitoringRepresentation = map[string]interface{}{
		"external_db_system_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_db_system.test_external_db_system.id}`},
		"is_enabled":              acctest.Representation{RepType: acctest.Required, Create: `false`},
		"enable_stack_monitoring": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	DatabaseManagementExternalDbSystemResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDbSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDbSystemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	agentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)

	dbHostName := utils.GetEnvSettingWithBlankDefault("db_host_name")
	dbHostNameVariableStr := fmt.Sprintf("variable \"db_host_name\" { default = \"%s\" }\n", dbHostName)

	dbPort := utils.GetEnvSettingWithBlankDefault("db_port")
	dbPortVariableStr := fmt.Sprintf("variable \"db_port\" { default = \"%s\" }\n", dbPort)

	dbServiceName := utils.GetEnvSettingWithBlankDefault("db_service_name")
	dbServiceNameVariableStr := fmt.Sprintf("variable \"db_service_name\" { default = \"%s\" }\n", dbServiceName)

	dbCredentialName := utils.GetEnvSettingWithBlankDefault("db_credential_name")
	dbCredentialNameVariableStr := fmt.Sprintf("variable \"db_credential_name\" { default = \"%s\" }\n", dbCredentialName)

	dbUserName := utils.GetEnvSettingWithBlankDefault("db_user_name")
	dbUserNameVariableStr := fmt.Sprintf("variable \"db_user_name\" { default = \"%s\" }\n", dbUserName)

	dbPasswordSecretId := utils.GetEnvSettingWithBlankDefault("db_password_secret_id")
	dbPasswordSecretIdVariableStr := fmt.Sprintf("variable \"db_password_secret_id\" { default = \"%s\" }\n", dbPasswordSecretId)

	discoveryResourceName := "oci_database_management_external_db_system_discovery.test_external_db_system_discovery"
	resourceName := "oci_database_management_external_db_system.test_external_db_system"
	datasourceName := "data.oci_database_management_external_db_systems.test_external_db_systems"
	singularDatasourceName := "data.oci_database_management_external_db_system.test_external_db_system"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+DatabaseManagementExternalDbSystemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Optional, acctest.Create, DatabaseManagementExternalDbSystemRepresentation), "databasemanagement", "externalDbSystem", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementExternalDbSystemDestroy, []resource.TestStep{
		// Patch discovery and add connector to CDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + dbHostNameVariableStr + dbPortVariableStr + dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr + DatabaseManagementExternalDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery",
					acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryCdbPatchOperationsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(discoveryResourceName, "agent_id"),
				resource.TestCheckResourceAttr(discoveryResourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, discoveryResourceName, "id")
					return err
				},
			),
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + dbHostNameVariableStr + dbPortVariableStr + dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr + DatabaseManagementExternalDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery",
					"test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_discovery_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + agentIdVariableStr + dbHostNameVariableStr + dbPortVariableStr + dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr + DatabaseManagementExternalDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery",
					"test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseManagementExternalDbSystemRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "database_management_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_management_config.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_discovery_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EXAMPLE-displayName-Value"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "stack_monitoring_config.#", "1"),
				// Commenting out the stack monitoring related testing pending env availability.
				//resource.TestCheckResourceAttr(resourceName, "stack_monitoring_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + agentIdVariableStr + dbHostNameVariableStr + dbPortVariableStr + dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr + DatabaseManagementExternalDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery",
					"test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_management_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_management_config.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_discovery_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "stack_monitoring_config.#", "1"),
				// Commenting out the stack monitoring related testing pending env availability.
				//resource.TestCheckResourceAttr(resourceName, "stack_monitoring_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_systems", "test_external_db_systems", acctest.Optional, acctest.Update, DatabaseManagementDatabaseManagementExternalDbSystemDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + dbHostNameVariableStr + dbPortVariableStr + dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr + DatabaseManagementExternalDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery",
					"test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "external_db_system_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_db_system_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbSystemSingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + dbHostNameVariableStr + dbPortVariableStr + dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr + DatabaseManagementExternalDbSystemResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery",
					"test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryPdbPatchOperationsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovery_agent_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cluster"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stack_monitoring_config.#", "1"),
				// Commenting out the stack monitoring related testing pending env availability.
				//resource.TestCheckResourceAttr(singularDatasourceName, "stack_monitoring_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// disable DB Management
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + dbHostNameVariableStr + dbPortVariableStr + dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr + DatabaseManagementExternalDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery",
					"test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_database_managements_management", "test_external_db_system_database_managements_management", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemDisableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_discovery_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Commenting out the stack monitoring related testing pending env availability.
		// 		// disable Stack Monitoring
		// 		{
		// 			Config: config + compartmentIdVariableStr + agentIdVariableStr + dbHostNameVariableStr + dbPortVariableStr + dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr + DatabaseManagementExternalDbSystemResourceDependencies +
		// 				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery",
		// 					"test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryPdbPatchOperationsRepresentation) +
		// 				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemRepresentation) +
		// 				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_database_managements_management", "test_external_db_system_database_managements_management", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemDisableDatabaseManagementRepresentation) +
		// 				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_stack_monitorings_management", "test_external_db_system_stack_monitoring_management", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemDisableStackMonitoringRepresentation),
		// 			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		// 				resource.TestCheckResourceAttrSet(resourceName, "db_system_discovery_id"),
		//
		// 				func(s *terraform.State) (err error) {
		// 					resId, err = acctest.FromInstanceState(s, resourceName, "id")
		// 					return err
		// 				},
		// 			),
		// 		},

		// verify resource import
		{
			Config:                  config + DatabaseManagementExternalDbSystemRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseManagementExternalDbSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_external_db_system" {
			noResourceFound = false
			request := oci_database_management.GetExternalDbSystemRequest{}

			tmp := rs.Primary.ID
			request.ExternalDbSystemId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetExternalDbSystem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.ExternalDbSystemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementExternalDbSystem") {
		resource.AddTestSweepers("DatabaseManagementExternalDbSystem", &resource.Sweeper{
			Name:         "DatabaseManagementExternalDbSystem",
			Dependencies: acctest.DependencyGraph["externalDbSystem"],
			F:            sweepDatabaseManagementExternalDbSystemResource,
		})
	}
}

func sweepDatabaseManagementExternalDbSystemResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	externalDbSystemIds, err := getDatabaseManagementExternalDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, externalDbSystemId := range externalDbSystemIds {
		if ok := acctest.SweeperDefaultResourceId[externalDbSystemId]; !ok {
			deleteExternalDbSystemRequest := oci_database_management.DeleteExternalDbSystemRequest{}

			deleteExternalDbSystemRequest.ExternalDbSystemId = &externalDbSystemId

			deleteExternalDbSystemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteExternalDbSystem(context.Background(), deleteExternalDbSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalDbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalDbSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalDbSystemId, DatabaseManagementExternalDbSystemSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementExternalDbSystemSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementExternalDbSystemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalDbSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listExternalDbSystemsRequest := oci_database_management.ListExternalDbSystemsRequest{}
	listExternalDbSystemsRequest.CompartmentId = &compartmentId
	listExternalDbSystemsResponse, err := dbManagementClient.ListExternalDbSystems(context.Background(), listExternalDbSystemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalDbSystem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalDbSystem := range listExternalDbSystemsResponse.Items {
		id := *externalDbSystem.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalDbSystemId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementExternalDbSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalDbSystemResponse, ok := response.Response.(oci_database_management.GetExternalDbSystemResponse); ok {
		return externalDbSystemResponse.LifecycleState != oci_database_management.ExternalDbSystemLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementExternalDbSystemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetExternalDbSystem(context.Background(), oci_database_management.GetExternalDbSystemRequest{
		ExternalDbSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
