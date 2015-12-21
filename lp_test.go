// Package golp gives Go bindings for LPSolve.
package golp

import (
	"testing"
)

// TestLP tests a real-valued linear programming example
func TestLP(t *testing.T) {
	lp := NewLP(0, 2)
	lp.SetVerboseLevel(NEUTRAL)
	lp.SetColName(0, "x")
	lp.SetColName(1, "y")
	if "x" != lp.ColName(0) {
		t.Errorf("x != %s", lp.ColName(0))
	}
	if "y" != lp.ColName(1) {
		t.Errorf("y != %s", lp.ColName(1))
	}

	lp.AddConstraint([]float64{120.0, 210.0}, LE, 15000)
	lp.AddConstraintSparse([]Entry{Entry{Col: 0, Val: 110.0}, Entry{Col: 1, Val: 30.0}}, LE, 4000)
	lp.AddConstraintSparse([]Entry{Entry{Col: 1, Val: 1.0}, Entry{Col: 0, Val: 1.0}}, LE, 75)

	lp.SetObjFn([]float64{143, 60})
	lp.SetMaximize()

	lpString := "/* Objective function */\nmax: +143 x +60 y;\n\n/* Constraints */\n+120 x +210 y <= 15000;\n+110 x +30 y <= 4000;\n+x +y <= 75;\n"
	if lpString != lp.WriteToString() {
		t.Errorf("lpString!=%s, but %s", lpString, lp.WriteToString())
	}

	lp.Solve()

	delta := 0.000001
	dt := 6315.625 - lp.Objective()
	if dt < -delta || dt > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", 6315.625, lp.Objective(), delta, dt)
	}

	vars := lp.Variables()
	if len(vars) != 2 {
		t.Errorf("len(vars) != 2, but %v", len(vars))
	}
	dt = 21.875 - vars[0]
	if dt < -delta || dt > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", 21.875, vars[0], delta, dt)
	}
	dt = 53.125 - vars[1]
	if dt < -delta || dt > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", 53.125, vars[1], delta, dt)
	}
}

// TestMIP tests a mixed-integer programming example
func TestMIP(t *testing.T) {
	lp := NewLP(0, 4)
	lp.AddConstraintSparse([]Entry{{0, 1.0}, {1, 1.0}}, LE, 5.0)
	lp.AddConstraintSparse([]Entry{{0, 2.0}, {1, -1.0}}, GE, 0.0)
	lp.AddConstraintSparse([]Entry{{0, 1.0}, {1, 3.0}}, GE, 0.0)
	lp.AddConstraintSparse([]Entry{{2, 1.0}, {3, 1.0}}, GE, 0.5)
	lp.AddConstraintSparse([]Entry{{2, 1.0}}, GE, 1.1)
	lp.SetObjFn([]float64{-1.0, -2.0, 0.1, 3.0})

	lp.SetInt(2, true)
	if !lp.IsInt(2) {
		t.Errorf("Col 2 is Not Int!")
	}

	lp.Solve()

	delta := 0.000001
	dt := -8.133333333 - lp.Objective()
	if dt < -delta || dt > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", -8.133333333, lp.Objective(), delta, dt)
	}

	vars := lp.Variables()
	if lp.NumCols() != 4 {
		t.Errorf("NumCols != 4, but %d", lp.NumCols())
	}
	if len(vars) != 4 {
		t.Errorf("len(vars) != 4, but %d", len(vars))
	}
	dt = 1.6666666666 - vars[0]
	if dt < -delta || dt > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", 1.6666666666, vars[0], delta, dt)
	}
	dt = 3.3333333333 - vars[1]
	if dt < -delta || dt > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", 3.3333333333, vars[1], delta, dt)
	}
	dt = 2.0 - vars[2]
	if dt < -delta || dt > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", 2.0, vars[2], delta, dt)
	}
	dt = 0.0 - vars[3]
	if dt < -delta || dt > delta {
		t.Errorf("Max difference between %v and %v allowed is %v, but difference was %v", 0.0, vars[3], delta, dt)
	}
}
