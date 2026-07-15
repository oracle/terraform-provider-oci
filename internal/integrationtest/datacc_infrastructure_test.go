// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const (
	infrastructureNumDnsServers                    = 2
	infrastructureNumContacts                      = 1
	infrastructureNumNtpServers                    = 1
	infrastructureNumFreeformTags                  = 1
	infrastructureNumMaintenanceWindowDaysOfWeek   = 2
	infrastructureNumMaintenanceWindowHoursOfDay   = 2
	infrastructureNumMaintenanceWindowMonths       = 2
	infrastructureNumMaintenanceWindowWeeksOfMonth = 2
	infrastructureMaintenanceWindowPreference      = "CUSTOM_PREFERENCE"
	infrastructureMaintenanceWindowEnabled         = "true"
	infrastructureMonthlyPatchingEnabled           = "true"
)

var (
	DataccInfrastructureRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Required, acctest.Create, DataccInfrastructureRepresentation)

	DataccInfrastructureResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Optional, acctest.Update, DataccInfrastructureRepresentation)

	DataccInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacc_infrastructure.test_infrastructure.id}`},
	}

	DataccInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `${var.infrastructure_display_name}`, Update: `${var.infrastructure_display_name_for_update}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: []string{string(oci_datacc.InfrastructureLifecycleStateRequiresValidation)}, Update: []string{string(oci_datacc.InfrastructureLifecycleStateRequiresValidation)}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DataccInfrastructureDataSourceFilterRepresentation}}
	DataccInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datacc_infrastructure.test_infrastructure.id}`}},
	}

	DataccInfrastructureRepresentation = map[string]interface{}{
		"cloud_control_plane_server1": acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_cloud_control_plane_server1}`},
		"cloud_control_plane_server2": acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_cloud_control_plane_server2}`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_display_name}`, Update: `${var.infrastructure_display_name_for_update}`},
		"dns_servers":                 acctest.Representation{RepType: acctest.Required, Create: []string{`${var.infrastructure_dns_server_0}`, `${var.infrastructure_dns_server_1}`}},
		"gateway":                     acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_gateway}`},
		"netmask":                     acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_netmask}`},
		"ntp_servers":                 acctest.Representation{RepType: acctest.Required, Create: []string{`${var.infrastructure_ntp_server_0}`}},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_shape}`},
		"system_model":                acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_system_model}`},
		"admin_networkcidr":           acctest.Representation{RepType: acctest.Optional, Create: `${var.infrastructure_admin_networkcidr}`},
		"contacts":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataccInfrastructureContactsRepresentation},
		"corporate_proxy":             acctest.Representation{RepType: acctest.Optional, Create: `${var.infrastructure_corporate_proxy}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.infrastructure_description}`, Update: `${var.infrastructure_description_for_update}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataccInfrastructureMaintenanceWindowRepresentation},
	}
	DataccInfrastructureContactsRepresentation = map[string]interface{}{
		"email":      acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_contact_0_email}`, Update: `${var.infrastructure_contact_0_email_for_update}`},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_contact_0_name}`, Update: `${var.infrastructure_contact_0_name_for_update}`},
		"is_primary": acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_contact_0_is_primary}`},

		"phone_number":             acctest.Representation{RepType: acctest.Optional, Create: `${var.infrastructure_contact_0_phone_number}`, Update: `${var.infrastructure_contact_0_phone_number_for_update}`},
		"is_contact_mos_validated": acctest.Representation{RepType: acctest.Optional, Create: `${var.infrastructure_contact_0_is_contact_mos_validated}`},
	}
	DataccInfrastructureMaintenanceWindowRepresentation = map[string]interface{}{
		"custom_action_timeout_in_mins": acctest.Representation{
			RepType: acctest.Optional,
			Create:  `${var.infrastructure_maintenance_window_custom_action_timeout_in_mins}`,
			Update:  `${var.infrastructure_maintenance_window_custom_action_timeout_in_mins_for_update}`,
		},
		"days_of_week": acctest.Representation{
			RepType: acctest.Optional,
			Create: []string{
				`${var.infrastructure_maintenance_window_days_of_week_0}`,
				`${var.infrastructure_maintenance_window_days_of_week_1}`,
			},
			Update: []string{
				`${var.infrastructure_maintenance_window_days_of_week_0_for_update}`,
				`${var.infrastructure_maintenance_window_days_of_week_1_for_update}`,
			},
		},
		"hours_of_day": acctest.Representation{
			RepType: acctest.Optional,
			Create: []string{
				`${var.infrastructure_maintenance_window_hours_of_day_0}`,
				`${var.infrastructure_maintenance_window_hours_of_day_1}`,
			},
			Update: []string{
				`${var.infrastructure_maintenance_window_hours_of_day_0_for_update}`,
				`${var.infrastructure_maintenance_window_hours_of_day_1_for_update}`,
			},
		},
		"is_custom_action_timeout_enabled": acctest.Representation{RepType: acctest.Optional, Create: infrastructureMaintenanceWindowEnabled},
		"is_monthly_patching_enabled":      acctest.Representation{RepType: acctest.Optional, Create: infrastructureMonthlyPatchingEnabled},
		"lead_time_in_weeks": acctest.Representation{
			RepType: acctest.Optional,
			Create:  `${var.infrastructure_maintenance_window_lead_time_in_weeks}`,
			Update:  `${var.infrastructure_maintenance_window_lead_time_in_weeks_for_update}`,
		},
		"months": acctest.Representation{
			RepType: acctest.Optional,
			Create: []string{
				`${var.infrastructure_maintenance_window_months_0}`,
				`${var.infrastructure_maintenance_window_months_1}`,
			},
			Update: []string{
				`${var.infrastructure_maintenance_window_months_0_for_update}`,
				`${var.infrastructure_maintenance_window_months_1_for_update}`,
			},
		},
		"patching_mode": acctest.Representation{
			RepType: acctest.Optional,
			Create:  `${var.infrastructure_maintenance_window_patching_mode}`,
		},
		"preference": acctest.Representation{RepType: acctest.Optional, Create: infrastructureMaintenanceWindowPreference},
		"weeks_of_month": acctest.Representation{
			RepType: acctest.Optional,
			Create: []string{
				`${var.infrastructure_maintenance_window_weeks_of_month_0}`,
				`${var.infrastructure_maintenance_window_weeks_of_month_1}`,
			},
			Update: []string{
				`${var.infrastructure_maintenance_window_weeks_of_month_0_for_update}`,
				`${var.infrastructure_maintenance_window_weeks_of_month_1_for_update}`,
			},
		},
	}
)

