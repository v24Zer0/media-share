package com.mediashare.authenticationservice.mock;

import com.mediashare.authenticationservice.token.TokenService;

public class MockTokenService implements TokenService {
    public String createToken(String userID) {
        return "";
    }

    public boolean verifyToken(String token) {
        return true;
    }
}
