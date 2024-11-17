package helper 

func WrapResponse(message string, code int, status string, data interface{}) map[string]interface{} {
    return map[string]interface{}{
        "meta": map[string]interface{}{
            "message": message,
            "code":    code,
            "status":  status,
        },
        "data": data,
    }
}