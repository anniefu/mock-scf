def hello(request):
    print("python", request.get_json())
    return request.get_json()