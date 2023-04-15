package com.mediashare.authenticationservice.token;

public class TokenVerifyRequest {
    private String token;

    public TokenVerifyRequest(String token) {
        this.token = token;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }
}
