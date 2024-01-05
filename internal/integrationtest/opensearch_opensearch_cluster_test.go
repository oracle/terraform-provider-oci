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
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"
)

var (
	OpensearchOpensearchClusterRequiredOnlyResource = OpensearchOpensearchClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Required, acctest.Create, OpensearchOpensearchClusterRepresentation)

	OpensearchOpensearchClusterResourceConfig = OpensearchOpensearchClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Optional, acctest.Update, OpensearchOpensearchClusterRepresentation)

	OpensearchOpensearchClusterSingularDataSourceRepresentation = map[string]interface{}{
		"opensearch_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opensearch_opensearch_cluster.test_opensearch_cluster.id}`},
	}

	OpensearchOpensearchClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `tf_provider_cluster_updated`, Update: `tf_provider_cluster_updated`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_opensearch_opensearch_cluster.test_opensearch_cluster.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OpensearchOpensearchClusterDataSourceFilterRepresentation},
	}
	OpensearchOpensearchClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opensearch_opensearch_cluster.test_opensearch_cluster.id}`}},
	}

	OpensearchOpensearchClusterRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_node_count":                    acctest.Representation{RepType: acctest.Required, Create: `1`},
		"data_node_host_memory_gb":           acctest.Representation{RepType: acctest.Required, Create: `20`},
		"data_node_host_ocpu_count":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"data_node_host_type":                acctest.Representation{RepType: acctest.Required, Create: `FLEX`},
		"data_node_storage_gb":               acctest.Representation{RepType: acctest.Required, Create: `50`},
		"display_name":                       acctest.Representation{RepType: acctest.Required, Create: `tf_provider_cluster_updated`, Update: `tf_provider_cluster_updated`},
		"master_node_count":                  acctest.Representation{RepType: acctest.Required, Create: `1`},
		"master_node_host_memory_gb":         acctest.Representation{RepType: acctest.Required, Create: `20`},
		"master_node_host_ocpu_count":        acctest.Representation{RepType: acctest.Required, Create: `1`},
		"master_node_host_type":              acctest.Representation{RepType: acctest.Required, Create: `FLEX`},
		"opendashboard_node_count":           acctest.Representation{RepType: acctest.Required, Create: `1`},
		"opendashboard_node_host_memory_gb":  acctest.Representation{RepType: acctest.Required, Create: `10`},
		"opendashboard_node_host_ocpu_count": acctest.Representation{RepType: acctest.Required, Create: `2`},
		"software_version":                   acctest.Representation{RepType: acctest.Required, Create: `1.2.4`, Update: `1.2.4`},
		"subnet_compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":                          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"system_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sys-namespace.tag-key": "value"}, Update: map[string]string{"sys-namespace.tag-key": "updatedValue"}},
		"security_mode":                      acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`, Update: `ENFORCING`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreOpenSearchSystemTagsChangesRep},
	}

	OpensearchOpensearchClusterRepresentationWithEnforcingSecurityMode = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_node_count":                    acctest.Representation{RepType: acctest.Required, Create: `1`},
		"data_node_host_memory_gb":           acctest.Representation{RepType: acctest.Required, Create: `20`},
		"data_node_host_ocpu_count":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"data_node_host_type":                acctest.Representation{RepType: acctest.Required, Create: `FLEX`},
		"data_node_storage_gb":               acctest.Representation{RepType: acctest.Required, Create: `50`},
		"display_name":                       acctest.Representation{RepType: acctest.Required, Create: `tf_provider_cluster_updated`, Update: `tf_provider_cluster_updated`},
		"master_node_count":                  acctest.Representation{RepType: acctest.Required, Create: `1`},
		"master_node_host_memory_gb":         acctest.Representation{RepType: acctest.Required, Create: `20`},
		"master_node_host_ocpu_count":        acctest.Representation{RepType: acctest.Required, Create: `1`},
		"master_node_host_type":              acctest.Representation{RepType: acctest.Required, Create: `FLEX`},
		"opendashboard_node_count":           acctest.Representation{RepType: acctest.Required, Create: `1`},
		"opendashboard_node_host_memory_gb":  acctest.Representation{RepType: acctest.Required, Create: `10`},
		"opendashboard_node_host_ocpu_count": acctest.Representation{RepType: acctest.Required, Create: `2`},
		"software_version":                   acctest.Representation{RepType: acctest.Required, Create: `1.2.4`, Update: `1.2.4`},
		"subnet_compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":                          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"system_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sys-namespace.tag-key": "value"}, Update: map[string]string{"sys-namespace.tag-key": "updatedValue"}},
		"security_mode":                      acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`, Update: `ENFORCING`},
		"security_master_user_name":          acctest.Representation{RepType: acctest.Optional, Update: `${oci_identity_user.test_user.name}`},
		"security_master_user_password_hash": acctest.Representation{RepType: acctest.Optional, Update: `securityMasterUserPasswordHash2`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreOpenSearchSystemTagsChangesRep},
	}

	OpensearchOpensearchClusterHorizontalResizeRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_node_count":                    acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"data_node_host_memory_gb":           acctest.Representation{RepType: acctest.Required, Create: `20`},
		"data_node_host_ocpu_count":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"data_node_host_type":                acctest.Representation{RepType: acctest.Required, Create: `FLEX`},
		"data_node_storage_gb":               acctest.Representation{RepType: acctest.Required, Create: `50`},
		"display_name":                       acctest.Representation{RepType: acctest.Required, Create: `tf_provider_cluster_updated`, Update: `tf_provider_cluster_updated`},
		"master_node_count":                  acctest.Representation{RepType: acctest.Required, Create: `1`},
		"master_node_host_memory_gb":         acctest.Representation{RepType: acctest.Required, Create: `20`},
		"master_node_host_ocpu_count":        acctest.Representation{RepType: acctest.Required, Create: `1`},
		"master_node_host_type":              acctest.Representation{RepType: acctest.Required, Create: `FLEX`},
		"opendashboard_node_count":           acctest.Representation{RepType: acctest.Required, Create: `1`},
		"opendashboard_node_host_memory_gb":  acctest.Representation{RepType: acctest.Required, Create: `10`},
		"opendashboard_node_host_ocpu_count": acctest.Representation{RepType: acctest.Required, Create: `2`},
		"software_version":                   acctest.Representation{RepType: acctest.Required, Create: `1.2.4`, Update: `1.2.4`},
		"subnet_compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":                          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"system_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sys-namespace.tag-key": "value"}, Update: map[string]string{"sys-namespace.tag-key": "updatedValue"}},
		"security_mode":                      acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreOpenSearchSystemTagsChangesRep},
	}

	OpensearchOpensearchClusterVerticalResizeRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_node_count":                    acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"data_node_host_memory_gb":           acctest.Representation{RepType: acctest.Required, Create: `20`, Update: `40`},
		"data_node_host_ocpu_count":          acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `4`},
		"data_node_host_type":                acctest.Representation{RepType: acctest.Required, Create: `FLEX`},
		"data_node_storage_gb":               acctest.Representation{RepType: acctest.Required, Create: `50`},
		"display_name":                       acctest.Representation{RepType: acctest.Required, Create: `tf_provider_cluster_updated`, Update: `tf_provider_cluster_updated`},
		"master_node_count":                  acctest.Representation{RepType: acctest.Required, Create: `1`},
		"master_node_host_memory_gb":         acctest.Representation{RepType: acctest.Required, Create: `20`},
		"master_node_host_ocpu_count":        acctest.Representation{RepType: acctest.Required, Create: `1`},
		"master_node_host_type":              acctest.Representation{RepType: acctest.Required, Create: `FLEX`},
		"opendashboard_node_count":           acctest.Representation{RepType: acctest.Required, Create: `1`},
		"opendashboard_node_host_memory_gb":  acctest.Representation{RepType: acctest.Required, Create: `10`},
		"opendashboard_node_host_ocpu_count": acctest.Representation{RepType: acctest.Required, Create: `2`},
		"software_version":                   acctest.Representation{RepType: acctest.Required, Create: `1.2.4`, Update: `1.2.4`},
		"subnet_compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":                          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"system_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sys-namespace.tag-key": "value"}, Update: map[string]string{"sys-namespace.tag-key": "updatedValue"}},
		"security_mode":                      acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreOpenSearchSystemTagsChangesRep},
	}

	ignoreOpenSearchSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`}},
	}

	OpensearchOpensearchClusterResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
)

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchOpensearchClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_opensearch_opensearch_cluster.test_opensearch_cluster"
	datasourceName := "data.oci_opensearch_opensearch_clusters.test_opensearch_clusters"
	singularDatasourceName := "data.oci_opensearch_opensearch_cluster.test_opensearch_cluster"

	var resId, resId2, resId3 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpensearchOpensearchClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Optional, acctest.Create, OpensearchOpensearchClusterRepresentation), "opensearch", "opensearchCluster", t)

	acctest.ResourceTest(t, testAccCheckOpensearchOpensearchClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Required, acctest.Create, OpensearchOpensearchClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_type", "FLEX"),
				resource.TestCheckResourceAttr(resourceName, "data_node_storage_gb", "50"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tf_provider_cluster_updated"),
				resource.TestCheckResourceAttr(resourceName, "master_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_ocpu_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_type", "FLEX"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_memory_gb", "10"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "1.2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify horizontal resize
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Optional, acctest.Update, OpensearchOpensearchClusterHorizontalResizeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_node_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_type", "FLEX"),
				resource.TestCheckResourceAttr(resourceName, "data_node_storage_gb", "50"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tf_provider_cluster_updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "master_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_ocpu_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_type", "FLEX"),
				resource.TestCheckResourceAttrSet(resourceName, "opendashboard_fqdn"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_memory_gb", "10"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "opendashboard_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "opensearch_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "opensearch_private_ip"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "1.2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_storage_gb"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_compartment_id"),
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

		// verify vertical resize
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Optional, acctest.Update, OpensearchOpensearchClusterVerticalResizeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_node_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_memory_gb", "40"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_ocpu_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_type", "FLEX"),
				resource.TestCheckResourceAttr(resourceName, "data_node_storage_gb", "50"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tf_provider_cluster_updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "master_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_ocpu_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_type", "FLEX"),
				resource.TestCheckResourceAttrSet(resourceName, "opendashboard_fqdn"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_memory_gb", "10"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "opendashboard_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "opensearch_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "opensearch_private_ip"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "1.2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_storage_gb"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId2 != resId3 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Optional, acctest.Create, OpensearchOpensearchClusterRepresentationWithEnforcingSecurityMode),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_type", "FLEX"),
				resource.TestCheckResourceAttr(resourceName, "data_node_storage_gb", "50"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tf_provider_cluster_updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "master_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_ocpu_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_type", "FLEX"),
				resource.TestCheckResourceAttrSet(resourceName, "opendashboard_fqdn"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_memory_gb", "10"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "opendashboard_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "opensearch_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "opensearch_private_ip"),
				resource.TestCheckResourceAttr(resourceName, "security_mode", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "1.2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_storage_gb"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_compartment_id"),
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
			Config: config + compartmentIdVariableStr + OpensearchOpensearchClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Optional, acctest.Update, OpensearchOpensearchClusterRepresentationWithEnforcingSecurityMode),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "data_node_host_type", "FLEX"),
				resource.TestCheckResourceAttr(resourceName, "data_node_storage_gb", "50"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tf_provider_cluster_updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "master_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_ocpu_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "master_node_host_type", "FLEX"),
				resource.TestCheckResourceAttrSet(resourceName, "opendashboard_fqdn"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_memory_gb", "10"),
				resource.TestCheckResourceAttr(resourceName, "opendashboard_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "opendashboard_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "opensearch_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "opensearch_private_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "security_master_user_name"),
				resource.TestCheckResourceAttr(resourceName, "security_master_user_password_hash", "securityMasterUserPasswordHash2"),
				resource.TestCheckResourceAttr(resourceName, "security_mode", "ENFORCING"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "1.2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_storage_gb"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_compartment_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opensearch_opensearch_clusters", "test_opensearch_clusters", acctest.Optional, acctest.Update, OpensearchOpensearchClusterDataSourceRepresentation) +
				compartmentIdVariableStr + OpensearchOpensearchClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Optional, acctest.Update, OpensearchOpensearchClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "tf_provider_cluster_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "opensearch_cluster_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "opensearch_cluster_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opensearch_opensearch_cluster", "test_opensearch_cluster", acctest.Required, acctest.Create, OpensearchOpensearchClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpensearchOpensearchClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opensearch_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_node_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_node_host_type", "FLEX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_node_storage_gb", "50"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "tf_provider_cluster_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "master_node_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "master_node_host_memory_gb", "20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "master_node_host_ocpu_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "master_node_host_type", "FLEX"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opendashboard_fqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "opendashboard_node_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "opendashboard_node_host_memory_gb", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "opendashboard_node_host_ocpu_count", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opendashboard_private_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opensearch_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opensearch_private_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_master_user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_master_user_password_hash", "securityMasterUserPasswordHash2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_mode", "ENFORCING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "software_version", "1.2.4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_storage_gb"),
			),
		},
		// verify resource import
		{
			Config:                  config + OpensearchOpensearchClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpensearchOpensearchClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OpensearchClusterClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opensearch_opensearch_cluster" {
			noResourceFound = false
			request := oci_opensearch.GetOpensearchClusterRequest{}

			tmp := rs.Primary.ID
			request.OpensearchClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opensearch")

			response, err := client.GetOpensearchCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opensearch.OpensearchClusterLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpensearchOpensearchCluster") {
		resource.AddTestSweepers("OpensearchOpensearchCluster", &resource.Sweeper{
			Name:         "OpensearchOpensearchCluster",
			Dependencies: acctest.DependencyGraph["opensearchCluster"],
			F:            sweepOpensearchOpensearchClusterResource,
		})
	}
}

func sweepOpensearchOpensearchClusterResource(compartment string) error {
	opensearchClusterClient := acctest.GetTestClients(&schema.ResourceData{}).OpensearchClusterClient()
	opensearchClusterIds, err := getOpensearchOpensearchClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, opensearchClusterId := range opensearchClusterIds {
		if ok := acctest.SweeperDefaultResourceId[opensearchClusterId]; !ok {
			deleteOpensearchClusterRequest := oci_opensearch.DeleteOpensearchClusterRequest{}

			deleteOpensearchClusterRequest.OpensearchClusterId = &opensearchClusterId

			deleteOpensearchClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opensearch")
			_, error := opensearchClusterClient.DeleteOpensearchCluster(context.Background(), deleteOpensearchClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting OpensearchCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", opensearchClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &opensearchClusterId, OpensearchOpensearchClusterSweepWaitCondition, time.Duration(3*time.Minute),
				OpensearchOpensearchClusterSweepResponseFetchOperation, "opensearch", true)
		}
	}
	return nil
}

func getOpensearchOpensearchClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OpensearchClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	opensearchClusterClient := acctest.GetTestClients(&schema.ResourceData{}).OpensearchClusterClient()

	listOpensearchClustersRequest := oci_opensearch.ListOpensearchClustersRequest{}
	listOpensearchClustersRequest.CompartmentId = &compartmentId
	listOpensearchClustersRequest.LifecycleState = oci_opensearch.OpensearchClusterLifecycleStateActive
	listOpensearchClustersResponse, err := opensearchClusterClient.ListOpensearchClusters(context.Background(), listOpensearchClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OpensearchCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, opensearchCluster := range listOpensearchClustersResponse.Items {
		id := *opensearchCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OpensearchClusterId", id)
	}
	return resourceIds, nil
}

func OpensearchOpensearchClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if opensearchClusterResponse, ok := response.Response.(oci_opensearch.GetOpensearchClusterResponse); ok {
		return opensearchClusterResponse.LifecycleState != oci_opensearch.OpensearchClusterLifecycleStateDeleted
	}
	return false
}

func OpensearchOpensearchClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OpensearchClusterClient().GetOpensearchCluster(context.Background(), oci_opensearch.GetOpensearchClusterRequest{
		OpensearchClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
