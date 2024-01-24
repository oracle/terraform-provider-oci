// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CrossConnectGroupRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"customer_reference_name": acctest.Representation{RepType: acctest.Optional, Create: `customerReferenceName`, Update: `customerReferenceName2`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"macsec_properties":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CrossConnectGroupMacsecPropertiesRepresentation},
	}

	CrossConnectGroupRepresentationUpdate = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"customer_reference_name": acctest.Representation{RepType: acctest.Optional, Create: `customerReferenceName`, Update: `customerReferenceName2`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"macsec_properties":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CrossConnectGroupMacsecPropertiesRepresentationUpdate},
	}
	CrossConnectGroupMacsecPropertiesRepresentation = map[string]interface{}{
		"state":             acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `ENABLED`},
		"encryption_cipher": acctest.Representation{RepType: acctest.Optional, Create: `AES256_GCM`, Update: `AES256_GCM_XPN`},
		"primary_key":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CrossConnectGroupMacsecPropertiesPrimaryKeyRepresentation},
	}

	CrossConnectGroupMacsecPropertiesRepresentationUpdate = map[string]interface{}{
		"state":             acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `ENABLED`},
		"encryption_cipher": acctest.Representation{RepType: acctest.Optional, Create: `AES256_GCM`, Update: `AES256_GCM_XPN`},
		"primary_key":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CrossConnectGroupMacsecPropertiesPrimaryKeyRepresentationUpdate},
	}
	CrossConnectGroupMacsecPropertiesPrimaryKeyRepresentation = map[string]interface{}{
		"connectivity_association_key_secret_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.secret_ocid_cak}`},
		"connectivity_association_name_secret_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.secret_ocid_ckn}`},
		"connectivity_association_key_secret_version":  acctest.Representation{RepType: acctest.Optional, Update: `${var.secret_version_cak_current}`},
		"connectivity_association_name_secret_version": acctest.Representation{RepType: acctest.Optional, Update: `${var.secret_version_ckn_current}`},
	}

	CrossConnectGroupMacsecPropertiesPrimaryKeyRepresentationUpdate = map[string]interface{}{
		"connectivity_association_key_secret_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.secret_ocid_cak}`},
		"connectivity_association_name_secret_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.secret_ocid_ckn}`},
		"connectivity_association_key_secret_version":  acctest.Representation{RepType: acctest.Optional, Update: `${var.secret_version_cak_prior}`},
		"connectivity_association_name_secret_version": acctest.Representation{RepType: acctest.Optional, Update: `${var.secret_version_ckn_prior}`},
	}

	CrossConnectGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/default
func TestResourceCoreCrossConnectGroupMACSecVersions(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreCrossConnectGroupMACSecVersions")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	secretIdCKN := utils.GetEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := utils.GetEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	secretVersionCAKCurrent := utils.GetEnvSettingWithBlankDefault("secret_version_cak_current")
	secretVersionStrCAKCurrent := fmt.Sprintf("variable \"secret_version_cak_current\" { default = \"%s\" }\n", secretVersionCAKCurrent)

	secretVersionCAKPrior := utils.GetEnvSettingWithBlankDefault("secret_version_cak_prior")
	secretVersionStrCAKPrior := fmt.Sprintf("variable \"secret_version_cak_prior\" { default = \"%s\" }\n", secretVersionCAKPrior)

	secretVersionCKNCurrent := utils.GetEnvSettingWithBlankDefault("secret_version_ckn_current")
	secretVersionStrCKNCurrent := fmt.Sprintf("variable \"secret_version_ckn_current\" { default = \"%s\" }\n", secretVersionCKNCurrent)

	secretVersionCKNPrior := utils.GetEnvSettingWithBlankDefault("secret_version_ckn_prior")
	secretVersionStrCKNPrior := fmt.Sprintf("variable \"secret_version_ckn_prior\" { default = \"%s\" }\n", secretVersionCKNPrior)

	resourceName := "oci_core_cross_connect_group.test_cross_connect_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreCrossConnectGroupDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + CoreCrossConnectGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK +
					acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Create, CrossConnectGroupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id", secretIdCAK),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id", secretIdCKN),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_version", secretVersionCAKCurrent),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_version", secretVersionCKNCurrent),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates - version don't change
			{
				Config: config + compartmentIdVariableStr + CoreCrossConnectGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAKCurrent + secretVersionStrCKNCurrent +
					acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Update, CrossConnectGroupRepresentation),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id", secretIdCAK),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id", secretIdCKN),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_version", secretVersionCAKCurrent),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_version", secretVersionCKNCurrent),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			//verify update - versions change
			{
				Config: config + compartmentIdVariableStr + CoreCrossConnectGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCKNPrior + secretVersionStrCAKPrior +
					acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Update, CrossConnectGroupRepresentationUpdate),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id", secretIdCAK),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id", secretIdCKN),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_version", secretVersionCAKPrior),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_version", secretVersionCKNPrior),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}
