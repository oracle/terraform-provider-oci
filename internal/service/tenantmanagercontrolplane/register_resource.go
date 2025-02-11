// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_tenantmanagercontrolplane_subscription_mapping", TenantmanagercontrolplaneSubscriptionMappingResource())
}
