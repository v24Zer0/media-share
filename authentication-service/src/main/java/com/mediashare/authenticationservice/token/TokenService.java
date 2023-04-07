package com.mediashare.authenticationservice.token;

import org.springframework.stereotype.Service;

@Service
public class TokenService {

    public TokenResponse getToken() {
        return new TokenResponse("token", "user");
    }
}
