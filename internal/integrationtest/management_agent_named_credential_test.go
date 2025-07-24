// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	oci_vault "github.com/oracle/oci-go-sdk/v65/vault"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ManagementAgentNamedCredentialRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_management_agent_named_credential", "test_named_credential", acctest.Required, acctest.Create, ManagementAgentNamedCredentialRepresentation)

	ManagementAgentNamedCredentialResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_management_agent_named_credential", "test_named_credential", acctest.Optional, acctest.Update, ManagementAgentNamedCredentialRepresentation)

	ManagementAgentNamedCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"named_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_management_agent_named_credential.test_named_credential.id}`},
	}

	ManagementAgentNamedCredentialDataSourceRepresentation = map[string]interface{}{
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id_nc}`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_management_agent_named_credential.test_named_credential.id}`}},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: []string{`name`}},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"type":                acctest.Representation{RepType: acctest.Optional, Create: []string{`DBCREDS`}},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ManagementAgentNamedCredentialDataSourceFilterRepresentation}}
	ManagementAgentNamedCredentialDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_management_agent_named_credential.test_named_credential.id}`}},
	}

	ManagementAgentNamedCredentialRepresentation = map[string]interface{}{
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id_nc}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`},
		"properties":          []acctest.RepresentationGroup{{RepType: acctest.Required, Group: ManagementAgentNamedCredentialPropertiesRepresentation}, {RepType: acctest.Required, Group: ManagementAgentNamedCredentialPropertiesRepresentation2}},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `DBCREDS`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	ManagementAgentNamedCredentialPropertiesRepresentation = map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: `DBUserName`, Update: `DBUserName`},
		"value":          acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`, Update: `${var.secret_id}`},
		"value_category": acctest.Representation{RepType: acctest.Required, Create: `SECRET_IDENTIFIER`, Update: `SECRET_IDENTIFIER`},
	}
	ManagementAgentNamedCredentialPropertiesRepresentation2 = map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: `DBPassword`, Update: `DBPassword`},
		"value":          acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`, Update: `${var.secret_id}`},
		"value_category": acctest.Representation{RepType: acctest.Required, Create: `SECRET_IDENTIFIER`, Update: `SECRET_IDENTIFIER`},
	}

	//	ManagementAgentNamedCredentialResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Required, acctest.Create, ManagementAgentManagementAgentRepresentation)
)

