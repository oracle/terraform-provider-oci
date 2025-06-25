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
	DataflowSqlEndpointVCNAccessRequiredOnlyResource = DataflowSqlEndpointResourceVCNDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Required, acctest.Create, DataflowSqlEndpointWithVCNRepresentation)

	DataflowSqlEndpointVCNAccessResourceConfig = DataflowSqlEndpointResourceVCNDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointWithVCNRepresentation)

	DataflowSqlEndpointVCNAccessSingularDataSourceRepresentation = map[string]interface{}{
		"sql_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_sql_endpoint.test_sql_endpoint.id}`},
	}

	DataflowSqlEndpointVCNAccessDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DataflowSqlEndpointVCNAccessDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_sql_endpoint.test_sql_endpoint.id}`}},
	}

	DataflowSqlEndpointWithVCNRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `test_sql_endpoint_terraform_vcn`, Update: `test_sql_endpoint_terraform_vcn_updated`},
		"driver_shape":                  acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"executor_shape":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"max_executor_count":            acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `2`},
		"metastore_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.metastore_id}`},
		"min_executor_count":            acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"network_configuration":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowSqlEndpointNetworkConfigurationRepresentationWithVCN},
		"sql_endpoint_version":          acctest.Representation{RepType: acctest.Required, Create: `3.2.1`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description updated`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSqlEndpointDefinedTagsRepresentation},
		"spark_advanced_configurations": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"testConfig": "testValue2"}, Update: map[string]string{"testConfig": "testValue"}},
	}

	DataflowSqlEndpointNetworkConfigurationRepresentationWithVCN = map[string]interface{}{
		"network_type":     acctest.Representation{RepType: acctest.Required, Create: `VCN`},
		"vcn_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.vcn_id}`},
		"subnet_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"host_name_prefix": acctest.Representation{RepType: acctest.Required, Create: `testHostNamePrefix`},
		"nsg_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${var.nsg_id}`}},
	}

	DataflowSqlEndpointResourceVCNDependencies = DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowSqlEndpointVCNAccess_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowSqlEndpointVCNAccess_basic")

	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	metastoreId := utils.GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	vcnId := utils.GetEnvSettingWithBlankDefault("vcn_id")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	nsgId := utils.GetEnvSettingWithBlankDefault("nsg_id")
	nsgIdVariableStr := fmt.Sprintf("variable \"nsg_id\" { default = \"%s\" }\n", nsgId)

	resourceName := "oci_dataflow_sql_endpoint.test_sql_endpoint"
	datasourceName := "data.oci_dataflow_sql_endpoints.test_sql_endpoints"
	singularDatasourceName := "data.oci_dataflow_sql_endpoint.test_sql_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+metastoreIdVariableStr+vcnIdVariableStr+subnetIdVariableStr+nsgIdVariableStr+DataflowSqlEndpointResourceVCNDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Create, DataflowSqlEndpointWithVCNRepresentation), "dataflow", "sqlEndpoint", t)

	acctest.ResourceTest(t, testAccCheckDataflowSqlEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + nsgIdVariableStr + DataflowSqlEndpointResourceVCNDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Required, acctest.Create, DataflowSqlEndpointWithVCNRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_terraform_vcn"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.host_name_prefix"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + nsgIdVariableStr + DataflowSqlEndpointResourceVCNDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + nsgIdVariableStr + DataflowSqlEndpointResourceVCNDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Create, DataflowSqlEndpointWithVCNRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_terraform_vcn"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.host_name_prefix"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					// 	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					// 		return errExport
					// 	}
					// }
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + metastoreIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + nsgIdVariableStr + DataflowSqlEndpointResourceVCNDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataflowSqlEndpointWithVCNRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_terraform_vcn"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.host_name_prefix"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters and switch the compartment back
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + nsgIdVariableStr + DataflowSqlEndpointResourceVCNDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataflowSqlEndpointWithVCNRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
					}),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description updated"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_terraform_vcn_updated"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.host_name_prefix"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_sql_endpoints", "test_sql_endpoints", acctest.Optional, acctest.Update, DataflowSqlEndpointVCNAccessDataSourceRepresentation) +
				compartmentIdVariableStr + metastoreIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + nsgIdVariableStr + DataflowSqlEndpointResourceVCNDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointWithVCNRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "sql_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sql_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointVCNAccessSingularDataSourceRepresentation) +
				compartmentIdVariableStr + metastoreIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + nsgIdVariableStr + DataflowSqlEndpointVCNAccessResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_sql_endpoint_terraform_vcn_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jdbc_endpoint_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_executor_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.0.network_type", "VCN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_configuration.0.host_name_prefix"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sql_endpoint_version", "3.2.1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataflowSqlEndpointVCNAccessRequiredOnlyResource,
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
	if !acctest.InSweeperExcludeList("DataflowSqlEndpointVCNAccessRules") {
		resource.AddTestSweepers("DataflowSqlEndpointVCNAccessRules", &resource.Sweeper{
			Name:         "DataflowSqlEndpointVCNAccessRules",
			Dependencies: acctest.DependencyGraph["sqlEndpointVCNAccessRules"],
			F:            sweepDataflowSqlEndpointResource,
		})
	}
}
