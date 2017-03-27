package libnecpp

/*
#include <libnecpp.h>
#cgo LDFLAGS: -lnecpp
*/
import "C"

import (
	"errors"
	"fmt"
	"strings"
)

// GainErrno
const GainErrno float64 = -999.0

var ErrNoPatternRequested = errors.New("no radiation pattern previously requested")

// PatchType is the type of patch for the SP card.
type PatchType int

// Patch shapes for Surface Patch (SP Card)
const (
	Arbitrary = PatchType(iota) // an arbitrary patch shape (the default)
	Rectangular 
	Triangular 
	Quadrilateral
)

type NecppCtx struct {
	necContext *C.nec_context
}

// New creates a new NEC context object, which contains the nec_context struct
// pointer from libnecpp. After it's done being used, call Delete() to free the
// struct.
func New() (*NecppCtx, error) {
	n := new(NecppCtx)
	nCtx := C.nec_create()
	if nCtx == nil {
		err := errors.New("nec_context was NULL")
		return nil, err
	}
	n.necContext = nCtx
	return n, nil
}

// if there's an error message, retrieve it. Only called by the wrapper
// functions below 

func (n *NecppCtx) errorMessage() error {
	// doesn't look like this should be freed
	cerr := C.nec_error_message()
	err := errors.New(C.GoString(cerr))
	return err
}

// these functions wrap around the various C functions from libnecpp - if they
// return a non-zero value, there's been an error of some kind. Get and return
// that error.

func (n *NecppCtx) errWrap(ret C.long) error {
	if ret != 0 {
		err := n.errorMessage()
		return err
	}
	return nil
}

// the gain functions are a little different, in that they return a meaningful
// number. If that number is -999.0, though, no radiation pattern as requested.

func (n *NecppCtx) gainErrWrap(gain C.double) (float64, error) {
	gainRet := float64(gain)
	if gainRet == GainErrno {
		return gainRet, ErrNoPatternRequested
	}
	return gainRet, nil
}

// Delete frees the nec_context struct. Call this after you're finished
// simulating the antenna.
func (n *NecppCtx) Delete() error {
	return n.errWrap(C.nec_delete(n.necContext))
}

// antenna geometry methods

// Wire creates a straight wire. The parameters are:
//
// 	tagId: the tag ID
// 	segmentCount: the number of segments
// 	xw1 The x coordinate of the wire starting point.
//	yw1 The y coordinate of the wire starting point.
//	zw1 The z coordinate of the wire starting point.
//	xw2 The x coordinate of the wire ending point.
//	yw2 The y coordinate of the wire ending point.
//	zw2 The z coordinate of the wire ending point.
//	rad The wire radius (meters)
//	rdel For tapered wires, the. Otherwise set to 1.0
//	rrad For tapered wires, the. Otherwise set to 1.0
//
// All co-ordinates are in meters.
func (n *NecppCtx) Wire(tagId int, segmentCount int, xw1 float64, yw1 float64, zw1 float64, xw2 float64, yw2 float64, zw2 float64, rad float64, rdel float64, rrad float64) error {
	return n.errWrap(C.nec_wire(n.necContext, C.int(tagId), C.int(segmentCount), C.double(xw1), C.double(yw1), C.double(zw1), C.double(xw2), C.double(yw2), C.double(zw2), C.double(rad), C.double(rdel), C.double(rrad)))
}

// SpCard makes a Surface Patch (SP) card.
//
// Parameters:
// 	ns PatchType - the type of Patch, see those constants
// 		Arbitrary
// 		Rectangular
// 		Triangular
// 		Quadrilateral
//	x1 The x coordinate of patch corner1.
//	y1 The y coordinate of patch corner1.
//	z1 The z coordinate of patch corner1.
//	x2 The x coordinate of patch corner2.
//	y2 The y coordinate of patch corner2.
//	z2 The z coordinate of patch corner2.
//
// All co-ordinates are in meters, except for arbitrary patches where the angles// are in degrees.
func (n *NecppCtx) SpCard(ns PatchType, x1 float64, y1 float64, z1 float64, x2 float64, y2 float64, z2 float64) error {
	return n.errWrap(C.nec_sp_card(n.necContext, C.int(ns), C.double(x1), C.double(y1), C.double(z1), C.double(x2), C.double(y2), C.double(z2)))
}

// ScCard makes a Surface Patch Continuation (SC) card.
//
// Parameters:
//	i2  Weird integer parameter.
//	x3 The x coordinate of patch corner 3.
//	y3 The y coordinate of patch corner 3.
//	z3 The z coordinate of patch corner 3.
//	x4 The x coordinate of patch corner 4.
//	y4 The y coordinate of patch corner 4.
//	z4 The z coordinate of patch corner 4.
//
// All co-ordinates are in meters.
func (n *NecppCtx) ScCard(i2 int, x3 float64, y3 float64, z3 float64, x4 float64, y4 float64, z4 float64) error {
	return n.errWrap(C.nec_sc_card(n.necContext, C.int(i2), C.double(x3), C.double(y3), C.double(z3), C.double(x4), C.double(y4), C.double(z4)))
}

