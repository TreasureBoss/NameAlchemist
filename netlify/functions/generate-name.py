import os

def handler(event, context):
    data = event.get('body', {})
    business_type = data.get('business_type', 'tech')
    names = generate_business_name(business_type)
    return {
        'statusCode': 200,
        'body': {"names": names}
    }

def generate_business_name(business_type):
    name_lists = {
        'tech': ['Quantum Innovations', 'Neon Circuits', 'DataSphere'],
        'food': ['Golden Crust Bakery', 'Urban Spice Kitchen', 'Harvest Table'],
        'fashion': ['Silk & Stitch', 'Velvet Horizon', 'Chroma Threads']
    }
    return name_lists.get(business_type.lower(), ['Creative Solutions Group'])