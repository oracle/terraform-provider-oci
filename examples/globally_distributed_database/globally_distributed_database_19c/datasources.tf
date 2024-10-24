data "oci_globally_distributed_database_private_endpoint" "this" {
  #Required
  private_endpoint_id = oci_globally_distributed_database_private_endpoint.this.id
}

data "oci_globally_distributed_database_private_endpoints" "this" {
  #Required
  compartment_id = oci_globally_distributed_database_private_endpoint.this.compartment_id

  #Optional
  display_name = oci_globally_distributed_database_private_endpoint.this.display_name
  state        = oci_globally_distributed_database_private_endpoint.this.state
}

data "oci_globally_distributed_database_sharded_database" "this" {
  #Required
  sharded_database_id = oci_globally_distributed_database_sharded_database.this.id

  #Optional
  metadata = "test" //oci_globally_distributed_database_sharded_database.this.metadata
}

data "oci_globally_distributed_database_sharded_databases" "this" {
  #Required
  compartment_id = oci_globally_distributed_database_sharded_database.this.compartment_id

  #Optional
  display_name = oci_globally_distributed_database_sharded_database.this.display_name
  state        = oci_globally_distributed_database_sharded_database.this.state
}