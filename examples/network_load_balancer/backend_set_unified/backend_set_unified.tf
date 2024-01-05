// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = ""
    us-ashburn-1   = ""
    eu-frankfurt-1 = ""
    uk-london-1    = ""
  }
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

variable "availability_domain" {
  default = 3
}


locals {
  ######################################################################################################################
  # IPV6 Constants
  ######################################################################################################################
  ipv6_cidr_block = oci_core_vcn.vcn1.ipv6cidr_blocks[0] // this ends in 0::/56
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_identity_availability_domain" "ad1" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

/* Network */

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "vcn1"
  dns_label      = "vcn1"
  is_ipv6enabled =  true
}

resource "oci_core_subnet" "subnet1" {
  cidr_block          = "10.1.20.0/24"
  display_name        = "subnet1"
  dns_label           = "subnet1"
  security_list_ids   = [oci_core_security_list.securitylist1.id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.vcn1.id
  route_table_id      = oci_core_route_table.routetable1.id
  dhcp_options_id     = oci_core_vcn.vcn1.default_dhcp_options_id

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_internet_gateway" "internetgateway1" {
  compartment_id = var.compartment_ocid
  display_name   = "internetgateway1"
  vcn_id         = oci_core_vcn.vcn1.id
}

resource "oci_core_route_table" "routetable1" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.vcn1.id
  display_name   = "routetable1"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.internetgateway1.id
  }
}

resource "oci_core_security_list" "securitylist1" {
  display_name   = "public"
  compartment_id = oci_core_vcn.vcn1.compartment_id
  vcn_id         = oci_core_vcn.vcn1.id

  egress_security_rules {
    protocol    = "all"
    destination = "0.0.0.0/0"
  }

  ingress_security_rules {
    protocol = "6"
    source   = "0.0.0.0/0"

    tcp_options {
      min = 80
      max = 80
    }
  }

  ingress_security_rules {
    protocol = "6"
    source   = "0.0.0.0/0"

    tcp_options {
      min = 443
      max = 443
    }
  }
}

/* Instances */

resource "oci_core_instance" "instance1" {
  availability_domain = data.oci_identity_availability_domain.ad1.name
  compartment_id      = var.compartment_ocid
  display_name        = "be-instance1"
  shape               = var.instance_shape

  metadata = {
    user_data = base64encode(var.user-data)
  }

  create_vnic_details {
    subnet_id      = oci_core_subnet.subnet1.id
    hostname_label = "be-instance1"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }
}

variable "user-data" {
  default = <<EOF
#!/bin/bash -x
echo '################### webserver userdata begins #####################'
touch ~opc/userdata.`date +%s`.start

# echo '########## yum update all ###############'
# yum update -y

echo '########## basic webserver ##############'
yum install -y httpd
systemctl enable  httpd.service
systemctl start  httpd.service
echo '<html><head></head><body><pre><code>' > /var/www/html/index.html
hostname >> /var/www/html/index.html
echo '' >> /var/www/html/index.html
cat /etc/os-release >> /var/www/html/index.html
echo '</code></pre></body></html>' >> /var/www/html/index.html
firewall-offline-cmd --add-service=http
systemctl enable  firewalld
systemctl restart  firewalld

touch ~opc/userdata.`date +%s`.finish
echo '################### webserver userdata ends #######################'
EOF

}

/* Network Load Balancer */

resource "oci_network_load_balancer_network_load_balancer" "nlb1" {
  compartment_id = var.compartment_ocid

  subnet_id = oci_core_subnet.subnet1.id

  display_name = "nlb1"
}

resource "oci_network_load_balancer_network_load_balancers_backend_sets_unified" "nlb-bes1" {
  name                     = "nlb-bes1"
  network_load_balancer_id = oci_network_load_balancer_network_load_balancer.nlb1.id
  policy                   = "TWO_TUPLE"

  health_checker {
    port                = "80"
    protocol            = "TCP"
    request_data        = "SGVsbG9Xb3JsZA=="
    response_data       = "SGVsbG9Xb3JsZA=="
    timeout_in_millis   = 10000
    interval_in_millis  = 10000
    retries             = 3
  }

  backends {
      ip_address               = "10.0.0.3"
      port                     = 80
      is_backup                = false
      is_drain                 = false
      is_offline               = false
      weight                   = 1
  }

   backends {
      target_id                = oci_core_instance.instance1.id
      port                     = 20
      is_backup                = false
      is_drain                 = false
      is_offline               = false
      weight                   = 1
   }
  depends_on = [oci_network_load_balancer_network_load_balancer.nlb1]
}

resource "oci_network_load_balancer_listener" "nlb-listener1" {
  network_load_balancer_id    = oci_network_load_balancer_network_load_balancer.nlb1.id
  name                        = "tcp_listener"
  default_backend_set_name    = oci_network_load_balancer_network_load_balancers_backend_sets_unified.nlb-bes1.name
  port                        = 80
  protocol                    = "TCP"

  depends_on = [oci_network_load_balancer_network_load_balancers_backend_sets_unified.nlb-bes1]
}