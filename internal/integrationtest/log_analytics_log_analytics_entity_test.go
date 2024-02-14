// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsLogAnalyticsEntityRequiredOnlyResource = LogAnalyticsLogAnalyticsEntityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsEntityRepresentation)

	LogAnalyticsLogAnalyticsEntityResourceConfig = LogAnalyticsLogAnalyticsEntityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsEntityRepresentation)

	LogAnalyticsLogAnalyticsLogAnalyticsEntitySingularDataSourceRepresentation = map[string]interface{}{
		"log_analytics_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id}`},
		"namespace":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsLogAnalyticsLogAnalyticsEntityDataSourceRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":                   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"cloud_resource_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"entity_type_name":            acctest.Representation{RepType: acctest.Optional, Create: []string{`Host (Linux)`}},
		"hostname":                    acctest.Representation{RepType: acctest.Optional, Create: `hostname`, Update: `hostname2`},
		"hostname_contains":           acctest.Representation{RepType: acctest.Optional, Create: `hostname`},
		"is_management_agent_id_null": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle_details_contains":  acctest.Representation{RepType: acctest.Optional, Create: `READY`},
		"metadata_equals":             acctest.Representation{RepType: acctest.Optional, Create: []string{`metadataName:metadataValue:metadataType`}, Update: []string{`metadataName:metadataValue1:metadataType`}},
		"name":                        acctest.Representation{RepType: acctest.Optional, Create: `TF_LA_ENTITY`},
		"name_contains":               acctest.Representation{RepType: acctest.Optional, Create: `TF_LA`},
		"source_id":                   acctest.Representation{RepType: acctest.Optional, Create: `source1`},
		"state":                       acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsLogAnalyticsEntityDataSourceFilterRepresentation}}
	LogAnalyticsLogAnalyticsEntityDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id}`}},
	}

	LogAnalyticsLogAnalyticsEntityRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_type_name":     acctest.Representation{RepType: acctest.Required, Create: `Host (Linux)`},
		"name":                 acctest.Representation{RepType: acctest.Required, Create: `TF_LA_ENTITY`},
		"namespace":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"cloud_resource_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname":             acctest.Representation{RepType: acctest.Optional, Create: `hostname`, Update: `hostname2`},
		"management_agent_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.managed_agent_id}`},
		"metadata":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: LogAnalyticsLogAnalyticsEntityMetadataRepresentation},
		"properties":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"properties": "properties"}, Update: map[string]string{"properties2": "properties2"}},
		"source_id":            acctest.Representation{RepType: acctest.Optional, Create: `source1`},
		"time_last_discovered": acctest.Representation{RepType: acctest.Optional, Create: `2023-09-10T08:20:50.052Z`, Update: `2023-09-10T08:20:51.052Z`},
		"timezone_region":      acctest.Representation{RepType: acctest.Optional, Create: `PST8PDT`, Update: `EST5EDT`},
	}
	LogAnalyticsLogAnalyticsEntityMetadataRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LogAnalyticsLogAnalyticsEntityMetadataItemsRepresentation},
	}
	LogAnalyticsLogAnalyticsEntityMetadataItemsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `metadataName`, Update: `metadataName`},
		"type":  acctest.Representation{RepType: acctest.Optional, Create: `metadataType`, Update: `metadataType`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `metadataValue`, Update: `metadataValue1`},
	}

	LogAnalyticsLogAnalyticsEntityResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsEntityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_log_analytics_log_analytics_entity.test_log_analytics_entity"
	datasourceName := "data.oci_log_analytics_log_analytics_entities.test_log_analytics_entities"
	singularDatasourceName := "data.oci_log_analytics_log_analytics_entity.test_log_analytics_entity"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsLogAnalyticsEntityResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsEntityRepresentation), "loganalytics", "logAnalyticsEntity", t)

	acctest.ResourceTest(t, testAccCheckLogAnalyticsLogAnalyticsEntityDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsEntityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsEntityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "entity_type_name", "Host (Linux)"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_LA_ENTITY"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsEntityResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
				LogAnalyticsLogAnalyticsEntityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsEntityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_resource_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "entity_type_internal_name"),
				resource.TestCheckResourceAttr(resourceName, "entity_type_name", "Host (Linux)"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "management_agent_id", managementAgentId),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.name", "metadataName"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.type", "metadataType"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.value", "metadataValue"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_LA_ENTITY"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_id", "source1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_last_discovered", "2023-09-10T08:20:50.052Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timezone_region", "PST8PDT"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + managementAgentIdVariableStr + LogAnalyticsLogAnalyticsEntityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsLogAnalyticsEntityRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_resource_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "entity_type_internal_name"),
				resource.TestCheckResourceAttr(resourceName, "entity_type_name", "Host (Linux)"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "management_agent_id", managementAgentId),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.#", "1"),
				// resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.name", "metadataName"),
				// resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.type", "metadataType"),
				// resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.value", "metadataValue"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_LA_ENTITY"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_id", "source1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_last_discovered"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timezone_region", "PST8PDT"),

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
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
				LogAnalyticsLogAnalyticsEntityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsEntityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_resource_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "entity_type_internal_name"),
				resource.TestCheckResourceAttr(resourceName, "entity_type_name", "Host (Linux)"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "hostname2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "management_agent_id", managementAgentId),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.name", "metadataName"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.type", "metadataType"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.items.0.value", "metadataValue1"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_LA_ENTITY"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_id", "source1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_last_discovered", "2023-09-10T08:20:51.052Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timezone_region", "EST5EDT"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entities", "test_log_analytics_entities", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsLogAnalyticsEntityDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr +
				LogAnalyticsLogAnalyticsEntityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsEntityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "entity_type_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "hostname", "hostname2"),
				resource.TestCheckResourceAttr(datasourceName, "hostname_contains", "hostname"),
				resource.TestCheckResourceAttr(datasourceName, "is_management_agent_id_null", "false"),
				resource.TestCheckResourceAttr(datasourceName, "lifecycle_details_contains", "READY"),
				resource.TestCheckResourceAttr(datasourceName, "metadata_equals.#", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsLogAnalyticsEntitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr +
				LogAnalyticsLogAnalyticsEntityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_analytics_entity_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "are_logs_collected"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entity_type_internal_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_type_name", "Host (Linux)"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hostname", "hostname2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_compartment_id"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.items.0.name", "metadataName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.items.0.type", "metadataType"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.items.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_last_discovered", "2023-09-10T08:20:51.052Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TF_LA_ENTITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timezone_region", "EST5EDT"),
			),
		},
		// verify resource import
		{
			Config:            config + LogAnalyticsLogAnalyticsEntityRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getLogAnalyticsEntityEndpointImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"namespace",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckLogAnalyticsLogAnalyticsEntityDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LogAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_log_analytics_entity" {
			noResourceFound = false
			request := oci_log_analytics.GetLogAnalyticsEntityRequest{}

			tmp := rs.Primary.ID
			request.LogAnalyticsEntityId = &tmp

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("LogAnalyticsLogAnalyticsEntity") {
		resource.AddTestSweepers("LogAnalyticsLogAnalyticsEntity", &resource.Sweeper{
			Name:         "LogAnalyticsLogAnalyticsEntity",
			Dependencies: acctest.DependencyGraph["logAnalyticsEntity"],
			F:            sweepLogAnalyticsLogAnalyticsEntityResource,
		})
	}
}

