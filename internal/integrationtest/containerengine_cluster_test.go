// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/containerengine"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineNativeVCNMigrationClusterRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-1]}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `test_cluster`},
		"vcn_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"endpoint_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineNativeVCNMIgrationClusterEndpointConfigRepresentation},
	}

	ContainerengineNativeVCNMIgrationClusterEndpointConfigRepresentation = map[string]interface{}{
		"nsg_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`, Update: `${oci_core_subnet.test_subnet2.id}`},
	}

	ContainerengineClusterRequiredOnlyResource = ContainerengineClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterRepresentation)

	ContainerengineClusterResourceConfig = ContainerengineClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Update, ContainerengineClusterRepresentation)

	ContainerengineClusterSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}

	ContainerengineContainerengineClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: []string{`CREATING`, `ACTIVE`, `FAILED`, `DELETING`, `DELETED`, `UPDATING`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterDataSourceFilterRepresentation}}
	ContainerengineClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_cluster.test_cluster.id}`}},
	}
	ContainerengineClusterRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-2]}`, Update: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-1]}`},
		"name":                        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"vcn_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"cluster_pod_network_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterClusterPodNetworkOptionsRepresentation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"endpoint_config":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterEndpointConfigRepresentation},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image_policy_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterImagePolicyConfigRepresentation},
		"kms_key_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"type":                        acctest.Representation{RepType: acctest.Optional, Create: `ENHANCED_CLUSTER`, Update: `ENHANCED_CLUSTER`},
		"options":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsRepresentation},
	}
	ContainerengineClusterClusterPodNetworkOptionsRepresentation = map[string]interface{}{
		"cni_type": acctest.Representation{RepType: acctest.Required, Create: `OCI_VCN_IP_NATIVE`},
	}
	ContainerengineClusterEndpointConfigRepresentation = map[string]interface{}{
		"nsg_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	ContainerengineClusterImagePolicyConfigRepresentation = map[string]interface{}{
		"is_policy_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterImagePolicyConfigKeyDetailsRepresentation},
	}
	ContainerengineClusterOptionsRepresentation = map[string]interface{}{
		"add_ons":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsAddOnsRepresentation},
		"admission_controller_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsAdmissionControllerOptionsRepresentation},
		"ip_families":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`IPv4`}},
		"kubernetes_network_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsKubernetesNetworkConfigRepresentation},
		"open_id_connect_token_authentication_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsOpenIdConnectTokenAuthenticationConfigRepresentation},
		"open_id_connect_discovery":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsOpenIdConnectDiscoveryRepresentation},
		"persistent_volume_config":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsPersistentVolumeConfigRepresentation},
		"service_lb_config":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsServiceLbConfigRepresentation},
		"service_lb_subnet_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_subnet.clusterSubnet_1.id}`, `${oci_core_subnet.clusterSubnet_2.id}`}},
	}
	ContainerengineClusterImagePolicyConfigKeyDetailsRepresentation = map[string]interface{}{
		"kms_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
	}
	ContainerengineClusterOptionsAddOnsRepresentation = map[string]interface{}{
		"is_kubernetes_dashboard_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_tiller_enabled":               acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	ContainerengineClusterOptionsAdmissionControllerOptionsRepresentation = map[string]interface{}{
		"is_pod_security_policy_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	ContainerengineClusterOptionsKubernetesNetworkConfigRepresentation = map[string]interface{}{
		"pods_cidr":     acctest.Representation{RepType: acctest.Optional, Create: `10.1.0.0/16`},
		"services_cidr": acctest.Representation{RepType: acctest.Optional, Create: `10.2.0.0/16`},
	}
	ContainerengineClusterOptionsOpenIdConnectTokenAuthenticationConfigRepresentation = map[string]interface{}{
		"is_open_id_connect_auth_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"ca_certificate":                  acctest.Representation{RepType: acctest.Optional, Create: `LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUY5RENDQTl5Z0F3SUJBZ0lVYjZUaGdGNThwZVR0QkQ3Q2dyWlVNbDRXMWZNd0RRWUpLb1pJaHZjTkFRRUwKQlFBd2dZQXhDekFKQmdOVkJBWVRBbGhZTVJJd0VBWURWUVFJREFsVGRHRjBaVTVoYldVeEVUQVBCZ05WQkFjTQpDRU5wZEhsT1lXMWxNUlF3RWdZRFZRUUtEQXREYjIxd1lXNTVUbUZ0WlRFYk1Ca0dBMVVFQ3d3U1EyOXRjR0Z1CmVWTmxZM1JwYjI1T1lXMWxNUmN3RlFZRFZRUUREQTR4TlRBdU1UTTJMak0zTGpJME16QWVGdzB5TkRBME1ETXcKTWpVMU1qQmFGdzB6TkRBME1ERXdNalUxTWpCYU1JR0FNUXN3Q1FZRFZRUUdFd0pZV0RFU01CQUdBMVVFQ0F3SgpVM1JoZEdWT1lXMWxNUkV3RHdZRFZRUUhEQWhEYVhSNVRtRnRaVEVVTUJJR0ExVUVDZ3dMUTI5dGNHRnVlVTVoCmJXVXhHekFaQmdOVkJBc01Fa052YlhCaGJubFRaV04wYVc5dVRtRnRaVEVYTUJVR0ExVUVBd3dPTVRVd0xqRXoKTmk0ek55NHlORE13Z2dJaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQ0R3QXdnZ0lLQW9JQ0FRRFBmWGFUaXE4ZAp3VW1QZjNNWUdZUnRYVlliUzhRblUzeGNUdnpqeExxRTlIZnVPUUgyVFBna0wvbjBQUTVvZnFURnlMRjd3Z1BvCks5Wmx2dG1mUzhNRUlsU1doUFdSWWpuNDRYVXhOVjBhTldUWi80bWFYcUlXOUM1aC9xaEhiVU1IbTAwNzZkQjMKby9GaTBEZnpwN3JDTGhZeTJUaG5oc1BOWXYzcFljVlNDbVFZVVpUNlh6eVR5Ym4vY3IvY2tTYXRWTkZKNEQ3UwpLK2xxdEtzNzF3bkMzTjhQd2xZemFyOWFaMnNlSmNrSXZRWWtKZ3phcktZK0hYTkg3SVZKa0h2N3QyY1NJdGJvClRUVTJHVE5icEJyK05YSFZlaWo0THpsZWdER0dPWkFjYWEwS094YWoxNXNISFlQSVBJYlZ2NXMwVXVodkVyUEIKOFUxRjVhRFU0L0MrbW9Lc3EwTlpEcDZNSkphZ1lBazEvUzduRCtXbEExem5rMzRyYXl0U2FDdDdFK1JTMlR2YQpaRkExQWFNNmlBVXp1eEY3ZWtnejhzZ1lZd0drMnJZTFRzQm9IdHloeUkxM0FwMVBlL3ROSVQxaHlYdmQzVzVICmFtQ09PSnFkYUFkYU1xR2g0V25lemZ6UjZUY2NHNmpaekVqV2N0amEyWXlPVlZaaWlRWWllajA2S1VYazI4U2wKUVlGamNXak53c0s5b1U4OFRFbkJ6U0FPZ29OVXBITkxEeFVzRHhEMDBuQm9aS2U2aExFVmFMckRiSEdzdU5MSgo1TDd3WE1xdUZIQ2RjUUhhTzMxTnp5UllrTGw3UUtiSkRvZjNWNzBQMmRtbXhDdU5ybTkwWFNoZ3o3L2xHK0krCnFiUkRVT21Pcjc1dmxvZEFDODB4SU5tKzZLcEdJcXVGaVFJREFRQUJvMlF3WWpBZEJnTlZIUTRFRmdRVU4zK0gKa3Z5SW9KVFVLSTlFNEFweTFSbEc3bjh3SHdZRFZSMGpCQmd3Rm9BVU4zK0hrdnlJb0pUVUtJOUU0QXB5MVJsRwo3bjh3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFQQmdOVkhSRUVDREFHaHdTV2lDWHpNQTBHQ1NxR1NJYjNEUUVCCkN3VUFBNElDQVFDQjR4SEljOG81MlpFaHoyYVJoNUl0QTdYeUs4bjdROVRqODJiTjZPaFB1ZDg0MDRNbUFldHIKNFAzTUtadWlJNDRIeUVZQ2tSenFNVXVRNS9yd2JJNU1iUStBN1Z3dEpDdjRMZEZjOEsxRnZKcEo0c2lXaWwzMgpaQ0YyS052elB6Vm1NVVc4QkZFSVFMMnIvdWhrelc2V2FvWHVVZTV0NkE3UDJIL1lYU3JvRUdDZGZDOGQ3TWo5CkppN1Nxc3BYTUM2WlNvZmhQUFZMTUtRWk44REh5SXorUlJHZS8zYkw2S2ZMUXpuZDJ4UzJVcVZ3eGIweERjZ04KK0lrNGRCcjdqS1Z1N3cwTjJvcUo5V2ZmVjgxRFV1OWgwZjROeGcxakc1bWlCZHRTU1dTWEdqNHlLRmdvWnJwTApMZEdGRGdiUFh1QjZLWGVBY3N1M21hRjZrOXBsU29WeFU1cC9zdGRVSXMycFphTmdsdzh2UmJxWFM3cXFKT0V2Ck5RSmZPcEszN2JrMjBCS1lsQ2NXUWxGVjFmejdRdUFJdS9hK0NlUXFwUXhGUVk5SmFsYStNWDFxaDRIYWJUNHgKNGNOaktBNC9pNVhkYkdaTUFJQVhxN2tnNVlSc2xyQmQ5ZFNaM0FMUVVocEppZ01TOFhDY1htQUZrcnR2Tm81YgplWXNjM1AyT0hnRjQ3a3Z1N00rUWdDNU43RkJvNkhSdkNLY1dwRG5oSVprd2JsRkRsVW1iejBicHJLRDVqc01wCi93dDl3OUhaWEFsOTRtU1JLeTFJdmFvMHdndEYzNW9Sait1eXpxRUl2RVdkMmpvc21LTmtGWGlXMU5lbGNrUnUKT1FjdFBNN3BJbHZ6ZWh2U3BENTVWN09NbHhLZHArMTQ0cVMrbDB2UStOUGJSMk91TkNZa1lBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=`},
		"client_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_client.test_client.id}`},
		"groups_claim":                    acctest.Representation{RepType: acctest.Optional, Create: `groupsClaim`},
		"groups_prefix":                   acctest.Representation{RepType: acctest.Optional, Create: `groupsPrefix`},
		"issuer_url":                      acctest.Representation{RepType: acctest.Optional, Create: `https://url1.com`},
		"required_claims":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterOptionsOpenIdConnectTokenAuthenticationConfigRequiredClaimsRepresentation},
		"signing_algorithms":              acctest.Representation{RepType: acctest.Optional, Create: []string{`RS256`}},
		"username_claim":                  acctest.Representation{RepType: acctest.Optional, Create: `sub`},
		"username_prefix":                 acctest.Representation{RepType: acctest.Optional, Create: `oidc:`},
	}
	ContainerengineClusterOptionsOpenIdConnectDiscoveryRepresentation = map[string]interface{}{
		"is_open_id_connect_discovery_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ContainerengineClusterOptionsPersistentVolumeConfigRepresentation = map[string]interface{}{
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	ContainerengineClusterOptionsServiceLbConfigRepresentation = map[string]interface{}{
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	ContainerengineClusterOptionsOpenIdConnectTokenAuthenticationConfigRequiredClaimsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}

	ContainerengineClusterResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.22.0/24`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	datasourceName := "data.oci_containerengine_clusters.test_clusters"
	singularDatasourceName := "data.oci_containerengine_cluster.test_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentation), "containerengine", "cluster", t)

	acctest.ResourceTest(t, testAccCheckContainerengineClusterDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_pod_network_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cluster_pod_network_options.0.cni_type", "OCI_VCN_IP_NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.is_policy_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_tiller_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.0.is_pod_security_policy_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "options.0.ip_families.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "options.0.open_id_connect_token_authentication_config.0.ca_certificate"),
				resource.TestCheckResourceAttrSet(resourceName, "options.0.open_id_connect_token_authentication_config.0.client_id"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.groups_claim", "groupsClaim"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.groups_prefix", "groupsPrefix"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.is_open_id_connect_auth_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.issuer_url", "https://url1.com"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.required_claims.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.required_claims.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.required_claims.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.signing_algorithms.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.username_claim", "RS256"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.0.username_prefix", "oidc:"),
				resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.0.is_pod_security_policy_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_discovery.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_discovery.0.is_open_id_connect_discovery_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "options.0.persistent_volume_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.persistent_volume_config.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_config.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_subnet_ids.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "type", "ENHANCED_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Update, ContainerengineClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_pod_network_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cluster_pod_network_options.0.cni_type", "OCI_VCN_IP_NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.is_policy_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image_policy_config.0.key_details.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_tiller_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.0.is_pod_security_policy_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "options.0.ip_families.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_token_authentication_config.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.0.is_pod_security_policy_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_discovery.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.open_id_connect_discovery.0.is_open_id_connect_discovery_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "options.0.persistent_volume_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.persistent_volume_config.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_config.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_subnet_ids.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "type", "ENHANCED_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_clusters", "test_clusters", acctest.Optional, acctest.Update, ContainerengineContainerengineClusterDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Update, ContainerengineClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "6"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.cluster_pod_network_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.cluster_pod_network_options.0.cni_type", "OCI_VCN_IP_NATIVE"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.available_kubernetes_upgrades.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.endpoint_config.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.image_policy_config.0.is_policy_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.image_policy_config.0.key_details.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.kubernetes_version"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.metadata.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.0.is_tiller_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.admission_controller_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.ip_families.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.ca_certificate", "caCertificate2"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.client_id"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.groups_claim", "groupsClaim2"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.groups_prefix", "groupsPrefix2"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.is_open_id_connect_auth_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.issuer_url", "issuerUrl2"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.required_claims.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.required_claims.0.key", "key2"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.required_claims.0.value", "value2"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.signing_algorithms.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.username_claim", "usernameClaim2"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_token_authentication_config.0.username_prefix", "usernamePrefix2"),
				resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.0.is_pod_security_policy_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_discovery.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.open_id_connect_discovery.0.is_open_id_connect_discovery_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.persistent_volume_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.persistent_volume_config.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.service_lb_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.service_lb_config.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.service_lb_subnet_ids.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.type", "ENHANCED_CLUSTER"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.vcn_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "available_kubernetes_upgrades.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_pod_network_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_pod_network_options.0.cni_type", "OCI_VCN_IP_NATIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_config.0.is_public_ip_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_policy_config.0.is_policy_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.add_ons.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.add_ons.0.is_tiller_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.admission_controller_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.ip_families.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.admission_controller_options.0.is_pod_security_policy_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.kubernetes_network_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.persistent_volume_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.persistent_volume_config.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.service_lb_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.service_lb_config.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "options.0.service_lb_subnet_ids.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ENHANCED_CLUSTER"),
			),
		},
		// verify resource import
		{
			Config:                  config + ContainerengineClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckContainerengineClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_cluster" {
			noResourceFound = false
			request := oci_containerengine.GetClusterRequest{}

			tmp := rs.Primary.ID
			request.ClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

			response, err := client.GetCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.ClusterLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ContainerengineCluster") {
		resource.AddTestSweepers("ContainerengineCluster", &resource.Sweeper{
			Name:         "ContainerengineCluster",
			Dependencies: acctest.DependencyGraph["cluster"],
			F:            sweepContainerengineClusterResource,
		})
	}
}

func sweepContainerengineClusterResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	clusterIds, err := getContainerengineClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterId := range clusterIds {
		if ok := acctest.SweeperDefaultResourceId[clusterId]; !ok {
			deleteClusterRequest := oci_containerengine.DeleteClusterRequest{}

			deleteClusterRequest.ClusterId = &clusterId

			deleteClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteCluster(context.Background(), deleteClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting Cluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterId, ContainerengineClusterSweepWaitCondition, time.Duration(3*time.Minute),
				ContainerengineClusterSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getContainerengineClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listClustersRequest := oci_containerengine.ListClustersRequest{}
	listClustersRequest.CompartmentId = &compartmentId
	listClustersRequest.LifecycleState = []containerengine.ClusterLifecycleStateEnum{oci_containerengine.ClusterLifecycleStateActive}
	listClustersResponse, err := containerEngineClient.ListClusters(context.Background(), listClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Cluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cluster := range listClustersResponse.Items {
		id := *cluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterId", id)
	}
	return resourceIds, nil
}

func ContainerengineClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterResponse, ok := response.Response.(oci_containerengine.GetClusterResponse); ok {
		return clusterResponse.LifecycleState != oci_containerengine.ClusterLifecycleStateDeleted
	}
	return false
}

func ContainerengineClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetCluster(context.Background(), oci_containerengine.GetClusterRequest{
		ClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// issue-routing-tag: containerengine/default
func TestContainerengineNativeVcnMigration_test(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNativeVcnMigration_test")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	singularDatasourceName := "data.oci_containerengine_migrate_to_native_vcn_status.test_migrate_to_native_vcn_status"

	var resId, resId2 string

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create V1h Cluster
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineNativeVCNMigrationClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_cluster"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify V1h Cluster migrates to V2
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Update, ContainerengineNativeVCNMigrationClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.is_public_ip_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Update, ContainerengineNativeVCNMigrationClusterRepresentation) + acctest.GenerateDataSourceFromRepresentationMap(
				"oci_containerengine_migrate_to_native_vcn_status", "test_migrate_to_native_vcn_status",
				acctest.Optional, acctest.Create, ContainerengineContainerengineMigrateToNativeVcnStatuSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_decommission_scheduled"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies,
		},
		// create v2 cluster
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineNativeVCNMigrationClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_cluster"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//verify recreation when changing endpointconfig.subnetId
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Update, ContainerengineNativeVCNMigrationClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.is_public_ip_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("resource is supposed to be recreated when endpoint_config.subnet_id updated")
					}
					return err
				},
			),
		},
	})
}
