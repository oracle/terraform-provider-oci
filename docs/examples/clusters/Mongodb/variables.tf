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
        // See https://docs.us-phoenix-1.oraclecloud.com/images/
        // Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
        us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
        us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
        eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
        uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
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

