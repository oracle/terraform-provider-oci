// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package analytics

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_analytics_analytics_instance", AnalyticsAnalyticsInstanceResource())
	tfresource.RegisterResource("oci_analytics_analytics_instance_private_access_channel", AnalyticsAnalyticsInstancePrivateAccessChannelResource())
	tfresource.RegisterResource("oci_analytics_analytics_instance_vanity_url", AnalyticsAnalyticsInstanceVanityUrlResource())
}
