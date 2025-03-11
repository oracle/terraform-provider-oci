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

func DatascienceModelDefinedMetadataArtifactResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceModelDefinedMetadataArtifact,
		Read:     readDatascienceModelDefinedMetadataArtifact,
		Update:   updateDatascienceModelDefinedMetadataArtifact,
		Delete:   deleteDatascienceModelDefinedMetadataArtifact,
		Schema: map[string]*schema.Schema{
			// Required
			"model_defined_metadatum_artifact": {
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

func createDatascienceModelDefinedMetadataArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDefinedMetadataArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceModelDefinedMetadataArtifact(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatascienceModelDefinedMetadataArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDefinedMetadataArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatascienceModelDefinedMetadataArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDefinedMetadataArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceModelDefinedMetadataArtifactResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	DisableNotFoundRetries bool
}

func (s *DatascienceModelDefinedMetadataArtifactResourceCrud) ID() string {
	return "nil"
}

func (s *DatascienceModelDefinedMetadataArtifactResourceCrud) Create() error {
	request := oci_datascience.CreateModelDefinedMetadatumArtifactRequest{}

	if modelDefinedMetadatumArtifact, ok := s.D.GetOkExists("model_defined_metadatum_artifact"); ok {
		tmp := []byte(modelDefinedMetadatumArtifact.(string))
		request.ModelDefinedMetadatumArtifact = ioutil.NopCloser(bytes.NewReader(tmp))
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

	_, err := s.Client.CreateModelDefinedMetadatumArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatascienceModelDefinedMetadataArtifactResourceCrud) Update() error {
	request := oci_datascience.UpdateModelDefinedMetadatumArtifactRequest{}

	if modelDefinedMetadatumArtifact, ok := s.D.GetOkExists("model_defined_metadatum_artifact"); ok {
		tmp := []byte(modelDefinedMetadatumArtifact.(string))
		request.ModelDefinedMetadatumArtifact = ioutil.NopCloser(bytes.NewReader(tmp))
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

	_, err := s.Client.UpdateModelDefinedMetadatumArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatascienceModelDefinedMetadataArtifactResourceCrud) Delete() error {
	request := oci_datascience.DeleteModelDefinedMetadatumArtifactRequest{}

	if metadatumKeyName, ok := s.D.GetOkExists("metadatum_key_name"); ok {
		tmp := metadatumKeyName.(string)
		request.MetadatumKeyName = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeleteModelDefinedMetadatumArtifact(context.Background(), request)
	return err
}

func (s *DatascienceModelDefinedMetadataArtifactResourceCrud) SetData() error {
	return nil
}
