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
		GenerateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", Required, Create, drgAttachmentsListRepresentation)

	drgAttachmentsListRepresentation = map[string]interface{}{
		"drg_id":           Representation{RepType: Required, Create: `${oci_core_drg.test_drg.id}`},
		"attachment_type":  Representation{RepType: Optional, Create: `VCN`},
		"is_cross_tenancy": Representation{RepType: Optional, Create: `false`},
	}

	DrgAttachmentsListResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/pnp
func TestCoreDrgAttachmentsListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentsListResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachments_list.test_drg_attachments_list"

	var resId string
	//Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DrgAttachmentsListResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", Optional, Create, drgAttachmentsListRepresentation), "core", "drgAttachmentsList", t)

	ResourceTest(t, testAccCheckCoreDrgAttachmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentsListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", Required, Create, drgAttachmentsListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentsListResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentsListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", Optional, Create, drgAttachmentsListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attachment_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_all_attachments.0.id"),
				resource.TestCheckResourceAttr(resourceName, "is_cross_tenancy", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
