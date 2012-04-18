"""
Copyright (C) 2012  Khanh Dao

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

from time import clock, time
import heapq
"""

POINTS_TABLE = { 'A': 1, 'E': 1, 'I': 1, 'L': 1, 'N': 1, 'O': 1, 'R': 1, 'S': 1,
           'T': 1, 'U': 1, 'D': 2, 'G': 2, 'B': 3, 'C': 3, 'M': 3, 'P': 3,
           'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4, 'K': 5, 'J': 8, 'X': 8,
           'Q': 10, 'Z': 10 }

class WordNode(object):

    def __init__(self, word):
        self.word = word
        self.score = 0
        for letter in word:
            self.score += POINTS_TABLE[letter]

    def getWord(self):
        return self.word

    def getScore(self):
        return self.score

    def __str__(self):
        return self.word

class Edge(object):

    def __init__(self, src, dst):
        self.src = src
        self.dst = dst

    def getSource(self):
        return self.src

    def getDestination(self):
        return self.dst

    def __str__(self):
        return self.src + ' -> ' + self.dst

class Graph(object):

    def __init__(self):
        self.nodes = set([])
        self.edges = {}
        self.depths = {}
        self.largest_paths = {}
        self.largest_ladders = {}

    def addNode(self, word_node):
        if word_node.getWord() in self.nodes:
            raise ValueError('Duplicate')
        else:
            self.nodes.add(word_node)
            self.edges[word_node] = set([])
            self.depths[word_node] = [0]
            self.largest_paths[word_node] = { 0 :  (word_node.score, []) }
            self.largest_ladders[word_node] = {0 : (True, word_node.score)}
            

    def addEdge(self, edge):
        src = edge.getSource()
        dst = edge.getDestination()
        self.edges[src].add(dst)

    def getNodes(self):
        return self.nodes

    def getEdges(self):
        return self.edges

    def getChildren(self, word_node):
        return self.edges[word_node]

    def hasWordNode(self, word_node):
        return word_node in self.nodes

    def getNodeDepths(self, word_node):
        return self.depths[word_node]

    def getLargestPath(self, word_node):
        return self.largest_paths[word_node]

    def getLargestLadders(self, word_node):
        return self.largest_ladders[word_node]

    def __str__(self):
        res = ''
        for node in self.nodes:
            res += str(node) + '\n'
            for depth in self.depths[node]:
                res += '{' + str(depth) + ': ' + str(self.largest_paths[node][depth][0]) + ', ['
                for n in self.largest_paths[node][depth][1]:
                    res += str(n) + ', '
                res = res + '] }' + '\n'
                        
        res += '\n'
        for node, neighbors in self.edges.iteritems():
            for n in neighbors:
                res += str(node) + '->' +  str(n) + '\n'
        return res[:-1]

def is1CharDiff(word1, word2):
    diff_count = 0
    for l1, l2 in zip(word1, word2):
        if l1 != l2:
            diff_count += 1
            if diff_count > 1:
                return False
    return True


def addToGraph(word_graph, node_lst):
    stack_node = sorted(node_lst, key=lambda x: x.score)

    while stack_node:
        node_src = stack_node.pop()
        word_graph.addNode(node_src)
        for node_dst in stack_node:
            if node_src.getScore() != node_dst.getScore() and \
               is1CharDiff(node_src.getWord(), node_dst.getWord()):
                edge = Edge(node_src, node_dst)
                word_graph.addEdge(edge)

def explore(node, depth, word_graph):
    if depth == 0:
        return True

    children = word_graph.getChildren(node)
    if not children:
        return False

    max_path = (0, [])
    for child in children:
        child_largest_path = word_graph.getLargestPath(child)
        if (depth-1) in child_largest_path or explore(child, depth-1, word_graph):
            child_largest_score = child_largest_path[depth-1][0]
            child_path = child_largest_path[depth-1][1]

            node_largest_score = child_largest_score + node.getScore()
            if node_largest_score > max_path[0]:
                max_path = (node_largest_score, [child] + child_path)
    if max_path[0] == 0:
        return False
    
    node_largest_path = word_graph.getLargestPath(node)
    node_largest_path[depth] = (max_path[0], max_path[1])
    return True

def calDepth(word_graph):
    for node in word_graph.getNodes():
        for depth in range(51):
            node_largest_path = word_graph.getLargestPath(node)
            if depth in node_largest_path:
                continue
            if explore(node, depth, word_graph): 
                word_graph.getNodeDepths(node).append(depth)
            else:
                break


def reExplore(node, depth, visited, word_graph):
    if depth == 0:
        return node.getScore()
    if node in visited:
        return None
    score = 0
    children = word_graph.getChildren(node)
    for child in children:
        if not child in visited:
            res_reExplore = reExplore(child, depth-1, visited, word_graph)
            if res_reExplore is not None:
                if (res_reExplore + node.getScore()) > score:
                    score = res_reExplore + node.getScore()
    if score != 0:
        return score
    else:
        return None

def calLadders(word_graph):
    max_score_ladder = 0
    for node in word_graph.getNodes():
        word_ladders = word_graph.getLargestLadders(node)
        children = word_graph.getChildren(node)
        depth_lst = word_graph.getNodeDepths(node)
        while depth_lst:
            depth = depth_lst.pop()
            if depth == 0:
                if word_ladders[depth][1] > max_score_ladder:
                    max_score_ladder = word_ladders[depth][1]
                continue
            if len(children) <= 1:
                word_ladders[depth] = (False)
                continue
            if len(children) > 1:
                score_ladder = calLaddersHelper(word_ladders, node, children, depth, word_graph)
                if score_ladder is None:
                    word_ladders[depth] = (False)
                else:
                    word_ladders[depth] = (True, score_ladder)
                    if score_ladder > max_score_ladder:
                        max_score_ladder = score_ladder
                        break
    return max_score_ladder 
                   

def calLaddersHelper(word_ladders, node, children, depth, word_graph):
    max_score = 0
    for childA in children:
        if not (depth-1) in word_graph.getLargestPath(childA):
            continue
        for childB in children:
            if childA is not childB:
                if not (depth-1) in word_graph.getLargestPath(childB):
                    continue
                childA_path_score, childA_path = word_graph.getLargestPath(childA)[depth-1]
                childB_path_score, childB_path = word_graph.getLargestPath(childB)[depth-1]
                pathA = [childA] + childA_path
                pathB = [childB] + childB_path
                intersection = set(filter(lambda x: x in pathB, pathA))
                if intersection:
                    new_childB_path_score = reExplore(childB, depth-1, intersection, word_graph)
                    if new_childB_path_score is not None:
                        if max_score < (childA_path_score + new_childB_path_score):
                            max_score = childA_path_score + new_childB_path_score
                else:
                    if max_score < (childA_path_score + childB_path_score):
                        max_score = childA_path_score + childB_path_score
    if max_score != 0:
        return max_score + node.getScore()
    return None

def getInput():
    k, n, node_lst = 0, 0, []
    i = 0
    while True:
        line = raw_input()
        if i == 0:
            k = int(line)
        elif i == 1:
            n = int(line)
        elif len(line) == k:
            word_node = WordNode(line.upper())
            node_lst.append(word_node)
        i += 1
        if i == n + 2:
            break

    return node_lst

                                
def main():
    node_lst = getInput()
    word_graph = Graph()

    addToGraph(word_graph, node_lst)
    calDepth(word_graph)
    print(calLadders(word_graph))
    

if __name__ == '__main__':
    main()
