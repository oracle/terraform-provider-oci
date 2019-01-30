// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func BucketResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createBucket,
		Read:     readBucket,
		Update:   updateBucket,
		Delete:   deleteBucket,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"access_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  string(oci_object_storage.CreateBucketDetailsPublicAccessTypeNopublicaccess),
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"storage_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"approximate_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"approximate_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_lifecycle_policy_etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBucket(d *schema.ResourceData, m interface{}) error {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return CreateResource(d, sync)
}

func readBucket(d *schema.ResourceData, m interface{}) error {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

func updateBucket(d *schema.ResourceData, m interface{}) error {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return UpdateResource(d, sync)
}

func deleteBucket(d *schema.ResourceData, m interface{}) error {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type BucketResourceCrud struct {
	BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.Bucket
	DisableNotFoundRetries bool
}

func (s *BucketResourceCrud) ID() string {
	if s.Res.Namespace == nil || s.Res.Name == nil {
		log.Printf("Could not get ID for bucket. The bucket namespace and/or name is nil")
	}

	return *s.Res.Namespace + "/" + *s.Res.Name
}

func (s *BucketResourceCrud) Create() error {
	request := oci_object_storage.CreateBucketRequest{}

	if accessType, ok := s.D.GetOkExists("access_type"); ok {
		request.PublicAccessType = oci_object_storage.CreateBucketDetailsPublicAccessTypeEnum(accessType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = resourceObjectStorageMapToMetadata(metadata.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if storageTier, ok := s.D.GetOkExists("storage_tier"); ok {
		request.StorageTier = oci_object_storage.CreateBucketDetailsStorageTierEnum(storageTier.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.CreateBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Bucket
	return nil
}

func (s *BucketResourceCrud) Get() error {
	request := oci_object_storage.GetBucketRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.Fields = oci_object_storage.GetGetBucketFieldsEnumValues()
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Bucket
	return nil
}

func (s *BucketResourceCrud) Update() error {
	request := oci_object_storage.UpdateBucketRequest{}

	if accessType, ok := s.D.GetOkExists("access_type"); ok {
		request.PublicAccessType = oci_object_storage.UpdateBucketDetailsPublicAccessTypeEnum(accessType.(string))
	}

	if bucket, ok := s.D.GetOkExists("name"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = resourceObjectStorageMapToMetadata(metadata.(map[string]interface{}))
	}

	// @CODEGEN 2/2018: This should be used to change the name of a bucket, but the "name" field
	// is already being used to identify the bucket. Should have a new field for this.
	// Existing provider omits this, so we will omit it for now to avoid a potential breaking change.
	//if name, ok := s.D.GetOkExists("name"); ok {
	//	tmp := name.(string)
	//	request.Name = &tmp
	//}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	// @CODEGEN 2/2018: This should be used to change the name of a bucket, but the "namespace" field
	// is already being used to identify the bucket. Should have a new field for this.
	// Existing provider omits this, so we will omit it for now to avoid a potential breaking change.
	//if namespace, ok := s.D.GetOkExists("namespace"); ok {
	//	tmp := namespace.(string)
	//	request.Namespace = &tmp
	//}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.UpdateBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Bucket
	return nil
}

func (s *BucketResourceCrud) Delete() error {
	request := oci_object_storage.DeleteBucketRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err := s.Client.DeleteBucket(context.Background(), request)
	return err
}

func (s *BucketResourceCrud) SetData() error {
	s.D.Set("access_type", s.Res.PublicAccessType)

	if s.Res.ApproximateCount != nil {
		s.D.Set("approximate_count", strconv.FormatInt(*s.Res.ApproximateCount, 10))
	}

	if s.Res.ApproximateSize != nil {
		s.D.Set("approximate_size", strconv.FormatInt(*s.Res.ApproximateSize, 10))
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", s.Res.Metadata)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ObjectLifecyclePolicyEtag != nil {
		s.D.Set("object_lifecycle_policy_etag", *s.Res.ObjectLifecyclePolicyEtag)
	}

	s.D.Set("storage_tier", s.Res.StorageTier)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
