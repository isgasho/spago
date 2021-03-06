// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import "github.com/nlpodyssey/spago/pkg/mat"

var _ Function = &Identity{}

// Identity function.
// y = x
type Identity struct {
	x Operand
}

func NewIdentity(x Operand) *Identity {
	return &Identity{x: x}
}

// Forward computes the output of the function.
func (r *Identity) Forward() mat.Matrix {
	return r.x.Value().Clone()
}

func (r *Identity) Backward(gy mat.Matrix) {
	if !(mat.SameDims(r.x.Value(), gy) || mat.VectorsOfSameSize(r.x.Value(), gy)) {
		panic("fn: matrices with not compatible size")
	}
	r.x.PropagateGrad(gy)
}
