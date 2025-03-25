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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	bdsInstanceRepresentation = map[string]interface{}{
		"cluster_admin_password":  acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
		"cluster_public_key":      acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDpUa4zUZKyU3AkW9yoJTBDO550wpWZOXdHswfRq75gbJ2ZYlMtifvwiO3qUL/RIZSC6e1wA5OL2LQ97UaHrLLPXgjvKGVIDRHqPkzTOayjJ4ZA7NPNhcu6f/OxhKkCYF3TAQObhMJmUSMrWSUeufaRIujDz1HHqazxOgFk09fj4i2dcGnfPcm32t8a9MzlsHSmgexYCUwxGisuuWTsnMgxbqsj6DaY51l+SEPi5tf10iFmUWqziF0eKDDQ/jHkwLJ8wgBJef9FSOmwJReHcBY+NviwFTatGj7Cwtnks6CVomsFD+rAMJ9uzM8SCv5agYunx07hnEXbR9r/TXqgXGfN bdsclusterkey@oracleoci.com`},
		"cluster_version":         acctest.Representation{RepType: acctest.Required, Create: `ODH2_0`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName`},
		"is_high_availability":    acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_secure":               acctest.Representation{RepType: acctest.Required, Create: `true`},
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
		"ignore_existing_nodes_shape": acctest.Representation{RepType: acctest.Optional, Create: []string{`worker`, `master`, `utility`, `compute_only_worker`, `edge`}, Update: []string{`worker`, `master`, `utility`, `compute_only_worker`, `edge`}},
	}

	bdsInstanceNodeGenericShapeRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}

	bdsInstanceNodesWorkerRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`, Update: `VM.Standard.E5.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`, Update: `300`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `3`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}

	bdsInstanceEdgeNodeRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`, Update: `VM.Standard.E5.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}

	bdsInstanceResourceConfig = acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance",
		"test_bds_instance",
		acctest.Optional,
		acctest.Create,
		bdsInstanceRepresentation,
	)

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

	bdsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	kafkaBrokerNodeGenericShapeRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_kafka_nodes":    acctest.Representation{RepType: acctest.Required, Create: `1`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}

	createKafkaProfileRepresentation = acctest.RepresentationCopyWithNewProperties(kafkaBrokerNodeGenericShapeRepresentation,
		map[string]interface{}{
			"number_of_kafka_nodes": acctest.Representation{RepType: acctest.Required, Create: `3`},
		})

	addBrokerRepresentation = acctest.RepresentationCopyWithNewProperties(kafkaBrokerNodeGenericShapeRepresentation,
		map[string]interface{}{
			"number_of_kafka_nodes": acctest.Representation{RepType: acctest.Required, Create: `4`},
		})

	addBlockStorageRepresentation = acctest.RepresentationCopyWithNewProperties(kafkaBrokerNodeGenericShapeRepresentation,
		map[string]interface{}{
			"number_of_kafka_nodes":    acctest.Representation{RepType: acctest.Required, Create: `4`},
			"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `300`},
		})
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

	bootstrapScriptUrl := utils.GetEnvSettingWithBlankDefault("bootstrap_script_url")
	bootstrapScriptUrlVariableStr := fmt.Sprintf("variable \"bootstrap_script_url\" { default = \"%s\" }\n", bootstrapScriptUrl)

	bootstrapScriptUrlU := utils.GetEnvSettingWithBlankDefault("bootstrap_script_urlU")
	bootstrapScriptUrlUVariableStr := fmt.Sprintf("variable \"bootstrap_script_urlU\" { default = \"%s\" }\n", bootstrapScriptUrlU)

	resourceName := "oci_bds_bds_instance.test_bds_instance"
	datasourceName := "data.oci_bds_bds_instances.test_bds_instances"
	singularDatasourceName := "data.oci_bds_bds_instance.test_bds_instance"

	var resId string
	// Save TF content to Create resource with optional properties.
	// This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+bootstrapScriptUrlVariableStr+subnetIdVariableStr+
		bdsInstanceResourceConfig, "bds", "bdsInstanceOdh", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceOdhDestroy, []resource.TestStep{

		// Verify Create HBase Profile Cluster
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_bds_bds_instance",
					"test_bds_instance",
					acctest.Optional,
					acctest.Create,
					acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
						"cluster_profile": acctest.Representation{RepType: acctest.Optional, Create: `HBASE`, Update: `HBASE`},
					}),
				),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "bootstrap_script_url", bootstrapScriptUrl),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_profile", "HBASE"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
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

		// Add Edge node
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				compartmentIdUVariableStr + subnetIdVariableStr + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Update,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"cluster_profile": acctest.Representation{RepType: acctest.Optional, Create: `HBASE`, Update: `HBASE`},
					"edge_node":       acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceEdgeNodeRepresentation},
				}),
			),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},

		// Add Kafka to cluster
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				compartmentIdUVariableStr + subnetIdVariableStr + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Update,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `HBASE`, Update: `HBASE`},
					"edge_node":           acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceEdgeNodeRepresentation},
					"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
					"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: kafkaBrokerNodeGenericShapeRepresentation},
				}),
			),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},

		// Remove Kafka from cluster
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				compartmentIdUVariableStr + subnetIdVariableStr + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Update,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `HBASE`, Update: `HBASE`},
					"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
					"edge_node":           acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceEdgeNodeRepresentation},
				}),
			),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "false"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},

		// Delete before next Create
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr,
		},

		// Create Kafka profile cluster
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				compartmentIdUVariableStr + subnetIdVariableStr + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Create,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `KAFKA`, Update: `KAFKA`},
					"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
					"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: createKafkaProfileRepresentation},
				}),
			),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "true"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},

		// Add broker to Kafka profile cluster
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				compartmentIdUVariableStr + subnetIdVariableStr + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Update,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `KAFKA`, Update: `KAFKA`},
					"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
					"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: addBrokerRepresentation},
				}),
			),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "true"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},

		// Add block storage to kafka Broker
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				compartmentIdUVariableStr + subnetIdVariableStr + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Update,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `KAFKA`, Update: `KAFKA`},
					"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
					"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: addBlockStorageRepresentation},
				}),
			),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "true"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},

		// Delete before next Create
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr,
		},

		// Verify Create Hadoop Profile Cluster
		{
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_bds_bds_instance",
					"test_bds_instance",
					acctest.Optional,
					acctest.Create,
					acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
						"cluster_profile": acctest.Representation{RepType: acctest.Optional, Create: `HADOOP`, Update: `HADOOP`},
					}),
				),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "bootstrap_script_url", bootstrapScriptUrl),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
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
			Config: config + compartmentIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				compartmentIdUVariableStr + subnetIdVariableStr + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Update,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
					"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: kafkaBrokerNodeGenericShapeRepresentation},
					"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `HADOOP`, Update: `HADOOP`},
				}),
			),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + bootstrapScriptUrlVariableStr +
				bootstrapScriptUrlUVariableStr + BdsGenerateDataSourceConfig +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_bds_bds_instance",
					"test_bds_instance",
					acctest.Optional,
					acctest.Update,
					acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
						"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
						"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: kafkaBrokerNodeGenericShapeRepresentation},
						"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `HADOOP`, Update: `HADOOP`},
					}),
				),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "true"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.cluster_version", "ODH2_0"),
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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + bootstrapScriptUrlVariableStr +
				bootstrapScriptUrlUVariableStr + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Update,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
					"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: kafkaBrokerNodeGenericShapeRepresentation},
					"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `HADOOP`, Update: `HADOOP`},
				})) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_bds_bds_instance",
					"test_bds_instance",
					acctest.Required,
					acctest.Create,
					bdsInstanceSingularDataSourceRepresentation,
				),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_sql_details.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_version", "ODH2_0"),
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
			Config: config + acctest.GenerateResourceFromRepresentationMap(
				"oci_bds_bds_instance",
				"test_bds_instance",
				acctest.Optional,
				acctest.Update,
				acctest.RepresentationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
					"is_kafka_configured": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
					"kafka_broker_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: kafkaBrokerNodeGenericShapeRepresentation},
					"cluster_profile":     acctest.Representation{RepType: acctest.Optional, Create: `HADOOP`, Update: `HADOOP`},
				}),
			),
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
