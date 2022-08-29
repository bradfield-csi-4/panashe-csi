# coding: utf-8
import csv
import statistics
        
amounts = []
with open('metrics/payments.csv') as csvfile:
    reader = csv.reader(csvfile)
    for l in reader:
        amounts.append(int(l[0]))
        
statistics.stdev(amounts)

def variance(data):
    T, ss = _ss(data)
    return convert(ss / len(data), T)

def stddev(data):
    return math.sqrt(variance(data))
