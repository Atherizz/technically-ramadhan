import re
def high(x):
    x = x.lower()
    separated = re.split(" ",x)
    letterCount = []
    
    for word in separated:
        sum = 0
        for char in word:
            num = ord(char) - ord('a') + 1
            sum += num
            
        letterCount.append(sum)
    highest = 0
    index = 0
    for i, num in enumerate(letterCount):
        if num > highest:
            highest = num
            index = i
            
    return separated[index]
    

print(high("what time are we climbing up the volcano"))
    
    