// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	AnalyticsAnalyticsInstanceRequiredOnlyResource = AnalyticsAnalyticsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Required, acctest.Create, analyticsInstanceRepresentation)

	AnalyticsAnalyticsInstanceResourceConfig = AnalyticsAnalyticsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, analyticsInstanceRepresentation)

	AnalyticsanalyticsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
	}

	AnalyticsanalyticsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"capacity_type":  acctest.Representation{RepType: acctest.Optional, Create: `OLPU_COUNT`},
		"feature_set":    acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_ANALYTICS`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: analyticsinstanceOptionalName},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceDataSourceFilterRepresentation}}
	analyticsInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_analytics_analytics_instance.test_analytics_instance.id}`}},
	}

	analyticsinstanceName         = utils.RandomString(15, utils.CharsetWithoutDigits)
	analyticsinstanceOptionalName = utils.RandomString(15, utils.CharsetWithoutDigits)
	vaultName                     = utils.RandomString(15, utils.CharsetWithoutDigits)

	analyticsInstanceRepresentation = map[string]interface{}{
		"capacity":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceCapacityRepresentation},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"feature_set":              acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_ANALYTICS`},
		"license_type":             acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`, Update: `BRING_YOUR_OWN_LICENSE`},
		"admin_user":               acctest.Representation{RepType: acctest.Optional, Create: `adminUser`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: analyticsinstanceOptionalName},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]map[string]string{"Oracle-Tags": {"CreatedBy": "rbm"}}, Update: map[string]map[string]string{"Oracle-Tags": {"CreatedBy": "dave"}}},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"domain_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_domain.test_domain.id}`},
		"email_notification":       acctest.Representation{RepType: acctest.Optional, Create: `emailNotification`, Update: `emailNotification2`},
		"feature_bundle":           acctest.Representation{RepType: acctest.Required, Create: `EE_EMBEDDED`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"update_channel":           acctest.Representation{RepType: acctest.Optional, Create: `REGULAR`, Update: `EARLY`},
		"kms_key_id":               acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`, Update: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"idcs_access_token":        acctest.Representation{RepType: acctest.Optional, Create: `${var.idcs_access_token}`},
		"network_endpoint_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: analyticsInstanceNetworkEndpointDetailsRepresentation},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
	}

	analyticsPublicInstanceRepresentation = map[string]interface{}{
		"capacity":           acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceCapacityRepresentation},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"feature_set":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_ANALYTICS`},
		"license_type":       acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`, Update: `BRING_YOUR_OWN_LICENSE`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: analyticsinstanceOptionalName},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]map[string]string{"Oracle-Tags": {"CreatedBy": "rbm"}}, Update: map[string]map[string]string{"Oracle-Tags": {"CreatedBy": "dave"}}},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"email_notification": acctest.Representation{RepType: acctest.Optional, Create: `emailNotification`, Update: `emailNotification2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance", "profileId": "oac76"}, Update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":         acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`, Update: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"idcs_access_token":  acctest.Representation{RepType: acctest.Optional, Create: `${var.idcs_access_token}`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
	}

	analyticsInstanceCapacityRepresentation = map[string]interface{}{
		"capacity_type":  acctest.Representation{RepType: acctest.Required, Create: `OLPU_COUNT`},
		"capacity_value": acctest.Representation{RepType: acctest.Required, Create: `2`},
	}
	analyticsInstanceNetworkEndpointDetailsRepresentation = map[string]interface{}{
		"network_endpoint_type":      acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"network_security_group_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"subnet_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	analyticsInstanceNetworkEndpointDetailsUpdateRepresentation = map[string]interface{}{
		"network_endpoint_type":      acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"network_security_group_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.nsg_update_id}`}},
		"subnet_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_update_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.vcn_update_id}`},
	}

	analyticsPublicInstanceNetworkEndpointDetailsUpdateRepresentation = map[string]interface{}{
		"network_endpoint_type": acctest.Representation{RepType: acctest.Required, Create: `PUBLIC`},
		"whitelisted_ips":       acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.whitelisted_ip}`}},
		"whitelisted_vcns":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: whiteListedVCNUpdateRepresentation},
	}

	whiteListedVCNUpdateRepresentation = map[string]interface{}{
		"id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.vcn_whitelist_id}`},
		"whitelisted_ips": acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.whitelisted_ip}`}},
	}

	analyticsInstanceCapacityUpdateRepresentation = map[string]interface{}{
		"capacity_type":  acctest.Representation{RepType: acctest.Required, Create: `OLPU_COUNT`},
		"capacity_value": acctest.Representation{RepType: acctest.Required, Create: `4`},
	}

	analyticsInstanceVaultRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: vaultName},
		"vault_type":     acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
	}

	AnalyticsAnalyticsInstanceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create, IdentityDomainRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: analytics/default
func TestAnalyticsAnalyticsInstanceResource_basic(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAnalyticsAnalyticsInstanceResource_basic") {
		t.Skip("Skipping suppressed TestAnalyticsAnalyticsInstanceResource_basic")
	}

	httpreplay.SetScenario("TestAnalyticsAnalyticsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_analytics_analytics_instance.test_analytics_instance"
	datasourceName := "data.oci_analytics_analytics_instances.test_analytics_instances"
	singularDatasourceName := "data.oci_analytics_analytics_instance.test_analytics_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AnalyticsAnalyticsInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create, analyticsInstanceRepresentation), "analytics", "analyticsInstance", t)

	acctest.ResourceTest(t, testAccCheckAnalyticsAnalyticsInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("name", acctest.Representation{RepType: acctest.Required, Create: analyticsinstanceName}, analyticsInstanceRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceName),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies,
		},

		//create public instance
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create, analyticsPublicInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_user", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification"),
				resource.TestCheckResourceAttr(resourceName, "feature_bundle", "EE_EMBEDDED"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		//update network-endpoint details for public instance
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(analyticsPublicInstanceRepresentation, map[string]interface{}{
						"network_endpoint_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsPublicInstanceNetworkEndpointDetailsUpdateRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		//delete before next Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies,
		},

		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create, analyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "update_channel", "REGULAR"),

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

		//verify change network endpoint
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
						"network_endpoint_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceNetworkEndpointDetailsUpdateRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_user", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		//verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),

			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_bundle", "EE_EMBEDDED"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "update_channel", "REGULAR"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Before testing update, set kms_key_id back to empty string.  In fact, this also effectively tests updating it, but we want to have a key in there after the final
		// update test, or there are problems with the import check at the end.
		// It is necessary to revert the compartment back first, since the new compartment may not have the required permissions to update the key.
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`}})),
		},

		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
						"kms_key_id": acctest.Representation{RepType: acctest.Optional, Create: ``}})),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, analyticsInstanceRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify up scaling
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(analyticsInstanceRepresentation, []string{"capacity"}), map[string]interface{}{
						"capacity": acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceCapacityUpdateRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "4"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify down scaling
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, analyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_user", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(resourceName, "feature_bundle", "EE_EMBEDDED"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify stop
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
					"state": acctest.Representation{RepType: acctest.Required, Create: `INACTIVE`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify start
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
					"state": acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "update_channel", "EARLY"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_analytics_analytics_instances", "test_analytics_instances", acctest.Optional, acctest.Update, AnalyticsanalyticsInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, analyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(datasourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.network_endpoint_details.0.network_endpoint_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.network_endpoint_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.network_endpoint_details.0.vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.service_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Required, acctest.Create, AnalyticsanalyticsInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "analytics_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "feature_bundle", "EE_EMBEDDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.network_endpoint_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_url"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "update_channel", "EARLY"),
			),
		},
		// verify resource import
		{
			Config:            config + AnalyticsAnalyticsInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"admin_user",
				"idcs_access_token",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckAnalyticsAnalyticsInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_analytics_analytics_instance" {
			noResourceFound = false
			request := oci_analytics.GetAnalyticsInstanceRequest{}

			tmp := rs.Primary.ID
			request.AnalyticsInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "analytics")

			response, err := client.GetAnalyticsInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_analytics.AnalyticsInstanceLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("AnalyticsAnalyticsInstance") {
		resource.AddTestSweepers("AnalyticsAnalyticsInstance", &resource.Sweeper{
			Name:         "AnalyticsAnalyticsInstance",
			Dependencies: acctest.DependencyGraph["analyticsInstance"],
			F:            sweepAnalyticsAnalyticsInstanceResource,
		})
	}
}

