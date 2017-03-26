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

const GainErrno float64 = -999.0
var ErrNoPatternRequested = errors.New("no radiation pattern previously requested")


type NecppCtx struct {
	necContext *C.nec_context
}

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

func (n *NecppCtx) errWrap(ret C.long) error {
	if ret != 0 {
		err := n.ErrorMessage()
		return err
	}
	return nil
}

func (n *NecppCtx) gainErrWrap(gain C.double) (float64, error) {
	gainRet := float64(gain)
	if gainRet == GainErrno {
		return gainRet, ErrNoPatternRequested
	}
	return gainRet, nil
}

func (n *NecppCtx) Delete() error {
	err := n.errWrap(C.nec_delete(n.necContext))
	if err != nil {
		return err
	}
	return nil
}

func (n *NecppCtx) ErrorMessage() error {
	// doesn't look like this should be freed
	cerr := C.nec_error_message()
	err := errors.New(C.GoString(cerr))
	return err
}

// antenna geometry methods

func (n *NecppCtx) Wire(tagId int, segmentCount int, xw1 float64, yw1 float64, zw1 float64, xw2 float64, yw2 float64, zw2 float64, rad float64, rdel float64, rrad float64) error {
	err := n.errWrap(C.nec_wire(n.necContext, C.int(tagId), C.int(segmentCount), C.double(xw1), C.double(yw1), C.double(zw1), C.double(xw2), C.double(yw2), C.double(zw2), C.double(rad), C.double(rdel), C.double(rrad)))
	if err != nil {
		return err
	}
	return nil
}

func (n *NecppCtx) SpCard(ns int, x1 float64, y1 float64, z1 float64, x2 float64, y2 float64, z2 float64) error {
	err := n.errWrap(C.nec_sp_card(n.necContext, C.int(ns), C.double(x1), C.double(y1), C.double(z1), C.double(x2), C.double(y2), C.double(z2)))
	if err != nil {
		return err
	}
	return nil
}

func (n *NecppCtx) ScCard(i2 int, x3 float64, y3 float64, z3 float64, x4 float64, y4 float64, z4 float64) error {
	err := n.errWrap(C.nec_sc_card(n.necContext, C.int(i2), C.double(x3), C.double(y3), C.double(z3), C.double(x4), C.double(y4), C.double(z4)))
	if err != nil {
		return err
	}
	return nil
}

func (n *NecppCtx) GmCard(itsi int, nrpt int, rox float64, roy float64, roz float64, xs float64, ys float64, zs float64, its int) error {
	return n.errWrap(C.nec_gm_card(n.necContext, C.int(itsi), C.int(nrpt), C.double(rox), C.double(roy), C.double(roz), C.double(xs), C.double(ys), C.double(zs), C.int(its)))
}

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
