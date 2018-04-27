variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "SubnetOCID" {}

# Choose an Availability Domain
variable "AD" {
  default = "1"
}

variable "InstanceShape" {
  default = "VM.Standard1.2"
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

variable "BootStrapFile" {
  default = "./userdata/bootstrap"
}


#Chef Configuration Variables

variable "chef_user" {
  description = "User name to access your Chef server"
}
variable "chef_key" {
  description = "Path to Private Key for your chef_user to access Chef server"
}

variable "chef_node_name" {
  description = "Chef Server Node Name, must be unique"
}

variable "chef_recipes" {
  description = "List of recipes for Chef to run"
  type = "list"
}

variable "chef_server" {
  description = "URL for your chef server"
}
