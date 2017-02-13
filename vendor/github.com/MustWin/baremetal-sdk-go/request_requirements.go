// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

type identityCreationRequirement struct {
	CompartmentID string `header:"-" json:"compartmentId" url:"-"`
	Description   string `header:"-" json:"description" url:"-"`
	Name          string `header:"-" json:"name" url:"-"`
}

type ocidRequirement struct {
	CompartmentID string `header:"-" json:"compartmentId" url:"compartmentId"`
}

type listOCIDRequirement struct {
	CompartmentID string `header:"-" json:"-" url:"compartmentId"`
}

// Body is handled explicitly during marshal.
type bodyMarshaller interface {
	body() []byte
}

type bodyRequirement struct {
	Body []byte `header:"-" json:"-" url:"-"`
}

func (b bodyRequirement) body() []byte {
	return b.Body
}
