public class divNumber {
    // DivNumber Challenge
    // Given three integers k, l and r
    // return the number of integers between l and r inclusive
    // that have exactly k divisors.
    // In Java
    public static void main(String[] args) {
        divNumber obj = new divNumber();
        // Example
        int k = 3;
        int l = 2;
        int r = 49;
        System.out.println(obj.divNumber(k, l, r)); // 4

        // We have 4 integers in that range
        // that have exactly 3 divisors - 4, 9, 25, 49.
    }

    int divNumber(int k, int l, int r) {
        // di is the number of Numbers
        // that has 'k' exact divisors
        int di = 0;
        while (l <= r) {
            // d is number of divisors of individual numbers
            int d = 0;
            // c is counter(Number)
            int c = 1;

            while (c <= l) {
                if (l % c == 0)
                    d++;
                c++;
            }
            if (d == k)
                di++;
            l++;
        }

        return di;
    }

}