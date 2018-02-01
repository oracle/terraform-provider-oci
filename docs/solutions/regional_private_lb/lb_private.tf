
# Create private loadbalancer in each of the two LB subnets

resource "oci_load_balancer" "lb1" {
  shape          = "${var.lb_shape}"
  compartment_id = "${var.compartment_ocid}"
  subnet_ids     = [
    "${oci_core_subnet.LBSubnet1.id}"
  ]
  display_name = "lb1"
  is_private = true
}

resource "oci_load_balancer_backendset" "lb1-bes1" {
  name             = "lb1-bes1"
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  policy           = "ROUND_ROBIN"

  health_checker {
    port     = "${var.backend_port}"
    protocol = "${var.backend_protocol}"
    response_body_regex = ".*"
    url_path = "/"
  }

  session_persistence_configuration {
    cookie_name = "SESSION-COOKIE"
    disable_fallback = true
  }
}

resource "oci_load_balancer_listener" "lb1-listener1" {
  load_balancer_id         = "${oci_load_balancer.lb1.id}"
  name                     = "http"
  default_backend_set_name = "${oci_load_balancer_backendset.lb1-bes1.id}"
  port                     = "${var.ha_app_port}"
  protocol                 = "${upper(var.ha_app_protocol)}"
}

resource "oci_load_balancer" "lb2" {
  shape          = "${var.lb_shape}"
  compartment_id = "${var.compartment_ocid}"
  subnet_ids     = [
    "${oci_core_subnet.LBSubnet2.id}"
  ]
  display_name = "lb2"
  is_private = true
}

resource "oci_load_balancer_backendset" "lb2-bes1" {
  name             = "lb2-bes1"
  load_balancer_id = "${oci_load_balancer.lb2.id}"
  policy           = "ROUND_ROBIN"

  health_checker {
    port     = "${var.backend_port}"
    protocol = "${var.backend_protocol}"
    response_body_regex = ".*"
    url_path = "/"
  }

  session_persistence_configuration {
    cookie_name = "SESSION-COOKIE"
    disable_fallback = true
  }
}

resource "oci_load_balancer_listener" "lb2-listener1" {
  load_balancer_id         = "${oci_load_balancer.lb2.id}"
  name                     = "http"
  default_backend_set_name = "${oci_load_balancer_backendset.lb2-bes1.id}"
  port                     = "${var.ha_app_port}"
  protocol                 = "${upper(var.ha_app_protocol)}"
}

