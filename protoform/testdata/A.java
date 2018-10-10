package com.foo.a

import java.io.Serializable;

public class A extends B implements Serializable {

    public A() {
    }

    public A(String aa, String bb, String cc, String dd, boolean ee) {
        this.aa = aa;
        this.bb = bb;
        this.cc = cc;
        this.dd = dd;
        this.ee = ee;
    }

}
