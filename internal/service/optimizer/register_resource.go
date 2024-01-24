// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_optimizer_enrollment_status", OptimizerEnrollmentStatusResource())
	tfresource.RegisterResource("oci_optimizer_profile", OptimizerProfileResource())
	tfresource.RegisterResource("oci_optimizer_recommendation", OptimizerRecommendationResource())
	tfresource.RegisterResource("oci_optimizer_resource_action", OptimizerResourceActionResource())
}
