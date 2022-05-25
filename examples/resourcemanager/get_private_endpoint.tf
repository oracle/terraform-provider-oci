variable "compartment_ocid" {}
variable "region" {}
// VCN holding the private subnet
variable "tenancy_ocid" {}
// Private subnet the compute instance will reside in
variable "subnet_ocid" {}
// Private endpoint that has already been created.
variable "orm_private_endpoint_ocid" {
  type = string
}

provider "oci" {
  region = "${var.region}"
}

// Local variables defined to increase readability.
// Only local variable that should remain consistent is tcp_protocol, as this specifies it will be used for SSH
locals {
  tcp_protocol = 6
  default_shape_name = "VM.Standard.E3.Flex"
  operating_system = "Oracle Linux"
}

data "oci_identity_availability_domains" "get_availability_domains" {
  compartment_id = var.tenancy_ocid
}


data "oci_core_images" "available_instance_images" {
  compartment_id = var.compartment_ocid
  operating_system = local.operating_system
  shape = local.default_shape_name
}

// Use a data source to get a pre-existing private endpoint. This private endpoint could already be created via CLI, SDK, console, etc
// in your tenancy
data "oci_resourcemanager_private_endpoint" "get_private_endpoint" {
  private_endpoint_id = var.orm_private_endpoint_ocid
}

// Resolves the private IP of the customer's private endpoint to a NAT IP. Used as the host address in the "remote-exec" resource
data "oci_resourcemanager_private_endpoint_reachable_ip" "test_private_endpoint_reachable_ips" {
  private_endpoint_id = data.oci_resourcemanager_private_endpoint.get_private_endpoint.id
  private_ip          = oci_core_instance.private_endpoint_instance.private_ip
}

// The public/private key used to SSH to the compute instance
resource "tls_private_key" "public_private_key_pair" {
  algorithm = "RSA"
}

// Compute instance the private endpoint will allow SSH communication to
resource "oci_core_instance" "private_endpoint_instance" {
  compartment_id = var.compartment_ocid
  display_name = "test script as one remote-exec instance"

  availability_domain = lookup(data.oci_identity_availability_domains.get_availability_domains.availability_domains[0], "name")
  shape = local.default_shape_name

  // specify this is a private by not assigning public ip
  create_vnic_details {
    subnet_id = var.subnet_ocid
    assign_public_ip = false
  }

  extended_metadata = {
    ssh_authorized_keys = tls_private_key.public_private_key_pair.public_key_openssh
  }

  source_details {
    source_id = data.oci_core_images.available_instance_images.images[0].id
    source_type = "image"
  }

  shape_config {
    memory_in_gbs = 4
    ocpus = 1
  }
}

// Resource to establish the SSH connection. Must have the compute instance created first.
resource "null_resource" "remote-exec" {
  depends_on = [oci_core_instance.private_endpoint_instance]

  provisioner "remote-exec" {
    connection {
      agent = false
      timeout = "30m"
      host = data.oci_resourcemanager_private_endpoint_reachable_ip.test_private_endpoint_reachable_ips.ip_address
      user = "opc"
      private_key = tls_private_key.public_private_key_pair.private_key_pem
    }
    // write to a file on the compute instance via the private access SSH connection
    inline = [
      "echo 'remote exec showcase' > ~/remoteExecTest.txt"
    ]
  }
}