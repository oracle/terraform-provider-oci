// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

var (
	ApplicationVipRequiredOnlyResource = ApplicationVipResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_application_vip", "test_application_vip", acctest.Required, acctest.Create, applicationVipRepresentation)

	ApplicationVipResourceConfig = ApplicationVipResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_application_vip", "test_application_vip", acctest.Optional, acctest.Update, applicationVipRepresentation)

	applicationVipSingularDataSourceRepresentation = map[string]interface{}{
		"application_vip_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_application_vip.test_application_vip.id}`},
	}

	applicationVipDataSourceRepresentation = map[string]interface{}{
		"cloud_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: applicationVipDataSourceFilterRepresentation}}
	applicationVipDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_application_vip.test_application_vip.id}`}},
	}

	applicationVipRepresentation = map[string]interface{}{
		"cloud_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
		"hostname_label":      acctest.Representation{RepType: acctest.Required, Create: `hostnameLabel`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet1.id}`},
		"db_node_id":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}`},
	}

	ApplicationVipResourceDependencies = DatabaseCloudVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudVmClusterRepresentation) +
		`data "oci_database_db_nodes" "test_db_nodes" {
			compartment_id = "${var.compartment_id}"
			vm_cluster_id = "${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}"
		}` +
		AvailabilityDomainConfig
)

// issue-routing-tag: database/default
func TestDatabaseApplicationVipResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseApplicationVipResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_application_vip.test_application_vip"
	datasourceName := "data.oci_database_application_vips.test_application_vips"
	singularDatasourceName := "data.oci_database_application_vip.test_application_vip"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApplicationVipResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_application_vip", "test_application_vip", acctest.Optional, acctest.Create, applicationVipRepresentation), "database", "applicationVip", t)

	acctest.ResourceTest(t, testAccCheckDatabaseApplicationVipDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApplicationVipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_application_vip", "test_application_vip", acctest.Required, acctest.Create, applicationVipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApplicationVipResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApplicationVipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_application_vip", "test_application_vip", acctest.Optional, acctest.Create, applicationVipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_assigned"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_application_vips", "test_application_vips", acctest.Optional, acctest.Update, applicationVipDataSourceRepresentation) +
				compartmentIdVariableStr + ApplicationVipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_application_vip", "test_application_vip", acctest.Optional, acctest.Update, applicationVipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "application_vips.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "application_vips.0.cloud_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "application_vips.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "application_vips.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(datasourceName, "application_vips.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "application_vips.0.ip_address"),
				resource.TestCheckResourceAttrSet(datasourceName, "application_vips.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "application_vips.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "application_vips.0.time_assigned"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_application_vip", "test_application_vip", acctest.Required, acctest.Create, applicationVipSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApplicationVipResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_vip_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_assigned"),
			),
		},
		// verify resource import
		{
			Config:            config + ApplicationVipRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"db_node_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseApplicationVipDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_application_vip" {
			noResourceFound = false
			request := oci_database.GetApplicationVipRequest{}

			tmp := rs.Primary.ID
			request.ApplicationVipId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetApplicationVip(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ApplicationVipLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseApplicationVip") {
		resource.AddTestSweepers("DatabaseApplicationVip", &resource.Sweeper{
			Name:         "DatabaseApplicationVip",
			Dependencies: acctest.DependencyGraph["applicationVip"],
			F:            sweepDatabaseApplicationVipResource,
		})
	}
}

func sweepDatabaseApplicationVipResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	applicationVipIds, err := getApplicationVipIds(compartment)
	if err != nil {
		return err
	}
	for _, applicationVipId := range applicationVipIds {
		if ok := acctest.SweeperDefaultResourceId[applicationVipId]; !ok {
			deleteApplicationVipRequest := oci_database.DeleteApplicationVipRequest{}

			deleteApplicationVipRequest.ApplicationVipId = &applicationVipId

			deleteApplicationVipRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteApplicationVip(context.Background(), deleteApplicationVipRequest)
			if error != nil {
				fmt.Printf("Error deleting ApplicationVip %s %s, It is possible that the resource is already deleted. Please verify manually \n", applicationVipId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &applicationVipId, applicationVipSweepWaitCondition, time.Duration(3*time.Minute),
				applicationVipSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getApplicationVipIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApplicationVipId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listApplicationVipsRequest := oci_database.ListApplicationVipsRequest{}
	listApplicationVipsRequest.CompartmentId = &compartmentId

	cloudVmClusterIds, error := getDatabaseCloudVmClusterIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting cloudVmClusterId required for ApplicationVip resource requests \n")
	}
	for _, cloudVmClusterId := range cloudVmClusterIds {
		listApplicationVipsRequest.CloudVmClusterId = &cloudVmClusterId

		listApplicationVipsRequest.LifecycleState = oci_database.ApplicationVipSummaryLifecycleStateAvailable
		listApplicationVipsResponse, err := databaseClient.ListApplicationVips(context.Background(), listApplicationVipsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ApplicationVip list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, applicationVip := range listApplicationVipsResponse.Items {
			id := *applicationVip.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApplicationVipId", id)
		}

	}
	return resourceIds, nil
}

func applicationVipSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if applicationVipResponse, ok := response.Response.(oci_database.GetApplicationVipResponse); ok {
		return applicationVipResponse.LifecycleState != oci_database.ApplicationVipLifecycleStateTerminated
	}
	return false
}

func applicationVipSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetApplicationVip(context.Background(), oci_database.GetApplicationVipRequest{
		ApplicationVipId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
