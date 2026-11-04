package containerengine

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestContainerengineSystemTags(t *testing.T) {
	t.Setenv("TF_ACC", "1")

	tagCases := []struct {
		name       string
		systemTags map[string]map[string]interface{}
		want       map[string]interface{}
	}{
		{
			name: "nil",
			want: map[string]interface{}{},
		},
		{
			name:       "empty",
			systemTags: map[string]map[string]interface{}{},
			want:       map[string]interface{}{},
		},
		{
			name: "populated",
			systemTags: map[string]map[string]interface{}{
				"orcl-cloud": {"free-tier-retained": "true"},
			},
			want: map[string]interface{}{"orcl-cloud.free-tier-retained": "true"},
		},
	}

	for _, tagCase := range tagCases {
		t.Run(tagCase.name, func(t *testing.T) {
			cluster := oci_containerengine.Cluster{
				Id:         common.String("cluster-id"),
				SystemTags: tagCase.systemTags,
			}
			nodePool := oci_containerengine.NodePool{
				Id:         common.String("node-pool-id"),
				SystemTags: tagCase.systemTags,
				Nodes:      []oci_containerengine.Node{{SystemTags: tagCase.systemTags}},
			}

			testCases := []struct {
				name        string
				newResource func() *schema.Resource
				setData     func(*schema.ResourceData) error
				paths       []string
			}{
				{
					name:        "cluster_resource",
					newResource: ContainerengineClusterResource,
					setData: func(data *schema.ResourceData) error {
						crud := ContainerengineClusterResourceCrud{
							BaseCrud: tfresource.BaseCrud{D: data},
							Res:      &cluster,
						}
						return crud.SetData()
					},
					paths: []string{"system_tags"},
				},
				{
					name:        "cluster_data_source",
					newResource: ContainerengineClusterDataSource,
					setData: func(data *schema.ResourceData) error {
						crud := ContainerengineClusterDataSourceCrud{
							D:   data,
							Res: &oci_containerengine.GetClusterResponse{Cluster: cluster},
						}
						return crud.SetData()
					},
					paths: []string{"system_tags"},
				},
				{
					name:        "clusters_data_source",
					newResource: ContainerengineClustersDataSource,
					setData: func(data *schema.ResourceData) error {
						crud := ContainerengineClustersDataSourceCrud{
							D: data,
							Res: &oci_containerengine.ListClustersResponse{
								Items: []oci_containerengine.ClusterSummary{{
									CompartmentId: common.String("compartment-id"),
									SystemTags:    tagCase.systemTags,
								}},
							},
						}
						return crud.SetData()
					},
					paths: []string{"clusters.0.system_tags"},
				},
				{
					name:        "node_pool_resource",
					newResource: ContainerengineNodePoolResource,
					setData: func(data *schema.ResourceData) error {
						crud := ContainerengineNodePoolResourceCrud{
							BaseCrud: tfresource.BaseCrud{D: data},
							Res:      &nodePool,
						}
						return crud.SetData()
					},
					paths: []string{"system_tags", "nodes.0.system_tags"},
				},
				{
					name:        "node_pool_data_source",
					newResource: ContainerengineNodePoolDataSource,
					setData: func(data *schema.ResourceData) error {
						crud := ContainerengineNodePoolDataSourceCrud{
							D:   data,
							Res: &oci_containerengine.GetNodePoolResponse{NodePool: nodePool},
						}
						return crud.SetData()
					},
					paths: []string{"system_tags", "nodes.0.system_tags"},
				},
				{
					name:        "node_pools_data_source",
					newResource: ContainerengineNodePoolsDataSource,
					setData: func(data *schema.ResourceData) error {
						crud := ContainerengineNodePoolsDataSourceCrud{
							D: data,
							Res: &oci_containerengine.ListNodePoolsResponse{
								Items: []oci_containerengine.NodePoolSummary{{
									CompartmentId: common.String("compartment-id"),
									SystemTags:    tagCase.systemTags,
								}},
							},
						}
						return crud.SetData()
					},
					paths: []string{"node_pools.0.system_tags"},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					data := schema.TestResourceDataRaw(t, testCase.newResource().Schema, nil)
					if err := testCase.setData(data); err != nil {
						t.Fatalf("SetData() failed: %v", err)
					}
					for _, path := range testCase.paths {
						if got := data.Get(path); !reflect.DeepEqual(got, tagCase.want) {
							t.Errorf("%s = %#v, want %#v", path, got, tagCase.want)
						}
					}
				})
			}
		})
	}
}
