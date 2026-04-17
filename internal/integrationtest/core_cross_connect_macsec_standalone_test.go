// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// issue-routing-tag: core/default
func TestResourceCoreCrossConnectStandaloneMacsecUpdateBehavior(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreCrossConnectStandaloneMacsecUpdateBehavior")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	secretIdCKN := utils.GetEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := utils.GetEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	if secretIdCKN == "" || secretIdCAK == "" {
		t.Skip("requires secret_ocid_ckn and secret_ocid_cak (or TF_VAR_ prefixed equivalents)")
	}

	locationNameVariableStr := ""

	standaloneDeps := acctest.GenerateDataSourceFromRepresentationMap(
		"oci_core_cross_connect_locations",
		"test_cross_connect_locations",
		acctest.Required,
		acctest.Create,
		CoreCoreCrossConnectLocationDataSourceRepresentation,
	)

	standaloneCrossConnectWithoutMacsecRepresentation := map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"location_name":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_cross_connect_locations.test_cross_connect_locations.cross_connect_locations.0.name}`}, // not all locations support macsec
		"port_speed_shape_name":   acctest.Representation{RepType: acctest.Required, Create: `10 Gbps`},
		"customer_reference_name": acctest.Representation{RepType: acctest.Optional, Create: `standaloneRef`, Update: `standaloneRef2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `standaloneDisplay`, Update: `standaloneDisplay2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_active":               acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	standaloneCrossConnectWithMacsecRepresentation := map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"location_name":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_cross_connect_locations.test_cross_connect_locations.cross_connect_locations.0.name}`}, // not all locations support macsec
		"port_speed_shape_name":   acctest.Representation{RepType: acctest.Required, Create: `10 Gbps`},
		"customer_reference_name": acctest.Representation{RepType: acctest.Optional, Create: `standaloneRef`, Update: `standaloneRef2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `standaloneDisplay`, Update: `standaloneDisplay2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_active":               acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"macsec_properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
			"state": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `ENABLED`},
			"encryption_cipher": acctest.Representation{
				RepType: acctest.Optional,
				Create:  `AES256_GCM`,
				Update:  `AES256_GCM_XPN`,
			},
			"is_unprotected_traffic_allowed": acctest.Representation{
				RepType: acctest.Optional,
				Create:  `false`,
				Update:  `true`,
			},
			"primary_key": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
				"connectivity_association_key_secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_ocid_cak}`},
				"connectivity_association_name_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.secret_ocid_ckn}`},
			}},
		}},
	}

	resourceName := "oci_core_cross_connect.test_cross_connect"
	var resId, resId2 string

	buildConfig := func(representation map[string]interface{}, lifecycle acctest.RepresentationMode) string {
		return config + compartmentIdVariableStr + locationNameVariableStr + secretIdVariableStrCKN + secretIdVariableStrCAK + standaloneDeps +
			acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Optional, lifecycle, representation)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreCrossConnectDestroy,
		Steps: []resource.TestStep{
			// Case 1: Standalone CC without MACsec => create/update should succeed.
			{
				Config: buildConfig(standaloneCrossConnectWithoutMacsecRepresentation, acctest.Create),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckNoResourceAttr(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "standaloneRef"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "standaloneDisplay"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			{
				Config: buildConfig(standaloneCrossConnectWithoutMacsecRepresentation, acctest.Update),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckNoResourceAttr(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "standaloneRef2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "standaloneDisplay2"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// Reset before MACsec-enabled case.
			{
				Config: config + compartmentIdVariableStr + locationNameVariableStr + secretIdVariableStrCKN + secretIdVariableStrCAK + standaloneDeps,
			},

			// Case 2: Standalone CC with MACsec => create/update should succeed.
			{
				Config: buildConfig(standaloneCrossConnectWithMacsecRepresentation, acctest.Create),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckNoResourceAttr(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "standaloneRef"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "standaloneDisplay"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.is_unprotected_traffic_allowed", "false"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			{
				Config: buildConfig(standaloneCrossConnectWithMacsecRepresentation, acctest.Update),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckNoResourceAttr(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "standaloneRef2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "standaloneDisplay2"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
					resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.is_unprotected_traffic_allowed", "true"),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
		},
	})
}
