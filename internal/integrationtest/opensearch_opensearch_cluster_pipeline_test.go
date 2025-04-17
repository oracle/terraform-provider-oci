// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpensearchOpensearchClusterPipelineRequiredOnlyResource = OpensearchOpensearchClusterPipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipeline", "test_opensearch_cluster_pipeline", acctest.Required, acctest.Create, OpensearchOpensearchClusterPipelineRepresentation2)

	OpensearchOpensearchClusterPipelineResourceConfig = OpensearchOpensearchClusterPipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipeline", "test_opensearch_cluster_pipeline", acctest.Optional, acctest.Update, OpensearchOpensearchClusterPipelineRepresentation2)

	OpensearchOpensearchClusterPipelineSingularDataSourceRepresentation = map[string]interface{}{
		"opensearch_cluster_pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opensearch_opensearch_cluster_pipeline.test_opensearch_cluster_pipeline.id}`},
	}

	OpensearchOpensearchClusterPipelineDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `tf_provider_pipeline_updated`, Update: `tf_provider_pipeline_updated`},
		"id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_opensearch_opensearch_cluster_pipeline.test_opensearch_cluster_pipeline.id}`},
		"pipeline_component_id": acctest.Representation{RepType: acctest.Optional, Create: `{{ClusterOCID}}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: OpensearchOpensearchClusterPipelineDataSourceFilterRepresentation}}
	OpensearchOpensearchClusterPipelineDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opensearch_opensearch_cluster_pipeline.test_opensearch_cluster_pipeline.id}`}},
	}
	// TODO: Before running this test, make sure to create cluster first with username and password
	// Then add username and password to vault
	// replace vault secret for username and password, and clusterId
	// replace {{clusterOCID}} {{username-vaultsecret}} {{password-vaultsecret}}
	OpensearchOpensearchClusterPipelineRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_prepper_configuration_body": acctest.Representation{RepType: acctest.Required, Create: "source_coordination:\\n  store:\\n    oci-object-bucket:\\n      name: data-prepper-source-coordination-testing\\n      namespace: idv3bncjikjv"},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `tf_provider_pipeline_updated`, Update: `tf_provider_pipeline_updated`},
		"memory_gb":                       acctest.Representation{RepType: acctest.Required, Create: `8`, Update: `16`},
		"node_count":                      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"node_shape":                      acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.A1.Flex`},
		"ocpu_count":                      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"pipeline_configuration_body":     acctest.Representation{RepType: acctest.Required, Create: "version: 2\\npipeline_configurations:\\n  oci:\\n    secrets:\\n      opensearch-username:\\n        secret_id: {{username-vaultsecret}}\\n        refresh_interval: PT2H\\n      opensearch-password:\\n        secret_id: {{password-vaultsecret}}\\n        refresh_interval: PT2H\\nsimple-sample-pipeline:\\n  source:\\n    oci-object:\\n      acknowledgments: true\\n      codec:\\n        newline:\\n      compression: none\\n      scan:\\n        scheduling:\\n          interval: PT30S\\n        buckets:\\n          - bucket:\\n              namespace: idv3bncjikjv\\n              name: data_prepper_integration_test_object_storage_source_bucket_0\\n              region: us-ashburn-1\\n  sink:\\n    - opensearch:\\n        hosts: [ {{ClusterOCID}} ]\\n        username: $${{oci_secrets:opensearch-username}}\\n        password: $${{oci_secrets:opensearch-password}}\\n        insecure: false\\n        index: pipeline-stage-testing-index-1"},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"opc_dry_run":                     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"nsg_id":                          acctest.Representation{RepType: acctest.Optional, Create: `customerNsgId`},
		"reverse_connection_endpoints":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: OpensearchOpensearchClusterPipelineReverseConnectionEndpointsRepresentation},
		"subnet_compartment_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"subnet_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_compartment_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"vcn_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreOpenSearchPipelineSystemTagsChangesRep},
	}

	OpensearchOpensearchClusterPipelineRepresentation2 = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_prepper_configuration_body": acctest.Representation{RepType: acctest.Required, Create: "source_coordination:\\n  store:\\n    oci-object-bucket:\\n      name: data-prepper-source-coordination-testing\\n      namespace: idv3bncjikjv"},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `tf_provider_pipeline_updated`, Update: `tf_provider_pipeline_updated`},
		"memory_gb":                       acctest.Representation{RepType: acctest.Required, Create: `8`, Update: `16`},
		"node_count":                      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"ocpu_count":                      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"pipeline_configuration_body":     acctest.Representation{RepType: acctest.Required, Create: "version: 2\\npipeline_configurations:\\n  oci:\\n    secrets:\\n      opensearch-username:\\n        secret_id: {{username-vaultsecret}}\\n        refresh_interval: PT2H\\n      opensearch-password:\\n        secret_id: {{password-vaultsecret}}\\n        refresh_interval: PT2H\\nsimple-sample-pipeline:\\n  source:\\n    oci-object:\\n      acknowledgments: true\\n      codec:\\n        newline:\\n      compression: none\\n      scan:\\n        scheduling:\\n          interval: PT30S\\n        buckets:\\n          - bucket:\\n              namespace: idv3bncjikjv\\n              name: data_prepper_integration_test_object_storage_source_bucket_0\\n              region: us-ashburn-1\\n  sink:\\n    - opensearch:\\n        hosts: [ {{ClusterOCID}} ]\\n        username: $${{oci_secrets:opensearch-username}}\\n        password: $${{oci_secrets:opensearch-password}}\\n        insecure: false\\n        index: pipeline-stage-testing-index-1"},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"opc_dry_run":                     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"nsg_id":                          acctest.Representation{RepType: acctest.Optional, Create: `customerNsgId`},
		"subnet_compartment_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"subnet_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_compartment_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"vcn_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreOpenSearchPipelineSystemTagsChangesRep},
	}

	ignoreOpenSearchPipelineSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	OpensearchOpensearchClusterPipelineReverseConnectionEndpointsRepresentation = map[string]interface{}{
		"customer_fqdn": acctest.Representation{RepType: acctest.Required, Create: `gzltoensueoa.streaming.us-ashburn-1.oci.oraclecloud.com`, Update: `gzltoensueoa.streaming.us-ashburn-1.oci.oraclecloud.com`},
		"customer_ip":   acctest.Representation{RepType: acctest.Required, Create: `10.0.0.211`, Update: `10.0.0.211`},
	}

	OpensearchOpensearchClusterPipelineResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchClusterPipelineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchOpensearchClusterPipelineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_opensearch_opensearch_cluster_pipeline.test_opensearch_cluster_pipeline"
	datasourceName := "data.oci_opensearch_opensearch_cluster_pipelines.test_opensearch_cluster_pipelines"
	singularDatasourceName := "data.oci_opensearch_opensearch_cluster_pipeline.test_opensearch_cluster_pipeline"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpensearchOpensearchClusterPipelineResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipeline", "test_opensearch_cluster_pipeline", acctest.Optional, acctest.Create, OpensearchOpensearchClusterPipelineRepresentation), "opensearch", "opensearchClusterPipeline", t)

	acctest.ResourceTest(t, testAccCheckOpensearchOpensearchClusterPipelineDestroy, []resource.TestStep{
		// verify Create with Object Source
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipeline", "test_opensearch_cluster_pipeline", acctest.Required, acctest.Create, OpensearchOpensearchClusterPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_prepper_configuration_body", "source_coordination:\n  store:\n    oci-object-bucket:\n      name: data-prepper-source-coordination-testing\n      namespace: idv3bncjikjv"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tf_provider_pipeline_updated"),
				resource.TestCheckResourceAttr(resourceName, "memory_gb", "8"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard.A1.Flex"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "pipeline_configuration_body", "version: 2\npipeline_configurations:\n  oci:\n    secrets:\n      opensearch-username:\n        secret_id: {{username-vaultsecret}}\n        refresh_interval: PT2H\n      opensearch-password:\n        secret_id: {{password-vaultsecret}}\n        refresh_interval: PT2H\nsimple-sample-pipeline:\n  source:\n    oci-object:\n      acknowledgments: true\n      codec:\n        newline:\n      compression: none\n      scan:\n        scheduling:\n          interval: PT30S\n        buckets:\n          - bucket:\n              namespace: idv3bncjikjv\n              name: data_prepper_integration_test_object_storage_source_bucket_0\n              region: us-ashburn-1\n  sink:\n    - opensearch:\n        hosts: [ {{ClusterOCID}} ]\n        username: ${{oci_secrets:opensearch-username}}\n        password: ${{oci_secrets:opensearch-password}}\n        insecure: false\n        index: pipeline-stage-testing-index-1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterPipelineResourceDependencies,
		},
		// verify Create with optionals Using Object source
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipeline", "test_opensearch_cluster_pipeline", acctest.Optional, acctest.Create, OpensearchOpensearchClusterPipelineRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_prepper_configuration_body", "source_coordination:\n  store:\n    oci-object-bucket:\n      name: data-prepper-source-coordination-testing\n      namespace: idv3bncjikjv"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tf_provider_pipeline_updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "opc_dry_run", "false"),
				resource.TestCheckResourceAttr(resourceName, "memory_gb", "8"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "nsg_id"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "pipeline_configuration_body", "version: 2\npipeline_configurations:\n  oci:\n    secrets:\n      opensearch-username:\n        secret_id: {{username-vaultsecret}}\n        refresh_interval: PT2H\n      opensearch-password:\n        secret_id: {{password-vaultsecret}}\n        refresh_interval: PT2H\nsimple-sample-pipeline:\n  source:\n    oci-object:\n      acknowledgments: true\n      codec:\n        newline:\n      compression: none\n      scan:\n        scheduling:\n          interval: PT30S\n        buckets:\n          - bucket:\n              namespace: idv3bncjikjv\n              name: data_prepper_integration_test_object_storage_source_bucket_0\n              region: us-ashburn-1\n  sink:\n    - opensearch:\n        hosts: [ {{ClusterOCID}} ]\n        username: ${{oci_secrets:opensearch-username}}\n        password: ${{oci_secrets:opensearch-password}}\n        insecure: false\n        index: pipeline-stage-testing-index-1"),
				resource.TestCheckResourceAttrSet(resourceName, "pipeline_mode"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipeline", "test_opensearch_cluster_pipeline", acctest.Optional, acctest.Update, OpensearchOpensearchClusterPipelineRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_prepper_configuration_body", "source_coordination:\n  store:\n    oci-object-bucket:\n      name: data-prepper-source-coordination-testing\n      namespace: idv3bncjikjv"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tf_provider_pipeline_updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "opc_dry_run", "false"),
				resource.TestCheckResourceAttr(resourceName, "memory_gb", "16"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "nsg_id"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "pipeline_configuration_body", "version: 2\npipeline_configurations:\n  oci:\n    secrets:\n      opensearch-username:\n        secret_id: {{username-vaultsecret}}\n        refresh_interval: PT2H\n      opensearch-password:\n        secret_id: {{password-vaultsecret}}\n        refresh_interval: PT2H\nsimple-sample-pipeline:\n  source:\n    oci-object:\n      acknowledgments: true\n      codec:\n        newline:\n      compression: none\n      scan:\n        scheduling:\n          interval: PT30S\n        buckets:\n          - bucket:\n              namespace: idv3bncjikjv\n              name: data_prepper_integration_test_object_storage_source_bucket_0\n              region: us-ashburn-1\n  sink:\n    - opensearch:\n        hosts: [ {{ClusterOCID}} ]\n        username: ${{oci_secrets:opensearch-username}}\n        password: ${{oci_secrets:opensearch-password}}\n        insecure: false\n        index: pipeline-stage-testing-index-1"),
				resource.TestCheckResourceAttrSet(resourceName, "pipeline_mode"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipelines", "test_opensearch_cluster_pipelines", acctest.Optional, acctest.Update, OpensearchOpensearchClusterPipelineDataSourceRepresentation) +
				compartmentIdVariableStr + OpensearchOpensearchClusterPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipeline", "test_opensearch_cluster_pipeline", acctest.Optional, acctest.Update, OpensearchOpensearchClusterPipelineRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "tf_provider_pipeline_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_component_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "opensearch_cluster_pipeline_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "opensearch_cluster_pipeline_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opensearch_opensearch_cluster_pipeline", "test_opensearch_cluster_pipeline", acctest.Required, acctest.Create, OpensearchOpensearchClusterPipelineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpensearchOpensearchClusterPipelineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opensearch_cluster_pipeline_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_prepper_configuration_body", "source_coordination:\n  store:\n    oci-object-bucket:\n      name: data-prepper-source-coordination-testing\n      namespace: idv3bncjikjv"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "tf_provider_pipeline_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "memory_gb", "16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pipeline_configuration_body", "version: 2\npipeline_configurations:\n  oci:\n    secrets:\n      opensearch-username:\n        secret_id: {{username-vaultsecret}}\n        refresh_interval: PT2H\n      opensearch-password:\n        secret_id: {{password-vaultsecret}}\n        refresh_interval: PT2H\nsimple-sample-pipeline:\n  source:\n    oci-object:\n      acknowledgments: true\n      codec:\n        newline:\n      compression: none\n      scan:\n        scheduling:\n          interval: PT30S\n        buckets:\n          - bucket:\n              namespace: idv3bncjikjv\n              name: data_prepper_integration_test_object_storage_source_bucket_0\n              region: us-ashburn-1\n  sink:\n    - opensearch:\n        hosts: [ {{ClusterOCID}} ]\n        username: ${{oci_secrets:opensearch-username}}\n        password: ${{oci_secrets:opensearch-password}}\n        insecure: false\n        index: pipeline-stage-testing-index-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pipeline_mode"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + OpensearchOpensearchClusterPipelineRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"opc_dry_run",
				"nsg_id",
				"subnet_compartment_id",
				"subnet_id",
				"vcn_compartment_id",
				"vcn_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOpensearchOpensearchClusterPipelineDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OpensearchClusterPipelineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opensearch_opensearch_cluster_pipeline" {
			noResourceFound = false
			request := oci_opensearch.GetOpensearchClusterPipelineRequest{}

			tmp := rs.Primary.ID
			request.OpensearchClusterPipelineId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opensearch")

			response, err := client.GetOpensearchClusterPipeline(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opensearch.OpensearchClusterPipelineLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpensearchOpensearchClusterPipeline") {
		resource.AddTestSweepers("OpensearchOpensearchClusterPipeline", &resource.Sweeper{
			Name:         "OpensearchOpensearchClusterPipeline",
			Dependencies: acctest.DependencyGraph["opensearchClusterPipeline"],
			F:            sweepOpensearchOpensearchClusterPipelineResource,
		})
	}
}

func sweepOpensearchOpensearchClusterPipelineResource(compartment string) error {
	opensearchClusterPipelineClient := acctest.GetTestClients(&schema.ResourceData{}).OpensearchClusterPipelineClient()
	opensearchClusterPipelineIds, err := getOpensearchOpensearchClusterPipelineIds(compartment)
	if err != nil {
		return err
	}
	for _, opensearchClusterPipelineId := range opensearchClusterPipelineIds {
		if ok := acctest.SweeperDefaultResourceId[opensearchClusterPipelineId]; !ok {
			deleteOpensearchClusterPipelineRequest := oci_opensearch.DeleteOpensearchClusterPipelineRequest{}

			deleteOpensearchClusterPipelineRequest.OpensearchClusterPipelineId = &opensearchClusterPipelineId

			deleteOpensearchClusterPipelineRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opensearch")
			_, error := opensearchClusterPipelineClient.DeleteOpensearchClusterPipeline(context.Background(), deleteOpensearchClusterPipelineRequest)
			if error != nil {
				fmt.Printf("Error deleting OpensearchClusterPipeline %s %s, It is possible that the resource is already deleted. Please verify manually \n", opensearchClusterPipelineId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &opensearchClusterPipelineId, OpensearchOpensearchClusterPipelineSweepWaitCondition, time.Duration(3*time.Minute),
				OpensearchOpensearchClusterPipelineSweepResponseFetchOperation, "opensearch", true)
		}
	}
	return nil
}

func getOpensearchOpensearchClusterPipelineIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OpensearchClusterPipelineId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	opensearchClusterPipelineClient := acctest.GetTestClients(&schema.ResourceData{}).OpensearchClusterPipelineClient()

	listOpensearchClusterPipelinesRequest := oci_opensearch.ListOpensearchClusterPipelinesRequest{}
	listOpensearchClusterPipelinesRequest.CompartmentId = &compartmentId
	listOpensearchClusterPipelinesRequest.LifecycleState = oci_opensearch.OpensearchClusterPipelineLifecycleStateActive
	listOpensearchClusterPipelinesResponse, err := opensearchClusterPipelineClient.ListOpensearchClusterPipelines(context.Background(), listOpensearchClusterPipelinesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OpensearchClusterPipeline list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, opensearchClusterPipeline := range listOpensearchClusterPipelinesResponse.Items {
		id := *opensearchClusterPipeline.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OpensearchClusterPipelineId", id)
	}
	return resourceIds, nil
}

func OpensearchOpensearchClusterPipelineSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if opensearchClusterPipelineResponse, ok := response.Response.(oci_opensearch.GetOpensearchClusterPipelineResponse); ok {
		return opensearchClusterPipelineResponse.LifecycleState != oci_opensearch.OpensearchClusterPipelineLifecycleStateDeleted
	}
	return false
}

func OpensearchOpensearchClusterPipelineSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OpensearchClusterPipelineClient().GetOpensearchClusterPipeline(context.Background(), oci_opensearch.GetOpensearchClusterPipelineRequest{
		OpensearchClusterPipelineId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
