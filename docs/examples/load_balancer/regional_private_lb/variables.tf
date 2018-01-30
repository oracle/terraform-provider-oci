variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "private_key_password" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

# Choose an Availability Domain
variable "AD1" {
    default = "1"
}

variable "AD2" {
    default = "2"
}

variable "InstanceShape" {
    default = "VM.Standard1.1"
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

variable "vcn_cidr" {
    default = "10.0.0.0/16"
}

variable "mgmt1_subnet_cidr" {
    default = "10.0.0.0/24"
}

variable "mgmt2_subnet_cidr" {
    default = "10.0.1.0/24"
}

variable "lb1_subnet_cidr" {
    default = "10.0.10.0/24"
}

variable "lb2_subnet_cidr" {
    default = "10.0.11.0/24"
}

variable "be1_subnet_cidr" {
    default = "10.0.20.0/24"
}

variable "be2_subnet_cidr" {
    default = "10.0.21.0/24"
}

variable "ha_app_domain" {
    default = "ha.oraclevcn.com"
}

variable "ha_app_name" {
    default = "app"
}

variable "ha_app_port" {
    default = "80"
}

variable "ha_app_protocol" {
    default = "http"
}

variable "backend_port" {
    default = "80"
}

variable "backend_protocol" {
    default = "HTTP"
}

variable "onprem_cidr" {
    default = "172.16.0.0/16"
}

variable "onprem_domain" {
    default = "customer.net"
}

variable "onprem_dns_server1" {
    default = "172.16.0.5"
}

variable "onprem_dns_server2" {
    default = "172.16.31.5"
}

variable "lb_shape" {
    default = "100Mbps"
}
