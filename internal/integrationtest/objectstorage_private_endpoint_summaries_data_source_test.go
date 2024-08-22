package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstoragePrivateEndpointSummaryTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *DatasourceObjectstoragePrivateEndpointSummaryTestSuite) SetupTest() {
	s.Token, s.TokenFn = acctest.TokenizeWithHttpReplay("objectstorage")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() +
		`
			data "oci_identity_availability_domains" "ads1" {
				compartment_id = "${var.compartment_id}"
			}
		
			data "oci_objectstorage_namespace" "t1" {
				compartment_id = "${var.compartment_id}"
			}
		
			resource "oci_core_virtual_network" "test_vcn_1" {
				cidr_block     = "10.0.0.0/16"
				compartment_id = "${var.compartment_id}"
				display_name   = "network_name"
				dns_label      = "myvcn"
			}
		
			resource "oci_objectstorage_private_endpoint" "testPe1" {
				compartment_id = "${var.compartment_id}"
				namespace = "${data.oci_objectstorage_namespace.t1.namespace}"
				name = "testPe1"
				subnet_id = "${oci_core_subnet.test_subnet_1.id}"
				prefix = "testPrefix1"
				access_targets  {
					namespace = "*"
					compartment_id = "*"
					bucket = "*"
				  }
			}
		
			resource "oci_core_subnet" "test_subnet_1" {
				availability_domain = "${data.oci_identity_availability_domains.ads1.availability_domains.0.name}"
				cidr_block          = "10.0.1.0/24"
				display_name        = "-tf-subnet-1"
				compartment_id      = "${var.compartment_id}"
				vcn_id              = "${oci_core_virtual_network.test_vcn_1.id}"
				route_table_id      = "${oci_core_virtual_network.test_vcn_1.default_route_table_id}"
				dhcp_options_id     = "${oci_core_virtual_network.test_vcn_1.default_dhcp_options_id}"
				dns_label           = "testsubnet1"
			}
		`
	s.ResourceName = "data.oci_objectstorage_private_endpoint_summaries.t"
}

func (s *DatasourceObjectstoragePrivateEndpointSummaryTestSuite) TestAccDatasourceObjectstoragePrivateEndpointSummaries_basic() {
	compartmentID := acctest.GetCompartmentIDForLegacyTests()
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			// Client-side filtering.
			{
				Config: s.Config + `
				data "oci_objectstorage_private_endpoint_summaries" "t" {
					compartment_id = "${var.compartment_id}"
					namespace = "${data.oci_objectstorage_namespace.t1.namespace}"
					filter {
						name = "name"
						values = [oci_objectstorage_private_endpoint.testPe1.name]
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_endpoint_summaries.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_endpoint_summaries.0.prefix", testPrefix),
					resource.TestCheckResourceAttr(s.ResourceName, "private_endpoint_summaries.0.name", "testPe1"),
				),
			},
			{
				Config: s.Config + `
				data "oci_objectstorage_private_endpoint_summaries" "t" {
					compartment_id = "${var.compartment_id}"
					namespace = "${data.oci_objectstorage_namespace.t1.namespace}"
					filter {
						name = "name"
						values = ["non-existent-private-endpoint"]
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "private_endpoint_summaries.#", "0"),
				),
			},
		},
	})
}

// issue-routing-tag: object_storage/default
func TestDatasourceObjectstoragePrivateEndpointSummaryTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceObjectstoragePrivateEndpointSummaryTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceObjectstoragePrivateEndpointSummaryTestSuite))
}
