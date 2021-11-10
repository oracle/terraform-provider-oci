// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v51/common"
	oci_database "github.com/oracle/oci-go-sdk/v51/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExadataInfrastructureRequiredOnlyResource = ExadataInfrastructureResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureRepresentation)

	ExadataInfrastructureResourceConfig = ExadataInfrastructureResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, exadataInfrastructureRepresentation)

	exadataInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
	}

	exadataInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `tstExaInfra`},
		"state":          Representation{RepType: Optional, Create: `REQUIRES_ACTIVATION`},
		"filter":         RepresentationGroup{Required, exadataInfrastructureDataSourceFilterRepresentation}}
	exadataInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`}},
	}

	exadataInfrastructureRepresentation = map[string]interface{}{
		"admin_network_cidr":          Representation{RepType: Required, Create: `192.168.0.0/16`, Update: `192.168.0.0/20`},
		"cloud_control_plane_server1": Representation{RepType: Required, Create: `10.32.88.1`, Update: `10.32.88.2`},
		"cloud_control_plane_server2": Representation{RepType: Required, Create: `10.32.88.3`, Update: `10.32.88.4`},
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":                Representation{RepType: Required, Create: `tstExaInfra`},
		"dns_server":                  Representation{RepType: Required, Create: []string{`10.231.225.65`}, Update: []string{`10.31.138.25`, `206.223.27.1`}},
		"gateway":                     Representation{RepType: Required, Create: `10.32.88.5`, Update: `10.32.88.6`},
		"infini_band_network_cidr":    Representation{RepType: Required, Create: `10.31.8.0/21`, Update: `10.31.8.0/22`},
		"netmask":                     Representation{RepType: Required, Create: `255.255.255.0`, Update: `255.255.254.0`},
		"ntp_server":                  Representation{RepType: Required, Create: []string{`10.231.225.76`}, Update: []string{`10.246.6.36`, `10.31.138.20`}},
		"shape":                       Representation{RepType: Required, Create: `ExadataCC.Quarter3.100`},
		"time_zone":                   Representation{RepType: Required, Create: `US/Pacific`, Update: `UTC`},
		"storage_count":               Representation{RepType: Optional, Create: `3`},
		"compute_count":               Representation{RepType: Optional, Create: `2`},
		"contacts":                    RepresentationGroup{Optional, exadataInfrastructureContactsRepresentation},
		"corporate_proxy":             Representation{RepType: Optional, Create: `http://192.168.19.1:80`, Update: `http://192.168.19.2:80`},
		"defined_tags":                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":          RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentation},
	}
	exadataInfrastructureContactsRepresentation = map[string]interface{}{
		"email":                    Representation{RepType: Required, Create: `testuser1@testdomain.com`, Update: `testuser2@testdomain.com`},
		"is_primary":               Representation{RepType: Required, Create: `true`},
		"name":                     Representation{RepType: Required, Create: `name`, Update: `name2`},
		"is_contact_mos_validated": Representation{RepType: Optional, Create: `false`},
		"phone_number":             Representation{RepType: Optional, Create: `1234567891`, Update: `1234567892`},
	}
	exadataInfrastructureMaintenanceWindowRepresentation = map[string]interface{}{
		"preference":         Representation{RepType: Required, Create: `NO_PREFERENCE`, Update: `CUSTOM_PREFERENCE`},
		"hours_of_day":       Representation{RepType: Optional, Update: []string{`15`, `20`}},
		"lead_time_in_weeks": Representation{RepType: Optional, Update: `11`},
		"weeks_of_month":     Representation{RepType: Optional, Update: []string{`2`, `3`}},
	}

	exadataInfrastructureMaintenanceWindowRepresentationComplete = RepresentationCopyWithNewProperties(exadataInfrastructureMaintenanceWindowRepresentation, map[string]interface{}{
		"days_of_week": []RepresentationGroup{{Optional, exadataInfrastructureMaintenanceWindowDaysOfWeekRepresentation}},
		"months":       []RepresentationGroup{{Optional, exadataInfrastructureMaintenanceWindowMonthsRepresentation1}, {Optional, exadataInfrastructureMaintenanceWindowMonthsRepresentation2}, {Optional, exadataInfrastructureMaintenanceWindowMonthsRepresentation3}, {Optional, exadataInfrastructureMaintenanceWindowMonthsRepresentation4}},
	})

	exadataInfrastructureMaintenanceWindowDaysOfWeekRepresentation = map[string]interface{}{
		"name": Representation{RepType: Required, Update: `TUESDAY`},
	}
	exadataInfrastructureMaintenanceWindowMonthsRepresentation1 = map[string]interface{}{
		"name": Representation{RepType: Required, Update: `DECEMBER`},
	}

	exadataInfrastructureMaintenanceWindowMonthsRepresentation2 = map[string]interface{}{
		"name": Representation{RepType: Required, Update: `MARCH`},
	}

	exadataInfrastructureMaintenanceWindowMonthsRepresentation3 = map[string]interface{}{
		"name": Representation{RepType: Required, Update: `JUNE`},
	}

	exadataInfrastructureMaintenanceWindowMonthsRepresentation4 = map[string]interface{}{
		"name": Representation{RepType: Required, Update: `SEPTEMBER`},
	}

	ExadataInfrastructureResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/ExaCC
func TestDatabaseExadataInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_exadata_infrastructure.test_exadata_infrastructure"
	datasourceName := "data.oci_database_exadata_infrastructures.test_exadata_infrastructures"
	singularDatasourceName := "data.oci_database_exadata_infrastructure.test_exadata_infrastructure"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ExadataInfrastructureResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create, exadataInfrastructureRepresentation), "database", "exadataInfrastructure", t)

	ResourceTest(t, testAccCheckDatabaseExadataInfrastructureDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.3"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(resourceName, "dns_server.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "gateway", "10.32.88.5"),
				resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.31.8.0/21"),
				resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.255.0"),
				resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create, exadataInfrastructureRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.3"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.email", "testuser1@testdomain.com"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_contact_mos_validated", "false"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.phone_number", "1234567891"),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.1:80"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(resourceName, "dns_server.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gateway", "10.32.88.5"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.31.8.0/21"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.255.0"),
				resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ExadataInfrastructureResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create,
					RepresentationCopyWithNewProperties(exadataInfrastructureRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.3"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.email", "testuser1@testdomain.com"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_contact_mos_validated", "false"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.phone_number", "1234567891"),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.1:80"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(resourceName, "dns_server.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gateway", "10.32.88.5"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.31.8.0/21"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.255.0"),
				resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
					RepresentationCopyWithNewProperties(exadataInfrastructureRepresentation, map[string]interface{}{
						"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/20"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.2"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.4"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.email", "testuser2@testdomain.com"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_contact_mos_validated", "false"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.is_primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "contacts.0.phone_number", "1234567892"),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.2:80"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(resourceName, "dns_server.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gateway", "10.32.88.6"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.31.8.0/22"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.lead_time_in_weeks", "11"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "DECEMBER"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.254.0"),
				resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_database_exadata_infrastructures", "test_exadata_infrastructures", Optional, Update, exadataInfrastructureDataSourceRepresentation) +
				compartmentIdVariableStr + ExadataInfrastructureResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
					RepresentationCopyWithNewProperties(exadataInfrastructureRepresentation, map[string]interface{}{
						"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(datasourceName, "state", "REQUIRES_ACTIVATION"),

				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.activated_storage_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.additional_storage_count"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.admin_network_cidr", "192.168.0.0/20"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.cloud_control_plane_server1", "10.32.88.2"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.cloud_control_plane_server2", "10.32.88.4"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.contacts.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.contacts.0.email", "testuser2@testdomain.com"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.contacts.0.is_contact_mos_validated", "false"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.contacts.0.is_primary", "true"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.contacts.0.name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.contacts.0.phone_number", "1234567892"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.corporate_proxy", "http://192.168.19.2:80"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.cpus_enabled"),
				//resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.csi_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.dns_server.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.gateway", "10.32.88.6"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.infini_band_network_cidr", "10.31.8.0/22"),
				//resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.maintenance_slo_status"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.0.hours_of_day.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.0.lead_time_in_weeks", "11"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.0.months.0.name", "DECEMBER"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.maintenance_window.0.weeks_of_month.#", "2"),

				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.max_cpu_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.max_data_storage_in_tbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.max_db_node_storage_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.max_memory_in_gbs"),

				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.memory_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.netmask", "255.255.254.0"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.ntp_server.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.storage_count", "3"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.compute_count", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.time_zone", "UTC"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ExadataInfrastructureResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "activated_storage_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "additional_storage_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "admin_network_cidr", "192.168.0.0/20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_control_plane_server1", "10.32.88.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_control_plane_server2", "10.32.88.4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.email", "testuser2@testdomain.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.is_contact_mos_validated", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.is_primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "contacts.0.phone_number", "1234567892"),
				resource.TestCheckResourceAttr(singularDatasourceName, "corporate_proxy", "http://192.168.19.2:80"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpus_enabled"),
				// csi_number is not avaliable immediately in the response
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "csi_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "tstExaInfra"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_server.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gateway", "10.32.88.6"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infini_band_network_cidr", "10.31.8.0/22"),
				// maintenance_slo_status is not avaliable immediately in the response
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_slo_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.lead_time_in_weeks", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "DECEMBER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "netmask", "255.255.254.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ntp_server.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_count", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_count", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "UTC"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseExadataInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_exadata_infrastructure" {
			noResourceFound = false
			request := oci_database.GetExadataInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.ExadataInfrastructureId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")

			response, err := client.GetExadataInfrastructure(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExadataInfrastructureLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DatabaseExadataInfrastructure") {
		resource.AddTestSweepers("DatabaseExadataInfrastructure", &resource.Sweeper{
			Name:         "DatabaseExadataInfrastructure",
			Dependencies: DependencyGraph["exadataInfrastructure"],
			F:            sweepDatabaseExadataInfrastructureResource,
		})
	}
}

func sweepDatabaseExadataInfrastructureResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	exadataInfrastructureIds, err := getExadataInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, exadataInfrastructureId := range exadataInfrastructureIds {
		if ok := SweeperDefaultResourceId[exadataInfrastructureId]; !ok {
			deleteExadataInfrastructureRequest := oci_database.DeleteExadataInfrastructureRequest{}

			deleteExadataInfrastructureRequest.ExadataInfrastructureId = &exadataInfrastructureId

			deleteExadataInfrastructureRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExadataInfrastructure(context.Background(), deleteExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting ExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", exadataInfrastructureId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &exadataInfrastructureId, exadataInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				exadataInfrastructureSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExadataInfrastructureIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ExadataInfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	resourceStatesForDestroy := []oci_database.ExadataInfrastructureSummaryLifecycleStateEnum{
		oci_database.ExadataInfrastructureSummaryLifecycleStateRequiresActivation,
		oci_database.ExadataInfrastructureSummaryLifecycleStateActive,
	}
	for _, state := range resourceStatesForDestroy {

		listExadataInfrastructuresRequest := oci_database.ListExadataInfrastructuresRequest{}
		listExadataInfrastructuresRequest.CompartmentId = &compartmentId
		listExadataInfrastructuresRequest.LifecycleState = state
		listExadataInfrastructuresResponse, err := databaseClient.ListExadataInfrastructures(context.Background(), listExadataInfrastructuresRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ExadataInfrastructure list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, exadataInfrastructure := range listExadataInfrastructuresResponse.Items {
			id := *exadataInfrastructure.Id
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "ExadataInfrastructureId", id)
		}
	}
	return resourceIds, nil
}

func exadataInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if exadataInfrastructureResponse, ok := response.Response.(oci_database.GetExadataInfrastructureResponse); ok {
		return exadataInfrastructureResponse.LifecycleState != oci_database.ExadataInfrastructureLifecycleStateDeleted
	}
	return false
}

func exadataInfrastructureSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetExadataInfrastructure(context.Background(), oci_database.GetExadataInfrastructureRequest{
		ExadataInfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
