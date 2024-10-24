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
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FileStorageFilesystemSnapshotPolicyRequiredOnlyResource = FileStorageFilesystemSnapshotPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Required, acctest.Create, FileStorageFilesystemSnapshotPolicyRepresentation)

	FileStorageFilesystemSnapshotPolicyResourceConfig = FileStorageFilesystemSnapshotPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Update, FileStorageFilesystemSnapshotPolicyRepresentation)

	FileStorageFilesystemSnapshotPolicySingularDataSourceRepresentation = map[string]interface{}{
		"filesystem_snapshot_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy.id}`},
	}

	FileStorageFilesystemSnapshotPolicyDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-policy-1`, Update: `displayName2`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageFilesystemSnapshotPolicyDataSourceFilterRepresentation}}
	FileStorageFilesystemSnapshotPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy.id}`}},
	}

	FileStorageFilesystemSnapshotPolicyRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-policy-1`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"policy_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `mp1`, Update: `policyPrefix2`},
		"schedules":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageFilesystemSnapshotPolicySchedulesRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	FileStorageFilesystemSnapshotPolicyRepresentationWithFullLock = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-policy-1`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageFilesystemSnapshotPolicyFullLocksRepresentation},
		"is_lock_override":    acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"policy_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `mp1`, Update: `policyPrefix2`},
		"schedules":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageFilesystemSnapshotPolicySchedulesRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	FileStorageFilesystemSnapshotPolicyRepresentationWithDeleteLock = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-policy-1`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageFilesystemSnapshotPolicyDeleteLocksRepresentation},
		"is_lock_override":    acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"policy_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `mp1`, Update: `policyPrefix2`},
		"schedules":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageFilesystemSnapshotPolicySchedulesRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	FileStorageFilesystemSnapshotPolicyFullLocksRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	FileStorageFilesystemSnapshotPolicyDeleteLocksRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `DELETE`},
		"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	FileStorageFilesystemSnapshotPolicySchedulesRepresentation = map[string]interface{}{
		"period":                        acctest.Representation{RepType: acctest.Required, Create: `YEARLY`, Update: `WEEKLY`},
		"time_zone":                     acctest.Representation{RepType: acctest.Required, Create: `UTC`, Update: `REGIONAL_DATA_CENTER_TIME`},
		"day_of_month":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `0`},
		"day_of_week":                   acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `TUESDAY`},
		"hour_of_day":                   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"month":                         acctest.Representation{RepType: acctest.Optional, Create: `JANUARY`, Update: ``},
		"retention_duration_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `3600`, Update: `7200`},
		"schedule_prefix":               acctest.Representation{RepType: acctest.Optional, Create: `schedulePrefix`, Update: `schedulePrefix2`},
		"time_schedule_start":           acctest.Representation{RepType: acctest.Optional, Create: TimeScheduleStartCreate, Update: TimeScheduleStartUpdate},
	}

	TimeFormat              = "2006-01-02T15:04:05Z"
	TimeScheduleStartCreate = time.Now().Add(10 * time.Hour).Format(TimeFormat)
	TimeScheduleStartUpdate = time.Now().Add(11 * time.Hour).Format(TimeFormat)

	FileStorageFilesystemSnapshotPolicyResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: file_storage/default
func TestFileStorageFilesystemSnapshotPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageFilesystemSnapshotPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy"
	datasourceName := "data.oci_file_storage_filesystem_snapshot_policies.test_filesystem_snapshot_policies"
	singularDatasourceName := "data.oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FileStorageFilesystemSnapshotPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Create, FileStorageFilesystemSnapshotPolicyRepresentation), "filestorage", "filesystemSnapshotPolicy", t)

	acctest.ResourceTest(t, testAccCheckFileStorageFilesystemSnapshotPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Required, acctest.Create, FileStorageFilesystemSnapshotPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies,
		},
		// verify Create with optionals and DELETE lock
		{
			Config: config + compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Create, FileStorageFilesystemSnapshotPolicyRepresentationWithDeleteLock),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-policy-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "DELETE"),
				resource.TestCheckResourceAttr(resourceName, "policy_prefix", "mp1"),
				resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.day_of_month", "10"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.day_of_week", ``),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.hour_of_day", "10"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.month", "JANUARY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.period", "YEARLY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.retention_duration_in_seconds", "3600"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule_prefix", "schedulePrefix"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.time_schedule_start", TimeScheduleStartCreate),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.time_zone", "UTC"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FileStorageFilesystemSnapshotPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "media-policy-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "policy_prefix", "mp1"),
				resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.day_of_month", "10"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.day_of_week", ``),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.hour_of_day", "10"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.month", "JANUARY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.period", "YEARLY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.retention_duration_in_seconds", "3600"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule_prefix", "schedulePrefix"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.time_schedule_start", TimeScheduleStartCreate),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.time_zone", "UTC"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Update, FileStorageFilesystemSnapshotPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "policy_prefix", "policyPrefix2"),
				resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.day_of_week", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.hour_of_day", "11"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.month", ``),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.period", "WEEKLY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.retention_duration_in_seconds", "7200"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule_prefix", "schedulePrefix2"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.time_schedule_start", TimeScheduleStartUpdate),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.time_zone", "REGIONAL_DATA_CENTER_TIME"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policies", "test_filesystem_snapshot_policies", acctest.Optional, acctest.Update, FileStorageFilesystemSnapshotPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Update, FileStorageFilesystemSnapshotPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "filesystem_snapshot_policies.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "filesystem_snapshot_policies.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "filesystem_snapshot_policies.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "filesystem_snapshot_policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "filesystem_snapshot_policies.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "filesystem_snapshot_policies.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "filesystem_snapshot_policies.0.locks.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "filesystem_snapshot_policies.0.locks.0.type", "DELETE"),
				resource.TestCheckResourceAttr(datasourceName, "filesystem_snapshot_policies.0.policy_prefix", "policyPrefix2"),
				resource.TestCheckResourceAttr(datasourceName, "filesystem_snapshot_policies.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "filesystem_snapshot_policies.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Required, acctest.Create, FileStorageFilesystemSnapshotPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "filesystem_snapshot_policy_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.0.type", "DELETE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_prefix", "policyPrefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.day_of_week", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.hour_of_day", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.month", ``),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.period", "WEEKLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.retention_duration_in_seconds", "7200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.schedule_prefix", "schedulePrefix2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedules.0.time_schedule_start"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.time_zone", "REGIONAL_DATA_CENTER_TIME"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageFileSystemResourceDependencies,
		},
		// verify Create with optionals and FULL Lock
		{
			Config: config + compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Create, FileStorageFilesystemSnapshotPolicyRepresentationWithFullLock),
			// FileStorageFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "FULL"),

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
		// verify resource import
		{
			Config:                  config + FileStorageFilesystemSnapshotPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"is_lock_override"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFileStorageFilesystemSnapshotPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_filesystem_snapshot_policy" {
			noResourceFound = false
			request := oci_file_storage.GetFilesystemSnapshotPolicyRequest{}

			tmp := rs.Primary.ID
			request.FilesystemSnapshotPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")

			response, err := client.GetFilesystemSnapshotPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FileStorageFilesystemSnapshotPolicy") {
		resource.AddTestSweepers("FileStorageFilesystemSnapshotPolicy", &resource.Sweeper{
			Name:         "FileStorageFilesystemSnapshotPolicy",
			Dependencies: acctest.DependencyGraph["filesystemSnapshotPolicy"],
			F:            sweepFileStorageFilesystemSnapshotPolicyResource,
		})
	}
}

func sweepFileStorageFilesystemSnapshotPolicyResource(compartment string) error {
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()
	filesystemSnapshotPolicyIds, err := getFileStorageFilesystemSnapshotPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, filesystemSnapshotPolicyId := range filesystemSnapshotPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[filesystemSnapshotPolicyId]; !ok {
			deleteFilesystemSnapshotPolicyRequest := oci_file_storage.DeleteFilesystemSnapshotPolicyRequest{}

			deleteFilesystemSnapshotPolicyRequest.FilesystemSnapshotPolicyId = &filesystemSnapshotPolicyId

			deleteFilesystemSnapshotPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteFilesystemSnapshotPolicy(context.Background(), deleteFilesystemSnapshotPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting FilesystemSnapshotPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", filesystemSnapshotPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &filesystemSnapshotPolicyId, FileStorageFilesystemSnapshotPolicySweepWaitCondition, time.Duration(3*time.Minute),
				FileStorageFilesystemSnapshotPolicySweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getFileStorageFilesystemSnapshotPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FilesystemSnapshotPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()

	listFilesystemSnapshotPoliciesRequest := oci_file_storage.ListFilesystemSnapshotPoliciesRequest{}
	listFilesystemSnapshotPoliciesRequest.CompartmentId = &compartmentId

	availabilityDomains, err := acctest.GetAvalabilityDomains(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting availabilityDomains required for FilesystemSnapshotPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listFilesystemSnapshotPoliciesRequest.AvailabilityDomain = &availabilityDomainName

		listFilesystemSnapshotPoliciesRequest.LifecycleState = oci_file_storage.ListFilesystemSnapshotPoliciesLifecycleStateActive
		listFilesystemSnapshotPoliciesResponse, err := fileStorageClient.ListFilesystemSnapshotPolicies(context.Background(), listFilesystemSnapshotPoliciesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting FilesystemSnapshotPolicy list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, filesystemSnapshotPolicy := range listFilesystemSnapshotPoliciesResponse.Items {
			id := *filesystemSnapshotPolicy.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FilesystemSnapshotPolicyId", id)
		}

	}
	return resourceIds, nil
}

func FileStorageFilesystemSnapshotPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if filesystemSnapshotPolicyResponse, ok := response.Response.(oci_file_storage.GetFilesystemSnapshotPolicyResponse); ok {
		return filesystemSnapshotPolicyResponse.LifecycleState != oci_file_storage.FilesystemSnapshotPolicyLifecycleStateDeleted
	}
	return false
}

func FileStorageFilesystemSnapshotPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FileStorageClient().GetFilesystemSnapshotPolicy(context.Background(), oci_file_storage.GetFilesystemSnapshotPolicyRequest{
		FilesystemSnapshotPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
