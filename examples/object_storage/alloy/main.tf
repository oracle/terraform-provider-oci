#variable "compartment_ocid" {
#  default = "ocid1.tenancy.oc17..aaaaaaaa5u4pwayzuvledyxkajssb77kpotf3jyew6tuolmhgss4abwpo4ea"
#}

provider "oci" {

}

#data "oci_identity_compartments" "test_compartments" {
#  #Required
#  compartment_id = "ocid1.tenancy.oc17..aaaaaaaa5u4pwayzuvledyxkajssb77kpotf3jyew6tuolmhgss4abwpo4ea"
#}
#data "oci_objectstorage_namespace" "ns" {
#  #Optional
#  compartment_id = "ocid1.tenancy.oc17..aaaaaaaa5u4pwayzuvledyxkajssb77kpotf3jyew6tuolmhgss4abwpo4ea" #data.oci_identity_compartments.test_compartments.compartments.0.compartment_id
#}
resource "oci_objectstorage_bucket" "bucket1" {
  compartment_id = "ocid1.tenancy.oc17..aaaaaaaa5u4pwayzuvledyxkajssb77kpotf3jyew6tuolmhgss4abwpo4ea"# data.oci_identity_compartments.test_compartments.compartments.0.compartment_id
  namespace      = "axefgkprufno"
  name           = "tf-example-alloy"
  access_type    = "NoPublicAccess"
  auto_tiering = "Disabled"
}

resource "oci_objectstorage_preauthrequest" "bucket_par" {
  namespace    =  "axefgkprufno"
  bucket       = oci_objectstorage_bucket.bucket1.name
  name         = "parOnBucket"
  access_type  = "AnyObjectWrite" //Other configurations accepted are ObjectWrite, ObjectRead, ObjectReadWrite, AnyObjectRead, AnyObjectReadWrite,
  time_expires = "2025-12-10T23:00:00Z"
}
/*
resource "oci_apm_apm_domain" "test_apm_domain" {
  #Required
  compartment_id = "ocid1.tenancy.oc17..aaaaaaaa5u4pwayzuvledyxkajssb77kpotf3jyew6tuolmhgss4abwpo4ea"
  display_name   = "var.apm_domain_display_name"

}
*/
output "par_request_url" {
  value = "https://objectstorage.us-dcc-phoenix-1.oraclecloud17.com${oci_objectstorage_preauthrequest.bucket_par.access_uri}"
}