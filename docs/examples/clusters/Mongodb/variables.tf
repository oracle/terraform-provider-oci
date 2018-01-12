variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}


variable "BastionShape" {
    default = "VM.Standard1.1"
}

variable "MongoDBShape" {
    default = "BM.DenseIO1.36"
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

variable "VPC-CIDR" {
    default = "10.0.0.0/26"
}

variable "PubSubnetAD1CIDR" {
    default = "10.0.0.0/28"
}

variable "PrivSubnetAD1CIDR" {
    default = "10.0.0.16/28"
}

variable "PrivSubnetAD2CIDR" {
    default = "10.0.0.32/28"
}

variable "BastSubnetAD1CIDR" {
    default = "10.0.0.48/28"
}

variable "BastionBootStrap" {
    default = "./userdata/bastion"
}

variable "MongoDBBootStrap" {
    default = "./userdata/MongoDB"
}

