// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLustreFileStorageAvailableOverrideMaintenanceStartTimes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"available_override_maintenance_start_time_collection": {
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
									"start_times": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"time_date_available": {
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

func readLustreFileStorageAvailableOverrideMaintenanceStartTimes(d *schema.ResourceData, m interface{}) error {
	sync := &LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.ReadResource(sync)
}

type LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_lustre_file_storage.LustreFileStorageClient
	Res    *oci_lustre_file_storage.ListAvailableOverrideMaintenanceStartTimesResponse
}

func (s *LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSourceCrud) Get() error {
	request := oci_lustre_file_storage.ListAvailableOverrideMaintenanceStartTimesRequest{}

	if date, ok := s.D.GetOkExists("date"); ok {
		tmp := date.(string)
		request.Date = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "lustre_file_storage")

	response, err := s.Client.ListAvailableOverrideMaintenanceStartTimes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAvailableOverrideMaintenanceStartTimes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSource-", LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSource(), s.D))
	resources := []map[string]interface{}{}
	availableOverrideMaintenanceStartTime := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AvailableOverrideMaintenanceStartTimeSummaryToMap(item))
	}
	availableOverrideMaintenanceStartTime["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSource().Schema["available_override_maintenance_start_time_collection"].Elem.(*schema.Resource).Schema)
		availableOverrideMaintenanceStartTime["items"] = items
	}

	resources = append(resources, availableOverrideMaintenanceStartTime)
	if err := s.D.Set("available_override_maintenance_start_time_collection", resources); err != nil {
		return err
	}

	return nil
}

func AvailableOverrideMaintenanceStartTimeSummaryToMap(obj oci_lustre_file_storage.AvailableOverrideMaintenanceStartTimeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["start_times"] = obj.StartTimes

	if obj.TimeDateAvailable != nil {
		result["time_date_available"] = obj.TimeDateAvailable.String()
	}

	return result
}
