import java.util.*;

class Solution {

    public static void main( String args[] ) {

        Scanner in = new Scanner( System.in );

        int N, K;
        N = in.nextInt();
        K = in.nextInt();

        Long C[] = new Long[N];
        for ( int i=0; i<N; i++ ) {
            C[i] = in.nextLong();
        }

        long result = calculate( C, N, K );
        System.out.println( result );
    }

    private static long calculate( Long [] costs, int flowersNum, int friendsNum ) {
        long minExpense = 0;
        int currFriend = 0;
        Arrays.sort( costs, Collections.reverseOrder() );
        for( int i=0; i<flowersNum; i++ ) {
            long x = (long) ( currFriend / friendsNum ); 
            minExpense += ( ( x + 1 ) * costs[ i ].longValue() );
            currFriend++;
        }
        return minExpense;
    }

}
