namespace go user

struct RegisterRequest{
    1: string username;
    2: string password;
}
struct RegisterResponse{
    1: i32 code;
    2: string message;
}
struct LoginRequest{
    1: string username;
    2: string password;
}
struct LoginResponse{
    1: i32 code;
    2: string message;
    3: string token;
}

struct User{
    1: string username;
    2: string password;
}


service UserService{
    RegisterResponse Register(1:RegisterRequest req);
    LoginResponse Login(1:LoginRequest req);

}