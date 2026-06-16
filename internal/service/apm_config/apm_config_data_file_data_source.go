// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_config

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmConfigDataFileDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["apm_type"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["data_file_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["base64_encode_content"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(ApmConfigDataFileResource(), fieldMap, readSingularApmConfigDataFileWithContext)
}

func readSingularApmConfigDataFileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ApmConfigDataFileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ApmConfigDataFileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_config.ConfigClient
	Res    *oci_apm_config.GetDataFileResponse
}

func (s *ApmConfigDataFileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmConfigDataFileDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_apm_config.GetDataFileRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if apmType, ok := s.D.GetOkExists("apm_type"); ok {
		tmp := apmType.(string)
		request.ApmType = &tmp
	}

	if dataFileName, ok := s.D.GetOkExists("data_file_name"); ok {
		tmp := dataFileName.(string)
		request.DataFileName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_config")

	response, err := s.Client.GetDataFile(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmConfigDataFileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	base64EncodeContent := false
	if tmp, ok := s.D.GetOkExists("base64_encode_content"); ok {
		base64EncodeContent = tmp.(bool)
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmConfigDataFileDataSource-", ApmConfigDataFileDataSource(), s.D))

	contentReader := s.Res.Content
	contentArray, err := ioutil.ReadAll(contentReader)
	if err != nil {
		log.Printf("unable to read 'content' from response. Error: %v", err)
	} else if base64EncodeContent {
		s.D.Set("content", base64.StdEncoding.EncodeToString(contentArray))
	} else {
		s.D.Set("content", string(contentArray))
	}

	if s.Res.ContentType != nil {
		s.D.Set("content_type", *s.Res.ContentType)
	}

	if s.Res.ContentLength != nil {
		s.D.Set("content_length", *s.Res.ContentLength)
	}

	if s.Res.ContentMd5 != nil {
		s.D.Set("content_md5", *s.Res.ContentMd5)
	}

	if s.Res.ContentEncoding != nil {
		s.D.Set("content_encoding", *s.Res.ContentEncoding)
	}

	if s.Res.ContentLanguage != nil {
		s.D.Set("content_language", *s.Res.ContentLanguage)
	}

	if s.Res.ContentDisposition != nil {
		s.D.Set("content_disposition", *s.Res.ContentDisposition)
	}

	if s.Res.LastModified != nil {
		s.D.Set("time_last_modified", s.Res.LastModified.String())
	}

	if s.Res.Metadata != nil {
		metadata, err := jsonStringToMetaDataStringMap(*s.Res.Metadata)
		if err != nil {
			return fmt.Errorf("could not set metadata map: %q", err)
		}
		s.D.Set("metadata", metadata)
	}

	return nil
}
