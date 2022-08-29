import math
import statistics
import random
# For a new value newValue, compute the new count, new mean, the new M2.
# mean accumulates the mean of the entire dataset
# M2 aggregates the squared distance from the mean
# count aggregates the number of samples seen so far

# Retrieve the mean, variance and sample variance from an aggregate
def finalize(existingAggregate):
    (count, mean, M2) = existingAggregate
    if count < 2:
        return float("nan")
    else:
        return M2 / (count - 1)

def stddev(l):
    count, mean, M2 = (0, 0, 0)
    for el in l:
        count += 1
        delta = el - mean
        mean += delta / count
        delta2 = el - mean
        M2 += delta * delta2

    return math.sqrt(M2 / (count - 1))

if __name__ == "__main__":
    l = [random.randint(0, 10) for _ in range(10)]
    print(f"{statistics.stdev(l)=}")
    print(f"{stddev(l)=}")
