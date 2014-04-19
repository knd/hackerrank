import java.util.*;
import java.lang.Math;

class Solution {

    public static void main(String [] args) {
        
        Scanner in = new Scanner(System.in);
        int T = in.nextInt();
        long D, M;
        int index = 0;
        List<Long> maxLateness = new LinkedList<Long>();
        SortedArrayList<Task> tasks = new SortedArrayList<Task>();

        for (int i=0; i<T; i++) {
            D = in.nextLong();
            M = in.nextLong();
            index = tasks.insertSorted(new Task(D, M));
            Task.maxLateness = Math.max( Task.maxLateness, calculate(index, tasks) );
            maxLateness.add(Task.maxLateness);
        }

        for (Long lateness : maxLateness) {
            System.out.println(lateness);
        }
    }

    public static long calculate(int index, SortedArrayList<Task> tasks) {
        long max = Integer.MIN_VALUE;
        long currentTime, finishTime, lateness;
        Task currentTask;
        if (index == 0) {
            currentTime = 0;
        } else {
            currentTime = tasks.get(index-1).startTime + tasks.get(index-1).runTime;
        }
        for (int i=index; i<tasks.size(); i++) {
            currentTask = tasks.get(i);
            currentTask.startTime = currentTime;
            finishTime = currentTime + currentTask.runTime;
            lateness = Math.max( 0, finishTime - currentTask.deadLine );
            if (max < lateness) {
                max = lateness;
            }
            currentTime = finishTime;
        }
        
        return max;
    }

    public static class Task implements Comparable<Task> {
        public long deadLine;
        public long runTime;
        public static long maxLateness = Integer.MIN_VALUE;
        public long startTime;

        public Task(long deadLine, long runTime) {
            this.deadLine = deadLine;
            this.runTime = runTime;
        }

        public void setMaxLateness(long l) {
            this.maxLateness = l;
        }

        public int compareTo(Task t) {
            if (this.deadLine < t.deadLine) {
                return -1;
            } else if (this.deadLine > t.deadLine) {
                return 1;
            } else {
                return 0;
            }
        }
    }

    public static class SortedArrayList<T> extends ArrayList<T> {
        @SuppressWarnings("unchecked")
        public int insertSorted(T value) {
            add(value);
            int index = size() - 1;
            Comparable<T> cmp = (Comparable<T>) value;
            for (int i=size()-1; i > 0 && cmp.compareTo(get(i-1)) < 0; i--) {
                Collections.swap(this, i, i-1);
                index = i - 1;
            }
            return index;
        }
    }

}
