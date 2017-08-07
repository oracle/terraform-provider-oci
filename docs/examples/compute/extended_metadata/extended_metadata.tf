variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "SubnetOCID" {}

# Choose an Availability Domain
variable "AD" {
    default = "1"
}

provider "baremetal" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
}

# Gets a list of Availability Domains
data "baremetal_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}

# Gets the OCID of the OS image to use
data "baremetal_core_images" "OLImageOCID" {
    compartment_id = "${var.compartment_ocid}"
    operating_system = "Oracle Linux"
    operating_system_version = "7.3"
}

# Gets a list of vNIC attachments on the instance
data "baremetal_core_vnic_attachments" "InstanceVnics" { 
    compartment_id = "${var.compartment_ocid}" 
    availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}" 
    instance_id = "${baremetal_core_instance.TFInstance1.id}" 
} 

# Gets the OCID of the first (default) vNIC
data "baremetal_core_vnic" "InstanceVnic" { 
    vnic_id = "${lookup(data.baremetal_core_vnic_attachments.InstanceVnics.vnic_attachments[0],"vnic_id")}" 
}

resource "baremetal_core_instance" "TFInstance1" {
    availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}" 
    compartment_id = "${var.compartment_ocid}"
    display_name = "TFInstance"
    hostname_label = "instance3"
    image = "${lookup(data.baremetal_core_images.OLImageOCID.images[0], "id")}"
    shape = "VM.Standard1.2"
    subnet_id = "${var.SubnetOCID}"
    extended_metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
        some_string = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
    }
}

resource "null_resource" "remote-exec" {
    depends_on = ["baremetal_core_instance.TFInstance1"]
    provisioner "remote-exec" {
        connection {
            agent = false
            timeout = "30m"
            host = "${data.baremetal_core_vnic.InstanceVnic.public_ip_address}"
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
    value = ["${data.baremetal_core_vnic.InstanceVnic.public_ip_address}"]
}


