// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"

	//	"strconv"
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

	//	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ManagedKafkaKafkaClusterRequiredOnlyResource = ManagedKafkaKafkaClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterRepresentation)

	ManagedKafkaKafkaClusterResourceConfig = ManagedKafkaKafkaClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterRepresentation)

	ManagedKafkaKafkaClusterSingularDataSourceRepresentation = map[string]interface{}{
		"kafka_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_managed_kafka_kafka_cluster.test_kafka_cluster.id}`},
	}

	ManagedKafkaKafkaClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_managed_kafka_kafka_cluster.test_kafka_cluster.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ManagedKafkaKafkaClusterDataSourceFilterRepresentation}}
	ManagedKafkaKafkaClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_managed_kafka_kafka_cluster.test_kafka_cluster.id}`}},
	}

	ManagedKafkaKafkaClusterRepresentation = map[string]interface{}{
		"access_subnets":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ManagedKafkaKafkaClusterAccessSubnetsRepresentation},
		"broker_shape":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ManagedKafkaKafkaClusterBrokerShapeRepresentation},
		"cluster_config_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}`},
		"cluster_config_version":    acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"cluster_type":              acctest.Representation{RepType: acctest.Required, Create: `DEVELOPMENT`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"coordination_type":         acctest.Representation{RepType: acctest.Required, Create: `ZOOKEEPER`, Update: `ZOOKEEPER`},
		"kafka_version":             acctest.Representation{RepType: acctest.Required, Create: `3.7.0`},
		"client_certificate_bundle": acctest.Representation{RepType: acctest.Optional, Create: `${var.certificate}`, Update: `${var.certificate}`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	ManagedKafkaKafkaClusterAccessSubnetsRepresentation = map[string]interface{}{
		"subnets": acctest.Representation{RepType: acctest.Required, Create: []string{`subnet`}, Update: []string{`subnet1`}},
	}
	// replace subnet with subnet-id
	ManagedKafkaKafkaClusterBrokerShapeRepresentation = map[string]interface{}{
		"node_count":          acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `4`},
		"ocpu_count":          acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `50`},
	}

	ManagedKafkaKafkaClusterResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: managed_kafka/default
