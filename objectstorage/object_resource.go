// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package objectstorage

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func ObjectResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: crud.DefaultTimeout,
		Create:   createObject,
		Read:     readObject,
		Update:   updateObject,
		Delete:   deleteObject,
		Schema:   objectSchema,
	}
}

func createObject(d *schema.ResourceData, m interface{}) (e error) {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readObject(d *schema.ResourceData, m interface{}) (e error) {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func updateObject(d *schema.ResourceData, m interface{}) (e error) {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.UpdateResource(d, sync)
}

func deleteObject(d *schema.ResourceData, m interface{}) (e error) {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}
