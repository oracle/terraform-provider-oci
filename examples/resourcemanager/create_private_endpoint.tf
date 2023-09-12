variable "compartment_ocid" {}
variable "region" {}
variable "tenancy_ocid" {}

provider "oci" {
  region = var.region
}

data "oci_identity_availability_domains" "get_availability_domains" {
  compartment_id = var.tenancy_ocid
}

// Local variables defined to increase readability
locals {
  private_endpoint_integ_test_vcn_cidr_block   = "10.12.0.0/16"
  private_endpoint_integ_test_vcn_subnet_count = 1
  tcp_protocol                                 = 6
  default_shape_name                           = "VM.Standard.E3.Flex"
  operating_system                             = "Oracle Linux"
  operating_system_version                     = "8"
}

// VCN holding the private subnet
resource "oci_core_vcn" "private_endpoint_integ_test_temp_vcn" {
  cidr_block     = "10.12.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "VCN for remote-exec integ test"
}

// Private subnet the compute instance will reside in
resource "oci_core_subnet" "private_endpoint_integ_test_temp_subnet" {
  compartment_id = var.compartment_ocid
  display_name   = "private_endpoint_integ_test_subnet"
  vcn_id         = oci_core_vcn.private_endpoint_integ_test_temp_vcn.id

  prohibit_public_ip_on_vnic = true
  cidr_block = cidrsubnet(
    local.private_endpoint_integ_test_vcn_cidr_block,
    8,
    1,
  )

  security_list_ids = [oci_core_security_list.terraform_private_endpoint_security_list.id]
}

data "oci_core_images" "available_instance_images" {
  compartment_id           = var.compartment_ocid
  operating_system         = local.operating_system
  operating_system_version = local.operating_system_version
  shape                    = local.default_shape_name
}

// Compute instance that our SSH connection will be established with.
resource "oci_core_instance" "private_endpoint_instance" {
  compartment_id = var.compartment_ocid
  display_name   = "test script as one remote-exec instance"

  availability_domain = lookup(data.oci_identity_availability_domains.get_availability_domains.availability_domains[0], "name")
  shape               = local.default_shape_name

  // specify the subnet and that there is no public IP assigned to the instance
  create_vnic_details {
    subnet_id        = oci_core_subnet.private_endpoint_integ_test_temp_subnet.id
    assign_public_ip = false
  }

  extended_metadata = {
    ssh_authorized_keys = tls_private_key.public_private_key_pair.public_key_openssh
  }

  // use latest oracle linux image via data source
  source_details {
    source_id   = data.oci_core_images.available_instance_images.images[0].id
    source_type = "image"
  }

  shape_config {
    memory_in_gbs = 4
    ocpus         = 1
  }
}

// The RMS private endpoint resource. Requires a VCN with a private subnet
resource "oci_resourcemanager_private_endpoint" "rms_private_endpoint" {
  compartment_id = var.compartment_ocid
  display_name   = "rms_private_endpoint"
  description    = "rms_private_endpoint_description"
  vcn_id         = oci_core_vcn.private_endpoint_integ_test_temp_vcn.id
  subnet_id      = oci_core_subnet.private_endpoint_integ_test_temp_subnet.id
}

// Resolves the private IP of the customer's private endpoint to a NAT IP. Used as the host address in the "remote-exec" resource
data "oci_resourcemanager_private_endpoint_reachable_ip" "test_private_endpoint_reachable_ips" {
  private_endpoint_id = oci_resourcemanager_private_endpoint.rms_private_endpoint.id
  private_ip          = oci_core_instance.private_endpoint_instance.private_ip
}

resource "oci_core_security_list" "terraform_private_endpoint_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.private_endpoint_integ_test_temp_vcn.id
  display_name   = "integ test vcn security list"

  // Lock down ingress and egress traffic to the VCN cidr block. Can be restricted further to be subnet cidr range
  // Only allow SSH communication on specific port
  ingress_security_rules {
    protocol = local.tcp_protocol
    source   = local.private_endpoint_integ_test_vcn_cidr_block
    tcp_options {
      min = 22
      max = 22
    }
  }

  egress_security_rules {
    destination = local.private_endpoint_integ_test_vcn_cidr_block
    protocol    = local.tcp_protocol
    tcp_options {
      min = 22
      max = 22
    }
  }
}

// The public/private key used to SSH to the compute instance
resource "tls_private_key" "public_private_key_pair" {
  algorithm = "RSA"
}

// Resource to establish the SSH connection. Must have the compute instance created first.
resource "null_resource" "remote-exec" {
  depends_on = [oci_core_instance.private_endpoint_instance]

  provisioner "remote-exec" {
    connection {
      agent       = false
      timeout     = "30m"
      host        = data.oci_resourcemanager_private_endpoint_reachable_ip.test_private_endpoint_reachable_ips.ip_address
      user        = "opc"
      private_key = tls_private_key.public_private_key_pair.private_key_pem
    }
    // write to a file on the compute instance via the private access SSH connection
    inline = [
      "echo 'remote exec showcase' > ~/remoteExecTest.txt"
    ]
  }
}