// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OcvpClusterRequiredOnlyResource = OcvpClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Required, acctest.Create, OcvpClusterRepresentation)

	OcvpClusterResourceConfig = OcvpClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Optional, acctest.Update, OcvpClusterRepresentation)

	OcvpClusterSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_cluster.test_cluster.id}`},
	}

	OcvpClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"sddc_id":        acctest.Representation{RepType: acctest.Optional, Create: `${local.upgraded_sddc_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpClusterDataSourceFilterRepresentation}}
	OcvpClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_cluster.test_cluster.id}`}},
	}

	OcvpClusterRepresentation = map[string]interface{}{
		"compute_availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}`},
		"esxi_hosts_count":             acctest.Representation{RepType: acctest.Required, Create: `1`},
		"network_configuration":        acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpClusterNetworkConfigurationRepresentation},
		"sddc_id":                      acctest.Representation{RepType: acctest.Required, Create: `${local.upgraded_sddc_id}`},
		"capacity_reservation_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"datastores":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: OcvpClusterDatastoresRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"esxi_software_version":        acctest.Representation{RepType: acctest.Optional, Create: `esxi7u3k-21313628-1`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"initial_commitment":           acctest.Representation{RepType: acctest.Optional, Create: `HOUR`},
		"initial_host_ocpu_count":      acctest.Representation{RepType: acctest.Optional, Create: `12`},
		"initial_host_shape_name":      acctest.Representation{RepType: acctest.Optional, Create: `BM.Standard2.52`},
		"instance_display_name_prefix": acctest.Representation{RepType: acctest.Optional, Create: `tf-test-`},
		"is_shielded_instance_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"vmware_software_version":      acctest.Representation{RepType: acctest.Optional, Create: noInstanceVmwareVersionV7},
		"workload_network_cidr":        acctest.Representation{RepType: acctest.Optional, Create: `172.20.0.0/24`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
	}
	OcvpClusterNetworkConfigurationRepresentation = map[string]interface{}{
		"nsx_edge_vtep_vlan_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.nsx_edge_vtep_vlan_id}`},
		"nsx_vtep_vlan_id":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.nsx_vtep_vlan_id}`},
		"provisioning_subnet_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.provisioning_subnet_id}`},
		"vmotion_vlan_id":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.vmotion_vlan_id}`},
		"vsan_vlan_id":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.vsan_vlan_id}`},
		"hcx_vlan_id":             acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.hcx_vlan_id}`},
		"nsx_edge_uplink1vlan_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.nsx_edge_uplink1vlan_id}`},
		"nsx_edge_uplink2vlan_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.nsx_edge_uplink2vlan_id}`},
		"provisioning_vlan_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.provisioning_vlan_id}`},
		"replication_vlan_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.replication_vlan_id}`},
		"vsphere_vlan_id":         acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_cluster.v7_sddc_management_cluster.network_configuration.0.vsphere_vlan_id}`},
	}
	OcvpClusterDatastoresRepresentation = map[string]interface{}{
		"block_volume_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume.test_volume.id}`}},
		"datastore_type":   acctest.Representation{RepType: acctest.Required, Create: `WORKLOAD`},
	}

	OcvpClusterResourceDependencies = EsxiHostResourceDependencies + ocvpAvailabilityDomainDependency + `
		data "oci_ocvp_cluster" "v7_sddc_management_cluster" {
		  cluster_id = data.oci_ocvp_clusters.test_clusters_v7_management.cluster_collection[0].items[0].id
		}
`
	OcvpClusterOptionalResourceDependencies = OcvpClusterResourceDependencies + DefinedTagsDependencies + `
	resource "oci_core_compute_capacity_reservation" "test_compute_capacity_reservation" {
		compartment_id = var.compartment_id
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
		display_name   = "tf-esxi-host-test-capacity-reservation"
		instance_reservation_configs {
			instance_shape = "BM.Standard2.52"
			reserved_count = 1
			fault_domain = "FAULT-DOMAIN-1"
		}
		instance_reservation_configs {
			instance_shape = "BM.Standard2.52"
			reserved_count = 1
			fault_domain = "FAULT-DOMAIN-2"
		}
		instance_reservation_configs {
			instance_shape = "BM.Standard2.52"
			reserved_count = 1
			fault_domain = "FAULT-DOMAIN-3"
		}
	}

	resource "oci_core_volume" "test_volume" {
	  display_name		  = "test_volume_management_cluster"
	  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
	  compartment_id      = var.compartment_id
	  vpus_per_gb		  = 10
	  size_in_gbs         = 4096
	}
`
)

