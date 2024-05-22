// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	AutonomousDatabaseInsightRequiredOnlyResource = AutonomousDatabaseInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, autonomousDatabaseInsightRepresentation)

	AutonomousDatabaseInsightResourceConfig = AutonomousDatabaseInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, autonomousDatabaseInsightRepresentation)

	autonomousDatabaseInsightSingularDataSourceRepresentation = map[string]interface{}{
		"database_insight_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_database_insight.test_database_insight.id}`},
	}

	autonomousDatabaseInsightDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"database_type":             acctest.Representation{RepType: acctest.Optional, Create: []string{`ADW-S`}},
		"fields":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`databaseName`, `databaseType`, `compartmentId`, `databaseDisplayName`, `freeformTags`, `definedTags`}},
		"id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_database_insight.test_database_insight.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`ENABLED`}, Update: []string{`DISABLED`}},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseInsightDataSourceFilterRepresentation},
	}

	autonomousDatabaseInsightDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_database_insight.test_database_insight.id}`}},
	}

	autonomousDatabaseInsightRepresentation = map[string]interface{}{
		"database_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.autonomous_database_id}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_source":                acctest.Representation{RepType: acctest.Required, Create: `AUTONOMOUS_DATABASE`, Update: `AUTONOMOUS_DATABASE`},
		"is_advanced_features_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"database_resource_type":       acctest.Representation{RepType: acctest.Required, Create: `autonomousdatabase`},
		"credential_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseInsightCredentialDetailsRepresentation},
		"connection_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseInsightConnectionDetailsRepresentation},
		"status":                       acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesAutonomousDatabaseInsightRepresentation},
	}

	autonomousDatabaseInsightCredentialDetailsRepresentation = map[string]interface{}{
		"credential_type": acctest.Representation{RepType: acctest.Required, Create: `CREDENTIALS_BY_IAM`},
	}

	autonomousDatabaseInsightCredentialDetailsForUpdateRepresentation = map[string]interface{}{
		"credential_type":    acctest.Representation{RepType: acctest.Required, Update: `CREDENTIALS_BY_VAULT`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Update: `${var.secret_id}`},
		"role":               acctest.Representation{RepType: acctest.Optional, Update: `NORMAL`},
		"user_name":          acctest.Representation{RepType: acctest.Required, Update: `${var.user_name}`},
	}

	autonomousDatabaseInsightConnectionDetailsRepresentation = map[string]interface{}{
		"host_name":    acctest.Representation{RepType: acctest.Required, Create: `${var.adb_host}`},
		"port":         acctest.Representation{RepType: acctest.Required, Create: `${var.adb_port}`},
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `TCPS`},
		"service_name": acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
	}

	ignoreChangesAutonomousDatabaseInsightRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	AutonomousDatabaseInsightResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAutonomousDatabaseInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAutonomousDatabaseInsightResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	autonomousDatabaseId := utils.GetEnvSettingWithBlankDefault("autonomous_database_id")
	autonomousDatabaseIdVariableStr := fmt.Sprintf("variable \"autonomous_database_id\" { default = \"%s\" }\n", autonomousDatabaseId)

	adbHostName := utils.GetEnvSettingWithBlankDefault("adb_host")
	adbHostNameVariableStr := fmt.Sprintf("variable \"adb_host\" { default = \"%s\" }\n", adbHostName)

	adbPort := utils.GetEnvSettingWithBlankDefault("adb_port")
	adbPortVariableStr := fmt.Sprintf("variable \"adb_port\" { default = \"%s\" }\n", adbPort)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	secretId := utils.GetEnvSettingWithBlankDefault("secret_id")
	secretIdVariableStr := fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", secretId)

	userName := utils.GetEnvSettingWithBlankDefault("user_name")
	userNamedVariableStr := fmt.Sprintf("variable \"user_name\" { default = \"%s\" }\n", userName)

	resourceName := "oci_opsi_database_insight.test_database_insight"
	datasourceName := "data.oci_opsi_database_insights.test_database_insights"
	singularDatasourceName := "data.oci_opsi_database_insight.test_database_insight"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+autonomousDatabaseIdVariableStr+adbHostNameVariableStr+adbPortVariableStr+serviceNameVariableStr+AutonomousDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Create, autonomousDatabaseInsightRepresentation), "opsi", "databaseInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiAutonomousDatabaseInsightDestroy, []resource.TestStep{
		// verify create with optional opsiPrivateEndpointId
		{
			Config: config + compartmentIdVariableStr + autonomousDatabaseIdVariableStr + adbHostNameVariableStr + adbPortVariableStr + serviceNameVariableStr + AutonomousDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Create, autonomousDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.credential_type", "CREDENTIALS_BY_IAM"),
				resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.host_name", adbHostName),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.port", adbPort),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.service_name", serviceName),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "database_resource_type", "autonomousdatabase"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "is_advanced_features_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
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
		// verify update to the credential by vault
		{
			Config: config + compartmentIdVariableStr + autonomousDatabaseIdVariableStr + adbHostNameVariableStr + adbPortVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + AutonomousDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseInsightRepresentation, map[string]interface{}{
						"credential_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseInsightCredentialDetailsForUpdateRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.credential_type", "CREDENTIALS_BY_VAULT"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.password_secret_id", secretId),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.role", "NORMAL"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.host_name", adbHostName),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.port", adbPort),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.service_name", serviceName),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "database_resource_type", "autonomousdatabase"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
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
			Config: config + compartmentIdVariableStr + autonomousDatabaseIdVariableStr + adbHostNameVariableStr + adbPortVariableStr + serviceNameVariableStr + AutonomousDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, autonomousDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.credential_type", "CREDENTIALS_BY_IAM"),
				resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.host_name", adbHostName),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.port", adbPort),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.service_name", serviceName),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "database_resource_type", "autonomousdatabase"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_database_insights", "test_database_insights", acctest.Optional, acctest.Update, autonomousDatabaseInsightDataSourceRepresentation) +
				compartmentIdVariableStr + autonomousDatabaseIdVariableStr + adbHostNameVariableStr + adbPortVariableStr + serviceNameVariableStr + AutonomousDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, autonomousDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "database_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fields.#", "6"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "database_insights_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_insights_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, autonomousDatabaseInsightSingularDataSourceRepresentation) +
				compartmentIdVariableStr + autonomousDatabaseIdVariableStr + adbHostNameVariableStr + adbPortVariableStr + serviceNameVariableStr + AutonomousDatabaseInsightResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.credential_type", "CREDENTIALS_BY_IAM"),
				resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.host_name", adbHostName),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.port", adbPort),
				resource.TestCheckResourceAttr(resourceName, "connection_details.0.service_name", serviceName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_resource_type", "autonomousdatabase"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_source", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + autonomousDatabaseIdVariableStr + adbHostNameVariableStr + adbPortVariableStr + serviceNameVariableStr + AutonomousDatabaseInsightResourceConfig,
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + autonomousDatabaseIdVariableStr + adbHostNameVariableStr + adbPortVariableStr + serviceNameVariableStr + AutonomousDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseInsightRepresentation, map[string]interface{}{
						"status": acctest.Representation{RepType: acctest.Required, Update: `ENABLED`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:            config + AutonomousDatabaseInsightRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"entity_source",
				"is_advanced_features_enable",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOpsiAutonomousDatabaseInsightDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_database_insight" {
			noResourceFound = false
			request := oci_opsi.GetDatabaseInsightRequest{}

			tmp := rs.Primary.ID
			request.DatabaseInsightId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetDatabaseInsight(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpsiAutonomousDatabaseInsight") {
		resource.AddTestSweepers("OpsiAutonomousDatabaseInsight", &resource.Sweeper{
			Name:         "OpsiAutonomousDatabaseInsight",
			Dependencies: acctest.DependencyGraph["databaseInsight"],
			F:            sweepOpsiAutonomousDatabaseInsightResource,
		})
	}
}

func sweepOpsiAutonomousDatabaseInsightResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	databaseInsightIds, err := getAutonomousDatabaseInsightIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseInsightId := range databaseInsightIds {
		if ok := acctest.SweeperDefaultResourceId[databaseInsightId]; !ok {
			deleteDatabaseInsightRequest := oci_opsi.DeleteDatabaseInsightRequest{}

			deleteDatabaseInsightRequest.DatabaseInsightId = &databaseInsightId

			deleteDatabaseInsightRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteDatabaseInsight(context.Background(), deleteDatabaseInsightRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseInsight %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseInsightId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseInsightId, autonomousDatabaseInsightSweepWaitCondition, time.Duration(3*time.Minute),
				autonomousDatabaseInsightSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getAutonomousDatabaseInsightIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseInsightId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listDatabaseInsightsRequest := oci_opsi.ListDatabaseInsightsRequest{}
	listDatabaseInsightsRequest.CompartmentId = &compartmentId
	listDatabaseInsightsRequest.LifecycleState = []oci_opsi.LifecycleStateEnum{oci_opsi.LifecycleStateActive}
	listDatabaseInsightsResponse, err := operationsInsightsClient.ListDatabaseInsights(context.Background(), listDatabaseInsightsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseInsight list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseInsight := range listDatabaseInsightsResponse.Items {
		id := *databaseInsight.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseInsightId", id)
	}
	return resourceIds, nil
}

func autonomousDatabaseInsightSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseInsightResponse, ok := response.Response.(oci_opsi.GetDatabaseInsightResponse); ok {
		return databaseInsightResponse.GetLifecycleState() != oci_opsi.LifecycleStateDeleted
	}
	return false
}

func autonomousDatabaseInsightSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetDatabaseInsight(context.Background(), oci_opsi.GetDatabaseInsightRequest{
		DatabaseInsightId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
