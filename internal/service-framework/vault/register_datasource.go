// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vault

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterFrameworkDataSource() {
	tfresource.RegisterFrameworkDatasource(NewVaultSecretVersionDataSource)
}
