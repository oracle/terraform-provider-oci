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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreDatabaseManagementCloudExadataInfrastructureRunDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `license_model`}},
	}

	DatabaseManagementCloudExadataInfrastructureRequiredOnlyResource = DatabaseManagementCloudExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementCloudExadataInfrastructureRepresentation)

	DatabaseManagementCloudExadataInfrastructureResourceConfig = DatabaseManagementCloudExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Update, DatabaseManagementCloudExadataInfrastructureRepresentation)

	DatabaseManagementCloudExadataInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
	}

	DatabaseManagementCloudExadataInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `${var.display_name}`, Update: `${var.display_name}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementCloudExadataInfrastructureDataSourceFilterRepresentation}}
	DatabaseManagementCloudExadataInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`}},
	}

	DatabaseManagementCloudExadataInfrastructureRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vm_cluster_ids":       acctest.Representation{RepType: acctest.Required, Create: []string{`${var.vm_cluster_id}`}},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("Oracle-Tags.CreatedBy", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"discovery_key":        acctest.Representation{RepType: acctest.Optional, Create: `${var.discovery_key}`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `${var.display_name}`, Update: `${var.display_name}`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Required, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"license_model":        acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"storage_server_names": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.storage_server_name}`}},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatabaseManagementCloudExadataInfrastructureRunDefinedTagsRepresentation},
	}

	DatabaseManagementCloudExadataInfrastructureResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudExadataInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	storageServerName := utils.GetEnvSettingWithBlankDefault("storage_server_name")
	storageServerNameVariableStr := fmt.Sprintf("variable \"storage_server_name\" { default = \"%s\" }\n", storageServerName)

	displayName := utils.GetEnvSettingWithBlankDefault("display_name")
	displayNameVariableStr := fmt.Sprintf("variable \"display_name\" { default = \"%s\" }\n", displayName)

	vmClusterId := utils.GetEnvSettingWithBlankDefault("vm_cluster_id")
	vmClusterIdVariableStr := fmt.Sprintf("variable \"vm_cluster_id\" { default = \"%s\" }\n", vmClusterId)

	discoveryKey := utils.GetEnvSettingWithBlankDefault("discovery_key")
	discoveryKeyStr := fmt.Sprintf("variable \"discovery_key\" { default = \"%s\" }\n", discoveryKey)

	resourceName := "oci_database_management_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure"
	datasourceName := "data.oci_database_management_cloud_exadata_infrastructures.test_cloud_exadata_infrastructures"
	singularDatasourceName := "data.oci_database_management_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementCloudExadataInfrastructureResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Create, DatabaseManagementCloudExadataInfrastructureRepresentation), "databasemanagement", "cloudExadataInfrastructure", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementCloudExadataInfrastructureDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + storageServerNameVariableStr + displayNameVariableStr + vmClusterIdVariableStr + DatabaseManagementCloudExadataInfrastructureResourceDependencies + discoveryKeyStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementCloudExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "storage_server_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm_cluster_ids.#", "1"),
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
			Config: config + compartmentIdVariableStr + storageServerNameVariableStr + displayNameVariableStr + vmClusterIdVariableStr + DatabaseManagementCloudExadataInfrastructureResourceDependencies,
		},

		// verify Create
		{
			Config: config + compartmentIdVariableStr + storageServerNameVariableStr + displayNameVariableStr + vmClusterIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementCloudExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "vm_cluster_ids.#", "1"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + storageServerNameVariableStr + displayNameVariableStr + vmClusterIdVariableStr + DatabaseManagementCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Update, DatabaseManagementCloudExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "storage_server_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm_cluster_ids.#", "1"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructures", "test_cloud_exadata_infrastructures", acctest.Required, acctest.Create, DatabaseManagementCloudExadataInfrastructureDataSourceRepresentation) +
				compartmentIdVariableStr + storageServerNameVariableStr + displayNameVariableStr + vmClusterIdVariableStr + DatabaseManagementCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Update, DatabaseManagementCloudExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructure_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructure_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementCloudExadataInfrastructureSingularDataSourceRepresentation) +
				compartmentIdVariableStr + storageServerNameVariableStr + displayNameVariableStr + vmClusterIdVariableStr + DatabaseManagementCloudExadataInfrastructureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_compartments.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_grid.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_clusters.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudExadataInfrastructureRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"discovery_key",
				"storage_server_names",
				"vm_cluster_ids",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementCloudExadataInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_cloud_exadata_infrastructure" {
			noResourceFound = false
			request := oci_database_management.GetCloudExadataInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.CloudExadataInfrastructureId = &tmp

			fmt.Println("Inside testAccCheckDatabaseManagementCloudExadataInfrastructureDestroy")

			deleteCloudExadataInfrastructureRequest := oci_database_management.DeleteCloudExadataInfrastructureRequest{}

			deleteCloudExadataInfrastructureRequest.CloudExadataInfrastructureId = request.CloudExadataInfrastructureId

			deleteCloudExadataInfrastructureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			fmt.Println("Inside testAccCheckDatabaseManagementCloudExadataInfrastructureDestroy  starting delete for ", *request.CloudExadataInfrastructureId)
			_, error := client.DeleteCloudExadataInfrastructure(context.Background(), deleteCloudExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", *request.CloudExadataInfrastructureId, error)
				continue
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

			response, err := client.GetCloudExadataInfrastructure(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DatabaseManagementCloudExadataInfrastructure") {
		resource.AddTestSweepers("DatabaseManagementCloudExadataInfrastructure", &resource.Sweeper{
			Name:         "DatabaseManagementCloudExadataInfrastructure",
			Dependencies: acctest.DependencyGraph["cloudExadataInfrastructure"],
			F:            sweepDatabaseManagementCloudExadataInfrastructureResource,
		})
	}
}

