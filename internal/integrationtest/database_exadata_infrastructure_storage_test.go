// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExadataInfrastructureStorageRequiredOnlyResource = ExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure_storage", "test_exadata_infrastructure_storage", acctest.Required, acctest.Create, exadataInfrastructureStorageRepresentation)

	ExadataInfrastructureStorageResourceConfig = ExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure_storage", "test_exadata_infrastructure_storage", acctest.Optional, acctest.Update, exadataInfrastructureStorageRepresentation)

	exadataInfrastructureStorageSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
	}

	exadataInfrastructureStorageDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `tstExaInfra`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `REQUIRES_ACTIVATION`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: exadataInfrastructureDataSourceFilterRepresentation}}
	exadataInfrastructureStorageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`}},
	}

	exadataInfrastructureStorageRepresentation = map[string]interface{}{
		"admin_network_cidr":          acctest.Representation{RepType: acctest.Required, Create: `192.168.0.0/16`, Update: `192.168.0.0/20`},
		"cloud_control_plane_server1": acctest.Representation{RepType: acctest.Required, Create: `10.32.88.1`, Update: `10.32.88.2`},
		"cloud_control_plane_server2": acctest.Representation{RepType: acctest.Required, Create: `10.32.88.3`, Update: `10.32.88.4`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `tstExaInfra`},
		"dns_server":                  acctest.Representation{RepType: acctest.Required, Create: []string{`10.231.225.65`}, Update: []string{`10.31.138.25`, `206.223.27.1`}},
		"gateway":                     acctest.Representation{RepType: acctest.Required, Create: `10.32.88.5`, Update: `10.32.88.6`},
		"infini_band_network_cidr":    acctest.Representation{RepType: acctest.Required, Create: `10.31.8.0/21`, Update: `10.31.8.0/22`},
		"netmask":                     acctest.Representation{RepType: acctest.Required, Create: `255.255.255.0`, Update: `255.255.254.0`},
		"ntp_server":                  acctest.Representation{RepType: acctest.Required, Create: []string{`10.231.225.76`}, Update: []string{`10.246.6.36`, `10.31.138.20`}},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `ExadataCC.Quarter3.100`},
		"time_zone":                   acctest.Representation{RepType: acctest.Required, Create: `US/Pacific`, Update: `UTC`},
		"compute_count":               acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"contacts":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureContactsRepresentation},
		"corporate_proxy":             acctest.Representation{RepType: acctest.Optional, Create: `http://192.168.19.1:80`, Update: `http://192.168.19.2:80`},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentation},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreShapeRepresentation},
	}

	ignoreShapeRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`shape`}},
	}

	ExadataInfrastructureStorageResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/ExaCC
func TestDatabaseExadataInfrastructureStorageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureStorageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_exadata_infrastructure.test_exadata_infrastructure"

	activationFilePath, err := createTmpActivationFile()
	if err != nil {
		t.Fatalf("Unable to Create files for invocation. Error: %q", err)
	}

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{

		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureStorageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(exadataInfrastructureStorageRepresentation, map[string]interface{}{
						"storage_count": 3,
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify activation
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
						"activation_file":    acctest.Representation{RepType: acctest.Optional, Create: activationFilePath},
						"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.3"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.1:80"),
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Update/ scale up infrastructure with additional storage servers
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(exadataInfrastructureStorageRepresentation, map[string]interface{}{
						//"activation_file":    acctest.Representation{RepType: acctest.Optional, Create: activationFilePath},
						"additional_storage_count": acctest.Representation{RepType: acctest.Optional, Update: `3`},
						"maintenance_window":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/20"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.2"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.4"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.2:80"),
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(exadataInfrastructureStorageRepresentation, map[string]interface{}{
						"activation_file":    acctest.Representation{RepType: acctest.Optional, Create: activationFilePath},
						"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.0.0/20"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "10.32.88.2"),
				resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "10.32.88.4"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.2:80"),
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseExadataInfrastructureStorage") {
		resource.AddTestSweepers("DatabaseExadataInfrastructureStorage", &resource.Sweeper{
			Name:         "DatabaseExadataInfrastructureStorage",
			Dependencies: acctest.DependencyGraph["exadataInfrastructureStorage"],
			F:            sweepDatabaseExadataInfrastructureResource,
		})
	}
}
