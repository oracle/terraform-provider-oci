// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ClusterRequiredOnlyResource = ClusterResourceDependencies + `
resource "oci_containerengine_cluster" "test_cluster" {
	#Required
	compartment_id = "${var.compartment_id}"
	kubernetes_version = "${var.cluster_kubernetes_version}"
	name = "${var.cluster_name}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
	options {
		service_lb_subnet_ids = ["${oci_core_subnet.clusterSubnet_1.id}", "${oci_core_subnet.clusterSubnet_2.id}"]
	}
}
`
	ClusterResourceConfig = ClusterResourceDependencies + `
resource "oci_containerengine_cluster" "test_cluster" {
	#Required
	compartment_id = "${var.compartment_id}"
	kubernetes_version = "${var.cluster_kubernetes_version}"
	name = "${var.cluster_name}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	options {
		service_lb_subnet_ids = ["${oci_core_subnet.clusterSubnet_1.id}", "${oci_core_subnet.clusterSubnet_2.id}"]

		#Optional
		add_ons {
			#Optional
			is_kubernetes_dashboard_enabled = "${var.cluster_options_add_ons_is_kubernetes_dashboard_enabled}"
			is_tiller_enabled = "${var.cluster_options_add_ons_is_tiller_enabled}"
		}
		kubernetes_network_config {
			#Optional
			pods_cidr = "${var.cluster_options_kubernetes_network_config_pods_cidr}"
			services_cidr = "${var.cluster_options_kubernetes_network_config_services_cidr}"
		}
	}
}
`
	ClusterPropertyVariables = `
variable "cluster_kubernetes_version" { default = "v1.8.11" }
variable "cluster_name" { default = "name" }
variable "cluster_options_add_ons_is_kubernetes_dashboard_enabled" { default = true }
variable "cluster_options_add_ons_is_tiller_enabled" { default = true }
variable "cluster_options_kubernetes_network_config_pods_cidr" { default = "10.1.0.0/16" }
variable "cluster_options_kubernetes_network_config_services_cidr" { default = "10.2.0.0/16" }
variable "cluster_state" { default = [] }

`
	ClusterResourceDependencies = VcnPropertyVariables + VcnResourceConfig + AvailabilityDomainConfig + `
resource "oci_core_subnet" "clusterSubnet_1" {
       #Required
	   availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
	   cidr_block = "10.0.20.0/24"
	   compartment_id = "${var.compartment_id}"
	   vcn_id = "${oci_core_vcn.test_vcn.id}"
       security_list_ids = ["${oci_core_vcn.test_vcn.default_security_list_id}"] # Provider code tries to maintain compatibility with old versions.
       display_name = "tfSubNet1ForClusters"
}
resource "oci_core_subnet" "clusterSubnet_2" {
       #Required
       availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
       cidr_block = "10.0.21.0/24"
       compartment_id = "${var.compartment_id}"
       vcn_id = "${oci_core_vcn.test_vcn.id}"
       display_name = "tfSubNet1ForClusters"
    security_list_ids = ["${oci_core_vcn.test_vcn.default_security_list_id}"] # Provider code tries to maintain compatibility with old versions.
}`
)

func TestContainerengineClusterResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	datasourceName := "data.oci_containerengine_clusters.test_clusters"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + ClusterPropertyVariables + compartmentIdVariableStr + ClusterRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "kubernetes_version", "v1.8.11"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ClusterResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + ClusterPropertyVariables + compartmentIdVariableStr + ClusterResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "kubernetes_version", "v1.8.11"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_tiller_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "cluster_kubernetes_version" { default = "v1.9.7" }
variable "cluster_name" { default = "name2" }
variable "cluster_options_add_ons_is_kubernetes_dashboard_enabled" { default = true }
variable "cluster_options_add_ons_is_tiller_enabled" { default = true }
variable "cluster_options_kubernetes_network_config_pods_cidr" { default = "10.1.0.0/16" }
variable "cluster_options_kubernetes_network_config_services_cidr" { default = "10.2.0.0/16" }
variable "cluster_options_service_lb_subnet_ids" { default = [] }
variable "cluster_state" { default = [] }

                ` + compartmentIdVariableStr + ClusterResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "kubernetes_version", "v1.9.7"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_tiller_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				Config: config + `
variable "cluster_kubernetes_version" { default = "v1.9.7" }
variable "cluster_name" { default = "name2" }
variable "cluster_options_add_ons_is_kubernetes_dashboard_enabled" { default = true }
variable "cluster_options_add_ons_is_tiller_enabled" { default = true }
variable "cluster_options_kubernetes_network_config_pods_cidr" { default = "10.1.0.0/16" }
variable "cluster_options_kubernetes_network_config_services_cidr" { default = "10.2.0.0/16" }
variable "cluster_options_service_lb_subnet_ids" { default = [] }
variable "cluster_state" { default = [] }

data "oci_containerengine_clusters" "test_clusters" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	name = "${var.cluster_name}"

    filter {
    	name = "id"
    	values = ["${oci_containerengine_cluster.test_cluster.id}"]
    }
}
                ` + compartmentIdVariableStr + ClusterResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "name2"),

					resource.TestCheckResourceAttr(datasourceName, "clusters.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.kubernetes_version", "v1.9.7"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.name", "name2"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.0.is_tiller_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.service_lb_subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.vcn_id"),
				),
			},
		},
	})
}
