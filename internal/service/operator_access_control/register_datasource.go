// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_operator_access_control_access_request", OperatorAccessControlAccessRequestDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_access_request_audit_log_report", OperatorAccessControlAccessRequestAuditLogReportDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_access_request_history", OperatorAccessControlAccessRequestHistoryDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_access_requests", OperatorAccessControlAccessRequestsDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_operator_action", OperatorAccessControlOperatorActionDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_operator_actions", OperatorAccessControlOperatorActionsDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_operator_control", OperatorAccessControlOperatorControlDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_operator_control_assignment", OperatorAccessControlOperatorControlAssignmentDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_operator_control_assignments", OperatorAccessControlOperatorControlAssignmentsDataSource())
	tfresource.RegisterDatasource("oci_operator_access_control_operator_controls", OperatorAccessControlOperatorControlsDataSource())
}
