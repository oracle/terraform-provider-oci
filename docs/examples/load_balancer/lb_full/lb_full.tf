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
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    // Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
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

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

/* Network */

resource "oci_core_virtual_network" "vcn1" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "vcn1"
  dns_label      = "vcn1"
}

resource "oci_core_subnet" "subnet1" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain -2],"name")}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "subnet1"
  dns_label           = "subnet1"
  security_list_ids   = ["${oci_core_security_list.securitylist1.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.vcn1.id}"
  route_table_id      = "${oci_core_route_table.routetable1.id}"
  dhcp_options_id     = "${oci_core_virtual_network.vcn1.default_dhcp_options_id}"

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_subnet" "subnet2" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain -1],"name")}"
  cidr_block          = "10.1.21.0/24"
  display_name        = "subnet2"
  dns_label           = "subnet2"
  security_list_ids   = ["${oci_core_security_list.securitylist1.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.vcn1.id}"
  route_table_id      = "${oci_core_route_table.routetable1.id}"
  dhcp_options_id     = "${oci_core_virtual_network.vcn1.default_dhcp_options_id}"

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_internet_gateway" "internetgateway1" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "internetgateway1"
  vcn_id         = "${oci_core_virtual_network.vcn1.id}"
}

resource "oci_core_route_table" "routetable1" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.vcn1.id}"
  display_name   = "routetable1"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.internetgateway1.id}"
  }
}

resource "oci_core_security_list" "securitylist1" {
  display_name   = "public"
  compartment_id = "${oci_core_virtual_network.vcn1.compartment_id}"
  vcn_id         = "${oci_core_virtual_network.vcn1.id}"

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

resource "oci_core_instance" "instance1" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain -2],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "be-instance1"
  shape               = "${var.instance_shape}"
  subnet_id           = "${oci_core_subnet.subnet1.id}"
  hostname_label      = "be-instance1"

  metadata {
    user_data = "${base64encode(var.user-data)}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }
}

resource "oci_core_instance" "instance2" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain -1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "be-instance2"
  shape               = "${var.instance_shape}"
  subnet_id           = "${oci_core_subnet.subnet2.id}"
  hostname_label      = "be-instance2"

  metadata {
    user_data = "${base64encode(var.user-data)}"
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
  ca_certificate     = "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----"
  certificate_name   = "certificate1"
  private_key        = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1vqH90VOIZHh/FL9pEH23op3t+FcMBSKQ+ijCoDZfil5m12X\nUQWUTS/ZFauGo0q1O4esHI8vusL6eyBVDOiTvyGfgXBYWq+yPkIGlqcuOdxN/bQv\nhfAJN5hcJYjosT/KhlwuuYO7V4ZWcgpacoyETh5bMinEPJ1k7e6dF4/4p3lisERP\n2zkxDtP98ZlUbo0ZF6sXPe9PJqYUCk0DXkH1GkffX6Zp+39q+ywJVP6Up54USuZ5\nmCJRyrGn0hq5+ha/UNLm6ggxa3sE2+/Rzx9sFrWLuxuEKQJJ4KPBuNreIHQk2z6K\nTH316SplmZRgTXqotTpYgWF9Ml/14ljvlDo6RwIDAQABAoIBAQCFzHXVQ1BGenpR\nRgHROrEAfuPWETAESLRpYaAgCGPVLtEeDpj/913+0FnnL9NjTDsR6vYG7GNDdNja\nyxvEJfjWy4Fv2VFUV+ey8fsRxslxf5kW3w946BWEgZJQVi6lKtPM3hDCq6ds6RJi\ndeknRCeQSzptNSuKoldP8uPY52VWLYTyy/ODwtSCFZKTm7iTD3RkpAqNMPs2V8EA\nRbjiu70q4Kk+ozHQ/0wtOSjZinR6LW7e+6bXxdVV0hdt52h2YzdQrQkHu5ATm54P\nm1S9PbiyV06BSuU37oZbyWiblP6rsJIqucEXSSmTY/5PCG+huzTseig1bot19eD0\nwEs9YAixAoGBAPgKMZ+VhA3eoAx21R62SmJJKdCYZ+qjPmvUt4tW3dJxYTweJo6A\nwyh+p3VsBwE6L8hPD1OyEvUrl64uQEMYSMmNaLpaMqNbytP8h6uiFGyRUSnkYvxF\nPcKW6UjKZNyGWT+dOnJN07DnwvZgOx2JZJkgFpaAEL6g0Hzm1bI/6dkPAoGBAN3g\nt+0sDfgYr3raGAPLQJCVgHV3MdIDHP/ebZBT1NTcxL7Wf/+0WAvnwD7DPsXj84JG\nzNMk7+EzwClAWGAQJMymC9NltfygyI0JjI+88nVk3mrpDq/zqR/vkn2R1T78No/X\nEWTnMmHlzDg4HsGTWKDg7jrmYSas0NvMHPtBbVtJAoGAKV0B03wKjomOpSV3+uwp\nUWSkDX4s7isU8MSDa0AsM7jmnzDj+yWr5efhIyrFrEW4zC2q/6kVkj8Xx1s9KjM1\niC8FxPXftfBLzbgyI8QepdBB+bt1al5do0KpWpMt6Lyay4n7wi4KXFj54T5A/Xb5\nCLQaMDThFfkZa4rPHi+cXq8CgYAfk665a06lo2W99zH5wEB1E0HP9eG6QMUsyQwQ\nwU2F6dF6U26uBo2NTDM4+3KAmVt7i/X0iso047eSZ1zsdv+1vF/sewo2ZO+F2vkN\nL9fVy0A4OOjlM6k7KU5Q3qNZrm1ZdUM9eAXclubEjYAbDoxLgReGfGkRJwEmdtsd\nCwe0OQKBgQD3m85OXSSf+xlm7tGO66bcxHifkp4XfkqWxFwpfYkxNtYZfFpN8jL5\niS0OyLldmJbCVB6EIs5ylW86aeZMH/JecPTxOnaT4qc43PMrLi4MSa65Gp/Zgs1U\nyO0hfWlpH2ncUIuQEksXEPSKQUjvdQl7pD8kghCbDYbm3zsjw3rkyA==\n-----END RSA PRIVATE KEY-----"
  public_certificate = "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----"

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

  load_balancer_id = "${oci_load_balancer.lb1.id}"
  name             = "example-rule-set-name"
}

output "lb_public_ip" {
  value = ["${oci_load_balancer.lb1.ip_addresses}"]
}
