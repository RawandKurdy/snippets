package main;

public class prodExcSelf {
    // Product Except Self [in O(n)] - Interview Question -
    // Asked by Amazon, LinkedIn, Facebook, Microsoft & Apple
    // In Java

    public int productExceptSelf(int[] nums, int m) {
        int s = 0; // Sum
        int p = 1; // Product

        for (int num : nums) {
            s = (p + s * num) % m;
            p = (p * num) % m;
        }
        return s;
    }
    public static void main(String args[]) {
        prodExcSelf obj = new prodExcSelf();
        // Example
        int[] nums = { 3, 3, 9, 5, 5, 4, 2, 8, 5, 9 }; // Input Array
        int m = 17; // Mod value

        System.out.println(obj.productExceptSelf(nums, m)); // 16
    }
}