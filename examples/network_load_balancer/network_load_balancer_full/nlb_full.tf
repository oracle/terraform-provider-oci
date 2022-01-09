// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

variable "availability_domain" {
  default = 3
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

resource "oci_core_public_ip" "test_reserved_ip" {
  compartment_id = "${var.compartment_ocid}"
  lifetime       = "RESERVED"

  lifecycle {
    ignore_changes = [private_ip_id]
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

resource "oci_network_load_balancer_backend_set" "nlb-bes1" {
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
}

resource "oci_network_load_balancer_backend_set" "nlb-bes2" {
  name                     = "nlb-bes2"
  network_load_balancer_id = oci_network_load_balancer_network_load_balancer.nlb1.id
  policy                   = "THREE_TUPLE"

  health_checker {
    port                = "443"
    protocol            = "HTTPS"
    url_path            = "/testPath"
    return_code         = 200
    response_body_regex = "^(?i)(true)$"
    timeout_in_millis   = 10000
    interval_in_millis  = 10000
    retries             = 3
  }
}

resource "oci_network_load_balancer_backend_set" "nlb-bes3" {
  name                     = "nlb-bes3"
  network_load_balancer_id = oci_network_load_balancer_network_load_balancer.nlb1.id
  policy                   = "THREE_TUPLE"

  health_checker {
    port                = "443"
    protocol            = "HTTPS"
    url_path            = "/testPath"
    return_code         = 200
    response_body_regex = "^(?i)(true)$"
    timeout_in_millis   = 10000
    interval_in_millis  = 10000
    retries             = 3
  }
}

resource "oci_network_load_balancer_listener" "nlb-listener1" {
  network_load_balancer_id    = oci_network_load_balancer_network_load_balancer.nlb1.id
  name                        = "tcp_listener"
  default_backend_set_name    = oci_network_load_balancer_backend_set.nlb-bes1.name
  port                        = 80
  protocol                    = "TCP"
}

resource "oci_network_load_balancer_listener" "nlb-listener2" {
  network_load_balancer_id    = oci_network_load_balancer_network_load_balancer.nlb1.id
  name                        = "udp_listener"
  default_backend_set_name    = oci_network_load_balancer_backend_set.nlb-bes2.name
  port                        = 80
  protocol                    = "UDP"
}

resource "oci_network_load_balancer_listener" "nlb-listener3" {
  network_load_balancer_id    = oci_network_load_balancer_network_load_balancer.nlb1.id
  name                        = "tcp_and_udp_listener"
  default_backend_set_name    = oci_network_load_balancer_backend_set.nlb-bes3.name
  port                        = 80
  protocol                    = "TCP_AND_UDP"
}

resource "oci_network_load_balancer_backend" "nlb-be1" {
  network_load_balancer_id = oci_network_load_balancer_network_load_balancer.nlb1.id
  backend_set_name         = oci_network_load_balancer_backend_set.nlb-bes1.name
  ip_address               = "10.0.0.3"
  port                     = 80
  is_backup                = false
  is_drain                 = false
  is_offline               = false
  weight                   = 1
}

resource "oci_network_load_balancer_backend" "nlb-be2" {
  network_load_balancer_id = oci_network_load_balancer_network_load_balancer.nlb1.id
  backend_set_name         = oci_network_load_balancer_backend_set.nlb-bes2.name
  target_id                = oci_core_instance.instance1.id
  port                     = 20
  is_backup                = false
  is_drain                 = false
  is_offline               = false
  weight                   = 1
}