func sweepAnalyticsAnalyticsInstanceResource(compartment string) error {
	analyticsClient := acctest.GetTestClients(&schema.ResourceData{}).AnalyticsClient()
	analyticsInstanceIds, err := getAnalyticsAnalyticsInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, analyticsInstanceId := range analyticsInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[analyticsInstanceId]; !ok {
			deleteAnalyticsInstanceRequest := oci_analytics.DeleteAnalyticsInstanceRequest{}

			deleteAnalyticsInstanceRequest.AnalyticsInstanceId = &analyticsInstanceId

			deleteAnalyticsInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "analytics")
			_, error := analyticsClient.DeleteAnalyticsInstance(context.Background(), deleteAnalyticsInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting AnalyticsInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", analyticsInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &analyticsInstanceId, AnalyticsanalyticsInstancesSweepWaitCondition, time.Duration(3*time.Minute),
				AnalyticsanalyticsInstancesSweepResponseFetchOperation, "analytics", true)
		}
	}
	return nil
}

func getAnalyticsAnalyticsInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AnalyticsInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	analyticsClient := acctest.GetTestClients(&schema.ResourceData{}).AnalyticsClient()

	listAnalyticsInstancesRequest := oci_analytics.ListAnalyticsInstancesRequest{}
	listAnalyticsInstancesRequest.CompartmentId = &compartmentId
	listAnalyticsInstancesRequest.LifecycleState = oci_analytics.ListAnalyticsInstancesLifecycleStateActive
	listAnalyticsInstancesResponse, err := analyticsClient.ListAnalyticsInstances(context.Background(), listAnalyticsInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AnalyticsInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, analyticsInstance := range listAnalyticsInstancesResponse.Items {
		id := *analyticsInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AnalyticsInstanceId", id)
	}
	return resourceIds, nil
}

func AnalyticsanalyticsInstancesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if analyticsInstanceResponse, ok := response.Response.(oci_analytics.GetAnalyticsInstanceResponse); ok {
		return analyticsInstanceResponse.LifecycleState != oci_analytics.AnalyticsInstanceLifecycleStateDeleted
	}
	return false
}

func AnalyticsanalyticsInstancesSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AnalyticsClient().GetAnalyticsInstance(context.Background(), oci_analytics.GetAnalyticsInstanceRequest{
		AnalyticsInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
