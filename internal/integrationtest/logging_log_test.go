// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/logging"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_logging "github.com/oracle/oci-go-sdk/v56/logging"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	LogRequiredOnlyResource = LogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, logRepresentation)

	LogResourceConfig = LogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Update, logRepresentation)

	logSingularDataSourceRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_log.id}`},
	}

	logDataSourceRepresentation = map[string]interface{}{
		"log_group_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"log_type":        acctest.Representation{RepType: acctest.Optional, Create: `SERVICE`},
		"source_resource": acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"source_service":  acctest.Representation{RepType: acctest.Optional, Create: `objectstorage`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: logDataSourceFilterRepresentation}}
	logDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_logging_log.test_log.id}`}},
	}

	logRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `SERVICE`},
		"configuration":      acctest.RepresentationGroup{RepType: acctest.Required, Group: logConfigurationRepresentation},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `60`},
	}
	logConfigurationRepresentation = map[string]interface{}{
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: logConfigurationSourceRepresentation},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}
	logConfigurationSourceRepresentation = map[string]interface{}{
		"category":    acctest.Representation{RepType: acctest.Required, Create: `write`},
		"resource":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"service":     acctest.Representation{RepType: acctest.Required, Create: `objectstorage`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OCISERVICE`},
	}

	LogResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Create, logGroupUpdateRepresentation)
)

// issue-routing-tag: logging/default
func TestLoggingLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_logging_log.test_log"
	datasourceName := "data.oci_logging_logs.test_logs"
	singularDatasourceName := "data.oci_logging_log.test_log"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Create, logRepresentation), "logging", "log", t)

	var compositeId string

	acctest.ResourceTest(t, testAccCheckLoggingLogDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, logRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_type", "SERVICE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LogResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Create, logRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.0.category", "write"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.0.resource", testBucketName),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.0.service", "objectstorage"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.0.source_type", "OCISERVICE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_type", "SERVICE"),
				resource.TestCheckResourceAttr(resourceName, "retention_duration", "30"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						logGroupId, _ := acctest.FromInstanceState(s, resourceName, "log_group_id")
						compositeId = logging.GetLogCompositeId(logGroupId, resId)
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + LogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Update, logRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.0.category", "write"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.0.resource", testBucketName),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.0.service", "objectstorage"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.source.0.source_type", "OCISERVICE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_type", "SERVICE"),
				resource.TestCheckResourceAttr(resourceName, "retention_duration", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_logs", "test_logs", acctest.Optional, acctest.Update, logDataSourceRepresentation) +
				compartmentIdVariableStr + LogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Update, logRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "log_type", "SERVICE"),
				resource.TestCheckResourceAttr(datasourceName, "source_resource", testBucketName),
				resource.TestCheckResourceAttr(datasourceName, "source_service", "objectstorage"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "logs.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "logs.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.configuration.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.configuration.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.configuration.0.source.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.configuration.0.source.0.category", "write"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.configuration.0.source.0.resource", testBucketName),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.configuration.0.source.0.service", "objectstorage"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.configuration.0.source.0.source_type", "OCISERVICE"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "logs.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "logs.0.log_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.log_type", "SERVICE"),
				resource.TestCheckResourceAttr(datasourceName, "logs.0.retention_duration", "60"),
				resource.TestCheckResourceAttrSet(datasourceName, "logs.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "logs.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "logs.0.time_last_modified"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, logSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_group_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.source.0.category", "write"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.source.0.resource", testBucketName),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.source.0.service", "objectstorage"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.source.0.source_type", "OCISERVICE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_type", "SERVICE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retention_duration", "60"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + LogResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ImportStateIdFunc:       GetLogResourceCompositeId(resourceName),
			ResourceName:            resourceName,
		},
	})
}

func GetLogResourceCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("logGroupId/%s/logId/%s", rs.Primary.Attributes["log_group_id"], rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckLoggingLogDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoggingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_logging_log" {
			noResourceFound = false
			request := oci_logging.GetLogRequest{}

			if value, ok := rs.Primary.Attributes["log_group_id"]; ok {
				request.LogGroupId = &value
			}

			tmp := rs.Primary.ID
			request.LogId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")

			_, err := client.GetLog(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("LoggingLog") {
		resource.AddTestSweepers("LoggingLog", &resource.Sweeper{
			Name:         "LoggingLog",
			Dependencies: acctest.DependencyGraph["log"],
			F:            sweepLoggingLogResource,
		})
	}
}

func sweepLoggingLogResource(compartment string) error {
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()
	logIds, err := getLogIds(compartment)
	if err != nil {
		return err
	}
	for _, logId := range logIds {
		if ok := acctest.SweeperDefaultResourceId[logId]; !ok {
			deleteLogRequest := oci_logging.DeleteLogRequest{}

			deleteLogRequest.LogId = &logId

			deleteLogRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")
			_, error := loggingManagementClient.DeleteLog(context.Background(), deleteLogRequest)
			if error != nil {
				fmt.Printf("Error deleting Log %s %s, It is possible that the resource is already deleted. Please verify manually \n", logId, error)
				continue
			}
		}
	}
	return nil
}

func getLogIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LogId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()

	listLogsRequest := oci_logging.ListLogsRequest{}

	logGroupIds, error := getLogGroupIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting logGroupId required for Log resource requests \n")
	}
	for _, logGroupId := range logGroupIds {
		listLogsRequest.LogGroupId = &logGroupId

		listLogsResponse, err := loggingManagementClient.ListLogs(context.Background(), listLogsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Log list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, log := range listLogsResponse.Items {
			id := *log.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LogId", id)
		}

	}
	return resourceIds, nil
}