func TestManagedKafkaKafkaClusterResource_basic(t *testing.T) {

	//replace certificate with actual certificate data so that test's go smoothly
	certificate := "certificate"

	httpreplay.SetScenario("TestManagedKafkaKafkaClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	certificateId := certificate
	certificateIdUVariableStr := fmt.Sprintf("variable \"certificate\" { default = \"%s\" }\n", certificateId)

	resourceName := "oci_managed_kafka_kafka_cluster.test_kafka_cluster"
	datasourceName := "data.oci_managed_kafka_kafka_clusters.test_kafka_clusters"
	singularDatasourceName := "data.oci_managed_kafka_kafka_cluster.test_kafka_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagedKafkaKafkaClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Optional, acctest.Create, ManagedKafkaKafkaClusterRepresentation), "managedkafka", "kafkaCluster", t)

	acctest.ResourceTest(t, testAccCheckManagedKafkaKafkaClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagedKafkaKafkaClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_subnets.0.subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.node_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.ocpu_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_config_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_config_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "cluster_type", "DEVELOPMENT"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "coordination_type", "ZOOKEEPER"),
				resource.TestCheckResourceAttr(resourceName, "kafka_version", "3.7.0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//delete before next Create
		{
			Config: config + compartmentIdVariableStr + ManagedKafkaKafkaClusterResourceDependencies,
		},

		//verify Create with optionals
		{
			Config: config + certificateIdUVariableStr + compartmentIdVariableStr + ManagedKafkaKafkaClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Optional, acctest.Create, ManagedKafkaKafkaClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_subnets.0.subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.node_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "client_certificate_bundle"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_config_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_config_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "cluster_type", "DEVELOPMENT"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "coordination_type", "ZOOKEEPER"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "kafka_version", "3.7.0"),
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
			Config: config + certificateIdUVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedKafkaKafkaClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ManagedKafkaKafkaClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_subnets.0.subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.node_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "client_certificate_bundle"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_config_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_config_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "cluster_type", "DEVELOPMENT"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "coordination_type", "ZOOKEEPER"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "kafka_version", "3.7.0"),
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
			Config: config + certificateIdUVariableStr + compartmentIdVariableStr + ManagedKafkaKafkaClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_subnets.0.subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.node_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.ocpu_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "broker_shape.0.storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "client_certificate_bundle"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_config_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_config_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "cluster_type", "DEVELOPMENT"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "coordination_type", "ZOOKEEPER"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "kafka_version", "3.7.0"),
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
			Config: config + certificateIdUVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_managed_kafka_kafka_clusters", "test_kafka_clusters", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedKafkaKafkaClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Optional, acctest.Update, ManagedKafkaKafkaClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "kafka_cluster_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "kafka_cluster_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + certificateIdUVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_managed_kafka_kafka_cluster", "test_kafka_cluster", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedKafkaKafkaClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kafka_cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "access_subnets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "access_subnets.0.subnets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "broker_shape.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "broker_shape.0.node_count", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "broker_shape.0.ocpu_count", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "broker_shape.0.storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "client_certificate_bundle"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_config_version", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_type", "DEVELOPMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "coordination_type", "ZOOKEEPER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kafka_bootstrap_urls.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kafka_version", "3.7.0"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ManagedKafkaKafkaClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckManagedKafkaKafkaClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).KafkaClusterClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_managed_kafka_kafka_cluster" {
			noResourceFound = false
			request := oci_managed_kafka.GetKafkaClusterRequest{}

			tmp := rs.Primary.ID
			request.KafkaClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "managed_kafka")

			response, err := client.GetKafkaCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_managed_kafka.KafkaClusterLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ManagedKafkaKafkaCluster") {
		resource.AddTestSweepers("ManagedKafkaKafkaCluster", &resource.Sweeper{
			Name:         "ManagedKafkaKafkaCluster",
			Dependencies: acctest.DependencyGraph["kafkaCluster"],
			F:            sweepManagedKafkaKafkaClusterResource,
		})
	}
}

func sweepManagedKafkaKafkaClusterResource(compartment string) error {
	kafkaClusterClient := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()
	kafkaClusterIds, err := getManagedKafkaKafkaClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, kafkaClusterId := range kafkaClusterIds {
		if ok := acctest.SweeperDefaultResourceId[kafkaClusterId]; !ok {
			deleteKafkaClusterRequest := oci_managed_kafka.DeleteKafkaClusterRequest{}

			deleteKafkaClusterRequest.KafkaClusterId = &kafkaClusterId

			deleteKafkaClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "managed_kafka")
			_, error := kafkaClusterClient.DeleteKafkaCluster(context.Background(), deleteKafkaClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting KafkaCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", kafkaClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &kafkaClusterId, ManagedKafkaKafkaClusterSweepWaitCondition, time.Duration(3*time.Minute),
				ManagedKafkaKafkaClusterSweepResponseFetchOperation, "managed_kafka", true)
		}
	}
	return nil
}

func getManagedKafkaKafkaClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "KafkaClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	kafkaClusterClient := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()

	listKafkaClustersRequest := oci_managed_kafka.ListKafkaClustersRequest{}
	listKafkaClustersRequest.CompartmentId = &compartmentId
	listKafkaClustersRequest.LifecycleState = oci_managed_kafka.KafkaClusterLifecycleStateActive
	listKafkaClustersResponse, err := kafkaClusterClient.ListKafkaClusters(context.Background(), listKafkaClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting KafkaCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, kafkaCluster := range listKafkaClustersResponse.Items {
		id := *kafkaCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "KafkaClusterId", id)
	}
	return resourceIds, nil
}

func ManagedKafkaKafkaClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if kafkaClusterResponse, ok := response.Response.(oci_managed_kafka.GetKafkaClusterResponse); ok {
		return kafkaClusterResponse.LifecycleState != oci_managed_kafka.KafkaClusterLifecycleStateDeleted
	}
	return false
}

func ManagedKafkaKafkaClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.KafkaClusterClient().GetKafkaCluster(context.Background(), oci_managed_kafka.GetKafkaClusterRequest{
		KafkaClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
