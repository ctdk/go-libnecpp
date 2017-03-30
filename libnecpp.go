package necpp

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

// GainErrno is the number returned by the Gain* functions when no radiation
// pattern was previously requested.
const GainErrno float64 = -999.0

var ErrNoPatternRequested = errors.New("no radiation pattern previously requested")

// PatchType is the shape of a patch for the Surface Patch (SP Card).
type PatchType int

const (
	Arbitrary PatchType = iota // an arbitrary patch shape (the default)
	Rectangular
	Triangular
	Quadrilateral
)

// GeoGroundPlaneFlag is used to indicate the type of ground plane to use with
// the antenna when indicating the geometry is complete.
//
// The types of ground plane to use are:
//
// • NoGroundPlane - no ground plane is present. (Fairly self-explanatory.)
//
// • CurrentExpansionModified - Structure symmetry is modified as required, and
// the current expansion is modified so that the currents and segments touching
// the ground (x, Y plane) are interpolated to their images below the ground
// (charge at base is zero)
//
// • CurrentExpansionUnmodified - indicates a ground is present. Structure
// symmetry is modified as required. Current expansion, however, is not
// modified, Thus, currents on segments touching the ground will go to zero at
// the ground.
type GeoGroundPlaneFlag int

const (
	CurrentExpansionUnmodified GeoGroundPlaneFlag = iota - 1
	NoGroundPlane
	CurrentExpansionModified
	MooMoo
)

// GroundTypeFlag indicates the general type of ground for the antenna.
//
// The flags for ground types are:
//
// • Nullified - Nullifies ground parameters previously used and sets free-space
// condition. The remainder of the parameters are ignored in this case.
//
// • Finite - Finite ground, reflection coefficient approximation.
//
// • Perfect - Perfectly conducting ground.
//
// • FiniteSomNorton - Finite ground, Sommerfeld/Norton method.
type GroundTypeFlag int

const (
	Nullified GroundTypeFlag = iota - 1
	Finite
	Perfect
	FiniteSomNorton
)

// FrequencyRange is used to set the type of frequency range for FR cards.
type FrequencyRange int

const (
	Linear      FrequencyRange = iota // a linear range
	Logarithmic                       // a logarithmic range
)

// WireKernel sets the type of wire kernel to use with EkCard
//
// • ReturnToNormal - Return to normal kernel
//
// • ExtendedThinWire - Use extended thin wire kernel
type WireKernel int

const (
	ReturnToNormal WireKernel = iota - 1
	ExtendedThinWire
)

// Excitation sets the type of excitation for ExCard
type Excitation int

const (
	VoltageApplied    Excitation = iota // voltage source (applied-E-field source)
	IncidentLinear                      // incident plane wave, linear polarization.
	IncidentRightHand                   // incident plane wave, right-hand (thumb along the incident k vector) elliptic polarization.
	IncidentLeftHand                    // incident plane wave, left-hand elliptic polarization.
	Elementary                          // elementary current source
	VoltageSlope                        // voltage source (current-slope-discontinuity)
)

// ExecutionOption control the generation of radiation patterns with XqCard()
//
// Options for radiation patterns:
//
// • NoPattern - no patterns requested (the normal case).
//
// • XZPlane - generates a pattern cut in the XZ plane, i.e., phi = 0 degrees
// and theta varies from 0 degrees to 90 degrees in 1 degree steps.
//
// • YZPlane - generates a pattern cut in the YZ plane, i.e., phi = 90 degrees
// theta varies from 0 degrees to 90 degrees in 1 degree steps.
//
// • BothPlane - generates both of the cuts described for XZPlane and YZPlane.
type ExecutionOption int

const (
	NoPattern ExecutionOption = iota
	XZPlane
	YZPlane
	BothPlane
)

// wow, there are a lot of flags for RP cards. :-/

