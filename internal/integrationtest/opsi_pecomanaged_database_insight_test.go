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
	PecomanagedDatabaseInsightRequiredOnlyResource = OpsiDatabaseInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, pecomanagedDatabaseInsightRepresentation)

	PecomanagedDatabaseInsightResourceConfig = OpsiDatabaseInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, pecomanagedDatabaseInsightRepresentation)

	pecomanagedDatabaseInsightSingularDataSourceRepresentation = map[string]interface{}{
		"database_insight_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_database_insight.test_database_insight.id}`},
	}

	pecomanagedDatabaseInsightDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"database_type":             acctest.Representation{RepType: acctest.Optional, Create: []string{`COMANAGED-VM-CDB`}},
		"fields":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`databaseName`, `databaseType`, `compartmentId`, `databaseDisplayName`, `freeformTags`, `definedTags`}},
		"id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_database_insight.test_database_insight.id}`},
		"opsi_private_endpoint_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.opsi_private_endpoint_id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`ENABLED`}, Update: []string{`DISABLED`}},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: pecomanagedDatabaseInsightDataSourceFilterRepresentation},
	}

	pecomanagedDatabaseInsightDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_database_insight.test_database_insight.id}`}},
	}

	pecomanagedDatabaseInsightRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_source":            acctest.Representation{RepType: acctest.Required, Create: `PE_COMANAGED_DATABASE`, Update: `PE_COMANAGED_DATABASE`},
		"credential_details":       acctest.RepresentationGroup{RepType: acctest.Required, Group: pecomanagedDatabaseInsightCredentialDetailsRepresentation},
		"database_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.dbsystem_database_id}`},
		"database_resource_type":   acctest.Representation{RepType: acctest.Required, Create: `database`},
		"deployment_type":          acctest.Representation{RepType: acctest.Required, Create: `VIRTUAL_MACHINE`},
		"opsi_private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opsi_private_endpoint_id}`},
		"service_name":             acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
		"status":                   acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesPecomanagedDatabaseInsightRepresentation},
	}

	pecomanagedDatabaseInsightCredentialDetailsRepresentation = map[string]interface{}{
		"credential_type":    acctest.Representation{RepType: acctest.Required, Create: `CREDENTIALS_BY_VAULT`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
		"role":               acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `${var.user_name}`},
	}

	pecomanagedDatabaseInsightCredentialDetailsForUpdateRepresentation = map[string]interface{}{
		"credential_type":    acctest.Representation{RepType: acctest.Required, Create: `CREDENTIALS_BY_VAULT`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id_for_update}`},
		"role":               acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `${var.user_name}`},
	}

	ignoreChangesPecomanagedDatabaseInsightRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	PecomanagedDatabaseInsightResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiPecomanagedDatabaseInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiPecomanagedDatabaseInsightResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	opsiPrivateEndpointId := utils.GetEnvSettingWithBlankDefault("opsi_private_endpoint_id")
	opsiPrivateEndpointIdVariableStr := fmt.Sprintf("variable \"opsi_private_endpoint_id\" { default = \"%s\" }\n", opsiPrivateEndpointId)

	dbsystemDatabaseId := utils.GetEnvSettingWithBlankDefault("dbsystem_database_id")
	dbsystemDatabaseIdVariableStr := fmt.Sprintf("variable \"dbsystem_database_id\" { default = \"%s\" }\n", dbsystemDatabaseId)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	secretId := utils.GetEnvSettingWithBlankDefault("secret_id")
	secretIdVariableStr := fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", secretId)

	secretIdU := utils.GetEnvSettingWithDefault("secret_id_for_update", secretId)
	secretIdUVariableStr := fmt.Sprintf("variable \"secret_id_for_update\" { default = \"%s\" }\n", secretIdU)

	userName := utils.GetEnvSettingWithBlankDefault("user_name")
	userNamedVariableStr := fmt.Sprintf("variable \"user_name\" { default = \"%s\" }\n", userName)

	resourceName := "oci_opsi_database_insight.test_database_insight"
	datasourceName := "data.oci_opsi_database_insights.test_database_insights"
	singularDatasourceName := "data.oci_opsi_database_insight.test_database_insight"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbsystemDatabaseIdVariableStr+opsiPrivateEndpointIdVariableStr+serviceNameVariableStr+secretIdVariableStr+userNamedVariableStr+PecomanagedDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Create, pecomanagedDatabaseInsightRepresentation), "opsi", "databaseInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiPecomanagedDatabaseInsightDestroy, []resource.TestStep{
		// verify create with optional opsiPrivateEndpointId
		{
			Config: config + compartmentIdVariableStr + dbsystemDatabaseIdVariableStr + opsiPrivateEndpointIdVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + PecomanagedDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Create, pecomanagedDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.credential_type", "CREDENTIALS_BY_VAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.role", "NORMAL"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.user_name"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "database_resource_type", "database"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "VIRTUAL_MACHINE"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "PE_COMANAGED_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "opsi_private_endpoint_id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
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
		// verify update to the password_secret_id (the secret_id will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + dbsystemDatabaseIdVariableStr + opsiPrivateEndpointIdVariableStr + serviceNameVariableStr + secretIdUVariableStr + userNamedVariableStr + PecomanagedDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(pecomanagedDatabaseInsightRepresentation, map[string]interface{}{
						"credential_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: pecomanagedDatabaseInsightCredentialDetailsForUpdateRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.credential_type", "CREDENTIALS_BY_VAULT"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.password_secret_id", secretIdU),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.role", "NORMAL"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.user_name"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "database_resource_type", "database"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "VIRTUAL_MACHINE"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "PE_COMANAGED_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "opsi_private_endpoint_id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
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
			Config: config + compartmentIdVariableStr + dbsystemDatabaseIdVariableStr + opsiPrivateEndpointIdVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + PecomanagedDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, pecomanagedDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "credential_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.credential_type", "CREDENTIALS_BY_VAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "credential_details.0.role", "NORMAL"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_details.0.user_name"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "database_resource_type", "database"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "VIRTUAL_MACHINE"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "PE_COMANAGED_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "opsi_private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "service_name", serviceName),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_database_insights", "test_database_insights", acctest.Optional, acctest.Update, pecomanagedDatabaseInsightDataSourceRepresentation) +
				compartmentIdVariableStr + dbsystemDatabaseIdVariableStr + opsiPrivateEndpointIdVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + PecomanagedDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, pecomanagedDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				//resource.TestCheckResourceAttr(datasourceName, "database_id.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_type.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "id.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "opsi_private_endpoint_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, pecomanagedDatabaseInsightSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbsystemDatabaseIdVariableStr + opsiPrivateEndpointIdVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + PecomanagedDatabaseInsightResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_details.#", "1"), //??
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_details.0.user_name", "system"),
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_details.0.credential_type", "CREDENTIALS_BY_VAULT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_details.0.role", "NORMAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_resource_type", "database"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_source", "PE_COMANAGED_DATABASE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + dbsystemDatabaseIdVariableStr + opsiPrivateEndpointIdVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + PecomanagedDatabaseInsightResourceConfig,
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + dbsystemDatabaseIdVariableStr + opsiPrivateEndpointIdVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + PecomanagedDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(pecomanagedDatabaseInsightRepresentation, map[string]interface{}{
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
			Config:            config + PecomanagedDatabaseInsightRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"deployment_type",
				"service_name",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOpsiPecomanagedDatabaseInsightDestroy(s *terraform.State) error {
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
	if !acctest.InSweeperExcludeList("OpsiPecomanagedDatabaseInsight") {
		resource.AddTestSweepers("OpsiPecomanagedDatabaseInsight", &resource.Sweeper{
			Name:         "OpsiPecomanagedDatabaseInsight",
			Dependencies: acctest.DependencyGraph["databaseInsight"],
			F:            sweepOpsiPecomanagedDatabaseInsightResource,
		})
	}
}

func sweepOpsiPecomanagedDatabaseInsightResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	databaseInsightIds, err := getPecomanagedDatabaseInsightIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseInsightId, pecomanagedDatabaseInsightSweepWaitCondition, time.Duration(3*time.Minute),
				pecomanagedDatabaseInsightSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getPecomanagedDatabaseInsightIds(compartment string) ([]string, error) {
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

func pecomanagedDatabaseInsightSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseInsightResponse, ok := response.Response.(oci_opsi.GetDatabaseInsightResponse); ok {
		return databaseInsightResponse.GetLifecycleState() != oci_opsi.LifecycleStateDeleted
	}
	return false
}

func pecomanagedDatabaseInsightSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetDatabaseInsight(context.Background(), oci_opsi.GetDatabaseInsightRequest{
		DatabaseInsightId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
