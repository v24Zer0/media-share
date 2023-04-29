package com.mediashare.authenticationservice.token;

public class TokenRequest {
    private String userID;

    public TokenRequest(String userID) {
        this.userID = userID;
    }
    public String getUserID() {
        return userID;
    }

    public void setUserID(String userID) {
        this.userID = userID;
    }
}
