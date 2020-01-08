// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

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
variable "region" {}

variable "instance_image_ocid" {
  type = "map"

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"

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
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

data "oci_identity_availability_domain" "ad1" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 2
}

/* Network */

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "vcn1"
  dns_label      = "vcn1"
}

resource "oci_core_subnet" "subnet1" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "subnet1"
  dns_label           = "subnet1"
  security_list_ids   = ["${oci_core_security_list.securitylist1.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn1.id}"
  route_table_id      = "${oci_core_route_table.routetable1.id}"
  dhcp_options_id     = "${oci_core_vcn.vcn1.default_dhcp_options_id}"

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_subnet" "subnet2" {
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  cidr_block          = "10.1.21.0/24"
  display_name        = "subnet2"
  dns_label           = "subnet2"
  security_list_ids   = ["${oci_core_security_list.securitylist1.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn1.id}"
  route_table_id      = "${oci_core_route_table.routetable1.id}"
  dhcp_options_id     = "${oci_core_vcn.vcn1.default_dhcp_options_id}"

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_internet_gateway" "internetgateway1" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "internetgateway1"
  vcn_id         = "${oci_core_vcn.vcn1.id}"
}

resource "oci_core_route_table" "routetable1" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn1.id}"
  display_name   = "routetable1"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.internetgateway1.id}"
  }
}

resource "oci_core_security_list" "securitylist1" {
  display_name   = "public"
  compartment_id = "${oci_core_vcn.vcn1.compartment_id}"
  vcn_id         = "${oci_core_vcn.vcn1.id}"

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
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "be-instance1"
  shape               = "${var.instance_shape}"

  metadata = {
    user_data = "${base64encode(var.user-data)}"
  }

  create_vnic_details {
    subnet_id      = "${oci_core_subnet.subnet1.id}"
    hostname_label = "be-instance1"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }
}

