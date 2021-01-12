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

variable "ssh_public_key" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

variable "ad_region_mapping" {
  type = map(string)

  default = {
    us-phoenix-1 = 2
    us-ashburn-1 = 3
    us-seattle-1 = 2
  }
}

variable "images" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaadtmpmfm77czi5ghi5zh7uvkguu6dsecsg7kuo3eigc5663und4za"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaayuihpsm2nfkxztdkottbjtfjqhgod7hfuirt2rqlewxrmdlgg75q"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
    us-seattle-1   = "ocid1.image.region1.sea.aaaaaaaau2iablsucjq62zapd3fr4laaoiqviwsujgois2wcaejaazjnhivq"
  }
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = var.ad_region_mapping[var.region]
}

resource "oci_core_virtual_network" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "testVCN"
  dns_label      = "testvcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block        = "10.1.20.0/24"
  display_name      = "testSubnet"
  dns_label         = "testsubnet"
  security_list_ids = [oci_core_security_list.test_security_list.id]
  compartment_id    = var.compartment_ocid
  vcn_id            = oci_core_virtual_network.test_vcn.id
  route_table_id    = oci_core_route_table.test_route_table.id
  dhcp_options_id   = oci_core_virtual_network.test_vcn.default_dhcp_options_id
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "testIG"
  vcn_id         = oci_core_virtual_network.test_vcn.id
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_virtual_network.test_vcn.id
  display_name   = "testRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_virtual_network.test_vcn.id
  display_name   = "testSecurityList"

  egress_security_rules {
    protocol    = "6"
    destination = "0.0.0.0/0"
  }

  ingress_security_rules {
    protocol = "6"
    source   = "0.0.0.0/0"

    tcp_options {
      max = "22"
      min = "22"
    }
  }

  ingress_security_rules {
    protocol = "6"
    source   = "0.0.0.0/0"

    tcp_options {
      max = "3000"
      min = "3000"
    }
  }

  ingress_security_rules {
    protocol = "6"
    source   = "0.0.0.0/0"

    tcp_options {
      max = "3005"
      min = "3005"
    }
  }

  ingress_security_rules {
    protocol = "6"
    source   = "0.0.0.0/0"

    tcp_options {
      max = "80"
      min = "80"
    }
  }
}

resource "oci_core_instance" "free_instance0" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "freeInstance0"
  shape               = "VM.Standard.E2.1.Micro"

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "freeinstance0"
  }

  source_details {
    source_type = "image"
    source_id   = var.images[var.region]
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }
}

resource "oci_core_instance" "free_instance1" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "freeInstance1"
  shape               = "VM.Standard.E2.1.Micro"

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "freeinstance1"
  }

  source_details {
    source_type = "image"
    source_id   = var.images[var.region]
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }
}

resource "oci_load_balancer" "free_load_balancer" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "alwaysFreeLoadBalancer"
  shape          = "10Mbps-Micro"

  subnet_ids = [
    oci_core_subnet.test_subnet.id,
  ]
}

resource "oci_load_balancer_backend_set" "free_load_balancer_backend_set" {
  name             = "lbBackendSet1"
  load_balancer_id = oci_load_balancer.free_load_balancer.id
  policy           = "ROUND_ROBIN"

  health_checker {
    port                = "80"
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path            = "/"
  }

  session_persistence_configuration {
    cookie_name      = "lb-session1"
    disable_fallback = true
  }
}

resource "oci_load_balancer_backend" "free_load_balancer_test_backend0" {
  #Required
  backendset_name  = oci_load_balancer_backend_set.free_load_balancer_backend_set.name
  ip_address       = oci_core_instance.free_instance0.public_ip
  load_balancer_id = oci_load_balancer.free_load_balancer.id
  port             = "80"
}

resource "oci_load_balancer_backend" "free_load_balancer_test_backend1" {
  #Required
  backendset_name  = oci_load_balancer_backend_set.free_load_balancer_backend_set.name
  ip_address       = oci_core_instance.free_instance1.public_ip
  load_balancer_id = oci_load_balancer.free_load_balancer.id
  port             = "80"
}

resource "oci_load_balancer_hostname" "test_hostname1" {
  #Required
  hostname         = "app.free.com"
  load_balancer_id = oci_load_balancer.free_load_balancer.id
  name             = "hostname1"
}

resource "oci_load_balancer_listener" "load_balancer_listener0" {
  load_balancer_id         = oci_load_balancer.free_load_balancer.id
  name                     = "http"
  default_backend_set_name = oci_load_balancer_backend_set.free_load_balancer_backend_set.name
  hostname_names           = [oci_load_balancer_hostname.test_hostname1.name]
  port                     = 80
  protocol                 = "HTTP"
  rule_set_names           = [oci_load_balancer_rule_set.test_rule_set.name]

  connection_configuration {
    idle_timeout_in_seconds = "240"
  }
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

  load_balancer_id = oci_load_balancer.free_load_balancer.id
  name             = "test_rule_set_name"
}

