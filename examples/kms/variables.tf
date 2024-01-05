// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "destination_region" {
}

variable "compartment_id" {
}

variable "vault_id" {
}

variable "key_id" {
}

variable "key_display_name" {
  default = "Key C"
}

variable "vault_display_name" {
  default = "Vault C"
}

variable "key_key_shape_algorithm" {
  default = "AES"
}

variable "key_key_shape_length" {
  default = 32
}

variable "crypto_endpoint" {
}

variable "test_rsa_key_id" {
}

variable "test_rsa_key_version" {
}

# Refer https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingkeypairs.htm on how to setup SSH key pairs for compute instances
variable "ssh_public_key" {
}

variable "ssh_private_key" {
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

variable "volume_size" {
  default = "50"
}

variable "vault_type" {
  type    = list(string)
  default = ["DEFAULT", "VIRTUAL_PRIVATE"]
}

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

variable "virtual_vault_ids" {
  type = map(string)

  default = {
    us-phoenix-1 = ""
    us-ashburn-1 = ""
  }
}

variable "destination" {
  type    = list(string)
  default = ["BUCKET", "PRE_AUTHENTICATED_REQUEST_URI"]
}

variable "key_restore_trigger" {
  default = false
}

variable "vault_restore_trigger" {
  default = false
}
