// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentPluginCountDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularManagementAgentManagementAgentPluginCount,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_by": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dimensions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"plugin_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"plugin_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readSingularManagementAgentManagementAgentPluginCount(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentPluginCountDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentPluginCountDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.SummarizeManagementAgentPluginCountsResponse
}

func (s *ManagementAgentManagementAgentPluginCountDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentPluginCountDataSourceCrud) Get() error {
	request := oci_management_agent.SummarizeManagementAgentPluginCountsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_management_agent.SummarizeManagementAgentPluginCountsGroupByEnum(groupBy.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.SummarizeManagementAgentPluginCounts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementAgentManagementAgentPluginCountDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentPluginCountDataSource-", ManagementAgentManagementAgentPluginCountDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagementAgentPluginAggregationToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func ManagementAgentPluginAggregationToMap(obj oci_management_agent.ManagementAgentPluginAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = int(*obj.Count)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{ManagementAgentPluginAggregationDimensionsToMap(obj.Dimensions)}
	}

	return result
}

func ManagementAgentPluginAggregationDimensionsToMap(obj *oci_management_agent.ManagementAgentPluginAggregationDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PluginDisplayName != nil {
		result["plugin_display_name"] = string(*obj.PluginDisplayName)
	}

	if obj.PluginName != nil {
		result["plugin_name"] = string(*obj.PluginName)
	}

	return result
}
