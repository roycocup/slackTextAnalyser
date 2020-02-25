
late = []
with open('dev_announcements.txt') as f:
    content = f.readlines()
    for line in content:
        if 'late' in line.lower(): 
            late.append(line)

latePeople = dict()
for line in late:
    personName = str(line.split(' - ')[1:2])
    if personName in latePeople.keys():
        latePeople[personName] += 1
    else:
        latePeople[personName] = 1

sorted_x = sorted(latePeople.items())
print(sorted_x)
    