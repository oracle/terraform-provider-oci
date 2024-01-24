// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package em_warehouse

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_em_warehouse "github.com/oracle/oci-go-sdk/v65/emwarehouse"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmWarehouseEmWarehouseEtlRunsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readEmWarehouseEmWarehouseEtlRuns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"em_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"etl_run_collection": {
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
												"data_read_in_bytes": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"data_written": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"defined_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"lifecycle_details": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"run_duration_in_milliseconds": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_created": {
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
				},
			},
		},
	}
}

func readEmWarehouseEmWarehouseEtlRuns(d *schema.ResourceData, m interface{}) error {
	sync := &EmWarehouseEmWarehouseEtlRunsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmWarehouseClient()

	return tfresource.ReadResource(sync)
}

type EmWarehouseEmWarehouseEtlRunsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_em_warehouse.EmWarehouseClient
	Res    *oci_em_warehouse.ListEtlRunsResponse
}

func (s *EmWarehouseEmWarehouseEtlRunsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmWarehouseEmWarehouseEtlRunsDataSourceCrud) Get() error {
	request := oci_em_warehouse.ListEtlRunsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if emWarehouseId, ok := s.D.GetOkExists("em_warehouse_id"); ok {
		tmp := emWarehouseId.(string)
		request.EmWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "em_warehouse")

	response, err := s.Client.ListEtlRuns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEtlRuns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *EmWarehouseEmWarehouseEtlRunsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmWarehouseEmWarehouseEtlRunsDataSource-", EmWarehouseEmWarehouseEtlRunsDataSource(), s.D))
	resources := []map[string]interface{}{}
	emWarehouseEtlRun := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EtlRunSummaryToMap(item))
	}
	emWarehouseEtlRun["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, EmWarehouseEmWarehouseEtlRunsDataSource().Schema["etl_run_collection"].Elem.(*schema.Resource).Schema)
		emWarehouseEtlRun["items"] = items
	}

	resources = append(resources, emWarehouseEtlRun)
	if err := s.D.Set("etl_run_collection", resources); err != nil {
		return err
	}

	return nil
}

func EtlRunSummaryToMap(obj oci_em_warehouse.EtlRunSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DataReadInBytes != nil {
		result["data_read_in_bytes"] = strconv.FormatInt(*obj.DataReadInBytes, 10)
	}

	if obj.DataWritten != nil {
		result["data_written"] = strconv.FormatInt(*obj.DataWritten, 10)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.RunDurationInMilliseconds != nil {
		result["run_duration_in_milliseconds"] = strconv.FormatInt(*obj.RunDurationInMilliseconds, 10)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
