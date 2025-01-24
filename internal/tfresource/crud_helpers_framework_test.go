package tfresource

import (
	"errors"
	"sync"
	"testing"
	"time"
)

type createResourceCrudFw struct {
	id                     string
	Mutex                  *sync.Mutex
	DisableNotFoundRetries bool
	SyncResource
	statefullyCreatedResource
	extraWaitPostCreateDelete
}

func (b createResourceCrudFw) Create() error {
	if b.id == "4" {
		return errors.New("")
	}
	return nil
}
func (b createResourceCrudFw) ID() string {
	return ""
}
func (b createResourceCrudFw) SetData() error {
	return nil
}
func (b createResourceCrudFw) VoidState() {}
func (b createResourceCrudFw) GetMutex() *sync.Mutex {
	return &sync.Mutex{}
}
func (b createResourceCrudFw) State() string {
	if b.id == "1" {
		return "SUCCEEDED"
	}
	return "FAILED"
}
func (b createResourceCrudFw) Get() error {
	if b.id == "2" {
		return errors.New("")
	}
	return nil
}
func (b createResourceCrudFw) CreatedPending() []string {
	return []string{"FAILED", "b"}
}
func (b createResourceCrudFw) CreatedTarget() []string {
	return []string{"FAILED", "b"}
}
func (b createResourceCrudFw) setState(s StatefulResource) error {
	if b.id == "3" {
		return errors.New("")
	}
	return nil
}
func (b createResourceCrudFw) ExtraWaitPostCreateDelete() time.Duration {
	timeoutDuration, _ := time.ParseDuration("1s")
	return timeoutDuration
}
func (b *createResourceCrudFw) GetOperationTimeout() time.Duration {
	return 1
}

type readResourceCrudFw struct {
	id                     string
	D                      *mockResourceData
	Mutex                  *sync.Mutex
	DisableNotFoundRetries bool
	SyncResource
	statefullyDeletedResource
}

func (b readResourceCrudFw) Create() error {
	return nil
}
func (b readResourceCrudFw) ID() string {
	return ""
}
func (b readResourceCrudFw) SetData() error {
	return nil
}
func (b readResourceCrudFw) VoidState() {}
func (b readResourceCrudFw) State() string {
	if b.id == "1" {
		return "a"
	}
	return "FAILED"
}
func (b readResourceCrudFw) Get() error {
	return nil
}
func (b readResourceCrudFw) DeletedPending() []string {
	return []string{"a", "b"}
}
func (b readResourceCrudFw) DeletedTarget() []string {
	return []string{"a", "b"}
}
func (b *readResourceCrudFw) GetOperationTimeout() time.Duration {
	return 1
}

type updateResourceCrudFw struct {
	D                      *mockResourceData
	Mutex                  *sync.Mutex
	DisableNotFoundRetries bool
	SyncResource
	statefullyUpdatedResource
}

func (b updateResourceCrudFw) Update() error {
	return nil
}
func (b updateResourceCrudFw) ID() string {
	return ""
}
func (b updateResourceCrudFw) SetData() error {
	return nil
}
func (b updateResourceCrudFw) VoidState() {}
func (b updateResourceCrudFw) State() string {
	return "FAILED"
}
func (b updateResourceCrudFw) Get() error {
	return nil
}
func (b updateResourceCrudFw) GetMutex() *sync.Mutex {
	return &sync.Mutex{}
}
func (b updateResourceCrudFw) UpdatedPending() []string {
	return []string{"a", "b"}
}
func (b updateResourceCrudFw) UpdatedTarget() []string {
	return []string{"a", "b"}
}
func (b *updateResourceCrudFw) GetOperationTimeout() time.Duration {
	return 1
}

type deleteResourceCrudFw struct {
	D                      *mockResourceData
	Mutex                  *sync.Mutex
	DisableNotFoundRetries bool
	SyncResource
	statefullyDeletedResource
	extraWaitPostCreateDelete
	extraWaitPostDelete
}

