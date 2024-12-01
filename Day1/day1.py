
firstList = []
secondList = []

# fileName = "testData.txt"
fileName = "Data.txt"

with open(fileName, 'r') as file:
    for line in file:
        line = line.rstrip()
        nums = line.split("   ")
        firstList.append(int(nums[0]))
        secondList.append(int(nums[1]))


#? First Star

firstList.sort()
secondList.sort()

sum = 0

for i in range(len(firstList)):
    firstValue = firstList[i]
    secondValue = secondList[i]
    sum += abs(firstValue - secondValue)

print(sum)


#? Second Star

simSum = 0

for firstValue in firstList:
    counter = 0
    for secondValue in secondList:
        if firstValue == secondValue:
            counter += 1
    simSum += counter * firstValue

print(simSum)
