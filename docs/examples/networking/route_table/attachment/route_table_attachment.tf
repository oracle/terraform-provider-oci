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

variable "availability_domain" {
  default = 3
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
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

  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
}

resource "oci_core_route_table_attachment" "ExampleRouteTableAttachment" {
  subnet_id      = "${oci_core_subnet.ExampleSubnet.id}"
  route_table_id = "${oci_core_route_table.ExampleRouteTable.id}"
}

resource "oci_core_private_ip" "TFPrivateIP" {
  vnic_id        = "${oci_core_vnic_attachment.ExampleVnicAttachment.vnic_id}"
  display_name   = "someDisplayName"
  hostname_label = "somehostnamelabel"
}

resource "oci_core_route_table" "ExampleRouteTable" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
  display_name   = "TFExampleRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_private_ip.TFPrivateIP.id}"
  }
}

resource "oci_core_vnic_attachment" "ExampleVnicAttachment" {
  create_vnic_details {
    subnet_id              = "${oci_core_subnet.ExampleSubnet.id}"
    skip_source_dest_check = true
  }

  instance_id = "${oci_core_instance.ExampleInstance.id}"
}

# Create Instance
resource "oci_core_instance" "ExampleInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFInstance"
  hostname_label      = "instance"
  image               = "${var.instance_image_ocid[var.region]}"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id              = "${oci_core_subnet.ExampleSubnet.id}"
    skip_source_dest_check = true
    assign_public_ip       = true
  }
}
