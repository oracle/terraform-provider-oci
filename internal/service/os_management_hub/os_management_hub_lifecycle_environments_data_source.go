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

func OsManagementHubLifecycleEnvironmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubLifecycleEnvironments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"arch_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_environment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_environment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OsManagementHubLifecycleEnvironmentResource()),
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubLifecycleEnvironments(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleEnvironmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubLifecycleEnvironmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.LifecycleEnvironmentClient
	Res    *oci_os_management_hub.ListLifecycleEnvironmentsResponse
}

func (s *OsManagementHubLifecycleEnvironmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubLifecycleEnvironmentsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListLifecycleEnvironmentsRequest{}

	if archType, ok := s.D.GetOkExists("arch_type"); ok {
		request.ArchType = oci_os_management_hub.ListLifecycleEnvironmentsArchTypeEnum(archType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		interfaces := displayName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_name") {
			request.DisplayName = tmp
		}
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if lifecycleEnvironmentId, ok := s.D.GetOkExists("id"); ok {
		tmp := lifecycleEnvironmentId.(string)
		request.LifecycleEnvironmentId = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_os_management_hub.ListLifecycleEnvironmentsOsFamilyEnum(osFamily.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_os_management_hub.LifecycleEnvironmentLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListLifecycleEnvironments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLifecycleEnvironments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubLifecycleEnvironmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubLifecycleEnvironmentsDataSource-", OsManagementHubLifecycleEnvironmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	lifecycleEnvironment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LifecycleEnvironmentSummaryToMap(item))
	}
	lifecycleEnvironment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubLifecycleEnvironmentsDataSource().Schema["lifecycle_environment_collection"].Elem.(*schema.Resource).Schema)
		lifecycleEnvironment["items"] = items
	}

	resources = append(resources, lifecycleEnvironment)
	if err := s.D.Set("lifecycle_environment_collection", resources); err != nil {
		return err
	}

	return nil
}
