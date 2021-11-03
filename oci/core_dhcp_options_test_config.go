package oci

var (
	DhcpOptionsRequiredOnlyResource = GenerateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Required, Create, dhcpOptionsRepresentation)

	dhcpOptionsDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `MyDhcpOptions`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"vcn_id":         Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, dhcpOptionsDataSourceFilterRepresentation}}
	dhcpOptionsDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_dhcp_options.test_dhcp_options.id}`}},
	}

	dhcpOptionsRepresentation = map[string]interface{}{
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"options":          []RepresentationGroup{{Required, optionsRepresentation1}, {Required, optionsRepresentation2}},
		"vcn_id":           Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":     Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     Representation{RepType: Optional, Create: `MyDhcpOptions`, Update: `displayName2`},
		"domain_name_type": Representation{RepType: Optional, Create: `CUSTOM_DOMAIN`, Update: `VCN_DOMAIN`},
		"freeform_tags":    Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	optionsRepresentation1 = map[string]interface{}{
		"type":        Representation{RepType: Required, Create: `DomainNameServer`},
		"server_type": Representation{RepType: Required, Create: `VcnLocalPlusInternet`},
	}

	optionsRepresentation2 = map[string]interface{}{
		"type":                Representation{RepType: Required, Create: `SearchDomain`},
		"search_domain_names": Representation{RepType: Required, Create: []string{"test.com"}},
	}

	DhcpOptionsResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		DefinedTagsDependencies
)