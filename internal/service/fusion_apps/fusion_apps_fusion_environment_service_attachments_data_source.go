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

func FusionAppsFusionEnvironmentServiceAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFusionAppsFusionEnvironmentServiceAttachments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_instance_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_attachment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FusionAppsFusionEnvironmentServiceAttachmentResource()),
						},
					},
				},
			},
		},
	}
}

func readFusionAppsFusionEnvironmentServiceAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentServiceAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentServiceAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListServiceAttachmentsResponse
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentsDataSourceCrud) Get() error {
	request := oci_fusion_apps.ListServiceAttachmentsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if serviceInstanceType, ok := s.D.GetOkExists("service_instance_type"); ok {
		request.ServiceInstanceType = oci_fusion_apps.ServiceAttachmentServiceInstanceTypeEnum(serviceInstanceType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fusion_apps.ServiceAttachmentLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.ListServiceAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServiceAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentServiceAttachmentsDataSource-", FusionAppsFusionEnvironmentServiceAttachmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fusionEnvironmentServiceAttachment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ServiceAttachmentSummaryToMap(item))
	}
	fusionEnvironmentServiceAttachment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FusionAppsFusionEnvironmentServiceAttachmentsDataSource().Schema["service_attachment_collection"].Elem.(*schema.Resource).Schema)
		fusionEnvironmentServiceAttachment["items"] = items
	}

	resources = append(resources, fusionEnvironmentServiceAttachment)
	if err := s.D.Set("service_attachment_collection", resources); err != nil {
		return err
	}

	return nil
}
