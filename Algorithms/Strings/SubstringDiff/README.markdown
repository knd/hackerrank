## Substring Diff

Given two strings of length n, P = p1...pn and Q = q1..qn, we define M(i,j,k) as the number of mismatches between p(i), p(i+1), ...p(i+k-1) and q(j), q(j+1)...,q(j+k-1). That is in set notation, M(i,j,k) refers to the size of the set

        {0<=x<k,|p(i+x)|b    |q(j+x)|}

Given an integer K, your task is to find the maximum length L, such that, there exists pair of indices (i,j) for which we have M(i, j, L) <= K. Of course, we should also have i+L-1 <=n and j+L-1 <=n.

**Input**

The first line of input contains a single integer, T (1 <=T <=10). T test cases follow. 
Each test case consists of an integer, K, and two strings P and Q separated by a single space.B

**Output**

For each test case output a single integer L which is the maximum value for which there exists pair of indices (i,j) such that M(i, j, L) <=K.

**Constraints**

0 <= K <= length of the string P 
Both P & Q would have the same length 
The size of each of the string would be at the max 1500 
All characters in P & Q are lower-case English letters.

**Sample Input**

        3 
        2 tabriz torino 
        0 abacba abcaba 
        3 helloworld yellomarin

**Sample Output**

        4 
        3 
        8B

**Explanation:**

* First test-case: If we take "briz" from the first string, and "orin" from the second string, then the number of mismatches between these two substrings is equal to 2, and the length of these substrings are 4. Hence we have chosen i=3, j=2, L=4, and we have M(3,2,4) = 2.

* Second test-case: Since K=0, we should find the longest common substring for the given input strings. We can choose "aba" as the result, and we don't have longer common substring between two strings. So, the answer is 3 for this test-case. That's we have chosen i=1, j=4, and L=3, and we have M(1,4,3)=0.

* Third test-case: We can choose "hellowor" from first string and "yellomar" from the second string. So, we have chosen i=1, j=1, and L=8, and we have M(1,1,8)=3. Of course we can also choose i=2, j=2, and L=8 and we still have M(2,2,8)=3.