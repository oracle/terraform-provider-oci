// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardDataSourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_source_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardDataSourceResource(), fieldMap, readSingularCloudGuardDataSource)
}

func readSingularCloudGuardDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardDataSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetDataSourceResponse
}

func (s *CloudGuardDataSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardDataSourceDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetDataSourceRequest{}

	if dataSourceId, ok := s.D.GetOkExists("data_source_id"); ok {
		tmp := dataSourceId.(string)
		request.DataSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardDataSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataSourceDetails != nil {
		dataSourceDetailsArray := []interface{}{}
		if dataSourceDetailsMap := DataSourceDetailsToMap(&s.Res.DataSourceDetails); dataSourceDetailsMap != nil {
			dataSourceDetailsArray = append(dataSourceDetailsArray, dataSourceDetailsMap)
		}
		s.D.Set("data_source_details", dataSourceDetailsArray)
	} else {
		s.D.Set("data_source_details", nil)
	}

	dataSourceDetectorMappingInfo := []interface{}{}
	for _, item := range s.Res.DataSourceDetectorMappingInfo {
		dataSourceDetectorMappingInfo = append(dataSourceDetectorMappingInfo, DataSourceMappingInfoToMap(item))
	}
	s.D.Set("data_source_detector_mapping_info", dataSourceDetectorMappingInfo)

	s.D.Set("data_source_feed_provider", s.Res.DataSourceFeedProvider)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	regionStatusDetail := []interface{}{}
	for _, item := range s.Res.RegionStatusDetail {
		regionStatusDetail = append(regionStatusDetail, RegionStatusDetailToMap(item))
	}
	s.D.Set("region_status_detail", regionStatusDetail)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
