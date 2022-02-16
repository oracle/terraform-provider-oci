// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v58/mysql"
)

func MysqlChannelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlChannels,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"channel_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"channels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(MysqlChannelResource()),
			},
		},
	}
}

func readMysqlChannels(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlChannelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ChannelsClient()

	return tfresource.ReadResource(sync)
}

type MysqlChannelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.ChannelsClient
	Res    *oci_mysql.ListChannelsResponse
}

func (s *MysqlChannelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlChannelsDataSourceCrud) Get() error {
	request := oci_mysql.ListChannelsRequest{}

	if channelId, ok := s.D.GetOkExists("id"); ok {
		tmp := channelId.(string)
		request.ChannelId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.ChannelLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.ListChannels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListChannels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MysqlChannelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlChannelsDataSource-", MysqlChannelsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mysqlChannel := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			mysqlChannel["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			mysqlChannel["display_name"] = *r.DisplayName
		}

		mysqlChannel["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			mysqlChannel["id"] = *r.Id
		}

		if r.IsEnabled != nil {
			mysqlChannel["is_enabled"] = *r.IsEnabled
		}

		if r.LifecycleDetails != nil {
			mysqlChannel["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Source != nil {
			sourceArray := []interface{}{}
			if sourceMap := ChannelSourceToMap(&r.Source); sourceMap != nil {
				sourceArray = append(sourceArray, sourceMap)
			}
			mysqlChannel["source"] = sourceArray
		} else {
			mysqlChannel["source"] = nil
		}

		mysqlChannel["state"] = r.LifecycleState

		if r.Target != nil {
			targetArray := []interface{}{}
			if targetMap := ChannelTargetToMap(&r.Target); targetMap != nil {
				targetArray = append(targetArray, targetMap)
			}
			mysqlChannel["target"] = targetArray
		} else {
			mysqlChannel["target"] = nil
		}

		if r.TimeCreated != nil {
			mysqlChannel["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			mysqlChannel["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, mysqlChannel)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MysqlChannelsDataSource().Schema["channels"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("channels", resources); err != nil {
		return err
	}

	return nil
}
