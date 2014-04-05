/* Sample program illustrating input and output */

import java.util.*;
import java.text.DecimalFormat;

class Solution1{
	public static void main( String args[] ){
		
// helpers for input/output		

		Scanner in = new Scanner(System.in);
		
		int N;
		N = in.nextInt();
	
		String s[] = new String[N];
		int x[] = new int[N];
		
		for(int i=0; i<N; i++){
			s[i] = in.next();
			x[i] = in.nextInt();
		}

        OrderedLinkedList ol = new OrderedLinkedList();
        DecimalFormat df = new DecimalFormat();
        boolean removed;
        for(int i=0; i<N; i++) {
            if ( s[i].equals("a") ) {
                ol.add( x[i] );
                System.out.println( df.format( median( ol ) ) );
            } else {
                removed = ol.remove( new Integer( x[i] ) );
                if ( removed && ol.size() > 0 ) {
                    System.out.println( df.format( median( ol ) ) );
                } else {
                    System.out.println( "Wrong!" );
                }
            }
        }

	}

    public static float median( OrderedLinkedList ol ) {
        int length = ol.size();
        if ( length % 2 == 0 ) {
            return ( ol.get( length / 2 ) + ol.get( length / 2 - 1 ) ) / (float) 2;
        } 
        return (float) ol.get( (length-1) / 2 );
    }

    public static class OrderedLinkedList extends LinkedList<Integer> {

        public boolean add( Integer item ) {
            for (int i=0; i < size(); i++) {
                Integer itemOfList = get( i );
                if ( item.compareTo( itemOfList ) < 0 ) {
                    add( i, item );
                    return true;
                }
            }
            addLast( item );
            return true;
        }
    }

}
