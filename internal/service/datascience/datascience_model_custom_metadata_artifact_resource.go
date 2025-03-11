// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelCustomMetadataArtifactResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceModelCustomMetadataArtifact,
		Read:     readDatascienceModelCustomMetadataArtifact,
		Update:   updateDatascienceModelCustomMetadataArtifact,
		Delete:   deleteDatascienceModelCustomMetadataArtifact,
		Schema: map[string]*schema.Schema{
			// Required
			"model_custom_metadatum_artifact": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content_length": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"metadatum_key_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"model_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"content_disposition": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createDatascienceModelCustomMetadataArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelCustomMetadataArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceModelCustomMetadataArtifact(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatascienceModelCustomMetadataArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelCustomMetadataArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatascienceModelCustomMetadataArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelCustomMetadataArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceModelCustomMetadataArtifactResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	DisableNotFoundRetries bool
}

func (s *DatascienceModelCustomMetadataArtifactResourceCrud) ID() string {
	return "nil"
}

func (s *DatascienceModelCustomMetadataArtifactResourceCrud) Create() error {
	request := oci_datascience.CreateModelCustomMetadatumArtifactRequest{}

	if modelCustomMetadatumArtifact, ok := s.D.GetOkExists("model_custom_metadatum_artifact"); ok {
		tmp := []byte(modelCustomMetadatumArtifact.(string))
		request.ModelCustomMetadatumArtifact = ioutil.NopCloser(bytes.NewReader(tmp))
	}

	if contentDisposition, ok := s.D.GetOkExists("content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentLength, ok := s.D.GetOkExists("content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert content-length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if metadatumKeyName, ok := s.D.GetOkExists("metadatum_key_name"); ok {
		tmp := metadatumKeyName.(string)
		request.MetadatumKeyName = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.CreateModelCustomMetadatumArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatascienceModelCustomMetadataArtifactResourceCrud) Update() error {
	request := oci_datascience.UpdateModelCustomMetadatumArtifactRequest{}

	if modelCustomMetadatumArtifact, ok := s.D.GetOkExists("model_custom_metadatum_artifact"); ok {
		tmp := []byte(modelCustomMetadatumArtifact.(string))
		request.ModelCustomMetadatumArtifact = ioutil.NopCloser(bytes.NewReader(tmp))
	}

	if contentDisposition, ok := s.D.GetOkExists("content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentLength, ok := s.D.GetOkExists("content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert content-length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if metadatumKeyName, ok := s.D.GetOkExists("metadatum_key_name"); ok {
		tmp := metadatumKeyName.(string)
		request.MetadatumKeyName = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.UpdateModelCustomMetadatumArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatascienceModelCustomMetadataArtifactResourceCrud) Delete() error {
	request := oci_datascience.DeleteModelCustomMetadatumArtifactRequest{}

	if metadatumKeyName, ok := s.D.GetOkExists("metadatum_key_name"); ok {
		tmp := metadatumKeyName.(string)
		request.MetadatumKeyName = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeleteModelCustomMetadatumArtifact(context.Background(), request)
	return err
}

func (s *DatascienceModelCustomMetadataArtifactResourceCrud) SetData() error {
	return nil
}
