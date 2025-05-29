// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ManagedKafkaKafkaClusterConfigRequiredOnlyResource = ManagedKafkaKafkaClusterConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterConfigRepresentation)

	ManagedKafkaKafkaClusterConfigResourceConfig = ManagedKafkaKafkaClusterConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterConfigRepresentation)

	ManagedKafkaKafkaClusterConfigSingularDataSourceRepresentation = map[string]interface{}{
		"kafka_cluster_config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}`},
	}

	ManagedKafkaKafkaClusterConfigDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ManagedKafkaKafkaClusterConfigDataSourceFilterRepresentation}}
	ManagedKafkaKafkaClusterConfigDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}`}},
	}

	ManagedKafkaKafkaClusterConfigRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"latest_config":  acctest.RepresentationGroup{RepType: acctest.Required, Group: ManagedKafkaKafkaClusterConfigLatestConfigRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	ManagedKafkaKafkaClusterConfigLatestConfigRepresentation = map[string]interface{}{
		"properties": acctest.Representation{RepType: acctest.Required, Create: map[string]string{"properties": "properties"}, Update: map[string]string{"properties2": "properties2"}},
		// Self-referential block
		//"config_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}`},
		"version_number": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}

	ManagedKafkaKafkaClusterConfigResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_config", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterConfigRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: managed_kafka/default
func TestManagedKafkaKafkaClusterConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagedKafkaKafkaClusterConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config"
	datasourceName := "data.oci_managed_kafka_kafka_cluster_configs.test_kafka_cluster_configs"
	singularDatasourceName := "data.oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagedKafkaKafkaClusterConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Optional, acctest.Create, ManagedKafkaKafkaClusterConfigRepresentation), "managedkafka", "kafkaClusterConfig", t)

	acctest.ResourceTest(t, testAccCheckManagedKafkaKafkaClusterConfigDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagedKafkaKafkaClusterConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "latest_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.0.properties.%", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ManagedKafkaKafkaClusterConfigResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ManagedKafkaKafkaClusterConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Optional, acctest.Create, ManagedKafkaKafkaClusterConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "latest_config.0.config_id"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.0.properties.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "latest_config.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.0.version_number", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedKafkaKafkaClusterConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ManagedKafkaKafkaClusterConfigRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "latest_config.0.config_id"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.0.properties.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "latest_config.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.0.version_number", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + ManagedKafkaKafkaClusterConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "latest_config.0.config_id"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.0.properties.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "latest_config.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "latest_config.0.version_number", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_configs", "test_kafka_cluster_configs", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterConfigDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedKafkaKafkaClusterConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "kafka_cluster_config_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "kafka_cluster_config_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedKafkaKafkaClusterConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kafka_cluster_config_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "latest_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "latest_config.0.properties.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "latest_config.0.time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "latest_config.0.version_number", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ManagedKafkaKafkaClusterConfigRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckManagedKafkaKafkaClusterConfigDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).KafkaClusterClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_managed_kafka_kafka_cluster_config" {
			noResourceFound = false
			request := oci_managed_kafka.GetKafkaClusterConfigRequest{}

			tmp := rs.Primary.ID
			request.KafkaClusterConfigId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "managed_kafka")

			response, err := client.GetKafkaClusterConfig(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_managed_kafka.KafkaClusterConfigLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ManagedKafkaKafkaClusterConfig") {
		resource.AddTestSweepers("ManagedKafkaKafkaClusterConfig", &resource.Sweeper{
			Name:         "ManagedKafkaKafkaClusterConfig",
			Dependencies: acctest.DependencyGraph["kafkaClusterConfig"],
			F:            sweepManagedKafkaKafkaClusterConfigResource,
		})
	}
}

func sweepManagedKafkaKafkaClusterConfigResource(compartment string) error {
	kafkaClusterClient := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()
	kafkaClusterConfigIds, err := getManagedKafkaKafkaClusterConfigIds(compartment)
	if err != nil {
		return err
	}
	for _, kafkaClusterConfigId := range kafkaClusterConfigIds {
		if ok := acctest.SweeperDefaultResourceId[kafkaClusterConfigId]; !ok {
			deleteKafkaClusterConfigRequest := oci_managed_kafka.DeleteKafkaClusterConfigRequest{}

			deleteKafkaClusterConfigRequest.KafkaClusterConfigId = &kafkaClusterConfigId

			deleteKafkaClusterConfigRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "managed_kafka")
			_, error := kafkaClusterClient.DeleteKafkaClusterConfig(context.Background(), deleteKafkaClusterConfigRequest)
			if error != nil {
				fmt.Printf("Error deleting KafkaClusterConfig %s %s, It is possible that the resource is already deleted. Please verify manually \n", kafkaClusterConfigId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &kafkaClusterConfigId, ManagedKafkaKafkaClusterConfigSweepWaitCondition, time.Duration(3*time.Minute),
				ManagedKafkaKafkaClusterConfigSweepResponseFetchOperation, "managed_kafka", true)
		}
	}
	return nil
}

func getManagedKafkaKafkaClusterConfigIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "KafkaClusterConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	kafkaClusterClient := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()

	listKafkaClusterConfigsRequest := oci_managed_kafka.ListKafkaClusterConfigsRequest{}
	listKafkaClusterConfigsRequest.CompartmentId = &compartmentId
	listKafkaClusterConfigsRequest.LifecycleState = oci_managed_kafka.KafkaClusterConfigLifecycleStateActive
	listKafkaClusterConfigsResponse, err := kafkaClusterClient.ListKafkaClusterConfigs(context.Background(), listKafkaClusterConfigsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting KafkaClusterConfig list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, kafkaClusterConfig := range listKafkaClusterConfigsResponse.Items {
		id := *kafkaClusterConfig.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "KafkaClusterConfigId", id)
	}
	return resourceIds, nil
}

func ManagedKafkaKafkaClusterConfigSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if kafkaClusterConfigResponse, ok := response.Response.(oci_managed_kafka.GetKafkaClusterConfigResponse); ok {
		return kafkaClusterConfigResponse.LifecycleState != oci_managed_kafka.KafkaClusterConfigLifecycleStateDeleted
	}
	return false
}

func ManagedKafkaKafkaClusterConfigSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.KafkaClusterClient().GetKafkaClusterConfig(context.Background(), oci_managed_kafka.GetKafkaClusterConfigRequest{
		KafkaClusterConfigId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