resource "oci_load_balancer_certificate" "load_balancer_certificate" {
  load_balancer_id   = oci_load_balancer.free_load_balancer.id
  ca_certificate     = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"
  certificate_name   = "certificate1"
  private_key        = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA30+wt7OlUB/YpmWbTRkxnLG0lKWiV+oupNKj8luXmC5jvOFT\nUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU+DWVV2So2B/obYxpiiyWF2tcF/cY\n\ni1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oCMQ2985/MTdCXONgnbmePU64GrJwf\nvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOOjLKRM68KXC5us4879IrSA77NQr1K\nwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6ytM66P/1CTpk1YpbI4gqiG0HBbuX\nG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc2wIDAQABAoIBAGQznukfG/uS/qTT\njNcQifl0p8HXfLwUIa/lsJkMTj6D+k8DkF59tVMGjv3NQSQ26JVX4J1L8XiAj+fc\nUtYr1Ap4CLX5PeYUkzesvKK6lPKXQvCh+Ip2eq9PVrvL2WcdDpb5695cy7suXD7c\n05aUtS0LrINH3eXAxkpEe5UHtQFni5YLrCLEXd+SSA3OKdCB+23HRELc1iCTvqjK\n5AtR916uHTBhtREHRMvWIdu4InRUsedlJhaJOLJ8G8r64JUtfm3wLUK1U8HFOsd0\nLAx9ZURU6cXl4osTWiy1vigGaM8Xuish2HkOLNYZADDUiDBB3SshmW5IDAJ5XTn5\nqVrszRECgYEA79j1y+WLTyV7yz7XkWk3OqoQXG4b2JfKItJI1M95UwllzQ8U/krM\n+QZjP3NTtB9i1YoHyaEfic103hV9Fkgz8jvKS5ocLGJulpN4CgqbHN6v9EJ3dqTk\no6X8mpx2eP2E0ngRekFyC/OCp0Zhe2KR9PXhijMa5eB2LTeCMIS/tzkCgYEA7lmk\nIdVjcpfqY7UFJ2R8zqPJHOne2+llrl9vzo6N5kx4DzAg7MP6XO9MekOvfmD1X1Lm\nFckXWFEF+0TlN5YvCTR/+OmVufYM3xp4GBT8RZdLFbyI4+xpAAeSC4SeM0ZkC9Jt\nrKqCS24+Kqy/+qSqtkxiPLQrXSdCSfCUlmn0ALMCgYBB7SLy3q+CG82BOk7Km18g\n8un4XhOtX1uiYqa+SCETH/wpd0HP/AOHV6gkIrEZS59BDuXBGFaw7BZ5jPKLE2Gj\n7adXTI797Dh1jydpqyyjrNo0i6iGpiBqkw9x+Bvged7ucy5qql6MxmxdSk01Owzf\nhk5uTEnScfZJy34vk+2WkQKBgBXx5uy+iuN4HTqE5i6UT/FunwusdLpmqNf/LXol\nIed8TumHEuD5wklgNvhi1vuZzb2zEkAbPa0B+L0DwN73UulUDhxK1WBDyTeZZklB\nVWDK5zzfGPNzRs+b4tRwp2gtKPT1sOde45QyWELxmNNo6dbS/ZB9Pijbfnz0S5n1\ns2OFAoGBAJUohI1+d2hKlkSUzpCorQarDe8lFVEbGMu0kX0JNDI7QU+H8vDp9NOl\nGqLm3sCVBYypT8sWfchgZpcVaLRLQCQtWy4+CbMN6DT3j/uBWeDpayU5Gvqt0/no\nvwqbG6b0NEYLRPLEdsS/c8TV9mMlvb0EW+GXfmkpTrTNt3hyXniu\n-----END RSA PRIVATE KEY-----"
  public_certificate = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"

  lifecycle {
    create_before_destroy = true
  }
}

resource "oci_load_balancer_listener" "load_balancer_listener1" {
  load_balancer_id         = oci_load_balancer.free_load_balancer.id
  name                     = "https"
  default_backend_set_name = oci_load_balancer_backend_set.free_load_balancer_backend_set.name
  port                     = 443
  protocol                 = "HTTP"

  ssl_configuration {
    certificate_name        = oci_load_balancer_certificate.load_balancer_certificate.certificate_name
    verify_peer_certificate = false
  }
}

output "lb_public_ip" {
  value = [oci_load_balancer.free_load_balancer.ip_address_details]
}

data "oci_core_vnic_attachments" "app_vnics" {
  compartment_id      = var.compartment_ocid
  availability_domain = data.oci_identity_availability_domain.ad.name
  instance_id         = oci_core_instance.free_instance0.id
}

data "oci_core_vnic" "app_vnic" {
  vnic_id = data.oci_core_vnic_attachments.app_vnics.vnic_attachments[0]["vnic_id"]
}

data "oci_core_images" "test_images" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  shape = "VM.Standard.E2.1.Micro"
}

output "app" {
  value = "http://${data.oci_core_vnic.app_vnic.public_ip_address}"
}

data "oci_database_autonomous_databases" "test_autonomous_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  db_workload  = "OLTP"
  is_free_tier = "true"
}

resource "oci_database_autonomous_database" "test_autonomous_database" {
  #Required
  admin_password           = "Testalwaysfree1"
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "testadb"

  #Optional
  db_workload  = "OLTP"
  display_name = "test_autonomous_database"

  freeform_tags = {
    "Department" = "Finance"
  }

  is_auto_scaling_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  is_free_tier            = "true"
}

