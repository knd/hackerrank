import java.util.*;
import java.lang.Math;
import java.math.BigInteger;


class Solution {
	public static void main( String args[] ){
		
// helpers for input/output		

		Scanner in = new Scanner(System.in);
		
		int N;
		N = in.nextInt();
	
		long [] xs = new long [N];
        long [] ys = new long [N];
        double xsAvg = 0.0;
        double ysAvg = 0.0;
        double minDist;
        int index = 0;
        BigInteger sum = new BigInteger("0");

		for(int i=0; i<N; i++){
			xs[i] = in.nextLong();
            ys[i] = in.nextLong();
            xsAvg += (double) xs[i] / (double) N;
            ysAvg += (double) ys[i] / (double) N;
		}
        
        minDist = Math.sqrt( Math.pow( xsAvg-xs[0], 2 ) + Math.pow( ysAvg-ys[0], 2 ) );
        

        for (int i=1; i<N; i++) {
            double dist = Math.sqrt( Math.pow( xsAvg-xs[i], 2 ) + Math.pow( ysAvg-ys[i], 2 ) );
            if ( minDist > dist ) {
                minDist = dist;
                index = i;
            }
        }
        
        for (int i=0; i<N; i++) {
            sum = sum.add(  BigInteger.valueOf( Math.max( Math.abs( xs[i]-xs[index] ), Math.abs( ys[i]-ys[index] ) ) ) );
        }

        System.out.println(sum.toString());
    }
}
