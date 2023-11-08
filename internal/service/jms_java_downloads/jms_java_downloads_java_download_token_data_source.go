// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaDownloadTokenDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["java_download_token_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsJavaDownloadsJavaDownloadTokenResource(), fieldMap, readSingularJmsJavaDownloadsJavaDownloadToken)
}

func readSingularJmsJavaDownloadsJavaDownloadToken(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadTokenDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaDownloadTokenDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.GetJavaDownloadTokenResponse
}

func (s *JmsJavaDownloadsJavaDownloadTokenDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaDownloadTokenDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.GetJavaDownloadTokenRequest{}

	if javaDownloadTokenId, ok := s.D.GetOkExists("java_download_token_id"); ok {
		tmp := javaDownloadTokenId.(string)
		request.JavaDownloadTokenId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.GetJavaDownloadToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJavaDownloadsJavaDownloadTokenDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", []interface{}{PrincipalToMap(s.Res.CreatedBy)})
	} else {
		s.D.Set("created_by", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.JavaVersion != nil {
		s.D.Set("java_version", *s.Res.JavaVersion)
	}

	if s.Res.LastUpdatedBy != nil {
		s.D.Set("last_updated_by", []interface{}{PrincipalToMap(s.Res.LastUpdatedBy)})
	} else {
		s.D.Set("last_updated_by", nil)
	}

	licenseType := []interface{}{}
	for _, item := range s.Res.LicenseType {
		licenseType = append(licenseType, item)
	}
	s.D.Set("license_type", licenseType)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.Format(time.RFC3339Nano))
	}

	if s.Res.TimeLastUsed != nil {
		s.D.Set("time_last_used", s.Res.TimeLastUsed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Value != nil {
		s.D.Set("value", *s.Res.Value)
	}

	return nil
}
