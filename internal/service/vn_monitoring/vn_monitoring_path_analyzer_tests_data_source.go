// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vn_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_vn_monitoring "github.com/oracle/oci-go-sdk/v65/vnmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func VnMonitoringPathAnalyzerTestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVnMonitoringPathAnalyzerTests,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"path_analyzer_test_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(VnMonitoringPathAnalyzerTestResource()),
						},
					},
				},
			},
		},
	}
}

func readVnMonitoringPathAnalyzerTests(d *schema.ResourceData, m interface{}) error {
	sync := &VnMonitoringPathAnalyzerTestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VnMonitoringClient()

	return tfresource.ReadResource(sync)
}

type VnMonitoringPathAnalyzerTestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_vn_monitoring.VnMonitoringClient
	Res    *oci_vn_monitoring.ListPathAnalyzerTestsResponse
}

func (s *VnMonitoringPathAnalyzerTestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VnMonitoringPathAnalyzerTestsDataSourceCrud) Get() error {
	request := oci_vn_monitoring.ListPathAnalyzerTestsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_vn_monitoring.PathAnalyzerTestLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "vn_monitoring")

	response, err := s.Client.ListPathAnalyzerTests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPathAnalyzerTests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VnMonitoringPathAnalyzerTestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("VnMonitoringPathAnalyzerTestsDataSource-", VnMonitoringPathAnalyzerTestsDataSource(), s.D))
	resources := []map[string]interface{}{}
	pathAnalyzerTest := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PathAnalyzerTestSummaryToMap(item))
	}
	pathAnalyzerTest["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, VnMonitoringPathAnalyzerTestsDataSource().Schema["path_analyzer_test_collection"].Elem.(*schema.Resource).Schema)
		pathAnalyzerTest["items"] = items
	}

	resources = append(resources, pathAnalyzerTest)
	if err := s.D.Set("path_analyzer_test_collection", resources); err != nil {
		return err
	}

	return nil
}
