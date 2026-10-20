# OpenSearch master user authentication with an OCI Vault secret

This example creates an OpenSearch cluster whose master user password comes from an existing OCI Vault secret. The secret must contain a valid OpenSearch master user password and be accessible to the OpenSearch service.

Set `TF_VAR_region`, `TF_VAR_config_file_profile`, `TF_VAR_compartment_id`, and `TF_VAR_vault_secret_id` for your test tenancy. Authenticate the named OCI CLI security-token profile, then run `terraform init`, `terraform apply`, and `terraform destroy` from this directory.

The example creates a VCN, subnet, and cluster. `terraform destroy` removes those resources; it does not delete the pre-existing Vault secret.
