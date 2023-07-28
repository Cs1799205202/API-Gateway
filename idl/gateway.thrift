namespace go gateway

struct APIRequest {
    1: string ServiceName (api.path="ServiceName")
    2: string MethodName (api.path="MethodName")
    3: string message (api.body="message")
}

struct APIResponce {
    1: string message
}

service GatewayService {
    APIResponce Route(1: APIRequest APIRequest) (api.post="/agw/:ServiceName/:MethodName")
}