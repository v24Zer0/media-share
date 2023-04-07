package com.mediashare.authenticationservice.token;

public class TokenResponse {
    String token;
    String userID;

    public TokenResponse(String token, String userID) {
        this.token = token;
        this.userID = userID;
    }
    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

    public String getUserID() {
        return userID;
    }

    public void setUserID(String userID) {
        this.userID = userID;
    }
}
