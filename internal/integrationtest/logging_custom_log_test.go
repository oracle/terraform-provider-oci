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
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_logging "github.com/oracle/oci-go-sdk/v58/logging"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CustomLogRequiredOnlyResource = LogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, customLogRepresentation)

	CustomLogResourceConfig = LogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Update, customLogRepresentation)

	customLogDataSourceRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"log_type":     acctest.Representation{RepType: acctest.Optional, Create: `CUSTOM`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: logDataSourceFilterRepresentation}}

	customLogRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `log`, Update: `displayName2`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `60`},
	}

	CustomLogResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupRepresentation)
)

// issue-routing-tag: logging/default
func TestLoggingCustomLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingLogResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_logging_log.test_log"
	datasourceName := "data.oci_logging_logs.test_logs"
	singularDatasourceName := "data.oci_logging_log.test_log"

	var compositeId string
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoggingCustomLogDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + CustomLogResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, customLogRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "display_name", "log"),
					resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
					resource.TestCheckResourceAttr(resourceName, "log_type", "CUSTOM"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
					acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Create, customLogRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "display_name", "log"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
					resource.TestCheckResourceAttr(resourceName, "log_type", "CUSTOM"),
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
				Config: config + compartmentIdVariableStr + CustomLogResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Update, customLogRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Update, logGroupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
					resource.TestCheckResourceAttr(resourceName, "log_type", "CUSTOM"),
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_logging_logs", "test_logs", acctest.Optional, acctest.Update, customLogDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Update, logGroupRepresentation) +
					compartmentIdVariableStr + CustomLogResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Update, customLogRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, logSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CustomLogResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	if !acctest.InSweeperExcludeList("CustomLoggingLog") {
		resource.AddTestSweepers("CustomLoggingLog", &resource.Sweeper{
			Name:         "CustomLoggingLog",
			Dependencies: acctest.DependencyGraph["log"],
			F:            sweepLoggingCustomLogResource,
		})
	}
}

func sweepLoggingCustomLogResource(compartment string) error {
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
