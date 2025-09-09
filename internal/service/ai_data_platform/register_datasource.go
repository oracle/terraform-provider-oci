// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_data_platform

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_ai_data_platform_ai_data_platform", AiDataPlatformAiDataPlatformDataSource())
	tfresource.RegisterDatasource("oci_ai_data_platform_ai_data_platforms", AiDataPlatformAiDataPlatformsDataSource())
}
