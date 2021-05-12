https://github.com/avamsykiran/GoLang_11May2021_28May2021_09001130.git
==========================================================================

IDE:        VS Code     https://code.visualstudio.com/download
Compile:    Go          https://golang.org/dl/

--------------------------------------------------------------

Go  Lang
-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

        GPL

Features

        1. Simple Familiar Syntax.
        2. Concurrent
        3. Resource 
        4. Distributed
        5. Compiler + Linker

Package         is a logical block of code.
                is a go approach to modularization.

Data Types:
-------------------------

    Primitive:
        Numeric:
            uint8   unsigned 8-bit integers (0 to 255)
            uint16  Unsigned 16-bit integers (0 to 65535)
            uint32  Unsigned 32-bit integers (0 to 4294967295)
            uint64  Unsigned 64-bit integers (0 to 18446744073709551615)
            int8    Signed 8-bit integers (-128 to 127)     
            int16   Signed 16-bit integers (-32768 to 32767)	
            int32   Signed 32-bit integers (-2147483648 to 2147483647)
            int64   Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
            
            byte    same as uint8
            rune    same as int32
            uint    32 or 64 bits
            int     same size as uint
            uintptr an unsigned integer 
                    to store the uninterpreted bits of a pointer value

            float32     IEEE-754 32-bit floating-point numbers
            float64     IEEE-754 64-bit floating-point numbers
            complex64   Complex numbers with float32 real and imaginary parts
            complex128  Complex numbers with float64 real and imaginary parts

        Non Numeric:
            bool
            string

Escape Sequences
------------------------------------------------

        \n              new line
        \t              tab space
        \r              carrage return
        \\              escape a \
        \"              escape a "

format specifiers
-------------------------------------------------

        %d
        %ld
        %f
        %c
        %s

Operators
-------------------------------------------------

        Binary
                Arithemtic
                        +
                        -
                        *
                        /
                        %
                Relational 
                        ==
                        <
                        >
                        <=
                        >=
                        !=
                Logical
                        &&
                        ||
                        !
                Assignment Operatos
                        =
                        +=              a+=b;           a = a + b;
                        -=
                        *=
                        /=
                        %=

                        &=
                        |=
                        ^=
                        ~=
        Bitwise
                &               0&0 1&0 0&1             0
                                1&1                     1
                
                |               0|0                     0
                                1|1 0|1 1|0             1

                ^               0^0 1^1                 1
                                0^1 1^0                 0

                ~               ~0                      1
                                ~1                      0

        Unary
                        ++
                        --

                        var x,y,z int;

                        x = 10;

                        y = ++x;    // x->11, y=11
                        z = x++;    // z->11, x->12

        Short hand Declaration Operator

                :=

                var a int = 45;

                var a int;
                a = 45;

                a := 45; //type is infered.

Decision Making Control Structures
---------------------------------------------

                if (bool_exp) {

                }

                if (bool_exp) {
                        //true blockm of code
                }else{
                        //false block of code.
                }

                Switch Statement
                ------------------------------
                                switch(exp){
                                        case val1:
                                                code-blcok-1;
                                        case val2:
                                                code-block-2;
                                        case val3:
                                                code-block-3;

                                        default:
                                                defalut-code-block;
                                }

                                switch {
                                        case exp:
                                                code-1;
                                        case exp2:
                                                code-2;
                                        case exp3:
                                                code-3...;

                                        default:
                                                defualt-code;
                                }

                                switch x.(type) {
                                        case type1:
                                         code...
                                        case type2:
                                         code..
                                        defualt:
                                         default-code..
                                }

                                select {
                                        case exp:
                                                code-1;
                                        case exp2:
                                                code-2;
                                        case exp3:
                                                code-3...;

                                        default:
                                                defualt-code;
                                }
                                

Looping Controls Strucutres
-------------------------------------

                loop goto

                for 

                        Indefinite loop
                                
                                for bool_exp {
                                        ...............
                                }

                        Definite loop

                                for i:=initalVal;bool_exp;inc/dec {
                                        ...............
                                }

Non-Conditional Control Statements

                        break
                        continue
                        goto 

        Functions
        ---------------------------------------------

                sub-programs - self contianed block of code-
                targeting a specific action/operation.

                func funName(param1 type,param2 type..)  {
                        
                }

                func funName(param1 type,param2 type..) returnType {
                        return soemthing;
                }

                func funName(param1 type,param2 type..) (returnType1,returnType2) {
                        return value1,value2;
                }

                Anonymous Functions
                        function with no name
                
                High Order Functions
                        function that return another function

                Closures
                        High order function that can
                        encapsulate a persisting state.
                

        Arrays
        -----------------------------------------------------------

                var arr [size]type

                var nums [10]int

                var nums = [5]int{1,2,3,4,5}

                len(array)              number of eles
                cap(array)              saem as Len

                nums[4] ----------> 5

                range           this is keyword that returns 
                                        the index and ele from the 
                                        givne array until no more eles are left.


        Slice
        ------------------------------------------------------------
                 slice is a small part of an array.

                 s1 := nums[:7]; //s1 contains 0 to 7th eles of the nums array
                 s2 := nums[3:]; //s2 contains 3rd to last ele of the nums array
                 s3 := nums[4:8] //s3 contains 4th to 8th ele of the nums array


                len(slice)              number of eles in that slice
                cap(slice)              number of eles in the array of te slice, starting from the first
                                                ele in the slice
