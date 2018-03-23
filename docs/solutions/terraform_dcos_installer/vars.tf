#OCI variables
variable "tenancy_ocid" {}
variable "compartment_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}

variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "region" {
  default = "us-phoenix-1"
}

# Choose an Availability Domain
variable "BootstrapAD" {
    default = "1"
}

variable "BootstrapInstanceShape" {
    default = "VM.Standard1.4"
}

variable "MasterInstanceShape" {
    default = "VM.Standard1.4"
}

variable "AgentInstanceShape" {
    default = "VM.Standard1.4"
}

variable "PublicAgentInstanceShape" {
    default = "VM.Standard1.4"
}

variable "InstanceImageOCID" {
    type = "map"
    default = {
        // CentOS-7-2018.01.04-0
        // See https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/OracleProvidedImageOCIDs.pdf
        us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaajycoi24gyc4tajpwwxjo63yu76cnhtg5a5cfope4tpalnjnhbjqq"
        us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaafrubf4l6e456z4mqn3bj5dpv3s6czfjmyt2m3ukkugzzaosz2fnq"
        eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt1.aaaaaaaaw2qeuo2g4flwz5uieo7hkt6a5wa7ol2z6y23yeqgixcinxmxg7ja"
    }
}

variable "GPUInstanceImageOCID" {
    type = "map"
    default = {
        // Oracle-Linux-7.4-Gen2-GPU-2018.01.20-0
        // See https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/OracleProvidedImageOCIDs.pdf
        us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaa6e56ujmzdgbahcjnz463nkcx7y6eoxjn4eye6yu72tww6r4wxpnq"
        us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaadqerbh7ydhptr353aac3vxbmle7uoadyo7o6asp72nezf4udff6q"
        eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt1.aaaaaaaaadylgm2ydkjz3wj5mzto6koguvamzsfsnjuhhya4ngtz5weud2ca"
    }
}

variable "GPUInstanceShape" {
    default = "BM.GPU2.2"
}

variable "network_cidrs" {
  type = "map"

  default = {
    VCN-CIDR          = "10.0.0.0/16"
    PublicSubnetAD1   = "10.0.10.0/24"
    PublicSubnetAD2   = "10.0.11.0/24"
    PublicSubnetAD3   = "10.0.12.0/24"
  }
}


variable "dcos_cluster_name" {
  description = "Name of your cluster. Alpha-numeric and hyphens only, please."
  default     = "oci-dcos"
}

variable "dcos_installer_url" {
  description = "Path to get DCOS"
  default     = "https://downloads.dcos.io/dcos/EarlyAccess/dcos_generate_config.sh"
}

variable "dcos_master_ad1_count" {
  description = "Number of master nodes. 1, 3, or 5."
  default     = "1"
}

variable "dcos_master_ad2_count" {
  description = "Number of master nodes. 1, 3, or 5."
  default     = "0"
}

variable "dcos_master_ad3_count" {
  description = "Number of master nodes. 1, 3, or 5."
  default     = "0"
}

variable "dcos_agent_ad1_count" {
  description = "Number of agents to deploy in AD1"
  default     = "1"
}

variable "dcos_agent_ad2_count" {
  description = "Number of agents to deploy in AD2"
  default     = "1"
}

variable "dcos_agent_ad3_count" {
  description = "Number of agents to deploy in AD3"
  default     = "1"
}

variable "dcos_public_agent_ad1_count" {
  description = "Number of public agents to deploy"
  default     = "1"
}

variable "dcos_public_agent_ad2_count" {
  description = "Number of public agents to deploy"
  default     = "0"
}

variable "dcos_public_agent_ad3_count" {
  description = "Number of public agents to deploy"
  default     = "0"
}

variable "dcos_gpu_agent_ad1_count" {
    default = "1"
}

variable "dcos_gpu_agent_ad2_count" {
    default = "0"
}

variable "dcos_gpu_agent_ad3_count" {
    default = "0"
}

