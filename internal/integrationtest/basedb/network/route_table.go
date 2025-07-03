package network

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	RouteTableResourceConfig         = acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, RouteTableResourceRepresentation)
	RouteTableResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfRouteTable`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_rules":    acctest.RepresentationGroup{RepType: acctest.Required, Group: routeRulesGroup},
	}

	routeRulesGroup = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"description":       acctest.Representation{RepType: acctest.Required, Create: `Internal traffic for OCI Services`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
	}
)
