// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardWlpAgentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardWlpAgents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wlp_agent_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudGuardWlpAgentResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudGuardWlpAgents(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardWlpAgentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardWlpAgentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListWlpAgentsResponse
}

func (s *CloudGuardWlpAgentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardWlpAgentsDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListWlpAgentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListWlpAgents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWlpAgents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardWlpAgentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardWlpAgentsDataSource-", CloudGuardWlpAgentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	wlpAgent := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WlpAgentSummaryToMap(item))
	}
	wlpAgent["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardWlpAgentsDataSource().Schema["wlp_agent_collection"].Elem.(*schema.Resource).Schema)
		wlpAgent["items"] = items
	}

	resources = append(resources, wlpAgent)
	if err := s.D.Set("wlp_agent_collection", resources); err != nil {
		return err
	}

	return nil
}
