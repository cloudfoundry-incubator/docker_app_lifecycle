package spec_test

import (
	. "github.com/cloudfoundry-incubator/docker-circus/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry-incubator/docker-circus/Godeps/_workspace/src/github.com/onsi/gomega"
	"testing"
)

func TestSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Suite")
}