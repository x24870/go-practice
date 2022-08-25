import requests
from requests.structures import CaseInsensitiveDict

# url = "https://api.reqbin.com/api/v1/requests"
url = "https://grafana.jigentec.com"
# url = "http://34.117.76.130"
# url = "http://34.81.218.145/b.png"

headers = CaseInsensitiveDict()
headers["Access-Control-Request-Method"] = "GET"
headers["Access-Control-Request-Headers"] = "Content-Type"
headers["Origin"] = "*"


resp = requests.options(url, headers=headers)
# resp = requests.post(url, headers=headers)
# resp = requests.get(url, headers=headers)

print(resp.status_code, resp.headers)
# print(resp.content)
