// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseCloudExadataInfrastructureRequiredOnlyResource = DatabaseCloudExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation)

	DatabaseCloudExadataInfrastructureResourceConfig = DatabaseCloudExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Update, DatabaseCloudExadataInfrastructureRepresentation)

	PeerCeiRepresentation                                                      = acctest.GetUpdatedRepresentationCopy("display_name", acctest.Representation{RepType: acctest.Required, Create: "peerCEInfra"}, DatabaseCloudExadataInfrastructureRepresentation)
	DatabaseDatabaseCloudExadataInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
	}

	DatabaseDatabaseCloudExadataInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_placement_group_id": acctest.Representation{RepType: acctest.Optional, Create: `FakeClusterPlacementGroupId`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `tstExaInfra`, Update: `displayName2`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseCloudExadataInfrastructureDataSourceFilterRepresentation}}
	DatabaseCloudExadataInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`}},
	}

	CloudExadataInfrastructureIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatabaseCloudExadataInfrastructureRepresentation = map[string]interface{}{
		"availability_domain":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `tstExaInfra`, Update: `displayName2`},
		"shape":                      acctest.Representation{RepType: acctest.Required, Create: `Exadata.X8M`},
		"compute_count":              acctest.Representation{RepType: acctest.Required, Create: `2`}, // required for shape Exadata.X8M
		"cluster_placement_group_id": acctest.Representation{RepType: acctest.Optional, Create: `FakeClusterPlacementGroupId`},
		"customer_contacts":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureCustomerContactsRepresentation},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureMaintenanceWindowRepresentation},
		"storage_count":              acctest.Representation{RepType: acctest.Required, Create: `3`}, // required for shape Exadata.X8M
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudExadataInfrastructureIgnoreDefinedTagsRepresentation},
	}

	DatabaseCloudExadataInfrastructureMVMRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `JPzK:SEA-AD-2`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `tstExaInfra`, Update: `displayName2`},
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `Exadata.X8M`},
		"compute_count":       acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`}, // required for shape Exadata.X8M
		"customer_contacts":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureCustomerContactsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureMaintenanceWindowRepresentation},
		"storage_count":       acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `4`}, // required for shape Exadata.X8M
	}

	DatabaseCloudExadataInfrastructureCustomerContactsRepresentation = map[string]interface{}{
		"email": acctest.Representation{RepType: acctest.Optional, Create: `test@oracle.com`, Update: `test2@oracle.com`},
	}
	DatabaseCloudExadataInfrastructureMaintenanceWindowRepresentation = map[string]interface{}{
		"preference": acctest.Representation{RepType: acctest.Optional, Create: `CUSTOM_PREFERENCE`},
		//"custom_action_timeout_in_mins":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"days_of_week": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureMaintenanceWindowDaysOfWeekRepresentation},
		"hours_of_day": acctest.Representation{RepType: acctest.Optional, Create: []string{`4`}, Update: []string{`8`}},
		//"is_custom_action_timeout_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lead_time_in_weeks": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"months":             []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureMaintenanceWindowMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureMaintenanceWindowMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureMaintenanceWindowMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseCloudExadataInfrastructureMaintenanceWindowMonthsRepresentation4}},
		//"patching_mode":                    acctest.Representation{RepType: acctest.Optional, Create: `ROLLING`, Update: `NONROLLING`},
		"weeks_of_month": acctest.Representation{RepType: acctest.Optional, Create: []string{`1`}, Update: []string{`2`}},
	}
	DatabaseCloudExadataInfrastructureMaintenanceWindowDaysOfWeekRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `MONDAY`, Update: `TUESDAY`},
	}
	DatabaseCloudExadataInfrastructureMaintenanceWindowMonthsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `MAY`, Update: `JUNE`},
	}
	DatabaseCloudExadataInfrastructureMaintenanceWindowMonthsRepresentation2 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `FEBRUARY`, Update: `MARCH`},
	}
	DatabaseCloudExadataInfrastructureMaintenanceWindowMonthsRepresentation3 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `AUGUST`, Update: `SEPTEMBER`},
	}
	DatabaseCloudExadataInfrastructureMaintenanceWindowMonthsRepresentation4 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `NOVEMBER`, Update: `DECEMBER`},
	}

	DatabaseCloudExadataInfrastructureResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudExadataInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure"
	datasourceName := "data.oci_database_cloud_exadata_infrastructures.test_cloud_exadata_infrastructures"
	singularDatasourceName := "data.oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseCloudExadataInfrastructureResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation), "database", "cloudExadataInfrastructure", t)

	acctest.ResourceTest(t, testAccCheckDatabaseCloudExadataInfrastructureDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudExadataInfrastructureResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cluster_placement_group_id", "FakeClusterPlacementGroupId"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.custom_action_timeout_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_custom_action_timeout_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "10"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "MAY"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "ROLLING"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseCloudExadataInfrastructureRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cluster_placement_group_id", "FakeClusterPlacementGroupId"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.custom_action_timeout_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_custom_action_timeout_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "10"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "MAY"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "ROLLING"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),

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
			Config: config + compartmentIdVariableStr + DatabaseCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Update, DatabaseCloudExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cluster_placement_group_id", "FakeClusterPlacementGroupId"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_custom_action_timeout_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "11"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "JUNE"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_exadata_infrastructures", "test_cloud_exadata_infrastructures", acctest.Optional, acctest.Update, DatabaseDatabaseCloudExadataInfrastructureDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Update, DatabaseCloudExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "cluster_placement_group_id", "FakeClusterPlacementGroupId"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.available_storage_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.cluster_placement_group_id", "FakeClusterPlacementGroupId"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.compute_count", "2"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.customer_contacts.#", "1"),
				//resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.db_server_version"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.last_maintenance_run_id"), // null for fake resource
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.is_custom_action_timeout_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.lead_time_in_weeks", "11"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.months.0.name", "JUNE"),
				//resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.maintenance_window.0.weeks_of_month.#", "1"),
				//resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.monthly_db_server_version"),
				//resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.monthly_storage_server_version"),
				//resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.next_maintenance_run_id"), // null for fake resource
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_infrastructures.0.storage_count", "3"),
				//resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.storage_server_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructures.0.total_storage_size_in_gbs"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseDatabaseCloudExadataInfrastructureSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseCloudExadataInfrastructureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_exadata_infrastructure_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_storage_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_contacts.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "db_server_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "last_maintenance_run_id"), // null for fake resource
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.is_custom_action_timeout_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.lead_time_in_weeks", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "JUNE"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "monthly_db_server_version"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "monthly_storage_server_version"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "next_maintenance_run_id"), // null for fake resource
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_count", "3"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_server_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_storage_size_in_gbs"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseCloudExadataInfrastructureRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseCloudExadataInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_cloud_exadata_infrastructure" {
			noResourceFound = false
			request := oci_database.GetCloudExadataInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.CloudExadataInfrastructureId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetCloudExadataInfrastructure(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.CloudExadataInfrastructureLifecycleStateTerminated): true,
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

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudExadataInfrastructureResourceMVM(t *testing.T) {
	t.Skip("Skip test because Exacs MVM needs specific compartments enabled`")
	httpreplay.SetScenario("TestDatabaseCloudExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseCloudExadataInfrastructureResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation), "database", "cloudExadataInfrastructure", t)

	acctest.ResourceTest(t, testAccCheckDatabaseCloudExadataInfrastructureDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Create, DatabaseCloudExadataInfrastructureMVMRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.custom_action_timeout_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_custom_action_timeout_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "10"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "MAY"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "ROLLING"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),

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
			Config: config + compartmentIdVariableStr + DatabaseCloudExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Optional, acctest.Update, DatabaseCloudExadataInfrastructureMVMRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "activated_storage_count"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "cpu_count"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttrSet(resourceName, "data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(resourceName, "db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.is_custom_action_timeout_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "11"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "JUNE"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "max_cpu_count"),
				resource.TestCheckResourceAttrSet(resourceName, "max_data_storage_in_tbs"),
				resource.TestCheckResourceAttrSet(resourceName, "max_db_node_storage_in_gbs"),
				resource.TestCheckResourceAttrSet(resourceName, "max_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(resourceName, "memory_size_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_count", "4"),

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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseCloudExadataInfrastructure") {
		resource.AddTestSweepers("DatabaseCloudExadataInfrastructure", &resource.Sweeper{
			Name:         "DatabaseCloudExadataInfrastructure",
			Dependencies: acctest.DependencyGraph["cloudExadataInfrastructure"],
			F:            sweepDatabaseCloudExadataInfrastructureResource,
		})
	}
}

func sweepDatabaseCloudExadataInfrastructureResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	cloudExadataInfrastructureIds, err := getDatabaseCloudExadataInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudExadataInfrastructureId := range cloudExadataInfrastructureIds {
		if ok := acctest.SweeperDefaultResourceId[cloudExadataInfrastructureId]; !ok {
			deleteCloudExadataInfrastructureRequest := oci_database.DeleteCloudExadataInfrastructureRequest{}

			deleteCloudExadataInfrastructureRequest.CloudExadataInfrastructureId = &cloudExadataInfrastructureId

			deleteCloudExadataInfrastructureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteCloudExadataInfrastructure(context.Background(), deleteCloudExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudExadataInfrastructureId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cloudExadataInfrastructureId, DatabaseCloudExadataInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseCloudExadataInfrastructureSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseCloudExadataInfrastructureIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudExadataInfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listCloudExadataInfrastructuresRequest := oci_database.ListCloudExadataInfrastructuresRequest{}
	listCloudExadataInfrastructuresRequest.CompartmentId = &compartmentId
	listCloudExadataInfrastructuresRequest.LifecycleState = oci_database.CloudExadataInfrastructureSummaryLifecycleStateAvailable
	listCloudExadataInfrastructuresResponse, err := databaseClient.ListCloudExadataInfrastructures(context.Background(), listCloudExadataInfrastructuresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudExadataInfrastructure list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudExadataInfrastructure := range listCloudExadataInfrastructuresResponse.Items {
		id := *cloudExadataInfrastructure.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudExadataInfrastructureId", id)
	}
	return resourceIds, nil
}

func DatabaseCloudExadataInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudExadataInfrastructureResponse, ok := response.Response.(oci_database.GetCloudExadataInfrastructureResponse); ok {
		return cloudExadataInfrastructureResponse.LifecycleState != oci_database.CloudExadataInfrastructureLifecycleStateTerminated
	}
	return false
}

func DatabaseCloudExadataInfrastructureSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetCloudExadataInfrastructure(context.Background(), oci_database.GetCloudExadataInfrastructureRequest{
		CloudExadataInfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
