package test

import (
	"testing"

	bplustest "https://gitlab.intelligentb.com/devops/bplus/test"
)

func TestMain(m *testing.M) {
	bplustest.BDD(m, FeatureContext)
}
