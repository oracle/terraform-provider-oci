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

func WlmsWlsDomainServersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsWlsDomainServers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"wls_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_collection": {
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
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_admin": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"jdk_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"jdk_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"latest_patches_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"middleware_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"middleware_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patch_readiness_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"restart_order": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"status": {
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
									"weblogic_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"wls_domain_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"wls_domain_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"wls_domain_path": {
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

func readWlmsWlsDomainServers(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainServersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainServersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListWlsDomainServersResponse
}

func (s *WlmsWlsDomainServersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainServersDataSourceCrud) Get() error {
	request := oci_wlms.ListWlsDomainServersRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListWlsDomainServers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWlsDomainServers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsWlsDomainServersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsWlsDomainServersDataSource-", WlmsWlsDomainServersDataSource(), s.D))
	resources := []map[string]interface{}{}
	wlsDomainServer := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WlsDomainServerSummaryToMap(item))
	}
	wlsDomainServer["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsWlsDomainServersDataSource().Schema["server_collection"].Elem.(*schema.Resource).Schema)
		wlsDomainServer["items"] = items
	}

	resources = append(resources, wlsDomainServer)
	if err := s.D.Set("server_collection", resources); err != nil {
		return err
	}

	return nil
}

func WlsDomainServerSummaryToMap(obj oci_wlms.ServerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAdmin != nil {
		result["is_admin"] = bool(*obj.IsAdmin)
	}

	if obj.JdkPath != nil {
		result["jdk_path"] = string(*obj.JdkPath)
	}

	if obj.JdkVersion != nil {
		result["jdk_version"] = string(*obj.JdkVersion)
	}

	result["latest_patches_status"] = string(obj.LatestPatchesStatus)

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.MiddlewarePath != nil {
		result["middleware_path"] = string(*obj.MiddlewarePath)
	}

	if obj.MiddlewareType != nil {
		result["middleware_type"] = string(*obj.MiddlewareType)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["patch_readiness_status"] = string(obj.PatchReadinessStatus)

	if obj.RestartOrder != nil {
		result["restart_order"] = int(*obj.RestartOrder)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.WeblogicVersion != nil {
		result["weblogic_version"] = string(*obj.WeblogicVersion)
	}

	if obj.WlsDomainId != nil {
		result["wls_domain_id"] = string(*obj.WlsDomainId)
	}

	if obj.WlsDomainName != nil {
		result["wls_domain_name"] = string(*obj.WlsDomainName)
	}

	if obj.WlsDomainPath != nil {
		result["wls_domain_path"] = string(*obj.WlsDomainPath)
	}

	return result
}
