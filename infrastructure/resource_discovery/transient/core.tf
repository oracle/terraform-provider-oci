// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Create Instance
resource "oci_core_instance" "instance_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "instanceRD"
  shape               = "VM.Standard2.1"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.subnet_core_rd.id}"
    hostname_label   = "instance"
    assign_public_ip = true
    display_name     = "Primaryvnic"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }
}

/* instance console connection*/
resource "oci_core_instance_console_connection" "instance_console_connection_rd" {
  #Required
  instance_id = "${oci_core_instance.instance_rd.id}"
  public_key  = "${var.ssh_public_key}"
}

/* instance configuration */
resource "oci_core_instance_configuration" "instance_configuration_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "instanceConfigurationRD"

  instance_details {
    instance_type = "compute"

    launch_details {
      compartment_id                      = "${var.compartment_ocid}"
      ipxe_script                         = "ipxeScript"
      shape                               = "${var.instance_shape}"
      display_name                        = "TestInstanceConfigurationLaunchDetails"
      is_pv_encryption_in_transit_enabled = false
      preferred_maintenance_action        = "LIVE_MIGRATE"
      launch_mode                         = "NATIVE"

      agent_config {
        is_management_disabled = false
        is_monitoring_disabled = false
      }

      launch_options {
        network_type = "PARAVIRTUALIZED"
      }

      shape_config {
        ocpus = 1
      }

      create_vnic_details {
        assign_public_ip       = true
        display_name           = "TestInstanceConfigurationVNIC"
        skip_source_dest_check = false
      }

      extended_metadata = {
        some_string   = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
      }

      source_details {
        source_type = "image"
        image_id    = "${oci_core_image.image_rd.id}"
      }
    }
  }
}

/* instance pool */
resource "oci_core_instance_pool" "test_instance_pool_rd" {
  compartment_id            = "${var.compartment_ocid}"
  instance_configuration_id = "${oci_core_instance_configuration.instance_configuration_rd.id}"
  size                      = 2
  state                     = "RUNNING"
  display_name              = "testInstancePoolRD"

  placement_configurations {
    availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
    fault_domains       = ["FAULT-DOMAIN-1"]
    primary_subnet_id   = "${oci_core_subnet.subnet_core_rd.id}"
  }

  load_balancers {
    backend_set_name = "${oci_load_balancer_backend_set.lb_backend_set_rd.name}"
    load_balancer_id = "${oci_load_balancer.load_balancer_rd.id}"
    port             = 80
    vnic_selection   = "primaryvnic"
  }
}

/* image */
resource "oci_core_image" "image_rd" {
  compartment_id = "${var.compartment_ocid}"
  instance_id    = "${oci_core_instance.instance_rd.id}"
  display_name   = "imageRD"

  launch_mode = "NATIVE"

  timeouts {
    create = "30m"
  }
}

/* volume */
resource "oci_core_volume" "volume_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "volumeRD"
}

/* volume attachment */
resource "oci_core_volume_attachment" "block_volume_attach_paravirtualized_rd" {
  count           = "1"
  attachment_type = "paravirtualized"
  instance_id     = "${oci_core_instance.instance_rd.*.id[0]}"
  volume_id       = "${oci_core_volume.volume_rd.*.id[0]}"

  # Set this to attach the volume as read-only.
  #is_read_only = true
}

/* volume backup */
resource "oci_core_volume_backup" "volume_backup_rd" {
  #Required
  volume_id = "${oci_core_volume.volume_rd.id}"

  #Optional
  display_name  = "volumeBackupRD"
  freeform_tags = "${var.freeform_tags}"
  type          = "${var.volume_backup_type}"
}

/* volume_backup_policy */
resource "oci_core_volume_backup_policy" "volume_backup_policy_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
  display_name = "BackupPolicy1"

  freeform_tags = {
    "Department" = "Finance"
  }

  schedules {
    #Required
    backup_type       = "INCREMENTAL"
    period            = "ONE_YEAR"
    retention_seconds = "604800"

    #Optional
    day_of_month   = "10"
    day_of_week    = "TUESDAY"
    hour_of_day    = "10"
    month          = "FEBRUARY"
    offset_seconds = "0"
    offset_type    = "STRUCTURED"
    time_zone      = "UTC"
  }
}

