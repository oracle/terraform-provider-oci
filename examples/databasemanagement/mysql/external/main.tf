variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
   //version = "5.36.0"
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "compartment_id" {  
  default =  "ocid1.test.oc1..<unique_ID>EXAMPLE-compartmentId-Value"
}

/*Creates an external MySQL database resource */
resource "oci_database_management_external_my_sql_database" "test_external_my_sql_database" {
  #Required
  compartment_id = var.compartment_id
  db_name        = "ExampleNameTest"
}

data "oci_database_management_external_my_sql_databases" "test_external_my_sql_databases" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = "ExampleNameTest"
}

/*Creates a database connecotor resource. Connector requires an external MySQL database to associate with. */
resource "oci_database_management_external_my_sql_database_connector" "test_external_my_sql_database_connector" {
  #Required
  compartment_id = "ocid1.test.oc1..<unique_ID>EXAMPLE-compartmentId-Value"
  connector_details {
    #Required
    credential_type      = "MYSQL_EXTERNAL_NON_SSL_CREDENTIALS"
    display_name         = "EXAMPLE-Name-Test"
    external_database_id = "ocid1.test.oc1..<unique_ID>EXAMPLE-externalDatabase-Value"
    host_name            = "exampleHost"
    macs_agent_id        = "ocid1.test.oc1..<unique_ID>EXAMPLE-agent-Value"
    network_protocol     = "TCP"
    port                 = "10"
    ssl_secret_id        = "ocid1.test.oc1..<unique_ID>EXAMPLE-secret-Value"
  }
  is_test_connection_param = "false"
}

data "oci_database_management_external_my_sql_database_connectors" "test_external_my_sql_database_connectors" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = "ExampleTest"
}

/*To enable/disable database management, provide an external MySQL database OCID and connector OCID. 
  "true" to enable and "false" to disable. In this example database management is being enabled for the external MySQL database 
  using the given connector.*/
resource "oci_database_management_external_my_sql_database_external_mysql_databases_management" "test_external_my_sql_database_external_mysql_databases_management" {
  #Required
 
  external_my_sql_database_id    = "ocid1.test.oc1..<unique_ID>EXAMPLE-database-Value"
  
  enable_external_mysql_database = "true"

  #Optional
  connector_id ="ocid1.test.oc1..<unique_ID>EXAMPLE-connector-Value"
}
