// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardAdhocQueriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardAdhocQueries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"adhoc_query_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"time_ended_filter_query_param": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_started_filter_query_param": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"adhoc_query_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudGuardAdhocQueryResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudGuardAdhocQueries(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardAdhocQueriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardAdhocQueriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListAdhocQueriesResponse
}

func (s *CloudGuardAdhocQueriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardAdhocQueriesDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListAdhocQueriesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_cloud_guard.ListAdhocQueriesAccessLevelEnum(accessLevel.(string))
	}

	if adhocQueryStatus, ok := s.D.GetOkExists("adhoc_query_status"); ok {
		request.AdhocQueryStatus = oci_cloud_guard.ListAdhocQueriesAdhocQueryStatusEnum(adhocQueryStatus.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if timeEndedFilterQueryParam, ok := s.D.GetOkExists("time_ended_filter_query_param"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEndedFilterQueryParam.(string))
		if err != nil {
			return err
		}
		request.TimeEndedFilterQueryParam = &oci_common.SDKTime{Time: tmp}
	}

	if timeStartedFilterQueryParam, ok := s.D.GetOkExists("time_started_filter_query_param"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStartedFilterQueryParam.(string))
		if err != nil {
			return err
		}
		request.TimeStartedFilterQueryParam = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListAdhocQueries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAdhocQueries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardAdhocQueriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardAdhocQueriesDataSource-", CloudGuardAdhocQueriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	adhocQuery := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdhocQuerySummaryToMap(item))
	}
	adhocQuery["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardAdhocQueriesDataSource().Schema["adhoc_query_collection"].Elem.(*schema.Resource).Schema)
		adhocQuery["items"] = items
	}

	resources = append(resources, adhocQuery)
	if err := s.D.Set("adhoc_query_collection", resources); err != nil {
		return err
	}

	return nil
}
