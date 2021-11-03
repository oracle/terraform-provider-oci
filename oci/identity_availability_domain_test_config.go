package oci

var (
	availabilityDomainSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"ad_number":      Representation{RepType: Optional, Create: `2`},
	}

	availabilityDomainDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	AvailabilityDomainResourceConfig = ""

	AvailabilityDomainConfig = GenerateDataSourceFromRepresentationMap("oci_identity_availability_domains", "test_availability_domains", Required, Create, availabilityDomainDataSourceRepresentation)
)