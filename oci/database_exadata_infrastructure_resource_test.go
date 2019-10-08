// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExadataInfrastructureActivatedResourceConfig = ExadataInfrastructureResourceActivateDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, exadataInfrastructureActivateRepresentation)

	exadataInfrastructureActivateRepresentation = map[string]interface{}{
		"admin_network_cidr":          Representation{repType: Required, create: `192.168.19.2/16`, update: `192.168.19.2/20`},
		"cloud_control_plane_server1": Representation{repType: Required, create: `192.168.19.1`, update: `192.168.19.3`},
		"cloud_control_plane_server2": Representation{repType: Required, create: `192.168.19.2`, update: `192.168.19.4`},
		"compartment_id":              Representation{repType: Required, create: `${var.compartment_id}`},
		"corporate_proxy":             Representation{repType: Required, create: `http://192.168.19.1:80`, update: `http://192.168.19.2:80`},
		"display_name":                Representation{repType: Required, create: `tstExaInfra`},
		"dns_server":                  Representation{repType: Required, create: []string{`192.168.10.10`}, update: []string{`192.168.10.11`, `192.168.10.12`}},
		"gateway":                     Representation{repType: Required, create: `192.168.20.1`, update: `192.168.20.2`},
		"infini_band_network_cidr":    Representation{repType: Required, create: `10.172.19.1/24`, update: `10.172.19.1/20`},
		"netmask":                     Representation{repType: Required, create: `255.255.0.0`, update: `255.254.0.0`},
		"ntp_server":                  Representation{repType: Required, create: []string{`192.168.10.20`}, update: []string{`192.168.10.22`, `192.168.10.24`}},
		"shape":                       Representation{repType: Required, create: `ExadataCC.Quarter3.100`},
		"time_zone":                   Representation{repType: Required, create: `US/Pacific`, update: `UTC`},
		"defined_tags":                Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	ExadataInfrastructureResourceActivateDependencies = DefinedTagsDependencies
)

func TestResourceDatabaseExadataInfrastructure_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseExadataInfrastructure_basic")
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
		Steps: []resource.TestStep{
			// verify create with activation
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
						representationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{repType: Optional, update: activationFilePath}})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.19.2/20"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "192.168.19.3"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "192.168.19.4"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.2:80"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "dns_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gateway", "192.168.20.2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.172.19.1/20"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "255.254.0.0"),
					resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies,
			},
			// verify create without activation
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
						representationCopyWithRemovedProperties(exadataInfrastructureActivateRepresentation, []string{`activation_file`})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.19.2/20"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "192.168.19.3"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "192.168.19.4"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.2:80"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "dns_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gateway", "192.168.20.2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.172.19.1/20"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "255.254.0.0"),
					resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "state", "REQUIRES_ACTIVATION"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify activation
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
						representationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{repType: Optional, update: activationFilePath}})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_network_cidr", "192.168.19.2/20"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server1", "192.168.19.3"),
					resource.TestCheckResourceAttr(resourceName, "cloud_control_plane_server2", "192.168.19.4"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "corporate_proxy", "http://192.168.19.2:80"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tstExaInfra"),
					resource.TestCheckResourceAttr(resourceName, "dns_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gateway", "192.168.20.2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "infini_band_network_cidr", "10.172.19.1/20"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "255.254.0.0"),
					resource.TestCheckResourceAttr(resourceName, "ntp_server.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.Quarter3.100"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
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
			// verify update on activated infrastructure
			{
				Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
					generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Create,
						representationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{repType: Optional, update: activationFilePath}})),
				ExpectError: regexp.MustCompile("update not allowed on activated exadata infrastructure"),
			},
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
	log.Printf("activationFile.Name() %s ", activationFile.Name())
	return activationFile.Name(), nil
}
