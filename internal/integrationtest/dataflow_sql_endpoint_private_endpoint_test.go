// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataflowSqlEndpointPrivateEndpointRequiredOnlyResource = DataflowSqlEndpointResourcePrivateEndpointDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Required, acctest.Create, DataflowSqlEndpointPrivateEndpointRepresentation)

	DataflowSqlEndpointPrivateEndpointResourceConfig = DataflowSqlEndpointResourcePrivateEndpointDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointPrivateEndpointRepresentation)

	DataflowSqlEndpointPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"sql_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_sql_endpoint.test_sql_endpoint_private_endpoint.id}`},
	}

	DataflowSqlEndpointPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"sql_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataflow_sql_endpoint.test_sql_endpoint_private_endpoint.id}`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowSqlEndpointDataSourceFilterRepresentation},
	}

	DataflowSqlEndpointPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `test_sql_endpoint_private_endpoint_terraform`},
		"driver_shape":                  acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"executor_shape":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"max_executor_count":            acctest.Representation{RepType: acctest.Required, Create: `2`},
		"metastore_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.metastore_id}`},
		"min_executor_count":            acctest.Representation{RepType: acctest.Required, Create: `1`},
		"network_configuration":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowSqlEndpointPrivateEndpointNetworkConfigurationRepresentation},
		"sql_endpoint_version":          acctest.Representation{RepType: acctest.Required, Create: `3.2.1`},
		"warehouse_bucket_uri":          acctest.Representation{RepType: acctest.Required, Create: `${var.dataflow_warehouse_bucket_uri}`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"driver_shape_config":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowSqlEndpointDriverShapeConfigRepresentation},
		"executor_shape_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowSqlEndpointExecutorShapeConfigRepresentation},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"spark_advanced_configurations": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sparkAdvancedConfigurations": "sparkAdvancedConfigurations"}},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSqlEndpointDefinedTagsRepresentation},
	}

	DataflowSqlEndpointPrivateEndpointNetworkConfigurationRepresentation = map[string]interface{}{
		"network_type":     acctest.Representation{RepType: acctest.Required, Create: `VCN`},
		"host_name_prefix": acctest.Representation{RepType: acctest.Optional, Create: `hostNamePrefix`},
		"nsg_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		// with network_type VCN -> vcn_id, subnet_id are required
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	DataflowSqlEndpointResourcePrivateEndpointDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowSqlEndpointPrivateEndpointResource_basic(t *testing.T) {
	t.Skip("Skip this test as this is not supported yet. It will be taken care post GA")
	httpreplay.SetScenario("TestDataflowSqlEndpointPrivateEndpointResource_basic")

	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	metastoreId := utils.GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	warehouseBucketUri := utils.GetEnvSettingWithBlankDefault("dataflow_warehouse_bucket_uri")
	warehouseBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_warehouse_bucket_uri\" { default = \"%s\" }\n", warehouseBucketUri)

	resourceName := "oci_dataflow_sql_endpoint.test_sql_endpoint_private_endpoint"
	datasourceName := "data.oci_dataflow_sql_endpoints.test_sql_endpoints"
	singularDatasourceName := "data.oci_dataflow_sql_endpoint.test_sql_endpoint_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+metastoreIdVariableStr+warehouseBucketUriVariableStr+DataflowSqlEndpointResourcePrivateEndpointDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Optional, acctest.Create, DataflowSqlEndpointPrivateEndpointRepresentation), "dataflow", "sqlEndpointPrivateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckDataflowSqlEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + warehouseBucketUriVariableStr + DataflowSqlEndpointResourcePrivateEndpointDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Required, acctest.Create, DataflowSqlEndpointPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_private_endpoint_terraform"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),
				resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + warehouseBucketUriVariableStr + DataflowSqlEndpointResourcePrivateEndpointDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + warehouseBucketUriVariableStr + DataflowSqlEndpointResourcePrivateEndpointDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Optional, acctest.Create, DataflowSqlEndpointPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_private_endpoint_terraform"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.host_name_prefix", "hostNamePrefix"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.private_endpoint_ip"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),
				resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),

				// func(s *terraform.State) (err error) {
				// 	resId, err = acctest.FromInstanceState(s, resourceName, "id")
				// 	if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
				// 		if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
				// 			return errExport
				// 		}
				// 	}
				// 	return err
				// },
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + metastoreIdVariableStr + warehouseBucketUriVariableStr + DataflowSqlEndpointResourcePrivateEndpointDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataflowSqlEndpointPrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_private_endpoint_terraform"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.host_name_prefix", "hostNamePrefix"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.private_endpoint_ip"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),
				resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),

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
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + warehouseBucketUriVariableStr + DataflowSqlEndpointResourcePrivateEndpointDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_private_endpoint_terraform"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.host_name_prefix", "hostNamePrefix"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.private_endpoint_ip"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),
				resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_sql_endpoints", "test_sql_endpoints", acctest.Optional, acctest.Update, DataflowSqlEndpointPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + metastoreIdVariableStr + warehouseBucketUriVariableStr + DataflowSqlEndpointResourcePrivateEndpointDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "test_sql_endpoint_private_endpoint_terraform"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_endpoint_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "sql_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sql_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint_private_endpoint", acctest.Required, acctest.Create, DataflowSqlEndpointPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + metastoreIdVariableStr + warehouseBucketUriVariableStr + DataflowSqlEndpointPrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_sql_endpoint_private_endpoint_terraform"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jdbc_endpoint_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_accepted_request_token"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.0.host_name_prefix", "hostNamePrefix"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sql_endpoint_version", "3.2.1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state_message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "warehouse_bucket_uri", warehouseBucketUri),
			),
		},
		// verify resource import
		{
			Config:                  config + DataflowSqlEndpointPrivateEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataflowSqlEndpointPrivateEndpoint") {
		resource.AddTestSweepers("DataflowSqlEndpointPrivateEndpoint", &resource.Sweeper{
			Name:         "DataflowSqlEndpointPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["sqlEndpointPrivateEndpoint"],
			F:            sweepDataflowSqlEndpointResource,
		})
	}
}