/* volume_backup_policy_assignment */
resource "oci_core_volume_backup_policy_assignment" "volume_backup_policy_rd" {
  count     = 2
  asset_id  = "${oci_core_instance.instance_rd.*.boot_volume_id[0]}"
  policy_id = "${data.oci_core_volume_backup_policies.test_predefined_volume_backup_policies.volume_backup_policies.0.id}"
}

data "oci_core_volume_backup_policies" "test_predefined_volume_backup_policies" {
  filter {
    name   = "display_name"
    values = ["silver"]
  }
}

/* volume_group */
resource "oci_core_volume" "volume2_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "-tf-volume"
  size_in_gbs         = "50"
}

resource "oci_core_volume_group" "test_volume_group_from_vol_ids_rd" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"

  source_details {
    #Required
    type = "volumeIds"

    // Mix of named volume and splatted multiple volumes
    // If using with Terraform v0.12, supply the following line instead.
    // volume_ids = concat(["${oci_core_volume.t.id}"], "${oci_core_volume.test_volume.*.id}")
    volume_ids = ["${oci_core_volume.volume2_rd.id}", "${oci_core_volume.volume_rd.*.id}"]
  }

  #Optional
  display_name = "testVolumeGroupFromVolIdsRD"
}

/* volume_group_backup */
resource "oci_core_volume_group_backup" "test_volume_group_backup" {
  #Required
  volume_group_id = "${oci_core_volume_group.test_volume_group_from_vol_ids_rd.id}"

  #Optional
  display_name = "tf-volume-group-backup"
  type         = "INCREMENTAL"
}

/* vnic attachment */
data "oci_core_vnic" "instance_vnic2" {
  vnic_id = "${oci_core_vnic_attachment.vnic_attachment_rd.vnic_id}"
}

resource "oci_core_vnic_attachment" "vnic_attachment_rd" {
  instance_id  = "${oci_core_instance.instance_core_vnic_attachment.id}"
  display_name = "vnicAttachmentRD"

  create_vnic_details {
    assign_public_ip       = false
    display_name           = "TFSecondaryVnic"
    skip_source_dest_check = true
    subnet_id              = "${oci_core_subnet.subnet_core_rd.id}"
  }
}

resource "oci_core_instance" "instance_core_vnic_attachment" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "instanceCoreVnicAttachment"
  shape               = "VM.Standard2.1"

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  create_vnic_details {
    subnet_id      = "${oci_core_subnet.subnet_core_rd.id}"
    hostname_label = "testinstance"
  }

  metadata = {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "60m"
  }
}

/* public id */
resource "oci_core_public_ip" "public_ip_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "publicIPAssignedRD"
  lifetime       = "RESERVED"
  private_ip_id  = "${lookup(data.oci_core_private_ips.private_ips2.private_ips[0],"id")}"
}

# Gets a list of private IPs on the second VNIC
data "oci_core_private_ips" "private_ips2" {
  vnic_id = "${data.oci_core_vnic.instance_vnic2.id}"
}

/* private id */
resource "oci_core_private_ip" "private_ip_rd" {
  vnic_id        = "${lookup(data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0],"vnic_id")}"
  display_name   = "privateIpRD"
  hostname_label = "somehostnamelabel"
}

# Gets a list of IPv6s on the second VNIC
data "oci_core_ipv6s" "ipv6s2" {
  vnic_id = "${data.oci_core_vnic.instance_vnic2.id}"
}

/* ipv6 id */
resource "oci_core_ipv6" "ipv6_rd" {
  vnic_id        = "${lookup(data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0],"vnic_id")}"
  display_name   = "ipv6RD"
}

data "oci_core_vnic_attachments" "instance_vnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  instance_id         = "${oci_core_instance.instance_rd.id}"
}

# Gets the OCID of the first (default) VNIC
data "oci_core_vnic" "instance_vnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0],"vnic_id")}"
}

/* boot volume */
resource "oci_core_boot_volume" "boot_volume_from_source_boot_volume_rd" {
  availability_domain = "${oci_core_instance.instance_rd.availability_domain}"
  compartment_id      = "${oci_core_instance.instance_rd.compartment_id}"
  display_name        = "bootVolumeRD"

  source_details {
    #Required
    id   = "${oci_core_instance.instance_rd.boot_volume_id}"
    type = "bootVolume"
  }
}

