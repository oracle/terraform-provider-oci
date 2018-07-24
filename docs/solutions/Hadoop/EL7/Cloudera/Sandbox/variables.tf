###
## Variables here are sourced from env, but still need to be initialized for Terraform
###

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" { default = "us-phoenix-1" }

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "AD" { default = "2" }

variable "boot_volume_size" { default = "256" }

variable "InstanceImageOCID" {
    type = "map"
    default = {
        // See https://docs.us-phoenix-1.oraclecloud.com/images/
        // Oracle-provided image "CentOS-7.5-2018.06.22-0"
        eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaasdvfvvgzjhqpuwmjbypgovachdgwvcvus5n4p64fajmbassg2pqa"
        us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaa5o7kjzy7gqtmu5pxuhnh6yoi3kmzazlk65trhpjx5xg3hfbuqvgq"
        uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa74er3gyrjg3fiesftpc42viplbhp7gdafqzv33kyyx3jrazruta"
        us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaapnnv2phiyw7apcgtg6kmn572b2mux56ll6j6mck5xti3aw4bnwrq"
    }
}

variable "blocksize_in_gbs" { default = "1024" }
