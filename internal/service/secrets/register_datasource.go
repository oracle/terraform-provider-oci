// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package secrets

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_secrets_secretbundle", SecretsSecretbundleDataSource())
	tfresource.RegisterDatasource("oci_secrets_secretbundle_versions", SecretsSecretbundleVersionsDataSource())
}
