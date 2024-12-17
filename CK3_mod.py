folder_path = 'viet_korea/'
file_path = 'viet_misc_l_korean'

with open(folder_path+file_path + ".yml", 'r', encoding='utf-8') as file:
    lines = file.readlines()

event_dict = {}
tooltip_dict = {}

for i in range(len(lines)):
    line = lines[i].strip()
    if line == '' or line[0] == '#' or i == 0:
        continue
    first_colon = line.find(':')
    num_dot = line[:first_colon].find('.')
    type_dot = line[num_dot + 1:].find('.') + num_dot + 1  # 절대 위치 계산

    if num_dot != -1 and line[num_dot + 1:type_dot].isdigit(): # 이벤트인 경우
        # exceptional_dot = line[type_dot + 1:first_colon].find('.')
        # if exceptional_dot != -1:  # exceptional_dot이 존재하는 경우에만 절대 위치 계산
        #     exceptional_dot += type_dot + 1
        #     event_num = line[num_dot + 1:type_dot]
        #     event_type = line[type_dot + 1:exceptional_dot]
        #     event_exceptional = line[exceptional_dot + 1:first_colon]
        # else:
        #     event_num = line[num_dot + 1:type_dot]
        #     event_type = line[type_dot + 1:first_colon]
        #     event_exceptional = ''

        event_num = line[num_dot + 1:type_dot]
        event_type = line[type_dot + 1:first_colon]
        
        content = line[first_colon+4:-1] # 본문 긁어오기

        if content != 'xxxxx': # 이벤트가 존재할때만 딕셔너리에 추가
            if event_type == "t":
                event_dict[event_num] = {"t": content}
            else:
                event_dict[event_num][event_type] = content
    else: # 툴팁인 경우
        tooltip = line[:first_colon]
        content = line[first_colon+4:-1]

        tooltip_dict[tooltip] = content


import pymongo

client = pymongo.MongoClient("mongodb://localhost:27017/")

db = client["CK3_mod"]

collection = db[file_path]

merged_data = {}
for key, value in event_dict.items():
    merged_data.update({key: value})

collection.insert_one(merged_data)