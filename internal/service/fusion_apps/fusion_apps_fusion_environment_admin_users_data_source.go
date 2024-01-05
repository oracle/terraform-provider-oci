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

func FusionAppsFusionEnvironmentAdminUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFusionAppsFusionEnvironmentAdminUsers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"admin_user_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FusionAppsFusionEnvironmentAdminUserResource()),
						},
					},
				},
			},
		},
	}
}

func readFusionAppsFusionEnvironmentAdminUsers(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentAdminUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentAdminUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListAdminUsersResponse
}

func (s *FusionAppsFusionEnvironmentAdminUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentAdminUsersDataSourceCrud) Get() error {
	request := oci_fusion_apps.ListAdminUsersRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.ListAdminUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentAdminUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentAdminUsersDataSource-", FusionAppsFusionEnvironmentAdminUsersDataSource(), s.D))
	resources := []map[string]interface{}{}
	fusionEnvironmentAdminUser := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdminUserSummaryToMap(item))
	}
	fusionEnvironmentAdminUser["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FusionAppsFusionEnvironmentAdminUsersDataSource().Schema["admin_user_collection"].Elem.(*schema.Resource).Schema)
		fusionEnvironmentAdminUser["items"] = items
	}

	resources = append(resources, fusionEnvironmentAdminUser)
	if err := s.D.Set("admin_user_collection", resources); err != nil {
		return err
	}

	return nil
}
