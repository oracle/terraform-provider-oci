variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "subnet_ocid" {}

# Choose an Availability Domain
variable "AD" {
    default = "1"
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

# Gets the OCID of the image. This technique is for example purposes only. The results of oci_core_images may
# change over time for Oracle-provided images, so the only sure way to get the correct OCID is to supply it directly.
data "oci_core_images" "OLImageOCID" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Oracle-Linux-7.4-2017.10.25-0"
}

# Create Instance
resource "oci_core_instance" "TFInstance1" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}" 
    compartment_id = "${var.compartment_ocid}"
    display_name = "TFInstance"
    hostname_label = "instance"
    image = "${lookup(data.oci_core_images.OLImageOCID.images[0], "id")}"
    shape = "VM.Standard1.2"
    subnet_id = "${var.subnet_ocid}"
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
