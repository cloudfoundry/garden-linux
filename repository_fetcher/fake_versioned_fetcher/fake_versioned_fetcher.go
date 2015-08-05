// This file was generated by counterfeiter
package fake_versioned_fetcher

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/repository_fetcher"
)

type FakeVersionedFetcher struct {
	FetchStub        func(*repository_fetcher.FetchRequest) (*repository_fetcher.FetchResponse, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		arg1 *repository_fetcher.FetchRequest
	}
	fetchReturns struct {
		result1 *repository_fetcher.FetchResponse
		result2 error
	}
}

func (fake *FakeVersionedFetcher) Fetch(arg1 *repository_fetcher.FetchRequest) (*repository_fetcher.FetchResponse, error) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		arg1 *repository_fetcher.FetchRequest
	}{arg1})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(arg1)
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2
	}
}

func (fake *FakeVersionedFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeVersionedFetcher) FetchArgsForCall(i int) *repository_fetcher.FetchRequest {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].arg1
}

func (fake *FakeVersionedFetcher) FetchReturns(result1 *repository_fetcher.FetchResponse, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 *repository_fetcher.FetchResponse
		result2 error
	}{result1, result2}
}

var _ repository_fetcher.VersionedFetcher = new(FakeVersionedFetcher)