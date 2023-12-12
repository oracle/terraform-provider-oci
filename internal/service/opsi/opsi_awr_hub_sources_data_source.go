// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiAwrHubSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpsiAwrHubSources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"awr_hub_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"awr_hub_source_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"awr_hub_source_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpsiAwrHubSourceResource()),
						},
					},
				},
			},
		},
	}
}

func readOpsiAwrHubSources(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiAwrHubSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListAwrHubSourcesResponse
}

func (s *OpsiAwrHubSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiAwrHubSourcesDataSourceCrud) Get() error {
	request := oci_opsi.ListAwrHubSourcesRequest{}

	if awrHubId, ok := s.D.GetOkExists("awr_hub_id"); ok {
		tmp := awrHubId.(string)
		request.AwrHubId = &tmp
	}

	if awrHubSourceId, ok := s.D.GetOkExists("id"); ok {
		tmp := awrHubSourceId.(string)
		request.AwrHubSourceId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		interfaces := sourceType.([]interface{})
		tmp := make([]oci_opsi.AwrHubSourceTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.AwrHubSourceTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("source_type") {
			request.SourceType = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		interfaces := state.([]interface{})
		tmp := make([]oci_opsi.AwrHubSourceLifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.AwrHubSourceLifecycleStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		interfaces := status.([]interface{})
		tmp := make([]oci_opsi.AwrHubSourceStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.AwrHubSourceStatusEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("state") {
			request.Status = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListAwrHubSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAwrHubSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiAwrHubSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiAwrHubSourcesDataSource-", OpsiAwrHubSourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	awrHubSource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AwrHubSourceSummaryToMap(item))
	}
	awrHubSource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiAwrHubSourcesDataSource().Schema["awr_hub_source_summary_collection"].Elem.(*schema.Resource).Schema)
		awrHubSource["items"] = items
	}

	resources = append(resources, awrHubSource)
	if err := s.D.Set("awr_hub_source_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
