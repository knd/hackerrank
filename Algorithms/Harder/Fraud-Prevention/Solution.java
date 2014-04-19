/*
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
*/

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Hashtable;
import java.util.TreeSet;
import java.util.List;
import java.util.ArrayList;
import java.util.Iterator;


public class Solution {

    public static Hashtable<String, List<Record>> fraudsEmail = new Hashtable<String, List<Record>>();
    public static Hashtable<String, List<Record>> fraudsAddr = new Hashtable<String, List<Record>>();
    public static TreeSet<Integer> fraudsOrder = new TreeSet<Integer>();

    /** 
    *   Address class to bookkeep the address information.
    */
    public static class Address {
        protected String streetAddr;
        protected String city;
        protected String state;
        protected String zip;
    }

    /**
    *   Record class to bookkeep the deal record information.
    */
    public static class Record {
        protected String orderID;
        protected String dealID;
        protected String email;
        protected Address addr;
        protected String creditCard;
    }

    /**
    *   Function to create a new instance of a record based on the information
    *   provided.
    */
    private static Record createRecord(String[] inputTokens) {
        Address addr = new Address();
        Record aRec = new Record();
        addr.streetAddr = inputTokens[3];
        addr.city = inputTokens[4];
        addr.state = inputTokens[5];
        addr.zip = inputTokens[6];
        aRec.addr = addr;
        aRec.orderID = inputTokens[0];
        aRec.dealID = inputTokens[1];
        aRec.email = inputTokens[2];
        aRec.creditCard = inputTokens[7];
        return aRec;
    }

    /**
    *   Function to generate a unified key from the given email.
    */
    private static String createEmailKey(String email) {
        String [] emailTokens = email.split("@");
        String username = emailTokens[0].toLowerCase();
        String emailKey = username.contains("+") ? username.split("\\+")[0] : username;
        emailKey = emailKey.replace(".", "");
        return emailKey;
    }

    /**
    *   Function to detect frauds type: same email, same deal id, different
    *   credit card.
    */
    private static void detectFraudsEmail(Record rec) {
        String emailKey = createEmailKey(rec.email);
        String fraudsEmailKey = emailKey + rec.dealID;
        if ( fraudsEmail.containsKey(fraudsEmailKey) ) {
            List<Record> recList = fraudsEmail.get(fraudsEmailKey);
            for (Record r : recList) {
                if ( !rec.creditCard.equals(r.creditCard) ) {
                    fraudsOrder.add( Integer.parseInt(rec.orderID) );
                    fraudsOrder.add( Integer.parseInt(r.orderID) );
                }
            }
            recList.add( rec );
        } else {
            fraudsEmail.put( fraudsEmailKey, new ArrayList<Record> () );
            fraudsEmail.get( fraudsEmailKey ).add( rec );
        }
    }

    /**
    *   Function to generate a unified address based on given address.
    */
    private static String createAddrKey(Address addr) {
        String stAddr = addr.streetAddr.toLowerCase();
        String city = addr.city.toLowerCase();
        String state = addr.state.toLowerCase();
        String zip = addr.zip;
        
        int length = stAddr.length();
        String lastTokenStAddr = stAddr.substring( length-3 );
        if ( lastTokenStAddr.equals("st.") ) {
            stAddr = stAddr.substring( 0, length-3 ) + "street";
        } else if ( lastTokenStAddr.equals("rd.") ) {
            stAddr = stAddr.substring( 0, length-3 ) + "road";
        }

        if ( state.equals("new york") ) {
            state = "ny";
        } else if ( state.equals("california") ) {
            state = "ca";
        } else if ( state.equals("illinois") ) {
            state = "il";
        }
        return stAddr + city + state + zip;
    }

    /**
    *   Function to detect frauds type: same address, same deal id, different
    *   credit card.
    */
    private static void detectFraudsAddr(Record rec) {
        String addrKey = createAddrKey(rec.addr);
        String fraudsAddrKey = addrKey + rec.dealID;
        if ( fraudsAddr.containsKey(fraudsAddrKey) ) {
            List<Record> recList = fraudsAddr.get( fraudsAddrKey );
            for (Record r : recList) {
                if ( !rec.creditCard.equals(r.creditCard) ) {
                    fraudsOrder.add( Integer.parseInt(rec.orderID) );
                    fraudsOrder.add( Integer.parseInt(r.orderID) );
                }
            }
            recList.add( rec );
        } else {
            fraudsAddr.put( fraudsAddrKey, new ArrayList<Record> () );
            fraudsAddr.get( fraudsAddrKey ).add( rec );
        }
    }


      public static void main(String [] args) throws IOException {
        BufferedReader in = new BufferedReader( new InputStreamReader ( System.in ) );
        int testCase = Integer.parseInt( in.readLine() );
        for (int i = 0; i < testCase; i++) {
            String input = in.readLine();
            String [] inputTokens = input.split(",");
            Record rec = createRecord( inputTokens );
            detectFraudsEmail(rec);
            detectFraudsAddr(rec);
        }
        
        String lineToPrint = "";
        Iterator<Integer> iter = fraudsOrder.iterator();
        while ( iter.hasNext() ) {
            lineToPrint += iter.next() + ",";
        }
        if ( !lineToPrint.equals("") ) {
            lineToPrint = lineToPrint.substring( 0, lineToPrint.length()-1);
        }
        System.out.println( lineToPrint );

    }
}
