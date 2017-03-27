/*
 * Golang example using libnecpp as a library.
 *
 * This is a translation of the C example in nec++, found at
 * https://github.com/tmolteno/necpp/blob/master/example/test_nec.c
 *
 */
package main

import (
	"fmt"
	"github.com/ctdk/go-necpp/libnecpp"
)

func main() {
	simpleExample()
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
	n, err := libnecpp.New()
	if err != nil {
		panic(err)
	}
	defer n.Delete()

	fmt.Printf("simple antenna example")
	fmt.Printf("----------------------")

	// skipping err checks here because it's a) an example and b) this very
	// same antenna is tested in the go tests in libnecpp/necpp_test.go
	n.Wire(0, 9, 0, 0, 2, 0, 0, 7, 0.1, 1, 1)
	n.GeometryComplete(libnecpp.CurrentExpansionModified)
	n.GnCard(1, 0, 0, 0, 0, 0, 0, 0)
	n.FrCard(0, 1, 30, 0)
	n.ExCard(0, 0, 5, 0, 1.0, 0, 0, 0, 0, 0)
	n.RpCard(0, 90, 1, 0,5,0,0, 0, 90, 1, 0, 0, 0)
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
	fmt.Printf("LHCP Gain: %f, %f +/- %f dB\n", maxL, meanL, sdL)
}

func example3() {

}
