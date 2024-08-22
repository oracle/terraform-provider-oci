package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

func TestResourcePrivateEndpoint(t *testing.T) {
	httpreplay.SetScenario("TestObjectStoragePrivateEndpointResource_resource")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.LegacyTestProviderConfig()
	singularDatasourceName := "oci_objectstorage_private_endpoint.testPe"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStoragePrivateEndpointDestroy,
		Steps: []resource.TestStep{
			// Verify Create with optionals
			{
				Config: config +
					`
						data "oci_identity_availability_domains" "ADs" {
							compartment_id = "${var.compartment_id}"
						}
				
						data "oci_objectstorage_namespace" "t" {
							compartment_id = "${var.compartment_id}"
						}
				
						resource "oci_core_virtual_network" "test_vcn" {
							cidr_block     = "10.0.0.0/16"
							compartment_id = "${var.compartment_id}"
							display_name   = "network_name"
							dns_label      = "myvcn"
						}
				
						resource "oci_objectstorage_private_endpoint" "testPe" {
							compartment_id = "${var.compartment_id}"
							namespace = "${data.oci_objectstorage_namespace.t.namespace}"
							name = "testPe"
							subnet_id = "${oci_core_subnet.test_subnet.id}"
							prefix = "testPrefix"
							access_targets  {
								namespace = "*"
								compartment_id = "*"
								bucket = "*"
							  }
						}
				
						resource "oci_core_subnet" "test_subnet" {
							availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
							cidr_block          = "10.0.1.0/24"
							display_name        = "-tf-subnet"
							compartment_id      = "${var.compartment_id}"
							vcn_id              = "${oci_core_virtual_network.test_vcn.id}"
							route_table_id      = "${oci_core_virtual_network.test_vcn.default_route_table_id}"
							dhcp_options_id     = "${oci_core_virtual_network.test_vcn.default_dhcp_options_id}"
							dns_label           = "testsubnet"
						}
					`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "testPe"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "prefix", "testPrefix"),
					acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "access_targets",
						map[string]string{
							"namespace":      "*",
							"compartment_id": "*",
							"bucket":         "*",
						},
						[]string{}),
				),
			},
		},
	})
}
