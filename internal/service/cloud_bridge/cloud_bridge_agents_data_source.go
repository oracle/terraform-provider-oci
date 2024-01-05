// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAgentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudBridgeAgents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"agent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"environment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"agent_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudBridgeAgentResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudBridgeAgents(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeAgentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.OcbAgentSvcClient
	Res    *oci_cloud_bridge.ListAgentsResponse
}

func (s *CloudBridgeAgentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeAgentsDataSourceCrud) Get() error {
	request := oci_cloud_bridge.ListAgentsRequest{}

	if agentId, ok := s.D.GetOkExists("id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if environmentId, ok := s.D.GetOkExists("environment_id"); ok {
		tmp := environmentId.(string)
		request.EnvironmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_bridge.AgentLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.ListAgents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAgents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudBridgeAgentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudBridgeAgentsDataSource-", CloudBridgeAgentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	agent := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AgentSummaryToMap(item))
	}
	agent["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudBridgeAgentsDataSource().Schema["agent_collection"].Elem.(*schema.Resource).Schema)
		agent["items"] = items
	}

	resources = append(resources, agent)
	if err := s.D.Set("agent_collection", resources); err != nil {
		return err
	}

	return nil
}
