// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v25/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v25/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MysqlDbSystemRequiredOnlyResource = MysqlDbSystemResourceDependencies +
		generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Required, Create, mysqlDbSystemRepresentation)

	MysqlDbSystemResourceConfig = MysqlDbSystemResourceDependencies +
		generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, mysqlDbSystemRepresentation)

	mysqlDbSystemSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": Representation{repType: Required, create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
	}

	mysqlDbSystemDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"configuration_id": Representation{repType: Optional, create: `${var.MysqlConfigurationOCID[var.region]}`},
		"db_system_id":     Representation{repType: Optional, create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"display_name":     Representation{repType: Optional, create: `DBSystem001`, update: `displayName2`},
		"is_up_to_date":    Representation{repType: Optional, create: `false`},
		"state":            Representation{repType: Optional, create: `ACTIVE`},
		"filter":           RepresentationGroup{Required, mysqlDbSystemDataSourceFilterRepresentation}}
	mysqlDbSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_mysql_mysql_db_system.test_mysql_db_system.id}`}},
	}

	mysqlDbSystemRepresentation = map[string]interface{}{
		"admin_password":          Representation{repType: Required, create: `BEstrO0ng_#11`},
		"admin_username":          Representation{repType: Required, create: `adminUser`},
		"availability_domain":     Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"configuration_id":        Representation{repType: Required, create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":              Representation{repType: Required, create: `VM.Standard.E2.2`},
		"subnet_id":               Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           RepresentationGroup{Optional, mysqlDbSystemBackupPolicyRepresentation},
		"data_storage_size_in_gb": Representation{repType: Required, create: `50`},
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             Representation{repType: Optional, create: `MySQL Database Service`, update: `description2`},
		"display_name":            Representation{repType: Optional, create: `DBSystem001`, update: `displayName2`},
		"fault_domain":            Representation{repType: Optional, create: `FAULT-DOMAIN-1`},
		"freeform_tags":           Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"hostname_label":          Representation{repType: Optional, create: `hostnameLabel`},
		"ip_address":              Representation{repType: Optional, create: `10.0.0.3`},
		"maintenance":             RepresentationGroup{Optional, mysqlDbSystemMaintenanceRepresentation},
		"mysql_version":           Representation{repType: Optional, create: `${data.oci_mysql_mysql_versions.test_mysql_versions.versions.0.versions.0.version}`},
		"port":                    Representation{repType: Optional, create: `3306`},
		"port_x":                  Representation{repType: Optional, create: `33306`},
	}
	mysqlDbSystemBackupPolicyRepresentation = map[string]interface{}{
		"defined_tags":      Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":     Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_enabled":        Representation{repType: Optional, create: `false`, update: `true`},
		"retention_in_days": Representation{repType: Optional, create: `10`, update: `11`},
		"window_start_time": Representation{repType: Optional, create: `01:00-00:00`, update: `02:00-00:00`},
	}
	mysqlDbSystemMaintenanceRepresentation = map[string]interface{}{
		"window_start_time": Representation{repType: Required, create: `sun 01:00`},
	}

	MysqlDbSystemResourceDependencies = MysqlConfigurationResourceConfig +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		MysqlVersionResourceConfig +
		DefinedTagsDependencies
)

func TestMysqlMysqlDbSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"
	datasourceName := "data.oci_mysql_mysql_db_systems.test_mysql_db_systems"
	singularDatasourceName := "data.oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckMysqlMysqlDbSystemDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + MysqlDbSystemResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Required, Create, mysqlDbSystemRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
					resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + MysqlDbSystemResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + MysqlDbSystemResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Create, mysqlDbSystemRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "01:00-00:00"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "50"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "MySQL Database Service"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "DBSystem001"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
					resource.TestCheckResourceAttr(resourceName, "mysql_version", "8.0.20"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, mysqlDbSystemRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "admin_username", "adminUser"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.retention_in_days", "11"),
					resource.TestCheckResourceAttr(resourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "50"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),
					resource.TestCheckResourceAttr(resourceName, "mysql_version", "8.0.20"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_mysql_mysql_db_systems", "test_mysql_db_systems", Optional, Update, mysqlDbSystemDataSourceRepresentation) +
					compartmentIdVariableStr + MysqlDbSystemResourceDependencies +
					generateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Optional, Update, mysqlDbSystemRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "configuration_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "is_up_to_date", "false"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "db_systems.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "db_systems.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "db_systems.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "db_systems.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "db_systems.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "db_systems.0.endpoints.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "db_systems.0.fault_domain", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(datasourceName, "db_systems.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "db_systems.0.mysql_version", "8.0.20"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_systems.0.time_updated"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Required, Create, mysqlDbSystemSingularDataSourceRepresentation) +
					compartmentIdVariableStr + MysqlDbSystemResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.retention_in_days", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backup_policy.0.window_start_time", "02:00-00:00"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_gb", "50"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "endpoints.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "NONE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "maintenance.0.window_start_time", "sun 01:00"),
					resource.TestCheckResourceAttr(singularDatasourceName, "mysql_version", "8.0.20"),
					resource.TestCheckResourceAttr(singularDatasourceName, "port", "3306"),
					resource.TestCheckResourceAttr(singularDatasourceName, "port_x", "33306"),
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
				Config:            config + generateResourceImportConfig("oci_mysql_mysql_db_system", "test_mysql_db_system"),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"admin_password",
					"admin_username",
					"shutdown_type",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckMysqlMysqlDbSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dbSystemClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_mysql_db_system" {
			noResourceFound = false
			request := oci_mysql.GetDbSystemRequest{}

			tmp := rs.Primary.ID
			request.DbSystemId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "mysql")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("MysqlMysqlDbSystem") {
		resource.AddTestSweepers("MysqlMysqlDbSystem", &resource.Sweeper{
			Name:         "MysqlMysqlDbSystem",
			Dependencies: DependencyGraph["mysqlDbSystem"],
			F:            sweepMysqlMysqlDbSystemResource,
		})
	}
}

func sweepMysqlMysqlDbSystemResource(compartment string) error {
	dbSystemClient := GetTestClients(&schema.ResourceData{}).dbSystemClient()
	mysqlDbSystemIds, err := getMysqlDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, mysqlDbSystemId := range mysqlDbSystemIds {
		if ok := SweeperDefaultResourceId[mysqlDbSystemId]; !ok {
			deleteDbSystemRequest := oci_mysql.DeleteDbSystemRequest{}
			deleteDbSystemRequest.DbSystemId = &mysqlDbSystemId

			deleteDbSystemRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "mysql")
			_, error := dbSystemClient.DeleteDbSystem(context.Background(), deleteDbSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting MysqlDbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", mysqlDbSystemId, error)
				continue
			}
			waitTillCondition(testAccProvider, &mysqlDbSystemId, mysqlDbSystemSweepWaitCondition, time.Duration(3*time.Minute),
				mysqlDbSystemSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getMysqlDbSystemIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "MysqlDbSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbSystemClient := GetTestClients(&schema.ResourceData{}).dbSystemClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "MysqlDbSystemId", id)
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

func mysqlDbSystemSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dbSystemClient().GetDbSystem(context.Background(), oci_mysql.GetDbSystemRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
