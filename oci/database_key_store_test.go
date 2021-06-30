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
	"github.com/oracle/oci-go-sdk/v43/common"
	oci_database "github.com/oracle/oci-go-sdk/v43/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	KeyStoreRequiredOnlyResource = KeyStoreResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Required, Create, keyStoreRepresentation)

	KeyStoreResourceConfig = KeyStoreResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Optional, Update, keyStoreRepresentation)

	keyStoreSingularDataSourceRepresentation = map[string]interface{}{
		"key_store_id": Representation{repType: Required, create: `${oci_database_key_store.test_key_store.id}`},
	}

	keyStoreDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, keyStoreDataSourceFilterRepresentation}}
	keyStoreDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_key_store.test_key_store.id}`}},
	}

	keyStoreRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `Key Store1`},
		"type_details":   RepresentationGroup{Required, keyStoreTypeDetailsRepresentation},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	keyStoreTypeDetailsRepresentation = map[string]interface{}{
		"admin_username": Representation{repType: Required, create: `username1`, update: `adminUsername2`},
		"connection_ips": Representation{repType: Required, create: []string{`192.1.1.1`}, update: []string{`192.1.1.1`, `192.1.1.2`}},
		"secret_id":      Representation{repType: Required, create: `${var.okv_secret}`},
		"type":           Representation{repType: Required, create: `ORACLE_KEY_VAULT`},
		"vault_id":       Representation{repType: Required, create: `${var.kms_vault_id}`},
	}

	okvSecret            = getEnvSettingWithBlankDefault("okv_secret")
	OkvSecretVariableStr = fmt.Sprintf("variable \"okv_secret\" { default = \"%s\" }\n", okvSecret)

	KeyStoreResourceDependencies = DefinedTagsDependencies + KmsVaultIdVariableStr + OkvSecretVariableStr
)

func TestDatabaseKeyStoreResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseKeyStoreResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_key_store.test_key_store"
	datasourceName := "data.oci_database_key_stores.test_key_stores"
	singularDatasourceName := "data.oci_database_key_store.test_key_store"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+KeyStoreResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Optional, Create, keyStoreRepresentation), "database", "keyStore", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseKeyStoreDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + KeyStoreResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Required, Create, keyStoreRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Key Store1"),
					resource.TestCheckResourceAttr(resourceName, "type_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "type_details.0.admin_username", "username1"),
					resource.TestCheckResourceAttrSet(resourceName, "type_details.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "type_details.0.type", "ORACLE_KEY_VAULT"),
					resource.TestCheckResourceAttrSet(resourceName, "type_details.0.vault_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + KeyStoreResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + KeyStoreResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Optional, Create, keyStoreRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + KeyStoreResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Optional, Create,
						representationCopyWithNewProperties(keyStoreRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + KeyStoreResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Optional, Update, keyStoreRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_database_key_stores", "test_key_stores", Optional, Update, keyStoreDataSourceRepresentation) +
					compartmentIdVariableStr + KeyStoreResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Optional, Update, keyStoreRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "key_stores.0.associated_databases.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "key_stores.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "key_stores.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_database_key_store", "test_key_store", Required, Create, keyStoreSingularDataSourceRepresentation) +
					compartmentIdVariableStr + KeyStoreResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_store_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "associated_databases.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		},
	})
}

func testAccCheckDatabaseKeyStoreDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_key_store" {
			noResourceFound = false
			request := oci_database.GetKeyStoreRequest{}

			tmp := rs.Primary.ID
			request.KeyStoreId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseKeyStore") {
		resource.AddTestSweepers("DatabaseKeyStore", &resource.Sweeper{
			Name:         "DatabaseKeyStore",
			Dependencies: DependencyGraph["keyStore"],
			F:            sweepDatabaseKeyStoreResource,
		})
	}
}

func sweepDatabaseKeyStoreResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	keyStoreIds, err := getKeyStoreIds(compartment)
	if err != nil {
		return err
	}
	for _, keyStoreId := range keyStoreIds {
		if ok := SweeperDefaultResourceId[keyStoreId]; !ok {
			deleteKeyStoreRequest := oci_database.DeleteKeyStoreRequest{}

			deleteKeyStoreRequest.KeyStoreId = &keyStoreId

			deleteKeyStoreRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteKeyStore(context.Background(), deleteKeyStoreRequest)
			if error != nil {
				fmt.Printf("Error deleting KeyStore %s %s, It is possible that the resource is already deleted. Please verify manually \n", keyStoreId, error)
				continue
			}
			waitTillCondition(testAccProvider, &keyStoreId, keyStoreSweepWaitCondition, time.Duration(3*time.Minute),
				keyStoreSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getKeyStoreIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "KeyStoreId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listKeyStoresRequest := oci_database.ListKeyStoresRequest{}
	listKeyStoresRequest.CompartmentId = &compartmentId
	listKeyStoresResponse, err := databaseClient.ListKeyStores(context.Background(), listKeyStoresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting KeyStore list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, keyStore := range listKeyStoresResponse.Items {
		id := *keyStore.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "KeyStoreId", id)
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

func keyStoreSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetKeyStore(context.Background(), oci_database.GetKeyStoreRequest{
		KeyStoreId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