resource "oci_core_boot_volume_backup" "boot_volume_backup_rd" {
  #Required
  boot_volume_id = "${oci_core_boot_volume.boot_volume_from_source_boot_volume_rd.id}"

  display_name = "bootVolumeBackupRD"

  // source_details: Cannot be defined if boot_volume_id is defined.
  #source_details = <<Optional value not found in discovery>>
}

/*network_security_group*/
resource "oci_core_network_security_group" "core_network_security_group_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn4_rd.id}"

  #Optional
  display_name = "coreNetworkSecurityGroupRD"
}

resource "oci_core_network_security_group_security_rule" "core_network_security_group_security_rule_rd" {
  network_security_group_id = "${oci_core_network_security_group.network_security_group_rd.id}"
  direction                 = "EGRESS"
  destination               = "0.0.0.0/16"
  protocol                  = "7"
  count                     = 5
}

/*dhcp*/
resource "oci_core_dhcp_options" "example_dhcp_options_rd" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn4_rd.id}"
  display_name   = "exampleDhcpOptionsRD"

  // required
  options {
    type               = "DomainNameServer"
    server_type        = "CustomDnsServer"
    custom_dns_servers = ["8.8.4.4", "8.8.8.8"]
  }

  // optional
  options {
    type                = "SearchDomain"
    search_domain_names = ["test.com"]
  }
}

/* internet gateway */
resource "oci_core_internet_gateway" "internet_gateway_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "internetGatewayRD"
  vcn_id         = "${oci_core_vcn.vcn4_rd.id}"
}

resource "oci_core_route_table" "route_table_rd" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn4_rd.id}"
  display_name   = "routeTableRD"

  route_rules {
    description       = "description"
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_local_peering_gateway.local_peering_gateway_rd.id}"
  }
}

/* local peering gateway */
resource "oci_core_local_peering_gateway" "local_peering_gateway1_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"

  #Optional
  display_name = "localPeeringGateway1RD"
}

resource "oci_core_local_peering_gateway" "local_peering_gateway_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn4_rd.id}"

  #Optional
  display_name = "localPeeringGatewayRD"
  peer_id      = "${oci_core_local_peering_gateway.local_peering_gateway1_rd.id}"
}

/* security list */
resource "oci_core_security_list" "security_list_rd" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn4_rd.id}"
  display_name   = "securityListRD"

  // allow outbound tcp traffic on all ports
  egress_security_rules {
    description = "allow outbound tcp traffic on all ports"
    destination = "0.0.0.0/0"
    protocol    = "6"
  }

  // allow outbound udp traffic on a port range
  egress_security_rules {
    description = "allow outbound udp traffic on a port range"
    destination = "0.0.0.0/0"
    protocol    = "17"                                         // udp
    stateless   = true

    udp_options {
      min = 319
      max = 320

      source_port_range {
        min = 100
        max = 100
      }
    }
  }

  // allow inbound ssh traffic
  ingress_security_rules {
    description = "allow inbound ssh traffic"
    protocol    = "6"                         // tcp
    source      = "0.0.0.0/0"
    stateless   = false

    tcp_options {
      min = 22
      max = 22

      source_port_range {
        min = 100
        max = 100
      }
    }
  }

  // allow inbound icmp traffic of a specific type
  ingress_security_rules {
    description = "allow inbound icmp traffic"
    protocol    = 1
    source      = "0.0.0.0/0"
    stateless   = true

    icmp_options {
      type = 3
      code = 4
    }
  }
}

/* ipsec */
resource "oci_core_ipsec" "ip_sec_connection_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  cpe_id         = "${oci_core_cpe.cpe_rd.id}"
  drg_id         = "${oci_core_drg.drg_rd.id}"
  static_routes  = ["10.0.0.0/16"]

  #Optional
  cpe_local_identifier      = "189.44.2.135"
  cpe_local_identifier_type = "IP_ADDRESS"
  defined_tags              = "${map("${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
  display_name              = "iPSecConnectionRD"

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_identity_tag_namespace" "tag_namespace1" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  description    = "Just a test"
  name           = "tagNamespace1"
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tag1"
  tag_namespace_id = "${oci_identity_tag_namespace.tag_namespace1.id}"
}

