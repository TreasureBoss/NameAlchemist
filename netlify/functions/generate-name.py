# netlify/functions/generate-name.py
import json

def handler(event, context):
    # 添加对空event的处理
    if not event.get('body'):
        return {
            'statusCode': 400,
            'body': json.dumps({'error': 'Missing request body'})
        }
    data = json.loads(event['body'])
    arabic_name = data.get('arabicName', '')
    
    # 添加空值校验
    if not arabic_name:
        return {
            'statusCode': 400,
            'body': json.dumps({'error': 'Missing arabicName parameter'})
        }
    
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

# 添加Netlify函数要求的main判断
__all__ = ["handler"]

if __name__ == '__main__':
    # 本地测试用示例
    class Context:
        function_version = '$LATEST'
    
    test_event = {
        'body': json.dumps({'arabicName': 'Ahmed'})
    }
    print(handler(test_event, Context()))