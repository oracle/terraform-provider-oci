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
	oci_streaming "github.com/oracle/oci-go-sdk/v65/streaming"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	StreamingStreamPoolRequiredOnlyResource = StreamingStreamPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Required, acctest.Create, StreamingStreamPoolRepresentation)

	StreamingStreamPoolResourceConfig = StreamingStreamPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Optional, acctest.Update, StreamingStreamPoolRepresentation)

	StreamingStreamingStreamPoolSingularDataSourceRepresentation = map[string]interface{}{
		"stream_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_streaming_stream_pool.test_stream_pool.id}`},
	}

	StreamingStreamingStreamPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_streaming_stream_pool.test_stream_pool.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `MyStreamPool`, Update: `name2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StreamingStreamPoolDataSourceFilterRepresentation}}
	StreamingStreamPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_streaming_stream_pool.test_stream_pool.id}`}},
	}

	StreamingStreamPoolRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                      acctest.Representation{RepType: acctest.Required, Create: `MyStreamPool`, Update: `name2`},
		"custom_encryption_key":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: StreamingStreamPoolCustomEncryptionKeyRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"kafka_settings":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: StreamingStreamPoolKafkaSettingsRepresentation},
		"private_endpoint_settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: StreamingStreamPoolPrivateEndpointSettingsRepresentation},
	}
	StreamingStreamPoolCustomEncryptionKeyRepresentation = map[string]interface{}{
		"kms_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id_for_create}`},
	}
	StreamingStreamPoolKafkaSettingsRepresentation = map[string]interface{}{
		"auto_create_topics_enable": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"log_retention_hours":       acctest.Representation{RepType: acctest.Optional, Create: `25`, Update: `30`},
		"num_partitions":            acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	StreamingStreamPoolPrivateEndpointSettingsRepresentation = map[string]interface{}{
		"nsg_ids":             acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"private_endpoint_ip": acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"subnet_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	StreamingStreamPoolResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr
)

// issue-routing-tag: streaming/default
func TestStreamingStreamPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStreamingStreamPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_streaming_stream_pool.test_stream_pool"
	datasourceName := "data.oci_streaming_stream_pools.test_stream_pools"
	singularDatasourceName := "data.oci_streaming_stream_pool.test_stream_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StreamingStreamPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Optional, acctest.Create, StreamingStreamPoolRepresentation), "streaming", "streamPool", t)

	acctest.ResourceTest(t, testAccCheckStreamingStreamPoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + StreamingStreamPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Required, acctest.Create, StreamingStreamPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "MyStreamPool"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + StreamingStreamPoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + StreamingStreamPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Optional, acctest.Create, StreamingStreamPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_encryption_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.auto_create_topics_enable", "false"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.log_retention_hours", "25"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.num_partitions", "10"),
				resource.TestCheckResourceAttr(resourceName, "name", "MyStreamPool"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_settings.0.private_endpoint_ip", "10.0.0.5"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_settings.0.subnet_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + StreamingStreamPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StreamingStreamPoolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "custom_encryption_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.auto_create_topics_enable", "false"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.log_retention_hours", "25"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.num_partitions", "10"),
				resource.TestCheckResourceAttr(resourceName, "name", "MyStreamPool"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_settings.0.private_endpoint_ip", "10.0.0.5"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_settings.0.subnet_id"),
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
			Config: config + compartmentIdVariableStr + StreamingStreamPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Optional, acctest.Update, StreamingStreamPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_encryption_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.auto_create_topics_enable", "true"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.log_retention_hours", "30"),
				resource.TestCheckResourceAttr(resourceName, "kafka_settings.0.num_partitions", "11"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_settings.0.private_endpoint_ip", "10.0.0.5"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_settings.0.subnet_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_streaming_stream_pools", "test_stream_pools", acctest.Optional, acctest.Update, StreamingStreamingStreamPoolDataSourceRepresentation) +
				compartmentIdVariableStr + StreamingStreamPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Optional, acctest.Update, StreamingStreamPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "stream_pools.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "stream_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "stream_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "stream_pools.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "stream_pools.0.is_private"),
				resource.TestCheckResourceAttr(datasourceName, "stream_pools.0.name", "name2"),
				resource.TestCheckResourceAttrSet(datasourceName, "stream_pools.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "stream_pools.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Required, acctest.Create, StreamingStreamingStreamPoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + StreamingStreamPoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_encryption_key.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_encryption_key.0.key_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "endpoint_fqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_private"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kafka_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kafka_settings.0.auto_create_topics_enable", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kafka_settings.0.bootstrap_servers"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kafka_settings.0.log_retention_hours", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kafka_settings.0.num_partitions", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_settings.0.private_endpoint_ip", "10.0.0.5"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + StreamingStreamPoolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckStreamingStreamPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StreamAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_streaming_stream_pool" {
			noResourceFound = false
			request := oci_streaming.GetStreamPoolRequest{}

			tmp := rs.Primary.ID
			request.StreamPoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "streaming")

			response, err := client.GetStreamPool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_streaming.StreamPoolLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("StreamingStreamPool") {
		resource.AddTestSweepers("StreamingStreamPool", &resource.Sweeper{
			Name:         "StreamingStreamPool",
			Dependencies: acctest.DependencyGraph["streamPool"],
			F:            sweepStreamingStreamPoolResource,
		})
	}
}

func sweepStreamingStreamPoolResource(compartment string) error {
	streamAdminClient := acctest.GetTestClients(&schema.ResourceData{}).StreamAdminClient()
	streamPoolIds, err := getStreamingStreamPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, streamPoolId := range streamPoolIds {
		if ok := acctest.SweeperDefaultResourceId[streamPoolId]; !ok {
			deleteStreamPoolRequest := oci_streaming.DeleteStreamPoolRequest{}

			deleteStreamPoolRequest.StreamPoolId = &streamPoolId

			deleteStreamPoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "streaming")
			_, error := streamAdminClient.DeleteStreamPool(context.Background(), deleteStreamPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting StreamPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamPoolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &streamPoolId, StreamingStreamPoolSweepWaitCondition, time.Duration(3*time.Minute),
				StreamingStreamPoolSweepResponseFetchOperation, "streaming", true)
		}
	}
	return nil
}

func getStreamingStreamPoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StreamPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	streamAdminClient := acctest.GetTestClients(&schema.ResourceData{}).StreamAdminClient()

	listStreamPoolsRequest := oci_streaming.ListStreamPoolsRequest{}
	listStreamPoolsRequest.CompartmentId = &compartmentId
	listStreamPoolsRequest.LifecycleState = oci_streaming.StreamPoolSummaryLifecycleStateActive
	listStreamPoolsResponse, err := streamAdminClient.ListStreamPools(context.Background(), listStreamPoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting StreamPool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, streamPool := range listStreamPoolsResponse.Items {
		id := *streamPool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StreamPoolId", id)
	}
	return resourceIds, nil
}

func StreamingStreamPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamPoolResponse, ok := response.Response.(oci_streaming.GetStreamPoolResponse); ok {
		return streamPoolResponse.LifecycleState != oci_streaming.StreamPoolLifecycleStateDeleted
	}
	return false
}

func StreamingStreamPoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StreamAdminClient().GetStreamPool(context.Background(), oci_streaming.GetStreamPoolRequest{
		StreamPoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
