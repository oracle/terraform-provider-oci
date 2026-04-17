// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// issue-routing-tag: core/default
func TestResourceCoreCrossConnectWithinGroupMacsecUpdateBehavior(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreCrossConnectWithinGroupMacsecUpdateBehavior")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	secretIdCKN := utils.GetEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := utils.GetEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	secretVersionCAK := utils.GetEnvSettingWithBlankDefault("secret_version_cak")
	secretVersionStrCAK := fmt.Sprintf("variable \"secret_version_cak\" { default = \"%s\" }\n", secretVersionCAK)

	secretVersionCKN := utils.GetEnvSettingWithBlankDefault("secret_version_ckn")
	secretVersionStrCKN := fmt.Sprintf("variable \"secret_version_ckn\" { default = \"%s\" }\n", secretVersionCKN)

	if secretIdCKN == "" || secretIdCAK == "" {
		t.Skip("requires secret_ocid_ckn and secret_ocid_cak (or TF_VAR_ prefixed equivalents)")
	}

	locationNameVariableStr := ""

	crossConnectGroupNoMacsecRepresentation := acctest.GetRepresentationCopyWithMultipleRemovedProperties(
		[]string{"defined_tags", "macsec_properties"},
		CoreCrossConnectGroupRepresentation,
	)
	crossConnectGroupWithMacsecRepresentation := map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: "${var.compartment_id}"},
		"customer_reference_name": acctest.Representation{RepType: acctest.Optional, Create: `customerReferenceName`, Update: `customerReferenceName2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName3`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"macsec_properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
			"state":                          acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `ENABLED`},
			"encryption_cipher":              acctest.Representation{RepType: acctest.Optional, Create: `AES256_GCM`, Update: `AES256_GCM_XPN`},
			"is_unprotected_traffic_allowed": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"primary_key": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
				"connectivity_association_key_secret_id":       acctest.Representation{RepType: acctest.Required, Create: "${var.secret_ocid_cak}"},
				"connectivity_association_name_secret_id":      acctest.Representation{RepType: acctest.Required, Create: "${var.secret_ocid_ckn}"},
				"connectivity_association_key_secret_version":  acctest.Representation{RepType: acctest.Optional, Create: "${var.secret_version_cak}", Update: "${var.secret_version_cak}"},
				"connectivity_association_name_secret_version": acctest.Representation{RepType: acctest.Optional, Create: "${var.secret_version_ckn}", Update: "${var.secret_version_ckn}"},
			}},
		}},
	}

	crossConnectWithGroupNoMacsecRepresentation := acctest.RepresentationCopyWithNewProperties(
		acctest.GetRepresentationCopyWithMultipleRemovedProperties(
			[]string{"defined_tags"},
			crossConnectWithGroupRepresentation,
		),
		map[string]interface{}{
			"location_name": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_cross_connect_locations.test_cross_connect_locations.cross_connect_locations.0.name}`}, // not all locations support macsec
		},
	)

	crossConnectWithGroupExplicitMacsecRepresentation := map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"location_name":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_cross_connect_locations.test_cross_connect_locations.cross_connect_locations.0.name}`}, // not all locations support macsec
		"port_speed_shape_name":   acctest.Representation{RepType: acctest.Required, Create: `10 Gbps`},
		"cross_connect_group_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
		"customer_reference_name": acctest.Representation{RepType: acctest.Optional, Create: `customerReferenceName`, Update: `customerReferenceName2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName3`},
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
	singularDatasourceName := "data.oci_core_cross_connect.test_cross_connect"
	var resId, resId2 string

	buildConfig := func(groupRepresentation map[string]interface{}, crossConnectRepresentation map[string]interface{}, lifecycle acctest.RepresentationMode) string {
		dependencies := acctest.GenerateResourceFromRepresentationMap(
			"oci_core_cross_connect_group",
			"test_cross_connect_group",
			acctest.Optional,
			lifecycle,
			groupRepresentation,
		) + acctest.GenerateDataSourceFromRepresentationMap(
			"oci_core_cross_connect_locations",
			"test_cross_connect_locations",
			acctest.Required,
			acctest.Create,
			CoreCoreCrossConnectLocationDataSourceRepresentation,
		)

		return config + compartmentIdVariableStr + locationNameVariableStr + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN + dependencies +
			acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Optional, lifecycle, crossConnectRepresentation)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreCrossConnectWithinGroupMacsecDestroy,
		Steps: []resource.TestStep{
			// Case 1: CCG without MACsec + CC without explicit MACsec => should succeed.
			{
				Config: buildConfig(crossConnectGroupNoMacsecRepresentation, crossConnectWithGroupNoMacsecRepresentation, acctest.Create),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			{
				Config: buildConfig(crossConnectGroupNoMacsecRepresentation, crossConnectWithGroupNoMacsecRepresentation, acctest.Update),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// Case 2: CCG without MACsec + CC with explicit MACsec => should fail (pass-through expected).
			{
				Config:      buildConfig(crossConnectGroupNoMacsecRepresentation, crossConnectWithGroupExplicitMacsecRepresentation, acctest.Update),
				ExpectError: regexp.MustCompile("(?i)(cross.?connect.?group|group|not.*allowed|must.*managed)"),
			},

			// Reset before next case matrix.
			{
				Config: config + compartmentIdVariableStr + locationNameVariableStr + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
			},

			// Case 3: CCG with MACsec + CC without explicit MACsec => should succeed.
			{
				Config: buildConfig(crossConnectGroupWithMacsecRepresentation, crossConnectWithGroupNoMacsecRepresentation, acctest.Create),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			{
				Config: buildConfig(crossConnectGroupWithMacsecRepresentation, crossConnectWithGroupNoMacsecRepresentation, acctest.Update) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Required, acctest.Create, CoreCoreCrossConnectSingularDataSourceRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "macsec_properties.0.state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "macsec_properties.0.encryption_cipher"),
					resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.0.primary_key.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// Case 4: CCG with MACsec + CC with explicit MACsec => should fail (pass-through expected).
			{
				Config:      buildConfig(crossConnectGroupWithMacsecRepresentation, crossConnectWithGroupExplicitMacsecRepresentation, acctest.Update),
				ExpectError: regexp.MustCompile("(?i)(cross.?connect.?group|group|not.*allowed|must.*managed)"),
			},
		},
	})
}

func testAccCheckCoreCrossConnectWithinGroupMacsecDestroy(s *terraform.State) error {
	hasCrossConnect := false
	hasCrossConnectGroup := false
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cross_connect" {
			hasCrossConnect = true
		}
		if rs.Type == "oci_core_cross_connect_group" {
			hasCrossConnectGroup = true
		}
	}

	if hasCrossConnect {
		if err := testAccCheckCoreCrossConnectDestroy(s); err != nil {
			return err
		}
	}
	if hasCrossConnectGroup {
		if err := testAccCheckCoreCrossConnectGroupDestroy(s); err != nil {
			return err
		}
	}

	return nil
}
