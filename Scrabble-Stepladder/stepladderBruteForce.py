import os, sys
from sets import Set
from collections import deque

POINTS = { 1: ['A', 'E', 'I', 'L', 'N', 'O', 'R', 'S', 'T', 'U'], 
           2: ['D', 'G'], 
           3: ['B', 'C', 'M', 'P'],
           4: ['F', 'H', 'V', 'W', 'Y'],
           5: ['K'], 
           8: ['J', 'X'],
          10: ['Q', 'Z'] }

def isOneLetterDiff(word_src, word_dst, k):
    diff_num = 0
    for i in range(k):
        if word_src[i] != word_dst[i]:
            diff_num += 1
        if diff_num > 1:
            return False
    return True
        

def getOneLetterDiff(word_src, stack, wordSet, k):
    next_words = []
    for word_dst in wordSet:
        if not word_dst in stack:
            if isOneLetterDiff(word_src, word_dst, k):
                next_words.append(word_dst)
    return next_words

def getWordScore(word):
    score = 0
    for letter in word:
        for point, letter_lst in POINTS.iteritems():
            if letter in letter_lst:
                score += point
                break
    return score


def isStepLadder(word_ladder):
    length = len(word_ladder)
    mid = length / 2
    valid, total_score, curr_score = True, 0, 0
    curr_score = getWordScore(word_ladder[0])
    total_score = curr_score
    for i in range(1,length):
        next_score = getWordScore(word_ladder[i])
        if i <= mid and next_score > curr_score:
            curr_score = next_score
            total_score += curr_score
        elif i > mid and next_score < curr_score:
            curr_score = next_score
            total_score += curr_score
        else:
            return False, total_score
    return valid, total_score


def main():
    k, n, wordSet = 0, 0, Set({})
    i = 0
    while True:
        line = raw_input()
        if i == 0:
            k = int(line)
        elif i == 1:
            n = int(line)
        elif len(line) == k:
            wordSet.add(line)
        i += 1
        if i == n + 2:
	        break  
   
    stack_lst = map(lambda x: [x], wordSet)
    queue = deque(stack_lst)
    word_ladders = []

    while queue:
        stack = queue.popleft()
        word = stack[-1]
        next_words = getOneLetterDiff(word, stack, wordSet, k)
        if not next_words:
            word_ladders.append(stack)
        else:
            for each_word in next_words:
                queue.append(stack + [each_word])

    word_ladders = filter(lambda x: len(x) % 2 != 0, word_ladders)
    highest_score = 0
    for each_ladder in word_ladders:
        valid, total_score = isStepLadder(each_ladder)
        if valid and total_score > highest_score:
            highest_score = total_score
    print(highest_score)


if __name__ == '__main__':
    main()
