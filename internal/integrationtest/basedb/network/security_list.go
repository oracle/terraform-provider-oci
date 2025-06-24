package network

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	SecurityListResourceConfig         = acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Optional, acctest.Create, SecurityListResourceRepresentation)
	SecurityListResourceRepresentation = map[string]interface{}{
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `tfSecurityList`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"egress_security_rules":  acctest.RepresentationGroup{RepType: acctest.Required, Group: egressSecurityRulesGroup},
		"ingress_security_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: ingressSecurityRulesGroup},
	}

	egressSecurityRulesGroup = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Optional, Create: `6`},
		"destination": acctest.Representation{RepType: acctest.Optional, Create: `0.0.0.0/0`},
	}

	ingressSecurityRulesGroup = map[string]interface{}{
		"protocol": acctest.Representation{RepType: acctest.Optional, Create: `6`},
		"source":   acctest.Representation{RepType: acctest.Optional, Create: `0.0.0.0/0`},
	}
)