// RpCalcMode is used to select radiation patterns for RP cards.
//
// RP card calculation flags:
//
// • Normal - normal mode. Space-wave fields are computed. An infinite ground
// plane is included if it has been specified previously on a GN card;
// otherwise, the antenna is in free space.
//
// • SurfaceWave - surface wave propagating along ground is added to the normal
// space wave. This option changes the meaning of some of the other parameters
// on the RP card as explained below, and the results appear in a special output
// format. Ground parameters must have been input on a GN card. The following
// options cause calculation of only the space wave but with special ground
// conditions. Ground conditions include a two-medium ground (cliff where the
// media join in a circle or a line), and a radial wire ground screen. Ground
// parameters and dimensions must be input on a GN or GD card before the RP card
// is read. The RP card only selects the option for inclusion in the field
// calculation. (Refer to the GN and GD cards for further explanation.)
//
// • LinearCliff - linear cliff with antenna above upper level. Lower medium
// parameters are as specified for the second medium on the GN card or on the
// GD card.
//
// • CircularCliff - circular cliff centered at origin of coordinate system:
// with antenna above upper level. Lower medium parameters are as specified for
// the second medium on the GN card or on the GD card.
//
// • RadialScreen - radial wire ground screen centered at origin.
//
// • RadialLinearCliff - both radial wire ground screen and linear cliff.
//
// • RadialCircularCliff - both radial wire ground screen ant circular cliff.
type RpCalcMode int

const (
	Normal RpCalcMode = iota
	SurfaceWave
	LinearCliff
	CircularCliff
	RadialScreen
	RadialLinearCliff
	RadialCircularCliff
)

// RpOutputFormat is used to select the output format for RP cards.
//
// • MajorMinor - major axis, minor axis and total gain printed.
//
// • VerticalHorizontal - vertical, horizontal ant total gain printed.
type RpOutputFormat int

const (
	MajorMinor RpOutputFormat = iota
	VerticalHorizontal
)

// RpNormalization is used to select the normalization type for RP cards.
//
// Radiation pattern normalization flags:
//
// • NoNormalization - no normalized gain.
//
// • MajorAxisNorm - major axis gain normalized.
//
// • MinorAxisNorm - minor axis gain normalized.
//
// • VerticalAxisNorm - vertical axis gain normalized.
//
// • HorizontalAxisNorm - horizontal axis gain normalized.
//
// • TotalNormalized - total gain normalized.
type RpNormalization int

const (
	NoNormalization RpNormalization = iota
	MajorAxisNorm
	MinorAxisNorm
	VerticalAxisNorm
	HorizontalAxisNorm
	TotalNormalized
)

// RpGain is used to select the type of gain for RP cards for standard printing
// and normalization constants. These ones have self explanatory names.
type RpGain int

const (
	PowerGain RpGain = iota
	DirectiveGain
)

// RpAveraging is used to select the type of averaging for RP cards to set the
// calculation of average power gain over the region covered by field points.
//
// RpAveraging flags:
//
// • NoAvg - no averaging
//
// • AvgGain - average gain computed.
//
// • AvgGainPrtSuppressed - average gain computed, printing of gain at the field
// points used for averaging is suppressed. If nTheta or NPH is equal to one,
// average gain will not be computed for any value of A since the area of the
// region covered by field points vanishes.
type RpAveraging int

const (
	NoAvg RpAveraging = iota
	AvgGain
	AvgGainPrtSuppressed
)

// NecppCtx is the nec context, and contains the libnecpp nec_context struct
// within itself.
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

// GeometryComplete indicates the antenna geometry is complete - makes a GE
// card. See GeoGroundPlaneFlag for details on that parameter.
func (n *NecppCtx) GeometryComplete(gpflag GeoGroundPlaneFlag) error {
	return n.errWrap(C.nec_geometry_complete(n.necContext, C.int(gpflag)))
}

// antenna environment methods

// MediumParameters sets the permittivity and permeability of the medium.
//
// Parameters:
// 	permittivity - The electric permittivity of the medium (in farads per
// 		meter)
// 	permeability - The magnetic permeability of the medium (in henries per
// 		meter)
func (n *NecppCtx) MediumParameters(permittivity float64, permeability float64) error {
	return n.errWrap(C.nec_medium_parameters(n.necContext, C.double(permittivity), C.double(permeability)))
}

