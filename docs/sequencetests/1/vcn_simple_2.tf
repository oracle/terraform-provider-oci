variable "private_key_path" {}
variable "fingerprint" {}
variable "region" {}
variable "compartment_ocid" {}
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "SubNet" {
	default = "ocid1.subnet.oc1.iad.aaaaaaaas4zpdv7ucss4dnhyextxiy2iatoyawvpd3lgfdelflbch6xamxza"
}
variable "VCN" {
	default = "ocid1.vcn.oc1.iad.aaaaaaaaeix65p7um76s3rekjl7didnwybufbtmcmlcnjkkumtto5gyyd55q"
}
variable "InstanceShape" {
	default = "VM.Standard1.2"
}
variable "Image_ocid" {
	default = "ocid1.image.oc1.iad.aaaaaaaa52leqhcbinlgzf2lfkhpqbi4k27h2ux7drnrwr53mdpbnnkbuicq"
}
variable "AD" {
	default = "LOil:US-ASHBURN-AD-1"
}

provider "oci" {
	user_ocid = "${var.user_ocid}"
	tenancy_ocid = "${var.tenancy_ocid}"
	region = "${var.region}"
	fingerprint = "${var.fingerprint}"
	private_key_path = "${var.private_key_path}"
}

resource "oci_core_virtual_network" "VCN1" {
	dns_label = "vcn1"
	compartment_id = "${var.compartment_ocid}"
	cidr_block = "10.0.0.0/16"
}

resource "oci_core_internet_gateway" "InternetGateway1" {
	vcn_id = "${oci_core_virtual_network.VCN1.id}"
	compartment_id = "${var.compartment_ocid}"
}

resource "oci_core_subnet" "Subnet1" {
	vcn_id = "${oci_core_virtual_network.VCN1.id}"
	route_table_id = "${oci_core_virtual_network.VCN1.default_route_table_id}"
	security_list_ids = ["${oci_core_virtual_network.VCN1.default_security_list_id}"]
	dhcp_options_id = "${oci_core_virtual_network.VCN1.default_dhcp_options_id}"
	availability_domain = "${var.AD}"
	cidr_block = "10.0.1.0/24"
	compartment_id = "${var.compartment_ocid}"
}

resource "oci_core_subnet" "Subnet2" {
	vcn_id = "${oci_core_virtual_network.VCN1.id}"
	route_table_id = "${oci_core_virtual_network.VCN1.default_route_table_id}"
	security_list_ids = ["${oci_core_virtual_network.VCN1.default_security_list_id}"]
	dhcp_options_id = "${oci_core_virtual_network.VCN1.default_dhcp_options_id}"
	availability_domain = "${var.AD}"
	cidr_block = "10.0.2.0/24"
	compartment_id = "${var.compartment_ocid}"
}
