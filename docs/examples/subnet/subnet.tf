variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  private_key = "${var.private_key}"
}

# Because you can specify multiple security lists/subnet the security_list_ids value must be specified as a list in []'s.
# See https://www.terraform.io/docs/configuration/syntax.html

# Generally you wouldn't specify a subnet without first specifying a VNC. Once the VNC has been created you would get the compartment_id, vcn_id, route_table_id, and security_list_id(s) from that resource and use Terraform attributes below to populate those values. 
# See https://www.terraform.io/docs/configuration/interpolation.html

resource "baremetal_core_subnet" "a_TF_managed_subnet" {
  availability_domain = "<some AD>"
  cidr_block = "10.0.1.0/24"
  display_name = "a_TF_managed_subnet"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "<the ocid of the VCN to create the subnet in>"
  route_table_id = "<the ocid of the route table to attach to the subnet>"
  security_list_ids = ["<the ocid(s) of the security list(s) to attach to the subnet>"]
}
