// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingReportSubsettingErrorsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingReportSubsettingErrorsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"step_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subsetting_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subsetting_error_collection": {
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

func readDataSafeSubsettingReportSubsettingErrorsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingReportSubsettingErrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingReportSubsettingErrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingErrorsResponse
}

func (s *DataSafeSubsettingReportSubsettingErrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingReportSubsettingErrorsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingErrorsRequest{}

	if stepName, ok := s.D.GetOkExists("step_name"); ok {
		request.StepName = oci_data_safe.ListSubsettingErrorsStepNameEnum(stepName.(string))
	}

	if subsettingReportId, ok := s.D.GetOkExists("subsetting_report_id"); ok {
		tmp := subsettingReportId.(string)
		request.SubsettingReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettingErrors(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingErrors(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingReportSubsettingErrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingReportSubsettingErrorsDataSource-", DataSafeSubsettingReportSubsettingErrorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingReportSubsettingError := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingErrorSummaryToMap(item))
	}
	subsettingReportSubsettingError["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingReportSubsettingErrorsDataSource().Schema["subsetting_error_collection"].Elem.(*schema.Resource).Schema)
		subsettingReportSubsettingError["items"] = items
	}

	resources = append(resources, subsettingReportSubsettingError)
	if err := s.D.Set("subsetting_error_collection", resources); err != nil {
		return err
	}

	return nil
}

func SubsettingErrorSummaryToMap(obj oci_data_safe.SubsettingErrorSummary) map[string]interface{} {
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
