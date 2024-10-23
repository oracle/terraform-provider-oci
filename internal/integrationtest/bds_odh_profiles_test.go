// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	bdsInstanceRepresentation = map[string]interface{}{
		"cluster_admin_password":  acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
		"cluster_public_key":      acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDpUa4zUZKyU3AkW9yoJTBDO550wpWZOXdHswfRq75gbJ2ZYlMtifvwiO3qUL/RIZSC6e1wA5OL2LQ97UaHrLLPXgjvKGVIDRHqPkzTOayjJ4ZA7NPNhcu6f/OxhKkCYF3TAQObhMJmUSMrWSUeufaRIujDz1HHqazxOgFk09fj4i2dcGnfPcm32t8a9MzlsHSmgexYCUwxGisuuWTsnMgxbqsj6DaY51l+SEPi5tf10iFmUWqziF0eKDDQ/jHkwLJ8wgBJef9FSOmwJReHcBY+NviwFTatGj7Cwtnks6CVomsFD+rAMJ9uzM8SCv5agYunx07hnEXbR9r/TXqgXGfN bdsclusterkey@oracleoci.com`},
		"cluster_version":         acctest.Representation{RepType: acctest.Required, Create: `ODH1`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName`},
		"is_high_availability":    acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_secure":               acctest.Representation{RepType: acctest.Required, Create: `true`},
		"cluster_profile":         acctest.Representation{RepType: acctest.Optional, Create: `HADOOP`},
		"kerberos_realm_name":     acctest.Representation{RepType: acctest.Optional, Create: `BDSCLOUDSERVICE.ORACLE.COM`},
		"master_node":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeGenericShapeRepresentation},
		"util_node":               acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeGenericShapeRepresentation},
		"worker_node":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesWorkerRepresentation},
		"bootstrap_script_url":    acctest.Representation{RepType: acctest.Optional, Create: `${var.bootstrap_script_url}`, Update: `${var.bootstrap_script_urlU}`},
		"is_cloud_sql_configured": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		// Uncomment below when running in home region (PHX)
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"network_config":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: bdsInstanceOdhNetworkConfigRepresentation},
		"ignore_existing_nodes_shape": acctest.Representation{RepType: acctest.Optional, Create: []string{`worker`, `master`, `utility`, `compute_only_worker`, `edge`, `kafka_broker`}, Update: []string{`worker`, `master`, `utility`, `compute_only_worker`, `edge`, `kafka_broker`}},
	}

	bdsInstanceNodeGenericShapeRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.Generic`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}

	bdsInstanceNodesWorkerRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.Generic`, Update: `VM.Standard.Generic`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `3`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesWorkerShapeConfigRepresentation},
	}

	bdsInstanceResourceConfig = acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance",
		"test_bds_instance",
		acctest.Optional,
		acctest.Create,
		bdsInstanceRepresentation)

	BdsInstanceWithAddKafkaResourceConfig = acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance",
		"test_bds_instance",
		acctest.Optional,
		acctest.Update,
		bdsInstanceWithAddKafkaRepresentation)

	BdsGenerateDataSourceConfig = acctest.GenerateDataSourceFromRepresentationMap(
		"oci_bds_bds_instances",
		"test_bds_instances",
		acctest.Optional,
		acctest.Update,
		bdsInstanceDataSourceRepresentation)

	bdsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceOdhDataSourceFilterRepresentation},
	}

	BdsSingularDataSourceConfig = acctest.GenerateDataSourceFromRepresentationMap(
		"oci_bds_bds_instance",
		"test_bds_instance",
		acctest.Required,
		acctest.Create,
		bdsInstanceSingularDataSourceRepresentation)

	bdsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	bdsInstanceWithAddKafkaRepresentation = acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation,
		map[string]interface{}{
			"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
			"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `HADOOP`, Update: `HADOOP`},
			"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName`},
			"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: kafkaBrokerNodeGenericShapeRepresentation},
		})

	kafkaBrokerNodeGenericShapeRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.Generic`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_kafka_nodes":    acctest.Representation{RepType: acctest.Required, Create: `3`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}
)

// issue-routing-tag: bds/default
func TestBdsOdhProfile(t *testing.T) {
	httpreplay.SetScenario("TestBdsOdhProfile")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyIdU := utils.GetEnvSettingWithBlankDefault("kms_key_id_for_update")
	kmsKeyIdUVariableStr := fmt.Sprintf("variable \"kms_key_id_for_update\" { default = \"%s\" }\n", kmsKeyIdU)

	bootstrapScriptUrl := utils.GetEnvSettingWithBlankDefault("bootstrap_script_url")
	bootstrapScriptUrlVariableStr := fmt.Sprintf("variable \"bootstrap_script_url\" { default = \"%s\" }\n", bootstrapScriptUrl)

	bootstrapScriptUrlU := utils.GetEnvSettingWithBlankDefault("bootstrap_script_urlU")
	bootstrapScriptUrlUVariableStr := fmt.Sprintf("variable \"bootstrap_script_urlU\" { default = \"%s\" }\n", bootstrapScriptUrlU)

	resourceName := "oci_bds_bds_instance.test_bds_instance"
	datasourceName := "data.oci_bds_bds_instances.test_bds_instances"
	singularDatasourceName := "data.oci_bds_bds_instance.test_bds_instance"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+kmsKeyIdVariableStr+bdsInstanceResourceConfig,
		"bds", "bdsInstanceOdh", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceOdhDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + subnetIdVariableStr +
				bootstrapScriptUrlVariableStr + bdsInstanceResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "bootstrap_script_url", bootstrapScriptUrl),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "kerberos_realm_name", "BDSCLOUDSERVICE.ORACLE.COM"),
				resource.TestCheckResourceAttr(resourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.is_nat_gateway_required", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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

		// Add Kafka to cluster
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + kmsKeyIdVariableStr +
				subnetIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				kmsKeyIdUVariableStr + BdsInstanceWithAddKafkaResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH1"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},

		// verify datasource
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdUVariableStr + subnetIdVariableStr +
				bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr + BdsInstanceWithAddKafkaResourceConfig +
				BdsGenerateDataSourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "true"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.cluster_version", "ODH1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.is_high_availability", "true"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.is_secure", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.number_of_nodes"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.number_of_nodes_requiring_maintenance_reboot"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.time_created"),
			),
		},

		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + subnetIdVariableStr +
				kmsKeyIdUVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				BdsInstanceWithAddKafkaResourceConfig + BdsSingularDataSourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_sql_details.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_version", "ODH1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.bd_cell_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.bds_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.csql_cell_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.db_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.os_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.0.is_nat_gateway_required", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.hostname"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.is_reboot_required"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.local_disks_total_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.memory_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nodes.0.node_type", "MASTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.nvmes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.ocpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_nodes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_nodes_requiring_maintenance_reboot"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_nodes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		// verify resource import
		{
			Config:            config + BdsInstanceWithAddKafkaResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cluster_admin_password",
				"cluster_public_key",
				"kerberos_realm_name",
				// shape_config should only be passed to backend API for flex shape. From get response, we can't know if
				// shape is flex or not. So we don't save it, unless user specified, to avoid passing wrong values during update.
				"master_node.0.shape_config",
				"util_node.0.shape_config",
				"worker_node.0.shape_config",
				"compute_only_worker_node.0.shape_config",
				"edge_node.0.shape_config",
				"kafka_broker_node.0.shape_config",
				"ignore_existing_nodes_shape.#",
				"ignore_existing_nodes_shape.0",
				"ignore_existing_nodes_shape.1",
				"ignore_existing_nodes_shape.2",
				"ignore_existing_nodes_shape.3",
				"ignore_existing_nodes_shape.4",
				"ignore_existing_nodes_shape.5",
				"kafka_broker_node.#",
				"kafka_broker_node.0.%",
				"kafka_broker_node.0.block_volume_size_in_gbs",
				"kafka_broker_node.0.number_of_kafka_nodes",
				"kafka_broker_node.0.shape",
				"kafka_broker_node.0.subnet_id",
			},
			ResourceName: resourceName,
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("BdsBdsInstance") {
		resource.AddTestSweepers("BdsBdsInstance", &resource.Sweeper{
			Name:         "BdsBdsInstance",
			Dependencies: acctest.DependencyGraph["bdsInstance"],
			F:            sweepBdsBdsInstanceResource,
		})
	}
}

/*
func sweepBdsBdsInstanceResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceIds, err := getBdsInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceId]; !ok {
			deleteBdsInstanceRequest := oci_bds.DeleteBdsInstanceRequest{}

			deleteBdsInstanceRequest.BdsInstanceId = &bdsInstanceId

			deleteBdsInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteBdsInstance(context.Background(), deleteBdsInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceId, bdsInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				bdsInstanceSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listBdsInstancesRequest := oci_bds.ListBdsInstancesRequest{}
	listBdsInstancesRequest.CompartmentId = &compartmentId
	listBdsInstancesRequest.LifecycleState = oci_bds.BdsInstanceLifecycleStateActive
	listBdsInstancesResponse, err := bdsClient.ListBdsInstances(context.Background(), listBdsInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BdsInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, bdsInstance := range listBdsInstancesResponse.Items {
		id := *bdsInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceId", id)
	}
	return resourceIds, nil
}

func bdsInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceResponse, ok := response.Response.(oci_bds.GetBdsInstanceResponse); ok {
		return bdsInstanceResponse.LifecycleState != oci_bds.BdsInstanceLifecycleStateDeleted
	}
	return false
}

func bdsInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetBdsInstance(context.Background(), oci_bds.GetBdsInstanceRequest{
		BdsInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
*/
