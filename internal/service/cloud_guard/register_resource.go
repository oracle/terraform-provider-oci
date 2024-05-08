// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_cloud_guard_adhoc_query", CloudGuardAdhocQueryResource())
	tfresource.RegisterResource("oci_cloud_guard_cloud_guard_configuration", CloudGuardCloudGuardConfigurationResource())
	tfresource.RegisterResource("oci_cloud_guard_data_mask_rule", CloudGuardDataMaskRuleResource())
	tfresource.RegisterResource("oci_cloud_guard_data_source", CloudGuardDataSourceResource())
	tfresource.RegisterResource("oci_cloud_guard_detector_recipe", CloudGuardDetectorRecipeResource())
	tfresource.RegisterResource("oci_cloud_guard_managed_list", CloudGuardManagedListResource())
	tfresource.RegisterResource("oci_cloud_guard_responder_recipe", CloudGuardResponderRecipeResource())
	tfresource.RegisterResource("oci_cloud_guard_saved_query", CloudGuardSavedQueryResource())
	tfresource.RegisterResource("oci_cloud_guard_security_recipe", CloudGuardSecurityRecipeResource())
	tfresource.RegisterResource("oci_cloud_guard_security_zone", CloudGuardSecurityZoneResource())
	tfresource.RegisterResource("oci_cloud_guard_target", CloudGuardTargetResource())
	tfresource.RegisterResource("oci_cloud_guard_wlp_agent", CloudGuardWlpAgentResource())
}
