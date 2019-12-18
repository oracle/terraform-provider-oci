// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	KeyResourceVirtualDependencyConfig = KeyResourceVirtualDependencies + DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create, keyVirtualRepresentation)

	keyVirtualRepresentation = getMultipleUpdatedRepresenationCopy([]string{"management_endpoint", "time_of_deletion"}, []interface{}{
		Representation{repType: Required, create: `${data.oci_kms_vault.test_virtual_vault.management_endpoint}`},
		Representation{repType: Required, create: deletionTime.Format(time.RFC3339Nano)}}, keyRepresentation)
	keyVirtualDataSourceRepresentation         = getUpdatedRepresentationCopy("management_endpoint", Representation{repType: Required, create: `${data.oci_kms_vault.test_virtual_vault.management_endpoint}`}, keyDataSourceRepresentation)
	keyVirtualSingularDataSourceRepresentation = getUpdatedRepresentationCopy("management_endpoint", Representation{repType: Required, create: `${data.oci_kms_vault.test_virtual_vault.management_endpoint}`}, keySingularDataSourceRepresentation)

	KeyResourceVirtualDependencies = `
	variable "virtual_vault_ids" {
		type = "map"
		default = {
			us-phoenix-1 = "ocid1.vault.oc1.phx.a5ov6eneaafna.abyhqljs563nxjpbh7hrx73ivi6mjc2xg3q7ljtxmaczys6qoxwg3vyordxq"
			us-ashburn-1 = "ocid1.vault.oc1.iad.bbo7h3seaaeug.abuwcljskeh5tco23lpk5hkijjrlrj64q5afzz3qbe25ku3b4n5ozn66qr5a"	
		}
	}
	data "oci_kms_vault" "test_virtual_vault" {
		#Required
		vault_id = "${var.virtual_vault_ids[var.region]}"
	}
	`
)

func TestResourceKmsKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceKmsKeyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_kms_key.test_key"
	datasourceName := "data.oci_kms_keys.test_keys"
	singularDatasourceName := "data.oci_kms_key.test_key"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckKMSKeyDestroy,
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + KeyResourceVirtualDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keyVirtualRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + KeyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + KeyResourceVirtualDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create, keyVirtualRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + KeyResourceVirtualDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
						representationCopyWithNewProperties(keyVirtualRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + KeyResourceVirtualDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyVirtualRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_kms_keys", "test_keys", Optional, Update, keyVirtualDataSourceRepresentation) +
					compartmentIdVariableStr + KeyResourceVirtualDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyVirtualRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "keys.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "keys.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "keys.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "keys.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "keys.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "keys.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "keys.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "keys.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "keys.0.vault_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keyVirtualSingularDataSourceRepresentation) +
					compartmentIdVariableStr + KeyResourceVirtualDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyVirtualRepresentation) + DefinedTagsDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "current_key_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + KeyResourceVirtualDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyVirtualRepresentation) + DefinedTagsDependencies,
			},
			// revert the updates
			{
				Config: config + compartmentIdVariableStr + KeyResourceVirtualDependencies + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create, keyVirtualRepresentation),
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
					"time_of_deletion",
				},
				ResourceName: resourceName,
			},
		},
	})
}
