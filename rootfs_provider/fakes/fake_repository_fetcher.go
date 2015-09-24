// This file was generated by counterfeiter
package fakes

import (
	"net/url"
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/repository_fetcher"
	"github.com/cloudfoundry-incubator/garden-linux/rootfs_provider"
)

type FakeRepositoryFetcher struct {
	FetchStub        func(*url.URL, int64) (*repository_fetcher.Image, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		arg1 *url.URL
		arg2 int64
	}
	fetchReturns struct {
		result1 *repository_fetcher.Image
		result2 error
	}
}

func (fake *FakeRepositoryFetcher) Fetch(arg1 *url.URL, arg2 int64) (*repository_fetcher.Image, error) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		arg1 *url.URL
		arg2 int64
	}{arg1, arg2})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(arg1, arg2)
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2
	}
}

func (fake *FakeRepositoryFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeRepositoryFetcher) FetchArgsForCall(i int) (*url.URL, int64) {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].arg1, fake.fetchArgsForCall[i].arg2
}

func (fake *FakeRepositoryFetcher) FetchReturns(result1 *repository_fetcher.Image, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 *repository_fetcher.Image
		result2 error
	}{result1, result2}
}

var _ rootfs_provider.RepositoryFetcher = new(FakeRepositoryFetcher)
