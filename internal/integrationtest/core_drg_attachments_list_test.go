// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
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
	DrgAttachmentsListRequiredOnlyResource = DrgAttachmentsListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", acctest.Required, acctest.Create, drgAttachmentsListRepresentation)

	drgAttachmentsListRepresentation = map[string]interface{}{
		"drg_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"attachment_type":  acctest.Representation{RepType: acctest.Optional, Create: `VCN`},
		"is_cross_tenancy": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DrgAttachmentsListResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Required, acctest.Create, drgAttachmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, drgRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, routeTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/pnp
func TestCoreDrgAttachmentsListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentsListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachments_list.test_drg_attachments_list"

	var resId string
	//Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DrgAttachmentsListResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", acctest.Optional, acctest.Create, drgAttachmentsListRepresentation), "core", "drgAttachmentsList", t)

	acctest.ResourceTest(t, testAccCheckCoreDrgAttachmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentsListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", acctest.Required, acctest.Create, drgAttachmentsListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachments_list", "test_drg_attachments_list", acctest.Optional, acctest.Create, drgAttachmentsListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attachment_type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_all_attachments.0.id"),
				resource.TestCheckResourceAttr(resourceName, "is_cross_tenancy", "false"),

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