// GnCard makes a ground card.
//
// Examples (TODO: check these examples out more - not sure if they're quite
// right)
//
// 1) Infinite ground plane
// 	n.GnCard(Perfect, 0, 0, 0, 0, 0, 0, 0)
// 2) Radial Wire Ground Plane (4 wires, 2 meters long, 5mm in radius)
// (This is the example I'm unsure of)
// 	n.GnCard(Finite, 4, 0.0, 0.0, 2.0, 0.005, 0.0, 0.0)
// 	(example from libnecpp was nec_gn_card(nec, 4, 0, 0.0, 0.0, 2.0, 0.005,
// 	 0.0, 0.0))
//
// Parameters (some - not all were detailed in the upstream documentation)
//
// 	iperf - Ground type flag. See the GroundTypeFlag constants for what
// 	goes here.
// 	epse - Relative dielectric constant for ground in the vicinity of the
// 	antenna. Zero in the case of perfect ground.
// 	sig - Conductivity in mhos/meter of the ground in the vicinity of the
// 	antenna. Use zero in the case of a perfect ground. If SIG is input as a
// 	negative number, the complex dielectric constant Ec = Er -j sigma/omega
// 	epsilon is set to EPSR - |SIG|.
func (n *NecppCtx) GnCard(iperf GroundTypeFlag, nradl int, epse float64, sig float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_gn_card(n.necContext, C.int(iperf), C.int(nradl), C.double(epse), C.double(sig), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

// FrCard makes a FR Card for frequency ranges.
//
// Parameters:
// 	inIfrq - a FrequencyRange constant, Linear or Logarithmic, for linear or
// 	logarithmic range of frequencies.
// 	inNfreq - the number of frequencies
// 	inFreqMhz - the starting frequency in MHz.
// 	inDelFreq - the frequency step in MHz (for inIfreq == Linear)
func (n *NecppCtx) FrCard(inIfrq FrequencyRange, inNfrq int, inFreqMhz float64, inDelFreq float64) error {
	return n.errWrap(C.nec_fr_card(n.necContext, C.int(inIfrq), C.int(inNfrq), C.double(inFreqMhz), C.double(inDelFreq)))
}

// EkCard controls the use of the external thin-wire kernel approximation.
func (n *NecppCtx) EkCard(itmp1 WireKernel) error {
	return n.errWrap(C.nec_ek_card(n.necContext, C.int(itmp1)))
}

// LdCard - loading.
//
// Parameters:
//
//	ldtyp - Type of loading (5 = segment conductivity)
//	ldtag - Tag (zero for absolute segment numbers, or in conjunction with 0 for next parameter, for all segments)
//	ldtagf - Equal to m specifies the mth segment of the set of segments
// 	whose tag numbers equal the tag number specified in the previous
// 	parameter. If the previous parameter (LDTAG) is zero, LDTAGF then
// 	specifies an absolute segment number. If both LDTAG and LDTAGF are zero,
// 	all segments will be loaded.
//	ldtagt - Equal to n specifies the nth segment of the set of segments
// 	whose tag numbers equal the tag number specified in the parameter LDTAG.
// 	This parameter must be greater than or equal to	the previous parameter.
// 	The loading specified is applied to each of the	mth through nth segments
// 	of the set of segments having tags equal to LDTAG. Again if LDTAG is
// 	zero, these parameters refer to absolute segment numbers. If LDTAGT is
// 	left blank, it is set equal to the previous parameter (LDTAGF).
//	tmp1 Resistance in Ohms, OR (A) Ohms per meter, OR (B) Resistance. OR
// 	(C) Conductivity (ldtyp=5)
//	tmp2 IND., HENRY, OR (A) HY/LENGTH OR (B) REACT. OR (C) Set to 0.0
//	tmp3 CAP,. FARAD, OR (A,B) BLANK (set to 0.0)
func (n *NecppCtx) LdCard(ldtype int, ldtag int, ldtagf int, ldtagt int, tmp1 float64, tmp2 float64, tmp3 float64) error {
	return n.errWrap(C.nec_ld_card(n.necContext, C.int(ldtype), C.int(ldtag), C.int(ldtagf), C.int(ldtagt), C.double(tmp1), C.double(tmp2), C.double(tmp3)))
}

// ExCard applies a source of excitation to the antenna, making an EX card.
//
// Parameters:
// 	extype - one of the Excitation constants - VoltageApplied,
// 	 	IncidentLinear, IncidentLeftHand, IncidentRightHand, Elementary,
// 		or VoltageSlope
// 	i2 - Tag number the source segment. This tag number along with the
// 	number to be given in (i3), which identifies the position of the
// 	segment in a set of equal tag numbers, uniquely definer the source
// 	segment. A 0 in field i2 implies that the Source segment will be
// 	identified by using the absolute segment number in the next field (i3).
// 	i3 - Equal to m, specifies the mth segment of the set of segments whose
// 	tag numbers are equal to the number set by the previous parameter. If
// 	the previous parameter is zero, the number in (i3) must be the absolute
// 	segment number of the source.
// 	i4 -  Meaning Depends on the extype parameter. See http://www.nec2.org/part_3/cards/ex.html
//
// The meaning of the floating point parameter depends on the excitation type.
// See http://www.nec2.org/part_3/cards/ex.html for more details.
//
// Simpler versions of the function are provided for common uses. These are
// ExcitationVoltage, ExcitationCurrent, and ExcitationPlanewave.
func (n *NecppCtx) ExCard(extype Excitation, i2 int, i3 int, i4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_ex_card(n.necContext, C.int(extype), C.int(i2), C.int(i3), C.int(i4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

// ExcitationVoltage makes a voltage source excitation source for the antenna.
// It is one of the simpler versions of ExCard().
//
// Parameters:
//	tag - Tag number of the source segment. This tag number along with the
// 	number to be given in (segment), which identifies the position of the
// 	segment in a set of equal tag numbers, uniquely definer the source
// 	segment. A zero in field (tag) implies that the Source segment will be
// 	identified by using the absolute segment number in the next field
// 	(segment).
//	segment - Equal to m, specifies the mth segment of the set of segments
// 	whose tag numbers are equal to the number
//	set by the previous parameter. If the previous parameter is zero, the
// 	number in (segment) must be the absolute segment number of the source.
// 	voltageExcitation - a complex128 number representing the voltage
// 	excitation.
// Only one incident plane wave or one elementary current source is allowed at
// a time. Also plane-wave or current-source excitation is not allowed with
// voltage sources.  If the excitation types are mixed, the program will use the
// last excitation type encountered.
func (n *NecppCtx) ExcitationVoltage(tag int, segment int, voltageExcitation complex128) error {
	return n.errWrap(C.nec_excitation_voltage(n.necContext, C.int(tag), C.int(segment), C.double(real(voltageExcitation)), C.double(imag(voltageExcitation))))
}

// ExcitationCurrent makes a current source excitation for the antenna. It is
// one of the simpler versions of ExCard().
//
// Parameters:
//	x - X position in meters.
//	y - Y position in meters.
//	z - Z position in meters.
//	a - a in degrees. a is the angle the current source makes with the XY
// 	plane as illustrated on figure 15.
//	beta - beta in degrees. beta is the angle the projection of the current
// 	source on the XY plane makes with the X axis.
//	moment - "Current moment" of the source. This parameter is equal to the
// 	product Il in amp meters.
//
// Only one incident plane wave or one elementary current source is allowed at
// a time. Also plane-wave or current-source excitation is not allowed with
// voltage sources.  If the excitation types are mixed, the program will use the
// last excitation type encountered.
func (n *NecppCtx) ExcitationCurrent(x float64, y float64, z float64, a float64, beta float64, moment float64) error {
	return n.errWrap(C.nec_excitation_current(n.necContext, C.double(x), C.double(y), C.double(z), C.double(a), C.double(beta), C.double(moment)))
}

// ExcitationPlanewave makes a linear polarized planewave excitation source. It
// is one of the simpler versions of ExCard().
//
// Parameters:
//	nTheta - Number of theta angles desired for the incident plane wave .
//	nPhi - Number of phi angles desired for the incident plane wave.
//	theta - Theta in degrees. Theta 19 defined in standard spherical coordinates as illustrated
//	phi - Phi in degrees. Phi is the standard spherical angle defined lned in the XY plane.
//	eta - Eta in degrees. Eta is the polarization angle defined as the angle
// 	between the theta unit vector and the direction of the electric field
// 	for linear polarization or the major ellipse axis for elliptical
// 	polarization.
//	dTheta - Theta angle stepping increment in degrees.
//	dPhi - Phi angle stepping increment in degrees.
//	polRatio - Ratio of minor axis to major axis for elliptic polarization
// 	(major axis field strength - 1 V/m).
//
// Only one incident plane wave or one elementary current source is allowed at
// a time. Also plane-wave or current-source excitation is not allowed with
// voltage sources.  If the excitation types are mixed, the program will use the
// last excitation type encountered.
func (n *NecppCtx) ExcitationPlanewave(nTheta int, nPhi int, theta float64, phi float64, eta float64, dTheta float64, dPhi float64, polRatio float64) error {
	return n.errWrap(C.nec_excitation_planewave(n.necContext, C.int(nTheta), C.int(nPhi), C.double(theta), C.double(phi), C.double(eta), C.double(dTheta), C.double(dPhi), C.double(polRatio)))
}

// TlCard, presumably, makes an NEC2 TL Card.
func (n *NecppCtx) TlCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_tl_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

// NtCard, presumably, makes an NEC2 NT Card.
func (n *NecppCtx) NtCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_nt_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

// XqCard causes program execution at points in the data stream where execution
// is not automatic. Options on the card also allow for automatic generation of
// radiation patterns in either of two vertical cuts.
//
// Parameter:
// 	itmp1 - an ExecutionOption flag, per the ExecutionOption consts.
func (n *NecppCtx) XqCard(itmp1 ExecutionOption) error {
	return n.errWrap(C.nec_xq_card(n.necContext, C.int(itmp1)))
}

// GdCard, presumably, makes a GD card.
func (n *NecppCtx) GdCard(tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64) error {
	return n.errWrap(C.nec_gd_card(n.necContext, C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4)))
}

// simulation output

// RpCard calculates the radiation patterns for the antenna.
//
// Parameters:
// 	calcMode - a RpCalcMode flag. See the RpCalcMode constants for those
// 	definitions.
// 	nTheta - the number of theta angles
// 	nPhi - the number of phi angles
// 	outputFormat - a RpOutputFormat flag. See the RpOutputFormat constants
// 	for those definitions.
// 	normalization - a RpNormalization flag. See the RpNormalization
// 	constants for those definitions.
// 	d - a RpGain flag, either PowerGain or DirectiveGain.
// 	a - a RpAveraging flag. See the RpAveraging constants for those
// 	definitions.
// 	theta0 - Initial theta angle in degrees (initial z coordinate in meters
// 	if calc_mode = 1).
// 	phi0 - Initial phi angle in degrees.
// 	deltaTheta - Increment for theta in degrees (increment for z in meters
// 	if calc_mode = 1).
// 	deltaPhi - Increment for phi in degrees.
// 	radialDistance - Radial distance (R) of field point from the origin in
// 	meters. radial_distance is optional. If it is zero, the radiated
// 	electric field will have the factor exp(-jkR)/R omitted. If a value of
// 	R is specified, it should represent a point in the far-field region
// 	since near components of the field cannot be obtained with an RP card.
// 	(If calc_mode = 1, then radial_distance represents the cylindrical
// 	coordinate phi in meters and is not optional. It must be greater than
// 	about one wavelength.)
// 	gainNorm - Determines the gain normalization factor if normalization has
// 	been requested in the normalization parameter. If gain_norm is zero,
// 	the gain will be normalized to its maximum value. If gain_norm is not
// 	zero, the gain wi11 be normalized to the value of gain_norm.
//
// The field point is specified in spherical coordinates (R, sigma, theta),
// except when the surface wave is computed. For computing the surface wave
// field (calcMode = l), cylindrical coordinates (phi, theta, z) are used to
// accurately define points near the ground plane at large radial distances.
//
// The RpCard() function allows automatic stepping of the field point to
// compute the field over a region about the antenna at uniformly spaced points.
//
// The integers nTheta and nPhi, and floating point numbers theta0, phi0,
// deltaTheta, deltaPhi, radialDistance, and gainNorm control the field-point
// stepping.
//
// The RpCard() method will cause the interaction matrix to be computed and
// factored and the structure currents to be computed if these operations have
// not already been performed. Hence, all required input parameters must be set
// before the RpCard() method is called.
//
// At a single frequency, any number of RpCard() calls may occur in sequence so
// that different field-point spacings may be used over different regions of
// space. If automatic frequency stepping is being used (i.e., inNfrq on the
// FrCard() method is greater than one), only one RpCard() method will act as
// data inside the loop. Subsequent calls to RpCard() will calculate patterns
// at the final frequency.
//
// When both nTheta and nPhi are greater than one, the angle theta (or Z) will
// be stepped faster than phi.
//
// When a ground plane has been specified, field points should not be requested
// below the ground (theta greater than 90 degrees or Z less than zero.)
func (n *NecppCtx) RpCard(calcMode RpCalcMode, nTheta int, nPhi int, outputFormat RpOutputFormat, normalization RpNormalization, d RpGain, a RpAveraging, theta0 float64, phi0 float64, deltaTheta float64, deltaPhi float64, radialDistance float64, gainNorm float64) error {
	return n.errWrap(C.nec_rp_card(n.necContext, C.int(calcMode), C.int(nTheta), C.int(nPhi), C.int(outputFormat), C.int(normalization), C.int(d), C.int(a), C.double(theta0), C.double(phi0), C.double(deltaTheta), C.double(deltaPhi), C.double(radialDistance), C.double(gainNorm)))
}

// PtCard makes a PT Card for printing of currents. This methods documentation
// needs to be checked against the NEC2 user manual before renaming these
// variables and making a new type for a flag. This is what was in libnecpp.h.
//
// IPTFLG Print control flag, specifies the type of format used in printing segment currents. The options are:
// 	-2 - all currents printed. This it a default value for the program if the card is Omitted.
// 	-1 - suppress printing of all wire segment currents.
// 	0 - current printing will be limited to the segments specified by the next three parameters.
// 	1 - currents are printed by using a format designed for a receiving pattern (refer to output section in this manual Only currents for the segments specified by the next three parameters are printed.
// 	2 - same as for 1 above; in addition, however, the current for one Segment will Cue normalized to its maximum, ant the normalized values along with the relative strength in tB will be printed in a table. If the currents for more than one segment are being printed, only currents from the last segment in the group appear in the normalized table.
// 	3 - only normalized currents from one segment are printed for the receiving pattern case.
//
// IPTAG - Tag number of the segments for which currents will be printed.
//
// IPTAGF - Equal to m, specifies the mth segment of the set of segments having the tag numbers of IPTAG, at which printing of currents starts. If IPTAG is zero or blank, then IPTAGF refers to an absolute segment number. If IPTAGF is blank, the current is printed for all segments.
//
// IPTAGT - Equal to n specifies the nth segment of the set of segments having tag numbers of IPTAG. Currents are printed for segments having tag number IPTAG starting at the m th segment in the set and ending at the nth segment. If IPTAG is zero or blank, then IPTAGF and IPTAGT refer to absoulte segment numbers. In IPTAGT is left blank, it is set to IPTAGF.
func (n *NecppCtx) PtCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int) error {
	return n.errWrap(C.nec_pt_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4)))
}

// PqCard makes a PQ Card. Needs documentation from the NEC2 user manual.
func (n *NecppCtx) PqCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int) error {
	return n.errWrap(C.nec_pq_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4)))
}

