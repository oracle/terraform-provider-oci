// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsaPsaWorkRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readPsaPsaWorkRequestsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"work_request_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"percent_complete": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"resources": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"action_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"entity_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"entity_uri": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"metadata": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
											},
										},
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_accepted": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_finished": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readPsaPsaWorkRequestsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsaPsaWorkRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type PsaPsaWorkRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psa.PrivateServiceAccessClient
	Res    *oci_psa.ListPsaWorkRequestsResponse
}

func (s *PsaPsaWorkRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsaPsaWorkRequestsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_psa.ListPsaWorkRequestsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_psa.ListPsaWorkRequestsStatusEnum(status.(string))
	}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psa")

	response, err := s.Client.ListPsaWorkRequests(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPsaWorkRequests(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PsaPsaWorkRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsaPsaWorkRequestsDataSource-", PsaPsaWorkRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}
	psaWorkRequest := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WorkRequestSummaryToMap(item))
	}
	psaWorkRequest["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, PsaPsaWorkRequestsDataSource().Schema["work_request_summary_collection"].Elem.(*schema.Resource).Schema)
		psaWorkRequest["items"] = items
	}

	resources = append(resources, psaWorkRequest)
	if err := s.D.Set("work_request_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func WorkRequestResourceToMap(obj oci_psa.WorkRequestResource) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_type"] = string(obj.ActionType)

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	if obj.EntityUri != nil {
		result["entity_uri"] = string(*obj.EntityUri)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	result["metadata"] = obj.Metadata

	return result
}

func WorkRequestSummaryToMap(obj oci_psa.WorkRequestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["operation_type"] = string(obj.OperationType)

	if obj.PercentComplete != nil {
		result["percent_complete"] = float32(*obj.PercentComplete)
	}

	resources := []interface{}{}
	for _, item := range obj.Resources {
		resources = append(resources, WorkRequestResourceToMap(item))
	}
	result["resources"] = resources

	result["status"] = string(obj.Status)

	if obj.TimeAccepted != nil {
		result["time_accepted"] = obj.TimeAccepted.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
