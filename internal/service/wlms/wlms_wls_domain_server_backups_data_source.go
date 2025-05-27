// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsWlsDomainServerBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsWlsDomainServerBackups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wls_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup_collection": {
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
									"backup_location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"content_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
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

func readWlmsWlsDomainServerBackups(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainServerBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainServerBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListWlsDomainServerBackupsResponse
}

func (s *WlmsWlsDomainServerBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainServerBackupsDataSourceCrud) Get() error {
	request := oci_wlms.ListWlsDomainServerBackupsRequest{}

	if serverId, ok := s.D.GetOkExists("server_id"); ok {
		tmp := serverId.(string)
		request.ServerId = &tmp
	}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListWlsDomainServerBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWlsDomainServerBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsWlsDomainServerBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsWlsDomainServerBackupsDataSource-", WlmsWlsDomainServerBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	wlsDomainServerBackup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BackupSummaryToMap(item))
	}
	wlsDomainServerBackup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsWlsDomainServerBackupsDataSource().Schema["backup_collection"].Elem.(*schema.Resource).Schema)
		wlsDomainServerBackup["items"] = items
	}

	resources = append(resources, wlsDomainServerBackup)
	if err := s.D.Set("backup_collection", resources); err != nil {
		return err
	}

	return nil
}

func BackupSummaryToMap(obj oci_wlms.BackupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupLocation != nil {
		result["backup_location"] = string(*obj.BackupLocation)
	}

	result["content_type"] = string(obj.ContentType)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	result["type"] = string(obj.Type)

	return result
}
