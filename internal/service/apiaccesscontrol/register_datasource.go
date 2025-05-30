// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apiaccesscontrol

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_apiaccesscontrol_api_metadata", ApiaccesscontrolApiMetadataDataSource())
	tfresource.RegisterDatasource("oci_apiaccesscontrol_api_metadata_by_entity_types", ApiaccesscontrolApiMetadataByEntityTypesDataSource())
	tfresource.RegisterDatasource("oci_apiaccesscontrol_api_metadatas", ApiaccesscontrolApiMetadatasDataSource())
	tfresource.RegisterDatasource("oci_apiaccesscontrol_privileged_api_control", ApiaccesscontrolPrivilegedApiControlDataSource())
	tfresource.RegisterDatasource("oci_apiaccesscontrol_privileged_api_controls", ApiaccesscontrolPrivilegedApiControlsDataSource())
	tfresource.RegisterDatasource("oci_apiaccesscontrol_privileged_api_request", ApiaccesscontrolPrivilegedApiRequestDataSource())
	tfresource.RegisterDatasource("oci_apiaccesscontrol_privileged_api_requests", ApiaccesscontrolPrivilegedApiRequestsDataSource())
}
