// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceSoftwareUpdatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceSoftwareUpdates,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"software_update_collection": {
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
									"software_update_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_update_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_update_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_due": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_released": {
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

func readBdsBdsInstanceSoftwareUpdates(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceSoftwareUpdatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceSoftwareUpdatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListSoftwareUpdatesResponse
}

func (s *BdsBdsInstanceSoftwareUpdatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceSoftwareUpdatesDataSourceCrud) Get() error {
	request := oci_bds.ListSoftwareUpdatesRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListSoftwareUpdates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSoftwareUpdates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceSoftwareUpdatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceSoftwareUpdatesDataSource-", BdsBdsInstanceSoftwareUpdatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	bdsInstanceSoftwareUpdate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SoftwareUpdateSummaryToMap(item))
	}
	bdsInstanceSoftwareUpdate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BdsBdsInstanceSoftwareUpdatesDataSource().Schema["software_update_collection"].Elem.(*schema.Resource).Schema)
		bdsInstanceSoftwareUpdate["items"] = items
	}

	resources = append(resources, bdsInstanceSoftwareUpdate)
	if err := s.D.Set("software_update_collection", resources); err != nil {
		return err
	}

	return nil
}

func SoftwareUpdateSummaryToMap(obj oci_bds.SoftwareUpdateSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SoftwareUpdateKey != nil {
		result["software_update_key"] = string(*obj.SoftwareUpdateKey)
	}

	result["software_update_type"] = string(obj.SoftwareUpdateType)

	if obj.SoftwareUpdateVersion != nil {
		result["software_update_version"] = string(*obj.SoftwareUpdateVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeReleased != nil {
		result["time_released"] = obj.TimeReleased.String()
	}

	return result
}
