package network

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	VcnResourceConfig         = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, VcnResourceRepresentation)
	VcnResourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfVcn`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.1.0.0/16`},
		"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `tfvcn`},
	}
)
