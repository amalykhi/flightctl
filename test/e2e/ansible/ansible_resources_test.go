package ansible

import (
	"fmt"
	"os/exec"

	"github.com/flightctl/flightctl/test/harness/e2e"
	"github.com/flightctl/flightctl/test/login"
	. "github.com/onsi/ginkgo/v2"
	"github.com/sirupsen/logrus"
)

// _ is a blank identifier used to ignore values or expressions, often applied to satisfy interface or assignment requirements.
var _ = Describe("Flightctl Ansible Integration", func() {
	var (
		harness *e2e.Harness
	)

	BeforeEach(func() {
		harness = e2e.NewTestHarness()
		login.LoginToAPIWithToken(harness)
	})

	AfterEach(func() {
		harness.Cleanup(false) // do not print console on error
	})

	Context("resource tests", func() {
		It("should get device information using Ansible", Label("00000"), func() {
			testCmd := exec.Command("ansible-test", "integration", "--target", "targets/flightctl_resource_info/tasks/devices.yml") // Adjust
			testOutput, err := testCmd.CombinedOutput()
			if err != nil {
				logrus.Infof("ansible-test output: %s", testOutput)
				Fail(fmt.Sprintf("Ansible test failed: %v\nOutput:\n%s", err, testOutput))
			}
			logrus.Infof("ansible-test output: %s", testOutput)
		})
	})
})