// KhCard makes a KH Card. Needs documentation from the NEC2 user manual.
func (n *NecppCtx) KhCard(tmp1 float64) error {
	return n.errWrap(C.nec_kh_card(n.necContext, C.double(tmp1)))
}

// NeCard makes a NE Card. Needs documentation from the NEC2 user manual.
func (n *NecppCtx) NeCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_ne_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

// NhCard makes a NH Card. Needs documentation from the NEC2 user manual.
func (n *NecppCtx) NhCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int, tmp1 float64, tmp2 float64, tmp3 float64, tmp4 float64, tmp5 float64, tmp6 float64) error {
	return n.errWrap(C.nec_nh_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4), C.double(tmp1), C.double(tmp2), C.double(tmp3), C.double(tmp4), C.double(tmp5), C.double(tmp6)))
}

// CpCard makes a CP Card. Needs documentation from the NEC2 user manual.
func (n *NecppCtx) CpCard(itmp1 int, itmp2 int, itmp3 int, itmp4 int) error {
	return n.errWrap(C.nec_cp_card(n.necContext, C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4)))
}

// PlCard makes a PL Card. Needs documentation from the NEC2 user manual.
func (n *NecppCtx) PlCard(ploutputFilename string, itmp1 int, itmp2 int, itmp3 int, itmp4 int) error {
	return n.errWrap(C.nec_pl_card(n.necContext, C.CString(ploutputFilename), C.int(itmp1), C.int(itmp2), C.int(itmp3), C.int(itmp4)))
}

