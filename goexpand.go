package goexpand

const (
    L1 = 1024
)

type Size uint

type Expand struct {
    slen         Size
    new_, old_   Size
    capa, offset Size
    incr         float64
}

func NewExpand() *Expand {
    return &Expand{}
}

func (e *Expand) Expand(expanded bool, needs Size) {
    if expanded == true {
        e.expand(needs)
    }
}

func (e *Expand) expand(needs Size) {
    e.incr = 0.25

    if needs > 2 * e.old_ {
        e.capa = needs
    } else if e.old_ < L1 {
        e.capa = e.old_ * 2
    } else if e.old_ >= L1 {
        for i := e.old_; i < e.slen; {
            e.capa = e.old_ * Size(1 + e.incr)
        }
    } else {
        if e.capa > e.offset {
            e.capa = needs
        }
    }
}
