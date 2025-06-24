package network

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	SubnetResourceConfig         = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, SubnetResourceRepresentation)
	SubnetResourceRepresentation = map[string]interface{}{
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `tfSubnet`},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `10.1.20.0/24`},
		"vcn_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"route_table_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}},
		"dhcp_options_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"dns_label":         acctest.Representation{RepType: acctest.Required, Create: `tfsubnet`},
	}
)
