// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExadataInfrastructureStorageRequiredOnlyResource = ExadataInfrastructureResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure_storage", "test_exadata_infrastructure_storage", Required, Create, exadataInfrastructureStorageRepresentation)

	ExadataInfrastructureStorageResourceConfig = ExadataInfrastructureResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure_storage", "test_exadata_infrastructure_storage", Optional, Update, exadataInfrastructureStorageRepresentation)

	exadataInfrastructureStorageSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
	}

	exadataInfrastructureStorageDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `tstExaInfra`},
		"state":          Representation{repType: Optional, create: `REQUIRES_ACTIVATION`},
		"filter":         RepresentationGroup{Required, exadataInfrastructureDataSourceFilterRepresentation}}
	exadataInfrastructureStorageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`}},
	}

	exadataInfrastructureStorageRepresentation = map[string]interface{}{
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
		"compute_count":               Representation{repType: Optional, create: `2`},
		"contacts":                    RepresentationGroup{Optional, exadataInfrastructureContactsRepresentation},
		"corporate_proxy":             Representation{repType: Optional, create: `http://192.168.19.1:80`, update: `http://192.168.19.2:80`},
		"defined_tags":                Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":          RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentation},
		"lifecycle":                   RepresentationGroup{Optional, ignoreShapeRepresentation},
	}

	ignoreShapeRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`shape`}},
	}

	ExadataInfrastructureStorageResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/ExaCC
func TestDatabaseExadataInfrastructureStorageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureStorageResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_exadata_infrastructure.test_exadata_infrastructure"

	activationFilePath, err := createTmpActivationFile()
	if err != nil {
		t.Fatalf("Unable to create files for invocation. Error: %q", err)
	}

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		//CheckDestroy: testAccCheckDatabaseExadataInfrastructureDestroy,
		Steps: []resource.TestStep{

			//verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureStorageResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create,
						representationCopyWithNewProperties(exadataInfrastructureStorageRepresentation, map[string]interface{}{
							"storage_count": 3,
						})),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify activation
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create,
						representationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
							"activation_file":    Representation{repType: Optional, create: activationFilePath},
							"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentation},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.1"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.3"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.1:80"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "dns_server.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gateway", "10.32.88.5"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.31.8.0/21"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// update/ scale up infrastructure with additional storage servers
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
						representationCopyWithNewProperties(exadataInfrastructureStorageRepresentation, map[string]interface{}{
							//"activation_file":    Representation{repType: Optional, create: activationFilePath},
							"additional_storage_count": Representation{repType: Optional, update: `3`},
							"maintenance_window":       RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/20"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.2"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.4"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.2:80"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "dns_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gateway", "10.32.88.6"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.31.8.0/22"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.254.0"),
					resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.X8"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "additional_storage_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "activated_storage_count", "0"),
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
			//activate additional storage cells
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
						representationCopyWithNewProperties(exadataInfrastructureStorageRepresentation, map[string]interface{}{
							"activation_file":    Representation{repType: Optional, create: activationFilePath},
							"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/20"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.2"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.4"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.2:80"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "dns_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gateway", "10.32.88.6"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.31.8.0/22"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.254.0"),
					resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.X8"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// delete
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureStorageResourceDependencies,
			},
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseExadataInfrastructureStorage") {
		resource.AddTestSweepers("DatabaseExadataInfrastructureStorage", &resource.Sweeper{
			Name:         "DatabaseExadataInfrastructureStorage",
			Dependencies: DependencyGraph["exadataInfrastructureStorage"],
			F:            sweepDatabaseExadataInfrastructureResource,
		})
	}
}
