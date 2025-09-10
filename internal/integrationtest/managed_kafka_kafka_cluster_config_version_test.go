// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ManagedKafkaKafkaClusterConfigVersionSingularDataSourceRepresentation = map[string]interface{}{
		"kafka_cluster_config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}`},
		"version_number":          acctest.Representation{RepType: acctest.Required, Create: `1`},
	}

	ManagedKafkaKafkaClusterConfigVersionDataSourceRepresentation = map[string]interface{}{
		"kafka_cluster_config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}`},
	}

	ManagedKafkaKafkaClusterConfigVersionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config", "test_kafka_cluster_config", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterConfigRepresentation)
)

// issue-routing-tag: managed_kafka/default
func TestManagedKafkaKafkaClusterConfigVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagedKafkaKafkaClusterConfigVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_managed_kafka_kafka_cluster_config_versions.test_kafka_cluster_config_versions"
	singularDatasourceName := "data.oci_managed_kafka_kafka_cluster_config_version.test_kafka_cluster_config_version"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config_versions", "test_kafka_cluster_config_versions", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterConfigVersionDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedKafkaKafkaClusterConfigVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "kafka_cluster_config_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "kafka_cluster_config_version_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_managed_kafka_kafka_cluster_config_version", "test_kafka_cluster_config_version", acctest.Required, acctest.Create, ManagedKafkaKafkaClusterConfigVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedKafkaKafkaClusterConfigVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kafka_cluster_config_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version_number"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version_number"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ManagedKafkaKafkaClusterConfigVersion") {
		resource.AddTestSweepers("ManagedKafkaKafkaClusterConfigVersion", &resource.Sweeper{
			Name:         "ManagedKafkaKafkaClusterConfigVersion",
			Dependencies: acctest.DependencyGraph["kafkaClusterConfigVersion"],
			F:            sweepManagedKafkaKafkaClusterConfigVersionResource,
		})
	}
}

func sweepManagedKafkaKafkaClusterConfigVersionResource(compartment string) error {
	kafkaClusterClient := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()
	kafkaClusterConfigVersionIds, err := getManagedKafkaKafkaClusterConfigVersionIds(compartment)
	if err != nil {
		return err
	}
	for _, kafkaClusterConfigVersionId := range kafkaClusterConfigVersionIds {
		if ok := acctest.SweeperDefaultResourceId[kafkaClusterConfigVersionId]; !ok {
			deleteKafkaClusterConfigVersionRequest := oci_managed_kafka.DeleteKafkaClusterConfigVersionRequest{}

			deleteKafkaClusterConfigVersionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "managed_kafka")
			_, error := kafkaClusterClient.DeleteKafkaClusterConfigVersion(context.Background(), deleteKafkaClusterConfigVersionRequest)
			if error != nil {
				fmt.Printf("Error deleting KafkaClusterConfigVersion %s %s, It is possible that the resource is already deleted. Please verify manually \n", kafkaClusterConfigVersionId, error)
				continue
			}
		}
	}
	return nil
}

func getManagedKafkaKafkaClusterConfigVersionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "KafkaClusterConfigVersionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	kafkaClusterClient := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()

	listKafkaClusterConfigVersionsRequest := oci_managed_kafka.ListKafkaClusterConfigVersionsRequest{}

	kafkaClusterConfigIds, error := getManagedKafkaKafkaClusterConfigIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting kafkaClusterConfigId required for KafkaClusterConfigVersion resource requests \n")
	}
	for _, kafkaClusterConfigId := range kafkaClusterConfigIds {
		listKafkaClusterConfigVersionsRequest.KafkaClusterConfigId = &kafkaClusterConfigId

		listKafkaClusterConfigVersionsResponse, err := kafkaClusterClient.ListKafkaClusterConfigVersions(context.Background(), listKafkaClusterConfigVersionsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting KafkaClusterConfigVersion list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, kafkaClusterConfigVersion := range listKafkaClusterConfigVersionsResponse.Items {
			id := *kafkaClusterConfigVersion.ConfigId
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "KafkaClusterConfigVersionId", id)
		}
	}
	return resourceIds, nil
}
