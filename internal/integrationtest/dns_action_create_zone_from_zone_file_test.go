// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DnsActionCreateZoneFromZoneFileRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_dns_action_create_zone_from_zone_file", "test_action_create_zone_from_zone_file", acctest.Required, acctest.Create, DnsActionCreateZoneFromZoneFileRepresentation)

	DnsActionCreateZoneFromZoneFileRepresentation = map[string]interface{}{
		"create_zone_from_zone_file_details": acctest.Representation{RepType: acctest.Required, Create: `createZoneFromZoneFileDetails`},
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"scope":                              acctest.Representation{RepType: acctest.Optional, Create: `GLOBAL`},
	}
)

// issue-routing-tag: dns/default
func TestDnsActionCreateZoneFromZoneFileResource_basic(t *testing.T) {
	t.Skip("This test will not work as the oci_dns_action_create_zone_from_zone_file must be imported to a oci_dns_zone resource to manage the zone beyond creation")
	httpreplay.SetScenario("TestDnsActionCreateZoneFromZoneFileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_action_create_zone_from_zone_file.test_action_create_zone_from_zone_file"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_dns_action_create_zone_from_zone_file", "test_action_create_zone_from_zone_file", acctest.Optional, acctest.Create, DnsActionCreateZoneFromZoneFileRepresentation), "dns", "actionCreateZoneFromZoneFile", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_action_create_zone_from_zone_file", "test_action_create_zone_from_zone_file", acctest.Required, acctest.Create, DnsActionCreateZoneFromZoneFileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "create_zone_from_zone_file_details", "createZoneFromZoneFileDetails"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_action_create_zone_from_zone_file", "test_action_create_zone_from_zone_file", acctest.Optional, acctest.Create, DnsActionCreateZoneFromZoneFileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "create_zone_from_zone_file_details", "createZoneFromZoneFileDetails"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "scope", "GLOBAL"),

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
