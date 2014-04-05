/* Sample program illustrating input and output */

import java.util.*;
import java.text.DecimalFormat;
import java.lang.Math;

class Solution{
	public static void main( String args[] ){
		
// helpers for input/output		

		Scanner in = new Scanner(System.in);
		
		int N;
		N = in.nextInt();
	
		String s[] = new String[N];
		long x[] = new long[N];
		
		for(int i=0; i<N; i++){
			s[i] = in.next();
			x[i] = in.nextLong();
		}

        // solution code starts here
        Queue<Long> lowerHalfMaxHeap = new PriorityQueue<Long>( 1, new MaxComparator() );
        Queue<Long> upperHalfMinHeap = new PriorityQueue<Long>();
        boolean removed;
        int lowerHalfSize, upperHalfSize;

        for( int i = 0; i < N; i++ ) {
            removed = true;
            // if operation is add
            if ( s[i].equals( "a" ) ) {
                if ( lowerHalfMaxHeap.size() == 0 ) {
                    lowerHalfMaxHeap.add( x[i] );
                    
                } else {
                    if ( x[i] > lowerHalfMaxHeap.peek() ) {
                        upperHalfMinHeap.add( x[i] );
                    } else {
                        lowerHalfMaxHeap.add( x[i] );
                    }
                }
            // if operation is remove
            } else {
                removed = false;
                if ( upperHalfMinHeap.size() != 0 && x[i] >= upperHalfMinHeap.peek() ) {
                    if ( upperHalfMinHeap.remove( x[i] ) ) {
                        removed = true;
                    }
                } else if ( lowerHalfMaxHeap.size() != 0 && x[i] <= lowerHalfMaxHeap.peek() ) {
                    if ( lowerHalfMaxHeap.remove( x[i] ) ) {
                        removed = true;
                    }
                } 

            }

            // adjustment
            if ( upperHalfMinHeap.size() > lowerHalfMaxHeap.size() ) {
                long min = upperHalfMinHeap.poll();
                lowerHalfMaxHeap.add( min );
            } if ( lowerHalfMaxHeap.size() > upperHalfMinHeap.size() + 1 ) {
                long max = lowerHalfMaxHeap.poll();
                upperHalfMinHeap.add( max );
            }
            
            
            // find median
            if ( !removed || ( lowerHalfMaxHeap.size() == 0 && upperHalfMinHeap.size() == 0 ) ) {
                System.out.println( "Wrong!" );
            } else if ( lowerHalfMaxHeap.size() == upperHalfMinHeap.size() ) {
                long sum = lowerHalfMaxHeap.peek() + upperHalfMinHeap.peek();
                if ( Math.abs(sum) % 2 == 0 ) {
                    System.out.println( sum / 2 );
                } else {
                    if ( sum >= 0 ) {
                        System.out.println( ( sum / 2 ) + ".5" );
                    } else {
                        System.out.println( "-" + ( Math.abs(sum) / 2 ) + ".5" );
                    }
                }
            } else if ( lowerHalfMaxHeap.size() != upperHalfMinHeap.size() ) {
                System.out.println( lowerHalfMaxHeap.peek() );
            } else {
                System.out.println( "Wrong!" );
            }
        }

    }
	
	
	public static class MaxComparator implements Comparator<Long> {
		@Override
		public int compare(Long a, Long b) {
			if (a.longValue() < b.longValue()) {
				return 1;
			} else if (a.longValue() > b.longValue()) {
				return -1;
			} else {
				return 0;
			}
		}

	}
}
