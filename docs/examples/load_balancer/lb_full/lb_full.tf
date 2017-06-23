/*
 * This example demonstrates round robin load balancing behavior by creating two instances, a configured
 * vcn and a load balancer. The public IP of the load balancer is outputted after a successful run, curl
 * this address to see the hostname change as different instances handle the request.
 *
 * NOTE: The https listener is included for completeness but should not be expected to work,
 * it uses dummy certs.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}


provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

data "baremetal_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}


/* Network */

resource "baremetal_core_virtual_network" "vcn1" {
    cidr_block = "10.1.0.0/16"
    compartment_id = "${var.compartment_ocid}"
    display_name = "vcn1"
    dns_label = "vcn1"
}

resource "baremetal_core_subnet" "subnet1" {
    availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.1.20.0/24"
    display_name = "subnet1"
    dns_label = "subnet1"
    security_list_ids = ["${baremetal_core_security_list.securitylist1.id}"]
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${baremetal_core_virtual_network.vcn1.id}"
    route_table_id = "${baremetal_core_route_table.routetable1.id}"

    provisioner "local-exec" {
        command = "sleep 5"
    }
}

resource "baremetal_core_subnet" "subnet2" {
    availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[1],"name")}"
    cidr_block = "10.1.21.0/24"
    display_name = "subnet2"
    dns_label = "subnet2"
    security_list_ids = ["${baremetal_core_security_list.securitylist1.id}"]
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${baremetal_core_virtual_network.vcn1.id}"
    route_table_id = "${baremetal_core_route_table.routetable1.id}"

    provisioner "local-exec" {
        command = "sleep 5"
    }
}

resource "baremetal_core_internet_gateway" "internetgateway1" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "internetgateway1"
    vcn_id = "${baremetal_core_virtual_network.vcn1.id}"
}

resource "baremetal_core_route_table" "routetable1" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${baremetal_core_virtual_network.vcn1.id}"
    display_name = "routetable1"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${baremetal_core_internet_gateway.internetgateway1.id}"
    }
}

resource "baremetal_core_security_list" "securitylist1" {
  display_name   = "public"
  compartment_id = "${baremetal_core_virtual_network.vcn1.compartment_id}"
  vcn_id         = "${baremetal_core_virtual_network.vcn1.id}"

  egress_security_rules = [{
    protocol    = "all"
    destination = "0.0.0.0/0"
  }]

  ingress_security_rules = [
    {
      protocol = "6"
      source   = "0.0.0.0/0"

      tcp_options {
        "min" = 80
        "max" = 80
      }
    },
    {
      protocol = "6"
      source   = "0.0.0.0/0"

      tcp_options {
        "min" = 443
        "max" = 443
      }
    },
  ]
}


/* Instances */

data "baremetal_core_images" "image-list" {
  compartment_id = "${var.compartment_ocid}"
  operating_system = "Oracle Linux"
  operating_system_version = "7.3"
}

resource "baremetal_core_instance" "instance1" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}" 
  compartment_id = "${var.compartment_ocid}"
  display_name = "be-instance1"
  image = "${lookup(data.baremetal_core_images.image-list.images[0], "id")}"
  shape = "VM.Standard1.2"
  subnet_id = "${baremetal_core_subnet.subnet1.id}"
  hostname_label = "be-instance1"
  metadata {
    user_data = "${base64encode(var.user-data)}"
  }
}

resource "baremetal_core_instance" "instance2" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[1],"name")}" 
  compartment_id = "${var.compartment_ocid}"
  display_name = "be-instance2"
  image = "${lookup(data.baremetal_core_images.image-list.images[0], "id")}"
  shape = "VM.Standard1.2"
  subnet_id = "${baremetal_core_subnet.subnet2.id}"
  hostname_label = "be-instance2"
  metadata {
    user_data = "${base64encode(var.user-data)}"
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


/* Load Balancer */

resource "baremetal_load_balancer" "lb1" {
  shape          = "100Mbps"
  compartment_id = "${var.compartment_ocid}"
  subnet_ids     = [
    "${baremetal_core_subnet.subnet1.id}",
    "${baremetal_core_subnet.subnet2.id}"
  ]
  display_name   = "lb1"
}

resource "baremetal_load_balancer_backendset" "lb-bes1" {
  name             = "lb-bes1"
  load_balancer_id = "${baremetal_load_balancer.lb1.id}"
  policy           = "ROUND_ROBIN"

  health_checker {
    port     = "80"
    protocol = "HTTP"
    response_body_regex = ".*"
    url_path = "/"
  }
}

resource "baremetal_load_balancer_certificate" "lb-cert1" {
  load_balancer_id   = "${baremetal_load_balancer.lb1.id}"
  ca_certificate     = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
  certificate_name   = "certificate1"
  private_key        = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"
  public_certificate = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
}

resource "baremetal_load_balancer_listener" "lb-listener1" {
  load_balancer_id         = "${baremetal_load_balancer.lb1.id}"
  name                     = "http"
  default_backend_set_name = "${baremetal_load_balancer_backendset.lb-bes1.id}"
  port                     = 80
  protocol                 = "HTTP"
}

resource "baremetal_load_balancer_listener" "lb-listener2" {
  load_balancer_id         = "${baremetal_load_balancer.lb1.id}"
  name                     = "https"
  default_backend_set_name = "${baremetal_load_balancer_backendset.lb-bes1.id}"
  port                     = 443
  protocol                 = "HTTP"

  ssl_configuration {
    certificate_name        = "${baremetal_load_balancer_certificate.lb-cert1.certificate_name}"
    verify_peer_certificate = false
  }
}

resource "baremetal_load_balancer_backend" "lb-be1" {
  load_balancer_id = "${baremetal_load_balancer.lb1.id}"
  backendset_name  = "${baremetal_load_balancer_backendset.lb-bes1.id}"
  ip_address       = "${baremetal_core_instance.instance1.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

resource "baremetal_load_balancer_backend" "lb-be2" {
  load_balancer_id = "${baremetal_load_balancer.lb1.id}"
  backendset_name  = "${baremetal_load_balancer_backendset.lb-bes1.id}"
  ip_address       = "${baremetal_core_instance.instance2.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}


output "lb_public_ip" {
  value = ["${baremetal_load_balancer.lb1.ip_addresses}"]
}
