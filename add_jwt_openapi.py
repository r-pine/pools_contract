import yaml

def add_jwt_auth(yaml_file):
    # Открываем YAML-файл
    with open(yaml_file, 'r') as file:
        data = yaml.safe_load(file)

    # Добавляем JWT Bearer Security Scheme
    if 'components' not in data:
        data['components'] = {'securitySchemes': {}}
    if 'securitySchemes' not in data['components']:
        data['components']['securitySchemes'] = {}
    data['components']['securitySchemes']['tokenAuth'] = {
        'type': 'http',
        'scheme': 'bearer',
        'bearerFormat': 'JWT'
    }

    # Добавляем глобальные требования безопасности, если они не существуют
    if 'security' not in data:
        data['security'] = [{'tokenAuth': []}]
    else:
        if not any('tokenAuth' in item for item in data['security']):
            data['security'].append({'tokenAuth': []})

    # Добавляем требования безопасности ко всем путям
    if 'paths' in data:
        for path, methods in data['paths'].items():
            for method, details in methods.items():
                if 'security' not in details:
                    details['security'] = [{'tokenAuth': []}]
                else:
                    if not any('tokenAuth' in item for item in details['security']):
                        details['security'].append({'tokenAuth': []})

    # Сохраняем изменения обратно в файл
    with open(yaml_file, 'w') as file:
        yaml.safe_dump(data, file, sort_keys=False)

if __name__ == "__main__":
    import sys
    if len(sys.argv) != 2:
        print("Usage: python add_auth_jwt.py <path_to_yaml>")
        sys.exit(1)
    yaml_file = sys.argv[1]
    add_jwt_auth(yaml_file)