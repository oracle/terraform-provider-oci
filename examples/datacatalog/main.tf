provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_datacatalog_catalog" "test_oci_datacatalog_catalog" {
  compartment_id                     = var.compartment_id
  attached_catalog_private_endpoints = [oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint.id]
  lifecycle {
    ignore_changes = [
      system_tags,
    ]
  }
}

resource "oci_datacatalog_catalog_private_endpoint" "test_catalog_private_endpoint" {
  #Required
  compartment_id = var.compartment_id
  dns_zones      = ["custpvtsubnet.oraclevcn.com"]
  subnet_id      = oci_core_subnet.test_subnet.id
  lifecycle {
    ignore_changes = [
      system_tags,
    ]
  }
}

data "oci_datacatalog_catalog_private_endpoints" "test_catalog_private_endpoints" {
  #Required
  compartment_id = var.compartment_id
}

data "oci_datacatalog_catalogs" "test_oci_datacatalog_catalogs" {
  compartment_id = var.compartment_ocid
}

resource "oci_datacatalog_data_asset" "test_oci_datacatalog_dataAsset" {
  display_name = "test_data_assets"
  type_key     = var.dataAsset_type_key
  catalog_id   = oci_datacatalog_catalog.test_oci_datacatalog_catalog.id

  properties = {
    "default.host"     = "jbanford-host"
    "default.port"     = "1521"
    "default.database" = "SID"
  }
}

#
resource "oci_datacatalog_connection" "test_connection" {
  catalog_id     = oci_datacatalog_catalog.test_oci_datacatalog_catalog.id
  type_key       = var.connection_type_key
  data_asset_key = oci_datacatalog_data_asset.test_oci_datacatalog_dataAsset.id
  display_name   = "connection_name"

  properties = {
    "default.username" = "scott"
    "default.passwordAndSecrets" = "passwordField"
  }

  enc_properties = {
    "default.password" = var.connection_password
  }

  is_default = true
}

/*
Needs to be used after terraform core issue on datasource response 
data "oci_datacatalog_catalog_types" "test_catalog_types_dataAsset" {
  catalog_id    = oci_datacatalog_catalog.test_oci_datacatalog_catalog.id
  type_category = "dataAsset"
  name          = "Oracle Database"
  state         = "ACTIVE"
}
//data.oci_datacatalog_catalog_types.test_catalog_types_dataAsset.type_collection[0].items[0].key

data "oci_datacatalog_catalog_types" "test_catalog_types_connection" {
  catalog_id    = oci_datacatalog_catalog.test_oci_datacatalog_catalog.id
  type_category = "connection"
  name          = "JDBC"
  state         = "ACTIVE"
}
//data.oci_datacatalog_catalog_types.test_catalog_types_connection.type_collection[0].items[0].key
*/