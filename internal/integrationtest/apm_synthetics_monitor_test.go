// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ApmSyntheticsMonitorSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"monitor_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_monitor.test_monitor.id}`},
	}

	ApmSyntheticsMonitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"monitor_type":  acctest.Representation{RepType: acctest.Optional, Create: `SCRIPTED_BROWSER`},
		"script_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_apm_synthetics_script.test_script.id}`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorDataSourceFilterRepresentation},
	}

	ApmSyntheticsMonitorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `display_name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_synthetics_monitor.test_monitor.display_name}`}},
	}

	ApmSyntheticsMonitorVantagePointsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `OraclePublic-us-ashburn-1`},
	}

	ApmSyntheticsMonitorAvailabilityConfigurationRepresentation = map[string]interface{}{
		"max_allowed_failures_per_interval": acctest.Representation{RepType: acctest.Optional, Create: `0`},
		"min_allowed_runs_per_interval":     acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}
	TimeStarted1                                                = time.Now().AddDate(0, 1, 0).UTC().Format(time.RFC3339)
	TimeStarted2                                                = time.Now().AddDate(0, 1, 1).UTC().Format(time.RFC3339)
	TimeEnded1                                                  = time.Now().AddDate(0, 1, 2).UTC().Format(time.RFC3339)
	TimeEnded2                                                  = time.Now().AddDate(0, 1, 3).UTC().Format(time.RFC3339)
	ApmSyntheticsMonitorMaintenanceWindowScheduleRepresentation = map[string]interface{}{
		"time_ended":   acctest.Representation{RepType: acctest.Optional, Create: TimeEnded1, Update: TimeEnded2},
		"time_started": acctest.Representation{RepType: acctest.Optional, Create: TimeStarted1, Update: TimeStarted2},
	}

	ApmSyntheticsMonitorConfigurationDnsConfigurationRepresentation = map[string]interface{}{
		"is_override_dns": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"override_dns_ip": acctest.Representation{RepType: acctest.Optional, Create: `12.1.21.1`, Update: `12.1.21.2`},
	}

	ApmSyntheticsMonitorConfigurationNetworkConfigurationRepresentation = map[string]interface{}{
		"number_of_hops":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"probe_mode":        acctest.Representation{RepType: acctest.Optional, Create: `SACK`, Update: `SYN`},
		"probe_per_hop":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `9`},
		"protocol":          acctest.Representation{RepType: acctest.Optional, Create: `TCP`},
		"transmission_rate": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	ApmSyntheticsMonitorScriptParametersRepresentation = map[string]interface{}{
		"param_name":  acctest.Representation{RepType: acctest.Required, Create: `testName`, Update: `testName`},
		"param_value": acctest.Representation{RepType: acctest.Required, Create: `myTest`, Update: `myTest1`},
	}

	ApmSyntheticsBrowserMonitorResourceConfig = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsBrowserMonitorRepresentation)

	ApmSyntheticsBrowserMonitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"monitor_type":  acctest.Representation{RepType: acctest.Optional, Create: `BROWSER`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorDataSourceFilterRepresentation}}

	ApmSyntheticsBrowserMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":               acctest.Representation{RepType: acctest.Required, Create: `BROWSER`},
		"repeat_interval_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorVantagePointsRepresentation},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                     acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                     acctest.Representation{RepType: acctest.Optional, Create: `https://console.us-ashburn-1.oraclecloud.com`, Update: `https://console.us-phoenix-1.oraclecloud.com`},
		"timeout_in_seconds":         acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
	}

	ApmSyntheticsRestMonitorResourceConfig = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsRestMonitorRepresentation)

	ApmSyntheticsRestMonitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"monitor_type":  acctest.Representation{RepType: acctest.Optional, Create: `REST`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorDataSourceFilterRepresentation},
	}

	ApmSyntheticsRestMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":               acctest.Representation{RepType: acctest.Required, Create: `REST`},
		"repeat_interval_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorVantagePointsRepresentation},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                     acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                     acctest.Representation{RepType: acctest.Optional, Create: `https://console.us-ashburn-1.oraclecloud.com`, Update: `https://console.us-phoenix-1.oraclecloud.com`},
		"timeout_in_seconds":         acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
		"configuration":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsRestMonitorConfigurationRepresentation},
	}

	ApmSyntheticsRestMonitorConfigurationRepresentation = map[string]interface{}{
		"config_type":                       acctest.Representation{RepType: acctest.Optional, Create: `REST_CONFIG`},
		"is_certificate_validation_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_failure_retried":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_redirection_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"req_authentication_details":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsRestMonitorConfigurationReqAuthenticationDetailsRepresentation},
		"req_authentication_scheme":         acctest.Representation{RepType: acctest.Optional, Create: `OAUTH`},
		"request_headers":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsRestMonitorConfigurationRequestHeadersRepresentation},
		"request_method":                    acctest.Representation{RepType: acctest.Optional, Create: `POST`},
		"request_post_body":                 acctest.Representation{RepType: acctest.Optional, Create: `requestPostBody`, Update: `requestPostBody2`},
		"request_query_params":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsRestMonitorConfigurationRequestQueryParamsRepresentation},
		"verify_response_codes":             acctest.Representation{RepType: acctest.Optional, Create: []string{`200`, `300`, `400`}},
		"verify_response_content":           acctest.Representation{RepType: acctest.Optional, Create: `verifyResponseContent`, Update: `verifyResponseContent2`},
	}

	ApmSyntheticsRestMonitorConfigurationReqAuthenticationDetailsRepresentation = map[string]interface{}{
		"auth_headers":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsRestMonitorConfigurationReqAuthenticationDetailsAuthHeadersRepresentation},
		"auth_request_method":    acctest.Representation{RepType: acctest.Optional, Create: `POST`},
		"auth_request_post_body": acctest.Representation{RepType: acctest.Optional, Create: `authRequestPostBody`, Update: `authRequestPostBody2`},
		"auth_url":               acctest.Representation{RepType: acctest.Optional, Create: `http://authUrl`, Update: `http://authUrl2`},
		"oauth_scheme":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}
	ApmSyntheticsRestMonitorConfigurationRequestHeadersRepresentation = map[string]interface{}{
		"header_name":  acctest.Representation{RepType: acctest.Optional, Create: `content-type`},
		"header_value": acctest.Representation{RepType: acctest.Optional, Create: `json`},
	}
	ApmSyntheticsRestMonitorConfigurationRequestQueryParamsRepresentation = map[string]interface{}{
		"param_name":  acctest.Representation{RepType: acctest.Optional, Create: `paramName`, Update: `paramName2`},
		"param_value": acctest.Representation{RepType: acctest.Optional, Create: `paramValue`, Update: `paramValue2`},
	}

	ApmSyntheticsRestMonitorConfigurationReqAuthenticationDetailsAuthHeadersRepresentation = map[string]interface{}{
		"header_name":  acctest.Representation{RepType: acctest.Optional, Create: `content-type`},
		"header_value": acctest.Representation{RepType: acctest.Optional, Create: `json`},
	}

	ApmSyntheticsScriptedBrowserMonitorRequiredOnlyResource = ApmSyntheticsScriptedMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsScriptedBrowserMonitorRepresentation)

	ApmSyntheticsScriptedBrowserMonitorResourceConfig = ApmSyntheticsScriptedMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsScriptedBrowserMonitorRepresentation)

	ApmSyntheticsScriptedBrowserMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":                acctest.Representation{RepType: acctest.Required, Create: `SCRIPTED_BROWSER`},
		"repeat_interval_in_seconds":  acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorVantagePointsRepresentation},
		"availability_configuration":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorAvailabilityConfigurationRepresentation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_run_once":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_run_now":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"maintenance_window_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorMaintenanceWindowScheduleRepresentation},
		"scheduling_policy":           acctest.Representation{RepType: acctest.Optional, Create: `ALL`, Update: `ROUND_ROBIN`},
		"script_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_apm_synthetics_script.test_script.id}`},
		"status":                      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                      acctest.Representation{RepType: acctest.Optional, Create: `https://console.us-ashburn-1.oraclecloud.com`, Update: `https://console.us-phoenix-1.oraclecloud.com`},
		"timeout_in_seconds":          acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
		"configuration":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsScriptedBrowserMonitorConfigurationRepresentation},
		"script_parameters":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorScriptParametersRepresentation},
	}

	ApmSyntheticsScriptedBrowserMonitorConfigurationRepresentation = map[string]interface{}{
		"dns_configuration":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorConfigurationDnsConfigurationRepresentation},
		"config_type":                       acctest.Representation{RepType: acctest.Optional, Create: `SCRIPTED_BROWSER_CONFIG`},
		"is_certificate_validation_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_default_snapshot_enabled":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_failure_retried":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"network_configuration":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorConfigurationNetworkConfigurationRepresentation},
	}

	ApmSyntheticsNetworkMonitorRequiredOnlyResource = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsNetworkMonitorRepresentation)

	ApmSyntheticsNetworkMonitorResourceConfig = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsNetworkMonitorRepresentation)

	ApmSyntheticsNetworkMonitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"monitor_type":  acctest.Representation{RepType: acctest.Optional, Create: `NETWORK`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorDataSourceFilterRepresentation}}

	ApmSyntheticsNetworkMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":                acctest.Representation{RepType: acctest.Required, Create: `NETWORK`},
		"repeat_interval_in_seconds":  acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorVantagePointsRepresentation},
		"availability_configuration":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorAvailabilityConfigurationRepresentation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_run_once":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_run_now":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"maintenance_window_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorMaintenanceWindowScheduleRepresentation},
		"scheduling_policy":           acctest.Representation{RepType: acctest.Optional, Create: `ALL`, Update: `ROUND_ROBIN`},
		"status":                      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                      acctest.Representation{RepType: acctest.Optional, Create: `www.oracle.com:80`, Update: `www.oracle.com:80`},
		"timeout_in_seconds":          acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
		"configuration":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsNetworkMonitorConfigurationRepresentation},
	}

	ApmSyntheticsNetworkMonitorConfigurationRepresentation = map[string]interface{}{
		"dns_configuration":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorConfigurationDnsConfigurationRepresentation},
		"config_type":           acctest.Representation{RepType: acctest.Optional, Create: `NETWORK_CONFIG`},
		"is_failure_retried":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"network_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorConfigurationNetworkConfigurationRepresentation},
	}

	ApmSyntheticsApmSyntheticsDNSMonitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"monitor_type":  acctest.Representation{RepType: acctest.Optional, Create: `DNS`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorDataSourceFilterRepresentation}}

	ApmSyntheticsDNSSecMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":                acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"repeat_interval_in_seconds":  acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorVantagePointsRepresentation},
		"availability_configuration":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorAvailabilityConfigurationRepresentation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_run_once":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_run_now":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"maintenance_window_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorMaintenanceWindowScheduleRepresentation},
		"scheduling_policy":           acctest.Representation{RepType: acctest.Optional, Create: `ALL`, Update: `ROUND_ROBIN`},
		"status":                      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                      acctest.Representation{RepType: acctest.Optional, Create: `www.oracle.com`, Update: `www.oracle.com`},
		"timeout_in_seconds":          acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
		"configuration":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsDNSSecMonitorConfigurationRepresentation},
	}

	ApmSyntheticsDNSSecMonitorRequiredOnlyResource = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsDNSSecMonitorRepresentation)

	ApmSyntheticsDNSSecMonitorResourceConfig = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSSecMonitorRepresentation)

	ApmSyntheticsDNSSecMonitorConfigurationRepresentation = map[string]interface{}{
		"dns_configuration":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorConfigurationDnsConfigurationRepresentation},
		"config_type":        acctest.Representation{RepType: acctest.Optional, Create: `DNSSEC_CONFIG`},
		"is_failure_retried": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"record_type":        acctest.Representation{RepType: acctest.Optional, Create: `A`, Update: `ANY`},
	}

	ApmSyntheticsDNSServerMonitorRequiredOnlyResource = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsDNSServerMonitorRepresentation)

	ApmSyntheticsDNSServerMonitorResourceConfig = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSServerMonitorRepresentation)

	ApmSyntheticsDNSServerMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":                acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"repeat_interval_in_seconds":  acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorVantagePointsRepresentation},
		"availability_configuration":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorAvailabilityConfigurationRepresentation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_run_once":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_run_now":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"maintenance_window_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorMaintenanceWindowScheduleRepresentation},
		"scheduling_policy":           acctest.Representation{RepType: acctest.Optional, Create: `ALL`, Update: `ROUND_ROBIN`},
		"status":                      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                      acctest.Representation{RepType: acctest.Optional, Create: `www.oracle.com`, Update: `www.oracle.com`},
		"timeout_in_seconds":          acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
		"configuration":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsDNSServerMonitorConfigurationRepresentation},
	}

	ApmSyntheticsDNSServerMonitorConfigurationRepresentation = map[string]interface{}{
		"dns_configuration":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorConfigurationDnsConfigurationRepresentation},
		"config_type":           acctest.Representation{RepType: acctest.Optional, Create: `DNS_SERVER_CONFIG`},
		"is_failure_retried":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_query_recursive":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"name_server":           acctest.Representation{RepType: acctest.Optional, Create: `8.8.8.8`, Update: `8.8.4.4`},
		"network_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorConfigurationNetworkConfigurationRepresentation},
		"protocol":              acctest.Representation{RepType: acctest.Optional, Create: `TCP`, Update: `UDP`},
		"record_type":           acctest.Representation{RepType: acctest.Optional, Create: `A`, Update: `ANY`},
	}

	ApmSyntheticsDNSTraceMonitorRequiredOnlyResource = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsDNSTraceMonitorRepresentation)

	ApmSyntheticsDNSTraceMonitorResourceConfig = ApmSyntheticsMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSTraceMonitorRepresentation)

	ApmSyntheticsDNSTraceMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":                acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"repeat_interval_in_seconds":  acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsMonitorVantagePointsRepresentation},
		"availability_configuration":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorAvailabilityConfigurationRepresentation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_run_once":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_run_now":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"maintenance_window_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorMaintenanceWindowScheduleRepresentation},
		"scheduling_policy":           acctest.Representation{RepType: acctest.Optional, Create: `ALL`, Update: `ROUND_ROBIN`},
		"status":                      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                      acctest.Representation{RepType: acctest.Optional, Create: `www.oracle.com`, Update: `www.oracle.com`},
		"timeout_in_seconds":          acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
		"configuration":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsDNSTraceMonitorConfigurationRepresentation},
	}

	ApmSyntheticsDNSTraceMonitorConfigurationRepresentation = map[string]interface{}{
		"dns_configuration":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmSyntheticsMonitorConfigurationDnsConfigurationRepresentation},
		"config_type":        acctest.Representation{RepType: acctest.Optional, Create: `DNS_TRACE_CONFIG`},
		"is_failure_retried": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"protocol":           acctest.Representation{RepType: acctest.Optional, Create: `TCP`, Update: `UDP`},
		"record_type":        acctest.Representation{RepType: acctest.Optional, Create: `A`, Update: `ANY`},
	}

	ApmSyntheticsScriptedMonitorResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Optional, acctest.Create, ApmSyntheticsscriptRepresentation)

	ApmSyntheticsMonitorResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsScriptedBrowserMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsScriptedBrowserMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsScriptedMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsScriptedBrowserMonitorRepresentation), "apmsynthetics", "monitor", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsScriptedMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsScriptedBrowserMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "SCRIPTED_BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_default_snapshot_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.number_of_hops", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_mode", "SACK"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_per_hop", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.transmission_rate", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted1),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
				resource.TestCheckResourceAttrSet(resourceName, "script_id"),
				resource.TestCheckResourceAttrSet(resourceName, "script_name"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_name", "testName"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_value", "myTest"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-ashburn-1.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsScriptedMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsScriptedBrowserMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "SCRIPTED_BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_default_snapshot_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.number_of_hops", "11"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_mode", "SYN"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_per_hop", "9"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.transmission_rate", "11"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded2),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted2),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttrSet(resourceName, "script_id"),
				resource.TestCheckResourceAttrSet(resourceName, "script_name"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_name", "testName"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_value", "myTest1"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, ApmSyntheticsMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsScriptedMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsScriptedBrowserMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttrSet(datasourceName, "script_id"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsScriptedBrowserMonitorResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "SCRIPTED_BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_default_snapshot_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.number_of_hops", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.probe_mode", "SYN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.probe_per_hop", "9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.transmission_rate", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_started"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(singularDatasourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0.name"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmSyntheticsScriptedBrowserMonitorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
				"batch_interval_in_seconds", //ignore as it does not apply to this case
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsBrowserMonitorResource(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsMonitorResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsBrowserMonitorRepresentation), "apmsynthetics", "monitor", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmSyntheticsMonitorDestroy,
		Steps: []resource.TestStep{

			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsBrowserMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "BROWSER_CONFIG"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_type", "BROWSER"),
					resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-ashburn-1.oraclecloud.com"),
					resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
					resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsBrowserMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "BROWSER_CONFIG"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_type", "BROWSER"),
					resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
					resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
					resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, ApmSyntheticsBrowserMonitorDataSourceRepresentation) +
					compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsBrowserMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "monitor_type", "BROWSER"),
					resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

					resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsMonitorSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ApmSyntheticsBrowserMonitorResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "BROWSER_CONFIG"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "BROWSER"),
					resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
					resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
					resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),
				),
			},
			// verify resource import
			{
				Config:            config + ApmSyntheticsScriptedBrowserMonitorRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"apm_domain_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsRestMonitorResource(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsRestMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsRestMonitorRepresentation), "apmsynthetics", "monitor", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsRestMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "REST_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_redirection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_url", "http://authUrl"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_scheme", "OAUTH"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_post_body", "requestPostBody"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_name", "paramName"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_value", "paramValue"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_codes.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_content", "verifyResponseContent"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "REST"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-ashburn-1.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsRestMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "REST_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_redirection_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.auth_url", "http://authUrl2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.req_authentication_scheme", "OAUTH"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_post_body", "requestPostBody2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_name", "paramName2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.request_query_params.0.param_value", "paramValue2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_codes.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.verify_response_content", "verifyResponseContent2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "REST"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, ApmSyntheticsRestMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsRestMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "REST"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsRestMonitorResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "REST_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_redirection_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_request_method", "POST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_request_post_body", "authRequestPostBody2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.auth_url", "http://authUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_details.0.oauth_scheme", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.req_authentication_scheme", "OAUTH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.0.header_name", "content-type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_headers.0.header_value", "json"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_method", "POST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_post_body", "requestPostBody2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.0.param_name", "paramName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.request_query_params.0.param_value", "paramValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.verify_response_codes.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.verify_response_content", "verifyResponseContent2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "REST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmSyntheticsScriptedBrowserMonitorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsNetworkMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsNetworkMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsNetworkMonitorRepresentation), "apmsynthetics", "monitor", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsNetworkMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "NETWORK_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.number_of_hops", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_mode", "SACK"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_per_hop", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.transmission_rate", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted1),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "NETWORK"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "www.oracle.com:80"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsNetworkMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "NETWORK_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.number_of_hops", "11"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_mode", "SYN"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_per_hop", "9"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.transmission_rate", "11"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded2),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted2),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "NETWORK"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "www.oracle.com:80"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, ApmSyntheticsNetworkMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsNetworkMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "NETWORK"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsNetworkMonitorResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "NETWORK_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.number_of_hops", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.probe_mode", "SYN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.probe_per_hop", "9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.transmission_rate", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_started"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "NETWORK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "www.oracle.com:80"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0.name"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmSyntheticsNetworkMonitorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
				"batch_interval_in_seconds", //ignore as it does not apply to this case
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsDNSSecMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsDNSSecMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsDNSSecMonitorRepresentation), "apmsynthetics", "monitor", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsDNSSecMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "DNSSEC_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.record_type", "A"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted1),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSSecMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "DNSSEC_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.record_type", "ANY"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded2),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted2),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, ApmSyntheticsApmSyntheticsDNSMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSSecMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsDNSSecMonitorResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "DNSSEC_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.record_type", "ANY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_started"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0.name"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmSyntheticsDNSSecMonitorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
				"batch_interval_in_seconds", //ignore as it does not apply to this case
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsDNSServerMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsDNSServerMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsDNSServerMonitorRepresentation), "apmsynthetics", "monitor", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsDNSServerMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "DNS_SERVER_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_query_recursive", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.name_server", "8.8.8.8"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.number_of_hops", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_mode", "SACK"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_per_hop", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.transmission_rate", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.record_type", "A"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted1),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSServerMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "DNS_SERVER_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_query_recursive", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.name_server", "8.8.4.4"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.number_of_hops", "11"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_mode", "SYN"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.probe_per_hop", "9"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.network_configuration.0.transmission_rate", "11"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.protocol", "UDP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.record_type", "ANY"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded2),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted2),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, ApmSyntheticsApmSyntheticsDNSMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSServerMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsDNSServerMonitorResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "DNS_SERVER_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_query_recursive", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.name_server", "8.8.4.4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.number_of_hops", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.probe_mode", "SYN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.probe_per_hop", "9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.network_configuration.0.transmission_rate", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.protocol", "UDP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.record_type", "ANY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_started"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0.name"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmSyntheticsDNSServerMonitorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
				"batch_interval_in_seconds", //ignore as it does not apply to this case
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsDNSTraceMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsDNSTraceMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsDNSTraceMonitorRepresentation), "apmsynthetics", "monitor", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, ApmSyntheticsDNSTraceMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "DNS_TRACE_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.record_type", "A"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted1),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSTraceMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(resourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "DNS_TRACE_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.protocol", "UDP"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.record_type", "ANY"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_ended", TimeEnded2),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_schedule.0.time_started", TimeStarted2),
				resource.TestCheckResourceAttr(resourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(resourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0.name"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, ApmSyntheticsApmSyntheticsDNSMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, ApmSyntheticsDNSTraceMonitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, ApmSyntheticsMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsDNSTraceMonitorResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.max_allowed_failures_per_interval", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_configuration.0.min_allowed_runs_per_interval", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "DNS_TRACE_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.is_override_dns", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.dns_configuration.0.override_dns_ip", "12.1.21.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.protocol", "UDP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.record_type", "ANY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_schedule.0.time_started"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_now", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduling_policy", "ROUND_ROBIN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_run_once", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "DNS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "www.oracle.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0.name"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmSyntheticsDNSTraceMonitorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
				"batch_interval_in_seconds", //ignore as it does not apply to this case
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApmSyntheticsMonitorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApmSyntheticClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_synthetics_monitor" {
			noResourceFound = false
			request := oci_apm_synthetics.GetMonitorRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp := rs.Primary.ID
			request.MonitorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")

			_, err := client.GetMonitor(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ApmSyntheticsMonitor") {
		resource.AddTestSweepers("ApmSyntheticsMonitor", &resource.Sweeper{
			Name:         "ApmSyntheticsMonitor",
			Dependencies: acctest.DependencyGraph["monitor"],
			F:            sweepApmSyntheticsMonitorResource,
		})
	}
}

func sweepApmSyntheticsMonitorResource(compartment string) error {
	apmSyntheticClient := acctest.GetTestClients(&schema.ResourceData{}).ApmSyntheticClient()
	monitorIds, err := getApmSyntheticsMonitorIds(compartment)
	if err != nil {
		return err
	}
	for _, monitorId := range monitorIds {
		if ok := acctest.SweeperDefaultResourceId[monitorId]; !ok {
			deleteMonitorRequest := oci_apm_synthetics.DeleteMonitorRequest{}

			deleteMonitorRequest.MonitorId = &monitorId

			deleteMonitorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")
			_, error := apmSyntheticClient.DeleteMonitor(context.Background(), deleteMonitorRequest)
			if error != nil {
				fmt.Printf("Error deleting Monitor %s %s, It is possible that the resource is already deleted. Please verify manually \n", monitorId, error)
				continue
			}
		}
	}
	return nil
}

func getApmSyntheticsMonitorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MonitorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apmSyntheticClient := acctest.GetTestClients(&schema.ResourceData{}).ApmSyntheticClient()

	listMonitorsRequest := oci_apm_synthetics.ListMonitorsRequest{}

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for Monitor resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listMonitorsRequest.ApmDomainId = &apmDomainId

		listMonitorsResponse, err := apmSyntheticClient.ListMonitors(context.Background(), listMonitorsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Monitor list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, monitor := range listMonitorsResponse.Items {
			id := *monitor.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MonitorId", id)
		}

	}
	return resourceIds, nil
}
