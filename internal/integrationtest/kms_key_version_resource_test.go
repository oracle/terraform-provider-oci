// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	KeyResourceVersionResourceDiscoveryDependencies = KeyVersionResourceDependencies + `
	data "oci_kms_key" "test_key" {
		key_id = "${var.kms_key_id}"
		management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"

	}
	`
)

// issue-routing-tag: kms/default
func TestKmsKeyVersionResource_ResourceDiscovery(t *testing.T) {
	httpreplay.SetScenario("TestKmsKeyVersionResource_ResourceDiscovery")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	os.Setenv("disable_kms_version_delete", "true")

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	resourceName := "oci_kms_key_version.test_key_version"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify resource discovery for KMS Key Versions
		// Our vault is in root compartment, so we need to run resource discovery in root compartment, as first RD tries to find the vault and then keys, versions inside the vault
		{
			Config: config + kmsKeyIdVariableStr + KeyResourceVersionResourceDiscoveryDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					managementEndpoint, errRead := acctest.FromInstanceState(s, "data.oci_kms_vault.test_vault", "management_endpoint")
					if errRead != nil {
						return errRead
					}

					keyVersionId, errRead := acctest.FromInstanceState(s, "data.oci_kms_key.test_key", "current_key_version")
					if errRead != nil {
						return errRead
					}

					compositeId := "managementEndpoint/" + managementEndpoint + "/keys/" + kmsKeyId + "/keyVersions/" + keyVersionId

					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
