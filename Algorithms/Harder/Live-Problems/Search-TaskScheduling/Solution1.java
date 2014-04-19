import java.util.*;
import java.lang.Math;

class Solution {

    public static void main(String [] args) {
        
        Scanner in = new Scanner(System.in);
        int T = in.nextInt();
        long D, M;
        Queue<Task> taskQueueA = new PriorityQueue<Task>(1, new TaskComparator());
        Queue<Task> taskQueueB;
        List<Long> maxLateness = new LinkedList<Long>();

        for (int i=0; i<T; i++) {
            D = in.nextLong();
            M = in.nextLong();
            taskQueueA.add(new Task(D, M));
            taskQueueB = new PriorityQueue<Task>(1, new TaskComparator());
            maxLateness.add(calculate(taskQueueA, taskQueueB));
            taskQueueA = taskQueueB;
        }

        for (Long lateness : maxLateness) {
            System.out.println(lateness);
        }
    }

    public static long calculate(Queue<Task> A, Queue<Task> B) {
        long maxLateness = Integer.MIN_VALUE;
        Task t = A.poll();
        long timeCurrent = 0, timeStart = 0, timeFinish = 0, lateTime = 0;
        while (t != null) {
            B.add(t);
            timeFinish = timeCurrent + t.runTime;
            lateTime = Math.max(0, timeFinish - t.deadLine);
            if (lateTime > maxLateness) {
                maxLateness = lateTime;
            }
            timeCurrent += t.runTime;
            t = A.poll();
        }

        return maxLateness;
    }

    public static class Task {
        public long deadLine;
        public long runTime;

        public Task(long d, long r) {
            this.deadLine = d;
            this.runTime = r;
        }
    }

    public static class TaskComparator implements Comparator<Task> {
        @Override
        public int compare(Task a, Task b) {
            if (a.deadLine < b.deadLine) {
                return -1;
            } else if (a.deadLine > b.deadLine) {
                return 1;
            } else {
                return 0;
            }
        }
    }

}
