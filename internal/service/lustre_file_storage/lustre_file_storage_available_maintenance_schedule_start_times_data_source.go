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

func LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLustreFileStorageAvailableMaintenanceScheduleStartTimes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"day_of_week": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"available_maintenance_schedule_start_time_collection": {
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
									"day_of_week": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_times": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
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

func readLustreFileStorageAvailableMaintenanceScheduleStartTimes(d *schema.ResourceData, m interface{}) error {
	sync := &LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.ReadResource(sync)
}

type LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_lustre_file_storage.LustreFileStorageClient
	Res    *oci_lustre_file_storage.ListAvailableMaintenanceScheduleStartTimesResponse
}

func (s *LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSourceCrud) Get() error {
	request := oci_lustre_file_storage.ListAvailableMaintenanceScheduleStartTimesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dayOfWeek, ok := s.D.GetOkExists("day_of_week"); ok {
		request.DayOfWeek = oci_lustre_file_storage.ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum(dayOfWeek.(string))
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "lustre_file_storage")

	response, err := s.Client.ListAvailableMaintenanceScheduleStartTimes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAvailableMaintenanceScheduleStartTimes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSource-", LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSource(), s.D))
	resources := []map[string]interface{}{}
	availableMaintenanceScheduleStartTime := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AvailableMaintenanceScheduleStartTimeSummaryToMap(item))
	}
	availableMaintenanceScheduleStartTime["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSource().Schema["available_maintenance_schedule_start_time_collection"].Elem.(*schema.Resource).Schema)
		availableMaintenanceScheduleStartTime["items"] = items
	}

	resources = append(resources, availableMaintenanceScheduleStartTime)
	if err := s.D.Set("available_maintenance_schedule_start_time_collection", resources); err != nil {
		return err
	}

	return nil
}

func AvailableMaintenanceScheduleStartTimeSummaryToMap(obj oci_lustre_file_storage.AvailableMaintenanceScheduleStartTimeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["day_of_week"] = string(obj.DayOfWeek)

	result["start_times"] = obj.StartTimes

	return result
}
