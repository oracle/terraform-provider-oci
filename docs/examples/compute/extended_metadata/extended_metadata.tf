variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

# Choose an Availability Domain
variable "availability_domain" {
  default = "3"
}

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

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_virtual_network" "ExampleVCN" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "ExampleSubnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "TFExampleSubnet"
  dns_label           = "tfexamplesubnet"
  security_list_ids   = ["${oci_core_virtual_network.ExampleVCN.default_security_list_id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.ExampleVCN.id}"
  route_table_id      = "${oci_core_virtual_network.ExampleVCN.default_route_table_id}"
  dhcp_options_id     = "${oci_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
}

resource "oci_core_internet_gateway" "ExampleIG" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleIG"
  vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
}

resource "oci_core_default_route_table" "ExampleRT" {
  manage_default_resource_id = "${oci_core_virtual_network.ExampleVCN.default_route_table_id}"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.ExampleIG.id}"
  }
}

# Gets a list of vNIC attachments on the instance
data "oci_core_vnic_attachments" "InstanceVnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  instance_id         = "${oci_core_instance.TFInstance1.id}"
}

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "InstanceVnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.InstanceVnics.vnic_attachments[0],"vnic_id")}"
}

resource "oci_core_instance" "TFInstance1" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFInstance"
  hostname_label      = "instance3"
  shape               = "${var.instance_shape}"
  subnet_id           = "${oci_core_subnet.ExampleSubnet.id}"

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  extended_metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    some_string         = "stringA"
    nested_object       = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
  }
}

resource "null_resource" "remote-exec" {
  depends_on = ["oci_core_instance.TFInstance1"]

  provisioner "remote-exec" {
    connection {
      agent       = false
      timeout     = "30m"
      host        = "${data.oci_core_vnic.InstanceVnic.public_ip_address}"
      user        = "opc"
      private_key = "${var.ssh_private_key}"
    }

    inline = [
      #For more info on accessing metadata see https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/gettingmetadata.htm
      "export STR=$(curl http://169.254.169.254/opc/v1/instance/metadata/nested_object/object/some_string)",

      "if [[ \"$STR\" != \"stringC\" ]]; then exit 1; fi",
    ]
  }
}

output "InstancePublicIP" {
  value = ["${data.oci_core_vnic.InstanceVnic.public_ip_address}"]
}
