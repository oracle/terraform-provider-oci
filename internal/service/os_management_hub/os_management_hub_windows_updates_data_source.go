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

func OsManagementHubWindowsUpdatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubWindowsUpdates,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"classification_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"windows_update_collection": {
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
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"installable": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"installation_requirements": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_reboot_required_for_installation": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"kb_article_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_in_bytes": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"update_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"update_type": {
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

func readOsManagementHubWindowsUpdates(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubWindowsUpdatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubWindowsUpdatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceClient
	Res    *oci_os_management_hub.ListWindowsUpdatesResponse
}

func (s *OsManagementHubWindowsUpdatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubWindowsUpdatesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListWindowsUpdatesRequest{}

	if classificationType, ok := s.D.GetOkExists("classification_type"); ok {
		interfaces := classificationType.([]interface{})
		tmp := make([]oci_os_management_hub.ClassificationTypesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ClassificationTypesEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("classification_type") {
			request.ClassificationType = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		interfaces := name.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("name") {
			request.Name = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListWindowsUpdates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWindowsUpdates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubWindowsUpdatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubWindowsUpdatesDataSource-", OsManagementHubWindowsUpdatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	windowsUpdate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WindowsUpdateSummaryToMap(item))
	}
	windowsUpdate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubWindowsUpdatesDataSource().Schema["windows_update_collection"].Elem.(*schema.Resource).Schema)
		windowsUpdate["items"] = items
	}

	resources = append(resources, windowsUpdate)
	if err := s.D.Set("windows_update_collection", resources); err != nil {
		return err
	}

	return nil
}

func WindowsUpdateSummaryToMap(obj oci_os_management_hub.WindowsUpdateSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["installable"] = string(obj.Installable)

	if obj.IsRebootRequiredForInstallation != nil {
		result["is_reboot_required_for_installation"] = bool(*obj.IsRebootRequiredForInstallation)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.UpdateId != nil {
		result["update_id"] = string(*obj.UpdateId)
	}

	result["update_type"] = string(obj.UpdateType)

	return result
}