// issue-routing-tag: datacc/default
func TestDataccInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataccInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// override terraform-federation-test profile with our own user profile
	if overrideProfile := os.Getenv("datacc_custom_config_file_profile_override"); overrideProfile != "" {
		t.Setenv(globalvar.TfEnvPrefix+globalvar.ConfigFileProfileAttrName, overrideProfile)
		t.Setenv(globalvar.TfEnvPrefix+globalvar.AuthAttrName, "")
		t.Setenv(globalvar.AuthAttrName, globalvar.AuthSecurityToken)
	}

	const testResourceType = "infrastructure"
	tfVariableStr := GenerateTFVariableStrings(testResourceType)
	getTFVar := func(variableName string) string {
		return os.Getenv(globalvar.TfEnvPrefix + testResourceType + "_" + variableName)
	}

	compartmentId := getTFVar("compartment_id")
	compartmentIdU := getTFVar("compartment_id_for_update")
	displayName := getTFVar("display_name")
	displayNameU := getTFVar("display_name_for_update")
	description := getTFVar("description")
	descriptionU := getTFVar("description_for_update")
	systemModel := getTFVar("system_model")
	shape := getTFVar("shape")
	cloudControlPlaneServer1 := getTFVar("cloud_control_plane_server1")
	cloudControlPlaneServer2 := getTFVar("cloud_control_plane_server2")
	netmask := getTFVar("netmask")
	gateway := getTFVar("gateway")
	adminNetworkcidr := getTFVar("admin_networkcidr")
	corporateProxy := getTFVar("corporate_proxy")
	contact0Name := getTFVar("contact_0_name")
	contact0NameU := getTFVar("contact_0_name_for_update")
	contact0PhoneNumber := getTFVar("contact_0_phone_number")
	contact0PhoneNumberU := getTFVar("contact_0_phone_number_for_update")
	contact0Email := getTFVar("contact_0_email")
	contact0EmailU := getTFVar("contact_0_email_for_update")
	contact0IsPrimary := getTFVar("contact_0_is_primary")
	contact0IsContactMosValidated := getTFVar("contact_0_is_contact_mos_validated")
	maintenanceWindowCustomActionTimeoutInMins := getTFVar("maintenance_window_custom_action_timeout_in_mins")
	maintenanceWindowCustomActionTimeoutInMinsU := getTFVar("maintenance_window_custom_action_timeout_in_mins_for_update")
	maintenanceWindowDaysOfWeek0 := getTFVar("maintenance_window_days_of_week_0")
	maintenanceWindowDaysOfWeek0U := getTFVar("maintenance_window_days_of_week_0_for_update")
	maintenanceWindowDaysOfWeek1 := getTFVar("maintenance_window_days_of_week_1")
	maintenanceWindowDaysOfWeek1U := getTFVar("maintenance_window_days_of_week_1_for_update")
	maintenanceWindowHoursOfDay0 := getTFVar("maintenance_window_hours_of_day_0")
	maintenanceWindowHoursOfDay0U := getTFVar("maintenance_window_hours_of_day_0_for_update")
	maintenanceWindowHoursOfDay1 := getTFVar("maintenance_window_hours_of_day_1")
	maintenanceWindowHoursOfDay1U := getTFVar("maintenance_window_hours_of_day_1_for_update")
	maintenanceWindowLeadTimeInWeeks := getTFVar("maintenance_window_lead_time_in_weeks")
	maintenanceWindowLeadTimeInWeeksU := getTFVar("maintenance_window_lead_time_in_weeks_for_update")
	maintenanceWindowMonths0 := getTFVar("maintenance_window_months_0")
	maintenanceWindowMonths0U := getTFVar("maintenance_window_months_0_for_update")
	maintenanceWindowMonths1 := getTFVar("maintenance_window_months_1")
	maintenanceWindowMonths1U := getTFVar("maintenance_window_months_1_for_update")
	maintenanceWindowPatchingMode := getTFVar("maintenance_window_patching_mode")
	maintenanceWindowWeeksOfMonth0 := getTFVar("maintenance_window_weeks_of_month_0")
	maintenanceWindowWeeksOfMonth0U := getTFVar("maintenance_window_weeks_of_month_0_for_update")
	maintenanceWindowWeeksOfMonth1 := getTFVar("maintenance_window_weeks_of_month_1")
	maintenanceWindowWeeksOfMonth1U := getTFVar("maintenance_window_weeks_of_month_1_for_update")

	resourceName := "oci_datacc_infrastructure.test_infrastructure"
	datasourceName := "data.oci_datacc_infrastructures.test_infrastructures"
	singularDatasourceName := "data.oci_datacc_infrastructure.test_infrastructure"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+tfVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Optional, acctest.Create, DataccInfrastructureRepresentation), "datacc", "infrastructure", t)

	acctest.ResourceTest(t, testAccCheckDataccInfrastructureDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + tfVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Required, acctest.Create, DataccInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", cloudControlPlaneServer1),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", cloudControlPlaneServer2),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "dns_servers.#", strconv.Itoa(infrastructureNumDnsServers)),
				resource.TestCheckResourceAttr(resourceName, "gateway", gateway),
				resource.TestCheckResourceAttr(resourceName, "netmask", netmask),
				resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", strconv.Itoa(infrastructureNumNtpServers)),
				resource.TestCheckResourceAttr(resourceName, "shape", shape),
				resource.TestCheckResourceAttr(resourceName, "system_model", systemModel),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + tfVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + tfVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Optional, acctest.Create, DataccInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_networkcidr", adminNetworkcidr),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", cloudControlPlaneServer1),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", cloudControlPlaneServer2),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contacts.#", strconv.Itoa(infrastructureNumContacts)),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.email", contact0Email),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_contact_mos_validated", contact0IsContactMosValidated),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_primary", contact0IsPrimary),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.name", contact0Name),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.phone_number", contact0PhoneNumber),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", corporateProxy),
				resource.TestCheckResourceAttr(resourceName, "description", description),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "dns_servers.#", strconv.Itoa(infrastructureNumDnsServers)),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(infrastructureNumFreeformTags)),
				resource.TestCheckResourceAttr(resourceName, "gateway", gateway),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.custom_action_timeout_in_mins", maintenanceWindowCustomActionTimeoutInMins),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", strconv.Itoa(infrastructureNumMaintenanceWindowDaysOfWeek)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0", maintenanceWindowDaysOfWeek0),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.1", maintenanceWindowDaysOfWeek1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", strconv.Itoa(infrastructureNumMaintenanceWindowHoursOfDay)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.0", maintenanceWindowHoursOfDay0),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.1", maintenanceWindowHoursOfDay1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_custom_action_timeout_enabled", infrastructureMaintenanceWindowEnabled),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_monthly_patching_enabled", infrastructureMonthlyPatchingEnabled),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", maintenanceWindowLeadTimeInWeeks),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", strconv.Itoa(infrastructureNumMaintenanceWindowMonths)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0", maintenanceWindowMonths0),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.1", maintenanceWindowMonths1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", maintenanceWindowPatchingMode),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", infrastructureMaintenanceWindowPreference),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", strconv.Itoa(infrastructureNumMaintenanceWindowWeeksOfMonth)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.0", maintenanceWindowWeeksOfMonth0),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.1", maintenanceWindowWeeksOfMonth1),
				resource.TestCheckResourceAttr(resourceName, "netmask", netmask),
				resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", strconv.Itoa(infrastructureNumNtpServers)),
				resource.TestCheckResourceAttr(resourceName, "shape", shape),
				resource.TestCheckResourceAttr(resourceName, "system_model", systemModel),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + tfVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataccInfrastructureRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_networkcidr", adminNetworkcidr),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", cloudControlPlaneServer1),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", cloudControlPlaneServer2),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "contacts.#", strconv.Itoa(infrastructureNumContacts)),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.email", contact0Email),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_contact_mos_validated", contact0IsContactMosValidated),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_primary", contact0IsPrimary),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.name", contact0Name),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.phone_number", contact0PhoneNumber),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", corporateProxy),
				resource.TestCheckResourceAttr(resourceName, "description", description),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "dns_servers.#", strconv.Itoa(infrastructureNumDnsServers)),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(infrastructureNumFreeformTags)),
				resource.TestCheckResourceAttr(resourceName, "gateway", gateway),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.custom_action_timeout_in_mins", maintenanceWindowCustomActionTimeoutInMins),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", strconv.Itoa(infrastructureNumMaintenanceWindowDaysOfWeek)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0", maintenanceWindowDaysOfWeek0),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.1", maintenanceWindowDaysOfWeek1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", strconv.Itoa(infrastructureNumMaintenanceWindowHoursOfDay)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.0", maintenanceWindowHoursOfDay0),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.1", maintenanceWindowHoursOfDay1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_custom_action_timeout_enabled", infrastructureMaintenanceWindowEnabled),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_monthly_patching_enabled", infrastructureMonthlyPatchingEnabled),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", maintenanceWindowLeadTimeInWeeks),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", strconv.Itoa(infrastructureNumMaintenanceWindowMonths)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0", maintenanceWindowMonths0),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.1", maintenanceWindowMonths1),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", maintenanceWindowPatchingMode),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", infrastructureMaintenanceWindowPreference),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", strconv.Itoa(infrastructureNumMaintenanceWindowWeeksOfMonth)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.0", maintenanceWindowWeeksOfMonth0),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.1", maintenanceWindowWeeksOfMonth1),
				resource.TestCheckResourceAttr(resourceName, "netmask", netmask),
				resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", strconv.Itoa(infrastructureNumNtpServers)),
				resource.TestCheckResourceAttr(resourceName, "shape", shape),
				resource.TestCheckResourceAttr(resourceName, "system_model", systemModel),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + tfVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Optional, acctest.Update, DataccInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_networkcidr", adminNetworkcidr),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", cloudControlPlaneServer1),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", cloudControlPlaneServer2),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contacts.#", strconv.Itoa(infrastructureNumContacts)),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.email", contact0EmailU),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_contact_mos_validated", contact0IsContactMosValidated),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_primary", contact0IsPrimary),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.name", contact0NameU),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.phone_number", contact0PhoneNumberU),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", corporateProxy),
				resource.TestCheckResourceAttr(resourceName, "description", descriptionU),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(resourceName, "dns_servers.#", strconv.Itoa(infrastructureNumDnsServers)),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(infrastructureNumFreeformTags)),
				resource.TestCheckResourceAttr(resourceName, "gateway", gateway),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.custom_action_timeout_in_mins", maintenanceWindowCustomActionTimeoutInMinsU),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", strconv.Itoa(infrastructureNumMaintenanceWindowDaysOfWeek)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0", maintenanceWindowDaysOfWeek0U),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.1", maintenanceWindowDaysOfWeek1U),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", strconv.Itoa(infrastructureNumMaintenanceWindowHoursOfDay)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.0", maintenanceWindowHoursOfDay0U),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.1", maintenanceWindowHoursOfDay1U),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_custom_action_timeout_enabled", infrastructureMaintenanceWindowEnabled),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_monthly_patching_enabled", infrastructureMonthlyPatchingEnabled),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", maintenanceWindowLeadTimeInWeeksU),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", strconv.Itoa(infrastructureNumMaintenanceWindowMonths)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0", maintenanceWindowMonths0U),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.1", maintenanceWindowMonths1U),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", maintenanceWindowPatchingMode),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", infrastructureMaintenanceWindowPreference),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", strconv.Itoa(infrastructureNumMaintenanceWindowWeeksOfMonth)),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.0", maintenanceWindowWeeksOfMonth0U),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.1", maintenanceWindowWeeksOfMonth1U),
				resource.TestCheckResourceAttr(resourceName, "netmask", netmask),
				resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", strconv.Itoa(infrastructureNumNtpServers)),
				resource.TestCheckResourceAttr(resourceName, "shape", shape),
				resource.TestCheckResourceAttr(resourceName, "system_model", systemModel),

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
			Config: config + tfVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_infrastructures", "test_infrastructures", acctest.Optional, acctest.Update, DataccInfrastructureDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Optional, acctest.Update, DataccInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.custom_action_timeout_in_mins", maintenanceWindowCustomActionTimeoutInMinsU),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.days_of_week.#", strconv.Itoa(infrastructureNumMaintenanceWindowDaysOfWeek)),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.days_of_week.0", maintenanceWindowDaysOfWeek0U),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.days_of_week.1", maintenanceWindowDaysOfWeek1U),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.hours_of_day.#", strconv.Itoa(infrastructureNumMaintenanceWindowHoursOfDay)),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.hours_of_day.0", maintenanceWindowHoursOfDay0U),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.hours_of_day.1", maintenanceWindowHoursOfDay1U),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.is_custom_action_timeout_enabled", infrastructureMaintenanceWindowEnabled),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.is_monthly_patching_enabled", infrastructureMonthlyPatchingEnabled),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.lead_time_in_weeks", maintenanceWindowLeadTimeInWeeksU),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.months.#", strconv.Itoa(infrastructureNumMaintenanceWindowMonths)),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.months.0", maintenanceWindowMonths0U),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.months.1", maintenanceWindowMonths1U),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.patching_mode", maintenanceWindowPatchingMode),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.preference", infrastructureMaintenanceWindowPreference),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.weeks_of_month.#", strconv.Itoa(infrastructureNumMaintenanceWindowWeeksOfMonth)),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.weeks_of_month.0", maintenanceWindowWeeksOfMonth0U),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_collection.0.items.0.maintenance_window.0.weeks_of_month.1", maintenanceWindowWeeksOfMonth1U),
			),
		},
		// verify singular datasource
		{
			Config: config + tfVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Required, acctest.Create, DataccInfrastructureSingularDataSourceRepresentation) +
				DataccInfrastructureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "infrastructure_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "admin_networkcidr", adminNetworkcidr),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_control_plane_server1", cloudControlPlaneServer1),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_control_plane_server2", cloudControlPlaneServer2),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_capacity.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.#", strconv.Itoa(infrastructureNumContacts)),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.email", contact0EmailU),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.is_contact_mos_validated", contact0IsContactMosValidated),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.is_primary", contact0IsPrimary),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.name", contact0NameU),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.phone_number", contact0PhoneNumberU),
				resource.TestCheckResourceAttr(singularDatasourceName, "corporate_proxy", corporateProxy),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", descriptionU),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_servers.#", strconv.Itoa(infrastructureNumDnsServers)),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", strconv.Itoa(infrastructureNumFreeformTags)),
				resource.TestCheckResourceAttr(singularDatasourceName, "gateway", gateway),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.custom_action_timeout_in_mins", maintenanceWindowCustomActionTimeoutInMinsU),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", strconv.Itoa(infrastructureNumMaintenanceWindowDaysOfWeek)),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0", maintenanceWindowDaysOfWeek0U),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.1", maintenanceWindowDaysOfWeek1U),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", strconv.Itoa(infrastructureNumMaintenanceWindowHoursOfDay)),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.0", maintenanceWindowHoursOfDay0U),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.1", maintenanceWindowHoursOfDay1U),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.is_custom_action_timeout_enabled", infrastructureMaintenanceWindowEnabled),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.is_monthly_patching_enabled", infrastructureMonthlyPatchingEnabled),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.lead_time_in_weeks", maintenanceWindowLeadTimeInWeeksU),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", strconv.Itoa(infrastructureNumMaintenanceWindowMonths)),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0", maintenanceWindowMonths0U),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.1", maintenanceWindowMonths1U),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.patching_mode", maintenanceWindowPatchingMode),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", infrastructureMaintenanceWindowPreference),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", strconv.Itoa(infrastructureNumMaintenanceWindowWeeksOfMonth)),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.0", maintenanceWindowWeeksOfMonth0U),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.1", maintenanceWindowWeeksOfMonth1U),
				resource.TestCheckResourceAttr(singularDatasourceName, "netmask", netmask),
				resource.TestCheckResourceAttr(singularDatasourceName, "ntp_servers.#", strconv.Itoa(infrastructureNumNtpServers)),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reco_disk_percentage"),
				resource.TestCheckResourceAttr(singularDatasourceName, "servers.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", shape),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_capacity.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_model", systemModel),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_storage_capacity.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_state_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + tfVariableStr + DataccInfrastructureRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataccInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BaseinfraClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacc_infrastructure" {
			noResourceFound = false
			request := oci_datacc.GetInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.InfrastructureId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacc")

			response, err := client.GetInfrastructure(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datacc.InfrastructureLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if !acctest.InSweeperExcludeList("DataccInfrastructure") {
		resource.AddTestSweepers("DataccInfrastructure", &resource.Sweeper{
			Name:         "DataccInfrastructure",
			Dependencies: acctest.DependencyGraph["infrastructure"],
			F:            sweepDataccInfrastructureResource,
		})
	}
}

func sweepDataccInfrastructureResource(compartment string) error {
	baseinfraClient := acctest.GetTestClients(&schema.ResourceData{}).BaseinfraClient()
	infrastructureIds, err := getDataccInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, infrastructureId := range infrastructureIds {
		if ok := acctest.SweeperDefaultResourceId[infrastructureId]; !ok {
			deleteInfrastructureRequest := oci_datacc.DeleteInfrastructureRequest{}

			deleteInfrastructureRequest.InfrastructureId = &infrastructureId

			deleteInfrastructureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacc")
			_, error := baseinfraClient.DeleteInfrastructure(context.Background(), deleteInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting Infrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", infrastructureId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &infrastructureId, DataccInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				DataccInfrastructureSweepResponseFetchOperation, "datacc", true)
		}
	}
	return nil
}

func getDataccInfrastructureIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	baseinfraClient := acctest.GetTestClients(&schema.ResourceData{}).BaseinfraClient()

	listInfrastructuresRequest := oci_datacc.ListInfrastructuresRequest{}
	listInfrastructuresRequest.CompartmentId = &compartmentId
	listInfrastructuresRequest.LifecycleState = []oci_datacc.InfrastructureLifecycleStateEnum{oci_datacc.InfrastructureLifecycleStateRequiresValidation}
	listInfrastructuresResponse, err := baseinfraClient.ListInfrastructures(context.Background(), listInfrastructuresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Infrastructure list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, infrastructure := range listInfrastructuresResponse.Items {
		id := *infrastructure.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InfrastructureId", id)
	}
	return resourceIds, nil
}

func DataccInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if infrastructureResponse, ok := response.Response.(oci_datacc.GetInfrastructureResponse); ok {
		return infrastructureResponse.LifecycleState != oci_datacc.InfrastructureLifecycleStateDeleted
	}
	return false
}

func DataccInfrastructureSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BaseinfraClient().GetInfrastructure(context.Background(), oci_datacc.GetInfrastructureRequest{
		InfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
