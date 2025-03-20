import requests
from bs4 import BeautifulSoup
import random

def generate_business_name(business_type):
    # 实现名称生成逻辑（示例代码）
    prefixes = ['Creative', 'Smart', 'Global', 'NextGen']
    suffixes = ['Solutions', 'Hub', 'Labs', 'Innovations']
    
    names = [
        f'{random.choice(prefixes)}{business_type.capitalize()}{random.choice(suffixes)}',
        f'{business_type.capitalize()}{random.choice(["Pro", "Max", "Plus"])}'
    ]
    return names[:3]