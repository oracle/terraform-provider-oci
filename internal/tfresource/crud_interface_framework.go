// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"time"
)

// Common interfaces

type StatefullyCreatedResourceFw interface {
	StatefullyCreatedResource
	GetOperationTimeout() time.Duration
}

type StatefullyDeletedResourceFw interface {
	StatefullyDeletedResource
	GetOperationTimeout() time.Duration
}

type StatefullyUpdatedResourceFw interface {
	StatefullyUpdatedResource
	GetOperationTimeout() time.Duration
}
