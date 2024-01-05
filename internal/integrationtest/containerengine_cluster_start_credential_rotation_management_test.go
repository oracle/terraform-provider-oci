// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineClusterRepresentationForCredentialRotation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-2]}`},
		"name":                        acctest.Representation{RepType: acctest.Required, Create: `name`},
		"vcn_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"cluster_pod_network_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterClusterPodNetworkOptionsRepresentationForCredentialRotation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"endpoint_config":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterEndpointConfigRepresentationForCredentialRotation},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image_policy_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterImagePolicyConfigRepresentationForCredentialRotation},
		"kms_key_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"options":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsRepresentationForCredentialRotation},
	}
	clusterClusterPodNetworkOptionsRepresentationForCredentialRotation = map[string]interface{}{
		"cni_type": acctest.Representation{RepType: acctest.Required, Create: `FLANNEL_OVERLAY`},
	}
	ContainerengineClusterEndpointConfigRepresentationForCredentialRotation = map[string]interface{}{
		"is_public_ip_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"nsg_ids":              acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"subnet_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	ContainerengineClusterImagePolicyConfigRepresentationForCredentialRotation = map[string]interface{}{
		"is_policy_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"key_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterImagePolicyConfigKeyDetailsRepresentationForCredentialRotation},
	}
	ContainerengineClusterOptionsRepresentationForCredentialRotation = map[string]interface{}{
		"add_ons":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsAddOnsRepresentationForCredentialRotation},
		"admission_controller_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsAdmissionControllerOptionsRepresentationForCredentialRotation},
		"kubernetes_network_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsKubernetesNetworkConfigRepresentationForCredentialRotation},
		"persistent_volume_config":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsPersistentVolumeConfigRepresentationForCredentialRotation},
		"service_lb_config":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsServiceLbConfigRepresentationForCredentialRotation},
		"service_lb_subnet_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_subnet.clusterSubnet_1.id}`, `${oci_core_subnet.clusterSubnet_2.id}`}},
	}
	ContainerengineClusterImagePolicyConfigKeyDetailsRepresentationForCredentialRotation = map[string]interface{}{
		"kms_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
	}
	ContainerengineClusterOptionsAddOnsRepresentationForCredentialRotation = map[string]interface{}{
		"is_kubernetes_dashboard_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_tiller_enabled":               acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	ContainerengineClusterOptionsAdmissionControllerOptionsRepresentationForCredentialRotation = map[string]interface{}{
		"is_pod_security_policy_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	ContainerengineClusterOptionsKubernetesNetworkConfigRepresentationForCredentialRotation = map[string]interface{}{
		"pods_cidr":     acctest.Representation{RepType: acctest.Optional, Create: `10.1.0.0/16`},
		"services_cidr": acctest.Representation{RepType: acctest.Optional, Create: `10.2.0.0/16`},
	}
	ContainerengineClusterOptionsPersistentVolumeConfigRepresentationForCredentialRotation = map[string]interface{}{
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
	}
	ContainerengineClusterOptionsServiceLbConfigRepresentationForCredentialRotation = map[string]interface{}{
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
	}

	ContainerengineStartCredentialRotationRepresentationForStartCredentialRotationForStartCredentialRotation = map[string]interface{}{
		"start_credential_rotation": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineStartCredentialRotationConfigRepresentationForStartCredentialRotationForCredentialRotation},
	}

	ContainerengineStartCredentialRotationConfigRepresentationForStartCredentialRotationForCredentialRotation = map[string]interface{}{
		"auto_completion_delay_duration":  acctest.Representation{RepType: acctest.Optional, Create: `NONE`, Update: `P5D`},
		"start_credential_rotation_state": acctest.Representation{RepType: acctest.Optional, Create: `NONE`, Update: `NEW_CREDENTIALS_ISSUED`},
	}

	ContainerengineClusterStartCredentialRotationManagementRepresentation = map[string]interface{}{
		"auto_completion_delay_duration": acctest.Representation{RepType: acctest.Required, Create: `P5D`},
		"cluster_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterStartCredentialRotationManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterStartCredentialRotationManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	singularDatasourceName := "data.oci_containerengine_cluster_credential_rotation_status.test_cluster_credential_rotation_status"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_start_credential_rotation_management", "test_cluster_start_credential_rotation_management", acctest.Required, acctest.Create, ContainerengineClusterStartCredentialRotationManagementRepresentation), "containerengine", "clusterStartCredentialRotationManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create cluster
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentationForCredentialRotation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.time_credential_expiration"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// start rotation
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentationForCredentialRotation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_start_credential_rotation_management", "test_cluster_start_credential_rotation_management", acctest.Required, acctest.Create, ContainerengineClusterStartCredentialRotationManagementRepresentation),
		},

		// verify rotation status
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentationForCredentialRotation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_start_credential_rotation_management", "test_cluster_start_credential_rotation_management", acctest.Required, acctest.Create, ContainerengineClusterStartCredentialRotationManagementRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_credential_rotation_status", "test_cluster_credential_rotation_status",
					acctest.Optional, acctest.Create, ContainerengineClusterCredentialRotationStatusSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "status", "WAITING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status_details", "NEW_CREDENTIALS_ISSUED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_auto_completion_scheduled"),
			),
		},
	})
}