func sweepDatabaseManagementCloudExadataInfrastructureResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	cloudExadataInfrastructureIds, err := getDatabaseManagementCloudExadataInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudExadataInfrastructureId := range cloudExadataInfrastructureIds {
		if ok := acctest.SweeperDefaultResourceId[cloudExadataInfrastructureId]; !ok {
			deleteCloudExadataInfrastructureRequest := oci_database_management.DeleteCloudExadataInfrastructureRequest{}

			deleteCloudExadataInfrastructureRequest.CloudExadataInfrastructureId = &cloudExadataInfrastructureId

			deleteCloudExadataInfrastructureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			fmt.Println("Inside sweepDatabaseManagementCloudExadataInfrastructureResource ", *(deleteCloudExadataInfrastructureRequest.CloudExadataInfrastructureId))
			_, error := dbManagementClient.DeleteCloudExadataInfrastructure(context.Background(), deleteCloudExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudExadataInfrastructureId, error)
				continue
			}

			acctest.WaitTillCondition(acctest.TestAccProvider, &cloudExadataInfrastructureId, DatabaseManagementCloudExadataInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementCloudExadataInfrastructureSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementCloudExadataInfrastructureIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudExadataInfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listCloudExadataInfrastructuresRequest := oci_database_management.ListCloudExadataInfrastructuresRequest{}
	listCloudExadataInfrastructuresRequest.CompartmentId = &compartmentId
	listCloudExadataInfrastructuresResponse, err := dbManagementClient.ListCloudExadataInfrastructures(context.Background(), listCloudExadataInfrastructuresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudExadataInfrastructure list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudExadataInfrastructure := range listCloudExadataInfrastructuresResponse.Items {
		id := *cloudExadataInfrastructure.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudExadataInfrastructureId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementCloudExadataInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudExadataInfrastructureResponse, ok := response.Response.(oci_database_management.GetCloudExadataInfrastructureResponse); ok {
		return cloudExadataInfrastructureResponse.GetLifecycleState() != oci_database_management.DbmResourceLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementCloudExadataInfrastructureSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetCloudExadataInfrastructure(context.Background(), oci_database_management.GetCloudExadataInfrastructureRequest{
		CloudExadataInfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
