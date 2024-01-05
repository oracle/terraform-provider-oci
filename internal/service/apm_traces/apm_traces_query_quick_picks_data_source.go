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

func ApmTracesQueryQuickPicksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmTracesQueryQuickPicks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quick_picks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"quick_pick_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quick_pick_query": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readApmTracesQueryQuickPicks(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesQueryQuickPicksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueryClient()

	return tfresource.ReadResource(sync)
}

type ApmTracesQueryQuickPicksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_traces.QueryClient
	Res    *oci_apm_traces.ListQuickPicksResponse
}

func (s *ApmTracesQueryQuickPicksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmTracesQueryQuickPicksDataSourceCrud) Get() error {
	request := oci_apm_traces.ListQuickPicksRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_traces")

	response, err := s.Client.ListQuickPicks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListQuickPicks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmTracesQueryQuickPicksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmTracesQueryQuickPicksDataSource-", ApmTracesQueryQuickPicksDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		queryQuickPick := map[string]interface{}{}

		if r.QuickPickName != nil {
			queryQuickPick["quick_pick_name"] = *r.QuickPickName
		}

		if r.QuickPickQuery != nil {
			queryQuickPick["quick_pick_query"] = *r.QuickPickQuery
		}

		resources = append(resources, queryQuickPick)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ApmTracesQueryQuickPicksDataSource().Schema["quick_picks"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("quick_picks", resources); err != nil {
		return err
	}

	return nil
}
