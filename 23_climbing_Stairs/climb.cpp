#include <iostream>
#include <vector>

using namespace std;

// Climbing Stairs - Interview Question -
// Asked By Apple and Adobe
// in C++

int climbingStairs(int n) {
    vector<int> r;
    // Initial Result Set
    r.push_back(1);
    r.push_back(1);
    r.push_back(2);

    if (n <= 2)
        return r.at(n);

    for (int step = 3; step < n + 1; step++)
        r.push_back(r.at(step - 1) + r.at(step - 2));

    return r.at(n);
}

int main() {
    int n = 26; // Steps
    int ways = climbingStairs(n);
    cout << ways << endl; // 196418
}
