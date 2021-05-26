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
	"github.com/oracle/oci-go-sdk/v41/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v41/loganalytics"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsEntityRequiredOnlyResource = LogAnalyticsEntityResourceDependencies +
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Required, Create, logAnalyticsEntityRepresentation)

	LogAnalyticsEntityResourceConfig = LogAnalyticsEntityResourceDependencies +
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Optional, Update, logAnalyticsEntityRepresentation)

	logAnalyticsEntitySingularDataSourceRepresentation = map[string]interface{}{
		"log_analytics_entity_id": Representation{repType: Required, create: `${oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id}`},
		"namespace":               Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	logAnalyticsEntityDataSourceRepresentation = map[string]interface{}{
		"compartment_id":              Representation{repType: Required, create: `${var.compartment_id}`},
		"namespace":                   Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"cloud_resource_id":           Representation{repType: Optional, create: `${var.compartment_id}`},
		"entity_type_name":            Representation{repType: Optional, create: []string{`Host (Linux)`}},
		"hostname":                    Representation{repType: Optional, create: `hostname`, update: `hostname2`},
		"hostname_contains":           Representation{repType: Optional, create: `hostname`},
		"is_management_agent_id_null": Representation{repType: Optional, create: `false`},
		"lifecycle_details_contains":  Representation{repType: Optional, create: `READY`},
		"name":                        Representation{repType: Optional, create: `TF_LA_ENTITY`},
		"name_contains":               Representation{repType: Optional, create: `TF_LA`},
		"source_id":                   Representation{repType: Optional, create: `source1`},
		"state":                       Representation{repType: Optional, create: `ACTIVE`},
		"filter":                      RepresentationGroup{Required, logAnalyticsEntityDataSourceFilterRepresentation}}
	logAnalyticsEntityDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id}`}},
	}

	logAnalyticsEntityRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"entity_type_name":    Representation{repType: Required, create: `Host (Linux)`},
		"name":                Representation{repType: Required, create: `TF_LA_ENTITY`},
		"namespace":           Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"cloud_resource_id":   Representation{repType: Optional, create: `${var.compartment_id}`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"hostname":            Representation{repType: Optional, create: `hostname`, update: `hostname2`},
		"management_agent_id": Representation{repType: Optional, create: `${var.managed_agent_id}`},
		"properties":          Representation{repType: Optional, create: map[string]string{"properties": "properties"}, update: map[string]string{"properties2": "properties2"}},
		"source_id":           Representation{repType: Optional, create: `source1`},
		"timezone_region":     Representation{repType: Optional, create: `PST8PDT`, update: `EST5EDT`},
	}

	LogAnalyticsEntityResourceDependencies = DefinedTagsDependencies +
		generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

func TestLogAnalyticsLogAnalyticsEntityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntityResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	managementAgentId := getEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_log_analytics_log_analytics_entity.test_log_analytics_entity"
	datasourceName := "data.oci_log_analytics_log_analytics_entities.test_log_analytics_entities"
	singularDatasourceName := "data.oci_log_analytics_log_analytics_entity.test_log_analytics_entity"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+LogAnalyticsEntityResourceDependencies+
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Optional, Create, logAnalyticsEntityRepresentation), "loganalytics", "logAnalyticsEntity", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLogAnalyticsLogAnalyticsEntityDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + LogAnalyticsEntityResourceDependencies +
					generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Required, Create, logAnalyticsEntityRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "entity_type_name", "Host (Linux)"),
					resource.TestCheckResourceAttr(resourceName, "name", "TF_LA_ENTITY"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + LogAnalyticsEntityResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
					LogAnalyticsEntityResourceDependencies +
					generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Optional, Create, logAnalyticsEntityRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cloud_resource_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "entity_type_internal_name"),
					resource.TestCheckResourceAttr(resourceName, "entity_type_name", "Host (Linux)"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "hostname"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "management_agent_id", managementAgentId),
					resource.TestCheckResourceAttr(resourceName, "name", "TF_LA_ENTITY"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_id", "source1"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "timezone_region", "PST8PDT"),

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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + managementAgentIdVariableStr + LogAnalyticsEntityResourceDependencies +
					generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Optional, Create,
						representationCopyWithNewProperties(logAnalyticsEntityRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cloud_resource_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "entity_type_internal_name"),
					resource.TestCheckResourceAttr(resourceName, "entity_type_name", "Host (Linux)"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "hostname"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "management_agent_id", managementAgentId),
					resource.TestCheckResourceAttr(resourceName, "name", "TF_LA_ENTITY"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_id", "source1"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "timezone_region", "PST8PDT"),

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
				Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
					LogAnalyticsEntityResourceDependencies +
					generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Optional, Update, logAnalyticsEntityRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cloud_resource_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "entity_type_internal_name"),
					resource.TestCheckResourceAttr(resourceName, "entity_type_name", "Host (Linux)"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "hostname2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "management_agent_id", managementAgentId),
					resource.TestCheckResourceAttr(resourceName, "name", "TF_LA_ENTITY"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_id", "source1"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "timezone_region", "EST5EDT"),

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
					generateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entities", "test_log_analytics_entities", Optional, Update, logAnalyticsEntityDataSourceRepresentation) +
					compartmentIdVariableStr + managementAgentIdVariableStr +
					LogAnalyticsEntityResourceDependencies +
					generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Optional, Update, logAnalyticsEntityRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_resource_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "entity_type_name.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "hostname", "hostname2"),
					resource.TestCheckResourceAttr(datasourceName, "hostname_contains", "hostname"),
					resource.TestCheckResourceAttr(datasourceName, "is_management_agent_id_null", "false"),
					resource.TestCheckResourceAttr(datasourceName, "lifecycle_details_contains", "READY"),
					resource.TestCheckResourceAttr(datasourceName, "name", "TF_LA_ENTITY"),
					resource.TestCheckResourceAttr(datasourceName, "name_contains", "TF_LA"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
					resource.TestCheckResourceAttrSet(datasourceName, "source_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "log_analytics_entity_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "log_analytics_entity_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Required, Create, logAnalyticsEntitySingularDataSourceRepresentation) +
					compartmentIdVariableStr + managementAgentIdVariableStr +
					LogAnalyticsEntityResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "log_analytics_entity_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "are_logs_collected"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "entity_type_internal_name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "entity_type_name", "Host (Linux)"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "hostname", "hostname2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_display_name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "TF_LA_ENTITY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "timezone_region", "EST5EDT"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsEntityResourceConfig,
			},
			// verify resource import
			{
				Config:            config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsEntityResourceConfig,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: getLogAnalyticsEntityEndpointImportId(resourceName),
				ImportStateVerifyIgnore: []string{
					"namespace",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckLogAnalyticsLogAnalyticsEntityDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).logAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_log_analytics_entity" {
			noResourceFound = false
			request := oci_log_analytics.GetLogAnalyticsEntityRequest{}

			tmp := rs.Primary.ID
			request.LogAnalyticsEntityId = &tmp

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "log_analytics")

			response, err := client.GetLogAnalyticsEntity(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_log_analytics.EntityLifecycleStatesDeleted): true,
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
	if !inSweeperExcludeList("LogAnalyticsLogAnalyticsEntity") {
		resource.AddTestSweepers("LogAnalyticsLogAnalyticsEntity", &resource.Sweeper{
			Name:         "LogAnalyticsLogAnalyticsEntity",
			Dependencies: DependencyGraph["logAnalyticsEntity"],
			F:            sweepLogAnalyticsLogAnalyticsEntityResource,
		})
	}
}

func sweepLogAnalyticsLogAnalyticsEntityResource(compartment string) error {
	logAnalyticsClient := GetTestClients(&schema.ResourceData{}).logAnalyticsClient()
	logAnalyticsEntityIds, err := getLogAnalyticsEntityIds(compartment)
	if err != nil {
		return err
	}
	for _, logAnalyticsEntityId := range logAnalyticsEntityIds {
		if ok := SweeperDefaultResourceId[logAnalyticsEntityId]; !ok {
			deleteLogAnalyticsEntityRequest := oci_log_analytics.DeleteLogAnalyticsEntityRequest{}

			deleteLogAnalyticsEntityRequest.LogAnalyticsEntityId = &logAnalyticsEntityId

			deleteLogAnalyticsEntityRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "log_analytics")
			_, error := logAnalyticsClient.DeleteLogAnalyticsEntity(context.Background(), deleteLogAnalyticsEntityRequest)
			if error != nil {
				fmt.Printf("Error deleting LogAnalyticsEntity %s %s, It is possible that the resource is already deleted. Please verify manually \n", logAnalyticsEntityId, error)
				continue
			}
			waitTillCondition(testAccProvider, &logAnalyticsEntityId, logAnalyticsEntitySweepWaitCondition, time.Duration(3*time.Minute),
				logAnalyticsEntitySweepResponseFetchOperation, "log_analytics", true)
		}
	}
	return nil
}

func getLogAnalyticsEntityIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "LogAnalyticsEntityId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	logAnalyticsClient := GetTestClients(&schema.ResourceData{}).logAnalyticsClient()

	listLogAnalyticsEntitiesRequest := oci_log_analytics.ListLogAnalyticsEntitiesRequest{}
	listLogAnalyticsEntitiesRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for LogAnalyticsEntity resource requests \n")
	}
	for _, namespace := range namespaces {
		listLogAnalyticsEntitiesRequest.NamespaceName = &namespace

		listLogAnalyticsEntitiesRequest.LifecycleState = oci_log_analytics.ListLogAnalyticsEntitiesLifecycleStateActive
		listLogAnalyticsEntitiesResponse, err := logAnalyticsClient.ListLogAnalyticsEntities(context.Background(), listLogAnalyticsEntitiesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting LogAnalyticsEntity list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, logAnalyticsEntity := range listLogAnalyticsEntitiesResponse.Items {
			id := *logAnalyticsEntity.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "LogAnalyticsEntityId", id)
		}

	}
	return resourceIds, nil
}

func logAnalyticsEntitySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if logAnalyticsEntityResponse, ok := response.Response.(oci_log_analytics.GetLogAnalyticsEntityResponse); ok {
		return logAnalyticsEntityResponse.LifecycleState != oci_log_analytics.EntityLifecycleStatesDeleted
	}
	return false
}

func logAnalyticsEntitySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.logAnalyticsClient().GetLogAnalyticsEntity(context.Background(), oci_log_analytics.GetLogAnalyticsEntityRequest{
		LogAnalyticsEntityId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getLogAnalyticsEntityEndpointImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("namespaces/" + rs.Primary.Attributes["namespace"] + "/logAnalyticsEntities/" + rs.Primary.Attributes["id"]), nil
	}
}
