
# Launch and configure DNS VMs

resource "oci_core_instance" "DnsVM1" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "DnsVM1"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "${var.InstanceShape}"
    create_vnic_details {
        subnet_id = "${oci_core_subnet.MgmtSubnet1.id}"
    }
    metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
    }
    timeouts {
        create = "10m"
    }
}

resource "oci_core_instance" "DnsVM2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "DnsVM2"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "${var.InstanceShape}"
    create_vnic_details {
        subnet_id = "${oci_core_subnet.MgmtSubnet2.id}"
    }
    metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
    }
    timeouts {
        create = "10m"
    }
}

# Gets a list of VNIC attachments on the DNS instance
data "oci_core_vnic_attachments" "DnsVM1Vnics" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    instance_id = "${oci_core_instance.DnsVM1.id}"
}

data "oci_core_vnic_attachments" "DnsVM2Vnics" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    instance_id = "${oci_core_instance.DnsVM2.id}"
}

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "DnsVMVnic" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.DnsVM1Vnics.vnic_attachments[0],"vnic_id")}"
}

data "oci_core_vnic" "DnsVMVnic2" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.DnsVM2Vnics.vnic_attachments[0],"vnic_id")}"
}

# Update the default DHCP options to use custom DNS servers
resource "oci_core_default_dhcp_options" "default-dhcp-options" {
    manage_default_resource_id = "${oci_core_virtual_network.MgmtVcn.default_dhcp_options_id}"

    options {
        type = "DomainNameServer"
        server_type = "CustomDnsServer"
        custom_dns_servers = [  "${data.oci_core_vnic.DnsVMVnic.private_ip_address}",
                                "${data.oci_core_vnic.DnsVMVnic2.private_ip_address}" ]
    }

    options {
        type = "SearchDomain"
        search_domain_names = [ "${oci_core_virtual_network.MgmtVcn.dns_label}.oraclevcn.com" ]
    }
}

data "template_file" "generate_named_conf" {
    template = "${file("named.conf.tpl")}"
 
    vars {
      vcn_cidr           = "${var.vcn_cidr}"
      zone               = "${var.ha_app_domain}"
      onprem_cidr        = "${var.onprem_cidr}"
      onprem_domain      = "${var.onprem_domain}"
      onprem_dns_server1 = "${var.onprem_dns_server1}"
      onprem_dns_server2 = "${var.onprem_dns_server2}"
    }
}

data "template_file" "generate_db_zone_vm1" {
    template = "${file("db.zone.tmpl")}"
 
    vars {
      dns_server_ip = "${data.oci_core_vnic.DnsVMVnic.private_ip_address}"
      app           = "${var.ha_app_name}"
      lb_ip1        = "${oci_load_balancer.lb1.ip_addresses[0]}"
      lb_ip2        = "${oci_load_balancer.lb2.ip_addresses[0]}"
    }
}

data "template_file" "generate_db_zone_vm2" {
    template = "${file("db.zone.tmpl")}"
 
    vars {
      dns_server_ip = "${data.oci_core_vnic.DnsVMVnic2.private_ip_address}"
      app           = "${var.ha_app_name}"
      lb_ip1        = "${oci_load_balancer.lb1.ip_addresses[0]}"
      lb_ip2        = "${oci_load_balancer.lb2.ip_addresses[0]}"
    }
}

data "template_file" "generate_monitrc" {
    template = "${file("monitrc.tmpl")}"
 
    vars {
      lb_ip1      = "${oci_load_balancer.lb1.ip_addresses[0]}"
      lb_ip2      = "${oci_load_balancer.lb2.ip_addresses[0]}"
      lb_port     = "${var.ha_app_port}"
      lb_protocol = "${var.ha_app_protocol}"
      app         = "${var.ha_app_name}"
      zone_file   = "/etc/named/db.${var.ha_app_domain}"
    }
}

resource "null_resource" "configure-bind-vm1" {
  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.ssh_private_key}"
    host        = "${data.oci_core_vnic.DnsVMVnic.public_ip_address}"
    timeout     = "30m"
  }

  provisioner "file" {
    content     = "${data.template_file.generate_named_conf.rendered}"
    destination = "~/named.conf"
  }

  provisioner "file" {
    content     = "${data.template_file.generate_db_zone_vm1.rendered}"
    destination = "~/db.${var.ha_app_domain}"
  }

  provisioner "file" {
    content     = "${data.template_file.generate_monitrc.rendered}"
    destination = "~/monitrc"
  }

  provisioner "file" {
    source      = "dnsops"
    destination = "~/dnsops"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",

      "sudo yum install bind -y",
      "sudo cp ~/named.conf /etc/named.conf",
      "sudo cp ~/db.${var.ha_app_domain} /etc/named/db.${var.ha_app_domain}",
      "sudo systemctl enable named",
      "sudo systemctl restart named",

      "sudo firewall-offline-cmd --add-port=53/udp",
      "sudo firewall-offline-cmd --add-port=53/tcp",
      "sudo /bin/systemctl restart firewalld",

      "sudo yum install monit -y",
      "sudo cp ~/monitrc /etc/monitrc",
      "sudo chmod +x ~/dnsops/*",
      "sudo systemctl enable monit",
      "sudo systemctl restart monit"
    ]
  }
}

resource "null_resource" "configure-bind-vm2" {
  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.ssh_private_key}"
    host        = "${data.oci_core_vnic.DnsVMVnic2.public_ip_address}"
    timeout     = "30m"
  }

  provisioner "file" {
    content     = "${data.template_file.generate_named_conf.rendered}"
    destination = "~/named.conf"
  }

  provisioner "file" {
    content     = "${data.template_file.generate_db_zone_vm2.rendered}"
    destination = "~/db.${var.ha_app_domain}"
  }

  provisioner "file" {
    content     = "${data.template_file.generate_monitrc.rendered}"
    destination = "~/monitrc"
  }

  provisioner "file" {
    source      = "dnsops"
    destination = "~/dnsops"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",

      "sudo yum install bind -y",
      "sudo cp ~/named.conf /etc/named.conf",
      "sudo cp ~/db.${var.ha_app_domain} /etc/named/db.${var.ha_app_domain}",
      "sudo systemctl enable named",
      "sudo systemctl restart named",

      "sudo firewall-offline-cmd --add-port=53/udp",
      "sudo firewall-offline-cmd --add-port=53/tcp",
      "sudo /bin/systemctl restart firewalld",

      "sudo yum install monit -y",
      "sudo cp ~/monitrc /etc/monitrc",
      "sudo chmod +x ~/dnsops/*",
      "sudo systemctl enable monit",
      "sudo systemctl restart monit"
    ]
  }
}

