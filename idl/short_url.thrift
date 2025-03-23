namespace go shorturl

struct GenerateReq {
    1: string LongUrl
}

struct GenerateResp {
    1: string ShortUrl
}

struct RedirectReq {
    1: string ShortUrl
}

struct RedirectResp {
    1: string LongUrl
}

service ShortUrlService {
    GenerateResp Generate(1: GenerateReq req)
    RedirectResp Redirect(1: RedirectReq req)
}