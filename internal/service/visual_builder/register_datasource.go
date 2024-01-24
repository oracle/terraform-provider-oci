// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package visual_builder

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_visual_builder_vb_instance", VisualBuilderVbInstanceDataSource())
	tfresource.RegisterDatasource("oci_visual_builder_vb_instances", VisualBuilderVbInstancesDataSource())
	tfresource.RegisterDatasource("oci_visual_builder_vb_instance_applications", VisualBuilderVbInstanceApplicationsDataSource())
}
