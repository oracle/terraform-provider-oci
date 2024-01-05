// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package autoscaling

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_autoscaling_auto_scaling_configuration", AutoScalingAutoScalingConfigurationDataSource())
	tfresource.RegisterDatasource("oci_autoscaling_auto_scaling_configurations", AutoScalingAutoScalingConfigurationsDataSource())
}
