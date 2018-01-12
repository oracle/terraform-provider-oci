variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

# Choose an Availability Domain
variable "AD" {
    default = "1"
}

variable "InstanceImageOCID" {
  type = "map"
  default = {
    // Oracle-provided image "Oracle-Linux-7.4-2017.12.18-0"
    // See https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/OracleProvidedImageOCIDs.pdf
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
  }
}

provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
	region = "${var.region}"
	disable_auto_retries = true
}

# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_virtual_network" "ExampleVCN" {
	cidr_block = "10.1.0.0/16"
	compartment_id = "${var.compartment_ocid}"
	display_name = "TFExampleVCN"
	dns_label = "tfexamplevcn"
}

resource "oci_core_subnet" "ExampleSubnet" {
	availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
	cidr_block = "10.1.20.0/24"
	display_name = "TFExampleSubnet"
	dns_label = "tfexamplesubnet"
	security_list_ids = ["${oci_core_virtual_network.ExampleVCN.default_security_list_id}"]
	compartment_id = "${var.compartment_ocid}"
	vcn_id = "${oci_core_virtual_network.ExampleVCN.id}"
	route_table_id = "${oci_core_virtual_network.ExampleVCN.default_route_table_id}"
	dhcp_options_id = "${oci_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
}

# Create Instance
resource "oci_core_instance" "TFInstance1" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}" 
    compartment_id = "${var.compartment_ocid}"
    display_name = "TFInstance"
    hostname_label = "instance"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "VM.Standard1.2"
  create_vnic_details {
    subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
  },
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "InstanceVnics" { 
    compartment_id = "${var.compartment_ocid}" 
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}" 
    instance_id = "${oci_core_instance.TFInstance1.id}" 
} 

# Gets the OCID of the first (default) VNIC
data "oci_core_vnic" "InstanceVnic" { 
    vnic_id = "${lookup(data.oci_core_vnic_attachments.InstanceVnics.vnic_attachments[0],"vnic_id")}" 
}

# Create PrivateIP
resource "oci_core_private_ip" "TFPrivateIP" {
	vnic_id = "${lookup(data.oci_core_vnic_attachments.InstanceVnics.vnic_attachments[0],"vnic_id")}"
	display_name = "someDisplayName"
	hostname_label = "somehostnamelabel"
}

# List Private IPs
data "oci_core_private_ips" "privateIpDatasource" {
	depends_on = ["oci_core_private_ip.TFPrivateIP"]
	vnic_id = "${oci_core_private_ip.TFPrivateIP.vnic_id}"
} 

output "privateIPs" {
	value = ["${data.oci_core_private_ips.privateIpDatasource.private_ips}"]
}
