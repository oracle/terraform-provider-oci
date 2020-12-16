// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v31/common"
	oci_logging "github.com/oracle/oci-go-sdk/v31/logging"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	LogGroupRequiredOnlyResource = LogGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation)

	LogGroupResourceConfig = LogGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Optional, Update, logGroupRepresentation)

	logGroupSingularDataSourceRepresentation = map[string]interface{}{
		"log_group_id": Representation{repType: Required, create: `${oci_logging_log_group.test_log_group.id}`},
	}

	logGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                 Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"is_compartment_id_in_subtree": Representation{repType: Optional, create: `false`},
		"filter":                       RepresentationGroup{Required, logGroupDataSourceFilterRepresentation}}
	logGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_logging_log_group.test_log_group.id}`}},
	}

	logGroupRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	logGroupUpdateRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `displayNameUpdate`, update: `displayNameUpdate2`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	LogGroupResourceDependencies = DefinedTagsDependencies
)

func TestLoggingLogGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingLogGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_logging_log_group.test_log_group"
	datasourceName := "data.oci_logging_log_groups.test_log_groups"
	singularDatasourceName := "data.oci_logging_log_group.test_log_group"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoggingLogGroupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + LogGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + LogGroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + LogGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Optional, Create, logGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LogGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Optional, Create,
						representationCopyWithNewProperties(logGroupRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Optional, Update, logGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

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
					generateDataSourceFromRepresentationMap("oci_logging_log_groups", "test_log_groups", Optional, Update, logGroupDataSourceRepresentation) +
					compartmentIdVariableStr + LogGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Optional, Update, logGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "is_compartment_id_in_subtree", "false"),

					resource.TestCheckResourceAttr(datasourceName, "log_groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "log_groups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "log_groups.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + LogGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "log_group_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		},
	})
}

func testAccCheckLoggingLogGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loggingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_logging_log_group" {
			noResourceFound = false
			request := oci_logging.GetLogGroupRequest{}

			tmp := rs.Primary.ID
			request.LogGroupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "logging")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("LoggingLogGroup") {
		resource.AddTestSweepers("LoggingLogGroup", &resource.Sweeper{
			Name:         "LoggingLogGroup",
			Dependencies: DependencyGraph["logGroup"],
			F:            sweepLoggingLogGroupResource,
		})
	}
}

func sweepLoggingLogGroupResource(compartment string) error {
	loggingManagementClient := GetTestClients(&schema.ResourceData{}).loggingManagementClient()
	logGroupIds, err := getLogGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, logGroupId := range logGroupIds {
		if ok := SweeperDefaultResourceId[logGroupId]; !ok {
			deleteLogGroupRequest := oci_logging.DeleteLogGroupRequest{}

			deleteLogGroupRequest.LogGroupId = &logGroupId

			deleteLogGroupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "logging")
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
	ids := getResourceIdsToSweep(compartment, "LogGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loggingManagementClient := GetTestClients(&schema.ResourceData{}).loggingManagementClient()

	listLogGroupsRequest := oci_logging.ListLogGroupsRequest{}
	listLogGroupsRequest.CompartmentId = &compartmentId
	listLogGroupsResponse, err := loggingManagementClient.ListLogGroups(context.Background(), listLogGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LogGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, logGroup := range listLogGroupsResponse.Items {
		id := *logGroup.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "LogGroupId", id)
	}
	return resourceIds, nil
}
