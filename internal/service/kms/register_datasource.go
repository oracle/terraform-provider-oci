// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_kms_ekms_private_endpoint", KmsEkmsPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_kms_ekms_private_endpoints", KmsEkmsPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_kms_key", KmsKeyDataSource())
	tfresource.RegisterDatasource("oci_kms_key_version", KmsKeyVersionDataSource())
	tfresource.RegisterDatasource("oci_kms_key_versions", KmsKeyVersionsDataSource())
	tfresource.RegisterDatasource("oci_kms_keys", KmsKeysDataSource())
	tfresource.RegisterDatasource("oci_kms_replication_status", KmsReplicationStatusDataSource())
	tfresource.RegisterDatasource("oci_kms_vault", KmsVaultDataSource())
	tfresource.RegisterDatasource("oci_kms_vault_replicas", KmsVaultReplicasDataSource())
	tfresource.RegisterDatasource("oci_kms_vault_usage", KmsVaultUsageDataSource())
	tfresource.RegisterDatasource("oci_kms_vaults", KmsVaultsDataSource())
	tfresource.RegisterDatasource("oci_kms_decrypted_data", KmsDecryptedDataDataSource())
	tfresource.RegisterDatasource("oci_kms_encrypted_data", KmsEncryptedDataDataSource())
}
