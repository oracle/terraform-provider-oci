// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_fusion_apps_fusion_environment", FusionAppsFusionEnvironmentResource())
	tfresource.RegisterResource("oci_fusion_apps_fusion_environment_admin_user", FusionAppsFusionEnvironmentAdminUserResource())
	tfresource.RegisterResource("oci_fusion_apps_fusion_environment_data_masking_activity", FusionAppsFusionEnvironmentDataMaskingActivityResource())
	tfresource.RegisterResource("oci_fusion_apps_fusion_environment_family", FusionAppsFusionEnvironmentFamilyResource())
	tfresource.RegisterResource("oci_fusion_apps_fusion_environment_refresh_activity", FusionAppsFusionEnvironmentRefreshActivityResource())
	tfresource.RegisterResource("oci_fusion_apps_fusion_environment_service_attachment", FusionAppsFusionEnvironmentServiceAttachmentResource())
}
