// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsDatabaseToolsIdentityRequiredOnlyResource = DatabaseToolsDatabaseToolsIdentityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsIdentityRepresentation)

	DatabaseToolsDatabaseToolsIdentityResourceConfig = DatabaseToolsDatabaseToolsIdentityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsIdentityRepresentation)

	DatabaseToolsDatabaseToolsIdentitySingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_identity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_identity.test_database_tools_identity.id}`},
	}

	DatabaseToolsDatabaseToolsIdentityDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `Identity1`, Update: `displayName2`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_DATABASE_RESOURCE_PRINCIPAL`}},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsIdentityDataSourceFilterRepresentation}}
	DatabaseToolsDatabaseToolsIdentityDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_identity.test_database_tools_identity.id}`}},
	}

	DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsIdentityRepresentation = map[string]interface{}{ // This may vary depending on the tenancy settings
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatabaseToolsDatabaseToolsIdentityRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `Key1`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `Identity1`, Update: `displayName2`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE_RESOURCE_PRINCIPAL`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{`${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}`: `value`}, Update: map[string]string{`${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}`: `updatedValue`}},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsIdentityRepresentation},
	}

	DatabaseToolsTestAdbWalletContent = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		// The secret_id here is associated with a secret that was manually created for a manually created DB.
		"secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.wallet_secret_id}`},
	}
	DatabaseToolsTestADBWalletKeyStore = map[string]interface{}{
		"key_store_content": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsTestAdbWalletContent},
		"key_store_type":    acctest.Representation{RepType: acctest.Required, Create: `SSO`},
	}

	DatabaseToolsOracleResourcePrincipalConnection = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `tf_connection_for_identity_name`},
		"type":              acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"connection_string": acctest.Representation{RepType: acctest.Required, Create: `${var.connection_string}`},
		"runtime_identity":  acctest.Representation{RepType: acctest.Required, Create: `RESOURCE_PRINCIPAL`},
		"runtime_support":   acctest.Representation{RepType: acctest.Required, Create: `SUPPORTED`},
		"user_name":         acctest.Representation{RepType: acctest.Required, Create: `admin`},
		"user_password":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation},
		"key_stores":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsTestADBWalletKeyStore},
	}

	DatabaseToolsDatabaseToolsIdentityResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsOracleResourcePrincipalConnection) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsIdentityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsIdentityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	allVars := databaseToolsStandardVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_identity.test_database_tools_identity"
	datasourceName := "data.oci_database_tools_database_tools_identities.test_database_tools_identities"
	singularDatasourceName := "data.oci_database_tools_database_tools_identity.test_database_tools_identity"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+allVars+DatabaseToolsDatabaseToolsIdentityResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsIdentityRepresentation), "databasetools", "databaseToolsIdentity", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsIdentityDestroy, []resource.TestStep{
		// 1. verify Create
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsIdentityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsIdentityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_key", "Key1"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Identity1"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE_RESOURCE_PRINCIPAL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// 2. delete before next Create
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsIdentityResourceDependencies,
		},
		// 3. verify Create with optionals
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsIdentityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsIdentityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_key", "Key1"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Identity1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE_RESOURCE_PRINCIPAL"),

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

		// 4. verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + allVars + compartmentIdUVariableStr + DatabaseToolsDatabaseToolsIdentityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsIdentityRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "credential_key", "Key1"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Identity1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE_RESOURCE_PRINCIPAL"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// 5. verify updates to updatable parameters
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsIdentityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsIdentityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_key", "Key1"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE_RESOURCE_PRINCIPAL"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// 6. verify datasource
		{
			Config: config + allVars +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_identities", "test_database_tools_identities", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsIdentityDataSourceRepresentation) +
				DatabaseToolsDatabaseToolsIdentityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsIdentityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "database_tools_identity_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_identity_collection.0.items.#", "1"),
			),
		},
		// 7. verify singular datasource
		{
			Config: config + allVars +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_identity", "test_database_tools_identity", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsIdentitySingularDataSourceRepresentation) +
				DatabaseToolsDatabaseToolsIdentityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_identity_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_key", "Key1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ORACLE_DATABASE_RESOURCE_PRINCIPAL"),
			),
		},
		// 8. verify resource import
		{
			Config:                  config + DatabaseToolsDatabaseToolsIdentityRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseToolsDatabaseToolsIdentityDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_identity" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsIdentityRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsIdentityId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsIdentity(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.DatabaseToolsIdentityLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseToolsDatabaseToolsIdentity") {
		resource.AddTestSweepers("DatabaseToolsDatabaseToolsIdentity", &resource.Sweeper{
			Name:         "DatabaseToolsDatabaseToolsIdentity",
			Dependencies: acctest.DependencyGraph["databaseToolsIdentity"],
			F:            sweepDatabaseToolsDatabaseToolsIdentityResource,
		})
	}
}

func sweepDatabaseToolsDatabaseToolsIdentityResource(compartment string) error {
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	databaseToolsIdentityIds, err := getDatabaseToolsDatabaseToolsIdentityIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsIdentityId := range databaseToolsIdentityIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsIdentityId]; !ok {
			deleteDatabaseToolsIdentityRequest := oci_database_tools.DeleteDatabaseToolsIdentityRequest{}

			deleteDatabaseToolsIdentityRequest.DatabaseToolsIdentityId = &databaseToolsIdentityId

			deleteDatabaseToolsIdentityRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsIdentity(context.Background(), deleteDatabaseToolsIdentityRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsIdentity %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsIdentityId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseToolsIdentityId, DatabaseToolsDatabaseToolsIdentitySweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseToolsDatabaseToolsIdentitySweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsDatabaseToolsIdentityIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsIdentityId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()

	listDatabaseToolsIdentitiesRequest := oci_database_tools.ListDatabaseToolsIdentitiesRequest{}
	listDatabaseToolsIdentitiesRequest.CompartmentId = &compartmentId
	listDatabaseToolsIdentitiesRequest.LifecycleState = oci_database_tools.ListDatabaseToolsIdentitiesLifecycleStateActive
	listDatabaseToolsIdentitiesResponse, err := databaseToolsClient.ListDatabaseToolsIdentities(context.Background(), listDatabaseToolsIdentitiesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsIdentity list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsIdentity := range listDatabaseToolsIdentitiesResponse.Items {
		id := *databaseToolsIdentity.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsIdentityId", id)
	}
	return resourceIds, nil
}

func DatabaseToolsDatabaseToolsIdentitySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseToolsIdentityResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsIdentityResponse); ok {
		return databaseToolsIdentityResponse.GetLifecycleState() != oci_database_tools.DatabaseToolsIdentityLifecycleStateDeleted
	}
	return false
}

func DatabaseToolsDatabaseToolsIdentitySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseToolsClient().GetDatabaseToolsIdentity(context.Background(), oci_database_tools.GetDatabaseToolsIdentityRequest{
		DatabaseToolsIdentityId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
