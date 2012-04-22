Information:
--------------

### Scenario: 
-------------
This **Fraud Prevention** problem is located at [Interview Street] (http://cs2.interviewstreet.com/recruit/challenges#problems).
I paid attention to and solved it after a job interview with a [Groupon] () engineer who is
working on this fraud matter. I am not going tell here how the interview was
though. In order to view the problem, you should sign in [Interview Street] (http://cs2.interviewstreet.com/recruit/challenges#problems) 
with your Google/Facebook/Twitter account since I am not going to describe the
problem here.

### Result:
-------------
I wish they had provided a greater number of test cases.

![correctResult.png] (http://dl.dropbox.com/u/54031414/CodeChallenges/FraudPrevention/correctResult.png)


### What I learned:
------------------
* **Java Regex**. Initially, I used `String.split("+")` and here is the result:
![firstResult.png] (http://dl.dropbox.com/u/54031414/CodeChallenges/FraudPrevention/firstResult.png)

I went to and read some docs about **Java Regex**. `String.split("\\+")` is the
correct way.

* **Difference between Hashtable methods: contains(Object value),
containsKey(Object key), and containsValue(Object value)**. I should have not
made any assumption before using them interchangeably. Each has different
purpose. I certainly got wrong result.

* **Odd piece of code** that I am still not sure why. If I were to combine the two lines of code 81, 82:

    emailKey = emailKey.replace(".", "");
    return emailKey;

into:

    return emailKey.replace(".", "");

I would get the result:
![oddResult.png] (http://dl.dropbox.com/u/54031414/CodeChallenges/FraudPrevention/oddResult.png)

* **Check empty string**. `String.equals("")` is definitely faster than `String.length() > 0`

* **StringTokenizer**. It is faster than `String.split(Object delimiter)` if a long string is encountered. The solution code I presented here has not been implemented with `StringTokenizer` yet.


    
    



