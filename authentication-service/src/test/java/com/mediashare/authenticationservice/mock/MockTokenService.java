package com.mediashare.authenticationservice.mock;

import com.mediashare.authenticationservice.token.TokenService;

public class MockTokenService implements TokenService {
    public String createToken(String userID) {
        if(userID.equals("invalid_user"))
            return "";
        return "token";
    }

    public boolean verifyToken(String token) {
        if(token.equals("valid_token"))
            return true;
        return false;
    }
}
