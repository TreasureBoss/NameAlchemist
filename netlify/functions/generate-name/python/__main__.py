import random
import json
import os

def handler(event, context):
    try:
        with open('chinese_words.json') as f:
            word_bank = json.load(f)
        
        names = [' '.join(random.sample(word_bank['adjectives'] + word_bank['nouns'], 2)) for _ in range(20)]
        return {
            'statusCode': 200,
            'body': '\n'.join(names)
        }
    except Exception as e:
        return {
            'statusCode': 500,
            'body': f'生成失败: {str(e)}'
        }

if __name__ == "__main__":
    print(handler(None, None))