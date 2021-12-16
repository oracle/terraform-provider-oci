// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ExadataInfrastructureActivatedResourceConfig = ExadataInfrastructureResourceActivateDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update, exadataInfrastructureActivateRepresentation)

	exadataInfrastructureActivateRepresentation = map[string]interface{}{
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
		"contacts":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureContactsRepresentation},
		"corporate_proxy":             acctest.Representation{RepType: acctest.Optional, Create: `http://192.168.19.1:80`, Update: `http://192.168.19.2:80`},
		//"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
	}

	exadataInfrastructureRepresentationWithContacts = acctest.GetUpdatedRepresentationCopy("contacts", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("phone_number", acctest.Representation{RepType: acctest.Required, Create: `1234567891`, Update: `1234567892`}, exadataInfrastructureContactsRepresentation)}, exadataInfrastructureRepresentation)

	ExadataInfrastructureResourceActivateDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/ExaCC
func TestResourceDatabaseExadataInfrastructure_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseExadataInfrastructure_basic")
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
		// verify Create with activation
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
						"activation_file":    acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
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
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies,
		},
		// verify Create without activation
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(exadataInfrastructureActivateRepresentation, []string{`activation_file`}), map[string]interface{}{
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
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
				resource.TestCheckResourceAttr(resourceName, "state", "REQUIRES_ACTIVATION"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify activation
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
						"activation_file":    acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
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
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
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
	})
}

func createTmpActivationFile() (string, error) {
	activationFile, err := ioutil.TempFile(os.TempDir(), "source-")
	if err != nil {
		return "", err
	}

	text := []byte("dummy activation key")
	if _, err = activationFile.Write(text); err != nil {
		return "", err
	}

	// Close the file
	if err := activationFile.Close(); err != nil {
		return "", err
	}
	return activationFile.Name(), nil
}
