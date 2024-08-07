// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_jms_fleet", JmsFleetResource())
	tfresource.RegisterResource("oci_jms_fleet_advanced_feature_configuration", JmsFleetAdvancedFeatureConfigurationResource())
	tfresource.RegisterResource("oci_jms_jms_plugin", JmsJmsPluginResource())
}
