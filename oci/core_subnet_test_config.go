package oci

var (
	SubnetRequiredOnlyResource = SubnetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, SubnetRepresentation)

	SubnetResourceConfig = SubnetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Update, SubnetRepresentation)

	subnetSingularDataSourceRepresentation = map[string]interface{}{
		"subnet_id": Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	subnetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `MySubnet`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"vcn_id":         Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, subnetDataSourceFilterRepresentation}}
	subnetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_subnet.test_subnet.id}`}},
	}
	ignoreDefinedTagsChangesRep = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}

	SubnetRepresentation = map[string]interface{}{
		"cidr_block":                 Representation{RepType: Required, Create: `10.0.0.0/24`, Update: "10.0.0.0/16"},
		"compartment_id":             Representation{RepType: Required, Create: `${var.compartment_id}`},
		"vcn_id":                     Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain":        Representation{RepType: Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"defined_tags":               Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dhcp_options_id":            Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`, Update: `${oci_core_dhcp_options.test_dhcp_options.id}`},
		"display_name":               Representation{RepType: Optional, Create: `MySubnet`, Update: `displayName2`},
		"dns_label":                  Representation{RepType: Optional, Create: `dnslabel`},
		"freeform_tags":              Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"prohibit_public_ip_on_vnic": Representation{RepType: Optional, Create: `false`},
		"prohibit_internet_ingress":  Representation{RepType: Optional, Create: `false`},
		"route_table_id":             Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`, Update: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids":          Representation{RepType: Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}, Update: []string{`${oci_core_security_list.test_security_list.id}`}},
		"lifecycle":                  RepresentationGroup{Required, ignoreDefinedTagsChangesRep},
	}

	SubnetResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Required, Create, dhcpOptionsRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, RepresentationCopyWithNewProperties(VcnRepresentation, map[string]interface{}{
			"dns_label":      Representation{RepType: Required, Create: `dnslabel`},
			"is_ipv6enabled": Representation{RepType: Optional, Create: `true`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
	AnotherSecurityListRequiredOnlyResource = GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation)
	SubnetRequiredOnlyResourceDependencies  = AvailabilityDomainConfig + VcnResourceConfig
)