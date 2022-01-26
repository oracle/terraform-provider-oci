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
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	MysqlDbSystemRequiredOnlyResource = MysqlDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, mysqlDbSystemRepresentation)

	MysqlDbSystemResourceConfig = MysqlDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemRepresentation)

	mysqlDbSystemSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
	}

	mysqlDbSystemDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"db_system_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `DBSystem001`, Update: `displayName2`},
		"is_analytics_cluster_attached": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_heat_wave_cluster_attached": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_up_to_date":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: mysqlDbSystemDataSourceFilterRepresentation}}
	mysqlDbSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_mysql_mysql_db_system.test_mysql_db_system.id}`}},
	}

	mysqlDbSystemRepresentation = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E2.2`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemBackupPolicyRepresentation},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `MySQL Database Service`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `DBSystem001`, Update: `displayName2`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":          acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"ip_address":              acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.3`},
		"is_highly_available":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"maintenance":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemMaintenanceRepresentation},
		"port":                    acctest.Representation{RepType: acctest.Optional, Create: `3306`},
		"port_x":                  acctest.Representation{RepType: acctest.Optional, Create: `33306`},
	}
	mysqlDbSystemBackupPolicyRepresentation = map[string]interface{}{
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"retention_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"window_start_time": acctest.Representation{RepType: acctest.Optional, Create: `01:00-00:00`, Update: `02:00-00:00`},
	}
	mysqlDbSystemMaintenanceRepresentation = map[string]interface{}{
		"window_start_time": acctest.Representation{RepType: acctest.Required, Create: `sun 01:00`},
	}

	MysqlDbSystemResourceDependencies = MysqlConfigurationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		MysqlVersionResourceConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"
	datasourceName := "data.oci_mysql_mysql_db_systems.test_mysql_db_systems"
	singularDatasourceName := "data.oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MysqlDbSystemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlDbSystemRepresentation), "mysql", "mysqlDbSystem", t)

	acctest.ResourceTest(t, testAccCheckMysqlMysqlDbSystemDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, mysqlDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlDbSystemRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Required, acctest.Create, analyticsClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", acctest.Required, acctest.Create, channelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "01:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "50"),
				resource.TestCheckResourceAttr(resourceName, "description", "MySQL Database Service"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DBSystem001"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(resourceName, "is_highly_available", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Required, acctest.Create, analyticsClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", acctest.Required, acctest.Create, channelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "50"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(resourceName, "is_highly_available", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(resourceName, "port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "port_x", "33306"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_db_systems", "test_mysql_db_systems", acctest.Optional, acctest.Update, mysqlDbSystemDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Required, acctest.Create, analyticsClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", acctest.Required, acctest.Create, channelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "configuration_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_analytics_cluster_attached", "true"),
				resource.TestCheckResourceAttr(datasourceName, "is_heat_wave_cluster_attached", "true"),
				resource.TestCheckResourceAttr(datasourceName, "is_up_to_date", "false"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "db_systems.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.analytics_cluster.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.current_placement.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.fault_domain", "FAULT-DOMAIN-1"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.heat_wave_cluster.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.is_analytics_cluster_attached", "true"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.is_heat_wave_cluster_attached", "true"),
				resource.TestCheckResourceAttr(datasourceName, "db_systems.0.is_highly_available", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, mysqlDbSystemSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlDbSystemResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Required, acctest.Create, analyticsClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", acctest.Required, acctest.Create, channelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "analytics_cluster.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.retention_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
				resource.TestCheckResourceAttr(singularDatasourceName, "channels.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "current_placement.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_gb", "50"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "heat_wave_cluster.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_analytics_cluster_attached", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_heat_wave_cluster_attached", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_highly_available", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance.0.window_start_time", "sun 01:00"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port", "3306"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port_x", "33306"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.source_type", "NONE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + MysqlDbSystemResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"admin_password",
				"admin_username",
				"shutdown_type",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMysqlMysqlDbSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbSystemClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_mysql_db_system" {
			noResourceFound = false
			request := oci_mysql.GetDbSystemRequest{}

			tmp := rs.Primary.ID
			request.DbSystemId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")

			response, err := client.GetDbSystem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.DbSystemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MysqlMysqlDbSystem") {
		resource.AddTestSweepers("MysqlMysqlDbSystem", &resource.Sweeper{
			Name:         "MysqlMysqlDbSystem",
			Dependencies: acctest.DependencyGraph["mysqlDbSystem"],
			F:            sweepMysqlMysqlDbSystemResource,
		})
	}
}

func sweepMysqlMysqlDbSystemResource(compartment string) error {
	dbSystemClient := acctest.GetTestClients(&schema.ResourceData{}).DbSystemClient()
	mysqlDbSystemIds, err := getMysqlDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, mysqlDbSystemId := range mysqlDbSystemIds {
		if ok := acctest.SweeperDefaultResourceId[mysqlDbSystemId]; !ok {
			deleteDbSystemRequest := oci_mysql.DeleteDbSystemRequest{}
			deleteDbSystemRequest.DbSystemId = &mysqlDbSystemId

			deleteDbSystemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")
			_, error := dbSystemClient.DeleteDbSystem(context.Background(), deleteDbSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting MysqlDbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", mysqlDbSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mysqlDbSystemId, mysqlDbSystemSweepWaitCondition, time.Duration(3*time.Minute),
				mysqlDbSystemSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getMysqlDbSystemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MysqlDbSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbSystemClient := acctest.GetTestClients(&schema.ResourceData{}).DbSystemClient()

	listDbSystemsRequest := oci_mysql.ListDbSystemsRequest{}
	listDbSystemsRequest.CompartmentId = &compartmentId
	listDbSystemsRequest.LifecycleState = oci_mysql.DbSystemLifecycleStateActive
	listDbSystemsResponse, err := dbSystemClient.ListDbSystems(context.Background(), listDbSystemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MysqlDbSystem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mysqlDbSystem := range listDbSystemsResponse.Items {
		id := *mysqlDbSystem.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MysqlDbSystemId", id)
	}
	return resourceIds, nil
}

func mysqlDbSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mysqlDbSystemResponse, ok := response.Response.(oci_mysql.GetDbSystemResponse); ok {
		return mysqlDbSystemResponse.LifecycleState != oci_mysql.DbSystemLifecycleStateDeleted
	}
	return false
}

func mysqlDbSystemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbSystemClient().GetDbSystem(context.Background(), oci_mysql.GetDbSystemRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
