// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentDataMaskingActivityDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_masking_activity_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["fusion_environment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FusionAppsFusionEnvironmentDataMaskingActivityResource(), fieldMap, readSingularFusionAppsFusionEnvironmentDataMaskingActivity)
}

func readSingularFusionAppsFusionEnvironmentDataMaskingActivity(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentDataMaskingActivityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentDataMaskingActivityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetDataMaskingActivityResponse
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetDataMaskingActivityRequest{}

	if dataMaskingActivityId, ok := s.D.GetOkExists("data_masking_activity_id"); ok {
		tmp := dataMaskingActivityId.(string)
		request.DataMaskingActivityId = &tmp
	}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetDataMaskingActivity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeMaskingFinish != nil {
		s.D.Set("time_masking_finish", s.Res.TimeMaskingFinish.String())
	}

	if s.Res.TimeMaskingStart != nil {
		s.D.Set("time_masking_start", s.Res.TimeMaskingStart.String())
	}

	return nil
}