// analysis of output

// Gain gets the gain from a radiation pattern.
//
// Parameters:
// 	freqIndex - The rp_card frequency index. If this parameter is 0, then
// 	the first simulation results are used. Subsequent simulations will store
// 	their results at higher indices.
// 	thetaIndex - The theta index (starting at zero) of the radiation pattern
// 	phiIndex - The phi index (starting at zero) of the radiation pattern
//
// This method returns the gain in db, or -999.0 and an error if no radiation
// pattern had been previously requested.
//
// This function requires a previous RpCard() method to have been called
// (with the gain normalization set to TotalNormalized).
func (n *NecppCtx) Gain(freqIndex int, thetaIndex int, phiIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain(n.necContext, C.int(freqIndex), C.int(thetaIndex), C.int(phiIndex)))
}

// GainMax gets the maximum gain from a radiation pattern.
//
// Parameters:
// 	freqIndex - The rp_card frequency index. If this parameter is 0, then
// 	the first simulation results are used. Subsequent simulations will store
// 	their results at higher indices.
//
// This method returns the gain in db, or -999.0 and an error if no radiation
// pattern had been previously requested.
//
// This function requires a previous RpCard() method to have been called
// (with the gain normalization set to TotalNormalized).
func (n *NecppCtx) GainMax(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_max(n.necContext, C.int(freqIndex)))
}

