// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageOutboundConnectorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["outbound_connector_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FileStorageOutboundConnectorResource(), fieldMap, readSingularFileStorageOutboundConnector)
}

func readSingularFileStorageOutboundConnector(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageOutboundConnectorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageOutboundConnectorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.GetOutboundConnectorResponse
}

func (s *FileStorageOutboundConnectorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageOutboundConnectorDataSourceCrud) Get() error {
	request := oci_file_storage.GetOutboundConnectorRequest{}

	if outboundConnectorId, ok := s.D.GetOkExists("outbound_connector_id"); ok {
		tmp := outboundConnectorId.(string)
		request.OutboundConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.GetOutboundConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FileStorageOutboundConnectorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.OutboundConnector).(type) {
	case oci_file_storage.LdapBindAccount:
		s.D.Set("connector_type", "LDAPBIND")

		if v.BindDistinguishedName != nil {
			s.D.Set("bind_distinguished_name", *v.BindDistinguishedName)
		}

		endpoints := []interface{}{}
		for _, item := range v.Endpoints {
			endpoints = append(endpoints, EndpointToMap(item))
		}
		s.D.Set("endpoints", endpoints)

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PasswordSecretVersion != nil {
			s.D.Set("password_secret_version", *v.PasswordSecretVersion)
		}

		if v.AvailabilityDomain != nil {
			s.D.Set("availability_domain", *v.AvailabilityDomain)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}
	default:
		log.Printf("[WARN] Received 'connector_type' of unknown type %v", s.Res.OutboundConnector)
		return nil
	}

	return nil
}
