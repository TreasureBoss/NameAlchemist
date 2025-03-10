# netlify/functions/generate-name.py
import json

def handler(event, context):
    data = json.loads(event['body'])
    arabic_name = data['arabicName']
    
    # 核心逻辑：将阿拉伯名转为中文名
    chinese_name = arabic_to_chinese(arabic_name)
    
    return {
        'statusCode': 200,
        'body': json.dumps({'chineseName': chinese_name})
    }

def arabic_to_chinese(arabic_name):
    # 实现音译逻辑（示例）
    mapping = {
        'Ahmed': '艾哈迈德',
        'Fatima': '法蒂玛'
    }
    return mapping.get(arabic_name.capitalize(), '未找到匹配名')