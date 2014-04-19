import java.util.*;
import java.lang.Math;

class Solution {

    public static void main(String [] args) {
        
        Scanner in = new Scanner(System.in);
        int T = in.nextInt();
        long D, M;
        
    }

    public static class Node {
        public long runTime;
        public long deadLine;
        public Node left;
        public Node right;
        
        public Node( long runTime, long deadLine ) {
            this.runTime = runTime;
            this.deadLine = deadLine;
        }

        public Node( Node n ) {
            this.runTime = n.runTime;
            this.deadLine = n.deadLine;
        }
    }

    public static class BinaryTree {
        public Node root;

        public BinaryTree() {
            root = null;
        }

        public void insert( Node n ) {
            if ( root == null ) {
                root = new Node(n);
                return;
            }
            Node tmp = root;
            while ( tmp != null ) {
                if ( n.deadLine < tmp.deadLine ) {
                    if ( tmp.left == null ) {
                        tmp.left = new Node(n);
                        return;
                    } else {
                        tmp = tmp.left;
                    }
                } else {
                    if ( tmp.right == null ) {
                        tmp.right = new Node(n);
                        return;
                    } else {
                        tmp = tmp.right;
                    }
                }
            }
        }

        public Node remove( Node n ) {
            Node tmp = root;
            if ( tmp == null ) {
                return null;
            }
            while ( tmp != null ) {
                if ( n.deadLine < tmp.deadLine ) {
                    tmp = tmp.left;
                } else if ( n.deadLine > tmp.deadLine ) {
                    tmp = tmp.right;
                } else if ( tmp.left != null && tmp.right != null ) {
                    Node min = findMin( tmp.right );
                    tmp.runTime = min.runTime;
                    tmp.deadLine = min.deadLine;
                    tmp.right = removeMin(tmp.right);
                } else {
                    
                }
            }
            return tmp;
        }

        public Node getNode( Node n ) {
            Node tmp = root;
            while ( tmp != null ) {
                if ( n.deadLine < tmp.deadLine ) {
                    tmp = tmp.left;
                } else if ( n.deadLine > tmp.deadLine ) {
                    tmp = tmp.right;
                } else {
                    return tmp;
                }
            }
            return null;
        }

        public Node findMin( Node n ) {
            if ( n != null ) {
                while ( n.left != null ) {
                    n = n.left;
                }
            }
            return n;
        }

        public Node removeMin( Node n ) {
            if ( n == null ) {
                return null;
            } else if ( n.left != null ) {
                n.left = removeMin( n.left );
                return n;
            } else {
                return n.right;
            }
        }
    }


}
