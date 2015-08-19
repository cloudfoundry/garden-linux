package debug_test

import (
	"net/http"

	"github.com/cloudfoundry-incubator/garden-linux/debug"
	"github.com/pivotal-golang/lager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Debug", func() {
	It("reports an expvar debug metric on /debug path", func() {
		sink := lager.NewReconfigurableSink(lager.NewWriterSink(GinkgoWriter, lager.DEBUG), lager.DEBUG)
		_, err := debug.Run("127.0.0.1:5123", sink)
		Expect(err).ToNot(HaveOccurred())

		resp, err := http.Get("http://127.0.0.1:5123/debug/vars")
		Expect(err).ToNot(HaveOccurred())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})
})