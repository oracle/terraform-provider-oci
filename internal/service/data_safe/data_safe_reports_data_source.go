// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeReportsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeReports,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mime_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_definition_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_generated_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_generated_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeReportResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeReports(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeReportsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListReportsResponse
}

func (s *DataSafeReportsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeReportsDataSourceCrud) Get() error {
	request := oci_data_safe.ListReportsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListReportsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if mimeType, ok := s.D.GetOkExists("mime_type"); ok {
		request.MimeType = oci_data_safe.ListReportsMimeTypeEnum(mimeType.(string))
	}

	if reportDefinitionId, ok := s.D.GetOkExists("report_definition_id"); ok {
		tmp := reportDefinitionId.(string)
		request.ReportDefinitionId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListReportsLifecycleStateEnum(state.(string))
	}

	if timeGeneratedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_generated_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeGeneratedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeGeneratedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeGeneratedLessThan, ok := s.D.GetOkExists("time_generated_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeGeneratedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeGeneratedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_data_safe.ListReportsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListReports(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReports(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeReportsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeReportsDataSource-", DataSafeReportsDataSource(), s.D))
	resources := []map[string]interface{}{}
	report := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ReportSummaryToMap(item))
	}
	report["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeReportsDataSource().Schema["report_collection"].Elem.(*schema.Resource).Schema)
		report["items"] = items
	}

	resources = append(resources, report)
	if err := s.D.Set("report_collection", resources); err != nil {
		return err
	}

	return nil
}

func ReportSummaryToMap(obj oci_data_safe.ReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["mime_type"] = string(obj.MimeType)

	if obj.ReportDefinitionId != nil {
		result["report_definition_id"] = string(*obj.ReportDefinitionId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeGenerated != nil {
		result["time_generated"] = obj.TimeGenerated.String()
	}

	result["type"] = string(obj.Type)

	return result
}
