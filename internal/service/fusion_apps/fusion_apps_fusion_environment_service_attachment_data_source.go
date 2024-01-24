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

func FusionAppsFusionEnvironmentServiceAttachmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fusion_environment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["service_attachment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FusionAppsFusionEnvironmentServiceAttachmentResource(), fieldMap, readSingularFusionAppsFusionEnvironmentServiceAttachment)
}

func readSingularFusionAppsFusionEnvironmentServiceAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentServiceAttachmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentServiceAttachmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetServiceAttachmentResponse
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetServiceAttachmentRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	// the service_attachment_id is a composite id
	if serviceAttachmentId, ok := s.D.GetOkExists("service_attachment_id"); ok {
		_, serviceAttachmentIdStr, err := parseFusionEnvironmentServiceAttachmentCompositeId(serviceAttachmentId.(string))
		if err == nil {
			request.ServiceAttachmentId = &serviceAttachmentIdStr
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetServiceAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSkuBased != nil {
		s.D.Set("is_sku_based", *s.Res.IsSkuBased)
	}

	if s.Res.ServiceInstanceId != nil {
		s.D.Set("service_instance_id", *s.Res.ServiceInstanceId)
	}

	s.D.Set("service_instance_type", s.Res.ServiceInstanceType)

	if s.Res.ServiceUrl != nil {
		s.D.Set("service_url", *s.Res.ServiceUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
