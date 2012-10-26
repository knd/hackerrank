Changing Bits
=============

The problem can be found [here](https://www.interviewstreet.com/challenges/dashboard/#problem/4f1c739a6ea3a).
8/11 test cases passed. I decided not to handle low level bit traversing, checking and computation to boost
up the performance because I think it's not **Java oriented**. That should be if someone is to solve in **C** code.
I believe `BigInteger` class should take care of this for me, though. One thing that I learned:

* Use all the bit methods of `BitInteger` such as **testBit**, **clearBit**, and **setBit**, it's super convenient.
* Could optimize solution by this tip [here](http://stackoverflow.com/questions/9338608/java-manipulating-large-bit-numbers).
