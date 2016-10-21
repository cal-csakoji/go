// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=loop -id Vi -d VecType=viVec -d Type=Vi github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loop

import (
	"github.com/platinasystems/go/elib"
)

type viVec []Vi

func (p *viVec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]Vi, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *viVec) validate(new_len uint, zero *Vi) *Vi {
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

func (p *viVec) validateSlowPath(zero *Vi,
	c, l, lʹ elib.Index) *Vi {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]Vi, cNext, cNext)
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

func (p *viVec) Validate(i uint) *Vi {
	return p.validate(i+1, (*Vi)(nil))
}

func (p *viVec) ValidateInit(i uint, zero Vi) *Vi {
	return p.validate(i+1, &zero)
}

func (p *viVec) ValidateLen(l uint) (v *Vi) {
	if l > 0 {
		v = p.validate(l, (*Vi)(nil))
	}
	return
}

func (p *viVec) ValidateLenInit(l uint, zero Vi) (v *Vi) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p viVec) Len() uint { return uint(len(p)) }
