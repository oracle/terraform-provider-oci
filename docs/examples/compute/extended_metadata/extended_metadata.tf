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

provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
    region = "${var.region}"
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

# Gets the OCID of the image. This technique is for example purposes only. The results of oci_core_images may
# change over time for Oracle-provided images, so the only sure way to get the correct OCID is to supply it directly.
data "oci_core_images" "OLImageOCID" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Oracle-Linux-7.4-2017.10.25-0"
}

# Gets a list of vNIC attachments on the instance
data "oci_core_vnic_attachments" "InstanceVnics" { 
    compartment_id = "${var.compartment_ocid}" 
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}" 
    instance_id = "${oci_core_instance.TFInstance1.id}" 
} 

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "InstanceVnic" { 
    vnic_id = "${lookup(data.oci_core_vnic_attachments.InstanceVnics.vnic_attachments[0],"vnic_id")}" 
}

resource "oci_core_instance" "TFInstance1" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}" 
    compartment_id = "${var.compartment_ocid}"
    display_name = "TFInstance"
    hostname_label = "instance3"
    image = "${lookup(data.oci_core_images.OLImageOCID.images[0], "id")}"
    shape = "VM.Standard1.2"
    subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
    extended_metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
        some_string = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
    }
}

resource "null_resource" "remote-exec" {
    depends_on = ["oci_core_instance.TFInstance1"]
    provisioner "remote-exec" {
        connection {
            agent = false
            timeout = "30m"
            host = "${data.oci_core_vnic.InstanceVnic.public_ip_address}"
            user = "opc"
            private_key = "${var.ssh_private_key}"
        }
        inline = [
            #For more info on accessing metadata see https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/gettingmetadata.htm
            "export STR=$(curl http://169.254.169.254/opc/v1/instance/metadata/nested_object/object/some_string)",
            "if [[ \"$STR\" != \"stringC\" ]]; then exit 1; fi"
        ]
    }
}

output "InstancePublicIP" {
    value = ["${data.oci_core_vnic.InstanceVnic.public_ip_address}"]
}


