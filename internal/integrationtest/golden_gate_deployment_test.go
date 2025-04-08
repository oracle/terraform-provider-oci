// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID                   = "compartment_id"
		COMPARTMENT_ID_FOR_MOVE          = "compartment_id_for_move"
		TEST_SUBNET_ID                   = "test_subnet_id"
		LOAD_BALANCER_SUBNET_ID          = "load_balancer_subnet_id"
		CERTIFICATE                      = "certificate"
		KEY                              = "key"
		BASE_OGG_VERSION                 = "base_ogg_version"
		UPGRADED_OGG_VERSION             = "upgraded_ogg_version"
		PASSWORD                         = "password"
		NEW_PASSWORD                     = "new_password"
		IDENTITY_DOMAIN_ID               = "identity_domain_id"
		PASSWORD_SECRET_ID               = "password_secret_id"
		PASSWORD_SECRET_ID_2             = "password_secret_id_2"
		GROUP_ID                         = "group_id"
		OBJECTSTORAGE_BUCKET_NAME        = "objectstorage_bucket_name"
		OBJECTSTORAGE_NAMESPACE          = "objectstorage_namespace"
		OBJECTSTORAGE_UPDATE_BUCKET_NAME = "objectstorage_update_bucket_name"
		AVAILABILITY_DOMAIN              = "availability_domain"
		FAULT_DOMAIN                     = "fault_domain"
	)

	var (
		resourceName           = "oci_golden_gate_deployment.depl_test_ggs_deployment"
		datasourceName         = "data.oci_golden_gate_deployments.depl_test_ggs_deployments"
		singularDatasourceName = "data.oci_golden_gate_deployment.depl_test_ggs_deployment"

		GoldenGateDeploymentDataSourceFilterRepresentation = map[string]interface{}{
			"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
			"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_golden_gate_deployment.test_deployment.id}`}},
		}
		GoldenGateGoldenGateDeploymentDataSourceRepresentation = map[string]interface{}{
			"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"assignable_connection_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_connection.test_connection.id}`},
			"assigned_connection_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_connection.test_connection.id}`},
			"display_name":              acctest.Representation{RepType: acctest.Required, Create: `Terraform_integration_test`, Update: `Terraform_integration_test2`},
			"fqdn":                      acctest.Representation{RepType: acctest.Required, Update: `fqdn1.oggdevops.us`},
			"lifecycle_sub_state":       acctest.Representation{RepType: acctest.Required, Create: `RECOVERING`},
			"state":                     acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`},
			"supported_connection_type": acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
			"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: GoldenGateDeploymentDataSourceFilterRepresentation}}

		compartmentId                = utils.GetEnvSettingWithBlankDefault(COMPARTMENT_ID)
		compartmentIdForMove         = utils.GetEnvSettingWithBlankDefault(COMPARTMENT_ID_FOR_MOVE)
		subnetId                     = utils.GetEnvSettingWithBlankDefault(TEST_SUBNET_ID)
		identityDomainId             = utils.GetEnvSettingWithBlankDefault(IDENTITY_DOMAIN_ID)
		passwordSecretId             = utils.GetEnvSettingWithBlankDefault(PASSWORD_SECRET_ID)
		passwordSecretId2            = utils.GetEnvSettingWithBlankDefault(PASSWORD_SECRET_ID_2)
		baseOggVersion               = utils.GetEnvSettingWithBlankDefault(BASE_OGG_VERSION)
		upgradedOggVersion           = utils.GetEnvSettingWithBlankDefault(UPGRADED_OGG_VERSION)
		groupId                      = utils.GetEnvSettingWithBlankDefault(GROUP_ID)
		timeBackupScheduledForCreate = time.Now().UTC().Add(time.Hour * 3).Truncate(time.Millisecond).Format(time.RFC3339Nano)
		timeBackupScheduledForUpdate = time.Now().UTC().Add(time.Hour * 6).Truncate(time.Millisecond).Format(time.RFC3339Nano)

		resId  string
		resId2 string
	)

	var (
		GoldenGateDeploymentResourceDependencies = ""

		ignoreDefinedTagsChangesRepresentation = map[string]interface{}{
			"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
		}

		goldenGateDeploymentOggDataRepresentation = map[string]interface{}{
			"admin_password":  acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
			"admin_username":  acctest.Representation{RepType: acctest.Required, Create: `adminUsername`, Update: `adminUsername2`},
			"deployment_name": acctest.Representation{RepType: acctest.Required, Create: `depl_test_ggs_deployment_name`},
			"certificate":     acctest.Representation{RepType: acctest.Optional, Update: `${var.certificate}`},
			"key":             acctest.Representation{RepType: acctest.Optional, Update: `${var.key}`},
		}

		groupToRolesMappingRepresentation = map[string]interface{}{
			"security_group_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.group_id}`},
			"administrator_group_id": acctest.Representation{RepType: acctest.Optional, Update: `${var.group_id}`},
			"operator_group_id":      acctest.Representation{RepType: acctest.Optional, Update: `${var.group_id}`},
			"user_group_id":          acctest.Representation{RepType: acctest.Optional, Update: `${var.group_id}`},
		}

		deploymentBackupScheduleRepresentation = map[string]interface{}{
			"bucket":                     acctest.Representation{RepType: acctest.Required, Create: `${var.objectstorage_bucket_name}`, Update: `${var.objectstorage_update_bucket_name}`},
			"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"frequency_backup_scheduled": acctest.Representation{RepType: acctest.Required, Create: `DAILY`, Update: `WEEKLY`},
			"is_metadata_only":           acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
			"namespace":                  acctest.Representation{RepType: acctest.Required, Create: `${var.objectstorage_namespace}`, Update: `${var.objectstorage_namespace}`},
			"time_backup_scheduled":      acctest.Representation{RepType: acctest.Required, Create: timeBackupScheduledForCreate, Update: timeBackupScheduledForUpdate},
		}

		goldenGateDeploymentOggDataWithGroupRoleMappingRepresentation = map[string]interface{}{
			"admin_password":         acctest.Representation{RepType: acctest.Required, Create: `${var.password}`},
			"admin_username":         acctest.Representation{RepType: acctest.Required, Create: `adminUsername`},
			"deployment_name":        acctest.Representation{RepType: acctest.Required, Create: `depl_test_ggs_deployment_name`},
			"group_to_roles_mapping": acctest.RepresentationGroup{RepType: acctest.Required, Group: groupToRolesMappingRepresentation},
		}

		deploymentMaintenanceConfigurationRepresentation = map[string]interface{}{
			"bundle_release_upgrade_period_in_days":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
			"interim_release_upgrade_period_in_days":  acctest.Representation{RepType: acctest.Optional, Create: `5`, Update: `6`},
			"is_interim_release_auto_upgrade_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
			"major_release_upgrade_period_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `20`, Update: `21`},
			"security_patch_upgrade_period_in_days":   acctest.Representation{RepType: acctest.Optional, Create: `7`, Update: `8`},
		}

		oggDataForUpgradeRepresentation = acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentOggDataRepresentation, map[string]interface{}{
			"ogg_version": acctest.Representation{RepType: acctest.Optional, Create: `${var.base_ogg_version}`, Update: `${var.upgraded_ogg_version}`},
		})

		oggDataRepresentationForIam = map[string]interface{}{
			"deployment_name":    acctest.Representation{RepType: acctest.Required, Create: `IamDeployment`},
			"credential_store":   acctest.Representation{RepType: acctest.Required, Create: `IAM`},
			"identity_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.identity_domain_id}`},
			"ogg_version":        acctest.Representation{RepType: acctest.Required, Create: `${var.base_ogg_version}`},
		}

		oggDataRepresentationForGoldenGate = map[string]interface{}{
			"deployment_name":    acctest.Representation{RepType: acctest.Required, Create: `GoldengateDeployment`},
			"credential_store":   acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
			"admin_username":     acctest.Representation{RepType: acctest.Required, Create: `adminUsername`, Update: `adminUsername2`},
			"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.password_secret_id}`, Update: `${var.password_secret_id_2}`},
			"certificate":        acctest.Representation{RepType: acctest.Optional, Update: `${var.certificate}`},
			"key":                acctest.Representation{RepType: acctest.Optional, Update: `${var.key}`},
		}

		deploymentMaintenanceWindowRepresentation = map[string]interface{}{
			"day":        acctest.Representation{RepType: acctest.Required, Create: `MONDAY`, Update: `TUESDAY`},
			"start_hour": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		}

		goldenGateDeploymentRepresentation = map[string]interface{}{
			"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"cpu_core_count":            acctest.Representation{RepType: acctest.Optional, Create: `1`},
			"deployment_type":           acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_ORACLE`},
			"display_name":              acctest.Representation{RepType: acctest.Required, Create: `Terraform_integration_test`, Update: `Terraform_integration_test2`},
			"is_auto_scaling_enabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.test_subnet_id}`},
			"license_model":             acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
			"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
			"fqdn":                      acctest.Representation{RepType: acctest.Optional, Update: `fqdn1.oggdevops.us`},
			"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
			"is_public":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
			"ogg_data":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDeploymentOggDataRepresentation},
			"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
			"maintenance_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentMaintenanceConfigurationRepresentation},
			"maintenance_window":        acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentMaintenanceWindowRepresentation},
			"backup_schedule":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentBackupScheduleRepresentation},
			"availability_domain":       acctest.Representation{RepType: acctest.Optional, Create: `${var.availability_domain}`},
			"fault_domain":              acctest.Representation{RepType: acctest.Optional, Create: `${var.fault_domain}`},
			"placements":                []acctest.RepresentationGroup{}, // start with empty peer list
			"source_deployment_id":      acctest.Representation{RepType: acctest.Optional, Create: nil},
		}

		deploymentLocksRepresentation = map[string]interface{}{
			"type":    acctest.Representation{RepType: acctest.Required, Create: `FULL`},
			"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
		}

		ignoreDefinedTagsAndLocks = map[string]interface{}{
			"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`locks`, `defined_tags`, `is_lock_override`}},
		}

		lockedDeploymentRepresentation = acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
			"locks":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentLocksRepresentation},
			"is_lock_override": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
			"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsAndLocks},
		})

		goldenGateDeploymentSingularDataSourceRepresentation = map[string]interface{}{
			"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.depl_test_ggs_deployment.id}`},
		}

		DeploymentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update, goldenGateDeploymentRepresentation)
	)

	testDeploymentId := utils.GetEnvSettingWithBlankDefault("deployment_ocid")
	testDeploymentIdVariableStr := fmt.Sprintf("variable \"test_deployment_id\" { default = \"%s\" }\n", testDeploymentId)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(COMPARTMENT_ID_FOR_MOVE, t) +
		makeVariableStr(TEST_SUBNET_ID, t) +
		makeVariableStr(LOAD_BALANCER_SUBNET_ID, t) +
		makeVariableStr(CERTIFICATE, t) +
		makeVariableStr(KEY, t) +
		makeVariableStr(BASE_OGG_VERSION, t) +
		makeVariableStr(UPGRADED_OGG_VERSION, t) +
		makeVariableStr(PASSWORD, t) +
		makeVariableStr(NEW_PASSWORD, t) +
		makeVariableStr(PASSWORD_SECRET_ID, t) +
		makeVariableStr(PASSWORD_SECRET_ID_2, t) +
		makeVariableStr(GROUP_ID, t) +
		makeVariableStr(OBJECTSTORAGE_BUCKET_NAME, t) +
		makeVariableStr(OBJECTSTORAGE_UPDATE_BUCKET_NAME, t) +
		makeVariableStr(OBJECTSTORAGE_NAMESPACE, t) +
		makeVariableStr(AVAILABILITY_DOMAIN, t) +
		makeVariableStr(FAULT_DOMAIN, t) +
		GoldenGateDeploymentResourceDependencies

	if identityDomainId != "" {
		config = config + makeVariableStr(IDENTITY_DOMAIN_ID, t)
	}

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+testDeploymentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create, goldenGateDeploymentRepresentation), "goldengate", "deployment", t)

	var steps = []resource.TestStep{
		//		verify Create
		{
			Config: config + testDeploymentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
		// check groupToRolesMapping attribute set
		{
			Config: config + testDeploymentIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create,
				acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
					"deployment_type": acctest.Representation{RepType: acctest.Required, Create: `OGG`},
					"ogg_data":        acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDeploymentOggDataWithGroupRoleMappingRepresentation},
				})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "OGG"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.group_to_roles_mapping.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.group_to_roles_mapping.0.security_group_id", groupId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//delete before next Create
		{
			Config: config,
		},

		{
			Config: config + testDeploymentIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create,
				acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
					"deployment_type": acctest.Representation{RepType: acctest.Required, Create: `GGSA`},
				})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "GGSA"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// optional step: verify Create with IAM credential store and identity domain id
		{
			PreConfig: func() {
				fmt.Println("This step will run only if TF_VAR_identity_domain_id env variable is set.")
			},
			SkipFunc: func() (bool, error) {
				if identityDomainId == "" {
					return true, nil
				}
				return false, nil
			},
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create,
				acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
					"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataRepresentationForIam},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "IAM"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.identity_domain_id", identityDomainId),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.deployment_name", "IamDeployment"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Create with GOLDENGATE credential store and password secret id
		{
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create,
				acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
					"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataRepresentationForGoldenGate},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.password_secret_id", passwordSecretId),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.deployment_name", "GoldengateDeployment"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify parameter updates for the deployment including password secret id
		{
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update,
				acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
					"ogg_data":                acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataRepresentationForGoldenGate},
					"is_public":               acctest.Representation{RepType: acctest.Optional, Update: `true`},
					"load_balancer_subnet_id": acctest.Representation{RepType: acctest.Optional, Update: `${var.load_balancer_subnet_id}`}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test2"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername2"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.password_secret_id", passwordSecretId2),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.deployment_name", "GoldengateDeployment"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource has been recreated meanwhile it was supposed to be updated")
					}
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config,
		},

		// verify Create with optionals
		{
			Config: config + testDeploymentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create, goldenGateDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_schedule.0.bucket"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.frequency_backup_scheduled", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.is_metadata_only", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_schedule.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.time_backup_scheduled", timeBackupScheduledForCreate),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "fault_domain"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.bundle_release_upgrade_period_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.interim_release_upgrade_period_in_days", "5"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.is_interim_release_auto_upgrade_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.major_release_upgrade_period_in_days", "20"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.security_patch_upgrade_period_in_days", "7"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.day", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.start_hour", "10"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "placements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					time.Sleep(1 * time.Minute)
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
			Config: config + testDeploymentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_move}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_schedule.0.bucket"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.frequency_backup_scheduled", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.is_metadata_only", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_schedule.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.time_backup_scheduled", timeBackupScheduledForCreate),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "fault_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdForMove),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.bundle_release_upgrade_period_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.interim_release_upgrade_period_in_days", "5"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.is_interim_release_auto_upgrade_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.major_release_upgrade_period_in_days", "20"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.security_patch_upgrade_period_in_days", "7"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.day", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.start_hour", "10"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "placements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + testDeploymentIdVariableStr + DeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_schedule.0.bucket"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.frequency_backup_scheduled", "WEEKLY"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.is_metadata_only", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_schedule.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "backup_schedule.0.time_backup_scheduled", timeBackupScheduledForUpdate),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "fault_domain"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Terraform_integration_test2"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.bundle_release_upgrade_period_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.interim_release_upgrade_period_in_days", "6"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.is_interim_release_auto_upgrade_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.major_release_upgrade_period_in_days", "21"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.security_patch_upgrade_period_in_days", "8"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.day", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.start_hour", "11"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername2"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.certificate"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "placements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placements.0.fault_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + DeploymentResourceConfig + testDeploymentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"fqdn":         acctest.Representation{RepType: acctest.Required, Create: `fqdn100.oggdevops.us`},
						"display_name": acctest.Representation{RepType: acctest.Required, Create: `Terraform_integration_test - DataSource test`},
						"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentOggDataRepresentation, map[string]interface{}{
							"certificate": acctest.Representation{RepType: acctest.Required, Create: `${var.certificate}`},
							"key":         acctest.Representation{RepType: acctest.Required, Create: `${var.key}`},
						})},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(GoldenGateConnectionRepresentation, map[string]interface{}{
					"host": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.127`, Update: `10.0.0.128`},
					"port": acctest.Representation{RepType: acctest.Required, Create: `12`, Update: `13`},
				})) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployments", "depl_test_ggs_deployments", acctest.Required, acctest.Update, GoldenGateGoldenGateDeploymentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "assignable_connection_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "assigned_connection_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Terraform_integration_test2"),
				resource.TestCheckResourceAttrSet(datasourceName, "fqdn"),
				resource.TestCheckResourceAttr(datasourceName, "lifecycle_sub_state", "RECOVERING"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "supported_connection_type", "GOLDENGATE"),

				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.0.items.#", "0"),
			),
		},

		// verify singular datasource
		{
			Config: config + DeploymentResourceConfig + testDeploymentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_schedule.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_schedule.0.bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_schedule.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_schedule.0.frequency_backup_scheduled", "WEEKLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_schedule.0.is_metadata_only", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_schedule.0.namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_schedule.0.time_backup_scheduled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fault_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Terraform_integration_test2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_healthy"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_latest_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_public"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_storage_utilization_limit_exceeded"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.bundle_release_upgrade_period_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.interim_release_upgrade_period_in_days", "6"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.is_interim_release_auto_upgrade_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.major_release_upgrade_period_in_days", "21"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_configuration.0.security_patch_upgrade_period_in_days", "8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.day", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.start_hour", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ogg_data.0.credential_store", "GOLDENGATE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ogg_data.0.admin_username", "adminUsername2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placements.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placements.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placements.0.fault_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ogg_data.0.certificate"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ogg_data.0.ogg_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_utilization_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_next_backup_scheduled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DeploymentResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"ogg_data.0.admin_password",
				"ogg_data.0.key",
			},
			ResourceName: resourceName,
		},
		{
			Config: config,
		},
		/* Start/stop/upgrade test*/
		// 0. create a new and locked deployment and stop it right after the creation
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(lockedDeploymentRepresentation, map[string]interface{}{
						"state":    acctest.Representation{RepType: acctest.Optional, Create: string(oci_golden_gate.LifecycleStateInactive)},
						"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataForUpgradeRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_golden_gate.LifecycleStateInactive)),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.ogg_version", baseOggVersion),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// 1. start the locked deployment and upgrade it at the same time
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(lockedDeploymentRepresentation, map[string]interface{}{
						"state":    acctest.Representation{RepType: acctest.Optional, Create: string(oci_golden_gate.LifecycleStateActive)},
						"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataForUpgradeRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_golden_gate.LifecycleStateActive)),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.ogg_version", upgradedOggVersion),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be upgraded.")
					}
					return err
				},
			),
		},
		// 2. step clear
		{
			Config: config,
		},
		// 3. create a new deployment on an older version
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataForUpgradeRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_golden_gate.LifecycleStateActive)),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.ogg_version", baseOggVersion),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// 4. upgrade deployment based on var.upgraded_ogg_version and stop it
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataForUpgradeRepresentation},
						"state":    acctest.Representation{RepType: acctest.Optional, Create: string(oci_golden_gate.LifecycleStateInactive)},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.ogg_version", upgradedOggVersion),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_golden_gate.LifecycleStateInactive)),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be upgraded.")
					}
					return err
				},
			),
		},
		// 5. No upgraded is required, no start/stop is required, nothing should happen here.
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataForUpgradeRepresentation},
						"state":    acctest.Representation{RepType: acctest.Optional, Create: string(oci_golden_gate.LifecycleStateInactive)},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.ogg_version", upgradedOggVersion),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_golden_gate.LifecycleStateInactive)),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be upgraded.")
					}
					return err
				},
			),
		},
		{
			Config: config,
		},
	}
	acctest.ResourceTest(t, testAccCheckGoldenGateDeploymentDestroy, steps)
}

func testAccCheckGoldenGateDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_deployment" {
			noResourceFound = false
			request := oci_golden_gate.GetDeploymentRequest{}

			tmp := rs.Primary.ID
			request.DeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			response, err := client.GetDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_golden_gate.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GoldenGateDeployment") {
		resource.AddTestSweepers("GoldenGateDeployment", &resource.Sweeper{
			Name:         "GoldenGateDeployment",
			Dependencies: acctest.DependencyGraph["deployment"],
			F:            sweepGoldenGateDeploymentResource,
		})
	}
}

func sweepGoldenGateDeploymentResource(compartment string) error {
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()
	deploymentIds, err := getGoldenGateDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, deploymentId := range deploymentIds {
		if ok := acctest.SweeperDefaultResourceId[deploymentId]; !ok {
			deleteDeploymentRequest := oci_golden_gate.DeleteDeploymentRequest{}

			deleteDeploymentRequest.DeploymentId = &deploymentId

			var overrideLock = true
			deleteDeploymentRequest.IsLockOverride = &overrideLock

			deleteDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeleteDeployment(context.Background(), deleteDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting Deployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", deploymentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &deploymentId, goldenGateDeploymentSweepWaitCondition, time.Duration(3*time.Minute),
				goldenGateDeploymentSweepResponseFetchOperation, "golden_gate", true)
		}
	}
	return nil
}

func getGoldenGateDeploymentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()

	listDeploymentsRequest := oci_golden_gate.ListDeploymentsRequest{}
	listDeploymentsRequest.CompartmentId = &compartmentId
	listDeploymentsRequest.LifecycleState = oci_golden_gate.ListDeploymentsLifecycleStateActive
	listDeploymentsResponse, err := goldenGateClient.ListDeployments(context.Background(), listDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Deployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deployment := range listDeploymentsResponse.Items {
		id := *deployment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DeploymentId", id)
	}
	return resourceIds, nil
}

func goldenGateDeploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deploymentResponse, ok := response.Response.(oci_golden_gate.GetDeploymentResponse); ok {
		return deploymentResponse.LifecycleState != oci_golden_gate.LifecycleStateDeleted
	}
	return false
}

func goldenGateDeploymentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GoldenGateClient().GetDeployment(context.Background(), oci_golden_gate.GetDeploymentRequest{
		DeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
