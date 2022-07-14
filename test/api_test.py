import json
import requests


def signup(username = "test", password = "test@123"):
    try:
        hdr = {
            "Content-Type": "application/json"
        }
        req_body = {
            "username": username,
            "password": password
        }
        resp = requests.post("http://localhost:8081/app/v1/user/signup", json=req_body, headers=hdr)
        if resp.status_code == 200:
            print("Sign up successfully!\n{}".format(json.dumps(resp.json(), indent=2)))

        else:
            print("Failed to signgup (username is already used?)!\n{}".format(json.dumps(resp.json(), indent=2)))
    except Exception as e:
        print(e)

def login(username = "test", password = "test@123"):
    try:
        hdr = {
            "Content-Type": "application/json"
        }
        req_body = {
            "username": username,
            "password": password
        }
        resp = requests.post("http://localhost:8081/app/v1/user/login", json=req_body, headers=hdr)
        if resp.status_code == 200:
            #print("Login: {}".format(json.dumps(resp.json(), indent=2)))
            return resp.json()["data"]["access_token"]
        else:
            print("Failed to login!\n{}".format(json.dumps(resp.json(), indent=2)))
    except Exception as e:
        print(e)

def createShape(token, shape, edges):
    shape_id = None
    try:
        hdr = {
            "Authorization": "Bearer {}".format(token),
            "Content-Type": "application/json"
        }
        req_body = {
            "shape": shape,
            "edges": edges
        }
        resp = requests.post("http://localhost:8081/app/v1/shape/create", json=req_body, headers=hdr)
        if resp.status_code == 200:
            shape_id = resp.json()["data"]["shape_id"]
            print("Create shape OK\n{}".format(json.dumps(resp.json(), indent=2)))
        else:
            print("Failed to create shape!\n{}".format(json.dumps(resp.json(), indent=2)))
    except Exception as e:
        print(e)
    return shape_id

def getAllShapes(token):
    try:
        hdr = {
            "Authorization": "Bearer {}".format(token)
        }
        resp = requests.get("http://localhost:8081/app/v1/shape", headers=hdr)
        if resp.status_code == 200:
            print("All Shapes:\n{}".format(json.dumps(resp.json(), indent=2)))
        else:
            print("Failed to get all shapes!\n{}".format(json.dumps(resp.json(), indent=2)))
    except Exception as e:
        print(e)

def getShapeById(token, shapeId=1):
    try:
        hdr = {
            "Authorization": "Bearer {}".format(token)
        }
        resp = requests.get("http://localhost:8081/app/v1/shape/{}".format(shapeId), headers=hdr)
        if resp.status_code == 200:
            print("Shape {} info:\n{}".format(shapeId, json.dumps(resp.json(), indent=2)))
        else:
            print("Failed to get shape {}!\n{}".format(shapeId, json.dumps(resp.json(), indent=2)))
    except Exception as e:
        print(e)

def calculateShapeById(token, shapeId=1):
    try:
        hdr = {
            "Authorization": "Bearer {}".format(token),
            "Content-Type": "application/json"
        }
        x = "area(shape_id:{}) perimeter(shape_id:{})".format(shapeId, shapeId)
        query = {
            "query" : '{} {} {}'.format('{', x, '}')
        }
        resp = requests.post("http://localhost:8081/app/v1/shape/calculate", json=query, headers=hdr)
        if resp.status_code == 200:
            print("Shape Value:\n{}".format(json.dumps(resp.json(), indent=2)))
        else:
            print("Failed to calculate shape {}!\n{}".format(shapeId, json.dumps(resp.json(), indent=2)))
    except Exception as e:
        print(e)

def updateShapeById(token, shapeId, edges):
    try:
        hdr = {
            "Authorization": "Bearer {}".format(token),
            "Content-Type": "application/json"
        }
        body = {
            "edges" : edges
        }
        resp = requests.put("http://localhost:8081/app/v1/shape/{}".format(shapeId), json=body, headers=hdr)
        if resp.status_code == 200:
            print("Shape Value:\n{}".format(json.dumps(resp.json(), indent=2)))
        else:
            print("Failed to update shape {}!\n{}".format(shapeId, json.dumps(resp.json(), indent=2)))
    except Exception as e:
        print(e)

def deleteShapeById(token, shapeId):
    try:
        hdr = {
            "Authorization": "Bearer {}".format(token),
            "Content-Type": "application/json"
        }
        resp = requests.delete("http://localhost:8081/app/v1/shape/{}".format(shapeId), headers=hdr)
        if resp.status_code == 200:
            print("Deleted shape {} successfully!".format(shapeId))
        else:
            print("Failed to delete shape {}!\n{}".format(shapeId, json.dumps(resp.json(), indent=2)))
    except Exception as e:
        print(e)

if __name__ == "__main__":
    signup("test", "test@123")
    access_token = login("test", "test@123")
    if access_token != None:
        print("access_token: {}".format(access_token))

    # Create shape
    shape_id = createShape(access_token, "square", ["5"])

    # Get all shape
    getAllShapes(access_token)

    # Get shape by ID
    getShapeById(access_token, shape_id)

    # Get shape value
    calculateShapeById(access_token, shape_id)

    # Update
    updateShapeById(access_token, shape_id, ["7"])

    # Get shape value again
    calculateShapeById(access_token, shape_id)

    # Delete shape
    deleteShapeById(access_token, shape_id)

    # Get shape again
    getShapeById(access_token, shape_id)


