
# Create vcn, mgmt and related resources
# MgmtSubnets for DNS VMs
# LBSubnets for LBs
# BackendSubnets for Backend VMs

resource "oci_core_virtual_network" "MgmtVcn" {
    cidr_block = "${var.vcn_cidr}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "MgmtVcn"
    dns_label = "mgmtvcn"
}

resource "oci_core_internet_gateway" "MgmtIgw" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "MgmtIgw"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
}

resource "oci_core_route_table" "MgmtRouteTable" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    display_name = "MgmtRouteTable"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${oci_core_internet_gateway.MgmtIgw.id}"
    }
}

resource "oci_core_security_list" "MgmtSecurityList" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "MgmtSecurityList"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"

    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]

    ingress_security_rules = [{
        tcp_options {
            "max" = 53
            "min" = 53
        }
        protocol = "6"
        source = "${var.vcn_cidr}"
    },
    {
        udp_options {
            "max" = 53
            "min" = 53
        }
        protocol = "17"
        source = "${var.vcn_cidr}"
    },
    {
        tcp_options {
            "max" = 53
            "min" = 53
        }
        protocol = "6"
        source = "${var.onprem_cidr}"
    },
    {
        udp_options {
            "max" = 53
            "min" = 53
        }
        protocol = "17"
        source = "${var.onprem_cidr}"
    },
	{
        protocol = "all"
        source = "${var.vcn_cidr}"
    },
    {
        protocol = "6"
        source = "0.0.0.0/0"
        tcp_options {
            "min" = 22
            "max" = 22
        }
    },
    {
        protocol = "1"
        source = "0.0.0.0/0"
        icmp_options {
            "type" = 3
            "code" = 4
        }
    }]
}

resource "oci_core_dhcp_options" "MgmtDhcpOptions" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    display_name = "MgmtDhcpOptions"
 
    options {
      type = "DomainNameServer"
      server_type = "VcnLocalPlusInternet"
    }
}

resource "oci_core_subnet" "MgmtSubnet1" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    cidr_block = "${var.mgmt1_subnet_cidr}"
    display_name = "MgmtSubnet1"
    dns_label = "mgmtsubnet"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    route_table_id = "${oci_core_route_table.MgmtRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.MgmtSecurityList.id}"]
    dhcp_options_id = "${oci_core_dhcp_options.MgmtDhcpOptions.id}"
}

resource "oci_core_subnet" "MgmtSubnet2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    cidr_block = "${var.mgmt2_subnet_cidr}"
    display_name = "MgmtSubnet2"
    dns_label = "mgmtsubnet2"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    route_table_id = "${oci_core_route_table.MgmtRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.MgmtSecurityList.id}"]
    dhcp_options_id = "${oci_core_dhcp_options.MgmtDhcpOptions.id}"
}


# Create LB subnets are related resources

resource "oci_core_route_table" "LBRouteTable" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    display_name = "LBRouteTable"
}

resource "oci_core_security_list" "LBSecurityList" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    display_name = "LBSecurityList"

    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]

    ingress_security_rules = [{
        tcp_options {
            "max" = "${var.ha_app_port}"
            "min" = "${var.ha_app_port}"
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
    {
        protocol = "1"
        source = "0.0.0.0/0"
        stateless = true
        icmp_options {
            "type" = 3
            "code" = 4
        }
    }]
}

resource "oci_core_subnet" "LBSubnet1" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    cidr_block = "${var.lb1_subnet_cidr}"
    display_name = "LBSubnet1"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    route_table_id = "${oci_core_route_table.LBRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.LBSecurityList.id}"]
    dhcp_options_id = "${oci_core_virtual_network.MgmtVcn.default_dhcp_options_id}"
    prohibit_public_ip_on_vnic = true

    provisioner "local-exec" {
        command = "sleep 5"
    }
}

resource "oci_core_subnet" "LBSubnet2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    cidr_block = "${var.lb2_subnet_cidr}"
    display_name = "LBSubnet2"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    route_table_id = "${oci_core_route_table.LBRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.LBSecurityList.id}"]
    dhcp_options_id = "${oci_core_virtual_network.MgmtVcn.default_dhcp_options_id}"
    prohibit_public_ip_on_vnic = true

    provisioner "local-exec" {
        command = "sleep 5"
    }
}

# Create backend subnets for LB backend VMs

resource "oci_core_subnet" "BESubnet1" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    cidr_block = "${var.be1_subnet_cidr}"
    display_name = "BESubnet1"
    dns_label = "besubnet1"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    route_table_id = "${oci_core_route_table.BERouteTable.id}"
    security_list_ids = ["${oci_core_security_list.BESecurityList.id}"]
    dhcp_options_id = "${oci_core_virtual_network.MgmtVcn.default_dhcp_options_id}"

    provisioner "local-exec" {
        command = "sleep 5"
    }
}

resource "oci_core_subnet" "BESubnet2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    cidr_block = "${var.be2_subnet_cidr}"
    display_name = "BESubnet2"
    dns_label = "besubnet2"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    route_table_id = "${oci_core_route_table.BERouteTable.id}"
    security_list_ids = ["${oci_core_security_list.BESecurityList.id}"]
    dhcp_options_id = "${oci_core_virtual_network.MgmtVcn.default_dhcp_options_id}"

    provisioner "local-exec" {
        command = "sleep 5"
    }
}

resource "oci_core_route_table" "BERouteTable" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    display_name = "BERouteTable"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${oci_core_internet_gateway.MgmtIgw.id}"
    }
}

resource "oci_core_security_list" "BESecurityList" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.MgmtVcn.id}"
    display_name = "BESecurityList"

    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]

    ingress_security_rules = [{
        tcp_options {
            "max" = "${var.ha_app_port}"
            "min" = "${var.ha_app_port}"
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
    {
        tcp_options {
            "min" = 22
            "max" = 22
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
    {
        protocol = "1"
        source = "0.0.0.0/0"
        stateless = true
        icmp_options {
            "type" = 3
            "code" = 4
        }
    }]
}

