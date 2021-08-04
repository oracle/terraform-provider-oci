// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DrgAttachmentsListRequiredOnlyResource = DrgAttachmentsListResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", Required, Create, drgAttachmentsListRepresentation)

	drgAttachmentsListRepresentation = map[string]interface{}{
		"drg_id":           Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"attachment_type":  Representation{repType: Optional, create: `VCN`},
		"is_cross_tenancy": Representation{repType: Optional, create: `false`},
	}

	DrgAttachmentsListResourceDependencies = generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation) +
		generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/pnp
func TestCoreDrgAttachmentsListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentsListResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachments_list.test_drg_attachments_list"

	var resId string
	//Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DrgAttachmentsListResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", Optional, Create, drgAttachmentsListRepresentation), "core", "drgAttachmentsList", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreDrgAttachmentDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentsListResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", Required, Create, drgAttachmentsListRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentsListResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentsListResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", Optional, Create, drgAttachmentsListRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "VCN"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_all_attachments.0.id"),
					resource.TestCheckResourceAttr(resourceName, "is_cross_tenancy", "false"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