func (b deleteResourceCrudFw) Delete() error {
	return nil
}
func (b deleteResourceCrudFw) ID() string {
	return ""
}
func (b deleteResourceCrudFw) SetData() error {
	return nil
}
func (b deleteResourceCrudFw) VoidState() {}
func (b deleteResourceCrudFw) State() string {
	return "FAILED"
}
func (b deleteResourceCrudFw) Get() error {
	return nil
}
func (b deleteResourceCrudFw) GetMutex() *sync.Mutex {
	return &sync.Mutex{}
}
func (b deleteResourceCrudFw) DeletedPending() []string {
	return []string{"a", "b"}
}
func (b deleteResourceCrudFw) DeletedTarget() []string {
	return []string{"a", "b"}
}
func (b deleteResourceCrudFw) ExtraWaitPostCreateDelete() time.Duration {
	timeoutDuration, _ := time.ParseDuration("1s")
	return timeoutDuration
}
func (b deleteResourceCrudFw) ExtraWaitPostDelete() time.Duration {
	timeoutDuration, _ := time.ParseDuration("1s")
	return timeoutDuration
}
func (b *deleteResourceCrudFw) GetOperationTimeout() time.Duration {
	return 1
}
func TestUnitCreateResourceFw(t *testing.T) {
	s1 := &createResourceCrudFw{}
	s2 := &createResourceCrudFw{id: "4"}
	s3 := &createResourceCrudFw{id: "2"}

	type args struct {
		sync ResourceCreator
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
		mockFunc func()
	}
	tests := []testFormat{
		{
			name:     "Test error is returned from create",
			args:     args{sync: s2},
			gotError: true,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return nil
				}
			},
		},
		{
			name:     "Test error is returned from waitForStateRefreshVar and state failed",
			args:     args{sync: s3},
			gotError: true,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return errors.New("default")
				}
			},
		},
		{
			name:     "Test no error is returned",
			args:     args{sync: s1},
			gotError: false,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return nil
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if res := CreateResourceFw(test.args.sync); (res != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", res, test.gotError)
		}
	}
}

func TestUnitReadResourceFw(t *testing.T) {
	s := &readResourceCrudFw{}
	s1 := &createResourceCrudFw{id: "1"}

	type args struct {
		sync ResourceReader
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
	}
	tests := []testFormat{
		{
			name:     "Test",
			args:     args{sync: s},
			gotError: false,
		},
		{
			name:     "Test resource already deleted",
			args:     args{sync: s1},
			gotError: false,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if res := ReadResourceFw(test.args.sync); (res != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", res, test.gotError)
		}
	}
}

func TestUnitUpdateResourceFw(t *testing.T) {
	s := &updateResourceCrudFw{}

	type args struct {
		sync ResourceUpdater
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
		mockFunc func()
	}
	tests := []testFormat{
		{
			name:     "Test error in waitForStateRefreshVar",
			args:     args{sync: s},
			gotError: true,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return errors.New("default")
				}
			},
		},
		{
			name:     "Test",
			args:     args{sync: s},
			gotError: false,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return nil
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if res := UpdateResourceFw(test.args.sync); (res != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", res, test.gotError)
		}
	}
}

func TestUnitDeleteResourceFw(t *testing.T) {
	s := &deleteResourceCrudFw{}

	type args struct {
		sync ResourceDeleter
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
		mockFunc func()
	}
	tests := []testFormat{
		{
			name:     "Test",
			args:     args{sync: s},
			gotError: true,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return errors.New("default")
				}
			},
		},
		{
			name:     "Test",
			args:     args{sync: s},
			gotError: false,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return nil
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if res := DeleteResourceFw(test.args.sync); (res != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", res, test.gotError)
		}
	}
}

func TestUnitWaitForUpdatedStateFw(t *testing.T) {
	s := &updateResourceCrudFw{}

	type args struct {
		sync ResourceUpdater
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
		mockFunc func()
	}
	tests := []testFormat{
		{
			name:     "Test error is returned",
			args:     args{sync: s},
			gotError: true,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return errors.New("default")
				}
			},
		},
		{
			name:     "Test no error is returned",
			args:     args{sync: s},
			gotError: false,
			mockFunc: func() {
				waitForStateRefreshVar = func(sr StatefulResource, timeout time.Duration, operationName string, pending []string, target []string) error {
					return nil
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if res := WaitForUpdatedStateFw(test.args.sync); (res != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", res, test.gotError)
		}
	}
}
