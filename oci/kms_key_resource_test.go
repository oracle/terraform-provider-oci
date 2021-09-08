// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

// issue-routing-tag: kms/default
func TestKmsKeyResource_ResourceDiscovery(t *testing.T) {
	httpreplay.SetScenario("TestKmsKeyResource_ResourceDiscovery")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")
	kmsKeyId := getEnvSettingWithBlankDefault("kms_key_ocid")

	resourceName := "oci_kms_key.test_key"

	ResourceTest(t, nil, []resource.TestStep{
		// verify resource discovery for KMS Keys
		// Our vault is in root compartment, so we need to run Keys resource discovery in root compartment, as first RD tries to find the vault and then keys inside the vault
		{
			Config: config + compartmentIdVariableStr + KeyResourceDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					managementEndpoint, errRead := fromInstanceState(s, "data.oci_kms_vault.test_vault", "management_endpoint")
					if errRead != nil {
						return errRead
					}

					compositeId := "managementEndpoint/" + managementEndpoint + "/keys/" + kmsKeyId
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&compositeId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
