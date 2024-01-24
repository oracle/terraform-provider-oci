// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/service/ocvp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OcvpSddcRequiredOnlyResource = OcvpSddcResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Create, OcvpSddcRepresentation)

	OcvpSddcResourceConfig = OcvpSddcOptionalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Optional, acctest.Update, OcvpSddcRepresentation)

	SddcV7ResourceConfig = OcvpSddcOptionalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Optional, acctest.Update, sddcV7Representation)

	OcvpSddcUpgradeResource = OcvpSddcResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Update, sddcUpgradedRepresentation)

	OcvpOcvpSddcSingularDataSourceRepresentation = map[string]interface{}{
		"sddc_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_sddc.test_sddc.id}`},
	}

	// use random name to avoid conflict in parallel tests
	sddcDisplayName1 = fmt.Sprintf("%s-%d", "test", rand.Intn(10000))
	sddcDisplayName2 = sddcDisplayName1 + "u"

	OcvpOcvpSddcDataSourceRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: sddcDisplayName1, Update: sddcDisplayName2},
		"state":                       acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpSddcDataSourceFilterRepresentation}}
	OcvpSddcDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_sddc.test_sddc.id}`}},
	}

	noInstanceVmwareVersionV6 = "test-no-instance-no-grace-period"
	noInstanceVmwareVersionV7 = "7.0 test-no-instance-no-grace-period"

	sshKey        = `ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAH1y0s/y8/nd6qsDcwrga3x7W6MlMK7u6Mx+iy9sI7GVpzuVeFj+6xxCnf6vsI+4p6wsYDktYfyggMsGMvHNcpuzwFQ2tb4HDxLFDNUsUUBuItns9GBU0sWOBgmP2s5M82ueKf6vTPky5M4mGMPDD4wvjK5hIe6SqdisKiJgP6AGg19iw==`
	sshKeyUpdated = `ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBACS6C+rdwWTs/j/Yn1XAS7+U0HdAAFHIua9Y+2vbe5kwY6WaHycxp4ntTNHPeeOC3QEX+3guiUY0j26LStWf2moewGH2q3i9MQGdrk0ojH0hy/ToAntFQcB8Ghh3RjCS+Dy8pIRtnb79am58ykaKq6pf9qM/f0rqYT0rly8DjW1uM5G6g==`

	esxiSoftwareVersion        = `esxi6.7-19195723-2`
	esxiSoftwareVersionUpdated = `esxi7u3k-21313628-1`

	OcvpSddcRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"initial_configuration":   acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpSddcInitialConfigurationRepresentation},
		"ssh_authorized_keys":     acctest.Representation{RepType: acctest.Required, Create: sshKey, Update: sshKeyUpdated},
		"vmware_software_version": acctest.Representation{RepType: acctest.Required, Create: noInstanceVmwareVersionV6, Update: noInstanceVmwareVersionV7},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: sddcDisplayName1, Update: sddcDisplayName2},
		"esxi_software_version":   acctest.Representation{RepType: acctest.Optional, Create: esxiSoftwareVersion, Update: esxiSoftwareVersionUpdated},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_single_host_sddc":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_hcx_enabled":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"hcx_action":              acctest.Representation{RepType: acctest.Optional, Create: ocvp.UpgradeHcxAction},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
	}
	OcvpSddcInitialConfigurationRepresentation = map[string]interface{}{
		"initial_cluster_configurations": acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpSddcInitialConfigurationInitialClusterConfigurationsRepresentation},
	}

	OcvpSddcInitialConfigurationUpdateRepresentation = map[string]interface{}{
		"initial_cluster_configurations": acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpSddcInitialConfigurationInitialClusterConfigurationsUpdateRepresentation},
	}

	OcvpSddcInitialConfigurationInitialClusterConfigurationsRepresentation = map[string]interface{}{
		"compute_availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}`},
		"esxi_hosts_count":             acctest.Representation{RepType: acctest.Required, Create: `3`},
		"network_configuration":        acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpSddcInitialConfigurationInitialClusterConfigurationsNetworkConfigurationRepresentation},
		"vsphere_type":                 acctest.Representation{RepType: acctest.Required, Create: `MANAGEMENT`},
		"capacity_reservation_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"datastores":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: OcvpSddcDatastoresRepresentation},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: "displayName"},
		"initial_commitment":           acctest.Representation{RepType: acctest.Optional, Create: `HOUR`},
		"initial_host_ocpu_count":      acctest.Representation{RepType: acctest.Optional, Create: `12`},
		"initial_host_shape_name":      acctest.Representation{RepType: acctest.Optional, Create: `BM.Standard2.52`},
		"instance_display_name_prefix": acctest.Representation{RepType: acctest.Optional, Create: `tf-test-`},
		"is_shielded_instance_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"workload_network_cidr":        acctest.Representation{RepType: acctest.Optional, Create: `172.20.0.0/24`},
	}

	OcvpSddcInitialConfigurationInitialClusterConfigurationsUpdateRepresentation = map[string]interface{}{
		"network_configuration":        acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpSddcInitialConfigurationInitialClusterConfigurationsNetworkConfigurationUpdateRepresentation},
		"compute_availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}`},
		"esxi_hosts_count":             acctest.Representation{RepType: acctest.Required, Create: `3`},
		"vsphere_type":                 acctest.Representation{RepType: acctest.Required, Create: `MANAGEMENT`},
		"capacity_reservation_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"datastores":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: OcvpSddcDatastoresRepresentation},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: "displayName"},
		"initial_commitment":           acctest.Representation{RepType: acctest.Optional, Create: `HOUR`},
		"initial_host_ocpu_count":      acctest.Representation{RepType: acctest.Optional, Create: `12`},
		"initial_host_shape_name":      acctest.Representation{RepType: acctest.Optional, Create: `BM.Standard2.52`},
		"instance_display_name_prefix": acctest.Representation{RepType: acctest.Optional, Create: `tf-test-`},
		"is_shielded_instance_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"workload_network_cidr":        acctest.Representation{RepType: acctest.Optional, Create: `172.20.0.0/24`},
	}

	OcvpSddcInitialConfigurationInitialClusterConfigurationsNetworkConfigurationRepresentation = map[string]interface{}{
		"nsx_edge_vtep_vlan_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_nsx_edge_vtep_vlan.id}`},
		"nsx_vtep_vlan_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_nsx_vtep_vlan.id}`},
		"provisioning_subnet_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_provisioning_subnet.id}`},
		"vmotion_vlan_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vmotion_net_vlan.id}`},
		"vsan_vlan_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vsan_net_vlan.id}`},
		"hcx_vlan_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vlan.test_hcx_vlan.id}`},
		"nsx_edge_uplink1vlan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_nsx_edge_uplink1_vlan.id}`},
		"nsx_edge_uplink2vlan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_nsx_edge_uplink2_vlan.id}`},
		"provisioning_vlan_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vlan.test_provisioning_vlan.id}`},
		"replication_vlan_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vlan.test_replication_vlan.id}`},
		"vsphere_vlan_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vsphere_net_vlan.id}`},
	}

	OcvpSddcInitialConfigurationInitialClusterConfigurationsNetworkConfigurationUpdateRepresentation = map[string]interface{}{
		"provisioning_vlan_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_provisioning_vlan.id}`},
		"replication_vlan_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_replication_vlan.id}`},
		"nsx_edge_vtep_vlan_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_nsx_edge_vtep_vlan.id}`},
		"nsx_vtep_vlan_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_nsx_vtep_vlan.id}`},
		"provisioning_subnet_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_provisioning_subnet.id}`},
		"vmotion_vlan_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vmotion_net_vlan.id}`},
		"vsan_vlan_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vsan_net_vlan.id}`},
		"hcx_vlan_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vlan.test_hcx_vlan.id}`},
		"nsx_edge_uplink1vlan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_nsx_edge_uplink1_vlan.id}`},
		"nsx_edge_uplink2vlan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_nsx_edge_uplink2_vlan.id}`},
		"vsphere_vlan_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vsphere_net_vlan.id}`},
	}

	ignoreDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	OcvpSddcDatastoresRepresentation = map[string]interface{}{
		"block_volume_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume.test_volume.id}`}},
		"datastore_type":   acctest.Representation{RepType: acctest.Required, Create: `MANAGEMENT`},
	}

	sddcV7Representation = acctest.RepresentationCopyWithNewProperties(OcvpSddcRepresentation, map[string]interface{}{
		"vmware_software_version": acctest.Representation{RepType: acctest.Required, Create: noInstanceVmwareVersionV7},
		"esxi_software_version":   acctest.Representation{RepType: acctest.Required, Create: esxiSoftwareVersionUpdated},
	})

	sddcUpgradedRepresentation = acctest.RepresentationCopyWithNewProperties(OcvpSddcRepresentation, map[string]interface{}{
		"vmware_software_version": acctest.Representation{RepType: acctest.Required, Create: noInstanceVmwareVersionV7},
		"initial_configuration":   acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpSddcInitialConfigurationUpdateRepresentation},
	})

	OcvpSddcResourceDependencies = DefinedTagsDependencies + `

data "oci_core_services" "test_services" {}

data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.compartment_id}"
}

resource "oci_core_vcn" "test_vcn_ocvp" {
    cidr_block = "10.0.0.0/16"
    compartment_id = "${var.compartment_id}"
    display_name = "VmWareOCVP"
    dns_label = "vmwareocvp"
}


resource oci_core_nat_gateway test_nat_gateway_ocvp {
  block_traffic  = "false"
  compartment_id = var.compartment_id

  display_name = "NAT Gateway OCVP"
  freeform_tags = {
    "VCN" = "VCN-2020-09-11T00:43:42"
  }
  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource oci_core_route_table test_route_table_for_vsphere_vlan {
  compartment_id = var.compartment_id

  display_name = "Route Table for VLAN-grk-vSphere"
  freeform_tags = {
    "VMware" = "VMware-2020-09-11T00:47:02"
  }
  route_rules {
    #description = <<Optional value not found in discovery>>
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_nat_gateway.test_nat_gateway_ocvp.id
  }
  vcn_id = oci_core_vcn.test_vcn_ocvp.id

  lifecycle {
    ignore_changes = [ route_rules ]
  }
}

resource "oci_core_network_security_group" "test_nsg_allow_all" {
    compartment_id = "${var.compartment_id}"
    display_name = "nsg-allow-all"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource oci_core_network_security_group_security_rule test_nsg_security_rule_1 {
  destination_type          = ""
  direction                 = "INGRESS"
  network_security_group_id = "${oci_core_network_security_group.test_nsg_allow_all.id}"
  protocol                  = "all"
  source                    = "0.0.0.0/0"
  source_type               = "CIDR_BLOCK"
}

resource oci_core_network_security_group_security_rule test_nsg_security_rule_2 {
  destination               = "0.0.0.0/0"
  destination_type          = "CIDR_BLOCK"
  direction                 = "EGRESS"
  network_security_group_id = "${oci_core_network_security_group.test_nsg_allow_all.id}"
  protocol                  = "all"
  source_type = ""
}

resource "oci_core_service_gateway" "export_sgw" {
    compartment_id = "${var.compartment_id}"
    display_name = "sgw"
    services {
        service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
    }
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource "oci_core_default_dhcp_options" "default_dhcp_options_ocvp"{
    display_name = "Default DHCP Options for OCVP"
    manage_default_resource_id = "${oci_core_vcn.test_vcn_ocvp.default_dhcp_options_id}"
    options {
        custom_dns_servers = []
        server_type = "VcnLocalPlusInternet"
        type = "DomainNameServer"
    }
    options {
            search_domain_names = ["vmwareocvp.oraclevcn.com"]
            type = "SearchDomain"
    }
}

resource "oci_core_route_table" "private_rt" {
    compartment_id = "${var.compartment_id}"
    display_name = "private-rt"
    route_rules {
        destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
        destination_type = "SERVICE_CIDR_BLOCK"
        network_entity_id = "${oci_core_service_gateway.export_sgw.id}"
    }
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"

  lifecycle {
    ignore_changes = [ route_rules ]
  }
}

resource "oci_core_security_list" "private_sl" {
    compartment_id = "${var.compartment_id}"
    display_name = "private-sl"
    egress_security_rules {
        destination = "0.0.0.0/0"
        destination_type = "CIDR_BLOCK"
        protocol = "all"
        stateless = "false"
    }
    ingress_security_rules {
        description = "TCP traffic for ports: 22 SSH Remote Login Protocol"
        protocol = "6"
        source = "10.0.0.0/16"
        source_type = "CIDR_BLOCK"
        stateless = "false"
        tcp_options {
            max = "22"
            min = "22"
        }
    }
    ingress_security_rules {
        description = "ICMP traffic for: 3 Destination Unreachable"
        icmp_options {
            code = "3"
            type = "3"
        }
        protocol = "1"
        source = "10.0.0.0/16"
        source_type = "CIDR_BLOCK"
        stateless = "false"
    }
    ingress_security_rules {
        protocol = "all"
        source = "0.0.0.0/0"
        source_type = "CIDR_BLOCK"
        stateless = "false"
    }
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource "oci_core_default_security_list" "default_security_list_ocvp" {
    display_name = "Default Security List for OCVP"
    egress_security_rules {
        destination = "0.0.0.0/0"
        destination_type = "CIDR_BLOCK"
        protocol = "all"
        stateless = "false"
    }
    ingress_security_rules {
        protocol = "6"
        source = "0.0.0.0/0"
        source_type = "CIDR_BLOCK"
        stateless = "false"
        tcp_options {
            max = "22"
            min = "22"
        }
    }
    ingress_security_rules {
        icmp_options {
            code = "4"
            type = "3"
        }
        protocol = "1"
        source = "0.0.0.0/0"
        source_type = "CIDR_BLOCK"
        stateless = "false"
    }
    ingress_security_rules {
        icmp_options {
            code = "-1"
            type = "3"
        }
        protocol = "1"
        source = "10.0.0.0/16"
        source_type = "CIDR_BLOCK"
        stateless = "false"
    }
    manage_default_resource_id = "${oci_core_vcn.test_vcn_ocvp.default_security_list_id}"
}

resource "oci_core_subnet" "test_provisioning_subnet" {
    cidr_block = "10.0.103.128/25"
    compartment_id = "${var.compartment_id}"
    dhcp_options_id = "${oci_core_vcn.test_vcn_ocvp.default_dhcp_options_id}"
    display_name = "provisioning-subnet"
    dns_label = "provisioningsub"
    prohibit_public_ip_on_vnic = "true"
    route_table_id = "${oci_core_route_table.private_rt.id}"
    security_list_ids = ["${oci_core_security_list.private_sl.id}"]
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource "oci_core_vlan" "test_nsx_edge_uplink2_vlan" {
    display_name = "NSX-Edge-UP2"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.103.0/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_nsx_edge_uplink1_vlan" {
    display_name = "NSX-Edge-UP1"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.100.0/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_nsx_vtep_vlan" {
    display_name = "NSX-vTep"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.101.0/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_nsx_edge_vtep_vlan" {
    display_name = "NSX Edge-vTep"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.102.0/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_vsan_net_vlan" {
    display_name = "vSAN-Net"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.101.128/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_vmotion_net_vlan" {
    display_name = "vMotion-Net"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.102.128/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_vsphere_net_vlan" {
  display_name        = "vSphere-Net"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.100.128/26"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id = oci_core_route_table.test_route_table_for_vsphere_vlan.id
}

resource "oci_core_vlan" "test_hcx_vlan" {
  display_name        = "hcx"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.100.192/26"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_provisioning_vlan" {
  display_name        = "provisioning-vlan"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.104.128/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_replication_vlan" {
  display_name        = "replication-vlan"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.104.0/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_volume" "test_volume" {
  display_name		  = "test_volume_management"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  compartment_id      = var.compartment_id
  vpus_per_gb		  = 10
  size_in_gbs         = 4096
}
`
	OcvpSddcCapacityReservationResource = `
resource "oci_core_compute_capacity_reservation" "test_compute_capacity_reservation" {
  compartment_id = var.compartment_id
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  display_name   = "tf-esxi-host-test-capacity-reservation"
  instance_reservation_configs {
    instance_shape = "BM.Standard2.52"
    reserved_count = 1
    fault_domain = "FAULT-DOMAIN-1"
  }
  instance_reservation_configs {
    instance_shape = "BM.Standard2.52"
    reserved_count = 1
    fault_domain = "FAULT-DOMAIN-2"
  }
  instance_reservation_configs {
    instance_shape = "BM.Standard2.52"
    reserved_count = 1
    fault_domain = "FAULT-DOMAIN-3"
  }
}
`

	OcvpSddcOptionalResourceDependencies = OcvpSddcResourceDependencies + OcvpSddcCapacityReservationResource
)

