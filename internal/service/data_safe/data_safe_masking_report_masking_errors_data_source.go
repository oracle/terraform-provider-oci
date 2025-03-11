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

func DataSafeMaskingReportMaskingErrorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeMaskingReportMaskingErrors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"masking_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"step_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masking_error_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"error": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"failed_statement": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"step_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
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

func readDataSafeMaskingReportMaskingErrors(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingReportMaskingErrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingReportMaskingErrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListMaskingErrorsResponse
}

func (s *DataSafeMaskingReportMaskingErrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingReportMaskingErrorsDataSourceCrud) Get() error {
	request := oci_data_safe.ListMaskingErrorsRequest{}

	if maskingReportId, ok := s.D.GetOkExists("masking_report_id"); ok {
		tmp := maskingReportId.(string)
		request.MaskingReportId = &tmp
	}

	if stepName, ok := s.D.GetOkExists("step_name"); ok {
		request.StepName = oci_data_safe.ListMaskingErrorsStepNameEnum(stepName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListMaskingErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaskingErrors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeMaskingReportMaskingErrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeMaskingReportMaskingErrorsDataSource-", DataSafeMaskingReportMaskingErrorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	maskingReportMaskingError := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaskingErrorSummaryToMap(item))
	}
	maskingReportMaskingError["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeMaskingReportMaskingErrorsDataSource().Schema["masking_error_collection"].Elem.(*schema.Resource).Schema)
		maskingReportMaskingError["items"] = items
	}

	resources = append(resources, maskingReportMaskingError)
	if err := s.D.Set("masking_error_collection", resources); err != nil {
		return err
	}

	return nil
}

func MaskingErrorSummaryToMap(obj oci_data_safe.MaskingErrorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Error != nil {
		result["error"] = string(*obj.Error)
	}

	if obj.FailedStatement != nil {
		result["failed_statement"] = string(*obj.FailedStatement)
	}

	result["step_name"] = string(obj.StepName)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
