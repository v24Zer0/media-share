package com.mediashare.authenticationservice.token;

public interface TokenService {
    String createToken(String userID);
    boolean verifyToken(String token);
}
