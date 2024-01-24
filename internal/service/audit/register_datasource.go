// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package audit

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_audit_events", AuditAuditEventsDataSource())
	tfresource.RegisterDatasource("oci_audit_configuration", AuditConfigurationDataSource())
}
