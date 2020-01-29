package test

import (
	"testing"

	bplustest "github.com/MenaEnergyVentures/bplus/test"
)

func TestMain(m *testing.M) {
	bplustest.BDD(m, FeatureContext)
}
