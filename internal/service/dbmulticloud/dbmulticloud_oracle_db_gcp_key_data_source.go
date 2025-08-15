// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbGcpKeyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDbmulticloudOracleDbGcpKey,
		Schema: map[string]*schema.Schema{
			"oracle_db_gcp_key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"gcp_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gcp_key_properties": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_db_gcp_key_ring_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDbmulticloudOracleDbGcpKey(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbGcpKeyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudGCPProviderClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbGcpKeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.DbMulticloudGCPProviderClient
	Res    *oci_dbmulticloud.GetOracleDbGcpKeyResponse
}

func (s *DbmulticloudOracleDbGcpKeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbGcpKeyDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbGcpKeyRequest{}

	if oracleDbGcpKeyId, ok := s.D.GetOkExists("oracle_db_gcp_key_id"); ok {
		tmp := oracleDbGcpKeyId.(string)
		request.OracleDbGcpKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbGcpKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbGcpKeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GcpKeyId != nil {
		s.D.Set("gcp_key_id", *s.Res.GcpKeyId)
	}

	if s.Res.GcpKeyProperties != nil {
		// Dereference
		keyPropsInterface := *s.Res.GcpKeyProperties

		// Assert it to a map[string]interface{}
		keyPropsMap, _ := keyPropsInterface.(map[string]interface{})

		props := map[string]interface{}{}
		for k, v := range keyPropsMap {
			props[k] = v
		}
		s.D.Set("gcp_key_properties", props)
	} else {
		s.D.Set("gcp_key_properties", nil)
	}

	//if s.Res.GcpKeyProperties != nil {
	//	s.D.Set("gcp_key_properties", []interface{}{s.Res.GcpKeyProperties})
	//} else {
	//	s.D.Set("gcp_key_properties", nil)
	//}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.OracleDbGcpKeyRingId != nil {
		s.D.Set("oracle_db_gcp_key_ring_id", *s.Res.OracleDbGcpKeyRingId)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
