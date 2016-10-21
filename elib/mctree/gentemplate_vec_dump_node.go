// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=mctree -id dump_node -d VecType=dump_node_vec -d Type=dump_node github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mctree

import (
	"github.com/platinasystems/go/elib"
)

type dump_node_vec []dump_node

func (p *dump_node_vec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]dump_node, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *dump_node_vec) validate(new_len uint, zero *dump_node) *dump_node {
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

func (p *dump_node_vec) validateSlowPath(zero *dump_node,
	c, l, lʹ elib.Index) *dump_node {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]dump_node, cNext, cNext)
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

func (p *dump_node_vec) Validate(i uint) *dump_node {
	return p.validate(i+1, (*dump_node)(nil))
}

func (p *dump_node_vec) ValidateInit(i uint, zero dump_node) *dump_node {
	return p.validate(i+1, &zero)
}

func (p *dump_node_vec) ValidateLen(l uint) (v *dump_node) {
	if l > 0 {
		v = p.validate(l, (*dump_node)(nil))
	}
	return
}

func (p *dump_node_vec) ValidateLenInit(l uint, zero dump_node) (v *dump_node) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p dump_node_vec) Len() uint { return uint(len(p)) }
