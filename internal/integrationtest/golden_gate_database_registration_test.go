// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

func makeVariableStr(envVar string, t *testing.T) string {
	var envVarVal = utils.GetEnvSettingWithBlankDefault(envVar)
	if envVarVal == "" {
		t.Skip(fmt.Sprintf("Missing environment %s: define \"TF_VAR_%s\"\n", envVar, envVar))
		return ""
	}
	return fmt.Sprintf("variable \"%s\" { default = \"%s\" }\n", envVar, envVarVal)
}

// issue-routing-tag: golden_gate/default
func TestGoldenGateDatabaseRegistrationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDatabaseRegistrationResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID          = "compartment_id"
		COMPARTMENT_ID_FOR_MOVE = "compartment_id_for_move"
		TEST_SUBNET_ID          = "test_subnet_id"
		TEST_DB_ID              = "test_db_id"
		KMS_KEY_ID              = "kms_key_id"
		KMS_VAULT_ID            = "kms_vault_id"
		PASSWORD                = "password"
		NEW_PASSWORD            = "new_password"
	)

	var (
		resourceName           = "oci_golden_gate_database_registration.test_database_registration"
		datasourceName         = "data.oci_golden_gate_database_registrations.test_database_registrations"
		singularDatasourceName = "data.oci_golden_gate_database_registration.test_database_registration"
	)

	var (
		compartmentId        = utils.GetEnvSettingWithBlankDefault(COMPARTMENT_ID)
		compartmentIdForMove = utils.GetEnvSettingWithBlankDefault(COMPARTMENT_ID_FOR_MOVE)
		resId                string
		resId2               string
	)

	var (
		DatabaseRegistrationResourceDependencies = DefinedTagsDependencies

		ignoreDefinedTagsChangesRepresentation = map[string]interface{}{
			"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
		}

		databaseRegistrationDataSourceFilterRepresentation = map[string]interface{}{
			"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
			"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_golden_gate_database_registration.test_database_registration.id}`}},
		}

		// NOTE: The connection string needs to use the FQDN for the hostname for passing API validation.
		databaseRegistrationRepresentation = map[string]interface{}{
			"alias_name":            acctest.Representation{RepType: acctest.Required, Create: `aliasName1`, Update: `aliasName2`},
			"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
			"fqdn":                  acctest.Representation{RepType: acctest.Required, Create: `fqdn.example.com`, Update: `fqdn2.example.com`},
			"password":              acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
			"username":              acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
			"connection_string":     acctest.Representation{RepType: acctest.Optional, Create: `fqdn.example.com:1521/ORION_phx1gq.example.com`, Update: `fqdn2.example.com:1521/ORION_phx1gq.example.com`},
			"database_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.test_db_id}`},
			"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
			"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
			"ip_address":            acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.10`},
			"key_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
			"secret_compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
			"session_mode":          acctest.Representation{RepType: acctest.Optional, Create: `DIRECT`, Update: `REDIRECT`},
			"subnet_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.test_subnet_id}`},
			"vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_vault_id}`},
			"wallet":                acctest.Representation{RepType: acctest.Optional, Create: `wallet`, Update: `wallet2`},
			"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
		}

		GoldenGateDatabaseRegistrationRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_database_registration", "test_database_registration", acctest.Required, acctest.Create, databaseRegistrationRepresentation)

		DatabaseRegistrationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_database_registration", "test_database_registration", acctest.Optional, acctest.Update, databaseRegistrationRepresentation)

		databaseRegistrationSingularDataSourceRepresentation = map[string]interface{}{
			"database_registration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_database_registration.test_database_registration.id}`},
		}

		databaseRegistrationDataSourceRepresentation = map[string]interface{}{
			"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
			"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
			"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseRegistrationDataSourceFilterRepresentation},
		}
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(COMPARTMENT_ID_FOR_MOVE, t) +
		makeVariableStr(TEST_SUBNET_ID, t) +
		makeVariableStr(TEST_DB_ID, t) +
		makeVariableStr(KMS_KEY_ID, t) +
		makeVariableStr(KMS_VAULT_ID, t) +
		makeVariableStr(PASSWORD, t) +
		makeVariableStr(NEW_PASSWORD, t) +
		DatabaseRegistrationResourceDependencies

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_database_registration", "test_database_registration", acctest.Optional, acctest.Create, databaseRegistrationRepresentation), "goldengate", "databaseRegistration", t)

	acctest.ResourceTest(t, testAccCheckGoldenGateDatabaseRegistrationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_database_registration", "test_database_registration", acctest.Required, acctest.Create, databaseRegistrationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alias_name", "aliasName1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fqdn", "fqdn.example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "password"),
				resource.TestCheckResourceAttr(resourceName, "username", "username"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config,
		},

		// verify Create with optionals
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_database_registration", "test_database_registration", acctest.Optional, acctest.Create, databaseRegistrationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alias_name", "aliasName1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "fqdn.example.com:1521/ORION_phx1gq.example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fqdn", "fqdn.example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "password"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "session_mode", "DIRECT"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "username", "username"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "wallet", "wallet"),

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
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_database_registration", "test_database_registration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(databaseRegistrationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_move}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alias_name", "aliasName1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdForMove),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "fqdn.example.com:1521/ORION_phx1gq.example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fqdn", "fqdn.example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "session_mode", "DIRECT"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "username", "username"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "wallet", "wallet"),

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
			Config: config + DatabaseRegistrationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alias_name", "aliasName2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "fqdn2.example.com:1521/ORION_phx1gq.example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "fqdn", "fqdn2.example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "secret_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "session_mode", "REDIRECT"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "username", "username2"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "wallet", "wallet2"),

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
			Config: config + DatabaseRegistrationResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_database_registrations", "test_database_registrations", acctest.Optional, acctest.Update, databaseRegistrationDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "database_registration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_registration_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config + DatabaseRegistrationResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_database_registration", "test_database_registration", acctest.Required, acctest.Create, databaseRegistrationSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_registration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "alias_name", "aliasName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string", "fqdn2.example.com:1521/ORION_phx1gq.example.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fqdn", "fqdn2.example.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "session_mode", "REDIRECT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "username", "username2"),
			),
		},
		// verify resource import
		{
			Config:            config + GoldenGateDatabaseRegistrationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"password",
				"wallet",
				"key_id",
				"secret_compartment_id",
				"vault_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckGoldenGateDatabaseRegistrationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_database_registration" {
			noResourceFound = false
			request := oci_golden_gate.GetDatabaseRegistrationRequest{}

			tmp := rs.Primary.ID
			request.DatabaseRegistrationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			response, err := client.GetDatabaseRegistration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_golden_gate.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GoldenGateDatabaseRegistration") {
		resource.AddTestSweepers("GoldenGateDatabaseRegistration", &resource.Sweeper{
			Name:         "GoldenGateDatabaseRegistration",
			Dependencies: acctest.DependencyGraph["databaseRegistration"],
			F:            sweepGoldenGateDatabaseRegistrationResource,
		})
	}
}

