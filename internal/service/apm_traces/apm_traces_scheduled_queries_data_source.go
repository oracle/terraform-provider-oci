// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmTracesScheduledQueriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmTracesScheduledQueries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduled_query_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApmTracesScheduledQueryResource()),
						},
					},
				},
			},
		},
	}
}

func readApmTracesScheduledQueries(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesScheduledQueriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledQueryClient()

	return tfresource.ReadResource(sync)
}

type ApmTracesScheduledQueriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_traces.ScheduledQueryClient
	Res    *oci_apm_traces.ListScheduledQueriesResponse
}

func (s *ApmTracesScheduledQueriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmTracesScheduledQueriesDataSourceCrud) Get() error {
	request := oci_apm_traces.ListScheduledQueriesRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_traces")

	response, err := s.Client.ListScheduledQueries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScheduledQueries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmTracesScheduledQueriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmTracesScheduledQueriesDataSource-", ApmTracesScheduledQueriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	scheduledQuery := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduledQuerySummaryToMap(item))
	}
	scheduledQuery["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApmTracesScheduledQueriesDataSource().Schema["scheduled_query_collection"].Elem.(*schema.Resource).Schema)
		scheduledQuery["items"] = items
	}

	resources = append(resources, scheduledQuery)
	if err := s.D.Set("scheduled_query_collection", resources); err != nil {
		return err
	}

	return nil
}
