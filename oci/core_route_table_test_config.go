package oci

var (
	RouteTableRequiredOnlyResource = RouteTableResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation)

	RouteTableResource = RouteTableResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Create, routeTableRepresentation)

	routeTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `MyRouteTable`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"vcn_id":         Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, routeTableDataSourceFilterRepresentation}}
	routeTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_route_table.test_route_table.id}`}},
	}

	routeTableRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"vcn_id":         Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Optional, Create: `MyRouteTable`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"route_rules":    RepresentationGroup{Optional, routeTableRouteRulesRepresentation},
	}
	routeTableRouteRulesRepresentation = map[string]interface{}{
		"network_entity_id": Representation{RepType: Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"description":       Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"destination":       Representation{RepType: Optional, Create: `0.0.0.0/0`, Update: `10.0.0.0/8`},
		"destination_type":  Representation{RepType: Optional, Create: `CIDR_BLOCK`},
	}

	RouteTableResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		DefinedTagsDependencies
)