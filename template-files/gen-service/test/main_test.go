package test

import (
	"testing"

	bplustest "gitlab.intelligentb.com/devops/bplus/test"
)

func TestMain(m *testing.M) {
	bplustest.BDD(m, FeatureContext)
}
