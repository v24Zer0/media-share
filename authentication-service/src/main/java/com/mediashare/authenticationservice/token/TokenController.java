package com.mediashare.authenticationservice.token;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TokenController {
    private final TokenService tokenService;

    @Autowired
    public TokenController(TokenService tokenService) {
        this.tokenService = tokenService;
    }

    @GetMapping("/")
    public TokenResponse getToken() {
        String token = this.tokenService.getToken();
        return new TokenResponse(token, "user");
    }
}
