package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"os/exec"
	"path/filepath"
)

var _ = Describe("goto", func() {
	var cliPath string

	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/EngineerBetter/goto/workto")
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})

	It("requires an argument", func() {
		command := exec.Command(cliPath)
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(1))
		Ω(session.Err).Should(Say("directory to look for was not specified"))
	})

	It("finds the given directory", func() {
		command := exec.Command(cliPath, "bosh-lite")
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit())

		expectedOutput := filepath.Join("/Users/deejay/workspace/bosh-lite")

		Ω(session.Out).Should(Say(expectedOutput))
	})
})