/* cpe */
resource oci_core_cpe "cpe_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "cpeRD"
  ip_address     = "189.44.2.135"
}

/* drg */
resource oci_core_drg "drg_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "drgRD"
}

resource "oci_core_drg_attachment" "drg_attachment_rd" {
  #Required
  drg_id = "${oci_core_drg.drg_rd.id}"
  vcn_id = "${oci_core_vcn.vcn4_rd.id}"

  #Optional
  display_name   = "drgAttachmentRD"
  route_table_id = "${oci_core_route_table.route_table_rd.id}"
}

/* nat gateway */
resource "oci_core_nat_gateway" "nat_gateway_rd" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn4_rd.id}"
  display_name   = "natGatewayRD"
}

/* service gateway */
resource "oci_core_service_gateway" "service_gateway_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  services {
    service_id = "${lookup(data.oci_core_services.services_rd.services[0], "id")}"
  }

  vcn_id = "${oci_core_vcn.vcn4_rd.id}"

  #Optional
  display_name   = "serviceGatewayRD"
  route_table_id = "${oci_core_route_table.test_route_table_transit_routing.id}"
}

/* services */
data "oci_core_services" "services_rd" {
  filter {
    name   = "name"
    values = ["All .* Services In Oracle Services Network"]
    regex  = true
  }
}

resource "oci_core_route_table" "test_route_table_transit_routing" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn4_rd.id}"
  display_name   = "testRouteTableTransitRouting"
}

/* dedicated_vm_host */
resource "oci_core_dedicated_vm_host" "dedicated_vm_host_rd" {
  #Required
  availability_domain     = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id          = "${var.compartment_ocid}"
  dedicated_vm_host_shape = "DVH.Standard2.52"

  #Optional
  display_name = "dedicatedVMHostRD"
}

/* cross connect */
resource "oci_core_cross_connect" "cross_connect_rd" {
  #Required
  compartment_id        = "${var.compartment_ocid}"
  location_name         = "${data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations.0.name}"
  port_speed_shape_name = "${data.oci_core_cross_connect_port_speed_shapes.cross_connect_port_speed_shapes.cross_connect_port_speed_shapes.0.name}"

  #Optional
  cross_connect_group_id = "${oci_core_cross_connect_group.cross_connect_group_rd.id}"
  display_name           = "crossConnectRD"

  #customer_reference_name
  #far_cross_connect_or_cross_connect_group_id = "${oci_core_far_cross_connect_or_cross_connect_group.far_cross_connect_or_cross_connect_group.id}"
  #near_cross_connect_or_cross_connect_group_id = "${oci_core_near_cross_connect_or_cross_connect_group.near_cross_connect_or_cross_connect_group.id}"

  #Set Cross Connect to Active to provision (required to provision virtual circuits).
  #You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up
  is_active = true
}

data "oci_core_cross_connect_locations" "cross_connect_locations" {
  #Required
  compartment_id = "${var.compartment_ocid}"
}

data "oci_core_cross_connect_port_speed_shapes" "cross_connect_port_speed_shapes" {
  #Required
  compartment_id = "${var.compartment_ocid}"
}

/* cross connect group */
resource "oci_core_cross_connect_group" "cross_connect_group_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  # customer_reference_name

  #Optional
  display_name = "crossConnectGroupRD"
}

/* remote peering connection */
resource "oci_core_remote_peering_connection" "remote_peering_connection_rd" {
  compartment_id = "${var.compartment_ocid}"
  drg_id         = "${oci_core_drg.drg_rd.id}"
  display_name   = "remotePeeringConnectionRequestorRD"
}

resource "oci_core_subnet" "subnet_core_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "10.0.1.0/24"
  display_name        = "subnetCoreRD"
  dns_label           = "subnetRD"

  security_list_ids = [
    "${oci_core_vcn.vcn4_rd.default_security_list_id}",
  ]

  compartment_id  = "${var.compartment_ocid}"
  vcn_id          = "${oci_core_vcn.vcn4_rd.id}"
  route_table_id  = "${oci_core_vcn.vcn4_rd.default_route_table_id}"
  dhcp_options_id = "${oci_core_vcn.vcn4_rd.default_dhcp_options_id}"
}
