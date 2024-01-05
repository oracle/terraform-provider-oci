// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_operator_access_control_operator_control", OperatorAccessControlOperatorControlResource())
	tfresource.RegisterResource("oci_operator_access_control_operator_control_assignment", OperatorAccessControlOperatorControlAssignmentResource())
}
