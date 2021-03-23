// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v37/common"
	oci_streaming "github.com/oracle/oci-go-sdk/v37/streaming"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	StreamPoolRequiredOnlyResource = StreamPoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Required, Create, streamPoolRepresentation)

	StreamPoolResourceConfig = StreamPoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Optional, Update, streamPoolRepresentation)

	streamPoolSingularDataSourceRepresentation = map[string]interface{}{
		"stream_pool_id": Representation{repType: Required, create: `${oci_streaming_stream_pool.test_stream_pool.id}`},
	}

	streamPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"id":             Representation{repType: Optional, create: `${oci_streaming_stream_pool.test_stream_pool.id}`},
		"name":           Representation{repType: Optional, create: `MyStreamPool`, update: `name2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, streamPoolDataSourceFilterRepresentation}}
	streamPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_streaming_stream_pool.test_stream_pool.id}`}},
	}

	streamPoolRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"name":                      Representation{repType: Required, create: `MyStreamPool`, update: `name2`},
		"custom_encryption_key":     RepresentationGroup{Optional, streamPoolCustomEncryptionKeyRepresentation},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"kafka_settings":            RepresentationGroup{Optional, streamPoolKafkaSettingsRepresentation},
		"private_endpoint_settings": RepresentationGroup{Optional, streamPoolPrivateEndpointSettingsRepresentation},
	}
	streamPoolCustomEncryptionKeyRepresentation = map[string]interface{}{
		"kms_key_id": Representation{repType: Optional, create: `${var.kms_key_id_for_create}`},
	}
	streamPoolKafkaSettingsRepresentation = map[string]interface{}{
		"auto_create_topics_enable": Representation{repType: Optional, create: `false`, update: `true`},
		"log_retention_hours":       Representation{repType: Optional, create: `25`, update: `30`},
		"num_partitions":            Representation{repType: Optional, create: `10`, update: `11`},
	}
	streamPoolPrivateEndpointSettingsRepresentation = map[string]interface{}{
		"nsg_ids":             Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"private_endpoint_ip": Representation{repType: Optional, create: `10.0.0.5`},
		"subnet_id":           Representation{repType: Optional, create: `${oci_core_subnet.test_subnet.id}`},
	}

	StreamPoolResourceDependencies = generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr
)

func TestStreamingStreamPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStreamingStreamPoolResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_streaming_stream_pool.test_stream_pool"
	datasourceName := "data.oci_streaming_stream_pools.test_stream_pools"
	singularDatasourceName := "data.oci_streaming_stream_pool.test_stream_pool"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+StreamPoolResourceDependencies+
		generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Optional, Create, streamPoolRepresentation), "streaming", "streamPool", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckStreamingStreamPoolDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + StreamPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Required, Create, streamPoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", "MyStreamPool"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + StreamPoolResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + StreamPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Optional, Create, streamPoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "custom_encryption_key.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key.0.kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + StreamPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Optional, Create,
						representationCopyWithNewProperties(streamPoolRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "custom_encryption_key.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key.0.kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + StreamPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Optional, Update, streamPoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "custom_encryption_key.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key.0.kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_streaming_stream_pools", "test_stream_pools", Optional, Update, streamPoolDataSourceRepresentation) +
					compartmentIdVariableStr + StreamPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Optional, Update, streamPoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "stream_pools.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "stream_pools.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "stream_pools.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Required, Create, streamPoolSingularDataSourceRepresentation) +
					compartmentIdVariableStr + StreamPoolResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_pool_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "custom_encryption_key.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_encryption_key.0.key_state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + StreamPoolResourceConfig,
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

func testAccCheckStreamingStreamPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).streamAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_streaming_stream_pool" {
			noResourceFound = false
			request := oci_streaming.GetStreamPoolRequest{}

			tmp := rs.Primary.ID
			request.StreamPoolId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "streaming")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("StreamingStreamPool") {
		resource.AddTestSweepers("StreamingStreamPool", &resource.Sweeper{
			Name:         "StreamingStreamPool",
			Dependencies: DependencyGraph["streamPool"],
			F:            sweepStreamingStreamPoolResource,
		})
	}
}

func sweepStreamingStreamPoolResource(compartment string) error {
	streamAdminClient := GetTestClients(&schema.ResourceData{}).streamAdminClient()
	streamPoolIds, err := getStreamPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, streamPoolId := range streamPoolIds {
		if ok := SweeperDefaultResourceId[streamPoolId]; !ok {
			deleteStreamPoolRequest := oci_streaming.DeleteStreamPoolRequest{}

			deleteStreamPoolRequest.StreamPoolId = &streamPoolId

			deleteStreamPoolRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "streaming")
			_, error := streamAdminClient.DeleteStreamPool(context.Background(), deleteStreamPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting StreamPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamPoolId, error)
				continue
			}
			waitTillCondition(testAccProvider, &streamPoolId, streamPoolSweepWaitCondition, time.Duration(3*time.Minute),
				streamPoolSweepResponseFetchOperation, "streaming", true)
		}
	}
	return nil
}

func getStreamPoolIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "StreamPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	streamAdminClient := GetTestClients(&schema.ResourceData{}).streamAdminClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "StreamPoolId", id)
	}
	return resourceIds, nil
}

func streamPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamPoolResponse, ok := response.Response.(oci_streaming.GetStreamPoolResponse); ok {
		return streamPoolResponse.LifecycleState != oci_streaming.StreamPoolLifecycleStateDeleted
	}
	return false
}

func streamPoolSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.streamAdminClient().GetStreamPool(context.Background(), oci_streaming.GetStreamPoolRequest{
		StreamPoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