resource "oci_core_instance" "instance2" {
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "be-instance2"
  shape               = "${var.instance_shape}"

  metadata = {
    user_data = "${base64encode(var.user-data)}"
  }

  create_vnic_details {
    subnet_id      = "${oci_core_subnet.subnet2.id}"
    hostname_label = "be-instance2"
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

resource "oci_load_balancer" "lb1" {
  shape          = "100Mbps"
  compartment_id = "${var.compartment_ocid}"

  subnet_ids = [
    "${oci_core_subnet.subnet1.id}",
    "${oci_core_subnet.subnet2.id}",
  ]

  display_name = "lb1"
}

resource "oci_load_balancer_backend_set" "lb-bes1" {
  name             = "lb-bes1"
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  policy           = "ROUND_ROBIN"

  health_checker {
    port                = "80"
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path            = "/"
  }
}

resource "oci_load_balancer_certificate" "lb-cert1" {
  load_balancer_id   = "${oci_load_balancer.lb1.id}"
  ca_certificate     = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"
  certificate_name   = "certificate1"
  private_key        = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA30+wt7OlUB/YpmWbTRkxnLG0lKWiV+oupNKj8luXmC5jvOFT\nUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU+DWVV2So2B/obYxpiiyWF2tcF/cY\n\ni1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oCMQ2985/MTdCXONgnbmePU64GrJwf\nvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOOjLKRM68KXC5us4879IrSA77NQr1K\nwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6ytM66P/1CTpk1YpbI4gqiG0HBbuX\nG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc2wIDAQABAoIBAGQznukfG/uS/qTT\njNcQifl0p8HXfLwUIa/lsJkMTj6D+k8DkF59tVMGjv3NQSQ26JVX4J1L8XiAj+fc\nUtYr1Ap4CLX5PeYUkzesvKK6lPKXQvCh+Ip2eq9PVrvL2WcdDpb5695cy7suXD7c\n05aUtS0LrINH3eXAxkpEe5UHtQFni5YLrCLEXd+SSA3OKdCB+23HRELc1iCTvqjK\n5AtR916uHTBhtREHRMvWIdu4InRUsedlJhaJOLJ8G8r64JUtfm3wLUK1U8HFOsd0\nLAx9ZURU6cXl4osTWiy1vigGaM8Xuish2HkOLNYZADDUiDBB3SshmW5IDAJ5XTn5\nqVrszRECgYEA79j1y+WLTyV7yz7XkWk3OqoQXG4b2JfKItJI1M95UwllzQ8U/krM\n+QZjP3NTtB9i1YoHyaEfic103hV9Fkgz8jvKS5ocLGJulpN4CgqbHN6v9EJ3dqTk\no6X8mpx2eP2E0ngRekFyC/OCp0Zhe2KR9PXhijMa5eB2LTeCMIS/tzkCgYEA7lmk\nIdVjcpfqY7UFJ2R8zqPJHOne2+llrl9vzo6N5kx4DzAg7MP6XO9MekOvfmD1X1Lm\nFckXWFEF+0TlN5YvCTR/+OmVufYM3xp4GBT8RZdLFbyI4+xpAAeSC4SeM0ZkC9Jt\nrKqCS24+Kqy/+qSqtkxiPLQrXSdCSfCUlmn0ALMCgYBB7SLy3q+CG82BOk7Km18g\n8un4XhOtX1uiYqa+SCETH/wpd0HP/AOHV6gkIrEZS59BDuXBGFaw7BZ5jPKLE2Gj\n7adXTI797Dh1jydpqyyjrNo0i6iGpiBqkw9x+Bvged7ucy5qql6MxmxdSk01Owzf\nhk5uTEnScfZJy34vk+2WkQKBgBXx5uy+iuN4HTqE5i6UT/FunwusdLpmqNf/LXol\nIed8TumHEuD5wklgNvhi1vuZzb2zEkAbPa0B+L0DwN73UulUDhxK1WBDyTeZZklB\nVWDK5zzfGPNzRs+b4tRwp2gtKPT1sOde45QyWELxmNNo6dbS/ZB9Pijbfnz0S5n1\ns2OFAoGBAJUohI1+d2hKlkSUzpCorQarDe8lFVEbGMu0kX0JNDI7QU+H8vDp9NOl\nGqLm3sCVBYypT8sWfchgZpcVaLRLQCQtWy4+CbMN6DT3j/uBWeDpayU5Gvqt0/no\nvwqbG6b0NEYLRPLEdsS/c8TV9mMlvb0EW+GXfmkpTrTNt3hyXniu\n-----END RSA PRIVATE KEY-----"
  public_certificate = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"

  lifecycle {
    create_before_destroy = true
  }
}

resource "oci_load_balancer_path_route_set" "test_path_route_set" {
  #Required
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  name             = "pr-set1"

  path_routes {
    #Required
    backend_set_name = "${oci_load_balancer_backend_set.lb-bes1.name}"
    path             = "/example/video/123"

    path_match_type {
      #Required
      match_type = "EXACT_MATCH"
    }
  }
}

resource "oci_load_balancer_hostname" "test_hostname1" {
  #Required
  hostname         = "app.example.com"
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  name             = "hostname1"
}

resource "oci_load_balancer_hostname" "test_hostname2" {
  #Required
  hostname         = "app2.example.com"
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  name             = "hostname2"
}

resource "oci_load_balancer_listener" "lb-listener1" {
  load_balancer_id         = "${oci_load_balancer.lb1.id}"
  name                     = "http"
  default_backend_set_name = "${oci_load_balancer_backend_set.lb-bes1.name}"
  hostname_names           = ["${oci_load_balancer_hostname.test_hostname1.name}", "${oci_load_balancer_hostname.test_hostname2.name}"]
  port                     = 80
  protocol                 = "HTTP"
  rule_set_names           = ["${oci_load_balancer_rule_set.test_rule_set.name}"]

  connection_configuration {
    idle_timeout_in_seconds = "2"
  }
}

resource "oci_load_balancer_listener" "lb-listener2" {
  load_balancer_id         = "${oci_load_balancer.lb1.id}"
  name                     = "https"
  default_backend_set_name = "${oci_load_balancer_backend_set.lb-bes1.name}"
  port                     = 443
  protocol                 = "HTTP"

  ssl_configuration {
    certificate_name        = "${oci_load_balancer_certificate.lb-cert1.certificate_name}"
    verify_peer_certificate = false
  }
}

resource "oci_load_balancer_listener" "lb-listener3" {
  load_balancer_id         = "${oci_load_balancer.lb1.id}"
  name                     = "tcp"
  default_backend_set_name = "${oci_load_balancer_backend_set.lb-bes1.name}"
  port                     = 80
  protocol                 = "TCP"
  rule_set_names           = ["${oci_load_balancer_rule_set.test_rule_set.name}"]

  connection_configuration {
    idle_timeout_in_seconds            = "2"
    backend_tcp_proxy_protocol_version = "1"
  }
}

resource "oci_load_balancer_backend" "lb-be1" {
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  backendset_name  = "${oci_load_balancer_backend_set.lb-bes1.name}"
  ip_address       = "${oci_core_instance.instance1.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

resource "oci_load_balancer_backend" "lb-be2" {
  load_balancer_id = "${oci_load_balancer.lb1.id}"
  backendset_name  = "${oci_load_balancer_backend_set.lb-bes1.name}"
  ip_address       = "${oci_core_instance.instance2.private_ip}"
  port             = 80
  backup           = false
  drain            = false
  offline          = false
  weight           = 1
}

resource "oci_load_balancer_rule_set" "test_rule_set" {
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
      attribute_value = "${oci_core_vcn.vcn1.id}"
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

  load_balancer_id = "${oci_load_balancer.lb1.id}"
  name             = "example_rule_set_name"
}

output "lb_public_ip" {
  value = ["${oci_load_balancer.lb1.ip_address_details}"]
}
