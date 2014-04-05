Median
======

The problem is located at this [link](https://www.interviewstreet.com/challenges/dashboard/#problem/4fcf919f11817). Enjoy! I had
enough fun. I got 8/10 test cases passed for this problem. A few things I learned and watched out:

* There could potentially be `long` integers. So use `long` instead of `int` as needed.
* 32-bit integer also means **negative** numbers. This was my mistake for not carefully examining the boundaries.
I jumped from 3 to 7 test cases passed for this fix.
* `PriorityQueue` in Java is convenient but not enough. I wanted to use `remove(Object e)` but the operation takes linear time. Dang!
* `TreeMap` and `TreeSet` obviously don't accept duplicate keys. C++ does with `multiset`, what a shame! There was nothing I could
leverage here. 
* `TreeBag` from [Apache Common Collections](http://commons.apache.org/collections/userguide.html) seems helpful here. However, 
I wanted to do Interview Street problems with just core Java utilities. 
* There are 3 basic operations for this problem on the high-level view. Those are **add**, **remove**, and **median**. I was able
to achieve **O(log n)** for add, **0(n)** for remove (optimized a bit to pass 1 more test case so I could say **O(n/2)** :p), and
**O(1)** for median. There were a several experiments with data structures that I tried. 

I decided to stop after 8 test cases passed. `Solution1.java` was one of the initial failed sketches.
