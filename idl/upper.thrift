namespace go upper

struct NormalRequest {
    1: string message
}

struct UpperResponse {
    1: string result
}

service UpperService {
    UpperResponse toupper(1: NormalRequest req) (api.post="/internal/toupper")
}