// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v54/common"

	oci_kms "github.com/oracle/oci-go-sdk/v54/keymanagement"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	KeyRequiredOnlyResource = KeyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keyRepresentation)

	KeyResourceConfig = KeyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyRepresentation)

	keySingularDataSourceRepresentation = map[string]interface{}{
		"key_id":              Representation{RepType: Required, Create: `${oci_kms_key.test_key.id}`},
		"management_endpoint": Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	keyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"management_endpoint": Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"protection_mode":     Representation{RepType: Optional, Create: `SOFTWARE`},
		"algorithm":           Representation{RepType: Optional, Create: `AES`},
		"length":              Representation{RepType: Optional, Create: `16`},
		"filter":              RepresentationGroup{Required, keyDataSourceFilterRepresentation}}
	keyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_kms_key.test_key.id}`}},
	}

	deletionTime = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)

	keyRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":        Representation{RepType: Required, Create: `Key C`, Update: `displayName2`},
		"key_shape":           RepresentationGroup{Required, keyKeyShapeRepresentation},
		"management_endpoint": Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"desired_state":       Representation{RepType: Optional, Create: `ENABLED`, Update: `DISABLED`},
		"freeform_tags":       Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"protection_mode":     Representation{RepType: Optional, Create: `SOFTWARE`},
		"time_of_deletion":    Representation{RepType: Required, Create: deletionTime.Format(time.RFC3339Nano)},
	}
	keyKeyShapeRepresentation = map[string]interface{}{
		"algorithm": Representation{RepType: Required, Create: `AES`},
		"length":    Representation{RepType: Required, Create: `16`},
	}

	kmsVaultId                = GetEnvSettingWithBlankDefault("kms_vault_ocid")
	KmsVaultIdVariableStr     = fmt.Sprintf("variable \"kms_vault_id\" { default = \"%s\" }\n", kmsVaultId)
	kmsKeyIdForCreate         = GetEnvSettingWithBlankDefault("key_ocid_for_create")
	kmsKeyIdCreateVariableStr = fmt.Sprintf("variable \"kms_key_id_for_create\" { default = \"%s\" }\n", kmsKeyIdForCreate)

	kmsKeyIdForUpdate         = GetEnvSettingWithBlankDefault("key_ocid_for_update")
	kmsKeyIdUpdateVariableStr = fmt.Sprintf("variable \"kms_key_id_for_update\" { default = \"%s\" }\n", kmsKeyIdForUpdate)

	kmsKeyCompartmentId            = GetEnvSettingWithBlankDefault("compartment_ocid")
	kmsKeyCompartmentIdVariableStr = fmt.Sprintf("variable \"kms_key_compartment_id\" { default = \"%s\" }\n", kmsKeyCompartmentId)

	// Should deprecate use of tenancy level resources
	KeyResourceDependencies = KmsVaultIdVariableStr + `
	data "oci_kms_vault" "test_vault" {
		#Required
		vault_id = "${var.kms_vault_id}"
	}
	`
	KeyResourceDependencyConfig = KeyResourceDependencies + `
	data "oci_kms_keys" "test_keys_dependency" {
		#Required
		compartment_id = "${var.tenancy_ocid}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
		algorithm = "AES"

		filter {
    		name = "state"
    		values = ["ENABLED", "UPDATING"]
        }
	}
	data "oci_kms_keys" "test_keys_dependency_RSA" {
		#Required
		compartment_id = "${var.tenancy_ocid}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
		algorithm = "RSA"

		filter {
    		name = "state"
    		values = ["ENABLED", "UPDATING"]
        }
	}
	`

	KeyResourceDependencyConfig2 = KeyResourceDependencies + kmsKeyCompartmentIdVariableStr + `
	data "oci_kms_keys" "test_keys_dependency" {
		#Required
		compartment_id = "${var.kms_key_compartment_id}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
		algorithm = "AES"

		filter {
    		name = "state"
    		values = ["ENABLED", "UPDATING"]
        }
	}
	data "oci_kms_keys" "test_keys_dependency_RSA" {
		#Required
		compartment_id = "${var.kms_key_compartment_id}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
		algorithm = "RSA"

		filter {
    		name = "state"
    		values = ["ENABLED", "UPDATING"]
        }
	}
	`
)

// issue-routing-tag: kms/default
func TestKmsKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_kms_key.test_key"
	datasourceName := "data.oci_kms_keys.test_keys"
	singularDatasourceName := "data.oci_kms_key.test_key"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+KeyResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create, keyRepresentation), "keymanagement", "key", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KeyResourceDependencies + DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + KeyResourceDependencies + DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create, keyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + KeyResourceDependencies + DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
					RepresentationCopyWithNewProperties(keyRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
			Config: config + compartmentIdVariableStr + KeyResourceDependencies + DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
				GenerateDataSourceFromRepresentationMap("oci_kms_keys", "test_keys", Optional, Update, keyDataSourceRepresentation) +
				compartmentIdVariableStr + KeyResourceDependencies + DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Update, keyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "algorithm", "AES"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "management_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttr(datasourceName, "length", "16"),

				resource.TestCheckResourceAttr(datasourceName, "keys.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.vault_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keySingularDataSourceRepresentation) +
				compartmentIdVariableStr + KeyResourceConfig + DefinedTagsDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_key_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.0.length", "16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + KeyResourceConfig + DefinedTagsDependencies,
		},
		// revert the updates
		{
			Config: config + compartmentIdVariableStr + KeyResourceDependencies + DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create, keyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
				resource.TestCheckResourceAttr(resourceName, "state", "ENABLED"),
				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				"replica_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckKMSKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_kms_key" {
			client, err := TestAccProvider.Meta().(*OracleClients).KmsManagementClient(rs.Primary.Attributes["management_endpoint"])
			if err != nil {
				return err
			}

			noResourceFound = false
			request := oci_kms.GetKeyRequest{}

			tmp := rs.Primary.ID
			request.KeyId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "kms")

			response, err := client.GetKey(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_kms.KeyLifecycleStateSchedulingDeletion): true,
					string(oci_kms.KeyLifecycleStatePendingDeletion):    true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}

				if !response.TimeOfDeletion.Equal(deletionTime) && !httpreplay.ModeRecordReplay() {
					return fmt.Errorf("resource time_of_deletion: %s is not set to %s", response.TimeOfDeletion.Format(time.RFC3339Nano), deletionTime.Format(time.RFC3339Nano))
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

func keyImportId(state *terraform.State) (string, error) {
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "oci_kms_key" {
			return fmt.Sprintf("managementEndpoint/%s/keys/%s", rs.Primary.Attributes["management_endpoint"], rs.Primary.ID), nil
		}
	}

	return "", fmt.Errorf("unable to Create import id as no resource of type oci_kms_key in state")
}
