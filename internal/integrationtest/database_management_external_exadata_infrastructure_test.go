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
	DatabaseManagementExternalExadataInfrastructureRequiredOnlyResource = DatabaseManagementExternalExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementExternalExadataInfrastructureRepresentation)

	DatabaseManagementExternalExadataInfrastructureResourceConfig = DatabaseManagementExternalExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Update, DatabaseManagementExternalExadataInfrastructureRepresentation)

	DatabaseManagementDatabaseManagementExternalExadataInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"external_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id}`},
	}

	DatabaseManagementDatabaseManagementExternalExadataInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `exadata-exaInfra01`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalExadataInfrastructureDataSourceFilterRepresentation}}
	DatabaseManagementExternalExadataInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id}`}},
	}

	DatabaseManagementExternalExadataInfrastructureRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_system_ids":        acctest.Representation{RepType: acctest.Required, Create: []string{`${var.db_system_id}`}},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `exadataInfra_Terraform_testing`, Update: `exadataInfra_Terraform_testingUpdate`},
		"discovery_key":        acctest.Representation{RepType: acctest.Optional, Create: `${var.discovery_key}`},
		"license_model":        acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"storage_server_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`scaqan10celadm07`}, Update: []string{`scaqan10celadm07`, `scaqan10celadm08`}},
	}

	DatabaseManagementExternalExadataInfrastructureResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalExadataInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	discoveryKey := utils.GetEnvSettingWithBlankDefault("discovery_key")
	discoveryKeyStr := fmt.Sprintf("variable \"discovery_key\" { default = \"%s\" }\n", discoveryKey)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("db_system_id")
	dbSystemIdStr := fmt.Sprintf("variable \"db_system_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure"
	datasourceName := "data.oci_database_management_external_exadata_infrastructures.test_external_exadata_infrastructures"
	singularDatasourceName := "data.oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementExternalExadataInfrastructureResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Optional, acctest.Create, DatabaseManagementExternalExadataInfrastructureRepresentation), "databasemanagement", "externalExadataInfrastructure", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementExternalExadataInfrastructureDestroy, []resource.TestStep{

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementExternalExadataInfrastructureResourceDependencies + discoveryKeyStr + dbSystemIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementExternalExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_system_ids.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "discovery_key"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "exadataInfra_Terraform_testing"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementExternalExadataInfrastructureResourceDependencies,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementExternalExadataInfrastructureResourceDependencies + dbSystemIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementExternalExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_system_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "exadataInfra_Terraform_testing"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseManagementExternalExadataInfrastructureResourceDependencies + dbSystemIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseManagementExternalExadataInfrastructureRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "db_system_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "exadataInfra_Terraform_testing"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),

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
			Config: config + compartmentIdVariableStr + DatabaseManagementExternalExadataInfrastructureResourceDependencies + dbSystemIdStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Update, DatabaseManagementExternalExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_system_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "exadataInfra_Terraform_testingUpdate"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),

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
			Config: config + dbSystemIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_infrastructures", "test_external_exadata_infrastructures", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalExadataInfrastructureDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementExternalExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "external_exadata_infrastructure_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_exadata_infrastructure_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + dbSystemIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalExadataInfrastructureSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalExadataInfrastructureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_exadata_infrastructure_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_compartments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_systems.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "exadataInfra_Terraform_testingUpdate"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_grid.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalExadataInfrastructureRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"db_system_ids",
				"discovery_key",
				"storage_server_names",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementExternalExadataInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_external_exadata_infrastructure" {
			noResourceFound = false
			request := oci_database_management.GetExternalExadataInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.ExternalExadataInfrastructureId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetExternalExadataInfrastructure(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.DbmResourceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementExternalExadataInfrastructure") {
		resource.AddTestSweepers("DatabaseManagementExternalExadataInfrastructure", &resource.Sweeper{
			Name:         "DatabaseManagementExternalExadataInfrastructure",
			Dependencies: acctest.DependencyGraph["externalExadataInfrastructure"],
			F:            sweepDatabaseManagementExternalExadataInfrastructureResource,
		})
	}
}

func sweepDatabaseManagementExternalExadataInfrastructureResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	externalExadataInfrastructureIds, err := getDatabaseManagementExternalExadataInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, externalExadataInfrastructureId := range externalExadataInfrastructureIds {
		if ok := acctest.SweeperDefaultResourceId[externalExadataInfrastructureId]; !ok {
			deleteExternalExadataInfrastructureRequest := oci_database_management.DeleteExternalExadataInfrastructureRequest{}

			deleteExternalExadataInfrastructureRequest.ExternalExadataInfrastructureId = &externalExadataInfrastructureId

			deleteExternalExadataInfrastructureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteExternalExadataInfrastructure(context.Background(), deleteExternalExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalExadataInfrastructureId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalExadataInfrastructureId, DatabaseManagementExternalExadataInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementExternalExadataInfrastructureSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementExternalExadataInfrastructureIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalExadataInfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listExternalExadataInfrastructuresRequest := oci_database_management.ListExternalExadataInfrastructuresRequest{}
	listExternalExadataInfrastructuresRequest.CompartmentId = &compartmentId
	listExternalExadataInfrastructuresResponse, err := dbManagementClient.ListExternalExadataInfrastructures(context.Background(), listExternalExadataInfrastructuresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalExadataInfrastructure list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalExadataInfrastructure := range listExternalExadataInfrastructuresResponse.Items {
		id := *externalExadataInfrastructure.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalExadataInfrastructureId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementExternalExadataInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalExadataInfrastructureResponse, ok := response.Response.(oci_database_management.GetExternalExadataInfrastructureResponse); ok {
		return externalExadataInfrastructureResponse.GetLifecycleState() != oci_database_management.DbmResourceLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementExternalExadataInfrastructureSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetExternalExadataInfrastructure(context.Background(), oci_database_management.GetExternalExadataInfrastructureRequest{
		ExternalExadataInfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
