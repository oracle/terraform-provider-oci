// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vault

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_vault_secret", VaultSecretDataSource())
	tfresource.RegisterDatasource("oci_vault_secret_version", VaultSecretVersionDataSource())
	tfresource.RegisterDatasource("oci_vault_secrets", VaultSecretsDataSource())
}
