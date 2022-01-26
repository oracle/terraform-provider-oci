// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v56/bds"
	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BdsInstanceOdhRequiredOnlyResource = BdsInstanceOdhResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation)

	BdsInstanceOdhResourceConfig = BdsInstanceOdhResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Update, bdsInstanceOdhRepresentation)

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

	bdsInstanceOdhRepresentation = map[string]interface{}{
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `V2VsY29tZTE=`},
		"cluster_public_key":     acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDpUa4zUZKyU3AkW9yoJTBDO550wpWZOXdHswfRq75gbJ2ZYlMtifvwiO3qUL/RIZSC6e1wA5OL2LQ97UaHrLLPXgjvKGVIDRHqPkzTOayjJ4ZA7NPNhcu6f/OxhKkCYF3TAQObhMJmUSMrWSUeufaRIujDz1HHqazxOgFk09fj4i2dcGnfPcm32t8a9MzlsHSmgexYCUwxGisuuWTsnMgxbqsj6DaY51l+SEPi5tf10iFmUWqziF0eKDDQ/jHkwLJ8wgBJef9FSOmwJReHcBY+NviwFTatGj7Cwtnks6CVomsFD+rAMJ9uzM8SCv5agYunx07hnEXbR9r/TXqgXGfN bdsclusterkey@oracleoci.com`},
		"cluster_version":        acctest.Representation{RepType: acctest.Required, Create: `ODH1`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_high_availability":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_secure":              acctest.Representation{RepType: acctest.Required, Create: `false`},
		"master_node":            acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesOdhMasterRepresentation},
		"util_node":              acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesOdhUtilRepresentation},
		"worker_node":            acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesOdhWorkerRepresentation},

		"is_cloud_sql_configured": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		//"cloud_sql_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: bdsInstanceNodesOdhCloudSqlRepresentation}, // capacity issue

		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"network_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: bdsInstanceOdhNetworkConfigRepresentation},
	}

	bdsInstanceNodesOdhCloudSqlRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.4`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `1000`},
	}

	bdsInstanceNodesOdhMasterRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.4`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	bdsInstanceNodesOdhUtilRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.4`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	bdsInstanceNodesOdhWorkerRepresentation = map[string]interface{}{
		"shape":                    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"block_volume_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `150`},
		"number_of_nodes":          acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `4`},
	}
	bdsInstanceOdhNetworkConfigRepresentation = map[string]interface{}{
		"cidr_block":              acctest.Representation{RepType: acctest.Optional, Create: `111.112.0.0/16`},
		"is_nat_gateway_required": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	BdsInstanceOdhResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create,
		acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{"cidr_block", "dns_label"},
			[]interface{}{acctest.Representation{RepType: acctest.Required, Create: `111.111.0.0/24`}, acctest.Representation{RepType: acctest.Required, Create: `bdssubnet`}},
			subnetRegionalRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{"cidr_block", "dns_label"},
			[]interface{}{acctest.Representation{RepType: acctest.Required, Create: `111.111.0.0/16`}, acctest.Representation{RepType: acctest.Required, Create: `bdsvcn`}},
			vcnRepresentation)) +
		DefinedTagsDependencies
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

	resourceName := "oci_bds_bds_instance.test_bds_instance"
	datasourceName := "data.oci_bds_bds_instances.test_bds_instances"
	singularDatasourceName := "data.oci_bds_bds_instance.test_bds_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsInstanceOdhResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceOdhRepresentation), "bds", "bdsInstanceOdh", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceOdhDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "false"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "5"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOdhResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceOdhRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "false"),
				resource.TestCheckResourceAttr(resourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.is_nat_gateway_required", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "number_of_nodes", "5"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "false"),
				resource.TestCheckResourceAttr(resourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.is_nat_gateway_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "5"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "number_of_nodes", "5"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Update, bdsInstanceOdhRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
				resource.TestCheckResourceAttr(resourceName, "cluster_version", "ODH1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(resourceName, "is_high_availability", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_secure", "false"),
				resource.TestCheckResourceAttr(resourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "network_config.0.is_nat_gateway_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "nodes.#", "6"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "number_of_nodes", "6"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				// Change shape not supported for ODH on GA
				//resource.TestCheckResourceAttr(resourceName, "master_node.0.shape", "VM.Standard2.8"),
				//resource.TestCheckResourceAttr(resourceName, "worker_node.0.shape", "VM.Standard2.4"),
				//resource.TestCheckResourceAttr(resourceName, "util_node.0.shape", "VM.Standard2.8"),

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
				compartmentIdVariableStr + BdsInstanceOdhResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Update, bdsInstanceOdhRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "bds_instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.cluster_version", "ODH1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.is_high_availability", "false"),
				resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.is_secure", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.number_of_nodes"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instances.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BdsInstanceOdhResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_sql_details.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_version", "ODH1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.bd_cell_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.bds_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.csql_cell_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.db_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.os_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"), //empty
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cloud_sql_configured"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_high_availability", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_secure", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.0.cidr_block", "111.112.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_config.0.is_nat_gateway_required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nodes.#", "6"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.fault_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.hostname"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.node_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_nodes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOdhResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cluster_admin_password",
				"cluster_public_key",
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
