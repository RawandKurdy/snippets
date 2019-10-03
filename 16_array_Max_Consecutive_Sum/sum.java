package sum;

public class sum {
    // Array Max Consecutive Sum - Interview Question -
    // Asked by Microsoft, LinkedIn, Amazon and Samsung
    // Maximum possible sum you can get from one of contiguous subarrays.
    // In Java
    int arrayMaxConsecutiveSum2(int[] inputArray) {
        int sumMax = 0;
        int sumSteps = 0;
        int maxNumber = 0;

        for (int index = 0; index < inputArray.length; index++) {
            int number = inputArray[index];
            sumSteps += number;

            if (maxNumber == 0 || maxNumber < number)
                maxNumber = number;
            if (sumSteps < 0)
                sumSteps = 0;
            if (sumMax < sumSteps)
                sumMax = sumSteps;
        }
        return sumMax != 0 ? sumMax : maxNumber;
    }

public static void main(String[] args){
    sum obj = new sum();
 // Example
 int[] arr = {1, -2, 3, -4, 5, -3, 2, 2, 2};
 System.out.println(obj.arrayMaxConsecutiveSum2(arr)); // 8
}
}