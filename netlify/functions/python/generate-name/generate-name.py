import json
import requests

def handler(event, context):
    try:
        # 函数逻辑代码
        return {
            'statusCode': 200,
            'body': json.dumps({'name': 'Generated Name'})
        }
    except Exception as e:
        return {
            'statusCode': 500,
            'body': json.dumps({'error': str(e)})
        }