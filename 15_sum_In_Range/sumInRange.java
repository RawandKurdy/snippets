package sumInRange;

import java.util.Map;
import java.util.HashMap;

public class sumInRange {
    // Sum In Range - Interview Question - Asked by Palantir Technologies
    // In Java
    // In O(n) time
    int sumInRange(int[] nums, int[][] queries) {
        int totalSum = 0;
        Map<Integer, Integer> prefixSums = new HashMap<>();
        int modulo = 1000000000 + 7;

        // calculates the prefix-sums
        for (int index = 0; index < nums.length; index++) {
            if (index == 0) {
                prefixSums.put(index, nums[index] % modulo);
            } else {

                prefixSums.put(index,
                (prefixSums.get(index - 1) + nums[index]) % modulo);
            }
        }

        // Now adds up each separate range
        for (int[] query : queries) {
            int fromIndex = query[0];
            int toIndex = query[1];
            totalSum += prefixSums.get(toIndex);

            if (fromIndex != 0) {
                totalSum -= prefixSums.get(fromIndex - 1);
            }
            totalSum %= modulo;
        }
        // if totalSum is negative we just add the modulo
        if (totalSum < 0) {
            totalSum += modulo;
        }
        return totalSum;
    }

    public static void main(String[] args){
        // Example
        int[] nums =  {34, 19, 21, 5, 1, 10, 26, 46, 33, 10};
        int[][] queries = {{3,7}, 
        {3,4}, 
        {3,7}, 
        {4,5}, 
        {0,5}};

        sumInRange obj = new sumInRange();
        System.out.println(obj.sumInRange(nums, queries)); // 283

    }
}