func sweepGoldenGateDatabaseRegistrationResource(compartment string) error {
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()
	databaseRegistrationIds, err := getDatabaseRegistrationIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseRegistrationId := range databaseRegistrationIds {
		if ok := acctest.SweeperDefaultResourceId[databaseRegistrationId]; !ok {
			deleteDatabaseRegistrationRequest := oci_golden_gate.DeleteDatabaseRegistrationRequest{}

			deleteDatabaseRegistrationRequest.DatabaseRegistrationId = &databaseRegistrationId

			deleteDatabaseRegistrationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeleteDatabaseRegistration(context.Background(), deleteDatabaseRegistrationRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseRegistration %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseRegistrationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseRegistrationId, databaseRegistrationSweepWaitCondition, time.Duration(3*time.Minute),
				databaseRegistrationSweepResponseFetchOperation, "golden_gate", true)
		}
	}
	return nil
}

func getDatabaseRegistrationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseRegistrationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()

	listDatabaseRegistrationsRequest := oci_golden_gate.ListDatabaseRegistrationsRequest{}
	listDatabaseRegistrationsRequest.CompartmentId = &compartmentId
	listDatabaseRegistrationsRequest.LifecycleState = oci_golden_gate.ListDatabaseRegistrationsLifecycleStateActive
	listDatabaseRegistrationsResponse, err := goldenGateClient.ListDatabaseRegistrations(context.Background(), listDatabaseRegistrationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseRegistration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseRegistration := range listDatabaseRegistrationsResponse.Items {
		id := *databaseRegistration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseRegistrationId", id)
	}
	return resourceIds, nil
}

func databaseRegistrationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseRegistrationResponse, ok := response.Response.(oci_golden_gate.GetDatabaseRegistrationResponse); ok {
		return databaseRegistrationResponse.LifecycleState != oci_golden_gate.LifecycleStateDeleted
	}
	return false
}

func databaseRegistrationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GoldenGateClient().GetDatabaseRegistration(context.Background(), oci_golden_gate.GetDatabaseRegistrationRequest{
		DatabaseRegistrationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
