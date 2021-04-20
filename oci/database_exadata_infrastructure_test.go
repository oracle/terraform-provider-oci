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
	"github.com/oracle/oci-go-sdk/v40/common"
	oci_database "github.com/oracle/oci-go-sdk/v40/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExadataInfrastructureRequiredOnlyResource = ExadataInfrastructureResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureRepresentation)

	ExadataInfrastructureResourceConfig = ExadataInfrastructureResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, exadataInfrastructureRepresentation)

	exadataInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
	}

	exadataInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `tstExaInfra`},
		"state":          Representation{repType: Optional, create: `REQUIRES_ACTIVATION`},
		"filter":         RepresentationGroup{Required, exadataInfrastructureDataSourceFilterRepresentation}}
	exadataInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`}},
	}

	exadataInfrastructureRepresentation = map[string]interface{}{
		"admin_network_cidr":          Representation{repType: Required, create: `192.168.0.0/16`, update: `192.168.0.0/20`},
		"cloud_control_plane_server1": Representation{repType: Required, create: `10.32.88.1`, update: `10.32.88.2`},
		"cloud_control_plane_server2": Representation{repType: Required, create: `10.32.88.3`, update: `10.32.88.4`},
		"compartment_id":              Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                Representation{repType: Required, create: `tstExaInfra`},
		"dns_server":                  Representation{repType: Required, create: []string{`10.231.225.65`}, update: []string{`10.31.138.25`, `206.223.27.1`}},
		"gateway":                     Representation{repType: Required, create: `10.32.88.5`, update: `10.32.88.6`},
		"infini_band_network_cidr":    Representation{repType: Required, create: `10.31.8.0/21`, update: `10.31.8.0/22`},
		"netmask":                     Representation{repType: Required, create: `255.255.255.0`, update: `255.255.254.0`},
		"ntp_server":                  Representation{repType: Required, create: []string{`10.231.225.76`}, update: []string{`10.246.6.36`, `10.31.138.20`}},
		"shape":                       Representation{repType: Required, create: `ExadataCC.Quarter3.100`},
		"time_zone":                   Representation{repType: Required, create: `US/Pacific`, update: `UTC`},
		"contacts":                    RepresentationGroup{Optional, exadataInfrastructureContactsRepresentation},
		"corporate_proxy":             Representation{repType: Optional, create: `http://192.168.19.1:80`, update: `http://192.168.19.2:80`},
		"defined_tags":                Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":          RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentation},
	}
	exadataInfrastructureContactsRepresentation = map[string]interface{}{
		"email":                    Representation{repType: Required, create: `testuser1@testdomain.com`, update: `testuser2@testdomain.com`},
		"is_primary":               Representation{repType: Required, create: `true`},
		"name":                     Representation{repType: Required, create: `name`, update: `name2`},
		"is_contact_mos_validated": Representation{repType: Optional, create: `false`},
		"phone_number":             Representation{repType: Optional, create: `1234567891`, update: `1234567892`},
	}
	exadataInfrastructureMaintenanceWindowRepresentation = map[string]interface{}{
		"preference":         Representation{repType: Required, create: `NO_PREFERENCE`, update: `CUSTOM_PREFERENCE`},
		"hours_of_day":       Representation{repType: Optional, update: []string{`15`, `20`}},
		"lead_time_in_weeks": Representation{repType: Optional, update: `11`},
		"weeks_of_month":     Representation{repType: Optional, update: []string{`2`, `3`}},
	}

	exadataInfrastructureMaintenanceWindowRepresentationComplete = representationCopyWithNewProperties(exadataInfrastructureMaintenanceWindowRepresentation, map[string]interface{}{
		"days_of_week": []RepresentationGroup{{Optional, exadataInfrastructureMaintenanceWindowDaysOfWeekRepresentation}},
		"months":       []RepresentationGroup{{Optional, exadataInfrastructureMaintenanceWindowMonthsRepresentation1}, {Optional, exadataInfrastructureMaintenanceWindowMonthsRepresentation2}, {Optional, exadataInfrastructureMaintenanceWindowMonthsRepresentation3}, {Optional, exadataInfrastructureMaintenanceWindowMonthsRepresentation4}},
	})

	exadataInfrastructureMaintenanceWindowDaysOfWeekRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, update: `TUESDAY`},
	}
	exadataInfrastructureMaintenanceWindowMonthsRepresentation1 = map[string]interface{}{
		"name": Representation{repType: Required, update: `DECEMBER`},
	}

	exadataInfrastructureMaintenanceWindowMonthsRepresentation2 = map[string]interface{}{
		"name": Representation{repType: Required, update: `MARCH`},
	}

	exadataInfrastructureMaintenanceWindowMonthsRepresentation3 = map[string]interface{}{
		"name": Representation{repType: Required, update: `JUNE`},
	}

	exadataInfrastructureMaintenanceWindowMonthsRepresentation4 = map[string]interface{}{
		"name": Representation{repType: Required, update: `SEPTEMBER`},
	}

	ExadataInfrastructureResourceDependencies = DefinedTagsDependencies
)

func TestDatabaseExadataInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_exadata_infrastructure.test_exadata_infrastructure"
	datasourceName := "data.oci_database_exadata_infrastructures.test_exadata_infrastructures"
	singularDatasourceName := "data.oci_database_exadata_infrastructure.test_exadata_infrastructure"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ExadataInfrastructureResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create, exadataInfrastructureRepresentation), "database", "exadataInfrastructure", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseExadataInfrastructureDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create, exadataInfrastructureRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create,
						representationCopyWithNewProperties(exadataInfrastructureRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
						representationCopyWithNewProperties(exadataInfrastructureRepresentation, map[string]interface{}{
							"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_exadata_infrastructures", "test_exadata_infrastructures", Optional, Update, exadataInfrastructureDataSourceRepresentation) +
					compartmentIdVariableStr + ExadataInfrastructureResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
						representationCopyWithNewProperties(exadataInfrastructureRepresentation, map[string]interface{}{
							"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(datasourceName, "state", "REQUIRES_ACTIVATION"),

					resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.#", "1"),
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
					resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructures.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "exadata_infrastructures.0.time_zone", "UTC"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ExadataInfrastructureResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),

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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

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
	if !inSweeperExcludeList("DatabaseExadataInfrastructure") {
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

			deleteExadataInfrastructureRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExadataInfrastructure(context.Background(), deleteExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting ExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", exadataInfrastructureId, error)
				continue
			}
			waitTillCondition(testAccProvider, &exadataInfrastructureId, exadataInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				exadataInfrastructureSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExadataInfrastructureIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ExadataInfrastructureId")
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
			addResourceIdToSweeperResourceIdMap(compartmentId, "ExadataInfrastructureId", id)
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
