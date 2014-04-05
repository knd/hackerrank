Meeting Points
==============

The problem can be found [here](https://www.interviewstreet.com/challenges/dashboard/#problem/4e14b2cc47f78).
All test cases passed. A few things that I learned and watched out:

* Optimize and optimize fron **O(n^2)** down to **O(n)**. The brute-force way won't work because of given contrainsts.
* `long` in *Java* is 64-bit signed integer.
* `double` in *Java* is 64-bit signed floating point.
* Use `BigInteger` for the sum of all the steps from all the other points to meeting point. This number is ridiculously **huge**
and I was so stubborn sticking with `double`. This drove me crazy with 4/13 test cases passed, thinking
that the correctness of the algorithm was invalid.
* `BigInteger` is **immutable**. You can't change it, just only reassign it. Otherwise, you constantly get 0 or whatever
you assign at the beginning.
* Credits are given to these guys on [here](http://stackoverflow.com/questions/7027753/minimum-sum-of-all-travel-times) for
pointers on the **O(n)** solution, which uses [Geometric Median](http://en.wikipedia.org/wiki/Geometric_median) 
and [Chebyshev Distance](http://en.wikipedia.org/wiki/Chebyshev_distance).
