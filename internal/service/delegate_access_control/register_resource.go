// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package delegate_access_control

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_delegate_access_control_delegation_control", DelegateAccessControlDelegationControlResource())
	tfresource.RegisterResource("oci_delegate_access_control_delegation_subscription", DelegateAccessControlDelegationSubscriptionResource())
}
