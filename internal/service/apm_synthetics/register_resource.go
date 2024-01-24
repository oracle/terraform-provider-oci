// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_apm_synthetics_dedicated_vantage_point", ApmSyntheticsDedicatedVantagePointResource())
	tfresource.RegisterResource("oci_apm_synthetics_monitor", ApmSyntheticsMonitorResource())
	tfresource.RegisterResource("oci_apm_synthetics_on_premise_vantage_point", ApmSyntheticsOnPremiseVantagePointResource())
	tfresource.RegisterResource("oci_apm_synthetics_on_premise_vantage_point_worker", ApmSyntheticsOnPremiseVantagePointWorkerResource())
	tfresource.RegisterResource("oci_apm_synthetics_script", ApmSyntheticsScriptResource())
}
