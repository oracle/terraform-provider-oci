/*
 * This example demonstrates basic load balancer configuration and requires an existing instance and subnets.
 * It should be configured with a proper cert.crt and cert.key for ssl configuration, but dummy certs are
 * included for demonstration purposes.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "instance_ip" {}
variable "subnet1" {}
variable "subnet2" {}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

/* Load Balancer */

resource "baremetal_load_balancer" "lb1" {
  shape          = "100Mbps"
  compartment_id = "${var.compartment_ocid}"
  subnet_ids     = [
    "${var.subnet1}",
    "${var.subnet2}"
  ]
  display_name   = "lb1"
}

resource "baremetal_load_balancer_backendset" "lb-bes1" {
  name             = "lb-bes1"
  load_balancer_id = "${baremetal_load_balancer.lb1.id}"
  policy           = "ROUND_ROBIN"

  health_checker {
    port     = "8080"
    protocol = "HTTP"
    response_body_regex = ".*"
    url_path = "/"
  }
}

resource "baremetal_load_balancer_listener" "lb-listener1" {
  load_balancer_id         = "${baremetal_load_balancer.lb1.id}"
  name                     = "lb-listener1"
  default_backend_set_name = "${baremetal_load_balancer_backendset.lb-bes1.id}"
  port                     = 80
  protocol                 = "HTTP"

  ssl_configuration {
    certificate_name        = "${baremetal_load_balancer_certificate.lb-cert1.certificate_name}"
    verify_peer_certificate = false
  }
}

resource "baremetal_load_balancer_backend" "lb-be1" {
  load_balancer_id = "${baremetal_load_balancer.lb1.id}"
  backendset_name  = "${baremetal_load_balancer_backendset.lb-bes1.id}"
  ip_address       = "${var.instance_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

resource "baremetal_load_balancer_certificate" "lb-cert1" {
  load_balancer_id   = "${baremetal_load_balancer.lb1.id}"
  certificate_name   = "certificate1"
  ca_certificate     = ""
  public_certificate = "${file("${path.module}/cert.crt")}"
  private_key        = "${file("${path.module}/cert.key")}"
}