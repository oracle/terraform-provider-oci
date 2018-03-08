# configuration.tf - General configuration of SQL Always On deployment
# Set variables here to meet your needs within your particular environment
# Check both this file and sql.tf!

# image_id - OCID of the Windows image to use.
# BE SURE TO UPDATE WITH MOST CURRENT IMAGE.
variable "image_id" {
  default = "ocid1.image.oc1.iad.aaaaaaaajlfsi5npxguvhad3v5d5lu7dc3zcylr2csfdexgd6kor3f6zeqeq"
}

# vcn_dns_name - Set the domain zone, e.g. <vcn_dns_name>.oraclevcn.com .
variable "vcn_dns_name" {
  default = "SQLAlwaysOn"
}

# label_prefix - Set a unique prefix for all resources. Only required if 
# deploying multiple clusters within the same compartment.
variable "label_prefix" {
  default = ""
}

# vcn_cidr_block - Set IP address pool for the VCN being created for SQL AO.
variable "vcn_cidr_block" {
  default = "10.0.0.0/19"
}

# compartment_name - Set the compartment to use for the configuration.
# If the compartment does not exist, it will be created.
# NOTE: Compartments CANNOT be destroyed once created
variable "compartment_name" {
  default = "SQLAlwaysOn"
}

# compartment_description - Set a short description for the compartment. 
# Only used for new compartments.
variable "compartment_description" {
  default = "Compartment created for SQL Always On deployment"
}

# ad_count - Set number of ADs to use. Minimum 2, Maximum 3.
variable "ad_count" {
  default = "3"
}

# local_dns_server - Set the address to use for the local DNS server.
variable "local_dns_server" {
  default = "1.1.1.1"
}

# Adminsitrative instance shapes definition.
# Change only if necessary.
# dmz_shape - Shape to use for DMZ server
variable "dmz_shape" {
  default = "VM.Standard1.2"
}

# admin_shape - Shape to use for Administrative server.
variable "admin_shape" {
  default = "VM.Standard1.4"
}
