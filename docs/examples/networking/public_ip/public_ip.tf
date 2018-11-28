variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}

# variable "compartment_ocid" {}
variable "compartment_ocid" {
  default = "ocid1.compartment.oc1..aaaaaaaa5ndoui3at7y3ge4w6xodngv7dnetng5u4p2uffwfkwsfyl7karmq"
}

variable "region" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

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

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Creates a VCN
resource "oci_core_virtual_network" "TFVcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFVcn"
  dns_label      = "tfVcn"
}

# Creates a subnet
resource "oci_core_subnet" "TFSubnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "TFSubnet"
  dns_label           = "tfsubnet"
  security_list_ids   = ["${oci_core_virtual_network.TFVcn.default_security_list_id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.TFVcn.id}"
  route_table_id      = "${oci_core_virtual_network.TFVcn.default_route_table_id}"
  dhcp_options_id     = "${oci_core_virtual_network.TFVcn.default_dhcp_options_id}"
}

# Creates an instance (without assigning a public IP to the primary private IP on the VNIC)
resource "oci_core_instance" "TFInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFInstance"
  hostname_label      = "instance"
  shape               = "${var.instance_shape}"

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  create_vnic_details {
    assign_public_ip = false
    display_name     = "TFPrimaryVnic"
    subnet_id        = "${oci_core_subnet.TFSubnet.id}"
  }
}

# Creates a secondary VNIC on the instance using a VNIC attachment
resource "oci_core_vnic_attachment" "TFSecondaryVnicAttachment" {
  instance_id  = "${oci_core_instance.TFInstance.id}"
  display_name = "TFTFSecondaryVnicAttachment"

  create_vnic_details {
    assign_public_ip       = false
    display_name           = "TFSecondaryVnic"
    skip_source_dest_check = true
    subnet_id              = "${oci_core_subnet.TFSubnet.id}"
  }
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "TFInstanceVnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  instance_id         = "${oci_core_instance.TFInstance.id}"
}

# Gets the OCID of the first VNIC
data "oci_core_vnic" "TFInstanceVnic1" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.TFInstanceVnics.vnic_attachments[0],"vnic_id")}"
}

# Gets the OCID of the second VNIC
data "oci_core_vnic" "TFInstanceVnic2" {
  vnic_id = "${oci_core_vnic_attachment.TFSecondaryVnicAttachment.vnic_id}"
}

# Gets a list of private IPs on the first VNIC
data "oci_core_private_ips" "TFPrivateIps1" {
  vnic_id = "${data.oci_core_vnic.TFInstanceVnic1.id}"
}

# Gets a list of private IPs on the second VNIC
data "oci_core_private_ips" "TFPrivateIps2" {
  vnic_id = "${data.oci_core_vnic.TFInstanceVnic2.id}"
}

# Creates 3 public IPs: 
#  - Assigned ephemeral public IP (assigned to the first private IP on the first VNIC)
#  - Assigned reserved public IP (assigned to the first private IP on the second VNIC)
#  - Unssigned reserved public IP (available for assignment)

resource "oci_core_public_ip" "EphemeralPublicIPAssigned" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFEphemeralPublicIPAssigned"
  lifetime       = "EPHEMERAL"
  private_ip_id  = "${lookup(data.oci_core_private_ips.TFPrivateIps1.private_ips[0],"id")}"
}

resource "oci_core_public_ip" "ReservedPublicIPAssigned" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFReservedPublicIPAssigned"
  lifetime       = "RESERVED"
  private_ip_id  = "${lookup(data.oci_core_private_ips.TFPrivateIps2.private_ips[0],"id")}"
}

resource "oci_core_public_ip" "ReservedPublicIPUnassigned" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFReservedPublicIPUnassigned"
  lifetime       = "RESERVED"
}

# List public IPs: 
#  - Public IP with availability domain scope (ephemeral).
#  - Public IP with regional scope (reserved).

data "oci_core_public_ips" "AvailabilityDomainPublicIPsList" {
  compartment_id      = "${var.compartment_ocid}"
  scope               = "AVAILABILITY_DOMAIN"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"

  filter {
    name   = "id"
    values = ["${oci_core_public_ip.EphemeralPublicIPAssigned.id}"]
  }
}

data "oci_core_public_ips" "RegionPublicIPsList" {
  compartment_id = "${var.compartment_ocid}"
  scope          = "REGION"

  filter {
    name   = "id"
    values = ["${oci_core_public_ip.ReservedPublicIPAssigned.id}", "${oci_core_public_ip.ReservedPublicIPUnassigned.id}"]
  }
}

output "PublicIPs" {
  value = [
    "${data.oci_core_public_ips.AvailabilityDomainPublicIPsList.public_ips}",
    "${data.oci_core_public_ips.RegionPublicIPsList.public_ips}",
  ]
}
