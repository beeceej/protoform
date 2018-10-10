# Protoform

Convert Java POJO's to [https://developers.google.com/protocol-buffers/](Protocol Buffers).

## Why Protocol Buffers:
Protocol Buffers are awesome for a couple reasons.
* Language Agnostic
 * Create a protobuf file, compile/generate code for many languages
* Efficient Serialization (`[]byte` over the wire)
  * Protobuf bakes in serialization, and the code it generates gives you everything needed to work with the format
For more info, read here. [https://developers.google.com/protocol-buffers/docs/overview](https://developers.google.com/protocol-buffers/docs/overview)


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
`$ ./pfrm --in-file="$SOME_JAVA_FILE"`  
Will take the java file and spit out the resulting proto in the working directory

`$ ./pfrm --in-file="$SOME_JAVA_FILE" --out-file=proto`  
Will take the java file and spit out the resulting proto in the directory `proto/`

`$ ./pfrm --in-file="$SOME_JAVA_FILE" --out-file=proto --package="$SOME_PACKAGE"`  
Will take the java file and spit out the resulting proto in the directory `proto/"$SOME_PACKAGE"`


# Current Limitations
* Map Types are not fully supported and in most cases will not parse correctly
* Unable to follow inheritance Trees
* Does not generate imports for files with user created data types
* It is likely there are types that exist which are not supported

With the above limitations in mind I will gladly accept PR's / help with refactors/cleanups/better handling of specific cases.