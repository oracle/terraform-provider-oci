// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentPeerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentPeerResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID          = "compartment_id"
		TEST_SUBNET_ID          = "test_subnet_id"
		LOAD_BALANCER_SUBNET_ID = "load_balancer_subnet_id"
		CERTIFICATE             = "certificate"
		KEY                     = "key"
		PASSWORD                = "password"
		SOURCE_DEPLOYMENT_ID    = "source_deployment_id"
		AVAILABILITY_DOMAIN     = "availability_domain"
		AVAILABILITY_DOMAIN_2   = "availability_domain_2"
		FAULT_DOMAIN            = "fault_domain"
		FAULT_DOMAIN_2          = "fault_domain_2"
	)

	var (
		resourceName = "oci_golden_gate_deployment.test_deployment_local_peers"

		resId  string
		resId2 string

		goldenGateDeploymentPeer1 = map[string]interface{}{
			"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${var.availability_domain}`},
			"fault_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${var.fault_domain}`},
		}
		goldenGateDeploymentPeer2 = map[string]interface{}{
			"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${var.availability_domain_2}`},
			"fault_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${var.fault_domain_2}`},
		}

		ignoreDefinedTagsChangesRepresentation = map[string]interface{}{
			"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
		}
		deploymentMaintenanceConfigurationRepresentation = map[string]interface{}{
			"bundle_release_upgrade_period_in_days":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
			"interim_release_upgrade_period_in_days":  acctest.Representation{RepType: acctest.Optional, Create: `5`},
			"is_interim_release_auto_upgrade_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
			"major_release_upgrade_period_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `20`},
			"security_patch_upgrade_period_in_days":   acctest.Representation{RepType: acctest.Optional, Create: `7`},
		}
		deploymentMaintenanceWindowRepresentation = map[string]interface{}{
			"day":        acctest.Representation{RepType: acctest.Required, Create: `MONDAY`},
			"start_hour": acctest.Representation{RepType: acctest.Required, Create: `10`},
		}
		goldenGateDeploymentOggDataRepresentation = map[string]interface{}{
			"admin_password":  acctest.Representation{RepType: acctest.Required, Create: `${var.password}`},
			"admin_username":  acctest.Representation{RepType: acctest.Required, Create: `adminUsername`},
			"deployment_name": acctest.Representation{RepType: acctest.Required, Create: `test_deployment_local_peers_name`},
			"certificate":     acctest.Representation{RepType: acctest.Optional, Update: `${var.certificate}`},
			"key":             acctest.Representation{RepType: acctest.Optional, Update: `${var.key}`},
		}

		goldenGateDeploymentRepresentation = map[string]interface{}{
			"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"cpu_core_count":            acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
			"deployment_type":           acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_ORACLE`},
			"display_name":              acctest.Representation{RepType: acctest.Required, Create: `Terraform_integration_test_peers`, Update: `Terraform_integration_test_peers_2`},
			"is_auto_scaling_enabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.test_subnet_id}`},
			"license_model":             acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
			"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`},
			"fqdn":                      acctest.Representation{RepType: acctest.Optional, Update: `fqdn1.oggdevops.us`},
			"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
			"is_public":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
			"ogg_data":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDeploymentOggDataRepresentation},
			"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
			"maintenance_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentMaintenanceConfigurationRepresentation},
			"maintenance_window":        acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentMaintenanceWindowRepresentation},
			"availability_domain":       acctest.Representation{RepType: acctest.Optional, Create: `${var.availability_domain}`},
			"fault_domain":              acctest.Representation{RepType: acctest.Optional, Create: `${var.fault_domain}`},
			"placements":                []acctest.RepresentationGroup{}, // start with empty peer list
			"source_deployment_id":      acctest.Representation{RepType: acctest.Optional, Create: nil},
		}

		compartmentId       = utils.GetEnvSettingWithBlankDefault(COMPARTMENT_ID)
		subnetId            = utils.GetEnvSettingWithBlankDefault(TEST_SUBNET_ID)
		sourceDeploymentId  = utils.GetEnvSettingWithBlankDefault(SOURCE_DEPLOYMENT_ID)
		availabilityDomain  = utils.GetEnvSettingWithBlankDefault(AVAILABILITY_DOMAIN)
		faultDomain         = utils.GetEnvSettingWithBlankDefault(FAULT_DOMAIN)
		availabilityDomain2 = utils.GetEnvSettingWithBlankDefault(AVAILABILITY_DOMAIN_2)
		faultDomain2        = utils.GetEnvSettingWithBlankDefault(FAULT_DOMAIN_2)
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(TEST_SUBNET_ID, t) +
		makeVariableStr(LOAD_BALANCER_SUBNET_ID, t) +
		makeVariableStr(CERTIFICATE, t) +
		makeVariableStr(KEY, t) +
		makeVariableStr(PASSWORD, t) +
		makeVariableStr(SOURCE_DEPLOYMENT_ID, t) +
		makeVariableStr(AVAILABILITY_DOMAIN, t) +
		makeVariableStr(AVAILABILITY_DOMAIN_2, t) +
		makeVariableStr(FAULT_DOMAIN, t) +
		makeVariableStr(FAULT_DOMAIN_2, t)

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify standBy deployment create
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment_local_peers", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"source_deployment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.source_deployment_id}`}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test_peers"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "source_deployment_id", sourceDeploymentId),
				resource.TestCheckResourceAttr(resourceName, "availability_domain", availabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", faultDomain),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_id"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config,
		},

		// verify create deployment
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment_local_peers", acctest.Optional,
					acctest.Create, goldenGateDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "availability_domain", availabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", faultDomain),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test_peers"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_id"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify add local peer
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment_local_peers", acctest.Optional,
					acctest.Update, acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"placements": []acctest.RepresentationGroup{
							{RepType: acctest.Optional, Group: goldenGateDeploymentPeer1},
							{RepType: acctest.Optional, Group: goldenGateDeploymentPeer2},
						},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "availability_domain", availabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", faultDomain),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test_peers_2"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_id"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placements.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "placements.0.availability_domain", availabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "placements.0.fault_domain", faultDomain),
				resource.TestCheckResourceAttr(resourceName, "placements.1.availability_domain", availabilityDomain2),
				resource.TestCheckResourceAttr(resourceName, "placements.1.fault_domain", faultDomain2),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource has been recreated meanwhile it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify remove local peer
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment_local_peers", acctest.Optional,
					acctest.Update, acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"placements": []acctest.RepresentationGroup{
							{RepType: acctest.Optional, Group: goldenGateDeploymentPeer1},
						},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "availability_domain", availabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", faultDomain),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test_peers_2"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_id"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placements.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "placements.0.availability_domain", availabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "placements.0.fault_domain", faultDomain),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource has been recreated meanwhile it was supposed to be updated")
					}
					return err
				},
			),
		},

		// add local peer to perform switchover
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment_local_peers", acctest.Optional,
					acctest.Update, acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"placements": []acctest.RepresentationGroup{
							{RepType: acctest.Optional, Group: goldenGateDeploymentPeer1},
							{RepType: acctest.Optional, Group: goldenGateDeploymentPeer2},
						},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "availability_domain", availabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", faultDomain),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test_peers_2"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_id"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placements.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "placements.0.availability_domain", availabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "placements.0.fault_domain", faultDomain),
				resource.TestCheckResourceAttr(resourceName, "placements.1.availability_domain", availabilityDomain2),
				resource.TestCheckResourceAttr(resourceName, "placements.1.fault_domain", faultDomain2),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource has been recreated meanwhile it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify switchover
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment_local_peers", acctest.Optional,
					acctest.Update, acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"availability_domain": acctest.Representation{RepType: acctest.Optional, Update: `${var.availability_domain_2}`},
						"fault_domain":        acctest.Representation{RepType: acctest.Optional, Update: `${var.fault_domain_2}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "availability_domain", availabilityDomain2),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", faultDomain2),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test_peers_2"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_id"),
				resource.TestCheckNoResourceAttr(resourceName, "load_balancer_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placements.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.1.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.1.fault_domain"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource has been recreated meanwhile it was supposed to be updated")
					}
					return err
				},
			),
		},

		{
			Config: config,
		},
	})
}
