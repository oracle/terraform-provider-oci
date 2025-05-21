data "oci_database_key_store" "test_key_store" {
  key_store_id = oci_database_key_store.test_key_store.id
}

data "oci_database_key_stores" "test_key_stores" {
  compartment_id = oci_database_key_store.test_key_store.compartment_id
}