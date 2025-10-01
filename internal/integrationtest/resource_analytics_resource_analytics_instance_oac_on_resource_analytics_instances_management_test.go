// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation = map[string]interface{}{
		"attachment_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourceAnalyticsResourceAnalyticsInstanceOacManagementAttachmentDetailsRepresentation},
		"attachment_type":                acctest.Representation{RepType: acctest.Required, Create: `MANAGED`},
		"resource_analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id}`},
		"enable_oac":                     acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	ResourceAnalyticsResourceAnalyticsInstanceOacManagementAttachmentDetailsRepresentation = map[string]interface{}{
		"idcs_domain_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_domain_id}`},
		"license_model":   acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"network_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ResourceAnalyticsResourceAnalyticsInstanceOacManagementAttachmentDetailsNetworkDetailsRepresentation},
		"nsg_ids":         acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.nsg_id}`}},
		"subnet_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.subnet_id}`},
	}
	ResourceAnalyticsResourceAnalyticsInstanceOacManagementAttachmentDetailsNetworkDetailsRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"nsg_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.nsg_id}`}},
	}

	ResourceAnalyticsInstanceOacManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance", "test_resource_analytics_instance", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceRepresentation)
)

// issue-routing-tag: resource_analytics/default
func TestResourceAnalyticsResourceAnalyticsInstanceOacManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceAnalyticsResourceAnalyticsInstanceOacManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	idcsDomainId := utils.GetEnvSettingWithBlankDefault("idcs_domain_id")
	idcsDomainIdStr := fmt.Sprintf("variable \"idcs_domain_id\" { default = \"%s\" }\n", idcsDomainId)

	nsgId := utils.GetEnvSettingWithBlankDefault("nsg_id")
	nsgIdStr := fmt.Sprintf("variable \"nsg_id\" { default = \"%s\" }\n", nsgId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	defaultVarsStr := nsgIdStr + subnetIdVariableStr + idcsDomainIdStr + fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_resource_analytics_resource_analytics_instance_oac_management.test_resource_analytics_instance_oac_management"
	parentResourceName := "oci_resource_analytics_resource_analytics_instance_oac_management.test_resource_analytics_instance_oac_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+defaultVarsStr+ResourceAnalyticsInstanceOacManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance_oac_management", "test_resource_analytics_instance_oac_management", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation), "resourceanalytics", "resourceAnalyticsInstanceOacManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Step 0 - create with enable
		{
			PreConfig: func() {
				fmt.Printf("-=- PreConfig Step 0: %s", config+defaultVarsStr+ResourceAnalyticsInstanceOacManagementResourceDependencies+
					acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance_oac_management", "test_resource_analytics_instance_oac_management", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation))
			},
			Config: config + defaultVarsStr + ResourceAnalyticsInstanceOacManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance_oac_management", "test_resource_analytics_instance_oac_management", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attachment_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "attachment_details.0.idcs_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "attachment_type", "MANAGED"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_analytics_instance_id"),
			),
		},
		// Step 1 - verify enable
		{
			PreConfig: func() {
				fmt.Printf("-=- PreConfig Step 1: %s", config+defaultVarsStr+ResourceAnalyticsInstanceOacManagementResourceDependencies+
					acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance_oac_management", "test_resource_analytics_instance_oac_management", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation))
			},
			Config: config + defaultVarsStr + ResourceAnalyticsInstanceOacManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance_oac_management", "test_resource_analytics_instance_oac_management", acctest.Required, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_oac", "true"),
			),
		},
		// Step 2 - delete before next Create
		{
			Config: config + defaultVarsStr + ResourceAnalyticsInstanceOacManagementResourceDependencies,
		},
		// Step 3 - create with enable and optional fields
		{
			Config: config + defaultVarsStr + ResourceAnalyticsInstanceOacManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance_oac_management", "test_resource_analytics_instance_oac_management", acctest.Optional, acctest.Create, ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attachment_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "attachment_details.0.idcs_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "attachment_type", "MANAGED"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_analytics_instance_id"),
			),
		},
		// Step 4 - update to disable
		{
			Config: config + defaultVarsStr + ResourceAnalyticsInstanceOacManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance_oac_management", "test_resource_analytics_instance_oac_management", acctest.Optional, acctest.Update, ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attachment_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "attachment_details.0.idcs_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "attachment_type", "MANAGED"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_analytics_instance_id"),
			),
		},
		// Step 5 - verify disable
		{
			Config: config + defaultVarsStr + ResourceAnalyticsInstanceOacManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_analytics_resource_analytics_instance_oac_management", "test_resource_analytics_instance_oac_management", acctest.Optional, acctest.Update, ResourceAnalyticsResourceAnalyticsInstanceOacManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_oac", "false"),
			),
		},
	})
}
