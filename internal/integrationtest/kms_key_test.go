// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"

	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	KmsKeyRequiredOnlyResource = KmsKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Required, acctest.Create, KmsKeyRepresentation)

	KmsKeyResourceConfig = KmsKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Update, KmsKeyRepresentation)

	KmsExternalKeyResourceConfig = KmsExternalKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_ext_key", acctest.Optional, acctest.Update, KmsExternalKeyRepresentation)

	KmsKmsKeySingularDataSourceRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key.test_key.id}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	KmsKmsExternalKeySingularDataSourceRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key.test_ext_key.id}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_ext_vault.management_endpoint}`},
	}

	KmsKmsKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"protection_mode":     acctest.Representation{RepType: acctest.Optional, Create: `SOFTWARE`},
		"algorithm":           acctest.Representation{RepType: acctest.Optional, Create: `AES`},
		"length":              acctest.Representation{RepType: acctest.Optional, Create: `16`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsKeyDataSourceFilterRepresentation}}
	KmsKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_kms_key.test_key.id}`}},
	}

	KmsKmsExternalKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_ext_vault.management_endpoint}`},
		"protection_mode":     acctest.Representation{RepType: acctest.Optional, Create: `EXTERNAL`},
		"algorithm":           acctest.Representation{RepType: acctest.Optional, Create: `AES`},
		"length":              acctest.Representation{RepType: acctest.Optional, Create: `32`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsKeyExternalDataSourceFilterRepresentation}}
	KmsKeyExternalDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_kms_key.test_ext_key.id}`}},
	}

	deletionTime = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)

	KmsKeyRepresentation = map[string]interface{}{
		"auto_key_rotation_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: KmsKeyAutoKeyRotationDetailsRepresentation},
		"is_auto_rotation_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `Key C`, Update: `displayName2`},
		"key_shape":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsKeyKeyShapeRepresentation},
		"management_endpoint":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"desired_state":             acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		//"external_key_reference": acctest.RepresentationGroup{RepType: acctest.Optional, Group: KmsKeyExternalKeyReferenceRepresentation},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"protection_mode":  acctest.Representation{RepType: acctest.Optional, Create: `SOFTWARE`},
		"time_of_deletion": acctest.Representation{RepType: acctest.Required, Create: deletionTime.Format(time.RFC3339Nano)},
	}

	KmsExternalKeyRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `External Key C`, Update: `displayName2`},
		"key_shape":              acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsExternalKeyKeyShapeRepresentation},
		"management_endpoint":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_ext_vault.management_endpoint}`},
		"desired_state":          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"external_key_reference": acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsKeyExternalKeyReferenceRepresentation},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"protection_mode":        acctest.Representation{RepType: acctest.Required, Create: `EXTERNAL`},
		"time_of_deletion":       acctest.Representation{RepType: acctest.Required, Create: deletionTime.Format(time.RFC3339Nano)},
	}

	KmsExternalKeyKeyShapeRepresentation = map[string]interface{}{
		"algorithm": acctest.Representation{RepType: acctest.Required, Create: `AES`},
		"length":    acctest.Representation{RepType: acctest.Required, Create: `32`},
	}

	KmsKeyKeyShapeRepresentation = map[string]interface{}{
		"algorithm": acctest.Representation{RepType: acctest.Required, Create: `AES`},
		"length":    acctest.Representation{RepType: acctest.Required, Create: `16`},
	}
	scheduleStartTime                          = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)
	KmsKeyAutoKeyRotationDetailsRepresentation = map[string]interface{}{
		"last_rotation_message":     acctest.Representation{RepType: acctest.Optional, Create: `lastRotationMessage`, Update: `lastRotationMessage2`},
		"last_rotation_status":      acctest.Representation{RepType: acctest.Optional, Create: `SUCCESS`, Update: `FAILED`},
		"rotation_interval_in_days": acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
		"time_of_last_rotation":     acctest.Representation{RepType: acctest.Optional, Create: `timeOfLastRotation`, Update: `timeOfLastRotation2`},
		"time_of_next_rotation":     acctest.Representation{RepType: acctest.Optional, Create: `timeOfNextRotation`, Update: `timeOfNextRotation2`},
		"time_of_schedule_start":    acctest.Representation{RepType: acctest.Optional, Create: scheduleStartTime.Format(time.RFC3339Nano), Update: scheduleStartTime.Format(time.RFC3339Nano)},
	}
	KmsKeyExternalKeyReferenceRepresentation = map[string]interface{}{
		"external_key_id": acctest.Representation{RepType: acctest.Required, Create: `f3cf68ae-659c-4e9e-8be7-ee39fa9ffa3c`},
	}

	kmsVaultId                = utils.GetEnvSettingWithBlankDefault("kms_vault_ocid")
	KmsVaultIdVariableStr     = fmt.Sprintf("variable \"kms_vault_id\" { default = \"%s\" }\n", kmsVaultId)
	kmsKeyIdForCreate         = utils.GetEnvSettingWithBlankDefault("key_ocid_for_create")
	kmsKeyIdCreateVariableStr = fmt.Sprintf("variable \"kms_key_id_for_create\" { default = \"%s\" }\n", kmsKeyIdForCreate)

	kmsKeyIdForUpdate         = utils.GetEnvSettingWithBlankDefault("key_ocid_for_update")
	kmsKeyIdUpdateVariableStr = fmt.Sprintf("variable \"kms_key_id_for_update\" { default = \"%s\" }\n", kmsKeyIdForUpdate)

	kmsKeyCompartmentId            = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	kmsKeyCompartmentIdVariableStr = fmt.Sprintf("variable \"kms_key_compartment_id\" { default = \"%s\" }\n", kmsKeyCompartmentId)
	kmsExternalVaultId             = utils.GetEnvSettingWithBlankDefault("kms_external_vault_ocid")
	KmsExternalVaultIdVariableStr  = fmt.Sprintf("variable \"kms_external_vault_id\" { default = \"%s\" }\n", kmsExternalVaultId)

	// Should deprecate use of tenancy level resources
	KmsKeyResourceDependencies = KmsVaultIdVariableStr + `
	data "oci_kms_vault" "test_vault" {
		#Required
		vault_id = "${var.kms_vault_id}"
	}
	`
	KmsExternalKeyResourceDependencies = KmsExternalVaultIdVariableStr + `
	data "oci_kms_vault" "test_ext_vault" {
		#Required
		vault_id = "${var.kms_external_vault_id}"
	}
	`

	KeyResourceDependencyConfig = KmsKeyResourceDependencies + `
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

	KeyResourceDependencyConfig2 = KmsKeyResourceDependencies + kmsKeyCompartmentIdVariableStr + `
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

func TestExternalKmsKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsExternalKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_key.test_ext_key"
	datasourceName := "data.oci_kms_keys.test_ext_key"
	singularDatasourceName := "data.oci_kms_key.test_ext_key"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KmsExternalKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_ext_key", acctest.Optional, acctest.Create, KmsExternalKeyRepresentation), "keymanagement", "key", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KmsExternalKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_ext_key", acctest.Required, acctest.Create, KmsExternalKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "External Key C"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "32"),
				resource.TestCheckResourceAttr(resourceName, "external_key_reference.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KmsExternalKeyResourceDependencies,
		},

		{
			Config: config + compartmentIdVariableStr + KmsExternalKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_ext_key", acctest.Optional, acctest.Create, KmsExternalKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "External Key C"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "32"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "external_key_reference.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + KmsExternalKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_ext_key", acctest.Optional, acctest.Update, KmsExternalKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "32"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "external_key_reference.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_keys", "test_ext_key", acctest.Optional, acctest.Update, KmsKmsExternalKeyDataSourceRepresentation) +
				compartmentIdVariableStr + KmsExternalKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_ext_key", acctest.Optional, acctest.Update, KmsExternalKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "algorithm", "AES"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "management_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "protection_mode", "EXTERNAL"),
				resource.TestCheckResourceAttr(datasourceName, "length", "32"),

				resource.TestCheckResourceAttr(datasourceName, "keys.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.protection_mode", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.vault_id"),
				resource.TestCheckResourceAttr(datasourceName, "external_key_reference.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_key", "test_ext_key", acctest.Required, acctest.Create, KmsKmsExternalKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + KmsExternalKeyResourceConfig + DefinedTagsDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_key_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_shape.0.length", "32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protection_mode", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
				resource.TestCheckResourceAttr(datasourceName, "external_key_reference.#", "1"),
			),
		},
	})

}

// issue-routing-tag: kms/default
func TestKmsKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_kms_key.test_key"
	datasourceName := "data.oci_kms_keys.test_keys"
	singularDatasourceName := "data.oci_kms_key.test_key"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KmsKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Create, KmsKeyRepresentation), "keymanagement", "key", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KmsKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Required, acctest.Create, KmsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KmsKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + KmsKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Create, KmsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.last_rotation_message", "lastRotationMessage"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.last_rotation_status", "SUCCESS"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_last_rotation", "timeOfLastRotation"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_next_rotation", "timeOfNextRotation"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_schedule_start", scheduleStartTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_rotation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + KmsKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(KmsKeyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.last_rotation_message", "lastRotationMessage"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.last_rotation_status", "SUCCESS"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.rotation_interval_in_days", "60"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_last_rotation", "timeOfLastRotation"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_next_rotation", "timeOfNextRotation"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_schedule_start", scheduleStartTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_rotation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
			Config: config + compartmentIdVariableStr + KmsKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Update, KmsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.last_rotation_message", "lastRotationMessage2"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.last_rotation_status", "FAILED"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.rotation_interval_in_days", "120"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_last_rotation", "timeOfLastRotation2"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_next_rotation", "timeOfNextRotation2"),
				resource.TestCheckResourceAttr(resourceName, "auto_key_rotation_details.0.time_of_schedule_start", scheduleStartTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "current_key_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_rotation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_keys", "test_keys", acctest.Optional, acctest.Update, KmsKmsKeyDataSourceRepresentation) +
				compartmentIdVariableStr + KmsKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Update, KmsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "algorithm", "AES"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "management_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttr(datasourceName, "length", "16"),

				resource.TestCheckResourceAttr(datasourceName, "keys.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.display_name", "displayName2"),
				//resource.TestCheckResourceAttr(datasourceName, "keys.0.external_key_reference_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.is_auto_rotation_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "keys.0.protection_mode", "SOFTWARE"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "keys.0.vault_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Required, acctest.Create, KmsKmsKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + KmsKeyResourceConfig + DefinedTagsDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "auto_key_rotation_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auto_key_rotation_details.0.rotation_interval_in_days", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_key_rotation_details.0.time_of_next_rotation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_key_rotation_details.0.time_of_schedule_start"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_key_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_rotation_enabled", "true"),
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
		// revert the updates
		{
			Config: config + compartmentIdVariableStr + KmsKeyResourceDependencies + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Create, KmsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "Key C"),
				resource.TestCheckResourceAttr(resourceName, "state", "ENABLED"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:            config + KmsKeyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: keyImportId,
			ImportStateVerifyIgnore: []string{
				"external_key_reference",
				"desired_state",
				"time_of_deletion",
				"replica_details",
				"is_auto_rotation_enabled",
				"auto_key_rotation_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckKMSKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_kms_key" {
			client, err := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).KmsManagementClientWithEndpoint(rs.Primary.Attributes["management_endpoint"])
			if err != nil {
				return err
			}

			noResourceFound = false
			request := oci_kms.GetKeyRequest{}

			tmp := rs.Primary.ID
			request.KeyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "kms")

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
