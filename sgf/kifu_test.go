package sgf

import (
	"testing"
)

func TestKifu_ToSGF(t *testing.T) {
	kifu := ParseSgf("(;GM[1]PB[]BR[7.2D]PW[]WR[5.4D]RE[暂无结果]KM[7.5]HA[0]TM[600]DT[2017-12-28]GN[标准区即时]SO[弈客围棋]OT[3/0.5]FF[4]CA[UTF-8]RU[zh]AP[WGo.js:2]BL[600]WL[600]SZ[19]BL[600]WL[600];B[dd];W[pp];B[cp];W[pc];B[pe];W[qe])")
	kifu.Last()
	sAll := kifu.ToCurSgf()
	if sAll == "(;SZ[19]KM[7.5]HA[0];B[dd];W[pp];B[cp];W[pc];B[pe];W[qe])" {
		t.Log(sAll)
	} else {
		t.Error("kifu to sgf all move fail")
	}

}
func TestKifu_Play(t *testing.T) {
	kifu:=ParseSgf("(;SZ[19]AP[WGo.js:2]FF[4]GM[1]CA[UTF-8];B[ba];W[rn];B[ab])")
	kifu.Last()
	node:=Node{
		X:0,
		Y:0,
		C:W,
	}
	result:=kifu.Play(node)
	if result{
		t.Error("move error")
	}
}

