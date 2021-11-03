package oci

var (
	SecurityListRequiredOnlyResource = SecurityListResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation)

	securityListDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `MyPrivateSubnetSecurityList`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"vcn_id":         Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, securityListDataSourceFilterRepresentation}}
	securityListDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}},
	}

	securityListRepresentation = map[string]interface{}{
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"vcn_id":                 Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":           Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           Representation{RepType: Optional, Create: `MyPrivateSubnetSecurityList`, Update: `displayName2`},
		"egress_security_rules":  []RepresentationGroup{{Required, securityListEgressSecurityRulesICMPRepresentation}, {Optional, securityListEgressSecurityRulesTCPRepresentation}, {Optional, securityListEgressSecurityRulesUDPRepresentation}},
		"freeform_tags":          Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ingress_security_rules": []RepresentationGroup{{Required, securityListIngressSecurityRulesICMPRepresentation}, {Optional, securityListIngressSecurityRulesTCPRepresentation}, {Optional, securityListIngressSecurityRulesUDPRepresentation}},
	}
	securityListEgressSecurityRulesICMPRepresentation = map[string]interface{}{
		"destination":      Representation{RepType: Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         Representation{RepType: Required, Create: `1`},
		"description":      Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"destination_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"icmp_options":     RepresentationGroup{Optional, securityListEgressSecurityRulesIcmpOptionsRepresentation},
		"stateless":        Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	securityListEgressSecurityRulesTCPRepresentation = map[string]interface{}{
		"destination":      Representation{RepType: Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         Representation{RepType: Required, Create: `6`},
		"destination_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":        Representation{RepType: Optional, Create: `false`, Update: `true`},
		"tcp_options":      RepresentationGroup{Optional, securityListEgressSecurityRulesTcpOptionsRepresentation},
	}
	securityListEgressSecurityRulesUDPRepresentation = map[string]interface{}{
		"destination":      Representation{RepType: Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         Representation{RepType: Required, Create: `17`},
		"destination_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":        Representation{RepType: Optional, Create: `false`, Update: `true`},
		"udp_options":      RepresentationGroup{Optional, securityListEgressSecurityRulesUdpOptionsRepresentation},
	}
	securityListIngressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol":     Representation{RepType: Required, Create: `1`},
		"description":  Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"source":       Representation{RepType: Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"icmp_options": RepresentationGroup{Optional, securityListIngressSecurityRulesIcmpOptionsRepresentation},
		"source_type":  Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":    Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	securityListIngressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol":    Representation{RepType: Required, Create: `6`},
		"source":      Representation{RepType: Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   Representation{RepType: Optional, Create: `false`, Update: `true`},
		"tcp_options": RepresentationGroup{Optional, securityListIngressSecurityRulesTcpOptionsRepresentation},
	}
	securityListIngressSecurityRulesUDPRepresentation = map[string]interface{}{
		"protocol":    Representation{RepType: Required, Create: `17`},
		"source":      Representation{RepType: Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   Representation{RepType: Optional, Create: `false`, Update: `true`},
		"udp_options": RepresentationGroup{Optional, securityListIngressSecurityRulesUdpOptionsRepresentation},
	}
	securityListEgressSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": Representation{RepType: Required, Create: `3`},
		"code": Representation{RepType: Optional, Create: `4`, Update: `0`},
	}
	securityListEgressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"min":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"source_port_range": RepresentationGroup{Optional, securityListEgressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	securityListEgressSecurityRulesUdpOptionsRepresentation = map[string]interface{}{
		"max":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"min":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"source_port_range": RepresentationGroup{Optional, securityListEgressSecurityRulesUdpOptionsSourcePortRangeRepresentation},
	}
	securityListIngressSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": Representation{RepType: Required, Create: `3`},
		"code": Representation{RepType: Optional, Create: `4`, Update: `0`},
	}
	securityListIngressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"min":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"source_port_range": RepresentationGroup{Optional, securityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	securityListIngressSecurityRulesUdpOptionsRepresentation = map[string]interface{}{
		"max":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"min":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"source_port_range": RepresentationGroup{Optional, securityListIngressSecurityRulesUdpOptionsSourcePortRangeRepresentation},
	}
	securityListEgressSecurityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityListEgressSecurityRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityListIngressSecurityRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}

	SecurityListResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		DefinedTagsDependencies
)
