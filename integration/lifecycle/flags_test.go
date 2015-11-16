package lifecycle_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudfoundry-incubator/garden"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Garden startup flags", func() {

	var debugAddr string

	BeforeEach(func() {
		debugAddr = fmt.Sprintf("0.0.0.0:%d", 15000+GinkgoParallelNode())
	})

	Context("when starting without the --debugAddr flag", func() {
		BeforeEach(func() {
			client = startGarden()
		})

		It("does not expose the pprof debug endpoint", func() {
			_, err := http.Get(fmt.Sprintf("http://%s/debug/pprof/?debug=1", debugAddr))
			Expect(err).To(HaveOccurred())
		})

		It("does not expose the log level adjustment endpoint", func() {
			_, err := http.Get(fmt.Sprintf("http://%s/log-level -X PUT -d debug", debugAddr))
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when started with the --maxContainers flag", func() {
		Context("when maxContainers is lower than the subnet pool capacity", func() {
			BeforeEach(func() {
				client = startGarden("--maxContainers", "1")
			})

			Context("when getting the capacity", func() {
				It("returns the maxContainers flag value", func() {
					capacity, err := client.Capacity()
					Expect(err).ToNot(HaveOccurred())
					Expect(capacity.MaxContainers).To(Equal(uint64(1)))
				})
			})
		})

		Context("when maxContainers is higher than the subnet pool capacity", func() {
			BeforeEach(func() {
				client = startGarden("--maxContainers", "1000")
			})

			Context("when getting the capacity", func() {
				It("returns the capacity of the subnet pool", func() {
					capacity, err := client.Capacity()
					Expect(err).ToNot(HaveOccurred())
					Expect(capacity.MaxContainers).To(Equal(uint64(64)))
				})
			})
		})
	})

	Context("when starting with the --debugAddr flag", func() {
		BeforeEach(func() {
			client = startGarden("--debugAddr", debugAddr)
		})

		It("exposes the pprof debug endpoint", func() {
			_, err := http.Get(fmt.Sprintf("http://%s/debug/pprof/?debug=1", debugAddr))
			Expect(err).ToNot(HaveOccurred())
		})

		It("exposes the log level adjustment endpoint", func() {
			_, err := http.Get(fmt.Sprintf("http://%s/log-level -X PUT -d debug", debugAddr))
			Expect(err).ToNot(HaveOccurred())

			_, err = http.Get(fmt.Sprintf("http://%s/log-level -X PUT -d info", debugAddr))
			Expect(err).ToNot(HaveOccurred())

			_, err = http.Get(fmt.Sprintf("http://%s/log-level -X PUT -d error", debugAddr))
			Expect(err).ToNot(HaveOccurred())

			_, err = http.Get(fmt.Sprintf("http://%s/log-level -X PUT -d fatal", debugAddr))
			Expect(err).ToNot(HaveOccurred())
		})

		Describe("vars", func() {
			var (
				diskLimits garden.DiskLimits
				container  garden.Container
				vars       map[string]interface{}
			)

			BeforeEach(func() {
				diskLimits = garden.DiskLimits{
					ByteHard: 10 * 1024 * 1024,
					Scope:    garden.DiskLimitScopeExclusive,
				}
			})

			JustBeforeEach(func() {
				var err error

				container, err = client.Create(garden.ContainerSpec{
					Limits: garden.Limits{
						Disk: diskLimits,
					},
					RootFSPath: "docker:///busybox",
				})
				Expect(err).NotTo(HaveOccurred())

				response, err := http.Get(fmt.Sprintf("http://%s/debug/vars", debugAddr))
				Expect(err).ToNot(HaveOccurred())

				contents, err := ioutil.ReadAll(response.Body)
				Expect(err).ToNot(HaveOccurred())

				vars = make(map[string]interface{})
				Expect(json.Unmarshal(contents, &vars)).To(Succeed())
			})

			AfterEach(func() {
				Expect(client.Destroy(container.Handle())).To(Succeed())
			})

			It("exposes the number of loop devices", func() {
				Expect(vars["loopDevices"]).To(BeNumerically(">=", float64(1)))
			})

			It("exposes the number of depot directories", func() {
				Expect(vars["depotDirs"]).To(Equal(float64(1)))
			})

			It("exposes the number of backing stores", func() {
				Expect(vars["backingStores"]).To(Equal(float64(1)))
			})

			Context("when the container does not have a limit", func() {
				BeforeEach(func() {
					diskLimits = garden.DiskLimits{}
				})

				It("should not have any backing stores", func() {
					Expect(vars["depotDirs"]).To(Equal(float64(1)))
					Expect(vars["backingStores"]).To(Equal(float64(0)))
				})
			})
		})
	})
})
