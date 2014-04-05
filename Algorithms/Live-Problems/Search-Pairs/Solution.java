import java.util.*;


class Solution{
	public static void main( String args[] ){
		
		// helpers for input/output		

		Scanner in = new Scanner(System.in);
		
		int N, count = 0;
        long K;
		N = in.nextInt();
        K = in.nextLong();
	
		TreeSet<Long> x = new TreeSet<Long>();
		for(int i=0; i<N; i++){
			x.add( in.nextLong() );
		}

        for( Iterator<Long> iter = x.iterator(); iter.hasNext(); ) {
            long num = iter.next();
            if ( x.contains( num + K ) ) {
                count++;
            }
        }
        System.out.println( count );

	}
}