// issue-routing-tag: ocvp/default
func TestOcvpSddcResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpSddcResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ocvp_sddc.test_sddc"
	datasourceName := "data.oci_ocvp_sddcs.test_sddcs"
	singularDatasourceName := "data.oci_ocvp_sddc.test_sddc"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OcvpSddcResourceDependencies+OcvpSddcCapacityReservationResource+
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Optional, acctest.Create, OcvpSddcRepresentation), "ocvp", "sddc", t)

	acctest.ResourceTest(t, testAccCheckOcvpSddcDestroy, []resource.TestStep{
		//  verify Create
		{
			Config: config + compartmentIdVariableStr + OcvpSddcResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Create, OcvpSddcRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV6),
				resource.TestCheckResourceAttr(resourceName, "ssh_authorized_keys", sshKey),
				resource.TestCheckResourceAttr(resourceName, "hcx_mode", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.vsphere_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsphere_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enterprise_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_pending_downgrade", "false"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Upgrade VMware version
		{
			Config: config + compartmentIdVariableStr + OcvpSddcResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Update, sddcUpgradedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttr(resourceName, "ssh_authorized_keys", sshKeyUpdated),
				resource.TestCheckResourceAttr(resourceName, "hcx_mode", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.vsphere_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsphere_vlan_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource for Upgrade
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_sddcs", "test_sddcs", acctest.Required, acctest.Update, OcvpOcvpSddcDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpSddcUpgradeResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.vmware_software_version"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.clusters_count", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.freeform_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.defined_tags.%"),
			),
		},

		// verify singular datasource for Upgrade
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Update, OcvpOcvpSddcSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpSddcUpgradeResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttr(resourceName, "ssh_authorized_keys", sshKeyUpdated),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "clusters_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcenter_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcenter_private_ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_private_ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_edge_uplink_ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcenter_username"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_username"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hcx_mode"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_hcx_pending_downgrade"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_single_host_sddc"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.vsphere_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsphere_vlan_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OcvpSddcResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OcvpSddcOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Optional, acctest.Create, sddcV7Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "clusters_count"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", sddcDisplayName1),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttr(resourceName, "ssh_authorized_keys", sshKey),
				resource.TestCheckResourceAttr(resourceName, "esxi_software_version", esxiSoftwareVersionUpdated),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hcx_mode", "ENTERPRISE"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.0.block_volume_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.0.datastore_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_commitment", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_host_ocpu_count", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_host_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.instance_display_name_prefix", "tf-test-"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.is_shielded_instance_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.replication_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsphere_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.vsphere_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.workload_network_cidr", "172.20.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.actual_esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "is_single_host_sddc", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "hcx_on_prem_licenses.#"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_pending_downgrade", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "hcx_action", ocvp.UpgradeHcxAction),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OcvpSddcOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(sddcV7Representation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "clusters_count"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", sddcDisplayName1),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttr(resourceName, "ssh_authorized_keys", sshKey),
				resource.TestCheckResourceAttr(resourceName, "esxi_software_version", esxiSoftwareVersionUpdated),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hcx_mode", "ENTERPRISE"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.0.block_volume_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.0.datastore_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_commitment", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_host_ocpu_count", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_host_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.instance_display_name_prefix", "tf-test-"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.is_shielded_instance_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.replication_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsphere_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.vsphere_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.workload_network_cidr", "172.20.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "is_single_host_sddc", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "hcx_on_prem_licenses.#"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_pending_downgrade", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "hcx_action", ocvp.UpgradeHcxAction),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		// Cannot Update VMware version here because some of the optional arguments are not applicable to VMware version less than 7.0
		{
			Config: config + compartmentIdVariableStr + OcvpSddcOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Optional, acctest.Update, sddcV7Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "clusters_count"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", sddcDisplayName2),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttr(resourceName, "ssh_authorized_keys", sshKeyUpdated),
				resource.TestCheckResourceAttr(resourceName, "esxi_software_version", esxiSoftwareVersionUpdated),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hcx_mode", "ENTERPRISE"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.0.block_volume_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.0.datastore_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_commitment", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_host_ocpu_count", "12"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_host_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.instance_display_name_prefix", "tf-test-"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.is_shielded_instance_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.provisioning_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.replication_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vmotion_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.0.vsphere_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.vsphere_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "initial_configuration.0.initial_cluster_configurations.0.workload_network_cidr", "172.20.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "is_single_host_sddc", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_private_ip_id"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_pending_downgrade", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "hcx_action", ocvp.UpgradeHcxAction),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config + compartmentIdVariableStr + SddcV7ResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_sddcs", "test_sddcs", acctest.Optional, acctest.Update, OcvpOcvpSddcDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.vmware_software_version"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.clusters_count", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.freeform_tags.%"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.defined_tags.%"),
			),
		},

		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + SddcV7ResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Create, OcvpOcvpSddcSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "clusters_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", sddcDisplayName2),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_authorized_keys", sshKeyUpdated),
				resource.TestCheckResourceAttr(singularDatasourceName, "esxi_software_version", esxiSoftwareVersionUpdated),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hcx_mode", "ENTERPRISE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hcx_on_prem_licenses.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.compute_availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.0.block_volume_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.datastores.0.datastore_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_commitment", "HOUR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.initial_host_ocpu_count", "12"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.instance_display_name_prefix", "tf-test-"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.is_shielded_instance_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.vsphere_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_configuration.0.initial_cluster_configurations.0.workload_network_cidr", "172.20.0.0/24"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_hcx_pending_downgrade"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_single_host_sddc", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_edge_uplink_ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_private_ip_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_username"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_overlay_segment_name"),
			),
		},
		//  verify resource import
		{
			Config:                  config + OcvpSddcRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"hcx_action", "refresh_hcx_license_status"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpSddcDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).SddcClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_sddc" {
			noResourceFound = false
			request := oci_ocvp.GetSddcRequest{}

			tmp := rs.Primary.ID
			request.SddcId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetSddc(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.LifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("OcvpSddc") {
		resource.AddTestSweepers("OcvpSddc", &resource.Sweeper{
			Name:         "OcvpSddc",
			Dependencies: acctest.DependencyGraph["sddc"],
			F:            sweepOcvpSddcResource,
		})
	}
}

