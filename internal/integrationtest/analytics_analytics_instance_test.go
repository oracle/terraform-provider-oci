// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_analytics "github.com/oracle/oci-go-sdk/v58/analytics"
	"github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AnalyticsInstanceRequiredOnlyResource = AnalyticsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Required, acctest.Create, analyticsInstanceRepresentation)

	AnalyticsInstanceResourceConfig = AnalyticsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, analyticsInstanceRepresentation)

	analyticsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
	}

	analyticsInstanceDataSourceRepresentation = map[string]interface{}{
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

	analyticsinstanceName           = utils.RandomString(15, utils.CharsetWithoutDigits)
	analyticsinstanceOptionalName   = utils.RandomString(15, utils.CharsetWithoutDigits)
	analyticsInstanceRepresentation = map[string]interface{}{
		"capacity":       acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceCapacityRepresentation},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"feature_set":    acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_ANALYTICS`},
		"license_type":   acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`, Update: `BRING_YOUR_OWN_LICENSE`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: analyticsinstanceOptionalName},
		//"defined_tags":       acctest.Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"email_notification":       acctest.Representation{RepType: acctest.Optional, Create: `emailNotification`, Update: `emailNotification2`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"idcs_access_token":        acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"network_endpoint_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceNetworkEndpointDetailsRepresentation},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
	}
	analyticsInstanceCapacityRepresentation = map[string]interface{}{
		"capacity_type":  acctest.Representation{RepType: acctest.Required, Create: `OLPU_COUNT`},
		"capacity_value": acctest.Representation{RepType: acctest.Required, Create: `2`},
	}
	analyticsInstanceNetworkEndpointDetailsRepresentation = map[string]interface{}{
		"network_endpoint_type": acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"subnet_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	analyticsInstanceCapacityUpdateRepresentation = map[string]interface{}{
		"capacity_type":  acctest.Representation{RepType: acctest.Required, Create: `OLPU_COUNT`},
		"capacity_value": acctest.Representation{RepType: acctest.Required, Create: `4`},
	}

	AnalyticsInstanceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation)
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AnalyticsInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create, analyticsInstanceRepresentation), "analytics", "analyticsInstance", t)

	acctest.ResourceTest(t, testAccCheckAnalyticsAnalyticsInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("name", acctest.Representation{RepType: acctest.Required, Create: analyticsinstanceName}, analyticsInstanceRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceName),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create, analyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.vcn_id"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_endpoint_details.0.network_endpoint_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_endpoint_details.0.vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, analyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(analyticsInstanceRepresentation, []string{"capacity"}), map[string]interface{}{
						"capacity": acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceCapacityUpdateRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "4"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, analyticsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
					"state": acctest.Representation{RepType: acctest.Required, Create: `INACTIVE`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
					"state": acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_analytics_analytics_instances", "test_analytics_instances", acctest.Optional, acctest.Update, analyticsInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Required, acctest.Create, analyticsInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "analytics_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity.0.capacity_value", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_notification", "emailNotification2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", analyticsinstanceOptionalName),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_endpoint_details.0.network_endpoint_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
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
	analyticsInstanceIds, err := getAnalyticsInstanceIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &analyticsInstanceId, analyticsInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				analyticsInstanceSweepResponseFetchOperation, "analytics", true)
		}
	}
	return nil
}

func getAnalyticsInstanceIds(compartment string) ([]string, error) {
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

func analyticsInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if analyticsInstanceResponse, ok := response.Response.(oci_analytics.GetAnalyticsInstanceResponse); ok {
		return analyticsInstanceResponse.LifecycleState != oci_analytics.AnalyticsInstanceLifecycleStateDeleted
	}
	return false
}

func analyticsInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AnalyticsClient().GetAnalyticsInstance(context.Background(), oci_analytics.GetAnalyticsInstanceRequest{
		AnalyticsInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
