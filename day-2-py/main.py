from functools import reduce
import operator


file1 = open('input.txt', 'r')
Lines = file1.readlines()
 
def colorToIndex(s):
  if s == "red":
    return 0
  if s == "green":
    return 1
  if s == "blue":
    return 2

def readGame(s):
  limits = [12, 13, 14]
  gameSets = s.split(';')
  minLimits = [0,0,0]
  for gameSet in gameSets:
    chunks = gameSet.split(',')
    for chunk in chunks:
      parts = chunk.split(' ')
      n = int(parts[1])
      i = colorToIndex(parts[2])
      if minLimits[i] < n:
        minLimits[i] = n
  return reduce(operator.mul, minLimits, 1)

count = 0
res = 0
for line in Lines:
    count += 1
    x = line.strip()
    parts = x.split(':')
    res += readGame(parts[1])
    
print(res)