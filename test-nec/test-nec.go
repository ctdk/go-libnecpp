/*
Golang example using libnecpp as a library.

This is a translation of the C example in nec++, found at
https://github.com/tmolteno/necpp/blob/master/example/test_nec.c

*/
package main

import (
	"fmt"
	"github.com/ctdk/go-libnecpp"
)

func main() {
	simpleExample()
	example3()
}

func sevenWireAntenna() {

}

func simpleExample() {
	/*  GW 0 9 0. 0. 2. 0. 0. 7 .1
	    GE 1
	    FR 0 1 0 30.
	    EX 0 5 0 1.
	    GN 1
	    RP 0 90 1 0000 0 90 1 0
	*/
	n, err := necpp.New()
	if err != nil {
		panic(err)
	}
	defer n.Delete()

	fmt.Println("simple antenna example")
	fmt.Println("----------------------")

	// skipping err checks here because it's a) an example and b) this very
	// same antenna is tested in the go tests in libnecpp/necpp_test.go
	n.Wire(0, 9, 0, 0, 2, 0, 0, 7, 0.1, 1, 1)
	n.GeometryComplete(necpp.CurrentExpansionModified)
	n.GnCard(necpp.Perfect, 0, 0, 0, 0, 0, 0, 0)
	n.FrCard(necpp.Linear, 1, 30, 0)
	n.ExCard(necpp.VoltageApplied, 0, 5, 0, 1.0, 0, 0, 0, 0, 0)
	n.RpCard(necpp.Normal, 90, 1, necpp.MajorMinor, necpp.TotalNormalized, necpp.PowerGain, necpp.NoAvg, 0, 90, 1, 0, 0, 0)
	max, _ := n.GainMax(0)
	mean, _ := n.GainMean(0)
	sd, _ := n.GainSd(0)
	fmt.Printf("Gain: %f, %f +/- %f dB\n", max, mean, sd)

	maxR, _ := n.GainRhcpMax(0)
	meanR, _ := n.GainRhcpMean(0)
	sdR, _ := n.GainRhcpSd(0)
	fmt.Printf("RHCP Gain: %f, %f +/- %f dB\n", maxR, meanR, sdR)

	maxL, _ := n.GainLhcpMax(0)
	meanL, _ := n.GainLhcpMean(0)
	sdL, _ := n.GainLhcpSd(0)
	fmt.Printf("LHCP Gain: %f, %f +/- %f dB\n\n", maxL, meanL, sdL)
}

func example3() {
	/*
	   CMEXAMPLE 3. VERTICAL HALF WAVELENGTH ANTENNA OVER GROUND
	   CM           EXTENDED THIN WIRE KERNEL USED
	   CM           1. PERFECT GROUND
	   CM           2. IMPERFECT GROUND INCLUDING GROUND WAVE AND RECEIVING
	   CE              PATTERN CALCULATIONS
	   GW 0 9 0. 0. 2. 0. 0. 7. .03
	   GE 1
	   EK
	   FR 0 1 0 0 30.
	   EX 0 0 5 0 1.
	   GN 1
	   RP 0 10 2 1301 0. 0. 10. 90.
	   GN 0 0 0 0 6. 1.000E-03
	   RP 0 10 2 1301 0. 0. 10. 90.
	   RP 1 10 1 0 1. 0. 2. 0. 1.000E+05
	   EX 1 10 1 0 0. 0. 0. 10.
	   PT 2 0 5 5
	   XQ
	   EN
	*/

	n, err := necpp.New()
	if err != nil {
		panic(err)
	}
	defer n.Delete()

	fmt.Printf("example3\n----------\n")

	n.Wire(0, 9, 0., 0.0, 2.0, 0.0, 0.0, 7.0, 0.03, 1.0, 1.0)
	n.GeometryComplete(necpp.CurrentExpansionModified)
	n.EkCard(necpp.ReturnToNormal)
	n.FrCard(necpp.Linear, 1, 30.0, 0)
	n.ExCard(necpp.VoltageApplied, 0, 5, 0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	n.GnCard(necpp.Perfect, 0, 0, 0, 0, 0, 0, 0)
	n.RpCard(necpp.Normal, 10, 2, necpp.VerticalHorizontal, necpp.VerticalAxisNorm, necpp.PowerGain, necpp.AvgGain, 0.0, 0.0, 10.0, 90.0, 0, 0)

	imp, _ := n.Impedance(0)
	fmt.Printf("Impedance: %g\n", imp)

	n.GnCard(necpp.Nullified, 0, 6.0, 1.000E-03, 0.0, 0.0, 0.0, 0.0)
	n.RpCard(necpp.Normal, 10, 2, necpp.VerticalHorizontal, necpp.VerticalAxisNorm, necpp.PowerGain, necpp.AvgGain, 0.0, 0.0, 10.0, 90.0, 0, 0)

	imp2, _ := n.Impedance(0)
	fmt.Printf("Impedance 2: %g\n", imp2)
}
