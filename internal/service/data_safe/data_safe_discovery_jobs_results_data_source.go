// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeDiscoveryJobsResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeDiscoveryJobsResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"column_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"discovery_job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"discovery_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_result_applied": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"object": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"planned_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schema_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"discovery_job_result_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataSafeDiscoveryJobsResultResource(),
						},
					},
				},
			},
		},
	}
}

func readDataSafeDiscoveryJobsResults(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobsResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeDiscoveryJobsResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListDiscoveryJobResultsResponse
}

func (s *DataSafeDiscoveryJobsResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeDiscoveryJobsResultsDataSourceCrud) Get() error {
	request := oci_data_safe.ListDiscoveryJobResultsRequest{}

	if columnName, ok := s.D.GetOkExists("column_name"); ok {
		interfaces := columnName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("column_name") {
			request.ColumnName = tmp
		}
	}

	if discoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := discoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}

	if discoveryType, ok := s.D.GetOkExists("discovery_type"); ok {
		request.DiscoveryType = oci_data_safe.DiscoveryJobDiscoveryTypeEnum(discoveryType.(string))
	}

	if isResultApplied, ok := s.D.GetOkExists("is_result_applied"); ok {
		tmp := isResultApplied.(bool)
		request.IsResultApplied = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		interfaces := object.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object") {
			request.ObjectName = tmp
		}
	}

	if plannedAction, ok := s.D.GetOkExists("planned_action"); ok {
		request.PlannedAction = oci_data_safe.DiscoveryJobResultPlannedActionEnum(plannedAction.(string))
	}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		interfaces := schemaName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schema_name") {
			request.SchemaName = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListDiscoveryJobResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDiscoveryJobResults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeDiscoveryJobsResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeDiscoveryJobsResultsDataSource-", DataSafeDiscoveryJobsResultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	discoveryJobsResult := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DiscoveryJobResultSummaryToMap(item))
	}
	discoveryJobsResult["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeDiscoveryJobsResultsDataSource().Schema["discovery_job_result_collection"].Elem.(*schema.Resource).Schema)
		discoveryJobsResult["items"] = items
	}

	resources = append(resources, discoveryJobsResult)
	if err := s.D.Set("discovery_job_result_collection", resources); err != nil {
		return err
	}

	return nil
}
