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

func CloudBridgeApplianceImageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCloudBridgeApplianceImage,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_cloud_bridge_appliance_image", "oci_cloud_bridge_appliance_images"),
	}
}

func readSingularCloudBridgeApplianceImage(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeApplianceImageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeApplianceImageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.OcbAgentSvcClient
	Res    *oci_cloud_bridge.ListApplianceImagesResponse
}

func (s *CloudBridgeApplianceImageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeApplianceImageDataSourceCrud) Get() error {
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
	return nil
}

func (s *CloudBridgeApplianceImageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudBridgeApplianceImageDataSource-", CloudBridgeApplianceImageDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ApplianceImageSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