// GmCard makes a GM card for Coordinate Transformation
//
// Parameters:
//	 itsi  Tag number increment.
//	 nprt  The number of new Structures to be generated
//	 rox   Angle in degrees through which the structure is rotated about
//             the X-axis.  A positive angle causes a right-hand rotation.
//	 roy   Angle of rotation about Y-axis.
//	 roz   Angle of rotation about Z-axis.
//	 xs    X, Y. Z components of vector by which
//	 ys    structure is translated with respect to
//	 zs    the coordinate system.
//	 its   This number is input as a decimal number but is rounded
//             to an integer before use.  Tag numbers are searched sequentially
//             until a segment having a tag of this segment through the end of
//             the sequence of segments is moved by the card.  If ITS is zero 
//             the entire structure is moved.
func (n *NecppCtx) GmCard(itsi int, nrpt int, rox float64, roy float64, roz float64, xs float64, ys float64, zs float64, its int) error {
	return n.errWrap(C.nec_gm_card(n.necContext, C.int(itsi), C.int(nrpt), C.double(rox), C.double(roy), C.double(roz), C.double(xs), C.double(ys), C.double(zs), C.int(its)))
}

// GxCard creates a GX card for Reflection in coordinate Planes.
//
/* Parameters:
	i1 - Tag number increment.
   	i2 - This integer is divided into three independent digits, in
        columns 8, 9, and 10 of the card, which control reflection
        in the three orthogonal coordinate planes.  A one in column
        8 causes reflection along the X-axis (reflection in Y, Z
        plane); a one in column 9 causes reflection along the Y-axis;
        and a one in column 10 causes reflection along the Z axis.
        A zero or blank in any of these columns causes the corres-
        ponding reflection to be skipped.

Any combination of reflections along the X, Y and Z axes may be used. 
For example, 101 for i2 will cause reflection along axes X and Z, and 111 will
cause reflection along axes X, Y and Z. When combinations of reflections are requested, 
the reflections are done in reverse alphabetical order. That is, if a structure is 
generated in a single octant of space and a GX card is then read with i2 equal to 111, 
the structure is first reflected along the Z-axis; the structure and its image are 
then reflected along the Y-axis; and, finally, these four structures are reflected 
along the X-axis to fill all octants. This order determines the position of a segment 
in the sequence and, hence, the absolute segment numbers.

The tag increment i1 is used to avoid duplication of tag numbers in the image 
segments. All valid tags on the original structure are incremented by i1 on the image.
When combinations of reflections are employed, the tag increment is doubled after each 
reflection. Thus, a tag increment greater than or equal to the largest tag an the 
original structure will ensure that no duplicate tags are generated. For example, 
if tags from 1 to 100 are used on the original structure with i2 equal to 011 and 
a tag increment of 100, the first reflection, along the Z-axis, will produce tags 
from 101 to 200; and the second reflection, along the Y-axis, will produce tags 
rom 201 to 400, as a result of the increment being doubled to 200.
*/
func (n *NecppCtx) GxCard(i1 int, i2 int) error {
	return n.errWrap(C.nec_gx_card(n.necContext, C.int(i1), C.int(i2)))
}

func (n *NecppCtx) GeometryComplete(gplfag int) error {
	return n.errWrap(C.nec_geometry_complete(n.necContext, C.int(gplfag)))
}

// antenna environment methods

func (n *NecppCtx) MediumParameters(permittivity float64, permeability float64) error {
	return n.errWrap(C.nec_medium_parameters(n.necContext, C.double(permittivity), C.double(permeability)))
}

