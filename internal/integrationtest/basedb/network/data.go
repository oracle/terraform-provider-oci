package network

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	AvailabilityDomainsDatasourceConfig         = acctest.GenerateDataSourceFromRepresentationMap("oci_identity_availability_domains", "test_availability_domains", acctest.Optional, acctest.Create, AvailabilityDomainsDatasourceRepresentation)
	AvailabilityDomainsDatasourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}
)
