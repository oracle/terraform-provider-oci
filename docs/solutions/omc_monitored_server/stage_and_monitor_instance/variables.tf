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
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaasez4lk2lucxcm52nslj5nhkvbvjtfies4yopwoy4b3vysg5iwjra"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaa2tq67tvbeavcmioghquci6p3pvqwbneq3vfy7fe7m7geiga4cnxa"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaakzrywmh7kwt7ugj5xqi5r4a7xoxsrxtc7nlsdyhmhqyp7ntobjwq"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaalsdgd47nl5tgb55sihdpqmqu2sbvvccjs6tmbkr4nx2pq5gkn63a"
  }
}

variable "compartment_name" {
  description = "Compartment that the OMC managed server will be created in"
}

variable "server_display_name" {
  description = "Display name for your server instance"
}

variable "subnet_id" {
  description = "OCID for the subnet in which the OMC managed server instance will run"
}

#Oracle Management Cloud Specific Variables
variable "omc_registration_key" {
  description = "OMC Agent registration key"
}

variable "omc_tennant_name" {
  description = "OMC Tennant ID"
}

variable "omc_url" {
  description = "OMC Cloud instance"
}

variable "omc_agent_repo_url" {
  description = "OMC Agent source"
}
