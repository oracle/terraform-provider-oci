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
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	timeNextExecution                               = "2024-05-12T15:00:00Z"
	timeNextExecution2                              = "2024-06-12T15:00:00Z"
	timeStart                                       = "2024-05-05T15:00:00Z"
	timeEnd                                         = "2024-07-12T15:00:00Z"
	OsManagementHubScheduledJobRequiredOnlyResource = OsManagementHubScheduledJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Required, acctest.Create, OsManagementHubScheduledJobRepresentationMIG)

	OsManagementHubScheduledJobResourceConfig = OsManagementHubScheduledJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Update, OsManagementHubScheduledJobRepresentationMIG)

	OsManagementHubScheduledJobSingularDataSourceRepresentation = map[string]interface{}{
		"scheduled_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_scheduled_job.test_scheduled_job.id}`},
	}

	OsManagementHubScheduledJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"display_name_contains":          acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
		"id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_scheduled_job.test_scheduled_job.id}`},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_restricted":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"location":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`OCI_COMPUTE`}},
		"location_not_equal_to":          acctest.Representation{RepType: acctest.Optional, Create: []string{`EC2`}},
		"managed_instance_group_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"operation_type":                 acctest.Representation{RepType: acctest.Optional, Create: `UPDATE_ALL`},
		"schedule_type":                  acctest.Representation{RepType: acctest.Optional, Create: `ONETIME`, Update: `ONETIME`},
		"state":                          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"time_end":                       acctest.Representation{RepType: acctest.Optional, Create: timeEnd},
		"time_start":                     acctest.Representation{RepType: acctest.Optional, Create: timeStart},
		"filter":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobDataSourceFilterRepresentation}}
	OsManagementHubScheduledJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_scheduled_job.test_scheduled_job.id}`}},
	}

	scheduledJobDefinedTagsIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OsManagementHubScheduledJobRepresentationMIG = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationUpdateAll},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `RECURRING`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"managed_instance_group_ids":     acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`}},
		"recurring_rule":                 acctest.Representation{RepType: acctest.Required, Create: ``, Update: `FREQ=DAILY;INTERVAL=1;BYHOUR=10`},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobRepresentationMC = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationUpdateAll},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `RECURRING`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_subcompartment_included":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"locations":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`OCI_COMPUTE`}},
		"managed_compartment_ids":        acctest.Representation{RepType: acctest.Required, Create: []string{`${var.compartment_id}`}},
		"recurring_rule":                 acctest.Representation{RepType: acctest.Required, Create: ``, Update: `FREQ=DAILY;INTERVAL=1;BYHOUR=10`},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobRepresentationLS = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationPromoteLifecycle},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `RECURRING`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle_stage_ids":            acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id}`}},
		"recurring_rule":                 acctest.Representation{RepType: acctest.Required, Create: ``, Update: `FREQ=DAILY;INTERVAL=1;BYHOUR=10`},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobRepresentationMIUpdateAll = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationUpdateAll},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `RECURRING`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"managed_instance_ids":           acctest.Representation{RepType: acctest.Required, Create: []string{`${var.managed_instance_id}`}},
		"recurring_rule":                 acctest.Representation{RepType: acctest.Required, Create: ``, Update: `FREQ=DAILY;INTERVAL=1;BYHOUR=10`},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobRepresentationMIManage = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationManage},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `ONETIME`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"managed_instance_ids":           acctest.Representation{RepType: acctest.Required, Create: []string{`${var.managed_instance_id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobRepresentationMISwitch = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationSwitch},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `ONETIME`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"managed_instance_ids":           acctest.Representation{RepType: acctest.Required, Create: []string{`${var.managed_instance_id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobRepresentationMIPackages = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationPackages},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `ONETIME`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"managed_instance_ids":           acctest.Representation{RepType: acctest.Required, Create: []string{`${var.managed_instance_id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobRepresentationMISoftwareSources = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationSoftwareSources},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `ONETIME`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"managed_instance_ids":           acctest.Representation{RepType: acctest.Required, Create: []string{`${var.managed_instance_id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobRepresentationMIWindows = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operations":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsRepresentationWindows},
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONETIME`, Update: `ONETIME`},
		"time_next_execution":            acctest.Representation{RepType: acctest.Required, Create: timeNextExecution, Update: timeNextExecution2},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_managed_by_autonomous_linux": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"managed_instance_ids":           acctest.Representation{RepType: acctest.Required, Create: []string{`${var.managed_instance_windows_id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: scheduledJobDefinedTagsIgnoreChangesRepresentation},
		"retry_intervals":                acctest.Representation{RepType: acctest.Optional, Create: []int{1, 3}, Update: []int{1, 5}},
	}
	OsManagementHubScheduledJobOperationsRepresentationUpdateAll = map[string]interface{}{
		"operation_type": acctest.Representation{RepType: acctest.Required, Create: `UPDATE_ALL`, Update: nil},
	}
	OsManagementHubScheduledJobOperationsRepresentationManage = map[string]interface{}{
		"operation_type":                acctest.Representation{RepType: acctest.Required, Create: `MANAGE_MODULE_STREAMS`, Update: `MANAGE_MODULE_STREAMS`},
		"manage_module_streams_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsRepresentation},
	}
	OsManagementHubScheduledJobOperationsRepresentationSwitch = map[string]interface{}{
		"operation_type":                acctest.Representation{RepType: acctest.Required, Create: `SWITCH_MODULE_STREAM`, Update: `SWITCH_MODULE_STREAM`},
		"switch_module_streams_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsSwitchModuleStreamsDetailsRepresentation},
	}
	OsManagementHubScheduledJobOperationsRepresentationPackages = map[string]interface{}{
		"operation_type": acctest.Representation{RepType: acctest.Required, Create: `INSTALL_PACKAGES`, Update: `INSTALL_PACKAGES`},
		"package_names":  acctest.Representation{RepType: acctest.Optional, Create: []string{`packageNames`}, Update: []string{`packageNames2`}},
	}
	OsManagementHubScheduledJobOperationsRepresentationSoftwareSources = map[string]interface{}{
		"operation_type":      acctest.Representation{RepType: acctest.Required, Create: `ATTACH_SOFTWARE_SOURCES`, Update: `ATTACH_SOFTWARE_SOURCES`},
		"software_source_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`}, Update: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`}},
	}
	OsManagementHubScheduledJobOperationsRepresentationPromoteLifecycle = map[string]interface{}{
		"operation_type": acctest.Representation{RepType: acctest.Required, Create: `PROMOTE_LIFECYCLE`, Update: `PROMOTE_LIFECYCLE`},
	}
	OsManagementHubScheduledJobOperationsRepresentationWindows = map[string]interface{}{
		"operation_type":       acctest.Representation{RepType: acctest.Required, Create: `INSTALL_WINDOWS_UPDATES`, Update: `INSTALL_WINDOWS_UPDATES`},
		"windows_update_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`windowsUpdateNames`}, Update: []string{`windowsUpdateNames2`}},
	}
	OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsRepresentation = map[string]interface{}{
		"disable": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsDisableRepresentation},
		"enable":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsEnableRepresentation},
		"install": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsInstallRepresentation},
		"remove":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsRemoveRepresentation},
	}
	OsManagementHubScheduledJobOperationsSwitchModuleStreamsDetailsRepresentation = map[string]interface{}{
		"module_name": acctest.Representation{RepType: acctest.Required, Create: `moduleName`, Update: `moduleName2`},
		"stream_name": acctest.Representation{RepType: acctest.Required, Create: `8.0`},
	}
	OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsDisableRepresentation = map[string]interface{}{
		"module_name": acctest.Representation{RepType: acctest.Required, Create: `moduleName`, Update: `moduleName2`},
		"stream_name": acctest.Representation{RepType: acctest.Required, Create: `8.0`},
	}
	OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsEnableRepresentation = map[string]interface{}{
		"module_name": acctest.Representation{RepType: acctest.Required, Create: `moduleName`, Update: `moduleName2`},
		"stream_name": acctest.Representation{RepType: acctest.Required, Create: `8.0`},
	}
	OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsInstallRepresentation = map[string]interface{}{
		"module_name":  acctest.Representation{RepType: acctest.Required, Create: `moduleName`, Update: `moduleName2`},
		"profile_name": acctest.Representation{RepType: acctest.Required, Create: `common`},
		"stream_name":  acctest.Representation{RepType: acctest.Required, Create: `8.0`},
	}
	OsManagementHubScheduledJobOperationsManageModuleStreamsDetailsRemoveRepresentation = map[string]interface{}{
		"module_name":  acctest.Representation{RepType: acctest.Required, Create: `moduleName`, Update: `moduleName2`},
		"profile_name": acctest.Representation{RepType: acctest.Required, Create: `common`},
		"stream_name":  acctest.Representation{RepType: acctest.Required, Create: `8.0`},
	}

	OsManagementHubScheduledJobResourceDependencies = OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubScheduledJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubScheduledJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	managedInstanceId := utils.GetEnvSettingWithBlankDefault("TF_VAR_managed_instance_for_scheduled_job_ocid")
	managedInstanceIdVariableStr := fmt.Sprintf("variable \"managed_instance_id\" { default = \"%s\" }\n", managedInstanceId)

	managedInstanceWindowsId := utils.GetEnvSettingWithBlankDefault("TF_VAR_osmh_managed_instance_windows_ocid")
	managedInstanceWindowsIdVariableStr := fmt.Sprintf("variable \"managed_instance_windows_id\" { default = \"%s\" }\n", managedInstanceWindowsId)

	resourceName := "oci_os_management_hub_scheduled_job.test_scheduled_job"
	datasourceName := "data.oci_os_management_hub_scheduled_jobs.test_scheduled_jobs"
	singularDatasourceName := "data.oci_os_management_hub_scheduled_job.test_scheduled_job"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managedInstanceIdVariableStr+OsManagementHubScheduledJobResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMIUpdateAll), "osmanagementhub", "scheduledJob", t)

	acctest.ResourceTest(t, testAccCheckOsManagementHubScheduledJobDestroy, []resource.TestStep{
		// 1. verify Create
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Required, acctest.Create, OsManagementHubScheduledJobRepresentationMIUpdateAll),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "UPDATE_ALL"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// 2. delete before next Create
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},
		// 3. verify Create with optionals - MI + UPDATE_ALL
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMIUpdateAll),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "UPDATE_ALL"),
				// resource.TestCheckResourceAttr(resourceName, "retry_intervals.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// 4. verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OsManagementHubScheduledJobRepresentationMIUpdateAll, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "UPDATE_ALL"),
				// resource.TestCheckResourceAttr(resourceName, "retry_intervals.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// 5. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Update, OsManagementHubScheduledJobRepresentationMIUpdateAll),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "UPDATE_ALL"),
				resource.TestCheckResourceAttr(resourceName, "recurring_rule", "FREQ=DAILY;INTERVAL=1;BYHOUR=10"),
				// resource.TestCheckResourceAttr(resourceName, "retry_intervals.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "RECURRING"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution2),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// 6. delete before next Create
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},
		// 7. verify Create with optionals - MI + MANAGE_MODULE_STREAMS
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMIManage),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.0.disable.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.0.disable.0.module_name", "moduleName"),
				resource.TestCheckResourceAttrSet(resourceName, "operations.0.manage_module_streams_details.0.disable.0.stream_name"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.0.enable.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.0.enable.0.module_name", "moduleName"),
				resource.TestCheckResourceAttrSet(resourceName, "operations.0.manage_module_streams_details.0.enable.0.stream_name"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.0.install.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.0.install.0.module_name", "moduleName"),
				resource.TestCheckResourceAttrSet(resourceName, "operations.0.manage_module_streams_details.0.install.0.profile_name"),
				resource.TestCheckResourceAttrSet(resourceName, "operations.0.manage_module_streams_details.0.install.0.stream_name"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.0.remove.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.manage_module_streams_details.0.remove.0.module_name", "moduleName"),
				resource.TestCheckResourceAttrSet(resourceName, "operations.0.manage_module_streams_details.0.remove.0.profile_name"),
				resource.TestCheckResourceAttrSet(resourceName, "operations.0.manage_module_streams_details.0.remove.0.stream_name"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "MANAGE_MODULE_STREAMS"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
		// 8. delete before next Create
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},

		// 9. verify Create with optionals - MI + SWITCH_MODULE_STREAM
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMISwitch),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.switch_module_streams_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.switch_module_streams_details.0.module_name", "moduleName"),
				resource.TestCheckResourceAttrSet(resourceName, "operations.0.switch_module_streams_details.0.stream_name"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "SWITCH_MODULE_STREAM"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
		// 10. delete before next Create
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},
		// 11. verify Create with optionals - MI + INSTALL_PACKAGES
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMIPackages),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.package_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "INSTALL_PACKAGES"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
		// 12. delete before next Create
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},
		// 13. verify Create with optionals - MI + INSTALL_WINDOWS_UPDATES
		{
			Config: config + compartmentIdVariableStr + managedInstanceWindowsIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMIWindows),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "INSTALL_WINDOWS_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.windows_update_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
		// 14. delete before next Create
		{
			Config: config + compartmentIdVariableStr + managedInstanceWindowsIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},
		// 15. verify Create with optionals - MI + ATTACH_SOFTWARE_SOURCES
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMISoftwareSources),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "ATTACH_SOFTWARE_SOURCES"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.software_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
		// 16. delete before next Create
		{
			Config: config + compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},

		// 17. verify Create with optionals - MC (ManagedCompartment) + UPDATE_ALL
		{
			Config: config + compartmentIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMC),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_subcompartment_included", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "managed_compartment_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "UPDATE_ALL"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// 18. delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},

		// 19. verify Create with optionals - LS (LifecycleStage) + PROMOTE_LIFECYCLE
		{
			Config: config + compartmentIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationLS),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "lifecycle_stage_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "PROMOTE_LIFECYCLE"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// 20. delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubScheduledJobResourceDependencies,
		},

		// 21. verify Create with optionals - MIG + UPDATE_ALL
		{
			Config: config + compartmentIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Create, OsManagementHubScheduledJobRepresentationMIG),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_group_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "operations.0.operation_type", "UPDATE_ALL"),
				resource.TestCheckResourceAttr(resourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_next_execution", timeNextExecution),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// 22. verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_scheduled_jobs", "test_scheduled_jobs", acctest.Optional, acctest.Update, OsManagementHubScheduledJobDataSourceRepresentation) +
				compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Optional, acctest.Update, OsManagementHubScheduledJobRepresentationMIG),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNameContains"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttr(datasourceName, "is_restricted", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "operation_type", "UPDATE_ALL"),
				resource.TestCheckResourceAttr(datasourceName, "schedule_type", "ONETIME"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_start"),

				// resource.TestCheckResourceAttr(datasourceName, "scheduled_job_collection.#", "1"),
				// resource.TestCheckResourceAttr(datasourceName, "scheduled_job_collection.0.items.#", "1"),
			),
		},
		// 23. verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_scheduled_job", "test_scheduled_job", acctest.Required, acctest.Create, OsManagementHubScheduledJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedInstanceIdVariableStr + OsManagementHubScheduledJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_job_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_managed_by_autonomous_linux", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_restricted"),
				resource.TestCheckResourceAttr(singularDatasourceName, "managed_instance_group_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "operations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "operations.0.operation_type", "UPDATE_ALL"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "retry_intervals.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_type", "RECURRING"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_execution"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_next_execution"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "work_request_ids.#", "1"),
			),
		},
		// 24. verify resource import
		{
			Config:                  config + OsManagementHubScheduledJobRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOsManagementHubScheduledJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ScheduledJobClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_os_management_hub_scheduled_job" {
			noResourceFound = false
			request := oci_os_management_hub.GetScheduledJobRequest{}

			tmp := rs.Primary.ID
			request.ScheduledJobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")

			response, err := client.GetScheduledJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_os_management_hub.ScheduledJobLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OsManagementHubScheduledJob") {
		resource.AddTestSweepers("OsManagementHubScheduledJob", &resource.Sweeper{
			Name:         "OsManagementHubScheduledJob",
			Dependencies: acctest.DependencyGraph["scheduledJob"],
			F:            sweepOsManagementHubScheduledJobResource,
		})
	}
}

