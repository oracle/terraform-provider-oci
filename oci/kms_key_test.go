// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	KeyResourceConfig = KeyResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keyRepresentation)

	keySingularDataSourceRepresentation = map[string]interface{}{
		"key_id":              Representation{repType: Required, create: `${oci_kms_key.test_key.id}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	keyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"filter":              RepresentationGroup{Required, keyDataSourceFilterRepresentation}}
	keyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_kms_key.test_key.id}`}},
	}

	keyRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"display_name":        Representation{repType: Required, create: `Key C`, update: `displayName2`},
		"key_shape":           RepresentationGroup{Required, keyKeyShapeRepresentation},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"desired_state":       Representation{repType: Optional, create: `ENABLED`, update: `DISABLED`},
	}
	keyKeyShapeRepresentation = map[string]interface{}{
		"algorithm": Representation{repType: Required, create: `AES`},
		"length":    Representation{repType: Required, create: `16`},
	}
	KeyResourceDependencies = `
	variable "vault_ids" {
		type = "map"
		default = {
			us-phoenix-1 = "ocid1.vault.oc1.phx.avnzdivwaadfa.abyhqljrmb7herjt4gz64avusyue25grswwsqc5x75im7vtg4x7yfgszqkfa"
			us-ashburn-1 = "ocid1.vault.oc1.iad.annnb3f4aacuu.abuwcljrumuxamzquswnwvgvqdyc76v4e6lyo4372wcjvtdxrhxdc6qxlupq"
		}
	}
	data "oci_kms_vault" "test_vault" {
		#Required
		vault_id = "${var.vault_ids[var.region]}"
	}
	`
	KeyResourceDependencyConfig = KeyResourceDependencies + `
	data "oci_kms_keys" "test_keys_dependency" {
		#Required
		compartment_id = "${var.tenancy_ocid}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"

		filter {
    		name = "state"
    		values = ["ENABLED"]
        }
	}
	`
)

func TestKmsKeyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_kms_key.test_key"
	datasourceName := "data.oci_kms_keys.test_keys"
	singularDatasourceName := "data.oci_kms_key.test_key"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + KeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + KeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
					generateDataSourceFromRepresentationMap("oci_kms_keys", "test_keys", Optional, Update, keyDataSourceRepresentation) +
					compartmentIdVariableStr + KeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(datasourceName, "keys.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "keys.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "keys.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "keys.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "keys.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "keys.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "keys.0.vault_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keySingularDataSourceRepresentation) +
					compartmentIdVariableStr + KeyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "current_key_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.0.algorithm", "AES"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.0.length", "16"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + KeyResourceConfig,
			},
			// revert the updates
			{
				Config: config + compartmentIdVariableStr + KeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create, keyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
					resource.TestCheckResourceAttr(resourceName, "state", "ENABLED"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: keyImportId,
				ImportStateVerifyIgnore: []string{
					"desired_state",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func keyImportId(state *terraform.State) (string, error) {
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "oci_kms_key" {
			return fmt.Sprintf("managementEndpoint/%s/keys/%s", rs.Primary.Attributes["management_endpoint"], rs.Primary.ID), nil
		}
	}

	return "", fmt.Errorf("unable to create import id as no resource of type oci_kms_key in state")
}
