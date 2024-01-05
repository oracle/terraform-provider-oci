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

func FusionAppsFusionEnvironmentAdminUserDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fusion_environment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FusionAppsFusionEnvironmentAdminUserResource(), fieldMap, readSingularFusionAppsFusionEnvironmentAdminUser)
}

func readSingularFusionAppsFusionEnvironmentAdminUser(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentAdminUserDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentAdminUserDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListAdminUsersResponse
}

func (s *FusionAppsFusionEnvironmentAdminUserDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentAdminUserDataSourceCrud) Get() error {
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

func (s *FusionAppsFusionEnvironmentAdminUserDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentAdminUserDataSource-", FusionAppsFusionEnvironmentAdminUserDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdminUserSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