func sweepOsManagementHubScheduledJobResource(compartment string) error {
	scheduledJobClient := acctest.GetTestClients(&schema.ResourceData{}).ScheduledJobClient()
	scheduledJobIds, err := getOsManagementHubScheduledJobIds(compartment)
	if err != nil {
		return err
	}
	for _, scheduledJobId := range scheduledJobIds {
		if ok := acctest.SweeperDefaultResourceId[scheduledJobId]; !ok {
			deleteScheduledJobRequest := oci_os_management_hub.DeleteScheduledJobRequest{}

			deleteScheduledJobRequest.ScheduledJobId = &scheduledJobId

			deleteScheduledJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")
			_, error := scheduledJobClient.DeleteScheduledJob(context.Background(), deleteScheduledJobRequest)
			if error != nil {
				fmt.Printf("Error deleting ScheduledJob %s %s, It is possible that the resource is already deleted. Please verify manually \n", scheduledJobId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &scheduledJobId, OsManagementHubScheduledJobSweepWaitCondition, time.Duration(3*time.Minute),
				OsManagementHubScheduledJobSweepResponseFetchOperation, "os_management_hub", true)
		}
	}
	return nil
}

func getOsManagementHubScheduledJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ScheduledJobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	scheduledJobClient := acctest.GetTestClients(&schema.ResourceData{}).ScheduledJobClient()

	listScheduledJobsRequest := oci_os_management_hub.ListScheduledJobsRequest{}
	listScheduledJobsRequest.CompartmentId = &compartmentId
	listScheduledJobsRequest.LifecycleState = oci_os_management_hub.ScheduledJobLifecycleStateActive
	listScheduledJobsResponse, err := scheduledJobClient.ListScheduledJobs(context.Background(), listScheduledJobsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ScheduledJob list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, scheduledJob := range listScheduledJobsResponse.Items {
		id := *scheduledJob.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ScheduledJobId", id)
	}
	return resourceIds, nil
}

func OsManagementHubScheduledJobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if scheduledJobResponse, ok := response.Response.(oci_os_management_hub.GetScheduledJobResponse); ok {
		return scheduledJobResponse.LifecycleState != oci_os_management_hub.ScheduledJobLifecycleStateDeleted
	}
	return false
}

func OsManagementHubScheduledJobSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ScheduledJobClient().GetScheduledJob(context.Background(), oci_os_management_hub.GetScheduledJobRequest{
		ScheduledJobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
