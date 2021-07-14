// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	KeyResourceVersionResourceDiscoveryDependencies = KeyVersionResourceDependencies + `
	data "oci_kms_key" "test_key" {
		key_id = "${var.kms_key_id}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"

	}
	`
)

func TestKmsKeyVersionResource_ResourceDiscovery(t *testing.T) {
	httpreplay.SetScenario("TestKmsKeyVersionResource_ResourceDiscovery")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()
	os.Setenv("disable_kms_version_delete", "true")

	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")
	kmsKeyId := getEnvSettingWithBlankDefault("kms_key_ocid")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	resourceName := "oci_kms_key_version.test_key_version"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify resource discovery for KMS Key Versions
			// Our vault is in root compartment, so we need to run resource discovery in root compartment, as first RD tries to find the vault and then keys, versions inside the vault
			{
				Config: config + kmsKeyIdVariableStr + KeyResourceVersionResourceDiscoveryDependencies,
				Check: ComposeAggregateTestCheckFuncWrapper(

					func(s *terraform.State) (err error) {
						managementEndpoint, errRead := fromInstanceState(s, "data.oci_kms_vault.test_vault", "management_endpoint")
						if errRead != nil {
							return errRead
						}

						keyVersionId, errRead := fromInstanceState(s, "data.oci_kms_key.test_key", "current_key_version")
						if errRead != nil {
							return errRead
						}

						compositeId := "managementEndpoint/" + managementEndpoint + "/keys/" + kmsKeyId + "/keyVersions/" + keyVersionId

						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&compositeId, &tenancyId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
		},
	})
}
