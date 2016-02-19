// This file was generated by counterfeiter
package fake_rootfs_cleaner

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/resource_pool"
	"github.com/pivotal-golang/lager"
)

type FakeRootFSCleaner struct {
	CleanStub        func(log lager.Logger, path string) error
	cleanMutex       sync.RWMutex
	cleanArgsForCall []struct {
		log  lager.Logger
		path string
	}
	cleanReturns struct {
		result1 error
	}
}

func (fake *FakeRootFSCleaner) Clean(log lager.Logger, path string) error {
	fake.cleanMutex.Lock()
	fake.cleanArgsForCall = append(fake.cleanArgsForCall, struct {
		log  lager.Logger
		path string
	}{log, path})
	fake.cleanMutex.Unlock()
	if fake.CleanStub != nil {
		return fake.CleanStub(log, path)
	} else {
		return fake.cleanReturns.result1
	}
}

func (fake *FakeRootFSCleaner) CleanCallCount() int {
	fake.cleanMutex.RLock()
	defer fake.cleanMutex.RUnlock()
	return len(fake.cleanArgsForCall)
}

func (fake *FakeRootFSCleaner) CleanArgsForCall(i int) (lager.Logger, string) {
	fake.cleanMutex.RLock()
	defer fake.cleanMutex.RUnlock()
	return fake.cleanArgsForCall[i].log, fake.cleanArgsForCall[i].path
}

func (fake *FakeRootFSCleaner) CleanReturns(result1 error) {
	fake.CleanStub = nil
	fake.cleanReturns = struct {
		result1 error
	}{result1}
}

var _ resource_pool.RootFSCleaner = new(FakeRootFSCleaner)