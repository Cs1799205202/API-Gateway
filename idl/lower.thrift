namespace go lower

struct NormalRequest {
    1: string message
}

struct LowerResponse {
    1: string result
}

service LowerService {
    LowerResponse tolower(1: NormalRequest req) (api.post="/internal/tolower")
}