data "oci_datascience_ml_application_implementation_versions" "test_ml_application_implementation_versions" {
  #Required
  ml_application_implementation_id = oci_datascience_ml_application_implementation.test_ml_application_implementation.id
  #Optional
  state = "ACTIVE"
}

data "oci_datascience_ml_application_implementation_version" "test_ml_application_implementation_version" {
  #Required
  ml_application_implementation_version_id = data.oci_datascience_ml_application_implementation_versions.test_ml_application_implementation_versions.ml_application_implementation_version_collection[0].items[0].id
}

data "oci_datascience_ml_application_implementation_versions" "test_ml_application_implementation_versions_with_filter" {
  #Required
  ml_application_implementation_id = oci_datascience_ml_application_implementation.test_ml_application_implementation.id
  state                            = "ACTIVE"
  filter {
    name   = "package_version"
    values = ["1.1"]
  }
}
