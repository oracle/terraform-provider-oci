#############################################
# Resource outputs
#############################################

output "distributed_database_id" {
  value       = oci_distributed_database_distributed_database.ddb.id
  description = "OCID of the created Distributed Database (EXADB_XS)"
}

output "wallet_zip_base64" {
  description = "Base64-encoded wallet ZIP returned by the Generate Wallet action"
  value       = oci_distributed_database_distributed_database.ddb.generate_wallet_downloaded_wallet_zip_base64
  sensitive   = true
}

output "wallet_etag" {
  description = "ETag returned by Generate Wallet"
  value       = oci_distributed_database_distributed_database.ddb.generate_wallet_downloaded_wallet_etag
}

output "wallet_last_modified" {
  description = "Last-Modified timestamp of generated wallet"
  value       = oci_distributed_database_distributed_database.ddb.generate_wallet_downloaded_wallet_last_modified
}

output "wallet_content_length" {
  description = "Content-Length of generated wallet ZIP"
  value       = oci_distributed_database_distributed_database.ddb.generate_wallet_downloaded_wallet_content_length
}
#############################################
# Singular datasource outputs
#############################################

output "ds_exascale_distributed_database_by_id" {
  description = "Exascale Distributed Database singular datasource fetched by OCID."
  value       = data.oci_distributed_database_distributed_database.gdd_by_id
}
#############################################
# Plural datasource outputs
#############################################

output "ds_exascale_distributed_databases_list" {
  description = "Exascale Distributed Databases plural datasource (list in compartment)."
  value       = data.oci_distributed_database_distributed_databases.gdd_list
}
#############################################
# Optional: JSON-friendly outputs
#############################################

output "resource_exascale_distributed_database_json" {
  description = "Exascale Distributed Database resource as JSON string."
  value       = jsonencode(oci_distributed_database_distributed_database.ddb)
}
