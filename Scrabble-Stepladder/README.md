Information:
------------

### Note:
There are 20 people who could solve this problem out of 762 attemps.
Solving means having the solution that passes 10 out of 10 test cases.
I passed 9 out of 10 test cases. However, this is an interesting challenge to remember.

### Initial ideas: `stepladderBruteForce.py`
*   Naive method using stacks and queues. This is actually a homework assignment that I found in a CS course at Carnegie Mellon University. The naive method is suggested in here at [http://www.cs.cmu.edu/~adamchik/15-121/labs/HW-4%20Word%20Ladder/lab.html] (http://www.cs.cmu.edu/~adamchik/15-121/labs/HW-4%20Word%20Ladder/lab.html)
*   The result is not too bad. At least it guarantees the correctness for small input sizes 

![result of brute force method] (http://dl.dropbox.com/u/54031414/CodeChallenges/ScrabbleStepladder/bruteForceResult.png)

### Later idea: `stepladder.py`
*   Dynamic programming is a powerful tool. Words are represented by nodes and they contruct a graph with directed edges from higher scores to lower scores. With graph, we can effectively traverse with graph algorithms and have better run time.
*   The result is more promising. At least it passed almost all the test cases 

![result of dynamic prog method] (http://dl.dropbox.com/u/54031414/CodeChallenges/ScrabbleStepladder/finalResult.png)

