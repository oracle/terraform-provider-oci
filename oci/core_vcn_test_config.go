package oci

var (
	VcnRequiredOnlyResource = VcnRequiredOnlyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation)

	VcnResourceConfig = GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, VcnRepresentation)

	vcnSingularDataSourceRepresentation = map[string]interface{}{
		"vcn_id": Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	vcnDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, vcnDataSourceFilterRepresentation}}
	vcnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_vcn.test_vcn.id}`}},
	}

	VcnRepresentation = map[string]interface{}{
		"cidr_block":     Representation{RepType: Required, Create: `10.0.0.0/16`},
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"dns_label":      Representation{RepType: Optional, Create: `dnslabel`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      RepresentationGroup{Required, ignoreDefinedTagsChangesRep},
	}

	VcnRequiredOnlyResourceDependencies = ``
	VcnResourceDependencies             = DefinedTagsDependencies
)
