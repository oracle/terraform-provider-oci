#!/bin/bash

# Directory to store the terraform files
dir=configure_default_dhcp

# Get required values from terraform output 
vcnid=$(terraform output --json | jq '.VcnId.value[0]')
dns1=$(terraform output --json | jq '.DnsServer1.value[0]')
dns2=$(terraform output --json | jq '.DnsServer2.value[0]')
defaultDHCPId=$(terraform output --json | jq '.DefaultDHCPOptions.value[0]' | sed -e 's/"//g')

# Create the directory
mkdir $dir

# Move into the directory
cd $dir

# Create the env-vars files with only the new env variables to be added
cat > env-vars <<EOF
export TF_VAR_vcn_id=$vcnid
export TF_VAR_dns_server1=$dns1
export TF_VAR_dns_server2=$dns2
EOF

# Create the terraform template file
cat > dhcp.tf <<EOF
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "private_key_password" {}
variable "region" {}

variable "compartment_ocid" {}
variable "vcn_id" {}
variable "dns_server1" {}
variable "dns_server2" {}

provider "oci" {
    tenancy_ocid = "\${var.tenancy_ocid}"
    user_ocid = "\${var.user_ocid}"
    fingerprint = "\${var.fingerprint}"
    private_key_path = "\${var.private_key_path}"
    private_key_password = "\${var.private_key_password}"
    region = "\${var.region}"
}

resource "oci_core_dhcp_options" "DefaultDHCPOptions" {
    compartment_id = "\${var.compartment_ocid}"
    vcn_id = "\${var.vcn_id}"
    display_name = "Default DHCP Options"

    options {
      type = "DomainNameServer"
      server_type = "CustomDnsServer"
      custom_dns_servers = [ "\${var.dns_server1}", "\${var.dns_server2}" ]
    }
}
EOF

# Source the environment variables
. env-vars

# Initialize terraform environment
terraform init

# Import the default DHCP options resource
terraform import oci_core_dhcp_options.DefaultDHCPOptions $defaultDHCPId

# Rest of the steps are manual as the terraform.tfstate file cannot be read by 'jq'. 

cat <<EOF

* Go to $dir directory.
$ cd $dir

* Add the following line to terraform.tfstate as an attribute under "oci_core_dhcp_options.DefaultDHCPOptions"
"vcn_id": $vcnid

* Then run 'terraform plan' and 'terraform apply'
$ terraform plan
$ terraform apply
EOF

