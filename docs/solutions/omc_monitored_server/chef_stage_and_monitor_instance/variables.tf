#OCI Provider Configuration
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}

variable "region" {
  description = "Region to create your instance, valid values are us-phoenix-1, or us-ashburn-1"
}

#SSH access to the server
variable "ssh_public_key" {
  description = "Public key to load onto a server during creation to allow for OPC user ssh access"
}

variable "ssh_private_key" {
  description = "Private key for terraform to use to ssh to the server for post creation instance configuration"
}

#Instance specific variables
variable "ad" {
  description = "Value of 1,2 or 3 expected to represent the AD your start your server instance in"
}

variable "shape_name" {
  description = "OCI server shape common name, find valid values in the OCI console drop down"
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

variable "compartment_name" {
  description = "Compartment that the OMC managed server will be created in"
}

variable "server_display_name" {
  description = "Display name for your server instance"
}

variable "hostname" {
  description = "DNS hostname for your server instance"
}

variable "subnet_id" {
  description = "OCID for the subnet in which the OMC managed server instance will run"
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

variable "json_attributes" {
  description = "Path Chef recipe configuration attributes file"
}

variable "chef_recipes" {
  description = "List of recipes for Chef to run"
  type = "list"
}

variable "chef_server" {
  description = "URL for your chef server"
}

#Oracle Management Cloud Specific Variables
variable "omc_custom_image_id" {
  description = "OCID for the OMC custom image"
  default = "notset"
}
