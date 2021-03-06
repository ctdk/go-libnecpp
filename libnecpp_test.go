package necpp

import (
	"math"
	"testing"
)

func TestNecCtxCreate(t *testing.T) {
	n, err := New()
	if err != nil {
		t.Error(err)
	}
	n.Delete()
}

func TestNecCtxDeletion(t *testing.T) {
	n, _ := New()
	err := n.Delete()
	if err != nil {
		t.Error(err)
	}
}

func TestSimpleAntenna(t *testing.T) {
	var expMax float64 = 8.407404
	var expMin float64 = -999.99
	var expMean float64 = -1.958236
	var expSd float64 = 16.108163

	n, _ := New()
	defer n.Delete()

	err := n.Wire(0, 9, 0, 0, 2, 0, 0, 7, 0.1, 1, 1)
	if err != nil {
		t.Error(err)
	}
	err = n.GeometryComplete(CurrentExpansionModified)
	if err != nil {
		t.Error(err)
	}
	err = n.GnCard(Perfect, 0, 0, 0, 0, 0, 0, 0)
	if err != nil {
		t.Error(err)
	}
	err = n.FrCard(Linear, 1, 30, 0)
	if err != nil {
		t.Error(err)
	}
	err = n.ExCard(VoltageApplied, 0, 5, 0, 1.0, 0, 0, 0, 0, 0)
	if err != nil {
		t.Error(err)
	}
	err = n.RpCard(Normal, 90, 1, MajorMinor, TotalNormalized, PowerGain, NoAvg, 0, 90, 1, 0, 0, 0)
	if err != nil {
		t.Error(err)
	}
	max, err := n.GainMax(0)
	if err != nil {
		t.Error(err)
	}
	min, err := n.GainMin(0)
	if err != nil {
		t.Error(err)
	}
	mean, err := n.GainMean(0)
	if err != nil {
		t.Error(err)
	}
	sd, err := n.GainSd(0)
	if err != nil {
		t.Error(err)
	}
	// not going to even try and figure out how to compare the number
	_, err = n.Impedance(0)
	if err != nil {
		t.Error(err)
	}
	if expMax != roundFloat(max, 6) {
		t.Errorf("max gain was %f, should have been %f", roundFloat(max, 6), expMax)
	}
	if expMin != roundFloat(min, 6) {
		t.Errorf("min gain was %f, should have been %f", roundFloat(min, 6), expMin)
	}
	if expMean != roundFloat(mean, 6) {
		t.Errorf("mean gain was %f, should have been %f", roundFloat(mean, 6), expMean)
	}
	if expSd != roundFloat(sd, 6) {
		t.Errorf("sd gain was %f, should have been %f", roundFloat(sd, 6), expSd)
	}
}

func roundFloat(n float64, d int) float64 {
	p := math.Pow(10, float64(d))
	return float64(int(n*float64(p))) / p
}

func TestCrossedWires(t *testing.T) {
	n, _ := New()
	defer n.Delete()
	if err := n.Wire(4, 11, -0.0318, -0.0287, 0.0775, -0.0318, 0.0439, 0.014, 0.001, 1.0, 1.0); err != nil {
		t.Errorf("This shouldn't have caused an error: %s", err.Error())
	}
	if err := n.Wire(5, 7, -0.0318, 0.0439, 0.014, -0.0318, 0.0045, 0.0624, 0.001, 1.0, 1.0); err == nil {
		t.Errorf("The second wire *should* have caused an error because it crossed over the first wire, but didn't")
	}

}