func sweepLogAnalyticsLogAnalyticsEntityResource(compartment string) error {
	logAnalyticsClient := acctest.GetTestClients(&schema.ResourceData{}).LogAnalyticsClient()
	logAnalyticsEntityIds, err := getLogAnalyticsLogAnalyticsEntityIds(compartment)
	if err != nil {
		return err
	}
	for _, logAnalyticsEntityId := range logAnalyticsEntityIds {
		if ok := acctest.SweeperDefaultResourceId[logAnalyticsEntityId]; !ok {
			deleteLogAnalyticsEntityRequest := oci_log_analytics.DeleteLogAnalyticsEntityRequest{}

			deleteLogAnalyticsEntityRequest.LogAnalyticsEntityId = &logAnalyticsEntityId

			deleteLogAnalyticsEntityRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")
			_, error := logAnalyticsClient.DeleteLogAnalyticsEntity(context.Background(), deleteLogAnalyticsEntityRequest)
			if error != nil {
				fmt.Printf("Error deleting LogAnalyticsEntity %s %s, It is possible that the resource is already deleted. Please verify manually \n", logAnalyticsEntityId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &logAnalyticsEntityId, LogAnalyticsLogAnalyticsEntitySweepWaitCondition, time.Duration(3*time.Minute),
				LogAnalyticsLogAnalyticsEntitySweepResponseFetchOperation, "log_analytics", true)
		}
	}
	return nil
}

func getLogAnalyticsLogAnalyticsEntityIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LogAnalyticsEntityId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	logAnalyticsClient := acctest.GetTestClients(&schema.ResourceData{}).LogAnalyticsClient()

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
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LogAnalyticsEntityId", id)
		}

	}
	return resourceIds, nil
}

func LogAnalyticsLogAnalyticsEntitySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if logAnalyticsEntityResponse, ok := response.Response.(oci_log_analytics.GetLogAnalyticsEntityResponse); ok {
		return logAnalyticsEntityResponse.LifecycleState != oci_log_analytics.EntityLifecycleStatesDeleted
	}
	return false
}

func LogAnalyticsLogAnalyticsEntitySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LogAnalyticsClient().GetLogAnalyticsEntity(context.Background(), oci_log_analytics.GetLogAnalyticsEntityRequest{
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
