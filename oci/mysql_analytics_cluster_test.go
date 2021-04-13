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
	"github.com/oracle/oci-go-sdk/v39/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v39/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AnalyticsClusterRequiredOnlyResource = AnalyticsClusterResourceDependencies +
		generateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", Required, Create, analyticsClusterRepresentation)

	AnalyticsClusterResourceConfig = AnalyticsClusterResourceDependencies +
		generateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", Optional, Update, analyticsClusterRepresentation)

	analyticsClusterSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": Representation{repType: Required, create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
	}

	analyticsClusterRepresentation = map[string]interface{}{
		"db_system_id": Representation{repType: Required, create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"cluster_size": Representation{repType: Required, create: `2`, update: `3`},
		"shape_name":   Representation{repType: Required, create: `VM.Standard.E2.2`, update: `VM.Standard.E2.4`},
		"state":        Representation{repType: Optional, create: `INACTIVE`, update: `ACTIVE`}, // testing stop & start actions
	}

	AnalyticsClusterResourceDependencies = MysqlConfigurationResourceConfig +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		generateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", Required, Create, mysqlShapeDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Required, Create, mysqlDbSystemRepresentation)
)

func TestMysqlAnalyticsClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlAnalyticsClusterResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_analytics_cluster.test_analytics_cluster"
	singularDatasourceName := "data.oci_mysql_analytics_cluster.test_analytics_cluster"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+AnalyticsClusterResourceDependencies+
		generateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", Optional, Create, analyticsClusterRepresentation), "mysql", "analyticsCluster", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckMysqlAnalyticsClusterDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AnalyticsClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", Required, Create, analyticsClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AnalyticsClusterResourceDependencies,
			},

			// verify create & stop
			{
				Config: config + compartmentIdVariableStr + AnalyticsClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", Optional, Create, analyticsClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "cluster_size", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape_name", "VM.Standard.E2.2"),
					resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

					resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

			// verify start & updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AnalyticsClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", Optional, Update, analyticsClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "cluster_size", "3"),
					resource.TestCheckResourceAttr(resourceName, "shape_name", "VM.Standard.E2.4"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", Required, Create, analyticsClusterSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AnalyticsClusterResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "cluster_nodes.#", "3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cluster_size", "3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_name", "VM.Standard.E2.4"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AnalyticsClusterResourceConfig,
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

func testAccCheckMysqlAnalyticsClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dbSystemClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_analytics_cluster" {
			noResourceFound = false
			request := oci_mysql.GetAnalyticsClusterRequest{}

			if value, ok := rs.Primary.Attributes["db_system_id"]; ok {
				request.DbSystemId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "mysql")

			response, err := client.GetAnalyticsCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.AnalyticsClusterLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("MysqlAnalyticsCluster") {
		resource.AddTestSweepers("MysqlAnalyticsCluster", &resource.Sweeper{
			Name:         "MysqlAnalyticsCluster",
			Dependencies: DependencyGraph["analyticsCluster"],
			F:            sweepMysqlAnalyticsClusterResource,
		})
	}
}

func sweepMysqlAnalyticsClusterResource(compartment string) error {
	dbSystemClient := GetTestClients(&schema.ResourceData{}).dbSystemClient()
	mysqlDbSystemIds, err := getMysqlDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, mysqlDbSystemId := range mysqlDbSystemIds {
		if ok := SweeperDefaultResourceId[mysqlDbSystemId]; !ok {
			deleteAnalyticsClusterRequest := oci_mysql.DeleteAnalyticsClusterRequest{}
			deleteAnalyticsClusterRequest.DbSystemId = &mysqlDbSystemId

			deleteAnalyticsClusterRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "mysql")
			_, error := dbSystemClient.DeleteAnalyticsCluster(context.Background(), deleteAnalyticsClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting AnalyticsCluster of DbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", mysqlDbSystemId, error)
				continue
			}
			waitTillCondition(testAccProvider, &mysqlDbSystemId, analyticsClusterSweepWaitCondition, time.Duration(3*time.Minute),
				analyticsClusterSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func analyticsClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if analyticsClusterResponse, ok := response.Response.(oci_mysql.GetAnalyticsClusterResponse); ok {
		return analyticsClusterResponse.LifecycleState != oci_mysql.AnalyticsClusterLifecycleStateDeleted
	}
	return false
}

func analyticsClusterSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dbSystemClient().GetAnalyticsCluster(context.Background(), oci_mysql.GetAnalyticsClusterRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
