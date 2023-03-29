package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	FileStorageFilesystemSnapshotPolicyInactiveRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-policy-1`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"policy_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `mp1`, Update: `policyPrefix2`},
		"schedules":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageFilesystemSnapshotPolicySchedulesRepresentation},
		"state":               acctest.Representation{RepType: acctest.Optional, Update: `INACTIVE`},
	}

	FileStorageFilesystemSnapshotPolicyActiveRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `media-policy-1`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"policy_prefix":       acctest.Representation{RepType: acctest.Optional, Create: `mp1`, Update: `policyPrefix2`},
		"schedules":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageFilesystemSnapshotPolicyHourlyScheduleRepresentation},
		"state":               acctest.Representation{RepType: acctest.Optional, Update: `ACTIVE`},
	}

	FileStorageFilesystemSnapshotPolicyHourlyScheduleRepresentation = map[string]interface{}{
		"period":                        acctest.Representation{RepType: acctest.Required, Create: `YEARLY`, Update: `HOURLY`},
		"time_zone":                     acctest.Representation{RepType: acctest.Required, Create: `UTC`, Update: `REGIONAL_DATA_CENTER_TIME`},
		"day_of_month":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `0`},
		"day_of_week":                   acctest.Representation{RepType: acctest.Optional, Create: ``, Update: ``},
		"hour_of_day":                   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `0`},
		"month":                         acctest.Representation{RepType: acctest.Optional, Create: `JANUARY`, Update: ``},
		"retention_duration_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `3600`, Update: `4200`},
		"schedule_prefix":               acctest.Representation{RepType: acctest.Optional, Create: `schedulePrefix`, Update: `schedulePrefix2`},
		"time_schedule_start":           acctest.Representation{RepType: acctest.Optional, Create: TimeScheduleStartCreate, Update: TimeScheduleStartUpdate},
	}
)

// issue-routing-tag: file_storage/default
func TestFileStorageFilesystemSnapshotPolicy_pause_unpause(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageFilesystemSnapshotPolicy_pause_unpause")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy"
	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckFileStorageFilesystemSnapshotPolicyDestroy, []resource.TestStep{
		// Create
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
		// Pause policy
		{
			Config: config + compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Update, FileStorageFilesystemSnapshotPolicyInactiveRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.period", "WEEKLY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.retention_duration_in_seconds", "7200"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					// Check that policy is updated, not recreated
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					// Check that inactive policy is discoverable
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// Unpause policy and change schedule to hourly
		{
			Config: config + compartmentIdVariableStr + FileStorageFilesystemSnapshotPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Optional, acctest.Update, FileStorageFilesystemSnapshotPolicyActiveRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.period", "HOURLY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.retention_duration_in_seconds", "4200"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					// Check that policy is updated, not recreated
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					// Check that unpaused policy is discoverable
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
