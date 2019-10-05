package longestSubarray;

import java.util.Map;
import java.util.Arrays;
import java.util.HashMap;

public class longestSubarray {
    // Longest Subarray By Sum in an Array - Interview Question -
    // Asked by Facebook & Palantir Technologies
    // In Java
    int[] findLongestSubarrayBySum(int sum, int[] arr) {
        Map<Integer, Integer> prefixSumsToIndex = new HashMap<>();
        int fromIndex = -1;
        int toIndex = -1;
        int sumExistAt = -1;
        
        // Calculates the prefixSums
        int previousSum = 0;
        for (int index = 0; index < arr.length; index++){
            int number = arr[index];
            int newSum = previousSum + number;
            previousSum = newSum;
            prefixSumsToIndex.put(newSum, index + 1);
            if (number == sum)
                sumExistAt = index + 1;
        }

        // Sees if such sum exists within the array
	    // initializes the indices with initial range if exists
        if (prefixSumsToIndex.get(sum) != null){
            fromIndex = 1;
            toIndex = prefixSumsToIndex.get(sum);
        }
    
        // Searches the array for a wider range if available
        for (int numSum : prefixSumsToIndex.keySet()){
            Integer indexofSum = prefixSumsToIndex.get(numSum);
            Integer indexOfSumWithS = prefixSumsToIndex.get(numSum+sum);
            if (!(indexOfSumWithS == null)){
                int tmpFromIndex = indexofSum + 1;
                int tmpToIndex = indexOfSumWithS;
                if ((toIndex - fromIndex) < (tmpToIndex - tmpFromIndex)){
                    toIndex = tmpToIndex;
                    fromIndex = tmpFromIndex;
                }
            }
        }

        // Returns depending on if indices has been found or not
        if (fromIndex != -1 && toIndex != -1)
            return new int[]{fromIndex, toIndex};
        else if (sumExistAt != -1)
            return new int[]{sumExistAt, sumExistAt};
        else return new int[]{-1};
    }

    public static void main(String[] args){
        longestSubarray obj = new longestSubarray();

        // Example
        int sum = 15;
        int[] arr = {6, 7, 8, 1, 2, 3, 0 , 4, 5, 0, 0, 9, 10};
        String result = Arrays.toString(
            obj.findLongestSubarrayBySum(sum, arr));
        System.out.println(result); // [4, 11]
    }
}