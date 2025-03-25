// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BdsInstanceOdhWithRegularComputeRequiredOnlyResource = BdsInstanceOdhResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhWithRegularComputeAndFlexMasterUtilRepresentation)

	BdsInstanceOdhWithRegularComputeWorkerResourceConfig = BdsInstanceOdhResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Update, bdsInstanceOdhWithRegularComputeAndFlexMasterUtilRepresentation)

	BdsInstanceOdhWithAddMasterUtilRequiredOnlyResource = BdsInstanceOdhResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Update, bdsInstanceOdhWithAddMasterUtilRepresentation)

	BdsInstanceOdhWithAddMasterUtilWorkerResourceConfig = BdsInstanceOdhResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Update, bdsInstanceOdhWithAddMasterUtilRepresentation)

	bdsInstanceOdhSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	bdsInstanceOdhDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceOdhDataSourceFilterRepresentation}}
	bdsInstanceOdhDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_instance.test_bds_instance.id}`}},
	}

	BdsBdsInstanceNetworkConfigRepresentation = map[string]interface{}{
		"cidr_block":              acctest.Representation{RepType: acctest.Optional, Create: `111.112.0.0/16`},
		"is_nat_gateway_required": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	BdsBdsInstanceNodesRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`, Update: `VM.Standard2.4`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `4`},
	}

	bdsInstanceOdhRepresentation = map[string]interface{}{
		"cluster_admin_password":      acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
		"cluster_public_key":          acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDpUa4zUZKyU3AkW9yoJTBDO550wpWZOXdHswfRq75gbJ2ZYlMtifvwiO3qUL/RIZSC6e1wA5OL2LQ97UaHrLLPXgjvKGVIDRHqPkzTOayjJ4ZA7NPNhcu6f/OxhKkCYF3TAQObhMJmUSMrWSUeufaRIujDz1HHqazxOgFk09fj4i2dcGnfPcm32t8a9MzlsHSmgexYCUwxGisuuWTsnMgxbqsj6DaY51l+SEPi5tf10iFmUWqziF0eKDDQ/jHkwLJ8wgBJef9FSOmwJReHcBY+NviwFTatGj7Cwtnks6CVomsFD+rAMJ9uzM8SCv5agYunx07hnEXbR9r/TXqgXGfN bdsclusterkey@oracleoci.com`},
		"cluster_version":             acctest.Representation{RepType: acctest.Required, Create: `ODH2_0`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_high_availability":        acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_secure":                   acctest.Representation{RepType: acctest.Required, Create: `true`},
		"bds_cluster_version_summary": acctest.RepresentationGroup{RepType: acctest.Optional, Group: BdsBdsInstanceBdsClusterVersionSummaryRepresentation},
		"cluster_profile":             acctest.Representation{RepType: acctest.Optional, Create: `HADOOP`},
		"kerberos_realm_name":         acctest.Representation{RepType: acctest.Optional, Create: `BDSCLOUDSERVICE.ORACLE.COM`},
		"master_node":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeFlexShapeRepresentation},
		"util_node":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeFlexShapeRepresentation},
		"worker_node":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesOdhWorkerRepresentation},
		"bootstrap_script_url":        acctest.Representation{RepType: acctest.Optional, Create: `${var.bootstrap_script_url}`, Update: `${var.bootstrap_script_urlU}`},
		"compute_only_worker_node":    acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeFlexShapeRepresentation},
		"edge_node":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeFlexShapeRepresentation},

		"is_cloud_sql_configured": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`, Update: `${var.kms_key_id_for_update}`},
		//"cloud_sql_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: bdsInstanceNodesOdhCloudSqlRepresentation}, // capacity issue

		//Uncomment this when running in home region (PHX)
		//	"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"network_config":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: bdsInstanceOdhNetworkConfigRepresentation},
		"ignore_existing_nodes_shape": acctest.Representation{RepType: acctest.Optional, Update: []string{`worker`}},
		//"os_patch_version": acctest.Representation{RepType: acctest.Optional, Update: `ol7.9-x86_64-1.28.0.619-0.0`}, // Test when patch is available
	}

	bdsInstanceStartClusterShapeConfigRepresentation = map[string]interface{}{
		"node_type_shape_configs": acctest.RepresentationGroup{RepType: acctest.Optional, Group: bdsInstanceNodeTypeShapeConfigsRepresentation},
	}

	bdsInstanceNodeTypeShapeConfigsRepresentation = map[string]interface{}{
		"node_type": acctest.Representation{RepType: acctest.Optional, Create: `WORKER`},
		"shape":     acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E5.Flex`},
	}

	BdsBdsInstanceBdsClusterVersionSummaryRepresentation = map[string]interface{}{
		"bds_version": acctest.Representation{RepType: acctest.Required, Create: `3.0.26`},
		"odh_version": acctest.Representation{RepType: acctest.Optional, Create: `2.0.7`},
	}

	bdsInstanceOdhWithFlexComputeAndRegularMasterUtilRepresentation = acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhRepresentation,
		map[string]interface{}{
			// Master & Util shape should be same
			"master_node":              acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesOdhMasterRepresentation},
			"util_node":                acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesOdhUtilRepresentation},
			"compute_only_worker_node": acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeFlexShapeRepresentation},
			"edge_node":                acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeFlexShapeRepresentation},
		})

	bdsInstanceOdhWithRegularComputeAndFlexMasterUtilRepresentation = acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhRepresentation,
		map[string]interface{}{
			"compute_only_worker_node": acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesOdhUtilRepresentation}, // Regular util shape representation usable for compute worker
		})

	bdsInstanceNodesOdhMasterRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.4`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `2`},
	}
	bdsInstanceNodesOdhUtilRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.4`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `2`},
	}
	bdsInstanceNodesOdhWorkerRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`, Update: `VM.Standard.E5.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `4`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesWorkerShapeConfigRepresentation},
	}
	bdsInstanceNodesWorkerShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `120`, Update: `120`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `8`, Update: `8`},
	}

	bdsInstanceNodeFlexShapeRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}

	bdsInstanceNodeFlex3ShapeRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard3.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}
	bdsInstanceNodesShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `32`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `3`},
	}

	bdsInstanceNodesDenseShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `128`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `8`},
		"nvmes":         acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	bdsInstanceOdhWithAddMasterUtilRepresentation = acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhRepresentation,
		map[string]interface{}{
			// Master & Util shape should be same
			"master_node": acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceOdhWithUpdateMasterUtilRepresentation},
			"util_node":   acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceOdhWithUpdateMasterUtilRepresentation},
		})
	bdsInstanceOdhWithUpdateMasterUtilRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Update: `3`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}
	bdsInstanceOdhNetworkConfigRepresentation = map[string]interface{}{
		"cidr_block":              acctest.Representation{RepType: acctest.Optional, Create: `111.112.0.0/16`},
		"is_nat_gateway_required": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
	}
	bdsInstanceKafkaBrokerNodeFlexShapeRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_kafka_nodes":    acctest.Representation{RepType: acctest.Required, Create: `3`},
		"shape_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigRepresentation},
	}

	BdsInstanceOdhResourceDependencies = KeyResourceDependencyConfig
)

// issue-routing-tag: bds/default
func TestResourceBdsOdhInstance(t *testing.T) {
	httpreplay.SetScenario("TestResourceBdsOdhInstance")
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

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+kmsKeyIdVariableStr+BdsInstanceOdhResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceOdhRepresentation), "bds", "bdsInstanceOdh", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceOdhDestroy, []resource.TestStep{
		// verify Create with required fields
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + subnetIdVariableStr + BdsInstanceOdhResourceDependencies + bootstrapScriptUrlVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOdhResourceDependencies + bootstrapScriptUrlVariableStr,
		},

		// verify Create, cluster will be force stopped after create
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + subnetIdVariableStr + BdsInstanceOdhResourceDependencies + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhRepresentation, map[string]interface{}{
						"is_force_stop_jobs": acctest.Representation{RepType: acctest.Required, Create: `true`},
						"state":              acctest.Representation{RepType: acctest.Required, Create: `INACTIVE`},
					})),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// start the cluster
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + subnetIdVariableStr + BdsInstanceOdhResourceDependencies + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhRepresentation, map[string]interface{}{
						"is_force_stop_jobs":          acctest.Representation{RepType: acctest.Required, Create: `true`},
						"state":                       acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`},
						"start_cluster_shape_configs": acctest.RepresentationGroup{RepType: acctest.Optional, Group: bdsInstanceStartClusterShapeConfigRepresentation},
					})),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOdhResourceDependencies + bootstrapScriptUrlVariableStr,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + subnetIdVariableStr + BdsInstanceOdhResourceDependencies + bootstrapScriptUrlVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceOdhRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "bootstrap_script_url", bootstrapScriptUrl),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.0.bds_version", "3.0.26"),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.0.odh_version", "2.0.7"),
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
				resource.TestCheckResourceAttr(resourceName, "kms_key_id", kmsKeyId),
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
				resource.TestCheckResourceAttr(resourceName, "number_of_nodes", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "util_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "master_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "compute_only_worker_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "edge_node.0.shape", "VM.Standard.E5.Flex"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step) and change shapes
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + kmsKeyIdVariableStr + subnetIdVariableStr + bootstrapScriptUrlVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhWithFlexComputeAndRegularMasterUtilRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.0.bds_version", "3.0.26"),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.0.odh_version", "2.0.7"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "kerberos_realm_name", "BDSCLOUDSERVICE.ORACLE.COM"),
				resource.TestCheckResourceAttr(resourceName, "kms_key_id", kmsKeyId),
				resource.TestCheckResourceAttr(resourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.is_nat_gateway_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "number_of_nodes", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "util_node.0.shape", "VM.Standard2.4"),
				resource.TestCheckResourceAttr(resourceName, "master_node.0.shape", "VM.Standard2.4"),
				resource.TestCheckResourceAttr(resourceName, "compute_only_worker_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "edge_node.0.shape", "VM.Standard.E5.Flex"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters, add a worker, update compute worker flex->regular, update util regular -> flex
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdUVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr + subnetIdVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Update, bdsInstanceOdhWithRegularComputeAndFlexMasterUtilRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.0.bds_version", "3.0.26"),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.0.odh_version", "2.0.7"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "kerberos_realm_name", "BDSCLOUDSERVICE.ORACLE.COM"),
				resource.TestCheckResourceAttr(resourceName, "kms_key_id", kmsKeyIdU),
				resource.TestCheckResourceAttr(resourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.is_nat_gateway_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "number_of_nodes", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "util_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "master_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "compute_only_worker_node.0.shape", "VM.Standard2.4"),
				resource.TestCheckResourceAttr(resourceName, "edge_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "worker_node.0.number_of_nodes", "4"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// Add Master and Utility Node
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdUVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr + subnetIdVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Update, bdsInstanceOdhWithAddMasterUtilRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.0.bds_version", "3.0.26"),
				resource.TestCheckResourceAttr(resourceName, "bds_cluster_version_summary.0.odh_version", "2.0.7"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "kerberos_realm_name", "BDSCLOUDSERVICE.ORACLE.COM"),
				resource.TestCheckResourceAttr(resourceName, "kms_key_id", kmsKeyIdU),
				resource.TestCheckResourceAttr(resourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.is_nat_gateway_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "14"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "number_of_nodes", "14"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "util_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "master_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "compute_only_worker_node.0.shape", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "edge_node.0.shape", "VM.Standard.E5.Flex"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instances", "test_bds_instances", acctest.Optional, acctest.Update, bdsInstanceOdhDataSourceRepresentation) +
				compartmentIdVariableStr + kmsKeyIdUVariableStr + subnetIdVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Update, bdsInstanceOdhWithAddMasterUtilRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "false"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.cluster_profile", "HADOOP"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.cluster_version", "ODH2_0"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.display_name", "displayName2"),
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhSingularDataSourceRepresentation) +
				compartmentIdVariableStr + kmsKeyIdVariableStr + subnetIdVariableStr + kmsKeyIdUVariableStr + bootstrapScriptUrlVariableStr + bootstrapScriptUrlUVariableStr + BdsInstanceOdhWithAddMasterUtilWorkerResourceConfig,
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
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"), //empty
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kms_key_id", kmsKeyIdU),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_kafka_configured", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_high_availability", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_secure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.0.is_nat_gateway_required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nodes.#", "14"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "nodes.0.shape", "VM.Standard.E5.Flex"),
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
			Config:            config + BdsInstanceOdhWithAddMasterUtilRequiredOnlyResource,
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
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckBdsBdsInstanceOdhDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_instance" {
			noResourceFound = false
			request := oci_bds.GetBdsInstanceRequest{}

			tmp := rs.Primary.ID
			request.BdsInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

			response, err := client.GetBdsInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bds.BdsInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BdsBdsInstanceOdh") {
		resource.AddTestSweepers("BdsBdsInstanceOdh", &resource.Sweeper{
			Name:         "BdsBdsInstanceOdh",
			Dependencies: acctest.DependencyGraph["bdsInstanceOdh"],
			F:            sweepBdsBdsInstanceOdhResource,
		})
	}
}

func sweepBdsBdsInstanceOdhResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceOdhIds, err := getBdsInstanceOdhIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceOdhId := range bdsInstanceOdhIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceOdhId]; !ok {
			deleteBdsInstanceRequest := oci_bds.DeleteBdsInstanceRequest{}

			deleteBdsInstanceRequest.BdsInstanceId = &bdsInstanceOdhId

			deleteBdsInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteBdsInstance(context.Background(), deleteBdsInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceOdhId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceOdhId, bdsInstanceOdhSweepWaitCondition, time.Duration(3*time.Minute),
				bdsInstanceOdhSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsInstanceOdhIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceOdhId")
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
	for _, bdsInstanceOdh := range listBdsInstancesResponse.Items {
		id := *bdsInstanceOdh.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceOdhId", id)
	}
	return resourceIds, nil
}

func bdsInstanceOdhSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceOdhResponse, ok := response.Response.(oci_bds.GetBdsInstanceResponse); ok {
		return bdsInstanceOdhResponse.LifecycleState != oci_bds.BdsInstanceLifecycleStateDeleted
	}
	return false
}

func bdsInstanceOdhSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetBdsInstance(context.Background(), oci_bds.GetBdsInstanceRequest{
		BdsInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
