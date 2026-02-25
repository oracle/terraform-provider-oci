// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// telemetry is only supported with certain mysql versions
	mysqlMysqlDbSystemTelemetryMinMysqlVersion = "9.6.1"

	mysqlMysqlDbSystemWithTelemetry = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.2`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"mysql_version":           acctest.Representation{RepType: acctest.Required, Create: `${var.mysql_version}`},

		// avoid waiting for DBM to make test faster
		"database_management": acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		// disable backup policy to make test faster
		"backup_policy": acctest.RepresentationGroup{RepType: acctest.Required, Group: disabledBackupPolicy},

		"telemetry_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryConfigurationComplex},
	}

	mysqlMysqlDbSystemTelemetryConfigurationComplex = map[string]interface{}{
		"logs": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryConfigurationLogDestinationComplex},
			{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryConfigurationLogDestinationSimple},
		},
	}

	mysqlMysqlDbSystemTelemetryConfigurationSimple = map[string]interface{}{
		"logs": acctest.RepresentationGroup{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryConfigurationLogDestinationSimple},
	}

	mysqlMysqlDbSystemTelemetryConfigurationEmpty = map[string]interface{}{}

	mysqlMysqlDbSystemTelemetryConfigurationLogDestinationSimple = map[string]interface{}{
		"destination":                acctest.Representation{RepType: acctest.Required, Create: `LOG_ANALYTICS`},
		"log_types":                  acctest.Representation{RepType: acctest.Required, Create: []string{`ERROR_LOG`}},
		"destination_configurations": acctest.RepresentationGroup{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryDestinationConfigLogGroup},
	}

	mysqlMysqlDbSystemTelemetryConfigurationLogDestinationComplex = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `LOG_ANALYTICS`},
		"log_types":   acctest.Representation{RepType: acctest.Required, Create: []string{`AUDIT_LOG`, `GENERAL_LOG`, `SLOW_QUERY_LOG`}},
		"destination_configurations": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryDestinationConfigLogSet},
			{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryDestinationConfigLogGroup},
		},
	}

	mysqlMysqlDbSystemTelemetryDestinationConfigLogGroup = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `log-group-id`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${var.log_group_id}`},
	}

	mysqlMysqlDbSystemTelemetryDestinationConfigLogSet = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `log-set`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `some-log-set-value`},
	}

	mysqlMysqlDbSystemSingularDataSourceWithTelemetry = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
	}
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_telemetryTest(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_mydasTest")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	mysqlVersion := utils.GetEnvSettingWithDefault("mysql_version", mysqlMysqlDbSystemTelemetryMinMysqlVersion)
	mysqlVersionVariableStr := fmt.Sprintf("variable \"mysql_version\" { default = \"%s\" }\n", mysqlVersion)
	logGroupId := utils.GetEnvSettingWithBlankDefault("log_group_ocid")
	if logGroupId == "" {
		t.Fatal("Test requires OCID of existing OCI Log Analytics Log Group passed via environment")
	}
	logGroupIdVariableStr := fmt.Sprintf("variable \"log_group_id\" { default = \"%s\" }\n", logGroupId)
	extraVariables := compartmentIdVariableStr + mysqlVersionVariableStr + logGroupIdVariableStr

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"
	singularDatasourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var launchedDbSystemId string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Verify create with complex telemetry configuration
		{
			Config: config + extraVariables + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Update, mysqlMysqlDbSystemWithTelemetry),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.#", "2"),

				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.destination", "LOG_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.log_types.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.log_types.0", "AUDIT_LOG"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.log_types.1", "GENERAL_LOG"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.log_types.2", "SLOW_QUERY_LOG"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.destination_configurations.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(
					resourceName,
					"telemetry_configuration.0.logs.0.destination_configurations",
					map[string]string{"key": "log-group-id", "value": logGroupId},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(
					resourceName,
					"telemetry_configuration.0.logs.0.destination_configurations",
					map[string]string{"key": "log-set", "value": "some-log-set-value"},
					[]string{}),

				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.1.destination", "LOG_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.1.log_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.1.log_types.0", "ERROR_LOG"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.1.destination_configurations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(
					resourceName,
					"telemetry_configuration.0.logs.1.destination_configurations",
					map[string]string{"key": "log-group-id", "value": logGroupId},
					[]string{}),

				func(s *terraform.State) (err error) {
					launchedDbSystemId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Verify telemetry configuration in singular data source
		{
			Config: config + extraVariables + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Update, mysqlMysqlDbSystemWithTelemetry) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, mysqlMysqlDbSystemSingularDataSourceWithTelemetry),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.#", "2"),

				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.destination", "LOG_ANALYTICS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.log_types.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.log_types.0", "AUDIT_LOG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.log_types.1", "GENERAL_LOG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.log_types.2", "SLOW_QUERY_LOG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.destination_configurations.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.destination_configurations.0.key", "log-group-id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.destination_configurations.0.value", logGroupId),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.destination_configurations.1.key", "log-set"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.0.destination_configurations.1.value", "some-log-set-value"),

				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.1.destination", "LOG_ANALYTICS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.1.log_types.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.1.log_types.0", "ERROR_LOG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.1.destination_configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.1.destination_configurations.0.key", "log-group-id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "telemetry_configuration.0.logs.1.destination_configurations.0.value", logGroupId),
			),
		},

		// Verify update with simple telemetry configuration
		{
			Config: config + extraVariables + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Update,
					acctest.GetUpdatedRepresentationCopy(
						"telemetry_configuration",
						acctest.RepresentationGroup{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryConfigurationSimple},
						mysqlMysqlDbSystemWithTelemetry)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.destination", "LOG_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.log_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.log_types.0", "ERROR_LOG"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.0.destination_configurations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(
					resourceName,
					"telemetry_configuration.0.logs.0.destination_configurations",
					map[string]string{"key": "log-group-id", "value": logGroupId},
					[]string{}),

				func(s *terraform.State) (err error) {
					currentDbSystemId, err := acctest.FromInstanceState(s, resourceName, "id")
					if launchedDbSystemId != currentDbSystemId {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Verify update with empty telemetry configuration (disable all telemetry for logs)
		{
			Config: config + extraVariables + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Update,
					acctest.GetUpdatedRepresentationCopy(
						"telemetry_configuration",
						acctest.RepresentationGroup{RepType: acctest.Required, Group: mysqlMysqlDbSystemTelemetryConfigurationEmpty},
						mysqlMysqlDbSystemWithTelemetry)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "telemetry_configuration.0.logs.#", "0"),

				func(s *terraform.State) (err error) {
					currentDbSystemId, err := acctest.FromInstanceState(s, resourceName, "id")
					if launchedDbSystemId != currentDbSystemId {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
	})
}
