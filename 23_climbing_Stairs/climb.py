# Climbing Stairs - Interview Question -
# Asked By Apple and Adobe
# in Python


def climbingStairs(n):
    r = [1, 1, 2]  # Initial Result Set
    if n <= 2: return r[n]

    for step in range(3, n + 1):
        r.append(r[step - 1] + r[step - 2])
    return r[n]


n = 26  # Steps
ways = climbingStairs(n)
print(ways)  # 196418
