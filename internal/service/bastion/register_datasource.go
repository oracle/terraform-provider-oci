// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bastion

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_bastion_bastion", BastionBastionDataSource())
	tfresource.RegisterDatasource("oci_bastion_bastions", BastionBastionsDataSource())
	tfresource.RegisterDatasource("oci_bastion_session", BastionSessionDataSource())
	tfresource.RegisterDatasource("oci_bastion_sessions", BastionSessionsDataSource())
}
