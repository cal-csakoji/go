// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=loop -id nextNode -d VecType=nextNodeVec -d Type=nextNode github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loop

import (
	"github.com/platinasystems/go/elib"
)

type nextNodeVec []nextNode

func (p *nextNodeVec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]nextNode, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *nextNodeVec) validate(new_len uint, zero *nextNode) *nextNode {
	c := elib.Index(cap(*p))
	lʹ := elib.Index(len(*p))
	l := elib.Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l >= lʹ {
			*p = (*p)[:l]
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *nextNodeVec) validateSlowPath(zero *nextNode,
	c, l, lʹ elib.Index) *nextNode {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]nextNode, cNext, cNext)
		copy(q, *p)
		if zero != nil {
			for i := c; i < cNext; i++ {
				q[i] = *zero
			}
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *nextNodeVec) Validate(i uint) *nextNode {
	return p.validate(i+1, (*nextNode)(nil))
}

func (p *nextNodeVec) ValidateInit(i uint, zero nextNode) *nextNode {
	return p.validate(i+1, &zero)
}

func (p *nextNodeVec) ValidateLen(l uint) (v *nextNode) {
	if l > 0 {
		v = p.validate(l, (*nextNode)(nil))
	}
	return
}

func (p *nextNodeVec) ValidateLenInit(l uint, zero nextNode) (v *nextNode) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p nextNodeVec) Len() uint { return uint(len(p)) }
