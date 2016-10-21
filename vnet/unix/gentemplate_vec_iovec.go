// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=unix -id iovec -d VecType=iovecVec -d Type=iovec github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/platinasystems/go/elib"
)

type iovecVec []iovec

func (p *iovecVec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]iovec, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *iovecVec) validate(new_len uint, zero *iovec) *iovec {
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

func (p *iovecVec) validateSlowPath(zero *iovec,
	c, l, lʹ elib.Index) *iovec {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]iovec, cNext, cNext)
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

func (p *iovecVec) Validate(i uint) *iovec {
	return p.validate(i+1, (*iovec)(nil))
}

func (p *iovecVec) ValidateInit(i uint, zero iovec) *iovec {
	return p.validate(i+1, &zero)
}

func (p *iovecVec) ValidateLen(l uint) (v *iovec) {
	if l > 0 {
		v = p.validate(l, (*iovec)(nil))
	}
	return
}

func (p *iovecVec) ValidateLenInit(l uint, zero iovec) (v *iovec) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p iovecVec) Len() uint { return uint(len(p)) }
