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

func EmWarehouseEmWarehouseEtlRunDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularEmWarehouseEmWarehouseEtlRun,
		Schema: map[string]*schema.Schema{
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_em_warehouse_em_warehouse_etl_run", "oci_em_warehouse_em_warehouse_etl_runs"),
	}
}

func readSingularEmWarehouseEmWarehouseEtlRun(d *schema.ResourceData, m interface{}) error {
	sync := &EmWarehouseEmWarehouseEtlRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmWarehouseClient()

	return tfresource.ReadResource(sync)
}

type EmWarehouseEmWarehouseEtlRunDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_em_warehouse.EmWarehouseClient
	Res    *oci_em_warehouse.ListEtlRunsResponse
}

func (s *EmWarehouseEmWarehouseEtlRunDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmWarehouseEmWarehouseEtlRunDataSourceCrud) Get() error {
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
	return nil
}

func (s *EmWarehouseEmWarehouseEtlRunDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmWarehouseEmWarehouseEtlRunDataSource-", EmWarehouseEmWarehouseEtlRunDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EtlRunSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func EM_WarehouseEtlRunSummaryToMap(obj oci_em_warehouse.EtlRunSummary) map[string]interface{} {
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