// issue-routing-tag: management_agent/default
func TestManagementAgentNamedCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentNamedCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	log.Printf("[DEBUG] NamedCredTest compartmentIdVariableStr: %s", compartmentIdVariableStr)

	//  1. List all agents in compartment (env:TF_VAR_compartment_ocid) with status=ACTIVE and displayName=terraformTest
	managementAgentIds, err := getManagementAgentIdsWithNoNamedCredentials(compartmentId)
	log.Printf("[DEBUG] NamedCredTest after managementAgentIds: %s", managementAgentIds)

	if err != nil {
		t.Errorf("Failed to get agents in compartment %s", err)
	}
	if len(managementAgentIds) == 0 {
		t.Errorf("Failed to find any active agents in compartment %s", compartmentId)
	}
	managementAgentId := managementAgentIds[0]
	log.Printf("[DEBUG] NamedCredTest Management Agent ID for NamedCred: %s", managementAgentId)

	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id_nc\" { default = \"%s\" }\n", managementAgentId)

	secretIds, err := getVaultSecrets()
	if err != nil {
		t.Errorf("Failed to get secrets in compartment %s", err)
	}
	if len(secretIds) == 0 {
		t.Errorf("Failed to find any active vaults in compartment %s", compartmentId)
	}
	secretIdVariableStr := fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", secretIds[0])
	resourceName := "oci_management_agent_named_credential.test_named_credential"
	datasourceName := "data.oci_management_agent_named_credentials.test_named_credentials"
	singularDatasourceName := "data.oci_management_agent_named_credential.test_named_credential"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+secretIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_named_credential", "test_named_credential", acctest.Optional, acctest.Create, ManagementAgentNamedCredentialRepresentation)+managementAgentIdVariableStr, "managementagent", "namedCredential", t)

	acctest.ResourceTest(t, testAccCheckManagementAgentNamedCredentialDestroy, []resource.TestStep{
		// verify Create   Test step 0
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + secretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_named_credential", "test_named_credential", acctest.Required, acctest.Create, ManagementAgentNamedCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.name", "DBUserName"),
				resource.TestCheckResourceAttrSet(resourceName, "properties.0.value"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.value_category", "SECRET_IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "type", "DBCREDS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create    Test Step 1
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + secretIdVariableStr,
		},
		// verify Create with optionals   Test Step 2
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + secretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_named_credential", "test_named_credential", acctest.Optional, acctest.Create, ManagementAgentNamedCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.name", "DBUserName"),
				resource.TestCheckResourceAttrSet(resourceName, "properties.0.value"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.value_category", "SECRET_IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "type", "DBCREDS"),

				func(s *terraform.State) (err error) {
					fmt.Printf("EXPORT here %s %s", resourceName, s)
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Printf("EXPORT here2 %s", resId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters   Test Step 3
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + secretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_named_credential", "test_named_credential", acctest.Optional, acctest.Update, ManagementAgentNamedCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.name", "DBUserName"),
				resource.TestCheckResourceAttrSet(resourceName, "properties.0.value"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.value_category", "SECRET_IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "type", "DBCREDS"),

				func(s *terraform.State) (err error) {
					fmt.Printf("EXPORT here3 %s %s", resourceName, s)
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource    Test step 4
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_named_credentials", "test_named_credentials", acctest.Optional, acctest.Update, ManagementAgentNamedCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + secretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_named_credential", "test_named_credential", acctest.Optional, acctest.Update, ManagementAgentNamedCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "id.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(datasourceName, "name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "named_credential_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "named_credential_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource   Test step 5
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_named_credential", "test_named_credential", acctest.Required, acctest.Create, ManagementAgentNamedCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + secretIdVariableStr + ManagementAgentNamedCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "named_credential_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.0.name", "DBUserName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "properties.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.0.value_category", "SECRET_IDENTIFIER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "DBCREDS"),
			),
		},
		// verify resource import   Test step 6
		{
			Config:                  config + managementAgentIdVariableStr + ManagementAgentNamedCredentialRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckManagementAgentNamedCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_management_agent_named_credential" {
			noResourceFound = false
			request := oci_management_agent.GetNamedCredentialRequest{}

			tmp := rs.Primary.ID
			request.NamedCredentialId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")

			response, err := client.GetNamedCredential(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_management_agent.NamedCredentialLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ManagementAgentNamedCredential") {
		resource.AddTestSweepers("ManagementAgentNamedCredential", &resource.Sweeper{
			Name:         "ManagementAgentNamedCredential",
			Dependencies: acctest.DependencyGraph["namedCredential"],
			F:            sweepManagementAgentNamedCredentialResource,
		})
	}
}

func sweepManagementAgentNamedCredentialResource(compartment string) error {
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()
	namedCredentialIds, err := getManagementAgentNamedCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, namedCredentialId := range namedCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[namedCredentialId]; !ok {
			deleteNamedCredentialRequest := oci_management_agent.DeleteNamedCredentialRequest{}

			deleteNamedCredentialRequest.NamedCredentialId = &namedCredentialId

			deleteNamedCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")
			_, error := managementAgentClient.DeleteNamedCredential(context.Background(), deleteNamedCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting NamedCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", namedCredentialId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &namedCredentialId, ManagementAgentNamedCredentialSweepWaitCondition, time.Duration(3*time.Minute),
				ManagementAgentNamedCredentialSweepResponseFetchOperation, "management_agent", true)
		}
	}
	return nil
}

