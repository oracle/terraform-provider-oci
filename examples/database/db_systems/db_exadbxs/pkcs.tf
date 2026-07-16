// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# PKCS registration example for an ExaDB-XS VM cluster.
#
# 1. Provision this example normally, or import an existing ExaDB-XS VM cluster.
# 2. Complete the Azure or GCP KMS setup for the cluster outside Terraform.
# 3. Apply with pkcs_operation = "register" and pkcs_trigger = 1.
# 4. Increment pkcs_trigger for every later register or unregister operation.
#
# Only one action is selected at a time by the conditional trigger arguments in
# exadb_vm_cluster.tf. The provider obtains tde_key_store_type from the service,
# so it must not be supplied in configuration.
#
# Example terraform.tfvars values:
# pkcs_operation = "register"
# pkcs_trigger   = 1
#
# To unregister instead:
# pkcs_operation = "unregister"
# pkcs_trigger   = 2

# The provider exposes the active keystore and multi-cloud identity connectors
# returned by the ExaDB-XS service.
output "exadb_vm_cluster_tde_key_store_type" {
  value = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.tde_key_store_type
}
