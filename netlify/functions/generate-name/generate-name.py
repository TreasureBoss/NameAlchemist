import json
# 假设name_generator模块在当前目录下，确保name_generator.py文件存在
from name_generator import generate_business_name

def handler(event, context):
    try:
        business_type = json.loads(event['body'])['type']
        names = generate_business_name(business_type)
        return {
            'statusCode': 200,
            'body': json.dumps({'names': names})
        }
    except Exception as e:
        return {
            'statusCode': 500,
            'body': json.dumps({'error': str(e)})
        }