func getManagementAgentIdsWithNoNamedCredentials(compartment string) ([]string, error) {

	var resourceIds []string
	compartmentId := compartment
	terraformTest := "terraformTestNC"
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()

	listManagementAgentsRequest := oci_management_agent.ListManagementAgentsRequest{}
	listManagementAgentsRequest.CompartmentId = &compartmentId
	listManagementAgentsRequest.LifecycleState = oci_management_agent.ListManagementAgentsLifecycleStateActive
	listManagementAgentsRequest.DisplayName = &terraformTest
	listManagementAgentsRequest.SortOrder = oci_management_agent.ListManagementAgentsSortOrderAsc
	listManagementAgentsResponse, err := managementAgentClient.ListManagementAgents(context.Background(), listManagementAgentsRequest)
	log.Printf("[DEBUG] NamedCredTest List response: %s", listManagementAgentsResponse)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementAgent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementAgent := range listManagementAgentsResponse.Items {
		log.Printf("[DEBUG] NamedCredTest Management Agent ID find for NamedCred: %s", managementAgent)

		id := *managementAgent.Id
		listNCRequest := oci_management_agent.ListNamedCredentialsRequest{
			ManagementAgentId: managementAgent.Id,
		}
		listNCResponse, err2 := managementAgentClient.ListNamedCredentials(context.Background(), listNCRequest)
		if err2 != nil {
			return resourceIds, fmt.Errorf("Error getting ManagementAgent named credential list for agent id : %s , %s \n", id, err)
		}
		if len(listNCResponse.Items) == 0 { // Only add agents which have 0 named credentials so far, this will stop duplicate names
			resourceIds = append(resourceIds, id)
			log.Printf("[DEBUG] NamedCredTest Management Agent ID find for NamedCred adding: %s", id)

			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementAgentId", id)
		} else {
			fmt.Printf("Skipping agent %s as it has named creds", id)
		}
		if len(resourceIds) >= 3 {
			break
		}
	}

	log.Printf("[DEBUG] NamedCredTest List return: %s", resourceIds)

	return resourceIds, nil
}
func getManagementAgentNamedCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NamedCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()

	listNamedCredentialsRequest := oci_management_agent.ListNamedCredentialsRequest{}

	managementAgentIds, error := getManagementAgentManagementAgentIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting managementAgentId required for NamedCredential resource requests \n")
	}
	for _, managementAgentId := range managementAgentIds {
		listNamedCredentialsRequest.ManagementAgentId = &managementAgentId

		listNamedCredentialsRequest.LifecycleState = []oci_management_agent.NamedCredentialLifecycleStateEnum{oci_management_agent.NamedCredentialLifecycleStateActive}
		listNamedCredentialsResponse, err := managementAgentClient.ListNamedCredentials(context.Background(), listNamedCredentialsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NamedCredential list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, namedCredential := range listNamedCredentialsResponse.Items {
			id := *namedCredential.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NamedCredentialId", id)
		}

	}
	return resourceIds, nil
}

func ManagementAgentNamedCredentialSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if namedCredentialResponse, ok := response.Response.(oci_management_agent.GetNamedCredentialResponse); ok {
		return namedCredentialResponse.LifecycleState != oci_management_agent.NamedCredentialLifecycleStateDeleted
	}
	return false
}

func ManagementAgentNamedCredentialSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementAgentClient().GetNamedCredential(context.Background(), oci_management_agent.GetNamedCredentialRequest{
		NamedCredentialId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getVaultSecrets() ([]string, error) {
	// Find secret in a vault, the vault and secrets will have already been created by MACS S1 test cases
	compartment := utils.GetEnvSettingWithBlankDefault("secret_compartment_ocid")

	ids := acctest.GetResourceIdsToSweep(compartment, "SecretId")
	if ids != nil {
		return ids, nil
	}

	var resourceIds []string
	compartmentId := compartment

	vaultsClient := acctest.GetTestClients(&schema.ResourceData{}).VaultsClient()
	vault := utils.GetEnvSettingWithBlankDefault("vault_ocid")
	fmt.Printf("Using vault %s \n", vault)
	listSecretsRequest := oci_vault.ListSecretsRequest{}
	listSecretsRequest.CompartmentId = &compartmentId
	listSecretsRequest.LifecycleState = oci_vault.SecretSummaryLifecycleStateActive
	listSecretsRequest.VaultId = &vault
	secretsResponse, err := vaultsClient.ListSecrets(context.Background(), listSecretsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Secret list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, item := range secretsResponse.Items {
		id := *item.Id
		resourceIds = append(resourceIds, id)
		fmt.Printf("Found SecretId %s \n", id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecretId", id)
	}
	return resourceIds, nil
}
