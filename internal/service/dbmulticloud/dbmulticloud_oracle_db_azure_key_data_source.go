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

func DbmulticloudOracleDbAzureKeyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDbmulticloudOracleDbAzureKey,
		Schema: map[string]*schema.Schema{
			"oracle_db_azure_key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"azure_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"key_properties": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"last_modification": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_db_azure_vault_id": {
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

func readSingularDbmulticloudOracleDbAzureKey(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureKeyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDbAzureKeyClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureKeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDbAzureKeyClient
	Res    *oci_dbmulticloud.GetOracleDbAzureKeyResponse
}

func (s *DbmulticloudOracleDbAzureKeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureKeyDataSourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbAzureKeyRequest{}

	if oracleDbAzureKeyId, ok := s.D.GetOkExists("oracle_db_azure_key_id"); ok {
		tmp := oracleDbAzureKeyId.(string)
		request.OracleDbAzureKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.GetOracleDbAzureKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbmulticloudOracleDbAzureKeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AzureKeyId != nil {
		s.D.Set("azure_key_id", *s.Res.AzureKeyId)
	}

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

	if s.Res.KeyProperties != nil {
		// Dereference
		keyPropsInterface := *s.Res.KeyProperties

		// Assert it to a map[string]interface{}
		keyPropsMap, _ := keyPropsInterface.(map[string]interface{})

		props := map[string]interface{}{}
		for k, v := range keyPropsMap {
			props[k] = v
		}
		s.D.Set("key_properties", props)
	} else {
		s.D.Set("key_properties", nil)
	}

	if s.Res.LastModification != nil {
		s.D.Set("last_modification", *s.Res.LastModification)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.OracleDbAzureVaultId != nil {
		s.D.Set("oracle_db_azure_vault_id", *s.Res.OracleDbAzureVaultId)
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