// GainMin gets the minimum gain from a radiation pattern.
//
// Parameters:
// 	freqIndex - The rp_card frequency index. If this parameter is 0, then
// 	the first simulation results are used. Subsequent simulations will store
// 	their results at higher indices.
//
// This method returns the gain in db, or -999.0 and an error if no radiation
// pattern had been previously requested.
//
// This function requires a previous RpCard() method to have been called
// (with the gain normalization set to TotalNormalized).
func (n *NecppCtx) GainMin(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_min(n.necContext, C.int(freqIndex)))
}

// GainMean gets the mean gain from a radiation pattern.
//
// Parameters:
// 	freqIndex - The rp_card frequency index. If this parameter is 0, then
// 	the first simulation results are used. Subsequent simulations will store
// 	their results at higher indices.
//
// This method returns the gain in db, or -999.0 and an error if no radiation
// pattern had been previously requested.
//
// This function requires a previous RpCard() method to have been called
// (with the gain normalization set to TotalNormalized).
func (n *NecppCtx) GainMean(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_gain_mean(n.necContext, C.int(freqIndex)))
}

// GainSd gets the standard deviation of the gain from a radiation pattern.
//
// Parameters:
// 	freqIndex - The rp_card frequency index. If this parameter is 0, then
// 	the first simulation results are used. Subsequent simulations will store
// 	their results at higher indices.
//
// This method returns the gain in db, or -999.0 and an error if no radiation
// pattern had been previously requested.
//
// This function requires a previous RpCard() method to have been called
// (with the gain normalization set to TotalNormalized).
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

func (n *NecppCtx) impedanceReal(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_impedance_real(n.necContext, C.int(freqIndex)))
}

func (n *NecppCtx) impedanceImag(freqIndex int) (float64, error) {
	return n.gainErrWrap(C.nec_impedance_imag(n.necContext, C.int(freqIndex)))
}

// Impedance gets the impedance of the antenna. It returns a complex128 number,
// and takes the place of two separate C library functions that returned the
// real and imaginary portions of the impedance, respectively.
func (n *NecppCtx) Impedance(freqIndex int) (complex128, error) {
	r, rerr := n.impedanceReal(freqIndex)
	i, ierr := n.impedanceImag(freqIndex)
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
