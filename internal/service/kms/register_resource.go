// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_kms_ekms_private_endpoint", KmsEkmsPrivateEndpointResource())
	tfresource.RegisterResource("oci_kms_encrypted_data", KmsEncryptedDataResource())
	tfresource.RegisterResource("oci_kms_generated_key", KmsGeneratedKeyResource())
	tfresource.RegisterResource("oci_kms_key", KmsKeyResource())
	tfresource.RegisterResource("oci_kms_key_version", KmsKeyVersionResource())
	tfresource.RegisterResource("oci_kms_sign", KmsSignResource())
	tfresource.RegisterResource("oci_kms_vault", KmsVaultResource())
	tfresource.RegisterResource("oci_kms_verify", KmsVerifyResource())
	tfresource.RegisterResource("oci_kms_vault_replication", KmsVaultReplicationResource())
}
