// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_desktops "github.com/oracle/oci-go-sdk/v65/desktops"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DesktopsDesktopPoolRequiredOnlyResource = DesktopsDesktopPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Required, acctest.Create, DesktopsDesktopPoolRepresentation)

	DesktopsDesktopPoolResourceConfig = DesktopsDesktopPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Update, DesktopsDesktopPoolRepresentation)

	DesktopsDesktopPoolSingularDataSourceRepresentation = map[string]interface{}{
		"desktop_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_desktops_desktop_pool.test_desktop_pool.id}`},
	}

	DesktopsDesktopPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `testPool1`, Update: `testPool2`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_desktops_desktop_pool.test_desktop_pool.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolDataSourceFilterRepresentation}}
	DesktopsDesktopPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_desktops_desktop_pool.test_desktop_pool.id}`}},
	}

	DesktopsDesktopPoolRepresentation = map[string]interface{}{
		"are_privileged_users":      acctest.Representation{RepType: acctest.Required, Create: `false`},
		"availability_domain":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"availability_policy":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolAvailabilityPolicyRepresentation},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"contact_details":           acctest.Representation{RepType: acctest.Required, Create: `contactDetails`, Update: `contactDetails2`},
		"device_policy":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolDevicePolicyRepresentation},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `testPool1`, Update: `testPool2`},
		"image":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolImageRepresentation},
		"is_storage_enabled":        acctest.Representation{RepType: acctest.Required, Create: `true`},
		"maximum_size":              acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"network_configuration":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolNetworkConfigurationRepresentation},
		"shape_name":                acctest.Representation{RepType: acctest.Required, Create: `${var.test_shape_name}`},
		"standby_size":              acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"storage_backup_policy_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.test_storage_backup_policy_id}`},
		"storage_size_in_gbs":       acctest.Representation{RepType: acctest.Required, Create: `50`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.test_nsg_id}`}},
		"shape_config":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolShapeConfigRepresentation},
		"use_dedicated_vm_host":     acctest.Representation{RepType: acctest.Optional, Create: `${var.test_use_dedicated_vm_host}`},
		"private_access_details":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolPrivateAccessDetailsRepresentation},
		"session_lifecycle_actions": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolSessionLifecycleActionsRepresentation},
	}

	DesktopsDesktopPoolAllSessionLifecycleActionsNoAvailPolicySchedulesRepresentation = map[string]interface{}{
		"are_privileged_users":      acctest.Representation{RepType: acctest.Required, Create: `false`},
		"availability_domain":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"availability_policy":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolAvailabilityPolicyNoStartStopSchedulesRepresentation},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"contact_details":           acctest.Representation{RepType: acctest.Required, Create: `contactDetails`, Update: `contactDetails2`},
		"device_policy":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolDevicePolicyRepresentation},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `testPool1`, Update: `testPool2`},
		"image":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolImageRepresentation},
		"is_storage_enabled":        acctest.Representation{RepType: acctest.Required, Create: `true`},
		"maximum_size":              acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"network_configuration":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DesktopsDesktopPoolNetworkConfigurationRepresentation},
		"shape_name":                acctest.Representation{RepType: acctest.Required, Create: `${var.test_shape_name}`},
		"standby_size":              acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"storage_backup_policy_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.test_storage_backup_policy_id}`},
		"storage_size_in_gbs":       acctest.Representation{RepType: acctest.Required, Create: `50`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.test_nsg_id}`}},
		"session_lifecycle_actions": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolSessionLifecycleActionsAllRepresentation},
	}

	DesktopsDesktopPoolAvailabilityPolicyRepresentation = map[string]interface{}{
		"start_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolAvailabilityPolicyStartScheduleRepresentation},
		"stop_schedule":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolAvailabilityPolicyStopScheduleRepresentation},
	}

	// This is the only way for not to not pass availability policy schedules and pass server side validation
	DesktopsDesktopPoolAvailabilityPolicyNoStartStopSchedulesRepresentation = map[string]interface{}{}

	DesktopsDesktopPoolDevicePolicyRepresentation = map[string]interface{}{
		"audio_mode":          acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `TODESKTOP`},
		"cdm_mode":            acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `READONLY`},
		"clipboard_mode":      acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `TODESKTOP`},
		"is_display_enabled":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_keyboard_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_pointer_enabled":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_printing_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}
	DesktopsDesktopPoolImageRepresentation = map[string]interface{}{
		"image_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.test_image_id}`},
		"image_name": acctest.Representation{RepType: acctest.Required, Create: `${var.test_image_name}`},
	}
	DesktopsDesktopPoolNetworkConfigurationRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_subnet_id}`},
		"vcn_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.test_vcn_id}`},
	}

	DesktopsDesktopPoolShapeConfigRepresentation = map[string]interface{}{
		"baseline_ocpu_utilization": acctest.Representation{RepType: acctest.Optional, Create: `${var.test_shape_config_baseline_ocpu_utilization}`},
		"memory_in_gbs":             acctest.Representation{RepType: acctest.Optional, Create: `${var.test_shape_config_memory_in_gbs}`},
		"ocpus":                     acctest.Representation{RepType: acctest.Optional, Create: `${var.test_shape_config_ocpus}`},
	}
	DesktopsDesktopPoolPrivateAccessDetailsRepresentation = map[string]interface{}{
		"subnet_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.test_private_access_subnet_id}`},
		"nsg_ids":    acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.test_private_access_nsg_id}`}},
		"private_ip": acctest.Representation{RepType: acctest.Optional, Create: `${var.test_private_access_private_ip}`},
	}

	DesktopsDesktopPoolSessionLifecycleActionsRepresentation = map[string]interface{}{
		"inactivity": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolSessionLifecycleActionsInactivityRepresentation},
	}

	DesktopsDesktopPoolSessionLifecycleActionsAllRepresentation = map[string]interface{}{
		"inactivity": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolSessionLifecycleActionsInactivityRepresentation},
		"disconnect": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DesktopsDesktopPoolSessionLifecycleActionsDisconnectRepresentation},
	}

	DesktopsDesktopPoolAvailabilityPolicyStartScheduleRepresentation = map[string]interface{}{
		"cron_expression": acctest.Representation{RepType: acctest.Required, Create: `${var.test_start_schedule_cron_expr_create}`, Update: `${var.test_start_schedule_cron_expr_update}`},
		"timezone":        acctest.Representation{RepType: acctest.Required, Create: `${var.test_start_schedule_timezone_create}`, Update: `${var.test_start_schedule_timezone_update}`},
	}
	DesktopsDesktopPoolAvailabilityPolicyStopScheduleRepresentation = map[string]interface{}{
		"cron_expression": acctest.Representation{RepType: acctest.Required, Create: `${var.test_stop_schedule_cron_expr_create}`, Update: `${var.test_stop_schedule_cron_expr_update}`},
		"timezone":        acctest.Representation{RepType: acctest.Required, Create: `${var.test_stop_schedule_timezone_create}`, Update: `${var.test_stop_schedule_timezone_update}`},
	}

	DesktopsDesktopPoolSessionLifecycleActionsInactivityRepresentation = map[string]interface{}{
		"action":                  acctest.Representation{RepType: acctest.Required, Create: `${var.test_slm_inactivity_action_create}`},
		"grace_period_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `${var.test_slm_inactivity_grace_period_create}`},
	}

	DesktopsDesktopPoolSessionLifecycleActionsDisconnectRepresentation = map[string]interface{}{
		"action":                  acctest.Representation{RepType: acctest.Required, Create: `${var.test_slm_disconnect_action_create}`},
		"grace_period_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `${var.test_slm_disconnect_grace_period_create}`},
	}

	test_vcn_id      = utils.GetEnvSettingWithBlankDefault("test_vcn_id")
	vcnIdVariableStr = fmt.Sprintf("variable \"test_vcn_id\" { default = \"%s\" }\n", test_vcn_id)

	test_subnet_id      = utils.GetEnvSettingWithBlankDefault("test_subnet_id")
	subnetIdVariableStr = fmt.Sprintf("variable \"test_subnet_id\" { default = \"%s\" }\n", test_subnet_id)

	test_shape_name      = utils.GetEnvSettingWithBlankDefault("test_shape_name")
	shapeNameVariableStr = fmt.Sprintf("variable \"test_shape_name\" { default = \"%s\" }\n", test_shape_name)

	test_image_id      = utils.GetEnvSettingWithBlankDefault("test_image_id")
	imageIdVariableStr = fmt.Sprintf("variable \"test_image_id\" { default = \"%s\" }\n", test_image_id)

	test_image_name      = utils.GetEnvSettingWithBlankDefault("test_image_name")
	imageNameVariableStr = fmt.Sprintf("variable \"test_image_name\" { default = \"%s\" }\n", test_image_name)

	test_storage_backup_policy_id    = utils.GetEnvSettingWithBlankDefault("test_storage_backup_policy_id")
	storageBackupPolicyIdVariableStr = fmt.Sprintf("variable \"test_storage_backup_policy_id\" { default = \"%s\" }\n", test_storage_backup_policy_id)

	test_start_schedule_cron_expr_create   = utils.GetEnvSettingWithBlankDefault("test_start_schedule_cron_expr_create")
	startScheduleCronExprCreateVariableStr = fmt.Sprintf("variable \"test_start_schedule_cron_expr_create\" { default = \"%s\" }\n", test_start_schedule_cron_expr_create)

	test_start_schedule_cron_expr_update   = utils.GetEnvSettingWithBlankDefault("test_start_schedule_cron_expr_update")
	startScheduleCronExprUpdateVariableStr = fmt.Sprintf("variable \"test_start_schedule_cron_expr_update\" { default = \"%s\" }\n", test_start_schedule_cron_expr_update)

	test_start_schedule_timezone_create    = utils.GetEnvSettingWithBlankDefault("test_start_schedule_timezone_create")
	startScheduleTimezoneCreateVariableStr = fmt.Sprintf("variable \"test_start_schedule_timezone_create\" { default = \"%s\" }\n", test_start_schedule_timezone_create)

	test_start_schedule_timezone_update    = utils.GetEnvSettingWithBlankDefault("test_start_schedule_timezone_update")
	startScheduleTimezoneUpdateVariableStr = fmt.Sprintf("variable \"test_start_schedule_timezone_update\" { default = \"%s\" }\n", test_start_schedule_timezone_update)

	test_stop_schedule_cron_expr_create   = utils.GetEnvSettingWithBlankDefault("test_stop_schedule_cron_expr_create")
	stopScheduleCronExprCreateVariableStr = fmt.Sprintf("variable \"test_stop_schedule_cron_expr_create\" { default = \"%s\" }\n", test_stop_schedule_cron_expr_create)

	test_stop_schedule_cron_expr_update   = utils.GetEnvSettingWithBlankDefault("test_stop_schedule_cron_expr_update")
	stopScheduleCronExprUpdateVariableStr = fmt.Sprintf("variable \"test_stop_schedule_cron_expr_update\" { default = \"%s\" }\n", test_stop_schedule_cron_expr_update)

	test_stop_schedule_timezone_create    = utils.GetEnvSettingWithBlankDefault("test_stop_schedule_timezone_create")
	stopScheduleTimezoneCreateVariableStr = fmt.Sprintf("variable \"test_stop_schedule_timezone_create\" { default = \"%s\" }\n", test_stop_schedule_timezone_create)

	test_stop_schedule_timezone_update    = utils.GetEnvSettingWithBlankDefault("test_stop_schedule_timezone_update")
	stopScheduleTimezoneUpdateVariableStr = fmt.Sprintf("variable \"test_stop_schedule_timezone_update\" { default = \"%s\" }\n", test_stop_schedule_timezone_update)

	test_nsg_id      = utils.GetEnvSettingWithBlankDefault("test_nsg_id")
	nsgIdVariableStr = fmt.Sprintf("variable \"test_nsg_id\" { default = \"%s\" }\n", test_nsg_id)

	test_use_dedicated_vm_host    = utils.GetEnvSettingWithBlankDefault("test_use_dedicated_vm_host")
	useDedicatedVmHostVariableStr = fmt.Sprintf("variable \"test_use_dedicated_vm_host\" { default = \"%s\" }\n", test_use_dedicated_vm_host)

	test_shape_config_baseline_ocpu_utilization   = utils.GetEnvSettingWithBlankDefault("test_shape_config_baseline_ocpu_utilization")
	shapeConfigBaselineOcpuUtilizationVariableStr = fmt.Sprintf("variable \"test_shape_config_baseline_ocpu_utilization\" { default = \"%s\" }\n", test_shape_config_baseline_ocpu_utilization)

	test_shape_config_memory_in_gbs   = utils.GetEnvSettingWithBlankDefault("test_shape_config_memory_in_gbs")
	shapeConfigMemoryInGbsVariableStr = fmt.Sprintf("variable \"test_shape_config_memory_in_gbs\" { default = \"%s\" }\n", test_shape_config_memory_in_gbs)

	test_shape_config_ocpus     = utils.GetEnvSettingWithBlankDefault("test_shape_config_ocpus")
	shapeConfigOcpusVariableStr = fmt.Sprintf("variable \"test_shape_config_ocpus\" { default = \"%s\" }\n", test_shape_config_ocpus)

	test_private_access_subnet_id    = utils.GetEnvSettingWithBlankDefault("test_private_access_subnet_id")
	privateAccessSubnetIdVariableStr = fmt.Sprintf("variable \"test_private_access_subnet_id\" { default = \"%s\" }\n", test_private_access_subnet_id)

	test_private_access_nsg_id    = utils.GetEnvSettingWithBlankDefault("test_private_access_nsg_id")
	privateAccessNsgIdVariableStr = fmt.Sprintf("variable \"test_private_access_nsg_id\" { default = \"%s\" }\n", test_private_access_nsg_id)

	test_private_access_private_ip    = utils.GetEnvSettingWithBlankDefault("test_private_access_private_ip")
	privateAccessPrivateIpVariableStr = fmt.Sprintf("variable \"test_private_access_private_ip\" { default = \"%s\" }\n", test_private_access_private_ip)

	test_slm_disconnect_action_create    = utils.GetEnvSettingWithBlankDefault("test_slm_disconnect_action_create")
	slmDisconnectActionCreateVariableStr = fmt.Sprintf("variable \"test_slm_disconnect_action_create\" { default = \"%s\" }\n", test_slm_disconnect_action_create)

	test_slm_disconnect_action_update    = utils.GetEnvSettingWithBlankDefault("test_slm_disconnect_action_update")
	slmDisconnectActionUpdateVariableStr = fmt.Sprintf("variable \"test_slm_disconnect_action_update\" { default = \"%s\" }\n", test_slm_disconnect_action_update)

	test_slm_disconnect_grace_period_create   = utils.GetEnvSettingWithBlankDefault("test_slm_disconnect_grace_period_create")
	slmDisconnectGracePeriodCreateVariableStr = fmt.Sprintf("variable \"test_slm_disconnect_grace_period_create\" { default = \"%s\" }\n", test_slm_disconnect_grace_period_create)

	test_slm_inactivity_action_create    = utils.GetEnvSettingWithBlankDefault("test_slm_inactivity_action_create")
	slmInactivityActionCreateVariableStr = fmt.Sprintf("variable \"test_slm_inactivity_action_create\" { default = \"%s\" }\n", test_slm_inactivity_action_create)

	test_slm_inactivity_action_update    = utils.GetEnvSettingWithBlankDefault("test_slm_inactivity_action_update")
	slmInactivityActionUpdateVariableStr = fmt.Sprintf("variable \"test_slm_inactivity_action_update\" { default = \"%s\" }\n", test_slm_inactivity_action_update)

	test_slm_inactivity_grace_period_create   = utils.GetEnvSettingWithBlankDefault("test_slm_inactivity_grace_period_create")
	slmInactivityGracePeriodCreateVariableStr = fmt.Sprintf("variable \"test_slm_inactivity_grace_period_create\" { default = \"%s\" }\n", test_slm_inactivity_grace_period_create)

	DesktopsDesktopPoolResourceDependencies = vcnIdVariableStr +
		subnetIdVariableStr +
		shapeNameVariableStr +
		imageIdVariableStr +
		imageNameVariableStr +
		storageBackupPolicyIdVariableStr +
		startScheduleCronExprCreateVariableStr +
		startScheduleCronExprUpdateVariableStr +
		startScheduleTimezoneCreateVariableStr +
		startScheduleTimezoneUpdateVariableStr +
		stopScheduleCronExprCreateVariableStr +
		stopScheduleCronExprUpdateVariableStr +
		stopScheduleTimezoneCreateVariableStr +
		stopScheduleTimezoneUpdateVariableStr +
		nsgIdVariableStr +
		useDedicatedVmHostVariableStr +
		shapeConfigBaselineOcpuUtilizationVariableStr +
		shapeConfigMemoryInGbsVariableStr +
		shapeConfigOcpusVariableStr +
		privateAccessSubnetIdVariableStr +
		privateAccessNsgIdVariableStr +
		privateAccessPrivateIpVariableStr +
		slmDisconnectActionCreateVariableStr +
		slmDisconnectActionUpdateVariableStr +
		slmDisconnectGracePeriodCreateVariableStr +
		slmInactivityActionCreateVariableStr +
		slmInactivityActionUpdateVariableStr +
		slmInactivityGracePeriodCreateVariableStr +
		AvailabilityDomainConfig
)

// issue-routing-tag: desktops/default
func TestDesktopsDesktopPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDesktopsDesktopPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_desktops_desktop_pool.test_desktop_pool"
	datasourceName := "data.oci_desktops_desktop_pools.test_desktop_pools"
	singularDatasourceName := "data.oci_desktops_desktop_pool.test_desktop_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DesktopsDesktopPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Create, DesktopsDesktopPoolRepresentation), "desktops", "desktopPool", t)

	acctest.ResourceTest(t, testAccCheckDesktopsDesktopPoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DesktopsDesktopPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Required, acctest.Create, DesktopsDesktopPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_privileged_users", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contact_details", "contactDetails"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.audio_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.cdm_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.clipboard_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_display_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_keyboard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_pointer_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_printing_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testPool1"),
				resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_name"),
				resource.TestCheckResourceAttr(resourceName, "is_storage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maximum_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "standby_size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_gbs", "50"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DesktopsDesktopPoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DesktopsDesktopPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Create, DesktopsDesktopPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_privileged_users", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.0.start_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.start_schedule.0.cron_expression"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.start_schedule.0.timezone"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.0.stop_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.stop_schedule.0.cron_expression"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.stop_schedule.0.timezone"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contact_details", "contactDetails"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.audio_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.cdm_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.clipboard_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_display_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_keyboard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_pointer_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_printing_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testPool1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_name"),
				resource.TestCheckResourceAttr(resourceName, "is_storage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maximum_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.baseline_ocpu_utilization"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.ocpus"),
				resource.TestCheckResourceAttr(resourceName, "private_access_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "session_lifecycle_actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "session_lifecycle_actions.0.inactivity.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "session_lifecycle_actions.0.inactivity.0.action"),
				resource.TestCheckResourceAttrSet(resourceName, "session_lifecycle_actions.0.inactivity.0.grace_period_in_minutes"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "standby_size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "use_dedicated_vm_host"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DesktopsDesktopPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DesktopsDesktopPoolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_privileged_users", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.0.start_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.start_schedule.0.cron_expression"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.start_schedule.0.timezone"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.0.stop_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.stop_schedule.0.cron_expression"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.stop_schedule.0.timezone"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "contact_details", "contactDetails"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.audio_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.cdm_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.clipboard_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_display_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_keyboard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_pointer_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_printing_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testPool1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_name"),
				resource.TestCheckResourceAttr(resourceName, "is_storage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maximum_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.baseline_ocpu_utilization"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.ocpus"),
				resource.TestCheckResourceAttr(resourceName, "private_access_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "session_lifecycle_actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "session_lifecycle_actions.0.inactivity.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "session_lifecycle_actions.0.inactivity.0.action"),
				resource.TestCheckResourceAttrSet(resourceName, "session_lifecycle_actions.0.inactivity.0.grace_period_in_minutes"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "standby_size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "use_dedicated_vm_host"),

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
			Config: config + compartmentIdVariableStr + DesktopsDesktopPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Update, DesktopsDesktopPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_privileged_users", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.0.start_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.start_schedule.0.cron_expression"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.start_schedule.0.timezone"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.0.stop_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.stop_schedule.0.cron_expression"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_policy.0.stop_schedule.0.timezone"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contact_details", "contactDetails2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.audio_mode", "TODESKTOP"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.cdm_mode", "READONLY"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.clipboard_mode", "TODESKTOP"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_display_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_keyboard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_pointer_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_printing_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testPool2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_name"),
				resource.TestCheckResourceAttr(resourceName, "is_storage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maximum_size", "11"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.baseline_ocpu_utilization"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_config.0.ocpus"),
				resource.TestCheckResourceAttr(resourceName, "private_access_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_access_details.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "standby_size", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "use_dedicated_vm_host"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktop_pools", "test_desktop_pools", acctest.Optional, acctest.Update, DesktopsDesktopPoolDataSourceRepresentation) +
				compartmentIdVariableStr + DesktopsDesktopPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Update, DesktopsDesktopPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "testPool2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "desktop_pool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "desktop_pool_collection.0.items.0.display_name", "testPool2"),
				resource.TestCheckResourceAttr(datasourceName, "desktop_pool_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "desktop_pool_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_collection.0.items.0.maximum_size"),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_collection.0.items.0.active_desktops"),
				resource.TestCheckResourceAttr(datasourceName, "desktop_pool_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Required, acctest.Create, DesktopsDesktopPoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DesktopsDesktopPoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "desktop_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "are_privileged_users", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_policy.0.start_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_policy.0.start_schedule.0.cron_expression"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_policy.0.start_schedule.0.timezone"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_policy.0.stop_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_policy.0.stop_schedule.0.cron_expression"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_policy.0.stop_schedule.0.timezone"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "contact_details", "contactDetails2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.0.audio_mode", "TODESKTOP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.0.cdm_mode", "READONLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.0.clipboard_mode", "TODESKTOP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.0.is_display_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.0.is_keyboard_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.0.is_pointer_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.0.is_printing_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "testPool2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image.0.image_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_storage_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maximum_size", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_access_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_access_details.0.endpoint_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_access_details.0.private_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_access_details.0.vcn_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.baseline_ocpu_utilization"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.ocpus"),
				resource.TestCheckResourceAttr(singularDatasourceName, "session_lifecycle_actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "session_lifecycle_actions.0.inactivity.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "session_lifecycle_actions.0.inactivity.0.action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "session_lifecycle_actions.0.inactivity.0.grace_period_in_minutes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "standby_size", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "use_dedicated_vm_host"),
			),
		},
		// verify resource import
		{
			Config:                  config + DesktopsDesktopPoolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func TestDesktopsDesktopPoolResource_session_lifecycle_disconnect(t *testing.T) {
	httpreplay.SetScenario("TestDesktopsDesktopPoolResource_session_lifecycle_disconnect")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_desktops_desktop_pool.test_desktop_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DesktopsDesktopPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Create, DesktopsDesktopPoolAllSessionLifecycleActionsNoAvailPolicySchedulesRepresentation), "desktops", "desktopPool", t)

	acctest.ResourceTest(t, testAccCheckDesktopsDesktopPoolDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DesktopsDesktopPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Create, DesktopsDesktopPoolAllSessionLifecycleActionsNoAvailPolicySchedulesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_privileged_users", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contact_details", "contactDetails"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.audio_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.cdm_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.clipboard_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_display_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_keyboard_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_pointer_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_printing_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testPool1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_name"),
				resource.TestCheckResourceAttr(resourceName, "is_storage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maximum_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "session_lifecycle_actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "session_lifecycle_actions.0.inactivity.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "session_lifecycle_actions.0.inactivity.0.action"),
				resource.TestCheckResourceAttrSet(resourceName, "session_lifecycle_actions.0.inactivity.0.grace_period_in_minutes"),
				resource.TestCheckResourceAttr(resourceName, "session_lifecycle_actions.0.disconnect.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "session_lifecycle_actions.0.disconnect.0.action"),
				resource.TestCheckResourceAttrSet(resourceName, "session_lifecycle_actions.0.disconnect.0.grace_period_in_minutes"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "standby_size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + DesktopsDesktopPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Optional, acctest.Update, DesktopsDesktopPoolAllSessionLifecycleActionsNoAvailPolicySchedulesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_privileged_users", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "availability_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contact_details", "contactDetails2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.audio_mode", "TODESKTOP"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.cdm_mode", "READONLY"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.clipboard_mode", "TODESKTOP"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_display_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_keyboard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_pointer_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "device_policy.0.is_printing_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testPool2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "image.0.image_name"),
				resource.TestCheckResourceAttr(resourceName, "is_storage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maximum_size", "11"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "standby_size", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckDesktopsDesktopPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DesktopServiceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_desktops_desktop_pool" {
			noResourceFound = false
			request := oci_desktops.GetDesktopPoolRequest{}

			tmp := rs.Primary.ID
			request.DesktopPoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "desktops")

			response, err := client.GetDesktopPool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_desktops.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DesktopsDesktopPool") {
		resource.AddTestSweepers("DesktopsDesktopPool", &resource.Sweeper{
			Name:         "DesktopsDesktopPool",
			Dependencies: acctest.DependencyGraph["desktopPool"],
			F:            sweepDesktopsDesktopPoolResource,
		})
	}
}

func sweepDesktopsDesktopPoolResource(compartment string) error {
	desktopServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DesktopServiceClient()
	desktopPoolIds, err := getDesktopsDesktopPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, desktopPoolId := range desktopPoolIds {
		if ok := acctest.SweeperDefaultResourceId[desktopPoolId]; !ok {
			deleteDesktopPoolRequest := oci_desktops.DeleteDesktopPoolRequest{}

			deleteDesktopPoolRequest.DesktopPoolId = &desktopPoolId

			deleteDesktopPoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "desktops")
			_, error := desktopServiceClient.DeleteDesktopPool(context.Background(), deleteDesktopPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting DesktopPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", desktopPoolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &desktopPoolId, DesktopsDesktopPoolSweepWaitCondition, time.Duration(3*time.Minute),
				DesktopsDesktopPoolSweepResponseFetchOperation, "desktops", true)
		}
	}
	return nil
}

func getDesktopsDesktopPoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DesktopPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	desktopServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DesktopServiceClient()

	listDesktopPoolsRequest := oci_desktops.ListDesktopPoolsRequest{}
	listDesktopPoolsRequest.CompartmentId = &compartmentId
	listDesktopPoolsRequest.LifecycleState = GetLifecycleStateEnumStringValue(oci_desktops.LifecycleStateActive)
	listDesktopPoolsResponse, err := desktopServiceClient.ListDesktopPools(context.Background(), listDesktopPoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DesktopPool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, desktopPool := range listDesktopPoolsResponse.Items {
		id := *desktopPool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DesktopPoolId", id)
	}
	return resourceIds, nil
}

func DesktopsDesktopPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if desktopPoolResponse, ok := response.Response.(oci_desktops.GetDesktopPoolResponse); ok {
		return desktopPoolResponse.LifecycleState != oci_desktops.LifecycleStateDeleted
	}
	return false
}

func DesktopsDesktopPoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DesktopServiceClient().GetDesktopPool(context.Background(), oci_desktops.GetDesktopPoolRequest{
		DesktopPoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func String(v string) *string { return &v }

func GetLifecycleStateEnumStringValue(v oci_desktops.LifecycleStateEnum) *string {
	return String(string(v))
}