// issue-routing-tag: ocvp/default
func TestOcvpClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_ocvp_cluster.test_cluster"
	datasourceName := "data.oci_ocvp_clusters.test_clusters"
	singularDatasourceName := "data.oci_ocvp_cluster.test_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OcvpClusterOptionalResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Optional, acctest.Create, OcvpClusterRepresentation), "ocvp", "cluster", t)

	acctest.ResourceTest(t, testAccCheckOcvpClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OcvpClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Required, acctest.Create, OcvpClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "esxi_hosts_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OcvpClusterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OcvpClusterOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Optional, acctest.Create, OcvpClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "datastores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datastores.0.block_volume_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "datastores.0.capacity"),
				resource.TestCheckResourceAttr(resourceName, "datastores.0.datastore_type", "WORKLOAD"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "esxi_hosts_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "esxi_software_version", "esxi7u3k-21313628-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_commitment", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "initial_host_ocpu_count", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_host_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "instance_display_name_prefix", "tf-test-"),
				resource.TestCheckResourceAttr(resourceName, "is_shielded_instance_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.provisioning_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.replication_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vsphere_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttrSet(resourceName, "vsphere_type"),
				resource.TestCheckResourceAttr(resourceName, "workload_network_cidr", "172.20.0.0/24"),

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
			Config: config + compartmentIdVariableStr + OcvpClusterOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Optional, acctest.Update, OcvpClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "datastores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datastores.0.block_volume_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "datastores.0.capacity"),
				resource.TestCheckResourceAttr(resourceName, "datastores.0.datastore_type", "WORKLOAD"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "esxi_hosts_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_commitment", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "initial_host_ocpu_count", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_host_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "instance_display_name_prefix", "tf-test-"),
				resource.TestCheckResourceAttr(resourceName, "is_shielded_instance_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.provisioning_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.replication_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vsphere_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttrSet(resourceName, "vsphere_type"),
				resource.TestCheckResourceAttr(resourceName, "workload_network_cidr", "172.20.0.0/24"),

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
			Config: config + compartmentIdVariableStr + OcvpClusterOptionalResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_clusters", "test_clusters", acctest.Optional, acctest.Update, OcvpClusterDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Optional, acctest.Update, OcvpClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "cluster_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + OcvpClusterOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Optional, acctest.Update, OcvpClusterRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_cluster", "test_cluster", acctest.Required, acctest.Create, OcvpClusterSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datastores.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datastores.0.block_volume_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "datastores.0.capacity"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datastores.0.datastore_type", "WORKLOAD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "esxi_hosts_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "esxi_software_version", "esxi7u3k-21313628-1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_commitment", "HOUR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_host_ocpu_count", "12"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_display_name_prefix", "tf-test-"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_shielded_instance_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "upgrade_licenses.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vsphere_upgrade_objects.#"),

				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vsphere_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "workload_network_cidr", "172.20.0.0/24"),
			),
		},
		// verify resource import
		{
			Config:                  config + OcvpClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ClusterClient()
	clusterDependencyDataSourceId, err := acctest.FromInstanceState(s, "data.oci_ocvp_cluster.v7_sddc_management_cluster", "id")
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_cluster" && rs.Primary.ID != clusterDependencyDataSourceId {

			noResourceFound = false
			request := oci_ocvp.GetClusterRequest{}

			tmp := rs.Primary.ID
			request.ClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("OcvpCluster") {
		resource.AddTestSweepers("OcvpCluster", &resource.Sweeper{
			Name:         "OcvpCluster",
			Dependencies: acctest.DependencyGraph["cluster"],
			F:            sweepOcvpClusterResource,
		})
	}
}

func sweepOcvpClusterResource(compartment string) error {
	clusterClient := acctest.GetTestClients(&schema.ResourceData{}).ClusterClient()
	clusterIds, err := getOcvpClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterId := range clusterIds {
		if ok := acctest.SweeperDefaultResourceId[clusterId]; !ok {
			deleteClusterRequest := oci_ocvp.DeleteClusterRequest{}

			deleteClusterRequest.ClusterId = &clusterId

			deleteClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := clusterClient.DeleteCluster(context.Background(), deleteClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting Cluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterId, OcvpClusterSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpClusterSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	clusterClient := acctest.GetTestClients(&schema.ResourceData{}).ClusterClient()

	listClustersRequest := oci_ocvp.ListClustersRequest{}
	listClustersRequest.CompartmentId = &compartmentId
	listClustersRequest.LifecycleState = oci_ocvp.ListClustersLifecycleStateActive
	listClustersResponse, err := clusterClient.ListClusters(context.Background(), listClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Cluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cluster := range listClustersResponse.Items {
		id := *cluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterId", id)
	}
	return resourceIds, nil
}

func OcvpClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterResponse, ok := response.Response.(oci_ocvp.GetClusterResponse); ok {
		return clusterResponse.LifecycleState != oci_ocvp.LifecycleStatesDeleted
	}
	return false
}

func OcvpClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ClusterClient().GetCluster(context.Background(), oci_ocvp.GetClusterRequest{
		ClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
