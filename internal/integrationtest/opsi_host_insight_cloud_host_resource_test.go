// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	cloudHostInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_source":  acctest.Representation{RepType: acctest.Required, Create: `MACS_MANAGED_CLOUD_HOST`},
		"compute_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compute_id}`},
		"status":         acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		//"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreCloudHIChangesRepresentation},
	}

	ignoreCloudHIChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// This test basically covers the required field
// issue-routing-tag: opsi/controlPlane
func TestResourceOpsiCloudHostInsight(t *testing.T) {
	httpreplay.SetScenario("TestResourceOpsiCloudHostInsight")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	computeId := utils.GetEnvSettingWithBlankDefault("compute_id")
	if computeId == "" {
		t.Skip("Provision compute instance and point oca/macs to hoth, set compute id = to compute instance ocid to run this test")
	}
	computeIdVariableStr := fmt.Sprintf("variable \"compute_id\" { default = \"%s\" }\n", computeId)

	resourceName := "oci_opsi_host_insight.test_host_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+computeIdVariableStr+OpsiHostInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Required, acctest.Create, cloudHostInsightRequiredRepresentation), "opsi", "hostInsight", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOpsiHostInsightDestroy,
		Steps: []resource.TestStep{
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + computeIdVariableStr + OpsiHostInsightResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Required, acctest.Create, cloudHostInsightRequiredRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

					func(s *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify resource import
			{
				Config:                  config + OpsiHostInsightRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
