// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v42/bds"
	"github.com/oracle/oci-go-sdk/v42/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BdsInstanceRequiredOnlyResource = BdsInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Required, Create, bdsInstanceRepresentation)

	BdsInstanceResourceConfig = BdsInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Optional, Update, bdsInstanceRepresentation)

	bdsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": Representation{repType: Required, create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	bdsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, bdsInstanceDataSourceFilterRepresentation}}
	bdsInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_bds_bds_instance.test_bds_instance.id}`}},
	}

	bdsInstanceRepresentation = map[string]interface{}{
		"cluster_admin_password": Representation{repType: Required, create: `V2VsY29tZTE=`},
		"cluster_public_key":     Representation{repType: Required, create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDpUa4zUZKyU3AkW9yoJTBDO550wpWZOXdHswfRq75gbJ2ZYlMtifvwiO3qUL/RIZSC6e1wA5OL2LQ97UaHrLLPXgjvKGVIDRHqPkzTOayjJ4ZA7NPNhcu6f/OxhKkCYF3TAQObhMJmUSMrWSUeufaRIujDz1HHqazxOgFk09fj4i2dcGnfPcm32t8a9MzlsHSmgexYCUwxGisuuWTsnMgxbqsj6DaY51l+SEPi5tf10iFmUWqziF0eKDDQ/jHkwLJ8wgBJef9FSOmwJReHcBY+NviwFTatGj7Cwtnks6CVomsFD+rAMJ9uzM8SCv5agYunx07hnEXbR9r/TXqgXGfN bdsclusterkey@oracleoci.com`},
		"cluster_version":        Representation{repType: Required, create: `CDH6`},
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":           Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"is_high_availability":   Representation{repType: Required, create: `false`},
		"is_secure":              Representation{repType: Required, create: `false`},
		"master_node":            RepresentationGroup{Required, bdsInstanceNodesMasterRepresentation},
		"util_node":              RepresentationGroup{Required, bdsInstanceNodesUtilRepresentation},
		"worker_node":            RepresentationGroup{Required, bdsInstanceNodesWorkerRepresentation},

		"is_cloud_sql_configured": Representation{repType: Optional, create: `false`},
		//"cloud_sql_details":       RepresentationGroup{Optional, bdsInstanceNodesCloudSqlRepresentation}, // capacity issue

		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"network_config": RepresentationGroup{Optional, bdsInstanceNetworkConfigRepresentation},
	}

	bdsInstanceNodesCloudSqlRepresentation = map[string]interface{}{
		"shape":                    Representation{repType: Required, create: `VM.Standard2.4`, update: `VM.Standard2.8`},
		"block_volume_size_in_gbs": Representation{repType: Required, create: `1000`},
	}

	bdsInstanceNodesMasterRepresentation = map[string]interface{}{
		"shape":                    Representation{repType: Required, create: `VM.Standard2.4`, update: `VM.Standard2.8`},
		"subnet_id":                Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"block_volume_size_in_gbs": Representation{repType: Required, create: `150`},
		"number_of_nodes":          Representation{repType: Required, create: `1`},
	}
	bdsInstanceNodesUtilRepresentation = map[string]interface{}{
		"shape":                    Representation{repType: Required, create: `VM.Standard2.4`, update: `VM.Standard2.8`},
		"subnet_id":                Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"block_volume_size_in_gbs": Representation{repType: Required, create: `150`},
		"number_of_nodes":          Representation{repType: Required, create: `1`},
	}
	bdsInstanceNodesWorkerRepresentation = map[string]interface{}{
		"shape":                    Representation{repType: Required, create: `VM.Standard2.1`, update: `VM.Standard2.4`},
		"subnet_id":                Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"block_volume_size_in_gbs": Representation{repType: Required, create: `150`},
		"number_of_nodes":          Representation{repType: Required, create: `3`, update: `4`},
	}
	bdsInstanceNetworkConfigRepresentation = map[string]interface{}{
		"cidr_block":              Representation{repType: Optional, create: `111.112.0.0/16`},
		"is_nat_gateway_required": Representation{repType: Optional, create: `false`},
	}

	BdsInstanceResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create,
		getMultipleUpdatedRepresenationCopy(
			[]string{"cidr_block", "dns_label"},
			[]interface{}{Representation{repType: Required, create: `111.111.0.0/24`}, Representation{repType: Required, create: `bdssubnet`}},
			subnetRegionalRepresentation)) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, getMultipleUpdatedRepresenationCopy(
			[]string{"cidr_block", "dns_label"},
			[]interface{}{Representation{repType: Required, create: `111.111.0.0/16`}, Representation{repType: Required, create: `bdsvcn`}},
			vcnRepresentation)) +
		DefinedTagsDependencies
)

func TestBdsBdsInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_bds_bds_instance.test_bds_instance"
	datasourceName := "data.oci_bds_bds_instances.test_bds_instances"
	singularDatasourceName := "data.oci_bds_bds_instance.test_bds_instance"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+BdsInstanceResourceDependencies+
		generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Optional, Create, bdsInstanceRepresentation), "bds", "bdsInstance", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckBdsBdsInstanceDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BdsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Required, Create, bdsInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
					resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
					resource.TestCheckResourceAttr(resourceName, "cluster_version", "CDH6"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "is_high_availability", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "nodes.#", "5"),
					resource.TestCheckResourceAttrSet(resourceName, "nodes.0.node_type"),
					resource.TestCheckResourceAttrSet(resourceName, "nodes.0.shape"),
					resource.TestCheckResourceAttrSet(resourceName, "nodes.0.subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BdsInstanceResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BdsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Optional, Create, bdsInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
					resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
					resource.TestCheckResourceAttr(resourceName, "cluster_version", "CDH6"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BdsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Optional, Create,
						representationCopyWithNewProperties(bdsInstanceRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
					resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
					resource.TestCheckResourceAttr(resourceName, "cluster_version", "CDH6"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + BdsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Optional, Update, bdsInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
					resource.TestCheckResourceAttrSet(resourceName, "cluster_public_key"),
					resource.TestCheckResourceAttr(resourceName, "cluster_version", "CDH6"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "master_node.0.shape", "VM.Standard2.8"),
					resource.TestCheckResourceAttr(resourceName, "worker_node.0.shape", "VM.Standard2.4"),
					resource.TestCheckResourceAttr(resourceName, "util_node.0.shape", "VM.Standard2.8"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_bds_bds_instances", "test_bds_instances", Optional, Update, bdsInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + BdsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Optional, Update, bdsInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "bds_instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.cluster_version", "CDH6"),
					resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "bds_instances.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Required, Create, bdsInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BdsInstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "cloud_sql_details.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cluster_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cluster_version", "CDH6"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.bd_cell_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.bds_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.csql_cell_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.db_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_details.0.os_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"), //empty
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + BdsInstanceResourceConfig,
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
		},
	})
}

func testAccCheckBdsBdsInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).bdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_instance" {
			noResourceFound = false
			request := oci_bds.GetBdsInstanceRequest{}

			tmp := rs.Primary.ID
			request.BdsInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "bds")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("BdsBdsInstance") {
		resource.AddTestSweepers("BdsBdsInstance", &resource.Sweeper{
			Name:         "BdsBdsInstance",
			Dependencies: DependencyGraph["bdsInstance"],
			F:            sweepBdsBdsInstanceResource,
		})
	}
}

func sweepBdsBdsInstanceResource(compartment string) error {
	bdsClient := GetTestClients(&schema.ResourceData{}).bdsClient()
	bdsInstanceIds, err := getBdsInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		if ok := SweeperDefaultResourceId[bdsInstanceId]; !ok {
			deleteBdsInstanceRequest := oci_bds.DeleteBdsInstanceRequest{}

			deleteBdsInstanceRequest.BdsInstanceId = &bdsInstanceId

			deleteBdsInstanceRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteBdsInstance(context.Background(), deleteBdsInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceId, error)
				continue
			}
			waitTillCondition(testAccProvider, &bdsInstanceId, bdsInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				bdsInstanceSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsInstanceIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "BdsInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := GetTestClients(&schema.ResourceData{}).bdsClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceId", id)
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

func bdsInstanceSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.bdsClient().GetBdsInstance(context.Background(), oci_bds.GetBdsInstanceRequest{
		BdsInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
