def generate_business_name(business_type):
    # 示例数据 - 后续可扩展更多类型和生成逻辑
    name_lists = {
        'tech': ['Quantum Innovations', 'Neon Circuits', 'DataSphere'],
        'food': ['Golden Crust Bakery', 'Urban Spice Kitchen', 'Harvest Table'],
        'fashion': ['Silk & Stitch', 'Velvet Horizon', 'Chroma Threads']
    }
    return name_lists.get(business_type.lower(), ['Creative Solutions Group'])