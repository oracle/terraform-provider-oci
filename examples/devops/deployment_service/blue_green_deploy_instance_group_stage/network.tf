data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = var.tenancy_ocid
  ad_number      = 2
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

resource "oci_core_vcn" "example_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "example_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "exampleSubnet"
  dns_label           = "tfexamplesubnet"

  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
}

//Regional loadbalancer needs subnets in atleast 2 ADs
resource "oci_core_subnet" "example_subnet_ad2" {
  availability_domain = data.oci_identity_availability_domain.ad2.name
  cidr_block          = "10.1.24.0/24"
  display_name        = "exampleSubnetAD2"
  dns_label           = "tfsubnet2"

  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
}

# See https://docs.oracle.com/iaas/images/
data "oci_core_images" "test_images" {
  compartment_id           = var.compartment_ocid
  operating_system         = "Oracle Linux"
  operating_system_version = "8"
  shape                    = var.instance_shape
  sort_by                  = "TIMECREATED"
  sort_order               = "DESC"
}

resource "oci_core_instance" "example_instance_a" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "testInstance"
  shape               = var.instance_shape

  create_vnic_details {
    hostname_label         = "instancea"
    subnet_id              = oci_core_subnet.example_subnet.id
    skip_source_dest_check = true
    assign_public_ip       = true
  }

  source_details {
    source_type = "image"
    source_id   = lookup(data.oci_core_images.test_images.images[0], "id")
  }
}

resource "oci_core_instance" "example_instance_b" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "testInstanceB"
  shape               = var.instance_shape

  create_vnic_details {
    hostname_label         = "instance"
    subnet_id              = oci_core_subnet.example_subnet.id
    skip_source_dest_check = true
    assign_public_ip       = true
  }

  source_details {
    source_type = "image"
    source_id   = lookup(data.oci_core_images.test_images.images[0], "id")
  }
}

/* Load Balancer */

resource "oci_load_balancer_load_balancer" "prod_load_balancer" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "ProdLoadBalancer"
  shape          = "flexible"
  shape_details {
    maximum_bandwidth_in_mbps = 10
    minimum_bandwidth_in_mbps = 10
  }

  subnet_ids = [
    oci_core_subnet.example_subnet.id, oci_core_subnet.example_subnet_ad2.id,
  ]
}

resource "oci_load_balancer_backend_set" "prod_load_balancer_backend_set" {
  name             = "lbBackendSet1"
  load_balancer_id = oci_load_balancer_load_balancer.prod_load_balancer.id
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

resource "oci_load_balancer_listener" "prod_load_balancer_listener" {
  load_balancer_id         = oci_load_balancer_load_balancer.prod_load_balancer.id
  name                     = "http"
  default_backend_set_name = oci_load_balancer_backend_set.prod_load_balancer_backend_set.name
  port                     = 80
  protocol                 = "HTTP"

  connection_configuration {
    idle_timeout_in_seconds = "240"
  }
}

resource "oci_load_balancer_load_balancer" "test_load_balancer" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "TestLoadBalancer"
  shape          = "flexible"
  shape_details {
    maximum_bandwidth_in_mbps = 10
    minimum_bandwidth_in_mbps = 10
  }

  subnet_ids = [
    oci_core_subnet.example_subnet.id, oci_core_subnet.example_subnet_ad2.id,
  ]
}

resource "oci_load_balancer_backend_set" "test_load_balancer_backend_set" {
  name             = "lbBackendSet2"
  load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
  policy           = "ROUND_ROBIN"

  health_checker {
    port                = "81"
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path            = "/"
  }

  session_persistence_configuration {
    cookie_name      = "lb-session1"
    disable_fallback = true
  }
}

resource "oci_load_balancer_listener" "test_load_balancer_listener" {
  load_balancer_id         = oci_load_balancer_load_balancer.test_load_balancer.id
  name                     = "http"
  default_backend_set_name = oci_load_balancer_backend_set.test_load_balancer_backend_set.name
  port                     = 81
  protocol                 = "HTTP"

  connection_configuration {
    idle_timeout_in_seconds = "240"
  }
}