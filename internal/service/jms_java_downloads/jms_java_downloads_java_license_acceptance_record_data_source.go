// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["java_license_acceptance_record_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsJavaDownloadsJavaLicenseAcceptanceRecordResource(), fieldMap, readSingularJmsJavaDownloadsJavaLicenseAcceptanceRecord)
}

func readSingularJmsJavaDownloadsJavaLicenseAcceptanceRecord(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.GetJavaLicenseAcceptanceRecordResponse
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.GetJavaLicenseAcceptanceRecordRequest{}

	if javaLicenseAcceptanceRecordId, ok := s.D.GetOkExists("java_license_acceptance_record_id"); ok {
		tmp := javaLicenseAcceptanceRecordId.(string)
		request.JavaLicenseAcceptanceRecordId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.GetJavaLicenseAcceptanceRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastUpdatedBy != nil {
		s.D.Set("last_updated_by", []interface{}{PrincipalToMap(s.Res.LastUpdatedBy)})
	} else {
		s.D.Set("last_updated_by", nil)
	}

	s.D.Set("license_acceptance_status", s.Res.LicenseAcceptanceStatus)

	s.D.Set("license_type", s.Res.LicenseType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeLastUpdated != nil {
		s.D.Set("time_last_updated", s.Res.TimeLastUpdated.String())
	}

	return nil
}
