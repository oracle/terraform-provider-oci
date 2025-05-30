// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apiaccesscontrol

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_apiaccesscontrol_privileged_api_control", ApiaccesscontrolPrivilegedApiControlResource())
	tfresource.RegisterResource("oci_apiaccesscontrol_privileged_api_request", ApiaccesscontrolPrivilegedApiRequestResource())
}
