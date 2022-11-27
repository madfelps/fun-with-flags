import json 

f = open('country.json')

fw = open("sql-formatted.txt", "a")

data = json.load(f)

for i in data: 
    text = '(\'' + i["name"] + '\', \'' + i["country-code"] + '\'),\n'
    fw.write(text)
    print(text)

f.close()
fw.close()