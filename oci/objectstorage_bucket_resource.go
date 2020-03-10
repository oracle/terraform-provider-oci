// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func init() {
	RegisterResource("oci_objectstorage_bucket", ObjectStorageBucketResource())
}

func ObjectStorageBucketResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createObjectStorageBucket,
		Read:     readObjectStorageBucket,
		Update:   updateObjectStorageBucket,
		Delete:   deleteObjectStorageBucket,
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
			"object_events_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
			"bucket_id": {
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
			"is_read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"object_lifecycle_policy_etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return CreateResource(d, sync)
}

func readObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

func updateObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return UpdateResource(d, sync)
}

func deleteObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ObjectStorageBucketResourceCrud struct {
	BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.Bucket
	DisableNotFoundRetries bool
}

func (s *ObjectStorageBucketResourceCrud) ID() string {

	return getBucketCompositeId(s.D.Get("name").(string), s.D.Get("namespace").(string))
}

func (s *ObjectStorageBucketResourceCrud) Create() error {
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

	if objectEventsEnabled, ok := s.D.GetOkExists("object_events_enabled"); ok {
		tmp := objectEventsEnabled.(bool)
		request.ObjectEventsEnabled = &tmp
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

func (s *ObjectStorageBucketResourceCrud) Get() error {
	request := oci_object_storage.GetBucketRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	bucket, namespace, err := parseBucketCompositeId(s.D.Id())
	if err == nil {
		request.BucketName = &bucket
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
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

func (s *ObjectStorageBucketResourceCrud) Update() error {
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

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if objectEventsEnabled, ok := s.D.GetOkExists("object_events_enabled"); ok {
		tmp := objectEventsEnabled.(bool)
		request.ObjectEventsEnabled = &tmp
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

func (s *ObjectStorageBucketResourceCrud) Delete() error {
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

func (s *ObjectStorageBucketResourceCrud) SetData() error {

	s.D.Set("bucket_id", *s.Res.Id)

	bucket, namespace, err := parseBucketCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bucket", &bucket)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

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

	if s.Res.IsReadOnly != nil {
		s.D.Set("is_read_only", *s.Res.IsReadOnly)
	}

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

	if s.Res.ObjectEventsEnabled != nil {
		s.D.Set("object_events_enabled", *s.Res.ObjectEventsEnabled)
	}

	if s.Res.ObjectLifecyclePolicyEtag != nil {
		s.D.Set("object_lifecycle_policy_etag", *s.Res.ObjectLifecyclePolicyEtag)
	}

	if s.Res.ReplicationEnabled != nil {
		s.D.Set("replication_enabled", *s.Res.ReplicationEnabled)
	}

	s.D.Set("storage_tier", s.Res.StorageTier)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func getBucketCompositeId(bucket string, namespace string) string {
	bucket = url.PathEscape(bucket)
	namespace = url.PathEscape(namespace)
	compositeId := "n/" + namespace + "/b/" + bucket
	return compositeId
}

func parseBucketCompositeId(compositeId string) (bucket string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("n/.*/b/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	bucket, _ = url.PathUnescape(parts[3])

	return
}
