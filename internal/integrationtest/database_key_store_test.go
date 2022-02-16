// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	KeyStoreRequiredOnlyResource = KeyStoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Required, acctest.Create, keyStoreRepresentation)

	KeyStoreResourceConfig = KeyStoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Update, keyStoreRepresentation)

	keyStoreSingularDataSourceRepresentation = map[string]interface{}{
		"key_store_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_key_store.test_key_store.id}`},
	}

	keyStoreDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: keyStoreDataSourceFilterRepresentation}}
	keyStoreDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_key_store.test_key_store.id}`}},
	}

	keyStoreRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `Key Store1`},
		"type_details":   acctest.RepresentationGroup{RepType: acctest.Required, Group: keyStoreTypeDetailsRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	keyStoreTypeDetailsRepresentation = map[string]interface{}{
		"admin_username": acctest.Representation{RepType: acctest.Required, Create: `username1`, Update: `adminUsername2`},
		"connection_ips": acctest.Representation{RepType: acctest.Required, Create: []string{`192.1.1.1`}, Update: []string{`192.1.1.1`, `192.1.1.2`}},
		"secret_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.okv_secret}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `ORACLE_KEY_VAULT`},
		"vault_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.kms_vault_id}`},
	}

	okvSecret            = utils.GetEnvSettingWithBlankDefault("okv_secret")
	OkvSecretVariableStr = fmt.Sprintf("variable \"okv_secret\" { default = \"%s\" }\n", okvSecret)

	KeyStoreResourceDependencies = DefinedTagsDependencies + KmsVaultIdVariableStr + OkvSecretVariableStr
)

// issue-routing-tag: database/ExaCC
func TestDatabaseKeyStoreResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseKeyStoreResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_key_store.test_key_store"
	datasourceName := "data.oci_database_key_stores.test_key_stores"
	singularDatasourceName := "data.oci_database_key_store.test_key_store"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KeyStoreResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, keyStoreRepresentation), "database", "keyStore", t)

	acctest.ResourceTest(t, testAccCheckDatabaseKeyStoreDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KeyStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Required, acctest.Create, keyStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key Store1"),
				resource.TestCheckResourceAttr(resourceName, "type_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type_details.0.admin_username", "username1"),
				resource.TestCheckResourceAttrSet(resourceName, "type_details.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "type_details.0.type", "ORACLE_KEY_VAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "type_details.0.vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KeyStoreResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + KeyStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, keyStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key Store1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type_details.0.admin_username", "username1"),
				resource.TestCheckResourceAttrSet(resourceName, "type_details.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "type_details.0.type", "ORACLE_KEY_VAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "type_details.0.vault_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + KeyStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(keyStoreRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key Store1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type_details.0.admin_username", "username1"),
				resource.TestCheckResourceAttrSet(resourceName, "type_details.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "type_details.0.type", "ORACLE_KEY_VAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "type_details.0.vault_id"),

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
			Config: config + compartmentIdVariableStr + KeyStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Update, keyStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key Store1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type_details.0.admin_username", "adminUsername2"),
				resource.TestCheckResourceAttrSet(resourceName, "type_details.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "type_details.0.type", "ORACLE_KEY_VAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "type_details.0.vault_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_key_stores", "test_key_stores", acctest.Optional, acctest.Update, keyStoreDataSourceRepresentation) +
				compartmentIdVariableStr + KeyStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Update, keyStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "key_stores.0.associated_databases.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "key_stores.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "key_stores.0.display_name", "Key Store1"),
				resource.TestCheckResourceAttr(datasourceName, "key_stores.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_stores.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_stores.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_stores.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "key_stores.0.type_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "key_stores.0.type_details.0.admin_username", "adminUsername2"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_stores.0.type_details.0.secret_id"),
				resource.TestCheckResourceAttr(datasourceName, "key_stores.0.type_details.0.type", "ORACLE_KEY_VAULT"),
				resource.TestCheckResourceAttrSet(datasourceName, "key_stores.0.type_details.0.vault_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Required, acctest.Create, keyStoreSingularDataSourceRepresentation) +
				compartmentIdVariableStr + KeyStoreResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_store_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "associated_databases.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Key Store1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type_details.0.admin_username", "adminUsername2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type_details.0.type", "ORACLE_KEY_VAULT"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + KeyStoreResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseKeyStoreDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_key_store" {
			noResourceFound = false
			request := oci_database.GetKeyStoreRequest{}

			tmp := rs.Primary.ID
			request.KeyStoreId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetKeyStore(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.KeyStoreLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseKeyStore") {
		resource.AddTestSweepers("DatabaseKeyStore", &resource.Sweeper{
			Name:         "DatabaseKeyStore",
			Dependencies: acctest.DependencyGraph["keyStore"],
			F:            sweepDatabaseKeyStoreResource,
		})
	}
}

func sweepDatabaseKeyStoreResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	keyStoreIds, err := getKeyStoreIds(compartment)
	if err != nil {
		return err
	}
	for _, keyStoreId := range keyStoreIds {
		if ok := acctest.SweeperDefaultResourceId[keyStoreId]; !ok {
			deleteKeyStoreRequest := oci_database.DeleteKeyStoreRequest{}

			deleteKeyStoreRequest.KeyStoreId = &keyStoreId

			deleteKeyStoreRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteKeyStore(context.Background(), deleteKeyStoreRequest)
			if error != nil {
				fmt.Printf("Error deleting KeyStore %s %s, It is possible that the resource is already deleted. Please verify manually \n", keyStoreId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &keyStoreId, keyStoreSweepWaitCondition, time.Duration(3*time.Minute),
				keyStoreSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getKeyStoreIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "KeyStoreId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listKeyStoresRequest := oci_database.ListKeyStoresRequest{}
	listKeyStoresRequest.CompartmentId = &compartmentId
	listKeyStoresResponse, err := databaseClient.ListKeyStores(context.Background(), listKeyStoresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting KeyStore list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, keyStore := range listKeyStoresResponse.Items {
		id := *keyStore.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "KeyStoreId", id)
	}
	return resourceIds, nil
}

func keyStoreSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if keyStoreResponse, ok := response.Response.(oci_database.GetKeyStoreResponse); ok {
		return keyStoreResponse.LifecycleState != oci_database.KeyStoreLifecycleStateDeleted
	}
	return false
}

func keyStoreSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetKeyStore(context.Background(), oci_database.GetKeyStoreRequest{
		KeyStoreId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
