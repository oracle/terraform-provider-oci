## Required connection credentials ##
variable "user_ocid" {
  description = "user OCID for the user account with which to connect to the API."
}

variable "private_key_path" {
  description = "full file path of the private key to use for API access with the user account. Does not support environment variables or ~ to abbreviate a user's home directory."
}

variable "fingerprint" {
  description = "PEM fingerprint for private key."
}

variable "tenancy_ocid" {
  description = "tenancy OCID in which to launch instances."
}

variable "compartment_ocid" {
  description = "compartment OCID in which to launch instances."
}

## Optional parameters ##
variable "instance_count" {
  description = "number of instances to launch."
  default = "1"
}

variable "instance_name_prefix" {
  default = "capacity_reservation_"
}

variable "region" {
  description = "region in which to launch instances."
  default = "us-phoenix-1"
}

# variable "image" {
#   description = "image OCID to launch instances from. This will be unique per region. This default is for region: us-phoenix-1."
#   default = "ocid1.image.oc1.phx.aaaaaaaapxvrtwbpgy3lchk2usn462ekarljwg4zou2acmundxlkzdty4bjq"
# }
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

variable "instance_shape" {
  default = "VM.Standard2.1"
}
