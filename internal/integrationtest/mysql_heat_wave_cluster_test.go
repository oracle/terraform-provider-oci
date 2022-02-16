// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v58/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	HeatWaveClusterRequiredOnlyResource = HeatWaveClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Required, acctest.Create, heatWaveClusterRepresentation)

	HeatWaveClusterResourceConfig = HeatWaveClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Optional, acctest.Update, heatWaveClusterRepresentation)

	heatWaveClusterSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
	}

	mysqlDbSystemHeatWaveRepresentation = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.HeatWave.VM.Standard.E3`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemBackupPolicyRepresentation},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `MySQL Database Service`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `HeatWave-DbSystem`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"hostname_label":          acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"ip_address":              acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.3`},
		"maintenance":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemMaintenanceRepresentation},
		"port":                    acctest.Representation{RepType: acctest.Optional, Create: `3306`},
		"port_x":                  acctest.Representation{RepType: acctest.Optional, Create: `33306`},
	}

	heatWaveClusterRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"cluster_size": acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"shape_name":   acctest.Representation{RepType: acctest.Required, Create: `MySQL.HeatWave.VM.Standard.E3`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`}, // testing stop & start actions
	}

	HeatWaveClusterResourceDependencies = MysqlConfigurationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", acctest.Required, acctest.Create, mysqlShapeDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, mysqlDbSystemHeatWaveRepresentation)
)

// issue-routing-tag: mysql/default
func TestMysqlHeatWaveClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlHeatWaveClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_heat_wave_cluster.test_heat_wave_cluster"

	singularDatasourceName := "data.oci_mysql_heat_wave_cluster.test_heat_wave_cluster"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckMysqlHeatWaveClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Required, acctest.Create, heatWaveClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceDependencies,
		},
		// verify Create & stop
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Optional, acctest.Create, heatWaveClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "2"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.HeatWave.VM.Standard.E3"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify start & updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Optional, acctest.Update, heatWaveClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "3"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.HeatWave.VM.Standard.E3"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Required, acctest.Create, heatWaveClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + HeatWaveClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_nodes.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_size", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_name", "MySQL.HeatWave.VM.Standard.E3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMysqlHeatWaveClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbSystemClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_heat_wave_cluster" {
			noResourceFound = false
			request := oci_mysql.GetHeatWaveClusterRequest{}

			if value, ok := rs.Primary.Attributes["db_system_id"]; ok {
				request.DbSystemId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")

			response, err := client.GetHeatWaveCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.HeatWaveClusterLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MysqlHeatWaveCluster") {
		resource.AddTestSweepers("MysqlHeatWaveCluster", &resource.Sweeper{
			Name:         "MysqlHeatWaveCluster",
			Dependencies: acctest.DependencyGraph["heatWaveCluster"],
			F:            sweepMysqlHeatWaveClusterResource,
		})
	}
}

func sweepMysqlHeatWaveClusterResource(compartment string) error {
	dbSystemClient := acctest.GetTestClients(&schema.ResourceData{}).DbSystemClient()
	mysqlDbSystemIds, err := getMysqlDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, mysqlDbSystemId := range mysqlDbSystemIds {
		if ok := acctest.SweeperDefaultResourceId[mysqlDbSystemId]; !ok {
			deleteHeatWaveClusterRequest := oci_mysql.DeleteHeatWaveClusterRequest{}
			deleteHeatWaveClusterRequest.DbSystemId = &mysqlDbSystemId

			deleteHeatWaveClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")
			_, error := dbSystemClient.DeleteHeatWaveCluster(context.Background(), deleteHeatWaveClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting HeatWaveCluster of DbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", mysqlDbSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mysqlDbSystemId, heatWaveClusterSweepWaitCondition, time.Duration(3*time.Minute),
				heatWaveClusterSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func heatWaveClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if heatWaveClusterResponse, ok := response.Response.(oci_mysql.GetHeatWaveClusterResponse); ok {
		return heatWaveClusterResponse.LifecycleState != oci_mysql.HeatWaveClusterLifecycleStateDeleted
	}
	return false
}

func heatWaveClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbSystemClient().GetHeatWaveCluster(context.Background(), oci_mysql.GetHeatWaveClusterRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
