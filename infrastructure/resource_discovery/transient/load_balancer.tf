// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_security_list" "securitylist_lb_rd" {
  display_name   = "public"
  compartment_id = "${oci_core_vcn.vcn_rd.compartment_id}"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"

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

resource "oci_core_subnet" "subnet1_lb_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "10.1.22.0/24"
  display_name        = "subnet1lbRD"
  dns_label           = "subnet1lbrd"
  security_list_ids   = ["${oci_core_security_list.securitylist_lb_rd.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn_rd.id}"
  route_table_id      = "${oci_core_route_table.routetable_rd.id}"

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_subnet" "subnet2_lb_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  cidr_block          = "10.1.21.0/24"
  display_name        = "subnet2lbRD"
  dns_label           = "subnet2lbrd"
  security_list_ids   = ["${oci_core_security_list.securitylist_lb_rd.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn_rd.id}"
  route_table_id      = "${oci_core_route_table.routetable_rd.id}"

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_instance" "instance_lb_rd" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "backendInstanceLBRD"
  shape               = "${var.instance_shape}"

  metadata = {
    user_data = "${base64encode(var.user-data)}"
  }

  create_vnic_details {
    subnet_id      = "${oci_core_subnet.subnet1_lb_rd.id}"
    hostname_label = "backendinstance"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
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

resource "oci_load_balancer" "load_balancer_rd" {
  shape          = "100Mbps"
  compartment_id = "${var.compartment_ocid}"

  subnet_ids = [
    "${oci_core_subnet.subnet1_lb_rd.id}",
    "${oci_core_subnet.subnet2_lb_rd.id}",
  ]

  display_name = "lbRD"
}

resource "oci_load_balancer_backend_set" "lb_backend_set_rd" {
  name             = "lb-besRD"
  load_balancer_id = "${oci_load_balancer.load_balancer_rd.id}"
  policy           = "ROUND_ROBIN"

  health_checker {
    port                = "80"
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path            = "/"
  }
}

resource "oci_load_balancer_certificate" "lb_cert_rd" {
  load_balancer_id   = "${oci_load_balancer.load_balancer_rd.id}"
  ca_certificate     = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"
  certificate_name   = "certificate1"
  private_key        = "${var.private_key_data}"
  public_certificate = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"

  lifecycle {
    create_before_destroy = true
  }
}

resource "oci_load_balancer_path_route_set" "test_path_route_set_rd" {
  #Required
  load_balancer_id = "${oci_load_balancer.load_balancer_rd.id}"
  name             = "pr-setLBRD"

  path_routes {
    #Required
    backend_set_name = "${oci_load_balancer_backend_set.lb_backend_set_rd.name}"
    path             = "/example/video/123"

    path_match_type {
      #Required
      match_type = "EXACT_MATCH"
    }
  }
}

resource "oci_load_balancer_hostname" "test_hostname_lb_rd" {
  #Required
  hostname         = "app.example.com"
  load_balancer_id = "${oci_load_balancer.load_balancer_rd.id}"
  name             = "hostnameLBRD"
}

resource "oci_load_balancer_listener" "lb_listener_rd" {
  load_balancer_id         = "${oci_load_balancer.load_balancer_rd.id}"
  name                     = "http"
  default_backend_set_name = "${oci_load_balancer_backend_set.lb_backend_set_rd.name}"
  hostname_names           = ["${oci_load_balancer_hostname.test_hostname_lb_rd.name}"]
  port                     = 80
  protocol                 = "HTTP"
  rule_set_names           = ["${oci_load_balancer_rule_set.test_rule_set_lb_rd.name}"]

  connection_configuration {
    idle_timeout_in_seconds = "2"
  }
}

resource "oci_load_balancer_backend" "lb-be1" {
  load_balancer_id = "${oci_load_balancer.load_balancer_rd.id}"
  backendset_name  = "${oci_load_balancer_backend_set.lb_backend_set_rd.name}"
  ip_address       = "${oci_core_instance.instance_lb_rd.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

resource "oci_load_balancer_rule_set" "test_rule_set_lb_rd" {
  items {
    action = "ADD_HTTP_REQUEST_HEADER"
    header = "example_header_name"
    value  = "example_header_value"
  }

  items {
    action          = "CONTROL_ACCESS_USING_HTTP_METHODS"
    allowed_methods = ["GET", "POST"]
    status_code     = "405"
  }

  items {
    action      = "ALLOW"
    description = "example vcn ACL"

    conditions {
      attribute_name  = "SOURCE_VCN_ID"
      attribute_value = "${oci_core_vcn.vcn_rd.id}"
    }

    conditions {
      attribute_name  = "SOURCE_VCN_IP_ADDRESS"
      attribute_value = "10.10.1.0/24"
    }
  }

  items {
    action = "REDIRECT"

    conditions {
      attribute_name  = "PATH"
      attribute_value = "/example"
      operator        = "FORCE_LONGEST_PREFIX_MATCH"
    }

    redirect_uri {
      protocol = "{protocol}"
      host     = "in{host}"
      port     = 8081
      path     = "{path}/video"
      query    = "?lang=en"
    }

    response_code = 302
  }

  load_balancer_id = "${oci_load_balancer.load_balancer_rd.id}"
  name             = "exampleRuleSetNameRD"
}
