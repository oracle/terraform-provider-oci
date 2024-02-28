// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
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
	DatabaseManagementNamedCredentialRequiredOnlyResource = DatabaseManagementNamedCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_named_credential", "test_named_credential", acctest.Required, acctest.Create, DatabaseManagementNamedCredentialRepresentation)

	DatabaseManagementNamedCredentialResourceConfig = DatabaseManagementNamedCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_named_credential", "test_named_credential", acctest.Optional, acctest.Update, DatabaseManagementNamedCredentialRepresentation)

	DatabaseManagementNamedCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"named_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_named_credential.test_named_credential.id}`},
	}

	DatabaseManagementNamedCredentialDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"associated_resource": acctest.Representation{RepType: acctest.Optional, Create: `${var.associated_resource_id}`, Update: `${var.associated_resource_updated_id}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `TestNamedCredential`},
		"scope":               acctest.Representation{RepType: acctest.Optional, Create: `RESOURCE`, Update: `GLOBAL`},
		"type":                acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_DB`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementNamedCredentialDataSourceFilterRepresentation}}

	DatabaseManagementNamedCredentialDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_named_credential.test_named_credential.id}`}},
	}

	DatabaseManagementNamedCredentialRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementNamedCredentialContentRepresentation},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `TestNamedCredential`},
		"scope":               acctest.Representation{RepType: acctest.Required, Create: `RESOURCE`, Update: `GLOBAL`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DB`},
		"associated_resource": acctest.Representation{RepType: acctest.Required, Create: `${var.associated_resource_id}`, Update: ``},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `Oracle DB named credential`, Update: `description2`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}
	DatabaseManagementNamedCredentialContentRepresentation = map[string]interface{}{
		"credential_type":             acctest.Representation{RepType: acctest.Required, Create: `BASIC`},
		"password_secret_access_mode": acctest.Representation{RepType: acctest.Required, Create: `USER_PRINCIPAL`, Update: `RESOURCE_PRINCIPAL`},
		"password_secret_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.key_id}`},
		"role":                        acctest.Representation{RepType: acctest.Required, Create: `${var.nc_user_role}`, Update: `${var.nc_user_role}`},
		"user_name":                   acctest.Representation{RepType: acctest.Required, Create: `${var.nc_user}`},
	}

	DatabaseManagementNamedCredentialResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementNamedCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementNamedCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	log.Printf("[INFO] Compartment for named credential is %v", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("dbmgmt_compartment_id_for_move", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	log.Printf("[INFO] Updated compartmentIdUVariableStr for named credential is %v", compartmentIdUVariableStr)

	key_id := utils.GetEnvSettingWithBlankDefault("dbmgmt_vault_secret_id")
	keyIdStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", key_id)
	log.Printf("[INFO] Secret for named credential is %v", key_id)

	associated_resource_id := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	associatedResourceIdStr := fmt.Sprintf("variable \"associated_resource_id\" { default = \"%s\" }\n", associated_resource_id)
	log.Printf("[INFO] Associated resource id for named credential is %v", associated_resource_id)

	nc_user := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_user")
	ncUserStr := fmt.Sprintf("variable \"nc_user\" { default = \"%s\" }\n", nc_user)
	log.Printf("[INFO] User name for named credential is %v", nc_user)

	nc_user_role := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_user_role")
	ncUserRoleStr := fmt.Sprintf("variable \"nc_user_role\" { default = \"%s\" }\n", nc_user_role)
	log.Printf("[INFO] User role for named credential is %v", nc_user_role)

	updated_associated_resource_id := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	updatedAssociatedResourceId := fmt.Sprintf("variable \"associated_resource_updated_id\" { default = \"%s\" }\n", updated_associated_resource_id)
	log.Printf("[INFO] Associated resource id for named credential is %v", associated_resource_id)

	commonVariable := keyIdStr + associatedResourceIdStr + ncUserStr + updatedAssociatedResourceId + ncUserRoleStr

	resourceName := "oci_database_management_named_credential.test_named_credential"
	datasourceName := "data.oci_database_management_named_credentials.test_named_credentials"
	singularDatasourceName := "data.oci_database_management_named_credential.test_named_credential"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementNamedCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + commonVariable + DatabaseManagementNamedCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_named_credential", "test_named_credential", acctest.Required, acctest.Create, DatabaseManagementNamedCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content.0.credential_type", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "content.0.password_secret_access_mode", "USER_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(resourceName, "content.0.password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "content.0.role", nc_user_role),
				resource.TestCheckResourceAttrSet(resourceName, "content.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestNamedCredential"),
				resource.TestCheckResourceAttr(resourceName, "scope", "RESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DB"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementNamedCredentialResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + commonVariable + DatabaseManagementNamedCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_named_credential", "test_named_credential", acctest.Optional, acctest.Create, DatabaseManagementNamedCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "associated_resource", associated_resource_id),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content.0.credential_type", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "content.0.password_secret_access_mode", "USER_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(resourceName, "content.0.password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "content.0.role", nc_user_role),
				resource.TestCheckResourceAttrSet(resourceName, "content.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "Oracle DB named credential"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestNamedCredential"),
				resource.TestCheckResourceAttr(resourceName, "scope", "RESOURCE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DB"),
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + commonVariable + DatabaseManagementNamedCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_named_credential", "test_named_credential", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseManagementNamedCredentialRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "associated_resource", associated_resource_id),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content.0.credential_type", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "content.0.password_secret_access_mode", "USER_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(resourceName, "content.0.password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "content.0.role", nc_user_role),
				resource.TestCheckResourceAttrSet(resourceName, "content.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "Oracle DB named credential"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestNamedCredential"),
				resource.TestCheckResourceAttr(resourceName, "scope", "RESOURCE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DB"),
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + commonVariable + DatabaseManagementNamedCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_named_credential", "test_named_credential", acctest.Optional, acctest.Update, DatabaseManagementNamedCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "associated_resource", updated_associated_resource_id),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content.0.credential_type", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "content.0.password_secret_access_mode", "RESOURCE_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(resourceName, "content.0.password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "content.0.role", nc_user_role),
				resource.TestCheckResourceAttrSet(resourceName, "content.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestNamedCredential"),
				resource.TestCheckResourceAttr(resourceName, "scope", "GLOBAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DB"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_named_credentials", "test_named_credentials", acctest.Optional, acctest.Update, DatabaseManagementNamedCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + commonVariable + DatabaseManagementNamedCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_named_credential", "test_named_credential", acctest.Optional, acctest.Update, DatabaseManagementNamedCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(datasourceName, "associated_resource", updated_associated_resource_id),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "TestNamedCredential"),
				resource.TestCheckResourceAttr(datasourceName, "scope", "GLOBAL"),
				resource.TestCheckResourceAttr(datasourceName, "type", "ORACLE_DB"),
				resource.TestCheckResourceAttr(datasourceName, "named_credential_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "named_credential_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_named_credential", "test_named_credential", acctest.Required, acctest.Create, DatabaseManagementNamedCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + commonVariable + DatabaseManagementNamedCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "named_credential_id"),

				//resource.TestCheckResourceAttr(singularDatasourceName, "associated_resource", updated_associated_resource_id),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "content.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content.0.credential_type", "BASIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content.0.password_secret_access_mode", "RESOURCE_PRINCIPAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content.0.role", nc_user_role),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TestNamedCredential"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope", "GLOBAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ORACLE_DB"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementNamedCredentialRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"associated_resource",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementNamedCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_named_credential" {
			noResourceFound = false
			request := oci_database_management.GetNamedCredentialRequest{}

			tmp := rs.Primary.ID
			request.NamedCredentialId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetNamedCredential(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementNamedCredential") {
		resource.AddTestSweepers("DatabaseManagementNamedCredential", &resource.Sweeper{
			Name:         "DatabaseManagementNamedCredential",
			Dependencies: acctest.DependencyGraph["namedCredential"],
			F:            sweepDatabaseManagementNamedCredentialResource,
		})
	}
}

func sweepDatabaseManagementNamedCredentialResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	namedCredentialIds, err := getDatabaseManagementNamedCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, namedCredentialId := range namedCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[namedCredentialId]; !ok {
			deleteNamedCredentialRequest := oci_database_management.DeleteNamedCredentialRequest{}

			deleteNamedCredentialRequest.NamedCredentialId = &namedCredentialId

			deleteNamedCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteNamedCredential(context.Background(), deleteNamedCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting NamedCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", namedCredentialId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &namedCredentialId, DatabaseManagementNamedCredentialSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementNamedCredentialSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementNamedCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NamedCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listNamedCredentialsRequest := oci_database_management.ListNamedCredentialsRequest{}
	listNamedCredentialsRequest.CompartmentId = &compartmentId
	listNamedCredentialsResponse, err := dbManagementClient.ListNamedCredentials(context.Background(), listNamedCredentialsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NamedCredential list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, namedCredential := range listNamedCredentialsResponse.Items {
		id := *namedCredential.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NamedCredentialId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementNamedCredentialSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if namedCredentialResponse, ok := response.Response.(oci_database_management.GetNamedCredentialResponse); ok {
		return namedCredentialResponse.LifecycleState != oci_database_management.LifecycleStatesDeleted
	}
	return false
}

func DatabaseManagementNamedCredentialSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetNamedCredential(context.Background(), oci_database_management.GetNamedCredentialRequest{
		NamedCredentialId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