func (n *NecppCtx) GnCard(iperf int, nradl int, epse float64, sig float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_gn_card(n.necContext, C.int(iperf), C.int(nradl), C.double(epse), C.double(sig), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

func (n *NecppCtx) FrCard(inIfrq int, inNfrq int, inFreqMhz float64, inDelFreq float64) error {
	return n.errWrap(C.nec_fr_card(n.necContext, C.int(inIfrq), C.int(inNfrq), C.double(inFreqMhz), C.double(inDelFreq)))
}

func (n *NecppCtx) EkCard(itmp1 int) error {
	return n.errWrap(C.nec_ek_card(n.necContext, C.int(itmp1)))
}

func (n *NecppCtx) LdCard(ldtype int, ldtag int, ldtagf int, ldtagt int, tmp1 float64, tmp2 float64, tmp3 float64) error {
	return n.errWrap(C.nec_ld_card(n.necContext, C.int(ldtype), C.int(ldtag), C.int(ldtagf), C.int(ldtagt), C.double(tmp1), C.double(tmp2), C.double(tmp3)))
}

func (n *NecppCtx) ExCard(extype int, i2 int, i3 int, i4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_ex_card(n.necContext, C.int(extype), C.int(i2), C.int(i3), C.int(i4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

func (n *NecppCtx) ExcitationVoltage(tag int, segment int, vReal float64, vImag float64) error {
	return n.errWrap(C.nec_excitation_voltage(n.necContext, C.int(tag), C.int(segment), C.double(vReal), C.double(vImag)))
}

func (n *NecppCtx) ExcitationCurrent(x float64, y float64, z float64, a float64, beta float64, moment float64) error {
	return n.errWrap(C.nec_excitation_current(n.necContext, C.double(x), C.double(y), C.double(z), C.double(a), C.double(beta), C.double(moment)))
}

func (n *NecppCtx) ExcitationPlanewave(nTheta int, nPhi int, theta float64, phi float64, eta float64, dtheta float64, dphi float64, polRatio float64) error {
	return n.errWrap(C.nec_excitation_planewave(n.necContext, C.int(nTheta), C.int(nPhi), C.double(theta), C.double(phi), C.double(eta), C.double(dtheta), C.double(dphi), C.double(polRatio)))
}

func (n *NecppCtx) TlCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_tl_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

func (n *NecppCtx) NtCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_nt_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

func (n *NecppCtx) XqCard(itmp1 int) error {
	return n.errWrap(C.nec_xq_card(n.necContext, C.int(itmp1)))
}

func (n *NecppCtx) GdCard(tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64) error {
	return n.errWrap(C.nec_gd_card(n.necContext, C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4)))
}

// simulation output

func (n *NecppCtx) RpCard(calcMode int, nTheta int, nPhi int, outputFormat int, normalization int, d int, a int, theta0 float64, phi0 float64, deltaTheta float64, deltaPhi float64, radialDistance float64, gainNorm float64) error {
	return n.errWrap(C.nec_rp_card(n.necContext, C.int(calcMode), C.int(nTheta), C.int(nPhi), C.int(outputFormat), C.int(normalization), C.int(d), C.int(a), C.double(theta0), C.double(phi0), C.double(deltaTheta), C.double(deltaPhi), C.double(radialDistance), C.double(gainNorm)))
}

func (n *NecppCtx) PtCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int) error {
	return n.errWrap(C.nec_pt_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4)))
}

func (n *NecppCtx) PqCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int) error {
	return n.errWrap(C.nec_pq_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4)))
}

func (n *NecppCtx) KhCard(tmp1 float64) error {
	return n.errWrap(C.nec_kh_card(n.necContext, C.double(tmp1)))
}

func (n *NecppCtx) NeCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_ne_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

func (n *NecppCtx) NhCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_nh_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

func (n *NecppCtx) CpCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int) error {
	return n.errWrap(C.nec_cp_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4)))
}

func (n *NecppCtx) PlCard(ploutputFilename string, itmp1 int, itmp2 int, itmp3 int, itmp4 int) error {
	return n.errWrap(C.nec_pl_card(n.necContext, C.CString(ploutputFilename), C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4)))
}

// analysis of output

func (n *NecppCtx) Gain(freqIndex int, thetaIndex int, phiIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain(n.necContext, C.int(freqIndex), C.int(thetaIndex), C.int(phiIndex)))
}

func (n *NecppCtx) GainMax(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_max(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainMin(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_min(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainMean(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_mean(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainSd(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_sd(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainRhcpMax(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_rhcp_max(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainRhcpMin(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_rhcp_min(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainRhcpMean(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_rhcp_mean(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainRhcpSd(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_rhcp_sd(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainLhcpMax(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_lhcp_max(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainLhcpMin(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_lhcp_min(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainLhcpMean(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_lhcp_mean(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) GainLhcpSd(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_lhcp_sd(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) ImpedanceReal(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_impedance_real(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) ImpedanceImag(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_impedance_imag(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) Impedance(freqIndex int) (complex128, error) {
	r, rerr := n.ImpedanceReal(freqIndex)
	i, ierr := n.ImpedanceImag(freqIndex)
	ret := complex(r, i)

	var errstrs []string
	if rerr != nil {
		errstrs = append(errstrs, fmt.Sprintf("real component error: %s", rerr.Error()))
	}
	if ierr != nil {
		errstrs = append(errstrs, fmt.Sprintf("imaginary component error: %s", rerr.Error()))
	}
	if errstrs != nil {
		err := errors.New(strings.Join(errstrs, " :: "))
		return ret, err
	}
	return ret, nil
}
