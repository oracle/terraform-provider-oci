// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDatabaseDbNodeSingularDataSourceRepresentation = map[string]interface{}{
		"db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}`},
	}

	DatabaseDatabaseDbNodeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vm_cluster_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`}}

	DatabaseDbNodeRepresentation = map[string]interface{}{
		"db_node_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseDbNodeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		AvailabilityDomainConfig + DatabaseVmClusterResourceDependencies

	DatabaseDbNodeResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		AvailabilityDomainConfig + DatabaseVmClusterResourceDependencies + `

	  data "oci_database_db_nodes" "test_db_nodes" {
	     compartment_id = "${var.compartment_id}"
	     vm_cluster_id = "${oci_database_vm_cluster.test_vm_cluster.id}"
	  }
	  data "oci_database_db_node" "test_db_node" {
	     db_node_id = "${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}"
	  }`
)

// issue-routing-tag: database/ExaCC
func TestDatabaseDbNodeResource_basic_exacc(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeResource_basic_exacc")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_nodes.test_db_nodes"
	resourceName := "oci_database_db_node.test_db_node"
	singularDatasourceName := "data.oci_database_db_node.test_db_node"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDbNodeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node", "test_db_node", acctest.Optional, acctest.Create, DatabaseDbNodeRepresentation), "database", "dbNode", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node", "test_db_node", acctest.Required, acctest.Create, DatabaseDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node", "test_db_node", acctest.Optional, acctest.Create, DatabaseDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + DatabaseDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node", "test_db_node", acctest.Optional, acctest.Update, DatabaseDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_nodes", "test_db_nodes", acctest.Required, acctest.Create, DatabaseDatabaseDbNodeDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbNodeResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node", "test_db_node", acctest.Required, acctest.Create, DatabaseDatabaseDbNodeSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_nodes", "test_db_nodes", acctest.Required, acctest.Create, DatabaseDatabaseDbNodeDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbNodeResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_id"),
			),
		},
	})
}
