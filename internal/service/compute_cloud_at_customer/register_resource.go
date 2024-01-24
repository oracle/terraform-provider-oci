// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package compute_cloud_at_customer

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_compute_cloud_at_customer_ccc_infrastructure", ComputeCloudAtCustomerCccInfrastructureResource())
	tfresource.RegisterResource("oci_compute_cloud_at_customer_ccc_upgrade_schedule", ComputeCloudAtCustomerCccUpgradeScheduleResource())
}