func sweepOcvpSddcResource(compartment string) error {
	sddcClient := acctest.GetTestClients(&schema.ResourceData{}).SddcClient()
	sddcIds, err := getOcvpSddcIds(compartment)
	if err != nil {
		return err
	}
	for _, sddcId := range sddcIds {
		if ok := acctest.SweeperDefaultResourceId[sddcId]; !ok {
			deleteSddcRequest := oci_ocvp.DeleteSddcRequest{}

			deleteSddcRequest.SddcId = &sddcId

			deleteSddcRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := sddcClient.DeleteSddc(context.Background(), deleteSddcRequest)
			if error != nil {
				fmt.Printf("Error deleting Sddc %s %s, It is possible that the resource is already deleted. Please verify manually \n", sddcId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sddcId, OcvpSddcSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpSddcSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpSddcIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SddcId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	sddcClient := acctest.GetTestClients(&schema.ResourceData{}).SddcClient()

	listSddcsRequest := oci_ocvp.ListSddcsRequest{}
	listSddcsRequest.CompartmentId = &compartmentId
	listSddcsRequest.LifecycleState = oci_ocvp.ListSddcsLifecycleStateActive
	listSddcsResponse, err := sddcClient.ListSddcs(context.Background(), listSddcsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Sddc list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sddc := range listSddcsResponse.Items {
		id := *sddc.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SddcId", id)
	}
	return resourceIds, nil
}

func OcvpSddcSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sddcResponse, ok := response.Response.(oci_ocvp.GetSddcResponse); ok {
		return sddcResponse.LifecycleState != oci_ocvp.LifecycleStatesDeleted
	}
	return false
}

func OcvpSddcSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.SddcClient().GetSddc(context.Background(), oci_ocvp.GetSddcRequest{
		SddcId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
