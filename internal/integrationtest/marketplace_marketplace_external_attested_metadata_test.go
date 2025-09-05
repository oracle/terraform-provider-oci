// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MarketplaceMarketplaceExternalAttestedMetadataRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.instance_id}`},
	}

	// 	MarketplaceMarketplaceExternalAttestedMetadataResourceDependencies = utils.OciImageIdsVariable +
	// 		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
	// 		AvailabilityDomainConfig + acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	// 		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: marketplace/default
func TestMarketplaceMarketplaceExternalAttestedMetadataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceMarketplaceExternalAttestedMetadataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	instanceId := utils.GetEnvSettingWithBlankDefault("instance_ocid")
	instanceIdVariableStr := fmt.Sprintf("variable \"instance_id\" { default = \"%s\" }\n", instanceId)

	resourceName := "oci_marketplace_marketplace_external_attested_metadata.test_marketplace_external_attested_metadata"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+instanceIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_marketplace_marketplace_external_attested_metadata", "test_marketplace_external_attested_metadata", acctest.Required, acctest.Create, MarketplaceMarketplaceExternalAttestedMetadataRepresentation), "marketplace", "marketplaceExternalAttestedMetadata", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + instanceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_marketplace_external_attested_metadata", "test_marketplace_external_attested_metadata", acctest.Required, acctest.Create, MarketplaceMarketplaceExternalAttestedMetadataRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_id", instanceId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
