# Protoform
Please feel free to submit issues / PR's if there are bugs (There almost certainly are many).
Thanks!


Convert Java POJO's to [Protobuf](https://developers.google.com/protocol-buffers/) with super hacky regexes

# What this project is able to do today:

* Convert simple Java POJO's to proto, for example:

given a Java source file:
```java
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

```

protoform will generate

```proto
syntax = "proto3";
package smartlabel;

message message {
	string aa = 1;
	string bb = 2;
	string cc = 3;
	string dd = 4;
	bool ee = 5;
}
```

# Building
`$ make build`

# Usage
`$ ./protoform --in-file="$SOME_JAVA_FILE"`  
Will take the java file and spit out the resulting proto in the working directory

`$ ./protoform --in-file="$SOME_JAVA_FILE" --out-file=proto`  
Will take the java file and spit out the resulting proto in the directory `proto/`

`$ ./protoform --in-file="$SOME_JAVA_FILE" --out-file=proto --package="$SOME_PACKAGE"`  
Will take the java file and spit out the resulting proto in the directory `proto/"$SOME_PACKAGE"`


# Current Limitations
* Map Types are not fully supported and in most cases will not parse correctly
* Unable to follow inheritance Trees
* Does not generate imports for files with user created data types
* It is likely there are types that exist which are not supported

With the above limitations in mind I will gladly accept PR's / help with refactors/cleanups/better handling of specific cases.
