// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MysqlHeatWaveClusterRequiredOnlyResource = MysqlHeatWaveWarehouseClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Required, acctest.Create, MysqlHeatWaveLakehouseClusterRepresentation)

	MysqlHeatWaveClusterResourceConfig = MysqlHeatWaveWarehouseClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Optional, acctest.Update, MysqlHeatWaveWarehouseClusterRepresentation)

	MysqlMysqlHeatWaveClusterSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
	}

	MysqlDbSystemHeatWaveRepresentation = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlDbSystemHeatWaveBackupPolicyRepresentation},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `MySQL Database Service`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `HeatWave-DbSystem`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"hostname_label":          acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"ip_address":              acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.3`},
		"maintenance":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemMaintenanceRepresentation},
		"port":                    acctest.Representation{RepType: acctest.Optional, Create: `3306`},
		"port_x":                  acctest.Representation{RepType: acctest.Optional, Create: `33306`},
	}

	MysqlDbSystemHeatWaveBackupPolicyRepresentation = map[string]interface{}{
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"is_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"pitr_policy":       acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlDbSystemHeatWavePitrPolicyRepresentation},
		"retention_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"window_start_time": acctest.Representation{RepType: acctest.Optional, Create: `01:00-00:00`},
	}

	// Lakehouse tests: Create - Warehouse tests: Update
	MysqlDbSystemHeatWavePitrPolicyRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}

	MysqlHeatWaveLakehouseClusterRepresentation = map[string]interface{}{
		"db_system_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"cluster_size":         acctest.Representation{RepType: acctest.Required, Create: `2`},
		"shape_name":           acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"state":                acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`},
		"is_lakehouse_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	MysqlHeatWaveWarehouseClusterRepresentation = map[string]interface{}{
		"db_system_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"cluster_size":         acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"shape_name":           acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"state":                acctest.Representation{RepType: acctest.Required, Create: `INACTIVE`, Update: `ACTIVE`}, // testing stop & start actions
		"is_lakehouse_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
	}

	MysqlHeatWaveClusterResourceDependenciesBase = MysqlMysqlConfigurationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", acctest.Required, acctest.Create, MysqlMysqlShapeDataSourceRepresentation)

	// DbSystem with PITR disabled
	MysqlHeatWaveLakehouseClusterResourceDependencies = MysqlHeatWaveClusterResourceDependenciesBase +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, MysqlDbSystemHeatWaveRepresentation)

	// DbSystem with PITR enabled
	MysqlHeatWaveWarehouseClusterResourceDependencies = MysqlHeatWaveClusterResourceDependenciesBase +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Update, MysqlDbSystemHeatWaveRepresentation)
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
			Config: config + compartmentIdVariableStr + MysqlHeatWaveLakehouseClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Required, acctest.Create, MysqlHeatWaveLakehouseClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "2"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.VM.Standard.E3.1.8GB"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "is_lakehouse_enabled", "true"),

				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Verify update enable Lakehouse
		{
			Config: config + compartmentIdVariableStr + MysqlHeatWaveLakehouseClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Required, acctest.Update, MysqlHeatWaveLakehouseClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "2"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.VM.Standard.E3.1.8GB"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "is_lakehouse_enabled", "false"),

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

		// deleting Lakehouse cluster before creating Warehouse cluster
		{
			Config: config + compartmentIdVariableStr + MysqlHeatWaveLakehouseClusterResourceDependencies,
		},
		// Update DbSystem to enable PITR policies
		{
			Config: config + compartmentIdVariableStr + MysqlHeatWaveWarehouseClusterResourceDependencies,
		},
		// verify Create & stop
		{
			Config: config + compartmentIdVariableStr + MysqlHeatWaveWarehouseClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Optional, acctest.Create, MysqlHeatWaveWarehouseClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "2"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.VM.Standard.E3.1.8GB"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "is_lakehouse_enabled", "false"),

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
			Config: config + compartmentIdVariableStr + MysqlHeatWaveWarehouseClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Optional, acctest.Update, MysqlHeatWaveWarehouseClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "3"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.VM.Standard.E3.1.8GB"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "is_lakehouse_enabled", "false"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", acctest.Required, acctest.Create, MysqlMysqlHeatWaveClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlHeatWaveClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_nodes.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_size", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_lakehouse_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_name", "MySQL.VM.Standard.E3.1.8GB"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + MysqlHeatWaveClusterRequiredOnlyResource,
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
	mysqlDbSystemIds, err := getMysqlMysqlDbSystemIds(compartment)
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
				MysqlHeatWaveClusterSweepResponseFetchOperation, "mysql", true)
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

func MysqlHeatWaveClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbSystemClient().GetHeatWaveCluster(context.Background(), oci_mysql.GetHeatWaveClusterRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
