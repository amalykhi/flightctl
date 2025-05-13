package ansible

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flightctl/flightctl/test/harness/e2e"
	"github.com/flightctl/flightctl/test/login"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

// _ is used as a blank identifier to ignore the return value of BeforeSuite, typically for initialization purposes.
var _ = BeforeSuite(func() {
	// This will be executed before all tests run.
	var h *e2e.Harness

	fmt.Println("Before all tests!")
	h = e2e.NewTestHarness()
	err := h.CleanUpAllResources()
	Expect(err).ToNot(HaveOccurred())
	login.LoginToAPIWithToken(h)
	//Install flightctl.core from Ansible Galaxy.
	galaxyCmd := exec.Command("ansible-galaxy", "collection", "install", "flightctl.core", "--force")
	output, err := galaxyCmd.CombinedOutput()
	if err != nil {
		logrus.Infof("Ansible Galaxy output: %s", output)
		Fail(fmt.Sprintf("Failed to install flightctl.core from Ansible Galaxy: %v", err))
	}
	logrus.Infof("Ansible Galaxy output: %s", output)
	logrus.Infof("Ansible Galaxy install successful")

	// 4. write-integration-config (assuming this is a shell script).
	//    You might need to adjust the path to the script.
	ansible_default_path := "/home/amalykhi/.ansible/collections/ansible_collections/flightctl/core"

	clientConfigPath := filepath.Join(os.Getenv("HOME"), ".config/flightctl/client.yaml")
	file, err := os.Open(clientConfigPath)
	if err != nil {
		Fail(fmt.Sprintf("Failed to open client.yaml: %v", err))
	}
	defer file.Close()

	var token, serviceAddr string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "  token:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				token = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "server:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				serviceAddr = strings.TrimSpace(parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		Fail(fmt.Sprintf("Failed to read client.yaml: %v", err))
	}

	configFile := filepath.Join(ansible_default_path, "./tests/integration/integration_config.yml")
	err = os.WriteFile(configFile, []byte(fmt.Sprintf("flightctl_token: %s\nflightctl_host: %s\n", token, serviceAddr)), 0644)
	if err != nil {
		Fail(fmt.Sprintf("Failed to write integration_config.yml: %v", err))
	}
	logrus.Infof("integration_config.yml written successfully")
})

// TestAnsible initializes and runs the suite of end-to-end tests for the Command Line Interface (CLI).
func TestAnsible(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ansible E2E Suite")
}
