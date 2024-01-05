// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeApplianceImagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudBridgeApplianceImages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"appliance_image_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"checksum": {
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
									"download_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"file_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"format": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"platform": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_in_mbs": {
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
									"version": {
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

func readCloudBridgeApplianceImages(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeApplianceImagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeApplianceImagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.OcbAgentSvcClient
	Res    *oci_cloud_bridge.ListApplianceImagesResponse
}

func (s *CloudBridgeApplianceImagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeApplianceImagesDataSourceCrud) Get() error {
	request := oci_cloud_bridge.ListApplianceImagesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.ListApplianceImages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListApplianceImages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudBridgeApplianceImagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudBridgeApplianceImagesDataSource-", CloudBridgeApplianceImagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	applianceImage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ApplianceImageSummaryToMap(item))
	}
	applianceImage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudBridgeApplianceImagesDataSource().Schema["appliance_image_collection"].Elem.(*schema.Resource).Schema)
		applianceImage["items"] = items
	}

	resources = append(resources, applianceImage)
	if err := s.D.Set("appliance_image_collection", resources); err != nil {
		return err
	}

	return nil
}

func ApplianceImageSummaryToMap(obj oci_cloud_bridge.ApplianceImageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Checksum != nil {
		result["checksum"] = string(*obj.Checksum)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DownloadUrl != nil {
		result["download_url"] = string(*obj.DownloadUrl)
	}

	if obj.FileName != nil {
		result["file_name"] = string(*obj.FileName)
	}

	if obj.Format != nil {
		result["format"] = string(*obj.Format)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Platform != nil {
		result["platform"] = string(*obj.Platform)
	}

	if obj.SizeInMBs != nil {
		result["size_in_mbs"] = string(*obj.SizeInMBs)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
