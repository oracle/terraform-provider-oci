// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagementStationMirrorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagementStationMirrors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_station_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mirror_states": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mirrors_collection": {
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
									"arch_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_family": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"percentage": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_synced": {
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

func readOsManagementHubManagementStationMirrors(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagementStationMirrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementStationClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagementStationMirrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagementStationClient
	Res    *oci_os_management_hub.ListMirrorsResponse
}

func (s *OsManagementHubManagementStationMirrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagementStationMirrorsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListMirrorsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if managementStationId, ok := s.D.GetOkExists("management_station_id"); ok {
		tmp := managementStationId.(string)
		request.ManagementStationId = &tmp
	}

	if mirrorStates, ok := s.D.GetOkExists("mirror_states"); ok {
		interfaces := mirrorStates.([]interface{})
		tmp := make([]oci_os_management_hub.MirrorStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.MirrorStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("mirror_states") {
			request.MirrorStates = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListMirrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMirrors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagementStationMirrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagementStationMirrorsDataSource-", OsManagementHubManagementStationMirrorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managementStationMirror := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MirrorSummaryToMap(item))
	}
	managementStationMirror["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagementStationMirrorsDataSource().Schema["mirrors_collection"].Elem.(*schema.Resource).Schema)
		managementStationMirror["items"] = items
	}

	resources = append(resources, managementStationMirror)
	if err := s.D.Set("mirrors_collection", resources); err != nil {
		return err
	}

	return nil
}

func MirrorSummaryToMap(obj oci_os_management_hub.MirrorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["arch_type"] = string(obj.ArchType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Log != nil {
		result["log"] = string(*obj.Log)
	}

	result["os_family"] = string(obj.OsFamily)

	if obj.Percentage != nil {
		result["percentage"] = int(*obj.Percentage)
	}

	result["state"] = string(obj.State)

	if obj.TimeLastSynced != nil {
		result["time_last_synced"] = obj.TimeLastSynced.String()
	}

	result["type"] = string(obj.Type)

	return result
}
