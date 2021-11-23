// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_logging "github.com/oracle/oci-go-sdk/v53/logging"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CustomLogRequiredOnlyResource = LogResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", Required, Create, customLogRepresentation)

	CustomLogResourceConfig = LogResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", Optional, Update, customLogRepresentation)

	customLogDataSourceRepresentation = map[string]interface{}{
		"log_group_id": Representation{RepType: Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"display_name": Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"log_type":     Representation{RepType: Optional, Create: `CUSTOM`},
		"state":        Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":       RepresentationGroup{Required, logDataSourceFilterRepresentation}}

	customLogRepresentation = map[string]interface{}{
		"display_name":       Representation{RepType: Required, Create: `log`, Update: `displayName2`},
		"log_group_id":       Representation{RepType: Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           Representation{RepType: Required, Create: `CUSTOM`},
		"defined_tags":       Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         Representation{RepType: Optional, Create: `false`, Update: `true`},
		"retention_duration": Representation{RepType: Optional, Create: `30`, Update: `60`},
	}

	CustomLogResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation)
)

// issue-routing-tag: logging/default
func TestLoggingCustomLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingLogResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_logging_log.test_log"
	datasourceName := "data.oci_logging_logs.test_logs"
	singularDatasourceName := "data.oci_logging_log.test_log"

	var compositeId string
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoggingCustomLogDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + CustomLogResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", Required, Create, customLogRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "display_name", "log"),
					resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
					resource.TestCheckResourceAttr(resourceName, "log_type", "CUSTOM"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + CustomLogResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + CustomLogResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", Optional, Create, customLogRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "display_name", "log"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
					resource.TestCheckResourceAttr(resourceName, "log_type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "retention_duration", "30"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							logGroupId, _ := FromInstanceState(s, resourceName, "log_group_id")
							compositeId = getLogCompositeId(logGroupId, resId)
							if errExport := TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CustomLogResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", Optional, Update, customLogRepresentation) +
					GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", Required, Update, logGroupRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
					resource.TestCheckResourceAttr(resourceName, "log_type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "retention_duration", "60"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
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
					GenerateDataSourceFromRepresentationMap("oci_logging_logs", "test_logs", Optional, Update, customLogDataSourceRepresentation) +
					GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", Required, Update, logGroupRepresentation) +
					compartmentIdVariableStr + CustomLogResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", Optional, Update, customLogRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "log_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "log_type", "CUSTOM"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "logs.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "logs.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "logs.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "logs.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "logs.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "logs.0.is_enabled", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "logs.0.log_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "logs.0.log_type", "CUSTOM"),
					resource.TestCheckResourceAttr(datasourceName, "logs.0.retention_duration", "60"),
					resource.TestCheckResourceAttrSet(datasourceName, "logs.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "logs.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "logs.0.time_last_modified"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_logging_log", "test_log", Required, Create, logSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CustomLogResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "log_group_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "log_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "log_type", "CUSTOM"),
					resource.TestCheckResourceAttr(singularDatasourceName, "retention_duration", "60"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CustomLogResourceConfig,
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
		},
	})
}

func testAccCheckLoggingCustomLogDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loggingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_logging_log" {
			noResourceFound = false
			request := oci_logging.GetLogRequest{}

			if value, ok := rs.Primary.Attributes["log_group_id"]; ok {
				request.LogGroupId = &value
			}

			tmp := rs.Primary.ID
			request.LogId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "logging")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("CustomLoggingLog") {
		resource.AddTestSweepers("CustomLoggingLog", &resource.Sweeper{
			Name:         "CustomLoggingLog",
			Dependencies: DependencyGraph["log"],
			F:            sweepLoggingCustomLogResource,
		})
	}
}

func sweepLoggingCustomLogResource(compartment string) error {
	loggingManagementClient := GetTestClients(&schema.ResourceData{}).loggingManagementClient()
	logIds, err := getLogIds(compartment)
	if err != nil {
		return err
	}
	for _, logId := range logIds {
		if ok := SweeperDefaultResourceId[logId]; !ok {
			deleteLogRequest := oci_logging.DeleteLogRequest{}

			deleteLogRequest.LogId = &logId

			deleteLogRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "logging")
			_, error := loggingManagementClient.DeleteLog(context.Background(), deleteLogRequest)
			if error != nil {
				fmt.Printf("Error deleting Log %s %s, It is possible that the resource is already deleted. Please verify manually \n", logId, error)
				continue
			}
		}
	}
	return nil
}
