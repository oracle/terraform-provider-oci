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
	LogGroupRequiredOnlyResource = LogGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupRepresentation)

	LogGroupResourceConfig = LogGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Optional, acctest.Update, logGroupRepresentation)

	logGroupSingularDataSourceRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
	}

	logGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"is_compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: logGroupDataSourceFilterRepresentation}}
	logGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_logging_log_group.test_log_group.id}`}},
	}

	logGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	logGroupUpdateRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayNameUpdate`, Update: `displayNameUpdate2`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	LogGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: logging/default
func TestLoggingLogGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingLogGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_logging_log_group.test_log_group"
	datasourceName := "data.oci_logging_log_groups.test_log_groups"
	singularDatasourceName := "data.oci_logging_log_group.test_log_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Optional, acctest.Create, logGroupRepresentation), "logging", "logGroup", t)

	acctest.ResourceTest(t, testAccCheckLoggingLogGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LogGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LogGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Optional, acctest.Create, logGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LogGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(logGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + LogGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Optional, acctest.Update, logGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_log_groups", "test_log_groups", acctest.Optional, acctest.Update, logGroupDataSourceRepresentation) +
				compartmentIdVariableStr + LogGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Optional, acctest.Update, logGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_compartment_id_in_subtree", "false"),

				resource.TestCheckResourceAttr(datasourceName, "log_groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "log_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "log_groups.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "log_groups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "log_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_groups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_groups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_groups.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_groups.0.time_last_modified"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + LogGroupResourceConfig,
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

func testAccCheckLoggingLogGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoggingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_logging_log_group" {
			noResourceFound = false
			request := oci_logging.GetLogGroupRequest{}

			tmp := rs.Primary.ID
			request.LogGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")

			_, err := client.GetLogGroup(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("LoggingLogGroup") {
		resource.AddTestSweepers("LoggingLogGroup", &resource.Sweeper{
			Name:         "LoggingLogGroup",
			Dependencies: acctest.DependencyGraph["logGroup"],
			F:            sweepLoggingLogGroupResource,
		})
	}
}

func sweepLoggingLogGroupResource(compartment string) error {
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()
	logGroupIds, err := getLogGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, logGroupId := range logGroupIds {
		if ok := acctest.SweeperDefaultResourceId[logGroupId]; !ok {
			deleteLogGroupRequest := oci_logging.DeleteLogGroupRequest{}

			deleteLogGroupRequest.LogGroupId = &logGroupId

			deleteLogGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")
			_, error := loggingManagementClient.DeleteLogGroup(context.Background(), deleteLogGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting LogGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", logGroupId, error)
				continue
			}
		}
	}
	return nil
}

func getLogGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LogGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()

	listLogGroupsRequest := oci_logging.ListLogGroupsRequest{}
	listLogGroupsRequest.CompartmentId = &compartmentId
	listLogGroupsResponse, err := loggingManagementClient.ListLogGroups(context.Background(), listLogGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LogGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, logGroup := range listLogGroupsResponse.Items {
		id := *logGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LogGroupId", id)
	}
	return resourceIds, nil
}
