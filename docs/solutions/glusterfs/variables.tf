// These settings can be populated here or read from your env-vars settings

// Settings for authentication
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "region" {}
variable "private_key_path" {}
variable "private_key_password" {}

variable "compartment_ocid" {}

// The SSH public key for connecting to the compute instances
variable "ssh_public_key" {}

// The name DNS label to use for the VCN
variable "DnsLabel" {}

variable "ServerInstanceShape" {
  default = "BM.DenseIO1.36"
}

variable "ClientInstanceShape" {
  default = "VM.Standard1.2"
}

variable "ClientInstanceImageOCID" {
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

variable "ServerInstanceImageOCID" {
  type = "map"
  default = {
    // Oracle-provided image "CentOS-7-2018.01.04-0"
    // See https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/OracleProvidedImageOCIDs.pdf
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaajycoi24gyc4tajpwwxjo63yu76cnhtg5a5cfope4tpalnjnhbjqq"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaafrubf4l6e456z4mqn3bj5dpv3s6czfjmyt2m3ukkugzzaosz2fnq"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt1.aaaaaaaaw2qeuo2g4flwz5uieo7hkt6a5wa7ol2z6y23yeqgixcinxmxg7ja"
  }
}

variable "ServerBootStrapFile" {
  default = "./userdata/bootstrap-server.sh"
}

variable "ClientBootStrapFile" {
  default = "./userdata/bootstrap-client.sh"
}
