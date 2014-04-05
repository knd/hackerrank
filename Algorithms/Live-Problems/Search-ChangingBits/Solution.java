import java.util.*;
import java.math.BigInteger;

class Solution {

    public static void main(String [] args) {
        Scanner in = new Scanner(System.in);

        int N, Q;
        N = in.nextInt();
        Q = in.nextInt();

        BigInteger A = new BigInteger( in.next(), 2 );
        BigInteger B = new BigInteger( in.next(), 2 );
        String command;
        int index = 0, bit = 0; 

        StringBuffer answer = new StringBuffer();

        for (int i=0; i<Q; i++) {
            command = in.next();

            if (command.equals("get_c")) {
                index = in.nextInt();
                answer.append(getCIdx(A, B, index));       

            } else {
                index = in.nextInt();
                bit = in.nextInt();
                if (command.equals("set_a")) {
                    A = changeBit( A, bit, index );
                } else {
                    B = changeBit( B, bit, index );
                }
            }
        }

        System.out.println(answer.toString());

    }

    public static String getCIdx(BigInteger A, BigInteger B, int index) {
        BigInteger C = A.add( B );
        if ( C.testBit( index ) ) {
            return new String("1");
        } 
        return new String("0");
    }

    public static BigInteger changeBit(BigInteger num, int bit, int index) {
        BigInteger bi = new BigInteger("0", 2);
        if ( bit == 0 ) {
            bi = num.clearBit( index );
        } else {
            bi = num.setBit( index );
        }
        return bi;
    }
    
    
}
