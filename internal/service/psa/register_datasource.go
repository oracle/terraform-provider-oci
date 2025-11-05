// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psa

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_psa_private_service_access", PsaPrivateServiceAccessDataSource())
	tfresource.RegisterDatasource("oci_psa_private_service_accesses", PsaPrivateServiceAccessesDataSource())
	tfresource.RegisterDatasource("oci_psa_psa_services", PsaPsaServicesDataSource())
	tfresource.RegisterDatasource("oci_psa_psa_work_request", PsaPsaWorkRequestDataSource())
	tfresource.RegisterDatasource("oci_psa_psa_work_requests", PsaPsaWorkRequestsDataSource())
	tfresource.RegisterDatasource("oci_psa_psa_work_request_errors", PsaWorkRequestErrorsDataSource())
	tfresource.RegisterDatasource("oci_psa_psa_work_request_logs", PsaWorkRequestLogEntriesDataSource())
}
